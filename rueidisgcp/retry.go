package rueidisgcp

import (
	"math/rand/v2"
	"time"

	"github.com/redis/rueidis"
)

// Backoff settings: full jitter over retryBase·2^attempt, capped at retryCap. After
// maxRetryAttempts attempts, memorystoreRetryDelay stops retrying.
const (
	retryBase        = 50 * time.Millisecond
	retryCap         = 2 * time.Second
	maxRetryAttempts = 20
)

// memorystoreRetryDelay is exponential backoff with full jitter, so clients that lose their
// connections together don't retry together. It is a rueidis.RetryDelayFn: negative stops
// retrying, 0 retries now, positive waits.
//
// rueidis calls it only for errors it retries: network errors, LOADING, and in cluster
// mode TRYAGAIN and CLUSTERDOWN. An error reply to a connection's HELLO, such as
// Memorystore's ERR_IAM_*, is not one of them: rueidis returns that error instead.
func memorystoreRetryDelay(attempts int, _ rueidis.Completed, _ error) time.Duration {
	if attempts >= maxRetryAttempts {
		return -1
	}
	return fullJitter(retryBase, retryCap, attempts)
}

// fullJitter returns a random duration in (0, min(base·2^attempt, limit)].
func fullJitter(base, limit time.Duration, attempt int) time.Duration {
	backoff := base << min(max(attempt, 0), 16)
	if backoff <= 0 || backoff > limit {
		backoff = limit
	}
	return time.Duration(rand.Int64N(int64(backoff)) + 1)
}
