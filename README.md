# rueidis

[![Go Reference](https://pkg.go.dev/badge/github.com/rueian/rueidis.svg)](https://pkg.go.dev/github.com/rueian/rueidis)
[![circleci](https://circleci.com/gh/rueian/rueidis.svg?style=shield)](https://app.circleci.com/pipelines/github/rueian/rueidis)
[![Go Report Card](https://goreportcard.com/badge/github.com/rueian/rueidis)](https://goreportcard.com/report/github.com/rueian/rueidis)
[![codecov](https://codecov.io/gh/rueian/rueidis/branch/master/graph/badge.svg?token=wGTB8GdY06)](https://codecov.io/gh/rueian/rueidis)
[![Maintainability](https://api.codeclimate.com/v1/badges/0d93d524c2b8497aacbe/maintainability)](https://codeclimate.com/github/rueian/rueidis/maintainability)

A Fast Golang Redis RESP3 client that does auto pipelining and supports client side caching.

## Features

* auto pipeline for non-blocking redis commands
* connection pooling for blocking redis commands
* opt-in client side caching
* redis cluster, sentinel, pub/sub, streams, TLS, RedisJSON, RedisBloom, RediSearch, RedisGraph, RedisTimeseries
* IDE friendly redis command builder
* Hash/RedisJSON Object Mapping with client side caching and optimistic locking
* OpenTelemetry tracing and metrics

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

	// SET key val NX
	c.Do(ctx, c.B().Set().Key("key").Value("val").Nx().Build()).Error()
	// GET key
	c.Do(ctx, c.B().Get().Key("key").Build()).ToString()
}
```

## Auto Pipeline

All non-blocking commands sending to a single redis node are automatically pipelined through one tcp connection,
which reduces the overall round trip costs, and gets higher throughput.

### Benchmark comparison with go-redis v8.11.4

Rueidis has higher throughput than go-redis v8.11.4 across 1, 8, and 64 parallelism settings.

It is even able to achieve ~14x throughput over go-redis in a local benchmark. (see `parallelism(64)-key(16)-value(64)-10`)

#### Single Client
![client_test_set](https://github.com/rueian/rueidis-benchmark/blob/master/client_test_set_6.png)
#### Cluster Client
![cluster_test_set](https://github.com/rueian/rueidis-benchmark/blob/master/cluster_test_set_5.png)

Benchmark source code: https://github.com/rueian/rueidis-benchmark

## Client Side Caching

The Opt-In mode of server-assisted client side caching is always enabled, and can be used by calling `DoCache()` with
an explicit client side TTL.

```golang
c.DoCache(ctx, c.B().Hmget().Key("myhash").Field("1", "2").Cache(), time.Minute).ToArray()
```

An explicit client side TTL is required because redis server may not send invalidation message in time when
a key is expired on the server. Please follow [#6833](https://github.com/redis/redis/issues/6833) and [#6867](https://github.com/redis/redis/issues/6867)

Although an explicit client side TTL is required, the `DoCache()` still sends a `PTTL` command to server and make sure that
the client side TTL is not longer than the TTL on server side.

Users can use `IsCacheHit()` to verify that if the response came from the client side memory.

```golang
c.DoCache(ctx, c.B().Hmget().Key("myhash").Field("1", "2").Cache(), time.Minute).IsCacheHit() == true
```

If the OpenTelemetry is enabled by the `rueidisotel.WithClient(client)`, then there are also two metrics instrumented:
* rueidis_do_cache_miss
* rueidis_do_cache_hits

### Benchmark

![client_test_get](https://github.com/rueian/rueidis-benchmark/blob/master/client_test_get_6.png)

Benchmark source code: https://github.com/rueian/rueidis-benchmark

### Supported Commands by Client Side Caching

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

## Redis Cluster, Single Redis and Sentinel

To connect to a redis cluster, the `NewClient` should be used:

```golang
c, _ := rueidis.NewClient(rueidis.ClientOption{
    InitAddress: []string{"127.0.0.1:7001", "127.0.0.1:7002", "127.0.0.1:7003"},
    ShuffleInit: true,
})
```

To connect to a single redis node, still use the `NewClient` with one InitAddress

To connect to sentinels, specify the required master set name:

```golang
c, _ := rueidis.NewClient(rueidis.ClientOption{
    InitAddress: []string{"127.0.0.1:26379", "127.0.0.1:26380", "127.0.0.1:26381"},
    Sentinel: rueidis.SentinelOption{
        MasterSet: "my_master",
    },
})
```

## Command Builder

Redis commands are very complex and their formats are very different from each other.

This library provides a type safe command builder within `client.B()` that can be used as
an entrypoint to construct a redis command. Once the command is completed, call the `Build()` or `Cache()` to get the actual command.
And then pass it to either `Client.Do()` or `Client.DoCache()`.

```golang
c.Do(ctx, c.B().Set().Key("mykey").Value("myval").Ex(10).Nx().Build())
```

**Once the command is passed to the `Client.Do()`, `Client.DoCache()`, the command will be recycled and should not be reused.**

**The `ClusterClient.B()` also checks if the command contains multiple keys belongs to different slots. If it does, then panic.**

## Object Mapping

The `NewHashRepository` and `NewJSONRepository` creates an OM repository backed by redis hash or RedisJSON.

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
    Key string `json:"key" redis:",key"` // the redis:",key" is required to indicate which field is the ULID key
    Ver int64  `json:"ver" redis:",ver"` // the redis:",ver" is required to do optimistic locking to prevent lost update
    Str string `json:"myStr"`            // both NewHashRepository and NewJSONRepository use json tag as field name
}

func main() {
    ctx := context.Background()
    c, _ := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
    // create the repo with NewHashRepository or NewJSONRepository
    repo := om.NewHashRepository("my_prefix", Example{}, c)

    exp := repo.NewEntity().(*Example)
    exp.Str = "mystr"
    fmt.Println(exp.Key) // output 01FNH4FCXV9JTB9WTVFAAKGSYB
    repo.Save(ctx, exp) // success

    // lookup "my_prefix:01FNH4FCXV9JTB9WTVFAAKGSYB" through client side caching
    cache, _ := repo.FetchCache(ctx, exp.Key, time.Second*5)
    exp2 := cache.(*Example)
    fmt.Println(exp2.Str) // output "mystr", which equals to exp.Str

    exp2.Ver = 0         // if someone changes the version during your GET then SET operation,
    repo.Save(ctx, exp2) // the save will fail with ErrVersionMismatch.
}

```

### Object Mapping + RediSearch

If you have RediSearch, you can create and search the repository against the index.

```golang

if _, ok := repo.(*om.HashRepository); ok {
    repo.CreateIndex(ctx, func(schema om.FtCreateSchema) om.Completed {
        return schema.FieldName("myStr").Text().Build() // Note that the Example.Str field is mapped to myStr on redis by its json tag
    })
}

if _, ok := repo.(*om.JSONRepository); ok {
    repo.CreateIndex(ctx, func(schema om.FtCreateSchema) om.Completed {
        return schema.FieldName("$.myStr").Text().Build() // the field name of json index should be a json path syntax
    })
}

exp := repo.NewEntity().(*Example)
exp.Str = "foo"
repo.Save(ctx, exp)

n, records, _ := repo.Search(ctx, func(search om.FtSearchIndex) om.Completed {
    return search.Query("foo").Build() // you have full query capability by building the command from om.FtSearchIndex
})

fmt.Println("total", n) // n is total number of results matched in redis, which is >= len(records)

for _, v := range records.([]*Example) {
    fmt.Println(v.Str) // print "foo"
}
```

### Object Mapping Limitation

`NewHashRepository` only accepts these field types:
* `string`, `*string`
* `int64`, `*int64`
* `bool`, `*bool`
* `[]byte`

Field projection by RediSearch is not supported.

## OpenTelemetry Tracing

Use `rueidisotel.WithClient` to create a client with OpenTelemetry Tracing enabled.

```golang
package main

import (
    "github.com/rueian/rueidis"
    "github.com/rueian/rueidis/rueidisotel"
)

func main() {
    client, _ := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
    client = rueidisotel.WithClient(client)
    defer client.Close()
}
```

## Command Response Cheatsheet

It is hard to remember what message type is returned from redis and which parsing method should be used with. So, here is some common examples:

```golang
// GET
client.Do(ctx, client.B().Get().Key("k").Build()).ToString()
client.Do(ctx, client.B().Get().Key("k").Build()).AsInt64()
// MGET
client.Do(ctx, client.B().Mget().Key("k1", "k2").Build()).ToArray()
// SET
client.Do(ctx, client.B().Set().Key("k").Value("v").Build()).Error()
// INCR
client.Do(ctx, client.B().Incr().Key("k").Build()).ToInt64()
// HGET
client.Do(ctx, client.B().Hget().Key("k").Field("f").Build()).ToString()
// HMGET
client.Do(ctx, client.B().Hmget().Key("h").Field("a", "b").Build()).ToArray()
// HGETALL
client.Do(ctx, client.B().Hgetall().Key("h").Build()).AsStrMap()
// ZRANGE
client.Do(ctx, client.B().Zrange().Key("k").Min("1").Max("2").Build()).AsStrSlice()
// ZRANK
client.Do(ctx, client.B().Zrank().Key("k").Member("m").Build()).ToInt64()
// ZSCORE
client.Do(ctx, client.B().Zscore().Key("k").Member("m").Build()).ToFloat64()
// SCARD
client.Do(ctx, client.B().Scard().Key("k").Build()).ToInt64()
// SMEMBERS
client.Do(ctx, client.B().Smembers().Key("k").Build()).AsStrSlice()
// LINDEX
client.Do(ctx, client.B().Lindex().Key("k").Index(0).Build()).ToString()
// LPOP
client.Do(ctx, client.B().Lpop().Key("k").Build()).ToString()
client.Do(ctx, client.B().Lpop().Key("k").Count(2).Build()).AsStrSlice()
```

## Not Yet Implement

The following subjects are not yet implemented.
* go-redis like api layer
* RESP2
