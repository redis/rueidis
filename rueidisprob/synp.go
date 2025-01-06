package rueidisprob

import "github.com/redis/rueidis/internal/util"

var bytesPool = util.NewPool(func(capacity int) *bytesContainer {
	return &bytesContainer{s: make([]byte, 0, capacity)}
})

type bytesContainer struct {
	s []byte
}

func (r *bytesContainer) Capacity() int {
	return cap(r.s)
}

func (r *bytesContainer) ResetLen(n int) {
	clear(r.s)
	r.s = r.s[:n]
}
