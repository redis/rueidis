package queue

// This is a minimum implementation of queue using golang channel
// that showing the performance difference with using ring buffer

import (
	"github.com/rueian/rueidis/internal/cmds"
	"github.com/rueian/rueidis/internal/proto"
	"github.com/rueian/rueidis/internal/queue"
)

func NewChan() *Chan {
	return &Chan{
		ch1: make(chan cmds.Completed, queue.RingSize),
		ch2: make(chan cmds.Completed, queue.RingSize),
	}
}

type Chan struct {
	ch1 chan cmds.Completed
	ch2 chan cmds.Completed
}

func (c *Chan) PutOne(m cmds.Completed) chan proto.Result {
	c.ch1 <- m
	return nil
}

func (c *Chan) NextCmd() (cmds.Completed, []cmds.Completed) {
	m := <-c.ch1
	c.ch2 <- m
	return cmds.Completed{}, nil
}

func (c *Chan) NextResultCh() (cmds.Completed, []cmds.Completed, chan proto.Result) {
	<-c.ch2
	return cmds.Completed{}, nil, nil
}
