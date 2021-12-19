package client

import (
	"errors"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/rueian/rueidis/internal/cmds"
	"github.com/rueian/rueidis/internal/mock"
	"github.com/rueian/rueidis/internal/proto"
	"github.com/rueian/rueidis/pkg/conn"
)

var slotsResp = proto.NewResult(proto.Message{Type: '*', Values: []proto.Message{
	{Type: '*', Values: []proto.Message{
		{Type: ':', Integer: 0},
		{Type: ':', Integer: 16383},
		{Type: '*', Values: []proto.Message{ // master
			{Type: '+', String: ""},
			{Type: ':', Integer: 0},
			{Type: '+', String: ""},
		}},
		{Type: '*', Values: []proto.Message{ // replica
			{Type: '+', String: ""},
			{Type: ':', Integer: 1},
			{Type: '+', String: ""},
		}},
	}},
}}, nil)

var singleSlotResp = proto.NewResult(proto.Message{Type: '*', Values: []proto.Message{
	{Type: '*', Values: []proto.Message{
		{Type: ':', Integer: 0},
		{Type: ':', Integer: 0},
		{Type: '*', Values: []proto.Message{ // master
			{Type: '+', String: ""},
			{Type: ':', Integer: 0},
			{Type: '+', String: ""},
		}},
	}},
}}, nil)

