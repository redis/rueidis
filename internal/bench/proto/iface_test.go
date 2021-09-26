package proto

import (
	"bufio"
	"bytes"
	"io"
	"testing"

	"github.com/rueian/rueidis/internal/proto"
)

func BenchmarkWriteArray(b *testing.B) {
	w := bufio.NewWriter(io.Discard)
	b.Run("Interface", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = (&Array{Val: []Message{&String{Val: "GET"}, &String{Val: "a"}}}).WriteTo(w)
		}
	})
	b.Run("PureSlice", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = proto.WriteCmd(w, []string{"GET", "a"})
		}
	})
}

func BenchmarkReadNext(b *testing.B) {
	prepare := func(n int) *bufio.Reader {
		buf := bytes.NewBuffer(nil)
		w := bufio.NewWriter(buf)
		for i := 0; i < n; i++ {
			_ = proto.WriteCmd(w, []string{"GET", "a"})
		}
		_ = w.Flush()
		return bufio.NewReader(buf)
	}
	b.Run("Interface", func(b *testing.B) {
		r := prepare(b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			if _, err := ReadNextInterfaceMessage(r); err != nil {
				panic(err)
			}
		}
	})
	b.Run("Struct", func(b *testing.B) {
		r := prepare(b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			if _, err := proto.ReadNextMessage(r); err != nil {
				panic(err)
			}
		}
	})
}
