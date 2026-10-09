# rueidiscompatrate

`rueidiscompatrate` provides a 100% drop-in replacement for [`github.com/go-redis/redis_rate/v10`](https://github.com/go-redis/redis_rate), powered by `rueidislimiter`'s high-performance Generic Cell Rate Algorithm (GCRA) engine.

## Key Benefits

- **Drop-in Migration**: Retains the exact same struct fields, rate helpers, and method signatures as `redis_rate/v10`.
- **Zero Cache Misses on Migration**: Uses the identical key format (`"rate:" + key`), avoiding cold starts during rolling migrations from `go-redis`.
- **Zero `CROSSSLOT` Errors**: Single-key architecture (`KEYS[1]`) runs safely in Redis Cluster mode with or without hash tags.
- **Server Clock Accuracy**: Uses server-side `TIME` in Lua, preventing client NTP drift discrepancies.
- **Bonus APIs**: Adds `AllowAtMost` (partial capacity grants) and `Reset` (state clearance).

## Installation

```bash
go get github.com/redis/rueidis/rueidiscompatrate
```

## Migration Example

To migrate from `go-redis/redis_rate/v10`, simply swap the import and initialize the limiter:

```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/rueidis"
	"github.com/redis/rueidis/rueidiscompat"
	"github.com/redis/rueidis/rueidiscompatrate"
)

func main() {
	client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// Initialize via rueidiscompat.Adapter:
	rdb := rueidiscompat.NewAdapter(client)
	limiter := rueidiscompatrate.NewLimiter(rdb)

	// Or initialize directly from rueidis.Client:
	// limiter := rueidiscompatrate.NewLimiterFromClient(client)

	ctx := context.Background()

	// Define rate limit using standard helpers:
	limit := rueidiscompatrate.PerMinute(10) // 10 req/min, burst 10

	// Allow a single request
	res, err := limiter.Allow(ctx, "project:123", limit)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Allowed: %d, Remaining: %d, RetryAfter: %v, ResetAfter: %v\n",
		res.Allowed, res.Remaining, res.RetryAfter, res.ResetAfter)

	// Consume multiple tokens (all-or-nothing)
	res, err = limiter.AllowN(ctx, "project:123", limit, 2)
	if err != nil {
		panic(err)
	}

	// Partial allowance (bonus API): grant up to 5 tokens based on available capacity
	res, err = limiter.AllowAtMost(ctx, "project:123", limit, 5)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Granted %d tokens!\n", res.Allowed)

	// Reset state (bonus API)
	_ = limiter.Reset(ctx, "project:123")
}
```

## Upstream Field Mapping Reference

| `redis_rate/v10` Field | `rueidiscompatrate` Value | Description |
| :--- | :--- | :--- |
| `Result.Allowed` (`int`) | `int(redisResult.Granted)` | Tokens granted for this request (`0` when rejected). |
| `Result.Remaining` (`int`) | `int(redisResult.Remaining)` | Remaining token count in the bucket. |
| `Result.RetryAfter` (`time.Duration`) | `redisResult.RetryAfter` | `-1` on success; duration to wait before retry when throttled. |
| `Result.ResetAfter` (`time.Duration`) | `redisResult.ResetAfter` | Duration until the bucket completely refills to maximum capacity. |

