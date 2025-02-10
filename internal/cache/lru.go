package cache

import (
	"runtime"
	"sync"
	"unsafe"
)

const LRUEntrySize = unsafe.Sizeof(linked[[]byte]{})

type linked[V any] struct {
	key  string
	head chain[V]
	next unsafe.Pointer
	prev unsafe.Pointer
	size int64
	ts   int64
	mark int64
	mu   sync.RWMutex
}

func (h *linked[V]) find(key string, ts int64) (v V, ok bool) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if h.ts > ts {
		return h.head.find(key)
	}
	return
}

func (h *linked[V]) close() {
	h.mu.Lock()
	h.ts = 0
	h.head = chain[V]{}
	h.mu.Unlock()
}

type LRUDoubleMap[V any] struct {
	ma    map[string]*linked[V]
	bp    sync.Pool
	mu    sync.RWMutex
	head  unsafe.Pointer
	tail  unsafe.Pointer
	total int64
	limit int64
	mark  int64
}

func (m *LRUDoubleMap[V]) Find(key1, key2 string, ts int64) (val V, ok bool) {
	m.mu.RLock()
	h := m.ma[key1]
	if h != nil {
		val, ok = h.find(key2, ts)
	}
	m.mu.RUnlock()
	if ok {
		b := m.bp.Get().(*ruBatch[V])
		b.m[h] = struct{}{}
		if len(b.m) >= bpsize {
			m.moveToTail(b.m)
			clear(b.m)
		}
		m.bp.Put(b)
	}
	return
}

func (m *LRUDoubleMap[V]) Insert(key1, key2 string, size, ts int64, v V) {
	// TODO: a RLock fast path?
	m.mu.Lock()
	m.total += size
	for m.head != nil {
		h := (*linked[V])(m.head)
		if m.total <= m.limit && h.ts != 0 { // TODO: clear expired entries?
			break
		}
		m.total -= h.size
		delete(m.ma, h.key)
		h.mark -= 1
		m.head = h.next
		if m.head != nil {
			(*linked[V])(m.head).prev = nil
		} else {
			m.tail = nil
			break
		}
	}

	h := m.ma[key1]
	if h == nil {
		h = &linked[V]{key: key1, ts: ts, mark: m.mark}
		m.ma[key1] = h
	} else if h.ts <= ts {
		m.total -= h.size
		h.size = 0
	}
	h.ts = ts
	h.size += size
	h.next = nil
	if m.tail != nil && m.tail != unsafe.Pointer(h) {
		h.prev = m.tail
		(*linked[V])(m.tail).next = unsafe.Pointer(h)
	}
	m.tail = unsafe.Pointer(h)
	if m.head == nil {
		m.head = unsafe.Pointer(h)
	}
	h.head.insert(key2, v)
	m.mu.Unlock()
}

func (m *LRUDoubleMap[V]) Delete(key1 string) {
	m.mu.RLock()
	if h := m.ma[key1]; h != nil {
		h.close()
	}
	m.mu.RUnlock()
}

func (m *LRUDoubleMap[V]) DeleteAll() {
	m.mu.Lock()
	m.ma = make(map[string]*linked[V], len(m.ma))
	m.head = nil
	m.tail = nil
	m.total = 0
	m.mark++
	m.mu.Unlock()
}

func (m *LRUDoubleMap[V]) moveToTail(b map[*linked[V]]struct{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for h := range b {
		if h.mark != m.mark {
			continue
		}
		prev := h.prev
		next := h.next
		if prev != nil {
			(*linked[V])(prev).next = next
		}
		if next != nil {
			(*linked[V])(next).prev = prev
		}
		h.next = nil
		if m.tail != nil && m.tail != unsafe.Pointer(h) {
			h.prev = m.tail
			(*linked[V])(m.tail).next = unsafe.Pointer(h)
		}
		m.tail = unsafe.Pointer(h)
		if m.head == unsafe.Pointer(h) && next != nil {
			m.head = next
		}
	}
}

type ruBatch[V any] struct {
	m map[*linked[V]]struct{}
}

func NewLRUDoubleMap[V any](hint, limit int64) *LRUDoubleMap[V] {
	m := &LRUDoubleMap[V]{
		ma:    make(map[string]*linked[V], hint),
		limit: limit,
	}
	m.bp.New = func() interface{} {
		b := &ruBatch[V]{m: make(map[*linked[V]]struct{}, bpsize)}
		runtime.SetFinalizer(b, func(b *ruBatch[V]) {
			if len(b.m) > 0 {
				m.moveToTail(b.m)
				clear(b.m)
			}
		})
		return b
	}
	return m
}