func TestClusterClientInit(t *testing.T) {
	t.Run("Init no nodes", func(t *testing.T) {
		if _, err := NewClusterClient(ClusterClientOption{InitAddress: []string{}}, func(dst string, opt conn.Option) conn.Conn { return nil }); err != ErrNoNodes {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("Init no dialable", func(t *testing.T) {
		v := errors.New("dial err")
		if _, err := NewClusterClient(ClusterClientOption{InitAddress: []string{":0"}}, func(dst string, opt conn.Option) conn.Conn {
			return &MockConn{DialableFn: func() error { return v }}
		}); err != v {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("Refresh err", func(t *testing.T) {
		v := errors.New("refresh err")
		if _, err := NewClusterClient(ClusterClientOption{InitAddress: []string{":0"}}, func(dst string, opt conn.Option) conn.Conn {
			return &MockConn{DoFn: func(cmd cmds.Completed) proto.Result { return proto.NewErrResult(v) }}
		}); err != v {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("Refresh retry", func(t *testing.T) {
		first := true
		if _, err := NewClusterClient(ClusterClientOption{InitAddress: []string{":0"}}, func(dst string, opt conn.Option) conn.Conn {
			return &MockConn{
				DoFn: func(cmd cmds.Completed) proto.Result {
					if first {
						first = false
						return proto.NewResult(proto.Message{Type: '*', Values: []proto.Message{}}, nil)
					}
					return slotsResp
				},
			}
		}); err != nil {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("Refresh retry err", func(t *testing.T) {
		v := errors.New("dial err")
		first := true
		if _, err := NewClusterClient(ClusterClientOption{InitAddress: []string{":0"}}, func(dst string, opt conn.Option) conn.Conn {
			return &MockConn{
				DoFn: func(cmd cmds.Completed) proto.Result {
					return proto.NewResult(proto.Message{Type: '*', Values: []proto.Message{}}, nil)
				},
				DialableFn: func() error {
					if first {
						first = false
						return nil
					}
					return v
				},
			}
		}); err != v {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("Refresh replace", func(t *testing.T) {
		if client, err := NewClusterClient(ClusterClientOption{InitAddress: []string{":1", ":2"}, ShuffleInit: true}, func(dst string, opt conn.Option) conn.Conn {
			return &MockConn{
				DoFn: func(cmd cmds.Completed) proto.Result {
					return slotsResp
				},
			}
		}); err != nil {
			t.Fatalf("unexpected err %v", err)
		} else if nodes := client.nodes(); len(nodes) != 1 || nodes[0] != ":0" {
			t.Fatalf("unexpected nodes %v", nodes)
		}
	})
}

func TestClusterClient(t *testing.T) {
	m := &MockConn{
		DoFn: func(cmd cmds.Completed) proto.Result {
			if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
				return slotsResp
			}
			return proto.Result{}
		},
	}

	client, err := NewClusterClient(ClusterClientOption{InitAddress: []string{":0"}, ShuffleInit: true}, func(dst string, opt conn.Option) conn.Conn {
		return m
	})
	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}

	t.Run("Delegate Do with no slot", func(t *testing.T) {
		c := client.Cmd.Info().Build()
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

	t.Run("Delegate Close", func(t *testing.T) {
		called := make(chan struct{})
		m.CloseFn = func() {
			close(called)
		}
		client.Close()
		<-called
	})

	t.Run("DedicatedWire Err", func(t *testing.T) {
		v := errors.New("fn err")
		if err := client.DedicatedWire(func(client *DedicatedClusterClient) error {
			return v
		}); err != v {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("DedicatedWire No Slot Err", func(t *testing.T) {
		defer func() {
			if err := recover(); err != "the first command in DedicatedClusterClient should contain the slot key" {
				t.Errorf("DedicatedWire should panic if no slot is selected")
			}
		}()
		client.DedicatedWire(func(c *DedicatedClusterClient) error {
			return c.Do(client.Cmd.Info().Build()).Error()
		})
	})

	t.Run("DedicatedWire Cross Slot Err", func(t *testing.T) {
		defer func() {
			if err := recover(); err != "cross slot command in DedicatedWire is prohibited" {
				t.Errorf("DedicatedWire should panic if cross slots is used")
			}
		}()
		m.AcquireFn = func() conn.Wire { return &mock.Wire{} }
		client.DedicatedWire(func(c *DedicatedClusterClient) error {
			c.Do(client.Cmd.Get().Key("a").Build()).Error()
			return c.Do(client.Cmd.Get().Key("b").Build()).Error()
		})
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
		m.AcquireFn = func() conn.Wire {
			return w
		}
		stored := false
		m.StoreFn = func(ww conn.Wire) {
			if ww != w {
				t.Fatalf("received unexpected wire %v", ww)
			}
			stored = true
		}
		if err := client.DedicatedWire(func(c *DedicatedClusterClient) error {
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

	t.Run("NewLuaScript Delegate", func(t *testing.T) {
		m.DoFn = func(cmd cmds.Completed) proto.Result {
			if cmd.Commands()[0] == "EVALSHA" {
				return proto.NewResult(proto.Message{Type: '-', String: "NOSCRIPT"}, nil)
			}
			if cmd.Commands()[0] != "EVAL" {
				t.Fatalf("unexpected command %v", cmd.Commands())
			}
			return proto.NewResult(proto.Message{Type: '+', String: strings.Join(cmd.Commands(), " ")}, nil)
		}
		if v, err := client.NewLuaScript("NewLuaScript").Exec([]string{"1"}, []string{"2", "3"}).ToString(); err != nil || v != "EVAL NewLuaScript 1 1 2 3" {
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
		if v, err := client.NewLuaScriptReadOnly("NewLuaScriptReadOnly").Exec([]string{"1"}, []string{"2", "3"}).ToString(); err != nil || v != "EVAL_RO NewLuaScriptReadOnly 1 1 2 3" {
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

func TestClusterClientErr(t *testing.T) {
	t.Run("refresh err on pick", func(t *testing.T) {
		first := true
		v := errors.New("refresh err")
		m := &MockConn{
			DoFn: func(cmd cmds.Completed) proto.Result {
				if first {
					first = false
					return singleSlotResp
				}
				return proto.NewErrResult(v)
			},
		}
		client, err := NewClusterClient(ClusterClientOption{InitAddress: []string{":0"}, ShuffleInit: true}, func(dst string, opt conn.Option) conn.Conn {
			return m
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if err := client.Do(client.Cmd.Get().Key("a").Build()).Error(); err != v {
			t.Fatalf("unexpected err %v", err)
		}
		if err := client.DoCache(client.Cmd.Get().Key("a").Cache(), 100).Error(); err != v {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("refresh empty on pick", func(t *testing.T) {
		m := &MockConn{DoFn: func(cmd cmds.Completed) proto.Result {
			return singleSlotResp
		}}
		client, err := NewClusterClient(ClusterClientOption{InitAddress: []string{":0"}, ShuffleInit: true}, func(dst string, opt conn.Option) conn.Conn {
			return m
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if err := client.Do(client.Cmd.Get().Key("a").Build()).Error(); err != ErrNoSlot {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("refresh empty on pick in dedicated wire", func(t *testing.T) {
		m := &MockConn{DoFn: func(cmd cmds.Completed) proto.Result {
			return singleSlotResp
		}}
		client, err := NewClusterClient(ClusterClientOption{InitAddress: []string{":0"}, ShuffleInit: true}, func(dst string, opt conn.Option) conn.Conn {
			return m
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if err := client.DedicatedWire(func(c *DedicatedClusterClient) error {
			return c.Do(client.Cmd.Get().Key("a").Build()).Error()
		}); err != ErrNoSlot {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("refresh empty on pick in dedicated wire (multi)", func(t *testing.T) {
		m := &MockConn{DoFn: func(cmd cmds.Completed) proto.Result {
			return singleSlotResp
		}}
		client, err := NewClusterClient(ClusterClientOption{InitAddress: []string{":0"}, ShuffleInit: true}, func(dst string, opt conn.Option) conn.Conn {
			return m
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if err := client.DedicatedWire(func(c *DedicatedClusterClient) error {
			for _, v := range c.DoMulti(client.Cmd.Get().Key("a").Build()) {
				if err := v.Error(); err != nil {
					return err
				}
			}
			return nil
		}); err != ErrNoSlot {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("slot moved", func(t *testing.T) {
		count := 0
		m := &MockConn{DoFn: func(cmd cmds.Completed) proto.Result {
			if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
				return slotsResp
			}
			if count < 3 {
				count++
				return proto.NewResult(proto.Message{Type: '-', String: "MOVED 0 :1"}, nil)
			}
			return proto.NewResult(proto.Message{Type: '+', String: "b"}, nil)
		}}
		client, err := NewClusterClient(ClusterClientOption{InitAddress: []string{":0"}}, func(dst string, opt conn.Option) conn.Conn {
			return m
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.Do(client.Cmd.Get().Key("a").Build()).ToString(); err != nil || v != "b" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot moved (cache)", func(t *testing.T) {
		count := 0
		m := &MockConn{
			DoFn: func(cmd cmds.Completed) proto.Result {
				return slotsResp
			},
			DoCacheFn: func(cmd cmds.Cacheable, ttl time.Duration) proto.Result {
				if count < 3 {
					count++
					return proto.NewResult(proto.Message{Type: '-', String: "MOVED 0 :1"}, nil)
				}
				return proto.NewResult(proto.Message{Type: '+', String: "b"}, nil)
			},
		}
		client, err := NewClusterClient(ClusterClientOption{InitAddress: []string{":0"}}, func(dst string, opt conn.Option) conn.Conn {
			return m
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.DoCache(client.Cmd.Get().Key("a").Cache(), 100).ToString(); err != nil || v != "b" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot asking", func(t *testing.T) {
		count := 0
		m := &MockConn{
			DoFn: func(cmd cmds.Completed) proto.Result {
				if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
					return slotsResp
				}
				return proto.NewResult(proto.Message{Type: '-', String: "ASK 0 :1"}, nil)
			},
			DoMultiFn: func(multi ...cmds.Completed) []proto.Result {
				if count < 3 {
					count++
					return []proto.Result{{}, proto.NewResult(proto.Message{Type: '-', String: "ASK 0 :1"}, nil)}
				}
				return []proto.Result{{}, proto.NewResult(proto.Message{Type: '+', String: "b"}, nil)}
			},
		}
		client, err := NewClusterClient(ClusterClientOption{InitAddress: []string{":0"}}, func(dst string, opt conn.Option) conn.Conn {
			return m
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.Do(client.Cmd.Get().Key("a").Build()).ToString(); err != nil || v != "b" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot asking (cache)", func(t *testing.T) {
		count := 0
		m := &MockConn{
			DoFn: func(cmd cmds.Completed) proto.Result {
				return slotsResp
			},
			DoCacheFn: func(cmd cmds.Cacheable, ttl time.Duration) proto.Result {
				return proto.NewResult(proto.Message{Type: '-', String: "ASK 0 :1"}, nil)
			},
			DoMultiFn: func(multi ...cmds.Completed) []proto.Result {
				if count < 3 {
					count++
					return []proto.Result{{}, proto.NewResult(proto.Message{Type: '-', String: "ASK 0 :1"}, nil)}
				}
				return []proto.Result{{}, proto.NewResult(proto.Message{Type: '+', String: "b"}, nil)}
			},
		}
		client, err := NewClusterClient(ClusterClientOption{InitAddress: []string{":0"}}, func(dst string, opt conn.Option) conn.Conn {
			return m
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.DoCache(client.Cmd.Get().Key("a").Cache(), 100).ToString(); err != nil || v != "b" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot try again", func(t *testing.T) {
		count := 0
		m := &MockConn{DoFn: func(cmd cmds.Completed) proto.Result {
			if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
				return slotsResp
			}
			if count < 3 {
				count++
				return proto.NewResult(proto.Message{Type: '-', String: "TRYAGAIN"}, nil)
			}
			return proto.NewResult(proto.Message{Type: '+', String: "b"}, nil)
		}}
		client, err := NewClusterClient(ClusterClientOption{InitAddress: []string{":0"}}, func(dst string, opt conn.Option) conn.Conn {
			return m
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.Do(client.Cmd.Get().Key("a").Build()).ToString(); err != nil || v != "b" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot try again (cache)", func(t *testing.T) {
		count := 0
		m := &MockConn{
			DoFn: func(cmd cmds.Completed) proto.Result {
				return slotsResp
			},
			DoCacheFn: func(cmd cmds.Cacheable, ttl time.Duration) proto.Result {
				if count < 3 {
					count++
					return proto.NewResult(proto.Message{Type: '-', String: "TRYAGAIN"}, nil)
				}
				return proto.NewResult(proto.Message{Type: '+', String: "b"}, nil)
			},
		}
		client, err := NewClusterClient(ClusterClientOption{InitAddress: []string{":0"}}, func(dst string, opt conn.Option) conn.Conn {
			return m
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.DoCache(client.Cmd.Get().Key("a").Cache(), 100).ToString(); err != nil || v != "b" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})
}

func TestHashObjectClusterClientAdapter(t *testing.T) {
	m := &MockConn{
		DoFn: func(cmd cmds.Completed) proto.Result {
			if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
				return slotsResp
			}
			return proto.Result{}
		},
	}
	client, err := NewClusterClient(ClusterClientOption{InitAddress: []string{":0"}}, func(dst string, opt conn.Option) conn.Conn {
		return m
	})
	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}
	adapter := &HashObjectClusterClientAdapter{c: client}

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
