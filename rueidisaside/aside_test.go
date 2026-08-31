package rueidisaside

import (
	"context"
	"errors"
	"math/rand"
	"strconv"
	"strings"
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

type cancelAfterAcquireClient struct {
	rueidis.Client
	key      string
	cancel   context.CancelFunc
	accepted atomic.Bool
}

func (c *cancelAfterAcquireClient) Do(ctx context.Context, cmd rueidis.Completed) rueidis.RedisResult {
	if isAcquireCommand(cmd.Commands(), c.key) {
		resp := c.Client.Do(ctx, cmd)
		if rueidis.IsRedisNil(resp.Error()) && c.accepted.CompareAndSwap(false, true) {
			c.cancel()
			return rueidis.NewErrorResult(ctx.Err())
		}
		return resp
	}
	return c.Client.Do(ctx, cmd)
}

type blockingAcquireClient struct {
	rueidis.Client
	key     string
	started chan struct{}
	release chan struct{}
	calls   atomic.Int64
}

func (c *blockingAcquireClient) Do(ctx context.Context, cmd rueidis.Completed) rueidis.RedisResult {
	if isAcquireCommand(cmd.Commands(), c.key) {
		if c.calls.Add(1) == 1 {
			close(c.started)
		}
		select {
		case <-c.release:
		case <-ctx.Done():
			return rueidis.NewErrorResult(ctx.Err())
		}
	}
	return c.Client.Do(ctx, cmd)
}

type blockingKeepaliveClient struct {
	rueidis.Client
	started chan struct{}
	release chan struct{}
	calls   atomic.Int64
}

func (c *blockingKeepaliveClient) Do(ctx context.Context, cmd rueidis.Completed) rueidis.RedisResult {
	commands := cmd.Commands()
	if len(commands) >= 2 && commands[0] == "SET" && strings.HasPrefix(commands[1], PlaceholderPrefix) {
		if c.calls.Add(1) == 1 {
			close(c.started)
			select {
			case <-c.release:
			case <-ctx.Done():
				return rueidis.NewErrorResult(ctx.Err())
			}
		}
	}
	return c.Client.Do(ctx, cmd)
}

func isAcquireCommand(commands []string, key string) bool {
	if len(commands) == 7 && commands[0] == "SET" {
		return commands[1] == key && commands[3] == "NX" && commands[4] == "GET"
	}
	if len(commands) == 6 && (commands[0] == "EVALSHA" || commands[0] == "EVAL") {
		return commands[2] == "1" && commands[3] == key
	}
	return false
}

type getResult struct {
	val string
	err error
}

func TestBeginFlightIsPerKey(t *testing.T) {
	c := &Client{
		flights: make(map[string]chan struct{}),
	}
	f, leader := c.beginFlight("key")
	if f == nil || !leader {
		t.Fatal("first caller did not create a flight")
	}
	c.id = "new-generation"
	follower, leader := c.beginFlight("key")
	if follower != f || leader {
		t.Fatal("same key did not join the active flight after the client id changed")
	}
	if len(c.flights) != 1 {
		t.Fatalf("follower changed the flights map: %d entries", len(c.flights))
	}
	if other, leader := c.beginFlight("other-key"); other == nil || !leader {
		t.Fatal("different key did not create an independent flight")
	}
	select {
	case <-f:
		t.Fatal("flight completed before the leader finished")
	default:
	}
	c.finishFlight("key", f)
	select {
	case <-f:
	default:
		t.Fatal("flight did not wake its follower")
	}

	next, leader := c.beginFlight("key")
	if next == f || !leader {
		t.Fatal("completed flight was reused")
	}
	c.finishFlight("key", f)
	if c.flights["key"] != next {
		t.Fatal("old flight removed the current flight")
	}
	select {
	case <-next:
		t.Fatal("old flight closed the current flight")
	default:
	}
	c.finishFlight("key", next)
	select {
	case <-next:
	default:
		t.Fatal("current flight did not wake its follower")
	}
}

func TestAcquireCancellationCleansPlaceholder(t *testing.T) {
	for _, useLuaLock := range []bool{false, true} {
		name := "redis-7"
		if useLuaLock {
			name = "legacy"
		}
		t.Run(name, func(t *testing.T) {
			key := "canceled-acquire-" + strconv.Itoa(rand.Int())
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			var wrapped *cancelAfterAcquireClient
			client, err := NewClient(ClientOption{
				ClientOption: rueidis.ClientOption{InitAddress: addr, DisableAutoPipelining: true, SelectDB: 5},
				ClientTTL:    time.Second,
				UseLuaLock:   useLuaLock,
				ClientBuilder: func(option rueidis.ClientOption) (rueidis.Client, error) {
					client, err := rueidis.NewClient(option)
					if err != nil {
						return nil, err
					}
					wrapped = &cancelAfterAcquireClient{
						Client: client,
						key:    key,
						cancel: cancel,
					}
					return wrapped, nil
				},
			})
			if err != nil {
				t.Fatal(err)
			}
			defer client.Close()
			defer client.Client().Do(context.Background(), client.Client().B().Del().Key(key).Build())

			loaderCalled := false
			_, err = client.Get(ctx, time.Second, key, func(context.Context, string) (string, error) {
				loaderCalled = true
				return "value", nil
			})
			if !errors.Is(err, context.Canceled) {
				t.Fatalf("expected context.Canceled, got %v", err)
			}
			if !wrapped.accepted.Load() {
				t.Fatal("Redis did not accept the lock command")
			}
			if loaderCalled {
				t.Fatal("loader ran after the canceled acquire")
			}

			val, err := client.Client().Do(context.Background(), client.Client().B().Get().Key(key).Build()).ToString()
			if !rueidis.IsRedisNil(err) {
				t.Fatalf("expected cleanup to remove the placeholder, got %q, %v", val, err)
			}
		})
	}
}

func TestSingleflightStartsBeforeKeepalive(t *testing.T) {
	key := "singleflight-keepalive-" + strconv.Itoa(rand.Int())

	var wrapped *blockingKeepaliveClient
	client, err := NewClient(ClientOption{
		ClientOption: rueidis.ClientOption{InitAddress: addr, SelectDB: 5},
		ClientTTL:    5 * time.Second,
		ClientBuilder: func(option rueidis.ClientOption) (rueidis.Client, error) {
			client, err := rueidis.NewClient(option)
			if err != nil {
				return nil, err
			}
			wrapped = &blockingKeepaliveClient{
				Client:  client,
				started: make(chan struct{}),
				release: make(chan struct{}),
			}
			return wrapped, nil
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()
	defer client.Client().Do(context.Background(), client.Client().B().Del().Key(key).Build())
	defer func() {
		select {
		case <-wrapped.release:
		default:
			close(wrapped.release)
		}
	}()

	first := make(chan getResult, 1)
	go func() {
		val, err := client.Get(context.Background(), time.Second, key, func(context.Context, string) (string, error) {
			return "value", nil
		})
		first <- getResult{val: val, err: err}
	}()
	select {
	case <-wrapped.started:
	case <-time.After(time.Second):
		t.Fatal("first keepalive did not start")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	var secondLoaderCalled atomic.Bool
	_, secondErr := client.Get(ctx, time.Second, key, func(context.Context, string) (string, error) {
		secondLoaderCalled.Store(true)
		return "second", nil
	})
	if !errors.Is(secondErr, context.DeadlineExceeded) {
		t.Fatalf("expected follower deadline, got %v", secondErr)
	}
	if secondLoaderCalled.Load() {
		t.Fatal("follower called the loader")
	}
	if calls := wrapped.calls.Load(); calls != 1 {
		t.Fatalf("expected one keepalive while the flight was active, got %d", calls)
	}

	close(wrapped.release)
	select {
	case result := <-first:
		if result.err != nil || result.val != "value" {
			t.Fatalf("leader returned %q, %v", result.val, result.err)
		}
	case <-time.After(time.Second):
		t.Fatal("leader did not finish")
	}
}

func TestSingleflightSerializesAcquire(t *testing.T) {
	for _, useLuaLock := range []bool{false, true} {
		name := "redis-7"
		if useLuaLock {
			name = "legacy"
		}
		t.Run(name, func(t *testing.T) {
			key := "singleflight-" + strconv.Itoa(rand.Int())

			var wrapped *blockingAcquireClient
			client, err := NewClient(ClientOption{
				ClientOption: rueidis.ClientOption{InitAddress: addr, SelectDB: 5},
				ClientTTL:    time.Second,
				UseLuaLock:   useLuaLock,
				ClientBuilder: func(option rueidis.ClientOption) (rueidis.Client, error) {
					client, err := rueidis.NewClient(option)
					if err != nil {
						return nil, err
					}
					wrapped = &blockingAcquireClient{
						Client:  client,
						key:     key,
						started: make(chan struct{}),
						release: make(chan struct{}),
					}
					return wrapped, nil
				},
			})
			if err != nil {
				t.Fatal(err)
			}
			defer client.Close()
			defer client.Client().Do(context.Background(), client.Client().B().Del().Key(key).Build())

			first := make(chan getResult, 1)
			go func() {
				val, err := client.Get(context.Background(), time.Second, key, func(context.Context, string) (string, error) {
					return "value", nil
				})
				first <- getResult{val: val, err: err}
			}()
			select {
			case <-wrapped.started:
			case <-time.After(time.Second):
				t.Fatal("first acquire did not start")
			}

			ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
			defer cancel()
			var secondLoaderCalled atomic.Bool
			_, secondErr := client.Get(ctx, time.Second, key, func(context.Context, string) (string, error) {
				secondLoaderCalled.Store(true)
				return "second", nil
			})
			if !errors.Is(secondErr, context.DeadlineExceeded) {
				t.Fatalf("expected follower deadline, got %v", secondErr)
			}
			if secondLoaderCalled.Load() {
				t.Fatal("follower called the loader")
			}
			if calls := wrapped.calls.Load(); calls != 1 {
				t.Fatalf("expected one acquire while the flight was active, got %d", calls)
			}

			close(wrapped.release)
			select {
			case result := <-first:
				if result.err != nil || result.val != "value" {
					t.Fatalf("leader returned %q, %v", result.val, result.err)
				}
			case <-time.After(time.Second):
				t.Fatal("leader did not finish")
			}
		})
	}
}

func TestSingleflightKeepsDifferentKeysConcurrent(t *testing.T) {
	client := makeClient(t, addr)
	defer client.Close()
	keys := []string{
		"parallel-a-" + strconv.Itoa(rand.Int()),
		"parallel-b-" + strconv.Itoa(rand.Int()),
	}
	started := make(chan string, len(keys))
	release := make(chan struct{})
	var releaseOnce sync.Once
	defer releaseOnce.Do(func() { close(release) })

	results := make(chan getResult, len(keys))
	for _, key := range keys {
		key := key
		go func() {
			val, err := client.Get(context.Background(), time.Second, key, func(context.Context, string) (string, error) {
				started <- key
				<-release
				return key, nil
			})
			results <- getResult{val: val, err: err}
		}()
	}
	for range keys {
		select {
		case <-started:
		case <-time.After(time.Second):
			t.Fatal("different keys were serialized")
		}
	}
	releaseOnce.Do(func() { close(release) })
	for range keys {
		select {
		case result := <-results:
			if result.err != nil {
				t.Fatal(result.err)
			}
		case <-time.After(time.Second):
			t.Fatal("parallel cache fill did not finish")
		}
	}
}

func TestLoaderPanicReleasesFlight(t *testing.T) {
	client := makeClient(t, addr)
	defer client.Close()
	key := "panic-" + strconv.Itoa(rand.Int())

	func() {
		defer func() {
			if recover() == nil {
				t.Fatal("loader did not panic")
			}
		}()
		_, _ = client.Get(context.Background(), time.Second, key, func(context.Context, string) (string, error) {
			panic("loader panic")
		})
	}()

	val, err := client.Get(context.Background(), time.Second, key, func(context.Context, string) (string, error) {
		return "recovered", nil
	})
	if err != nil || val != "recovered" {
		t.Fatalf("flight was not released after panic: %q, %v", val, err)
	}
}
