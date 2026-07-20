package rueidis

import (
	"testing"
)

// TestSwitchTargetDoesNotCloseReusedConn pins the invariant that a connection
// still referenced by mConn must never be closed on a failure path.
//
// Reachable in ordinary operation: Sentinel reports the address the client
// already holds, so _switchTarget reuses the live mConn as `target`; the node
// has just been demoted, so ROLE answers "slave"; the old code then called
// target.Close() and returned errNotMaster WITHOUT swapping. mConn was left
// pointing at a closed mux, so every later command failed with ErrClosing and
// never even attempted to dial — observed in production as minutes of total
// outage against a perfectly healthy Redis.
func TestSwitchTargetDoesNotCloseReusedConn(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())

	closed := false
	// Reports "slave": the demoted-master case.
	demoted := &mockConn{
		DoFn: func(cmd Completed) RedisResult {
			return RedisResult{val: slicemsg('*', []RedisMessage{strmsg('+', "slave")})}
		},
		CloseFn: func() { closed = true },
		ErrorFn: func() error {
			if closed {
				return ErrClosing
			}
			return nil
		},
	}

	c := &sentinelClient{
		mOpt:   &ClientOption{},
		connFn: func(dst string, opt *ClientOption) conn { return demoted },
	}
	c.mConn.Store(conn(demoted))
	c.mAddr.Store("127.0.0.1:6379")

	// Same address the client already holds -> the reuse path.
	if err := c._switchTarget("127.0.0.1:6379", true); err != errNotMaster {
		t.Fatalf("expected errNotMaster, got %v", err)
	}

	if closed {
		t.Fatal("_switchTarget closed the connection still referenced by mConn — " +
			"mConn now points at a closed mux and every command will fail with ErrClosing")
	}
	if got := c.mConn.Load().(conn); got.Error() != nil {
		t.Fatalf("mConn must remain usable after a failed switch, got error %v", got.Error())
	}
}
