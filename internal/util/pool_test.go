package util

import "testing"

type container struct {
	s []int
}

func (c *container) Capacity() int {
	return cap(c.s)
}

func (c *container) ResetLen(n int) {
	c.s = c.s[:n]
}

func TestPool(t *testing.T) {
	p := NewPool(func(capacity int) *container {
		return &container{s: make([]int, 0, capacity)}
	})
	c := p.Get(5, 10)
	if len(c.s) != 5 || cap(c.s) != 10 {
		t.Fatal("wrong length or capacity")
	}
	c.s[0] = 1
	p.Put(c)
	for {
		c = p.Get(5, 10)
		if c.s[0] == 1 {
			break
		}
		c.s[0] = 1
		p.Put(c)
	}
	c = p.Get(5, 20)
	if c.s[0] != 0 {
		t.Fatal("should not use recycled")
	}
}
