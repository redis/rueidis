// Package rueidisgcp connects rueidis to Google Cloud Memorystore for Valkey with native
// Google Cloud authentication.
//
// NewClient takes the same rueidis.ClientOption as rueidis.NewClient, plus Options, and
// returns a ready rueidis.Client:
//
//	client, err := rueidisgcp.NewClient(ctx, rueidis.ClientOption{
//		InitAddress: []string{"10.0.0.15:6379"}, // the instance's discovery endpoint
//	}, rueidisgcp.WithDefaultIAM())
//
// NewClient fills Memorystore defaults (retry backoff, cluster topology refresh through the
// discovery endpoint, reads sent to the reader endpoint) into the fields left empty, and
// adds the credentials its Options ask for:
//   - WithDefaultIAM authenticates with IAM access tokens from Application Default
//     Credentials, which are replaced on live connections before they expire.
//   - WithIAMOptions does the same with your own token source, username or refresh settings.
//
// ResolveClientOption returns the same rueidis.ClientOption without connecting, for
// rueidis.NewClient or other rueidis modules such as rueidisotel.
package rueidisgcp
