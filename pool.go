package rueidis

import (
	"sync"
	"time"
)

func newPool(cap int, dead wire, idleConnTTL time.Duration, minSize int, makeFn func() wire) *pool {
	if cap <= 0 {
		cap = DefaultPoolSize
	}

	p := &pool{
		size:        0,
		minSize:     minSize,
		cap:         cap,
		dead:        dead,
		make:        makeFn,
		list:        make([]wire, 0, 4),
		cond:        sync.NewCond(&sync.Mutex{}),
		idleConnTTL: idleConnTTL,
	}

	if idleConnTTL != 0 {
		p.timer = time.AfterFunc(idleConnTTL, p.removeIdleConns)
	}

	return p
}

type pool struct {
	dead        wire
	cond        *sync.Cond
	make        func() wire
	list        []wire
	size        int
	minSize     int
	cap         int
	down        bool
	idleConnTTL time.Duration
	timer       *time.Timer
}

func (p *pool) Acquire() (v wire) {
	p.cond.L.Lock()
	for len(p.list) == 0 && p.size == p.cap && !p.down {
		p.cond.Wait()
	}
	if p.down {
		v = p.dead
	} else if len(p.list) == 0 {
		p.size++
		v = p.make()
	} else {
		i := len(p.list) - 1
		v = p.list[i]
		p.list = p.list[:i]
	}
	p.cond.L.Unlock()
	return v
}

func (p *pool) Store(v wire) {
	p.cond.L.Lock()
	if !p.down && v.Error() == nil {
		p.list = append(p.list, v)
		p.resetTimerIfNeeded()
	} else {
		p.size--
		v.Close()
	}
	p.cond.L.Unlock()
	p.cond.Signal()
}

func (p *pool) Close() {
	p.cond.L.Lock()
	p.down = true
	p.stopTimer()
	for _, w := range p.list {
		w.Close()
	}
	p.cond.L.Unlock()
	p.cond.Broadcast()
}

func (p *pool) removeIdleConns() {
	if p.idleConnTTL == 0 {
		return
	}

	p.cond.L.Lock()
	defer p.cond.L.Unlock()

	if p.down || p.size <= p.minSize || len(p.list) == 0 {
		return
	}

	newLen := len(p.list) - min(p.size-p.minSize, len(p.list))
	for _, w := range p.list[newLen:] {
		w.Close()
	}

	p.list = p.list[:newLen]
	p.size = p.minSize

	p.stopTimer()
}

func (p *pool) resetTimerIfNeeded() {
	if p.timer != nil && p.size > p.minSize {
		p.timer.Reset(p.idleConnTTL)
	}
}

func (p *pool) stopTimer() {
	if p.timer != nil {
		p.timer.Stop()
	}
}
