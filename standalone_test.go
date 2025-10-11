package rueidis

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/redis/rueidis/internal/cmds"
)

func TestNewStandaloneClientNoNode(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	if _, err := newStandaloneClient(
		&ClientOption{}, func(dst string, opt *ClientOption) conn {
			return nil
		}, newRetryer(defaultRetryDelayFn),
	); err != ErrNoAddr {
		t.Fatalf("unexpected err %v", err)
	}
}

func TestNewStandaloneClientError(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	v := errors.New("dial err")
	if _, err := newStandaloneClient(
		&ClientOption{InitAddress: []string{""}}, func(dst string, opt *ClientOption) conn { return &mockConn{DialFn: func() error { return v }} }, newRetryer(defaultRetryDelayFn),
	); err != v {
		t.Fatalf("unexpected err %v", err)
	}
}

func TestNewStandaloneClientReplicasError(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	v := errors.New("dial err")
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
	defer ShouldNotLeak(SetupLeakDetection())

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
	defer ShouldNotLeak(SetupLeakDetection())

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

func TestStandaloneRedirectHandling(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())

	// Create a mock redirect response
	redirectErr := RedisError(strmsg('-', "REDIRECT 127.0.0.1:6380"))

	// Mock primary connection that returns redirect
	primaryConn := &mockConn{
		DoFn: func(cmd Completed) RedisResult {
			return newErrResult(&redirectErr)
		},
	}

	// Mock redirect target connection that returns success
	redirectConn := &mockConn{
		DoFn: func(cmd Completed) RedisResult {
			return RedisResult{val: strmsg('+', "OK")}
		},
	}

	// Track which connection is being used
	var connUsed string

	s, err := newStandaloneClient(&ClientOption{
		InitAddress: []string{"primary"},
		Standalone: StandaloneOption{
			EnableRedirect: true,
		},
		DisableRetry: true,
	}, func(dst string, opt *ClientOption) conn {
		connUsed = dst
		if dst == "primary" {
			return primaryConn
		}
		return redirectConn
	}, newRetryer(defaultRetryDelayFn))

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer s.Close()

	ctx := context.Background()
	result := s.Do(ctx, s.B().Get().Key("test").Build())

	if result.Error() != nil {
		t.Errorf("expected no error after redirect, got: %v", result.Error())
	}

	if str, _ := result.ToString(); str != "OK" {
		t.Errorf("expected OK response after redirect, got: %s", str)
	}

	// Verify that the redirect target was used
	if connUsed != "127.0.0.1:6380" {
		t.Errorf("expected redirect to use 127.0.0.1:6380, got: %s", connUsed)
	}
}

func TestStandaloneDoCacheRedirectHandling(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())

	// Create a mock redirect response
	redirectErr := RedisError(strmsg('-', "REDIRECT 127.0.0.1:6380"))

	// Mock primary connection that returns redirect
	primaryConn := &mockConn{
		DoCacheFn: func(cmd Cacheable, ttl time.Duration) RedisResult {
			return newErrResult(&redirectErr)
		},
	}

	// Mock redirect target connection that returns success
	redirectConn := &mockConn{
		DoCacheFn: func(cmd Cacheable, ttl time.Duration) RedisResult {
			return RedisResult{val: strmsg('+', "OK")}
		},
	}

	// Track which connection is being used
	var connUsed string

	s, err := newStandaloneClient(&ClientOption{
		InitAddress: []string{"primary"},
		Standalone: StandaloneOption{
			EnableRedirect: true,
		},
		DisableRetry: true,
	}, func(dst string, opt *ClientOption) conn {
		connUsed = dst
		if dst == "primary" {
			return primaryConn
		}
		return redirectConn
	}, newRetryer(defaultRetryDelayFn))

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer s.Close()

	ctx := context.Background()
	result := s.DoCache(ctx, s.B().Get().Key("test").Cache(), time.Second)

	if result.Error() != nil {
		t.Errorf("expected no error after redirect, got: %v", result.Error())
	}

	if str, _ := result.ToString(); str != "OK" {
		t.Errorf("expected OK response after redirect, got: %s", str)
	}

	// Verify that the redirect target was used
	if connUsed != "127.0.0.1:6380" {
		t.Errorf("expected redirect to use 127.0.0.1:6380, got: %s", connUsed)
	}
}

