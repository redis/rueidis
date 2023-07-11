# rueidis mock

Due to the design of the command builder, it is impossible for users to mock `rueidis.Client` for testing.

Therefore, rueidis provides an implemented one, based on the `gomock`, with some helpers
to make user writing tests more easily, including command matcher `mock.Match`, `mock.MatchFn` and `mock.Result` for faking redis responses.

## Examples

### Mock `client.Do`

```go
package main

import (
	"context"
	"testing"

	"go.uber.org/mock/gomock"
	"github.com/redis/rueidis/mock"
)

func TestWithRueidis(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	client := mock.NewClient(ctrl)

	client.EXPECT().Do(ctx, mock.Match("GET", "key")).Return(mock.Result(mock.RedisString("val")))
	if v, _ := client.Do(ctx, client.B().Get().Key("key").Build()).ToString(); v != "val" {
		t.Fatalf("unexpected val %v", v)
	}
	client.EXPECT().DoMulti(ctx, mock.Match("GET", "c"), mock.Match("GET", "d")).Return([]rueidis.RedisResult{
		mock.Result(mock.RedisNil()),
		mock.Result(mock.RedisNil()),
	})
	for _, resp := range client.DoMulti(ctx, client.B().Get().Key("c").Build(), client.B().Get().Key("d").Build()) {
		if err := resp.Error(); !rueidis.IsRedisNil(err) {
			t.Fatalf("unexpected err %v", err)
		}
	}
}
```

### Mock `client.Receive`

```go
package main

import (
	"context"
	"testing"

	"go.uber.org/mock/gomock"
	"github.com/redis/rueidis/mock"
)

func TestWithRueidisReceive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	client := mock.NewClient(ctrl)

	client.EXPECT().Receive(ctx, mock.Match("SUBSCRIBE", "ch"), gomock.Any()).Do(func(_, _ any, fn func(message rueidis.PubSubMessage)) {
		fn(rueidis.PubSubMessage{Message: "msg"})
	})

	client.Receive(ctx, client.B().Subscribe().Channel("ch").Build(), func(msg rueidis.PubSubMessage) {
		if msg.Message != "msg" {
			t.Fatalf("unexpected val %v", msg.Message)
		}
	})
}
```

