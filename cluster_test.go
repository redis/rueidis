package rueidis

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
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

var singleSlotResp2 = newResult(RedisMessage{typ: '*', values: []RedisMessage{
	{typ: '*', values: []RedisMessage{
		{typ: ':', integer: 0},
		{typ: ':', integer: 0},
		{typ: '*', values: []RedisMessage{ // master
			{typ: '+', string: ""},
			{typ: ':', integer: 3},
			{typ: '+', string: ""},
		}},
	}},
}}, nil)

//gocyclo:ignore
func TestClusterClientInit(t *testing.T) {
	t.Run("Init no nodes", func(t *testing.T) {
		if _, err := newClusterClient(&ClientOption{InitAddress: []string{}}, func(dst string, opt *ClientOption) conn { return nil }); err != ErrNoAddr {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("Init no dialable", func(t *testing.T) {
		v := errors.New("dial err")
		if _, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return &mockConn{DialFn: func() error { return v }}
		}); err != v {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("Refresh err", func(t *testing.T) {
		v := errors.New("refresh err")
		if _, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return &mockConn{DoFn: func(cmd cmds.Completed) RedisResult { return newErrResult(v) }}
		}); err != v {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("Refresh retry", func(t *testing.T) {
		var first int64
		if _, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return &mockConn{
				DoFn: func(cmd cmds.Completed) RedisResult {
					if atomic.AddInt64(&first, 1) == 1 {
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
		var first int64
		if _, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return &mockConn{
				DoFn: func(cmd cmds.Completed) RedisResult {
					return newResult(RedisMessage{typ: '*', values: []RedisMessage{}}, nil)
				},
				DialFn: func() error {
					if atomic.AddInt64(&first, 1) == 1 {
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
		var first int64
		client, err := newClusterClient(&ClientOption{InitAddress: []string{":1", ":2"}}, func(dst string, opt *ClientOption) conn {
			return &mockConn{
				DoFn: func(cmd cmds.Completed) RedisResult {
					if atomic.LoadInt64(&first) == 1 {
						return singleSlotResp2
					}
					return slotsResp
				},
			}
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}

		nodes := client.nodes()
		sort.Strings(nodes)
		if len(nodes) != 3 ||
			nodes[0] != ":0" ||
			nodes[1] != ":1" ||
			nodes[2] != ":2" {
			t.Fatalf("unexpected nodes %v", nodes)
		}

		atomic.AddInt64(&first, 1)

		if err = client.refresh(); err != nil {
			t.Fatalf("unexpected err %v", err)
		}

		nodes = client.nodes()
		sort.Strings(nodes)
		if len(nodes) != 3 ||
			nodes[0] != ":1" ||
			nodes[1] != ":2" ||
			nodes[2] != ":3" {
			t.Fatalf("unexpected nodes %v", nodes)
		}
	})
}

//gocyclo:ignore
func TestClusterClient(t *testing.T) {
	m := &mockConn{
		DoFn: func(cmd cmds.Completed) RedisResult {
			if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
				return slotsResp
			}
			return RedisResult{}
		},
		DoMultiFn: func(multi ...cmds.Completed) []RedisResult {
			resps := make([]RedisResult, len(multi))
			for i, cmd := range multi {
				resps[i] = newResult(RedisMessage{typ: '+', string: strings.Join(cmd.Commands(), " ")}, nil)
			}
			return resps
		},
		DoMultiCacheFn: func(multi ...CacheableTTL) []RedisResult {
			resps := make([]RedisResult, len(multi))
			for i, cmd := range multi {
				resps[i] = newResult(RedisMessage{typ: '+', string: strings.Join(cmd.Cmd.Commands(), " ")}, nil)
			}
			return resps
		},
		DoOverride: map[string]func(cmd cmds.Completed) RedisResult{
			"GET Do": func(cmd cmds.Completed) RedisResult {
				return newResult(RedisMessage{typ: '+', string: "Do"}, nil)
			},
			"INFO": func(cmd cmds.Completed) RedisResult {
				return newResult(RedisMessage{typ: '+', string: "Info"}, nil)
			},
		},
		DoCacheOverride: map[string]func(cmd cmds.Cacheable, ttl time.Duration) RedisResult{
			"GET DoCache": func(cmd cmds.Cacheable, ttl time.Duration) RedisResult {
				return newResult(RedisMessage{typ: '+', string: "DoCache"}, nil)
			},
		},
	}

	client, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
		return m
	})
	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}

	t.Run("Nodes", func(t *testing.T) {
		nodes := client.Nodes()
		if len(nodes) != 2 || nodes[":0"] == nil || nodes[":1"] == nil {
			t.Fatalf("unexpected Nodes")
		}
	})

	t.Run("Delegate Do with no slot", func(t *testing.T) {
		c := client.B().Info().Build()
		if v, err := client.Do(context.Background(), c).ToString(); err != nil || v != "Info" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("Delegate Do", func(t *testing.T) {
		c := client.B().Get().Key("Do").Build()
		if v, err := client.Do(context.Background(), c).ToString(); err != nil || v != "Do" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("Delegate DoMulti Empty", func(t *testing.T) {
		if resps := client.DoMulti(context.Background()); resps != nil {
			t.Fatalf("unexpected response %v", resps)
		}
	})

	t.Run("Delegate DoMulti Single Slot", func(t *testing.T) {
		c1 := client.B().Get().Key("K1{a}").Build()
		c2 := client.B().Get().Key("K2{a}").Build()
		resps := client.DoMulti(context.Background(), c1, c2)
		if v, err := resps[0].ToString(); err != nil || v != "GET K1{a}" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
		if v, err := resps[1].ToString(); err != nil || v != "GET K2{a}" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("Delegate DoMulti Single Slot + Init Slot", func(t *testing.T) {
		c1 := client.B().Get().Key("K1{a}").Build()
		c2 := client.B().Info().Build()
		resps := client.DoMulti(context.Background(), c1, c2)
		if v, err := resps[0].ToString(); err != nil || v != "GET K1{a}" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
		if v, err := resps[1].ToString(); err != nil || v != "INFO" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("Delegate DoMulti Cross Slot + Init Slot", func(t *testing.T) {
		defer func() {
			if err := recover(); err != panicMixCxSlot {
				t.Errorf("DoMulti should panic if Cross Slot + Init Slot")
			}
		}()
		c1 := client.B().Get().Key("K1{a}").Build()
		c2 := client.B().Get().Key("K1{b}").Build()
		c3 := client.B().Info().Build()
		client.DoMulti(context.Background(), c1, c2, c3)
	})

	t.Run("Delegate DoMulti Multi Slot", func(t *testing.T) {
		multi := make([]cmds.Completed, 500)
		for i := 0; i < len(multi); i++ {
			multi[i] = client.B().Get().Key(fmt.Sprintf("K1{%d}", i)).Build()
		}
		resps := client.DoMulti(context.Background(), multi...)
		for i := 0; i < len(multi); i++ {
			if v, err := resps[i].ToString(); err != nil || v != fmt.Sprintf("GET K1{%d}", i) {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		}
	})

	t.Run("Delegate DoCache", func(t *testing.T) {
		c := client.B().Get().Key("DoCache").Cache()
		if v, err := client.DoCache(context.Background(), c, 100).ToString(); err != nil || v != "DoCache" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("Delegate DoMultiCache Empty", func(t *testing.T) {
		if resps := client.DoMultiCache(context.Background()); resps != nil {
			t.Fatalf("unexpected response %v", resps)
		}
	})

	t.Run("Delegate DoMultiCache Single Slot", func(t *testing.T) {
		c1 := client.B().Get().Key("K1{a}").Cache()
		c2 := client.B().Get().Key("K2{a}").Cache()
		resps := client.DoMultiCache(context.Background(), CT(c1, time.Second), CT(c2, time.Second))
		if v, err := resps[0].ToString(); err != nil || v != "GET K1{a}" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
		if v, err := resps[1].ToString(); err != nil || v != "GET K2{a}" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("Delegate DoMultiCache Multi Slot", func(t *testing.T) {
		multi := make([]CacheableTTL, 500)
		for i := 0; i < len(multi); i++ {
			multi[i] = CT(client.B().Get().Key(fmt.Sprintf("K1{%d}", i)).Cache(), time.Second)
		}
		resps := client.DoMultiCache(context.Background(), multi...)
		for i := 0; i < len(multi); i++ {
			if v, err := resps[i].ToString(); err != nil || v != fmt.Sprintf("GET K1{%d}", i) {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		}
	})

	t.Run("Delegate Receive", func(t *testing.T) {
		c := client.B().Subscribe().Channel("ch").Build()
		hdl := func(message PubSubMessage) {}
		m.ReceiveFn = func(ctx context.Context, subscribe cmds.Completed, fn func(message PubSubMessage)) error {
			if !reflect.DeepEqual(subscribe.Commands(), c.Commands()) {
				t.Fatalf("unexpected command %v", subscribe)
			}
			return nil
		}
		if err := client.Receive(context.Background(), c, hdl); err != nil {
			t.Fatalf("unexpected response %v", err)
		}
	})

	t.Run("Delegate Receive Redis Err", func(t *testing.T) {
		c := client.B().Subscribe().Channel("ch").Build()
		e := &RedisError{}
		m.ReceiveFn = func(ctx context.Context, subscribe cmds.Completed, fn func(message PubSubMessage)) error {
			return e
		}
		if err := client.Receive(context.Background(), c, func(message PubSubMessage) {}); err != e {
			t.Fatalf("unexpected response %v", err)
		}
	})

	t.Run("Delegate Close", func(t *testing.T) {
		once := sync.Once{}
		called := make(chan struct{})
		m.CloseFn = func() {
			once.Do(func() { close(called) })
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

	t.Run("Dedicated Cross Slot Err Multi", func(t *testing.T) {
		defer func() {
			if err := recover(); err != panicMsgCxSlot {
				t.Errorf("Dedicated should panic if cross slots is used")
			}
		}()
		m.AcquireFn = func() wire {
			return &mockWire{
				DoMultiFn: func(multi ...cmds.Completed) []RedisResult {
					return []RedisResult{
						newResult(RedisMessage{typ: '+', string: "OK"}, nil),
						newResult(RedisMessage{typ: '+', string: "OK"}, nil),
						newResult(RedisMessage{typ: '*', values: []RedisMessage{{typ: '+', string: "a"}}}, nil),
					}
				},
			}
		}
		client.Dedicated(func(c DedicatedClient) (err error) {
			c.DoMulti(
				context.Background(),
				c.B().Multi().Build(),
				c.B().Get().Key("a").Build(),
				c.B().Exec().Build(),
			)
			c.DoMulti(
				context.Background(),
				c.B().Multi().Build(),
				c.B().Get().Key("b").Build(),
				c.B().Exec().Build(),
			)
			return nil
		})
	})

	t.Run("Dedicated Multi Cross Slot Err", func(t *testing.T) {
		m.AcquireFn = func() wire { return &mockWire{} }
		err := client.Dedicated(func(c DedicatedClient) (err error) {
			defer func() {
				err = errors.New(recover().(string))
			}()
			c.DoMulti(
				context.Background(),
				c.B().Get().Key("a").Build(),
				c.B().Get().Key("b").Build(),
			)
			return nil
		})
		if err == nil || err.Error() != panicMsgCxSlot {
			t.Errorf("Multi should panic if cross slots is used")
		}
	})

	t.Run("Dedicated Delegate Receive Redis Err", func(t *testing.T) {
		e := &RedisError{}
		w := &mockWire{
			ReceiveFn: func(ctx context.Context, subscribe cmds.Completed, fn func(message PubSubMessage)) error {
				return e
			},
		}
		m.AcquireFn = func() wire {
			return w
		}
		if err := client.Dedicated(func(c DedicatedClient) error {
			return c.Receive(context.Background(), c.B().Subscribe().Channel("a").Build(), func(msg PubSubMessage) {})
		}); err != e {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("Dedicated Delegate", func(t *testing.T) {
		closed := false
		w := &mockWire{
			DoFn: func(cmd cmds.Completed) RedisResult {
				return newResult(RedisMessage{typ: '+', string: "Delegate"}, nil)
			},
			DoMultiFn: func(cmd ...cmds.Completed) []RedisResult {
				if len(cmd) == 4 {
					return []RedisResult{
						newResult(RedisMessage{typ: '+', string: "OK"}, nil),
						newResult(RedisMessage{typ: '+', string: "OK"}, nil),
						newResult(RedisMessage{typ: '+', string: "OK"}, nil),
						newResult(RedisMessage{typ: '*', values: []RedisMessage{
							{typ: '+', string: "Delegate0"},
							{typ: '+', string: "Delegate1"},
						}}, nil),
					}
				}
				return []RedisResult{
					newResult(RedisMessage{typ: '+', string: "Delegate0"}, nil),
					newResult(RedisMessage{typ: '+', string: "Delegate1"}, nil),
				}
			},
			ReceiveFn: func(ctx context.Context, subscribe cmds.Completed, fn func(message PubSubMessage)) error {
				return ErrClosing
			},
			SetPubSubHooksFn: func(hooks PubSubHooks) <-chan error {
				ch := make(chan error, 1)
				ch <- ErrClosing
				close(ch)
				return ch
			},
			ErrorFn: func() error {
				return ErrClosing
			},
			CloseFn: func() {
				closed = true
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
			ch := c.SetPubSubHooks(PubSubHooks{OnMessage: func(m PubSubMessage) {}})
			if v, err := c.Do(context.Background(), c.B().Get().Key("a").Build()).ToString(); err != nil || v != "Delegate" {
				t.Fatalf("unexpected response %v %v", v, err)
			}
			if v := c.DoMulti(context.Background()); len(v) != 0 {
				t.Fatalf("received unexpected response %v", v)
			}
			for i, resp := range c.DoMulti(
				context.Background(),
				c.B().Get().Key("a").Build(),
				c.B().Get().Key("a").Build(),
			) {
				if v, err := resp.ToString(); err != nil || v != "Delegate"+strconv.Itoa(i) {
					t.Fatalf("unexpected response %v %v", v, err)
				}
			}
			for i, resp := range c.DoMulti(
				context.Background(),
				c.B().Multi().Build(),
				c.B().Get().Key("a").Build(),
				c.B().Get().Key("a").Build(),
				c.B().Exec().Build(),
			)[3].val.values {
				if v, err := resp.ToString(); err != nil || v != "Delegate"+strconv.Itoa(i) {
					t.Fatalf("unexpected response %v %v", v, err)
				}
			}
			if err := c.Receive(context.Background(), c.B().Ssubscribe().Channel("a").Build(), func(msg PubSubMessage) {}); err != ErrClosing {
				t.Fatalf("unexpected ret %v", err)
			}
			if err := <-ch; err != ErrClosing {
				t.Fatalf("unexpected ret %v", err)
			}
			if err := <-c.SetPubSubHooks(PubSubHooks{OnMessage: func(m PubSubMessage) {}}); err != ErrClosing {
				t.Fatalf("unexpected ret %v", err)
			}
			c.Close()
			return nil
		}); err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if !stored {
			t.Fatalf("Dedicated desn't put back the wire")
		}
		if !closed {
			t.Fatalf("Dedicated desn't delegate Close")
		}
	})

	t.Run("Dedicated Delegate", func(t *testing.T) {
		closed := false
		w := &mockWire{
			DoFn: func(cmd cmds.Completed) RedisResult {
				return newResult(RedisMessage{typ: '+', string: "Delegate"}, nil)
			},
			DoMultiFn: func(cmd ...cmds.Completed) []RedisResult {
				if len(cmd) == 4 {
					return []RedisResult{
						newResult(RedisMessage{typ: '+', string: "OK"}, nil),
						newResult(RedisMessage{typ: '+', string: "OK"}, nil),
						newResult(RedisMessage{typ: '+', string: "OK"}, nil),
						newResult(RedisMessage{typ: '*', values: []RedisMessage{
							{typ: '+', string: "Delegate0"},
							{typ: '+', string: "Delegate1"},
						}}, nil),
					}
				}
				return []RedisResult{
					newResult(RedisMessage{typ: '+', string: "Delegate0"}, nil),
					newResult(RedisMessage{typ: '+', string: "Delegate1"}, nil),
				}
			},
			ReceiveFn: func(ctx context.Context, subscribe cmds.Completed, fn func(message PubSubMessage)) error {
				return ErrClosing
			},
			SetPubSubHooksFn: func(hooks PubSubHooks) <-chan error {
				ch := make(chan error, 1)
				ch <- ErrClosing
				close(ch)
				return ch
			},
			ErrorFn: func() error {
				return ErrClosing
			},
			CloseFn: func() {
				closed = true
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
		c, cancel := client.Dedicate()
		ch := c.SetPubSubHooks(PubSubHooks{OnMessage: func(m PubSubMessage) {}})
		if v, err := c.Do(context.Background(), c.B().Get().Key("a").Build()).ToString(); err != nil || v != "Delegate" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
		if v := c.DoMulti(context.Background()); len(v) != 0 {
			t.Fatalf("received unexpected response %v", v)
		}
		for i, resp := range c.DoMulti(
			context.Background(),
			c.B().Get().Key("a").Build(),
			c.B().Get().Key("a").Build(),
		) {
			if v, err := resp.ToString(); err != nil || v != "Delegate"+strconv.Itoa(i) {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		}
		for i, resp := range c.DoMulti(
			context.Background(),
			c.B().Multi().Build(),
			c.B().Get().Key("a").Build(),
			c.B().Get().Key("a").Build(),
			c.B().Exec().Build(),
		)[3].val.values {
			if v, err := resp.ToString(); err != nil || v != "Delegate"+strconv.Itoa(i) {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		}
		if err := c.Receive(context.Background(), c.B().Ssubscribe().Channel("a").Build(), func(msg PubSubMessage) {}); err != ErrClosing {
			t.Fatalf("unexpected ret %v", err)
		}
		if err := <-ch; err != ErrClosing {
			t.Fatalf("unexpected ret %v", err)
		}
		if err := <-c.SetPubSubHooks(PubSubHooks{OnMessage: func(m PubSubMessage) {}}); err != ErrClosing {
			t.Fatalf("unexpected ret %v", err)
		}
		c.Close()
		cancel()

		if !stored {
			t.Fatalf("Dedicated desn't put back the wire")
		}
		if !closed {
			t.Fatalf("Dedicated desn't delegate Close")
		}
	})

	t.Run("Dedicate Delegate Release On Close", func(t *testing.T) {
		stored := 0
		w := &mockWire{}
		m.AcquireFn = func() wire { return w }
		m.StoreFn = func(ww wire) { stored++ }
		c, _ := client.Dedicate()
		c.Do(context.Background(), c.B().Get().Key("a").Build())

		c.Close()

		if stored != 1 {
			t.Fatalf("unexpected stored count %v", stored)
		}
	})

	t.Run("Dedicate Delegate No Duplicate Release", func(t *testing.T) {
		stored := 0
		w := &mockWire{}
		m.AcquireFn = func() wire { return w }
		m.StoreFn = func(ww wire) { stored++ }
		c, cancel := client.Dedicate()
		c.Do(context.Background(), c.B().Get().Key("a").Build())

		c.Close()
		c.Close() // should have no effect
		cancel()  // should have no effect
		cancel()  // should have no effect

		if stored != 1 {
			t.Fatalf("unexpected stored count %v", stored)
		}
	})

	t.Run("Dedicated SetPubSubHooks Released", func(t *testing.T) {
		c, cancel := client.Dedicate()
		ch1 := c.SetPubSubHooks(PubSubHooks{OnMessage: func(m PubSubMessage) {}})
		ch2 := c.SetPubSubHooks(PubSubHooks{OnMessage: func(m PubSubMessage) {}})
		<-ch1
		cancel()
		<-ch2
	})

	t.Run("Dedicated SetPubSubHooks Close", func(t *testing.T) {
		c, cancel := client.Dedicate()
		defer cancel()
		ch := c.SetPubSubHooks(PubSubHooks{OnMessage: func(m PubSubMessage) {}})
		c.Close()
		if err := <-ch; err != ErrClosing {
			t.Fatalf("unexpected ret %v", ch)
		}
	})

	t.Run("Dedicated SetPubSubHooks Released", func(t *testing.T) {
		c, cancel := client.Dedicate()
		defer cancel()
		if ch := c.SetPubSubHooks(PubSubHooks{}); ch != nil {
			t.Fatalf("unexpected ret %v", ch)
		}
	})
}

//gocyclo:ignore
func TestClusterClientErr(t *testing.T) {
	t.Run("refresh err on pick", func(t *testing.T) {
		var first int64
		v := errors.New("refresh err")
		m := &mockConn{
			DoFn: func(cmd cmds.Completed) RedisResult {
				if atomic.AddInt64(&first, 1) == 1 {
					return singleSlotResp
				}
				return newErrResult(v)
			},
			ReceiveFn: func(ctx context.Context, subscribe cmds.Completed, fn func(message PubSubMessage)) error {
				return v
			},
		}
		client, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return m
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if err := client.Do(context.Background(), client.B().Get().Key("a").Build()).Error(); err != v {
			t.Fatalf("unexpected err %v", err)
		}
		if err := client.DoMulti(context.Background(), client.B().Get().Key("a").Build())[0].Error(); err != v {
			t.Fatalf("unexpected err %v", err)
		}
		if err := client.DoCache(context.Background(), client.B().Get().Key("a").Cache(), 100).Error(); err != v {
			t.Fatalf("unexpected err %v", err)
		}
		if err := client.DoMultiCache(context.Background(), CT(client.B().Get().Key("a").Cache(), 100))[0].Error(); err != v {
			t.Fatalf("unexpected err %v", err)
		}
		if err := client.Receive(context.Background(), client.B().Ssubscribe().Channel("a").Build(), func(msg PubSubMessage) {}); err != v {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("refresh empty on pick", func(t *testing.T) {
		m := &mockConn{DoFn: func(cmd cmds.Completed) RedisResult {
			return singleSlotResp
		}}
		client, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return m
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if err := client.Do(context.Background(), client.B().Get().Key("a").Build()).Error(); err != ErrNoSlot {
			t.Fatalf("unexpected err %v", err)
		}
		if err := client.DoMulti(context.Background(), client.B().Get().Key("a").Build())[0].Error(); err != ErrNoSlot {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("refresh empty on pick in dedicated wire", func(t *testing.T) {
		m := &mockConn{DoFn: func(cmd cmds.Completed) RedisResult {
			return singleSlotResp
		}}
		client, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return m
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		var ch <-chan error
		if err := client.Dedicated(func(c DedicatedClient) error {
			ch = c.SetPubSubHooks(PubSubHooks{OnMessage: func(m PubSubMessage) {}})
			return c.Do(context.Background(), c.B().Get().Key("a").Build()).Error()
		}); err != ErrNoSlot {
			t.Fatalf("unexpected err %v", err)
		}
		if err := <-ch; err != ErrNoSlot {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("refresh empty on pick in dedicated wire (multi)", func(t *testing.T) {
		m := &mockConn{DoFn: func(cmd cmds.Completed) RedisResult {
			return singleSlotResp
		}}
		client, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return m
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		var ch <-chan error
		if err := client.Dedicated(func(c DedicatedClient) error {
			ch = c.SetPubSubHooks(PubSubHooks{OnMessage: func(m PubSubMessage) {}})
			for _, v := range c.DoMulti(context.Background(), c.B().Get().Key("a").Build()) {
				if err := v.Error(); err != nil {
					return err
				}
			}
			return nil
		}); err != ErrNoSlot {
			t.Fatalf("unexpected err %v", err)
		}
		if err := <-ch; err != ErrNoSlot {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("slot reconnect", func(t *testing.T) {
		var count, check int64
		m := &mockConn{DoFn: func(cmd cmds.Completed) RedisResult {
			if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
				return slotsResp
			}
			if atomic.AddInt64(&count, 1) <= 3 {
				return newResult(RedisMessage{typ: '-', string: "MOVED 0 :0"}, nil)
			}
			return newResult(RedisMessage{typ: '+', string: "b"}, nil)
		}, AddrFn: func() string {
			atomic.AddInt64(&check, 1)
			return ":0"
		}}

		client, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return m
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.Do(context.Background(), client.B().Get().Key("a").Build()).ToString(); err != nil || v != "b" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if atomic.LoadInt64(&check) != 6 {
			t.Fatalf("unexpected check count %v", check)
		}
	})

	t.Run("slot moved", func(t *testing.T) {
		var count int64
		m := &mockConn{DoFn: func(cmd cmds.Completed) RedisResult {
			if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
				return slotsResp
			}
			if atomic.AddInt64(&count, 1) <= 3 {
				return newResult(RedisMessage{typ: '-', string: "MOVED 0 :1"}, nil)
			}
			return newResult(RedisMessage{typ: '+', string: "b"}, nil)
		}}
		client, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return m
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.Do(context.Background(), client.B().Get().Key("a").Build()).ToString(); err != nil || v != "b" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot moved DoMulti", func(t *testing.T) {
		var count int64
		m := &mockConn{DoFn: func(cmd cmds.Completed) RedisResult {
			return slotsResp
		}, DoMultiFn: func(multi ...cmds.Completed) []RedisResult {
			if atomic.AddInt64(&count, 1) <= 3 {
				return []RedisResult{newResult(RedisMessage{typ: '-', string: "MOVED 0 :1"}, nil)}
			}
			ret := make([]RedisResult, len(multi))
			ret[0] = newResult(RedisMessage{typ: '+', string: "b"}, nil)
			return ret
		}}
		client, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return m
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.DoMulti(context.Background(), client.B().Get().Key("a").Build())[0].ToString(); err != nil || v != "b" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot moved new", func(t *testing.T) {
		var count, check int64
		m := &mockConn{DoFn: func(cmd cmds.Completed) RedisResult {
			if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
				return slotsResp
			}
			if atomic.AddInt64(&count, 1) <= 3 {
				return newResult(RedisMessage{typ: '-', string: "MOVED 0 :2"}, nil)
			}
			return newResult(RedisMessage{typ: '+', string: "b"}, nil)
		}}

		client, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			if dst == ":2" {
				atomic.AddInt64(&check, 1)
			}
			return m
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.Do(context.Background(), client.B().Get().Key("a").Build()).ToString(); err != nil || v != "b" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if atomic.LoadInt64(&check) == 0 {
			t.Fatalf("unexpected check value %v", check)
		}
	})

	t.Run("slot moved new (multi)", func(t *testing.T) {
		var count, check int64
		m := &mockConn{DoFn: func(cmd cmds.Completed) RedisResult {
			return slotsResp
		}, DoMultiFn: func(multi ...cmds.Completed) []RedisResult {
			if atomic.AddInt64(&count, 1) <= 3 {
				return []RedisResult{newResult(RedisMessage{typ: '-', string: "MOVED 0 :2"}, nil)}
			}
			ret := make([]RedisResult, len(multi))
			ret[0] = newResult(RedisMessage{typ: '+', string: "b"}, nil)
			return ret
		}}

		client, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			if dst == ":2" {
				atomic.AddInt64(&check, 1)
			}
			return m
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.DoMulti(context.Background(), client.B().Get().Key("a").Build())[0].ToString(); err != nil || v != "b" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if atomic.LoadInt64(&check) == 0 {
			t.Fatalf("unexpected check value %v", check)
		}
	})

	t.Run("slot moved (cache)", func(t *testing.T) {
		var count int64
		m := &mockConn{
			DoFn: func(cmd cmds.Completed) RedisResult {
				return slotsResp
			},
			DoCacheFn: func(cmd cmds.Cacheable, ttl time.Duration) RedisResult {
				if atomic.AddInt64(&count, 1) <= 3 {
					return newResult(RedisMessage{typ: '-', string: "MOVED 0 :1"}, nil)
				}
				return newResult(RedisMessage{typ: '+', string: "b"}, nil)
			},
		}
		client, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return m
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.DoCache(context.Background(), client.B().Get().Key("a").Cache(), 100).ToString(); err != nil || v != "b" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot moved (cache multi)", func(t *testing.T) {
		var count int64
		m := &mockConn{DoFn: func(cmd cmds.Completed) RedisResult {
			return slotsResp
		}, DoMultiCacheFn: func(multi ...CacheableTTL) []RedisResult {
			if atomic.AddInt64(&count, 1) <= 3 {
				return []RedisResult{newResult(RedisMessage{typ: '-', string: "MOVED 0 :1"}, nil)}
			}
			return []RedisResult{newResult(RedisMessage{typ: '+', string: "b"}, nil)}
		}}
		client, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return m
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.DoMultiCache(context.Background(), CT(client.B().Get().Key("a").Cache(), 100))[0].ToString(); err != nil || v != "b" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot asking", func(t *testing.T) {
		var count int64
		m := &mockConn{
			DoFn: func(cmd cmds.Completed) RedisResult {
				if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
					return slotsResp
				}
				return newResult(RedisMessage{typ: '-', string: "ASK 0 :1"}, nil)
			},
			DoMultiFn: func(multi ...cmds.Completed) []RedisResult {
				if atomic.AddInt64(&count, 1) <= 3 {
					return []RedisResult{{}, newResult(RedisMessage{typ: '-', string: "ASK 0 :1"}, nil)}
				}
				return []RedisResult{{}, newResult(RedisMessage{typ: '+', string: "b"}, nil)}
			},
		}
		client, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return m
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.Do(context.Background(), client.B().Get().Key("a").Build()).ToString(); err != nil || v != "b" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot asking DoMulti", func(t *testing.T) {
		var count int64
		m := &mockConn{
			DoFn: func(cmd cmds.Completed) RedisResult {
				return slotsResp
			},
			DoMultiFn: func(multi ...cmds.Completed) []RedisResult {
				if atomic.AddInt64(&count, 1) <= 3 {
					return []RedisResult{{}, newResult(RedisMessage{typ: '-', string: "ASK 0 :1"}, nil)}
				}
				ret := make([]RedisResult, len(multi))
				ret[0] = newResult(RedisMessage{typ: '+', string: "OK"}, nil)
				ret[1] = newResult(RedisMessage{typ: '+', string: "b"}, nil)
				return ret
			},
		}
		client, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return m
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.DoMulti(context.Background(), client.B().Get().Key("a").Build())[0].ToString(); err != nil || v != "b" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot asking (cache)", func(t *testing.T) {
		var count int64
		m := &mockConn{
			DoFn: func(cmd cmds.Completed) RedisResult {
				return slotsResp
			},
			DoCacheFn: func(cmd cmds.Cacheable, ttl time.Duration) RedisResult {
				return newResult(RedisMessage{typ: '-', string: "ASK 0 :1"}, nil)
			},
			DoMultiFn: func(multi ...cmds.Completed) []RedisResult {
				if atomic.AddInt64(&count, 1) <= 3 {
					return []RedisResult{{}, newResult(RedisMessage{typ: '-', string: "ASK 0 :1"}, nil)}
				}
				return []RedisResult{{}, newResult(RedisMessage{typ: '+', string: "b"}, nil)}
			},
		}
		client, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return m
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.DoCache(context.Background(), client.B().Get().Key("a").Cache(), 100).ToString(); err != nil || v != "b" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot asking (cache multi)", func(t *testing.T) {
		var count int64
		m := &mockConn{
			DoFn: func(cmd cmds.Completed) RedisResult {
				return slotsResp
			},
			DoMultiCacheFn: func(multi ...CacheableTTL) []RedisResult {
				return []RedisResult{newResult(RedisMessage{typ: '-', string: "ASK 0 :1"}, nil)}
			},
			DoMultiFn: func(multi ...cmds.Completed) []RedisResult {
				if atomic.AddInt64(&count, 1) <= 3 {
					return []RedisResult{{}, {}, {}, newResult(RedisMessage{typ: '-', string: "ASK 0 :1"}, nil)}
				}
				ret := make([]RedisResult, len(multi))
				for i := 0; i < len(multi)-1; i++ {
					ret[i] = newResult(RedisMessage{typ: '+', string: "QUEUED"}, nil)
				}
				ret[len(multi)-1] = newResult(RedisMessage{typ: '*', values: []RedisMessage{{typ: '+', string: "b"}}}, nil)
				return ret
			},
		}
		client, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return m
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.DoMultiCache(context.Background(), CT(client.B().Get().Key("a").Cache(), 100))[0].ToString(); err != nil || v != "b" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot try again", func(t *testing.T) {
		var count int64
		m := &mockConn{DoFn: func(cmd cmds.Completed) RedisResult {
			if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
				return slotsResp
			}
			if atomic.AddInt64(&count, 1) <= 3 {
				return newResult(RedisMessage{typ: '-', string: "TRYAGAIN"}, nil)
			}
			return newResult(RedisMessage{typ: '+', string: "b"}, nil)
		}}
		client, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return m
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.Do(context.Background(), client.B().Get().Key("a").Build()).ToString(); err != nil || v != "b" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot try again DoMulti", func(t *testing.T) {
		var count int64
		m := &mockConn{DoFn: func(cmd cmds.Completed) RedisResult {
			return slotsResp
		}, DoMultiFn: func(multi ...cmds.Completed) []RedisResult {
			if atomic.AddInt64(&count, 1) <= 3 {
				return []RedisResult{newResult(RedisMessage{typ: '-', string: "TRYAGAIN"}, nil)}
			}
			ret := make([]RedisResult, len(multi))
			ret[0] = newResult(RedisMessage{typ: '+', string: "b"}, nil)
			return ret
		}}
		client, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return m
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.DoMulti(context.Background(), client.B().Get().Key("a").Build())[0].ToString(); err != nil || v != "b" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot try again (cache)", func(t *testing.T) {
		var count int64
		m := &mockConn{
			DoFn: func(cmd cmds.Completed) RedisResult {
				return slotsResp
			},
			DoCacheFn: func(cmd cmds.Cacheable, ttl time.Duration) RedisResult {
				if atomic.AddInt64(&count, 1) <= 3 {
					return newResult(RedisMessage{typ: '-', string: "TRYAGAIN"}, nil)
				}
				return newResult(RedisMessage{typ: '+', string: "b"}, nil)
			},
		}
		client, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return m
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.DoCache(context.Background(), client.B().Get().Key("a").Cache(), 100).ToString(); err != nil || v != "b" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot try again (cache multi)", func(t *testing.T) {
		var count int64
		m := &mockConn{DoFn: func(cmd cmds.Completed) RedisResult {
			return slotsResp
		}, DoMultiCacheFn: func(multi ...CacheableTTL) []RedisResult {
			if atomic.AddInt64(&count, 1) <= 3 {
				return []RedisResult{newResult(RedisMessage{typ: '-', string: "TRYAGAIN"}, nil)}
			}
			return []RedisResult{newResult(RedisMessage{typ: '+', string: "b"}, nil)}
		}}
		client, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return m
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.DoMultiCache(context.Background(), CT(client.B().Get().Key("a").Cache(), 100))[0].ToString(); err != nil || v != "b" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})
}

func TestClusterClientRetry(t *testing.T) {
	SetupClientRetry(t, func(m *mockConn) Client {
		m.DoOverride = map[string]func(cmd cmds.Completed) RedisResult{
			"CLUSTER SLOTS": func(cmd cmds.Completed) RedisResult { return slotsResp },
		}
		c, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return m
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		return c
	})
}