func TestStandaloneRedirectDisabled(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())

	// Create a mock redirect response
	redirectErr := RedisError(strmsg('-', "REDIRECT 127.0.0.1:6380"))

	// Mock primary connection that returns redirect
	primaryConn := &mockConn{
		DoFn: func(cmd Completed) RedisResult {
			return newErrResult(&redirectErr)
		},
	}

	s, err := newStandaloneClient(&ClientOption{
		InitAddress: []string{"primary"},
		Standalone: StandaloneOption{
			EnableRedirect: false, // Redirect disabled
		},
		DisableRetry: true,
	}, func(dst string, opt *ClientOption) conn {
		return primaryConn
	}, newRetryer(defaultRetryDelayFn))

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer s.Close()

	ctx := context.Background()
	result := s.Do(ctx, s.B().Get().Key("test").Build())

	// Should return the original redirect error since redirect is disabled
	if result.Error() == nil {
		t.Error("expected redirect error to be returned when redirect is disabled")
	}

	if result.Error().Error() != "REDIRECT 127.0.0.1:6380" {
		t.Errorf("expected redirect error, got: %v", result.Error())
	}
}

func TestStandaloneDoCacheRedirectDisabled(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())

	// Create a mock redirect response
	redirectErr := RedisError(strmsg('-', "REDIRECT 127.0.0.1:6380"))

	// Mock primary connection that returns redirect
	primaryConn := &mockConn{
		DoCacheFn: func(cmd Cacheable, ttl time.Duration) RedisResult {
			return newErrResult(&redirectErr)
		},
	}

	s, err := newStandaloneClient(&ClientOption{
		InitAddress: []string{"primary"},
		Standalone: StandaloneOption{
			EnableRedirect: false, // Redirect disabled
		},
		DisableRetry: true,
	}, func(dst string, opt *ClientOption) conn {
		return primaryConn
	}, newRetryer(defaultRetryDelayFn))

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer s.Close()

	ctx := context.Background()
	result := s.DoCache(ctx, s.B().Get().Key("test").Cache(), time.Second)

	// Should return the original redirect error since redirect is disabled
	if result.Error() == nil {
		t.Error("expected redirect error to be returned when redirect is disabled")
	}

	if result.Error().Error() != "REDIRECT 127.0.0.1:6380" {
		t.Errorf("expected redirect error, got: %v", result.Error())
	}
}

func TestNewClientEnableRedirectPriority(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())

	// Test that EnableRedirect creates a standalone client
	s, err := newStandaloneClient(&ClientOption{
		InitAddress: []string{"primary"},
		Standalone: StandaloneOption{
			EnableRedirect: true,
		},
	}, func(dst string, opt *ClientOption) conn {
		return &mockConn{
			DialFn: func() error { return nil },
		}
	}, newRetryer(defaultRetryDelayFn))

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer s.Close()

	// Verify that we got a standalone client with redirect enabled
	if s.Mode() != ClientModeStandalone {
		t.Errorf("expected standalone client, got: %v", s.Mode())
	}

	// Verify that EnableRedirect is properly configured
	if !s.enableRedirect {
		t.Error("expected EnableRedirect to be true")
	}
}

func TestStandaloneDoStreamToReplica(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())

	replicaUsed := false
	primaryConn := &mockConn{
		DialFn: func() error { return nil },
		DoStreamFn: func(cmd Completed) RedisResultStream {
			return RedisResultStream{e: errors.New("primary")}
		},
	}

	replicaConn := &mockConn{
		DialFn: func() error { return nil },
		DoStreamFn: func(cmd Completed) RedisResultStream {
			replicaUsed = true
			return RedisResultStream{e: errors.New("replica")}
		},
	}

	s, err := newStandaloneClient(&ClientOption{
		InitAddress: []string{"primary"},
		Standalone: StandaloneOption{
			ReplicaAddress: []string{"replica"},
		},
		SendToReplicas: func(cmd Completed) bool { return cmd.IsReadOnly() },
		DisableRetry:   true,
	}, func(dst string, opt *ClientOption) conn {
		if dst == "primary" {
			return primaryConn
		}
		return replicaConn
	}, newRetryer(defaultRetryDelayFn))

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer s.Close()

	// Test DoStream to replica
	stream := s.DoStream(context.Background(), s.B().Get().Key("k").Build())
	if stream.Error() == nil || stream.Error().Error() != "replica" {
		t.Errorf("expected replica error, got %v", stream.Error())
	}

	if !replicaUsed {
		t.Error("expected replica to be used")
	}
}

