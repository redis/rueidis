package singleflight

import (
	"errors"
	"runtime"
	"sync/atomic"
	"testing"
)

func TestSingleFlight(t *testing.T) {
	var calls, waits, done, err int64

	sg := Call{}

	for i := 0; i < 1000; i++ {
		go func() {
			atomic.AddInt64(&waits, 1)

			if ret := sg.Do(func() error {
				atomic.AddInt64(&calls, 1)
				// wait for all goroutine invoked then return
				for atomic.LoadInt64(&waits) != 1000 {
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

	if atomic.LoadInt64(&calls) != 1 {
		t.Fatalf("singleflight should supress all concurrent calls")
	}

	if atomic.LoadInt64(&err) != 1 {
		t.Fatalf("singleflight should that one call get the return value")
	}
}
