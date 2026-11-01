package rueidis

import (
	"context"
	"net"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
)

// This file holds the benchmarks that .github/scripts/bench.sh tracks (see BENCH there). For a pull request, bench.sh
// builds the base with the head's version of this file, so that base and head run identical benchmark code and only
// the library differs. Keep the file self-contained: it must not use helpers from other _test.go files, because those
// may differ or be missing in the base checkout.

func newBenchmarkClient(b *testing.B) Client {
	client, err := NewClient(ClientOption{
		InitAddress: []string{"127.0.0.1:6379"},
		Dialer:      net.Dialer{KeepAlive: -1},
	})
	if err != nil {
		b.Skipf("skipping live Redis benchmark: %v", err)
	}
	b.Cleanup(func() { client.Close() })
	return client
}

// BenchmarkClient_Do measures client.Do() auto-pipelining throughput across core Redis data structures and commands.
//
// The LPushBLPop case is the exception: each op is an LPUSH followed by a BLPOP, and BLPOP goes through the
// dedicated blocking pool rather than auto-pipelining. Its result therefore includes the LPUSH, so it measures the
// blocking path plus one pipelined write, not BLPOP alone.
func BenchmarkClient_Do(b *testing.B) {
	client := newBenchmarkClient(b)
	ctx := context.Background()

	b.Run("Set", func(b *testing.B) {
		val := strings.Repeat("x", 128)
		b.ReportAllocs()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			i := 0
			for pb.Next() {
				cmd := client.B().Set().Key("bench_set_" + strconv.Itoa(i%1000)).Value(val).Build()
				if err := client.Do(ctx, cmd).Error(); err != nil {
					b.Errorf("unexpected error: %v", err)
				}
				i++
			}
		})
	})

	b.Run("Get", func(b *testing.B) {
		// Seed the key so GET returns a real value. If the SET failed, every GET would return a Nil error
		// and the benchmark would log b.N errors, so fail fast instead.
		if err := client.Do(ctx, client.B().Set().Key("bench_get").Value("val").Build()).Error(); err != nil {
			b.Fatalf("unexpected SET error: %v", err)
		}
		cmd := client.B().Get().Key("bench_get").Build().Pin()
		b.ReportAllocs()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if err := client.Do(ctx, cmd).Error(); err != nil {
					b.Errorf("unexpected error: %v", err)
				}
			}
		})
	})

	b.Run("LPushBLPop", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				// Push an item so BLPOP returns immediately, testing the blocking path without deadlocking.
				// The 1s timeout is only a safety net so a missing item fails fast instead of hanging.
				if err := client.Do(ctx, client.B().Lpush().Key("bench_blpop").Element("item").Build()).Error(); err != nil {
					b.Errorf("unexpected LPUSH error: %v", err)
					return
				}
				if err := client.Do(ctx, client.B().Blpop().Key("bench_blpop").Timeout(1).Build()).Error(); err != nil {
					b.Errorf("unexpected BLPOP error: %v", err)
				}
			}
		})
	})
}

// BenchmarkClient_Do_Parallelism measures client.Do() throughput with a fixed number of concurrent goroutines.
//
// Goroutines=N starts exactly N goroutines. b.SetParallelism(p) would start p*GOMAXPROCS of them, which would make
// the result depend on the runner's CPU count.
func BenchmarkClient_Do_Parallelism(b *testing.B) {
	client := newBenchmarkClient(b)
	ctx := context.Background()
	cmd := client.B().Get().Key("bench_conc").Build().Pin()
	// Seed the key so GET returns a real value; fail fast rather than logging b.N Nil errors.
	if err := client.Do(ctx, client.B().Set().Key("bench_conc").Value("val").Build()).Error(); err != nil {
		b.Fatalf("unexpected SET error: %v", err)
	}

	for _, n := range []int{1, 8, 64} {
		b.Run("Goroutines="+strconv.Itoa(n), func(b *testing.B) {
			total := int64(b.N)
			var next atomic.Int64
			var wg sync.WaitGroup
			b.ReportAllocs()
			b.ResetTimer()
			wg.Add(n)
			for range n {
				go func() {
					defer wg.Done()
					// Claim one op at a time from the shared counter, so exactly b.N ops run across all goroutines.
					for next.Add(1) <= total {
						if err := client.Do(ctx, cmd).Error(); err != nil {
							b.Errorf("unexpected error: %v", err)
						}
					}
				}()
			}
			wg.Wait()
		})
	}
}

// BenchmarkClient_DoMulti measures manual multi-command transaction pipelining throughput using client.DoMulti().
func BenchmarkClient_DoMulti(b *testing.B) {
	client := newBenchmarkClient(b)
	ctx := context.Background()

	b.Run("BatchGet/Count=10", func(b *testing.B) {
		// Seed 10 distinct keys so DoMulti measures real 128-byte replies rather than nil replies.
		val := strings.Repeat("x", 128)
		cmds := make([]Completed, 0, 10)
		for j := 0; j < 10; j++ {
			key := "bench_mget_" + strconv.Itoa(j)
			if err := client.Do(ctx, client.B().Set().Key(key).Value(val).Build()).Error(); err != nil {
				b.Fatalf("unexpected SET error: %v", err)
			}
			cmds = append(cmds, client.B().Get().Key(key).Build().Pin())
		}
		b.ReportAllocs()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				for _, resp := range client.DoMulti(ctx, cmds...) {
					if err := resp.Error(); err != nil {
						b.Errorf("unexpected error: %v", err)
					}
				}
			}
		})
	})
}
