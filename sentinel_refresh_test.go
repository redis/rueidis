package rueidis

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// TestSentinelTopologyRefreshIntervalRecoversMissedSwitch shows the client
// recovering from a missed +switch-master event.
//
// A subscribed client can still miss the event: the sentinel connection can
// drop at the same moment the master dies, and by the time a watch is
// re-established the +switch-master was already published. Redis PUB/SUB has no
// replay, so a subscriber that reconnects after the event never hears it.
// Without periodic reconciliation the client stays bound to the old master; the
// silent healthy watch in this mock reproduces that end state.
func TestSentinelTopologyRefreshIntervalRecoversMissedSwitch(t *testing.T) {
	var switched atomic.Bool
	var roleAsked atomic.Int64

	// Sentinel answers with :1 first, then :2 — i.e. the master moved, but the
	// client is never told via PUB/SUB.
	s0 := &mockConn{
		DoFn: func(cmd Completed) RedisResult { return RedisResult{} },
		DoMultiFn: func(multi ...Completed) *redisresults {
			addr := "1"
			if switched.Load() {
				addr = "2"
			}
			return &redisresults{s: []RedisResult{
				{val: slicemsg('*', []RedisMessage{})},
				{val: slicemsg('*', []RedisMessage{strmsg('+', ""), strmsg('+', addr)})},
			}}
		},
	}
	node := func(role string) *mockConn {
		return &mockConn{
			DoFn: func(cmd Completed) RedisResult {
				roleAsked.Add(1)
				return RedisResult{val: slicemsg('*', []RedisMessage{strmsg('+', role)})}
			},
		}
	}
	m1, m2 := node("master"), node("master")

	client, err := newSentinelClient(
		&ClientOption{
			InitAddress: []string{":0"},
			Sentinel:    SentinelOption{TopologyRefreshInterval: 50 * time.Millisecond},
		},
		func(dst string, opt *ClientOption) conn {
			switch dst {
			case ":0":
				return s0
			case ":1":
				return m1
			case ":2":
				return m2
			}
			return nil
		},
		newRetryer(defaultRetryDelayFn),
	)
	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}
	defer client.Close()

	if got := client.mAddr.Load().(string); got != ":1" {
		t.Fatalf("expected initial master :1, got %v", got)
	}

	// Master moves; NO +switch-master is delivered.
	switched.Store(true)

	deadline := time.Now().Add(5 * time.Second)
	for time.Now().Before(deadline) {
		if client.mAddr.Load().(string) == ":2" {
			return // reconciled without any event
		}
		time.Sleep(20 * time.Millisecond)
	}
	t.Fatal("client never reconciled to the new master :2 — with no periodic " +
		"refresh a missed +switch-master pins it to the old master indefinitely")
}

func TestSentinelNegativeTopologyRefreshInterval(t *testing.T) {
	_, err := newSentinelClient(
		&ClientOption{
			InitAddress: []string{":0"},
			Sentinel:    SentinelOption{TopologyRefreshInterval: -1},
		},
		func(dst string, opt *ClientOption) conn { return &mockConn{} },
		newRetryer(defaultRetryDelayFn),
	)
	if err != ErrInvalidTopologyRefreshInterval {
		t.Fatalf("expected ErrInvalidTopologyRefreshInterval, got %v", err)
	}
}

// TestSentinelTopologyRefreshmentStopsOnClose pins that the background
// reconciler exits immediately on Close rather than on its next tick. With a
// multi-second interval, polling c.stop alone kept the goroutine alive well
// past shutdown — which the suite's leak detector caught, and which would be
// a real shutdown delay for an application closing its client (e.g. on
// graceful process restart).
func TestSentinelTopologyRefreshmentStopsOnClose(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())

	s0 := &mockConn{
		DoFn: func(cmd Completed) RedisResult { return RedisResult{} },
		DoMultiFn: func(multi ...Completed) *redisresults {
			return &redisresults{s: []RedisResult{
				{val: slicemsg('*', []RedisMessage{})},
				{val: slicemsg('*', []RedisMessage{strmsg('+', ""), strmsg('+', "1")})},
			}}
		},
	}
	m := &mockConn{DoFn: func(cmd Completed) RedisResult {
		return RedisResult{val: slicemsg('*', []RedisMessage{strmsg('+', "master")})}
	}}

	client, err := newSentinelClient(
		&ClientOption{
			InitAddress: []string{":0"},
			// Far longer than the test: only prompt exit on Close can pass.
			Sentinel: SentinelOption{TopologyRefreshInterval: time.Hour},
		},
		func(dst string, opt *ClientOption) conn {
			if dst == ":0" {
				return s0
			}
			return m
		},
		newRetryer(defaultRetryDelayFn),
	)
	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}
	client.Close()
	// Close must be idempotent: closeCh is closed under a CAS on stop.
	client.Close()
}

