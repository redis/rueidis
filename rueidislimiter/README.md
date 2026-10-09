
# rueidislimiter

`rueidislimiter` is a high-performance, distributed rate limiting module for Redis.

By default, it uses the **Fixed Window** counter algorithm for simplicity and backward compatibility, while also providing the **Generic Cell Rate Algorithm (GCRA)** (leaky bucket) for smooth, continuous traffic pacing with burst tolerance.

## Features

- **Fixed Window by Default**: Proven counter-based rate limiting compatible with existing deployments.
- **Generic Cell Rate Algorithm (GCRA) Support**: Smooth leaky-bucket rate limiting that eliminates boundary burst spikes ("double-dipping") and paces requests evenly over time.
- **Single-Key Architecture for GCRA**: Operates on a single Redis key (`rueidislimiter:gcra:{<identifier>}`), cutting keyspace overhead in half compared to dual-key fixed-window limiters while keeping namespaces isolated to prevent key collision with Fixed Window.
- **Server-Authoritative Clock**: For GCRA, time is derived atomically inside Redis using `redis.call('TIME')`, eliminating synchronization discrepancies caused by client NTP drift.
- **Redis Cluster Safe (Zero `CROSSSLOT`)**: All operations use cluster hash tags (`{...}`) to ensure multi-key operations and single-key GCRA calls never trigger `CROSSSLOT` errors across cluster nodes.
- **Partial Allowance (`AllowAtMost`)**: Allows batch jobs or multi-token consumers to acquire "up to" available capacity without failing completely.
- **Programmatic Reset (`Reset`)**: Programmatically clears rate limiting state for an identifier.
- **Detailed Timing Diagnostics**: Provides `RetryAfter` (`-1` on success, duration to wait on throttle) and `ResetAfter` (duration until full capacity is restored).
- **Multi-Strategy Extensibility**: Seamlessly select between legacy Fixed Window (`AlgorithmFixedWindow = 0` default) and GCRA (`AlgorithmGCRA = 1`).
- **Zero-Allocation Hot Path**: Leverages pooled buffers (`sync.Pool`) for high-throughput, low-allocation execution.

## Installation

To install the `rueidislimiter` module, run:

```bash
go get github.com/redis/rueidis/rueidislimiter
```

## Usage

### Basic Rate Limiting Example (Fixed Window Default)

```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/rueidis"
	"github.com/redis/rueidis/rueidislimiter"
)

func main() {
	// Initialize a limiter: 10 requests per minute (Fixed Window by default)
	limiter, err := rueidislimiter.NewRateLimiter(rueidislimiter.RateLimiterOption{
		ClientOption: rueidis.ClientOption{InitAddress: []string{"127.0.0.1:6379"}},
		Limit:        10,
		Window:       time.Minute,
	})
	if err != nil {
		panic(err)
	}
	defer limiter.Close()

	ctx := context.Background()
	userID := "user_12345"

	// Check without consuming capacity
	chk, err := limiter.Check(ctx, userID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Initial state: Allowed=%v, Remaining=%d\n", chk.Allowed, chk.Remaining)

	// Consume 1 token
	res, err := limiter.Allow(ctx, userID)
	if err != nil {
		panic(err)
	}
	if res.Allowed {
		fmt.Printf("Allowed! Remaining: %d, ResetAfter: %v\n", res.Remaining, res.ResetAfter)
	} else {
		fmt.Printf("Rate limited! Retry after: %v\n", res.RetryAfter)
	}

	// Consume multiple tokens
	res, err = limiter.AllowN(ctx, userID, 3)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Allowed %d tokens! Remaining: %d\n", res.Granted, res.Remaining)
}
```

---

## Advanced Usage

### 1. Opting Into Generic Cell Rate Algorithm (GCRA)

To use smooth leaky-bucket rate limiting with burst tolerance, specify `Algorithm: rueidislimiter.AlgorithmGCRA`:

```go
limiter, err := rueidislimiter.NewRateLimiter(rueidislimiter.RateLimiterOption{
	ClientOption: rueidis.ClientOption{InitAddress: []string{"127.0.0.1:6379"}},
	Limit:        10,                                 // 10 req/s sustained
	Window:       time.Second,
	Burst:        50,                                 // burst capacity up to 50 tokens
	Algorithm:    rueidislimiter.AlgorithmGCRA,        // Opt-in GCRA
})
```

### 2. Burst Capacity (`Burst` & `WithBurst`)

GCRA separates the sustained **rate** from the **burst tolerance**. By default, `Burst` equals `Limit`. You can configure a larger burst buffer or override it dynamically:

```go
// Override burst dynamically per request:
res, err := limiter.Allow(ctx, "user_123", rueidislimiter.WithBurst(100))
```

### 3. Partial Token Consumption (`AllowAtMost`)

For batch tasks, `AllowAtMost` grants whatever capacity is currently available up to `n` tokens:

```go
// Request up to 10 tokens; if only 4 remain, 4 are granted rather than failing completely
res, err := limiter.AllowAtMost(ctx, "batch_job", 10)
if err != nil {
	panic(err)
}

fmt.Printf("Allowed: %v, Granted: %d, Remaining: %d\n", res.Allowed, res.Granted, res.Remaining)
```

### 4. Programmatic State Clearance (`Reset`)

Clear rate limiting state immediately for a specific identifier:

```go
err := limiter.Reset(ctx, "user_123")
if err != nil {
	panic(err)
}
```

### 5. Redis Cluster Usage (Zero `CROSSSLOT`)

All operations use cluster hash tags (`{...}`) to ensure multi-key operations and single-key GCRA calls never trigger `CROSSSLOT` errors across cluster nodes:

```go
limiter, err := rueidislimiter.NewRateLimiter(rueidislimiter.RateLimiterOption{
	ClientOption: rueidis.ClientOption{
		InitAddress: []string{"127.0.0.1:7010", "127.0.0.1:7011", "127.0.0.1:7012"},
	},
	Limit:     100,
	Window:    time.Minute,
	Algorithm: rueidislimiter.AlgorithmGCRA,
})
if err != nil {
	panic(err)
}
defer limiter.Close()

// Cluster routes by hash tag {tenant_1}; zero CROSSSLOT errors guaranteed
res, err := limiter.Allow(ctx, "{tenant_1}:user_42")
```

---

## Implementation Details

The `rueidislimiter` module executes an atomic Lua script inside Redis:
1. **Effects Replication (`redis.replicate_commands()`)**: Issued on Line 1 of the Lua script for compatibility with older Redis instances (e.g., Redis 6.2 on Memorystore/ElastiCache) where `redis.call('TIME')` is evaluated.
2. **Authoritative Server Time**: Derives the current timestamp via `redis.call('TIME')`, completely eliminating clock drift across distributed client hosts.
3. **Theoretical Arrival Time (TAT)**: Tracks the continuous arrival schedule using a single scalar timestamp. If an arrival is within burst tolerance ($\text{TAT} \le \text{now} + \tau$), tokens are granted and the new TAT is committed. Otherwise, the arrival is rejected without modifying state.
4. **Zero Keyspace Leaks**: Keys automatically expire based on the exact time required for the bucket to drain completely. Idle keys leave no residual state in Redis.
