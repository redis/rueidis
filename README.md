# rueidis

[![Go Reference](https://pkg.go.dev/badge/github.com/rueian/rueidis.svg)](https://pkg.go.dev/github.com/rueian/rueidis)
[![circleci](https://circleci.com/gh/rueian/rueidis.svg?style=shield)](https://app.circleci.com/pipelines/github/rueian/rueidis)
[![Maintainability](https://api.codeclimate.com/v1/badges/0d93d524c2b8497aacbe/maintainability)](https://codeclimate.com/github/rueian/rueidis/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/0d93d524c2b8497aacbe/test_coverage)](https://codeclimate.com/github/rueian/rueidis/test_coverage)

A Fast Golang Redis RESP3 client that does auto pipelining and supports client side caching.

## Features

* auto pipeline for non-blocking redis commands
* connection pooling for blocking redis commands
* opt-in client side caching
* redis cluster, pub/sub, streams, TLS, RedisJSON, RedisBloom, RediSearch, RedisGraph, RedisTimeseries
* IDE friendly redis command builder
* Hash Object Mapping with client side caching and optimistic locking

## Requirement

* Currently, only supports redis >= 6.x

## Getting Started

```golang
package main

import (
	"context"
	"github.com/rueian/rueidis"
)

func main() {
	c, _ := rueidis.NewClient(rueidis.ClientOption{
		InitAddress: []string{"127.0.0.1:6379"},
	})
	defer c.Close()

	ctx := context.Background()

	_ := c.Do(ctx, c.B().Set().Key("my_data").Value("my_value").Nx().Build()).Error()
	val, _ := c.Do(ctx, c.B().Get().Key("my_data").Build()).ToString()
	// val == "my_value"
}
```

## Auto Pipeline

All non-blocking commands sending to a single redis instance are automatically pipelined through one tcp connection,
which reduces the overall round trip costs, and gets higher throughput.

### Benchmark comparison with go-redis v8.11.4

Rueidis has higher throughput than go-redis v8.11.4 across 1, 8, and 64 parallelism settings.

It is even able to achieve ~14x throughput over go-redis in a local benchmark. (see `parallelism(64)-key(16)-value(64)-10`)

#### Single Client
![client_test_set](https://github.com/rueian/rueidis-benchmark/blob/master/client_test_set_4.png)
#### Cluster Client
![cluster_test_set](https://github.com/rueian/rueidis-benchmark/blob/master/cluster_test_set_4.png)

Benchmark source code: https://github.com/rueian/rueidis-benchmark

## Client Side Caching

The Opt-In mode of server-assisted client side caching is always enabled, and can be used by calling `DoCache()` with
an explicit client side TTL.

An explicit client side TTL is required because redis server may not send invalidation message in time when
a key is expired on the server. Please follow [#6833](https://github.com/redis/redis/issues/6833) and [#6867](https://github.com/redis/redis/issues/6867)

Although an explicit client side TTL is required, the `DoCache()` still sends a `PTTL` command to server and make sure that
the client side TTL is not longer than the TTL on server side.

### Benchmark

![client_test_get](https://github.com/rueian/rueidis-benchmark/blob/master/client_test_get_4.png)

Benchmark source code: https://github.com/rueian/rueidis-benchmark

### Supported Commands for Client Side Caching

* bitcount
* bitfieldro
* bitpos
* expiretime
* geodist
* geohash
* geopos
* georadiusro
* georadiusbymemberro
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
* pexpiretime
* pttl
* scard
* sismember
* smembers
* smismember
* sortro
* strlen
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
* jsonget
* jsonstrlen
* jsonarrindex
* jsonarrlen
* jsonobjkeys
* jsonobjlen
* jsontype
* jsonresp
* bfexists
* bfinfo
* cfexists
* cfcount
* cfinfo
* cmsquery
* cmsinfo
* topkquery
* topklist
* topkinfo

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
* bzmpop
* clientpause
* migrate
* wait

## Pub/Sub

To receive messages from channels, the message handler should be registered when creating the redis connection:

```golang
c, _ := rueidis.NewClient(rueidis.ClientOption{
    InitAddress: []string{"127.0.0.1:6379"},
    PubSubOption: rueidis.NewPubSubOption(func(prev error, client rueidis.DedicatedClient) {
        // Subscribe channels in this PubSubSetup hook for auto reconnecting after disconnected.
        // The "prev" err is previous disconnect error.
        err := client.Do(ctx, client.B().Subscribe().Channel("my_channel").Build()).Error()
    }, rueidis.PubSubHandler{
        OnMessage: func(channel, message string) {
            // handle the message
        },
    },
})
```

## CAS Pattern

To do a CAS operation (WATCH + MULTI + EXEC), a dedicated connection should be used, because there should be no
unintentional write commands between WATCH and EXEC. Otherwise, the EXEC may not fail as expected.

The dedicated connection shares the same connection pool with blocking commands.

```golang
c.Dedicated(func(client client.DedicatedClient) error {
    // watch keys first
    client.Do(ctx, client.B().Watch().Key("k1", "k2").Build())
    // perform read here
    client.Do(ctx, client.B().Mget().Key("k1", "k2").Build())
    // perform write with MULTI EXEC
    client.DoMulti(
        ctx,
        client.B().Multi().Build(),
        client.B().Set().Key("k1").Value("1").Build(),
        client.B().Set().Key("k2").Value("2").Build(),
        client.B().Exec().Build(),
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
script := rueidis.NewLuaScript("return {KEYS[1],KEYS[2],ARGV[1],ARGV[2]}")
// the script.Exec is safe for concurrent call
list, err := script.Exec(ctx, client, []string{"k1", "k2"}, []string{"a1", "a2"}).ToArray()
```

## Redis Cluster and Single Redis

To connect to a redis cluster, the `NewClient` should be used:

```golang
c, _ := rueidis.NewClient(rueidis.ClientOption{
    InitAddress: []string{"127.0.0.1:7001", "127.0.0.1:7002", "127.0.0.1:7003"},
    ShuffleInit: true,
})
```

To connect to a single redis node, still use the `NewClient` with one InitAddress

## Command Builder

Redis commands are very complex and their formats are very different from each other.

This library provides a type safe command builder within `client.B()` that can be used as
an entrypoint to construct a redis command. Once the command is completed, call the `Build()` or `Cache()` to get the actual command.
And then pass it to either `Client.Do()` or `Client.DoCache()`.

```golang
c.Do(ctx, c.B().Set().Key("mykey").Value("myval").Ex(10).Nx().Build())
c.DoCache(ctx, c.B().Hmget().Key("myhash").Field("1", "2").Cache(), time.Second*30)
```

**Once the command is passed to the one of above `Client.DoXXX()`, the command will be recycled and should not be reused.**

**The `ClusterClient.B()` also checks if the command contains multiple keys belongs to different slots. If it does, then panic.**

## Object Mapping

The `NewHashRepository` creates an OM repository backed by redis hash.

```golang
package main

import (
    "context"
    "fmt"
    "time"

    "github.com/rueian/rueidis"
    "github.com/rueian/rueidis/om"
)

type Example struct {
    ID     string   `redis:"-,pk"`     // the pk option indicate that this field is the ULID key
    Ver    int64    `redis:"_v"`       // the _v field is required for optimistic locking to prevent the lost update
    MyStr  string   `redis:"f1"`
    MyArr  []string `redis:"f2,sep=|"` // the sep=<ooo> option is required for converting the slice to/from a string
}

func main() {
    ctx := context.Background()
    c, _ := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
    // create the hash repo.
    repo := om.NewHashRepository("my_prefix", Example{}, c)

    exp := repo.NewEntity().(*Example)
    exp.MyArr = []string{"1", "2"}
    fmt.Println(exp.ID) // output 01FNH4FCXV9JTB9WTVFAAKGSYB
    repo.Save(ctx, exp) // success

    // lookup "my_prefix:01FNH4FCXV9JTB9WTVFAAKGSYB" through client side caching
    cache, _ := repo.FetchCache(ctx, exp.ID, time.Second*5)
    exp2 := cache.(*Example)
    fmt.Println(exp2.MyArr) // output [1 2], which equals to exp.MyArray

    exp2.Ver = 0         // if someone changes the version during your GET then SET operation,
    repo.Save(ctx, exp2) // the save will fail with ErrVersionMismatch.
}

```

### Object Mapping + RediSearch

If you have RediSearch, you can create and search the repository against the index.

```golang

repo.CreateIndex(ctx, func(schema om.FtCreateSchema) om.Completed {
    return schema.FieldName("f1").Text().Build() // you have full index capability by building the command from om.FtCreateSchema
})

exp := repo.NewEntity().(*Example)
exp.MyStr = "foo" // Note that MyStr is mapped to "f1" in redis by the `redis:"f1"` tag
repo.Save(ctx, exp)

n, records, _ := repo.Search(ctx, func(search om.FtSearchIndex) om.Completed {
    return search.Query("foo").Build() // you have full query capability by building the command from om.FtSearchIndex
})

fmt.Println("total", n) // n is total number of results matched in redis, which is >= len(records)

for _, v := range records.([]*Example) {
    fmt.Println(v.MyStr) // print "foo"
}
```

## Not Yet Implement

The following subjects are not yet implemented.

* RESP2
