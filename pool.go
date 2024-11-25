package rueidis

import (
	"sync"
	"time"
)

func newPool(cap int, dead wire, idleConnTTL time.Duration, minSize int, makeFn func() wire) *pool {
	if cap <= 0 {
		cap = DefaultPoolSize
	}

	return &pool{
		size:        0,
		minSize:     minSize,
		cap:         cap,
		dead:        dead,
		make:        makeFn,
		list:        make([]wire, 0, 4),
		cond:        sync.NewCond(&sync.Mutex{}),
		idleConnTTL: idleConnTTL,
	}
}

type pool struct {
	dead          wire
	cond          *sync.Cond
	make          func() wire
	list          []wire
	size          int
	minSize       int
	cap           int
	down          bool
	idleConnTTL   time.Duration
	timer         *time.Timer
	timerIsActive bool
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
		p.startTimerIfNeeded()
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

func (p *pool) startTimerIfNeeded() {
	if p.idleConnTTL == 0 || p.timerIsActive || len(p.list) <= p.minSize {
		return
	}

	p.timerIsActive = true
	if p.timer == nil {
		p.timer = time.AfterFunc(p.idleConnTTL, p.removeIdleConns)
	} else {
		p.timer.Reset(p.idleConnTTL)
	}
}

func (p *pool) removeIdleConns() {
	p.cond.L.Lock()
	defer p.cond.L.Unlock()

	if p.down || len(p.list) <= p.minSize {
		return
	}

	newLen := min(p.minSize, len(p.list))
	for i, w := range p.list[newLen:] {
		w.Close()
		p.list[newLen+i] = nil
		p.size--
	}

	p.list = p.list[:newLen]
	p.timerIsActive = false
}

func (p *pool) stopTimer() {
	p.timerIsActive = false
	if p.timer != nil {
		p.timer.Stop()
	}
}
