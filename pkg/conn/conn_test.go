package conn

import (
	"testing"
	"time"
)

func BenchmarkClientSideCaching(b *testing.B) {
	setup := func(b *testing.B) *Conn {
		c, err := NewConn("127.0.0.1:6379", Option{CacheSize: DefaultCacheBytes})
		if err != nil {
			panic(err)
		}
		b.SetParallelism(100)
		b.ResetTimer()
		return c
	}
	b.Run("Do", func(b *testing.B) {
		c := setup(b)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.Do(c.Cmd.Get().Key("a").Build())
			}
		})
	})
	b.Run("DoCache", func(b *testing.B) {
		c := setup(b)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.DoCache(c.Cmd.Get().Key("a").Cache(), time.Second*5)
			}
		})
	})
}
