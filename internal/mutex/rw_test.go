package mutex

import (
	"sync"
	"sync/atomic"
	"testing"
)

func BenchmarkRWMutex(b *testing.B) {
	b.Run("RWMutex", func(b *testing.B) {
		mu := sync.RWMutex{}
		stop := int32(0)
		go func() {
			for atomic.LoadInt32(&stop) == 0 {
				mu.Lock()
				mu.Unlock()
			}
		}()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				mu.RLock()
				mu.RUnlock()
			}
		})
		atomic.AddInt32(&stop, 1)
	})
	b.Run("Mutex", func(b *testing.B) {
		mu := sync.Mutex{}
		stop := int32(0)
		go func() {
			for atomic.LoadInt32(&stop) == 0 {
				mu.Lock()
				mu.Unlock()
			}
		}()
		for i := 0; i < b.N; i++ {
			mu.Lock()
			mu.Unlock()
		}
		atomic.AddInt32(&stop, 1)
	})
	b.Run("MutexParallel", func(b *testing.B) {
		mu := sync.Mutex{}
		stop := int32(0)
		go func() {
			for atomic.LoadInt32(&stop) == 0 {
				mu.Lock()
				mu.Unlock()
			}
		}()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				mu.Lock()
				mu.Unlock()
			}
		})
		atomic.AddInt32(&stop, 1)
	})
}
