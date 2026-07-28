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

func TestBeginFlightWaitsAcrossGenerations(t *testing.T) {
	c := &Client{
		id:      "new-generation",
		flights: make(map[string]*flight),
	}
	if f, leader := c.beginFlight("key", "old-generation"); f != nil || leader {
		t.Fatal("stale generation created a flight")
	}
	f, leader := c.beginFlight("key", "new-generation")
	if f == nil || !leader {
		t.Fatal("current generation did not create a flight")
	}
	staleFollower, leader := c.beginFlight("key", "old-generation")
	if staleFollower != f || leader {
		t.Fatal("stale generation did not join the current flight")
	}
	follower, leader := c.beginFlight("key", "new-generation")
	if follower != f || leader {
		t.Fatal("same key did not join the active flight")
	}
	if staleOther, leader := c.beginFlight("other-key", "old-generation"); staleOther != nil || leader {
		t.Fatal("stale generation joined a flight for another key")
	}
	if len(c.flights) != 1 {
		t.Fatalf("stale generation changed the flights map: %d entries", len(c.flights))
	}
	if other, leader := c.beginFlight("other-key", "new-generation"); other == nil || !leader {
		t.Fatal("different key did not create an independent flight")
	}
	select {
	case <-f.done:
		t.Fatal("flight completed before the leader finished")
	default:
	}
	c.finishFlight("key", f)
	select {
	case <-f.done:
	default:
		t.Fatal("flight did not wake its follower")
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
				ClientOption: rueidis.ClientOption{InitAddress: addr, SelectDB: 5},
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
