package om

import (
	"testing"

	"github.com/rueian/rueidis"
)

func setup(t *testing.T) rueidis.Client {
	client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
	if err != nil {
		t.Fatal(err)
	}
	return client
}
