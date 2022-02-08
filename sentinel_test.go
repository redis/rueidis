package rueidis

import (
	"context"
	"errors"
	"github.com/rueian/rueidis/internal/cmds"
	"reflect"
	"sync/atomic"
	"testing"
	"time"
)

//gocyclo:ignore
func TestSentinelClientInit(t *testing.T) {
	t.Run("Init no nodes", func(t *testing.T) {
		if _, err := newSentinelClient(&ClientOption{InitAddress: []string{}}, func(dst string, opt *ClientOption) conn { return nil }); err != ErrNoAddr {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("Init no dialable", func(t *testing.T) {
		v := errors.New("dial err")
		if _, err := newSentinelClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return &mockConn{DialFn: func() error { return v }}
		}); err != v {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("Refresh err", func(t *testing.T) {
		v := errors.New("refresh err")
		if _, err := newSentinelClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return &mockConn{
				DoMultiFn: func(cmd ...cmds.Completed) []RedisResult { return []RedisResult{newErrResult(v)} },
			}
		}); err != v {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("Refresh retry", func(t *testing.T) {
		v := errors.New("refresh err")
		s0 := &mockConn{
			DoFn: func(cmd cmds.Completed) RedisResult { return RedisResult{} },
			DoMultiFn: func(multi ...cmds.Completed) []RedisResult {
				return []RedisResult{
					newErrResult(v),
					newErrResult(v),
				}
			},
		}
		s1 := &mockConn{
			DoFn: func(cmd cmds.Completed) RedisResult { return RedisResult{} },
			DoMultiFn: func(multi ...cmds.Completed) []RedisResult {
				return []RedisResult{
					{val: RedisMessage{typ: '*', values: []RedisMessage{
						{typ: '%', values: []RedisMessage{
							{typ: '+', string: "ip"}, {typ: '+', string: ""},
							{typ: '+', string: "port"}, {typ: '+', string: "0"},
						}},
					}}},
					newErrResult(v),
				}
			},
		}
		s2 := &mockConn{
			DoFn: func(cmd cmds.Completed) RedisResult { return RedisResult{} },
			DoMultiFn: func(multi ...cmds.Completed) []RedisResult {
				return []RedisResult{
					{val: RedisMessage{typ: '*', values: []RedisMessage{
						{typ: '%', values: []RedisMessage{
							{typ: '+', string: "ip"}, {typ: '+', string: ""},
							{typ: '+', string: "port"}, {typ: '+', string: "3"},
						}},
					}}},
					{val: RedisMessage{typ: '*', values: []RedisMessage{
						{typ: '+', string: ""}, {typ: '+', string: "5"},
					}}},
				}
			},
		}
		s3 := &mockConn{
			DoFn: func(cmd cmds.Completed) RedisResult { return RedisResult{} },
			DoMultiFn: func(multi ...cmds.Completed) []RedisResult {
				return []RedisResult{
					{val: RedisMessage{typ: '*', values: []RedisMessage{
						{typ: '%', values: []RedisMessage{
							{typ: '+', string: "ip"}, {typ: '+', string: ""},
							{typ: '+', string: "port"}, {typ: '+', string: "4"},
						}},
					}}},
					{val: RedisMessage{typ: '*', values: []RedisMessage{
						{typ: '+', string: ""}, {typ: '+', string: "6"},
					}}},
				}
			},
		}
		s4 := &mockConn{
			DoFn: func(cmd cmds.Completed) RedisResult { return RedisResult{} },
			DoMultiFn: func(multi ...cmds.Completed) []RedisResult {
				return []RedisResult{
					{val: RedisMessage{typ: '*', values: []RedisMessage{
						{typ: '%', values: []RedisMessage{
							{typ: '+', string: "ip"}, {typ: '+', string: ""},
							{typ: '+', string: "port"}, {typ: '+', string: "2"},
						}},
					}}},
					{val: RedisMessage{typ: '*', values: []RedisMessage{
						{typ: '+', string: ""}, {typ: '+', string: "7"},
					}}},
				}
			},
		}
		client, err := newSentinelClient(&ClientOption{InitAddress: []string{":0", ":1", ":2"}}, func(dst string, opt *ClientOption) conn {
			if dst == ":0" {
				return s0
			}
			if dst == ":1" {
				return s1
			}
			if dst == ":2" {
				return s2
			}
			if dst == ":3" {
				return s3
			}
			if dst == ":4" {
				return s4
			}
			if dst == ":5" {
				return &mockConn{
					DialFn: func() error { return v },
				}
			}
			if dst == ":6" {
				return &mockConn{
					DoFn: func(cmd cmds.Completed) RedisResult {
						return RedisResult{val: RedisMessage{typ: '*', values: []RedisMessage{{typ: '+', string: "slave"}}}}
					},
				}
			}
			if dst == ":7" {
				return &mockConn{
					DoFn: func(cmd cmds.Completed) RedisResult {
						return RedisResult{val: RedisMessage{typ: '*', values: []RedisMessage{{typ: '+', string: "master"}}}}
					},
				}
			}
			return nil
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if client.sConn == nil {
			t.Fatalf("unexpected nil sentinel conn")
		}
		if client.mConn.Load() == nil {
			t.Fatalf("unexpected nil master conn")
		}
		client.Close()
	})

	t.Run("Refresh retry 2", func(t *testing.T) {
		v := errors.New("refresh err")
		s0 := &mockConn{
			DoFn: func(cmd cmds.Completed) RedisResult { return RedisResult{} },
			DoMultiFn: func(multi ...cmds.Completed) []RedisResult {
				return []RedisResult{
					{val: RedisMessage{typ: '*', values: []RedisMessage{
						{typ: '%', values: []RedisMessage{
							{typ: '+', string: "ip"}, {typ: '+', string: ""},
							{typ: '+', string: "port"}, {typ: '+', string: "1"},
						}},
					}}},
					{val: RedisMessage{typ: '*', values: []RedisMessage{
						{typ: '+', string: ""}, {typ: '+', string: "2"},
					}}},
				}
			},
		}
		s1 := &mockConn{
			DoFn: func(cmd cmds.Completed) RedisResult { return RedisResult{} },
			DoMultiFn: func(multi ...cmds.Completed) []RedisResult {
				return []RedisResult{
					{val: RedisMessage{typ: '*', values: []RedisMessage{
						{typ: '%', values: []RedisMessage{
							{typ: '+', string: "ip"}, {typ: '+', string: ""},
							{typ: '+', string: "port"}, {typ: '+', string: "0"},
						}},
					}}},
					{val: RedisMessage{typ: '*', values: []RedisMessage{
						{typ: '+', string: ""}, {typ: '+', string: "3"},
					}}},
				}
			},
		}
		client, err := newSentinelClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			if dst == ":0" {
				return s0
			}
			if dst == ":1" {
				return s1
			}
			if dst == ":2" {
				return &mockConn{
					DoFn: func(cmd cmds.Completed) RedisResult { return newErrResult(v) },
				}
			}
			if dst == ":3" {
				return &mockConn{
					DoFn: func(cmd cmds.Completed) RedisResult {
						return RedisResult{val: RedisMessage{typ: '*', values: []RedisMessage{{typ: '+', string: "master"}}}}
					},
				}
			}
			return nil
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if client.sConn == nil {
			t.Fatalf("unexpected nil sentinel conn")
		}
		if client.mConn.Load() == nil {
			t.Fatalf("unexpected nil master conn")
		}
		client.Close()
	})

	t.Run("sentinel disconnect", func(t *testing.T) {
		trigger := make(chan error)
		disconnect := int32(0)
		s0closed := int32(0)
		r3closed := int32(0)
		s0 := &mockConn{
			DoFn: func(cmd cmds.Completed) RedisResult {
				if atomic.LoadInt32(&disconnect) == 1 {
					return newErrResult(errors.New("die"))
				}
				return RedisResult{}
			},
			DoMultiFn: func(multi ...cmds.Completed) []RedisResult {
				return []RedisResult{
					{val: RedisMessage{typ: '*', values: []RedisMessage{
						{typ: '%', values: []RedisMessage{
							{typ: '+', string: "ip"}, {typ: '+', string: ""},
							{typ: '+', string: "port"}, {typ: '+', string: "1"},
						}},
					}}},
					{val: RedisMessage{typ: '*', values: []RedisMessage{
						{typ: '+', string: ""}, {typ: '+', string: "3"},
					}}},
				}
			},
			ReceiveFn: func(ctx context.Context, subscribe cmds.Completed, fn func(message PubSubMessage)) error {
				if err, ok := <-trigger; ok {
					return err
				}
				return ErrClosing
			},
			CloseFn: func() {
				atomic.StoreInt32(&s0closed, 1)
			},
		}
		s1 := &mockConn{
			DoFn: func(cmd cmds.Completed) RedisResult { return RedisResult{} },
			DoMultiFn: func(multi ...cmds.Completed) []RedisResult {
				return []RedisResult{
					{val: RedisMessage{typ: '*', values: []RedisMessage{
						{typ: '%', values: []RedisMessage{
							{typ: '+', string: "ip"}, {typ: '+', string: ""},
							{typ: '+', string: "port"}, {typ: '+', string: "2"},
						}},
					}}},
					{val: RedisMessage{typ: '*', values: []RedisMessage{
						{typ: '+', string: ""}, {typ: '+', string: "3"},
					}}},
				}
			},
		}
		s2 := &mockConn{
			DoFn: func(cmd cmds.Completed) RedisResult { return RedisResult{} },
			DoMultiFn: func(multi ...cmds.Completed) []RedisResult {
				return []RedisResult{
					{val: RedisMessage{typ: '*', values: []RedisMessage{
						{typ: '%', values: []RedisMessage{
							{typ: '+', string: "ip"}, {typ: '+', string: ""},
							{typ: '+', string: "port"}, {typ: '+', string: "0"},
						}},
					}}},
					{val: RedisMessage{typ: '*', values: []RedisMessage{
						{typ: '+', string: ""}, {typ: '+', string: "4"},
					}}},
				}
			},
		}
		r3 := &mockConn{
			DoFn: func(cmd cmds.Completed) RedisResult {
				if atomic.LoadInt32(&disconnect) == 1 {
					return newErrResult(errors.New("die"))
				}
				return RedisResult{val: RedisMessage{typ: '*', values: []RedisMessage{{typ: '+', string: "master"}}}}
			},
			CloseFn: func() {
				atomic.StoreInt32(&r3closed, 1)
			},
			ErrorFn: func() error {
				if atomic.LoadInt32(&disconnect) == 1 {
					return errClosing
				}
				return nil
			},
		}
		r4 := &mockConn{
			DoFn: func(cmd cmds.Completed) RedisResult {
				return RedisResult{val: RedisMessage{typ: '*', values: []RedisMessage{{typ: '+', string: "master"}}}}
			},
		}
		client, err := newSentinelClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			if dst == ":0" {
				return s0
			}
			if dst == ":1" {
				return s1
			}
			if dst == ":2" {
				return s2
			}
			if dst == ":3" {
				return r3
			}
			if dst == ":4" {
				return r4
			}
			return nil
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		atomic.StoreInt32(&disconnect, 1)
		trigger <- errors.New("reconnect")
		close(trigger)
		for {
			t.Log("wait switch master")
			if client.mConn.Load().(*mockConn) == r4 {
				break
			}
		}
		if atomic.LoadInt32(&s0closed) != 1 {
			t.Fatalf("s0 not closed")
		}
		if atomic.LoadInt32(&r3closed) != 1 {
			t.Fatalf("r3 not closed")
		}
		client.Close()
	})
}

//gocyclo:ignore
func TestSentinelClientDelegate(t *testing.T) {
	s0 := &mockConn{
		DoFn: func(cmd cmds.Completed) RedisResult { return RedisResult{} },
		DoMultiFn: func(multi ...cmds.Completed) []RedisResult {
			return []RedisResult{
				{val: RedisMessage{typ: '*', values: []RedisMessage{}}},
				{val: RedisMessage{typ: '*', values: []RedisMessage{
					{typ: '+', string: ""}, {typ: '+', string: "1"},
				}}},
			}
		},
	}
	m := &mockConn{
		DoFn: func(cmd cmds.Completed) RedisResult {
			return RedisResult{val: RedisMessage{typ: '*', values: []RedisMessage{{typ: '+', string: "master"}}}}
		},
	}
	client, err := newSentinelClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
		if dst == ":0" {
			return s0
		}
		if dst == ":1" {
			return m
		}
		return nil
	})
	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}
	defer client.Close()

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

//gocyclo:ignore
func TestSentinelClientDelegateRetry(t *testing.T) {
	setup := func() (client *sentinelClient, cb func()) {
		retry := uint32(0)
		trigger := make(chan error)
		s0 := &mockConn{
			DoFn: func(cmd cmds.Completed) RedisResult { return RedisResult{} },
			DoMultiFn: func(multi ...cmds.Completed) []RedisResult {
				if atomic.LoadUint32(&retry) == 0 {
					return []RedisResult{
						{val: RedisMessage{typ: '*', values: []RedisMessage{}}},
						{val: RedisMessage{typ: '*', values: []RedisMessage{
							{typ: '+', string: ""}, {typ: '+', string: "1"},
						}}},
					}
				}
				return []RedisResult{
					{val: RedisMessage{typ: '*', values: []RedisMessage{}}},
					{val: RedisMessage{typ: '*', values: []RedisMessage{
						{typ: '+', string: ""}, {typ: '+', string: "2"},
					}}},
				}
			},
			ReceiveFn: func(ctx context.Context, subscribe cmds.Completed, fn func(message PubSubMessage)) error {
				if err, ok := <-trigger; ok {
					return err
				}
				return ErrClosing
			},
		}
		m1 := &mockConn{
			DoFn: func(cmd cmds.Completed) RedisResult {
				if cmd == cmds.RoleCmd {
					return RedisResult{val: RedisMessage{typ: '*', values: []RedisMessage{{typ: '+', string: "master"}}}}
				}
				atomic.AddUint32(&retry, 1)
				return newErrResult(ErrClosing)
			},
			DoCacheFn: func(cmd cmds.Cacheable, ttl time.Duration) RedisResult {
				atomic.AddUint32(&retry, 1)
				return newErrResult(ErrClosing)
			},
		}
		m2 := &mockConn{
			DoFn: func(cmd cmds.Completed) RedisResult {
				if cmd == cmds.RoleCmd {
					return RedisResult{val: RedisMessage{typ: '*', values: []RedisMessage{{typ: '+', string: "master"}}}}
				}
				return RedisResult{val: RedisMessage{typ: '+', string: "OK"}}
			},
			DoCacheFn: func(cmd cmds.Cacheable, ttl time.Duration) RedisResult {
				return RedisResult{val: RedisMessage{typ: '+', string: "OK"}}
			},
		}
		client, err := newSentinelClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			if dst == ":0" {
				return s0
			}
			if dst == ":1" {
				return m1
			}
			if dst == ":2" {
				return m2
			}
			return nil
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		return client, func() {
			for atomic.LoadUint32(&retry) >= 10 {
				break
			}
			trigger <- errors.New("die")
			close(trigger)
		}
	}

	t.Run("Delegate Do", func(t *testing.T) {
		client, cb := setup()

		go func() {
			cb()
		}()

		v, err := client.Do(context.Background(), client.B().Get().Key("k").Build()).ToString()
		if err != nil || v != "OK" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}

		client.Close()
	})

	t.Run("Delegate DoCache", func(t *testing.T) {
		client, cb := setup()

		go func() {
			cb()
		}()

		v, err := client.DoCache(context.Background(), client.B().Get().Key("k").Cache(), time.Minute).ToString()
		if err != nil || v != "OK" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}

		client.Close()
	})
}

//gocyclo:ignore
func TestSentinelClientPubSub(t *testing.T) {
	var s0count, s0close, m1close, m2close, m4close int32

	messages := make(chan PubSubMessage)

	s0 := &mockConn{
		DoFn: func(cmd cmds.Completed) RedisResult { return RedisResult{} },
		DoMultiFn: func(multi ...cmds.Completed) []RedisResult {
			count := atomic.AddInt32(&s0count, 1)
			if (count-1)%2 == 0 {
				return []RedisResult{
					{val: RedisMessage{typ: '*', values: []RedisMessage{}}},
					{val: RedisMessage{typ: '*', values: []RedisMessage{
						{typ: '+', string: ""}, {typ: '+', string: "1"},
					}}},
				}
			}
			return []RedisResult{
				{val: RedisMessage{typ: '*', values: []RedisMessage{}}},
				{val: RedisMessage{typ: '*', values: []RedisMessage{
					{typ: '+', string: ""}, {typ: '+', string: "2"},
				}}},
			}
		},
		ReceiveFn: func(ctx context.Context, subscribe cmds.Completed, fn func(message PubSubMessage)) error {
			for msg := range messages {
				fn(msg)
			}
			return ErrClosing
		},
		CloseFn: func() { atomic.AddInt32(&s0close, 1) },
	}
	m1 := &mockConn{
		DoFn: func(cmd cmds.Completed) RedisResult {
			if cmd == cmds.RoleCmd {
				return RedisResult{val: RedisMessage{typ: '*', values: []RedisMessage{{typ: '+', string: "master"}}}}
			}
			return RedisResult{val: RedisMessage{typ: '+', string: "OK"}}
		},
		CloseFn: func() {
			atomic.AddInt32(&m1close, 1)
		},
	}
	m2 := &mockConn{
		DoFn: func(cmd cmds.Completed) RedisResult {
			return RedisResult{val: RedisMessage{typ: '*', values: []RedisMessage{{typ: '+', string: "slave"}}}}
		},
		CloseFn: func() { atomic.AddInt32(&m2close, 1) },
	}
	s3 := &mockConn{
		DoMultiFn: func(cmd ...cmds.Completed) []RedisResult { return []RedisResult{newErrResult(errClosing)} },
	}
	m4 := &mockConn{
		DoFn: func(cmd cmds.Completed) RedisResult {
			if cmd == cmds.RoleCmd {
				return RedisResult{val: RedisMessage{typ: '*', values: []RedisMessage{{typ: '+', string: "master"}}}}
			}
			return RedisResult{val: RedisMessage{typ: '+', string: "OK4"}}
		},
		CloseFn: func() { atomic.AddInt32(&m4close, 1) },
	}

	client, err := newSentinelClient(&ClientOption{
		InitAddress: []string{":0"},
		Sentinel: SentinelOption{
			MasterSet: "test",
		},
	}, func(dst string, opt *ClientOption) conn {
		if dst == ":0" {
			return s0
		}
		if dst == ":1" {
			return m1
		}
		if dst == ":2" {
			return m2
		}
		if dst == ":3" {
			return s3
		}
		if dst == ":4" {
			return m4
		}
		return nil
	})
	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}

	messages <- PubSubMessage{Channel: "+sentinel", Message: "sentinel 000000  3"}

	var added bool
	for !added {
		client.mu.Lock()
		added = client.sentinels.Front().Value.(string) == ":3"
		client.mu.Unlock()
		t.Log("wait +sentinel")
		time.Sleep(time.Millisecond * 100)
	}

	// switch to false master
	messages <- PubSubMessage{Channel: "+switch-master", Message: "test  1  2"}

	for atomic.LoadInt32(&m2close) != 2 {
		t.Log("wait false m2 to be close", atomic.LoadInt32(&m2close))
		time.Sleep(time.Millisecond * 100)
	}

	for atomic.LoadInt32(&s0count) != 3 {
		t.Log("wait s0 to be call third time", atomic.LoadInt32(&s0count))
		time.Sleep(time.Millisecond * 100)
	}

	v, err := client.Do(context.Background(), client.B().Get().Key("k").Build()).ToString()
	if err != nil || v != "OK" {
		t.Fatalf("unexpected resp %v %v", v, err)
	}

	// switch to master by reboot
	messages <- PubSubMessage{Channel: "+reboot", Message: "master test  4"}

	for atomic.LoadInt32(&m1close) != 1 {
		t.Log("wait old m1 to be close", atomic.LoadInt32(&m1close))
		time.Sleep(time.Millisecond * 100)
	}

	v, err = client.Do(context.Background(), client.B().Get().Key("k").Build()).ToString()
	if err != nil || v != "OK4" {
		t.Fatalf("unexpected resp %v %v", v, err)
	}

	close(messages)
	client.Close()

	for atomic.LoadInt32(&s0close) != 4 {
		t.Log("wait old s0 to be close", atomic.LoadInt32(&s0close))
		time.Sleep(time.Millisecond * 100)
	}
	for atomic.LoadInt32(&m4close) != 1 {
		t.Log("wait old m1 to be close", atomic.LoadInt32(&m4close))
		time.Sleep(time.Millisecond * 100)
	}
}
