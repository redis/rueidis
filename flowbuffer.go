package rueidis

import (
	"context"
	"sync"
)

type flowBuffer struct {
	f    chan queuedCmd
	r    chan queuedCmd
	w    chan queuedCmd
	size int
}

var _ queue = (*flowBuffer)(nil)

func newFlowBuffer(factor int) *flowBuffer {
	if factor <= 0 {
		factor = DefaultRingScale
	}
	size := 2 << (factor - 1)

	r := &flowBuffer{
		f:    make(chan queuedCmd, size),
		r:    make(chan queuedCmd, size),
		w:    make(chan queuedCmd, size),
		size: size,
	}
	for i := 0; i < size; i++ {
		r.f <- queuedCmd{
			ch: make(chan RedisResult),
		}
	}
	return r
}

func (b *flowBuffer) PutOne(ctx context.Context, m Completed) (chan RedisResult, error) {
	select {
	case cmd := <-b.f:
		cmd.one = m

		b.w <- cmd

		return cmd.ch, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (b *flowBuffer) PutMulti(ctx context.Context, m []Completed, resps []RedisResult) (chan RedisResult, error) {
	select {
	case cmd := <-b.f:
		cmd.multi, cmd.resps = m, resps

		b.w <- cmd

		return cmd.ch, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// NextWriteCmd should be only called by one dedicated thread
func (b *flowBuffer) NextWriteCmd() (one Completed, multi []Completed, ch chan RedisResult) {
	select {
	case cmd := <-b.w:
		one, multi, ch = cmd.one, cmd.multi, cmd.ch

		b.r <- cmd

		return
	default:
		return
	}
}

// WaitForWrite should be only called by one dedicated thread
func (b *flowBuffer) WaitForWrite() (one Completed, multi []Completed, ch chan RedisResult) {
	cmd := <-b.w
	one, multi, ch = cmd.one, cmd.multi, cmd.ch

	b.r <- cmd

	return
}

// NextResultCh should be only called by one dedicated thread
func (b *flowBuffer) NextResultCh() (queuedCmd, *sync.Cond, chan<- queuedCmd) {
	select {
	case cmd := <-b.r:
		return cmd, nil, b.f
	default:
		return queuedCmd{}, nil, nil
	}
}
