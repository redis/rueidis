package cache

import (
	"runtime"
	"sync"
	"sync/atomic"
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
	mu   sync.RWMutex
	cnt  uint32
	mark int32
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
	mi    []string
	bp    sync.Pool
	mu    sync.RWMutex
	head  unsafe.Pointer
	tail  unsafe.Pointer
	total int64
	limit int64
	mark  int32
}

func (m *LRUDoubleMap[V]) Find(key1, key2 string, ts int64) (val V, ok bool) {
	m.mu.RLock()
	h := m.ma[key1]
	if h != nil {
		val, ok = h.find(key2, ts)
	}
	m.mu.RUnlock()
	if ok && atomic.AddUint32(&h.cnt, 1)&3 == 0 {
		b := m.bp.Get().(*ruBatch[V])
		b.s = append(b.s, h)
		if len(b.s) < bpsize {
			m.bp.Put(b)
			return
		}
		go func(m *LRUDoubleMap[V], b *ruBatch[V]) {
			m.moveToTail(b.s)
			clear(b.s)
			b.s = b.s[:0]
			m.bp.Put(b)
		}(m, b)
	}
	return
}

func (m *LRUDoubleMap[V]) remove(h *linked[V]) {
	h.mark -= 1
	next := h.next
	prev := h.prev
	h.next = nil
	h.prev = nil
	if next != nil {
		(*linked[V])(next).prev = prev
	}
	if prev != nil {
		(*linked[V])(prev).next = next
	}
	if m.head == unsafe.Pointer(h) {
		m.head = next
	}
	if m.tail == unsafe.Pointer(h) {
		m.tail = prev
	}
	atomic.AddInt64(&m.total, -h.size)
	delete(m.ma, h.key)
}

func (m *LRUDoubleMap[V]) move(h *linked[V]) {
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

func (m *LRUDoubleMap[V]) Insert(key1, key2 string, size, ts, now int64, v V) {
	m.mu.RLock()
	if h := m.ma[key1]; h != nil {
		atomic.AddInt64(&m.total, size)
		h.mu.Lock()
		if h.ts <= now {
			atomic.AddInt64(&m.total, -h.size)
			h.size = 0
			h.head = chain[V]{}
		}
		h.ts = ts
		h.size += size
		h.head.insert(key2, v)
		h.mu.Unlock()
		m.mu.RUnlock()
		return
	}
	m.mu.RUnlock()
	m.mu.Lock()
	if m.ma == nil {
		m.mu.Unlock()
		return
	}
	atomic.AddInt64(&m.total, size)
	for m.head != nil {
		h := (*linked[V])(m.head)
		if h.ts != 0 && h.ts > now && atomic.LoadInt64(&m.total) <= m.limit {
			break
		}
		m.remove(h)
	}

	h := &linked[V]{key: key1, ts: ts, size: size, mark: m.mark}
	h.head.insert(key2, v)
	m.ma[key1] = h // h must not exist in the map because this Insert is called sequentially.
	m.move(h)
	if m.head == nil {
		m.head = unsafe.Pointer(h)
	}
	m.mu.Unlock()
}

func (m *LRUDoubleMap[V]) Delete(key1 string) {
	if m.mi == nil {
		m.mi = make([]string, 0, bpsize)
	} else if len(m.mi) == bpsize {
		m.mu.Lock()
		for _, key := range m.mi {
			if h := m.ma[key]; h != nil && h.ts == 0 {
				m.remove(h)
			}
		}
		if h := m.ma[key1]; h != nil {
			m.remove(h)
		}
		for m.head != nil {
			h := (*linked[V])(m.head)
			if h.ts != 0 && atomic.LoadInt64(&m.total) <= m.limit {
				break
			}
			m.remove(h)
		}
		m.mu.Unlock()
		clear(m.mi)
		return
	}
	m.mi = append(m.mi, key1)
	m.mu.RLock()
	if h := m.ma[key1]; h != nil {
		h.close()
	}
	m.mu.RUnlock()
}

func (m *LRUDoubleMap[V]) DeleteAll() {
	m.mu.Lock()
	m.ma = nil
	m.mi = nil
	m.head = nil
	m.tail = nil
	atomic.StoreInt64(&m.total, 0)
	m.mark++
	m.mu.Unlock()
}

func (m *LRUDoubleMap[V]) Reset() {
	m.mu.Lock()
	m.ma = make(map[string]*linked[V], len(m.ma))
	m.mi = nil
	m.head = nil
	m.tail = nil
	atomic.StoreInt64(&m.total, 0)
	m.mark++
	m.mu.Unlock()
}

func (m *LRUDoubleMap[V]) moveToTail(s []*linked[V]) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, h := range s {
		if h.mark == m.mark {
			m.move(h)
		}
	}
	for m.head != nil {
		h := (*linked[V])(m.head)
		if h.ts != 0 && atomic.LoadInt64(&m.total) <= m.limit {
			break
		}
		m.remove(h)
	}
}

type ruBatch[V any] struct {
	s []*linked[V]
}

func NewLRUDoubleMap[V any](hint, limit int64) *LRUDoubleMap[V] {
	m := &LRUDoubleMap[V]{
		ma:    make(map[string]*linked[V], hint),
		limit: limit,
	}
	m.bp.New = func() interface{} {
		b := &ruBatch[V]{s: make([]*linked[V], 0, bpsize)}
		runtime.SetFinalizer(b, func(b *ruBatch[V]) {
			if len(b.s) > 0 {
				m.moveToTail(b.s)
			}
		})
		return b
	}
	return m
}
