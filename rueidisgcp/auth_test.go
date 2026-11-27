package rueidisgcp

import (
	"context"
	"errors"
	"fmt"
	"net"
	"slices"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/redis/rueidis"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// mockTokenSource simulates a Google OAuth2 token source with configurable expiry and call tracking.
type mockTokenSource struct {
	mu         sync.Mutex
	calls      int64
	tokens     []*oauth2.Token
	defaultTok *oauth2.Token
	err        error
}

func (m *mockTokenSource) Token() (*oauth2.Token, error) {
	atomic.AddInt64(&m.calls, 1)
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.err != nil {
		return nil, m.err
	}
	if len(m.tokens) > 0 {
		tok := m.tokens[0]
		m.tokens = m.tokens[1:]
		return tok, nil
	}
	if m.defaultTok != nil {
		return m.defaultTok, nil
	}
	return &oauth2.Token{
		AccessToken: "mock-token-" + time.Now().Format(time.RFC3339Nano),
		TokenType:   "Bearer",
		Expiry:      time.Now().Add(1 * time.Hour),
	}, nil
}

func (m *mockTokenSource) CallCount() int64 {
	return atomic.LoadInt64(&m.calls)
}

// A provider built on a custom token source returns its token from Token, and
// AuthCredentialsFn turns that token into AUTH credentials: the configured username, the
// token as the password, and a re-AUTH time inside the refresh buffer.
func TestNewIAMTokenProvider_CustomTokenSource(t *testing.T) {
	ctx := context.Background()
	expectedToken := &oauth2.Token{
		AccessToken: "initial-secret-token",
		TokenType:   "Bearer",
		Expiry:      time.Now().Add(1 * time.Hour),
	}
	mock := &mockTokenSource{defaultTok: expectedToken}

	provider, err := NewIAMTokenProvider(ctx, IAMAuthOptions{
		Username:          "custom-user",
		RefreshBuffer:     10 * time.Minute,
		CustomTokenSource: mock,
	})
	if err != nil {
		t.Fatalf("unexpected error creating provider: %v", err)
	}

	tok, err := provider.Token(ctx)
	if err != nil {
		t.Fatalf("unexpected error fetching token: %v", err)
	}
	if tok.AccessToken != expectedToken.AccessToken {
		t.Errorf("got token %q, want %q", tok.AccessToken, expectedToken.AccessToken)
	}

	// Verify AuthCredentialsFn
	authFn := provider.AuthCredentialsFn()
	creds, err := authFn(rueidis.AuthCredentialsContext{})
	if err != nil {
		t.Fatalf("unexpected error in AuthCredentialsFn: %v", err)
	}
	if creds.Username != "custom-user" {
		t.Errorf("got username %q, want custom-user", creds.Username)
	}
	if creds.Password != expectedToken.AccessToken {
		t.Errorf("got password %q, want %q", creds.Password, expectedToken.AccessToken)
	}

	// RefreshAfter must fall in [Expiry-10m, Expiry-5m], the first half of the 10m refresh
	// buffer: inside the buffer, so the re-AUTH gets a new token from the cache, and at a
	// random point, so the connections of a client don't all re-AUTH at the same moment.
	windowStart := expectedToken.Expiry.Add(-10 * time.Minute)
	windowEnd := expectedToken.Expiry.Add(-5 * time.Minute)
	if creds.RefreshAfter.Before(windowStart) || creds.RefreshAfter.After(windowEnd) {
		t.Errorf("got RefreshAfter %v, want within [%v, %v]", creds.RefreshAfter, windowStart, windowEnd)
	}
}

// The token cache treats a token as expiring RefreshBuffer early, so a token with less time
// left than the buffer is replaced on the next call. With oauth2's default 10s margin, a
// re-AUTH scheduled inside the buffer would get the old token again and reschedule itself
// at ever shorter intervals (the "halving loop").
func TestTokenCache_RefreshesInsideBuffer(t *testing.T) {
	ctx := context.Background()

	// Token 1: Expires in 4 minutes (which is inside the 5-minute buffer)
	token1 := &oauth2.Token{
		AccessToken: "token-expiring-soon",
		TokenType:   "Bearer",
		Expiry:      time.Now().Add(4 * time.Minute),
	}
	// Token 2: Refreshed token expiring in 1 hour
	token2 := &oauth2.Token{
		AccessToken: "token-refreshed",
		TokenType:   "Bearer",
		Expiry:      time.Now().Add(1 * time.Hour),
	}

	mock := &mockTokenSource{
		tokens: []*oauth2.Token{token1, token2},
	}

	// Create provider with 5-minute buffer
	provider, err := NewIAMTokenProvider(ctx, IAMAuthOptions{
		RefreshBuffer:     5 * time.Minute,
		CustomTokenSource: mock,
	})
	if err != nil {
		t.Fatalf("failed to create provider: %v", err)
	}

	// First call consumes token1
	tok1, err := provider.Token(ctx)
	if err != nil {
		t.Fatalf("first token call failed: %v", err)
	}
	if tok1.AccessToken != "token-expiring-soon" {
		t.Errorf("got token %q, want token-expiring-soon", tok1.AccessToken)
	}
	if mock.CallCount() != 1 {
		t.Errorf("expected 1 underlying call, got %d", mock.CallCount())
	}

	// Under standard ReuseTokenSource, calling Token() now would see 4 minutes remaining (>10s)
	// and return the same token1 without calling mock.
	// But with ReuseTokenSourceWithExpiry(5m), remaining time (4m) < buffer (5m),
	// so it must fetch token2 right away.
	tok2, err := provider.Token(ctx)
	if err != nil {
		t.Fatalf("second token call failed: %v", err)
	}
	if tok2.AccessToken != "token-refreshed" {
		t.Errorf("got cached token %q, want the refreshed token %q", tok2.AccessToken, "token-refreshed")
	}
	if mock.CallCount() != 2 {
		t.Errorf("expected 2 underlying calls due to eager refresh, got %d", mock.CallCount())
	}
}

// When the token has less time left than the refresh buffer, the re-AUTH is scheduled after
// half the remaining time; when the token has already expired, 10s from now.
func TestAuthCredentialsFn_NearExpiryFallback(t *testing.T) {
	ctx := context.Background()

	// Case 1: Token expiring in 2 minutes with a 5-minute buffer
	nearExpiryToken := &oauth2.Token{
		AccessToken: "near-expiry-token",
		TokenType:   "Bearer",
		Expiry:      time.Now().Add(2 * time.Minute),
	}
	mock := &mockTokenSource{defaultTok: nearExpiryToken}

	provider, err := NewIAMTokenProvider(ctx, IAMAuthOptions{
		RefreshBuffer:     5 * time.Minute,
		CustomTokenSource: mock,
	})
	if err != nil {
		t.Fatalf("failed to create provider: %v", err)
	}

	authFn := provider.AuthCredentialsFn()
	creds, err := authFn(rueidis.AuthCredentialsContext{})
	if err != nil {
		t.Fatalf("authFn failed: %v", err)
	}

	// Remaining time is ~2m, so RefreshAfter should be ~1m in the future (remaining / 2)
	timeUntilRefresh := time.Until(creds.RefreshAfter)
	if timeUntilRefresh < 45*time.Second || timeUntilRefresh > 75*time.Second {
		t.Errorf("got timeUntilRefresh %v, want ~1m (half remaining)", timeUntilRefresh)
	}

	// Case 2: Already expired token -> should fall back to +10s
	expiredToken := &oauth2.Token{
		AccessToken: "expired-token",
		TokenType:   "Bearer",
		Expiry:      time.Now().Add(-1 * time.Minute),
	}
	mockExpired := &mockTokenSource{defaultTok: expiredToken}
	providerExpired, err := NewIAMTokenProvider(ctx, IAMAuthOptions{
		RefreshBuffer:     5 * time.Minute,
		CustomTokenSource: mockExpired,
	})
	if err != nil {
		t.Fatal(err)
	}
	credsExpired, err := providerExpired.AuthCredentialsFn()(rueidis.AuthCredentialsContext{})
	if err != nil {
		t.Fatalf("authFn on expired token failed: %v", err)
	}
	timeUntilExpiredRefresh := time.Until(credsExpired.RefreshAfter)
	if timeUntilExpiredRefresh < 8*time.Second || timeUntilExpiredRefresh > 12*time.Second {
		t.Errorf("got timeUntilExpiredRefresh %v, want ~10s", timeUntilExpiredRefresh)
	}
}

// Token and AuthCredentialsFn can be called from many goroutines at once: rueidis calls
// AuthCredentialsFn for each of its connections. Run with -race to catch data races.
func TestAuthCredentials_Concurrency(t *testing.T) {
	ctx := context.Background()
	mock := &mockTokenSource{
		defaultTok: &oauth2.Token{
			AccessToken: "concurrent-token",
			TokenType:   "Bearer",
			Expiry:      time.Now().Add(1 * time.Hour),
		},
	}
	provider, err := NewIAMTokenProvider(ctx, IAMAuthOptions{
		CustomTokenSource: mock,
	})
	if err != nil {
		t.Fatalf("failed to create provider: %v", err)
	}

	authFn := provider.AuthCredentialsFn()

	const goroutines = 20
	const iterations = 50
	var wg sync.WaitGroup
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				_, err := provider.Token(ctx)
				if err != nil {
					t.Errorf("concurrent Token() failed: %v", err)
				}
				creds, err := authFn(rueidis.AuthCredentialsContext{})
				if err != nil {
					t.Errorf("concurrent AuthCredentialsFn() failed: %v", err)
				}
				if creds.Username != DefaultIAMUsername {
					t.Errorf("unexpected username %q", creds.Username)
				}
			}
		}()
	}

	wg.Wait()
}

