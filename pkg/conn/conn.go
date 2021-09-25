package conn

import (
	"bufio"
	"errors"
	"net"
	"runtime"
	"sync/atomic"

	"github.com/rueian/rueidis/pkg/proto"
)

var ErrConnClosing = errors.New("conn is closing")

type Conn struct {
	waits int32
	state int32

	conn net.Conn
	err  atomic.Value

	r *bufio.Reader
	w *bufio.Writer
	q *ring
}

func NewConn(conn net.Conn) *Conn {
	c := &Conn{conn: conn, r: bufio.NewReader(conn), w: bufio.NewWriter(conn), q: newRing()}
	c.reading()
	return c
}

func (c *Conn) reading() {
	go func() {
		var err error
		defer func() {
			if err != nil {
				c.err.Store(err)
			}
			c.conn.Close()
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
		for atomic.LoadInt32(&c.state) != 2 {
			msg, err := proto.ReadNextMessage(c.r)
			if err == nil {
				if msg.Type == '>' {
					// TODO: handle push data
					continue
				}
			}
			ch := c.q.nextResultCh()
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
