package rueidis

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/rueian/rueidis/internal/cmds"
	"github.com/rueian/rueidis/internal/mock"
	"github.com/rueian/rueidis/internal/proto"
)

type MockConn struct {
	DoFn      func(cmd cmds.Completed) proto.Result
	DoCacheFn func(cmd cmds.Cacheable, ttl time.Duration) proto.Result
	DoMultiFn func(multi ...cmds.Completed) []proto.Result
	InfoFn    func() map[string]proto.Message
	ErrorFn   func() error
	CloseFn   func()
	DialFn    func() error
	AcquireFn func() wire
	StoreFn   func(w wire)

	disconnectedFn func(err error)
}

func (m *MockConn) Dial() error {
	if m.DialFn != nil {
		return m.DialFn()
	}
	return nil
}

func (m *MockConn) Acquire() wire {
	if m.AcquireFn != nil {
		return m.AcquireFn()
	}
	return nil
}

func (m *MockConn) Store(w wire) {
	if m.StoreFn != nil {
		m.StoreFn(w)
	}
}

func (m *MockConn) Do(cmd cmds.Completed) proto.Result {
	if m.DoFn != nil {
		return m.DoFn(cmd)
	}
	return proto.Result{}
}

func (m *MockConn) DoCache(cmd cmds.Cacheable, ttl time.Duration) proto.Result {
	if m.DoCacheFn != nil {
		return m.DoCacheFn(cmd, ttl)
	}
	return proto.Result{}
}

func (m *MockConn) DoMulti(multi ...cmds.Completed) []proto.Result {
	if m.DoMultiFn != nil {
		return m.DoMultiFn(multi...)
	}
	return nil
}

func (m *MockConn) Info() map[string]proto.Message {
	if m.InfoFn != nil {
		return m.InfoFn()
	}
	return nil
}

func (m *MockConn) Error() error {
	if m.ErrorFn != nil {
		return m.ErrorFn()
	}
	return nil
}

func (m *MockConn) Close() {
	if m.CloseFn != nil {
		m.CloseFn()
	}
}

func (m *MockConn) OnDisconnected(fn func(err error)) {
	m.disconnectedFn = fn
}

func (m *MockConn) TriggerDisconnect(err error) {
	if m.disconnectedFn != nil {
		m.disconnectedFn(err)
	}
}

func TestNewSingleClientError(t *testing.T) {
	v := errors.New("dail err")
	if _, err := newSingleClient(SingleClientOption{}, func(dst string, opt ConnOption) conn {
		return &MockConn{DialFn: func() error { return v }}
	}); err != v {
		t.Fatalf("unexpected err %v", err)
	}
}

func TestSingleClient(t *testing.T) {
	m := &MockConn{}
	client, err := newSingleClient(SingleClientOption{}, func(dst string, opt ConnOption) conn {
		return m
	})
	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}

	t.Run("Delegate Do", func(t *testing.T) {
		c := client.B().Get().Key("Do").Build()
		m.DoFn = func(cmd cmds.Completed) proto.Result {
			if !reflect.DeepEqual(cmd.Commands(), c.Commands()) {
				t.Fatalf("unexpected command %v", cmd)
			}
			return proto.NewResult(proto.Message{Type: '+', String: "Do"}, nil)
		}
		if v, err := client.Do(context.Background(), c).ToString(); err != nil || v != "Do" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("Delegate DoCache", func(t *testing.T) {
		c := client.B().Get().Key("DoCache").Cache()
		m.DoCacheFn = func(cmd cmds.Cacheable, ttl time.Duration) proto.Result {
			if !reflect.DeepEqual(cmd.Commands(), c.Commands()) || ttl != 100 {
				t.Fatalf("unexpected command %v, %v", cmd, ttl)
			}
			return proto.NewResult(proto.Message{Type: '+', String: "DoCache"}, nil)
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
		w := &mock.Wire{
			DoFn: func(cmd cmds.Completed) proto.Result {
				return proto.NewResult(proto.Message{Type: '+', String: "Delegate"}, nil)
			},
			DoMultiFn: func(cmd ...cmds.Completed) []proto.Result {
				return []proto.Result{proto.NewResult(proto.Message{Type: '+', String: "Delegate"}, nil)}
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
				t.Fatalf("unexpected respone %v %v", v, err)
			}
			if v := c.DoMulti(context.Background()); len(v) != 0 {
				t.Fatalf("received unexpected respone %v", v)
			}
			for _, resp := range c.DoMulti(context.Background(), c.B().Get().Key("a").Build()) {
				if v, err := resp.ToString(); err != nil || v != "Delegate" {
					t.Fatalf("unexpected respone %v %v", v, err)
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
