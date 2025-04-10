package rueidislimiter

import "github.com/redis/rueidis/internal/util"

var rateBuffersPool = util.NewPool(func(capacity int) *rateBuffersContainer {
	return &rateBuffersContainer{
		keyBuf: make([]byte, 0, capacity),
	}
})

type rateBuffersContainer struct {
	keyBuf []byte
}

func (r *rateBuffersContainer) Capacity() int {
	return cap(r.keyBuf)
}

func (r *rateBuffersContainer) ResetLen(n int) {
	r.keyBuf = r.keyBuf[:0]
}
