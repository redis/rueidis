# rueidislock

A [Redis Distributed Lock Pattern](https://redis.io/docs/latest/develop/use/patterns/distributed-locks/) enhanced by [Client Side Caching](https://redis.io/docs/manual/client-side-caching/).

```go
package main

import (
	"context"
	"github.com/redis/rueidis"
	"github.com/redis/rueidis/rueidislock"
)

func main() {
	locker, err := rueidislock.NewLocker(rueidislock.LockerOption{
		ClientOption:   rueidis.ClientOption{InitAddress: []string{"localhost:6379"}},
		KeyMajority:    1,    // Use KeyMajority=1 if you have only one Redis instance. Also make sure that all your `Locker`s share the same KeyMajority.
		NoLoopTracking: true, // Enable this to have better performance if all your Redis are >= 7.0.5.
	})
	if err != nil {
		panic(err)
	}
	defer locker.Close()

	// acquire the lock "my_lock"
	ctx, cancel, err := locker.WithContext(context.Background(), "my_lock")
	if err != nil {
		panic(err)
	}

	// "my_lock" is acquired. use the ctx as normal.
	doSomething(ctx)

	// invoke cancel() to release the lock.
	cancel()
}
```

## Default Configuration

When creating a new `Locker` without explicitly setting options, the following defaults are applied:

- **KeyPrefix**: `"rueidislock"`  
- **KeyValidity**: `5s` — lock validity period before it expires.  
- **ExtendInterval**: `KeyValidity / 2` (default: 2.5s) — how often the lock is automatically extended.  
- **TryNextAfter**: `20ms` — wait time before trying the next Redis key when acquiring locks.  
- **KeyMajority**: `2` — number of keys required out of `KeyMajority*2-1` total keys for a valid lock.  
- **NoLoopTracking**: `false` — disables NOLOOP in client tracking unless explicitly set.  
- **FallbackSETPX**: `false` — uses `SET PXAT` by default; set to `true` for Redis versions < 6.2.

These values can be overridden by providing a `LockerOption` when calling `NewLocker`.

## Features backed by the Redis Client Side Caching
* The returned `ctx` will be canceled automatically and immediately once the `KeyMajority` is not held anymore, for example:
  * Redis are down.
  * Acquired keys have been deleted by other programs or administrators.
* The waiting `Locker.WithContext` will try acquiring the lock again automatically and immediately once it has been released by someone or by another program.

## How it works

When the `locker.WithContext` is invoked, it will:

1. Try acquiring 3 keys (given that the default `KeyMajority` is 2), which are `rueidislock:0:my_lock`, `rueidislock:1:my_lock` and `rueidislock:2:my_lock`, by sending redis command `SET NX PXAT` or `SET NX PX` if `FallbackSETPX` is set.
2. If the `KeyMajority` is satisfied within the `KeyValidity` duration, the invocation is successful and a `ctx` is returned as the lock.
3. If the invocation is not successful, it will wait for client-side caching notifications to retry again.
4. If the invocation is successful, the `Locker` will extend the `ctx` validity periodically and also watch client-side caching notifications for canceling the `ctx` if the `KeyMajority` is not held anymore.

### Disable Client Side Caching

Some Redis providers don't support client-side caching, ex. Google Cloud Memorystore.
You can disable client-side caching by setting `ClientOption.DisableCache` to `true`.
Please note that when the client-side caching is disabled, rueidislock will only try to re-acquire locks for every ExtendInterval.

## Benchmark

```bash
▶ go test -bench=. -benchmem -run=.
goos: darwin
goarch: arm64
pkg: rueidis-benchmark/locker
Benchmark/rueidislock-10         	   20103	     57842 ns/op	    1849 B/op	      29 allocs/op
Benchmark/redislock-10           	   13209	     86285 ns/op	    8083 B/op	     225 allocs/op
PASS
ok  	rueidis-benchmark/locker	3.782s
```

```go
package locker

import (
	"context"
	"testing"
	"time"

	"github.com/bsm/redislock"
	"github.com/redis/go-redis/v9"
	"github.com/redis/rueidis"
	"github.com/redis/rueidis/rueidislock"
)

func Benchmark(b *testing.B) {
	b.Run("rueidislock", func(b *testing.B) {
		l, _ := rueidislock.NewLocker(rueidislock.LockerOption{
			ClientOption:   rueidis.ClientOption{InitAddress: []string{"127.0.0.1:6379"}},
			KeyMajority:    1,
			NoLoopTracking: true,
		})
		defer l.Close()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_, cancel, err := l.WithContext(context.Background(), "mylock")
				if err != nil {
					panic(err)
				}
				cancel()
			}
		})
		b.StopTimer()
	})
	b.Run("redislock", func(b *testing.B) {
		client := redis.NewUniversalClient(&redis.UniversalOptions{Addrs: []string{"127.0.0.1:6379"}})
		locker := redislock.New(client)
		defer client.Close()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
			retry:
				lock, err := locker.Obtain(context.Background(), "mylock", time.Minute, nil)
				if err == redislock.ErrNotObtained {
					goto retry
				} else if err != nil {
					panic(err)
				}
				lock.Release(context.Background())
			}
		})
		b.StopTimer()
	})
}
```
