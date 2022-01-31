package rueidis

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync/atomic"
	"testing"
	"time"

	"github.com/rueian/rueidis/internal/cmds"
)

func TestSingleClientPubSubReconnect(t *testing.T) {
	count := int64(0)
	errs := int64(0)
	m := &mockConn{
		DialFn: func() error { return nil },
		DoFn:   func(cmd cmds.Completed) RedisResult { return RedisResult{} },
	}
	_, err := newSingleClient(&ClientOption{
		InitAddress: []string{""},
		PubSubOption: NewPubSubOption(func(prev error, client DedicatedClient) {
			if prev != nil {
				atomic.AddInt64(&errs, 1)
			}
			if err := client.Do(context.Background(), client.B().Subscribe().Channel("a").Build()).Error(); err != nil {
				t.Errorf("unexpected subscribe err %v", err)
			}
			atomic.AddInt64(&count, 1)
		}, PubSubHandler{})}, nil, func(dst string, opt *ClientOption) conn {
		return m
	})
	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}
	m.TriggerDisconnect(errors.New("network")) // should trigger reconnect
	m.TriggerDisconnect(ErrClosing)            // should not trigger reconnect

	for atomic.LoadInt64(&count) != 2 {
		log.Printf("wait for pubsub reconnect count to be 2, got: %d\n", atomic.LoadInt64(&count))
		time.Sleep(time.Millisecond * 100)
	}

	if atomic.LoadInt64(&errs) != 1 {
		t.Fatalf("errs count should be 1")
	}
}

func TestClusterClientPubSubReconnect(t *testing.T) {
	count := int64(0)
	errs := int64(0)
	m := &mockConn{
		DialFn: func() error { return nil },
		DoFn:   func(cmd cmds.Completed) RedisResult { return slotsResp },
	}
	_, err := newClusterClient(&ClientOption{
		InitAddress: []string{":0"},
		PubSubOption: NewPubSubOption(func(prev error, client DedicatedClient) {
			if prev != nil {
				atomic.AddInt64(&errs, 1)
			}
			if err := client.Do(context.Background(), client.B().Subscribe().Channel("a").Build()).Error(); err != nil {
				t.Errorf("unexpected subscribe err %v", err)
			}
			atomic.AddInt64(&count, 1)
		}, PubSubHandler{}),
	}, func(dst string, opt *ClientOption) conn {
		return m
	})
	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}
	m.TriggerDisconnect(errors.New("network")) // should trigger reconnect
	m.TriggerDisconnect(ErrClosing)            // ErrClosing for cluster client should trigger reconnect

	for atomic.LoadInt64(&count) != 3 {
		log.Printf("wait for pubsub reconnect count to be 3, got: %d\n", atomic.LoadInt64(&count))
		time.Sleep(time.Millisecond * 100)
	}

	if atomic.LoadInt64(&errs) != 2 {
		t.Fatalf("errs count should be 2")
	}
}

func ExampleNewPubSubOption_subscribe() {
	messages := make(chan string, 100)

	ctx := context.Background()

	client, err := NewClient(ClientOption{
		InitAddress: []string{"127.0.0.1:6379"},
		PubSubOption: NewPubSubOption(func(prev error, client DedicatedClient) {
			if prev != nil {
				fmt.Printf("auto reconnected, previous err: %v\n", prev)
			}
			// do subscribe here
			client.Do(ctx, client.B().Subscribe().Channel("ch").Build())
		}, PubSubHandler{
			OnMessage: func(channel, message string) {
				// Users should avoid OnMessage blocking too long,
				// otherwise Client performance will decrease.
				messages <- message
			},
		}),
	})
	if err != nil {
		panic(err)
	}
	defer client.Close()

	for msg := range messages {
		fmt.Println(msg)
	}
}