// WithIAMOptions sets ClientOption.AuthCredentialsFn to the IAM provider, which hands out
// the configured username and the token from the token source.
func TestWithIAMOptions_ClientOptionWiring(t *testing.T) {
	mock := &mockTokenSource{
		defaultTok: &oauth2.Token{
			AccessToken: "wired-token",
			TokenType:   "Bearer",
			Expiry:      time.Now().Add(1 * time.Hour),
		},
	}

	opt, err := ResolveClientOption(context.Background(), rueidis.ClientOption{InitAddress: []string{"10.0.0.15:6379"}},
		WithIAMOptions(IAMAuthOptions{
			Username:          "memorystore-user",
			CustomTokenSource: mock,
		}),
	)
	if err != nil {
		t.Fatalf("ResolveClientOption with WithIAMOptions failed: %v", err)
	}

	if opt.AuthCredentialsFn == nil {
		t.Fatal("expected AuthCredentialsFn to be configured on ClientOption, got nil")
	}

	creds, err := opt.AuthCredentialsFn(rueidis.AuthCredentialsContext{})
	if err != nil {
		t.Fatalf("AuthCredentialsFn execution failed: %v", err)
	}
	if creds.Username != "memorystore-user" {
		t.Errorf("got username %q, want memorystore-user", creds.Username)
	}
	if creds.Password != "wired-token" {
		t.Errorf("got password %q, want wired-token", creds.Password)
	}
}

