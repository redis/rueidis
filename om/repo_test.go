package om

import (
	"testing"

	"github.com/redis/rueidis"
)

func setup(t *testing.T) rueidis.Client {
	client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{"127.0.0.1:6377"}})
	if err != nil {
		t.Fatal(err)
	}
	return client
}

type TestStruct struct {
	Key string `redis:",key"`
	Ver int64  `redis:",ver"`
}

func TestWithIndexName(t *testing.T) {
	client := setup(t)
	defer client.Close()

	for _, repo := range []Repository[TestStruct]{
		NewHashRepository("custom_prefix", TestStruct{}, client, WithIndexName("custom_index")),
		NewJSONRepository("custom_prefix", TestStruct{}, client, WithIndexName("custom_index")),
	} {
		if repo.IndexName() != "custom_index" {
			t.Fail()
		}
	}
}
