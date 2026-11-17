package rueidis

import (
	"context"
	"sync"
	"testing"
	"time"
)

func TestRingPutOne_ContextCancellation(t *testing.T) {
	r := newRing(DefaultRingScale)
	ctx := context.Background()

	// Fill the ring so subsequent puts will block
	for i := 0; i < len(r.store); i++ {
		ch, err := r.PutOne(ctx, Completed{})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		_ = ch
	}

	// Now put with a cancellable ctx — should return promptly on cancel
	cancelCtx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(50 * time.Millisecond)
		cancel()
	}()

	start := time.Now()
	_, err := r.PutOne(cancelCtx, Completed{})
	elapsed := time.Since(start)

	if err == nil {
		t.Fatal("expected context cancellation error, got nil")
	}
	if err != context.Canceled {
		t.Fatalf("expected context.Canceled, got %v", err)
	}
	if elapsed > 2*time.Second {
		t.Fatalf("PutOne did not unblock promptly on cancel: took %v", elapsed)
	}
}

func TestRingPutMulti_ContextCancellation(t *testing.T) {
	r := newRing(DefaultRingScale)
	ctx := context.Background()

	// Fill the ring
	for i := 0; i < len(r.store); i++ {
		ch, err := r.PutMulti(ctx, []Completed{{}}, []RedisResult{})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		_ = ch
	}

	// Cancelled put
	cancelCtx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(50 * time.Millisecond)
		cancel()
	}()

	start := time.Now()
	_, err := r.PutMulti(cancelCtx, []Completed{{}}, []RedisResult{})
	elapsed := time.Since(start)

	if err == nil {
		t.Fatal("expected context cancellation error, got nil")
	}
	if err != context.Canceled {
		t.Fatalf("expected context.Canceled, got %v", err)
	}
	if elapsed > 2*time.Second {
		t.Fatalf("PutMulti did not unblock promptly on cancel: took %v", elapsed)
	}
}

func TestRingPutOne_TimeoutCancellation(t *testing.T) {
	r := newRing(DefaultRingScale)

	// Fill the ring
	ctx := context.Background()
	for i := 0; i < len(r.store); i++ {
		ch, err := r.PutOne(ctx, Completed{})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		_ = ch
	}

	// Timeout context
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	start := time.Now()
	_, err := r.PutOne(timeoutCtx, Completed{})
	elapsed := time.Since(start)

	if err == nil {
		t.Fatal("expected context timeout error, got nil")
	}
	if err != context.DeadlineExceeded {
		t.Fatalf("expected context.DeadlineExceeded, got %v", err)
	}
	if elapsed > 2*time.Second {
		t.Fatalf("PutOne did not unblock promptly on timeout: took %v", elapsed)
	}
}

func TestRingPutOne_DoesNotConsumeWriteCounter(t *testing.T) {
	// After cancellation, the write counter must be decremented so the next
	// non-cancelled put can find its slot. This test verifies the invariant
	// holds across cancellation + success cycles.
	const rounds = 10
	for round := 0; round < rounds; round++ {
		r := newRing(DefaultRingScale)
		ctx := context.Background()

		// Fill all slots
		chans := make([]chan RedisResult, len(r.store))
		for i := range chans {
			var err error
			chans[i], err = r.PutOne(ctx, Completed{})
			if err != nil {
				t.Fatalf("round %d: unexpected error: %v", round, err)
			}
		}

		// Cancelled put — must NOT consume a write slot
		cancelCtx, cancel := context.WithCancel(context.Background())
		cancel() // cancel immediately
		_, err := r.PutOne(cancelCtx, Completed{})
		if err == nil {
			t.Fatalf("round %d: expected error on cancelled ctx", round)
		}

		// One of the live channels should still accept — verifies counter not
		// advanced by the cancelled put.
		gotSlot := false
		for _, ch := range chans {
			select {
			case ch <- RedisResult{}:
				gotSlot = true
			default:
			}
		}
		if !gotSlot {
			t.Fatalf("round %d: no slot available after cancelled put (write counter consumed)", round)
		}
	}
}

func TestWaitContextOrSignal_StopsOnCancel(t *testing.T) {
	m := &sync.Mutex{}
	c1 := sync.NewCond(m)
	c2 := sync.NewCond(m)

	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan error, 1)
	go func() {
		m.Lock()
		err := waitContextOrSignal(ctx, c1, c2)
		m.Unlock()
		done <- err
	}()

	// Cancel via context
	go func() {
		time.Sleep(50 * time.Millisecond)
		cancel()
	}()

	select {
	case err := <-done:
		if err != context.Canceled {
			t.Fatalf("expected context.Canceled, got %v", err)
		}
	case <-time.After(time.Second):
		t.Fatal("waitContextOrSignal did not return after context cancel")
	}
}
