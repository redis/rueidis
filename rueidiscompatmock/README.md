## Go-redis like API Mock Adapter

`rueidiscompat` exists as a very close interface to go-redis's `Cmdable` interface,
but there is not a test helper that mirrors go-redis's `ClientMock`  
interface. This package aims to do that.

### Usage example

```golang
package main

import (
	"context"
	"testing"

	"github.com/redis/rueidis/mock"
	"github.com/redis/rueidis/rueidiscompat"
	"github.com/redis/rueidis/rueidiscompatmock"
	"go.uber.org/mock/gomock"
)

func TestExample(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := mock.NewClient(ctrl)
	compatmock := rueidiscompatmock.NewAdapter(m)

	compatmock.ExpectSet("key", "val", 0).SetVal("OK")
	compatmock.ExpectGet("key").SetVal("val")

	rdb := rueidiscompat.NewAdapter(m)
	rdb.Set(context.Background(), "key", "val", 0)
	rdb.Get(context.Background(), "key")
}
```

### Pipeline example

Pipelined commands use the same `Expect*` calls as non-pipelined commands.
Expectations are matched in the order they are queued, mirroring `go-redis/redismock`.

```golang
compatmock.ExpectGet("k1").SetVal("v1")
compatmock.ExpectSet("k2", "v2", 0).SetVal("OK")

p := rdb.Pipeline()
p.Get(ctx, "k1")
p.Set(ctx, "k2", "v2", 0)
p.Exec(ctx)
```

### Client-side caching example

Cached commands (`Cache(ttl)`) use the same `Expect*` calls as non-cached commands.
The TTL argument is not matched against expectations.

```golang
compatmock.ExpectGet("key").SetVal("val")

cached := rdb.Cache(100 * time.Millisecond)
cached.Get(context.Background(), "key")
```

If a test needs to assert that a call went through the cache path specifically
(as opposed to a plain, non-cached call), chain `ViaCache()` onto the
expectation. A `ViaCache()` expectation is only satisfied by a call made
through `Cache(ttl)`; a plain call for the same command will not match it.
Expectations that don't call `ViaCache()` are unaffected and continue to match
calls from either origin.

```golang
compatmock.ExpectGet("key").ViaCache().SetVal("val")

cached := rdb.Cache(100 * time.Millisecond)
cached.Get(context.Background(), "key") // matches

rdb.Get(context.Background(), "key") // would NOT match the ViaCache() expectation
```