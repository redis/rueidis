package conn

import "sync"

const defaultPoolSize = 100

func newPool(size int, makeFn func() *wire) *pool {
	return &pool{
		size: 0,
		make: makeFn,
		list: make([]*wire, 0, size),
		cond: sync.NewCond(&sync.Mutex{}),
	}
}

type pool struct {
	list []*wire
	cond *sync.Cond
	make func() *wire
	size int
}

func (p *pool) Acquire() (v *wire) {
	p.cond.L.Lock()
	for len(p.list) == 0 && p.size == cap(p.list) {
		p.cond.Wait()
	}
	if len(p.list) == 0 {
		v = p.make()
		p.size++
	} else {
		i := len(p.list) - 1
		v = p.list[i]
		p.list = p.list[:i]
	}
	p.cond.L.Unlock()
	return
}

func (p *pool) Store(v *wire) {
	p.cond.L.Lock()
	if v.Error() == nil {
		p.list = append(p.list, v)
	} else {
		p.size--
	}
	p.cond.L.Unlock()
	p.cond.Signal()
}
