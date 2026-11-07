# rueidisgcp: Google Cloud Memorystore Integration for rueidis

[![Go Reference](https://pkg.go.dev/badge/github.com/redis/rueidis/rueidisgcp.svg)](https://pkg.go.dev/github.com/redis/rueidis/rueidisgcp)

`rueidisgcp` is the official Google Cloud sub-module for [`rueidis`](https://github.com/redis/rueidis), integrating it with **Google Cloud Memorystore for Valkey**.

---

## 1. Native Google Cloud IAM Authentication

`rueidisgcp` connects rueidis to Memorystore for Valkey with IAM authentication instead of a password. `rueidisgcp.NewClient` takes the same `rueidis.ClientOption` as `rueidis.NewClient`, as `rueidisotel.NewClient` does, fills in Memorystore-friendly defaults, and adds IAM credentials: `WithDefaultIAM()` uses Application Default Credentials (ADC), and `WithIAMOptions` lets you choose the token source, username and refresh settings.

### 1.1 Features

* **The rueidis way:** pass any `rueidis.ClientOption`, with as many addresses as you need in `InitAddress` and `Standalone.ReplicaAddress`.
* **IAM authentication:** OAuth2 access tokens from ADC (GKE Workload Identity, the metadata server, or a service account key in `GOOGLE_APPLICATION_CREDENTIALS`), or from your own `oauth2.TokenSource`.
* **Token rotation on live connections:** connections re-AUTH with a new token before the old one expires, spread out so they don't all re-AUTH at once.
* **Refresh failure tolerance:** if a refresh fails, live connections keep the last valid token, and `OnCredentialError` reports the error.
* **Bounded setup and token fetches:** `ctx` bounds setup, and every token fetch is capped by `FetchTimeout`.
* **Memorystore defaults:** retry backoff with jitter, regular cluster topology refresh through the discovery endpoint, and reads sent to the reader endpoint. Timeouts stay rueidis's.
* **Works with other rueidis modules:** `rueidislock`, `rueidislimiter`, `rueidisaside` and `rueidisotel`.
* **No new dependencies for rueidis:** the Google Cloud dependencies stay in this sub-module.

### 1.2 Installation

```bash
go get github.com/redis/rueidis/rueidisgcp
```

### 1.3 Quickstart

#### Connect with Default IAM (Cluster Mode)

```go
package main

import (
	"context"
	"log"

	"github.com/redis/rueidis"
	"github.com/redis/rueidis/rueidisgcp"
)

func main() {
	ctx := context.Background()

	// IAM access tokens from Application Default Credentials (ADC) or GKE Workload Identity:
	client, err := rueidisgcp.NewClient(ctx, rueidis.ClientOption{
		InitAddress: []string{"10.0.0.15:6379"}, // discovery endpoint
	}, rueidisgcp.WithDefaultIAM())
	if err != nil {
		log.Fatalf("Failed to connect to Memorystore: %v", err)
	}
	defer client.Close()

	// Execute Valkey commands normally - tokens are refreshed in the background:
	err = client.Do(ctx, client.B().Set().Key("mykey").Value("myval").Build()).Error()
	if err != nil {
		log.Fatalf("SET failed: %v", err)
	}
}
```

`InitAddress` holds the discovery endpoint of a cluster-mode instance, or the primary endpoint of a cluster-disabled instance; rueidis detects which one it is when it connects. Any other `rueidis.ClientOption` field works as with `rueidis.NewClient`.

#### Cluster Mode Disabled with a Reader Endpoint

```go
client, err := rueidisgcp.NewClient(ctx, rueidis.ClientOption{
	InitAddress: []string{"10.0.0.2:6379"},                                         // primary endpoint
	Standalone:  rueidis.StandaloneOption{ReplicaAddress: []string{"10.0.0.3:6379"}}, // reader endpoint
}, rueidisgcp.WithDefaultIAM())
```

Read-only commands that don't block go to the reader endpoint. Leave out `Standalone.ReplicaAddress` for an instance without replicas.

#### Connect with Custom IAM Options

`WithDefaultIAM()` is `WithIAMOptions(rueidisgcp.IAMAuthOptions{})`: ADC with every default. Pass `WithIAMOptions` to change any of them:

```go
client, err := rueidisgcp.NewClient(ctx, rueidis.ClientOption{
	InitAddress: []string{"10.0.0.15:6379"},
}, rueidisgcp.WithIAMOptions(rueidisgcp.IAMAuthOptions{
	CustomTokenSource: tokenSource, // any oauth2.TokenSource instead of ADC
	OnCredentialError: func(err error) { log.Printf("IAM token refresh failed, keeping the last token: %v", err) },
	MaxAuthPerSecond:  50, // pace AUTHs when many connections reconnect at once
}))
```

| `IAMAuthOptions` field   | Default                                                                | Meaning                                                                                                  |
| :------------------------- | :--------------------------------------------------------------------- | :------------------------------------------------------------------------------------------------------- |
| `CustomTokenSource`      | ADC                                                                    | Where access tokens come from                                                                            |
| `Username`               | `ClientOption.Username`, else `"default"` (`DefaultIAMUsername`) | The AUTH username                                                                                        |
| `Scopes`                 | `cloud-platform` (`DefaultCloudPlatformScope`)                     | OAuth2 scopes requested from ADC; not used with `CustomTokenSource`                                    |
| `RefreshBuffer`          | 3m45s (`DefaultRefreshBuffer`)                                       | How long before expiry a token is replaced; live connections re-AUTH in the first half of that window    |
| `FetchTimeout`           | 10s (`DefaultTokenFetchTimeout`)                                     | Cap on one token fetch; zero or negative means the default                                               |
| `DisableProactiveReauth` | `false`                                                              | When true, live connections aren't re-AUTHed; every new connection still gets a fresh token              |
| `MaxAuthPerSecond`       | 0 (off)                                                                | Limits how many AUTHs the client sends per second; N connections at once wait up to about N/rate seconds |
| `OnCredentialError`      | none                                                                   | Receives token fetch errors while live connections keep the last valid token                             |

#### Use rueidis.NewClient Directly

`NewIAMTokenProvider` builds the IAM token provider on its own. Its `AuthCredentialsFn()` plugs into `rueidis.ClientOption`:

```go
provider, err := rueidisgcp.NewIAMTokenProvider(ctx, rueidisgcp.IAMAuthOptions{})
if err != nil {
	log.Fatal(err)
}
client, err := rueidis.NewClient(rueidis.ClientOption{
	InitAddress:       []string{"10.0.0.15:6379"},
	AuthCredentialsFn: provider.AuthCredentialsFn(),
})
```

This skips the Memorystore defaults; `rueidisgcp.ResolveClientOption` returns a `rueidis.ClientOption` with both the defaults and the credentials. `provider.Token(ctx)` returns the current access token, for example to check that ADC works.

#### Use with Other rueidis Modules

```go
// rueidislock, rueidislimiter and rueidisaside build their client through ClientBuilder:
locker, err := rueidislock.NewLocker(rueidislock.LockerOption{
	ClientOption: rueidis.ClientOption{InitAddress: []string{"10.0.0.15:6379"}},
	ClientBuilder: func(option rueidis.ClientOption) (rueidis.Client, error) {
		return rueidisgcp.NewClient(ctx, option, rueidisgcp.WithDefaultIAM())
	},
})

// rueidisotel builds the client itself, so pass it the ClientOption rueidisgcp prepared:
option, err := rueidisgcp.ResolveClientOption(ctx, rueidis.ClientOption{InitAddress: []string{"10.0.0.15:6379"}},
	rueidisgcp.WithDefaultIAM())
client, err := rueidisotel.NewClient(option)
```

### 1.4 Defaults

`NewClient` and `ResolveClientOption` fill these into the `rueidis.ClientOption` fields you leave at zero:

| Field              | Default                                                                                                                                                                                                                                                                                    |
| :----------------- | :----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `RetryDelay`     | Exponential backoff with full jitter, up to 2s; stops after 20 retries                                                                                                                                                                                                                     |
| `ClusterOption`  | Only when left entirely at zero: `ShardsRefreshInterval` 10s, `PreferInitAddressRefresh` and `PreferClusterShards` on, so the topology is refreshed through the discovery endpoint with `CLUSTER SHARDS`. Set any `ClusterOption` field and the whole `ClusterOption` is yours |
| `SendToReplicas` | Read-only commands that don't block go to the replicas; set only with `Standalone.ReplicaAddress`                                                                                                                                                                                        |

The cluster settings are left out when `Standalone.ReplicaAddress` is set. rueidis ignores them when it ends up with a single-node client.

Each of these follows Google's [Valkey client best practices](https://docs.cloud.google.com/memorystore/docs/valkey/general-best-practices#valkey_client_best_practices): reconnect with exponential backoff and jitter, discover nodes through the discovery endpoint with `CLUSTER SHARDS`, refresh the topology periodically, and read from the reader endpoint. Timeouts are not changed: `Dialer.Timeout`, `Dialer.KeepAlive` and `ConnWriteTimeout` keep rueidis's defaults (5s, 1s and 10s).

### 1.5 Token Refresh and Failures

* **When tokens are refreshed.** A token is replaced `RefreshBuffer` (3m45s) before it expires. Live connections re-AUTH with the new token at a random time in the first half of that window, so they don't all re-AUTH in the same second. `DisableProactiveReauth` turns this off; new connections still get a fresh token.
* **Why 3m45s.** It is the margin Google's own auth library uses ([`defaultExpiryDelta` in `cloud.google.com/go/auth`](https://github.com/googleapis/google-cloud-go/blob/auth/v0.24.0/auth/auth.go#L47-L49)): the metadata server can keep handing out the old token until 4 minutes before it expires, so asking earlier could get the same token again. rueidisgcp passes the same margin to Application Default Credentials (`EarlyTokenRefresh`), so their own token cache renews at the same point. The Go OAuth2 library applies it only to metadata-server credentials.
* **Custom token sources.** rueidisgcp asks for a new token once the current one has less than `RefreshBuffer` left. A `CustomTokenSource` that caches tokens must hand out a new one by then; otherwise live connections re-AUTH with the old token.
* **When a refresh fails.** Live connections keep the last valid token, and the refresh is retried about every 30 seconds. `OnCredentialError` receives each error. If the last token expires before a refresh succeeds, live connections stay authenticated (their re-AUTH is skipped), but new connections fail with `NOAUTH` until a token can be fetched again. Those `NOAUTH` errors go to the commands that opened the connections, not to `OnCredentialError`. With `DisableProactiveReauth`, an expired token isn't handed out at all: new connections fail with the token fetch error itself (`rueidisgcp: token acquisition failed: …`), not `NOAUTH`, and `OnCredentialError` isn't called. The same happens before the first fetch has succeeded.
* **Token fetches are capped.** Each fetch waits at most `FetchTimeout` (10s); a fetch that times out counts as a failed one. While a timed-out fetch is still hanging, the next fetches fail at once instead of waiting another `FetchTimeout`, and a new fetch is tried after 3×`FetchTimeout`.
* **A rejected re-AUTH closes the connection.** If Memorystore rejects the re-AUTH of a live connection (for example with `ERR_IAM_EXHAUSTED`), rueidis closes that connection and the next command dials a new one, although Memorystore would have kept the authenticated connection working.

### 1.6 Contexts and Setup Failures

* **`NewClient`'s `ctx` bounds setup only.** Credential discovery stops when it is cancelled or its deadline passes, and after 15 seconds if it has no deadline; a `ctx` that has ended by then stops `NewClient` before it connects. Token refreshes never use it, because they run for the life of the client.
* **Setup needs working credentials.** rueidis connects and authenticates while it builds the client, so missing ADC fails `NewClient` immediately, with the underlying error. Without ADC, pass a token source in `IAMAuthOptions.CustomTokenSource`.
* **`InitAddress` is required.** `NewClient` and `ResolveClientOption` return an error without it.
* **Credentials on the `ClientOption`** (`Password` or `AuthCredentialsFn`) are used as is; passing them together with `WithIAMOptions` or `WithDefaultIAM` is an error. With an IAM option, `ClientOption.Username` is the IAM username unless `IAMAuthOptions.Username` sets one; if both are set, they must match.

### 1.7 Required Roles

| Mechanism | Role                                                                       |
| :-------- | :------------------------------------------------------------------------- |
| IAM Auth  | `roles/memorystore.dbConnectionUser` (`memorystore.instances.connect`) |

Grant it to the identity behind the access tokens: the ADC identity, or the identity of `IAMAuthOptions.CustomTokenSource` when you set one.

### 1.8 Testing

Run the unit tests with Go's race detector:

```bash
go test -v -race .
```

The unit tests need no cloud access. `TestRefreshTolerance_LiveConnection` also runs against a local Redis on `127.0.0.1:6379` (the repository's `docker compose` setup) and is skipped when none is running.

Programs for manual checks live in `examples/`, a separate Go module, so they aren't part of the `rueidisgcp` module. Run them from the `examples` directory, for example `go run ./accelerated_test`:

| Program              | What it does                                                               | Needs                                                                            |
| :------------------- | :------------------------------------------------------------------------- | :------------------------------------------------------------------------------- |
| `test_adc`         | Fetches one IAM access token through ADC and prints its type and expiry    | ADC                                                                              |
| `accelerated_test` | Simulates token rotation with tokens that expire after 10 seconds          | Nothing                                                                          |
| `live_e2e`         | Sends SET/GET traffic to an instance until stopped, across token rotations | ADC, network access to the instance, `MEMORYSTORE_TEST_DISCOVERY_ENDPOINT` set |
