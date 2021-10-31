package queue

// This is a minimum implementation of queue using ring buffer without padding
// that showing the performance difference with using padding

import (
	"github.com/rueian/rueidis/internal/cmds"
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
	mark uint32
	cmds []cmds.Completed
	ch   chan proto.Result
}

func (r *NoPadRing) PutOne(m cmds.Completed) chan proto.Result {
	n := &r.store[atomic.AddUint64(&r.write, 1)&r.mask]
	for !atomic.CompareAndSwapUint32(&n.mark, 0, 1) {
		runtime.Gosched()
	}
	atomic.StoreUint32(&n.mark, 2)
	return nil
}

func (r *NoPadRing) NextCmd() (cmds.Completed, []cmds.Completed, chan proto.Result) {
	r.read1 = (r.read1 + 1) & r.mask
	n := &r.store[r.read1]
	for !atomic.CompareAndSwapUint32(&n.mark, 2, 3) {
		runtime.Gosched()
	}
	return cmds.Completed{}, nil, nil
}

func (r *NoPadRing) NextResultCh() (cmds.Completed, []cmds.Completed, chan proto.Result) {
	r.read2++
	p := r.read2 & r.mask
	n := &r.store[p]
	atomic.CompareAndSwapUint32(&n.mark, 3, 0)
	return cmds.Completed{}, nil, nil
}
