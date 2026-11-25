package rueidis_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/redis/rueidis"
)

// streamValues is the Streaming Read example from the README.
func streamValues(ctx context.Context, client rueidis.Client, firstKey, secondKey string, dst io.Writer) error {
	s := client.DoMultiStream(ctx,
		client.B().Get().Key(firstKey).Build(),
		client.B().Get().Key(secondKey).Build(),
	)
	if err := s.Error(); err != nil {
		return err // The stream may fail before there are any replies.
	}

	w := &streamDestination{dst: dst}
	for s.HasNext() {
		_, err := s.WriteTo(w)
		if w.err == nil && err != nil && !rueidis.IsRedisNil(err) {
			w.err = err
		}
	}
	return w.err
}

type streamDestination struct {
	dst io.Writer
	err error
}

func (w *streamDestination) Write(p []byte) (int, error) {
	if w.err == nil {
		n, err := w.dst.Write(p)
		if err == nil && n != len(p) {
			err = io.ErrShortWrite
		}
		w.err = err
	}
	// Save the error for the caller, but consume all bytes from Redis.
	// Do not call the destination again once it has failed.
	return len(p), nil
}

type failAfterByteWriter struct {
	buf   bytes.Buffer
	err   error
	calls int
}

func (w *failAfterByteWriter) Write(p []byte) (int, error) {
	w.calls++
	if len(p) == 0 {
		return 0, w.err
	}
	n, _ := w.buf.Write(p[:1])
	return n, w.err
}

func TestStreamingReadExample(t *testing.T) {
	addr := os.Getenv("RUEIDIS_STREAM_TEST_ADDR")
	if addr == "" {
		addr = "127.0.0.1:6379"
	}
	client, err := rueidis.NewClient(rueidis.ClientOption{
		InitAddress:      []string{addr},
		BlockingPoolSize: 1,
		DisableCache:     true,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	prefix := "streaming-read-example-" + strconv.FormatInt(time.Now().UnixNano(), 10)
	firstKey, secondKey := prefix+":a{slot1}", prefix+":b{slot1}"
	setValue := func(key, value string) {
		t.Helper()
		if err := client.Do(ctx, client.B().Set().Key(key).Value(value).Build()).Error(); err != nil {
			t.Fatal(err)
		}
	}
	defer func() {
		if err := client.Do(ctx, client.B().Del().Key(firstKey, secondKey).Build()).Error(); err != nil {
			t.Error(err)
		}
	}()
	setValue(firstKey, "first")
	setValue(secondKey, "second")

	t.Run("both values", func(t *testing.T) {
		var dst bytes.Buffer
		if err := streamValues(ctx, client, firstKey, secondKey, &dst); err != nil || dst.String() != "firstsecond" {
			t.Fatalf("streamValues: got (%q, %v), want (%q, nil)", &dst, err, "firstsecond")
		}
	})

	t.Run("writer failure drains and reuses the connection", func(t *testing.T) {
		connectionID := func() string {
			t.Helper()
			s := client.DoStream(ctx, client.B().ClientId().Build())
			var dst bytes.Buffer
			if _, err := s.WriteTo(&dst); err != nil {
				t.Fatal(err)
			}
			return dst.String()
		}
		before := connectionID()
		writerErr := errors.New("destination is full")
		dst := &failAfterByteWriter{err: writerErr}
		if err := streamValues(ctx, client, firstKey, secondKey, dst); !errors.Is(err, writerErr) || dst.buf.String() != "f" || dst.calls != 1 {
			t.Fatalf("streamValues: got (%q, %v) in %d writes, want (%q, %v) in one write", &dst.buf, err, dst.calls, "f", writerErr)
		}
		if after := connectionID(); after != before {
			t.Fatalf("did not reuse the streaming connection: %s != %s", after, before)
		}
	})

	t.Run("missing first key does not suppress the second", func(t *testing.T) {
		if err := client.Do(ctx, client.B().Del().Key(firstKey).Build()).Error(); err != nil {
			t.Fatal(err)
		}
		var dst bytes.Buffer
		if err := streamValues(ctx, client, firstKey, secondKey, &dst); err != nil || dst.String() != "second" {
			t.Fatalf("streamValues: got (%q, %v), want (%q, nil)", &dst, err, "second")
		}
	})

	t.Run("failure before first reply", func(t *testing.T) {
		canceledCtx, cancel := context.WithCancel(ctx)
		cancel()
		if err := streamValues(canceledCtx, client, firstKey, secondKey, io.Discard); !errors.Is(err, context.Canceled) {
			t.Fatalf("streamValues: got %v, want context.Canceled", err)
		}
		s := client.DoStream(ctx, client.B().Ping().Build())
		var dst bytes.Buffer
		if _, err := s.WriteTo(&dst); err != nil || dst.String() != "PONG" {
			t.Fatalf("streaming PING after cancellation: got (%q, %v)", &dst, err)
		}
	})
}