// When the token source fails and no token has ever been fetched, AuthCredentialsFn returns
// the error, so rueidis fails the dial.
func TestAuthCredentialsFn_TokenError(t *testing.T) {
	ctx := context.Background()
	mock := &mockTokenSource{
		err: errors.New("network failure fetching oauth2 token"),
	}

	provider, err := NewIAMTokenProvider(ctx, IAMAuthOptions{
		CustomTokenSource: mock,
	})
	if err != nil {
		t.Fatalf("unexpected error creating provider: %v", err)
	}

	authFn := provider.AuthCredentialsFn()
	_, err = authFn(rueidis.AuthCredentialsContext{})
	if err == nil {
		t.Fatal("expected error from authFn when token acquisition fails, got nil")
	}
}

// Options left empty get their defaults: the AUTH username is DefaultIAMUsername.
func TestNewIAMTokenProvider_Defaults(t *testing.T) {
	ctx := context.Background()
	mock := &mockTokenSource{
		defaultTok: &oauth2.Token{
			AccessToken: "default-opts-token",
			TokenType:   "Bearer",
			Expiry:      time.Now().Add(1 * time.Hour),
		},
	}

	provider, err := NewIAMTokenProvider(ctx, IAMAuthOptions{
		CustomTokenSource: mock,
		// Leaving Username, RefreshBuffer, Scopes empty to trigger defaults
	})
	if err != nil {
		t.Fatalf("unexpected error creating provider: %v", err)
	}

	authFn := provider.AuthCredentialsFn()
	creds, err := authFn(rueidis.AuthCredentialsContext{})
	if err != nil {
		t.Fatalf("unexpected error in AuthCredentialsFn: %v", err)
	}
	if creds.Username != DefaultIAMUsername {
		t.Errorf("got username %q, want %q", creds.Username, DefaultIAMUsername)
	}
}

