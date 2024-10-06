package rueidis

import (
	"context"
	"time"

	"github.com/redis/rueidis/internal/util"
)

const (
	defaultMaxRetries    = 20
	defaultMaxRetryDelay = 1 * time.Second
)

// RetryDelay returns the delay that should be used before retrying the
// attempt. Will return negative delay if the delay could not be determined or do not retry.
type RetryDelay func(attempts int, err error) time.Duration

// defaultRetryDelay delays the next retry exponentially without considering the error.
// max delay is 1 second.
func defaultRetryDelay(attempts int, _ error) time.Duration {
	base := time.Microsecond * (1 << min(defaultMaxRetries, attempts))
	jitter := time.Microsecond * time.Duration(util.FastRand(1<<min(defaultMaxRetries, attempts)))
	return min(defaultMaxRetryDelay, base+jitter)
}

type retryHandler interface {
	// WaitUntilNextRetry waits until the next retry should be attempted.
	// Returns false immediately if the command should not be retried.
	// Returns true after the delay if the command should be retried.
	WaitUntilNextRetry(ctx context.Context, attempts int, err error) bool
}

type retryer struct {
	RetryDelay RetryDelay
}

func newRetryer(retryDelay RetryDelay) *retryer {
	return &retryer{RetryDelay: retryDelay}
}

func (r *retryer) WaitUntilNextRetry(
	ctx context.Context, attempts int, err error,
) bool {
	delay := r.RetryDelay(attempts, err)
	if delay < 0 {
		return false
	}

	if dl, ok := ctx.Deadline(); !ok || time.Until(dl) > delay {
		if ch := ctx.Done(); ch != nil {
			tm := time.NewTimer(delay)
			defer tm.Stop()
			select {
			case <-ch:
			case <-tm.C:
			}
		} else {
			time.Sleep(delay)
		}
		return true
	}
	return false
}
