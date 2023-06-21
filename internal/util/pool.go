package util

import (
	"sync"
	"sync/atomic"
)

type Unit interface {
	Capacity() int
	ResetLen(n int)
}

func NewPool[T Unit](fn func(capacity int) T) *Pool[T] {
	p := &Pool[T]{fn: fn}
	p.sp.New = func() any {
		return fn(int(atomic.LoadUint32(&p.ca)))
	}
	return p
}

type Pool[T Unit] struct {
	sp sync.Pool
	fn func(capacity int) T
	ca uint32
}

func (p *Pool[T]) Get(length, capacity int) T {
	s := p.sp.Get().(T)
	if s.Capacity() < capacity {
		atomic.StoreUint32(&p.ca, uint32(capacity))
		p.sp.Put(s)
		s = p.fn(capacity)
	}
	s.ResetLen(length)
	return s
}

func (p *Pool[T]) Put(s T) {
	p.sp.Put(s)
}
