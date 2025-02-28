package rueidis

import (
	"context"
	"sync"
	"time"

	"github.com/redis/rueidis/internal/cache"
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
	Update(key, cmd string, val RedisMessage, now time.Time) (pxat int64)
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

func (a *adapter) Update(key, cmd string, val RedisMessage, _ time.Time) (sxat int64) {
	a.mu.Lock()
	entries := a.flights[key]
	if flight, ok := entries[cmd].(*adapterEntry); ok {
		sxat = val.getExpireAt()
		if flight.xat < sxat || sxat == 0 {
			sxat = flight.xat
			val.setExpireAt(sxat)
		}
		a.store.Set(key+cmd, val)
		flight.setVal(val)
		entries[cmd] = nil
	}
	a.mu.Unlock()
	return
}

func (a *adapter) Cancel(key, cmd string, err error) {
	a.mu.Lock()
	entries := a.flights[key]
	if flight, ok := entries[cmd].(*adapterEntry); ok {
		flight.setErr(err)
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
				e.(*adapterEntry).setErr(err)
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

func (a *adapterEntry) setVal(val RedisMessage) {
	a.val = val
	close(a.ch)
}

func (a *adapterEntry) setErr(err error) {
	a.err = err
	close(a.ch)
}

func (a *adapterEntry) Wait(ctx context.Context) (RedisMessage, error) {
	ctxCh := ctx.Done()
	if ctxCh == nil {
		<-a.ch
		return a.val, a.err
	}
	select {
	case <-ctxCh:
		return RedisMessage{}, ctx.Err()
	case <-a.ch:
		return a.val, a.err
	}
}

func NewChainedCache(limit int) CacheStore {
	return &chained{
		flights: cache.NewDoubleMap[*adapterEntry](64),
		cache:   cache.NewLRUDoubleMap[[]byte](64, int64(limit)),
	}
}

type chained struct {
	flights *cache.DoubleMap[*adapterEntry]
	cache   *cache.LRUDoubleMap[[]byte]
}

func (f *chained) Flight(key, cmd string, ttl time.Duration, now time.Time) (RedisMessage, CacheEntry) {
	ts := now.UnixMilli()
	if e, ok := f.cache.Find(key, cmd, ts); ok {
		var ret RedisMessage
		_ = ret.CacheUnmarshalView(e)
		return ret, nil
	}
	xat := ts + ttl.Milliseconds()
	if af, ok := f.flights.FindOrInsert(key, cmd, func() *adapterEntry {
		return &adapterEntry{ch: make(chan struct{}), xat: xat}
	}); ok {
		return RedisMessage{}, af
	}
	return RedisMessage{}, nil
}

func (f *chained) Update(key, cmd string, val RedisMessage, now time.Time) (sxat int64) {
	if af, ok := f.flights.Find(key, cmd); ok {
		sxat = val.getExpireAt()
		if af.xat < sxat || sxat == 0 {
			sxat = af.xat
			val.setExpireAt(sxat)
		}
		bs := val.CacheMarshal(nil)
		f.cache.Insert(key, cmd, int64(len(bs)+len(key)+len(cmd))+int64(cache.LRUEntrySize)+64, sxat, now.UnixMilli(), bs)
		f.flights.Delete(key, cmd)
		af.setVal(val)
	}
	return sxat
}

func (f *chained) Cancel(key, cmd string, err error) {
	if af, ok := f.flights.Find(key, cmd); ok {
		f.flights.Delete(key, cmd)
		af.setErr(err)
	}
}

func (f *chained) Delete(keys []RedisMessage) {
	if keys == nil {
		f.cache.Reset()
	} else {
		for _, k := range keys {
			f.cache.Delete(k.string)
		}
	}
}

func (f *chained) Close(err error) {
	f.cache.DeleteAll()
	f.flights.Close(func(entry *adapterEntry) {
		entry.setErr(err)
	})
}
