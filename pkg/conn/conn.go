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
			if err = cmd.WriteTo(c.w); err != nil {
				return
			}
		}
	}()
	go func() {
		for atomic.LoadInt32(&c.state) != 2 {
			msg, err := proto.ReadNextRaw(c.r)
			if err == nil {
				if msg.Type == '>' {
					// TODO: handle push data
					continue
				}
			}
			ch := c.q.nextResultCh()
			ch <- result{val: msg, err: err}
		}
	}()
}

func (c *Conn) Write(cmd proto.StringArray) (proto.Raw, error) {
	atomic.AddInt32(&c.waits, 1)
	if atomic.LoadInt32(&c.state) != 0 {
		atomic.AddInt32(&c.waits, -1)
		return proto.Raw{}, ErrConnClosing
	}
	r := <-c.q.put(cmd)
	atomic.AddInt32(&c.waits, -1)
	return r.val, r.err
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
