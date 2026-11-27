package rueidisgcp

import (
	"context"
	"fmt"
	"math/rand/v2"
	"sync"
	"sync/atomic"
	"time"

	"github.com/redis/rueidis"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Defaults used when the matching IAMAuthOptions field is left empty.
const (
	// DefaultCloudPlatformScope is the OAuth2 scope requested for IAM access tokens.
	DefaultCloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"
	// DefaultRefreshBuffer is how long before a token expires a new one is used: 3m45s, the
	// margin Google's own auth library uses (defaultExpiryDelta in cloud.google.com/go/auth,
	// https://github.com/googleapis/google-cloud-go/blob/auth/v0.24.0/auth/auth.go#L47-L49).
	// The metadata server can keep handing out the old token until 4 minutes before it
	// expires, so asking earlier could get the same token again.
	DefaultRefreshBuffer = 3*time.Minute + 45*time.Second
	// DefaultIAMUsername is the username sent with AUTH. Memorystore accepts "default" or
	// the email of the IAM principal.
	DefaultIAMUsername = "default"
	// DefaultTokenFetchTimeout is how long one request for a new IAM token may take.
	DefaultTokenFetchTimeout = 10 * time.Second
)

// credentialRetry is how soon a failed token refresh is tried again while the last token
// that worked is still in use. jitter turns it into 27-30s, so connections don't all retry
// at the same moment.
const credentialRetry = 30 * time.Second

// IAMAuthOptions configures IAM authentication: where the access tokens come from, the
// username sent with them, and how they are refreshed. The zero value uses Application
// Default Credentials (ADC) and the username "default".
//
// The access tokens come from CustomTokenSource, or from ADC when it is nil; that identity
// needs roles/memorystore.dbConnectionUser on the instance.
type IAMAuthOptions struct {
	// Username is sent with AUTH. Memorystore accepts only "default" or the email of the
	// IAM principal. Empty means ClientOption.Username when NewClient or
	// ResolveClientOption gets one, otherwise DefaultIAMUsername.
	Username string
	// RefreshBuffer is how long before a token expires a new one is used. Live connections
	// re-AUTH with the new token at a random moment in the first half of this window, so
	// they don't all re-AUTH at once. A buffer longer than the token source's own refresh
	// margin makes re-AUTHs resend the old token (see CustomTokenSource). Zero or negative
	// means DefaultRefreshBuffer.
	RefreshBuffer time.Duration
	// CustomTokenSource supplies the access tokens instead of Application Default
	// Credentials, for tokens you get your own way. If it caches tokens, it must hand out a
	// new one once the current token has less than RefreshBuffer left: rueidisgcp asks for a
	// new token then, and live connections would otherwise re-AUTH with the old one.
	CustomTokenSource oauth2.TokenSource
	// Scopes are the OAuth2 scopes requested from Application Default Credentials;
	// DefaultCloudPlatformScope when empty. They aren't used with CustomTokenSource.
	Scopes []string

	// DisableProactiveReauth turns off re-AUTH on live connections. Every new connection
	// still gets a fresh token, but a live connection keeps the token it authenticated with:
	// the credentials carry no RefreshAfter, so rueidis schedules no re-AUTH.
	DisableProactiveReauth bool
	// MaxAuthPerSecond spaces out the AUTHs that send a token, at most this many per second,
	// so many connections reconnecting at once don't all ask for authentication together.
	// Each NewClient, ResolveClientOption or NewIAMTokenProvider call has its own limit.
	// 0 disables it.
	//
	// The queue has no upper bound: when N connections authenticate at once, the last one
	// waits about N/MaxAuthPerSecond seconds, and so does a command waiting for that
	// connection, since AuthCredentialsFn gets no context. rueidis starts the handshake
	// timeout (Dialer.Timeout) and the re-AUTH timeout (ConnWriteTimeout) only after
	// AuthCredentialsFn returns, so the wait doesn't make those time out. Pick a rate that
	// authenticates all of a client's connections within a delay you can accept.
	MaxAuthPerSecond float64
	// OnCredentialError receives the token fetch errors that rueidisgcp works around instead
	// of returning: when a refresh fails, connections keep the last token that worked, and
	// the fetch is retried about every 30s. Use it for logs and metrics.
	//
	// Once that token has expired, rueidisgcp has no token to give a new connection, so its
	// HELLO goes without AUTH and fails with NOAUTH. That NOAUTH error goes to the command
	// that opened the connection, not to OnCredentialError, which still gets the fetch errors.
	//
	// Some fetch errors are returned instead, and don't come here: the new connection fails
	// with "rueidisgcp: token acquisition failed: ...". That happens when no fetch has
	// succeeded yet, and with DisableProactiveReauth once the last token has expired: rueidisgcp
	// then runs only for new connections, so it returns the fetch error rather than NOAUTH.
	OnCredentialError func(error)
	// FetchTimeout limits how long one request for a new token may take, so a token endpoint
	// that hangs can't block new connections. A timed-out request counts as a failed fetch
	// (see OnCredentialError). Zero or negative means DefaultTokenFetchTimeout.
	FetchTimeout time.Duration
}

// IAMTokenProvider gets, caches and refreshes IAM access tokens. NewIAMTokenProvider makes
// one, for callers who build the rueidis.ClientOption themselves.
type IAMTokenProvider interface {
	// Token returns a valid access token: the cached one, or a new one when the cached one
	// is within RefreshBuffer of expiring.
	Token(ctx context.Context) (*oauth2.Token, error)
	// AuthCredentialsFn returns the function to set as rueidis.ClientOption.AuthCredentialsFn.
	// rueidis calls it for every new connection and at every scheduled re-AUTH.
	AuthCredentialsFn() func(rueidis.AuthCredentialsContext) (rueidis.AuthCredentials, error)
}

// tokenProvider is the IAMTokenProvider behind NewIAMTokenProvider and WithIAMOptions.
type tokenProvider struct {
	source        oauth2.TokenSource // cached tokens; a new one RefreshBuffer before expiry
	refreshBuffer time.Duration
	username      string
	lastGood      atomic.Pointer[oauth2.Token] // the last token fetched successfully
	disableReauth bool
	limiter       *authLimiter
	onErr         func(error)
}

// NewIAMTokenProvider returns an IAMTokenProvider for opt, for callers who set
// rueidis.ClientOption.AuthCredentialsFn themselves instead of using NewClient. Tokens are
// cached and replaced RefreshBuffer before they expire, so a re-AUTH sends a new token, as
// long as the token source hands one out by then (see IAMAuthOptions.CustomTokenSource).
//
// ctx bounds the one-off setup (finding the credentials), which also stops after 15s when
// ctx has no earlier deadline. Token refreshes don't use ctx.
func NewIAMTokenProvider(ctx context.Context, opt IAMAuthOptions) (IAMTokenProvider, error) {
	if opt.Username == "" {
		opt.Username = DefaultIAMUsername
	}
	if opt.RefreshBuffer <= 0 {
		opt.RefreshBuffer = DefaultRefreshBuffer
	}
	if len(opt.Scopes) == 0 {
		opt.Scopes = []string{DefaultCloudPlatformScope}
	}

	setupCtx, cancel := context.WithTimeout(ctx, setupTimeout)
	defer cancel()
	// Token sources refresh for as long as the client lives, so they must not inherit ctx's
	// cancellation (a NewClient timeout would otherwise break every later refresh). The
	// setup itself still stops when setupCtx ends: callWithContext stops waiting for it.
	srcCtx := context.WithoutCancel(ctx)
	base, err := callWithContext(setupCtx, "credential setup", func() (oauth2.TokenSource, error) {
		return newBaseSource(srcCtx, opt)
	})
	if err != nil {
		return nil, err
	}

	p := &tokenProvider{
		// The cache hands out the same token until RefreshBuffer before it expires, then
		// fetches a new one. Re-AUTHs are scheduled inside that window, so they send a new
		// token. With the default 10s margin of oauth2.ReuseTokenSource they would
		// resend the old token and reschedule at ever shorter intervals until it expired.
		// For ADC, newBaseSource gives ADC's own cache the same margin, so the fetch gets a
		// new token too. The fetch timeout sits under the cache, so it limits real fetches,
		// not cache hits.
		source:        oauth2.ReuseTokenSourceWithExpiry(nil, withFetchTimeout(base, opt.FetchTimeout), opt.RefreshBuffer),
		refreshBuffer: opt.RefreshBuffer,
		username:      opt.Username,
		disableReauth: opt.DisableProactiveReauth,
		onErr:         opt.OnCredentialError,
	}
	if opt.MaxAuthPerSecond > 0 {
		interval := min(float64(time.Second)/opt.MaxAuthPerSecond, float64(24*time.Hour)) // clamp: no int64 overflow
		p.limiter = &authLimiter{interval: time.Duration(interval)}
	}
	return p, nil
}

// findDefaultCredentials looks up Application Default Credentials. It is a variable so tests
// can check the parameters without a metadata server.
var findDefaultCredentials = google.FindDefaultCredentialsWithParams

// newBaseSource creates the token source the cache fetches from: CustomTokenSource when set,
// otherwise Application Default Credentials. The source keeps ctx for later refreshes, so
// the caller passes a context that is never cancelled.
func newBaseSource(ctx context.Context, opt IAMAuthOptions) (oauth2.TokenSource, error) {
	if opt.CustomTokenSource != nil {
		return opt.CustomTokenSource, nil
	}
	// ADC keeps its own token cache, which renews tokens 10s before they expire unless told
	// otherwise. EarlyTokenRefresh makes it renew at the same point as rueidisgcp's cache, so a
	// re-AUTH gets a new token. oauth2 honors it only for metadata-server credentials
	// (Compute Engine, GKE, Cloud Run). Key files, user credentials and Workload Identity
	// Federation keep the 10s cache; for them, early renewal still relies on oauth2's
	// ReuseTokenSource sharing its token object with rueidisgcp's cache.
	creds, err := findDefaultCredentials(ctx, google.CredentialsParams{
		Scopes:            opt.Scopes,
		EarlyTokenRefresh: opt.RefreshBuffer,
	})
	if err != nil {
		return nil, fmt.Errorf("rueidisgcp: Application Default Credentials not found: %w", err)
	}
	return creds.TokenSource, nil
}

// Token returns the cached token, or fetches a new one when the cached one is within
// RefreshBuffer of expiring. ctx bounds how long the caller waits; each fetch is also
// limited by FetchTimeout.
func (p *tokenProvider) Token(ctx context.Context) (*oauth2.Token, error) {
	return callWithContext(ctx, "token fetch", func() (*oauth2.Token, error) {
		tok, err := p.source.Token()
		if err == nil {
			p.lastGood.Store(tok) // also when the caller has stopped waiting
		}
		return tok, err
	})
}

// AuthCredentialsFn returns the rueidis.ClientOption.AuthCredentialsFn of this provider. Each
// call hands rueidis the username and a token, and schedules the next re-AUTH.
func (p *tokenProvider) AuthCredentialsFn() func(rueidis.AuthCredentialsContext) (rueidis.AuthCredentials, error) {
	return func(rueidis.AuthCredentialsContext) (rueidis.AuthCredentials, error) {
		p.limiter.wait() // waits for its turn when MaxAuthPerSecond is set
		return p.credentials()
	}
}

// credentials returns the AUTH credentials. When a fetch fails, tolerate decides what to
// hand out instead.
func (p *tokenProvider) credentials() (rueidis.AuthCredentials, error) {
	tok, err := p.source.Token() // the cache: a new token only inside RefreshBuffer
	now := time.Now()
	if err != nil {
		return p.tolerate(now, err)
	}
	p.lastGood.Store(tok)
	creds := rueidis.AuthCredentials{Username: p.username, Password: tok.AccessToken}
	// A zero Expiry means the token never expires, so there is nothing to refresh.
	if !p.disableReauth && !tok.Expiry.IsZero() {
		creds.RefreshAfter = refreshTime(now, tok.Expiry, p.refreshBuffer)
	}
	return creds, nil
}

// tolerate decides what to hand rueidis when a token fetch fails. rueidis closes a live
// connection when AuthCredentialsFn returns an error during a re-AUTH, so a short outage of
// the token source must not become an error while connections can keep working:
//   - The last token that worked is still valid: use it, and retry the fetch in about 30s.
//   - It has expired: return empty credentials. rueidis then skips the re-AUTH, and the
//     live connection stays authenticated; a new connection fails with NOAUTH until a fetch
//     succeeds.
//   - No fetch has succeeded yet, or the last token has expired and DisableProactiveReauth
//     is set: return the error, without calling OnCredentialError. The new connection fails
//     with it.
func (p *tokenProvider) tolerate(now time.Time, err error) (rueidis.AuthCredentials, error) {
	last := p.lastGood.Load()
	lastValid := last != nil && (last.Expiry.IsZero() || last.Expiry.After(now))
	// Never succeeded: surface the real error. With DisableProactiveReauth the provider only
	// runs on new dials, so an expired token has no live connection to protect either.
	if last == nil || (p.disableReauth && !lastValid) {
		return rueidis.AuthCredentials{}, fmt.Errorf("rueidisgcp: token acquisition failed: %w", err)
	}
	if p.onErr != nil {
		p.onErr(err)
	}
	if p.disableReauth { // re-AUTH is off: hand out the valid token without scheduling a re-AUTH
		return rueidis.AuthCredentials{Username: p.username, Password: last.AccessToken}, nil
	}
	retry := now.Add(jitter(credentialRetry))
	if lastValid { // still valid: keep using it and retry soon
		return rueidis.AuthCredentials{Username: p.username, Password: last.AccessToken, RefreshAfter: retry}, nil
	}
	// Expired: with empty credentials rueidis skips the re-AUTH, and the connection stays
	// authenticated. A new connection fails with NOAUTH until the source recovers.
	return rueidis.AuthCredentials{RefreshAfter: retry}, nil
}

// refreshTime returns when a live connection should re-AUTH, for a token that expires at
// expiry: at a random moment in the first half of the refresh buffer, so connections don't
// all re-AUTH at once. It is never before expiry-buffer, because until then the cache hands
// out the same token again.
func refreshTime(now, expiry time.Time, buffer time.Duration) time.Time {
	start := expiry.Add(-buffer)
	if now.Before(start) {
		return start.Add(time.Duration(rand.Int64N(int64(buffer/2) + 1)))
	}
	if remaining := expiry.Sub(now); remaining > 0 {
		return now.Add(remaining / 2) // already inside the buffer: halfway to expiry
	}
	return now.Add(10 * time.Second) // already expired: try again soon
}

// jitter returns d minus a random amount of up to 10%, so timers set at the same moment
// don't fire together. The result is never more than d.
func jitter(d time.Duration) time.Duration {
	if d <= 0 {
		return d
	}
	return d - time.Duration(rand.Int64N(int64(d)/10+1))
}

// authLimiter spaces out the AUTHs of one provider for IAMAuthOptions.MaxAuthPerSecond: each
// caller waits until interval after the previous one. A nil limiter doesn't wait.
type authLimiter struct {
	mu       sync.Mutex
	next     time.Time // when the next caller may go
	interval time.Duration
}

// wait blocks until the caller's turn.
func (l *authLimiter) wait() {
	if l == nil || l.interval <= 0 {
		return
	}
	l.mu.Lock()
	now := time.Now()
	if l.next.Before(now) {
		l.next = now
	}
	sleep := l.next.Sub(now)
	l.next = l.next.Add(l.interval)
	l.mu.Unlock()
	time.Sleep(sleep)
}

// callWithContext runs fn and stops waiting when ctx ends, for calls that take no context
// or keep the one they get for later refreshes. fn then finishes in the background and its
// result is dropped.
func callWithContext[T any](ctx context.Context, what string, fn func() (T, error)) (T, error) {
	type result struct {
		v   T
		err error
	}
	ch := make(chan result, 1) // buffered: an abandoned fn never blocks
	go func() {
		v, err := fn()
		ch <- result{v, err}
	}()
	select {
	case r := <-ch:
		return r.v, r.err
	case <-ctx.Done():
		var zero T
		return zero, fmt.Errorf("rueidisgcp: %s: %w", what, ctx.Err())
	}
}

// withFetchTimeout limits each fetch from src to timeout (IAMAuthOptions.FetchTimeout).
// Zero or negative means DefaultTokenFetchTimeout: an unbounded fetch would hold the token
// cache's lock, and so block every new connection and re-AUTH, for as long as it hangs.
func withFetchTimeout(src oauth2.TokenSource, timeout time.Duration) *timeoutSource {
	if timeout <= 0 {
		timeout = DefaultTokenFetchTimeout
	}
	return &timeoutSource{src: src, timeout: timeout}
}

// timeoutSource stops waiting for a token fetch after timeout, so a token endpoint that hangs
// can't block new connections, nor every other caller waiting on the token cache, which
// holds its lock while it fetches. oauth2.TokenSource can't be cancelled, so a timed-out
// fetch keeps running in the background. Later calls share that fetch instead of starting
// another until it is 3*timeout old, so a hung endpoint costs at most one goroutine per
// 3*timeout. A call waits only until the shared fetch is timeout old: once the fetch has
// timed out, later calls fail at once, and the provider falls back to the last token that
// worked, instead of each caller waiting another full timeout.
type timeoutSource struct {
	src     oauth2.TokenSource
	timeout time.Duration

	mu  sync.Mutex
	cur *tokenFetch // the latest fetch, shared while it runs
}

// tokenFetch is one fetch from the underlying source, shared by the callers waiting for it.
type tokenFetch struct {
	start time.Time
	done  chan struct{} // closed when tok and err are set
	tok   *oauth2.Token
	err   error
}

// finished reports whether the fetch has returned.
func (f *tokenFetch) finished() bool {
	select {
	case <-f.done:
		return true
	default:
		return false
	}
}

// Token waits for the running fetch, starting one when none is running, until the fetch is
// timeout old.
func (s *timeoutSource) Token() (*oauth2.Token, error) {
	s.mu.Lock()
	f := s.cur
	if f == nil || f.finished() || time.Since(f.start) > 3*s.timeout {
		f = &tokenFetch{start: time.Now(), done: make(chan struct{})}
		s.cur = f
		go func() {
			f.tok, f.err = s.src.Token()
			close(f.done)
		}()
	}
	s.mu.Unlock()

	// A shared fetch has already used part of its timeout; one that has used all of it is
	// stuck, so fail now and let the provider fall back to the last token that worked.
	wait := s.timeout - time.Since(f.start)
	if wait <= 0 {
		return nil, fmt.Errorf("token fetch timed out after %v", s.timeout)
	}
	timer := time.NewTimer(wait)
	defer timer.Stop()
	select {
	case <-f.done:
		return f.tok, f.err
	case <-timer.C:
		return nil, fmt.Errorf("token fetch timed out after %v", s.timeout)
	}
}

// WithDefaultIAM authenticates with IAM access tokens from Application Default Credentials,
// with the username "default". It is the same as WithIAMOptions(IAMAuthOptions{}).
func WithDefaultIAM() Option {
	return WithIAMOptions(IAMAuthOptions{})
}

// WithIAMOptions authenticates with IAM access tokens configured by iamOpt. NewClient's ctx
// bounds the token setup (see NewIAMTokenProvider).
func WithIAMOptions(iamOpt IAMAuthOptions) Option {
	return func(c *config) { c.iam = &iamOpt }
}
