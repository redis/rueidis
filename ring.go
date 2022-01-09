package rueidis

import (
	"runtime"
	"sync/atomic"

	"github.com/rueian/rueidis/internal/cmds"
)

type queue interface {
	PutOne(m cmds.Completed) chan RedisResult
	PutMulti(m []cmds.Completed) chan RedisResult
	NextWriteCmd() (cmds.Completed, []cmds.Completed, chan RedisResult)
	NextResultCh() (cmds.Completed, []cmds.Completed, chan RedisResult)
}

const ringSize = 4096

var _ queue = (*ring)(nil)

func newRing() *ring {
	r := &ring{}
	r.mask = uint64(len(r.store) - 1)
	for i := range r.store {
		r.store[i].ch = make(chan RedisResult, 1)
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
	store [ringSize]node // store's size must be 2^N to work with the mask
}

type node struct {
	mark  uint32
	one   cmds.Completed
	multi []cmds.Completed
	ch    chan RedisResult
}

func (r *ring) PutOne(m cmds.Completed) chan RedisResult {
	n := &r.store[atomic.AddUint64(&r.write, 1)&r.mask]
	for !atomic.CompareAndSwapUint32(&n.mark, 0, 1) {
		runtime.Gosched()
	}
	n.one = m
	n.multi = nil
	atomic.StoreUint32(&n.mark, 2)
	return n.ch
}

func (r *ring) PutMulti(m []cmds.Completed) chan RedisResult {
	n := &r.store[atomic.AddUint64(&r.write, 1)&r.mask]
	for !atomic.CompareAndSwapUint32(&n.mark, 0, 1) {
		runtime.Gosched()
	}
	n.one = cmds.Completed{}
	n.multi = m
	atomic.StoreUint32(&n.mark, 2)
	return n.ch
}

// NextWriteCmd should be only called by one dedicated thread
func (r *ring) NextWriteCmd() (one cmds.Completed, multi []cmds.Completed, ch chan RedisResult) {
	r.read1++
	p := r.read1 & r.mask
	n := &r.store[p]
	if atomic.LoadUint32(&n.mark) == 2 {
		one, multi, ch = n.one, n.multi, n.ch
		atomic.StoreUint32(&n.mark, 3)
		return
	}
	r.read1--
	return cmds.Completed{}, nil, nil
}

// NextResultCh should be only called by one dedicated thread
func (r *ring) NextResultCh() (one cmds.Completed, multi []cmds.Completed, ch chan RedisResult) {
	r.read2++
	p := r.read2 & r.mask
	n := &r.store[p]
	if atomic.LoadUint32(&n.mark) == 3 {
		one, multi, ch = n.one, n.multi, n.ch
		atomic.StoreUint32(&n.mark, 0)
		return
	}
	r.read2--
	return cmds.Completed{}, nil, nil
}
