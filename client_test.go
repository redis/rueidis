package rueidis

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/rueian/rueidis/internal/cmds"
)

type mockConn struct {
	DoFn      func(cmd cmds.Completed) RedisResult
	DoCacheFn func(cmd cmds.Cacheable, ttl time.Duration) RedisResult
	DoMultiFn func(multi ...cmds.Completed) []RedisResult
	InfoFn    func() map[string]RedisMessage
	ErrorFn   func() error
	CloseFn   func()
	DialFn    func() error
	AcquireFn func() wire
	StoreFn   func(w wire)

	disconnectedFn func(err error)
}

func (m *mockConn) Dial() error {
	if m.DialFn != nil {
		return m.DialFn()
	}
	return nil
}

func (m *mockConn) Acquire() wire {
	if m.AcquireFn != nil {
		return m.AcquireFn()
	}
	return nil
}

func (m *mockConn) Store(w wire) {
	if m.StoreFn != nil {
		m.StoreFn(w)
	}
}

func (m *mockConn) Do(cmd cmds.Completed) RedisResult {
	if m.DoFn != nil {
		return m.DoFn(cmd)
	}
	return RedisResult{}
}

func (m *mockConn) DoCache(cmd cmds.Cacheable, ttl time.Duration) RedisResult {
	if m.DoCacheFn != nil {
		return m.DoCacheFn(cmd, ttl)
	}
	return RedisResult{}
}

func (m *mockConn) DoMulti(multi ...cmds.Completed) []RedisResult {
	if m.DoMultiFn != nil {
		return m.DoMultiFn(multi...)
	}
	return nil
}

func (m *mockConn) Info() map[string]RedisMessage {
	if m.InfoFn != nil {
		return m.InfoFn()
	}
	return nil
}

func (m *mockConn) Error() error {
	if m.ErrorFn != nil {
		return m.ErrorFn()
	}
	return nil
}

func (m *mockConn) Close() {
	if m.CloseFn != nil {
		m.CloseFn()
	}
}

func (m *mockConn) OnDisconnected(fn func(err error)) {
	m.disconnectedFn = fn
}

func (m *mockConn) TriggerDisconnect(err error) {
	if m.disconnectedFn != nil {
		m.disconnectedFn(err)
	}
}

func TestNewSingleClientNoNode(t *testing.T) {
	if _, err := newSingleClient(ClientOption{}, func(dst string, opt ClientOption) conn {
		return nil
	}); err != ErrNoAddr {
		t.Fatalf("unexpected err %v", err)
	}
}

func TestNewSingleClientError(t *testing.T) {
	v := errors.New("dail err")
	if _, err := newSingleClient(ClientOption{InitAddress: []string{""}}, func(dst string, opt ClientOption) conn {
		return &mockConn{DialFn: func() error { return v }}
	}); err != v {
		t.Fatalf("unexpected err %v", err)
	}
}

func TestSingleClient(t *testing.T) {
	m := &mockConn{}
	client, err := newSingleClient(ClientOption{InitAddress: []string{""}}, func(dst string, opt ClientOption) conn {
		return m
	})
	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}

	t.Run("Delegate Do", func(t *testing.T) {
		c := client.B().Get().Key("Do").Build()
		m.DoFn = func(cmd cmds.Completed) RedisResult {
			if !reflect.DeepEqual(cmd.Commands(), c.Commands()) {
				t.Fatalf("unexpected command %v", cmd)
			}
			return newResult(RedisMessage{typ: '+', string: "Do"}, nil)
		}
		if v, err := client.Do(context.Background(), c).ToString(); err != nil || v != "Do" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("Delegate DoCache", func(t *testing.T) {
		c := client.B().Get().Key("DoCache").Cache()
		m.DoCacheFn = func(cmd cmds.Cacheable, ttl time.Duration) RedisResult {
			if !reflect.DeepEqual(cmd.Commands(), c.Commands()) || ttl != 100 {
				t.Fatalf("unexpected command %v, %v", cmd, ttl)
			}
			return newResult(RedisMessage{typ: '+', string: "DoCache"}, nil)
		}
		if v, err := client.DoCache(context.Background(), c, 100).ToString(); err != nil || v != "DoCache" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("Delegate Close", func(t *testing.T) {
		called := false
		m.CloseFn = func() { called = true }
		client.Close()
		if !called {
			t.Fatalf("Close is not delegated")
		}
	})

	t.Run("Dedicated Err", func(t *testing.T) {
		v := errors.New("fn err")
		if err := client.Dedicated(func(client DedicatedClient) error {
			return v
		}); err != v {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("Dedicated Delegate", func(t *testing.T) {
		w := &mockWire{
			DoFn: func(cmd cmds.Completed) RedisResult {
				return newResult(RedisMessage{typ: '+', string: "Delegate"}, nil)
			},
			DoMultiFn: func(cmd ...cmds.Completed) []RedisResult {
				return []RedisResult{newResult(RedisMessage{typ: '+', string: "Delegate"}, nil)}
			},
		}
		m.AcquireFn = func() wire {
			return w
		}
		stored := false
		m.StoreFn = func(ww wire) {
			if ww != w {
				t.Fatalf("received unexpected wire %v", ww)
			}
			stored = true
		}
		if err := client.Dedicated(func(c DedicatedClient) error {
			if v, err := c.Do(context.Background(), c.B().Get().Key("a").Build()).ToString(); err != nil || v != "Delegate" {
				t.Fatalf("unexpected response %v %v", v, err)
			}
			if v := c.DoMulti(context.Background()); len(v) != 0 {
				t.Fatalf("received unexpected response %v", v)
			}
			for _, resp := range c.DoMulti(context.Background(), c.B().Get().Key("a").Build()) {
				if v, err := resp.ToString(); err != nil || v != "Delegate" {
					t.Fatalf("unexpected response %v %v", v, err)
				}
			}
			return nil
		}); err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if !stored {
			t.Fatalf("Dedicated desn't put back the wire")
		}
	})
}
