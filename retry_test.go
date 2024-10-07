package rueidis

import (
	"context"
	"errors"
	"testing"
	"time"
)

type mockRetryHandler struct {
	WaitUntilNextRetryFunc func(ctx context.Context, attempts int, err error) bool
}

func (m *mockRetryHandler) WaitUntilNextRetry(ctx context.Context, attempts int, err error) bool {
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

		shouldRetry := r.WaitUntilNextRetry(nil, 0, nil)
		if shouldRetry {
			t.Error("WaitUntilNextRetry() = true; want false")
		}
	})

	t.Run("context is canceled", func(t *testing.T) {
		r := &retryer{
			RetryDelay: func(attempts int, err error) time.Duration {
				return time.Second
			},
		}

		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		shouldRetry := r.WaitUntilNextRetry(ctx, 0, nil)
		if !shouldRetry {
			t.Error("WaitUntilNextRetry() = false; want true")
		}
	})

	t.Run("context deadline is before delay", func(t *testing.T) {
		r := &retryer{
			RetryDelay: func(attempts int, err error) time.Duration {
				return time.Second
			},
		}

		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
		defer cancel()

		start := time.Now()
		shouldRetry := r.WaitUntilNextRetry(ctx, 0, nil)
		if shouldRetry {
			t.Error("WaitUntilNextRetry() = true; want false")
		}
		elapsed := time.Since(start)

		if elapsed > 100*time.Millisecond {
			t.Errorf("WaitUntilNextRetry() took %v; want < 100ms", elapsed)
		}
	})

	t.Run("wait until next retry", func(t *testing.T) {
		r := &retryer{
			RetryDelay: func(attempts int, err error) time.Duration {
				return 50 * time.Millisecond
			},
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		start := time.Now()
		shouldRetry := r.WaitUntilNextRetry(ctx, 0, nil)
		if !shouldRetry {
			t.Error("WaitUntilNextRetry() = false; want true")
		}
		elapsed := time.Since(start)

		if elapsed > 100*time.Millisecond {
			t.Errorf("WaitUntilNextRetry() took %v; want < 100ms", elapsed)
		}
	})

	t.Run("empty context", func(t *testing.T) {
		r := &retryer{
			RetryDelay: func(attempts int, err error) time.Duration {
				return 50 * time.Millisecond
			},
		}

		start := time.Now()
		shouldRetry := r.WaitUntilNextRetry(context.Background(), 0, nil)
		if !shouldRetry {
			t.Error("WaitUntilNextRetry() = false; want true")
		}
		elapsed := time.Since(start)

		if elapsed > 100*time.Millisecond {
			t.Errorf("WaitUntilNextRetry() took %v; want < 100ms", elapsed)
		}
	})
}
