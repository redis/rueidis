package conn

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
	"github.com/rueian/rueidis/internal/proto"
)

type mockWire struct {
	do      func(cmd cmds.Completed) proto.Result
	doCache func(cmd cmds.Cacheable, ttl time.Duration) proto.Result
	doMulti func(multi ...cmds.Completed) []proto.Result
	info    func() map[string]proto.Message
	error   func() error
	close   func()
}

func (m *mockWire) Do(cmd cmds.Completed) proto.Result {
	if m.do != nil {
		return m.do(cmd)
	}
	return proto.Result{}
}

func (m *mockWire) DoCache(cmd cmds.Cacheable, ttl time.Duration) proto.Result {
	if m.doCache != nil {
		return m.doCache(cmd, ttl)
	}
	return proto.Result{}
}

func (m *mockWire) DoMulti(multi ...cmds.Completed) []proto.Result {
	if m.doMulti != nil {
		return m.doMulti(multi...)
	}
	return nil
}

func (m *mockWire) Info() map[string]proto.Message {
	if m.info != nil {
		return m.info()
	}
	return nil
}

func (m *mockWire) Error() error {
	if m.error != nil {
		return m.error()
	}
	return nil
}

func (m *mockWire) Close() {
	if m.close != nil {
		m.close()
	}
}

