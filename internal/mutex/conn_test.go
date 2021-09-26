package mutex

import "testing"

func BenchmarkConnCacheMutex(b *testing.B) {
	bench := func(factory func() Conn) func(b *testing.B) {
		return func(b *testing.B) {
			conn := factory()
			b.SetParallelism(1000)
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					conn.WriteOne(nil)
				}
			})
			conn.Close()
		}
	}
	b.Run("NoCache", bench(func() Conn {
		return NewConnNoCache()
	}))
	b.Run("MutexOnWrite", bench(func() Conn {
		return NewConnMutexOnWrite(50, 10)
	}))
	b.Run("RWMutexOnWrite", bench(func() Conn {
		return NewConnRWMutexOnWrite(50, 10)
	}))
	b.Run("MutexInLoop", bench(func() Conn {
		return NewConnMutexInEventLoop(50, 10)
	}))
}
