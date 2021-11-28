package conn

import (
	"crypto/tls"
	"errors"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/rueian/rueidis/internal/cmds"
	"github.com/rueian/rueidis/internal/proto"
)

var ErrConnClosing = errors.New("connection is closing")

// DefaultCacheBytes = 128 MiB.
const DefaultCacheBytes = 128 * (1 << 20)

const DefaultPoolSize = 1000

type Option struct {
	// CacheSizeEachConn is redis client side cache size that bind to each TCP connection to a single redis instance.
	// The default is DefaultCacheBytes.
	CacheSizeEachConn int

	// BlockingPoolSize is the size of the connection pool shared by blocking commands (ex BLPOP, XREAD with BLOCK).
	// The default is DefaultPoolSize.
	BlockingPoolSize int

	// Redis AUTH parameters
	Username   string
	Password   string
	ClientName string
	SelectDB   int

	// TCP & TLS
	DialTimeout time.Duration
	TLSConfig   *tls.Config

	// Redis PubSub callbacks
	PubSubHandlers PubSubHandlers
}

type dialFn func(dst string, opt Option) (net.Conn, error)
type wireFn func(conn net.Conn, opt Option) (Wire, error)

var broken Wire = (*wire)(nil)

type Conn struct {
	dst  string
	opt  Option
	wire atomic.Value
	mu   sync.Mutex
	pool *pool

	dialFn dialFn
	wireFn wireFn
}

func NewConn(dst string, option Option) *Conn {
	return newConn(dst, option, func(dst string, opt Option) (conn net.Conn, err error) {
		dialer := &net.Dialer{Timeout: opt.DialTimeout}
		if opt.TLSConfig != nil {
			conn, err = tls.DialWithDialer(dialer, "tcp", dst, opt.TLSConfig)
		} else {
			conn, err = dialer.Dial("tcp", dst)
		}
		return conn, err
	}, func(conn net.Conn, opt Option) (Wire, error) {
		return newWire(conn, opt)
	})
}

func newConn(dst string, option Option, dialFn dialFn, wireFn wireFn) *Conn {
	conn := &Conn{dst: dst, opt: option, dialFn: dialFn, wireFn: wireFn}
	conn.wire.Store(broken)
	conn.pool = newPool(option.BlockingPoolSize, conn.dialRetry)
	return conn
}

func (c *Conn) connect() (Wire, error) {
	if wire := c.wire.Load().(Wire); wire != broken {
		return wire, nil
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if wire := c.wire.Load().(Wire); wire != broken {
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

func (c *Conn) dial() (w Wire, err error) {
	conn, err := c.dialFn(c.dst, c.opt)
	if err == nil {
		w, err = c.wireFn(conn, c.opt)
	}
	return w, err
}

func (c *Conn) dialRetry() Wire {
retry:
	if wire, err := c.dial(); err == nil {
		return wire
	}
	goto retry
}

func (c *Conn) acquire() Wire {
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
	c.pool.Store(w)
}

func (c *Conn) Close() {
	// TODO close pool
	c.acquire().Close()
}

func isNetworkErr(err error) bool {
	return err != nil && err != ErrConnClosing
}
