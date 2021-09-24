package conn

import (
	"github.com/rueian/rueidis/internal/proto"
	"io"
	"net"
	"testing"

	conn2 "github.com/rueian/rueidis/pkg/conn"
)

func BenchmarkNewConn(b *testing.B) {
	b.Run("Interface", func(b *testing.B) {
		server, client := net.Pipe()
		go io.Copy(server, server)
		conn := NewConn(client)
		m := proto.StringArray{"GET", "a"}
		b.SetParallelism(20)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				conn.Write(m)
			}
		})
	})

	b.Run("Struct", func(b *testing.B) {
		server, client := net.Pipe()
		go io.Copy(server, server)
		conn := NewStructConn(client)
		m := proto.StringArray{"GET", "a"}
		b.SetParallelism(20)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				conn.Write(m)
			}
		})
	})

	b.Run("Flat", func(b *testing.B) {
		server, client := net.Pipe()
		go io.Copy(server, server)
		conn := conn2.NewConn(client)
		m := proto.StringArray{"GET", "a"}
		b.SetParallelism(20)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				conn.Write(m)
			}
		})
		conn.Close()
	})
}
