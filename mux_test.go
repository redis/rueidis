package rueidis

import (
	"bufio"
	"context"
	"errors"
	"net"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/rueian/rueidis/internal/cmds"
)

func setupMux(wires []*mockWire) (conn *mux, checkClean func(t *testing.T)) {
	var mu sync.Mutex
	var count = -1
	return newMux("", &ClientOption{}, (*mockWire)(nil), (*mockWire)(nil), func() wire {
			mu.Lock()
			defer mu.Unlock()
			count++
			return wires[count]
		}), func(t *testing.T) {
			if count != len(wires)-1 {
				t.Fatalf("there is %d remaining unused wires", len(wires)-count-1)
			}
		}
}

func TestNewMuxDailErr(t *testing.T) {
	c := 0
	e := errors.New("any")
	m := makeMux("", &ClientOption{}, func(dst string, opt *ClientOption) (net.Conn, error) {
		c++
		return nil, e
	})
	if err := m.Dial(); err != e {
		t.Fatalf("unexpected return %v", err)
	}
	if c != 1 {
		t.Fatalf("dialFn not called")
	}
	if w := m.pipe(); w != m.dead { // c = 2
		t.Fatalf("unexpected wire %v", w)
	}
	if err := m.Dial(); err != e { // c = 3
		t.Fatalf("unexpected return %v", err)
	}
	if w := m.Acquire(); w != m.dead {
		t.Fatalf("unexpected wire %v", w)
	}
	if c != 4 {
		t.Fatalf("dialFn not called %v", c)
	}
}

func TestNewMux(t *testing.T) {
	n1, n2 := net.Pipe()
	mock := &redisMock{t: t, buf: bufio.NewReader(n2), conn: n2}
	go func() {
		mock.Expect("HELLO", "3").
			Reply(RedisMessage{
				typ: '%',
				values: []RedisMessage{
					{typ: '+', string: "key"},
					{typ: '+', string: "value"},
				},
			})
		mock.Expect("CLIENT", "TRACKING", "ON", "OPTIN").
			ReplyString("OK")
		mock.Expect("QUIT").ReplyString("OK")
	}()
	m := makeMux("", &ClientOption{}, func(dst string, opt *ClientOption) (net.Conn, error) {
		return n1, nil
	})
	if err := m.Dial(); err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	t.Run("Override with previous mux", func(t *testing.T) {
		m2 := makeMux("", &ClientOption{}, func(dst string, opt *ClientOption) (net.Conn, error) {
			return n1, nil
		})
		m2.Override(m)
		if err := m2.Dial(); err != nil {
			t.Fatalf("unexpected error %v", err)
		}
		m2.Close()
	})
}

func TestMuxIs(t *testing.T) {
	m := makeMux("dst1", &ClientOption{}, nil)
	if m.Is("dst1") != true {
		t.Fatalf("unexpected m.Is == false")
	}
	if m.Is("") == true {
		t.Fatalf("unexpected m.Is == true")
	}
}

func TestMuxDialSuppress(t *testing.T) {
	var wires, waits, done int64
	blocking := make(chan struct{})
	m := newMux("", &ClientOption{}, (*mockWire)(nil), (*mockWire)(nil), func() wire {
		atomic.AddInt64(&wires, 1)
		<-blocking
		return &mockWire{}
	})
	for i := 0; i < 1000; i++ {
		go func() {
			atomic.AddInt64(&waits, 1)
			m.Info()
			atomic.AddInt64(&done, 1)
		}()
	}
	for atomic.LoadInt64(&waits) != 1000 {
		runtime.Gosched()
	}
	close(blocking)
	for atomic.LoadInt64(&done) != 1000 {
		runtime.Gosched()
	}
	if atomic.LoadInt64(&wires) != 1 {
		t.Fatalf("wireFn is not suppressed")
	}
}

