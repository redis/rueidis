package rueidis

import (
	"errors"
	"runtime"
	"sync/atomic"
	"testing"
)

var dead = deadFn()

//gocyclo:ignore
func TestPool(t *testing.T) {
	setup := func(size int) (*pool, *int32) {
		var count int32
		return newPool(size, dead, func() wire {
			atomic.AddInt32(&count, 1)
			closed := false
			return &mockWire{
				CloseFn: func() {
					closed = true
				},
				ErrorFn: func() error {
					if closed {
						return ErrClosing
					}
					return nil
				},
			}
		}), &count
	}

	t.Run("DefaultPoolSize", func(t *testing.T) {
		p := newPool(0, dead, func() wire { return nil })
		if cap(p.list) == 0 {
			t.Fatalf("DefaultPoolSize is not applied")
		}
	})

	t.Run("Reuse", func(t *testing.T) {
		pool, count := setup(100)
		for i := 0; i < 1000; i++ {
			pool.Store(pool.Acquire())
		}
		if atomic.LoadInt32(count) != 1 {
			t.Fatalf("pool does not reuse connection")
		}
	})

	t.Run("NotExceed", func(t *testing.T) {
		conn := make([]wire, 100)
		pool, count := setup(len(conn))
		for i := 0; i < len(conn); i++ {
			conn[i] = pool.Acquire()
		}
		if atomic.LoadInt32(count) != 100 {
			t.Fatalf("unexpected acquire count")
		}
		go func() {
			for i := 0; i < len(conn); i++ {
				pool.Store(conn[i])
			}
		}()
		for i := 0; i < len(conn); i++ {
			pool.Acquire()
		}
		if atomic.LoadInt32(count) > 100 {
			t.Fatalf("pool must not exceed the size limit")
		}
	})

	t.Run("NoShare", func(t *testing.T) {
		conn := make([]wire, 100)
		pool, _ := setup(len(conn))
		for i := 0; i < len(conn); i++ {
			w := pool.Acquire()
			go pool.Store(w)
		}
		for i := 0; i < len(conn); i++ {
			conn[i] = pool.Acquire()
		}
		for i := 0; i < len(conn); i++ {
			for j := i + 1; j < len(conn); j++ {
				if conn[i] == conn[j] {
					t.Fatalf("pool must not output acquired connection")
				}
			}
		}
	})

	t.Run("Close", func(t *testing.T) {
		pool, count := setup(2)
		w1 := pool.Acquire()
		w2 := pool.Acquire()
		if w1.Error() != nil {
			t.Fatalf("unexpected err %v", w1.Error())
		}
		if w2.Error() != nil {
			t.Fatalf("unexpected err %v", w2.Error())
		}
		if atomic.LoadInt32(count) != 2 {
			t.Fatalf("pool does not make new wire")
		}
		pool.Store(w1)
		pool.Close()
		if w1.Error() != ErrClosing {
			t.Fatalf("pool does not close existing wire after Close()")
		}
		for i := 0; i < 100; i++ {
			if rw := pool.Acquire(); rw != dead {
				t.Fatalf("pool does not return the dead wire after Close()")
			}
		}
		pool.Store(w2)
		if w2.Error() != ErrClosing {
			t.Fatalf("pool does not close stored wire after Close()")
		}
	})

	t.Run("Close Empty", func(t *testing.T) {
		pool, count := setup(2)
		w1 := pool.Acquire()
		if w1.Error() != nil {
			t.Fatalf("unexpected err %v", w1.Error())
		}
		pool.Close()
		w2 := pool.Acquire()
		if w2.Error() != ErrClosing {
			t.Fatalf("pool does not close wire after Close()")
		}
		if atomic.LoadInt32(count) != 1 {
			t.Fatalf("pool should not make new wire")
		}
		for i := 0; i < 100; i++ {
			if rw := pool.Acquire(); rw != dead {
				t.Fatalf("pool does not return the dead wire after Close()")
			}
		}
		pool.Store(w1)
		if w1.Error() != ErrClosing {
			t.Fatalf("pool does not close existing wire after Close()")
		}
	})

	t.Run("Close Waiting", func(t *testing.T) {
		pool, count := setup(1)
		w1 := pool.Acquire()
		if w1.Error() != nil {
			t.Fatalf("unexpected err %v", w1.Error())
		}
		pending := int64(0)
		for i := 0; i < 100; i++ {
			go func() {
				atomic.AddInt64(&pending, 1)
				if rw := pool.Acquire(); rw != dead {
					t.Errorf("pool does not return the dead wire after Close()")
				}
				atomic.AddInt64(&pending, -1)
			}()
		}
		for atomic.LoadInt64(&pending) != 100 {
			runtime.Gosched()
		}
		if atomic.LoadInt32(count) != 1 {
			t.Fatalf("pool should not make new wire")
		}
		pool.Close()
		for atomic.LoadInt64(&pending) != 0 {
			runtime.Gosched()
		}
		pool.Store(w1)
		if w1.Error() != ErrClosing {
			t.Fatalf("pool does not close existing wire after Close()")
		}
	})
}

func TestPoolError(t *testing.T) {
	setup := func(size int) (*pool, *int32) {
		var count int32
		return newPool(size, dead, func() wire {
			w := &pipe{}
			w.pshks.Store(emptypshks)
			c := atomic.AddInt32(&count, 1)
			if c%2 == 0 {
				w.error.Store(&errs{error: errors.New("any")})
			}
			return w
		}), &count
	}

	t.Run("NotStoreErrConn", func(t *testing.T) {
		conn := make([]wire, 100)
		pool, count := setup(len(conn))
		for i := 0; i < len(conn); i++ {
			conn[i] = pool.Acquire()
		}
		if atomic.LoadInt32(count) != int32(len(conn)) {
			t.Fatalf("unexpected acquire count")
		}
		for i := 0; i < len(conn); i++ {
			pool.Store(conn[i])
		}
		for i := 0; i < len(conn); i++ {
			conn[i] = pool.Acquire()
		}
		if atomic.LoadInt32(count) != int32(len(conn)+len(conn)/2) {
			t.Fatalf("unexpected acquire count")
		}
	})
}
