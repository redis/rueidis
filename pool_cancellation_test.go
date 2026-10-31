package rueidis

import (
	"context"
	"errors"
	"sync/atomic"
	"testing"
	"time"

	"github.com/redis/rueidis/internal/cmds"
)

func TestPoolReturnedCanceledAcquirePreservesCapacity(t *testing.T) {
	for _, tc := range []struct {
		name     string
		attempts int
		deadline bool
	}{
		{name: "canceled", attempts: 1},
		{name: "deadline_while_waiting", attempts: 1, deadline: true},
		{name: "500_cancellations", attempts: 500},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var dials atomic.Int32
			p := newPool(2, deadFn(), 0, 0, func(context.Context) wire {
				dials.Add(1)
				return &mockWire{}
			})
			defer p.Close()
			first, second := p.Acquire(context.Background()), p.Acquire(context.Background())
			defer p.Store(first)
			defer p.Store(second)

			for range tc.attempts {
				ctx, cancel := context.WithCancel(context.Background())
				want := context.Canceled
				if tc.deadline {
					cancel()
					ctx, cancel = context.WithTimeout(context.Background(), 5*time.Millisecond)
					want = context.DeadlineExceeded
				} else {
					cancel()
				}
				w := p.Acquire(ctx)
				cancel()
				if !errors.Is(w.Error(), want) {
					t.Fatalf("Acquire error = %v, want %v", w.Error(), want)
				}
				// Blocking commands return the acquisition result even on failure.
				w.Close()
				p.Store(w)
			}
			if p.size != 2 {
				t.Errorf("pool size = %d, want 2 connections still checked out", p.size)
			}

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
			defer cancel()
			third := p.Acquire(ctx)
			if !errors.Is(third.Error(), context.DeadlineExceeded) {
				t.Errorf("third acquisition = %v, want deadline exceeded while both connections remain checked out", third.Error())
			}
			p.Store(third)
			if got := dials.Load(); got != 2 {
				t.Errorf("opened %d connections with capacity 2", got)
			}
		})
	}
}

func TestPoolReturnedClosedAcquirePreservesCount(t *testing.T) {
	p := newPool(2, deadFn(), 0, 0, func(context.Context) wire {
		t.Fatal("closed pool must not dial")
		return nil
	})
	p.Close()
	w := p.Acquire(context.Background())
	if !errors.Is(w.Error(), ErrClosing) {
		t.Fatalf("Acquire error = %v, want ErrClosing", w.Error())
	}
	p.Store(w)
	if p.size != 0 {
		t.Fatalf("closed pool size = %d, want 0", p.size)
	}
}

func TestPoolFailedConnectionStillReleasesCapacity(t *testing.T) {
	var dials int
	p := newPool(1, deadFn(), 0, 0, func(context.Context) wire {
		dials++
		if dials == 1 {
			return &mockWire{ErrorFn: func() error { return context.DeadlineExceeded }}
		}
		return &mockWire{}
	})
	defer p.Close()
	failed := p.Acquire(context.Background())
	if !errors.Is(failed.Error(), context.DeadlineExceeded) {
		t.Fatalf("dial error = %v, want deadline exceeded", failed.Error())
	}
	p.Store(failed)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	w := p.Acquire(ctx)
	if err := w.Error(); err != nil {
		t.Fatalf("replacement acquisition failed: %v", err)
	}
	defer p.Store(w)
	if dials != 2 || p.size != 1 {
		t.Fatalf("replacement: dials=%d, size=%d; want 2 dials and 1 connection", dials, p.size)
	}
}

func TestPoolUnacquiredStreamsPreserveCapacity(t *testing.T) {
	for _, closed := range []bool{false, true} {
		p := newPool(1, deadFn(), 0, 0, func(context.Context) wire {
			t.Fatal("failed acquisition must not dial")
			return nil
		})
		ctx, cancel := context.WithCancel(context.Background())
		want := context.Canceled
		if closed {
			p.Close()
			want = ErrClosing
		} else {
			cancel()
		}
		w := p.Acquire(ctx)
		single := w.DoStream(ctx, p, cmds.NewCompleted([]string{"PING"}))
		if !errors.Is(single.Error(), want) {
			t.Errorf("single stream error=%v, want %v", single.Error(), want)
		}
		multi := w.DoMultiStream(ctx, p, cmds.NewCompleted([]string{"PING"}))
		if !errors.Is(multi.Error(), want) {
			t.Errorf("multi stream error=%v, want %v", multi.Error(), want)
		}
		if p.size != 0 {
			t.Errorf("stream failures changed pool size to %d, want 0", p.size)
		}
		cancel()
		p.Close()
	}
}