//gocyclo:ignore
func TestMuxReuseWire(t *testing.T) {
	t.Run("reuse wire if no error", func(t *testing.T) {
		m, checkClean := setupMux([]*mockWire{
			{
				DoFn: func(cmd cmds.Completed) RedisResult {
					return newResult(RedisMessage{typ: '+', string: "PONG"}, nil)
				},
			},
		})
		defer checkClean(t)
		defer m.Close()
		for i := 0; i < 2; i++ {
			if err := m.Do(context.Background(), cmds.NewCompleted([]string{"PING"})).Error(); err != nil {
				t.Fatalf("unexpected error %v", err)
			}
		}
		m.Close()
	})

	t.Run("reuse blocking pool", func(t *testing.T) {
		blocking := make(chan struct{})
		response := make(chan RedisResult)
		m, checkClean := setupMux([]*mockWire{
			{
				// leave first wire for pipeline calls
			},
			{
				DoFn: func(cmd cmds.Completed) RedisResult {
					return newResult(RedisMessage{typ: '+', string: "ACQUIRED"}, nil)
				},
			},
			{
				DoFn: func(cmd cmds.Completed) RedisResult {
					blocking <- struct{}{}
					return <-response
				},
			},
		})
		defer checkClean(t)
		defer m.Close()
		if err := m.Dial(); err != nil {
			t.Fatalf("unexpected dial error %v", err)
		}

		wire1 := m.Acquire()

		go func() {
			// this should use the second wire
			if val, err := m.Do(context.Background(), cmds.NewBlockingCompleted([]string{"PING"})).ToString(); err != nil {
				t.Errorf("unexpected error %v", err)
			} else if val != "BLOCK_RESPONSE" {
				t.Errorf("unexpected response %v", val)
			}
			close(blocking)
		}()
		<-blocking

		m.Store(wire1)
		// this should use the first wire
		if val, err := m.Do(context.Background(), cmds.NewBlockingCompleted([]string{"PING"})).ToString(); err != nil {
			t.Fatalf("unexpected error %v", err)
		} else if val != "ACQUIRED" {
			t.Fatalf("unexpected response %v", val)
		}

		response <- newResult(RedisMessage{typ: '+', string: "BLOCK_RESPONSE"}, nil)
		<-blocking
	})

	t.Run("unsubscribe blocking pool", func(t *testing.T) {
		cleaned := false
		m, checkClean := setupMux([]*mockWire{
			{
				// leave first wire for pipeline calls
			},
			{
				CleanSubscriptionsFn: func() {
					cleaned = true
				},
			},
		})
		defer checkClean(t)
		defer m.Close()

		if err := m.Dial(); err != nil {
			t.Fatalf("unexpected dial error %v", err)
		}

		wire1 := m.Acquire()
		m.Store(wire1)

		if !cleaned {
			t.Fatalf("CleanSubscriptions not called")
		}
	})
}

