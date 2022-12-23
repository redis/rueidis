# rueidislock

A [Redis Distributed Lock Pattern](https://redis.io/docs/reference/patterns/distributed-locks/) enhanced by [Client Side Caching](https://redis.io/docs/manual/client-side-caching/).

```go
package main

import (
	"context"
	"github.com/rueian/rueidis"
	"github.com/rueian/rueidis/rueidislock"
)

func main() {
	locker, err := rueidislock.NewLocker(rueidislock.LockerOption{
		ClientOption: rueidis.ClientOption{InitAddress: []string{"node1:6379", "node2:6380", "node3:6379"}},
		KeyMajority:  2, // please make sure that all your locker share the same KeyMajority
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

## Features backed by the Redis Client Side Caching
* The returned `ctx` will be canceled automatically if the `KeyMajority` is not held anymore.
* The waiting `Locker.WithContext` will retry again if the lock is released by someone.

## How it works

When the `locker.WithContext` is invoked, it will:

1. Try acquiring 3 keys (given that `KeyMajority` is 2), which are `rueidislock:0:my_lock`, `rueidislock:1:my_lock` and `rueidislock:2:my_lock`, by sending redis command `SET NX PXAT`.
2. If the `KeyMajority` is satisfied within the `KeyValidity` duration, the invocation is successful and a `ctx` is returned.
3. If the invocation is not successful, it will wait for client-side caching notification to retry again.
4. If the invocation is successful, the `Locker` will extend the `ctx` validity periodically and also watch client-side caching notification for canceling the `ctx` if the `KeyMajority` is not held anymore.