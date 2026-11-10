package rueidis

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"strings"
	"testing"
)

// shortWriteWriter accepts only part of each write without reporting an error,
// mimicking an io.Writer that silently truncates its input.
type shortWriteWriter struct{ n int }

func (w *shortWriteWriter) Write(p []byte) (int, error) {
	if len(p) > w.n {
		return w.n, nil
	}
	return len(p), nil
}

// streamErrWriter reports a real error together with a short write. The writer
// error must be preserved instead of being replaced by io.ErrShortWrite.
type streamErrWriter struct {
	n   int
	err error
}

func (w *streamErrWriter) Write(p []byte) (int, error) {
	if len(p) > w.n {
		return w.n, w.err
	}
	return len(p), w.err
}

// Every branch of streamTo must report io.ErrShortWrite when the destination
// accepts fewer bytes than offered, matching the io.Copy-based blob path.
func TestStreamToShortWrite(t *testing.T) {
	t.Run("blob string already reports short write", func(t *testing.T) {
		w := &shortWriteWriter{n: 4}
		_, err, _ := streamTo(bufio.NewReader(strings.NewReader("$10\r\nHello word\r\n")), w)
		if !errors.Is(err, io.ErrShortWrite) {
			t.Fatalf("blob: got %v, want io.ErrShortWrite", err)
		}
	})
	t.Run("simple string swallows short write", func(t *testing.T) {
		w := &shortWriteWriter{n: 4}
		_, err, _ := streamTo(bufio.NewReader(strings.NewReader("+Hello word\r\n")), w)
		if !errors.Is(err, io.ErrShortWrite) {
			t.Fatalf("simple string: got %v, want io.ErrShortWrite (bytes silently lost)", err)
		}
	})
	t.Run("integer swallows short write", func(t *testing.T) {
		w := &shortWriteWriter{n: 4}
		_, err, _ := streamTo(bufio.NewReader(strings.NewReader(":1234567\r\n")), w)
		if !errors.Is(err, io.ErrShortWrite) {
			t.Fatalf("integer: got %v, want io.ErrShortWrite (bytes silently lost)", err)
		}
	})
	t.Run("writer error is preserved", func(t *testing.T) {
		want := errors.New("destination failed")
		w := &streamErrWriter{n: 4, err: want}
		_, err, _ := streamTo(bufio.NewReader(strings.NewReader("+Hello word\r\n")), w)
		if !errors.Is(err, want) {
			t.Fatalf("got %v, want the original writer error %v", err, want)
		}
	})
	t.Run("full write succeeds", func(t *testing.T) {
		var buf bytes.Buffer
		n, err, clean := streamTo(bufio.NewReader(strings.NewReader("+Hello word\r\n")), &buf)
		if err != nil || n != 10 || !clean || buf.String() != "Hello word" {
			t.Fatalf("got n=%d err=%v clean=%v data=%q", n, err, clean, buf.String())
		}
	})
}
