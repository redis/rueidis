package queue

import (
	"sync/atomic"
	"testing"

	"github.com/rueian/rueidis/internal/proto"
	"github.com/rueian/rueidis/internal/queue"
)

type Queue interface {
	PutOne(m []string) chan proto.Result
	NextCmd() []string
	NextResultCh() ([]string, chan proto.Result)
}

func BenchmarkQueue(b *testing.B) {
	bench := func(factory func() Queue) func(b *testing.B) {
		return func(b *testing.B) {
			q := factory()
			stop := int32(0)
			go func() {
				for atomic.LoadInt32(&stop) == 0 {
					q.NextCmd()
					q.NextResultCh()
				}
			}()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					q.PutOne(nil)
				}
			})
			atomic.StoreInt32(&stop, 1)
		}
	}
	b.Run("Ring", bench(func() Queue { return queue.NewRing() }))
	b.Run("NoPad", bench(func() Queue { return NewNoPadRing() }))
	b.Run("Chan", bench(func() Queue { return NewChan() }))
}
