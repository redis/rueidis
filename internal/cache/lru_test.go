package cache

import (
	"runtime"
	"strconv"
	"testing"
)

func TestLRUDoubleMap(t *testing.T) {
	m := NewLRUDoubleMap[int](bpsize, bpsize)
	if _, ok := m.Find("1", "a", 1); ok {
		t.Fatal("should not find 1")
	}
	m.Insert("1", "a", 1, 2, 1, 1)
	m.Insert("1", "b", 1, 2, 1, 2)
	m.Insert("2", "c", 1, 2, 1, 3)
	if v, ok := m.Find("1", "a", 1); !ok || v != 1 {
		t.Fatal("not find 1")
	}
	if v, ok := m.Find("1", "b", 1); !ok || v != 2 {
		t.Fatal("not find 2")
	}
	if v, ok := m.Find("2", "c", 1); !ok || v != 3 {
		t.Fatal("not find 3")
	}
	if _, ok := m.Find("1", "a", 2); ok {
		t.Fatal("should not find")
	}
	if _, ok := m.Find("1", "b", 2); ok {
		t.Fatal("should not find")
	}
	if _, ok := m.Find("2", "c", 2); ok {
		t.Fatal("should not find")
	}
	m.Delete("1")
	if _, ok := m.Find("1", "a", 1); ok {
		t.Fatal("should not find")
	}
	if _, ok := m.Find("1", "b", 1); ok {
		t.Fatal("should not find")
	}
	if v, ok := m.Find("2", "c", 1); !ok || v != 3 {
		t.Fatal("not find 3")
	}
	m.Delete("2")
	m.mu.Lock()
	heads := len(m.ma)
	m.mu.Unlock()
	if heads != 2 {
		t.Fatal("should have 2 heads")
	}

	m.Insert("1", "d", 1, 2, 1, 1)
	m.Insert("1", "e", 1, 2, 1, 2)
	m.Insert("2", "f", 1, 2, 1, 3)
	if v, ok := m.Find("1", "d", 1); !ok || v != 1 {
		t.Fatal("not find 1")
	}
	if v, ok := m.Find("1", "e", 1); !ok || v != 2 {
		t.Fatal("not find 2")
	}
	if v, ok := m.Find("2", "f", 1); !ok || v != 3 {
		t.Fatal("not find 3")
	}
	m.DeleteAll()
	if _, ok := m.Find("1", "d", 1); ok {
		t.Fatal("should not find")
	}
	if _, ok := m.Find("1", "e", 1); ok {
		t.Fatal("should not find")
	}
	if _, ok := m.Find("2", "f", 1); ok {
		t.Fatal("should not find")
	}
}

func TestLRUCache_LRU_1(t *testing.T) {
	m := NewLRUDoubleMap[int](bpsize, bpsize)
	for i := 0; i < bpsize; i++ {
		m.Insert(strconv.Itoa(i), "a", 2, 2, 1, i)
	}
	m.mu.Lock()
	heads := len(m.ma)
	m.mu.Unlock()
	if heads != (bpsize / 2) {
		t.Fatal("should have bpsize/2 heads", heads)
	}
	for i := 0; i < bpsize/2; i++ {
		if _, ok := m.Find(strconv.Itoa(i), "a", 1); ok {
			t.Fatal("should not find")
		}
	}
	for i := bpsize / 2; i < bpsize; i++ {
		if v, ok := m.Find(strconv.Itoa(i), "a", 1); !ok || v != i {
			t.Fatal("not find")
		}
	}
}

func TestLRUCache_LRU_2(t *testing.T) {
	m := NewLRUDoubleMap[int](bpsize*2, bpsize*2)
	for i := 0; i < bpsize*2; i++ {
		m.Insert(strconv.Itoa(i), "a", 1, 2, 1, i)
	}
	m.mu.Lock()
	heads := len(m.ma)
	m.mu.Unlock()
	if heads != (bpsize * 2) {
		t.Fatal("should have bpsize*2 heads", heads)
	}
	for i := 0; i < bpsize; i++ {
		for j := 0; j < 4; j++ {
			if v, ok := m.Find(strconv.Itoa(i), "a", 1); !ok || v != i {
				t.Fatal("not find")
			}
		}
	}
	runtime.GC()
	runtime.GC()
	for i := bpsize * 2; i < bpsize*3; i++ {
		m.Insert(strconv.Itoa(i), "a", 1, 2, 1, i)
	}
	for i := 0; i < bpsize; i++ {
		if v, ok := m.Find(strconv.Itoa(i), "a", 1); !ok || v != i {
			t.Fatal("not find", v, ok)
		}
	}
	for i := bpsize * 1; i < bpsize*2; i++ {
		if _, ok := m.Find(strconv.Itoa(i), "a", 1); ok {
			t.Fatal("should not find")
		}
	}
	for i := bpsize * 2; i < bpsize*3; i++ {
		if v, ok := m.Find(strconv.Itoa(i), "a", 1); !ok || v != i {
			t.Fatal("not find", v, ok)
		}
	}
}

func TestLRUCache_LRU_GC(t *testing.T) {
	m := NewLRUDoubleMap[int](bpsize, bpsize)
	for i := 0; i < bpsize; i++ {
		m.Insert(strconv.Itoa(i), "a", 1, 2, 1, i)
	}
	for j := 0; j < 4; j++ {
		if v, ok := m.Find(strconv.Itoa(bpsize/2), "a", 1); !ok || v != bpsize/2 {
			t.Fatal("not find")
		}
	}
	runtime.GC()
	runtime.GC()
	m.Insert("a", "a", bpsize-1, 2, 1, 0)
	m.mu.Lock()
	heads := len(m.ma)
	total := m.total
	m.mu.Unlock()
	if heads != 2 {
		t.Fatal("should have 2 heads", heads)
	}
	if total != bpsize {
		t.Fatal("should have bpsize", bpsize)
	}
	for i := 0; i < bpsize; i++ {
		if i == bpsize/2 {
			if v, ok := m.Find(strconv.Itoa(i), "a", 1); !ok || v != i {
				t.Fatal("not find")
			}
		} else {
			if _, ok := m.Find(strconv.Itoa(i), "a", 1); ok {
				t.Fatal("should not find")
			}
		}
	}
}

func TestLRUCache_LRU_GC_2(t *testing.T) {
	m := NewLRUDoubleMap[int](bpsize, bpsize)
	for i := 0; i < bpsize; i++ {
		m.Insert(strconv.Itoa(i), "a", 1, 2, 1, i)
	}
	for j := 0; j < 4; j++ {
		if v, ok := m.Find(strconv.Itoa(bpsize/2), "a", 1); !ok || v != bpsize/2 {
			t.Fatal("not find")
		}
	}
	m.Reset()
	runtime.GC()
	runtime.GC()
	m.Insert("a", "a", bpsize-1, 2, 1, 0)
	m.mu.Lock()
	heads := len(m.ma)
	total := m.total
	m.mu.Unlock()
	if heads != 1 {
		t.Fatal("should have 1 heads", heads)
	}
	if total != bpsize-1 {
		t.Fatal("should have bpsize-1", bpsize-1)
	}
}
