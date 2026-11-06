package rueidiscompat

import (
	"context"
	"testing"
	"time"

	"github.com/redis/rueidis/mock"
	"go.uber.org/mock/gomock"
)

func TestPExpireAtFarFuture(t *testing.T) {
	client := mock.NewClient(gomock.NewController(t))
	client.EXPECT().Do(gomock.Any(), mock.Match("PEXPIREAT", "key", "32503680000123")).Return(mock.Result(mock.RedisInt64(1)))
	expires := time.Date(3000, time.January, 1, 0, 0, 0, 123000000, time.UTC)
	if err := NewAdapter(client).PExpireAt(context.Background(), "key", expires).Err(); err != nil {
		t.Fatal(err)
	}
}
