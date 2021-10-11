package queue

// This is a minimum implementation of queue using golang channel
// that showing the performance difference with using ring buffer

import (
	"github.com/rueian/rueidis/internal/proto"
	"github.com/rueian/rueidis/internal/queue"
)

func NewChan() *Chan {
	return &Chan{
		ch1: make(chan []string, queue.RingSize),
		ch2: make(chan []string, queue.RingSize),
	}
}

type Chan struct {
	ch1 chan []string
	ch2 chan []string
}

func (c *Chan) PutOne(m []string) chan proto.Result {
	c.ch1 <- m
	return nil
}

func (c *Chan) NextCmd() ([]string, [][]string) {
	m := <-c.ch1
	c.ch2 <- m
	return nil, nil
}

func (c *Chan) NextResultCh() ([]string, [][]string, chan proto.Result) {
	<-c.ch2
	return nil, nil, nil
}
