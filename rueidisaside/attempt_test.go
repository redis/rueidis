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
	"github.com/redis/rueidis/internal/cmds"
)

type delayedAcquireClient struct {
	rueidis.Client
	key     string
	once    sync.Once
	blocked chan struct{}
	pending rueidis.Completed
	guard   string
}

func (c *delayedAcquireClient) Do(ctx context.Context, cmd rueidis.Completed) rueidis.RedisResult {
	delayed := false
	if isGuardedScript(cmd.Commands(), c.key, 2) {
		c.once.Do(func() {
			delayed = true
			c.pending = cmd
			c.guard = cmd.Commands()[4]
			close(c.blocked)
		})
	}
	if delayed {
		<-ctx.Done()
		return rueidis.NewErrorResult(ctx.Err())
	}
	return c.Client.Do(ctx, cmd)
}

func (c *delayedAcquireClient) executePending() rueidis.RedisResult {
	return c.Client.Do(context.Background(), c.pending)
}

type blockingAcquireClient struct {
	rueidis.Client
	key     string
	started chan struct{}
	release chan struct{}
	calls   atomic.Int64
}

type blockingCleanupClient struct {
	rueidis.Client
	key            string
	cancel         context.CancelFunc
	cleanupStarted chan struct{}
	cleanupOnce    sync.Once
	calls          atomic.Int64
}

func (c *blockingCleanupClient) Do(ctx context.Context, cmd rueidis.Completed) rueidis.RedisResult {
	if isGuardedScript(cmd.Commands(), c.key, 2) {
		if c.calls.Add(1) == 1 {
			resp := c.Client.Do(ctx, cmd)
			if resp.Error() == nil {
				c.cancel()
				return rueidis.NewErrorResult(ctx.Err())
			}
			return resp
		}
		c.cleanupOnce.Do(func() { close(c.cleanupStarted) })
		<-ctx.Done()
		return rueidis.NewErrorResult(ctx.Err())
	}
	return c.Client.Do(ctx, cmd)
}

