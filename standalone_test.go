package rueidis

import (
	"context"
	"errors"
	"strconv"
	"sync/atomic"
	"testing"
	"time"
)

func TestNewStandaloneClientNoNode(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	if _, err := newStandaloneClient(
		&ClientOption{}, func(dst string, opt *ClientOption) conn {
			return nil
		}, newRetryer(defaultRetryDelayFn),
	); err != ErrNoAddr {
		t.Fatalf("unexpected err %v", err)
	}
}

func TestNewStandaloneClientError(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	v := errors.New("dail err")
	if _, err := newStandaloneClient(
		&ClientOption{InitAddress: []string{""}}, func(dst string, opt *ClientOption) conn { return &mockConn{DialFn: func() error { return v }} }, newRetryer(defaultRetryDelayFn),
	); err != v {
		t.Fatalf("unexpected err %v", err)
	}
}

func TestNewStandaloneClientReplicasError(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	v := errors.New("dail err")
	if _, err := newStandaloneClient(
		&ClientOption{
			InitAddress: []string{"1"},
			Standalone: StandaloneOption{
				ReplicaAddress: []string{"2", "3"}, // two replicas
			},
		}, func(dst string, opt *ClientOption) conn {
			return &mockConn{DialFn: func() error {
				if dst == "3" {
					return v
				}
				return nil
			}}
		}, newRetryer(defaultRetryDelayFn),
	); err != v {
		t.Fatalf("unexpected err %v", err)
	}
}

