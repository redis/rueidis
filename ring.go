package rueidis

import (
	"sync"
	"sync/atomic"

	"github.com/rueian/rueidis/internal/cmds"
)

type queue interface {
	PutOne(m cmds.Completed) chan RedisResult
	PutMulti(m []cmds.Completed) chan RedisResult
	NextWriteCmd() (cmds.Completed, []cmds.Completed, chan RedisResult)
	NextResultCh() (cmds.Completed, []cmds.Completed, chan RedisResult, *sync.Cond)
	CleanNoReply()
}

const ringSize = 1024

var _ queue = (*ring)(nil)

func newRing() *ring {
	r := &ring{}
	r.mask = uint64(len(r.store) - 1)
	for i := range r.store {
		r.store[i].ch = make(chan RedisResult, 0) // this channel can't be buffered
		r.store[i].cond = sync.NewCond(&sync.Mutex{})
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
	cond  *sync.Cond
	ch    chan RedisResult
	one   cmds.Completed
	multi []cmds.Completed
	mark  uint32
}

func (r *ring) PutOne(m cmds.Completed) chan RedisResult {
	n := &r.store[atomic.AddUint64(&r.write, 1)&r.mask]
	n.cond.L.Lock()
	for n.mark != 0 {
		n.cond.Wait()
	}
	n.one = m
	n.multi = nil
	n.mark = 1
	n.cond.L.Unlock()
	return n.ch
}

func (r *ring) PutMulti(m []cmds.Completed) chan RedisResult {
	n := &r.store[atomic.AddUint64(&r.write, 1)&r.mask]
	n.cond.L.Lock()
	for n.mark != 0 {
		n.cond.Wait()
	}
	n.one = cmds.Completed{}
	n.multi = m
	n.mark = 1
	n.cond.L.Unlock()
	return n.ch
}

// NextWriteCmd should be only called by one dedicated thread
func (r *ring) NextWriteCmd() (one cmds.Completed, multi []cmds.Completed, ch chan RedisResult) {
	r.read1++
	p := r.read1 & r.mask
	n := &r.store[p]
	n.cond.L.Lock()
	if n.mark == 1 {
		one, multi, ch = n.one, n.multi, n.ch
		n.mark = 2
	} else {
		r.read1--
	}
	n.cond.L.Unlock()
	return
}

// NextResultCh should be only called by one dedicated thread
func (r *ring) NextResultCh() (one cmds.Completed, multi []cmds.Completed, ch chan RedisResult, cond *sync.Cond) {
	r.read2++
	p := r.read2 & r.mask
	n := &r.store[p]
	cond = n.cond
	n.cond.L.Lock()
	if n.mark == 2 {
		one, multi, ch = n.one, n.multi, n.ch
		n.mark = 0
	} else {
		r.read2--
	}
	return
}

// CleanNoReply should be only called by one dedicated thread
func (r *ring) CleanNoReply() {
	p := (r.read2 + 1) & r.mask
	n := &r.store[p]
	n.cond.L.Lock()
	if n.mark == 2 {
		mNoReply := len(n.multi) != 0
		for _, one := range n.multi {
			mNoReply = mNoReply && one.NoReply()
		}
		if mNoReply || n.one.NoReply() {
			n.mark = 0
			r.read2++
		}
	}
	n.cond.L.Unlock()
	n.cond.Signal()
}
