package conn

import (
	"testing"

	"github.com/rueian/rueidis/internal/cmds"
	"github.com/rueian/rueidis/internal/proto"
)

func BenchmarkConnCacheMutex(b *testing.B) {
	bench := func(factory func() *Conn, write func(*Conn, cmds.Completed) proto.Result) func(b *testing.B) {
		return func(b *testing.B) {
			conn := factory()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					write(conn, cmds.PingCmd)
				}
			})
			conn.Close()
		}
	}
	b.Run("NoMutex", bench(func() *Conn { return NewConnNoMutex() }, WriteNoMutex))
	b.Run("MutexOnWrite", bench(func() *Conn { return NewConnMutexOnWrite(20, 10) }, WriteWithMutex))
	b.Run("MutexInLoop", bench(func() *Conn { return NewConnMutexInEventLoop(20, 10) }, WriteNoMutex))
}
