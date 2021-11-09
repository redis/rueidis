package conn

import (
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/rueian/rueidis/internal/cmds"
	"github.com/rueian/rueidis/internal/proto"
)

var broken *wire

type Conn struct {
	dst  string
	opt  Option
	wire atomic.Value
	mu   sync.Mutex

	blocking *pool
}

func NewConn(dst string, option Option) *Conn {
	conn := &Conn{dst: dst, opt: option}
	conn.wire.Store(broken)
	conn.blocking = newPool(defaultPoolSize, conn.dialRetry)
	return conn
}

func (c *Conn) connect() (*wire, error) {
	if wire := c.wire.Load().(*wire); wire != broken {
		return wire, nil
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if wire := c.wire.Load().(*wire); wire != broken {
		return wire, nil
	}

	wire, err := c.dial()
	if err != nil {
		return nil, err
	}

	go func() {
		for resp := wire.Do(cmds.PingCmd); resp.Err == nil; resp = wire.Do(cmds.PingCmd) {
			time.Sleep(time.Second)
		}
	}()

	c.wire.Store(wire)
	return wire, nil
}

func (c *Conn) dial() (*wire, error) {
	conn, err := net.DialTimeout("tcp", c.dst, c.opt.DialTimeout)
	if err != nil {
		return nil, err
	}
	wire, err := newWire(conn, c.opt)
	if err != nil {
		return nil, err
	}
	return wire, nil
}

func (c *Conn) dialRetry() *wire {
retry:
	if wire, err := c.dial(); err == nil {
		return wire
	}
	goto retry
}

func (c *Conn) acquire() *wire {
retry:
	if wire, err := c.connect(); err == nil {
		return wire
	}
	goto retry
}

func (c *Conn) Init() error { // no retry
	_, err := c.connect()
	return err
}

func (c *Conn) Info() proto.Message {
	return c.acquire().Info()
}

func (c *Conn) Do(cmd cmds.Completed) (resp proto.Result) {
retry:
	if cmd.IsBlock() {
		wire := c.blocking.Acquire()
		resp = wire.Do(cmd)
		c.blocking.Store(wire)
		if retryable(resp.Err) {
			goto retry
		}
	} else {
		wire := c.acquire()
		resp = wire.Do(cmd)
		if retryable(resp.Err) {
			c.wire.CompareAndSwap(wire, broken)
			goto retry
		}
	}
	return resp
}

func (c *Conn) DoMulti(multi ...cmds.Completed) (resp []proto.Result) {
	var blocking bool
	for i := len(multi) - 1; i >= 0; i-- {
		if multi[i].IsBlock() {
			blocking = true
			break
		}
	}
retry:
	if blocking {
		wire := c.blocking.Acquire()
		resp = wire.DoMulti(multi...)
		c.blocking.Store(wire)
		for _, r := range resp {
			if retryable(r.Err) {
				goto retry
			}
		}
	} else {
		wire := c.acquire()
		resp := wire.DoMulti(multi...)
		for _, r := range resp {
			if retryable(r.Err) {
				c.wire.CompareAndSwap(wire, broken)
				goto retry
			}
		}
	}
	return resp
}

func (c *Conn) DoCache(cmd cmds.Cacheable, ttl time.Duration) proto.Result {
retry:
	wire := c.acquire()
	resp := wire.DoCache(cmd, ttl)
	if retryable(resp.Err) {
		c.wire.CompareAndSwap(wire, broken)
		goto retry
	}
	return resp
}

func (c *Conn) Close() {
	c.acquire().Close()
}

func retryable(err error) bool {
	return err != nil && err != ErrConnClosing
}
