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

// RetryDelayFn returns the delay that should be used before retrying the
// attempt. Will return negative delay if the delay could not be determined or do not retry.
type RetryDelayFn func(attempts int, err error) time.Duration

// defaultRetryDelayFn delays the next retry exponentially without considering the error.
// max delay is 1 second.
func defaultRetryDelayFn(attempts int, _ error) time.Duration {
	base := time.Microsecond * (1 << min(defaultMaxRetries, attempts))
	jitter := time.Microsecond * time.Duration(util.FastRand(1<<min(defaultMaxRetries, attempts)))
	return min(defaultMaxRetryDelay, base+jitter)
}

type retryHandler interface {
	// RetryDelay returns the delay that should be used before retrying the
	// attempt. Will return negative delay if the delay could not be determined or do
	// not retry.
	// If the delay is zero, the next retry should be attempted immediately.
	RetryDelay(attempts int, err error) time.Duration

	// WaitForRetry waits until the next retry should be attempted.
	WaitForRetry(ctx context.Context, duration time.Duration)

	// WaitOrSkipRetry waits until the next retry should be attempted
	// or returns false if the command should not be retried.
	// Returns false immediately if the command should not be retried.
	// Returns true after the delay if the command should be retried.
	WaitOrSkipRetry(ctx context.Context, attempts int, err error) bool
}

type retryer struct {
	RetryDelayFn RetryDelayFn
}

var _ retryHandler = (*retryer)(nil)

func newRetryer(retryDelayFn RetryDelayFn) *retryer {
	return &retryer{RetryDelayFn: retryDelayFn}
}

func (r *retryer) RetryDelay(attempts int, err error) time.Duration {
	return r.RetryDelayFn(attempts, err)
}

func (r *retryer) WaitForRetry(ctx context.Context, duration time.Duration) {
	if duration <= 0 {
		return
	}

	if ch := ctx.Done(); ch != nil {
		tm := time.NewTimer(duration)
		defer tm.Stop()
		select {
		case <-ch:
		case <-tm.C:
		}
	} else {
		time.Sleep(duration)
	}
}

func (r *retryer) WaitOrSkipRetry(
	ctx context.Context, attempts int, err error,
) bool {
	delay := r.RetryDelay(attempts, err)
	if delay < 0 {
		return false
	}
	if delay == 0 {
		return true
	}

	if dl, ok := ctx.Deadline(); !ok || time.Until(dl) > delay {
		r.WaitForRetry(ctx, delay)
		return true
	}
	return false
}
