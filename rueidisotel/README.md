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

These metrics can include additional labels using `Labeler` (see below).

Client side command metrics:
- `rueidis_command_duration_seconds`: histogram of command duration
- `rueidis_command_errors`: number of command errors

```golang
package main

import (
    "context"
    "time"

    "github.com/redis/rueidis"
    "github.com/redis/rueidis/rueidisotel"
    "go.opentelemetry.io/otel/attribute"
)

func main() {
    client, err := rueidisotel.NewClient(rueidis.ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
    if err != nil {
        panic(err)
    }
    defer client.Close()

    // Basic usage
    ctx := context.Background()
    client.DoCache(ctx, client.B().Get().Key("mykey").Cache(), time.Minute)

    // Add custom labels to cache metrics using Labeler
    // Check if labeler exists in context, create new context only if needed
    bookLabeler, ok := valkeyotel.LabelerFromContext(ctx)
    if !ok {
        ctx = valkeyotel.ContextWithLabeler(ctx, bookLabeler)
    }
    bookLabeler.Add(attribute.String("key_pattern", "book"))
    client.DoCache(ctx, client.B().Get().Key("book:123").Cache(), time.Minute)

    // Track with multiple attributes
    authorLabeler, ok := valkeyotel.LabelerFromContext(ctx)
    if !ok {
        ctx = valkeyotel.ContextWithLabeler(ctx, authorLabeler)
    }
    authorLabeler.Add(
        attribute.String("key_pattern", "author"),
        attribute.String("tenant", "acme"),
    )
    client.DoCache(ctx, client.B().Get().Key("author:456").Cache(), time.Minute)
}
```

See [rueidishook](../rueidishook) if you want more customizations.

Note: `rueidisotel.NewClient` is not supported on go1.18 and go1.19 builds. [Reference](https://github.com/redis/rueidis/issues/442#issuecomment-1886993707)