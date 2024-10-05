package rueidis

import (
	"context"
	"math"
	"time"

	"github.com/redis/rueidis/internal/util"
)

const (
	defaultMaxRetryDelay = 1 * time.Second
	maxAttemptsShift     = 63 // Avoid excessive shifts
)

// RetryDelay returns the delay that should be used before retrying the
// attempt. Will return negative delay if the delay could not be determined or do not retry.
type RetryDelay func(attempts int, err error) time.Duration

// defaultRetryDelay delays the next retry exponentially without considering the error.
// max delay is 1 second.
func defaultRetryDelay(attempts int, _ error) time.Duration {
	base := util.FastRandFloat64()
	backoff := uint64(1 << uint64(math.Min(float64(attempts), maxAttemptsShift)))
	jitter := base * float64(backoff) * float64(time.Millisecond)
	return time.Duration(math.Min(jitter, float64(defaultMaxRetryDelay)))
}

type retryHandler interface {
	// WaitUntilNextRetry waits until the next retry should be attempted.
	// Returns false without error immediately if the command should not be retried.
	// Returns false and an error if waiting for the next retry was interrupted.
	// Returns true after the delay if the command should be retried.
	WaitUntilNextRetry(ctx context.Context, attempts int, err error) (bool, error)
}

type retryer struct {
	RetryDelay RetryDelay
}

func newRetryer(retryDelay RetryDelay) *retryer {
	return &retryer{RetryDelay: retryDelay}
}

func (r *retryer) WaitUntilNextRetry(
	ctx context.Context, attempts int, err error,
) (bool, error) {
	delay := r.RetryDelay(attempts, err)
	if delay < 0 {
		return false, nil
	}

	timer := time.NewTimer(delay)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return false, ctx.Err()
	case <-timer.C:
		return true, nil
	}
}
