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

// TestSingleFlightJoinerKeepsItsOwnFlightError: a waiter must get the error of
// the flight it waited on even when the next flight overlaps with it.
//
// The call struct is reused. do() clears the flight before it closes the
// channel, so the next flight can start while the previous waiters are still
// waking up. Anything per-flight kept on call itself is overwritten in that
// window, and the waiters then read the next flight's result instead of their
// own — nil, which is the failure this fix is about. Keeping err on the flight
// avoids it, and also removes the need to synchronize the read: the write
// happens before close(ch), and the read happens after <-ch.
func TestSingleFlightJoinerKeepsItsOwnFlightError(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	flightErr := errors.New("flight failed")
	const joiners = 50

	for range 200 {
		var (
			sg    = call{}
			block = make(chan struct{})
			errs  = make([]error, joiners)
			done  int64
		)

		go func() {
			sg.Do(context.Background(), func() error {
				<-block
				return flightErr
			})
		}()
		for sg.suppressing() != 1 {
			runtime.Gosched()
		}

		for j := range joiners {
			go func(j int) {
				errs[j] = sg.Do(context.Background(), func() error { return nil })
				atomic.AddInt64(&done, 1)
			}(j)
		}
		for sg.suppressing() != joiners+1 {
			runtime.Gosched()
		}

		// The next flight starts as soon as the current one clears its
		// counter, which is before the waiters are released.
		next := make(chan struct{})
		go func() {
			defer close(next)
			for sg.suppressing() != 0 {
				runtime.Gosched()
			}
			sg.Do(context.Background(), func() error { return nil })
		}()

		close(block)
		for atomic.LoadInt64(&done) != joiners {
			runtime.Gosched()
		}
		<-next

		for j := range errs {
			if errs[j] != flightErr {
				t.Fatalf("joiner %v got %v, want %v", j, errs[j], flightErr)
			}
		}
	}
}

// TestSingleFlightCancellableJoinerAtCompletion drives callers into the moment
// a flight completes, half of them able to be cancelled and half not, since the
// two take different branches of Do.
//
// do() clears c.fl before it releases the waiters, so a caller arriving in that
// window finds no flight and starts its own instead of joining. Whichever side
// of it a caller lands on, it must come back with the error of the flight it
// waited on, and none may stay parked.
func TestSingleFlightCancellableJoinerAtCompletion(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	flightErr := errors.New("flight failed")
	const callers = 20

	for range 200 {
		var (
			sg      call
			done    int64
			results = make([]error, callers)
		)
		for j := range callers {
			go func(j int) {
				ctx := context.Background()
				if j%2 == 0 { // half wait through the ctx branch of Do
					c, cancel := context.WithCancel(ctx)
					defer cancel()
					ctx = c
				}
				results[j] = sg.Do(ctx, func() error {
					runtime.Gosched()
					return flightErr
				})
				atomic.AddInt64(&done, 1)
			}(j)
		}
		for atomic.LoadInt64(&done) != callers {
			runtime.Gosched()
		}
		for j := range results {
			if results[j] != flightErr {
				t.Fatalf("caller %v got %v, want %v", j, results[j], flightErr)
			}
		}
	}
}

// TestSingleFlightCancelledJoinerAtCompletion: the same window, with the
// cancellable waiters cancelled around the time the flight ends. Each must come
// back with either its flight's error or the cancellation, never nil and never
// parked, and the runner must still release the waiters that stayed.
func TestSingleFlightCancelledJoinerAtCompletion(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	flightErr := errors.New("flight failed")
	const callers = 20

	for range 200 {
		var (
			sg      call
			done    int64
			release = make(chan struct{})
			results = make([]error, callers)
		)
		for j := range callers {
			go func(j int) {
				ctx, cancel := context.WithCancel(context.Background())
				if j%2 == 0 {
					defer cancel()
				} else {
					cancel() // already cancelled on arrival
				}
				results[j] = sg.Do(ctx, func() error {
					<-release
					return flightErr
				})
				atomic.AddInt64(&done, 1)
			}(j)
		}
		close(release)
		for atomic.LoadInt64(&done) != callers {
			runtime.Gosched()
		}
		for j := range results {
			if results[j] != flightErr && results[j] != context.Canceled {
				t.Fatalf("caller %v got %v", j, results[j])
			}
		}
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