// Refresh-failure tolerance. rueidis closes a live connection when AuthCredentialsFn
// returns an error during a re-AUTH, so a failed token fetch must not reach it as an error
// while the connection can keep working.

// When a fetch fails while the last good token is still valid, the provider hands out that
// token again, schedules a retry in 27-30s, and passes the error to OnCredentialError.
func TestRefreshTolerance_CachedTokenValid(t *testing.T) {
	var fail atomic.Bool
	src := tokenSourceFunc(func() (*oauth2.Token, error) {
		if fail.Load() {
			return nil, errors.New("metadata server unreachable")
		}
		// Inside the default 3m45s refresh buffer, so the token cache fetches again on the next
		// call.
		return &oauth2.Token{AccessToken: "cached", Expiry: time.Now().Add(3 * time.Minute)}, nil
	})
	var absorbed atomic.Int32
	provider, err := NewIAMTokenProvider(context.Background(), IAMAuthOptions{
		CustomTokenSource: src,
		OnCredentialError: func(error) { absorbed.Add(1) },
	})
	if err != nil {
		t.Fatal(err)
	}
	authFn := provider.AuthCredentialsFn()
	if _, err := authFn(rueidis.AuthCredentialsContext{}); err != nil {
		t.Fatal(err)
	}
	fail.Store(true)
	before := time.Now()
	creds, err := authFn(rueidis.AuthCredentialsContext{})
	if err != nil {
		t.Fatalf("refresh failure surfaced as an error (would close the connection): %v", err)
	}
	if creds.Username != DefaultIAMUsername || creds.Password != "cached" {
		t.Errorf("creds = %+v, want the cached token", creds)
	}
	if d := creds.RefreshAfter.Sub(before); d < 27*time.Second || d > 30*time.Second+100*time.Millisecond {
		t.Errorf("retry in %v, want 27-30s", d)
	}
	if absorbed.Load() != 1 {
		t.Errorf("OnCredentialError called %d times, want 1", absorbed.Load())
	}
}

// Once the last good token has expired, a failed fetch hands out empty credentials, so
// rueidis skips the re-AUTH and the connection stays authenticated. A retry is still
// scheduled in 27-30s.
func TestRefreshTolerance_CachedTokenExpired(t *testing.T) {
	mock := &mockTokenSource{tokens: []*oauth2.Token{{AccessToken: "expired", Expiry: time.Now().Add(-time.Minute)}}}
	provider, err := NewIAMTokenProvider(context.Background(), IAMAuthOptions{CustomTokenSource: mock})
	if err != nil {
		t.Fatal(err)
	}
	authFn := provider.AuthCredentialsFn()
	if _, err := authFn(rueidis.AuthCredentialsContext{}); err != nil {
		t.Fatal(err)
	}
	mock.mu.Lock()
	mock.err = errors.New("sts unavailable")
	mock.mu.Unlock()
	before := time.Now()
	creds, err := authFn(rueidis.AuthCredentialsContext{})
	if err != nil {
		t.Fatalf("got error %v, want empty credentials", err)
	}
	if creds.Username != "" || creds.Password != "" {
		t.Errorf("creds = %+v, want empty (rueidis then skips the re-AUTH)", creds)
	}
	if d := creds.RefreshAfter.Sub(before); d < 27*time.Second || d > 30*time.Second+100*time.Millisecond {
		t.Errorf("retry in %v, want 27-30s", d)
	}
}

