package rueidis

import (
	"context"
	"errors"
	"testing"
	"time"
)

type mockRetryHandler struct {
	WaitUntilNextRetryFunc func(ctx context.Context, attempts int, err error) (bool, error)
}

func (m *mockRetryHandler) WaitUntilNextRetry(ctx context.Context, attempts int, err error) (bool, error) {
	return m.WaitUntilNextRetryFunc(ctx, attempts, err)
}

func TestDefaultRetryDelay(t *testing.T) {
	for i := 0; i < 100; i++ {
		err := errors.New("test")
		got := defaultRetryDelay(i, err)

		if got < 0 || got > defaultMaxRetryDelay {
			t.Errorf("defaultRetryDelay(%d, %v) = %v; want >= 0 and <= %v", i, err, got, defaultMaxRetryDelay)
		}
	}
}

func TestRetrier_WaitUntilNextRetry(t *testing.T) {
	t.Run("RetryDelay returns negative delay", func(t *testing.T) {
		r := &retryer{
			RetryDelay: func(attempts int, err error) time.Duration {
				return -1 * time.Second
			},
		}

		shouldRetry, err := r.WaitUntilNextRetry(nil, 0, nil)
		if shouldRetry {
			t.Error("WaitUntilNextRetry() = true; want false")
		}
		if err != nil {
			t.Errorf("WaitUntilNextRetry() = %v; want nil", err)
		}
	})

	t.Run("context is done", func(t *testing.T) {
		r := &retryer{
			RetryDelay: func(attempts int, err error) time.Duration {
				return time.Second
			},
		}

		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		shouldRetry, err := r.WaitUntilNextRetry(ctx, 0, nil)
		if shouldRetry {
			t.Error("WaitUntilNextRetry() = true; want false")
		}
		if !errors.Is(err, context.Canceled) {
			t.Error("WaitUntilNextRetry() = nil; want error")
		}
	})

	t.Run("wait until next retry", func(t *testing.T) {
		r := &retryer{
			RetryDelay: func(attempts int, err error) time.Duration {
				return 10 * time.Millisecond
			},
		}

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		start := time.Now()
		shouldRetry, err := r.WaitUntilNextRetry(ctx, 0, nil)
		if !shouldRetry {
			t.Error("WaitUntilNextRetry() = false; want true")
		}
		if err != nil {
			t.Errorf("WaitUntilNextRetry() = %v; want nil", err)
		}
		elapsed := time.Since(start)

		if elapsed > time.Second {
			t.Errorf("WaitUntilNextRetry() took %v; want >= 1s", elapsed)
		}
	})
}
