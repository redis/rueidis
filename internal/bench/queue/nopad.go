package queue

// This is a minimum implementation of queue using ring buffer without padding
// that showing the performance difference with using padding

import (
	"runtime"
	"sync/atomic"

	"github.com/rueian/rueidis/internal/proto"
	"github.com/rueian/rueidis/internal/queue"
)

func NewNoPadRing() *NoPadRing {
	r := &NoPadRing{}
	r.mask = uint64(len(r.store) - 1)
	return r
}

type NoPadRing struct {
	write uint64
	read1 uint64
	read2 uint64
	mask  uint64
	store [queue.RingSize]node
}

type node struct {
	r uint64
}

func (r *NoPadRing) PutOne(m []string) chan proto.Result {
	n := &r.store[atomic.AddUint64(&r.write, 1)&r.mask]
	for !atomic.CompareAndSwapUint64(&n.r, 0, 1) {
		runtime.Gosched()
	}
	atomic.StoreUint64(&n.r, 2)
	return nil
}

func (r *NoPadRing) NextCmd() []string {
	r.read1 = (r.read1 + 1) & r.mask
	n := &r.store[r.read1]
	for !atomic.CompareAndSwapUint64(&n.r, 2, 3) {
		runtime.Gosched()
	}
	return nil
}

func (r *NoPadRing) NextResultCh() ([]string, chan proto.Result) {
	r.read2++
	p := r.read2 & r.mask
	n := &r.store[p]
	atomic.CompareAndSwapUint64(&n.r, 3, 0)
	return nil, nil
}
