package queue

import (
	"runtime"
	"sync/atomic"
)

type Ring struct {
	_     [8]uint64
	write uint64
	_     [7]uint64
	read1 uint64
	_     [7]uint64
	read2 uint64
	_     [7]uint64
	mask  uint64
	_     [7]uint64
	store []*node
}

type node struct {
	_ [8]uint64
	r uint64
	_ [7]uint64
	v *Task
}

func (r *Ring) Put(m *Task) {
	p := atomic.AddUint64(&r.write, 1) & r.mask
	n := r.store[p]
	for !atomic.CompareAndSwapUint64(&n.r, 0, 1) {
		runtime.Gosched()
	}
	n.v = m
	atomic.StoreUint64(&n.r, 2)
}

func (r *Ring) Next1() *Task {
	p := atomic.AddUint64(&r.read1, 1) & r.mask
	n := r.store[p]
	for !atomic.CompareAndSwapUint64(&n.r, 2, 3) {
		runtime.Gosched()
	}
	return n.v
}

func (r *Ring) Next2() *Task {
	p := atomic.AddUint64(&r.read2, 1) & r.mask
	n := r.store[p]
	v := n.v
	if !atomic.CompareAndSwapUint64(&n.r, 3, 0) {
		panic("unexpected call on ring")
	}
	return v
}

func NewRing(size uint64) *Ring {
	size = roundUp(size)
	store := make([]*node, size)
	for i := range store {
		store[i] = &node{}
	}
	return &Ring{mask: size - 1, store: store}
}
