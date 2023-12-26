# OpenTelemetry Tracing & Connection Metrics

Use `rueidisotel.WithClient` to create a client with OpenTelemetry Tracing enabled.

Use `rueidisotel.NewClient` to create a client with OpenTelemetry Tracing and Connection Metrics enabled.
default metrics are:
- `rueidis_dial_attempt`: number of dial attempts
- `rueidis_dial_success`: number of successful dials
- `rueidis_dial_conns`: number of connections
- `rueidis_dial_latency`: dial latency in seconds

```golang
package main

import (
    "github.com/redis/rueidis"
    "github.com/redis/rueidis/rueidisotel"
)

func main() {
    client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
    if err != nil {
        panic(err)
    }
    client = rueidisotel.WithClient(client)
    defer client.Close()
}
```

See [rueidishook](../rueidishook) if you want more customizations.