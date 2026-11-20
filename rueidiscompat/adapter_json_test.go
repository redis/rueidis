package rueidiscompat

import (
	"context"
	"testing"

	"github.com/redis/rueidis"
	"github.com/redis/rueidis/mock"
	"go.uber.org/mock/gomock"
)

func TestJSONMSetSlot(t *testing.T) {
	for _, tt := range []struct {
		name string
		set  func(Cmdable, string, string) *StatusCmd
	}{
		{
			name: "JSONMSet",
			set: func(c Cmdable, first, second string) *StatusCmd {
				return c.JSONMSet(context.Background(), first, "$", `{"a":1}`, second, "$", `{"b":2}`)
			},
		},
		{
			name: "JSONMSetArgs",
			set: func(c Cmdable, first, second string) *StatusCmd {
				return c.JSONMSetArgs(context.Background(), []JSONSetArgs{
					{Key: first, Path: "$", Value: `{"a":1}`},
					{Key: second, Path: "$", Value: `{"b":2}`},
				})
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Run("same slot", func(t *testing.T) {
				client := mock.NewClient(gomock.NewController(t), mock.WithSlotCheck())
				expected := client.B().Get().Key("{json}:1").Build()
				wantSlot := expected.Slot()
				client.EXPECT().Do(gomock.Any(), mock.Match("JSON.MSET", "{json}:1", "$", `{"a":1}`, "{json}:2", "$", `{"b":2}`)).DoAndReturn(
					func(_ context.Context, cmd rueidis.Completed) rueidis.RedisResult {
						if cmd.Slot() != wantSlot {
							t.Errorf("slot = %d, want %d", cmd.Slot(), wantSlot)
						}
						return mock.Result(mock.RedisString("OK"))
					},
				)
				if err := tt.set(NewAdapter(client), "{json}:1", "{json}:2").Err(); err != nil {
					t.Fatal(err)
				}
			})
			t.Run("cross slot", func(t *testing.T) {
				client := mock.NewClient(gomock.NewController(t), mock.WithSlotCheck())
				client.EXPECT().Do(gomock.Any(), gomock.Any()).Return(mock.Result(mock.RedisString("OK"))).AnyTimes()
				defer func() {
					if recover() == nil {
						t.Error("expected a cross-slot panic")
					}
				}()
				tt.set(NewAdapter(client), "{first}:1", "{second}:2")
			})
		})
	}
}
