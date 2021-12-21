package rueidis

import (
	"errors"
	"reflect"
	"strings"
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
		c := client.Cmd.Get().Key("Do").Build()
		m.DoFn = func(cmd cmds.Completed) proto.Result {
			if !reflect.DeepEqual(cmd.Commands(), c.Commands()) {
				t.Fatalf("unexpected command %v", cmd)
			}
			return proto.NewResult(proto.Message{Type: '+', String: "Do"}, nil)
		}
		if v, err := client.Do(c).ToString(); err != nil || v != "Do" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("Delegate DoCache", func(t *testing.T) {
		c := client.Cmd.Get().Key("DoCache").Cache()
		m.DoCacheFn = func(cmd cmds.Cacheable, ttl time.Duration) proto.Result {
			if !reflect.DeepEqual(cmd.Commands(), c.Commands()) || ttl != 100 {
				t.Fatalf("unexpected command %v, %v", cmd, ttl)
			}
			return proto.NewResult(proto.Message{Type: '+', String: "DoCache"}, nil)
		}
		if v, err := client.DoCache(c, 100).ToString(); err != nil || v != "DoCache" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("Delegate Info", func(t *testing.T) {
		m.InfoFn = func() map[string]proto.Message {
			return map[string]proto.Message{}
		}
		if client.Info() == nil {
			t.Fatalf("unexpected nil info")
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

	t.Run("DedicatedWire Err", func(t *testing.T) {
		v := errors.New("fn err")
		if err := client.DedicatedWire(func(client *DedicatedSingleClient) error {
			return v
		}); err != v {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("DedicatedWire Delegate", func(t *testing.T) {
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
		if err := client.DedicatedWire(func(c *DedicatedSingleClient) error {
			if v, err := c.Do(client.Cmd.Get().Key("a").Build()).ToString(); err != nil || v != "Delegate" {
				t.Fatalf("unexpected respone %v %v", v, err)
			}
			if v := c.DoMulti(); len(v) != 0 {
				t.Fatalf("received unexpected respone %v", v)
			}
			for _, resp := range c.DoMulti(client.Cmd.Get().Key("a").Build()) {
				if v, err := resp.ToString(); err != nil || v != "Delegate" {
					t.Fatalf("unexpected respone %v %v", v, err)
				}
			}
			return nil
		}); err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if !stored {
			t.Fatalf("DedicatedWire desn't put back the wire")
		}
	})

	t.Run("newLuaScript Delegate", func(t *testing.T) {
		m.DoFn = func(cmd cmds.Completed) proto.Result {
			if cmd.Commands()[0] == "EVALSHA" {
				return proto.NewResult(proto.Message{Type: '-', String: "NOSCRIPT"}, nil)
			}
			if cmd.Commands()[0] != "EVAL" {
				t.Fatalf("unexpected command %v", cmd.Commands())
			}
			return proto.NewResult(proto.Message{Type: '+', String: strings.Join(cmd.Commands(), " ")}, nil)
		}
		if v, err := client.NewLuaScript("newLuaScript").Exec([]string{"1", "2"}, []string{"3", "4"}).ToString(); err != nil || v != "EVAL newLuaScript 2 1 2 3 4" {
			t.Fatalf("unexpected respone %v %v", v, err)
		}
	})

	t.Run("NewLuaScriptReadOnly Delegate", func(t *testing.T) {
		m.DoFn = func(cmd cmds.Completed) proto.Result {
			if cmd.Commands()[0] == "EVALSHA_RO" {
				return proto.NewResult(proto.Message{Type: '-', String: "NOSCRIPT"}, nil)
			}
			if cmd.Commands()[0] != "EVAL_RO" {
				t.Fatalf("unexpected command %v", cmd.Commands())
			}
			return proto.NewResult(proto.Message{Type: '+', String: strings.Join(cmd.Commands(), " ")}, nil)
		}
		if v, err := client.NewLuaScriptReadOnly("NewLuaScriptReadOnly").Exec([]string{"1", "2"}, []string{"3", "4"}).ToString(); err != nil || v != "EVAL_RO NewLuaScriptReadOnly 2 1 2 3 4" {
			t.Fatalf("unexpected respone %v %v", v, err)
		}
	})

	t.Run("NewHashRepository Delegate", func(t *testing.T) {
		repo := client.NewHashRepository("", schema{})
		if repo == nil {
			t.Fatalf("unexpected nil repo")
		}
	})
}

func TestHashObjectSingleClientAdapter(t *testing.T) {
	m := &MockConn{}
	client, err := newSingleClient(SingleClientOption{}, func(dst string, opt ConnOption) conn {
		return m
	})
	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}
	adapter := &hashObjectSingleClientAdapter{c: client}

	t.Run("Save Delegate", func(t *testing.T) {
		m.DoFn = func(cmd cmds.Completed) proto.Result {
			if v := strings.Join(cmd.Commands(), " "); v != "HSET k a b" {
				return proto.NewResult(proto.Message{Type: '-', String: "wrong command " + v}, nil)
			}
			return proto.NewResult(proto.Message{Type: '+', String: "OK"}, nil)
		}
		if err := adapter.Save("k", map[string]string{"a": "b"}); err != nil {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("Fetch Delegate", func(t *testing.T) {
		m.DoFn = func(cmd cmds.Completed) proto.Result {
			if v := strings.Join(cmd.Commands(), " "); v != "HGETALL k" {
				return proto.NewResult(proto.Message{Type: '-', String: "wrong command " + v}, nil)
			}
			return proto.NewResult(proto.Message{Type: '%', Values: []proto.Message{
				{Type: '+', String: "a"},
				{Type: '+', String: "b"},
			}}, nil)
		}
		if v, err := adapter.Fetch("k"); err != nil || v["a"].String != "b" {
			t.Fatalf("unexpected response %v", err)
		}
	})

	t.Run("FetchCache Delegate", func(t *testing.T) {
		m.DoCacheFn = func(cmd cmds.Cacheable, ttl time.Duration) proto.Result {
			if v := strings.Join(cmd.Commands(), " "); v != "HGETALL k" || ttl != 100 {
				return proto.NewResult(proto.Message{Type: '-', String: "wrong command " + v}, nil)
			}
			return proto.NewResult(proto.Message{Type: '%', Values: []proto.Message{
				{Type: '+', String: "a"},
				{Type: '+', String: "b"},
			}}, nil)
		}
		if v, err := adapter.FetchCache("k", 100); err != nil || v["a"].String != "b" {
			t.Fatalf("unexpected response %v", err)
		}
	})

	t.Run("Remove Delegate", func(t *testing.T) {
		m.DoFn = func(cmd cmds.Completed) proto.Result {
			if v := strings.Join(cmd.Commands(), " "); v != "DEL k" {
				return proto.NewResult(proto.Message{Type: '-', String: "wrong command " + v}, nil)
			}
			return proto.NewResult(proto.Message{Type: '+', String: "OK"}, nil)
		}
		if err := adapter.Remove("k"); err != nil {
			t.Fatalf("unexpected err %v", err)
		}
	})
}

type schema struct {
	ID  string `redis:"-,pk"`
	Ver int64  `redis:"_v"`
}
