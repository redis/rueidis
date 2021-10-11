package conn

import (
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/rueian/rueidis/internal/cmds"
	"github.com/rueian/rueidis/internal/proto"
)

var PingCmd = []string{"PING"}

type RetryConn struct {
	Cmd  *cmds.Builder
	dst  string
	opt  Option
	conn atomic.Value
	mu   sync.Mutex
}

func NewRetryConn(dst string, option Option) (*RetryConn, error) {
	conn := &RetryConn{
		Cmd: cmds.NewBuilder(),
		dst: dst, opt: option,
	}
	if _, err := conn.connect(); err != nil {
		return nil, err
	}
	return conn, nil
}

func (c *RetryConn) connect() (*Conn, error) {
	if v := c.conn.Load(); v != nil {
		return v.(*Conn), nil
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if v := c.conn.Load(); v != nil {
		return v.(*Conn), nil
	}

	tcp, err := net.Dial("tcp", c.dst)
	if err != nil {
		return nil, err
	}

	conn, err := newConn(tcp, c.opt)
	if err != nil {
		return nil, err
	}

	go func() {
		for res := conn.Do(PingCmd); res.Err == nil; res = conn.Do(PingCmd) {
			time.Sleep(time.Second)
		}
	}()

	c.conn.Store(conn)
	return conn, nil
}

func (c *RetryConn) acquire() *Conn {
retry:
	if conn, err := c.connect(); err == nil {
		return conn
	}
	goto retry
}

func (c *RetryConn) Info() proto.Message {
	return c.acquire().Info()
}

func (c *RetryConn) Do(cmd []string) (res proto.Result) {
retry:
	conn := c.acquire()
	res = conn.Do(cmd)
	if retryable(res.Err) {
		c.conn.CompareAndSwap(conn, nil)
		goto retry
	}
	c.Cmd.Put(cmd)
	return res
}

func (c *RetryConn) DoMulti(cmds ...[]string) []proto.Result {
retry:
	conn := c.acquire()
	res := conn.DoMulti(cmds...)
	for _, r := range res {
		if retryable(r.Err) {
			c.conn.CompareAndSwap(conn, nil)
			goto retry
		}
	}
	for _, cmd := range cmds {
		c.Cmd.Put(cmd)
	}
	return res
}

func (c *RetryConn) DoCache(cmd []string, ttl time.Duration) proto.Result {
retry:
	conn := c.acquire()
	res := conn.DoCache(cmd, ttl)
	if retryable(res.Err) {
		c.conn.CompareAndSwap(conn, nil)
		goto retry
	}
	c.Cmd.Put(cmd)
	return res
}

func (c *RetryConn) Close() {
	c.acquire().Close()
}

func retryable(err error) bool {
	if err == nil || err == ErrConnClosing {
		return false
	}
	return true
}