func TestNewStandaloneClientDelegation(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())

	w := &mockWire{}
	p := &mockConn{
		AddrFn: func() string {
			return "p"
		},
		DoFn: func(cmd Completed) RedisResult {
			return newErrResult(errors.New("primary"))
		},
		DoMultiFn: func(multi ...Completed) *redisresults {
			return &redisresults{s: []RedisResult{newErrResult(errors.New("primary"))}}
		},
		DoCacheFn: func(cmd Cacheable, ttl time.Duration) RedisResult {
			return newErrResult(errors.New("primary"))
		},
		DoMultiCacheFn: func(multi ...CacheableTTL) *redisresults {
			return &redisresults{s: []RedisResult{newErrResult(errors.New("primary"))}}
		},
		DoStreamFn: func(cmd Completed) RedisResultStream {
			return RedisResultStream{e: errors.New("primary")}
		},
		DoMultiStreamFn: func(cmd ...Completed) MultiRedisResultStream {
			return MultiRedisResultStream{e: errors.New("primary")}
		},
		ReceiveFn: func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
			return errors.New("primary")
		},
		AcquireFn: func() wire {
			return w
		},
	}
	r := &mockConn{
		AddrFn: func() string {
			return "r"
		},
		DoFn: func(cmd Completed) RedisResult {
			return newErrResult(errors.New("replica"))
		},
		DoMultiFn: func(multi ...Completed) *redisresults {
			return &redisresults{s: []RedisResult{newErrResult(errors.New("replica"))}}
		},
		DoStreamFn: func(cmd Completed) RedisResultStream {
			return RedisResultStream{e: errors.New("replica")}
		},
		DoMultiStreamFn: func(cmd ...Completed) MultiRedisResultStream {
			return MultiRedisResultStream{e: errors.New("replica")}
		},
		ReceiveFn: func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
			return errors.New("replica")
		},
	}

	c, err := newStandaloneClient(&ClientOption{
		InitAddress: []string{"p"},
		Standalone: StandaloneOption{
			ReplicaAddress: []string{"r"},
		},
		SendToReplicas: func(cmd Completed) bool {
			return cmd.IsReadOnly() && !cmd.IsUnsub()
		},
		DisableRetry: true,
	}, func(dst string, opt *ClientOption) conn {
		if dst == "p" {
			return p
		}
		return r
	}, newRetryer(defaultRetryDelayFn))

	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}

	defer c.Close()

	ctx := context.Background()
	if err := c.Do(ctx, c.B().Get().Key("k").Build()).Error(); err == nil || err.Error() != "replica" {
		t.Fatalf("unexpected err %v", err)
	}
	if err := c.Do(ctx, c.B().Set().Key("k").Value("v").Build()).Error(); err == nil || err.Error() != "primary" {
		t.Fatalf("unexpected err %v", err)
	}
	if err := c.DoCache(ctx, c.B().Get().Key("k").Cache(), time.Second).Error(); err == nil || err.Error() != "primary" {
		t.Fatalf("unexpected err %v", err)
	}
	if err := c.DoMulti(ctx, c.B().Get().Key("k").Build())[0].Error(); err == nil || err.Error() != "replica" {
		t.Fatalf("unexpected err %v", err)
	}
	if err := c.DoMulti(ctx, c.B().Set().Key("k").Value("v").Build())[0].Error(); err == nil || err.Error() != "primary" {
		t.Fatalf("unexpected err %v", err)
	}
	if err := c.DoMultiCache(ctx, CT(c.B().Get().Key("k").Cache(), time.Second))[0].Error(); err == nil || err.Error() != "primary" {
		t.Fatalf("unexpected err %v", err)
	}
	stream := c.DoStream(ctx, c.B().Get().Key("k").Build())
	if err := stream.Error(); err == nil || err.Error() != "replica" {
		t.Fatalf("unexpected err %v", err)
	}
	multiStream := c.DoMultiStream(ctx, c.B().Get().Key("k").Build())
	if err := multiStream.Error(); err == nil || err.Error() != "replica" {
		t.Fatalf("unexpected err %v", err)
	}
	stream = c.DoStream(ctx, c.B().Set().Key("k").Value("v").Build())
	if err := stream.Error(); err == nil || err.Error() != "primary" {
		t.Fatalf("unexpected err %v", err)
	}
	multiStream = c.DoMultiStream(ctx, c.B().Set().Key("k").Value("v").Build())
	if err := multiStream.Error(); err == nil || err.Error() != "primary" {
		t.Fatalf("unexpected err %v", err)
	}
	if err := c.Receive(ctx, c.B().Subscribe().Channel("ch").Build(), func(msg PubSubMessage) {}); err == nil || err.Error() != "replica" {
		t.Fatalf("unexpected err %v", err)
	}
	if err := c.Receive(ctx, c.B().Unsubscribe().Channel("ch").Build(), func(msg PubSubMessage) {}); err == nil || err.Error() != "primary" {
		t.Fatalf("unexpected err %v", err)
	}

	if err := c.Dedicated(func(dc DedicatedClient) error {
		if dc.(*dedicatedSingleClient).wire != w {
			return errors.New("wire")
		}
		return nil
	}); err != nil {
		t.Fatalf("unexpected err %v", err)
	}

	if dc, cancel := c.Dedicate(); dc.(*dedicatedSingleClient).wire != w {
		t.Fatalf("unexpected wire %v", dc.(*dedicatedSingleClient).wire)
	} else {
		cancel()
	}

	if c.Mode() != ClientModeStandalone {
		t.Fatalf("unexpected mode: %v", c.Mode())
	}

	nodes := c.Nodes()
	if len(nodes) != 2 && nodes["p"].(*singleClient).conn != p && nodes["r"].(*singleClient).conn != r {
		t.Fatalf("unexpected nodes %v", nodes)
	}
}

func TestNewStandaloneClientMultiReplicasDelegation(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())

	var counts [2]int32

	c, err := newStandaloneClient(&ClientOption{
		InitAddress: []string{"p"},
		Standalone: StandaloneOption{
			ReplicaAddress: []string{"0", "1"},
		},
		SendToReplicas: func(cmd Completed) bool {
			return cmd.IsReadOnly()
		},
		DisableRetry: true,
	}, func(dst string, opt *ClientOption) conn {
		if dst == "p" {
			return &mockConn{}
		}
		return &mockConn{
			DoFn: func(cmd Completed) RedisResult {
				i, _ := strconv.Atoi(dst)
				atomic.AddInt32(&counts[i], 1)
				return newErrResult(errors.New("replica"))
			},
		}
	}, newRetryer(defaultRetryDelayFn))

	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}

	defer c.Close()

	ctx := context.Background()

	for i := 0; i < 1000; i++ {
		if err := c.Do(ctx, c.B().Get().Key("k").Build()).Error(); err == nil || err.Error() != "replica" {
			t.Fatalf("unexpected err %v", err)
		}
	}
	for i := 0; i < len(counts); i++ {
		if atomic.LoadInt32(&counts[i]) == 0 {
			t.Fatalf("replica %d was not called", i)
		}
	}
}