// If no fetch has ever succeeded, there is no token to fall back on: AuthCredentialsFn
// returns the error, and OnCredentialError isn't called for it.
func TestRefreshTolerance_NeverSucceeded(t *testing.T) {
	var absorbed atomic.Int32
	provider, err := NewIAMTokenProvider(context.Background(), IAMAuthOptions{
		CustomTokenSource: &mockTokenSource{err: errors.New("no credentials")},
		OnCredentialError: func(error) { absorbed.Add(1) },
	})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := provider.AuthCredentialsFn()(rueidis.AuthCredentialsContext{}); err == nil {
		t.Fatal("want the real error when no token was ever fetched")
	}
	if absorbed.Load() != 0 {
		t.Error("a surfaced error must not also go to OnCredentialError")
	}
}

// Against a real server, a token source that goes down before the first re-AUTH leaves the
// live connection in place: commands keep running on it (same CLIENT ID). Needs the docker
// compose Redis on 127.0.0.1:6379 (dockertest.sh starts it in CI); skipped otherwise.
func TestRefreshTolerance_LiveConnection(t *testing.T) {
	const addr = "127.0.0.1:6379"
	dialer := net.Dialer{Timeout: 300 * time.Millisecond}
	if c, err := dialer.DialContext(context.Background(), "tcp", addr); err != nil {
		t.Skipf("no local Redis at %s: %v", addr, err)
	} else {
		_ = c.Close()
	}
	var fail atomic.Bool
	var issued atomic.Int32
	src := tokenSourceFunc(func() (*oauth2.Token, error) {
		if fail.Load() {
			return nil, errors.New("metadata server unreachable")
		}
		return &oauth2.Token{AccessToken: fmt.Sprintf("tok-%d", issued.Add(1)), Expiry: time.Now().Add(3 * time.Second)}, nil
	})
	var absorbed atomic.Int32
	provider, err := NewIAMTokenProvider(context.Background(), IAMAuthOptions{
		CustomTokenSource: src,
		RefreshBuffer:     2 * time.Second, // first re-AUTH 1-2s after the dial
		OnCredentialError: func(error) { absorbed.Add(1) },
	})
	if err != nil {
		t.Fatal(err)
	}
	// One connection, so CLIENT ID tells whether it was replaced. The default user is
	// nopass, so AUTH default <token> succeeds like an accepted IAM token.
	client, err := rueidis.NewClient(rueidis.ClientOption{
		InitAddress:       []string{addr},
		ForceSingleClient: true,
		PipelineMultiplex: -1,
		DisableCache:      true,
		AuthCredentialsFn: provider.AuthCredentialsFn(),
	})
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()
	ctx := context.Background()
	id, err := client.Do(ctx, client.B().ClientId().Build()).AsInt64()
	if err != nil {
		t.Fatal(err)
	}
	fail.Store(true) // the token source goes down before the first refresh
	deadline := time.Now().Add(5 * time.Second)
	for absorbed.Load() == 0 && time.Now().Before(deadline) {
		time.Sleep(50 * time.Millisecond)
	}
	if absorbed.Load() == 0 {
		t.Fatal("the refresh never ran")
	}
	after, err := client.Do(ctx, client.B().ClientId().Build()).AsInt64()
	if err != nil {
		t.Fatalf("command failed after the refresh failure: %v", err)
	}
	if after != id {
		t.Fatalf("connection was replaced (CLIENT ID %d -> %d)", id, after)
	}
	if err := client.Do(ctx, client.B().Set().Key("rueidisgcp:refresh-tolerance").Value("ok").Build()).Error(); err != nil {
		t.Fatal(err)
	}
}

// refreshTime picks a random point in the first half of the refresh buffer; for a token
// already inside the buffer it waits half the remaining time, and for an expired one 10s.
// jitter takes up to 10% off a duration and passes non-positive durations through.
func TestRefreshTolerance_RefreshTimeWindow(t *testing.T) {
	now := time.Now()
	expiry := now.Add(time.Hour)
	const buffer = 5 * time.Minute
	for range 1000 {
		at := refreshTime(now, expiry, buffer)
		if at.Before(expiry.Add(-buffer)) || at.After(expiry.Add(-buffer/2)) {
			t.Fatalf("refresh at Expiry-%v, want within [Expiry-%v, Expiry-%v]", expiry.Sub(at), buffer, buffer/2)
		}
	}
	if got := refreshTime(now, now.Add(2*time.Minute), buffer).Sub(now); got != time.Minute {
		t.Errorf("near expiry: %v, want half the remaining time", got)
	}
	if got := refreshTime(now, now.Add(-time.Minute), buffer).Sub(now); got != 10*time.Second {
		t.Errorf("expired: %v, want 10s", got)
	}
	for range 1000 {
		if j := jitter(30 * time.Second); j < 27*time.Second || j > 30*time.Second {
			t.Fatalf("jitter = %v", j)
		}
	}
	if jitter(0) != 0 || jitter(-time.Second) != -time.Second {
		t.Error("jitter must pass non-positive durations through")
	}
}