func TestStandaloneDoMultiStreamToReplica(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())

	replicaUsed := false
	primaryConn := &mockConn{
		DialFn: func() error { return nil },
		DoMultiStreamFn: func(multi ...Completed) MultiRedisResultStream {
			return MultiRedisResultStream{e: errors.New("primary")}
		},
	}

	replicaConn := &mockConn{
		DialFn: func() error { return nil },
		DoMultiStreamFn: func(multi ...Completed) MultiRedisResultStream {
			replicaUsed = true
			return MultiRedisResultStream{e: errors.New("replica")}
		},
	}

	s, err := newStandaloneClient(&ClientOption{
		InitAddress: []string{"primary"},
		Standalone: StandaloneOption{
			ReplicaAddress: []string{"replica"},
		},
		SendToReplicas: func(cmd Completed) bool { return cmd.IsReadOnly() },
		DisableRetry:   true,
	}, func(dst string, opt *ClientOption) conn {
		if dst == "primary" {
			return primaryConn
		}
		return replicaConn
	}, newRetryer(defaultRetryDelayFn))

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer s.Close()

	// Test DoMultiStream to replica
	stream := s.DoMultiStream(context.Background(), s.B().Get().Key("k").Build())
	if stream.Error() == nil || stream.Error().Error() != "replica" {
		t.Errorf("expected replica error, got %v", stream.Error())
	}

	if !replicaUsed {
		t.Error("expected replica to be used")
	}
}

func TestStandalonePickReplica(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())

	primaryConn := &mockConn{
		DialFn: func() error { return nil },
	}

	replicaConn := &mockConn{
		DialFn: func() error { return nil },
	}

	s, err := newStandaloneClient(&ClientOption{
		InitAddress: []string{"primary"},
		Standalone: StandaloneOption{
			ReplicaAddress: []string{"replica"},
		},
		DisableRetry: true,
	}, func(dst string, opt *ClientOption) conn {
		if dst == "primary" {
			return primaryConn
		}
		return replicaConn
	}, newRetryer(defaultRetryDelayFn))

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer s.Close()

	// Test that pick() returns 0 for single replica
	index := s.pick()
	if index != 0 {
		t.Errorf("expected index 0, got %d", index)
	}
}

func TestNewStandaloneClientWithReplicasPartialFailure(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())

	dialCount := 0
	primaryConn := &mockConn{
		DialFn:  func() error { return nil },
		CloseFn: func() {},
	}

	replicaConn := &mockConn{
		DialFn: func() error {
			dialCount++
			if dialCount == 2 { // Second replica fails
				return errors.New("replica 2 dial failed")
			}
			return nil
		},
		CloseFn: func() {},
	}

	_, err := newStandaloneClient(&ClientOption{
		InitAddress: []string{"primary"},
		Standalone: StandaloneOption{
			ReplicaAddress: []string{"replica1", "replica2"},
		},
		DisableRetry: true,
	}, func(dst string, opt *ClientOption) conn {
		if dst == "primary" {
			return primaryConn
		}
		return replicaConn
	}, newRetryer(defaultRetryDelayFn))

	if err == nil {
		t.Error("expected error due to replica failure")
	}

	if err.Error() != "replica 2 dial failed" {
		t.Errorf("expected replica 2 dial failed, got %v", err)
	}
}

func TestStandalonePickMultipleReplicas(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())

	primaryConn := &mockConn{
		DialFn: func() error { return nil },
	}

	replicaConn := &mockConn{
		DialFn: func() error { return nil },
	}

	s, err := newStandaloneClient(&ClientOption{
		InitAddress: []string{"primary"},
		Standalone: StandaloneOption{
			ReplicaAddress: []string{"replica1", "replica2"},
		},
		DisableRetry: true,
	}, func(dst string, opt *ClientOption) conn {
		if dst == "primary" {
			return primaryConn
		}
		return replicaConn
	}, newRetryer(defaultRetryDelayFn))

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer s.Close()

	// Test that pick() returns a valid index for multiple replicas
	for i := 0; i < 10; i++ {
		index := s.pick()
		if index < 0 || index >= 2 {
			t.Errorf("expected index 0 or 1, got %d", index)
		}
	}
}

