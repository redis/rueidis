package rueidis

import (
	"bufio"
	"net"
	"testing"

	"github.com/rueian/rueidis/internal/proto"
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
		Reply(proto.Message{
			Type: '%',
			Values: []proto.Message{
				{Type: '+', String: "key"},
				{Type: '+', String: "value"},
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
		slots, _ := slotsResp.Value()
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
		mock.Expect("CLUSTER", "SLOTS").Reply(proto.Message{Type: '-', String: redisErrMsgClusterDisabled})
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

func TestIsRedisNil(t *testing.T) {
	if !IsRedisNil(&proto.RedisError{Type: '_'}) {
		t.Fatal("IsRedisNil fail")
	}
}
