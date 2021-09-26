package conn

// This is a minimum implementation of client side caching locking simulation
// that showing the performance difference with using mutex on write path or in event loop

import (
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"

	"github.com/rueian/rueidis/internal/proto"
	"github.com/rueian/rueidis/internal/queue"
)

type Conn struct {
	waits int32
	state int32

	q  *queue.Ring
	mu sync.Mutex

	hits int
	evic int
}

func (c *Conn) Close() {
	atomic.CompareAndSwapInt32(&c.state, 0, 1)
	for atomic.LoadInt32(&c.waits) != 0 {
		runtime.Gosched()
	}
	atomic.CompareAndSwapInt32(&c.state, 1, 2)
}

func reading(c *Conn) {
	for atomic.LoadInt32(&c.state) != 2 {
		c.q.NextCmd()
		_, ch := c.q.NextResultCh()
		ch <- proto.Result{}
	}
}

func evicting(c *Conn) {
	for atomic.LoadInt32(&c.state) != 2 {
		if rand.Intn(100) < c.evic {
			c.mu.Lock()
			c.mu.Unlock()
		}
	}
}

func NewConnNoMutex() *Conn {
	c := &Conn{q: queue.NewRing()}
	go reading(c)
	return c
}

func NewConnMutexOnWrite(hits, evic int) *Conn {
	c := &Conn{q: queue.NewRing(), hits: hits, evic: evic}
	go reading(c)
	go evicting(c)
	return c
}

func NewConnMutexInEventLoop(hits, evic int) *Conn {
	c := &Conn{q: queue.NewRing(), hits: hits, evic: evic}
	go func() {
		for atomic.LoadInt32(&c.state) != 2 {
			c.q.NextCmd()
			c.mu.Lock()
			c.mu.Unlock()
			_, ch := c.q.NextResultCh()
			ch <- proto.Result{}
		}
	}()
	go evicting(c)
	return c
}

func WriteWithMutex(c *Conn, cmd []string) (res proto.Result) {
	atomic.AddInt32(&c.waits, 1)
	if atomic.LoadInt32(&c.state) == 0 {
		c.mu.Lock()
		c.mu.Unlock()
		if rand.Intn(100) >= c.hits {
			res = <-c.q.PutOne(cmd)
		}
	}
	atomic.AddInt32(&c.waits, -1)
	return res
}

func WriteNoMutex(c *Conn, cmd []string) (res proto.Result) {
	atomic.AddInt32(&c.waits, 1)
	if atomic.LoadInt32(&c.state) == 0 {
		res = <-c.q.PutOne(cmd)
	}
	atomic.AddInt32(&c.waits, -1)
	return res
}
