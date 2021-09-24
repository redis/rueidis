package conn

import (
	"runtime"
	"sync/atomic"

	"github.com/rueian/rueidis/pkg/proto"
)

func newRing() *ring {
	r := &ring{}
	r.mask = uint64(len(r.store) - 1)
	for i := range r.store {
		r.store[i].ch = make(chan result, 1)
	}
	return r
}

type ring struct {
	_     [8]uint64
	write uint64
	_     [7]uint64
	read1 uint64
	_     [7]uint64
	read2 uint64
	_     [7]uint64
	mask  uint64
	_     [7]uint64
	store [8192]node // store's size must be 2^N to work with the mask
}

type node struct {
	_   [8]uint64
	r   uint64
	_   [7]uint64
	cmd []string
	ch  chan result
}

type result struct {
	val proto.Message
	err error
}

func (r *ring) put(m []string) chan result {
	p := atomic.AddUint64(&r.write, 1) & r.mask
	n := &r.store[p]
	for !atomic.CompareAndSwapUint64(&n.r, 0, 1) {
		runtime.Gosched()
	}
	n.cmd = m
	atomic.StoreUint64(&n.r, 2)
	return n.ch
}

// tryNextCmd should be only called by one dedicated thread
func (r *ring) tryNextCmd() []string {
	r.read1++
	p := r.read1 & r.mask
	n := &r.store[p]
	if !atomic.CompareAndSwapUint64(&n.r, 2, 3) {
		r.read1--
		return nil
	}
	return n.cmd
}

// nextCmd should be only called by one dedicated thread
func (r *ring) nextCmd() []string {
	r.read1++
	p := r.read1 & r.mask
	n := &r.store[p]
	for !atomic.CompareAndSwapUint64(&n.r, 2, 3) {
		runtime.Gosched()
	}
	return n.cmd
}

// nextResultCh should be only called by one dedicated thread
func (r *ring) nextResultCh() chan result {
	r.read2++
	p := r.read2 & r.mask
	n := &r.store[p]
	v := n.ch
	if !atomic.CompareAndSwapUint64(&n.r, 3, 0) {
		panic("unexpected call on ring")
	}
	return v
}