func TestStandaloneDoMultiWithRedirectRetry(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())

	redirectErr := RedisError(strmsg('-', "REDIRECT 127.0.0.1:6380"))
	attempts := 0

	primaryConn := &mockConn{
		DialFn: func() error { return nil },
		DoMultiFn: func(multi ...Completed) *redisresults {
			attempts++
			// First attempt returns redirect error, second returns success
			if attempts == 1 {
				return &redisresults{s: []RedisResult{newErrResult(&redirectErr)}}
			}
			return &redisresults{s: []RedisResult{RedisResult{val: strmsg('+', "OK")}}}
		},
	}

	redirectConnCalled := false
	redirectConn := &mockConn{
		DialFn: func() error {
			redirectConnCalled = true
			return nil
		},
		DoMultiFn: func(multi ...Completed) *redisresults {
			return &redisresults{s: []RedisResult{RedisResult{val: strmsg('+', "OK")}}}
		},
		CloseFn: func() {},
	}

	// Mock retry handler that allows one retry
	mockRetry := &mockRetryHandler{
		WaitOrSkipRetryFunc: func(ctx context.Context, attempts int, cmd Completed, err error) bool {
			return attempts < 2 // Allow one retry
		},
		RetryDelayFn: func(attempts int, _ Completed, err error) time.Duration {
			return time.Millisecond
		},
		WaitForRetryFn: func(ctx context.Context, duration time.Duration) {
			time.Sleep(duration)
		},
	}

	s, err := newStandaloneClient(&ClientOption{
		InitAddress: []string{"primary"},
		Standalone: StandaloneOption{
			EnableRedirect: true,
		},
		DisableRetry: false,
	}, func(dst string, opt *ClientOption) conn {
		if dst == "primary" {
			return primaryConn
		}
		return redirectConn
	}, mockRetry)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer s.Close()

	// Create a simple command using the internal cmds package
	cmd := cmds.NewCompleted([]string{"SET", "k", "v"})

	// Test DoMulti with redirect and retry
	results := s.DoMulti(context.Background(), cmd)
	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}

	if results[0].Error() != nil {
		t.Errorf("expected success after retry, got error: %v", results[0].Error())
	}

	// The primary connection should have been called once, then redirected
	if attempts != 1 {
		t.Errorf("expected 1 attempt on primary before redirect, got %d", attempts)
	}

	if !redirectConnCalled {
		t.Error("expected redirect connection to be called")
	}
}

func TestStandaloneDoMultiWithRedirectRetryFailure(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())

	redirectErr := RedisError(strmsg('-', "REDIRECT 127.0.0.1:6380"))

	primaryConn := &mockConn{
		DialFn: func() error { return nil },
		DoMultiFn: func(multi ...Completed) *redisresults {
			return &redisresults{s: []RedisResult{newErrResult(&redirectErr)}}
		},
	}

	// Redirect connection fails to dial
	redirectConn := &mockConn{
		DialFn: func() error {
			return errors.New("redirect connection failed")
		},
	}

	// Mock retry handler that doesn't allow retries after connection failure
	mockRetry := &mockRetryHandler{
		WaitOrSkipRetryFunc: func(ctx context.Context, attempts int, cmd Completed, err error) bool {
			return false // Don't retry
		},
		RetryDelayFn: func(attempts int, _ Completed, err error) time.Duration {
			return time.Millisecond
		},
		WaitForRetryFn: func(ctx context.Context, duration time.Duration) {
			time.Sleep(duration)
		},
	}

	s, err := newStandaloneClient(&ClientOption{
		InitAddress: []string{"primary"},
		Standalone: StandaloneOption{
			EnableRedirect: true,
		},
		DisableRetry: false,
	}, func(dst string, opt *ClientOption) conn {
		if dst == "primary" {
			return primaryConn
		}
		return redirectConn
	}, mockRetry)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer s.Close()

	// Create a simple command using the internal cmds package
	cmd := cmds.NewCompleted([]string{"SET", "k", "v"})

	// Test DoMulti with redirect failure
	results := s.DoMulti(context.Background(), cmd)
	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}

	// Should return the original redirect error since retry is not allowed
	if results[0].Error() == nil {
		t.Error("expected error to be returned")
	}

	if verr, ok := results[0].Error().(*RedisError); !ok || !strings.Contains(verr.Error(), "REDIRECT") {
		t.Errorf("expected REDIRECT error, got %v", results[0].Error())
	}
}

