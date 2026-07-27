package rueidis

import (
	"context"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/redis/rueidis/internal/cmds"
)

// TestCloseReturnsOnFullRing: pipe.Close() must return in bounded time even
// when the command ring is full and the connection has stopped answering.
//
// Close() enqueues a final PING through the ring and waits for its reply with
// a one-second escape. The escape used to cover only the reply: the enqueue
// itself (queue.PutOne) blocks while every ring slot is occupied, and at
// Close() time nothing can free a slot when the peer has stopped replying but
// the TCP connection is still open:
//
//   - the keepalive ping would normally detect the stall and error the
//     connection, which drains the ring (see TestExitOnRingFullAndPingTimeout)
//     — but Close() increments blcksig, which makes backgroundPing skip its
//     check, and stores errClosing, which stops the ping timer from
//     re-arming;
//   - Close() closes the network connection only AFTER the enqueue, so the
//     background loops keep waiting and the drain never starts.
//
// So Close() blocked forever. That turns into a goroutine calling
// client.Close() that never returns (a hung shutdown), or — worse — the
// sentinel client closing a replaced connection while holding its mutex,
// wedging all of its topology handling.
//
// It takes a connection that stalls without erroring (network partition,
// frozen server), enough in-flight commands to fill the ring, and a Close()
// arriving before the keepalive ping notices the stall.
func TestCloseReturnsOnFullRing(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	p, mock, _, closeConn := setup(t, ClientOption{
		RingScaleEachConn: 1, // 2-slot ring, cheap to fill
		ConnWriteTimeout:  500 * time.Millisecond,
		Dialer:            net.Dialer{KeepAlive: 500 * time.Millisecond},
	})
	p.background()

	// Fill the ring: the writer sends the commands into the socket, the mock
	// reads them and never replies, so every slot stays occupied waiting for
	// a reply.
	ringLen := len(p.queue.(*ring).store)
	for i := 0; i < ringLen; i++ {
		go func() { _ = p.Do(context.Background(), cmds.NewCompleted([]string{"GET", "a"})).Error() }()
	}
	for i := 0; i < ringLen; i++ {
		mock.Expect("GET", "a")
	}
	time.Sleep(50 * time.Millisecond)

	done := make(chan struct{})
	go func() { p.Close(); close(done) }()

	// Close's internal escape is one second; three is a deadlock.
	select {
	case <-done:
	case <-time.After(3 * time.Second):
		closeConn() // unpark the wedged Close so the leak detector can finish
		<-done
		t.Fatal("pipe.Close() did not return within 3s with a full ring and a silent connection")
	}
}

// TestCloseAfterConnErrorOnFullRing pins the behavior that keeps the common
// failure path safe: when the connection ERRORS instead of stalling silently
// — the usual case when a server process dies and the OS resets the
// connection — the background loops exit and drain the ring, so a Close()
// that follows returns promptly. The silent-stall variant, where nothing
// drains the ring, is covered by TestCloseReturnsOnFullRing.
func TestCloseAfterConnErrorOnFullRing(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	p, mock, _, closeConn := setup(t, ClientOption{
		RingScaleEachConn: 1,
		ConnWriteTimeout:  500 * time.Millisecond,
		Dialer:            net.Dialer{KeepAlive: 500 * time.Millisecond},
	})
	p.background()

	ringLen := len(p.queue.(*ring).store)
	for i := 0; i < ringLen; i++ {
		go func() { _ = p.Do(context.Background(), cmds.NewCompleted([]string{"GET", "a"})).Error() }()
	}
	for i := 0; i < ringLen; i++ {
		mock.Expect("GET", "a")
	}
	time.Sleep(50 * time.Millisecond)

	// The connection errors first (like a killed server), then Close runs.
	closeConn()

	done := make(chan struct{})
	go func() { p.Close(); close(done) }()

	select {
	case <-done:
	case <-time.After(3 * time.Second):
		t.Fatal("pipe.Close() did not return although the connection had already errored")
	}
}

// TestCloseReturnsOnFullRingBeforeBackground pins the link Close relies on to
// escape a stuck enqueue: cutting the connection only unparks queue.PutOne
// because the background loops then error out and drain the queue. That
// requires a background worker, and this is the one path to the enqueue where
// none is running yet.
//
// The pipe is still in sync mode: one caller owns the connection in syncDo,
// the others have queued behind it and filled the ring, and the peer never
// answers. Nothing here has started _background, and Close does not start it
// either — waits != 1 tells it a sync read may be in flight, and racing a
// second reader onto the same connection is what 8503b21 fixed. So the drain
// hangs on the sync read breaking first: the connection Close cuts is what
// breaks it, and its error path (pipe.go, syncDo) is what finally starts the
// worker that drains the ring.
func TestCloseReturnsOnFullRingBeforeBackground(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	p, mock, _, closeConn := setup(t, ClientOption{
		RingScaleEachConn: 1, // 2-slot ring, cheap to fill
	})
	// No p.background() here: the pipe must stay in sync mode.

	var callers sync.WaitGroup

	// context.Background() carries neither a deadline nor a Done channel, so
	// this one takes the syncDo path and owns the connection. ConnWriteTimeout
	// is unset, so syncDo clears the connection deadline and its read has no
	// bound of its own.
	callers.Add(1)
	go func() {
		defer callers.Done()
		_ = p.Do(context.Background(), cmds.NewCompleted([]string{"GET", "a"})).Error()
	}()
	mock.Expect("GET", "a") // reaches the peer, which never answers
	time.Sleep(50 * time.Millisecond)

	// Fill the ring behind the sync command. These never reach the socket:
	// with no background writer they just sit in their slots.
	ringLen := len(p.queue.(*ring).store)
	callers.Add(ringLen)
	for i := 0; i < ringLen; i++ {
		go func() {
			defer callers.Done()
			_ = p.Do(context.Background(), cmds.NewCompleted([]string{"GET", "b"})).Error()
		}()
	}
	time.Sleep(50 * time.Millisecond)

	done := make(chan struct{})
	go func() { p.Close(); close(done) }()

	// Close's internal escape is one second; three is a deadlock.
	select {
	case <-done:
	case <-time.After(3 * time.Second):
		closeConn() // unpark the wedged Close so the leak detector can finish
		<-done
		t.Fatal("pipe.Close() did not return with a full ring and no background worker running")
	}
	callers.Wait()
}
