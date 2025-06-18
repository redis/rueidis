package rueidis

import (
	"context"
	"errors"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"

	"github.com/redis/rueidis/internal/cmds"
)

const PTTL = 50
const TTL = 100 * time.Millisecond
const Entries = 3

//gocyclo:ignore
func TestLRU(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	setup := func(t *testing.T) *lru {
		store := newLRU(CacheStoreOption{CacheSizeEachConn: entryMinSize * Entries})
		if v, entry := store.Flight("0", "GET", TTL, time.Now()); v.typ != 0 || entry != nil {
			t.Fatalf("got unexpected value from the first Flight: %v %v", v, entry)
		}
		m := RedisMessage{
			typ:    '+',
			bytes:  unsafe.StringData("0"),
			array:  unsafe.SliceData([]RedisMessage{{}}),
			intlen: 1,
		}
		m.setExpireAt(time.Now().Add(PTTL * time.Millisecond).UnixMilli())
		store.Update("0", "GET", m)
		return store.(*lru)
	}

	t.Run("Cache Hit & Expire", func(t *testing.T) {
		lru := setup(t)
		if v, _ := lru.Flight("0", "GET", TTL, time.Now()); v.typ == 0 {
			t.Fatalf("did not get the value from the second Flight")
		} else if v.string() != "0" {
			t.Fatalf("got unexpected value from the second Flight: %v", v)
		}
		time.Sleep(PTTL * time.Millisecond)
		if v, entry := lru.Flight("0", "GET", TTL, time.Now()); v.typ != 0 || entry != nil {
			t.Fatalf("got unexpected value from the Flight after pttl: %v %v", v, entry)
		}
	})

	t.Run("Cache Should Not Expire By PTTL -2", func(t *testing.T) {
		lru := setup(t)
		if v, entry := lru.Flight("1", "GET", TTL, time.Now()); v.typ != 0 || entry != nil {
			t.Fatalf("got unexpected value from the Flight after pttl: %v %v", v, entry)
		}
		m := strmsg('+', "1")
		lru.Update("1", "GET", m)
		if v, _ := lru.Flight("1", "GET", TTL, time.Now()); v.typ == 0 {
			t.Fatalf("did not get the value from the second Flight")
		} else if v.string() != "1" {
			t.Fatalf("got unexpected value from the second Flight: %v", v)
		}
	})

	t.Run("Cache Miss Suppress", func(t *testing.T) {
		count := 5000
		lru := setup(t)
		wg := sync.WaitGroup{}
		wg.Add(count)
		for i := 0; i < count; i++ {
			go func() {
				defer wg.Done()
				if v, _ := lru.Flight("1", "GET", TTL, time.Now()); v.typ != 0 {
					t.Errorf("got unexpected value from the first Flight: %v", v)
				}
				if v, _ := lru.Flight("2", "GET", TTL, time.Now()); v.typ != 0 {
					t.Errorf("got unexpected value from the first Flight: %v", v)
				}
			}()
		}
		wg.Wait()
		lru.mu.RLock()
		store1 := lru.store["1"]
		store2 := lru.store["2"]
		lru.mu.RUnlock()
		if miss := atomic.LoadUint32(&store1.miss); miss != 1 {
			t.Fatalf("unexpected miss count %v", miss)
		}
		if hits := atomic.LoadUint32(&store1.hits); hits != uint32(count-1) {
			t.Fatalf("unexpected hits count %v", hits)
		}
		if miss := atomic.LoadUint32(&store2.miss); miss != 1 {
			t.Fatalf("unexpected miss count %v", miss)
		}
		if hits := atomic.LoadUint32(&store2.hits); hits != uint32(count-1) {
			t.Fatalf("unexpected hits count %v", hits)
		}
	})

	t.Run("Cache Evict", func(t *testing.T) {
		lru := setup(t)
		for i := 1; i <= Entries; i++ {
			lru.Flight(strconv.Itoa(i), "GET", TTL, time.Now())
			m := strmsg('+', strconv.Itoa(i))
			m.setExpireAt(time.Now().Add(PTTL * time.Millisecond).UnixMilli())
			lru.Update(strconv.Itoa(i), "GET", m)
		}
		if v, entry := lru.Flight("1", "GET", TTL, time.Now()); v.typ != 0 {
			t.Fatalf("got evicted value from the first Flight: %v %v", v, entry)
		}
		if v, _ := lru.Flight(strconv.Itoa(Entries), "GET", TTL, time.Now()); v.typ == 0 {
			t.Fatalf("did not get the latest value from the Flight")
		} else if v.string() != strconv.Itoa(Entries) {
			t.Fatalf("got unexpected value from the Flight: %v", v)
		}
	})

	t.Run("Cache Delete", func(t *testing.T) {
		lru := setup(t)
		lru.Delete([]RedisMessage{strmsg(0x0, "0")})
		if v, _ := lru.Flight("0", "GET", TTL, time.Now()); v.typ != 0 {
			t.Fatalf("got unexpected value from the first Flight: %v", v)
		}
	})

	t.Run("Cache Flush", func(t *testing.T) {
		lru := setup(t)
		for i := 1; i < Entries; i++ {
			lru.Flight(strconv.Itoa(i), "GET", TTL, time.Now())
			m := strmsg('+', strconv.Itoa(i))
			lru.Update(strconv.Itoa(i), "GET", m)
		}
		for i := 1; i < Entries; i++ {
			if v, _ := lru.Flight(strconv.Itoa(i), "GET", TTL, time.Now()); v.string() != strconv.Itoa(i) {
				t.Fatalf("got unexpected value before flush all: %v", v)
			}
		}
		lru.Delete(nil)
		for i := 1; i <= Entries; i++ {
			if v, _ := lru.Flight(strconv.Itoa(i), "GET", TTL, time.Now()); v.typ != 0 {
				t.Fatalf("got unexpected value after flush all: %v", v)
			}
		}
	})

	t.Run("Cache Close", func(t *testing.T) {
		lru := setup(t)
		v, entry := lru.Flight("1", "GET", TTL, time.Now())
		if v.typ != 0 || entry != nil {
			t.Fatalf("got unexpected value from the first Flight: %v %v", v, entry)
		}
		v, entry = lru.Flight("1", "GET", TTL, time.Now())
		if v.typ != 0 || entry == nil { // entry should not be nil in second call
			t.Fatalf("got unexpected value from the second Flight: %v %v", v, entry)
		}

		lru.Close(ErrDoCacheAborted)

		if _, err := entry.Wait(context.Background()); err != ErrDoCacheAborted {
			t.Fatalf("got unexpected value after Close: %v", err)
		}

		m := strmsg('+', "this Update should have no effect")
		m.setExpireAt(time.Now().Add(PTTL * time.Millisecond).UnixMilli())
		lru.Update("1", "GET", m)
		for i := 0; i < 2; i++ { // entry should be always nil after the first call if Close
			if v, entry := lru.Flight("1", "GET", TTL, time.Now()); v.typ != 0 || entry != nil {
				t.Fatalf("got unexpected value from the first Flight: %v %v", v, entry)
			}
		}
	})

	t.Run("Cache Cancel", func(t *testing.T) {
		lru := setup(t)
		v, entry := lru.Flight("1", "GET", TTL, time.Now())
		if v.typ != 0 || entry != nil {
			t.Fatalf("got unexpected value from the first Flight: %v %v", v, entry)
		}
		v, entry = lru.Flight("1", "GET", TTL, time.Now())
		if v.typ != 0 || entry == nil { // entry should not be nil in second call
			t.Fatalf("got unexpected value from the second Flight: %v %v", v, entry)
		}
		err := errors.New("any")

		go func() {
			lru.Cancel("1", "GET", err)
		}()

		if _, err2 := entry.Wait(context.Background()); err2 != err {
			t.Fatalf("got unexpected value from the CacheEntry.Wait(): %v %v", err, err2)
		}
	})

	t.Run("GetTTL", func(t *testing.T) {
		lru := setup(t)
		if v := lru.GetTTL("empty", "cmd"); v != -2 {
			t.Fatalf("unexpected %v", v)
		}
		lru.Flight("key", "cmd", time.Second, time.Now())
		m := RedisMessage{typ: 1}
		m.setExpireAt(time.Now().Add(time.Second).UnixMilli())
		lru.Update("key", "cmd", m)
		if v := lru.GetTTL("key", "cmd"); !roughly(v, time.Second) {
			t.Fatalf("unexpected %v", v)
		}
	})

	t.Run("Update Message TTL", func(t *testing.T) {
		t.Run("client side TTL > server side TTL", func(t *testing.T) {
			lru := setup(t)
			lru.Flight("key", "cmd", 2*time.Second, time.Now())
			m := RedisMessage{typ: 1}
			m.setExpireAt(time.Now().Add(time.Second).UnixMilli())
			lru.Update("key", "cmd", m)
			if v, _ := lru.Flight("key", "cmd", 2*time.Second, time.Now()); v.CacheTTL() != 1 {
				t.Fatalf("unexpected %v", v.CacheTTL())
			}
		})
		t.Run("client side TTL < server side TTL", func(t *testing.T) {
			lru := setup(t)
			lru.Flight("key", "cmd", 2*time.Second, time.Now())
			m := RedisMessage{typ: 1}
			m.setExpireAt(time.Now().Add(3 * time.Second).UnixMilli())
			lru.Update("key", "cmd", m)
			if v, _ := lru.Flight("key", "cmd", 2*time.Second, time.Now()); v.CacheTTL() != 2 {
				t.Fatalf("unexpected %v", v.CacheTTL())
			}
		})
		t.Run("no server side TTL -1", func(t *testing.T) {
			lru := setup(t)
			lru.Flight("key", "cmd", 2*time.Second, time.Now())
			m := RedisMessage{typ: 1}
			lru.Update("key", "cmd", m)
			if v, _ := lru.Flight("key", "cmd", 2*time.Second, time.Now()); v.CacheTTL() != 2 {
				t.Fatalf("unexpected %v", v.CacheTTL())
			}
		})
		t.Run("no server side TTL -2", func(t *testing.T) {
			lru := setup(t)
			lru.Flight("key", "cmd", 2*time.Second, time.Now())
			m := RedisMessage{typ: 1}
			lru.Update("key", "cmd", m)
			if v, _ := lru.Flight("key", "cmd", 2*time.Second, time.Now()); v.CacheTTL() != 2 {
				t.Fatalf("unexpected %v", v.CacheTTL())
			}
		})
	})

	t.Run("Batch Cache Hit & Expire", func(t *testing.T) {
		lru := setup(t)
		if v, _ := lru.Flight("0", "GET", TTL, time.Now()); v.typ == 0 {
			t.Fatalf("did not get the value from the second Flight")
		} else if v.string() != "0" {
			t.Fatalf("got unexpected value from the second Flight: %v", v)
		}
		time.Sleep(PTTL * time.Millisecond)
		if v, entry := flights(lru, time.Now(), TTL, "GET", "0"); v.typ != 0 || entry != nil {
			t.Fatalf("got unexpected value from the Flight after pttl: %v %v", v, entry)
		}
	})

	t.Run("Batch Cache Should Not Expire By PTTL -2", func(t *testing.T) {
		lru := setup(t)
		if v, entry := lru.Flight("1", "GET", TTL, time.Now()); v.typ != 0 || entry != nil {
			t.Fatalf("got unexpected value from the Flight after pttl: %v %v", v, entry)
		}
		m := strmsg('+', "1")
		lru.Update("1", "GET", m)
		if v, _ := flights(lru, time.Now(), TTL, "GET", "1"); v.typ == 0 {
			t.Fatalf("did not get the value from the second Flight")
		} else if v.string() != "1" {
			t.Fatalf("got unexpected value from the second Flight: %v", v)
		}
	})

	t.Run("Batch Cache Miss Suppress", func(t *testing.T) {
		count := 5000
		lru := setup(t)
		wg := sync.WaitGroup{}
		wg.Add(count)
		for i := 0; i < count; i++ {
			go func() {
				defer wg.Done()
				if v, _ := flights(lru, time.Now(), TTL, "GET", "1"); v.typ != 0 {
					t.Errorf("got unexpected value from the first Flight: %v", v)
				}
				if v, _ := flights(lru, time.Now(), TTL, "GET", "2"); v.typ != 0 {
					t.Errorf("got unexpected value from the first Flight: %v", v)
				}
			}()
		}
		wg.Wait()
		lru.mu.RLock()
		store1 := lru.store["1"]
		store2 := lru.store["2"]
		lru.mu.RUnlock()
		if miss := atomic.LoadUint32(&store1.miss); miss != 1 {
			t.Fatalf("unexpected miss count %v", miss)
		}
		if hits := atomic.LoadUint32(&store1.hits); hits != uint32(count-1) {
			t.Fatalf("unexpected hits count %v", hits)
		}
		if miss := atomic.LoadUint32(&store2.miss); miss != 1 {
			t.Fatalf("unexpected miss count %v", miss)
		}
		if hits := atomic.LoadUint32(&store2.hits); hits != uint32(count-1) {
			t.Fatalf("unexpected hits count %v", hits)
		}
	})

	t.Run("Batch Cache Evict", func(t *testing.T) {
		lru := setup(t)
		for i := 1; i <= Entries; i++ {
			flights(lru, time.Now(), TTL, "GET", strconv.Itoa(i))
			m := strmsg('+', strconv.Itoa(i))
			m.setExpireAt(time.Now().Add(PTTL * time.Millisecond).UnixMilli())
			lru.Update(strconv.Itoa(i), "GET", m)
		}
		if v, entry := flights(lru, time.Now(), TTL, "GET", "1"); v.typ != 0 {
			t.Fatalf("got evicted value from the first Flight: %v %v", v, entry)
		}
		if v, _ := flights(lru, time.Now(), TTL, "GET", strconv.Itoa(Entries)); v.typ == 0 {
			t.Fatalf("did not get the latest value from the Flight")
		} else if v.string() != strconv.Itoa(Entries) {
			t.Fatalf("got unexpected value from the Flight: %v", v)
		}
	})

	t.Run("Batch Cache Delete", func(t *testing.T) {
		lru := setup(t)
		lru.Delete([]RedisMessage{strmsg(0x0, "0")})
		if v, _ := flights(lru, time.Now(), TTL, "GET", "0"); v.typ != 0 {
			t.Fatalf("got unexpected value from the first Flight: %v", v)
		}
	})

	t.Run("Batch Cache Flush", func(t *testing.T) {
		lru := setup(t)
		for i := 1; i < Entries; i++ {
			flights(lru, time.Now(), TTL, "GET", strconv.Itoa(i))
			m := strmsg('+', strconv.Itoa(i))
			lru.Update(strconv.Itoa(i), "GET", m)
		}
		for i := 1; i < Entries; i++ {
			if v, _ := flights(lru, time.Now(), TTL, "GET", strconv.Itoa(i)); v.string() != strconv.Itoa(i) {
				t.Fatalf("got unexpected value before flush all: %v", v)
			}
		}
		lru.Delete(nil)
		for i := 1; i <= Entries; i++ {
			if v, _ := flights(lru, time.Now(), TTL, "GET", strconv.Itoa(i)); v.typ != 0 {
				t.Fatalf("got unexpected value after flush all: %v", v)
			}
		}
	})

	t.Run("Batch Cache Close", func(t *testing.T) {
		lru := setup(t)
		v, entry := flights(lru, time.Now(), TTL, "GET", "1")
		if v.typ != 0 || entry != nil {
			t.Fatalf("got unexpected value from the first Flight: %v %v", v, entry)
		}
		v, entry = flights(lru, time.Now(), TTL, "GET", "1")
		if v.typ != 0 || entry == nil { // entry should not be nil in second call
			t.Fatalf("got unexpected value from the second Flight: %v %v", v, entry)
		}

		lru.Close(ErrDoCacheAborted)

		if _, err := entry.Wait(context.Background()); err != ErrDoCacheAborted {
			t.Fatalf("got unexpected value after Close: %v", err)
		}

		m := strmsg('+', "this Update should have no effect")
		m.setExpireAt(time.Now().Add(PTTL * time.Millisecond).UnixMilli())
		lru.Update("1", "GET", m)
		for i := 0; i < 2; i++ { // entry should be always nil after the first call if Close
			if v, entry := flights(lru, time.Now(), TTL, "GET", "1"); v.typ != 0 || entry != nil {
				t.Fatalf("got unexpected value from the first Flight: %v %v", v, entry)
			}
		}
	})

	t.Run("Batch Cache Cancel", func(t *testing.T) {
		lru := setup(t)
		v, entry := flights(lru, time.Now(), TTL, "GET", "1")
		if v.typ != 0 || entry != nil {
			t.Fatalf("got unexpected value from the first Flight: %v %v", v, entry)
		}
		v, entry = flights(lru, time.Now(), TTL, "GET", "1")
		if v.typ != 0 || entry == nil { // entry should not be nil in second call
			t.Fatalf("got unexpected value from the second Flight: %v %v", v, entry)
		}
		err := errors.New("any")

		go func() {
			lru.Cancel("1", "GET", err)
		}()

		if _, err2 := entry.Wait(context.Background()); err2 != err {
			t.Fatalf("got unexpected value from the CacheEntry.Wait(): %v %v", err, err2)
		}
	})

	t.Run("Batch GetTTL", func(t *testing.T) {
		lru := setup(t)
		if v := lru.GetTTL("empty", "cmd"); v != -2 {
			t.Fatalf("unexpected %v", v)
		}
		flights(lru, time.Now(), time.Second, "cmd", "key")
		m := RedisMessage{typ: 1}
		m.setExpireAt(time.Now().Add(time.Second).UnixMilli())
		lru.Update("key", "cmd", m)
		if v := lru.GetTTL("key", "cmd"); !roughly(v, time.Second) {
			t.Fatalf("unexpected %v", v)
		}
	})

	t.Run("Batch Update Message TTL", func(t *testing.T) {
		t.Run("client side TTL > server side TTL", func(t *testing.T) {
			lru := setup(t)
			flights(lru, time.Now(), time.Second*2, "cmd", "key")
			m := RedisMessage{typ: 1}
			m.setExpireAt(time.Now().Add(time.Second).UnixMilli())
			lru.Update("key", "cmd", m)
			if v, _ := flights(lru, time.Now(), time.Second*2, "cmd", "key"); v.CacheTTL() != 1 {
				t.Fatalf("unexpected %v", v.CacheTTL())
			}
		})
		t.Run("client side TTL < server side TTL", func(t *testing.T) {
			lru := setup(t)
			flights(lru, time.Now(), time.Second*2, "cmd", "key")
			m := RedisMessage{typ: 1}
			m.setExpireAt(time.Now().Add(3 * time.Second).UnixMilli())
			lru.Update("key", "cmd", m)
			if v, _ := flights(lru, time.Now(), time.Second*2, "cmd", "key"); v.CacheTTL() != 2 {
				t.Fatalf("unexpected %v", v.CacheTTL())
			}
		})
		t.Run("no server side TTL -1", func(t *testing.T) {
			lru := setup(t)
			flights(lru, time.Now(), time.Second*2, "cmd", "key")
			m := RedisMessage{typ: 1}
			lru.Update("key", "cmd", m)
			if v, _ := flights(lru, time.Now(), time.Second*2, "cmd", "key"); v.CacheTTL() != 2 {
				t.Fatalf("unexpected %v", v.CacheTTL())
			}
		})
		t.Run("no server side TTL -2", func(t *testing.T) {
			lru := setup(t)
			flights(lru, time.Now(), time.Second*2, "cmd", "key")
			m := RedisMessage{typ: 1}
			lru.Update("key", "cmd", m)
			if v, _ := flights(lru, time.Now(), time.Second*2, "cmd", "key"); v.CacheTTL() != 2 {
				t.Fatalf("unexpected %v", v.CacheTTL())
			}
		})
	})
}

