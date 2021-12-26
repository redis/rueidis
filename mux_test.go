package rueidis

import (
	"bufio"
	"errors"
	"net"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/rueian/rueidis/internal/cmds"
	"github.com/rueian/rueidis/internal/mock"
	"github.com/rueian/rueidis/internal/proto"
)

func setupMux(wires []*mock.Wire) (conn *mux, checkClean func(t *testing.T)) {
	var mu sync.Mutex
	var count = -1
	return newMux("", ConnOption{}, (*mock.Wire)(nil), func(fn func(err error)) (wire, error) {
			mu.Lock()
			defer mu.Unlock()
			count++
			return wires[count], nil
		}), func(t *testing.T) {
			if count != len(wires)-1 {
				t.Fatalf("there is %d remaining unused wires", len(wires)-count-1)
			}
		}
}

func TestNewMux(t *testing.T) {
	n1, n2 := net.Pipe()
	mock := &redisMock{t: t, buf: bufio.NewReader(n2), conn: n2}
	go func() {
		mock.Expect("HELLO", "3").
			Reply(proto.Message{
				Type: '%',
				Values: []proto.Message{
					{Type: '+', String: "key"},
					{Type: '+', String: "value"},
				},
			})
		mock.Expect("CLIENT", "TRACKING", "ON", "OPTIN").
			ReplyString("OK")
		mock.Expect("QUIT").ReplyString("OK")
	}()
	m := makeMux("", ConnOption{}, func(dst string, opt ConnOption) (net.Conn, error) {
		return n1, nil
	})
	if err := m.Dial(); err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	m.Close()
}

func TestMuxOnDisconnected(t *testing.T) {
	var trigger func(err error)
	m := newMux("", ConnOption{}, (*mock.Wire)(nil), func(fn func(err error)) (wire, error) {
		trigger = fn
		return &mock.Wire{}, nil
	})
	if err := m.Dial(); err != nil {
		t.Fatalf("unexpected err %v", err)
	}

	count := int64(0)

	trigger(errors.New("should have no effect before registering callback"))

	e := errors.New("should trigger")
	m.OnDisconnected(func(err error) {
		if err != e {
			t.Errorf("unexpected error %v", err)
		}
		atomic.AddInt64(&count, 1)
	})

	trigger(e)

	if atomic.LoadInt64(&count) != 1 {
		t.Fatalf("unxpected callback call count %d", atomic.LoadInt64(&count))
	}
}

