# OpenTelemetry Tracing & Connection Metrics

Use `rueidisotel.NewClient` to create a client with OpenTelemetry Tracing and Connection Metrics enabled.
Builtin connection metrics are:
- `rueidis_dial_attempt`: number of dial attempts
- `rueidis_dial_success`: number of successful dials
- `rueidis_dial_conns`: number of connections
- `rueidis_dial_latency`: dial latency in seconds

Client side caching metrics:
- `rueidis_do_cache_miss`: number of cache miss on client side
- `rueidis_do_cache_hits`: number of cache hits on client side

```golang
package main

import (
    "github.com/redis/rueidis"
    "github.com/redis/rueidis/rueidisotel"
)

func main() {
    client, err := rueidisotel.NewClient(rueidis.ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
    if err != nil {
        panic(err)
    }
    defer client.Close()
}
```

See [rueidishook](../rueidishook) if you want more customizations.

Note: `rueidisotel.NewClient` is not supported on go1.18 and go1.19 builds. [Reference](https://github.com/redis/rueidis/issues/442#issuecomment-1886993707)