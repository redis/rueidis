package queue

import (
	"sync/atomic"
	"testing"
)

func BenchmarkNewRing(b *testing.B) {
	bench := func(factory func() Queue) func(b *testing.B) {
		return func(b *testing.B) {
			q := factory()
			stop := int32(0)
			go func() {
				for atomic.LoadInt32(&stop) == 0 {
					q.Next1()
					q.Next2()
				}
			}()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					q.Put(&Task{})
				}
			})
			atomic.StoreInt32(&stop, 1)
		}
	}
	b.Run("Ring", bench(func() Queue { return NewRing(8192) }))
	b.Run("Chan", bench(func() Queue { return NewChan(8192) }))
}
