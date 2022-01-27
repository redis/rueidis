package rueidis

import (
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

const PTTL = 50
const TTL = 100 * time.Millisecond
const Entries = 3

//gocyclo:ignore
func TestLRU(t *testing.T) {

	setup := func(t *testing.T) *lru {
		lru := newLRU(entryMinSize * Entries)
		if v, entry := lru.GetOrPrepare("0", "GET", TTL); v.typ != 0 || entry != nil {
			t.Fatalf("got unexpected value from the first GetOrPrepare: %v %v", v, entry)
		}
		lru.Update("0", "GET", RedisMessage{typ: '+', string: "0", values: []RedisMessage{{}}}, PTTL)
		return lru
	}

	t.Run("Cache Hit & Expire", func(t *testing.T) {
		lru := setup(t)
		if v, _ := lru.GetOrPrepare("0", "GET", TTL); v.typ == 0 {
			t.Fatalf("did not get the value from the second GetOrPrepare")
		} else if v.string != "0" {
			t.Fatalf("got unexpected value from the second GetOrPrepare: %v", v)
		}
		time.Sleep(PTTL * time.Millisecond)
		if v, entry := lru.GetOrPrepare("0", "GET", TTL); v.typ != 0 || entry != nil {
			t.Fatalf("got unexpected value from the GetOrPrepare after pttl: %v %v", v, entry)
		}
	})

	t.Run("Cache Should Not Expire By PTTL -2", func(t *testing.T) {
		lru := setup(t)
		if v, entry := lru.GetOrPrepare("1", "GET", TTL); v.typ != 0 || entry != nil {
			t.Fatalf("got unexpected value from the GetOrPrepare after pttl: %v %v", v, entry)
		}
		lru.Update("1", "GET", RedisMessage{typ: '+', string: "1"}, -2)
		if v, _ := lru.GetOrPrepare("1", "GET", TTL); v.typ == 0 {
			t.Fatalf("did not get the value from the second GetOrPrepare")
		} else if v.string != "1" {
			t.Fatalf("got unexpected value from the second GetOrPrepare: %v", v)
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
				if v, _ := lru.GetOrPrepare("1", "GET", TTL); v.typ != 0 {
					t.Errorf("got unexpected value from the first GetOrPrepare: %v", v)
				}
				if v, _ := lru.GetOrPrepare("2", "GET", TTL); v.typ != 0 {
					t.Errorf("got unexpected value from the first GetOrPrepare: %v", v)
				}
			}()
		}
		wg.Wait()
		lru.mu.RLock()
		store1 := lru.store["1"]
		store2 := lru.store["2"]
		lru.mu.RUnlock()
		if miss := atomic.LoadUint64(&store1.miss); miss != 1 {
			t.Fatalf("unexpected miss count %v", miss)
		}
		if hits := atomic.LoadUint64(&store1.hits); hits != uint64(count-1) {
			t.Fatalf("unexpected hits count %v", hits)
		}
		if miss := atomic.LoadUint64(&store2.miss); miss != 1 {
			t.Fatalf("unexpected miss count %v", miss)
		}
		if hits := atomic.LoadUint64(&store2.hits); hits != uint64(count-1) {
			t.Fatalf("unexpected hits count %v", hits)
		}
	})

	t.Run("Cache Evict", func(t *testing.T) {
		lru := setup(t)
		for i := 1; i <= Entries; i++ {
			lru.GetOrPrepare(strconv.Itoa(i), "GET", TTL)
			lru.Update(strconv.Itoa(i), "GET", RedisMessage{typ: '+', string: strconv.Itoa(i)}, PTTL)
		}
		if v, entry := lru.GetOrPrepare("1", "GET", TTL); v.typ != 0 {
			t.Fatalf("got evicted value from the first GetOrPrepare: %v %v", v, entry)
		}
		if v, _ := lru.GetOrPrepare(strconv.Itoa(Entries), "GET", TTL); v.typ == 0 {
			t.Fatalf("did not get the latest value from the GetOrPrepare")
		} else if v.string != strconv.Itoa(Entries) {
			t.Fatalf("got unexpected value from the GetOrPrepare: %v", v)
		}
	})

	t.Run("Cache Delete", func(t *testing.T) {
		lru := setup(t)
		lru.Delete([]RedisMessage{{string: "0"}})
		if v, _ := lru.GetOrPrepare("0", "GET", TTL); v.typ != 0 {
			t.Fatalf("got unexpected value from the first GetOrPrepare: %v", v)
		}
	})

	t.Run("Cache Flush", func(t *testing.T) {
		lru := setup(t)
		for i := 1; i < Entries; i++ {
			lru.GetOrPrepare(strconv.Itoa(i), "GET", TTL)
			lru.Update(strconv.Itoa(i), "GET", RedisMessage{typ: '+', string: strconv.Itoa(i)}, -1)
		}
		for i := 1; i < Entries; i++ {
			if v, _ := lru.GetOrPrepare(strconv.Itoa(i), "GET", TTL); v.string != strconv.Itoa(i) {
				t.Fatalf("got unexpected value before flush all: %v", v)
			}
		}
		lru.Delete(nil)
		for i := 1; i <= Entries; i++ {
			if v, _ := lru.GetOrPrepare(strconv.Itoa(i), "GET", TTL); v.typ != 0 {
				t.Fatalf("got unexpected value after flush all: %v", v)
			}
		}
	})

	t.Run("Cache FreeAndClose", func(t *testing.T) {
		lru := setup(t)
		v, entry := lru.GetOrPrepare("1", "GET", TTL)
		if v.typ != 0 || entry != nil {
			t.Fatalf("got unexpected value from the first GetOrPrepare: %v %v", v, entry)
		}
		v, entry = lru.GetOrPrepare("1", "GET", TTL)
		if v.typ != 0 || entry == nil { // entry should not be nil in second call
			t.Fatalf("got unexpected value from the second GetOrPrepare: %v %v", v, entry)
		}

		lru.FreeAndClose(RedisMessage{typ: '-', string: "closed"})

		if resp := entry.Wait(); resp.typ != '-' || resp.string != "closed" {
			t.Fatalf("got unexpected value after FreeAndClose: %v", resp)
		}

		lru.Update("1", "GET", RedisMessage{typ: '+', string: "this Update should have no effect"}, PTTL)

		for i := 0; i < 2; i++ { // entry should be always nil after the first call if FreeAndClose
			if v, entry := lru.GetOrPrepare("1", "GET", TTL); v.typ != 0 || entry != nil {
				t.Fatalf("got unexpected value from the first GetOrPrepare: %v %v", v, entry)
			}
		}
	})
}

func BenchmarkLRU(b *testing.B) {
	lru := newLRU(entryMinSize * Entries)
	b.Run("GetOrPrepare", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				lru.GetOrPrepare("0", "GET", TTL)
			}
		})
	})
	b.Run("Update", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			key := strconv.Itoa(i)
			lru.GetOrPrepare(key, "GET", TTL)
			lru.Update(key, "GET", RedisMessage{}, PTTL)
		}
	})
}

func TestEntry(t *testing.T) {
	t.Run("Wait", func(t *testing.T) {
		e := entry{ch: make(chan struct{}, 1)}
		go func() {
			e.val = RedisMessage{typ: 1}
			close(e.ch)
		}()
		if v := e.Wait(); v.typ != 1 {
			t.Fatalf("got unexpected value from the Wait: %v", v.typ)
		}
	})
}
