package conn

import (
	"bufio"
	"errors"
	"net"
	"runtime"
	"sync/atomic"
	"time"

	"github.com/rueian/rueidis/pkg/proto"
)

var ErrConnClosing = errors.New("conn is closing")
var OptInCmd = []string{"CLIENT", "CACHING", "yes"}

type Conn struct {
	waits int32
	state int32

	conn net.Conn
	err  atomic.Value

	r *bufio.Reader
	w *bufio.Writer
	q *ring

	cache *cache
}

func NewConn(conn net.Conn, cacheSize int) *Conn {
	c := &Conn{
		conn:  conn,
		cache: NewCache(cacheSize),
		r:     bufio.NewReader(conn),
		w:     bufio.NewWriter(conn),
		q:     newRing(),
	}
	c.reading()
	return c
}

func (c *Conn) reading() {
	go func() {
		var err error
		defer func() {
			c.cache.DeleteAll()
			if err != nil {
				c.err.Store(err)
			}
			_ = c.conn.Close()
			c.Close()
		}()
		for atomic.LoadInt32(&c.state) != 2 {
			cmd := c.q.tryNextCmd()
			if cmd == nil {
				if err = c.w.Flush(); err != nil {
					return
				}
				cmd = c.q.nextCmd()
			}
			if err = proto.WriteCmd(c.w, cmd); err != nil {
				return
			}
		}
	}()
	go func() {
		var opted bool
		for atomic.LoadInt32(&c.state) != 2 {
			msg, err := proto.ReadNextMessage(c.r)
			if msg.Type == '>' {
				// TODO: handle other push data
				// tracking-redir-broken
				// message
				// pmessage
				// subscribe
				// unsubscribe
				// psubscribe
				// punsubscribe
				// server-cpu-usage
				if len(msg.Values) < 2 {
					continue
				}
				if msg.Values[0].String == "invalidate" {
					c.cache.Delete(msg.Values[1].String)
				}
				continue
			}
			cmd, ch := c.q.nextResultCh()
			if err == nil && msg.Type != '-' && msg.Type != '!' && opted {
				c.cache.Update(cmd[1], msg)
			}
			opted = cmdEqual(cmd, OptInCmd)
			ch <- proto.Result{Val: msg, Err: err}
		}
	}()
}

func (c *Conn) WriteOne(cmd []string) (res proto.Result) {
	atomic.AddInt32(&c.waits, 1)
	if atomic.LoadInt32(&c.state) == 0 {
		res = <-c.q.putOne(cmd)
	} else {
		res.Err = ErrConnClosing
	}
	atomic.AddInt32(&c.waits, -1)
	return res
}

func (c *Conn) WriteMulti(cmd [][]string) []proto.Result {
	res := make([]proto.Result, len(cmd))
	atomic.AddInt32(&c.waits, 1)
	if atomic.LoadInt32(&c.state) == 0 {
		for i, ch := range c.q.putMulti(cmd) {
			res[i] = <-ch
		}
	} else {
		for i := 0; i < len(res); i++ {
			res[i].Err = ErrConnClosing
		}
	}
	atomic.AddInt32(&c.waits, -1)
	return res
}

func (c *Conn) WriteOptInCache(cmd []string, ttl time.Duration) proto.Result {
	v := c.cache.GetOrPrepare(cmd[1], ttl)
	if v.Type != 0 {
		return proto.Result{Val: v}
	}
	return c.WriteMulti([][]string{OptInCmd, cmd})[1]
}

func (c *Conn) Close() {
	atomic.CompareAndSwapInt32(&c.state, 0, 1)
	for atomic.LoadInt32(&c.waits) != 0 {
		runtime.Gosched()
	}
	atomic.CompareAndSwapInt32(&c.state, 1, 2)
}

func (c *Conn) Err() error {
	v := c.err.Load()
	if v == nil {
		return nil
	}
	return v.(error)
}

func cmdEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, s := range a {
		if s != b[i] {
			return false
		}
	}
	return true
}
