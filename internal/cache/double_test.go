package cache

import (
	"runtime"
	"strconv"
	"testing"
)

func TestDoubleMap(t *testing.T) {
	m := NewDoubleMap[int](8)
	if _, ok := m.Find("1", "2"); ok {
		t.Fatalf("should not find 1 2")
	}
	if v, ok := m.FindOrInsert("1", "a", func() int {
		return 1
	}); ok || v != 1 {
		t.Fatalf("should insert 1 but not found")
	}
	if v, ok := m.FindOrInsert("1", "a", func() int {
		return 2
	}); !ok || v != 1 {
		t.Fatalf("should found 1")
	}
	m.Delete("1", "a")
	if _, ok := m.Find("1", "2"); ok {
		t.Fatalf("should not find 1 2")
	}
	if v, ok := m.FindOrInsert("1", "a", func() int {
		return 2
	}); ok || v != 2 {
		t.Fatalf("should insert 1 but not found")
	}
	if v, ok := m.FindOrInsert("1", "b", func() int {
		return 2
	}); ok || v != 2 {
		t.Fatalf("should insert 1 but not found")
	}
	if v, ok := m.FindOrInsert("2", "b", func() int {
		return 2
	}); ok || v != 2 {
		t.Fatalf("should insert 1 but not found")
	}
	c := 0
	m.Iterate(func(i int) {
		if i != 2 {
			t.Fatalf("should iterate 2")
		}
		c++
	})
	if c != 3 {
		t.Fatalf("should iterate 3 times")
	}
}

func TestDoubleMap_Delete(t *testing.T) {
	m := NewDoubleMap[int](bpsize)
	for i := 0; i < bpsize; i++ {
		m.FindOrInsert(strconv.Itoa(i), "a", func() int {
			return 1
		})
	}
	for i := 0; i < bpsize-1; i++ {
		m.Delete(strconv.Itoa(i), "a")
	}
	m.Delete(strconv.Itoa(bpsize-1), "a")
	runtime.GC()
	runtime.GC()
	m.mu.Lock()
	heads := len(m.ma)
	m.mu.Unlock()
	if heads != 0 {
		t.Fatalf("no shrink")
	}
}

func TestDoubleMap_DeleteGC(t *testing.T) {
	m := NewDoubleMap[int](bpsize)
	for i := 0; i < bpsize; i++ {
		m.FindOrInsert(strconv.Itoa(i), "a", func() int {
			return 1
		})
	}
	for i := 0; i < bpsize-1; i++ {
		m.Delete(strconv.Itoa(i), "a")
	}
	runtime.GC()
	runtime.GC()
	m.mu.Lock()
	heads := len(m.ma)
	m.mu.Unlock()
	if heads != 1 {
		t.Fatalf("no shrink")
	}
}