func TestStandaloneDoMultiCacheWithRedirectRetry(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())

	redirectErr := RedisError(strmsg('-', "REDIRECT 127.0.0.1:6380"))
	attempts := 0

	primaryConn := &mockConn{
		DialFn: func() error { return nil },
		DoMultiCacheFn: func(multi ...CacheableTTL) *redisresults {
			attempts++
			// First attempt returns redirect error, second returns success
			if attempts == 1 {
				return &redisresults{s: []RedisResult{newErrResult(&redirectErr)}}
			}
			return &redisresults{s: []RedisResult{RedisResult{val: strmsg('+', "OK")}}}
		},
	}

	redirectConnCalled := false
	redirectConn := &mockConn{
		DialFn: func() error {
			redirectConnCalled = true
			return nil
		},
		DoMultiCacheFn: func(multi ...CacheableTTL) *redisresults {
			return &redisresults{s: []RedisResult{RedisResult{val: strmsg('+', "OK")}}}
		},
		CloseFn: func() {},
	}

	// Mock retry handler that allows one retry
	mockRetry := &mockRetryHandler{
		WaitOrSkipRetryFunc: func(ctx context.Context, attempts int, cmd Completed, err error) bool {
			return attempts < 2 // Allow one retry
		},
		RetryDelayFn: func(attempts int, _ Completed, err error) time.Duration {
			return time.Millisecond
		},
		WaitForRetryFn: func(ctx context.Context, duration time.Duration) {
			time.Sleep(duration)
		},
	}

	s, err := newStandaloneClient(&ClientOption{
		InitAddress: []string{"primary"},
		Standalone: StandaloneOption{
			EnableRedirect: true,
		},
		DisableRetry: false,
	}, func(dst string, opt *ClientOption) conn {
		if dst == "primary" {
			return primaryConn
		}
		return redirectConn
	}, mockRetry)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer s.Close()

	// Create a simple command using the internal cmds package
	cmd := cmds.NewCompleted([]string{"SET", "k", "v"})

	// Test DoMulti with redirect and retry
	results := s.DoMultiCache(context.Background(), CT(Cacheable(cmd), time.Second))
	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}

	if results[0].Error() != nil {
		t.Errorf("expected success after retry, got error: %v", results[0].Error())
	}

	// The primary connection should have been called once, then redirected
	if attempts != 1 {
		t.Errorf("expected 1 attempt on primary before redirect, got %d", attempts)
	}

	if !redirectConnCalled {
		t.Error("expected redirect connection to be called")
	}
}

func TestStandaloneDoMultiCacheWithRedirectRetryFailure(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())

	redirectErr := RedisError(strmsg('-', "REDIRECT 127.0.0.1:6380"))

	primaryConn := &mockConn{
		DialFn: func() error { return nil },
		DoMultiCacheFn: func(multi ...CacheableTTL) *redisresults {
			return &redisresults{s: []RedisResult{newErrResult(&redirectErr)}}
		},
	}

	// Redirect connection fails to dial
	redirectConn := &mockConn{
		DialFn: func() error {
			return errors.New("redirect connection failed")
		},
	}

	// Mock retry handler that doesn't allow retries after connection failure
	mockRetry := &mockRetryHandler{
		WaitOrSkipRetryFunc: func(ctx context.Context, attempts int, cmd Completed, err error) bool {
			return false // Don't retry
		},
		RetryDelayFn: func(attempts int, _ Completed, err error) time.Duration {
			return time.Millisecond
		},
		WaitForRetryFn: func(ctx context.Context, duration time.Duration) {
			time.Sleep(duration)
		},
	}

	s, err := newStandaloneClient(&ClientOption{
		InitAddress: []string{"primary"},
		Standalone: StandaloneOption{
			EnableRedirect: true,
		},
		DisableRetry: false,
	}, func(dst string, opt *ClientOption) conn {
		if dst == "primary" {
			return primaryConn
		}
		return redirectConn
	}, mockRetry)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer s.Close()

	// Create a simple command using the internal cmds package
	cmd := cmds.NewCompleted([]string{"SET", "k", "v"})

	// Test DoMulti with redirect failure
	results := s.DoMultiCache(context.Background(), CT(Cacheable(cmd), time.Second))
	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}

	// Should return the original redirect error since retry is not allowed
	if results[0].Error() == nil {
		t.Error("expected error to be returned")
	}

	if verr, ok := results[0].Error().(*RedisError); !ok || !strings.Contains(verr.Error(), "REDIRECT") {
		t.Errorf("expected REDIRECT error, got %v", results[0].Error())
	}
}
