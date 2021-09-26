package queue

import (
	"runtime"
	"sync/atomic"

	"github.com/rueian/rueidis/internal/proto"
)

const RingSize = 8192

func NewRing() *Ring {
	r := &Ring{}
	r.mask = uint64(len(r.store) - 1)
	for i := range r.store {
		r.store[i].ch = make(chan proto.Result, 1)
	}
	return r
}

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
	store [RingSize]node // store's size must be 2^N to work with the mask
}

type node struct {
	r   uint64
	cmd []string
	ch  chan proto.Result
}

func (r *Ring) PutOne(m []string) chan proto.Result {
	return r.put(atomic.AddUint64(&r.write, 1)&r.mask, m)
}

func (r *Ring) PutMulti(m [][]string) []chan proto.Result {
	l := uint64(len(m))
	e := atomic.AddUint64(&r.write, l)
	s := e - l + 1

	chs := make([]chan proto.Result, len(m))
	for i := uint64(0); i < l; i++ {
		chs[i] = r.put((s+i)&r.mask, m[i])
	}
	return chs
}

func (r *Ring) put(position uint64, m []string) chan proto.Result {
	n := &r.store[position]
	for !atomic.CompareAndSwapUint64(&n.r, 0, 1) {
		runtime.Gosched()
	}
	n.cmd = m
	atomic.StoreUint64(&n.r, 2)
	return n.ch
}

// TryNextCmd should be only called by one dedicated thread
func (r *Ring) TryNextCmd() []string {
	r.read1++
	p := r.read1 & r.mask
	n := &r.store[p]
	if !atomic.CompareAndSwapUint64(&n.r, 2, 3) {
		r.read1--
		return nil
	}
	return n.cmd
}

// NextCmd should be only called by one dedicated thread
func (r *Ring) NextCmd() []string {
	r.read1 = (r.read1 + 1) & r.mask
	n := &r.store[r.read1]
	for !atomic.CompareAndSwapUint64(&n.r, 2, 3) {
		runtime.Gosched()
	}
	return n.cmd
}

// NextResultCh should be only called by one dedicated thread
func (r *Ring) NextResultCh() ([]string, chan proto.Result) {
	r.read2++
	p := r.read2 & r.mask
	n := &r.store[p]
	if atomic.CompareAndSwapUint64(&n.r, 3, 0) {
		return n.cmd, n.ch
	}
	panic("unexpected NextResultCh call on ring")
}
