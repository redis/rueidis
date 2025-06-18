# rueidishook

With `rueidishook.WithHook`, users can easily intercept `rueidis.Client` by implementing custom `rueidishook.Hook` handler.

This can be useful to change the behavior of `rueidis.Client` or add other integrations such as observability, APM, etc.

## Example

```go
package main

import (
	"context"
	"time"

	"github.com/redis/rueidis"
	"github.com/redis/rueidis/rueidishook"
)

type hook struct{}

func (h *hook) Do(client rueidis.Client, ctx context.Context, cmd rueidis.Completed) (resp rueidis.RedisResult) {
	// do whatever you want before a client.Do
	resp = client.Do(ctx, cmd)
	// do whatever you want after a client.Do
	return
}

func (h *hook) DoMulti(client rueidis.Client, ctx context.Context, multi ...rueidis.Completed) (resps []rueidis.RedisResult) {
	// do whatever you want before a client.DoMulti
	resps = client.DoMulti(ctx, multi...)
	// do whatever you want after a client.DoMulti
	return
}

func (h *hook) DoCache(client rueidis.Client, ctx context.Context, cmd rueidis.Cacheable, ttl time.Duration) (resp rueidis.RedisResult) {
	// do whatever you want before a client.DoCache
	resp = client.DoCache(ctx, cmd, ttl)
	// do whatever you want after a client.DoCache
	return
}

func (h *hook) DoMultiCache(client rueidis.Client, ctx context.Context, multi ...rueidis.CacheableTTL) (resps []rueidis.RedisResult) {
	// do whatever you want before a client.DoMultiCache
	resps = client.DoMultiCache(ctx, multi...)
	// do whatever you want after a client.DoMultiCache
	return
}

func (h *hook) Receive(client rueidis.Client, ctx context.Context, subscribe rueidis.Completed, fn func(msg rueidis.PubSubMessage)) (err error) {
	// do whatever you want before a client.Receive
	err = client.Receive(ctx, subscribe, fn)
	// do whatever you want after a client.Receive
	return
}

func main() {
	client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
	if err != nil {
		panic(err)
	}
	client = rueidishook.WithHook(client, &hook{})
	defer client.Close()
}
```