func flights(lru *lru, now time.Time, ttl time.Duration, args ...string) (RedisMessage, CacheEntry) {
	results := make([]RedisResult, 1)
	entries := make(map[int]CacheEntry, 1)
	lru.Flights(now, commands(ttl, args...), results, entries)
	return results[0].val, entries[0]
}

func commands(ttl time.Duration, args ...string) []CacheableTTL {
	return []CacheableTTL{CT(Cacheable(cmds.NewCompleted(args)), ttl)}
}

func roughly(ttl, expect time.Duration) bool {
	return ttl >= (expect/4) && ttl <= expect
}

func BenchmarkLRU(b *testing.B) {
	lru := newLRU(CacheStoreOption{CacheSizeEachConn: entryMinSize * Entries})
	b.Run("Flight", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				lru.Flight("0", "GET", TTL, time.Now())
			}
		})
	})
	b.Run("Update", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			key := strconv.Itoa(i)
			lru.Flight(key, "GET", TTL, time.Now())
			m := RedisMessage{}
			m.setExpireAt(time.Now().Add(PTTL * time.Millisecond).UnixMilli())
			lru.Update(key, "GET", m)
		}
	})
}

func TestEntry(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	t.Run("Wait", func(t *testing.T) {
		e := cacheEntry{ch: make(chan struct{})}
		err := errors.New("any")
		go func() {
			e.val = RedisMessage{typ: 1}
			e.err = err
			close(e.ch)
		}()
		if v, err2 := e.Wait(context.Background()); v.typ != 1 || err2 != err {
			t.Fatalf("got unexpected value from the Wait: %v %v", v.typ, err)
		}
	})
	t.Run("Wait with cancel", func(t *testing.T) {
		e := cacheEntry{ch: make(chan struct{})}
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		go func() {
			e.val = RedisMessage{typ: 1}
			close(e.ch)
		}()
		if v, err := e.Wait(ctx); v.typ != 1 || err != nil {
			t.Fatalf("got unexpected value from the Wait: %v %v", v.typ, err)
		}
	})
	t.Run("Wait with closed ctx", func(t *testing.T) {
		e := cacheEntry{ch: make(chan struct{})}
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		if v, err := e.Wait(ctx); err != context.Canceled {
			t.Fatalf("got unexpected value from the Wait: %v %v", v.typ, err)
		}
	})
}
