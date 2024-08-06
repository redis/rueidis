package rueidiscompat

import (
	"context"
	"fmt"

	"github.com/redis/rueidis"
)

func ExampleScript() {
	ctx = context.Background()
	IncrByXX := NewScript(`
		if redis.call("GET", KEYS[1]) ~= false then
			return redis.call("INCRBY", KEYS[1], ARGV[1])
		end
		return false
	`)
	client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
	if err != nil {
		panic(err)
	}
	defer client.Close()
	rdb := NewAdapter(client)
	n, err := IncrByXX.Run(ctx, rdb, []string{"xx_counter"}, 2).Result()
	fmt.Println(n, err)

	err = rdb.Set(ctx, "xx_counter", "40", 0).Err()
	if err != nil {
		panic(err)
	}

	n, err = IncrByXX.Run(ctx, rdb, []string{"xx_counter"}, 2).Result()
	fmt.Println(n, err)
}
