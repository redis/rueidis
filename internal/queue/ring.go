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

// Next1 should be only called by one dedicated thread
func (r *Ring) Next1(try bool) *Task {
	r.read1++
	p := r.read1 & r.mask
	n := r.store[p]
	if try {
		if !atomic.CompareAndSwapUint64(&n.r, 2, 3) {
			r.read1--
			return nil
		}
	} else {
		for !atomic.CompareAndSwapUint64(&n.r, 2, 3) {
			runtime.Gosched()
		}
	}
	return n.v
}

// Next2 should be only called by one dedicated thread
func (r *Ring) Next2() *Task {
	r.read2++
	p := r.read2 & r.mask
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
