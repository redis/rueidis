
# rueidislimiter

This module provides an interface for fixed window rate limiting with precise control over limits and time windows. Inspired by GitHub's approach to scaling their API with a sharded, replicated rate limiter in Redis ([github.blog](https://github.blog/engineering/infrastructure/how-we-scaled-github-api-sharded-replicated-rate-limiter-redis/)).

## Features

- **Fixed Window Algorithm**: Implements a fixed window algorithm to control the number of actions (e.g., API requests) a user can perform within a specified time window.
- **Customizable Limits**: Allows configuration of request limits and time windows to suit various application requirements.
- **Distributed Rate Limiting**: Leverages Redis to maintain rate limit counters, ensuring consistency across distributed environments.
- **Reset Information**: Provides `ResetAtMs` timestamps to inform clients when they can retry requests.

## Installation

To install the `rueidislimiter` module, run:

```bash
go get github.com/redis/rueidis/rueidislimiter
```

## Usage

### Basic Rate Limiting Example

The following example demonstrates how to initialize a rate limiter with a custom request limit and time window, and how to check and allow requests based on an identifier (e.g., a user ID or IP address):

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
	// Initialize a new rate limiter with a limit of 5 requests per minute
	limiter, err := rueidislimiter.NewRateLimiter(rueidislimiter.RateLimiterOption{
		ClientOption: rueidis.ClientOption{InitAddress: []string{"localhost:6379"}},
		KeyPrefix:    "api_rate_limit",
		Limit:        5,
		Window:       time.Minute,
	})
	if err != nil {
		panic(err)
	}

	identifier := "user_123"

	// Check if a request is allowed
	result, err := limiter.Check(context.Background(), identifier)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Allowed: %v, Remaining: %d, RetryAfter: %v\n", result.Allowed, result.Remaining, result.RetryAfter)

	// Allow a request
	result, err = limiter.Allow(context.Background(), identifier)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Allowed: %v, Remaining: %d, RetryAfter: %v\n", result.Allowed, result.Remaining, result.RetryAfter)

	// Allow multiple requests
	result, err = limiter.AllowN(context.Background(), identifier, 3)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Allowed: %v, Remaining: %d, RetryAfter: %v\n", result.Allowed, result.Remaining, result.RetryAfter)
}
```

### API

#### `NewRateLimiter`

Creates a new rate limiter with the specified options:

- `ClientOption`: Options to connect to Redis.
- `KeyPrefix`: Prefix for Redis keys used by this limiter.
- `Limit`: Maximum number of allowed requests per window.
- `Window`: Time window duration for rate limiting. Must be greater than 1 millisecond.

```go
limiter, err := rueidislimiter.NewRateLimiter(rueidislimiter.RateLimiterOption{
	ClientOption: rueidis.ClientOption{InitAddress: []string{"localhost:6379"}},
	KeyPrefix:    "api_rate_limit",
	Limit:        5,
	Window:       time.Second,
})
```

#### `Check`

Checks if a request is allowed under the rate limit without incrementing the count.

```go
result, err := limiter.Check(ctx, "user_identifier")
```

Returns a `Result` struct:

- `Allowed`: Whether the request is allowed.
- `Remaining`: Number of remaining requests in the current window.
- `ResetAtMs`: Unix timestamp in milliseconds at which the rate limit will reset.

#### `Allow`

Allows a single request, incrementing the counter if allowed.

```go
result, err := limiter.Allow(ctx, "user_identifier")
```

#### `AllowN`

Allows `n` requests, incrementing the counter accordingly if allowed.

```go
result, err := limiter.AllowN(ctx, "user_identifier", 3)
```

- `n`: The number of requests to allow.

## Implementation Details

The `rueidislimiter` module employs Lua scripts executed within Redis to ensure atomic operations for checking and updating rate limits. This approach minimizes race conditions and maintains consistency across distributed systems.

By utilizing Redis's expiration capabilities, the module automatically resets rate limits after the specified time window, ensuring efficient memory usage and accurate rate-limiting behavior.

For more information on the design and implementation of Redis-based rate limiters, refer to GitHub's detailed account of scaling their API with a sharded, replicated rate limiter in Redis ([github.blog](https://github.blog/engineering/infrastructure/how-we-scaled-github-api-sharded-replicated-rate-limiter-redis/)).
