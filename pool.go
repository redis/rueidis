package rueidis

import (
	"sync"
	"time"
)

func newPool(cap int, dead wire, idleConnTTL *time.Duration, minSize int, makeFn func() wire) *pool {
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
		timers:      make(map[wire]*time.Timer),
	}
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
	idleConnTTL *time.Duration
	timers      map[wire]*time.Timer
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
		delete(p.timers, v)
	}
	p.cond.L.Unlock()
	return v
}

func (p *pool) Store(v wire) {
	p.cond.L.Lock()
	if !p.down && v.Error() == nil {
		p.list = append(p.list, v)

		if p.idleConnTTL != nil {
			p.removeIdleConns()
			p.timers[v] = time.NewTimer(*p.idleConnTTL)
		}
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
	for _, w := range p.list {
		w.Close()
	}
	p.cond.L.Unlock()
	p.cond.Broadcast()
}

func (p *pool) removeIdleConns() {
	if p.idleConnTTL == nil {
		return
	}

	for p.size > p.minSize {
		w := p.list[0]
		timer, ok := p.timers[w]
		if !ok {
			return
		}

		select {
		case <-timer.C:
			w.Close()
			p.size--
			p.list = p.list[1:]
			delete(p.timers, w)

		default:
			return
		}
	}
}
