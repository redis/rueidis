package rueidis

import (
	"context"
	"testing"
	"time"
)

func TestWaitOrSkipRetry_DeadlineBuffer(t *testing.T) {
	r := newRetryer(func(attempts int, cmd Completed, err error) time.Duration {
		return 10 * time.Millisecond // This will force a 10ms delay.
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()

	start := time.Now() // Returns false immediately b/c 5ms < 10ms + buffer inherently
	retriable := r.WaitOrSkipRetry(ctx, 1, Completed{}, nil)
	duration := time.Since(start)

	if retriable {
		t.Errorf("Expected retriable to be false when deadline is shorter than delay")
	}
	if duration > 2*time.Millisecond {
		t.Errorf("Expected function to return immediately, but it took %v", duration)
	}
}
