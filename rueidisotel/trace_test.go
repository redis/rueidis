package rueidisotel

import (
	"github.com/rueian/rueidis"
)

func ExampleWithClient_openTelemetry() {
	client, _ := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
	client = WithClient(client)
	defer client.Close()
}
