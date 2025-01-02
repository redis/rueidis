package rueidis

import (
	"context"
	"runtime"
	"sync"
	"time"
	"unsafe"
)

// NewCacheStoreFn can be provided in ClientOption for using a custom CacheStore implementation
type NewCacheStoreFn func(CacheStoreOption) CacheStore

// CacheStoreOption will be passed to NewCacheStoreFn
type CacheStoreOption struct {
	// CacheSizeEachConn is redis client side cache size that bind to each TCP connection to a single redis instance.
	// The default is DefaultCacheBytes.
	CacheSizeEachConn int
}

// CacheStore is the store interface for the client side caching
// More detailed interface requirement can be found in cache_test.go
type CacheStore interface {
	// Flight is called when DoCache and DoMultiCache, with the requested client side ttl and the current time.
	// It should look up the store in single-flight manner and return one of the following three combinations:
	// Case 1: (empty RedisMessage, nil CacheEntry)     <- when cache missed, and rueidis will send the request to redis.
	// Case 2: (empty RedisMessage, non-nil CacheEntry) <- when cache missed, and rueidis will use CacheEntry.Wait to wait for response.
	// Case 3: (non-empty RedisMessage, nil CacheEntry) <- when cache hit
	Flight(key, cmd string, ttl time.Duration, now time.Time) (v RedisMessage, e CacheEntry)
	// Update is called when receiving the response of the request sent by the above Flight Case 1 from redis.
	// It should not only update the store but also deliver the response to all CacheEntry.Wait and return a desired client side PXAT of the response.
	// Note that the server side expire time can be retrieved from RedisMessage.CachePXAT.
	Update(key, cmd string, val RedisMessage) (pxat int64)
	// Cancel is called when the request sent by the above Flight Case 1 failed.
	// It should not only deliver the error to all CacheEntry.Wait but also remove the CacheEntry from the store.
	Cancel(key, cmd string, err error)
	// Delete is called when receiving invalidation notifications from redis.
	// If the keys is nil then it should delete all non-pending cached entries under all keys.
	// If the keys is not nil then it should delete all non-pending cached entries under those keys.
	Delete(keys []RedisMessage)
	// Close is called when connection between redis is broken.
	// It should flush all cached entries and deliver the error to all pending CacheEntry.Wait.
	Close(err error)
}

// CacheEntry should be used to wait for single-flight response when cache missed.
type CacheEntry interface {
	Wait(ctx context.Context) (RedisMessage, error)
}

// SimpleCache is an alternative interface should be paired with NewSimpleCacheAdapter to construct a CacheStore
type SimpleCache interface {
	Get(key string) RedisMessage
	Set(key string, val RedisMessage)
	Del(key string)
	Flush()
}

// NewSimpleCacheAdapter converts a SimpleCache into CacheStore
func NewSimpleCacheAdapter(store SimpleCache) CacheStore {
	return &adapter{store: store, flights: make(map[string]map[string]CacheEntry)}
}

type adapter struct {
	store   SimpleCache
	flights map[string]map[string]CacheEntry
	mu      sync.RWMutex
}

func (a *adapter) Flight(key, cmd string, ttl time.Duration, now time.Time) (RedisMessage, CacheEntry) {
	a.mu.RLock()
	if v := a.store.Get(key + cmd); v.typ != 0 && v.relativePTTL(now) > 0 {
		a.mu.RUnlock()
		return v, nil
	}
	flight := a.flights[key][cmd]
	a.mu.RUnlock()
	if flight != nil {
		return RedisMessage{}, flight
	}
	a.mu.Lock()
	entries := a.flights[key]
	if entries == nil && a.flights != nil {
		entries = make(map[string]CacheEntry, 1)
		a.flights[key] = entries
	}
	if flight = entries[cmd]; flight == nil && entries != nil {
		entries[cmd] = &adapterEntry{ch: make(chan struct{}), xat: now.Add(ttl).UnixMilli()}
	}
	a.mu.Unlock()
	return RedisMessage{}, flight
}

