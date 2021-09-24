package conn

import (
	"bufio"
	proto2 "github.com/rueian/rueidis/internal/proto"
	"net"

	"github.com/rueian/rueidis/internal/queue"
)

type InterfaceConn struct {
	r *bufio.Reader
	w *bufio.Writer
	q queue.Queue
}

func NewConn(conn net.Conn) *InterfaceConn {
	c := &InterfaceConn{r: bufio.NewReader(conn), w: bufio.NewWriter(conn), q: queue.NewRing(4096)}
	c.reading()
	return c
}

func (c *InterfaceConn) reading() {
	go func() {
		for {
			t := c.q.Next1(true)
			if t == nil {
				if err := c.w.Flush(); err != nil {
					panic(err)
				}
				t = c.q.Next1(false)
			}
			if err := t.C.WriteTo(c.w); err != nil {
				panic(err)
			}
		}
	}()
	go func() {
		for {
			msg, err := proto2.ReadNext(c.r)
			if msg != nil {
				if _, ok := msg.(*proto2.Push); ok {
					continue
				}
			}
			t := c.q.Next2()
			t.W <- &result{R: msg, E: err}
		}
	}()
}

func (c *InterfaceConn) Write(cmd proto2.StringArray) (proto2.Message, error) {
	t := queue.Task{C: cmd, W: make(chan interface{}, 1)}
	c.q.Put(&t)
	r := (<-t.W).(*result)
	return r.R, r.E
}

type result struct {
	R proto2.Message
	E error
}
