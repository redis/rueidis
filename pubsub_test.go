package rueidis

import (
	"context"
	"errors"
	"log"
	"sync/atomic"
	"testing"
	"time"

	"github.com/rueian/rueidis/internal/cmds"
	"github.com/rueian/rueidis/internal/proto"
)

func TestSingleClientPubSubReconnect(t *testing.T) {
	count := int64(0)
	errs := int64(0)
	m := &mockConn{
		DialFn: func() error { return nil },
		DoFn:   func(cmd cmds.Completed) proto.Result { return proto.Result{} },
	}
	_, err := newSingleClient(ClientOption{
		InitAddress: []string{""},
		PubSubOption: NewPubSubOption(func(prev error, client DedicatedClient) {
			if prev != nil {
				atomic.AddInt64(&errs, 1)
			}
			if err := client.Do(context.Background(), client.B().Subscribe().Channel("a").Build()).Error(); err != nil {
				t.Errorf("unexpected subscribe err %v", err)
			}
			atomic.AddInt64(&count, 1)
		}, PubSubHandler{})}, func(dst string, opt ClientOption) conn {
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
		DoFn:   func(cmd cmds.Completed) proto.Result { return slotsResp },
	}
	_, err := newClusterClient(ClientOption{
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
	}, func(dst string, opt ClientOption) conn {
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