// With DisableProactiveReauth, credentials never carry a re-AUTH time. After a failed
// fetch, new connections still get the last token while it is valid, but once it has
// expired the dial fails with the real error: there is no live connection to protect.
func TestRefreshTolerance_DisableProactiveReauth(t *testing.T) {
	var fail atomic.Bool
	expiry := time.Now().Add(3 * time.Minute)
	src := tokenSourceFunc(func() (*oauth2.Token, error) {
		if fail.Load() {
			return nil, errors.New("down")
		}
		return &oauth2.Token{AccessToken: "t", Expiry: expiry}, nil
	})
	provider, err := NewIAMTokenProvider(context.Background(), IAMAuthOptions{CustomTokenSource: src, DisableProactiveReauth: true})
	if err != nil {
		t.Fatal(err)
	}
	authFn := provider.AuthCredentialsFn()
	creds, err := authFn(rueidis.AuthCredentialsContext{})
	if err != nil || creds.Password != "t" || !creds.RefreshAfter.IsZero() {
		t.Fatalf("creds = %+v, %v; want zero RefreshAfter", creds, err)
	}
	// Dial-only mode: a valid cached token is still served, without arming a timer...
	fail.Store(true)
	creds, err = authFn(rueidis.AuthCredentialsContext{})
	if err != nil || creds.Password != "t" || !creds.RefreshAfter.IsZero() {
		t.Fatalf("cached: %+v, %v", creds, err)
	}
	// ...but an expired one fails the dial with the real error.
	provider.(*tokenProvider).lastGood.Store(&oauth2.Token{AccessToken: "old", Expiry: time.Now().Add(-time.Second)})
	if _, err := authFn(rueidis.AuthCredentialsContext{}); err == nil {
		t.Fatal("expired cached token in dial-only mode must fail")
	}
}

func TestAuthCredentials_NonExpiringToken(t *testing.T) {
	provider, err := NewIAMTokenProvider(context.Background(), IAMAuthOptions{
		CustomTokenSource: oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "forever"}),
	})
	if err != nil {
		t.Fatal(err)
	}
	creds, err := provider.AuthCredentialsFn()(rueidis.AuthCredentialsContext{})
	if err != nil || creds.Password != "forever" || !creds.RefreshAfter.IsZero() {
		t.Fatalf("creds = %+v, %v; a token without expiry needs no refresh", creds, err)
	}
}

// MaxAuthPerSecond spaces out the AUTHs of one provider, so a burst of new connections
// doesn't send all its AUTHs to Memorystore at once: at 20 per second, four AUTHs take at
// least 150ms. MaxAuthPerSecond 0 creates no limiter, and a nil limiter doesn't wait.
func TestMaxAuthPerSecond(t *testing.T) {
	provider, err := NewIAMTokenProvider(context.Background(), IAMAuthOptions{
		CustomTokenSource: &mockTokenSource{},
		MaxAuthPerSecond:  20, // one AUTH per 50ms
	})
	if err != nil {
		t.Fatal(err)
	}
	authFn := provider.AuthCredentialsFn()
	start := time.Now()
	for range 4 {
		if _, err := authFn(rueidis.AuthCredentialsContext{}); err != nil {
			t.Fatal(err)
		}
	}
	if elapsed := time.Since(start); elapsed < 140*time.Millisecond {
		t.Errorf("4 AUTHs took %v, want about 150ms or more: 3 waits of 50ms at 20/s", elapsed)
	}
	var nilLimiter *authLimiter
	nilLimiter.wait() // no-op, must not panic
	if p := provider.(*tokenProvider); p.limiter == nil || p.limiter.interval != 50*time.Millisecond {
		t.Errorf("limiter = %+v", p.limiter)
	}
	unpaced, err := NewIAMTokenProvider(context.Background(), IAMAuthOptions{CustomTokenSource: &mockTokenSource{}})
	if err != nil {
		t.Fatal(err)
	}
	if unpaced.(*tokenProvider).limiter != nil {
		t.Error("MaxAuthPerSecond 0 must not create a limiter")
	}
}

