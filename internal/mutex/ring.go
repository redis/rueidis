package mutex

import (
	"runtime"
	"strconv"
	"sync/atomic"

	"github.com/rueian/rueidis/pkg/proto"
)

func newRing() *ring {
	r := &ring{}
	r.mask = uint64(len(r.store) - 1)
	for i := range r.store {
		r.store[i].ch = make(chan proto.Result, 1)
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
	ch  chan proto.Result
}

func (r *ring) putOne(m []string) chan proto.Result {
	return r.put(atomic.AddUint64(&r.write, 1)&r.mask, m)
}

func (r *ring) putMulti(m [][]string) []chan proto.Result {
	l := uint64(len(m))
	e := atomic.AddUint64(&r.write, l)
	s := e - l + 1

	chs := make([]chan proto.Result, len(m))
	for i := uint64(0); i < l; i++ {
		chs[i] = r.put((s+i)&r.mask, m[i])
	}
	return chs
}

func (r *ring) put(position uint64, m []string) chan proto.Result {
	n := &r.store[position]
	for !atomic.CompareAndSwapUint64(&n.r, 0, 1) {
		runtime.Gosched()
	}
	n.cmd = m
	atomic.StoreUint64(&n.r, 2)
	return n.ch
}

// tryNextCmd should be only called by one dedicated thread
func (r *ring) tryNextCmd() ([]string, chan proto.Result) {
	r.read1++
	p := r.read1 & r.mask
	n := &r.store[p]
	if !atomic.CompareAndSwapUint64(&n.r, 2, 3) {
		r.read1--
		return nil, nil
	}
	return n.cmd, n.ch
}

// nextCmd should be only called by one dedicated thread
func (r *ring) nextCmd() ([]string, chan proto.Result) {
	r.read1 = (r.read1 + 1) & r.mask
	n := &r.store[r.read1]
	for !atomic.CompareAndSwapUint64(&n.r, 2, 3) {
		runtime.Gosched()
	}
	return n.cmd, n.ch
}

// skipLastOne should be only called by the same thread right after calling nextCmd or tryNextCmd
// this is related to client side cache
func (r *ring) skipLastOne() {
	n := &r.store[r.read1]
	for !atomic.CompareAndSwapUint64(&n.r, 3, 0) {
		panic("unexpected skipLastOne call on ring")
	}
}

// nextResultCh should be only called by one dedicated thread
func (r *ring) nextResultCh() ([]string, chan proto.Result) {
	for {
		r.read2++
		p := r.read2 & r.mask
		n := &r.store[p]
		if old := atomic.SwapUint64(&n.r, 0); old == 3 {
			return n.cmd, n.ch
		} else if old == 0 {
			// already marked skipped by skipLastOne, try next
			// this is related to client side cache
			continue
		} else {
			panic("unexpected nextResultCh call on ring" + " " + strconv.FormatInt(int64(old), 10))
		}
	}
}
