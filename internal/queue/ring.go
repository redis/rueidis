package queue

import (
	"runtime"
	"sync/atomic"

	"github.com/rueian/rueidis/internal/cmds"
	"github.com/rueian/rueidis/internal/proto"
)

const RingSize = 4096

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
	mark  uint32
	one   cmds.Completed
	multi []cmds.Completed
	ch    chan proto.Result
}

func (r *Ring) PutOne(m cmds.Completed) chan proto.Result {
	n := r.acquire(atomic.AddUint64(&r.write, 1) & r.mask)
	n.one = m
	n.multi = nil
	atomic.StoreUint32(&n.mark, 2)
	return n.ch
}

func (r *Ring) PutMulti(m []cmds.Completed) chan proto.Result {
	n := r.acquire(atomic.AddUint64(&r.write, 1) & r.mask)
	n.one = cmds.Completed{}
	n.multi = m
	atomic.StoreUint32(&n.mark, 2)
	return n.ch
}

func (r *Ring) acquire(position uint64) *node {
	n := &r.store[position]
	for !atomic.CompareAndSwapUint32(&n.mark, 0, 1) {
		runtime.Gosched()
	}
	return n
}

// NextWriteCmd should be only called by one dedicated thread
func (r *Ring) NextWriteCmd() (cmds.Completed, []cmds.Completed, chan proto.Result) {
	r.read1++
	p := r.read1 & r.mask
	n := &r.store[p]
	if !atomic.CompareAndSwapUint32(&n.mark, 2, 3) {
		r.read1--
		return cmds.Completed{}, nil, nil
	}
	return n.one, n.multi, n.ch
}

// NextResultCh should be only called by one dedicated thread
func (r *Ring) NextResultCh() (one cmds.Completed, multi []cmds.Completed, ch chan proto.Result) {
	r.read2++
	p := r.read2 & r.mask
	n := &r.store[p]
	one, multi, ch = n.one, n.multi, n.ch
	if atomic.CompareAndSwapUint32(&n.mark, 3, 0) {
		return
	}
	r.read2--
	return cmds.Completed{}, nil, nil
}
