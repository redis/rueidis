package proto

import (
	"bufio"
	"bytes"
	"io"
	"testing"

	"github.com/rueian/rueidis/pkg/proto"
)

func BenchmarkWriteArray(b *testing.B) {
	w := bufio.NewWriter(io.Discard)
	b.Run("Standard Array", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m := &Array{Val: []Message{&String{Val: "GET"}, &String{Val: "a"}}}
			m.WriteTo(w)
		}
	})
	b.Run("Simple Array", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			proto.WriteCmd(w, []string{"GET", "a"})
		}
	})
}

func BenchmarkReadNext(b *testing.B) {
	b.Run("Interface", func(b *testing.B) {
		buf := bytes.NewBuffer(nil)
		w := bufio.NewWriter(buf)
		for i := 0; i < b.N; i++ {
			proto.WriteCmd(w, []string{"GET", "a"})
		}
		w.Flush()
		r := bufio.NewReader(bytes.NewReader(buf.Bytes()))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			if _, err := ReadNext(r); err != nil {
				panic(err)
			}
		}
	})
	b.Run("Raw", func(b *testing.B) {
		buf := bytes.NewBuffer(nil)
		w := bufio.NewWriter(buf)
		for i := 0; i < b.N; i++ {
			proto.WriteCmd(w, []string{"GET", "a"})
		}
		w.Flush()
		r := bufio.NewReader(bytes.NewReader(buf.Bytes()))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			if _, err := proto.ReadNextMessage(r); err != nil {
				panic(err)
			}
		}
	})
}
