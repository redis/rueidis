package rueidis

import (
	"context"
	"sync"
	"sync/atomic"

	"golang.org/x/sys/cpu"
)

type queue interface {
	PutOne(ctx context.Context, m Completed) (chan RedisResult, error)
	PutMulti(ctx context.Context, m []Completed, resps []RedisResult) (chan RedisResult, error)
	NextWriteCmd() (Completed, []Completed, chan RedisResult)
	WaitForWrite() (Completed, []Completed, chan RedisResult)
	NextResultCh() (Completed, []Completed, chan RedisResult, []RedisResult)
	FinishResult()
}

var _ queue = (*ring)(nil)

func newRing(factor int) *ring {
	if factor <= 0 {
		factor = DefaultRingScale
	}
	r := &ring{store: make([]node, 2<<(factor-1))}
	r.mask = uint32(len(r.store) - 1)
	for i := range r.store {
		m := &sync.Mutex{}
		r.store[i].c1 = sync.NewCond(m)
		r.store[i].c2 = sync.NewCond(m)
		r.store[i].ch = make(chan RedisResult) // this channel can't be buffered
	}
	return r
}

type ring struct {
	resc     *sync.Cond
	wake     chan struct{}
	store    []node // store's size must be 2^N to work with the mask
	_        cpu.CacheLinePad
	write    uint32
	_        cpu.CacheLinePad
	recycled uint32
	waiters  int32
	waitMu   sync.Mutex
	_        cpu.CacheLinePad
	read1    uint32
	read2    uint32
	mask     uint32
	resOK    bool
}

type node struct {
	c1    *sync.Cond
	c2    *sync.Cond
	ch    chan RedisResult
	one   Completed
	multi []Completed
	resps []RedisResult
	mark  uint32
	slept bool
}

// reserve 先以 CAS 预留空闲槽位；环满时再等待唤醒信号，从而避免取消后留下空槽。
func (r *ring) reserve(ctx context.Context) (*node, error) {
	done := ctx.Done()
	size := uint32(len(r.store))
	for {
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		recycled := atomic.LoadUint32(&r.recycled)
		w := atomic.LoadUint32(&r.write)
		if w-recycled < size {
			if atomic.CompareAndSwapUint32(&r.write, w, w+1) {
				return &r.store[(w+1)&r.mask], nil
			}
			continue
		}

		// 注册等待者后再次检查容量，避免释放发生在首次检查与注册之间时丢失唤醒。
		r.waitMu.Lock()
		atomic.AddInt32(&r.waiters, 1)
		recycled = atomic.LoadUint32(&r.recycled)
		if atomic.LoadUint32(&r.write)-recycled < size {
			atomic.AddInt32(&r.waiters, -1)
			r.waitMu.Unlock()
			continue
		}
		if r.wake == nil {
			r.wake = make(chan struct{})
		}
		wake := r.wake
		r.waitMu.Unlock()
		select {
		case <-wake:
		case <-done:
		}
		atomic.AddInt32(&r.waiters, -1)
	}
}

func (r *ring) PutOne(ctx context.Context, m Completed) (chan RedisResult, error) {
	n, err := r.reserve(ctx)
	if err != nil {
		return nil, err
	}
	n.c1.L.Lock()
	n.one = m
	n.mark = 1
	s := n.slept
	n.c1.L.Unlock()
	if s {
		n.c2.Broadcast()
	}
	return n.ch, nil
}

func (r *ring) PutMulti(ctx context.Context, m []Completed, resps []RedisResult) (chan RedisResult, error) {
	n, err := r.reserve(ctx)
	if err != nil {
		return nil, err
	}
	n.c1.L.Lock()
	n.multi = m
	n.resps = resps
	n.mark = 1
	s := n.slept
	n.c1.L.Unlock()
	if s {
		n.c2.Broadcast()
	}
	return n.ch, nil
}

// NextWriteCmd should be only called by one dedicated thread
func (r *ring) NextWriteCmd() (one Completed, multi []Completed, ch chan RedisResult) {
	r.read1++
	p := r.read1 & r.mask
	n := &r.store[p]
	n.c1.L.Lock()
	if n.mark == 1 {
		one, multi, ch = n.one, n.multi, n.ch
		n.mark = 2
	} else {
		r.read1--
	}
	n.c1.L.Unlock()
	return
}

// WaitForWrite should be only called by one dedicated thread
func (r *ring) WaitForWrite() (one Completed, multi []Completed, ch chan RedisResult) {
	r.read1++
	p := r.read1 & r.mask
	n := &r.store[p]
	n.c1.L.Lock()
	for n.mark != 1 {
		n.slept = true
		n.c2.Wait() // c1 and c2 share the same mutex
		n.slept = false
	}
	one, multi, ch = n.one, n.multi, n.ch
	n.mark = 2
	n.c1.L.Unlock()
	return
}

// NextResultCh should be only called by one dedicated thread
func (r *ring) NextResultCh() (one Completed, multi []Completed, ch chan RedisResult, resps []RedisResult) {
	r.read2++
	p := r.read2 & r.mask
	n := &r.store[p]
	r.resc = n.c1
	r.resOK = false
	n.c1.L.Lock()
	if n.mark == 2 {
		r.resOK = true
		one, multi, ch, resps = n.one, n.multi, n.ch, n.resps
		n.mark = 0
		n.one = Completed{}
		n.multi = nil
		n.resps = nil
	} else {
		r.read2--
	}
	return
}

// FinishResult should be only called by one dedicated thread
func (r *ring) FinishResult() {
	if r.resc != nil {
		r.resc.L.Unlock()
		r.resc.Signal()
		r.resc = nil
		if r.resOK {
			// 响应交付且槽位解锁后才发布容量，防止复用仍属于上一条命令的槽位。
			atomic.AddUint32(&r.recycled, 1)
			if atomic.LoadInt32(&r.waiters) > 0 {
				r.waitMu.Lock()
				if r.wake != nil {
					// 关闭本轮通知通道，确保取消的等待者不会吞掉其他请求的唤醒。
					close(r.wake)
					r.wake = nil
				}
				r.waitMu.Unlock()
			}
			r.resOK = false
		}
	}
}
