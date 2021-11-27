# rueidis
A Fast Golang Redis RESP3 client that does auto pipelining and supports client side caching.

## Features

* auto pipeline for non-blocking redis commands
* connection pooling for blocking redis commands
* opt-in client side caching
* redis cluster, pub/sub, streams, TLS
* IDE friendly redis command builder

## Requirement

* Currently, only supports redis >= 6.x

## Getting Started

```golang
package main

import "github.com/rueian/rueidis"

func main() {
    c, _ := rueidis.NewClusterClient(rueidis.ClusterClientOption{
        InitAddress: []string{"127.0.0.1:6379"},
    })
    defer c.Close()

    _ := c.Do(c.Cmd.Set().Key("my_redis_data:1").Value("my_value").Nx().Build()).Error()
    val, _ := c.Do(c.Cmd.Get().Key("my_redis_data:1").Build()).ToString()
    // val == "my_value"
}
```

## Auto Pipeline

All non-blocking commands sending to a single redis instance are automatically pipelined through one tcp connection,
which reduces the overall round trip costs, and gets higher throughput.

### Benchmark comparison with go-redis v8.11.4

```shell
▶ # run redis-server 6.2.5 at 127.0.0.1:6379
▶ ./redis-server
▶ go test -bench=. -benchmem ./cmd/bench3/...
goos: darwin
goarch: amd64
pkg: github.com/rueian/rueidis/cmd/bench3
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkRedisClient/RueidisParallel100Get-12    1559809     744.2 ns/op    104 B/op    2 allocs/op
BenchmarkRedisClient/GoRedisParallel100Get-12     148611      7915 ns/op    208 B/op    6 allocs/op
PASS
ok  	github.com/rueian/rueidis/cmd/bench3	3.589s

```
Benchmark source code:
```golang
func BenchmarkRedisClient(b *testing.B) {
    b.Run("RueidisParallel100Get", func(b *testing.B) {
        c, _ := rueidis.NewSingleClient(rueidis.SingleClient{Address: "127.0.0.1:6379"})
        b.SetParallelism(100)
        b.ResetTimer()
        b.RunParallel(func(pb *testing.PB) {
            for pb.Next() {
                c.Do(c.Cmd.Get().Key("a").Build())
            }
        })
        c.Close()
    })
    b.Run("GoRedisParallel100Get", func(b *testing.B) {
        rdb := redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379", PoolSize: 100})
        ctx := context.Background()
        b.SetParallelism(100)
        b.ResetTimer()
        b.RunParallel(func(pb *testing.PB) {
            for pb.Next() {
                rdb.Get(ctx, "a")
            }
        })
        rdb.Close()
    })
}
```

## Client Side Caching

The Opt-In mode of server-assisted client side caching is always enabled, and can be used by calling `DoCache()` with
an explicit client side TTL.

