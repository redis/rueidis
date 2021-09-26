package mutex

import (
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"

	"github.com/rueian/rueidis/pkg/proto"
)

type Conn interface {
	WriteOne(cmd []string) (res proto.Result)
	Close()
}

type ConnNoCache struct {
	waits int32
	state int32

	q *ring
}

func NewConnNoCache() Conn {
	c := &ConnNoCache{q: newRing()}
	c.reading()
	return c
}

func (c *ConnNoCache) reading() {
	go func() {
		for atomic.LoadInt32(&c.state) != 2 {
			_, ch := c.q.nextCmd()
			c.q.nextResultCh()
			ch <- proto.Result{}
		}
	}()
}

func (c *ConnNoCache) WriteOne(cmd []string) (res proto.Result) {
	atomic.AddInt32(&c.waits, 1)
	if atomic.LoadInt32(&c.state) == 0 {
		res = <-c.q.putOne(cmd)
	}
	atomic.AddInt32(&c.waits, -1)
	return res
}

func (c *ConnNoCache) Close() {
	atomic.CompareAndSwapInt32(&c.state, 0, 1)
	for atomic.LoadInt32(&c.waits) != 0 {
		runtime.Gosched()
	}
	atomic.CompareAndSwapInt32(&c.state, 1, 2)
}

type ConnMutexInEventLoop struct {
	waits int32
	state int32

	q  *ring
	mu sync.Mutex

	hits int
	evic int
}

func NewConnMutexInEventLoop(hits, evic int) Conn {
	c := &ConnMutexInEventLoop{q: newRing(), hits: hits, evic: evic}
	c.reading()
	return c
}

func (c *ConnMutexInEventLoop) reading() {
	go func() {
		for atomic.LoadInt32(&c.state) != 2 {
			_, ch := c.q.nextCmd()
			c.mu.Lock()
			c.mu.Unlock()
			c.q.nextResultCh()
			ch <- proto.Result{}
		}
	}()
	go func() {
		for atomic.LoadInt32(&c.state) != 2 {
			if rand.Intn(100) < c.evic {
				c.mu.Lock()
				c.mu.Unlock()
			}
		}
	}()
}

func (c *ConnMutexInEventLoop) WriteOne(cmd []string) (res proto.Result) {
	atomic.AddInt32(&c.waits, 1)
	if atomic.LoadInt32(&c.state) == 0 {
		res = <-c.q.putOne(cmd)
	}
	atomic.AddInt32(&c.waits, -1)
	return res
}

func (c *ConnMutexInEventLoop) Close() {
	atomic.CompareAndSwapInt32(&c.state, 0, 1)
	for atomic.LoadInt32(&c.waits) != 0 {
		runtime.Gosched()
	}
	atomic.CompareAndSwapInt32(&c.state, 1, 2)
}

type ConnMutexOnWrite struct {
	waits int32
	state int32

	q  *ring
	mu sync.Mutex

	hits int
	evic int
}

func NewConnMutexOnWrite(hits, evic int) Conn {
	c := &ConnMutexOnWrite{q: newRing(), hits: hits, evic: evic}
	c.reading()
	return c
}

func (c *ConnMutexOnWrite) reading() {
	go func() {
		for atomic.LoadInt32(&c.state) != 2 {
			c.q.nextCmd()
			_, ch := c.q.nextResultCh()
			ch <- proto.Result{}
		}
	}()
	go func() {
		for atomic.LoadInt32(&c.state) != 2 {
			if rand.Intn(100) < c.evic {
				c.mu.Lock()
				c.mu.Unlock()
			}
		}
	}()
}

func (c *ConnMutexOnWrite) WriteOne(cmd []string) (res proto.Result) {
	atomic.AddInt32(&c.waits, 1)
	if atomic.LoadInt32(&c.state) == 0 {
		c.mu.Lock()
		c.mu.Unlock()
		if rand.Intn(100) < c.hits {

		} else {
			res = <-c.q.putOne(cmd)
		}
	}
	atomic.AddInt32(&c.waits, -1)
	return res
}

func (c *ConnMutexOnWrite) Close() {
	atomic.CompareAndSwapInt32(&c.state, 0, 1)
	for atomic.LoadInt32(&c.waits) != 0 {
		runtime.Gosched()
	}
	atomic.CompareAndSwapInt32(&c.state, 1, 2)
}

type ConnRWMutexOnWrite struct {
	waits int32
	state int32

	q  *ring
	mu sync.RWMutex

	hits int
	evic int
}

func NewConnRWMutexOnWrite(hits, evic int) Conn {
	c := &ConnRWMutexOnWrite{q: newRing(), hits: hits, evic: evic}
	c.reading()
	return c
}

func (c *ConnRWMutexOnWrite) reading() {
	go func() {
		for atomic.LoadInt32(&c.state) != 2 {
			c.q.nextCmd()
			_, ch := c.q.nextResultCh()
			ch <- proto.Result{}
		}
	}()
	go func() {
		for atomic.LoadInt32(&c.state) != 2 {
			if rand.Intn(100) < c.evic {
				c.mu.Lock()
				c.mu.Unlock()
			}
		}
	}()
}

func (c *ConnRWMutexOnWrite) WriteOne(cmd []string) (res proto.Result) {
	atomic.AddInt32(&c.waits, 1)
	if atomic.LoadInt32(&c.state) == 0 {
		c.mu.RLock()
		c.mu.RUnlock()
		if rand.Intn(100) < c.hits {

		} else {
			res = <-c.q.putOne(cmd)
		}
	}
	atomic.AddInt32(&c.waits, -1)
	return res
}

func (c *ConnRWMutexOnWrite) Close() {
	atomic.CompareAndSwapInt32(&c.state, 0, 1)
	for atomic.LoadInt32(&c.waits) != 0 {
		runtime.Gosched()
	}
	atomic.CompareAndSwapInt32(&c.state, 1, 2)
}
