package rueidis

import (
	"context"
)

type queue interface {
	PutOne(ctx context.Context, m Completed) (chan RedisResult, error)
	PutMulti(ctx context.Context, m []Completed, resps []RedisResult) (chan RedisResult, error)
	NextWriteCmd() (Completed, []Completed, chan RedisResult)
	WaitForWrite() (Completed, []Completed, chan RedisResult)
	NextResultCh() (node, chan<- node)
}

var _ queue = (*ring)(nil)

func newRing(factor int) *ring {
	if factor <= 0 {
		factor = DefaultRingScale
	}
	size := 2 << (factor - 1)

	r := &ring{
		f:    make(chan node, size),
		r:    make(chan node, size),
		w:    make(chan node, size),
		size: size,
	}
	for i := 0; i < size; i++ {
		r.f <- node{
			ch: make(chan RedisResult),
		}
	}
	return r
}

type ring struct {
	f    chan node
	r    chan node
	w    chan node
	size int
}

type node struct {
	ch    chan RedisResult
	one   Completed
	multi []Completed
	resps []RedisResult
}

func (n *node) reset() {
	n.one = Completed{}
	n.multi = nil
	n.resps = nil
}

func (r *ring) PutOne(ctx context.Context, m Completed) (chan RedisResult, error) {
	select {
	case n := <-r.f:
		n.one = m

		r.w <- n

		return n.ch, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (r *ring) PutMulti(ctx context.Context, m []Completed, resps []RedisResult) (chan RedisResult, error) {
	select {
	case n := <-r.f:
		n.multi, n.resps = m, resps

		r.w <- n

		return n.ch, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// NextWriteCmd should be only called by one dedicated thread
func (r *ring) NextWriteCmd() (one Completed, multi []Completed, ch chan RedisResult) {
	select {
	case n := <-r.w:
		one, multi, ch = n.one, n.multi, n.ch

		r.r <- n

		return
	default:
		return
	}
}

// WaitForWrite should be only called by one dedicated thread
func (r *ring) WaitForWrite() (one Completed, multi []Completed, ch chan RedisResult) {
	n := <-r.w
	one, multi, ch = n.one, n.multi, n.ch

	r.r <- n

	return
}

// NextResultCh should be only called by one dedicated thread
func (r *ring) NextResultCh() (node, chan<- node) {
	select {
	case n := <-r.r:
		return n, r.f
	default:
		return node{}, nil
	}
}