func (a *adapter) Update(key, cmd string, val RedisMessage) (sxat int64) {
	a.mu.Lock()
	entries := a.flights[key]
	if flight, ok := entries[cmd].(*adapterEntry); ok {
		sxat = val.getExpireAt()
		if flight.xat < sxat || sxat == 0 {
			sxat = flight.xat
			val.setExpireAt(sxat)
		}
		a.store.Set(key+cmd, val)
		flight.set(val, nil)
		entries[cmd] = nil
	}
	a.mu.Unlock()
	return
}

func (a *adapter) Cancel(key, cmd string, err error) {
	a.mu.Lock()
	entries := a.flights[key]
	if flight, ok := entries[cmd].(*adapterEntry); ok {
		flight.set(RedisMessage{}, err)
		entries[cmd] = nil
	}
	a.mu.Unlock()
}

func (a *adapter) del(key string) {
	entries := a.flights[key]
	for cmd, e := range entries {
		if e == nil {
			a.store.Del(key + cmd)
			delete(entries, cmd)
		}
	}
	if len(entries) == 0 {
		delete(a.flights, key)
	}
}

func (a *adapter) Delete(keys []RedisMessage) {
	a.mu.Lock()
	if keys == nil {
		for key := range a.flights {
			a.del(key)
		}
	} else {
		for _, k := range keys {
			a.del(k.string)
		}
	}
	a.mu.Unlock()
}

func (a *adapter) Close(err error) {
	a.mu.Lock()
	flights := a.flights
	a.flights = nil
	a.store.Flush()
	a.mu.Unlock()
	for _, entries := range flights {
		for _, e := range entries {
			if e != nil {
				e.(*adapterEntry).set(RedisMessage{}, err)
			}
		}
	}
}

type adapterEntry struct {
	err error
	ch  chan struct{}
	val RedisMessage
	xat int64
}

func (a *adapterEntry) set(val RedisMessage, err error) {
	a.err, a.val = err, val
	close(a.ch)
}

func (a *adapterEntry) Wait(ctx context.Context) (RedisMessage, error) {
	select {
	case <-ctx.Done():
		return RedisMessage{}, ctx.Err()
	case <-a.ch:
		return a.val, a.err
	}
}

type flatentry struct {
	ovfl *flatentry
	next unsafe.Pointer
	prev unsafe.Pointer
	cmd  string
	key  string
	val  []byte
	ttl  int64
	size int64
	mark int64
	mu   sync.Mutex
}

func (f *flatentry) insert(e *flatentry) {
	f.size += e.size
	f.mu.Lock()
	defer f.mu.Unlock()
	e.ovfl = f.ovfl
	f.ovfl = e
}

func (f *flatentry) find(cmd string, ts int64) (ret RedisMessage, expired bool) {
	if f == nil {
		return
	}
	if ts >= f.ttl {
		expired = true
		return
	}
	if cmd == f.cmd {
		_ = ret.CacheUnmarshalView(f.val)
		return
	}
	f.mu.Lock()
	ovfl := f.ovfl
	f.mu.Unlock()
	return ovfl.find(cmd, ts)
}

const lrBatchSize = 64

type lrBatch struct {
	m map[*flatentry]struct{}
}

func NewFlattenCache(limit int64) CacheStore {
	f := &flatten{
		flights: make(map[string]*adapterEntry),
		cache:   make(map[string]*flatentry),
		head:    &flatentry{},
		tail:    &flatentry{},
		size:    0,
		limit:   limit,
	}
	f.head.next = unsafe.Pointer(f.tail)
	f.tail.prev = unsafe.Pointer(f.head)
	f.lrup = sync.Pool{New: func() any {
		b := &lrBatch{m: make(map[*flatentry]struct{}, lrBatchSize)}
		runtime.SetFinalizer(b, func() {
			f.llTailBatch(b)
		})
		return b
	}}
	return f
}

type flatten struct {
	flights map[string]*adapterEntry
	cache   map[string]*flatentry
	head    *flatentry
	tail    *flatentry
	lrup    sync.Pool
	mark    int64
	size    int64
	limit   int64
	mu      sync.RWMutex
}

func (f *flatten) llAdd(e *flatentry) {
	e.mark = f.mark
	e.prev = f.tail.prev
	e.next = unsafe.Pointer(f.tail)
	f.tail.prev = unsafe.Pointer(e)
	(*flatentry)(e.prev).next = unsafe.Pointer(e)
}

