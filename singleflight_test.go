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

	for i := 0; i < 1000; i++ {
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

	if atomic.LoadInt64(&err) != 1 {
		t.Fatalf("singleflight should that one call get the return value")
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

func TestSingleFlightLazyDo(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	ch := make(chan struct{})
	sg := call{}
	sg.LazyDo(time.Second, func() error {
		<-ch
		return nil
	})
	cn := 0
	sg.LazyDo(time.Second, func() error {
		cn++ // this should not occur
		return nil
	})
	if cn != 0 {
		t.Fatalf("unexpected cn %v", cn)
	}
	if sc := sg.suppressing(); sc != 1 {
		t.Fatalf("unexpected suppressing %v", sc)
	}
	close(ch)
}
