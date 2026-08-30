package rueidisaside

import (
	"context"
	"errors"
	"math/rand"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/redis/rueidis"
)

var addr = []string{"127.0.0.1:6379"}

func makeClient(t *testing.T, addr []string) CacheAsideClient {
	client, err := NewClient(ClientOption{
		ClientOption: rueidis.ClientOption{InitAddress: addr, PipelineMultiplex: -1, SelectDB: 5},
		ClientTTL:    time.Second,
	})
	if err != nil {
		t.Fatal(err)
	}
	return client
}

func makeClientWithLuaLock(t *testing.T, addr []string) CacheAsideClient {
	client, err := NewClient(ClientOption{
		UseLuaLock:   true,
		ClientOption: rueidis.ClientOption{InitAddress: addr, PipelineMultiplex: -1, SelectDB: 5},
		ClientTTL:    time.Second,
	})
	if err != nil {
		t.Fatal(err)
	}
	return client
}

func TestClientErr(t *testing.T) {
	if _, err := NewClient(ClientOption{}); err == nil {
		t.Error(err)
	}
}

func TestWithClientBuilder(t *testing.T) {
	var client rueidis.Client
	var pipelineMultiplex int
	var disableAutoPipelining bool
	c, err := NewClient(ClientOption{
		ClientOption: rueidis.ClientOption{InitAddress: addr, PipelineMultiplex: 3, DisableAutoPipelining: true, SelectDB: 5},
		ClientBuilder: func(option rueidis.ClientOption) (_ rueidis.Client, err error) {
			pipelineMultiplex = option.PipelineMultiplex
			disableAutoPipelining = option.DisableAutoPipelining
			client, err = rueidis.NewClient(option)
			return client, err
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	if c.Client() != client {
		t.Fatal("client mismatched")
	}
	if pipelineMultiplex != -1 {
		t.Fatalf("expected PipelineMultiplex -1, got %d", pipelineMultiplex)
	}
	if disableAutoPipelining {
		t.Fatal("expected DisableAutoPipelining false")
	}
}

func TestCacheFilled(t *testing.T) {
	client := makeClient(t, addr)
	defer client.Close()
	key := strconv.Itoa(rand.Int())
	for i := 0; i < 2; i++ {
		val, err := client.Get(context.Background(), time.Millisecond*500, key, func(ctx context.Context, key string) (val string, err error) {
			return "1", nil
		})
		if err != nil || val != "1" {
			t.Fatal(err)
		}
		val, err = client.Get(context.Background(), time.Millisecond*500, key, nil)
		if err != nil || val != "1" {
			t.Fatal(err)
		}
		time.Sleep(time.Millisecond * 600)
		val, err = client.Get(context.Background(), time.Millisecond*500, key, nil) // should miss
		if !rueidis.IsRedisNil(err) {
			t.Fatal(err)
		}
	}
}

func TestCacheFilledLL(t *testing.T) {
	client := makeClientWithLuaLock(t, addr)
	defer client.Close()
	key := strconv.Itoa(rand.Int())
	for i := 0; i < 2; i++ {
		val, err := client.Get(context.Background(), time.Millisecond*500, key, func(ctx context.Context, key string) (val string, err error) {
			return "1", nil
		})
		if err != nil || val != "1" {
			t.Fatal(err)
		}
		val, err = client.Get(context.Background(), time.Millisecond*500, key, nil)
		if err != nil || val != "1" {
			t.Fatal(err)
		}
		time.Sleep(time.Millisecond * 600)
		val, err = client.Get(context.Background(), time.Millisecond*500, key, nil) // should miss
		if !rueidis.IsRedisNil(err) {
			t.Fatal(err)
		}
	}
}

func TestCacheDel(t *testing.T) {
	client := makeClient(t, addr)
	defer client.Close()
	key := strconv.Itoa(rand.Int())
	for i := 0; i < 2; i++ {
		val, err := client.Get(context.Background(), time.Millisecond*500, key, func(ctx context.Context, key string) (val string, err error) {
			return "1", nil
		})
		if err != nil || val != "1" {
			t.Fatal(err)
		}
		val, err = client.Get(context.Background(), time.Millisecond*500, key, nil)
		if err != nil || val != "1" {
			t.Fatal(err)
		}
		if err = client.Del(context.Background(), key); err != nil {
			t.Fatal(err)
		}
		time.Sleep(time.Millisecond * 50)
		val, err = client.Get(context.Background(), time.Millisecond*500, key, nil) // should miss
		if !rueidis.IsRedisNil(err) {
			t.Fatal(err)
		}
	}
}

func TestCacheDelLL(t *testing.T) {
	client := makeClientWithLuaLock(t, addr)
	defer client.Close()
	key := strconv.Itoa(rand.Int())
	for i := 0; i < 2; i++ {
		val, err := client.Get(context.Background(), time.Millisecond*500, key, func(ctx context.Context, key string) (val string, err error) {
			return "1", nil
		})
		if err != nil || val != "1" {
			t.Fatal(err)
		}
		val, err = client.Get(context.Background(), time.Millisecond*500, key, nil)
		if err != nil || val != "1" {
			t.Fatal(err)
		}
		if err = client.Del(context.Background(), key); err != nil {
			t.Fatal(err)
		}
		time.Sleep(time.Millisecond * 50)
		val, err = client.Get(context.Background(), time.Millisecond*500, key, nil) // should miss
		if !rueidis.IsRedisNil(err) {
			t.Fatal(err)
		}
	}
}

func TestClientRefresh(t *testing.T) {
	client := makeClient(t, addr).(*Client)
	defer client.Close()
	key := strconv.Itoa(rand.Int())
	_, _ = client.Get(context.Background(), time.Millisecond*500, key, func(ctx context.Context, key string) (val string, err error) {
		id, err := client.client.Do(context.Background(), client.client.B().Get().Key(key).Build()).ToString()
		if err != nil {
			t.Error(err)
		}
		for i := 0; i < 2; i++ {
			err = client.client.Do(context.Background(), client.client.B().Get().Key(id).Build()).Error()
			if err != nil {
				t.Error(err)
			}
			time.Sleep(client.ttl)
		}
		return "1", nil
	})
}

func TestClientRefreshLL(t *testing.T) {
	client := makeClientWithLuaLock(t, addr).(*Client)
	defer client.Close()
	key := strconv.Itoa(rand.Int())
	_, _ = client.Get(context.Background(), time.Millisecond*500, key, func(ctx context.Context, key string) (val string, err error) {
		id, err := client.client.Do(context.Background(), client.client.B().Get().Key(key).Build()).ToString()
		if err != nil {
			t.Error(err)
		}
		for i := 0; i < 2; i++ {
			err = client.client.Do(context.Background(), client.client.B().Get().Key(id).Build()).Error()
			if err != nil {
				t.Error(err)
			}
			time.Sleep(client.ttl)
		}
		return "1", nil
	})
}

func TestCloseCleanup(t *testing.T) {
	client := makeClient(t, addr).(*Client)
	key := strconv.Itoa(rand.Int())
	ch := make(chan string, 1)
	_, _ = client.Get(context.Background(), time.Millisecond*500, key, func(ctx context.Context, key string) (val string, err error) {
		id, err := client.client.Do(context.Background(), client.client.B().Get().Key(key).Build()).ToString()
		if err != nil {
			t.Error(err)
		}
		err = client.client.Do(context.Background(), client.client.B().Get().Key(id).Build()).Error()
		if err != nil {
			t.Error(err)
		}
		ch <- id
		return "1", nil
	})
	client.Close()
	client = makeClient(t, addr).(*Client)
	defer client.Close()
	err := client.client.Do(context.Background(), client.client.B().Get().Key(<-ch).Build()).Error()
	if !rueidis.IsRedisNil(err) {
		t.Error(err)
	}
}

func TestCloseCleanupLL(t *testing.T) {
	client := makeClientWithLuaLock(t, addr).(*Client)
	key := strconv.Itoa(rand.Int())
	ch := make(chan string, 1)
	_, _ = client.Get(context.Background(), time.Millisecond*500, key, func(ctx context.Context, key string) (val string, err error) {
		id, err := client.client.Do(context.Background(), client.client.B().Get().Key(key).Build()).ToString()
		if err != nil {
			t.Error(err)
		}
		err = client.client.Do(context.Background(), client.client.B().Get().Key(id).Build()).Error()
		if err != nil {
			t.Error(err)
		}
		ch <- id
		return "1", nil
	})
	client.Close()
	client = makeClient(t, addr).(*Client)
	defer client.Close()
	err := client.client.Do(context.Background(), client.client.B().Get().Key(<-ch).Build()).Error()
	if !rueidis.IsRedisNil(err) {
		t.Error(err)
	}
}

func TestWriteCancel(t *testing.T) {
	client := makeClient(t, addr).(*Client)
	defer client.Close()
	key := strconv.Itoa(rand.Int())
	ch := make(chan string, 1)
	ctx, cancel := context.WithCancel(context.Background())
	val, err := client.Get(ctx, time.Millisecond*500, key, func(ctx context.Context, key string) (val string, err error) {
		id, err := client.client.Do(context.Background(), client.client.B().Get().Key(key).Build()).ToString()
		if err != nil {
			t.Error(err)
		}
		cancel()
		ch <- id
		return "1", nil
	})
	if val != "1" {
		t.Fatal(err)
	}
	if err != context.Canceled {
		t.Fatal(err)
	}
	err = client.client.Do(context.Background(), client.client.B().Get().Key(key).Build()).Error()
	if !rueidis.IsRedisNil(err) {
		t.Error(err)
	}
}

func TestWriteCancelLL(t *testing.T) {
	client := makeClientWithLuaLock(t, addr).(*Client)
	defer client.Close()
	key := strconv.Itoa(rand.Int())
	ch := make(chan string, 1)
	ctx, cancel := context.WithCancel(context.Background())
	val, err := client.Get(ctx, time.Millisecond*500, key, func(ctx context.Context, key string) (val string, err error) {
		id, err := client.client.Do(context.Background(), client.client.B().Get().Key(key).Build()).ToString()
		if err != nil {
			t.Error(err)
		}
		cancel()
		ch <- id
		return "1", nil
	})
	if val != "1" {
		t.Fatal(err)
	}
	if err != context.Canceled {
		t.Fatal(err)
	}
	err = client.client.Do(context.Background(), client.client.B().Get().Key(key).Build()).Error()
	if !rueidis.IsRedisNil(err) {
		t.Error(err)
	}
}

func TestTimeout(t *testing.T) {
	client := makeClient(t, addr).(*Client)
	defer client.Close()
	key := strconv.Itoa(rand.Int())
	_, err := client.Get(context.Background(), time.Millisecond*500, key, func(ctx context.Context, key string) (val string, err error) {
		_, err = client.Get(context.Background(), time.Millisecond*500, key, func(ctx context.Context, key string) (val string, err error) {
			return "1", nil
		})
		if err != context.DeadlineExceeded {
			t.Error(err)
		}
		return "", err
	})
	if err != context.DeadlineExceeded {
		t.Fatal(err)
	}
}

func TestTimeoutLL(t *testing.T) {
	client := makeClientWithLuaLock(t, addr).(*Client)
	defer client.Close()
	key := strconv.Itoa(rand.Int())
	_, err := client.Get(context.Background(), time.Millisecond*500, key, func(ctx context.Context, key string) (val string, err error) {
		_, err = client.Get(context.Background(), time.Millisecond*500, key, func(ctx context.Context, key string) (val string, err error) {
			return "1", nil
		})
		if err != context.DeadlineExceeded {
			t.Error(err)
		}
		return "", err
	})
	if err != context.DeadlineExceeded {
		t.Fatal(err)
	}
}

func TestDisconnect(t *testing.T) {
	client := makeClient(t, addr).(*Client)
	testDisconnectWaitsForFlight(t, client)
}

func TestDisconnectLL(t *testing.T) {
	client := makeClientWithLuaLock(t, addr).(*Client)
	testDisconnectWaitsForFlight(t, client)
}

func testDisconnectWaitsForFlight(t *testing.T, client *Client) {
	t.Helper()
	defer client.Close()
	key := strconv.Itoa(rand.Int())
	defer client.client.Do(context.Background(), client.client.B().Del().Key(key).Build())

	leaderErr := errors.New("leader stopped after disconnect")
	leaderReady := make(chan string, 1)
	leaderRelease := make(chan struct{})
	var releaseOnce sync.Once
	releaseLeader := func() {
		releaseOnce.Do(func() { close(leaderRelease) })
	}
	defer releaseLeader()

	leaderResult := make(chan getResult, 1)
	go func() {
		val, err := client.Get(context.Background(), time.Second*5, key, func(ctx context.Context, key string) (string, error) {
			id1, err := client.client.Do(context.Background(), client.client.B().Get().Key(key).Build()).ToString()
			if err != nil {
				return "", err
			}
			client.onInvalidation(nil) // simulate disconnection
			leaderReady <- id1
			<-leaderRelease
			return "", leaderErr
		})
		leaderResult <- getResult{val: val, err: err}
	}()

	var id1 string
	select {
	case id1 = <-leaderReady:
	case <-time.After(time.Second):
		t.Fatal("leader did not reach the loader")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	var followerLoaderCalled atomic.Bool
	_, err := client.Get(ctx, time.Second*5, key, func(context.Context, string) (string, error) {
		followerLoaderCalled.Store(true)
		return "unexpected", nil
	})
	cancel()
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("expected follower to wait for the active flight, got %v", err)
	}
	if followerLoaderCalled.Load() {
		t.Fatal("follower started another loader while the old generation was active")
	}

	releaseLeader()
	select {
	case result := <-leaderResult:
		if !errors.Is(result.err, leaderErr) {
			t.Fatalf("expected leader error, got %q, %v", result.val, result.err)
		}
	case <-time.After(time.Second):
		t.Fatal("leader did not finish")
	}

	var id2 string
	val, err := client.Get(context.Background(), time.Second*5, key, func(context.Context, string) (string, error) {
		var err error
		id2, err = client.client.Do(context.Background(), client.client.B().Get().Key(key).Build()).ToString()
		return "2", err
	})
	if err != nil || val != "2" {
		t.Fatalf("new generation did not populate the cache: %q, %v", val, err)
	}
	if id1 == id2 {
		t.Fatal("client id did not change")
	}

	val, err = client.client.Do(context.Background(), client.client.B().Get().Key(key).Build()).ToString()
	if err != nil || val != "2" {
		t.Fatalf("unexpected cache value: %q, %v", val, err)
	}
	err = client.client.Do(context.Background(), client.client.B().Get().Key(id1).Build()).Error()
	if !rueidis.IsRedisNil(err) {
		t.Fatalf("old client marker still exists: %v", err)
	}
	err = client.client.Do(context.Background(), client.client.B().Get().Key(id2).Build()).Error()
	if err != nil {
		t.Fatalf("new client marker is missing: %v", err)
	}
}

func TestMultipleClient(t *testing.T) {
	clients := make([]CacheAsideClient, 10)
	for i := 0; i < len(clients); i++ {
		clients[i] = makeClient(t, addr)
	}
	defer func() {
		for _, client := range clients {
			client.Close()
		}
	}()
	cnt := 1000
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(len(clients))
		key := strconv.Itoa(rand.Int())
		sum := int64(0)
		for i, c := range clients {
			go func(i int, c CacheAsideClient) {
				defer wg.Done()
				for j := 0; j < cnt; j++ {
					v, err := c.Get(context.Background(), time.Second, key, func(ctx context.Context, key string) (val string, err error) {
						atomic.AddInt64(&sum, 1)
						return "1", nil
					})
					if err != nil || v != "1" {
						t.Error(err)
					}
				}
			}(i, c)
		}
		wg.Wait()
		if atomic.LoadInt64(&sum) != 1 {
			t.Fatalf("unexpected sum")
		}
	}
}

func TestMultipleClientLL(t *testing.T) {
	clients := make([]CacheAsideClient, 10)
	for i := 0; i < len(clients); i++ {
		clients[i] = makeClientWithLuaLock(t, addr)
	}
	defer func() {
		for _, client := range clients {
			client.Close()
		}
	}()
	cnt := 1000
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(len(clients))
		key := strconv.Itoa(rand.Int())
		sum := int64(0)
		for i, c := range clients {
			go func(i int, c CacheAsideClient) {
				defer wg.Done()
				for j := 0; j < cnt; j++ {
					v, err := c.Get(context.Background(), time.Second, key, func(ctx context.Context, key string) (val string, err error) {
						atomic.AddInt64(&sum, 1)
						return "1", nil
					})
					if err != nil || v != "1" {
						t.Error(err)
					}
				}
			}(i, c)
		}
		wg.Wait()
		if atomic.LoadInt64(&sum) != 1 {
			t.Fatalf("unexpected sum")
		}
	}
}

func TestOverrideCacheTTL(t *testing.T) {
	client := makeClient(t, addr)
	defer client.Close()
	key := strconv.Itoa(rand.Int())

	val, err := client.Get(context.Background(), time.Second*5, key, func(ctx context.Context, key string) (val string, err error) {
		OverrideCacheTTL(ctx, time.Millisecond*300)
		return "1", nil
	})
	if err != nil || val != "1" {
		t.Fatal(err)
	}

	val, err = client.Get(context.Background(), time.Second*5, key, nil)
	if err != nil || val != "1" {
		t.Fatal(err)
	}

	time.Sleep(time.Millisecond * 400)
	val, err = client.Get(context.Background(), time.Second*5, key, nil) // should miss
	if !rueidis.IsRedisNil(err) {
		t.Fatal("expected cache miss after overridden TTL expired")
	}
}

func TestOverrideCacheTTLLL(t *testing.T) {
	client := makeClientWithLuaLock(t, addr)
	defer client.Close()
	key := strconv.Itoa(rand.Int())

	val, err := client.Get(context.Background(), time.Second*5, key, func(ctx context.Context, key string) (val string, err error) {
		OverrideCacheTTL(ctx, time.Millisecond*300)
		return "1", nil
	})
	if err != nil || val != "1" {
		t.Fatal(err)
	}

	val, err = client.Get(context.Background(), time.Second*5, key, nil)
	if err != nil || val != "1" {
		t.Fatal(err)
	}

	time.Sleep(time.Millisecond * 400)
	val, err = client.Get(context.Background(), time.Second*5, key, nil) // should miss
	if !rueidis.IsRedisNil(err) {
		t.Fatal("expected cache miss after overridden TTL expired")
	}
}

func TestOverrideCacheTTLNegativeCaching(t *testing.T) {
	client := makeClient(t, addr)
	defer client.Close()
	key := strconv.Itoa(rand.Int())

	val, err := client.Get(context.Background(), time.Second*5, key, func(ctx context.Context, key string) (val string, err error) {
		OverrideCacheTTL(ctx, time.Millisecond*300)
		return "NOT_FOUND", nil
	})
	if err != nil || val != "NOT_FOUND" {
		t.Fatal(err)
	}

	val, err = client.Get(context.Background(), time.Second*5, key, nil)
	if err != nil || val != "NOT_FOUND" {
		t.Fatal(err)
	}

	time.Sleep(time.Millisecond * 400)
	val, err = client.Get(context.Background(), time.Second*5, key, func(ctx context.Context, key string) (val string, err error) {
		return "FOUND", nil
	})
	if err != nil || val != "FOUND" {
		t.Fatal(err)
	}
}

func TestOverrideCacheTTLNegativeCachingLL(t *testing.T) {
	client := makeClientWithLuaLock(t, addr)
	defer client.Close()
	key := strconv.Itoa(rand.Int())

	val, err := client.Get(context.Background(), time.Second*5, key, func(ctx context.Context, key string) (val string, err error) {
		OverrideCacheTTL(ctx, time.Millisecond*300)
		return "NOT_FOUND", nil
	})
	if err != nil || val != "NOT_FOUND" {
		t.Fatal(err)
	}

	val, err = client.Get(context.Background(), time.Second*5, key, nil)
	if err != nil || val != "NOT_FOUND" {
		t.Fatal(err)
	}

	time.Sleep(time.Millisecond * 400)
	val, err = client.Get(context.Background(), time.Second*5, key, func(ctx context.Context, key string) (val string, err error) {
		return "FOUND", nil
	})
	if err != nil || val != "FOUND" {
		t.Fatal(err)
	}
}

func TestGetSkipsContextWithTimeoutWhenParentDeadlineIsTighter(t *testing.T) {
	client := makeClient(t, addr).(*Client)
	defer client.Close()

	key := strconv.Itoa(rand.Int())
	parent, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	parentDone := parent.Done()

	var innerDone <-chan struct{}
	val, err := client.Get(parent, time.Hour, key, func(ctx context.Context, key string) (val string, err error) {
		innerDone = ctx.Done()
		return "v", nil
	})
	if err != nil || val != "v" {
		t.Fatalf("Get returned %q, %v", val, err)
	}
	if innerDone != parentDone {
		t.Fatal("expected ctx.Done() inside fn to equal parent.Done()")
	}
}

func BenchmarkGet(b *testing.B) {
	client, err := NewClient(ClientOption{
		ClientOption: rueidis.ClientOption{InitAddress: addr, PipelineMultiplex: -1, SelectDB: 5},
		ClientTTL:    time.Second,
	})
	if err != nil {
		b.Fatal(err)
	}
	defer client.Close()

	// Populate the key and warm the rueidis client-side cache so subsequent
	// Get calls hit the cache and exercise the rueidisaside fast path only.
	key := "bench-" + strconv.Itoa(rand.Int())
	if _, err := client.Get(context.Background(), time.Minute, key, func(context.Context, string) (string, error) {
		return "v", nil
	}); err != nil {
		b.Fatal(err)
	}
	if _, err := client.Get(context.Background(), time.Minute, key, nil); err != nil {
		b.Fatal(err)
	}

	b.Run("context.Background", func(b *testing.B) {
		ctx := context.Background()

		b.ReportAllocs()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if _, err := client.Get(ctx, time.Minute, key, nil); err != nil {
					b.Fatal(err)
				}
			}
		})
	})

	b.Run("parent.TTL", func(b *testing.B) {
		parent, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		b.ReportAllocs()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if _, err := client.Get(parent, time.Minute, key, nil); err != nil {
					b.Fatal(err)
				}
			}
		})
	})
}
