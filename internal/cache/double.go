package cache

import (
	"runtime"
	"sync"
)

const bpsize = 1024

type head[V any] struct {
	chain[V]
	mu sync.RWMutex
}

type DoubleMap[V any] struct {
	ma map[string]*head[V]
	bp sync.Pool
	mu sync.RWMutex
}

func (m *DoubleMap[V]) Find(key1, key2 string) (val V, ok bool) {
	m.mu.RLock()
	if h := m.ma[key1]; h != nil {
		h.mu.RLock()
		val, ok = h.find(key2)
		h.mu.RUnlock()
	}
	m.mu.RUnlock()
	return
}

func (m *DoubleMap[V]) FindOrInsert(key1, key2 string, fn func() V) (val V, ok bool) {
	m.mu.RLock()
	if h := m.ma[key1]; h != nil {
		h.mu.Lock()
		if val, ok = h.find(key2); !ok {
			val = fn()
			h.insert(key2, val)
		}
		h.mu.Unlock()
		m.mu.RUnlock()
		return
	}
	if m.ma == nil {
		m.mu.RUnlock()
		return
	}
	m.mu.RUnlock()
	m.mu.Lock()
	h := m.ma[key1]
	if h != nil {
		if val, ok = h.find(key2); ok {
			m.mu.Unlock()
			return
		}
	} else if m.ma == nil {
		m.mu.Unlock()
		return
	} else {
		h = &head[V]{}
		m.ma[key1] = h
	}
	val = fn()
	h.insert(key2, val)
	m.mu.Unlock()
	return
}

func (m *DoubleMap[V]) Delete(key1, key2 string) {
	var empty bool
	m.mu.RLock()
	if h := m.ma[key1]; h != nil {
		h.mu.Lock()
		empty = h.delete(key2)
		h.mu.Unlock()
	}
	m.mu.RUnlock()
	if empty {
		e := m.bp.Get().(*empties)
		e.s = append(e.s, key1)
		if len(e.s) >= bpsize {
			m.delete(e.s)
			clear(e.s)
			e.s = e.s[:0]
		}
		m.bp.Put(e)
		return
	}
}

func (m *DoubleMap[V]) delete(keys []string) {
	m.mu.Lock()
	for _, key := range keys {
		if h := m.ma[key]; h != nil {
			if h.empty() {
				delete(m.ma, key)
			}
		}
	}
	m.mu.Unlock()
}

func (m *DoubleMap[V]) Close(cb func(V)) {
	m.mu.Lock()
	for _, h := range m.ma {
		if h.node.key != "" {
			cb(h.node.val)
		}
		for curr := h.node.next; curr != nil; curr = curr.next {
			cb(curr.val)
		}
	}
	m.ma = nil
	m.mu.Unlock()
}

type empties struct {
	s []string
}

func NewDoubleMap[V any](hint int) *DoubleMap[V] {
	m := &DoubleMap[V]{ma: make(map[string]*head[V], hint)}
	m.bp.New = func() any {
		e := &empties{s: make([]string, 0, bpsize)}
		runtime.SetFinalizer(e, func(e *empties) {
			if len(e.s) >= 0 {
				m.delete(e.s)
				clear(e.s)
			}
		})
		return e
	}
	return m
}