// TestSentinelTopologyRefreshKeepsReplica pins that reconciling does not move a
// client off a replica that is still healthy.
//
// Sentinel reports every replica and the client picks one. That pick used to be
// re-drawn at random on every refresh, which was almost invisible while refresh
// only ran on a sentinel event. On an interval it is not: _switchTarget closes
// the connection it replaces immediately, so with three replicas roughly two
// ticks in three tore down a working connection, and the requests in flight on
// it, for no reason at all. The replica must only change when the one in use
// stops being eligible.
//
// The ticks are performed by calling client.refresh() directly — the same call
// runTopologyRefreshment makes — rather than running a real ticker, so the
// count of reconciliations is exact.
func TestSentinelTopologyRefreshKeepsReplica(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())

	replicaAddrs := []string{":10", ":11", ":12"}
	var down atomic.Value // address currently reported with s-down-time
	down.Store("")

	replicasReply := func() RedisResult {
		out := make([]RedisMessage, 0, len(replicaAddrs))
		for _, addr := range replicaAddrs {
			fields := []RedisMessage{
				strmsg('+', "ip"), strmsg('+', ""),
				strmsg('+', "port"), strmsg('+', addr[1:]),
			}
			if addr == down.Load().(string) {
				fields = append(fields, strmsg('+', "s-down-time"), strmsg('+', "1000"))
			}
			out = append(out, slicemsg('*', fields))
		}
		return RedisResult{val: slicemsg('*', out)}
	}

	s0 := &mockConn{
		DoFn: func(cmd Completed) RedisResult { return RedisResult{} },
		DoMultiFn: func(multi ...Completed) *redisresults {
			return &redisresults{s: []RedisResult{
				{val: slicemsg('*', []RedisMessage{})},
				{val: slicemsg('*', []RedisMessage{strmsg('+', ""), strmsg('+', "1")})},
				replicasReply(),
			}}
		},
	}

	var mu sync.Mutex
	dials := map[string]int{}
	closes := map[string]int{}
	node := func(dst, role string) *mockConn {
		return &mockConn{
			DoFn: func(cmd Completed) RedisResult {
				return RedisResult{val: slicemsg('*', []RedisMessage{strmsg('+', role)})}
			},
			CloseFn: func() {
				mu.Lock()
				defer mu.Unlock()
				closes[dst]++
			},
		}
	}

	client, err := newSentinelClient(
		&ClientOption{
			InitAddress:    []string{":0"},
			SendToReplicas: func(cmd Completed) bool { return true },
		},
		func(dst string, opt *ClientOption) conn {
			mu.Lock()
			dials[dst]++
			mu.Unlock()
			switch dst {
			case ":0":
				return s0
			case ":1":
				return node(dst, "master")
			}
			return node(dst, "slave")
		},
		newRetryer(defaultRetryDelayFn),
	)
	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}
	defer client.Close()

	first := client.rAddr.Load().(string)

	for i := 0; i < 20; i++ {
		if err := client.refresh(); err != nil {
			t.Fatalf("refresh %d: %v", i, err)
		}
		if got := client.rAddr.Load().(string); got != first {
			t.Fatalf("refresh %d moved the client from replica %s to %s while both were healthy", i, first, got)
		}
	}

	mu.Lock()
	replicaDials, replicaCloses := 0, 0
	for _, addr := range replicaAddrs {
		replicaDials += dials[addr]
		replicaCloses += closes[addr]
	}
	mu.Unlock()
	if replicaDials != 1 || replicaCloses != 0 {
		t.Fatalf("20 reconciliations should touch no replica connection, got %d dials and %d closes",
			replicaDials, replicaCloses)
	}

	// Stickiness must not outlive eligibility: once sentinel reports the replica
	// in use as subjectively down, the client has to move to another one.
	down.Store(first)
	if err := client.refresh(); err != nil {
		t.Fatalf("refresh after s-down: %v", err)
	}
	if got := client.rAddr.Load().(string); got == first {
		t.Fatalf("client stayed on replica %s after sentinel reported it down", got)
	}
}
