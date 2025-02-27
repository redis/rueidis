package rueidislimiter_test

import (
	"context"
	"encoding/binary"
	"encoding/hex"
	"math/rand"
	"testing"
	"time"
	"unsafe"

	"github.com/dannotripp/rueidis"
	"github.com/dannotripp/rueidis/rueidislimiter"
)

func setup(t testing.TB) rueidis.Client {
	client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
	if err != nil {
		t.Fatal(err)
	}
	return client
}

func TestRateLimiter(t *testing.T) {
	client := setup(t)
	t.Cleanup(client.Close)

	now := time.Now()
	window := 100 * time.Millisecond
	limiter, err := rueidislimiter.NewRateLimiter(rueidislimiter.RateLimiterOption{
		ClientBuilder: func(option rueidis.ClientOption) (rueidis.Client, error) {
			return client, nil
		},
		Limit:  3,
		Window: window,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Check defaults", func(t *testing.T) {
		limiter, err := rueidislimiter.NewRateLimiter(rueidislimiter.RateLimiterOption{
			ClientBuilder: func(option rueidis.ClientOption) (rueidis.Client, error) {
				return client, nil
			},
		})
		if err != nil {
			t.Fatal(err)
		}
		result, err := limiter.Check(context.Background(), randStr())
		if err != nil {
			t.Fatal(err)
		}
		if !result.Allowed || result.Remaining != 1 || result.ResetAtMs < now.UnixMilli() {
			t.Fatalf("Expected Allowed=true, Remaining=1, ResetAt >= now; got Allowed=%v, Remaining=%v, ResetAt=%v", result.Allowed, result.Remaining, result.ResetAtMs)
		}
	})

	t.Run("Check allowed within limit", func(t *testing.T) {
		result, err := limiter.Check(context.Background(), randStr())
		if err != nil {
			t.Fatal(err)
		}
		if !result.Allowed || result.Remaining != 3 || result.ResetAtMs < now.UnixMilli() {
			t.Fatalf("Expected Allowed=true, Remaining=3, ResetAt >= now; got Allowed=%v, Remaining=%v, ResetAt=%v", result.Allowed, result.Remaining, result.ResetAtMs)
		}
	})

	t.Run("Check denied after exceeding limit", func(t *testing.T) {
		key := randStr()
		generateLoad(t, limiter, key, 3)

		result, err := limiter.Check(context.Background(), key)
		if err != nil {
			t.Fatal(err)
		}
		if result.Allowed || result.Remaining != 0 || result.ResetAtMs < now.UnixMilli() {
			t.Fatalf("Expected Allowed=false, Remaining=0, ResetAt >= now; got Allowed=%v, Remaining=%v, ResetAt=%v", result.Allowed, result.Remaining, result.ResetAtMs)
		}
	})

	t.Run("Check allowed after window reset", func(t *testing.T) {
		key := randStr()
		generateLoad(t, limiter, key, 3)

		// Sleep for slightly longer than window duration to ensure reset
		time.Sleep(window * 2)
		result, err := limiter.Check(context.Background(), key)
		if err != nil {
			t.Fatal(err)
		}
		if !result.Allowed || result.Remaining != 3 || result.ResetAtMs < now.UnixMilli() {
			t.Fatalf("Expected Allowed=true, Remaining=3, ResetAt >= now after reset; got Allowed=%v, Remaining=%v, ResetAt=%v", result.Allowed, result.Remaining, result.ResetAtMs)
		}
	})

	t.Run("Check allowed with limit option", func(t *testing.T) {
		key := randStr()
		generateLoad(t, limiter, key, 3)

		result, err := limiter.Check(context.Background(), key)
		if err != nil {
			t.Fatal(err)
		}
		if result.Allowed {
			t.Fatalf("Expected Allowed=false; got Allowed=%v", result.Allowed)
		}

		result, err = limiter.Check(context.Background(), key, rueidislimiter.WithCustomRateLimit(10, time.Millisecond*100))
		if err != nil {
			t.Fatal(err)
		}
		if !result.Allowed || result.Remaining != 7 || result.ResetAtMs < now.UnixMilli() {
			t.Fatalf("Expected Allowed=true, Remaining=7, ResetAt >= now after reset; got Allowed=%v, Remaining=%v, ResetAt=%v", result.Allowed, result.Remaining, result.ResetAtMs)
		}
	})

	t.Run("AllowN defaults", func(t *testing.T) {
		limiter, err := rueidislimiter.NewRateLimiter(rueidislimiter.RateLimiterOption{
			ClientBuilder: func(option rueidis.ClientOption) (rueidis.Client, error) {
				return client, nil
			},
		})
		if err != nil {
			t.Fatal(err)
		}
		result, err := limiter.AllowN(context.Background(), randStr(), 1)
		if err != nil {
			t.Fatal(err)
		}
		if !result.Allowed || result.Remaining != 0 || result.ResetAtMs < now.UnixMilli() {
			t.Fatalf("Expected Allowed=true, Remaining=0, ResetAt >= now; got Allowed=%v, Remaining=%v, ResetAt=%v", result.Allowed, result.Remaining, result.ResetAtMs)
		}
	})

	t.Run("AllowN with tokens within limit", func(t *testing.T) {
		key := randStr()
		result, err := limiter.AllowN(context.Background(), key, 1)
		if err != nil {
			t.Fatal(err)
		}
		if !result.Allowed || result.Remaining != 2 || result.ResetAtMs < now.UnixMilli() {
			t.Fatalf("Expected Allowed=true, Remaining=2, ResetAt >= now; got Allowed=%v, Remaining=%v, ResetAt=%v", result.Allowed, result.Remaining, result.ResetAtMs)
		}
	})

	t.Run("AllowN denied after exceeding limit", func(t *testing.T) {
		key := randStr()
		generateLoad(t, limiter, key, 3)

		result, err := limiter.AllowN(context.Background(), key, 1)
		if err != nil {
			t.Fatal(err)
		}
		if result.Allowed || result.Remaining != 0 || result.ResetAtMs < now.UnixMilli() {
			t.Fatalf("Expected Allowed=false, Remaining=0, ResetAt >= now; got Allowed=%v, Remaining=%v, ResetAt=%v", result.Allowed, result.Remaining, result.ResetAtMs)
		}
	})

	t.Run("AllowN with zero tokens", func(t *testing.T) {
		key := randStr()
		result, err := limiter.AllowN(context.Background(), key, 0)
		if err != nil {
			t.Fatal(err)
		}
		if !result.Allowed {
			t.Fatalf("Expected Allowed=true when allowing zero tokens, but got false")
		}
	})

	t.Run("AllowN with negative tokens", func(t *testing.T) {
		key := randStr()
		result, err := limiter.AllowN(context.Background(), key, -1)
		if err == nil {
			t.Fatalf("Expected error for negative tokens, but got nil")
		}
		if result.Allowed {
			t.Fatalf("Expected Allowed=false when allowing negative tokens, but got true")
		}
	})
}

