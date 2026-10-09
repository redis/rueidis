package rueidiscompatrate_test

import (
	"context"
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/redis/rueidis"
	"github.com/redis/rueidis/rueidiscompat"
	"github.com/redis/rueidis/rueidiscompatrate"
)

func TestRediscompatrate(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rueidiscompatrate Suite")
}

func getTestClient() rueidis.Client {
	addrs := []string{"127.0.0.1:6379", "127.0.0.1:6378"}
	for _, addr := range addrs {
		client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{addr}})
		if err == nil {
			if err := client.Do(context.Background(), client.B().Ping().Build()).Error(); err == nil {
				return client
			}
			client.Close()
		}
	}
	return nil
}

func getClusterClient() rueidis.Client {
	client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{"127.0.0.1:7010"}})
	if err == nil {
		if err := client.Do(context.Background(), client.B().Ping().Build()).Error(); err == nil {
			return client
		}
		client.Close()
	}
	return nil
}

var _ = Describe("Rueidiscompatrate", func() {
	Describe("Limit Helpers", func() {
		It("formats and initializes PerSecond correctly", func() {
			s := rueidiscompatrate.PerSecond(10)
			Expect(s.Rate).To(Equal(10))
			Expect(s.Burst).To(Equal(10))
			Expect(s.Period).To(Equal(time.Second))
			Expect(s.String()).To(Equal("10 req/s (burst 10)"))
			Expect(s.IsZero()).To(BeFalse())
		})

		It("formats and initializes PerMinute correctly", func() {
			m := rueidiscompatrate.PerMinute(120)
			Expect(m.Rate).To(Equal(120))
			Expect(m.Burst).To(Equal(120))
			Expect(m.Period).To(Equal(time.Minute))
			Expect(m.String()).To(Equal("120 req/m (burst 120)"))
			Expect(m.IsZero()).To(BeFalse())
		})

		It("formats and initializes PerHour correctly", func() {
			h := rueidiscompatrate.PerHour(3600)
			Expect(h.Rate).To(Equal(3600))
			Expect(h.Burst).To(Equal(3600))
			Expect(h.Period).To(Equal(time.Hour))
			Expect(h.String()).To(Equal("3600 req/h (burst 3600)"))
			Expect(h.IsZero()).To(BeFalse())
		})

		It("formats custom limits and detects IsZero", func() {
			custom := rueidiscompatrate.Limit{Rate: 5, Burst: 15, Period: 2 * time.Second}
			Expect(custom.String()).To(Equal("5 req/2s (burst 15)"))

			empty := rueidiscompatrate.Limit{}
			Expect(empty.IsZero()).To(BeTrue())
		})
	})

	Describe("Limiter Constructors & Nil Handling", func() {
		It("creates limiter instances via NewLimiter and NewLimiterFromClient", func() {
			client := getTestClient()
			if client == nil {
				Skip("skipping test: no live Redis instance accessible on 127.0.0.1:6379 or 6378")
			}
			defer client.Close()

			adapter := rueidiscompat.NewAdapter(client)
			l1 := rueidiscompatrate.NewLimiter(adapter)
			Expect(l1).NotTo(BeNil())

			l2 := rueidiscompatrate.NewLimiterFromClient(client)
			Expect(l2).NotTo(BeNil())

			l3 := rueidiscompatrate.NewLimiter(nil)
			Expect(l3).NotTo(BeNil())
		})

		It("returns appropriate errors when client is nil", func() {
			l := rueidiscompatrate.NewLimiter(nil)
			ctx := context.Background()
			limit := rueidiscompatrate.PerSecond(10)

			_, err := l.Allow(ctx, "k", limit)
			Expect(err).To(HaveOccurred())

			_, err = l.AllowN(ctx, "k", limit, 1)
			Expect(err).To(HaveOccurred())

			_, err = l.AllowAtMost(ctx, "k", limit, 1)
			Expect(err).To(HaveOccurred())

			err = l.Reset(ctx, "k")
			Expect(err).To(HaveOccurred())
		})
	})

	Context("Live Redis Rate Limiting Operations", func() {
		var (
			client  rueidis.Client
			limiter *rueidiscompatrate.Limiter
			ctx     context.Context
			testID  string
		)

		BeforeEach(func() {
			client = getTestClient()
			if client == nil {
				Skip("skipping test: no live Redis instance accessible on 127.0.0.1:6379 or 6378")
			}
			limiter = rueidiscompatrate.NewLimiter(rueidiscompat.NewAdapter(client))
			ctx = context.Background()
			testID = fmt.Sprintf("ginkgo_rate_%d", time.Now().UnixNano())
		})

		AfterEach(func() {
			if client != nil {
				_ = limiter.Reset(ctx, testID)
				client.Close()
			}
		})

		It("executes Allow and Reset correctly", func() {
			limit := rueidiscompatrate.PerSecond(10)

			// 1. First request
			res, err := limiter.Allow(ctx, testID, limit)
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Allowed).To(Equal(1))
			Expect(res.Remaining).To(Equal(9))
			Expect(res.RetryAfter).To(Equal(time.Duration(-1)))
			Expect(res.ResetAfter).To(BeNumerically("~", 100*time.Millisecond, 30*time.Millisecond))

			// 2. Reset and request again
			err = limiter.Reset(ctx, testID)
			Expect(err).NotTo(HaveOccurred())

			res, err = limiter.Allow(ctx, testID, limit)
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Allowed).To(Equal(1))
			Expect(res.Remaining).To(Equal(9))
			Expect(res.RetryAfter).To(Equal(time.Duration(-1)))
			Expect(res.ResetAfter).To(BeNumerically("~", 100*time.Millisecond, 30*time.Millisecond))

			// 3. AllowN with 2 tokens
			res, err = limiter.AllowN(ctx, testID, limit, 2)
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Allowed).To(Equal(2))
			Expect(res.Remaining).To(Equal(7))
			Expect(res.RetryAfter).To(Equal(time.Duration(-1)))
			Expect(res.ResetAfter).To(BeNumerically("~", 300*time.Millisecond, 40*time.Millisecond))

			// 4. AllowN with remaining 7 tokens
			res, err = limiter.AllowN(ctx, testID, limit, 7)
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Allowed).To(Equal(7))
			Expect(res.Remaining).To(Equal(0))
			Expect(res.RetryAfter).To(Equal(time.Duration(-1)))
			Expect(res.ResetAfter).To(BeNumerically("~", 999*time.Millisecond, 60*time.Millisecond))

			// 5. AllowN with 1000 tokens (exceeded)
			res, err = limiter.AllowN(ctx, testID, limit, 1000)
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Allowed).To(Equal(0))
			Expect(res.Remaining).To(Equal(0))
			Expect(res.RetryAfter).To(BeNumerically("~", 99*time.Second, 2*time.Second))
			Expect(res.ResetAfter).To(BeNumerically("~", 999*time.Millisecond, 60*time.Millisecond))
		})

		It("peeks without consumption when increment is zero in AllowN", func() {
			limit := rueidiscompatrate.PerSecond(10)

			// Non-existent key
			res, err := limiter.AllowN(ctx, testID, limit, 0)
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Allowed).To(Equal(0))
			Expect(res.Remaining).To(Equal(10))
			Expect(res.RetryAfter).To(Equal(time.Duration(-1)))
			Expect(res.ResetAfter).To(Equal(time.Duration(0)))

			// Consume 1 token
			res, err = limiter.Allow(ctx, testID, limit)
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Allowed).To(Equal(1))
			Expect(res.Remaining).To(Equal(9))
			Expect(res.RetryAfter).To(Equal(time.Duration(-1)))
			Expect(res.ResetAfter).To(BeNumerically("~", 100*time.Millisecond, 30*time.Millisecond))

			// Peek again
			res, err = limiter.AllowN(ctx, testID, limit, 0)
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Allowed).To(Equal(0))
			Expect(res.Remaining).To(Equal(9))
			Expect(res.RetryAfter).To(Equal(time.Duration(-1)))
			Expect(res.ResetAfter).To(BeNumerically("~", 100*time.Millisecond, 30*time.Millisecond))
		})

		It("calculates RetryAfter accurately under high-frequency rates", func() {
			limit := rueidiscompatrate.Limit{
				Rate:   1,
				Period: time.Millisecond,
				Burst:  1,
			}

			for i := 0; i < 100; i++ {
				res, err := limiter.Allow(ctx, testID, limit)
				Expect(err).NotTo(HaveOccurred())

				if res.Allowed > 0 {
					continue
				}

				Expect(int64(res.RetryAfter)).To(BeNumerically("<=", int64(2*time.Millisecond)))
			}
		})

		It("grants partial capacity when calling AllowAtMost", func() {
			limit := rueidiscompatrate.PerSecond(10)

			// 1. Consume 1 token
			res, err := limiter.Allow(ctx, testID, limit)
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Allowed).To(Equal(1))
			Expect(res.Remaining).To(Equal(9))
			Expect(res.RetryAfter).To(Equal(time.Duration(-1)))
			Expect(res.ResetAfter).To(BeNumerically("~", 100*time.Millisecond, 30*time.Millisecond))

			// 2. Consume 2 tokens via AllowAtMost
			res, err = limiter.AllowAtMost(ctx, testID, limit, 2)
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Allowed).To(Equal(2))
			Expect(res.Remaining).To(Equal(7))
			Expect(res.RetryAfter).To(Equal(time.Duration(-1)))
			Expect(res.ResetAfter).To(BeNumerically("~", 300*time.Millisecond, 40*time.Millisecond))

			// 3. Peek with AllowN(0)
			res, err = limiter.AllowN(ctx, testID, limit, 0)
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Allowed).To(Equal(0))
			Expect(res.Remaining).To(Equal(7))
			Expect(res.RetryAfter).To(Equal(time.Duration(-1)))
			Expect(res.ResetAfter).To(BeNumerically("~", 300*time.Millisecond, 40*time.Millisecond))

			// 4. Request 10 with 7 remaining -> grants 7!
			res, err = limiter.AllowAtMost(ctx, testID, limit, 10)
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Allowed).To(Equal(7))
			Expect(res.Remaining).To(Equal(0))
			Expect(res.RetryAfter).To(Equal(time.Duration(-1)))
			Expect(res.ResetAfter).To(BeNumerically("~", 999*time.Millisecond, 60*time.Millisecond))

			// 5. Peek with AllowN(0)
			res, err = limiter.AllowN(ctx, testID, limit, 0)
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Allowed).To(Equal(0))
			Expect(res.Remaining).To(Equal(0))
			Expect(res.RetryAfter).To(Equal(time.Duration(-1)))
			Expect(res.ResetAfter).To(BeNumerically("~", 999*time.Millisecond, 60*time.Millisecond))

			// 6. Request 1000 with 0 remaining via AllowAtMost -> grants 0, RetryAfter ~99ms
			res, err = limiter.AllowAtMost(ctx, testID, limit, 1000)
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Allowed).To(Equal(0))
			Expect(res.Remaining).To(Equal(0))
			Expect(res.RetryAfter).To(BeNumerically("~", 99*time.Millisecond, 30*time.Millisecond))
			Expect(res.ResetAfter).To(BeNumerically("~", 999*time.Millisecond, 60*time.Millisecond))

			// 7. Request 1000 with 0 remaining via AllowN -> rejected, RetryAfter ~99s
			res, err = limiter.AllowN(ctx, testID, limit, 1000)
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Allowed).To(Equal(0))
			Expect(res.Remaining).To(Equal(0))
			Expect(res.RetryAfter).To(BeNumerically("~", 99*time.Second, 2*time.Second))
			Expect(res.ResetAfter).To(BeNumerically("~", 999*time.Millisecond, 60*time.Millisecond))
		})

		It("peeks without consumption when increment is zero in AllowAtMost", func() {
			limit := rueidiscompatrate.PerSecond(10)

			// Non-existent key
			res, err := limiter.AllowAtMost(ctx, testID, limit, 0)
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Allowed).To(Equal(0))
			Expect(res.Remaining).To(Equal(10))
			Expect(res.RetryAfter).To(Equal(time.Duration(-1)))
			Expect(res.ResetAfter).To(Equal(time.Duration(0)))

			// Consume 1 token
			res, err = limiter.Allow(ctx, testID, limit)
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Allowed).To(Equal(1))
			Expect(res.Remaining).To(Equal(9))
			Expect(res.RetryAfter).To(Equal(time.Duration(-1)))
			Expect(res.ResetAfter).To(BeNumerically("~", 100*time.Millisecond, 30*time.Millisecond))

			// Peek again
			res, err = limiter.AllowAtMost(ctx, testID, limit, 0)
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Allowed).To(Equal(0))
			Expect(res.Remaining).To(Equal(9))
			Expect(res.RetryAfter).To(Equal(time.Duration(-1)))
			Expect(res.ResetAfter).To(BeNumerically("~", 100*time.Millisecond, 30*time.Millisecond))
		})
	})

	Describe("High Concurrency Contention", func() {
		It("accurately throttles 100 simultaneous concurrent workers", func() {
			client := getTestClient()
			if client == nil {
				Skip("skipping test: no live Redis instance accessible")
			}
			defer client.Close()

			limiter := rueidiscompatrate.NewLimiterFromClient(client)
			ctx := context.Background()
			testID := fmt.Sprintf("rueidiscompatrate_concur_%d", time.Now().UnixNano())
			defer func() {
				_ = limiter.Reset(ctx, testID)
			}()

			limit := rueidiscompatrate.Limit{
				Rate:   20,
				Burst:  20,
				Period: 10 * time.Second,
			}

			const numWorkers = 100
			var allowedCount int64
			var rejectedCount int64

			startBarrier := make(chan struct{})
			doneCh := make(chan struct{}, numWorkers)

			for i := 0; i < numWorkers; i++ {
				go func() {
					<-startBarrier
					res, err := limiter.Allow(ctx, testID, limit)
					if err == nil {
						if res.Allowed > 0 {
							atomic.AddInt64(&allowedCount, int64(res.Allowed))
							Expect(res.RetryAfter).To(Equal(time.Duration(-1)))
						} else {
							atomic.AddInt64(&rejectedCount, 1)
							Expect(res.RetryAfter).To(BeNumerically(">", 0))
						}
					}
					doneCh <- struct{}{}
				}()
			}

			close(startBarrier)
			for i := 0; i < numWorkers; i++ {
				<-doneCh
			}

			Expect(allowedCount).To(Equal(int64(20)))
			Expect(rejectedCount).To(Equal(int64(80)))
		})
	})

	Describe("Redis Cluster Integration", func() {
		var (
			clusterClient rueidis.Client
			limiter       *rueidiscompatrate.Limiter
			ctx           context.Context
		)

		BeforeEach(func() {
			clusterClient = getClusterClient()
			if clusterClient == nil {
				Skip("skipping test: Redis Cluster not accessible on 127.0.0.1:7010")
			}
			limiter = rueidiscompatrate.NewLimiterFromClient(clusterClient)
			ctx = context.Background()
		})

		AfterEach(func() {
			if clusterClient != nil {
				clusterClient.Close()
			}
		})

		It("operates across multiple hash slots without CROSSSLOT errors", func() {
			limit := rueidiscompatrate.PerSecond(10)

			testKeys := []string{
				fmt.Sprintf("rueidiscompatrate_cluster_plain_1_%d", time.Now().UnixNano()),
				fmt.Sprintf("rueidiscompatrate_cluster_plain_2_%d", time.Now().UnixNano()),
				fmt.Sprintf("{slot_tenant_A}:api_calls_%d", time.Now().UnixNano()),
				fmt.Sprintf("{slot_tenant_B}:api_calls_%d", time.Now().UnixNano()),
				fmt.Sprintf("{slot_tenant_C}:api_calls_%d", time.Now().UnixNano()),
			}

			for _, key := range testKeys {
				defer func(k string) { _ = limiter.Reset(ctx, k) }(key)

				// 1. Allow 1
				res, err := limiter.Allow(ctx, key, limit)
				Expect(err).NotTo(HaveOccurred(), "cluster allow on %s should not error", key)
				Expect(res.Allowed).To(Equal(1))
				Expect(res.Remaining).To(Equal(9))
				Expect(res.RetryAfter).To(Equal(time.Duration(-1)))

				// 2. AllowN 3
				resN, err := limiter.AllowN(ctx, key, limit, 3)
				Expect(err).NotTo(HaveOccurred())
				Expect(resN.Allowed).To(Equal(3))
				Expect(resN.Remaining).To(Equal(6))

				// 3. AllowAtMost 10 (grants remaining 6)
				resAtMost, err := limiter.AllowAtMost(ctx, key, limit, 10)
				Expect(err).NotTo(HaveOccurred())
				Expect(resAtMost.Allowed).To(Equal(6))
				Expect(resAtMost.Remaining).To(Equal(0))

				// 4. Over limit
				resOver, err := limiter.Allow(ctx, key, limit)
				Expect(err).NotTo(HaveOccurred())
				Expect(resOver.Allowed).To(Equal(0))
				Expect(resOver.RetryAfter).To(BeNumerically(">", 0))

				// 5. Reset
				err = limiter.Reset(ctx, key)
				Expect(err).NotTo(HaveOccurred())

				resRestored, err := limiter.Allow(ctx, key, limit)
				Expect(err).NotTo(HaveOccurred())
				Expect(resRestored.Allowed).To(Equal(1))
				Expect(resRestored.Remaining).To(Equal(9))
			}
		})
	})
})

