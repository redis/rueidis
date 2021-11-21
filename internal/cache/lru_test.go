package cache

import (
	"github.com/rueian/rueidis/internal/proto"
	"strconv"
	"testing"
	"time"
)

const PTTL = 50
const TTL = 100 * time.Millisecond
const Entries = 3

func TestLRU(t *testing.T) {

	setup := func(t *testing.T) *LRU {
		lru := NewLRU(EntryMinSize * Entries)
		if v, entry := lru.GetOrPrepare("0", "GET", TTL); v.Type != 0 || entry != nil {
			t.Fatalf("got unexpected value from the first GetOrPrepare: %v %v", v, entry)
		}
		lru.Update("0", "GET", proto.Message{Type: '+', String: "0"}, PTTL)
		return lru
	}

	t.Run("Cache Hit & Expire", func(t *testing.T) {
		lru := setup(t)
		if v, _ := lru.GetOrPrepare("0", "GET", TTL); v.Type == 0 {
			t.Fatalf("did not get the value from the second GetOrPrepare")
		} else if v.String != "0" {
			t.Fatalf("got unexpected value from the second GetOrPrepare: %v", v)
		}
		time.Sleep(TTL)
		if v, entry := lru.GetOrPrepare("0", "GET", TTL); v.Type != 0 || entry != nil {
			t.Fatalf("got unexpected value from the GetOrPrepare after ttl: %v %v", v, entry)
		}
	})

	t.Run("Cache Expire By PTTL -2", func(t *testing.T) {
		lru := setup(t)
		lru.Update("0", "GET", proto.Message{Type: '+', String: "0"}, -2)
		if v, _ := lru.GetOrPrepare("1", "GET", TTL); v.Type != 0 {
			t.Fatalf("got unexpected value from the first GetOrPrepare: %v", v)
		}
	})

	t.Run("Cache Miss", func(t *testing.T) {
		lru := setup(t)
		if v, _ := lru.GetOrPrepare("1", "GET", TTL); v.Type != 0 {
			t.Fatalf("got unexpected value from the first GetOrPrepare: %v", v)
		}
	})

	t.Run("Cache Evict", func(t *testing.T) {
		lru := setup(t)
		for i := 1; i <= Entries; i++ {
			lru.GetOrPrepare(strconv.Itoa(i), "GET", TTL)
			lru.Update(strconv.Itoa(i), "GET", proto.Message{Type: '+', String: strconv.Itoa(i)}, PTTL)
		}
		if v, entry := lru.GetOrPrepare("1", "GET", TTL); v.Type != 0 {
			t.Fatalf("got evicted value from the first GetOrPrepare: %v %v", v, entry)
		}
		if v, _ := lru.GetOrPrepare(strconv.Itoa(Entries), "GET", TTL); v.Type == 0 {
			t.Fatalf("did not get the latest value from the GetOrPrepare")
		} else if v.String != strconv.Itoa(Entries) {
			t.Fatalf("got unexpected value from the GetOrPrepare: %v", v)
		}
	})

	t.Run("Cache Delete", func(t *testing.T) {
		lru := setup(t)
		lru.Delete([]proto.Message{{String: "0"}})
		if v, _ := lru.GetOrPrepare("0", "GET", TTL); v.Type != 0 {
			t.Fatalf("got unexpected value from the first GetOrPrepare: %v", v)
		}
	})

	t.Run("Cache DeleteAll", func(t *testing.T) {
		lru := setup(t)
		lru.DeleteAll()
		if v, _ := lru.GetOrPrepare("0", "GET", TTL); v.Type != 0 {
			t.Fatalf("got unexpected value from the first GetOrPrepare: %v", v)
		}
	})
}

func BenchmarkLRU(b *testing.B) {
	lru := NewLRU(EntryMinSize * Entries)
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
			lru.Update(key, "GET", proto.Message{}, PTTL)
		}
	})
}

func TestEntry(t *testing.T) {
	t.Run("Wait", func(t *testing.T) {
		e := Entry{ch: make(chan struct{}, 1)}
		go func() {
			e.val = proto.Message{Type: 1}
			close(e.ch)
		}()
		if v := e.Wait(); v.Type != 1 {
			t.Fatalf("got unexpected value from the Wait: %v", v.Type)
		}
	})
}