func BenchmarkRateLimiter(b *testing.B) {
	client := setup(b)
	defer client.Close()

	limiter, err := rueidislimiter.NewRateLimiter(rueidislimiter.RateLimiterOption{
		ClientBuilder: func(option rueidis.ClientOption) (rueidis.Client, error) {
			return client, nil
		},
	})
	if err != nil {
		b.Fatal(err)
	}
	key := randStr()

	b.ResetTimer()
	b.ReportAllocs()

	b.Run("Check", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			limiter.Check(context.Background(), key)
		}
	})

	b.Run("AllowN", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			limiter.AllowN(context.Background(), key, 1)
		}
	})
}

func generateLoad(t *testing.T, limiter rueidislimiter.RateLimiterClient, key string, n int) {
	for i := 0; i < n; i++ {
		_, err := limiter.Allow(context.Background(), key)
		if err != nil {
			t.Fatal(err)
		}
	}
}

// randStr generates a 24-byte long, random string.
func randStr() string {
	b := make([]byte, 24)
	binary.LittleEndian.PutUint64(b[12:], rand.Uint64())
	binary.LittleEndian.PutUint32(b[20:], rand.Uint32())
	hex.Encode(b, b[12:])

	return unsafe.String(unsafe.SliceData(b), len(b))
}
