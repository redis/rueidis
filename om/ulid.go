package om

import (
	"io"
	"math/rand"
	"sync"
	"time"

	"github.com/oklog/ulid/v2"
)

var entropies = sync.Pool{
	New: func() any {
		return ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	},
}

func id() string {
	n := time.Now()
	entropy := entropies.Get().(io.Reader)
	id := ulid.MustNew(ulid.Timestamp(n), entropy)
	entropies.Put(entropy)
	return id.String()
}
