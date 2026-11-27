// Command test_adc checks that Application Default Credentials can produce the IAM access
// token rueidisgcp authenticates with. It builds an IAM token provider with default options
// (NewIAMTokenProvider with an empty IAMAuthOptions, as WithDefaultIAM does), fetches one
// token and prints its length, type and expiry. The token itself is never printed.
//
// It doesn't connect to Memorystore, so it needs no instance and no role, only working
// Application Default Credentials: for example from gcloud auth application-default login,
// from GOOGLE_APPLICATION_CREDENTIALS, or from the metadata server on Compute Engine, GKE
// or Cloud Run. Run it from the rueidisgcp/examples directory:
//
//	go run ./test_adc
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/rueidis/rueidisgcp"
)

func main() {
	ctx := context.Background()
	provider, err := rueidisgcp.NewIAMTokenProvider(ctx, rueidisgcp.IAMAuthOptions{})
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize ADC provider: %v", err))
	}

	tok, err := provider.Token(ctx)
	if err != nil {
		panic(fmt.Sprintf("Failed to acquire token from ADC: %v", err))
	}

	fmt.Printf("✅ Successfully acquired Google OAuth2 Access Token!\n")
	// The token is a credential, so only its length is shown.
	fmt.Printf("Token Length: %d characters (not printed)\n", len(tok.AccessToken))
	fmt.Printf("Token Type:   %s\n", tok.TokenType)
	fmt.Printf("Expiration:   %s (in %v)\n", tok.Expiry.Format(time.RFC3339), time.Until(tok.Expiry).Round(time.Second))
}