// Context and timeouts: setup honors the caller's ctx, and each upstream token fetch is
// capped by FetchTimeout.

// A hung setup call (NewIAMTokenProvider runs credential discovery through callWithContext)
// stops at the caller's deadline.
func TestCallWithContext_HonorsDeadline(t *testing.T) {
	release := make(chan struct{})
	t.Cleanup(func() { close(release) })
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Millisecond)
	defer cancel()
	start := time.Now()
	_, err := callWithContext(ctx, "credential setup", func() (oauth2.TokenSource, error) {
		<-release // a credential lookup that never answers
		return nil, errors.New("released")
	})
	if !errors.Is(err, context.DeadlineExceeded) || !strings.Contains(err.Error(), "credential setup") {
		t.Fatalf("got %v, want the caller's deadline", err)
	}
	if d := time.Since(start); d > 2*time.Second {
		t.Fatalf("setup took %v despite a 30ms deadline", d)
	}
}

// A hung token endpoint can't block a dial for longer than FetchTimeout, and later calls
// share the running fetch instead of piling up goroutines. Once that fetch has timed out,
// later calls fail at once instead of each waiting another FetchTimeout. Once the fetch is
// 3*FetchTimeout old, it counts as stuck and the next call starts a new one.
func TestFetchTimeout_HungSource(t *testing.T) {
	const timeout = 40 * time.Millisecond
	release := make(chan struct{})
	t.Cleanup(func() { close(release) })
	var calls atomic.Int32
	hung := tokenSourceFunc(func() (*oauth2.Token, error) {
		calls.Add(1)
		<-release
		return nil, errors.New("released")
	})
	provider, err := NewIAMTokenProvider(context.Background(), IAMAuthOptions{CustomTokenSource: hung, FetchTimeout: timeout})
	if err != nil {
		t.Fatal(err)
	}
	authFn := provider.AuthCredentialsFn()
	start := time.Now()
	if _, err := authFn(rueidis.AuthCredentialsContext{}); err == nil || !strings.Contains(err.Error(), "timed out") {
		t.Fatalf("got %v, want a fetch timeout (no token was ever fetched)", err)
	}
	if d := time.Since(start); d > timeout+time.Second {
		t.Fatalf("dial blocked for %v", d)
	}
	start = time.Now()
	if _, err := authFn(rueidis.AuthCredentialsContext{}); err == nil {
		t.Fatal("want a timeout while the fetch is still hung")
	}
	if d := time.Since(start); d >= timeout {
		t.Fatalf("second call waited %v; a fetch that has already timed out must fail at once", d)
	}
	if n := calls.Load(); n != 1 {
		t.Fatalf("upstream fetches = %d, want 1 while the first is still running", n)
	}
	time.Sleep(3 * timeout) // the running fetch now counts as stuck
	_, _ = authFn(rueidis.AuthCredentialsContext{})
	if n := calls.Load(); n != 2 {
		t.Fatalf("upstream fetches = %d, want a fresh fetch after 3*FetchTimeout", n)
	}
	for _, d := range []time.Duration{0, -1} {
		if s := withFetchTimeout(hung, d); s.timeout != DefaultTokenFetchTimeout {
			t.Errorf("FetchTimeout %v: cap = %v, want DefaultTokenFetchTimeout", d, s.timeout)
		}
	}
}

