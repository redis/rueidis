package rueidis

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"testing"
	"time"
)

func accept(t *testing.T, ln net.Listener) (*redisMock, error) {
	conn, err := ln.Accept()
	if err != nil {
		t.Error(err)
		return nil, err
	}
	mock := &redisMock{
		t:    t,
		buf:  bufio.NewReader(conn),
		conn: conn,
	}
	mock.Expect("HELLO", "3").
		Reply(RedisMessage{
			typ: '%',
			values: []RedisMessage{
				{typ: '+', string: "key"},
				{typ: '+', string: "value"},
			},
		})
	mock.Expect("CLIENT", "TRACKING", "ON", "OPTIN").
		ReplyString("OK")
	return mock, nil
}

func TestNewClusterClient(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()
	go func() {
		mock, err := accept(t, ln)
		if err != nil {
			return
		}
		slots, _ := slotsResp.ToMessage()
		mock.Expect("CLUSTER", "SLOTS").Reply(slots)
		mock.Expect("QUIT").ReplyString("OK")
	}()

	_, port, _ := net.SplitHostPort(ln.Addr().String())
	client, err := NewClient(ClientOption{InitAddress: []string{"127.0.0.1:" + port}})
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()
	if _, ok := client.(*clusterClient); !ok {
		t.Fatal("client should be a singleClient")
	}
}

func TestFallBackSingleClient(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()
	go func() {
		mock, err := accept(t, ln)
		if err != nil {
			return
		}
		mock.Expect("CLUSTER", "SLOTS").Reply(RedisMessage{typ: '-', string: redisErrMsgClusterDisabled})
		mock.Expect("QUIT").ReplyString("OK")
		mock, err = accept(t, ln)
		if err != nil {
			return
		}
		mock.Expect("QUIT").ReplyString("OK")
	}()

	_, port, _ := net.SplitHostPort(ln.Addr().String())
	client, err := NewClient(ClientOption{InitAddress: []string{"127.0.0.1:" + port}})
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()
	if _, ok := client.(*singleClient); !ok {
		t.Fatal("client should be a singleClient")
	}
}

func ExampleIsRedisNil() {
	client, err := NewClient(ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
	if err != nil {
		panic(err)
	}
	defer client.Close()

	_, err = client.Do(context.Background(), client.B().Get().Key("not_exists").Build()).ToString()
	if err != nil && IsRedisNil(err) {
		fmt.Printf("it is a nil response")
	}
}

func ExampleClient_do() {
	client, err := NewClient(ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
	if err != nil {
		panic(err)
	}
	defer client.Close()

	ctx := context.Background()

	client.Do(ctx, client.B().Set().Key("k").Value("v").Build()).Error()

	client.Do(ctx, client.B().Get().Key("k").Build()).ToString()

	client.Do(ctx, client.B().Hmget().Key("h").Field("a", "b").Build()).AsMap()

	client.Do(ctx, client.B().Scard().Key("s").Build()).ToInt64()

	client.Do(ctx, client.B().Smembers().Key("s").Build()).ToArray()
}

func ExampleClient_doCache() {
	client, err := NewClient(ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
	if err != nil {
		panic(err)
	}
	defer client.Close()

	ctx := context.Background()

	client.DoCache(ctx, client.B().Get().Key("k").Cache(), time.Minute).ToString()

	client.DoCache(ctx, client.B().Hmget().Key("h").Field("a", "b").Cache(), time.Minute).ToMap()

	client.DoCache(ctx, client.B().Scard().Key("s").Cache(), time.Minute).ToInt64()

	client.DoCache(ctx, client.B().Smembers().Key("s").Cache(), time.Minute).ToArray()
}

func ExampleClient_dedicatedCAS() {
	client, err := NewClient(ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
	if err != nil {
		panic(err)
	}
	defer client.Close()

	ctx := context.Background()

	client.Dedicated(func(client DedicatedClient) error {
		// watch keys first
		if err := client.Do(ctx, client.B().Watch().Key("k1", "k2").Build()).Error(); err != nil {
			return err
		}
		// perform read here
		values, err := client.Do(ctx, client.B().Mget().Key("k1", "k2").Build()).ToArray()
		if err != nil {
			return err
		}
		v1, _ := values[0].ToString()
		v2, _ := values[1].ToString()
		// perform write with MULTI EXEC
		for _, resp := range client.DoMulti(
			ctx,
			client.B().Multi().Build(),
			client.B().Set().Key("k1").Value(v1+"1").Build(),
			client.B().Set().Key("k2").Value(v2+"2").Build(),
			client.B().Exec().Build(),
		) {
			if err := resp.Error(); err != nil {
				return err
			}
		}
		return nil
	})
}