func BenchmarkAllow(b *testing.B) {
	client := getTestClient()
	if client == nil {
		b.Skip("cannot connect to live Redis instance")
	}
	defer client.Close()

	l := rueidiscompatrate.NewLimiterFromClient(client)
	ctx := context.Background()
	limit := rueidiscompatrate.PerSecond(1e6)

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		res, err := l.Allow(ctx, "bench_allow", limit)
		if err != nil {
			b.Fatal(err)
		}
		if res.Allowed == 0 {
			b.Fatal("rate limit exceeded during benchmark")
		}
	}
}

func BenchmarkAllowN(b *testing.B) {
	client := getTestClient()
	if client == nil {
		b.Skip("cannot connect to live Redis instance")
	}
	defer client.Close()

	l := rueidiscompatrate.NewLimiterFromClient(client)
	ctx := context.Background()
	limit := rueidiscompatrate.PerSecond(1e6)

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		res, err := l.AllowN(ctx, "bench_allown", limit, 1)
		if err != nil {
			b.Fatal(err)
		}
		if res.Allowed == 0 {
			b.Fatal("rate limit exceeded during benchmark")
		}
	}
}

func BenchmarkAllowAtMost(b *testing.B) {
	client := getTestClient()
	if client == nil {
		b.Skip("cannot connect to live Redis instance")
	}
	defer client.Close()

	l := rueidiscompatrate.NewLimiterFromClient(client)
	ctx := context.Background()
	limit := rueidiscompatrate.PerSecond(1e6)

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		res, err := l.AllowAtMost(ctx, "bench_atmost", limit, 1)
		if err != nil {
			b.Fatal(err)
		}
		if res.Allowed == 0 {
			b.Fatal("rate limit exceeded during benchmark")
		}
	}
}

func BenchmarkAllow_Parallel(b *testing.B) {
	client := getTestClient()
	if client == nil {
		b.Skip("cannot connect to live Redis instance")
	}
	defer client.Close()

	l := rueidiscompatrate.NewLimiterFromClient(client)
	limit := rueidiscompatrate.PerSecond(1e6)

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		ctx := context.Background()
		for pb.Next() {
			res, err := l.Allow(ctx, "bench_allow_par", limit)
			if err != nil {
				b.Fatal(err)
			}
			if res.Allowed == 0 {
				b.Fatal("rate limit exceeded during benchmark")
			}
		}
	})
}