//gocyclo:ignore
func TestMuxDelegation(t *testing.T) {
	t.Run("wire info", func(t *testing.T) {
		m, checkClean := setupMux([]*mockWire{
			{
				InfoFn: func() map[string]RedisMessage {
					return map[string]RedisMessage{"key": {typ: '+', string: "value"}}
				},
			},
		})
		defer checkClean(t)
		defer m.Close()
		if info := m.Info(); info == nil || info["key"].string != "value" {
			t.Fatalf("unexpected info %v", info)
		}
	})

	t.Run("wire err", func(t *testing.T) {
		e := errors.New("err")
		m, checkClean := setupMux([]*mockWire{
			{
				ErrorFn: func() error {
					return e
				},
			},
		})
		defer checkClean(t)
		defer m.Close()
		if err := m.Error(); err != e {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("wire do", func(t *testing.T) {
		m, checkClean := setupMux([]*mockWire{
			{
				DoFn: func(cmd cmds.Completed) RedisResult {
					return newErrResult(context.DeadlineExceeded)
				},
				ErrorFn: func() error {
					return context.DeadlineExceeded
				},
			},
			{
				DoFn: func(cmd cmds.Completed) RedisResult {
					if cmd.Commands()[0] != "READONLY_COMMAND" {
						t.Fatalf("command should be READONLY_COMMAND")
					}
					return newResult(RedisMessage{typ: '+', string: "READONLY_COMMAND_RESPONSE"}, nil)
				},
			},
		})
		defer checkClean(t)
		defer m.Close()
		if err := m.Do(context.Background(), cmds.NewReadOnlyCompleted([]string{"READONLY_COMMAND"})).Error(); !errors.Is(err, context.DeadlineExceeded) {
			t.Fatalf("unexpected error %v", err)
		}
		if val, err := m.Do(context.Background(), cmds.NewReadOnlyCompleted([]string{"READONLY_COMMAND"})).ToString(); err != nil {
			t.Fatalf("unexpected error %v", err)
		} else if val != "READONLY_COMMAND_RESPONSE" {
			t.Fatalf("unexpected response %v", val)
		}
	})

	t.Run("wire do multi", func(t *testing.T) {
		m, checkClean := setupMux([]*mockWire{
			{
				DoMultiFn: func(multi ...cmds.Completed) []RedisResult {
					return []RedisResult{newErrResult(context.DeadlineExceeded)}
				},
				ErrorFn: func() error {
					return context.DeadlineExceeded
				},
			},
			{
				DoMultiFn: func(multi ...cmds.Completed) []RedisResult {
					return []RedisResult{newResult(RedisMessage{typ: '+', string: "MULTI_COMMANDS_RESPONSE"}, nil)}
				},
			},
		})
		defer checkClean(t)
		defer m.Close()
		if err := m.DoMulti(context.Background(), cmds.NewReadOnlyCompleted([]string{"READONLY_COMMAND"}))[0].Error(); !errors.Is(err, context.DeadlineExceeded) {
			t.Fatalf("unexpected error %v", err)
		}
		if val, err := m.DoMulti(context.Background(), cmds.NewReadOnlyCompleted([]string{"READONLY_COMMAND"}))[0].ToString(); err != nil {
			t.Fatalf("unexpected error %v", err)
		} else if val != "MULTI_COMMANDS_RESPONSE" {
			t.Fatalf("unexpected response %v", val)
		}
	})

	t.Run("wire do cache", func(t *testing.T) {
		m, checkClean := setupMux([]*mockWire{
			{
				DoCacheFn: func(cmd cmds.Cacheable, ttl time.Duration) RedisResult {
					return newErrResult(context.DeadlineExceeded)
				},
				ErrorFn: func() error {
					return context.DeadlineExceeded
				},
			},
			{
				DoCacheFn: func(cmd cmds.Cacheable, ttl time.Duration) RedisResult {
					return newResult(RedisMessage{typ: '+', string: "READONLY_COMMAND_RESPONSE"}, nil)
				},
			},
		})
		defer checkClean(t)
		defer m.Close()
		if err := m.DoCache(context.Background(), cmds.Cacheable(cmds.NewReadOnlyCompleted([]string{"READONLY_COMMAND"})), time.Second).Error(); !errors.Is(err, context.DeadlineExceeded) {
			t.Fatalf("unexpected error %v", err)
		}
		if val, err := m.DoCache(context.Background(), cmds.Cacheable(cmds.NewReadOnlyCompleted([]string{"READONLY_COMMAND"})), time.Second).ToString(); err != nil {
			t.Fatalf("unexpected error %v", err)
		} else if val != "READONLY_COMMAND_RESPONSE" {
			t.Fatalf("unexpected response %v", val)
		}
	})

	t.Run("wire receive", func(t *testing.T) {
		m, checkClean := setupMux([]*mockWire{
			{
				ReceiveFn: func(ctx context.Context, subscribe cmds.Completed, fn func(message PubSubMessage)) error {
					return context.DeadlineExceeded
				},
				ErrorFn: func() error {
					return context.DeadlineExceeded
				},
			},
			{
				ReceiveFn: func(ctx context.Context, subscribe cmds.Completed, fn func(message PubSubMessage)) error {
					if subscribe.Commands()[0] != "SUBSCRIBE" {
						t.Fatalf("command should be SUBSCRIBE")
					}
					return nil
				},
			},
		})
		defer checkClean(t)
		defer m.Close()
		if err := m.Receive(context.Background(), cmds.NewCompleted([]string{"SUBSCRIBE"}), func(message PubSubMessage) {}); !errors.Is(err, context.DeadlineExceeded) {
			t.Fatalf("unexpected error %v", err)
		}
		if err := m.Receive(context.Background(), cmds.NewCompleted([]string{"SUBSCRIBE"}), func(message PubSubMessage) {}); err != nil {
			t.Fatalf("unexpected error %v", err)
		}
	})

	t.Run("single blocking", func(t *testing.T) {
		blocked := make(chan struct{})
		responses := make(chan RedisResult)

		m, checkClean := setupMux([]*mockWire{
			{
				// leave first wire for pipeline calls
			},
			{
				DoFn: func(cmd cmds.Completed) RedisResult {
					blocked <- struct{}{}
					return <-responses
				},
			},
			{
				DoFn: func(cmd cmds.Completed) RedisResult {
					blocked <- struct{}{}
					return <-responses
				},
			},
		})
		defer checkClean(t)
		defer m.Close()
		if err := m.Dial(); err != nil {
			t.Fatalf("unexpected dial error %v", err)
		}

		wg := sync.WaitGroup{}
		wg.Add(2)
		for i := 0; i < 2; i++ {
			go func() {
				if val, err := m.Do(context.Background(), cmds.NewBlockingCompleted([]string{"BLOCK"})).ToString(); err != nil {
					t.Errorf("unexpected error %v", err)
				} else if val != "BLOCK_COMMANDS_RESPONSE" {
					t.Errorf("unexpected response %v", val)
				} else {
					wg.Done()
				}
			}()
		}
		for i := 0; i < 2; i++ {
			<-blocked
		}
		for i := 0; i < 2; i++ {
			responses <- newResult(RedisMessage{typ: '+', string: "BLOCK_COMMANDS_RESPONSE"}, nil)
		}
		wg.Wait()
	})

	t.Run("multiple blocking", func(t *testing.T) {
		blocked := make(chan struct{})
		responses := make(chan RedisResult)

		m, checkClean := setupMux([]*mockWire{
			{
				// leave first wire for pipeline calls
			},
			{
				DoMultiFn: func(cmd ...cmds.Completed) []RedisResult {
					blocked <- struct{}{}
					return []RedisResult{<-responses}
				},
			},
			{
				DoMultiFn: func(cmd ...cmds.Completed) []RedisResult {
					blocked <- struct{}{}
					return []RedisResult{<-responses}
				},
			},
		})
		defer checkClean(t)
		defer m.Close()
		if err := m.Dial(); err != nil {
			t.Fatalf("unexpected dial error %v", err)
		}

		wg := sync.WaitGroup{}
		wg.Add(2)
		for i := 0; i < 2; i++ {
			go func() {
				if val, err := m.DoMulti(
					context.Background(),
					cmds.NewReadOnlyCompleted([]string{"READONLY"}),
					cmds.NewBlockingCompleted([]string{"BLOCK"}),
				)[0].ToString(); err != nil {
					t.Errorf("unexpected error %v", err)
				} else if val != "BLOCK_COMMANDS_RESPONSE" {
					t.Errorf("unexpected response %v", val)
				} else {
					wg.Done()
				}
			}()
		}
		for i := 0; i < 2; i++ {
			<-blocked
		}
		for i := 0; i < 2; i++ {
			responses <- newResult(RedisMessage{typ: '+', string: "BLOCK_COMMANDS_RESPONSE"}, nil)
		}
		wg.Wait()
	})
}

func BenchmarkClientSideCaching(b *testing.B) {
	setup := func(b *testing.B) *mux {
		c := makeMux("127.0.0.1:6379", &ClientOption{CacheSizeEachConn: DefaultCacheBytes}, func(dst string, opt *ClientOption) (conn net.Conn, err error) {
			return net.Dial("tcp", dst)
		})
		if err := c.Dial(); err != nil {
			panic(err)
		}
		b.SetParallelism(100)
		b.ResetTimer()
		return c
	}
	b.Run("Do", func(b *testing.B) {
		m := setup(b)
		cmd := cmds.NewCompleted([]string{"GET", "a"})
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				m.Do(context.Background(), cmd)
			}
		})
	})
	b.Run("DoCache", func(b *testing.B) {
		m := setup(b)
		cmd := cmds.Cacheable(cmds.NewCompleted([]string{"GET", "a"}))
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				m.DoCache(context.Background(), cmd, time.Second*5)
			}
		})
	})
}

