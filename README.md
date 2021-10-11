# rueidis

A Golang Redis RESP3 client that does auto pipelining and supports client side caching.

## Auto Pipeline

All commands to a single redis are pipelined through one tcp connection, which reduces
the overall round trip costs, and gets higher throughput.

### Benchmark comparison with go-redis v8.11.4

```shell
▶ # run redis-server at 127.0.0.1:6379
▶ ./redis-6.2.5/src/redis-server
▶ go test -bench=. -benchmem
goos: darwin
goarch: amd64
pkg: github.com/rueian/rueidis/cmd/bench3
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkRedisClient/GoRedisParallel100Set1KB-12                   71458             21280 ns/op             855 B/op          9 allocs/op
BenchmarkRedisClient/RueidisParallel100Set1KB-12                  379381              2809 ns/op              34 B/op          3 allocs/op
PASS
ok      github.com/rueian/rueidis/cmd/bench3    3.973s

```
Benchmark source code:
```golang
func BenchmarkRedisClient(b *testing.B) {
	sb := strings.Builder{}
	sb.Write(make([]byte, 1024))

	b.Run("GoRedisParallel100Set1KB", func(b *testing.B) {
		rdb := redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379", PoolSize: 1000})
		ctx := context.Background()
		b.SetParallelism(100)
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				rdb.Set(ctx, "a", sb.String(), 0)
			}
		})
		rdb.Close()
	})
	b.Run("RueidisParallel100Set1KB", func(b *testing.B) {
		c, _ := conn.NewConn("127.0.0.1:6379", conn.Option{})
		b.SetParallelism(100)
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.Do(c.Cmd.Set().Key("a").Value(sb.String()).Build())
			}
		})
		c.Close()
	})
}
```

## Client Side Caching

The Opt-In mode of server-assisted client side caching are always enabled, and can be used by calling `DoCache()` with
a separated client side TTL.

A separated client side TTL is required because the current spec (redis 6.2) of Client Side Caching doesn't include notification of
key expiration on server in time.

### Benchmark

```shell
▶ go test -bench=. -benchmem
goos: darwin
goarch: amd64
pkg: github.com/rueian/rueidis/cmd/bench4
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkClientSideCache/Do-12                    594303            1920 ns/op      1048 B/op          2 allocs/op
BenchmarkClientSideCache/DoCache-12              3448129           347.1 ns/op        24 B/op          1 allocs/op
PASS
ok  	github.com/rueian/rueidis/cmd/bench4	3.801s
```
Benchmark source code:
```golang
func BenchmarkClientSideCache(b *testing.B) {
	b.Run("Do", func(b *testing.B) {
		c, _ := conn.NewConn("127.0.0.1:6379", conn.Option{CacheSize: conn.DefaultCacheBytes})
		b.SetParallelism(100)
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.Do(c.Cmd.Get().Key("a").Build())
			}
		})
	})
	b.Run("DoCache", func(b *testing.B) {
		c, _ := conn.NewConn("127.0.0.1:6379", conn.Option{CacheSize: conn.DefaultCacheBytes})
		b.SetParallelism(100)
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.DoCache(c.Cmd.Get().Key("a").Build(), time.Second*5)
			}
		})
	})
}
```

## Command Builder

Redis commands are very complex and their formats are very different from each other.

This library provides a type safe command builder with in `Conn.Cmd` that can be used as
an entrypoint to construct a redis command. Once the command is completed, call the `Build()` to get the actual command.
And then pass it to either `Conn.Do()` or `Conn.DoMulti()` or `Conn.DoCache()`.

```golang
c.Do(c.Cmd.Set().Key("mykey").Value("myval").Ex(10).Nx().Build())
```

Once the command is passed to the one of above `Conn.DoXXX()`, the command will be recycled and should not be reused.

## Not Yet Implement

The following subjects are not yet implemented.

* Better blocking commands supporting (ex: BLPOP) 
* PubSub commands
* Redis Cluster client
* ~~Auto Reconnect~~
* RESP2
