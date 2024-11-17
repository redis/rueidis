package rueidis

import (
	"sync"
	"time"
)

func newPool(cap int, dead wire, idleConnTTL *time.Duration, minSize int, makeFn func() wire) *pool {
	if cap <= 0 {
		cap = DefaultPoolSize
	}

	p := &pool{
		size:         0,
		minSize:      minSize,
		cap:          cap,
		dead:         dead,
		make:         makeFn,
		list:         make([]wire, 0, 4),
		cond:         sync.NewCond(&sync.Mutex{}),
		idleConnTTL:  idleConnTTL,
		lastUsage:    make(map[wire]time.Time),
		stopChan:     make(chan struct{}, 1),
		stopChanOnce: &sync.Once{},
	}

	go p.removeIdleConns()

	return p
}

type pool struct {
	dead         wire
	cond         *sync.Cond
	make         func() wire
	list         []wire
	size         int
	minSize      int
	cap          int
	down         bool
	idleConnTTL  *time.Duration
	stopChan     chan struct{}
	stopChanOnce *sync.Once
	lastUsage    map[wire]time.Time
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
		delete(p.lastUsage, v)
	}
	p.cond.L.Unlock()
	return v
}

func (p *pool) Store(v wire) {
	p.cond.L.Lock()
	if !p.down && v.Error() == nil {
		p.list = append(p.list, v)

		if p.idleConnTTL != nil {
			p.lastUsage[v] = time.Now()
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
	p.stopChanOnce.Do(func() {
		p.stopChan <- struct{}{}
	})
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

	ttl := *p.idleConnTTL
	ticker := time.NewTicker(ttl)

	for {
		select {
		case <-p.stopChan:
			ticker.Stop()
			return

		case <-ticker.C:
			p.cond.L.Lock()
			if p.down {
				p.cond.L.Unlock()
				return
			}

			for p.size > p.minSize && len(p.list) > 0 {
				w := p.list[0]
				lastUsageTime, ok := p.lastUsage[w]
				if !ok {
					break
				}

				if time.Since(lastUsageTime) <= ttl {
					break
				}

				w.Close()
				p.size--
				p.list = p.list[1:]
				delete(p.lastUsage, w)
			}

			p.cond.L.Unlock()
		}
	}
}
