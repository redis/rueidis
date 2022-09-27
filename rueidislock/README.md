# rueidislock

A Distributed Lock Pattern with Redis and Client Side Caching.

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
		KeyMajority:  2,
	})
	if err != nil {
		panic(err)
	}
	defer locker.Close()

	ctx, cancel, err := locker.WithContext(context.Background(), "my_lock")
	if err != nil {
		panic(err)
	}
	defer cancel()

	// use the ctx as normal
}
```

## Features backed by the Redis Client Side Caching
* The `ctx` returned will be canceled automatically if the `KeyMajority` is not hold anymore.
* The `Locker.WithContext` waiting for a lock will retry again as soon as possible when the lock is released by someone.

## How it works

When the `locker.WithContext` is invoked, it will:

1. Try acquiring 3 keys (given that `KeyMajority` is 2), which are `rueidislock:0:my_lock`, `rueidislock:1:my_lock` and `rueidislock:2:my_lock`, by sending redis command `SET NX PXAT`.
2. If the `KeyMajority` is satisfied within the `KeyValidity`, the invocation is successful and the `ctx` is returned.
3. If the invocation is not successful, it will watch client-side caching invalidation notification for retrying again.
4. If the invocation is successful, the `Locker` will extend the validity of the `ctx` periodically and also watch client-side caching invalidation notification for canceling the `ctx` if the `KeyMajority` is not hold anymore.