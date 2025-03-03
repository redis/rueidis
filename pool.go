package rueidis

import (
	"sync"
	"time"
)

func newPool(cap int, dead wire, cleanup time.Duration, minSize int, idleTm time.Duration, makeFn func() wire) *pool {
	if cap <= 0 {
		cap = DefaultPoolSize
	}

	return &pool{
		size:    0,
		minSize: minSize,
		cap:     cap,
		dead:    dead,
		make:    makeFn,
		list:    make([]wire, 0, 4),
		cond:    sync.NewCond(&sync.Mutex{}),
		cleanup: cleanup,
		idleTm:  idleTm,
	}
}

type pool struct {
	dead    wire
	cond    *sync.Cond
	timer   *time.Timer
	make    func() wire
	list    []wire
	cleanup time.Duration
	size    int
	minSize int
	cap     int
	down    bool
	timerOn bool
	idleTm  time.Duration
}

func (p *pool) Acquire() (v wire) {
	p.cond.L.Lock()
retry:
	for len(p.list) == 0 && p.size == p.cap && !p.down {
		p.cond.Wait()
	}
	if p.down {
		v = p.dead
		p.cond.L.Unlock()
		return v
	}
	if len(p.list) == 0 {
		p.size++
		// unlock before start to make a new wire
		// allowing others to make wires concurrently instead of waiting in line
		p.cond.L.Unlock()
		v = p.make()
		return v
	}

	i := len(p.list) - 1
	v = p.list[i]
	p.list[i] = nil
	p.list = p.list[:i]
	if v.Error() != nil {
		p.size--
		v.Close()
		goto retry
	}
	p.cond.L.Unlock()
	return v
}

func (p *pool) Store(v wire) {
	p.cond.L.Lock()
	if !p.down && v.Error() == nil {
		v.SetLastAccess(time.Now())
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
	if p.cleanup == 0 || p.timerOn || len(p.list) <= p.minSize {
		return
	}

	p.timerOn = true
	if p.timer == nil {
		p.timer = time.AfterFunc(p.cleanup, p.removeIdleConns)
	} else {
		p.timer.Reset(p.cleanup)
	}
}

func (p *pool) removeIdleConns() {
	p.cond.L.Lock()
	defer p.cond.L.Unlock()

	newLen := min(p.minSize, len(p.list))
	if p.idleTm > 0 && newLen < len(p.list) {
		t := time.Now()
		pos := len(p.list)-1
		for i := 0; i < pos && pos >= newLen; {
			if t.Sub(p.list[pos].LastAccess()) > p.idleTm {
				pos--
			} else if t.Sub(p.list[i].LastAccess()) > p.idleTm {
				p.list[i], p.list[pos] = p.list[pos], p.list[i]
				pos--
				i++
			} else {
				i++
			}
		}
		if t.Sub(p.list[pos].LastAccess()) > p.idleTm {
			newLen = max(newLen, pos)
		} else {
			newLen = max(newLen, pos+1)
		}
	}
	for i, w := range p.list[newLen:] {
		w.Close()
		p.list[newLen+i] = nil
		p.size--
	}

	p.list = p.list[:newLen]
	p.timerOn = false
}

func (p *pool) stopTimer() {
	p.timerOn = false
	if p.timer != nil {
		p.timer.Stop()
	}
}
