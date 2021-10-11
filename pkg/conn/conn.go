package conn

import (
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/rueian/rueidis/internal/cmds"
	"github.com/rueian/rueidis/internal/proto"
)

var ping = []string{"PING"}
var broken *wire

type Conn struct {
	Cmd  *cmds.Builder
	dst  string
	opt  Option
	wire atomic.Value
	mu   sync.Mutex
}

func NewConn(dst string, option Option) (*Conn, error) {
	conn := &Conn{
		Cmd: cmds.NewBuilder(),
		dst: dst, opt: option,
	}
	conn.wire.Store(broken)
	if _, err := conn.connect(); err != nil {
		return nil, err
	}
	return conn, nil
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

	conn, err := net.Dial("tcp", c.dst)
	if err != nil {
		return nil, err
	}

	wire, err := newWire(conn, c.opt)
	if err != nil {
		return nil, err
	}

	go func() {
		for resp := wire.Do(ping); resp.Err == nil; resp = wire.Do(ping) {
			time.Sleep(time.Second)
		}
	}()

	c.wire.Store(wire)
	return wire, nil
}

func (c *Conn) acquire() *wire {
retry:
	if wire, err := c.connect(); err == nil {
		return wire
	}
	goto retry
}

func (c *Conn) Info() proto.Message {
	return c.acquire().Info()
}

func (c *Conn) Do(cmd []string) proto.Result {
retry:
	wire := c.acquire()
	resp := wire.Do(cmd)
	if retryable(resp.Err) {
		c.wire.CompareAndSwap(wire, broken)
		goto retry
	}
	c.Cmd.Put(cmd)
	return resp
}

func (c *Conn) DoMulti(multi ...[]string) []proto.Result {
retry:
	wire := c.acquire()
	resp := wire.DoMulti(multi...)
	for _, r := range resp {
		if retryable(r.Err) {
			c.wire.CompareAndSwap(wire, broken)
			goto retry
		}
	}
	for _, cmd := range multi {
		c.Cmd.Put(cmd)
	}
	return resp
}

func (c *Conn) DoCache(cmd []string, ttl time.Duration) proto.Result {
retry:
	wire := c.acquire()
	resp := wire.DoCache(cmd, ttl)
	if retryable(resp.Err) {
		c.wire.CompareAndSwap(wire, broken)
		goto retry
	}
	c.Cmd.Put(cmd)
	return resp
}

func (c *Conn) Close() {
	c.acquire().Close()
}

func retryable(err error) bool {
	if err == nil || err == ErrConnClosing {
		return false
	}
	return true
}
