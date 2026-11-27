// Command accelerated_test shows how rueidisgcp's IAM token provider caches tokens and
// schedules the proactive re-AUTH of live connections. The tokens here expire after 10
// seconds instead of an hour, so the timing shows in seconds.
//
// A fake token source (IAMAuthOptions.CustomTokenSource) hands out the tokens, so the
// program needs no credentials, instance or network. Run it from the rueidisgcp/examples
// directory:
//
//	go run ./accelerated_test
//
// It asks the provider's AuthCredentialsFn for credentials 15 times, one second apart, as
// rueidis does for each new connection, and logs the number of the token it got (never
// the token itself) and when the next re-AUTH would be due (AuthCredentials.RefreshAfter).
// With RefreshBuffer set to 3s, the provider's cache treats each token as expired 3s
// before its real expiry, so each token is handed out for 7s and the run sees three
// tokens. RefreshAfter falls in the first half of those last 3s, 7s to 8.5s after the
// token was issued, picked at random so that connections don't re-AUTH at once. The cache
// already treats the token as expired by then, so the re-AUTH gets a new token, not the
// old one.
package main

import (
	"context"
	"fmt"
	"log"
	"sync/atomic"
	"time"

	"github.com/redis/rueidis"
	"github.com/redis/rueidis/rueidisgcp"
	"golang.org/x/oauth2"
)

// shortLivedTokenSource is a fake oauth2.TokenSource. Each call returns a new token that
// expires 10s later; rotations counts the tokens handed out.
type shortLivedTokenSource struct {
	rotations int64
}

func (s *shortLivedTokenSource) Token() (*oauth2.Token, error) {
	rot := atomic.AddInt64(&s.rotations, 1)
	log.Printf("🔄 [TokenSource] Emitting fresh token #%d (expires in 10s)", rot)
	return &oauth2.Token{
		AccessToken: fmt.Sprintf("token-version-%d", rot),
		TokenType:   "Bearer",
		Expiry:      time.Now().Add(10 * time.Second),
	}, nil
}

func main() {
	mockSource := &shortLivedTokenSource{}

	provider, err := rueidisgcp.NewIAMTokenProvider(context.Background(), rueidisgcp.IAMAuthOptions{
		RefreshBuffer:     3 * time.Second, // re-AUTH 7s to 8.5s after each token is issued
		CustomTokenSource: mockSource,
	})
	if err != nil {
		log.Fatal(err)
	}

	authFn := provider.AuthCredentialsFn()

	const calls = 15
	log.Printf("Asking for credentials %d times, 1s apart, as rueidis does for each new connection...", calls)
	for range calls {
		creds, err := authFn(rueidis.AuthCredentialsContext{})
		if err != nil {
			log.Fatalf("authFn error: %v", err)
		}
		// Credentials don't belong in logs, even fake ones, so log the token's number: the
		// cache hands out the newest token, which is the last one the fake source issued.
		log.Printf("Got token #%d. Next re-AUTH due in: %v", atomic.LoadInt64(&mockSource.rotations),
			time.Until(creds.RefreshAfter).Round(time.Millisecond),
		)
		time.Sleep(1 * time.Second)
	}
	issued := atomic.LoadInt64(&mockSource.rotations)
	log.Printf("✅ Done: %d tokens issued in %ds, so the cache rotated the token %d times", issued, calls, issued-1)
}