func setupConn(wires []*mockWire) (conn *Conn, checkClean func(t *testing.T)) {
	var mu sync.Mutex
	var count = -1
	return newConn("", Option{}, (*mockWire)(nil), func(dst string, opt Option) (net.Conn, error) {
			return nil, nil
		}, func(conn net.Conn, opt Option) (Wire, error) {
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

func TestNewConn(t *testing.T) {
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
	conn := NewConn("", Option{}, func(dst string, opt Option) (net.Conn, error) {
		return n1, nil
	})
	if err := conn.Dialable(); err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	conn.Close()
}

func TestConnDialSuppress(t *testing.T) {
	var dials, wires, waits, done int64
	blocking := make(chan struct{})
	conn := newConn("", Option{}, (*mockWire)(nil), func(dst string, opt Option) (net.Conn, error) {
		atomic.AddInt64(&dials, 1)
		return nil, nil
	}, func(conn net.Conn, opt Option) (Wire, error) {
		atomic.AddInt64(&wires, 1)
		<-blocking
		return &mockWire{}, nil
	})
	for i := 0; i < 1000; i++ {
		go func() {
			atomic.AddInt64(&waits, 1)
			conn.Info()
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
	if atomic.LoadInt64(&dials) != 1 {
		t.Fatalf("dailFn is not suppressed")
	}
	if atomic.LoadInt64(&wires) != 1 {
		t.Fatalf("wireFn is not suppressed")
	}
}

func TestConnReuseWire(t *testing.T) {
	t.Run("reuse wire if no error", func(t *testing.T) {
		conn, checkClean := setupConn([]*mockWire{
			{
				do: func(cmd cmds.Completed) proto.Result {
					return proto.NewResult(proto.Message{Type: '+', String: "PONG"}, nil)
				},
			},
		})
		defer checkClean(t)
		defer conn.Close()
		for i := 0; i < 2; i++ {
			if err := conn.Do(cmds.NewCompleted([]string{"PING"})).Error(); err != nil {
				t.Fatalf("unexpected error %v", err)
			}
		}
		conn.Close()
	})
	t.Run("reuse blocking pool", func(t *testing.T) {
		blocking := make(chan struct{})
		response := make(chan proto.Result)
		conn, checkClean := setupConn([]*mockWire{
			{
				// leave first wire for pipeline calls
			},
			{
				do: func(cmd cmds.Completed) proto.Result {
					return proto.NewResult(proto.Message{Type: '+', String: "ACQUIRED"}, nil)
				},
			},
			{
				do: func(cmd cmds.Completed) proto.Result {
					blocking <- struct{}{}
					return <-response
				},
			},
		})
		defer checkClean(t)
		defer conn.Close()
		if err := conn.Dialable(); err != nil {
			t.Fatalf("unexpected dial error %v", err)
		}

		wire1 := conn.Acquire()

		go func() {
			// this should use the second wire
			if val, err := conn.Do(cmds.NewBlockingCompleted([]string{"PING"})).ToString(); err != nil {
				t.Errorf("unexpected error %v", err)
			} else if val != "BLOCK_RESPONSE" {
				t.Errorf("unexpected response %v", val)
			}
			close(blocking)
		}()
		<-blocking

		conn.Store(wire1)
		// this should use the first wire
		if val, err := conn.Do(cmds.NewBlockingCompleted([]string{"PING"})).ToString(); err != nil {
			t.Fatalf("unexpected error %v", err)
		} else if val != "ACQUIRED" {
			t.Fatalf("unexpected response %v", val)
		}

		response <- proto.NewResult(proto.Message{Type: '+', String: "BLOCK_RESPONSE"}, nil)
		<-blocking
	})
}

func TestConnCMDRetry(t *testing.T) {
	t.Run("wire info", func(t *testing.T) {
		conn, checkClean := setupConn([]*mockWire{
			{
				info: func() map[string]proto.Message {
					return map[string]proto.Message{"key": {Type: '+', String: "value"}}
				},
			},
		})
		defer checkClean(t)
		defer conn.Close()
		if info := conn.Info(); info == nil || info["key"].String != "value" {
			t.Fatalf("unexpected info %v", info)
		}
	})

	t.Run("retry single read", func(t *testing.T) {
		conn, checkClean := setupConn([]*mockWire{
			{
				do: func(cmd cmds.Completed) proto.Result {
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
				do: func(cmd cmds.Completed) proto.Result {
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
		defer conn.Close()
		// this should automatically use the second wire
		if val, err := conn.Do(cmds.NewReadOnlyCompleted([]string{"READONLY_COMMAND"})).ToString(); err != nil {
			t.Fatalf("unexpected error %v", err)
		} else if val != "READONLY_COMMAND_RESPONSE" {
			t.Fatalf("unexpected response %v", val)
		}
	})

	t.Run("retry multi read", func(t *testing.T) {
		conn, checkClean := setupConn([]*mockWire{
			{
				doMulti: func(multi ...cmds.Completed) []proto.Result {
					return []proto.Result{proto.NewErrResult(errors.New("network error"))}
				},
			},
			{
				doMulti: func(multi ...cmds.Completed) []proto.Result {
					return []proto.Result{proto.NewResult(proto.Message{Type: '+', String: "MULTI_COMMANDS_RESPONSE"}, nil)}
				},
			},
		})
		defer checkClean(t)
		defer conn.Close()
		// this should automatically use the second wire
		if val, err := conn.DoMulti(cmds.NewReadOnlyCompleted([]string{"READONLY_COMMAND"}))[0].ToString(); err != nil {
			t.Fatalf("unexpected error %v", err)
		} else if val != "MULTI_COMMANDS_RESPONSE" {
			t.Fatalf("unexpected response %v", val)
		}
	})

	t.Run("retry single read cache", func(t *testing.T) {
		conn, checkClean := setupConn([]*mockWire{
			{
				doCache: func(cmd cmds.Cacheable, ttl time.Duration) proto.Result {
					return proto.NewErrResult(errors.New("network error"))
				},
			},
			{
				doCache: func(cmd cmds.Cacheable, ttl time.Duration) proto.Result {
					return proto.NewResult(proto.Message{Type: '+', String: "READONLY_COMMAND_RESPONSE"}, nil)
				},
			},
		})
		defer checkClean(t)
		defer conn.Close()
		// this should automatically use the second wire
		if val, err := conn.DoCache(cmds.Cacheable(cmds.NewReadOnlyCompleted([]string{"READONLY_COMMAND"})), time.Second).ToString(); err != nil {
			t.Fatalf("unexpected error %v", err)
		} else if val != "READONLY_COMMAND_RESPONSE" {
			t.Fatalf("unexpected response %v", val)
		}
	})

	t.Run("not retry single write", func(t *testing.T) {
		conn, checkClean := setupConn([]*mockWire{
			{
				do: func(cmd cmds.Completed) proto.Result {
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
				do: func(cmd cmds.Completed) proto.Result {
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
		defer conn.Close()
		// this should only use the first wire
		if _, err := conn.Do(cmds.NewCompleted([]string{"WRITE_COMMAND"})).ToString(); err == nil || err.Error() != "network error" {
			t.Fatalf("unexpected error %v", err)
		}
		// this should use the second wire
		if val, err := conn.Do(cmds.NewCompleted([]string{"WRITE_COMMAND"})).ToString(); err != nil {
			t.Fatalf("unexpected error %v", err)
		} else if val != "WRITE_COMMAND_RESPONSE" {
			t.Fatalf("unexpected response %v", val)
		}
	})

	t.Run("not retry multi write", func(t *testing.T) {
		conn, checkClean := setupConn([]*mockWire{
			{
				doMulti: func(multi ...cmds.Completed) []proto.Result {
					return []proto.Result{proto.NewErrResult(errors.New("network error"))}
				},
			},
			{
				doMulti: func(multi ...cmds.Completed) []proto.Result {
					return []proto.Result{proto.NewResult(proto.Message{Type: '+', String: "MULTI_COMMANDS_RESPONSE"}, nil)}
				},
			},
		})
		defer checkClean(t)
		defer conn.Close()
		// this should only use the first wire
		if _, err := conn.DoMulti(
			cmds.NewReadOnlyCompleted([]string{"READONLY_COMMAND"}),
			cmds.NewCompleted([]string{"WRITE_COMMAND"}),
		)[0].ToString(); err == nil || err.Error() != "network error" {
			t.Fatalf("unexpected error %v", err)
		}
		// this should use the second wire
		if val, err := conn.DoMulti(
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

		conn, checkClean := setupConn([]*mockWire{
			{
				// leave first wire for pipeline calls
			},
			{
				do: func(cmd cmds.Completed) proto.Result {
					blocked <- struct{}{}
					return <-responses
				},
			},
			{
				do: func(cmd cmds.Completed) proto.Result {
					blocked <- struct{}{}
					return <-responses
				},
			},
		})
		defer checkClean(t)
		defer conn.Close()
		if err := conn.Dialable(); err != nil {
			t.Fatalf("unexpected dial error %v", err)
		}

		wg := sync.WaitGroup{}
		wg.Add(2)
		for i := 0; i < 2; i++ {
			go func() {
				if val, err := conn.Do(cmds.NewBlockingCompleted([]string{"BLOCK"})).ToString(); err != nil {
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

		conn, checkClean := setupConn([]*mockWire{
			{
				// leave first wire for pipeline calls
			},
			{
				doMulti: func(cmd ...cmds.Completed) []proto.Result {
					blocked <- struct{}{}
					return []proto.Result{<-responses}
				},
			},
			{
				doMulti: func(cmd ...cmds.Completed) []proto.Result {
					blocked <- struct{}{}
					return []proto.Result{<-responses}
				},
			},
		})
		defer checkClean(t)
		defer conn.Close()
		if err := conn.Dialable(); err != nil {
			t.Fatalf("unexpected dial error %v", err)
		}

		wg := sync.WaitGroup{}
		wg.Add(2)
		for i := 0; i < 2; i++ {
			go func() {
				if val, err := conn.DoMulti(
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

func TestConnDialRetry(t *testing.T) {
	setup := func() (*Conn, *int64) {
		var count int64
		return newConn("", Option{}, (*mockWire)(nil), func(dst string, opt Option) (net.Conn, error) {
			if count == 1 {
				return nil, nil
			}
			count++
			return nil, errors.New("network err")
		}, func(conn net.Conn, opt Option) (Wire, error) {
			return &mockWire{
				do: func(cmd cmds.Completed) proto.Result {
					return proto.NewResult(proto.Message{Type: '+', String: "PONG"}, nil)
				},
			}, nil
		}), &count
	}
	t.Run("retry on auto pipeline", func(t *testing.T) {
		conn, count := setup()
		defer conn.Close()
		if val, err := conn.Do(cmds.NewCompleted([]string{"PING"})).ToString(); err != nil {
			t.Fatalf("unexpected err %v", err)
		} else if val != "PONG" {
			t.Fatalf("unexpected response %v", val)
		}
		if *count != 1 {
			t.Fatalf("no dial err at all")
		}
	})

	t.Run("retry on blocking pool", func(t *testing.T) {
		conn, count := setup()
		defer conn.Close()
		if val, err := conn.Do(cmds.NewBlockingCompleted([]string{"PING"})).ToString(); err != nil {
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
	setup := func(b *testing.B) *Conn {
		c := NewConn("127.0.0.1:6379", Option{CacheSizeEachConn: DefaultCacheBytes}, func(dst string, opt Option) (conn net.Conn, err error) {
			return net.Dial("tcp", dst)
		})
		if err := c.Dialable(); err != nil {
			panic(err)
		}
		b.SetParallelism(100)
		b.ResetTimer()
		return c
	}
	b.Run("Do", func(b *testing.B) {
		c := setup(b)
		cmd := cmds.NewCompleted([]string{"GET", "a"})
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.Do(cmd)
			}
		})
	})
	b.Run("DoCache", func(b *testing.B) {
		c := setup(b)
		cmd := cmds.Cacheable(cmds.NewCompleted([]string{"GET", "a"}))
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.DoCache(cmd, time.Second*5)
			}
		})
	})
}
