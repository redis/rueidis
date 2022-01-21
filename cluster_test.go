package rueidis

import (
	"context"
	"errors"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/rueian/rueidis/internal/cmds"
)

var slotsResp = newResult(RedisMessage{typ: '*', values: []RedisMessage{
	{typ: '*', values: []RedisMessage{
		{typ: ':', integer: 0},
		{typ: ':', integer: 16383},
		{typ: '*', values: []RedisMessage{ // master
			{typ: '+', string: ""},
			{typ: ':', integer: 0},
			{typ: '+', string: ""},
		}},
		{typ: '*', values: []RedisMessage{ // replica
			{typ: '+', string: ""},
			{typ: ':', integer: 1},
			{typ: '+', string: ""},
		}},
	}},
}}, nil)

var singleSlotResp = newResult(RedisMessage{typ: '*', values: []RedisMessage{
	{typ: '*', values: []RedisMessage{
		{typ: ':', integer: 0},
		{typ: ':', integer: 0},
		{typ: '*', values: []RedisMessage{ // master
			{typ: '+', string: ""},
			{typ: ':', integer: 0},
			{typ: '+', string: ""},
		}},
	}},
}}, nil)

func TestClusterClientInit(t *testing.T) {
	t.Run("Init no nodes", func(t *testing.T) {
		if _, err := newClusterClient(ClientOption{InitAddress: []string{}}, func(dst string, opt ClientOption) conn { return nil }); err != ErrNoAddr {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("Init no dialable", func(t *testing.T) {
		v := errors.New("dial err")
		if _, err := newClusterClient(ClientOption{InitAddress: []string{":0"}}, func(dst string, opt ClientOption) conn {
			return &mockConn{DialFn: func() error { return v }}
		}); err != v {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("Refresh err", func(t *testing.T) {
		v := errors.New("refresh err")
		if _, err := newClusterClient(ClientOption{InitAddress: []string{":0"}}, func(dst string, opt ClientOption) conn {
			return &mockConn{DoFn: func(cmd cmds.Completed) RedisResult { return newErrResult(v) }}
		}); err != v {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("Refresh retry", func(t *testing.T) {
		first := true
		if _, err := newClusterClient(ClientOption{InitAddress: []string{":0"}}, func(dst string, opt ClientOption) conn {
			return &mockConn{
				DoFn: func(cmd cmds.Completed) RedisResult {
					if first {
						first = false
						return newResult(RedisMessage{typ: '*', values: []RedisMessage{}}, nil)
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
		if _, err := newClusterClient(ClientOption{InitAddress: []string{":0"}}, func(dst string, opt ClientOption) conn {
			return &mockConn{
				DoFn: func(cmd cmds.Completed) RedisResult {
					return newResult(RedisMessage{typ: '*', values: []RedisMessage{}}, nil)
				},
				DialFn: func() error {
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
		if client, err := newClusterClient(ClientOption{InitAddress: []string{":1", ":2"}, ShuffleInit: true}, func(dst string, opt ClientOption) conn {
			return &mockConn{
				DoFn: func(cmd cmds.Completed) RedisResult {
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
	m := &mockConn{
		DoFn: func(cmd cmds.Completed) RedisResult {
			if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
				return slotsResp
			}
			return RedisResult{}
		},
	}

	client, err := newClusterClient(ClientOption{InitAddress: []string{":0"}, ShuffleInit: true}, func(dst string, opt ClientOption) conn {
		return m
	})
	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}

	t.Run("Delegate Do with no slot", func(t *testing.T) {
		c := client.B().Info().Build()
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
		called := make(chan struct{})
		m.CloseFn = func() {
			close(called)
		}
		client.Close()
		<-called
	})

	t.Run("Dedicated Err", func(t *testing.T) {
		v := errors.New("fn err")
		if err := client.Dedicated(func(client DedicatedClient) error {
			return v
		}); err != v {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("Dedicated No Slot Err", func(t *testing.T) {
		defer func() {
			if err := recover(); err != panicMsgNoSlot {
				t.Errorf("Dedicated should panic if no slot is selected")
			}
		}()
		builder := cmds.NewBuilder(cmds.NoSlot)
		client.Dedicated(func(c DedicatedClient) error {
			return c.Do(context.Background(), builder.Info().Build()).Error()
		})
	})

	t.Run("Dedicated Cross Slot Err", func(t *testing.T) {
		defer func() {
			if err := recover(); err != panicMsgCxSlot {
				t.Errorf("Dedicated should panic if cross slots is used")
			}
		}()
		m.AcquireFn = func() wire { return &mockWire{} }
		client.Dedicated(func(c DedicatedClient) error {
			c.Do(context.Background(), c.B().Get().Key("a").Build()).Error()
			return c.Do(context.Background(), c.B().Get().Key("b").Build()).Error()
		})
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

func TestClusterClientErr(t *testing.T) {
	t.Run("refresh err on pick", func(t *testing.T) {
		first := true
		v := errors.New("refresh err")
		m := &mockConn{
			DoFn: func(cmd cmds.Completed) RedisResult {
				if first {
					first = false
					return singleSlotResp
				}
				return newErrResult(v)
			},
		}
		client, err := newClusterClient(ClientOption{InitAddress: []string{":0"}, ShuffleInit: true}, func(dst string, opt ClientOption) conn {
			return m
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if err := client.Do(context.Background(), client.B().Get().Key("a").Build()).Error(); err != v {
			t.Fatalf("unexpected err %v", err)
		}
		if err := client.DoCache(context.Background(), client.B().Get().Key("a").Cache(), 100).Error(); err != v {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("refresh empty on pick", func(t *testing.T) {
		m := &mockConn{DoFn: func(cmd cmds.Completed) RedisResult {
			return singleSlotResp
		}}
		client, err := newClusterClient(ClientOption{InitAddress: []string{":0"}, ShuffleInit: true}, func(dst string, opt ClientOption) conn {
			return m
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if err := client.Do(context.Background(), client.B().Get().Key("a").Build()).Error(); err != ErrNoSlot {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("refresh empty on pick in dedicated wire", func(t *testing.T) {
		m := &mockConn{DoFn: func(cmd cmds.Completed) RedisResult {
			return singleSlotResp
		}}
		client, err := newClusterClient(ClientOption{InitAddress: []string{":0"}, ShuffleInit: true}, func(dst string, opt ClientOption) conn {
			return m
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if err := client.Dedicated(func(c DedicatedClient) error {
			return c.Do(context.Background(), c.B().Get().Key("a").Build()).Error()
		}); err != ErrNoSlot {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("refresh empty on pick in dedicated wire (multi)", func(t *testing.T) {
		m := &mockConn{DoFn: func(cmd cmds.Completed) RedisResult {
			return singleSlotResp
		}}
		client, err := newClusterClient(ClientOption{InitAddress: []string{":0"}, ShuffleInit: true}, func(dst string, opt ClientOption) conn {
			return m
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if err := client.Dedicated(func(c DedicatedClient) error {
			for _, v := range c.DoMulti(context.Background(), c.B().Get().Key("a").Build()) {
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
		m := &mockConn{DoFn: func(cmd cmds.Completed) RedisResult {
			if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
				return slotsResp
			}
			if count < 3 {
				count++
				return newResult(RedisMessage{typ: '-', string: "MOVED 0 :1"}, nil)
			}
			return newResult(RedisMessage{typ: '+', string: "b"}, nil)
		}}
		client, err := newClusterClient(ClientOption{InitAddress: []string{":0"}}, func(dst string, opt ClientOption) conn {
			return m
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.Do(context.Background(), client.B().Get().Key("a").Build()).ToString(); err != nil || v != "b" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot moved (cache)", func(t *testing.T) {
		count := 0
		m := &mockConn{
			DoFn: func(cmd cmds.Completed) RedisResult {
				return slotsResp
			},
			DoCacheFn: func(cmd cmds.Cacheable, ttl time.Duration) RedisResult {
				if count < 3 {
					count++
					return newResult(RedisMessage{typ: '-', string: "MOVED 0 :1"}, nil)
				}
				return newResult(RedisMessage{typ: '+', string: "b"}, nil)
			},
		}
		client, err := newClusterClient(ClientOption{InitAddress: []string{":0"}}, func(dst string, opt ClientOption) conn {
			return m
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.DoCache(context.Background(), client.B().Get().Key("a").Cache(), 100).ToString(); err != nil || v != "b" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot asking", func(t *testing.T) {
		count := 0
		m := &mockConn{
			DoFn: func(cmd cmds.Completed) RedisResult {
				if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
					return slotsResp
				}
				return newResult(RedisMessage{typ: '-', string: "ASK 0 :1"}, nil)
			},
			DoMultiFn: func(multi ...cmds.Completed) []RedisResult {
				if count < 3 {
					count++
					return []RedisResult{{}, newResult(RedisMessage{typ: '-', string: "ASK 0 :1"}, nil)}
				}
				return []RedisResult{{}, newResult(RedisMessage{typ: '+', string: "b"}, nil)}
			},
		}
		client, err := newClusterClient(ClientOption{InitAddress: []string{":0"}}, func(dst string, opt ClientOption) conn {
			return m
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.Do(context.Background(), client.B().Get().Key("a").Build()).ToString(); err != nil || v != "b" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot asking (cache)", func(t *testing.T) {
		count := 0
		m := &mockConn{
			DoFn: func(cmd cmds.Completed) RedisResult {
				return slotsResp
			},
			DoCacheFn: func(cmd cmds.Cacheable, ttl time.Duration) RedisResult {
				return newResult(RedisMessage{typ: '-', string: "ASK 0 :1"}, nil)
			},
			DoMultiFn: func(multi ...cmds.Completed) []RedisResult {
				if count < 3 {
					count++
					return []RedisResult{{}, newResult(RedisMessage{typ: '-', string: "ASK 0 :1"}, nil)}
				}
				return []RedisResult{{}, newResult(RedisMessage{typ: '+', string: "b"}, nil)}
			},
		}
		client, err := newClusterClient(ClientOption{InitAddress: []string{":0"}}, func(dst string, opt ClientOption) conn {
			return m
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.DoCache(context.Background(), client.B().Get().Key("a").Cache(), 100).ToString(); err != nil || v != "b" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot try again", func(t *testing.T) {
		count := 0
		m := &mockConn{DoFn: func(cmd cmds.Completed) RedisResult {
			if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
				return slotsResp
			}
			if count < 3 {
				count++
				return newResult(RedisMessage{typ: '-', string: "TRYAGAIN"}, nil)
			}
			return newResult(RedisMessage{typ: '+', string: "b"}, nil)
		}}
		client, err := newClusterClient(ClientOption{InitAddress: []string{":0"}}, func(dst string, opt ClientOption) conn {
			return m
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.Do(context.Background(), client.B().Get().Key("a").Build()).ToString(); err != nil || v != "b" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot try again (cache)", func(t *testing.T) {
		count := 0
		m := &mockConn{
			DoFn: func(cmd cmds.Completed) RedisResult {
				return slotsResp
			},
			DoCacheFn: func(cmd cmds.Cacheable, ttl time.Duration) RedisResult {
				if count < 3 {
					count++
					return newResult(RedisMessage{typ: '-', string: "TRYAGAIN"}, nil)
				}
				return newResult(RedisMessage{typ: '+', string: "b"}, nil)
			},
		}
		client, err := newClusterClient(ClientOption{InitAddress: []string{":0"}}, func(dst string, opt ClientOption) conn {
			return m
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.DoCache(context.Background(), client.B().Get().Key("a").Cache(), 100).ToString(); err != nil || v != "b" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})
}