func (c *blockingAcquireClient) Do(ctx context.Context, cmd rueidis.Completed) rueidis.RedisResult {
	if isGuardedScript(cmd.Commands(), c.key, 2) {
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

func isGuardedScript(commands []string, key string, argCount int) bool {
	if len(commands) != 5+argCount || (commands[0] != "EVALSHA" && commands[0] != "EVAL") {
		return false
	}
	return commands[2] == "2" && commands[3] == key
}

func preloadAcquire(client rueidis.Client, useLuaLock bool, key string) error {
	script := guardedAcquire
	if useLuaLock {
		script = guardedAcquireLegacy
	}
	guard := attemptGuardKey(key, "preload")
	_, err := script.Exec(
		context.Background(),
		client,
		[]string{key, guard},
		[]string{"preload-id", "preload"},
	).ToArray()
	return err
}

func newDelayedAcquireCache(t *testing.T, key string, useLuaLock bool) (*Client, *delayedAcquireClient) {
	t.Helper()
	var wrapped *delayedAcquireClient
	client, err := NewClient(ClientOption{
		ClientOption: rueidis.ClientOption{InitAddress: addr, PipelineMultiplex: -1, SelectDB: 5},
		ClientTTL:    time.Second,
		UseLuaLock:   useLuaLock,
		ClientBuilder: func(option rueidis.ClientOption) (rueidis.Client, error) {
			client, err := rueidis.NewClient(option)
			if err != nil {
				return nil, err
			}
			if err = preloadAcquire(client, useLuaLock, key); err != nil {
				client.Close()
				return nil, err
			}
			wrapped = &delayedAcquireClient{
				Client:  client,
				key:     key,
				blocked: make(chan struct{}),
			}
			return wrapped, nil
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	return client.(*Client), wrapped
}

func newBlockingAcquireCache(t *testing.T, key string, useLuaLock bool) (*Client, *blockingAcquireClient) {
	t.Helper()
	var wrapped *blockingAcquireClient
	client, err := NewClient(ClientOption{
		ClientOption: rueidis.ClientOption{InitAddress: addr, PipelineMultiplex: -1, SelectDB: 5},
		ClientTTL:    time.Second,
		UseLuaLock:   useLuaLock,
		ClientBuilder: func(option rueidis.ClientOption) (rueidis.Client, error) {
			client, err := rueidis.NewClient(option)
			if err != nil {
				return nil, err
			}
			if err = preloadAcquire(client, useLuaLock, key); err != nil {
				client.Close()
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
	return client.(*Client), wrapped
}

func newBlockingCleanupCache(
	t *testing.T,
	key string,
	useLuaLock bool,
	cancel context.CancelFunc,
) (*Client, *blockingCleanupClient) {
	t.Helper()
	var wrapped *blockingCleanupClient
	client, err := NewClient(ClientOption{
		ClientOption: rueidis.ClientOption{InitAddress: addr, PipelineMultiplex: -1, SelectDB: 5},
		ClientTTL:    100 * time.Millisecond,
		UseLuaLock:   useLuaLock,
		ClientBuilder: func(option rueidis.ClientOption) (rueidis.Client, error) {
			client, err := rueidis.NewClient(option)
			if err != nil {
				return nil, err
			}
			if err = preloadAcquire(client, useLuaLock, key); err != nil {
				client.Close()
				return nil, err
			}
			wrapped = &blockingCleanupClient{
				Client:         client,
				key:            key,
				cancel:         cancel,
				cleanupStarted: make(chan struct{}),
			}
			return wrapped, nil
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	return client.(*Client), wrapped
}

func TestCanceledAcquireCannotArriveLate(t *testing.T) {
	for _, useLuaLock := range []bool{false, true} {
		name := "redis-7"
		if useLuaLock {
			name = "legacy"
		}
		t.Run(name, func(t *testing.T) {
			key := "late-acquire-" + strconv.Itoa(rand.Int())
			client, wrapped := newDelayedAcquireCache(t, key, useLuaLock)
			defer client.Close()
			defer client.client.Do(context.Background(), client.client.B().Del().Key(key).Build())

			ctx, cancel := context.WithCancel(context.Background())
			first := make(chan getResult, 1)
			var firstLoaderCalled atomic.Bool
			go func() {
				val, err := client.Get(ctx, time.Second, key, func(context.Context, string) (string, error) {
					firstLoaderCalled.Store(true)
					return "old", nil
				})
				first <- getResult{val: val, err: err}
			}()

			select {
			case <-wrapped.blocked:
			case <-time.After(time.Second):
				t.Fatal("acquire did not reach the delay point")
			}

			second := make(chan getResult, 1)
			go func() {
				val, err := client.Get(context.Background(), time.Second, key, func(context.Context, string) (string, error) {
					return "new", nil
				})
				second <- getResult{val: val, err: err}
			}()

			cancel()
			select {
			case result := <-first:
				if !errors.Is(result.err, context.Canceled) {
					t.Fatalf("expected context.Canceled, got %v", result.err)
				}
			case <-time.After(time.Second):
				t.Fatal("canceled leader did not return")
			}
			if firstLoaderCalled.Load() {
				t.Fatal("canceled acquire called the loader")
			}

			select {
			case result := <-second:
				if result.err != nil || result.val != "new" {
					t.Fatalf("follower returned %q, %v", result.val, result.err)
				}
			case <-time.After(2 * time.Second):
				t.Fatal("follower did not retry after cleanup")
			}

			resp, err := wrapped.executePending().ToArray()
			if err != nil {
				t.Fatal(err)
			}
			status, err := resp[0].AsInt64()
			if err != nil || status != acquireRevoked {
				t.Fatalf("late acquire returned status %d, %v", status, err)
			}

			val, err := client.client.Do(context.Background(), client.client.B().Get().Key(key).Build()).ToString()
			if err != nil || val != "new" {
				t.Fatalf("late acquire changed the cached value to %q, %v", val, err)
			}
			if err = client.client.Do(context.Background(), client.client.B().Get().Key(wrapped.guard).Build()).Error(); !rueidis.IsRedisNil(err) {
				t.Fatalf("attempt guard was not revoked: %v", err)
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
			client, wrapped := newBlockingAcquireCache(t, key, useLuaLock)
			defer client.Close()
			defer client.client.Do(context.Background(), client.client.B().Del().Key(key).Build())

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
			close(wrapped.release)

			select {
			case result := <-first:
				if result.err != nil || result.val != "value" {
					t.Fatalf("leader returned %q, %v", result.val, result.err)
				}
			case <-time.After(time.Second):
				t.Fatal("leader did not finish")
			}
			if !errors.Is(secondErr, context.DeadlineExceeded) {
				t.Fatalf("expected follower deadline, got %v", secondErr)
			}
			if secondLoaderCalled.Load() {
				t.Fatal("follower called the loader")
			}
			if calls := wrapped.calls.Load(); calls != 1 {
				t.Fatalf("expected one acquire, got %d", calls)
			}
		})
	}
}

func TestCleanupTimeoutRotatesClientGeneration(t *testing.T) {
	for _, useLuaLock := range []bool{false, true} {
		name := "redis-7"
		if useLuaLock {
			name = "legacy"
		}
		t.Run(name, func(t *testing.T) {
			key := "cleanup-timeout-" + strconv.Itoa(rand.Int())
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			client, wrapped := newBlockingCleanupCache(t, key, useLuaLock, cancel)
			defer client.Close()
			defer client.client.Do(context.Background(), client.client.B().Del().Key(key).Build())

			started := time.Now()
			loaderCalled := false
			_, err := client.Get(ctx, time.Second, key, func(context.Context, string) (string, error) {
				loaderCalled = true
				return "value", nil
			})
			if !errors.Is(err, context.Canceled) {
				t.Fatalf("expected context.Canceled, got %v", err)
			}
			if elapsed := time.Since(started); elapsed > time.Second {
				t.Fatalf("cleanup exceeded its bound: %s", elapsed)
			}
			if loaderCalled {
				t.Fatal("loader ran after the canceled acquire")
			}
			select {
			case <-wrapped.cleanupStarted:
			default:
				t.Fatal("cleanup was not attempted")
			}

			client.mu.Lock()
			id := client.id
			client.mu.Unlock()
			if id != "" {
				t.Fatalf("client generation was not rotated: %q", id)
			}
		})
	}
}

type getResult struct {
	val string
	err error
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

func TestDelayedSetCannotOverwriteNewAttempt(t *testing.T) {
	for _, useLuaLock := range []bool{false, true} {
		name := "redis-7"
		if useLuaLock {
			name = "legacy"
		}
		t.Run(name, func(t *testing.T) {
			client := makeClient(t, addr).(*Client)
			client.useLuaLock = useLuaLock
			defer client.Close()

			key := "late-set-" + strconv.Itoa(rand.Int())
			id := PlaceholderPrefix + "late-set-owner"
			first := attempt{token: "first-token"}
			first.guard = attemptGuardKey(key, first.token)
			second := attempt{token: "second-token"}
			second.guard = attemptGuardKey(key, second.token)
			defer client.client.Do(
				context.Background(),
				client.client.B().Del().Key(key, first.guard, second.guard).Build(),
			)

			if err := client.client.Do(context.Background(), client.client.B().Set().Key(first.guard).Value(first.token).Px(time.Second).Build()).Error(); err != nil {
				t.Fatal(err)
			}
			if _, acquired, err := client.acquire(context.Background(), key, id, first); err != nil || !acquired {
				t.Fatalf("first acquire failed: acquired=%v err=%v", acquired, err)
			}
			if err := revokeAttempt.Exec(context.Background(), client.client, []string{key, first.guard}, []string{id, first.token}).Error(); err != nil {
				t.Fatal(err)
			}

			if err := client.client.Do(context.Background(), client.client.B().Set().Key(second.guard).Value(second.token).Px(time.Second).Build()).Error(); err != nil {
				t.Fatal(err)
			}
			if _, acquired, err := client.acquire(context.Background(), key, id, second); err != nil || !acquired {
				t.Fatalf("second acquire failed: acquired=%v err=%v", acquired, err)
			}
			if status, err := guardedSet.Exec(
				context.Background(),
				client.client,
				[]string{key, second.guard},
				[]string{id, second.token, "new", "1000"},
			).AsInt64(); err != nil || status != 1 {
				t.Fatalf("second set failed: status=%d err=%v", status, err)
			}
			if status, err := guardedSet.Exec(
				context.Background(),
				client.client,
				[]string{key, first.guard},
				[]string{id, first.token, "old", "1000"},
			).AsInt64(); err != nil || status != 0 {
				t.Fatalf("delayed set returned status=%d err=%v", status, err)
			}

			val, err := client.client.Do(context.Background(), client.client.B().Get().Key(key).Build()).ToString()
			if err != nil || val != "new" {
				t.Fatalf("delayed set changed the value to %q, %v", val, err)
			}
		})
	}
}

func TestCleanupRemovesAcceptedAcquire(t *testing.T) {
	for _, useLuaLock := range []bool{false, true} {
		name := "redis-7"
		if useLuaLock {
			name = "legacy"
		}
		t.Run(name, func(t *testing.T) {
			client := makeClient(t, addr).(*Client)
			client.useLuaLock = useLuaLock
			defer client.Close()

			key := "accepted-acquire-" + strconv.Itoa(rand.Int())
			id := PlaceholderPrefix + "accepted-owner"
			a := attempt{token: "accepted-token"}
			a.guard = attemptGuardKey(key, a.token)
			defer client.client.Do(context.Background(), client.client.B().Del().Key(key, a.guard).Build())

			if err := client.client.Do(context.Background(), client.client.B().Set().Key(a.guard).Value(a.token).Px(time.Second).Build()).Error(); err != nil {
				t.Fatal(err)
			}
			if _, acquired, err := client.acquire(context.Background(), key, id, a); err != nil || !acquired {
				t.Fatalf("acquire failed: acquired=%v err=%v", acquired, err)
			}
			if err := revokeAttempt.Exec(context.Background(), client.client, []string{key, a.guard}, []string{id, a.token}).Error(); err != nil {
				t.Fatal(err)
			}
			if err := client.client.Do(context.Background(), client.client.B().Get().Key(key).Build()).Error(); !rueidis.IsRedisNil(err) {
				t.Fatalf("placeholder was not removed: %v", err)
			}
			if err := client.client.Do(context.Background(), client.client.B().Get().Key(a.guard).Build()).Error(); !rueidis.IsRedisNil(err) {
				t.Fatalf("guard was not removed: %v", err)
			}
		})
	}
}

func TestAcquireTTLDoesNotOutliveGuard(t *testing.T) {
	for _, useLuaLock := range []bool{false, true} {
		name := "redis-7"
		if useLuaLock {
			name = "legacy"
		}
		t.Run(name, func(t *testing.T) {
			client := makeClient(t, addr).(*Client)
			client.useLuaLock = useLuaLock
			defer client.Close()

			key := "guard-ttl-" + strconv.Itoa(rand.Int())
			id := PlaceholderPrefix + "ttl-owner"
			a := attempt{token: "ttl-token"}
			a.guard = attemptGuardKey(key, a.token)
			defer client.client.Do(context.Background(), client.client.B().Del().Key(key, a.guard).Build())

			if err := client.client.Do(context.Background(), client.client.B().Set().Key(a.guard).Value(a.token).Px(time.Second).Build()).Error(); err != nil {
				t.Fatal(err)
			}
			if _, acquired, err := client.acquire(context.Background(), key, id, a); err != nil || !acquired {
				t.Fatalf("acquire failed: acquired=%v err=%v", acquired, err)
			}
			pttls := client.client.DoMulti(
				context.Background(),
				client.client.B().Pttl().Key(key).Build(),
				client.client.B().Pttl().Key(a.guard).Build(),
			)
			keyTTL, err := pttls[0].AsInt64()
			if err != nil {
				t.Fatal(err)
			}
			guardTTL, err := pttls[1].AsInt64()
			if err != nil {
				t.Fatal(err)
			}
			if keyTTL <= 0 || guardTTL <= 0 || keyTTL > guardTTL+20 {
				t.Fatalf("unexpected TTLs: key=%dms guard=%dms", keyTTL, guardTTL)
			}
		})
	}
}

func TestAttemptGuardKeyUsesSameSlot(t *testing.T) {
	slotTagsOnce.Do(initSlotTags)
	for slot, tag := range slotTags {
		if got := cmds.Slot(string(tag[:])); got != uint16(slot) {
			t.Fatalf("slot tag %q maps to %d, expected %d", tag, got, slot)
		}
	}

	keys := []string{
		"",
		"plain",
		"key{tag}suffix",
		"{}",
		"unclosed{tag",
		"empty{}tag",
		"\x00binary\xff",
		"nested{{tag}}",
	}
	for _, key := range keys {
		guard := attemptGuardKey(key, "token")
		if got, want := cmds.Slot(guard), cmds.Slot(key); got != want {
			t.Fatalf("guard for %q maps to slot %d, expected %d", key, got, want)
		}
	}
}