func (f *flatten) llDel(e *flatentry) {
	(*flatentry)(e.prev).next = e.next
	(*flatentry)(e.next).prev = e.prev
	e.mark = 0
}

func (f *flatten) llTail(e *flatentry) {
	if e.mark == f.mark {
		f.llDel(e)
		f.llAdd(e)
	}
}

func (f *flatten) llTailBatch(b *lrBatch) {
	f.mu.Lock()
	for e := range b.m {
		f.llTail(e)
	}
	f.mu.Unlock()
	clear(b.m)
}

func (f *flatten) remove(e *flatentry) {
	f.size -= e.size
	f.llDel(e)
	delete(f.cache, e.key)
}

func (f *flatten) Flight(key, cmd string, ttl time.Duration, now time.Time) (RedisMessage, CacheEntry) {
	f.mu.RLock()
	e := f.cache[key]
	f.mu.RUnlock()
	ts := now.UnixMilli()
	if v, _ := e.find(cmd, ts); v.typ != 0 {
		batch := f.lrup.Get().(*lrBatch)
		batch.m[e] = struct{}{}
		if len(batch.m) == lrBatchSize {
			f.llTailBatch(batch)
		}
		f.lrup.Put(batch)
		return v, nil
	}
	fk := key + cmd
	f.mu.RLock()
	af := f.flights[fk]
	f.mu.RUnlock()
	if af != nil {
		return RedisMessage{}, af
	}
	f.mu.Lock()
	defer f.mu.Unlock()
	e = f.cache[key]
	v, expired := e.find(cmd, ts)
	if v.typ != 0 {
		f.llTail(e)
		return v, nil
	}
	if expired {
		f.remove(e)
	}
	if af = f.flights[fk]; af != nil {
		return RedisMessage{}, af
	}
	f.flights[fk] = &adapterEntry{ch: make(chan struct{}), xat: ts + ttl.Milliseconds()}
	return RedisMessage{}, nil
}

func (f *flatten) Update(key, cmd string, val RedisMessage) int64 {
	fk := key + cmd
	bs := val.CacheMarshal(nil)
	fe := &flatentry{cmd: cmd, val: bs, ttl: val.CachePXAT(), size: int64(len(bs)+len(key)+len(cmd)) + int64(unsafe.Sizeof(flatentry{}))}
	f.mu.Lock()
	af := f.flights[fk]
	if af != nil {
		delete(f.flights, fk)
		if af.xat < fe.ttl {
			fe.ttl = af.xat
		}
	}
	f.size += fe.size
	for ep := f.head.next; f.size > f.limit && ep != unsafe.Pointer(f.tail); {
		e := (*flatentry)(ep)
		f.remove(e)
		ep = e.next
	}
	if e := f.cache[key]; e == nil {
		fe.key = key
		f.cache[key] = fe
		f.llAdd(fe)
	} else {
		e.insert(fe)
	}
	f.mu.Unlock()
	if af != nil {
		af.set(val, nil)
	}
	return fe.ttl
}

func (f *flatten) Cancel(key, cmd string, err error) {
	fk := key + cmd
	f.mu.Lock()
	defer f.mu.Unlock()
	if af := f.flights[fk]; af != nil {
		delete(f.flights, fk)
		af.set(RedisMessage{}, err)
	}
}

func (f *flatten) Delete(keys []RedisMessage) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if keys == nil {
		f.cache = make(map[string]*flatentry, len(f.cache))
		f.head.next = unsafe.Pointer(f.tail)
		f.tail.prev = unsafe.Pointer(f.head)
		f.mark++
		f.size = 0
	} else {
		for _, k := range keys {
			if e := f.cache[k.string]; e != nil {
				f.remove(e)
			}
		}
	}
}

func (f *flatten) Close(err error) {
	f.mu.Lock()
	flights := f.flights
	f.flights = nil
	f.cache = nil
	f.tail = nil
	f.head = nil
	f.mark++
	f.mu.Unlock()
	for _, entry := range flights {
		entry.set(RedisMessage{}, err)
	}
}