// A fetch that times out after an earlier success is handled like any failed fetch: the
// last valid token is handed out with a retry time, and OnCredentialError gets the timeout.
func TestFetchTimeout_FallsBackToLastGood(t *testing.T) {
	const timeout = 40 * time.Millisecond
	release := make(chan struct{})
	t.Cleanup(func() { close(release) })
	var calls atomic.Int32
	src := tokenSourceFunc(func() (*oauth2.Token, error) {
		if calls.Add(1) == 1 {
			// Inside the refresh buffer, so the token cache fetches again on the next call.
			return &oauth2.Token{AccessToken: "first", Expiry: time.Now().Add(time.Minute)}, nil
		}
		<-release
		return nil, errors.New("released")
	})
	var absorbed atomic.Int32
	provider, err := NewIAMTokenProvider(context.Background(), IAMAuthOptions{
		CustomTokenSource: src,
		FetchTimeout:      timeout,
		OnCredentialError: func(error) { absorbed.Add(1) },
	})
	if err != nil {
		t.Fatal(err)
	}
	authFn := provider.AuthCredentialsFn()
	if creds, err := authFn(rueidis.AuthCredentialsContext{}); err != nil || creds.Password != "first" {
		t.Fatalf("first AUTH: %+v, %v", creds, err)
	}
	creds, err := authFn(rueidis.AuthCredentialsContext{})
	if err != nil || creds.Password != "first" || creds.RefreshAfter.IsZero() {
		t.Fatalf("after a timeout: %+v, %v; want the last valid token and a retry", creds, err)
	}
	if absorbed.Load() != 1 {
		t.Fatalf("OnCredentialError called %d times, want 1", absorbed.Load())
	}
}

// IAMTokenProvider.Token stops waiting when its ctx ends, even when FetchTimeout is longer.
func TestTokenProvider_TokenHonorsContext(t *testing.T) {
	release := make(chan struct{})
	t.Cleanup(func() { close(release) })
	hung := tokenSourceFunc(func() (*oauth2.Token, error) {
		<-release
		return nil, errors.New("released")
	})
	provider, err := NewIAMTokenProvider(context.Background(), IAMAuthOptions{CustomTokenSource: hung, FetchTimeout: time.Minute})
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Millisecond)
	defer cancel()
	if _, err := provider.Token(ctx); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("got %v, want the caller's deadline", err)
	}
}

// Application Default Credentials are looked up with EarlyTokenRefresh set to RefreshBuffer,
// so ADC's own token cache renews at the same point as rueidisgcp's cache and a re-AUTH gets
// a new token, without relying on how oauth2's caches share token objects. A
// CustomTokenSource is used as is, with no ADC lookup.
func TestNewBaseSource_EarlyTokenRefresh(t *testing.T) {
	orig := findDefaultCredentials
	defer func() { findDefaultCredentials = orig }()
	var got google.CredentialsParams
	findDefaultCredentials = func(_ context.Context, params google.CredentialsParams) (*google.Credentials, error) {
		got = params
		return &google.Credentials{TokenSource: &mockTokenSource{}}, nil
	}
	for _, tc := range []struct {
		name   string
		opt    IAMAuthOptions
		buffer time.Duration
		scopes []string
	}{
		{"defaults", IAMAuthOptions{}, DefaultRefreshBuffer, []string{DefaultCloudPlatformScope}},
		{"custom", IAMAuthOptions{RefreshBuffer: 2 * time.Minute, Scopes: []string{"scope-a"}}, 2 * time.Minute, []string{"scope-a"}},
	} {
		got = google.CredentialsParams{}
		if _, err := NewIAMTokenProvider(context.Background(), tc.opt); err != nil {
			t.Fatalf("%s: %v", tc.name, err)
		}
		if got.EarlyTokenRefresh != tc.buffer || !slices.Equal(got.Scopes, tc.scopes) {
			t.Errorf("%s: EarlyTokenRefresh = %v, Scopes = %v; want %v, %v",
				tc.name, got.EarlyTokenRefresh, got.Scopes, tc.buffer, tc.scopes)
		}
	}

	var looked atomic.Bool
	findDefaultCredentials = func(context.Context, google.CredentialsParams) (*google.Credentials, error) {
		looked.Store(true)
		return nil, errors.New("ADC must not be looked up")
	}
	if _, err := NewIAMTokenProvider(context.Background(), IAMAuthOptions{CustomTokenSource: &mockTokenSource{}}); err != nil || looked.Load() {
		t.Fatalf("CustomTokenSource: err = %v, ADC looked up = %v", err, looked.Load())
	}
}