func TestMuxDialSuppress(t *testing.T) {
	var wires, waits, done int64
	blocking := make(chan struct{})
	m := newMux("", ConnOption{}, (*mock.Wire)(nil), func(fn func(err error)) (wire, error) {
		atomic.AddInt64(&wires, 1)
		<-blocking
		return &mock.Wire{}, nil
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

func TestMuxReuseWire(t *testing.T) {
	t.Run("reuse wire if no error", func(t *testing.T) {
		m, checkClean := setupMux([]*mock.Wire{
			{
				DoFn: func(cmd cmds.Completed) proto.Result {
					return proto.NewResult(proto.Message{Type: '+', String: "PONG"}, nil)
				},
			},
		})
		defer checkClean(t)
		defer m.Close()
		for i := 0; i < 2; i++ {
			if err := m.Do(cmds.NewCompleted([]string{"PING"})).Error(); err != nil {
				t.Fatalf("unexpected error %v", err)
			}
		}
		m.Close()
	})
	t.Run("reuse blocking pool", func(t *testing.T) {
		blocking := make(chan struct{})
		response := make(chan proto.Result)
		m, checkClean := setupMux([]*mock.Wire{
			{
				// leave first wire for pipeline calls
			},
			{
				DoFn: func(cmd cmds.Completed) proto.Result {
					return proto.NewResult(proto.Message{Type: '+', String: "ACQUIRED"}, nil)
				},
			},
			{
				DoFn: func(cmd cmds.Completed) proto.Result {
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
			if val, err := m.Do(cmds.NewBlockingCompleted([]string{"PING"})).ToString(); err != nil {
				t.Errorf("unexpected error %v", err)
			} else if val != "BLOCK_RESPONSE" {
				t.Errorf("unexpected response %v", val)
			}
			close(blocking)
		}()
		<-blocking

		m.Store(wire1)
		// this should use the first wire
		if val, err := m.Do(cmds.NewBlockingCompleted([]string{"PING"})).ToString(); err != nil {
			t.Fatalf("unexpected error %v", err)
		} else if val != "ACQUIRED" {
			t.Fatalf("unexpected response %v", val)
		}

		response <- proto.NewResult(proto.Message{Type: '+', String: "BLOCK_RESPONSE"}, nil)
		<-blocking
	})
}

func TestMuxCMDRetry(t *testing.T) {
	t.Run("wire info", func(t *testing.T) {
		m, checkClean := setupMux([]*mock.Wire{
			{
				InfoFn: func() map[string]proto.Message {
					return map[string]proto.Message{"key": {Type: '+', String: "value"}}
				},
			},
		})
		defer checkClean(t)
		defer m.Close()
		if info := m.Info(); info == nil || info["key"].String != "value" {
			t.Fatalf("unexpected info %v", info)
		}
	})

	t.Run("wire err", func(t *testing.T) {
		e := errors.New("err")
		m, checkClean := setupMux([]*mock.Wire{
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

	t.Run("retry single read", func(t *testing.T) {
		m, checkClean := setupMux([]*mock.Wire{
			{
				DoFn: func(cmd cmds.Completed) proto.Result {
					if cmd.Commands()[0] == "PING" {
						return proto.NewResult(proto.Message{Type: '+', String: "PONG"}, nil)
					}
					if cmd.Commands()[0] != "READONLY_COMMAND" {
						t.Fatalf("command should be READONLY_COMMAND")
					}
					return proto.NewErrResult(errors.New("network error"))
				},
			},
			{
				DoFn: func(cmd cmds.Completed) proto.Result {
					if cmd.Commands()[0] == "PING" {
						return proto.NewResult(proto.Message{Type: '+', String: "PONG"}, nil)
					}
					if cmd.Commands()[0] != "READONLY_COMMAND" {
						t.Fatalf("command should be READONLY_COMMAND")
					}
					return proto.NewResult(proto.Message{Type: '+', String: "READONLY_COMMAND_RESPONSE"}, nil)
				},
			},
		})
		defer checkClean(t)
		defer m.Close()
		// this should automatically use the second wire
		if val, err := m.Do(cmds.NewReadOnlyCompleted([]string{"READONLY_COMMAND"})).ToString(); err != nil {
			t.Fatalf("unexpected error %v", err)
		} else if val != "READONLY_COMMAND_RESPONSE" {
			t.Fatalf("unexpected response %v", val)
		}
	})

	t.Run("retry multi read", func(t *testing.T) {
		m, checkClean := setupMux([]*mock.Wire{
			{
				DoMultiFn: func(multi ...cmds.Completed) []proto.Result {
					return []proto.Result{proto.NewErrResult(errors.New("network error"))}
				},
			},
			{
				DoMultiFn: func(multi ...cmds.Completed) []proto.Result {
					return []proto.Result{proto.NewResult(proto.Message{Type: '+', String: "MULTI_COMMANDS_RESPONSE"}, nil)}
				},
			},
		})
		defer checkClean(t)
		defer m.Close()
		// this should automatically use the second wire
		if val, err := m.DoMulti(cmds.NewReadOnlyCompleted([]string{"READONLY_COMMAND"}))[0].ToString(); err != nil {
			t.Fatalf("unexpected error %v", err)
		} else if val != "MULTI_COMMANDS_RESPONSE" {
			t.Fatalf("unexpected response %v", val)
		}
	})

	t.Run("retry single read cache", func(t *testing.T) {
		m, checkClean := setupMux([]*mock.Wire{
			{
				DoCacheFn: func(cmd cmds.Cacheable, ttl time.Duration) proto.Result {
					return proto.NewErrResult(errors.New("network error"))
				},
			},
			{
				DoCacheFn: func(cmd cmds.Cacheable, ttl time.Duration) proto.Result {
					return proto.NewResult(proto.Message{Type: '+', String: "READONLY_COMMAND_RESPONSE"}, nil)
				},
			},
		})
		defer checkClean(t)
		defer m.Close()
		// this should automatically use the second wire
		if val, err := m.DoCache(cmds.Cacheable(cmds.NewReadOnlyCompleted([]string{"READONLY_COMMAND"})), time.Second).ToString(); err != nil {
			t.Fatalf("unexpected error %v", err)
		} else if val != "READONLY_COMMAND_RESPONSE" {
			t.Fatalf("unexpected response %v", val)
		}
	})

	t.Run("not retry single write", func(t *testing.T) {
		m, checkClean := setupMux([]*mock.Wire{
			{
				DoFn: func(cmd cmds.Completed) proto.Result {
					if cmd.Commands()[0] == "PING" {
						return proto.NewResult(proto.Message{Type: '+', String: "PONG"}, nil)
					}
					if cmd.Commands()[0] != "WRITE_COMMAND" {
						t.Fatalf("command should be WRITE_COMMAND")
					}
					return proto.NewErrResult(errors.New("network error"))
				},
			},
			{
				DoFn: func(cmd cmds.Completed) proto.Result {
					if cmd.Commands()[0] == "PING" {
						return proto.NewResult(proto.Message{Type: '+', String: "PONG"}, nil)
					}
					if cmd.Commands()[0] != "WRITE_COMMAND" {
						t.Fatalf("command should be WRITE_COMMAND")
					}
					return proto.NewResult(proto.Message{Type: '+', String: "WRITE_COMMAND_RESPONSE"}, nil)
				},
			},
		})
		defer checkClean(t)
		defer m.Close()
		// this should only use the first wire
		if _, err := m.Do(cmds.NewCompleted([]string{"WRITE_COMMAND"})).ToString(); err == nil || err.Error() != "network error" {
			t.Fatalf("unexpected error %v", err)
		}
		// this should use the second wire
		if val, err := m.Do(cmds.NewCompleted([]string{"WRITE_COMMAND"})).ToString(); err != nil {
			t.Fatalf("unexpected error %v", err)
		} else if val != "WRITE_COMMAND_RESPONSE" {
			t.Fatalf("unexpected response %v", val)
		}
	})

	t.Run("not retry multi write", func(t *testing.T) {
		m, checkClean := setupMux([]*mock.Wire{
			{
				DoMultiFn: func(multi ...cmds.Completed) []proto.Result {
					return []proto.Result{proto.NewErrResult(errors.New("network error"))}
				},
			},
			{
				DoMultiFn: func(multi ...cmds.Completed) []proto.Result {
					return []proto.Result{proto.NewResult(proto.Message{Type: '+', String: "MULTI_COMMANDS_RESPONSE"}, nil)}
				},
			},
		})
		defer checkClean(t)
		defer m.Close()
		// this should only use the first wire
		if _, err := m.DoMulti(
			cmds.NewReadOnlyCompleted([]string{"READONLY_COMMAND"}),
			cmds.NewCompleted([]string{"WRITE_COMMAND"}),
		)[0].ToString(); err == nil || err.Error() != "network error" {
			t.Fatalf("unexpected error %v", err)
		}
		// this should use the second wire
		if val, err := m.DoMulti(
			cmds.NewReadOnlyCompleted([]string{"READONLY_COMMAND"}),
			cmds.NewCompleted([]string{"WRITE_COMMAND"}),
		)[0].ToString(); err != nil {
			t.Fatalf("unexpected error %v", err)
		} else if val != "MULTI_COMMANDS_RESPONSE" {
			t.Fatalf("unexpected response %v", val)
		}
	})

	t.Run("single blocking", func(t *testing.T) {
		blocked := make(chan struct{})
		responses := make(chan proto.Result)

		m, checkClean := setupMux([]*mock.Wire{
			{
				// leave first wire for pipeline calls
			},
			{
				DoFn: func(cmd cmds.Completed) proto.Result {
					blocked <- struct{}{}
					return <-responses
				},
			},
			{
				DoFn: func(cmd cmds.Completed) proto.Result {
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
				if val, err := m.Do(cmds.NewBlockingCompleted([]string{"BLOCK"})).ToString(); err != nil {
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
			responses <- proto.NewResult(proto.Message{Type: '+', String: "BLOCK_COMMANDS_RESPONSE"}, nil)
		}
		wg.Wait()
	})

	t.Run("multiple blocking", func(t *testing.T) {
		blocked := make(chan struct{})
		responses := make(chan proto.Result)

		m, checkClean := setupMux([]*mock.Wire{
			{
				// leave first wire for pipeline calls
			},
			{
				DoMultiFn: func(cmd ...cmds.Completed) []proto.Result {
					blocked <- struct{}{}
					return []proto.Result{<-responses}
				},
			},
			{
				DoMultiFn: func(cmd ...cmds.Completed) []proto.Result {
					blocked <- struct{}{}
					return []proto.Result{<-responses}
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
			responses <- proto.NewResult(proto.Message{Type: '+', String: "BLOCK_COMMANDS_RESPONSE"}, nil)
		}
		wg.Wait()
	})
}

func TestMuxDialRetry(t *testing.T) {
	setup := func() (*mux, *int64) {
		var count int64
		return newMux("", ConnOption{}, (*mock.Wire)(nil), func(fn func(err error)) (wire, error) {
			if count == 1 {
				return &mock.Wire{
					DoFn: func(cmd cmds.Completed) proto.Result {
						return proto.NewResult(proto.Message{Type: '+', String: "PONG"}, nil)
					},
				}, nil
			}
			count++
			return nil, errors.New("network err")
		}), &count
	}
	t.Run("retry on auto pipeline", func(t *testing.T) {
		m, count := setup()
		defer m.Close()
		if val, err := m.Do(cmds.NewCompleted([]string{"PING"})).ToString(); err != nil {
			t.Fatalf("unexpected err %v", err)
		} else if val != "PONG" {
			t.Fatalf("unexpected response %v", val)
		}
		if *count != 1 {
			t.Fatalf("no dial err at all")
		}
	})

	t.Run("retry on blocking pool", func(t *testing.T) {
		m, count := setup()
		defer m.Close()
		if val, err := m.Do(cmds.NewBlockingCompleted([]string{"PING"})).ToString(); err != nil {
			t.Fatalf("unexpected err %v", err)
		} else if val != "PONG" {
			t.Fatalf("unexpected response %v", val)
		}
		if *count != 1 {
			t.Fatalf("no dial err at all")
		}
	})
}

func BenchmarkClientSideCaching(b *testing.B) {
	setup := func(b *testing.B) *mux {
		c := makeMux("127.0.0.1:6379", ConnOption{CacheSizeEachConn: DefaultCacheBytes}, func(dst string, opt ConnOption) (conn net.Conn, err error) {
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
				m.Do(cmd)
			}
		})
	})
	b.Run("DoCache", func(b *testing.B) {
		m := setup(b)
		cmd := cmds.Cacheable(cmds.NewCompleted([]string{"GET", "a"}))
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				m.DoCache(cmd, time.Second*5)
			}
		})
	})
}
