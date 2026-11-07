package rueidisgcp

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/redis/rueidis"
)

// Option configures what NewClient and ResolveClientOption add to a rueidis.ClientOption:
// IAM credentials (WithIAMOptions, WithDefaultIAM). Settings rueidis already has a field
// for, such as addresses and timeouts, are set on the rueidis.ClientOption itself.
// Later options override earlier ones of the same kind.
type Option func(*config)

// config collects the Options of one NewClient or ResolveClientOption call.
type config struct {
	iam *IAMAuthOptions // IAM settings (WithIAMOptions, WithDefaultIAM); nil when not given
}

// newConfig applies opts in order, skipping nil ones.
func newConfig(opts []Option) config {
	var c config
	for _, o := range opts {
		if o != nil {
			o(&c)
		}
	}
	return c
}

// Defaults NewClient fills into the rueidis.ClientOption fields left at zero. Timeouts are
// not among them: Dialer.Timeout, Dialer.KeepAlive and ConnWriteTimeout keep rueidis's
// defaults.
const (
	// defaultShardsRefreshInterval is how often a cluster client re-reads the cluster
	// topology. rueidis starts each refresh after a short random delay, so many clients
	// don't query the cluster at the same moment.
	defaultShardsRefreshInterval = 10 * time.Second
	// setupTimeout caps one-off setup calls (credential discovery) when the caller's context
	// has no earlier deadline.
	setupTimeout = 15 * time.Second
)

var errNoEndpoint = errors.New("rueidisgcp: no endpoint: set ClientOption.InitAddress")

// NewClient returns a rueidis.Client for Memorystore for Valkey. It takes the same
// rueidis.ClientOption as rueidis.NewClient: InitAddress holds the discovery endpoint (cluster
// mode) or the primary endpoint (cluster mode disabled), and Standalone.ReplicaAddress the
// reader endpoint, if any. NewClient fills the Memorystore defaults into the fields left at
// their zero value, adds the credentials opts ask for, and connects.
//
// ctx bounds credential setup; a ctx that has ended by then stops NewClient before it
// connects. rueidis connects and authenticates while it builds the client, so setup needs
// working credentials. Token refreshes never use ctx.
func NewClient(ctx context.Context, option rueidis.ClientOption, opts ...Option) (rueidis.Client, error) {
	opt, err := ResolveClientOption(ctx, option, opts...)
	if err != nil {
		return nil, err
	}
	if err := ctx.Err(); err != nil {
		return nil, err // rueidis.NewClient takes no context
	}
	return rueidis.NewClient(opt)
}

// ResolveClientOption returns the rueidis.ClientOption NewClient would connect with, without
// connecting, for callers who build the client themselves (rueidis.NewClient,
// rueidisotel.NewClient) or want to inspect it.
func ResolveClientOption(ctx context.Context, option rueidis.ClientOption, opts ...Option) (rueidis.ClientOption, error) {
	cfg := newConfig(opts)
	if len(option.InitAddress) == 0 {
		return rueidis.ClientOption{}, errNoEndpoint
	}
	if field := callerAuthField(option); field != "" && cfg.iam != nil {
		return rueidis.ClientOption{}, fmt.Errorf("rueidisgcp: ClientOption.%s conflicts with WithIAMOptions, which supplies the AUTH token", field)
	}
	opt := withDefaults(option)
	if cfg.iam != nil {
		iam, err := withCallerUsername(*cfg.iam, option.Username)
		if err != nil {
			return rueidis.ClientOption{}, err
		}
		p, err := NewIAMTokenProvider(ctx, iam)
		if err != nil {
			return rueidis.ClientOption{}, err
		}
		opt.AuthCredentialsFn = p.AuthCredentialsFn()
	}
	return opt, nil
}

// withDefaults fills the Memorystore defaults into the fields of opt left at their zero
// value.
func withDefaults(opt rueidis.ClientOption) rueidis.ClientOption {
	if len(opt.Standalone.ReplicaAddress) > 0 {
		if opt.SendToReplicas == nil {
			opt.SendToReplicas = sendReadsToReplicas
		}
	} else if opt.ClusterOption == (rueidis.ClusterOption{}) {
		// Cluster mode, or not known before connecting. rueidis reads these only in its
		// cluster client, so they are harmless when it falls back to a single-node client.
		// A false bool can't be told apart from an unset one, so ClusterOption is filled
		// only when the caller left all of it at zero: setting any of its fields leaves the
		// whole ClusterOption to the caller, who can then turn each of these off.
		opt.ClusterOption = rueidis.ClusterOption{
			ShardsRefreshInterval:    defaultShardsRefreshInterval,
			PreferInitAddressRefresh: true, // topology only through InitAddress
			PreferClusterShards:      true,
		}
	}
	if opt.RetryDelay == nil {
		opt.RetryDelay = memorystoreRetryDelay // exponential backoff with full jitter
	}
	return opt
}

// callerAuthField names the ClientOption credential field the caller set, if any.
// Caller-managed credentials are used as is; an IAM option next to them is an error.
// Username is not one of them: next to an IAM option it is the IAM username (see
// withCallerUsername).
func callerAuthField(o rueidis.ClientOption) string {
	switch {
	case o.AuthCredentialsFn != nil:
		return "AuthCredentialsFn"
	case o.Password != "":
		return "Password"
	}
	return ""
}

// withCallerUsername returns iam with ClientOption.Username as the IAM username when
// IAMAuthOptions.Username is empty. rueidis ignores ClientOption.Username once
// AuthCredentialsFn is set, so this keeps it from being dropped without a word. Two
// different usernames are an error: only one can be sent. iam is a copy, so the
// caller's Option is not changed.
func withCallerUsername(iam IAMAuthOptions, username string) (IAMAuthOptions, error) {
	switch {
	case username == "" || username == iam.Username:
	case iam.Username == "":
		iam.Username = username
	default:
		return IAMAuthOptions{}, fmt.Errorf("rueidisgcp: ClientOption.Username %q and IAMAuthOptions.Username %q differ; set only one of them", username, iam.Username)
	}
	return iam, nil
}

// sendReadsToReplicas routes read-only commands to the reader endpoint, except blocking
// ones.
func sendReadsToReplicas(cmd rueidis.Completed) bool {
	// A blocking read, such as XREAD with BLOCK, waits for new data. New data reaches the
	// primary first, so blocking reads stay on the primary.
	return cmd.IsReadOnly() && !cmd.IsBlock()
}
