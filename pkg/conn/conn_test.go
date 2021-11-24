package conn

import (
	"testing"
	"time"

	"github.com/rueian/rueidis/internal/cmds"
)

func BenchmarkClientSideCaching(b *testing.B) {
	setup := func(b *testing.B) *Conn {
		c := NewConn("127.0.0.1:6379", Option{CacheSizeEachConn: DefaultCacheBytes})
		if err := c.Dialable(); err != nil {
			panic(err)
		}
		b.SetParallelism(100)
		b.ResetTimer()
		return c
	}
	b.Run("Do", func(b *testing.B) {
		c := setup(b)
		cmd := cmds.NewCompleted([]string{"GET", "a"})
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.Do(cmd)
			}
		})
	})
	b.Run("DoCache", func(b *testing.B) {
		c := setup(b)
		cmd := cmds.Cacheable(cmds.NewCompleted([]string{"GET", "a"}))
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.DoCache(cmd, time.Second*5)
			}
		})
	})
}
