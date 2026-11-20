package rueidiscompat

import (
	"context"
	"reflect"
	"testing"

	"github.com/redis/rueidis"
	"github.com/redis/rueidis/mock"
	"go.uber.org/mock/gomock"
)

func TestTSMAddArguments(t *testing.T) {
	for _, tt := range []struct {
		name      string
		timestamp any
		value     any
		args      []string
	}{
		{"integers", 1, 2, []string{"TS.MADD", "{ts}:1", "1", "2", "{ts}:2", "1", "2"}},
		{"fractional value", 1, 1.25, []string{"TS.MADD", "{ts}:1", "1", "1.25", "{ts}:2", "1", "1.25"}},
		{"int64 timestamp", int64(1700000000000), 2, []string{"TS.MADD", "{ts}:1", "1700000000000", "2", "{ts}:2", "1700000000000", "2"}},
		{"server timestamp", "*", 1.25, []string{"TS.MADD", "{ts}:1", "*", "1.25", "{ts}:2", "*", "1.25"}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if v := recover(); v != nil {
					t.Errorf("TSMAdd panicked: %v", v)
				}
			}()
			client := mock.NewClient(gomock.NewController(t), mock.WithSlotCheck())
			expected := client.B().Get().Key("{ts}:1").Build()
			client.EXPECT().Do(gomock.Any(), mock.Match(tt.args...)).DoAndReturn(
				func(_ context.Context, cmd rueidis.Completed) rueidis.RedisResult {
					if cmd.Slot() != expected.Slot() {
						t.Errorf("slot = %d, want %d", cmd.Slot(), expected.Slot())
					}
					return mock.Result(mock.RedisArray(mock.RedisInt64(1), mock.RedisInt64(1)))
				},
			).AnyTimes()
			got, err := NewAdapter(client).TSMAdd(context.Background(), [][]interface{}{
				{"{ts}:1", tt.timestamp, tt.value},
				{"{ts}:2", tt.timestamp, tt.value},
			}).Result()
			if err != nil || !reflect.DeepEqual(got, []int64{1, 1}) {
				t.Fatalf("got %v, %v", got, err)
			}
		})
	}
}

func TestTSMAddCrossSlot(t *testing.T) {
	client := mock.NewClient(gomock.NewController(t), mock.WithSlotCheck())
	client.EXPECT().Do(gomock.Any(), gomock.Any()).Return(mock.Result(mock.RedisArray())).AnyTimes()
	defer func() {
		if v := recover(); v != "multi key command with different key slots are not allowed" {
			t.Errorf("expected cross-slot panic, got %v", v)
		}
	}()
	NewAdapter(client).TSMAdd(context.Background(), [][]interface{}{
		{"{first}:1", 1, 2},
		{"{second}:2", 1, 2},
	})
}