type mockWire struct {
	DoFn      func(cmd cmds.Completed) RedisResult
	DoCacheFn func(cmd cmds.Cacheable, ttl time.Duration) RedisResult
	DoMultiFn func(multi ...cmds.Completed) []RedisResult
	ReceiveFn func(ctx context.Context, subscribe cmds.Completed, fn func(message PubSubMessage)) error
	InfoFn    func() map[string]RedisMessage
	ErrorFn   func() error
	CloseFn   func()

	CleanSubscriptionsFn func()
	SetPubSubHooksFn     func(hooks PubSubHooks) <-chan error
}

func (m *mockWire) Do(ctx context.Context, cmd cmds.Completed) RedisResult {
	if m.DoFn != nil {
		return m.DoFn(cmd)
	}
	return RedisResult{}
}

func (m *mockWire) DoCache(ctx context.Context, cmd cmds.Cacheable, ttl time.Duration) RedisResult {
	if m.DoCacheFn != nil {
		return m.DoCacheFn(cmd, ttl)
	}
	return RedisResult{}
}

func (m *mockWire) DoMulti(ctx context.Context, multi ...cmds.Completed) []RedisResult {
	if m.DoMultiFn != nil {
		return m.DoMultiFn(multi...)
	}
	return nil
}

func (m *mockWire) Receive(ctx context.Context, subscribe cmds.Completed, fn func(message PubSubMessage)) error {
	if m.ReceiveFn != nil {
		return m.ReceiveFn(ctx, subscribe, fn)
	}
	return nil
}

func (m *mockWire) CleanSubscriptions() {
	if m.CleanSubscriptionsFn != nil {
		m.CleanSubscriptionsFn()
	}
	return
}

func (m *mockWire) SetPubSubHooks(hooks PubSubHooks) <-chan error {
	if m.SetPubSubHooksFn != nil {
		return m.SetPubSubHooksFn(hooks)
	}
	return nil
}

func (m *mockWire) Info() map[string]RedisMessage {
	if m.InfoFn != nil {
		return m.InfoFn()
	}
	return nil
}

func (m *mockWire) Error() error {
	if m == nil {
		return ErrClosing
	}
	if m.ErrorFn != nil {
		return m.ErrorFn()
	}
	return nil
}

func (m *mockWire) Close() {
	if m == nil {
		return
	}
	if m.CloseFn != nil {
		m.CloseFn()
	}
}
