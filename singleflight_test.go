package rueidis

import (
	"context"
	"errors"
	"runtime"
	"sync/atomic"
	"testing"
	"time"
)

func TestSingleFlight(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	var calls, done, err int64

	sg := call{}

	for range 1000 {
		go func() {
			if ret := sg.Do(context.Background(), func() error {
				atomic.AddInt64(&calls, 1)
				// wait for all goroutine invoked then return
				for sg.suppressing() != 1000 {
					runtime.Gosched()
				}
				return errors.New("I should be the only ret")
			}); ret != nil {
				atomic.AddInt64(&err, 1)
			}

			atomic.AddInt64(&done, 1)
		}()
	}

	for atomic.LoadInt64(&done) != 1000 {
		runtime.Gosched()
	}

	if atomic.LoadInt64(&calls) == 0 {
		t.Fatalf("singleflight not call at all")
	}

	if v := atomic.LoadInt64(&calls); v != 1 {
		t.Fatalf("singleflight should suppress all concurrent calls, got: %v", v)
	}

	// Every caller must see the error: the one that ran fn and everyone who
	// waited on it. Waiters used to get nil, and nil looks like success to
	// code that retries until an operation succeeds.
	if v := atomic.LoadInt64(&err); v != 1000 {
		t.Fatalf("all callers should get the error of the run they waited on, got: %v", v)
	}
}

// TestSingleFlightJoinerReceivesFlightError: a caller that waits for an
// already-running fn must get the error that fn actually returned.
//
// It used to get nil even when fn failed. nil looks like success, so code that
// retries an operation until it succeeds stopped retrying after a run that
// failed: sentinelClient.refreshRetry loops until refresh() returns nil, so
// joining someone else's failed refresh ended the retry loop with the master
// still unresolved. The test also checks the other direction: the error of a
// finished run must not leak into the next run.
func TestSingleFlightJoinerReceivesFlightError(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	block := make(chan struct{})
	flightErr := errors.New("flight failed")
	sg := call{}

	initiatorDone := make(chan error, 1)
	go func() {
		initiatorDone <- sg.Do(context.Background(), func() error {
			<-block
			return flightErr
		})
	}()
	for sg.suppressing() != 1 {
		runtime.Gosched()
	}

	joinerDone := make(chan error, 1)
	go func() {
		joinerDone <- sg.Do(context.Background(), func() error {
			t.Error("joiner fn must not run")
			return nil
		})
	}()
	for sg.suppressing() != 2 {
		runtime.Gosched()
	}

	close(block)
	if err := <-initiatorDone; err != flightErr {
		t.Fatalf("initiator: unexpected err %v", err)
	}
	if err := <-joinerDone; err != flightErr {
		t.Fatalf("joiner: unexpected err %v", err)
	}

	// A caller arriving after the flight completed starts a fresh flight and
	// gets its own result, not the previous flight's error.
	if err := sg.Do(context.Background(), func() error { return nil }); err != nil {
		t.Fatalf("fresh flight: unexpected err %v", err)
	}
}

func TestSingleFlightWithContext(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	ch := make(chan struct{})
	sg := call{}
	go func() {
		sg.Do(context.Background(), func() error {
			<-ch
			return nil
		})
	}()
	for sg.suppressing() != 1 {
		time.Sleep(time.Millisecond)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := sg.Do(ctx, func() error { return nil }); err != context.Canceled {
		t.Fatalf("unexpected err %v", err)
	}
	go func() {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		if err := sg.Do(ctx, func() error { return nil }); err != nil {
			t.Errorf("unexpected err %v", err)
		}
	}()
	for sg.suppressing() != 3 {
		time.Sleep(time.Millisecond)
	}
	close(ch)
	if err := sg.Do(context.Background(), func() error { return nil }); err != nil {
		t.Fatalf("unexpected err %v", err)
	}
}

func TestSingleFlightDelayDoDedupesInFlight(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	ch := make(chan struct{})
	sg := call{}
	sg.DelayDo(0, func() error {
		<-ch
		return nil
	})
	cn := 0
	sg.DelayDo(0, func() error {
		cn++ // dedupe: should not run while first is in-flight
		return nil
	})
	if cn != 0 {
		t.Fatalf("DelayDo did not dedupe, cn=%v", cn)
	}
	if sc := sg.suppressing(); sc != 1 {
		t.Fatalf("unexpected suppressing %v", sc)
	}
	close(ch)
}

func TestSingleFlightDelayDoHonorsDelay(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	sg := call{}
	delay := 75 * time.Millisecond
	start := time.Now()
	done := make(chan time.Time, 1)
	sg.DelayDo(delay, func() error {
		done <- time.Now()
		return nil
	})
	select {
	case ts := <-done:
		got := ts.Sub(start)
		if got < delay {
			t.Fatalf("DelayDo ran too early: waited %v, expected >= %v", got, delay)
		}
	case <-time.After(time.Second):
		t.Fatalf("DelayDo never ran")
	}
}