An explicit client side TTL is required because redis server may not send invalidation message in time when
a key is expired on the server. Please follow [#6833](https://github.com/redis/redis/issues/6833) and [#6867](https://github.com/redis/redis/issues/6867)

Although an explicit client side TTL is required, the `DoCache()` still sends a `PTTL` command to server and make sure that
the client side TTL is not longer than the TTL on server side.

### Benchmark [(source)](./pkg/conn/conn_test.go)

```shell
▶ ./redis-server
▶ go test -bench=BenchmarkClientSideCaching -benchmem ./pkg/conn
goos: darwin
goarch: amd64
pkg: github.com/rueian/rueidis/pkg/conn
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkClientSideCaching/Do-12         1378052    866.8 ns/op    0 B/op    0 allocs/op
BenchmarkClientSideCaching/DoCache-12    3485281    351.1 ns/op    0 B/op    0 allocs/op
PASS
ok  	github.com/rueian/rueidis/pkg/conn	3.057s
```

### Supported Commands for Client Side Caching

* bitcount
* bitfieldro
* bitpos
* geodist
* geohash
* geopos
* geosearch
* get
* getbit
* getrange
* hexists
* hget
* hgetall
* hkeys
* hlen
* hmget
* hstrlen
* hvals
* lindex
* llen
* lpos
* lrange
* pttl
* scard
* sismember
* smembers
* smismember
* strlen
* substr
* ttl
* type
* zcard
* zcount
* zlexcount
* zmscore
* zrange
* zrangebylex
* zrangebyscore
* zrank
* zrevrange
* zrevrangebylex
* zrevrangebyscore
* zrevrank
* zscore

## Blocking Commands

The following blocking commands use another connection pool and will not share the same connection
with non-blocking commands and thus will not cause the pipeline to be blocked:

* xread with block
* xreadgroup with block
* blpop
* brpop
* brpoplpush
* blmove
* blmpop
* bzpopmin
* bzpopmax
* client pause
* migrate
* wait

## Pub/Sub

To receive messages from channels, the message handler should be registered when creating the redis connection:

```golang
c, _ := rueidis.NewSingleClient(rueidis.SingleClient{
    Address: "127.0.0.1:6379",
    ConnOption: conn.Option{
        PubSubHandlers: conn.PubSubHandlers{
            OnMessage: func(channel, message string) {
                // handle the message
            },
        },
    },
})
c.Do(c.Cmd.Subscribe().Channel("my_channel").Build())
```

## CAS Pattern

To do a CAS operation (WATCH + MULTI + EXEC), a dedicated connection should be used, because there should be no
unintentional write commands between WATCH and EXEC. Otherwise, the EXEC may not fail as expected.

The dedicated connection shares the same connection pool with blocking commands.

```golang
c.DedicatedWire(func(client client.DedicatedSingleClient) error {
    // watch keys first
    client.Do(c.Cmd.Watch().Key("k1", "k2").Build())
    // perform read here
    client.Do(c.Cmd.Mget().Key("k1", "k2").Build())
    // perform write with MULTI EXEC
    client.DoMulti(
        c.Cmd.Multi().Build(),
        c.Cmd.Set().Key("k1").Value("1").Build(),
        c.Cmd.Set().Key("k2").Value("2").Build(),
        c.Cmd.Exec().Build(),
    )
    return nil
})
```

However, occupying a connection is not good in terms of throughput. It is better to use Lua script to perform
optimistic locking instead.

## Lua Script

The `NewLuaScript` or `NewLuaScriptReadOnly` will create a script which is safe for concurrent usage.

When calling the `script.Exec`, it will try sending EVALSHA to the client and if the server returns NOSCRIPT,
it will send EVAL to try again.

```golang
script := c.NewLuaScript("return {KEYS[1],KEYS[2],ARGV[1],ARGV[2]}")
// the script.Exec is safe for concurrent call
list, err := script.Exec([]string{"k1", "k2"}, []string{"a1", "a2"}).ToArray()
```

## Redis Cluster

To connect to a redis cluster, the `NewClusterClient` should be used:

```golang
c, _ := rueidis.NewClusterClient(rueidis.ClusterClientOption{
    InitAddress: []string{"127.0.0.1:7001", "127.0.0.1:7002", "127.0.0.1:7003"},
    ShuffleInit: false,
})
```

## Command Builder

Redis commands are very complex and their formats are very different from each other.

This library provides a type safe command builder within `SingleClient.Cmd` and `ClusterClient.Cmd` that can be used as
an entrypoint to construct a redis command. Once the command is completed, call the `Build()` or `Cache()` to get the actual command.
And then pass it to either `Client.Do()` or `Client.DoCache()`.

```golang
c.Do(c.Cmd.Set().Key("mykey").Value("myval").Ex(10).Nx().Build())
c.DoCache(c.Cmd.Hmget().Key("myhash").Field("1", "2").Cache(), time.Second*30)
```

**Once the command is passed to the one of above `Client.DoXXX()`, the command will be recycled and should not be reused.**

**The `ClusterClient.Cmd` also checks if the command contains multiple keys belongs to different slots. If it does, then panic.**

## Object Mapping

The `NewHashRepository` creates an OM repository backed by redis hash.

```golang
package main

import (
    "context"
    "fmt"
    "time"

    "github.com/rueian/rueidis"
)

type Example struct {
    ID       string   `redis:"-,pk"`     // the pk option indicate that this field is the ULID key
    Ver      int64    `redis:"_v"`       // the _v field is required for optimistic locking to prevent the lost update
    MyString string   `redis:"f1"`
    MyBool   bool     `redis:"f3"`
    MyArray  []string `redis:"f4,sep=|"` // the sep=<ooo> option is required for converting the slice to/from a string
}

func main() {
    ctx := context.Background()
    c, _ := rueidis.NewSingleClient(rueidis.SingleClientOption{Address: "127.0.0.1:6379"})
    // create the hash repo.
    repo := c.NewHashRepository("my_prefix", Example{})

    exp := repo.Make().(*Example)
    exp.MyArray = []string{"1", "2"}
    fmt.Println(exp.ID) // output 01FNH4FCXV9JTB9WTVFAAKGSYB
    repo.Save(ctx, exp) // success

    // lookup "my_prefix:01FNH4FCXV9JTB9WTVFAAKGSYB" through client side caching
    cache, _ := repo.FetchCache(ctx, exp.ID, time.Second*5)
    exp2 := cache.(*Example)
    fmt.Println(exp2.MyArray) // output [1 2], which equals to exp.MyArray

    exp2.Ver = 0         // if someone changes the version during your GET then SET operation,
    repo.Save(ctx, exp2) // the save will fail with ErrVersionMismatch.
}

```

## Not Yet Implement

The following subjects are not yet implemented.

* RESP2
