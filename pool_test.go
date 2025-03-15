package rueidis

import (
	"context"
	"errors"
	"runtime"
	"sync/atomic"
	"testing"
	"time"
)

var dead = deadFn()

//gocyclo:ignore
func TestPool(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	setup := func(size int) (*pool, *int32) {
		var count int32
		return newPool(size, dead, 0, 0, func(_ context.Context) wire {
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
		p := newPool(0, dead, 0, 0, func(_ context.Context) wire { return nil })
		if cap(p.list) == 0 {
			t.Fatalf("DefaultPoolSize is not applied")
		}
	})

	t.Run("Reuse", func(t *testing.T) {
		pool, count := setup(100)
		for i := 0; i < 1000; i++ {
			pool.Store(pool.Acquire(context.Background()))
		}
		if atomic.LoadInt32(count) != 1 {
			t.Fatalf("pool does not reuse connection")
		}
	})

	t.Run("Reuse without broken connections", func(t *testing.T) {
		pool, count := setup(100)
		c1 := pool.Acquire(context.Background())
		c2 := pool.Acquire(context.Background())
		pool.Store(c1)
		pool.Store(c2)
		pool.cond.L.Lock()
		for _, p := range pool.list {
			p.Close()
		}
		pool.cond.L.Unlock()
		c3 := pool.Acquire(context.Background())
		if c3.Error() != nil {
			t.Fatalf("c3.Error() is not nil")
		}
		if atomic.LoadInt32(count) != 3 {
			t.Fatalf("pool does not clean borken connections")
		}
		pool.cond.L.Lock()
		defer pool.cond.L.Unlock()
		if pool.size != 1 {
			t.Fatalf("pool size is not 1")
		}
		if len(pool.list) != 0 {
			t.Fatalf("pool list is not empty")
		}
	})

	t.Run("NotExceed", func(t *testing.T) {
		conn := make([]wire, 100)
		pool, count := setup(len(conn))
		for i := 0; i < len(conn); i++ {
			conn[i] = pool.Acquire(context.Background())
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
			pool.Acquire(context.Background())
		}
		if atomic.LoadInt32(count) > 100 {
			t.Fatalf("pool must not exceed the size limit")
		}
	})

	t.Run("NoShare", func(t *testing.T) {
		conn := make([]wire, 100)
		pool, _ := setup(len(conn))
		for i := 0; i < len(conn); i++ {
			w := pool.Acquire(context.Background())
			go pool.Store(w)
		}
		for i := 0; i < len(conn); i++ {
			conn[i] = pool.Acquire(context.Background())
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
		w1 := pool.Acquire(context.Background())
		w2 := pool.Acquire(context.Background())
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
			if rw := pool.Acquire(context.Background()); rw != dead {
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
		w1 := pool.Acquire(context.Background())
		if w1.Error() != nil {
			t.Fatalf("unexpected err %v", w1.Error())
		}
		pool.Close()
		w2 := pool.Acquire(context.Background())
		if w2.Error() != ErrClosing {
			t.Fatalf("pool does not close wire after Close()")
		}
		if atomic.LoadInt32(count) != 1 {
			t.Fatalf("pool should not make new wire")
		}
		for i := 0; i < 100; i++ {
			if rw := pool.Acquire(context.Background()); rw != dead {
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
		w1 := pool.Acquire(context.Background())
		if w1.Error() != nil {
			t.Fatalf("unexpected err %v", w1.Error())
		}
		pending := int64(0)
		for i := 0; i < 100; i++ {
			go func() {
				atomic.AddInt64(&pending, 1)
				if rw := pool.Acquire(context.Background()); rw != dead {
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
	defer ShouldNotLeaked(SetupLeakDetection())
	setup := func(size int) (*pool, *int32) {
		var count int32
		return newPool(size, dead, 0, 0, func(_ context.Context) wire {
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
			conn[i] = pool.Acquire(context.Background())
		}
		if atomic.LoadInt32(count) != int32(len(conn)) {
			t.Fatalf("unexpected acquire count")
		}
		for i := 0; i < len(conn); i++ {
			pool.Store(conn[i])
		}
		for i := 0; i < len(conn); i++ {
			conn[i] = pool.Acquire(context.Background())
		}
		if atomic.LoadInt32(count) != int32(len(conn)+len(conn)/2) {
			t.Fatalf("unexpected acquire count")
		}
	})
}

func TestPoolWithIdleTTL(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	setup := func(size int, ttl time.Duration, minSize int) *pool {
		return newPool(size, dead, ttl, minSize, func(_ context.Context) wire {
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
		})
	}

	t.Run("Removing idle conns. Min size is not 0", func(t *testing.T) {
		minSize := 3
		p := setup(0, time.Millisecond*50, minSize)
		conns := make([]wire, 10)

		for i := 0; i < 2; i++ {
			for i := range conns {
				w := p.Acquire(context.Background())
				conns[i] = w
			}

			for _, w := range conns {
				p.Store(w)
			}

			time.Sleep(time.Millisecond * 60)
			p.cond.Broadcast()
			time.Sleep(time.Millisecond * 40)

			p.cond.L.Lock()
			if p.size != minSize {
				defer p.cond.L.Unlock()
				t.Fatalf("size must be equal to %d, actual: %d", minSize, p.size)
			}

			if len(p.list) != minSize {
				defer p.cond.L.Unlock()
				t.Fatalf("pool len must equal to %d, actual: %d", minSize, len(p.list))
			}
			p.cond.L.Unlock()
		}

		p.Close()
	})

	t.Run("Removing idle conns. Min size is 0", func(t *testing.T) {
		p := setup(0, time.Millisecond*50, 0)
		conns := make([]wire, 10)

		for i := 0; i < 2; i++ {
			for i := range conns {
				w := p.Acquire(context.Background())
				conns[i] = w
			}

			for _, w := range conns {
				p.Store(w)
			}

			time.Sleep(time.Millisecond * 60)
			p.cond.Broadcast()
			time.Sleep(time.Millisecond * 40)

			p.cond.L.Lock()
			if p.size != 0 {
				defer p.cond.L.Unlock()
				t.Fatalf("size must be equal to 0, actual: %d", p.size)
			}

			if len(p.list) != 0 {
				defer p.cond.L.Unlock()
				t.Fatalf("pool len must equal to 0, actual: %d", len(p.list))
			}
			p.cond.L.Unlock()
		}

		p.Close()
	})
}

func TestPoolWithAcquireCtx(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	setup := func(size int, delay time.Duration) *pool {
		return newPool(size, dead, 0, 0, func(ctx context.Context) wire {
			var err error 
			closed := false
			timer := time.NewTimer(delay)
			defer timer.Stop()
			select {
			case <-ctx.Done():
				err = ctx.Err()
				closed = true
			case <-timer.C:
				// noop
			}
			
			return &mockWire{
				CloseFn: func() {
					closed = true
				},
				ErrorFn: func() error {
					if err != nil {
						return err
					} else if closed {
						return ErrClosing
					}
					return nil
				},
			}
		})
	}
	t.Run("Acquire connections, all exceed context deadline", func(t *testing.T) {
		p := setup(10, time.Millisecond*5)
		conns := make([]wire, 10)

		for i := 0; i < 2; i++ {
			for i := range conns {
				ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
				w := p.Acquire(ctx)
				conns[i] = w
				cancel()
			}

			for _, w := range conns {
				p.Store(w)
			}

			p.cond.L.Lock()
			if p.size != 0 {
				defer p.cond.L.Unlock()
				t.Fatalf("size must be equal to 0, actual: %d", p.size)
			}

			if len(p.list) != 0 {
				defer p.cond.L.Unlock()
				t.Fatalf("pool len must equal to 0, actual: %d", len(p.list))
			}
			p.cond.L.Unlock()
		}

		p.Close()
	})

	t.Run("Acquire connections, some exceed context deadline", func(t *testing.T) {
		p := setup(10, time.Millisecond*5)
		conns := make([]wire, 10)

		// size = 5
		for i := range conns {
			d := time.Millisecond
			if i % 2 == 0 {
				d = time.Millisecond * 8
			}
			ctx, cancel := context.WithTimeout(context.Background(), d)
			w := p.Acquire(ctx)
			conns[i] = w
			cancel()
		}
		for _, w := range conns {
			p.Store(w)
		}
		p.cond.L.Lock()
		if p.size != len(conns)/2 {
			defer p.cond.L.Unlock()
			t.Fatalf("size must be equal to %d, actual: %d", len(conns)/2, p.size)
		}

		if len(p.list) != len(conns)/2 {
			defer p.cond.L.Unlock()
			t.Fatalf("pool len must equal to %d, actual: %d", len(conns)/2, len(p.list))
		}
		p.cond.L.Unlock()

		// size = 10
		for i := range conns {
			ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond * 8)
			w := p.Acquire(ctx)
			conns[i] = w
			cancel()
		}
		for _, w := range conns {
			p.Store(w)
		}
		p.cond.L.Lock()
		if p.size != len(conns) {
			defer p.cond.L.Unlock()
			t.Fatalf("size must be equal to %d, actual: %d", len(conns), p.size)
		}

		if len(p.list) != len(conns) {
			defer p.cond.L.Unlock()
			t.Fatalf("pool len must equal to %d, actual: %d", len(conns), len(p.list))
		}
		p.cond.L.Unlock()
		
		p.Close()
	})
}