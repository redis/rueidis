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

type cancelAfterAcquireClient struct {
	rueidis.Client
	key      string
	cancel   context.CancelFunc
	accepted bool
}

func (c *cancelAfterAcquireClient) Do(ctx context.Context, cmd rueidis.Completed) rueidis.RedisResult {
	commands := cmd.Commands()
	if isAcquireCommand(commands, c.key) {
		resp := c.Client.Do(ctx, cmd)
		if !rueidis.IsRedisNil(resp.Error()) {
			return resp
		}
		c.accepted = true
		c.cancel()
		if err := ctx.Err(); err != nil {
			return rueidis.NewErrorResult(err)
		}
		return resp
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
	c, err := NewClient(ClientOption{
		ClientOption: rueidis.ClientOption{InitAddress: addr, SelectDB: 5},
		ClientBuilder: func(option rueidis.ClientOption) (_ rueidis.Client, err error) {
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

func TestAcquireCancelCleanup(t *testing.T) {
	for _, useLuaLock := range []bool{false, true} {
		name := "set"
		if useLuaLock {
			name = "lua"
		}
		t.Run(name, func(t *testing.T) {
			key := strconv.Itoa(rand.Int())
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			var wrapped *cancelAfterAcquireClient
			client, err := NewClient(ClientOption{
				ClientOption: rueidis.ClientOption{InitAddress: addr, PipelineMultiplex: -1, SelectDB: 5},
				ClientTTL:    time.Second,
				UseLuaLock:   useLuaLock,
				ClientBuilder: func(option rueidis.ClientOption) (rueidis.Client, error) {
					client, err := rueidis.NewClient(option)
					if err != nil {
						return nil, err
					}
					wrapped = &cancelAfterAcquireClient{Client: client, key: key, cancel: cancel}
					return wrapped, nil
				},
			})
			if err != nil {
				t.Fatal(err)
			}
			if err = client.Client().Do(context.Background(), client.Client().B().Del().Key(key).Build()).Error(); err != nil {
				t.Fatal(err)
			}
			defer func() {
				client.Client().Do(context.Background(), client.Client().B().Del().Key(key).Build())
				client.Close()
			}()

			loaderCalled := false
			_, err = client.Get(ctx, time.Minute, key, func(context.Context, string) (string, error) {
				loaderCalled = true
				return "value", nil
			})
			if !errors.Is(err, context.Canceled) {
				t.Fatalf("expected context.Canceled, got %v", err)
			}
			if !wrapped.accepted {
				t.Fatal("Redis did not accept the lock command")
			}
			if loaderCalled {
				t.Fatal("loader should not be called when lock acquisition is canceled")
			}

			val, err := client.Client().Do(context.Background(), client.Client().B().Get().Key(key).Build()).ToString()
			if !rueidis.IsRedisNil(err) {
				t.Fatalf("expected canceled lock acquisition to release the placeholder, got value %q and error %v", val, err)
			}
		})
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
	defer client.Close()
	key := strconv.Itoa(rand.Int())
	ch := make(chan string, 2)
	val, err := client.Get(context.Background(), time.Second*5, key, func(ctx context.Context, key string) (val string, err error) {
		id1, err := client.client.Do(context.Background(), client.client.B().Get().Key(key).Build()).ToString()
		if err != nil {
			t.Error(err)
		}
		go func() {
			val, err := client.Get(context.Background(), time.Second*5, key, func(ctx context.Context, key string) (val string, err error) {
				id2, err := client.client.Do(context.Background(), client.client.B().Get().Key(key).Build()).ToString()
				if err != nil {
					t.Error(err)
				}
				ch <- id2
				return "2", nil
			})
			if val != "2" {
				t.Error(err)
			}
		}()
		client.onInvalidation(nil) // simulate disconnection
		id2 := <-ch
		if id1 == id2 {
			t.Error("id not changed")
		}
		ch <- id1
		ch <- id2
		return "1", nil
	})
	if val != "1" {
		t.Fatal(err)
	}
	val, err = client.Get(context.Background(), time.Millisecond*500, key, nil)
	if val != "2" {
		t.Error(err)
	}
	err = client.client.Do(context.Background(), client.client.B().Get().Key(<-ch).Build()).Error() // id1
	if !rueidis.IsRedisNil(err) {
		t.Error(err)
	}
	err = client.client.Do(context.Background(), client.client.B().Get().Key(<-ch).Build()).Error() // id2
	if err != nil {
		t.Error(err)
	}
	time.Sleep(client.ttl) // wait old refresh goroutine exit
}

func TestDisconnectLL(t *testing.T) {
	client := makeClientWithLuaLock(t, addr).(*Client)
	defer client.Close()
	key := strconv.Itoa(rand.Int())
	ch := make(chan string, 2)
	val, err := client.Get(context.Background(), time.Second*5, key, func(ctx context.Context, key string) (val string, err error) {
		id1, err := client.client.Do(context.Background(), client.client.B().Get().Key(key).Build()).ToString()
		if err != nil {
			t.Error(err)
		}
		go func() {
			val, err := client.Get(context.Background(), time.Second*5, key, func(ctx context.Context, key string) (val string, err error) {
				id2, err := client.client.Do(context.Background(), client.client.B().Get().Key(key).Build()).ToString()
				if err != nil {
					t.Error(err)
				}
				ch <- id2
				return "2", nil
			})
			if val != "2" {
				t.Error(err)
			}
		}()
		client.onInvalidation(nil) // simulate disconnection
		id2 := <-ch
		if id1 == id2 {
			t.Error("id not changed")
		}
		ch <- id1
		ch <- id2
		return "1", nil
	})
	if val != "1" {
		t.Fatal(err)
	}
	val, err = client.Get(context.Background(), time.Millisecond*500, key, nil)
	if val != "2" {
		t.Error(err)
	}
	err = client.client.Do(context.Background(), client.client.B().Get().Key(<-ch).Build()).Error() // id1
	if !rueidis.IsRedisNil(err) {
		t.Error(err)
	}
	err = client.client.Do(context.Background(), client.client.B().Get().Key(<-ch).Build()).Error() // id2
	if err != nil {
		t.Error(err)
	}
	time.Sleep(client.ttl) // wait old refresh goroutine exit
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
