# OpenTelemetry Tracing

Use `rueidisotel.WithClient` to create a client with OpenTelemetry Tracing enabled.

```golang
package main

import (
    "github.com/rueian/rueidis"
    "github.com/rueian/rueidis/rueidisotel"
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