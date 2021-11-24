package conn

import (
	"crypto/tls"
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

	pool *pool
}

func NewConn(dst string, option Option) *Conn {
	conn := &Conn{dst: dst, opt: option}
	conn.wire.Store(broken)
	conn.pool = newPool(defaultPoolSize, conn.dialRetry)
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
		for resp := wire.Do(cmds.PingCmd); resp.NonRedisError() == nil; resp = wire.Do(cmds.PingCmd) {
			time.Sleep(time.Second)
		}
	}()

	c.wire.Store(wire)
	return wire, nil
}

func (c *Conn) dial() (w *wire, err error) {
	dialer := &net.Dialer{Timeout: c.opt.DialTimeout}

	var conn net.Conn
	if c.opt.TLSConfig != nil {
		conn, err = tls.DialWithDialer(dialer, "tcp", c.dst, c.opt.TLSConfig)
	} else {
		conn, err = dialer.Dial("tcp", c.dst)
	}
	if err == nil {
		w, err = newWire(conn, c.opt)
	}
	if err != nil {
		return nil, err
	}
	return w, nil
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

func (c *Conn) Dialable() error { // no retry
	_, err := c.connect()
	return err
}

func (c *Conn) Info() map[string]proto.Message {
	return c.acquire().Info()
}

func (c *Conn) Do(cmd cmds.Completed) (resp proto.Result) {
retry:
	if cmd.IsBlock() {
		resp = c.blocking(cmd)
	} else {
		resp = c.pipeline(cmd)
	}
	if cmd.IsReadOnly() && isNetworkErr(resp.NonRedisError()) {
		goto retry
	}
	return resp
}

func (c *Conn) DoMulti(multi ...cmds.Completed) (resp []proto.Result) {
	var block, write bool
	for _, cmd := range multi {
		block = block || cmd.IsBlock()
		write = write || cmd.IsWrite()
	}
retry:
	if block {
		resp = c.blockingMulti(multi)
	} else {
		resp = c.pipelineMulti(multi)
	}
	if !write {
		for _, r := range resp {
			if isNetworkErr(r.NonRedisError()) {
				goto retry
			}
		}
	}
	return resp
}

func (c *Conn) blocking(cmd cmds.Completed) (resp proto.Result) {
	wire := c.pool.Acquire()
	resp = wire.Do(cmd)
	c.pool.Store(wire)
	return resp
}

func (c *Conn) blockingMulti(cmd []cmds.Completed) (resp []proto.Result) {
	wire := c.pool.Acquire()
	resp = wire.DoMulti(cmd...)
	c.pool.Store(wire)
	return resp
}

func (c *Conn) pipeline(cmd cmds.Completed) (resp proto.Result) {
	wire := c.acquire()
	if resp = wire.Do(cmd); isNetworkErr(resp.NonRedisError()) {
		c.wire.CompareAndSwap(wire, broken)
	}
	return resp
}

func (c *Conn) pipelineMulti(cmd []cmds.Completed) (resp []proto.Result) {
	wire := c.acquire()
	resp = wire.DoMulti(cmd...)
	for _, r := range resp {
		if isNetworkErr(r.NonRedisError()) {
			c.wire.CompareAndSwap(wire, broken)
			return resp
		}
	}
	return resp
}

func (c *Conn) DoCache(cmd cmds.Cacheable, ttl time.Duration) proto.Result {
retry:
	wire := c.acquire()
	resp := wire.DoCache(cmd, ttl)
	if isNetworkErr(resp.NonRedisError()) {
		c.wire.CompareAndSwap(wire, broken)
		goto retry
	}
	return resp
}

func (c *Conn) Acquire() Wire {
	return c.pool.Acquire()
}

func (c *Conn) Store(w Wire) {
	c.pool.Store(w.(*wire))
}

func (c *Conn) Close() {
	// TODO close pool
	c.acquire().Close()
}

func isNetworkErr(err error) bool {
	return err != nil && err != ErrConnClosing
}
