package cache

import (
	"testing"
)

func TestChain(t *testing.T) {
	h := chain[int]{}
	if h.empty() != true {
		t.Fatal("chain is not empty")
	}
	if _, ok := h.find("any"); ok {
		t.Fatal("value is found")
	}
	if empty := h.delete("any"); !empty {
		t.Fatal("not empty")
	}
	h.insert("1", 1)
	h.insert("2", 2)
	h.insert("3", 3)
	if v, ok := h.find("1"); !ok || v != 1 {
		t.Fatal("value is not found")
	}
	if v, ok := h.find("2"); !ok || v != 2 {
		t.Fatal("value is not found")
	}
	if v, ok := h.find("3"); !ok || v != 3 {
		t.Fatal("value is not found")
	}
	if empty := h.delete("1"); empty {
		t.Fatal("empty")
	}
	if _, ok := h.find("1"); ok {
		t.Fatal("value is found")
	}
	if v, ok := h.find("2"); !ok || v != 2 {
		t.Fatal("value is not found")
	}
	if v, ok := h.find("3"); !ok || v != 3 {
		t.Fatal("value is not found")
	}
	if empty := h.delete("2"); empty {
		t.Fatal("empty")
	}
	if _, ok := h.find("2"); ok {
		t.Fatal("value is found")
	}
	if v, ok := h.find("3"); !ok || v != 3 {
		t.Fatal("value is not found")
	}
	h.insert("4", 4)
	if v, ok := h.find("3"); !ok || v != 3 {
		t.Fatal("value is not found")
	}
	if v, ok := h.find("4"); !ok || v != 4 {
		t.Fatal("value is not found")
	}
	if empty := h.delete("3"); empty {
		t.Fatal("empty")
	}
	if empty := h.delete("4"); !empty {
		t.Fatal("not empty")
	}
}
