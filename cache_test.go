package rueidis

import (
	"context"
	"errors"
	"testing"
	"time"
)

func test(t *testing.T, storeFn func() CacheStore) {
	t.Run("Flight and Update", func(t *testing.T) {
		var err error
		var now = time.Now()
		var store = storeFn()

		v, e := store.Flight("key", "cmd", time.Millisecond*100, now)
		if v.typ != 0 || e != nil {
			t.Fatal("first flight should return empty RedisMessage and nil CacheEntry")
		}

		v, e = store.Flight("key", "cmd", time.Millisecond*100, now)
		if v.typ != 0 || e == nil {
			t.Fatal("flights before Update should return empty RedisMessage and non-nil CacheEntry")
		}

		store.Delete([]RedisMessage{strmsg('+', "key")}) // Delete should not affect pending CacheEntry

		v2, e2 := store.Flight("key", "cmd", time.Millisecond*100, now)
		if v2.typ != 0 || e != e2 {
			t.Fatal("flights before Update should return empty RedisMessage and the same CacheEntry, not be affected by Delete")
		}

		v = strmsg('+', "val")
		v.setExpireAt(now.Add(time.Second).UnixMilli())
		if pttl := store.Update("key", "cmd", v); pttl < now.Add(90*time.Millisecond).UnixMilli() || pttl > now.Add(100*time.Millisecond).UnixMilli() {
			t.Fatal("Update should return a desired pttl")
		}

		v2, err = e.Wait(context.Background())
		if v2.typ != v.typ || v2.string() != v.string() || err != nil {
			t.Fatal("unexpected cache response")
		}
		if pttl := v2.CachePXAT(); pttl < now.Add(90*time.Millisecond).UnixMilli() || pttl > now.Add(100*time.Millisecond).UnixMilli() {
			t.Fatal("CachePXAT should return a desired pttl")
		}

		v2, _ = store.Flight("key", "cmd", time.Millisecond*100, now)
		if v2.typ != v.typ || v2.string() != v.string() {
			t.Fatal("flights after Update should return updated RedisMessage")
		}
		if pttl := v2.CachePXAT(); pttl < now.Add(90*time.Millisecond).UnixMilli() || pttl > now.Add(100*time.Millisecond).UnixMilli() {
			t.Fatal("CachePXAT should return a desired pttl")
		}

		store.Delete([]RedisMessage{strmsg('+', "key")})
		v, e = store.Flight("key", "cmd", time.Millisecond*100, now)
		if v.typ != 0 || e != nil {
			t.Fatal("flights after Delete should return empty RedisMessage and nil CacheEntry")
		}
	})

	t.Run("Flight and Cancel", func(t *testing.T) {
		var err error
		var now = time.Now()
		var store = storeFn()

		v, e := store.Flight("key", "cmd", time.Millisecond*100, now)
		if v.typ != 0 || e != nil {
			t.Fatal("first flight should return empty RedisMessage and nil CacheEntry")
		}

		v, e = store.Flight("key", "cmd", time.Millisecond*100, now)
		if v.typ != 0 || e == nil {
			t.Fatal("flights before Update should return empty RedisMessage and non-nil CacheEntry")
		}

		store.Delete([]RedisMessage{strmsg('+', "key")}) // Delete should not affect pending CacheEntry

		v2, e2 := store.Flight("key", "cmd", time.Millisecond*100, now)
		if v2.typ != 0 || e != e2 {
			t.Fatal("flights before Update should return empty RedisMessage and the same CacheEntry, not be affected by Delete")
		}

		store.Cancel("key", "cmd", errors.New("err"))

		v2, err = e.Wait(context.Background())
		if err.Error() != "err" {
			t.Fatal("unexpected cache response")
		}

		v, e = store.Flight("key", "cmd", time.Millisecond*100, now)
		if v.typ != 0 || e != nil {
			t.Fatal("flights after Cancel should return empty RedisMessage and nil CacheEntry")
		}
	})

	t.Run("Flight and Delete", func(t *testing.T) {
		var now = time.Now()
		var store = storeFn()

		for _, deletions := range [][]RedisMessage{
			{strmsg('+', "key")},
			nil,
		} {
			store.Flight("key", "cmd1", time.Millisecond*100, now)
			store.Flight("key", "cmd2", time.Millisecond*100, now)
			store.Update("key", "cmd1", strmsg('+', "val"))
			store.Update("key", "cmd2", strmsg('+', "val"))

			store.Delete(deletions)

			if v, e := store.Flight("key", "cmd1", time.Millisecond*100, now); v.typ != 0 || e != nil {
				t.Fatal("flight after delete should return empty RedisMessage and nil CacheEntry")
			}

			if v, e := store.Flight("key", "cmd2", time.Millisecond*100, now); v.typ != 0 || e != nil {
				t.Fatal("flight after delete should return empty RedisMessage and nil CacheEntry")
			}
		}
	})

	t.Run("Flight and TTL", func(t *testing.T) {
		var now = time.Now()
		var store = storeFn()

		v, e := store.Flight("key", "cmd", time.Second, now)
		if v.typ != 0 || e != nil {
			t.Fatal("first flight should return empty RedisMessage and nil CacheEntry")
		}

		v = strmsg('+', "val")
		v.setExpireAt(now.Add(time.Millisecond).UnixMilli())
		store.Update("key", "cmd", v)

		v, e = store.Flight("key", "cmd", time.Second, now.Add(time.Millisecond))
		if v.typ != 0 || e != nil {
			t.Fatal("flight after TTL should return empty RedisMessage and nil CacheEntry")
		}
	})

	t.Run("Flight and Close", func(t *testing.T) {
		var now = time.Now()
		var store = storeFn()

		_, _ = store.Flight("key", "cmd", time.Millisecond*100, now)
		_, e := store.Flight("key", "cmd", time.Millisecond*100, now)

		store.Close(errors.New("err"))

		if _, err := e.Wait(context.Background()); err.Error() != "err" {
			t.Fatal("unexpected cache response")
		}

		_, e = store.Flight("key", "cmd", time.Millisecond*100, now)
		if e != nil {
			t.Fatal("flight after Close should return empty RedisMessage and nil CacheEntry")
		}
	})

	t.Run("Flight timeout", func(t *testing.T) {
		var now = time.Now()
		var store = storeFn()

		_, _ = store.Flight("key", "cmd", time.Millisecond*100, now)
		_, e := store.Flight("key", "cmd", time.Millisecond*100, now)

		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		if _, err := e.Wait(ctx); err != context.Canceled {
			t.Fatal("Wait should honor context")
		}
	})
}

func TestCacheStore(t *testing.T) {
	t.Run("LRUCacheStore", func(t *testing.T) {
		test(t, func() CacheStore {
			return newLRU(CacheStoreOption{CacheSizeEachConn: DefaultCacheBytes})
		})
	})
	t.Run("SimpleCache", func(t *testing.T) {
		test(t, func() CacheStore {
			return NewSimpleCacheAdapter(&simple{store: map[string]RedisMessage{}})
		})
	})
}

type simple struct {
	store map[string]RedisMessage
}

func (s *simple) Get(key string) RedisMessage {
	return s.store[key]
}

func (s *simple) Set(key string, val RedisMessage) {
	s.store[key] = val
}

func (s *simple) Del(key string) {
	delete(s.store, key)
}

func (s *simple) Flush() {
	s.store = nil
}
