package rueidis

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

var slotsResp = newResult(RedisMessage{typ: '*', values: []RedisMessage{
	{typ: '*', values: []RedisMessage{
		{typ: ':', integer: 0},
		{typ: ':', integer: 16383},
		{typ: '*', values: []RedisMessage{ // master
			{typ: '+', string: "127.0.0.1"},
			{typ: ':', integer: 0},
			{typ: '+', string: ""},
		}},
		{typ: '*', values: []RedisMessage{ // replica
			{typ: '+', string: "127.0.1.1"},
			{typ: ':', integer: 1},
			{typ: '+', string: ""},
		}},
	}},
}}, nil)

var slotsMultiResp = newResult(RedisMessage{typ: '*', values: []RedisMessage{
	{typ: '*', values: []RedisMessage{
		{typ: ':', integer: 0},
		{typ: ':', integer: 8192},
		{typ: '*', values: []RedisMessage{ // master
			{typ: '+', string: "127.0.0.1"},
			{typ: ':', integer: 0},
			{typ: '+', string: ""},
		}},
		{typ: '*', values: []RedisMessage{ // replica
			{typ: '+', string: "127.0.1.1"},
			{typ: ':', integer: 1},
			{typ: '+', string: ""},
		}},
	}},
	{typ: '*', values: []RedisMessage{
		{typ: ':', integer: 8193},
		{typ: ':', integer: 16383},
		{typ: '*', values: []RedisMessage{ // master
			{typ: '+', string: "127.0.2.1"},
			{typ: ':', integer: 0},
			{typ: '+', string: ""},
		}},
		{typ: '*', values: []RedisMessage{ // replica
			{typ: '+', string: "127.0.3.1"},
			{typ: ':', integer: 1},
			{typ: '+', string: ""},
		}},
	}},
}}, nil)

var slotsMultiRespWithoutReplicas = newResult(RedisMessage{typ: '*', values: []RedisMessage{
	{typ: '*', values: []RedisMessage{
		{typ: ':', integer: 0},
		{typ: ':', integer: 8192},
		{typ: '*', values: []RedisMessage{ // master
			{typ: '+', string: "127.0.0.1"},
			{typ: ':', integer: 0},
			{typ: '+', string: ""},
		}},
	}},
	{typ: '*', values: []RedisMessage{
		{typ: ':', integer: 8193},
		{typ: ':', integer: 16383},
		{typ: '*', values: []RedisMessage{ // master
			{typ: '+', string: "127.0.1.1"},
			{typ: ':', integer: 0},
			{typ: '+', string: ""},
		}},
	}},
}}, nil)

var slotsMultiRespWithMultiReplicas = newResult(RedisMessage{typ: '*', values: []RedisMessage{
	{typ: '*', values: []RedisMessage{
		{typ: ':', integer: 0},
		{typ: ':', integer: 8192},
		{typ: '*', values: []RedisMessage{ // master
			{typ: '+', string: "127.0.0.1"},
			{typ: ':', integer: 0},
			{typ: '+', string: ""},
		}},
		{typ: '*', values: []RedisMessage{ // replica1
			{typ: '+', string: "127.0.0.2"},
			{typ: ':', integer: 1},
			{typ: '+', string: ""},
		}},
		{typ: '*', values: []RedisMessage{ // replica2
			{typ: '+', string: "127.0.0.3"},
			{typ: ':', integer: 2},
			{typ: '+', string: ""},
		}},
		{typ: '*', values: []RedisMessage{ // replica3
			{typ: '+', string: "127.0.0.4"},
			{typ: ':', integer: 3},
			{typ: '+', string: ""},
		}},
	}},
	{typ: '*', values: []RedisMessage{
		{typ: ':', integer: 8193},
		{typ: ':', integer: 16383},
		{typ: '*', values: []RedisMessage{ // master
			{typ: '+', string: "127.0.1.1"},
			{typ: ':', integer: 0},
			{typ: '+', string: ""},
		}},
		{typ: '*', values: []RedisMessage{ // replica1
			{typ: '+', string: "127.0.1.2"},
			{typ: ':', integer: 1},
			{typ: '+', string: ""},
		}},
		{typ: '*', values: []RedisMessage{ // replica2
			{typ: '+', string: "127.0.1.3"},
			{typ: ':', integer: 2},
			{typ: '+', string: ""},
		}},
		{typ: '*', values: []RedisMessage{ // replica3
			{typ: '+', string: "127.0.1.4"},
			{typ: ':', integer: 3},
			{typ: '+', string: ""},
		}},
	}},
}}, nil)

var singleSlotResp = newResult(RedisMessage{typ: '*', values: []RedisMessage{
	{typ: '*', values: []RedisMessage{
		{typ: ':', integer: 0},
		{typ: ':', integer: 0},
		{typ: '*', values: []RedisMessage{ // master
			{typ: '+', string: "127.0.0.1"},
			{typ: ':', integer: 0},
			{typ: '+', string: ""},
		}},
	}},
}}, nil)

var singleSlotResp2 = newResult(RedisMessage{typ: '*', values: []RedisMessage{
	{typ: '*', values: []RedisMessage{
		{typ: ':', integer: 0},
		{typ: ':', integer: 0},
		{typ: '*', values: []RedisMessage{ // master
			{typ: '+', string: "127.0.3.1"},
			{typ: ':', integer: 3},
			{typ: '+', string: ""},
		}},
	}},
}}, nil)

var singleSlotWithoutIP = newResult(RedisMessage{typ: '*', values: []RedisMessage{
	{typ: '*', values: []RedisMessage{
		{typ: ':', integer: 0},
		{typ: ':', integer: 0},
		{typ: '*', values: []RedisMessage{ // master
			{typ: '+', string: ""},
			{typ: ':', integer: 4},
			{typ: '+', string: ""},
		}},
		{typ: '*', values: []RedisMessage{ // replica
			{typ: '+', string: "?"},
			{typ: ':', integer: 1},
			{typ: '+', string: ""},
		}},
	}},
	{typ: '*', values: []RedisMessage{
		{typ: ':', integer: 0},
		{typ: ':', integer: 0},
		{typ: '*', values: []RedisMessage{ // master
			{typ: '+', string: "?"},
			{typ: ':', integer: 4},
			{typ: '+', string: ""},
		}},
	}},
}}, nil)

var shardsResp = newResult(RedisMessage{typ: typeArray, values: []RedisMessage{
	{typ: typeMap, values: []RedisMessage{
		{typ: typeBlobString, string: "slots"},
		{typ: typeArray, values: []RedisMessage{
			{typ: typeBlobString, string: "0"},
			{typ: typeBlobString, string: "16383"},
		}},
		{typ: typeBlobString, string: "nodes"},
		{typ: typeArray, values: []RedisMessage{
			{typ: typeMap, values: []RedisMessage{ // master
				{typ: typeBlobString, string: "id"},
				{typ: typeBlobString, string: ""},

				{typ: typeBlobString, string: "port"},
				{typ: typeInteger, integer: 0},

				{typ: typeBlobString, string: "ip"},
				{typ: typeBlobString, string: "127.0.0.1"},

				{typ: typeBlobString, string: "endpoint"},
				{typ: typeBlobString, string: "127.0.0.1"},

				{typ: typeBlobString, string: "role"},
				{typ: typeBlobString, string: "master"},

				{typ: typeBlobString, string: "replication-offset"},
				{typ: typeInteger, integer: 72156},

				{typ: typeBlobString, string: "health"},
				{typ: typeBlobString, string: "online"},
			}},
			{typ: typeMap, values: []RedisMessage{ // replica
				{typ: typeBlobString, string: "id"},
				{typ: typeBlobString, string: ""},

				{typ: typeBlobString, string: "port"},
				{typ: typeInteger, integer: 1},

				{typ: typeBlobString, string: "ip"},
				{typ: typeBlobString, string: "127.0.1.1"},

				{typ: typeBlobString, string: "endpoint"},
				{typ: typeBlobString, string: "127.0.1.1"},

				{typ: typeBlobString, string: "role"},
				{typ: typeBlobString, string: "replica"},

				{typ: typeBlobString, string: "replication-offset"},
				{typ: typeInteger, integer: 72156},

				{typ: typeBlobString, string: "health"},
				{typ: typeBlobString, string: "online"},
			}},
		}},
	}},
}}, nil)

var shardsRespTls = newResult(RedisMessage{typ: typeArray, values: []RedisMessage{
	{typ: typeMap, values: []RedisMessage{
		{typ: typeBlobString, string: "slots"},
		{typ: typeArray, values: []RedisMessage{
			{typ: typeBlobString, string: "0"},
			{typ: typeBlobString, string: "16383"},
		}},
		{typ: typeBlobString, string: "nodes"},
		{typ: typeArray, values: []RedisMessage{
			{typ: typeMap, values: []RedisMessage{ // replica, tls
				{typ: typeBlobString, string: "id"},
				{typ: typeBlobString, string: ""},

				{typ: typeBlobString, string: "tls-port"},
				{typ: typeInteger, integer: 2},

				{typ: typeBlobString, string: "ip"},
				{typ: typeBlobString, string: "127.0.2.1"},

				{typ: typeBlobString, string: "endpoint"},
				{typ: typeBlobString, string: "127.0.2.1"},

				{typ: typeBlobString, string: "role"},
				{typ: typeBlobString, string: "replica"},

				{typ: typeBlobString, string: "replication-offset"},
				{typ: typeInteger, integer: 72156},

				{typ: typeBlobString, string: "health"},
				{typ: typeBlobString, string: "online"},
			}},
			{typ: typeMap, values: []RedisMessage{ // master, tls + port
				{typ: typeBlobString, string: "id"},
				{typ: typeBlobString, string: ""},

				{typ: typeBlobString, string: "port"},
				{typ: typeInteger, integer: 0},

				{typ: typeBlobString, string: "tls-port"},
				{typ: typeInteger, integer: 1},

				{typ: typeBlobString, string: "ip"},
				{typ: typeBlobString, string: "127.0.1.1"},

				{typ: typeBlobString, string: "endpoint"},
				{typ: typeBlobString, string: "127.0.1.1"},

				{typ: typeBlobString, string: "role"},
				{typ: typeBlobString, string: "master"},

				{typ: typeBlobString, string: "replication-offset"},
				{typ: typeInteger, integer: 72156},

				{typ: typeBlobString, string: "health"},
				{typ: typeBlobString, string: "online"},
			}},
			{typ: typeMap, values: []RedisMessage{ // replica, port
				{typ: typeBlobString, string: "id"},
				{typ: typeBlobString, string: ""},

				{typ: typeBlobString, string: "port"},
				{typ: typeInteger, integer: 3},

				{typ: typeBlobString, string: "ip"},
				{typ: typeBlobString, string: "127.0.3.1"},

				{typ: typeBlobString, string: "endpoint"},
				{typ: typeBlobString, string: "127.0.3.1"},

				{typ: typeBlobString, string: "role"},
				{typ: typeBlobString, string: "replica"},

				{typ: typeBlobString, string: "replication-offset"},
				{typ: typeInteger, integer: 72156},

				{typ: typeBlobString, string: "health"},
				{typ: typeBlobString, string: "online"},
			}},
		}},
	}},
}}, nil)

var shardsMultiResp = newResult(RedisMessage{typ: '*', values: []RedisMessage{
	{typ: typeMap, values: []RedisMessage{
		{typ: typeBlobString, string: "slots"},
		{typ: typeArray, values: []RedisMessage{
			{typ: typeBlobString, string: "0"},
			{typ: typeBlobString, string: "8192"},
		}},
		{typ: typeBlobString, string: "nodes"},
		{typ: typeArray, values: []RedisMessage{
			{typ: typeMap, values: []RedisMessage{ // master
				{typ: typeBlobString, string: "id"},
				{typ: typeBlobString, string: ""},

				{typ: typeBlobString, string: "port"},
				{typ: typeInteger, integer: 0},

				{typ: typeBlobString, string: "ip"},
				{typ: typeBlobString, string: "127.0.0.1"},

				{typ: typeBlobString, string: "endpoint"},
				{typ: typeBlobString, string: "127.0.0.1"},

				{typ: typeBlobString, string: "role"},
				{typ: typeBlobString, string: "master"},

				{typ: typeBlobString, string: "replication-offset"},
				{typ: typeInteger, integer: 72156},

				{typ: typeBlobString, string: "health"},
				{typ: typeBlobString, string: "online"},
			}},
			{typ: typeMap, values: []RedisMessage{ // replica
				{typ: typeBlobString, string: "id"},
				{typ: typeBlobString, string: ""},

				{typ: typeBlobString, string: "port"},
				{typ: typeInteger, integer: 1},

				{typ: typeBlobString, string: "ip"},
				{typ: typeBlobString, string: "127.0.1.1"},

				{typ: typeBlobString, string: "endpoint"},
				{typ: typeBlobString, string: "127.0.1.1"},

				{typ: typeBlobString, string: "role"},
				{typ: typeBlobString, string: "replica"},

				{typ: typeBlobString, string: "replication-offset"},
				{typ: typeInteger, integer: 72156},

				{typ: typeBlobString, string: "health"},
				{typ: typeBlobString, string: "online"},
			}},
		}},
	}},
	{typ: typeMap, values: []RedisMessage{
		{typ: typeBlobString, string: "slots"},
		{typ: typeArray, values: []RedisMessage{
			{typ: typeBlobString, string: "8193"},
			{typ: typeBlobString, string: "16383"},
		}},
		{typ: typeBlobString, string: "nodes"},
		{typ: typeArray, values: []RedisMessage{
			{typ: typeMap, values: []RedisMessage{ // master
				{typ: typeBlobString, string: "id"},
				{typ: typeBlobString, string: ""},

				{typ: typeBlobString, string: "port"},
				{typ: typeInteger, integer: 0},

				{typ: typeBlobString, string: "ip"},
				{typ: typeBlobString, string: "127.0.2.1"},

				{typ: typeBlobString, string: "endpoint"},
				{typ: typeBlobString, string: "127.0.2.1"},

				{typ: typeBlobString, string: "role"},
				{typ: typeBlobString, string: "master"},

				{typ: typeBlobString, string: "replication-offset"},
				{typ: typeInteger, integer: 72156},

				{typ: typeBlobString, string: "health"},
				{typ: typeBlobString, string: "online"},
			}},
			{typ: typeMap, values: []RedisMessage{ // replica
				{typ: typeBlobString, string: "id"},
				{typ: typeBlobString, string: ""},

				{typ: typeBlobString, string: "port"},
				{typ: typeInteger, integer: 1},

				{typ: typeBlobString, string: "ip"},
				{typ: typeBlobString, string: "127.0.3.1"},

				{typ: typeBlobString, string: "endpoint"},
				{typ: typeBlobString, string: "127.0.3.1"},

				{typ: typeBlobString, string: "role"},
				{typ: typeBlobString, string: "replica"},

				{typ: typeBlobString, string: "replication-offset"},
				{typ: typeInteger, integer: 72156},

				{typ: typeBlobString, string: "health"},
				{typ: typeBlobString, string: "online"},
			}},
		}},
	}},
}}, nil)

var singleShardResp2 = newResult(RedisMessage{typ: '*', values: []RedisMessage{
	{typ: typeMap, values: []RedisMessage{
		{typ: typeBlobString, string: "slots"},
		{typ: typeArray, values: []RedisMessage{
			{typ: typeBlobString, string: "0"},
			{typ: typeBlobString, string: "0"},
		}},
		{typ: typeBlobString, string: "nodes"},
		{typ: typeArray, values: []RedisMessage{
			{typ: typeMap, values: []RedisMessage{ // master
				{typ: typeBlobString, string: "id"},
				{typ: typeBlobString, string: ""},

				{typ: typeBlobString, string: "port"},
				{typ: typeInteger, integer: 3},

				{typ: typeBlobString, string: "ip"},
				{typ: typeBlobString, string: "127.0.3.1"},

				{typ: typeBlobString, string: "endpoint"},
				{typ: typeBlobString, string: "127.0.3.1"},

				{typ: typeBlobString, string: "role"},
				{typ: typeBlobString, string: "master"},

				{typ: typeBlobString, string: "replication-offset"},
				{typ: typeInteger, integer: 72156},

				{typ: typeBlobString, string: "health"},
				{typ: typeBlobString, string: "online"},
			}},
		}},
	}},
}}, nil)

var singleShardWithoutIP = newResult(RedisMessage{typ: typeArray, values: []RedisMessage{
	{typ: typeMap, values: []RedisMessage{
		{typ: typeBlobString, string: "slots"},
		{typ: typeArray, values: []RedisMessage{
			{typ: typeBlobString, string: "0"},
			{typ: typeBlobString, string: "0"},
		}},
		{typ: typeBlobString, string: "nodes"},
		{typ: typeArray, values: []RedisMessage{
			{typ: typeMap, values: []RedisMessage{ // master
				{typ: typeBlobString, string: "id"},
				{typ: typeBlobString, string: ""},

				{typ: typeBlobString, string: "port"},
				{typ: typeInteger, integer: 4},

				{typ: typeBlobString, string: "ip"},
				{typ: typeBlobString, string: ""},

				{typ: typeBlobString, string: "endpoint"},
				{typ: typeBlobString, string: ""},

				{typ: typeBlobString, string: "role"},
				{typ: typeBlobString, string: "master"},

				{typ: typeBlobString, string: "replication-offset"},
				{typ: typeInteger, integer: 72156},

				{typ: typeBlobString, string: "health"},
				{typ: typeBlobString, string: "online"},
			}},
			{typ: typeMap, values: []RedisMessage{ // replica
				{typ: typeBlobString, string: "id"},
				{typ: typeBlobString, string: ""},

				{typ: typeBlobString, string: "port"},
				{typ: typeInteger, integer: 1},

				{typ: typeBlobString, string: "ip"},
				{typ: typeBlobString, string: "?"},

				{typ: typeBlobString, string: "endpoint"},
				{typ: typeBlobString, string: "?"},

				{typ: typeBlobString, string: "role"},
				{typ: typeBlobString, string: "replica"},

				{typ: typeBlobString, string: "replication-offset"},
				{typ: typeInteger, integer: 72156},

				{typ: typeBlobString, string: "health"},
				{typ: typeBlobString, string: "online"},
			}},
		}},
	}},
	{typ: typeMap, values: []RedisMessage{
		{typ: typeBlobString, string: "slots"},
		{typ: typeArray, values: []RedisMessage{
			{typ: typeBlobString, string: "0"},
			{typ: typeBlobString, string: "0"},
		}},
		{typ: typeBlobString, string: "nodes"},
		{typ: typeArray, values: []RedisMessage{
			{typ: typeMap, values: []RedisMessage{ // master
				{typ: typeBlobString, string: "id"},
				{typ: typeBlobString, string: ""},

				{typ: typeBlobString, string: "port"},
				{typ: typeInteger, integer: 4},

				{typ: typeBlobString, string: "ip"},
				{typ: typeBlobString, string: "?"},

				{typ: typeBlobString, string: "endpoint"},
				{typ: typeBlobString, string: "?"},

				{typ: typeBlobString, string: "role"},
				{typ: typeBlobString, string: "master"},

				{typ: typeBlobString, string: "replication-offset"},
				{typ: typeInteger, integer: 72156},

				{typ: typeBlobString, string: "health"},
				{typ: typeBlobString, string: "online"},
			}},
		}},
	}},
}}, nil)

//gocyclo:ignore
func TestClusterClientInit(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	t.Run("Init no nodes", func(t *testing.T) {
		if _, err := newClusterClient(&ClientOption{InitAddress: []string{}}, func(dst string, opt *ClientOption) conn { return nil }); err != ErrNoAddr {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("Init no dialable", func(t *testing.T) {
		v := errors.New("dial err")
		if _, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return &mockConn{DialFn: func() error { return v }}
		}); err != v {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("Refresh err", func(t *testing.T) {
		v := errors.New("refresh err")
		if _, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return &mockConn{DoFn: func(cmd Completed) RedisResult { return newErrResult(v) }}
		}); err != v {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("Refresh skip zero slots", func(t *testing.T) {
		var first int64
		if _, err := newClusterClient(&ClientOption{InitAddress: []string{"127.0.0.1:0", "127.0.1.1:1"}}, func(dst string, opt *ClientOption) conn {
			return &mockConn{
				DoFn: func(cmd Completed) RedisResult {
					if atomic.AddInt64(&first, 1) == 1 {
						return newResult(RedisMessage{typ: '*', values: []RedisMessage{}}, nil)
					}
					return slotsResp
				},
			}
		}); err != nil || atomic.AddInt64(&first, 1) < 2 {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("Refresh skip zero shards", func(t *testing.T) {
		var first int64
		if _, err := newClusterClient(&ClientOption{InitAddress: []string{"127.0.0.1:0", "127.0.1.1:1"}}, func(dst string, opt *ClientOption) conn {
			return &mockConn{
				DoFn: func(cmd Completed) RedisResult {
					if atomic.AddInt64(&first, 1) == 1 {
						return newResult(RedisMessage{typ: '*', values: []RedisMessage{}}, nil)
					}
					return shardsResp
				},
				VersionFn: func() int { return 7 },
			}
		}); err != nil || atomic.AddInt64(&first, 1) < 2 {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("Refresh no slots cluster", func(t *testing.T) {
		if _, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return &mockConn{
				DoFn: func(cmd Completed) RedisResult {
					return newResult(RedisMessage{typ: '*', values: []RedisMessage{}}, nil)
				},
			}
		}); err != nil {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("Refresh no shards cluster", func(t *testing.T) {
		if _, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return &mockConn{
				DoFn: func(cmd Completed) RedisResult {
					return newResult(RedisMessage{typ: '*', values: []RedisMessage{}}, nil)
				},
				VersionFn: func() int { return 7 },
			}
		}); err != nil {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("Refresh cluster of 1 node without knowing its own ip", func(t *testing.T) {
		getClient := func(version int) (client *clusterClient, err error) {
			return newClusterClient(&ClientOption{InitAddress: []string{"127.0.4.1:4"}}, func(dst string, opt *ClientOption) conn {
				return &mockConn{
					DoFn: func(cmd Completed) RedisResult {
						if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
							return singleSlotWithoutIP
						}
						return singleShardWithoutIP
					},
					AddrFn:    func() string { return "127.0.4.1:4" },
					VersionFn: func() int { return version },
				}
			})
		}

		t.Run("slots", func(t *testing.T) {
			client, err := getClient(6)
			if err != nil {
				t.Fatalf("unexpected err %v", err)
			}
			nodes := client.nodes()
			sort.Strings(nodes)
			if len(nodes) != 1 ||
				nodes[0] != "127.0.4.1:4" {
				t.Fatalf("unexpected nodes %v", nodes)
			}
		})

		t.Run("shards", func(t *testing.T) {
			client, err := getClient(7)
			if err != nil {
				t.Fatalf("unexpected err %v", err)
			}
			nodes := client.nodes()
			sort.Strings(nodes)
			if len(nodes) != 1 ||
				nodes[0] != "127.0.4.1:4" {
				t.Fatalf("unexpected nodes %v", nodes)
			}
		})
	})

	t.Run("Refresh replace", func(t *testing.T) {
		testFunc := func(t *testing.T, client *clusterClient, num *int64) {
			nodes := client.nodes()
			sort.Strings(nodes)
			if len(nodes) != 3 ||
				nodes[0] != "127.0.0.1:0" ||
				nodes[1] != "127.0.1.1:1" ||
				nodes[2] != "127.0.2.1:2" {
				t.Fatalf("unexpected nodes %v", nodes)
			}

			atomic.AddInt64(num, 1)

			if err := client.refresh(context.Background()); err != nil {
				t.Fatalf("unexpected err %v", err)
			}

			nodes = client.nodes()
			sort.Strings(nodes)
			if len(nodes) != 3 ||
				nodes[0] != "127.0.1.1:1" ||
				nodes[1] != "127.0.2.1:2" ||
				nodes[2] != "127.0.3.1:3" {
				t.Fatalf("unexpected nodes %v", nodes)
			}
		}

		t.Run("slots", func(t *testing.T) {
			var first int64
			client, err := newClusterClient(&ClientOption{InitAddress: []string{"127.0.1.1:1", "127.0.2.1:2"}}, func(dst string, opt *ClientOption) conn {
				return &mockConn{
					DoFn: func(cmd Completed) RedisResult {
						if atomic.LoadInt64(&first) == 1 {
							return singleSlotResp2
						}
						return slotsResp
					},
				}
			})
			if err != nil {
				t.Fatalf("unexpected err %v", err)
			}
			testFunc(t, client, &first)
		})

		t.Run("shards", func(t *testing.T) {
			var first int64
			client, err := newClusterClient(&ClientOption{InitAddress: []string{"127.0.1.1:1", "127.0.2.1:2"}}, func(dst string, opt *ClientOption) conn {
				return &mockConn{
					DoFn: func(cmd Completed) RedisResult {
						if atomic.LoadInt64(&first) == 1 {
							return singleShardResp2
						}
						return shardsResp
					},
					VersionFn: func() int { return 7 },
				}
			})
			if err != nil {
				t.Fatalf("unexpected err %v", err)
			}
			testFunc(t, client, &first)
		})
	})

	t.Run("Shards tls", func(t *testing.T) {
		client, err := newClusterClient(&ClientOption{InitAddress: []string{"127.0.0.1:0"}, TLSConfig: &tls.Config{}}, func(dst string, opt *ClientOption) conn {
			return &mockConn{
				DoFn: func(cmd Completed) RedisResult {
					return shardsRespTls
				},
				VersionFn: func() int { return 7 },
			}
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		nodes := client.nodes()
		sort.Strings(nodes)
		if len(nodes) != 4 ||
			nodes[0] != "127.0.0.1:0" ||
			nodes[1] != "127.0.1.1:1" ||
			nodes[2] != "127.0.2.1:2" ||
			nodes[3] != "127.0.3.1:3" {
			t.Fatalf("unexpected nodes %v", nodes)
		}
	})

	t.Run("Refresh aws cluster", func(t *testing.T) {
		getClient := func(version int) (client *clusterClient, err error) {
			return newClusterClient(&ClientOption{InitAddress: []string{"xxxxx.amazonaws.com:1"}}, func(dst string, opt *ClientOption) conn {
				return &mockConn{
					DoFn: func(cmd Completed) RedisResult {
						if dst == "xxxxx.amazonaws.com:1" && strings.Join(cmd.Commands(), " ") == "CLUSTER SHARDS" {
							return shardsResp
						}
						return newErrResult(errors.New("unexpected call"))
					},
					AddrFn:    func() string { return "xxxxx.amazonaws.com:1" },
					VersionFn: func() int { return version },
				}
			})
		}

		t.Run("shards", func(t *testing.T) {
			client, err := getClient(7)
			if err != nil {
				t.Fatalf("unexpected err %v", err)
			}
			nodes := client.nodes()
			sort.Strings(nodes)
			if len(nodes) != 3 ||
				nodes[0] != "127.0.0.1:0" ||
				nodes[1] != "127.0.1.1:1" ||
				nodes[2] != "xxxxx.amazonaws.com:1" {
				t.Fatalf("unexpected nodes %v", nodes)
			}
			client.Close()
		})
	})

	t.Run("Refresh cluster which has only primary node per shard with SendToReplica option", func(t *testing.T) {
		m := &mockConn{
			DoFn: func(cmd Completed) RedisResult {
				if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
					return slotsMultiRespWithoutReplicas
				}
				return RedisResult{}
			},
		}

		client, err := newClusterClient(
			&ClientOption{
				InitAddress: []string{"127.0.0.1:0"},
				SendToReplicas: func(cmd Completed) bool {
					return true
				},
			},
			func(dst string, opt *ClientOption) conn {
				copiedM := *m
				return &copiedM
			},
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}

		if client.pslots[0] != client.conns["127.0.0.1:0"].conn {
			t.Fatalf("unexpected node assigned to pslot 0")
		}
		if client.pslots[8192] != client.conns["127.0.0.1:0"].conn {
			t.Fatalf("unexpected node assigned to pslot 8192")
		}
		if client.pslots[8193] != client.conns["127.0.1.1:0"].conn {
			t.Fatalf("unexpected node assigned to pslot 8193")
		}
		if client.pslots[16383] != client.conns["127.0.1.1:0"].conn {
			t.Fatalf("unexpected node assigned to pslot 16383")
		}
		if client.rslots[0] != client.conns["127.0.0.1:0"].conn {
			t.Fatalf("unexpected node assigned to rslot 0")
		}
		if client.rslots[8192] != client.conns["127.0.0.1:0"].conn {
			t.Fatalf("unexpected node assigned to rslot 8192")
		}
		if client.rslots[8193] != client.conns["127.0.1.1:0"].conn {
			t.Fatalf("unexpected node assigned to rslot 8193")
		}
		if client.rslots[16383] != client.conns["127.0.1.1:0"].conn {
			t.Fatalf("unexpected node assigned to rslot 16383")
		}
	})

	t.Run("Refresh cluster which has multi nodes per shard with SendToReplica option", func(t *testing.T) {
		m := &mockConn{
			DoFn: func(cmd Completed) RedisResult {
				if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
					return slotsMultiResp
				}
				return RedisResult{}
			},
		}

		client, err := newClusterClient(
			&ClientOption{
				InitAddress: []string{"127.0.0.1:0"},
				SendToReplicas: func(cmd Completed) bool {
					return true
				},
			},
			func(dst string, opt *ClientOption) conn {
				copiedM := *m
				return &copiedM
			},
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}

		if client.pslots[0] != client.conns["127.0.0.1:0"].conn {
			t.Fatalf("unexpected node assigned to pslot 0")
		}
		if client.pslots[8192] != client.conns["127.0.0.1:0"].conn {
			t.Fatalf("unexpected node assigned to pslot 8192")
		}
		if client.pslots[8193] != client.conns["127.0.2.1:0"].conn {
			t.Fatalf("unexpected node assigned to pslot 8193")
		}
		if client.pslots[16383] != client.conns["127.0.2.1:0"].conn {
			t.Fatalf("unexpected node assigned to pslot 16383")
		}
		if client.rslots[0] != client.conns["127.0.1.1:1"].conn {
			t.Fatalf("unexpected node assigned to rslot 0")
		}
		if client.rslots[8192] != client.conns["127.0.1.1:1"].conn {
			t.Fatalf("unexpected node assigned to rslot 8192")
		}
		if client.rslots[8193] != client.conns["127.0.3.1:1"].conn {
			t.Fatalf("unexpected node assigned to rslot 8193")
		}
		if client.rslots[16383] != client.conns["127.0.3.1:1"].conn {
			t.Fatalf("unexpected node assigned to rslot 16383")
		}
	})
}

//gocyclo:ignore
func TestClusterClient(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	m := &mockConn{
		DoFn: func(cmd Completed) RedisResult {
			if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
				return slotsMultiResp
			}
			return RedisResult{}
		},
		DoMultiFn: func(multi ...Completed) *redisresults {
			resps := make([]RedisResult, len(multi))
			for i, cmd := range multi {
				resps[i] = newResult(RedisMessage{typ: '+', string: strings.Join(cmd.Commands(), " ")}, nil)
			}
			return &redisresults{s: resps}
		},
		DoMultiCacheFn: func(multi ...CacheableTTL) *redisresults {
			resps := make([]RedisResult, len(multi))
			for i, cmd := range multi {
				resps[i] = newResult(RedisMessage{typ: '+', string: strings.Join(cmd.Cmd.Commands(), " ")}, nil)
			}
			return &redisresults{s: resps}
		},
		DoOverride: map[string]func(cmd Completed) RedisResult{
			"GET Do": func(cmd Completed) RedisResult {
				return newResult(RedisMessage{typ: '+', string: "Do"}, nil)
			},
			"INFO": func(cmd Completed) RedisResult {
				return newResult(RedisMessage{typ: '+', string: "Info"}, nil)
			},
		},
		DoCacheOverride: map[string]func(cmd Cacheable, ttl time.Duration) RedisResult{
			"GET DoCache": func(cmd Cacheable, ttl time.Duration) RedisResult {
				return newResult(RedisMessage{typ: '+', string: "DoCache"}, nil)
			},
		},
	}

	client, err := newClusterClient(&ClientOption{InitAddress: []string{"127.0.0.1:0"}}, func(dst string, opt *ClientOption) conn {
		return m
	})
	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}

	t.Run("Nodes", func(t *testing.T) {
		nodes := client.Nodes()
		if len(nodes) != 4 || nodes["127.0.0.1:0"] == nil || nodes["127.0.1.1:1"] == nil ||
			nodes["127.0.2.1:0"] == nil || nodes["127.0.3.1:1"] == nil {
			t.Fatalf("unexpected Nodes")
		}
	})

	t.Run("Delegate Do with no slot", func(t *testing.T) {
		c := client.B().Info().Build()
		if v, err := client.Do(context.Background(), c).ToString(); err != nil || v != "Info" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("Delegate Do", func(t *testing.T) {
		c := client.B().Get().Key("Do").Build()
		if v, err := client.Do(context.Background(), c).ToString(); err != nil || v != "Do" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("Delegate DoMulti Empty", func(t *testing.T) {
		if resps := client.DoMulti(context.Background()); resps != nil {
			t.Fatalf("unexpected response %v", resps)
		}
	})

	t.Run("Delegate DoMulti Single Slot", func(t *testing.T) {
		c1 := client.B().Get().Key("K1{a}").Build()
		c2 := client.B().Get().Key("K2{a}").Build()
		resps := client.DoMulti(context.Background(), c1, c2)
		if v, err := resps[0].ToString(); err != nil || v != "GET K1{a}" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
		if v, err := resps[1].ToString(); err != nil || v != "GET K2{a}" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("Delegate DoMulti Single Slot + Init Slot", func(t *testing.T) {
		c1 := client.B().Get().Key("K1{a}").Build()
		c2 := client.B().Info().Build()
		resps := client.DoMulti(context.Background(), c1, c2)
		if v, err := resps[0].ToString(); err != nil || v != "GET K1{a}" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
		if v, err := resps[1].ToString(); err != nil || v != "INFO" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("Delegate DoMulti Cross Slot + Init Slot", func(t *testing.T) {
		defer func() {
			if err := recover(); err != panicMixCxSlot {
				t.Errorf("DoMulti should panic if Cross Slot + Init Slot")
			}
		}()
		c1 := client.B().Get().Key("K1{a}").Build()
		c2 := client.B().Get().Key("K1{b}").Build()
		c3 := client.B().Info().Build()
		client.DoMulti(context.Background(), c1, c2, c3)
	})

	t.Run("Delegate DoMulti Multi Slot", func(t *testing.T) {
		multi := make([]Completed, 500)
		for i := 0; i < len(multi); i++ {
			multi[i] = client.B().Get().Key(fmt.Sprintf("K1{%d}", i)).Build()
		}
		resps := client.DoMulti(context.Background(), multi...)
		for i := 0; i < len(multi); i++ {
			if v, err := resps[i].ToString(); err != nil || v != fmt.Sprintf("GET K1{%d}", i) {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		}
	})

	t.Run("Delegate DoMulti Multi Slot", func(t *testing.T) {
		multi := make([]Completed, 500)
		for i := 0; i < len(multi); i++ {
			multi[i] = client.B().Get().Key(fmt.Sprintf("K1{%d}", i)).Build()
		}
		resps := client.DoMulti(context.Background(), multi...)
		for i := 0; i < len(multi); i++ {
			if v, err := resps[i].ToString(); err != nil || v != fmt.Sprintf("GET K1{%d}", i) {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		}
	})

	t.Run("Delegate DoCache", func(t *testing.T) {
		c := client.B().Get().Key("DoCache").Cache()
		if v, err := client.DoCache(context.Background(), c, 100).ToString(); err != nil || v != "DoCache" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("Delegate DoMultiCache Empty", func(t *testing.T) {
		if resps := client.DoMultiCache(context.Background()); resps != nil {
			t.Fatalf("unexpected response %v", resps)
		}
	})

	t.Run("Delegate DoMultiCache Single Slot", func(t *testing.T) {
		c1 := client.B().Get().Key("K1{a}").Cache()
		c2 := client.B().Get().Key("K2{a}").Cache()
		resps := client.DoMultiCache(context.Background(), CT(c1, time.Second), CT(c2, time.Second))
		if v, err := resps[0].ToString(); err != nil || v != "GET K1{a}" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
		if v, err := resps[1].ToString(); err != nil || v != "GET K2{a}" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("Delegate DoMultiCache Multi Slot", func(t *testing.T) {
		multi := make([]CacheableTTL, 500)
		for i := 0; i < len(multi); i++ {
			multi[i] = CT(client.B().Get().Key(fmt.Sprintf("K1{%d}", i)).Cache(), time.Second)
		}
		resps := client.DoMultiCache(context.Background(), multi...)
		for i := 0; i < len(multi); i++ {
			if v, err := resps[i].ToString(); err != nil || v != fmt.Sprintf("GET K1{%d}", i) {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		}
	})

	t.Run("Delegate Receive", func(t *testing.T) {
		c := client.B().Subscribe().Channel("ch").Build()
		hdl := func(message PubSubMessage) {}
		m.ReceiveFn = func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
			if !reflect.DeepEqual(subscribe.Commands(), c.Commands()) {
				t.Fatalf("unexpected command %v", subscribe)
			}
			return nil
		}
		if err := client.Receive(context.Background(), c, hdl); err != nil {
			t.Fatalf("unexpected response %v", err)
		}
	})

	t.Run("Delegate Receive Redis Err", func(t *testing.T) {
		c := client.B().Subscribe().Channel("ch").Build()
		e := &RedisError{}
		m.ReceiveFn = func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
			return e
		}
		if err := client.Receive(context.Background(), c, func(message PubSubMessage) {}); err != e {
			t.Fatalf("unexpected response %v", err)
		}
	})

	t.Run("Delegate Close", func(t *testing.T) {
		once := sync.Once{}
		called := make(chan struct{})
		m.CloseFn = func() {
			once.Do(func() { close(called) })
		}
		client.Close()
		<-called
	})

	t.Run("Dedicated Err", func(t *testing.T) {
		v := errors.New("fn err")
		if err := client.Dedicated(func(client DedicatedClient) error {
			return v
		}); err != v {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("Dedicated Cross Slot Err", func(t *testing.T) {
		defer func() {
			if err := recover(); err != panicMsgCxSlot {
				t.Errorf("Dedicated should panic if cross slots is used")
			}
		}()
		m.AcquireFn = func() wire { return &mockWire{} }
		client.Dedicated(func(c DedicatedClient) error {
			c.Do(context.Background(), c.B().Get().Key("a").Build()).Error()
			return c.Do(context.Background(), c.B().Get().Key("b").Build()).Error()
		})
	})

	t.Run("Dedicated Cross Slot Err Multi", func(t *testing.T) {
		defer func() {
			if err := recover(); err != panicMsgCxSlot {
				t.Errorf("Dedicated should panic if cross slots is used")
			}
		}()
		m.AcquireFn = func() wire {
			return &mockWire{
				DoMultiFn: func(multi ...Completed) *redisresults {
					return &redisresults{s: []RedisResult{
						newResult(RedisMessage{typ: '+', string: "OK"}, nil),
						newResult(RedisMessage{typ: '+', string: "OK"}, nil),
						newResult(RedisMessage{typ: '*', values: []RedisMessage{{typ: '+', string: "a"}}}, nil),
					}}
				},
			}
		}
		client.Dedicated(func(c DedicatedClient) (err error) {
			c.DoMulti(
				context.Background(),
				c.B().Multi().Build(),
				c.B().Get().Key("a").Build(),
				c.B().Exec().Build(),
			)
			c.DoMulti(
				context.Background(),
				c.B().Multi().Build(),
				c.B().Get().Key("b").Build(),
				c.B().Exec().Build(),
			)
			return nil
		})
	})

	t.Run("Dedicated Multi Cross Slot Err", func(t *testing.T) {
		m.AcquireFn = func() wire { return &mockWire{} }
		err := client.Dedicated(func(c DedicatedClient) (err error) {
			defer func() {
				err = errors.New(recover().(string))
			}()
			c.DoMulti(
				context.Background(),
				c.B().Get().Key("a").Build(),
				c.B().Get().Key("b").Build(),
			)
			return nil
		})
		if err == nil || err.Error() != panicMsgCxSlot {
			t.Errorf("Multi should panic if cross slots is used")
		}
	})

	t.Run("Dedicated Delegate Receive Redis Err", func(t *testing.T) {
		e := &RedisError{}
		w := &mockWire{
			ReceiveFn: func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
				return e
			},
		}
		m.AcquireFn = func() wire {
			return w
		}
		if err := client.Dedicated(func(c DedicatedClient) error {
			return c.Receive(context.Background(), c.B().Subscribe().Channel("a").Build(), func(msg PubSubMessage) {})
		}); err != e {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("Dedicated Delegate", func(t *testing.T) {
		closed := false
		w := &mockWire{
			DoFn: func(cmd Completed) RedisResult {
				return newResult(RedisMessage{typ: '+', string: "Delegate"}, nil)
			},
			DoMultiFn: func(cmd ...Completed) *redisresults {
				if len(cmd) == 4 {
					return &redisresults{s: []RedisResult{
						newResult(RedisMessage{typ: '+', string: "OK"}, nil),
						newResult(RedisMessage{typ: '+', string: "OK"}, nil),
						newResult(RedisMessage{typ: '+', string: "OK"}, nil),
						newResult(RedisMessage{typ: '*', values: []RedisMessage{
							{typ: '+', string: "Delegate0"},
							{typ: '+', string: "Delegate1"},
						}}, nil),
					}}
				}
				return &redisresults{s: []RedisResult{
					newResult(RedisMessage{typ: '+', string: "Delegate0"}, nil),
					newResult(RedisMessage{typ: '+', string: "Delegate1"}, nil),
				}}
			},
			ReceiveFn: func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
				return ErrClosing
			},
			SetPubSubHooksFn: func(hooks PubSubHooks) <-chan error {
				ch := make(chan error, 1)
				ch <- ErrClosing
				close(ch)
				return ch
			},
			ErrorFn: func() error {
				return ErrClosing
			},
			CloseFn: func() {
				closed = true
			},
		}
		m.AcquireFn = func() wire {
			return w
		}
		stored := false
		m.StoreFn = func(ww wire) {
			if ww != w {
				t.Fatalf("received unexpected wire %v", ww)
			}
			stored = true
		}
		if err := client.Dedicated(func(c DedicatedClient) error {
			ch := c.SetPubSubHooks(PubSubHooks{OnMessage: func(m PubSubMessage) {}})
			if v, err := c.Do(context.Background(), c.B().Get().Key("a").Build()).ToString(); err != nil || v != "Delegate" {
				t.Fatalf("unexpected response %v %v", v, err)
			}
			if v := c.DoMulti(context.Background()); len(v) != 0 {
				t.Fatalf("received unexpected response %v", v)
			}
			for i, resp := range c.DoMulti(
				context.Background(),
				c.B().Info().Build(),
				c.B().Info().Build(),
			) {
				if v, err := resp.ToString(); err != nil || v != "Delegate"+strconv.Itoa(i) {
					t.Fatalf("unexpected response %v %v", v, err)
				}
			}
			for i, resp := range c.DoMulti(
				context.Background(),
				c.B().Get().Key("a").Build(),
				c.B().Get().Key("a").Build(),
			) {
				if v, err := resp.ToString(); err != nil || v != "Delegate"+strconv.Itoa(i) {
					t.Fatalf("unexpected response %v %v", v, err)
				}
			}
			for i, resp := range c.DoMulti(
				context.Background(),
				c.B().Multi().Build(),
				c.B().Get().Key("a").Build(),
				c.B().Get().Key("a").Build(),
				c.B().Exec().Build(),
			)[3].val.values {
				if v, err := resp.ToString(); err != nil || v != "Delegate"+strconv.Itoa(i) {
					t.Fatalf("unexpected response %v %v", v, err)
				}
			}
			if err := c.Receive(context.Background(), c.B().Ssubscribe().Channel("a").Build(), func(msg PubSubMessage) {}); err != ErrClosing {
				t.Fatalf("unexpected ret %v", err)
			}
			if err := <-ch; err != ErrClosing {
				t.Fatalf("unexpected ret %v", err)
			}
			if err := <-c.SetPubSubHooks(PubSubHooks{OnMessage: func(m PubSubMessage) {}}); err != ErrClosing {
				t.Fatalf("unexpected ret %v", err)
			}
			c.Close()
			return nil
		}); err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if !stored {
			t.Fatalf("Dedicated desn't put back the wire")
		}
		if !closed {
			t.Fatalf("Dedicated desn't delegate Close")
		}
	})

	t.Run("Dedicated Delegate", func(t *testing.T) {
		closed := false
		w := &mockWire{
			DoFn: func(cmd Completed) RedisResult {
				return newResult(RedisMessage{typ: '+', string: "Delegate"}, nil)
			},
			DoMultiFn: func(cmd ...Completed) *redisresults {
				if len(cmd) == 4 {
					return &redisresults{s: []RedisResult{
						newResult(RedisMessage{typ: '+', string: "OK"}, nil),
						newResult(RedisMessage{typ: '+', string: "OK"}, nil),
						newResult(RedisMessage{typ: '+', string: "OK"}, nil),
						newResult(RedisMessage{typ: '*', values: []RedisMessage{
							{typ: '+', string: "Delegate0"},
							{typ: '+', string: "Delegate1"},
						}}, nil),
					}}
				}
				return &redisresults{s: []RedisResult{
					newResult(RedisMessage{typ: '+', string: "Delegate0"}, nil),
					newResult(RedisMessage{typ: '+', string: "Delegate1"}, nil),
				}}
			},
			ReceiveFn: func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
				return ErrClosing
			},
			SetPubSubHooksFn: func(hooks PubSubHooks) <-chan error {
				ch := make(chan error, 1)
				ch <- ErrClosing
				close(ch)
				return ch
			},
			ErrorFn: func() error {
				return ErrClosing
			},
			CloseFn: func() {
				closed = true
			},
		}
		m.AcquireFn = func() wire {
			return w
		}
		stored := false
		m.StoreFn = func(ww wire) {
			if ww != w {
				t.Fatalf("received unexpected wire %v", ww)
			}
			stored = true
		}
		c, cancel := client.Dedicate()
		ch := c.SetPubSubHooks(PubSubHooks{OnMessage: func(m PubSubMessage) {}})
		if v, err := c.Do(context.Background(), c.B().Get().Key("a").Build()).ToString(); err != nil || v != "Delegate" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
		if v := c.DoMulti(context.Background()); len(v) != 0 {
			t.Fatalf("received unexpected response %v", v)
		}
		for i, resp := range c.DoMulti(
			context.Background(),
			c.B().Info().Build(),
			c.B().Info().Build(),
		) {
			if v, err := resp.ToString(); err != nil || v != "Delegate"+strconv.Itoa(i) {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		}
		for i, resp := range c.DoMulti(
			context.Background(),
			c.B().Get().Key("a").Build(),
			c.B().Get().Key("a").Build(),
		) {
			if v, err := resp.ToString(); err != nil || v != "Delegate"+strconv.Itoa(i) {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		}
		for i, resp := range c.DoMulti(
			context.Background(),
			c.B().Multi().Build(),
			c.B().Get().Key("a").Build(),
			c.B().Get().Key("a").Build(),
			c.B().Exec().Build(),
		)[3].val.values {
			if v, err := resp.ToString(); err != nil || v != "Delegate"+strconv.Itoa(i) {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		}
		if err := c.Receive(context.Background(), c.B().Ssubscribe().Channel("a").Build(), func(msg PubSubMessage) {}); err != ErrClosing {
			t.Fatalf("unexpected ret %v", err)
		}
		if err := <-ch; err != ErrClosing {
			t.Fatalf("unexpected ret %v", err)
		}
		if err := <-c.SetPubSubHooks(PubSubHooks{OnMessage: func(m PubSubMessage) {}}); err != ErrClosing {
			t.Fatalf("unexpected ret %v", err)
		}
		c.Close()
		cancel()

		if !stored {
			t.Fatalf("Dedicated desn't put back the wire")
		}
		if !closed {
			t.Fatalf("Dedicated desn't delegate Close")
		}
	})

	t.Run("Dedicate Delegate Release On Close", func(t *testing.T) {
		stored := 0
		w := &mockWire{}
		m.AcquireFn = func() wire { return w }
		m.StoreFn = func(ww wire) { stored++ }
		c, _ := client.Dedicate()
		c.Do(context.Background(), c.B().Get().Key("a").Build())

		c.Close()

		if stored != 1 {
			t.Fatalf("unexpected stored count %v", stored)
		}
	})

	t.Run("Dedicate Delegate No Duplicate Release", func(t *testing.T) {
		stored := 0
		w := &mockWire{}
		m.AcquireFn = func() wire { return w }
		m.StoreFn = func(ww wire) { stored++ }
		c, cancel := client.Dedicate()
		c.Do(context.Background(), c.B().Get().Key("a").Build())

		c.Close()
		c.Close() // should have no effect
		cancel()  // should have no effect
		cancel()  // should have no effect

		if stored != 1 {
			t.Fatalf("unexpected stored count %v", stored)
		}
	})

	t.Run("Dedicated SetPubSubHooks Released", func(t *testing.T) {
		c, cancel := client.Dedicate()
		ch1 := c.SetPubSubHooks(PubSubHooks{OnMessage: func(m PubSubMessage) {}})
		ch2 := c.SetPubSubHooks(PubSubHooks{OnMessage: func(m PubSubMessage) {}})
		<-ch1
		cancel()
		<-ch2
	})

	t.Run("Dedicated SetPubSubHooks Close", func(t *testing.T) {
		c, cancel := client.Dedicate()
		defer cancel()
		ch := c.SetPubSubHooks(PubSubHooks{OnMessage: func(m PubSubMessage) {}})
		c.Close()
		if err := <-ch; err != ErrClosing {
			t.Fatalf("unexpected ret %v", ch)
		}
	})

	t.Run("Dedicated SetPubSubHooks Released", func(t *testing.T) {
		c, cancel := client.Dedicate()
		defer cancel()
		if ch := c.SetPubSubHooks(PubSubHooks{}); ch != nil {
			t.Fatalf("unexpected ret %v", ch)
		}
	})

	t.Run("Dedicate panic after released", func(t *testing.T) {
		check := func() {
			if err := recover(); err != dedicatedClientUsedAfterReleased {
				t.Fatalf("unexpected err %v", err)
			}
		}
		for _, closeFn := range []func(client DedicatedClient, cancel func()){
			func(client DedicatedClient, cancel func()) {
				client.Close()
			},
			func(client DedicatedClient, cancel func()) {
				cancel()
			},
		} {
			c, cancel := client.Dedicate()
			closeFn(c, cancel)
			for _, fn := range []func(){
				func() {
					defer check()
					c.Do(context.Background(), c.B().Get().Key("k").Build())
				},
				func() {
					defer check()
					c.DoMulti(context.Background(), c.B().Get().Key("k").Build())
				},
				func() {
					defer check()
					c.Receive(context.Background(), c.B().Subscribe().Channel("k").Build(), func(msg PubSubMessage) {})
				},
				func() {
					defer check()
					c.SetPubSubHooks(PubSubHooks{})
				},
			} {
				fn()
			}
		}
	})
}

//gocyclo:ignore
func TestClusterClient_SendToOnlyPrimaryNodes(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())

	primaryNodeConn := &mockConn{
		DoOverride: map[string]func(cmd Completed) RedisResult{
			"CLUSTER SLOTS": func(cmd Completed) RedisResult {
				return slotsMultiResp
			},
			"GET Do": func(cmd Completed) RedisResult {
				return newResult(RedisMessage{typ: '+', string: "GET Do"}, nil)
			},
			"GET K1{a}": func(cmd Completed) RedisResult {
				return newResult(RedisMessage{typ: '+', string: "GET K1{a}"}, nil)
			},
			"GET K2{a}": func(cmd Completed) RedisResult {
				return newResult(RedisMessage{typ: '+', string: "GET K2{a}"}, nil)
			},
			"INFO": func(cmd Completed) RedisResult {
				return newResult(RedisMessage{typ: '+', string: "INFO"}, nil)
			},
		},
		DoMultiFn: func(multi ...Completed) *redisresults {
			resps := make([]RedisResult, len(multi))
			for i, cmd := range multi {
				resps[i] = newResult(RedisMessage{typ: '+', string: strings.Join(cmd.Commands(), " ")}, nil)
			}
			return &redisresults{s: resps}
		},
		DoCacheOverride: map[string]func(cmd Cacheable, ttl time.Duration) RedisResult{
			"GET DoCache": func(cmd Cacheable, ttl time.Duration) RedisResult {
				return newResult(RedisMessage{typ: '+', string: "GET DoCache"}, nil)
			},
			"GET K1{a}": func(cmd Cacheable, ttl time.Duration) RedisResult {
				return newResult(RedisMessage{typ: '+', string: "GET K1{a}"}, nil)
			},
			"GET K2{a}": func(cmd Cacheable, ttl time.Duration) RedisResult {
				return newResult(RedisMessage{typ: '+', string: "GET K2{a}"}, nil)
			},
		},
		DoMultiCacheFn: func(multi ...CacheableTTL) *redisresults {
			resps := make([]RedisResult, len(multi))
			for i, cmd := range multi {
				resps[i] = newResult(RedisMessage{typ: '+', string: strings.Join(cmd.Cmd.Commands(), " ")}, nil)
			}
			return &redisresults{s: resps}
		},
	}
	replicaNodeConn := &mockConn{}

	client, err := newClusterClient(
		&ClientOption{
			InitAddress: []string{"127.0.0.1:0"},
			SendToReplicas: func(cmd Completed) bool {
				return false
			},
		},
		func(dst string, opt *ClientOption) conn {
			if dst == "127.0.0.1:0" || dst == "127.0.2.1:0" { // primary nodes
				return primaryNodeConn
			} else { // replica nodes
				return replicaNodeConn
			}
		},
	)
	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}

	t.Run("Do with no slot", func(t *testing.T) {
		c := client.B().Info().Build()
		if v, err := client.Do(context.Background(), c).ToString(); err != nil || v != "INFO" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("Do", func(t *testing.T) {
		c := client.B().Get().Key("Do").Build()
		if v, err := client.Do(context.Background(), c).ToString(); err != nil || v != "GET Do" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("DoMulti Single Slot", func(t *testing.T) {
		c1 := client.B().Get().Key("K1{a}").Build()
		c2 := client.B().Get().Key("K2{a}").Build()
		resps := client.DoMulti(context.Background(), c1, c2)
		if v, err := resps[0].ToString(); err != nil || v != "GET K1{a}" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
		if v, err := resps[1].ToString(); err != nil || v != "GET K2{a}" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("DoMulti Single Slot + Init Slot", func(t *testing.T) {
		c1 := client.B().Get().Key("K1{a}").Build()
		c2 := client.B().Info().Build()
		resps := client.DoMulti(context.Background(), c1, c2)
		if v, err := resps[0].ToString(); err != nil || v != "GET K1{a}" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
		if v, err := resps[1].ToString(); err != nil || v != "INFO" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("DoMulti Cross Slot + Init Slot", func(t *testing.T) {
		defer func() {
			if err := recover(); err != panicMixCxSlot {
				t.Errorf("DoMulti should panic if Cross Slot + Init Slot")
			}
		}()
		c1 := client.B().Get().Key("K1{a}").Build()
		c2 := client.B().Get().Key("K1{b}").Build()
		c3 := client.B().Info().Build()
		client.DoMulti(context.Background(), c1, c2, c3)
	})

	t.Run("DoMulti Multi Slot", func(t *testing.T) {
		multi := make([]Completed, 500)
		for i := 0; i < len(multi); i++ {
			multi[i] = client.B().Get().Key(fmt.Sprintf("K1{%d}", i)).Build()
		}
		resps := client.DoMulti(context.Background(), multi...)
		for i := 0; i < len(multi); i++ {
			if v, err := resps[i].ToString(); err != nil || v != fmt.Sprintf("GET K1{%d}", i) {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		}
	})

	t.Run("DoCache", func(t *testing.T) {
		c := client.B().Get().Key("DoCache").Cache()
		if v, err := client.DoCache(context.Background(), c, 100).ToString(); err != nil || v != "GET DoCache" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("DoMultiCache Single Slot", func(t *testing.T) {
		c1 := client.B().Get().Key("K1{a}").Cache()
		c2 := client.B().Get().Key("K2{a}").Cache()
		resps := client.DoMultiCache(context.Background(), CT(c1, time.Second), CT(c2, time.Second))
		if v, err := resps[0].ToString(); err != nil || v != "GET K1{a}" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
		if v, err := resps[1].ToString(); err != nil || v != "GET K2{a}" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("DoMultiCache Multi Slot", func(t *testing.T) {
		multi := make([]CacheableTTL, 500)
		for i := 0; i < len(multi); i++ {
			multi[i] = CT(client.B().Get().Key(fmt.Sprintf("K1{%d}", i)).Cache(), time.Second)
		}
		resps := client.DoMultiCache(context.Background(), multi...)
		for i := 0; i < len(multi); i++ {
			if v, err := resps[i].ToString(); err != nil || v != fmt.Sprintf("GET K1{%d}", i) {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		}
	})

	t.Run("Receive", func(t *testing.T) {
		c := client.B().Subscribe().Channel("ch").Build()
		hdl := func(message PubSubMessage) {}
		primaryNodeConn.ReceiveFn = func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
			if !reflect.DeepEqual(subscribe.Commands(), c.Commands()) {
				t.Fatalf("unexpected command %v", subscribe)
			}
			return nil
		}
		replicaNodeConn.ReceiveFn = func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
			if !reflect.DeepEqual(subscribe.Commands(), c.Commands()) {
				t.Fatalf("unexpected command %v", subscribe)
			}
			return nil
		}

		if err := client.Receive(context.Background(), c, hdl); err != nil {
			t.Fatalf("unexpected response %v", err)
		}
	})

	t.Run("Receive Redis Err", func(t *testing.T) {
		c := client.B().Subscribe().Channel("ch").Build()
		e := &RedisError{}
		primaryNodeConn.ReceiveFn = func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
			return e
		}
		replicaNodeConn.ReceiveFn = func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
			return e
		}
		if err := client.Receive(context.Background(), c, func(message PubSubMessage) {}); err != e {
			t.Fatalf("unexpected response %v", err)
		}
	})

	t.Run("Dedicated Err", func(t *testing.T) {
		v := errors.New("fn err")
		if err := client.Dedicated(func(client DedicatedClient) error {
			return v
		}); err != v {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("Dedicated Cross Slot Err", func(t *testing.T) {
		defer func() {
			if err := recover(); err != panicMsgCxSlot {
				t.Errorf("Dedicated should panic if cross slots is used")
			}
		}()
		primaryNodeConn.AcquireFn = func() wire { return &mockWire{} }
		client.Dedicated(func(c DedicatedClient) error {
			c.Do(context.Background(), c.B().Get().Key("a").Build()).Error()
			return c.Do(context.Background(), c.B().Get().Key("b").Build()).Error()
		})
	})

	t.Run("Dedicated Cross Slot Err Multi", func(t *testing.T) {
		defer func() {
			if err := recover(); err != panicMsgCxSlot {
				t.Errorf("Dedicated should panic if cross slots is used")
			}
		}()
		primaryNodeConn.AcquireFn = func() wire {
			return &mockWire{
				DoMultiFn: func(multi ...Completed) *redisresults {
					return &redisresults{s: []RedisResult{
						newResult(RedisMessage{typ: '+', string: "OK"}, nil),
						newResult(RedisMessage{typ: '+', string: "OK"}, nil),
						newResult(RedisMessage{typ: '*', values: []RedisMessage{{typ: '+', string: "a"}}}, nil),
					}}
				},
			}
		}
		client.Dedicated(func(c DedicatedClient) (err error) {
			c.DoMulti(
				context.Background(),
				c.B().Multi().Build(),
				c.B().Get().Key("a").Build(),
				c.B().Exec().Build(),
			)
			c.DoMulti(
				context.Background(),
				c.B().Multi().Build(),
				c.B().Get().Key("b").Build(),
				c.B().Exec().Build(),
			)
			return nil
		})
	})

	t.Run("Dedicated Multi Cross Slot Err", func(t *testing.T) {
		primaryNodeConn.AcquireFn = func() wire { return &mockWire{} }
		err := client.Dedicated(func(c DedicatedClient) (err error) {
			defer func() {
				err = errors.New(recover().(string))
			}()
			c.DoMulti(
				context.Background(),
				c.B().Get().Key("a").Build(),
				c.B().Get().Key("b").Build(),
			)
			return nil
		})
		if err == nil || err.Error() != panicMsgCxSlot {
			t.Errorf("Multi should panic if cross slots is used")
		}
	})

	t.Run("Dedicated Receive Redis Err", func(t *testing.T) {
		e := &RedisError{}
		w := &mockWire{
			ReceiveFn: func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
				return e
			},
		}
		primaryNodeConn.AcquireFn = func() wire {
			return w
		}
		if err := client.Dedicated(func(c DedicatedClient) error {
			return c.Receive(context.Background(), c.B().Ssubscribe().Channel("a").Build(), func(msg PubSubMessage) {})
		}); err != e {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("Dedicated", func(t *testing.T) {
		closed := false
		w := &mockWire{
			DoFn: func(cmd Completed) RedisResult {
				return newResult(RedisMessage{typ: '+', string: "Delegate"}, nil)
			},
			DoMultiFn: func(cmd ...Completed) *redisresults {
				if len(cmd) == 4 {
					return &redisresults{s: []RedisResult{
						newResult(RedisMessage{typ: '+', string: "OK"}, nil),
						newResult(RedisMessage{typ: '+', string: "OK"}, nil),
						newResult(RedisMessage{typ: '+', string: "OK"}, nil),
						newResult(RedisMessage{typ: '*', values: []RedisMessage{
							{typ: '+', string: "Delegate0"},
							{typ: '+', string: "Delegate1"},
						}}, nil),
					}}
				}
				return &redisresults{s: []RedisResult{
					newResult(RedisMessage{typ: '+', string: "Delegate0"}, nil),
					newResult(RedisMessage{typ: '+', string: "Delegate1"}, nil),
				}}
			},
			ReceiveFn: func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
				return ErrClosing
			},
			SetPubSubHooksFn: func(hooks PubSubHooks) <-chan error {
				ch := make(chan error, 1)
				ch <- ErrClosing
				close(ch)
				return ch
			},
			ErrorFn: func() error {
				return ErrClosing
			},
			CloseFn: func() {
				closed = true
			},
		}
		primaryNodeConn.AcquireFn = func() wire {
			return w
		}
		stored := false
		primaryNodeConn.StoreFn = func(ww wire) {
			if ww != w {
				t.Fatalf("received unexpected wire %v", ww)
			}
			stored = true
		}
		if err := client.Dedicated(func(c DedicatedClient) error {
			ch := c.SetPubSubHooks(PubSubHooks{OnMessage: func(m PubSubMessage) {}})
			if v, err := c.Do(context.Background(), c.B().Get().Key("a").Build()).ToString(); err != nil || v != "Delegate" {
				t.Fatalf("unexpected response %v %v", v, err)
			}
			if v := c.DoMulti(context.Background()); len(v) != 0 {
				t.Fatalf("received unexpected response %v", v)
			}
			for i, resp := range c.DoMulti(
				context.Background(),
				c.B().Info().Build(),
				c.B().Info().Build(),
			) {
				if v, err := resp.ToString(); err != nil || v != "Delegate"+strconv.Itoa(i) {
					t.Fatalf("unexpected response %v %v", v, err)
				}
			}
			for i, resp := range c.DoMulti(
				context.Background(),
				c.B().Get().Key("a").Build(),
				c.B().Get().Key("a").Build(),
			) {
				if v, err := resp.ToString(); err != nil || v != "Delegate"+strconv.Itoa(i) {
					t.Fatalf("unexpected response %v %v", v, err)
				}
			}
			for i, resp := range c.DoMulti(
				context.Background(),
				c.B().Multi().Build(),
				c.B().Get().Key("a").Build(),
				c.B().Get().Key("a").Build(),
				c.B().Exec().Build(),
			)[3].val.values {
				if v, err := resp.ToString(); err != nil || v != "Delegate"+strconv.Itoa(i) {
					t.Fatalf("unexpected response %v %v", v, err)
				}
			}
			if err := c.Receive(context.Background(), c.B().Ssubscribe().Channel("a").Build(), func(msg PubSubMessage) {}); err != ErrClosing {
				t.Fatalf("unexpected ret %v", err)
			}
			if err := <-ch; err != ErrClosing {
				t.Fatalf("unexpected ret %v", err)
			}
			if err := <-c.SetPubSubHooks(PubSubHooks{OnMessage: func(m PubSubMessage) {}}); err != ErrClosing {
				t.Fatalf("unexpected ret %v", err)
			}
			c.Close()
			return nil
		}); err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if !stored {
			t.Fatalf("Dedicated desn't put back the wire")
		}
		if !closed {
			t.Fatalf("Dedicated desn't delegate Close")
		}
	})

	t.Run("Dedicate", func(t *testing.T) {
		closed := false
		w := &mockWire{
			DoFn: func(cmd Completed) RedisResult {
				return newResult(RedisMessage{typ: '+', string: "Delegate"}, nil)
			},
			DoMultiFn: func(cmd ...Completed) *redisresults {
				if len(cmd) == 4 {
					return &redisresults{s: []RedisResult{
						newResult(RedisMessage{typ: '+', string: "OK"}, nil),
						newResult(RedisMessage{typ: '+', string: "OK"}, nil),
						newResult(RedisMessage{typ: '+', string: "OK"}, nil),
						newResult(RedisMessage{typ: '*', values: []RedisMessage{
							{typ: '+', string: "Delegate0"},
							{typ: '+', string: "Delegate1"},
						}}, nil),
					}}
				}
				return &redisresults{s: []RedisResult{
					newResult(RedisMessage{typ: '+', string: "Delegate0"}, nil),
					newResult(RedisMessage{typ: '+', string: "Delegate1"}, nil),
				}}
			},
			ReceiveFn: func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
				return ErrClosing
			},
			SetPubSubHooksFn: func(hooks PubSubHooks) <-chan error {
				ch := make(chan error, 1)
				ch <- ErrClosing
				close(ch)
				return ch
			},
			ErrorFn: func() error {
				return ErrClosing
			},
			CloseFn: func() {
				closed = true
			},
		}
		primaryNodeConn.AcquireFn = func() wire {
			return w
		}
		stored := false
		primaryNodeConn.StoreFn = func(ww wire) {
			if ww != w {
				t.Fatalf("received unexpected wire %v", ww)
			}
			stored = true
		}
		c, cancel := client.Dedicate()
		ch := c.SetPubSubHooks(PubSubHooks{OnMessage: func(m PubSubMessage) {}})
		if v, err := c.Do(context.Background(), c.B().Get().Key("a").Build()).ToString(); err != nil || v != "Delegate" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
		if v := c.DoMulti(context.Background()); len(v) != 0 {
			t.Fatalf("received unexpected response %v", v)
		}
		for i, resp := range c.DoMulti(
			context.Background(),
			c.B().Info().Build(),
			c.B().Info().Build(),
		) {
			if v, err := resp.ToString(); err != nil || v != "Delegate"+strconv.Itoa(i) {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		}
		for i, resp := range c.DoMulti(
			context.Background(),
			c.B().Get().Key("a").Build(),
			c.B().Get().Key("a").Build(),
		) {
			if v, err := resp.ToString(); err != nil || v != "Delegate"+strconv.Itoa(i) {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		}
		for i, resp := range c.DoMulti(
			context.Background(),
			c.B().Multi().Build(),
			c.B().Get().Key("a").Build(),
			c.B().Get().Key("a").Build(),
			c.B().Exec().Build(),
		)[3].val.values {
			if v, err := resp.ToString(); err != nil || v != "Delegate"+strconv.Itoa(i) {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		}
		if err := c.Receive(context.Background(), c.B().Ssubscribe().Channel("a").Build(), func(msg PubSubMessage) {}); err != ErrClosing {
			t.Fatalf("unexpected ret %v", err)
		}
		if err := <-ch; err != ErrClosing {
			t.Fatalf("unexpected ret %v", err)
		}
		if err := <-c.SetPubSubHooks(PubSubHooks{OnMessage: func(m PubSubMessage) {}}); err != ErrClosing {
			t.Fatalf("unexpected ret %v", err)
		}
		c.Close()
		cancel()

		if !stored {
			t.Fatalf("Dedicated desn't put back the wire")
		}
		if !closed {
			t.Fatalf("Dedicated desn't delegate Close")
		}
	})
}

//gocyclo:ignore
func TestClusterClient_SendToOnlyReplicaNodes(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())

	primaryNodeConn := &mockConn{
		DoOverride: map[string]func(cmd Completed) RedisResult{
			"CLUSTER SLOTS": func(cmd Completed) RedisResult {
				return slotsMultiResp
			},
			"INFO": func(cmd Completed) RedisResult {
				return newResult(RedisMessage{typ: '+', string: "INFO"}, nil)
			},
			"GET K1{a}": func(cmd Completed) RedisResult {
				return newResult(RedisMessage{typ: '+', string: "GET K1{a}"}, nil)
			},
		},
	}
	replicaNodeConn := &mockConn{
		DoMultiFn: func(multi ...Completed) *redisresults {
			resps := make([]RedisResult, len(multi))
			for i, cmd := range multi {
				resps[i] = newResult(RedisMessage{typ: '+', string: strings.Join(cmd.Commands(), " ")}, nil)
			}
			return &redisresults{s: resps}
		},
		DoMultiCacheFn: func(multi ...CacheableTTL) *redisresults {
			resps := make([]RedisResult, len(multi))
			for i, cmd := range multi {
				resps[i] = newResult(RedisMessage{typ: '+', string: strings.Join(cmd.Cmd.Commands(), " ")}, nil)
			}
			return &redisresults{s: resps}
		},
		DoOverride: map[string]func(cmd Completed) RedisResult{
			"GET Do": func(cmd Completed) RedisResult {
				return newResult(RedisMessage{typ: '+', string: "GET Do"}, nil)
			},
			"GET K1{a}": func(cmd Completed) RedisResult {
				return newResult(RedisMessage{typ: '+', string: "GET K1{a}"}, nil)
			},
			"GET K2{a}": func(cmd Completed) RedisResult {
				return newResult(RedisMessage{typ: '+', string: "GET K2{a}"}, nil)
			},
		},
		DoCacheOverride: map[string]func(cmd Cacheable, ttl time.Duration) RedisResult{
			"GET DoCache": func(cmd Cacheable, ttl time.Duration) RedisResult {
				return newResult(RedisMessage{typ: '+', string: "GET DoCache"}, nil)
			},
		},
	}

	client, err := newClusterClient(
		&ClientOption{
			InitAddress: []string{"127.0.0.1:0"},
			SendToReplicas: func(cmd Completed) bool {
				return true
			},
		},
		func(dst string, opt *ClientOption) conn {
			if dst == "127.0.0.1:0" || dst == "127.0.2.1:0" { // primary nodes
				return primaryNodeConn
			} else { // replica nodes
				return replicaNodeConn
			}
		},
	)
	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}

	t.Run("Do with no slot", func(t *testing.T) {
		c := client.B().Info().Build()
		if v, err := client.Do(context.Background(), c).ToString(); err != nil || v != "INFO" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("Do", func(t *testing.T) {
		c := client.B().Get().Key("Do").Build()
		if v, err := client.Do(context.Background(), c).ToString(); err != nil || v != "GET Do" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("DoMulti Single Slot", func(t *testing.T) {
		c1 := client.B().Get().Key("K1{a}").Build()
		c2 := client.B().Get().Key("K2{a}").Build()
		resps := client.DoMulti(context.Background(), c1, c2)
		if v, err := resps[0].ToString(); err != nil || v != "GET K1{a}" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
		if v, err := resps[1].ToString(); err != nil || v != "GET K2{a}" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("DoMulti Single Slot + Init Slot", func(t *testing.T) {
		c1 := client.B().Get().Key("K1{a}").Build()
		c2 := client.B().Info().Build()
		resps := client.DoMulti(context.Background(), c1, c2)
		if v, err := resps[0].ToString(); err != nil || v != "GET K1{a}" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
		if v, err := resps[1].ToString(); err != nil || v != "INFO" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("DoMulti Cross Slot + Init Slot", func(t *testing.T) {
		defer func() {
			if err := recover(); err != panicMixCxSlot {
				t.Errorf("DoMulti should panic if Cross Slot + Init Slot")
			}
		}()
		c1 := client.B().Get().Key("K1{a}").Build()
		c2 := client.B().Get().Key("K1{b}").Build()
		c3 := client.B().Info().Build()
		client.DoMulti(context.Background(), c1, c2, c3)
	})

	t.Run("DoMulti Multi Slot", func(t *testing.T) {
		multi := make([]Completed, 500)
		for i := 0; i < len(multi); i++ {
			multi[i] = client.B().Get().Key(fmt.Sprintf("K1{%d}", i)).Build()
		}
		resps := client.DoMulti(context.Background(), multi...)
		for i := 0; i < len(multi); i++ {
			if v, err := resps[i].ToString(); err != nil || v != fmt.Sprintf("GET K1{%d}", i) {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		}
	})

	t.Run("DoCache", func(t *testing.T) {
		c := client.B().Get().Key("DoCache").Cache()
		if v, err := client.DoCache(context.Background(), c, 100).ToString(); err != nil || v != "GET DoCache" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("DoMultiCache Single Slot", func(t *testing.T) {
		c1 := client.B().Get().Key("K1{a}").Cache()
		c2 := client.B().Get().Key("K2{a}").Cache()
		resps := client.DoMultiCache(context.Background(), CT(c1, time.Second), CT(c2, time.Second))
		if v, err := resps[0].ToString(); err != nil || v != "GET K1{a}" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
		if v, err := resps[1].ToString(); err != nil || v != "GET K2{a}" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("DoMultiCache Multi Slot", func(t *testing.T) {
		multi := make([]CacheableTTL, 500)
		for i := 0; i < len(multi); i++ {
			multi[i] = CT(client.B().Get().Key(fmt.Sprintf("K1{%d}", i)).Cache(), time.Second)
		}
		resps := client.DoMultiCache(context.Background(), multi...)
		for i := 0; i < len(multi); i++ {
			if v, err := resps[i].ToString(); err != nil || v != fmt.Sprintf("GET K1{%d}", i) {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		}
	})

	t.Run("Receive", func(t *testing.T) {
		c := client.B().Subscribe().Channel("ch").Build()
		hdl := func(message PubSubMessage) {}
		primaryNodeConn.ReceiveFn = func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
			if !reflect.DeepEqual(subscribe.Commands(), c.Commands()) {
				t.Fatalf("unexpected command %v", subscribe)
			}
			return nil
		}
		replicaNodeConn.ReceiveFn = func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
			if !reflect.DeepEqual(subscribe.Commands(), c.Commands()) {
				t.Fatalf("unexpected command %v", subscribe)
			}
			return nil
		}
		if err := client.Receive(context.Background(), c, hdl); err != nil {
			t.Fatalf("unexpected response %v", err)
		}
	})

	t.Run("Receive Redis Err", func(t *testing.T) {
		c := client.B().Ssubscribe().Channel("ch").Build()
		e := &RedisError{}
		primaryNodeConn.ReceiveFn = func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
			return e
		}
		replicaNodeConn.ReceiveFn = func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
			return e
		}
		if err := client.Receive(context.Background(), c, func(message PubSubMessage) {}); err != e {
			t.Fatalf("unexpected response %v", err)
		}
	})

	t.Run("Dedicated Cross Slot Err", func(t *testing.T) {
		defer func() {
			if err := recover(); err != panicMsgCxSlot {
				t.Errorf("Dedicated should panic if cross slots is used")
			}
		}()
		primaryNodeConn.AcquireFn = func() wire { return &mockWire{} }
		client.Dedicated(func(c DedicatedClient) error {
			c.Do(context.Background(), c.B().Get().Key("a").Build()).Error()
			return c.Do(context.Background(), c.B().Get().Key("b").Build()).Error()
		})
	})

	t.Run("Dedicated Cross Slot Err Multi", func(t *testing.T) {
		defer func() {
			if err := recover(); err != panicMsgCxSlot {
				t.Errorf("Dedicated should panic if cross slots is used")
			}
		}()
		primaryNodeConn.AcquireFn = func() wire {
			return &mockWire{
				DoMultiFn: func(multi ...Completed) *redisresults {
					return &redisresults{s: []RedisResult{
						newResult(RedisMessage{typ: '+', string: "OK"}, nil),
						newResult(RedisMessage{typ: '+', string: "OK"}, nil),
						newResult(RedisMessage{typ: '*', values: []RedisMessage{{typ: '+', string: "a"}}}, nil),
					}}
				},
			}
		}
		client.Dedicated(func(c DedicatedClient) (err error) {
			c.DoMulti(
				context.Background(),
				c.B().Multi().Build(),
				c.B().Get().Key("a").Build(),
				c.B().Exec().Build(),
			)
			c.DoMulti(
				context.Background(),
				c.B().Multi().Build(),
				c.B().Get().Key("b").Build(),
				c.B().Exec().Build(),
			)
			return nil
		})
	})

	t.Run("Dedicated Multi Cross Slot Err", func(t *testing.T) {
		primaryNodeConn.AcquireFn = func() wire { return &mockWire{} }
		err := client.Dedicated(func(c DedicatedClient) (err error) {
			defer func() {
				err = errors.New(recover().(string))
			}()
			c.DoMulti(
				context.Background(),
				c.B().Get().Key("a").Build(),
				c.B().Get().Key("b").Build(),
			)
			return nil
		})
		if err == nil || err.Error() != panicMsgCxSlot {
			t.Errorf("Multi should panic if cross slots is used")
		}
	})

	t.Run("Dedicated Receive Redis Err", func(t *testing.T) {
		e := &RedisError{}
		w := &mockWire{
			ReceiveFn: func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
				return e
			},
		}
		primaryNodeConn.AcquireFn = func() wire {
			return w
		}
		if err := client.Dedicated(func(c DedicatedClient) error {
			return c.Receive(context.Background(), c.B().Subscribe().Channel("a").Build(), func(msg PubSubMessage) {})
		}); err != e {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("Dedicated", func(t *testing.T) {
		closed := false
		w := &mockWire{
			DoFn: func(cmd Completed) RedisResult {
				return newResult(RedisMessage{typ: '+', string: "Delegate"}, nil)
			},
			DoMultiFn: func(cmd ...Completed) *redisresults {
				if len(cmd) == 4 {
					return &redisresults{s: []RedisResult{
						newResult(RedisMessage{typ: '+', string: "OK"}, nil),
						newResult(RedisMessage{typ: '+', string: "OK"}, nil),
						newResult(RedisMessage{typ: '+', string: "OK"}, nil),
						newResult(RedisMessage{typ: '*', values: []RedisMessage{
							{typ: '+', string: "Delegate0"},
							{typ: '+', string: "Delegate1"},
						}}, nil),
					}}
				}
				return &redisresults{s: []RedisResult{
					newResult(RedisMessage{typ: '+', string: "Delegate0"}, nil),
					newResult(RedisMessage{typ: '+', string: "Delegate1"}, nil),
				}}
			},
			ReceiveFn: func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
				return ErrClosing
			},
			SetPubSubHooksFn: func(hooks PubSubHooks) <-chan error {
				ch := make(chan error, 1)
				ch <- ErrClosing
				close(ch)
				return ch
			},
			ErrorFn: func() error {
				return ErrClosing
			},
			CloseFn: func() {
				closed = true
			},
		}
		primaryNodeConn.AcquireFn = func() wire {
			return w
		}
		stored := false
		primaryNodeConn.StoreFn = func(ww wire) {
			if ww != w {
				t.Fatalf("received unexpected wire %v", ww)
			}
			stored = true
		}
		if err := client.Dedicated(func(c DedicatedClient) error {
			ch := c.SetPubSubHooks(PubSubHooks{OnMessage: func(m PubSubMessage) {}})
			if v, err := c.Do(context.Background(), c.B().Get().Key("a").Build()).ToString(); err != nil || v != "Delegate" {
				t.Fatalf("unexpected response %v %v", v, err)
			}
			if v := c.DoMulti(context.Background()); len(v) != 0 {
				t.Fatalf("received unexpected response %v", v)
			}
			for i, resp := range c.DoMulti(
				context.Background(),
				c.B().Info().Build(),
				c.B().Info().Build(),
			) {
				if v, err := resp.ToString(); err != nil || v != "Delegate"+strconv.Itoa(i) {
					t.Fatalf("unexpected response %v %v", v, err)
				}
			}
			for i, resp := range c.DoMulti(
				context.Background(),
				c.B().Get().Key("a").Build(),
				c.B().Get().Key("a").Build(),
			) {
				if v, err := resp.ToString(); err != nil || v != "Delegate"+strconv.Itoa(i) {
					t.Fatalf("unexpected response %v %v", v, err)
				}
			}
			for i, resp := range c.DoMulti(
				context.Background(),
				c.B().Multi().Build(),
				c.B().Get().Key("a").Build(),
				c.B().Get().Key("a").Build(),
				c.B().Exec().Build(),
			)[3].val.values {
				if v, err := resp.ToString(); err != nil || v != "Delegate"+strconv.Itoa(i) {
					t.Fatalf("unexpected response %v %v", v, err)
				}
			}
			if err := c.Receive(context.Background(), c.B().Ssubscribe().Channel("a").Build(), func(msg PubSubMessage) {}); err != ErrClosing {
				t.Fatalf("unexpected ret %v", err)
			}
			if err := <-ch; err != ErrClosing {
				t.Fatalf("unexpected ret %v", err)
			}
			if err := <-c.SetPubSubHooks(PubSubHooks{OnMessage: func(m PubSubMessage) {}}); err != ErrClosing {
				t.Fatalf("unexpected ret %v", err)
			}
			c.Close()
			return nil
		}); err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if !stored {
			t.Fatalf("Dedicated desn't put back the wire")
		}
		if !closed {
			t.Fatalf("Dedicated desn't delegate Close")
		}
	})

	t.Run("Dedicate", func(t *testing.T) {
		closed := false
		w := &mockWire{
			DoFn: func(cmd Completed) RedisResult {
				return newResult(RedisMessage{typ: '+', string: "Delegate"}, nil)
			},
			DoMultiFn: func(cmd ...Completed) *redisresults {
				if len(cmd) == 4 {
					return &redisresults{s: []RedisResult{
						newResult(RedisMessage{typ: '+', string: "OK"}, nil),
						newResult(RedisMessage{typ: '+', string: "OK"}, nil),
						newResult(RedisMessage{typ: '+', string: "OK"}, nil),
						newResult(RedisMessage{typ: '*', values: []RedisMessage{
							{typ: '+', string: "Delegate0"},
							{typ: '+', string: "Delegate1"},
						}}, nil),
					}}
				}
				return &redisresults{s: []RedisResult{
					newResult(RedisMessage{typ: '+', string: "Delegate0"}, nil),
					newResult(RedisMessage{typ: '+', string: "Delegate1"}, nil),
				}}
			},
			ReceiveFn: func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
				return ErrClosing
			},
			SetPubSubHooksFn: func(hooks PubSubHooks) <-chan error {
				ch := make(chan error, 1)
				ch <- ErrClosing
				close(ch)
				return ch
			},
			ErrorFn: func() error {
				return ErrClosing
			},
			CloseFn: func() {
				closed = true
			},
		}
		primaryNodeConn.AcquireFn = func() wire {
			return w
		}
		stored := false
		primaryNodeConn.StoreFn = func(ww wire) {
			if ww != w {
				t.Fatalf("received unexpected wire %v", ww)
			}
			stored = true
		}
		c, cancel := client.Dedicate()
		ch := c.SetPubSubHooks(PubSubHooks{OnMessage: func(m PubSubMessage) {}})
		if v, err := c.Do(context.Background(), c.B().Get().Key("a").Build()).ToString(); err != nil || v != "Delegate" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
		if v := c.DoMulti(context.Background()); len(v) != 0 {
			t.Fatalf("received unexpected response %v", v)
		}
		for i, resp := range c.DoMulti(
			context.Background(),
			c.B().Info().Build(),
			c.B().Info().Build(),
		) {
			if v, err := resp.ToString(); err != nil || v != "Delegate"+strconv.Itoa(i) {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		}
		for i, resp := range c.DoMulti(
			context.Background(),
			c.B().Get().Key("a").Build(),
			c.B().Get().Key("a").Build(),
		) {
			if v, err := resp.ToString(); err != nil || v != "Delegate"+strconv.Itoa(i) {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		}
		for i, resp := range c.DoMulti(
			context.Background(),
			c.B().Multi().Build(),
			c.B().Get().Key("a").Build(),
			c.B().Get().Key("a").Build(),
			c.B().Exec().Build(),
		)[3].val.values {
			if v, err := resp.ToString(); err != nil || v != "Delegate"+strconv.Itoa(i) {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		}
		if err := c.Receive(context.Background(), c.B().Ssubscribe().Channel("a").Build(), func(msg PubSubMessage) {}); err != ErrClosing {
			t.Fatalf("unexpected ret %v", err)
		}
		if err := <-ch; err != ErrClosing {
			t.Fatalf("unexpected ret %v", err)
		}
		if err := <-c.SetPubSubHooks(PubSubHooks{OnMessage: func(m PubSubMessage) {}}); err != ErrClosing {
			t.Fatalf("unexpected ret %v", err)
		}
		c.Close()
		cancel()

		if !stored {
			t.Fatalf("Dedicated desn't put back the wire")
		}
		if !closed {
			t.Fatalf("Dedicated desn't delegate Close")
		}
	})
}

//gocyclo:ignore
func TestClusterClient_SendReadOperationToReplicaNodesWriteOperationToPrimaryNodes(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())

	primaryNodeConn := &mockConn{
		DoOverride: map[string]func(cmd Completed) RedisResult{
			"CLUSTER SLOTS": func(cmd Completed) RedisResult {
				return slotsMultiResp
			},
			"INFO": func(cmd Completed) RedisResult {
				return newResult(RedisMessage{typ: '+', string: "INFO"}, nil)
			},
			"SET Do V": func(cmd Completed) RedisResult {
				return newResult(RedisMessage{typ: '+', string: "SET Do V"}, nil)
			},
			"SET K2{a} V2{a}": func(cmd Completed) RedisResult {
				return newResult(RedisMessage{typ: '+', string: "SET K2{a} V2{a}"}, nil)
			},
		},
		DoMultiFn: func(multi ...Completed) *redisresults {
			resps := make([]RedisResult, len(multi))
			for i, cmd := range multi {
				if strings.HasPrefix(strings.Join(cmd.Commands(), " "), "SET K1") {
					resps[i] = newResult(RedisMessage{typ: '+', string: strings.Join(cmd.Commands(), " ")}, nil)
					continue
				}
				if strings.HasPrefix(strings.Join(cmd.Commands(), " "), "SET K2") {
					resps[i] = newResult(RedisMessage{typ: '+', string: strings.Join(cmd.Commands(), " ")}, nil)
					continue
				}
				if strings.HasPrefix(strings.Join(cmd.Commands(), " "), "MULTI") {
					resps[i] = newResult(RedisMessage{typ: '+', string: "MULTI"}, nil)
					continue
				}
				if strings.HasPrefix(strings.Join(cmd.Commands(), " "), "EXEC") {
					resps[i] = newResult(RedisMessage{typ: '+', string: "EXEC"}, nil)
					continue
				}

				return &redisresults{
					s: []RedisResult{},
				}
			}
			return &redisresults{s: resps}
		},
	}
	replicaNodeConn := &mockConn{
		DoOverride: map[string]func(cmd Completed) RedisResult{
			"GET Do": func(cmd Completed) RedisResult {
				return newResult(RedisMessage{typ: '+', string: "GET Do"}, nil)
			},
			"GET K1{a}": func(cmd Completed) RedisResult {
				return newResult(RedisMessage{typ: '+', string: "GET K1{a}"}, nil)
			},
			"GET K2{a}": func(cmd Completed) RedisResult {
				return newResult(RedisMessage{typ: '+', string: "GET K2{a}"}, nil)
			},
		},
		DoMultiFn: func(multi ...Completed) *redisresults {
			resps := make([]RedisResult, len(multi))
			for i, cmd := range multi {
				if strings.HasPrefix(strings.Join(cmd.Commands(), " "), "GET K1") {
					resps[i] = newResult(RedisMessage{typ: '+', string: strings.Join(cmd.Commands(), " ")}, nil)
					continue
				}

				return &redisresults{
					s: []RedisResult{},
				}
			}
			return &redisresults{s: resps}
		},
		DoCacheOverride: map[string]func(cmd Cacheable, ttl time.Duration) RedisResult{
			"GET DoCache": func(cmd Cacheable, ttl time.Duration) RedisResult {
				return newResult(RedisMessage{typ: '+', string: "GET DoCache"}, nil)
			},
			"GET K1{a}": func(cmd Cacheable, ttl time.Duration) RedisResult {
				return newResult(RedisMessage{typ: '+', string: "GET K1{a}"}, nil)
			},
			"GET K2{a}": func(cmd Cacheable, ttl time.Duration) RedisResult {
				return newResult(RedisMessage{typ: '+', string: "GET K2{a}"}, nil)
			},
		},
		DoMultiCacheFn: func(multi ...CacheableTTL) *redisresults {
			resps := make([]RedisResult, len(multi))
			for i, cmd := range multi {
				if strings.HasPrefix(strings.Join(cmd.Cmd.Commands(), " "), "GET K1") {
					resps[i] = newResult(RedisMessage{typ: '+', string: strings.Join(cmd.Cmd.Commands(), " ")}, nil)
					continue
				}

				return &redisresults{
					s: []RedisResult{},
				}
			}
			return &redisresults{s: resps}
		},
	}

	client, err := newClusterClient(
		&ClientOption{
			InitAddress: []string{"127.0.0.1:0"},
			SendToReplicas: func(cmd Completed) bool {
				if cmd.IsReadOnly() {
					return true
				}
				return false
			},
		},
		func(dst string, opt *ClientOption) conn {
			if dst == "127.0.0.1:0" || dst == "127.0.2.1:0" { // primary nodes
				return primaryNodeConn
			} else { // replica nodes
				return replicaNodeConn
			}
		},
	)
	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}

	t.Run("Do with no slot", func(t *testing.T) {
		c := client.B().Info().Build()
		if v, err := client.Do(context.Background(), c).ToString(); err != nil || v != "INFO" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("Do read operation", func(t *testing.T) {
		c := client.B().Get().Key("Do").Build()
		if v, err := client.Do(context.Background(), c).ToString(); err != nil || v != "GET Do" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("Do write operation", func(t *testing.T) {
		c := client.B().Set().Key("Do").Value("V").Build()
		if v, err := client.Do(context.Background(), c).ToString(); err != nil || v != "SET Do V" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("DoMulti Single Slot All Read Operations", func(t *testing.T) {
		c1 := client.B().Get().Key("K1{a}").Build()
		c2 := client.B().Get().Key("K2{a}").Build()
		resps := client.DoMulti(context.Background(), c1, c2)
		if v, err := resps[0].ToString(); err != nil || v != "GET K1{a}" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
		if v, err := resps[1].ToString(); err != nil || v != "GET K2{a}" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("DoMulti Single Slot Read Operation And Write Operation", func(t *testing.T) {
		c1 := client.B().Get().Key("K1{a}").Build()
		c2 := client.B().Set().Key("K2{a}").Value("V2{a}").Build()
		resps := client.DoMulti(context.Background(), c1, c2)
		if v, err := resps[0].ToString(); err != nil || v != "GET K1{a}" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
		if v, err := resps[1].ToString(); err != nil || v != "SET K2{a} V2{a}" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("DoMulti Single Slot Operations + Init Slot", func(t *testing.T) {
		c1 := client.B().Multi().Build()
		c2 := client.B().Set().Key("K1{a}").Value("V1{a}").Build()
		c3 := client.B().Set().Key("K2{a}").Value("V2{a}").Build()
		c4 := client.B().Exec().Build()
		resps := client.DoMulti(context.Background(), c1, c2, c3, c4)
		if v, err := resps[0].ToString(); err != nil || v != "MULTI" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
		if v, err := resps[1].ToString(); err != nil || v != "SET K1{a} V1{a}" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
		if v, err := resps[2].ToString(); err != nil || v != "SET K2{a} V2{a}" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
		if v, err := resps[3].ToString(); err != nil || v != "EXEC" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("DoMulti Cross Slot + Init Slot", func(t *testing.T) {
		defer func() {
			if err := recover(); err != panicMixCxSlot {
				t.Errorf("DoMulti should panic if Cross Slot + Init Slot")
			}
		}()
		c1 := client.B().Get().Key("K1{a}").Build()
		c2 := client.B().Get().Key("K1{b}").Build()
		c3 := client.B().Info().Build()
		client.DoMulti(context.Background(), c1, c2, c3)
	})

	t.Run("DoMulti Multi Slot All Read Operations", func(t *testing.T) {
		multi := make([]Completed, 500)
		for i := 0; i < len(multi); i++ {
			multi[i] = client.B().Get().Key(fmt.Sprintf("K1{%d}", i)).Build()
		}
		resps := client.DoMulti(context.Background(), multi...)
		for i := 0; i < len(multi); i++ {
			if v, err := resps[i].ToString(); err != nil || v != fmt.Sprintf("GET K1{%d}", i) {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		}
	})
	t.Run("DoMulti Multi Slot Read & Write Operations", func(t *testing.T) {
		multi := make([]Completed, 500)
		for i := 0; i < len(multi); i++ {
			if i%2 == 0 {
				multi[i] = client.B().Get().Key(fmt.Sprintf("K1{%d}", i)).Build()
			} else {
				multi[i] = client.B().Set().Key(fmt.Sprintf("K2{%d}", i)).Value(fmt.Sprintf("V2{%d}", i)).Build()
			}
		}
		resps := client.DoMulti(context.Background(), multi...)
		for i := 0; i < len(multi); i++ {
			if i%2 == 0 {
				if v, err := resps[i].ToString(); err != nil || v != fmt.Sprintf("GET K1{%d}", i) {
					t.Fatalf("unexpected response %v %v", v, err)
				}
			} else {
				if v, err := resps[i].ToString(); err != nil || v != fmt.Sprintf("SET K2{%d} V2{%d}", i, i) {
					t.Fatalf("unexpected response %v %v", v, err)
				}
			}
		}
	})

	t.Run("DoCache Operation", func(t *testing.T) {
		c := client.B().Get().Key("DoCache").Cache()
		if v, err := client.DoCache(context.Background(), c, 100).ToString(); err != nil || v != "GET DoCache" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("DoMultiCache Single Slot", func(t *testing.T) {
		c1 := client.B().Get().Key("K1{a}").Cache()
		c2 := client.B().Get().Key("K2{a}").Cache()
		resps := client.DoMultiCache(context.Background(), CT(c1, time.Second), CT(c2, time.Second))
		if v, err := resps[0].ToString(); err != nil || v != "GET K1{a}" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
		if v, err := resps[1].ToString(); err != nil || v != "GET K2{a}" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("DoMultiCache Multi Slot", func(t *testing.T) {
		multi := make([]CacheableTTL, 500)
		for i := 0; i < len(multi); i++ {
			multi[i] = CT(client.B().Get().Key(fmt.Sprintf("K1{%d}", i)).Cache(), time.Second)
		}
		resps := client.DoMultiCache(context.Background(), multi...)
		for i := 0; i < len(multi); i++ {
			if v, err := resps[i].ToString(); err != nil || v != fmt.Sprintf("GET K1{%d}", i) {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		}
	})

	t.Run("Receive", func(t *testing.T) {
		c := client.B().Subscribe().Channel("ch").Build()
		hdl := func(message PubSubMessage) {}
		primaryNodeConn.ReceiveFn = func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
			if !reflect.DeepEqual(subscribe.Commands(), c.Commands()) {
				t.Fatalf("unexpected command %v", subscribe)
			}
			return nil
		}
		replicaNodeConn.ReceiveFn = func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
			if !reflect.DeepEqual(subscribe.Commands(), c.Commands()) {
				t.Fatalf("unexpected command %v", subscribe)
			}
			return nil
		}
		if err := client.Receive(context.Background(), c, hdl); err != nil {
			t.Fatalf("unexpected response %v", err)
		}
	})

	t.Run("Receive Redis Err", func(t *testing.T) {
		c := client.B().Ssubscribe().Channel("ch").Build()
		e := &RedisError{}
		primaryNodeConn.ReceiveFn = func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
			return e
		}
		replicaNodeConn.ReceiveFn = func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
			return e
		}
		if err := client.Receive(context.Background(), c, func(message PubSubMessage) {}); err != e {
			t.Fatalf("unexpected response %v", err)
		}
	})

	t.Run("Dedicated Cross Slot Err", func(t *testing.T) {
		defer func() {
			if err := recover(); err != panicMsgCxSlot {
				t.Errorf("Dedicated should panic if cross slots is used")
			}
		}()
		primaryNodeConn.AcquireFn = func() wire { return &mockWire{} }
		client.Dedicated(func(c DedicatedClient) error {
			c.Do(context.Background(), c.B().Get().Key("a").Build()).Error()
			return c.Do(context.Background(), c.B().Get().Key("b").Build()).Error()
		})
	})

	t.Run("Dedicated Cross Slot Err Multi", func(t *testing.T) {
		defer func() {
			if err := recover(); err != panicMsgCxSlot {
				t.Errorf("Dedicated should panic if cross slots is used")
			}
		}()
		primaryNodeConn.AcquireFn = func() wire {
			return &mockWire{
				DoMultiFn: func(multi ...Completed) *redisresults {
					return &redisresults{s: []RedisResult{
						newResult(RedisMessage{typ: '+', string: "OK"}, nil),
						newResult(RedisMessage{typ: '+', string: "OK"}, nil),
						newResult(RedisMessage{typ: '*', values: []RedisMessage{{typ: '+', string: "a"}}}, nil),
					}}
				},
			}
		}
		client.Dedicated(func(c DedicatedClient) (err error) {
			c.DoMulti(
				context.Background(),
				c.B().Multi().Build(),
				c.B().Get().Key("a").Build(),
				c.B().Exec().Build(),
			)
			c.DoMulti(
				context.Background(),
				c.B().Multi().Build(),
				c.B().Get().Key("b").Build(),
				c.B().Exec().Build(),
			)
			return nil
		})
	})

	t.Run("Dedicated Multi Cross Slot Err", func(t *testing.T) {
		primaryNodeConn.AcquireFn = func() wire { return &mockWire{} }
		err := client.Dedicated(func(c DedicatedClient) (err error) {
			defer func() {
				err = errors.New(recover().(string))
			}()
			c.DoMulti(
				context.Background(),
				c.B().Get().Key("a").Build(),
				c.B().Get().Key("b").Build(),
			)
			return nil
		})
		if err == nil || err.Error() != panicMsgCxSlot {
			t.Errorf("Multi should panic if cross slots is used")
		}
	})

	t.Run("Dedicated Receive Redis Err", func(t *testing.T) {
		e := &RedisError{}
		w := &mockWire{
			ReceiveFn: func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
				return e
			},
		}
		primaryNodeConn.AcquireFn = func() wire {
			return w
		}
		if err := client.Dedicated(func(c DedicatedClient) error {
			return c.Receive(context.Background(), c.B().Subscribe().Channel("a").Build(), func(msg PubSubMessage) {})
		}); err != e {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("Dedicated", func(t *testing.T) {
		closed := false
		w := &mockWire{
			DoFn: func(cmd Completed) RedisResult {
				return newResult(RedisMessage{typ: '+', string: "Delegate"}, nil)
			},
			DoMultiFn: func(cmd ...Completed) *redisresults {
				if len(cmd) == 4 {
					return &redisresults{s: []RedisResult{
						newResult(RedisMessage{typ: '+', string: "OK"}, nil),
						newResult(RedisMessage{typ: '+', string: "OK"}, nil),
						newResult(RedisMessage{typ: '+', string: "OK"}, nil),
						newResult(RedisMessage{typ: '*', values: []RedisMessage{
							{typ: '+', string: "Delegate0"},
							{typ: '+', string: "Delegate1"},
						}}, nil),
					}}
				}
				return &redisresults{s: []RedisResult{
					newResult(RedisMessage{typ: '+', string: "Delegate0"}, nil),
					newResult(RedisMessage{typ: '+', string: "Delegate1"}, nil),
				}}
			},
			ReceiveFn: func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
				return ErrClosing
			},
			SetPubSubHooksFn: func(hooks PubSubHooks) <-chan error {
				ch := make(chan error, 1)
				ch <- ErrClosing
				close(ch)
				return ch
			},
			ErrorFn: func() error {
				return ErrClosing
			},
			CloseFn: func() {
				closed = true
			},
		}
		primaryNodeConn.AcquireFn = func() wire {
			return w
		}
		stored := false
		primaryNodeConn.StoreFn = func(ww wire) {
			if ww != w {
				t.Fatalf("received unexpected wire %v", ww)
			}
			stored = true
		}
		if err := client.Dedicated(func(c DedicatedClient) error {
			ch := c.SetPubSubHooks(PubSubHooks{OnMessage: func(m PubSubMessage) {}})
			if v, err := c.Do(context.Background(), c.B().Get().Key("a").Build()).ToString(); err != nil || v != "Delegate" {
				t.Fatalf("unexpected response %v %v", v, err)
			}
			if v := c.DoMulti(context.Background()); len(v) != 0 {
				t.Fatalf("received unexpected response %v", v)
			}
			for i, resp := range c.DoMulti(
				context.Background(),
				c.B().Info().Build(),
				c.B().Info().Build(),
			) {
				if v, err := resp.ToString(); err != nil || v != "Delegate"+strconv.Itoa(i) {
					t.Fatalf("unexpected response %v %v", v, err)
				}
			}
			for i, resp := range c.DoMulti(
				context.Background(),
				c.B().Get().Key("a").Build(),
				c.B().Get().Key("a").Build(),
			) {
				if v, err := resp.ToString(); err != nil || v != "Delegate"+strconv.Itoa(i) {
					t.Fatalf("unexpected response %v %v", v, err)
				}
			}
			for i, resp := range c.DoMulti(
				context.Background(),
				c.B().Multi().Build(),
				c.B().Get().Key("a").Build(),
				c.B().Get().Key("a").Build(),
				c.B().Exec().Build(),
			)[3].val.values {
				if v, err := resp.ToString(); err != nil || v != "Delegate"+strconv.Itoa(i) {
					t.Fatalf("unexpected response %v %v", v, err)
				}
			}
			if err := c.Receive(context.Background(), c.B().Ssubscribe().Channel("a").Build(), func(msg PubSubMessage) {}); err != ErrClosing {
				t.Fatalf("unexpected ret %v", err)
			}
			if err := <-ch; err != ErrClosing {
				t.Fatalf("unexpected ret %v", err)
			}
			if err := <-c.SetPubSubHooks(PubSubHooks{OnMessage: func(m PubSubMessage) {}}); err != ErrClosing {
				t.Fatalf("unexpected ret %v", err)
			}
			c.Close()
			return nil
		}); err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if !stored {
			t.Fatalf("Dedicated desn't put back the wire")
		}
		if !closed {
			t.Fatalf("Dedicated desn't delegate Close")
		}
	})

	t.Run("Dedicate", func(t *testing.T) {
		closed := false
		w := &mockWire{
			DoFn: func(cmd Completed) RedisResult {
				return newResult(RedisMessage{typ: '+', string: "Delegate"}, nil)
			},
			DoMultiFn: func(cmd ...Completed) *redisresults {
				if len(cmd) == 4 {
					return &redisresults{s: []RedisResult{
						newResult(RedisMessage{typ: '+', string: "OK"}, nil),
						newResult(RedisMessage{typ: '+', string: "OK"}, nil),
						newResult(RedisMessage{typ: '+', string: "OK"}, nil),
						newResult(RedisMessage{typ: '*', values: []RedisMessage{
							{typ: '+', string: "Delegate0"},
							{typ: '+', string: "Delegate1"},
						}}, nil),
					}}
				}
				return &redisresults{s: []RedisResult{
					newResult(RedisMessage{typ: '+', string: "Delegate0"}, nil),
					newResult(RedisMessage{typ: '+', string: "Delegate1"}, nil),
				}}
			},
			ReceiveFn: func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
				return ErrClosing
			},
			SetPubSubHooksFn: func(hooks PubSubHooks) <-chan error {
				ch := make(chan error, 1)
				ch <- ErrClosing
				close(ch)
				return ch
			},
			ErrorFn: func() error {
				return ErrClosing
			},
			CloseFn: func() {
				closed = true
			},
		}
		primaryNodeConn.AcquireFn = func() wire {
			return w
		}
		stored := false
		primaryNodeConn.StoreFn = func(ww wire) {
			if ww != w {
				t.Fatalf("received unexpected wire %v", ww)
			}
			stored = true
		}
		c, cancel := client.Dedicate()
		ch := c.SetPubSubHooks(PubSubHooks{OnMessage: func(m PubSubMessage) {}})
		if v, err := c.Do(context.Background(), c.B().Get().Key("a").Build()).ToString(); err != nil || v != "Delegate" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
		if v := c.DoMulti(context.Background()); len(v) != 0 {
			t.Fatalf("received unexpected response %v", v)
		}
		for i, resp := range c.DoMulti(
			context.Background(),
			c.B().Info().Build(),
			c.B().Info().Build(),
		) {
			if v, err := resp.ToString(); err != nil || v != "Delegate"+strconv.Itoa(i) {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		}
		for i, resp := range c.DoMulti(
			context.Background(),
			c.B().Get().Key("a").Build(),
			c.B().Get().Key("a").Build(),
		) {
			if v, err := resp.ToString(); err != nil || v != "Delegate"+strconv.Itoa(i) {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		}
		for i, resp := range c.DoMulti(
			context.Background(),
			c.B().Multi().Build(),
			c.B().Get().Key("a").Build(),
			c.B().Get().Key("a").Build(),
			c.B().Exec().Build(),
		)[3].val.values {
			if v, err := resp.ToString(); err != nil || v != "Delegate"+strconv.Itoa(i) {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		}
		if err := c.Receive(context.Background(), c.B().Ssubscribe().Channel("a").Build(), func(msg PubSubMessage) {}); err != ErrClosing {
			t.Fatalf("unexpected ret %v", err)
		}
		if err := <-ch; err != ErrClosing {
			t.Fatalf("unexpected ret %v", err)
		}
		if err := <-c.SetPubSubHooks(PubSubHooks{OnMessage: func(m PubSubMessage) {}}); err != ErrClosing {
			t.Fatalf("unexpected ret %v", err)
		}
		c.Close()
		cancel()

		if !stored {
			t.Fatalf("Dedicated desn't put back the wire")
		}
		if !closed {
			t.Fatalf("Dedicated desn't delegate Close")
		}
	})
}

//gocyclo:ignore
func TestClusterClient_SendPrimaryNodeOnlyButOneSlotAssigned(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())

	primaryNodeConn := &mockConn{
		DoOverride: map[string]func(cmd Completed) RedisResult{
			"CLUSTER SLOTS": func(cmd Completed) RedisResult {
				return singleSlotResp
			},
			"INFO": func(cmd Completed) RedisResult {
				return newResult(RedisMessage{typ: '+', string: "INFO"}, nil)
			},
		},
		DoMultiFn: func(multi ...Completed) *redisresults {
			resps := make([]RedisResult, len(multi))
			for i, cmd := range multi {
				if strings.HasPrefix(strings.Join(cmd.Commands(), " "), "MULTI") {
					resps[i] = newResult(RedisMessage{typ: '+', string: "MULTI"}, nil)
					continue
				}
				if strings.HasPrefix(strings.Join(cmd.Commands(), " "), "EXEC") {
					resps[i] = newResult(RedisMessage{typ: '+', string: "EXEC"}, nil)
					continue
				}

				return &redisresults{
					s: []RedisResult{},
				}
			}
			return &redisresults{s: resps}
		},
	}

	client, err := newClusterClient(
		&ClientOption{
			InitAddress: []string{"127.0.0.1:0"},
			SendToReplicas: func(cmd Completed) bool {
				return false
			},
		},
		func(dst string, opt *ClientOption) conn {
			return primaryNodeConn
		},
	)
	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}

	t.Run("Do with no slot", func(t *testing.T) {
		c := client.B().Info().Build()
		if v, err := client.Do(context.Background(), c).ToString(); err != nil || v != "INFO" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("DoMulti Init Slot Operations", func(t *testing.T) {
		c1 := client.B().Multi().Build()
		c2 := client.B().Exec().Build()
		resps := client.DoMulti(context.Background(), c1, c2)
		if v, err := resps[0].ToString(); err != nil || v != "MULTI" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
		if v, err := resps[1].ToString(); err != nil || v != "EXEC" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})
}

//gocyclo:ignore
func TestClusterClientErr(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())

	t.Run("not refresh on context error", func(t *testing.T) {
		var count int64
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		v := ctx.Err()
		m := &mockConn{
			DoFn: func(cmd Completed) RedisResult {
				if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
					atomic.AddInt64(&count, 1)
					return slotsResp
				}
				return newErrResult(v)
			},
			DoMultiFn: func(multi ...Completed) *redisresults {
				res := make([]RedisResult, len(multi))
				for i := range res {
					res[i] = newErrResult(v)
				}
				return &redisresults{s: res}
			},
			DoCacheFn: func(cmd Cacheable, ttl time.Duration) RedisResult {
				return newErrResult(v)
			},
			DoMultiCacheFn: func(multi ...CacheableTTL) *redisresults {
				res := make([]RedisResult, len(multi))
				for i := range res {
					res[i] = newErrResult(v)
				}
				return &redisresults{s: res}
			},
			ReceiveFn: func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
				return v
			},
		}
		client, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return m
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if err := client.Do(ctx, client.B().Get().Key("a").Build()).Error(); err != v {
			t.Fatalf("unexpected err %v", err)
		}
		if err := client.DoMulti(ctx, client.B().Get().Key("a").Build())[0].Error(); err != v {
			t.Fatalf("unexpected err %v", err)
		}
		if err := client.DoCache(ctx, client.B().Get().Key("a").Cache(), 100).Error(); err != v {
			t.Fatalf("unexpected err %v", err)
		}
		if err := client.DoMultiCache(ctx, CT(client.B().Get().Key("a").Cache(), 100))[0].Error(); err != v {
			t.Fatalf("unexpected err %v", err)
		}
		if err := client.Receive(ctx, client.B().Ssubscribe().Channel("a").Build(), func(msg PubSubMessage) {}); err != v {
			t.Fatalf("unexpected err %v", err)
		}
		if c := atomic.LoadInt64(&count); c != 1 {
			t.Fatalf("unexpected refresh count %v", c)
		}
	})

	t.Run("refresh err on pick", func(t *testing.T) {
		var first int64
		v := errors.New("refresh err")
		m := &mockConn{
			DoFn: func(cmd Completed) RedisResult {
				if atomic.AddInt64(&first, 1) == 1 {
					return singleSlotResp
				}
				return newErrResult(v)
			},
			ReceiveFn: func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
				return v
			},
		}
		client, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return m
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if err := client.Do(context.Background(), client.B().Get().Key("a").Build()).Error(); err != v {
			t.Fatalf("unexpected err %v", err)
		}
		if err := client.DoMulti(context.Background(), client.B().Get().Key("a").Build())[0].Error(); err != v {
			t.Fatalf("unexpected err %v", err)
		}
		for _, resp := range client.DoMulti(context.Background(), client.B().Get().Key("a").Build(), client.B().Get().Key("b").Build()) {
			if err := resp.Error(); err != v {
				t.Fatalf("unexpected err %v", err)
			}
		}
		if err := client.DoCache(context.Background(), client.B().Get().Key("a").Cache(), 100).Error(); err != v {
			t.Fatalf("unexpected err %v", err)
		}
		if err := client.DoMultiCache(context.Background(), CT(client.B().Get().Key("a").Cache(), 100))[0].Error(); err != v {
			t.Fatalf("unexpected err %v", err)
		}
		for _, resp := range client.DoMultiCache(context.Background(), CT(client.B().Get().Key("a").Cache(), 100), CT(client.B().Get().Key("b").Cache(), 100)) {
			if err := resp.Error(); err != v {
				t.Fatalf("unexpected err %v", err)
			}
		}
		if err := client.Receive(context.Background(), client.B().Ssubscribe().Channel("a").Build(), func(msg PubSubMessage) {}); err != v {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("refresh empty on pick", func(t *testing.T) {
		m := &mockConn{DoFn: func(cmd Completed) RedisResult {
			return singleSlotResp
		}}
		client, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return m
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if err := client.Do(context.Background(), client.B().Get().Key("a").Build()).Error(); err != ErrNoSlot {
			t.Fatalf("unexpected err %v", err)
		}
		if err := client.DoMulti(context.Background(), client.B().Get().Key("a").Build())[0].Error(); err != ErrNoSlot {
			t.Fatalf("unexpected err %v", err)
		}
		for _, resp := range client.DoMulti(context.Background(), client.B().Get().Key("a").Build(), client.B().Get().Key("b").Build()) {
			if err := resp.Error(); err != ErrNoSlot {
				t.Fatalf("unexpected err %v", err)
			}
		}
		if err := client.DoMultiCache(context.Background(), CT(client.B().Get().Key("a").Cache(), 100))[0].Error(); err != ErrNoSlot {
			t.Fatalf("unexpected err %v", err)
		}
		for _, resp := range client.DoMultiCache(context.Background(), CT(client.B().Get().Key("a").Cache(), 100), CT(client.B().Get().Key("b").Cache(), 100)) {
			if err := resp.Error(); err != ErrNoSlot {
				t.Fatalf("unexpected err %v", err)
			}
		}
	})

	t.Run("refresh empty on pick in dedicated wire", func(t *testing.T) {
		m := &mockConn{DoFn: func(cmd Completed) RedisResult {
			return singleSlotResp
		}}
		client, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return m
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		var ch <-chan error
		if err := client.Dedicated(func(c DedicatedClient) error {
			ch = c.SetPubSubHooks(PubSubHooks{OnMessage: func(m PubSubMessage) {}})
			return c.Do(context.Background(), c.B().Get().Key("a").Build()).Error()
		}); err != ErrNoSlot {
			t.Fatalf("unexpected err %v", err)
		}
		if err := <-ch; err != ErrNoSlot {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("refresh empty on pick in dedicated wire (multi)", func(t *testing.T) {
		m := &mockConn{DoFn: func(cmd Completed) RedisResult {
			return singleSlotResp
		}}
		client, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return m
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		var ch <-chan error
		if err := client.Dedicated(func(c DedicatedClient) error {
			ch = c.SetPubSubHooks(PubSubHooks{OnMessage: func(m PubSubMessage) {}})
			for _, v := range c.DoMulti(context.Background(), c.B().Get().Key("a").Build()) {
				if err := v.Error(); err != nil {
					return err
				}
			}
			return nil
		}); err != ErrNoSlot {
			t.Fatalf("unexpected err %v", err)
		}
		if err := <-ch; err != ErrNoSlot {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("slot reconnect", func(t *testing.T) {
		var count, check int64
		client, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			atomic.AddInt64(&check, 1)
			return &mockConn{DoFn: func(cmd Completed) RedisResult {
				if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
					return slotsMultiResp
				}
				if atomic.AddInt64(&count, 1) <= 3 {
					return newResult(RedisMessage{typ: '-', string: "MOVED 0 :0"}, nil)
				}
				return newResult(RedisMessage{typ: '+', string: "b"}, nil)
			}}
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.Do(context.Background(), client.B().Get().Key("a").Build()).ToString(); err != nil || v != "b" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if atomic.LoadInt64(&check) != 6 {
			t.Fatalf("unexpected check count %v", check)
		}
	})

	t.Run("slot moved", func(t *testing.T) {
		var count int64
		client, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return &mockConn{DoFn: func(cmd Completed) RedisResult {
				if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
					return slotsMultiResp
				}
				if atomic.AddInt64(&count, 1) <= 3 {
					return newResult(RedisMessage{typ: '-', string: "MOVED 0 :1"}, nil)
				}
				return newResult(RedisMessage{typ: '+', string: "b"}, nil)
			}}
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.Do(context.Background(), client.B().Get().Key("a").Build()).ToString(); err != nil || v != "b" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot moved DoMulti (single)", func(t *testing.T) {
		var count int64
		client, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return &mockConn{DoFn: func(cmd Completed) RedisResult {
				return slotsMultiResp
			}, DoMultiFn: func(multi ...Completed) *redisresults {
				ret := make([]RedisResult, len(multi))
				if atomic.AddInt64(&count, 1) <= 3 {
					for i := range ret {
						ret[i] = newResult(RedisMessage{typ: '-', string: "MOVED 0 :1"}, nil)
					}
					return &redisresults{s: ret}
				}
				for i := range ret {
					ret[i] = newResult(RedisMessage{typ: '+', string: multi[i].Commands()[1]}, nil)
				}
				return &redisresults{s: ret}
			}}
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.DoMulti(context.Background(), client.B().Get().Key("a").Build())[0].ToString(); err != nil || v != "a" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot moved DoMulti (multi)", func(t *testing.T) {
		var count int64
		client, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return &mockConn{DoFn: func(cmd Completed) RedisResult {
				return slotsMultiResp
			}, DoMultiFn: func(multi ...Completed) *redisresults {
				ret := make([]RedisResult, len(multi))
				if atomic.AddInt64(&count, 1) <= 3 {
					for i := range ret {
						ret[i] = newResult(RedisMessage{typ: '-', string: "MOVED 0 :1"}, nil)
					}
					return &redisresults{s: ret}
				}
				for i := range ret {
					ret[i] = newResult(RedisMessage{typ: '+', string: multi[i].Commands()[1]}, nil)
				}
				return &redisresults{s: ret}
			}}
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		for i, resp := range client.DoMulti(context.Background(), client.B().Set().Key("a").Value("a").Build(), client.B().Get().Key("b").Build()) {
			if v, err := resp.ToString(); err != nil {
				t.Fatalf("unexpected resp %v %v", v, err)
			} else if i == 0 && v != "a" {
				t.Fatalf("unexpected resp %v %v", v, err)
			} else if i == 1 && v != "b" {
				t.Fatalf("unexpected resp %v %v", v, err)
			}
		}
	})

	t.Run("slot moved DoMulti (multi) TRYAGAIN", func(t *testing.T) {
		var count int64
		client, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return &mockConn{DoFn: func(cmd Completed) RedisResult {
				return slotsMultiResp
			}, DoMultiFn: func(multi ...Completed) *redisresults {
				ret := make([]RedisResult, len(multi))
				if atomic.AddInt64(&count, 1) <= 2 {
					for i := range ret {
						ret[i] = newResult(RedisMessage{typ: '-', string: "TRYAGAIN"}, nil)
					}
					return &redisresults{s: ret}
				}
				for i := range ret {
					ret[i] = newResult(RedisMessage{typ: '+', string: multi[i].Commands()[1]}, nil)
				}
				return &redisresults{s: ret}
			}}
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		for i, resp := range client.DoMulti(context.Background(), client.B().Set().Key("a").Value("a").Build(), client.B().Get().Key("b").Build()) {
			if v, err := resp.ToString(); err != nil && i != 0 && !strings.Contains(err.Error(), "TRYAGAIN") {
				t.Fatalf("unexpected resp %v %v", v, err)
			} else if i == 1 && v != "b" {
				t.Fatalf("unexpected resp %v %v", v, err)
			}
		}
	})

	t.Run("slot moved new", func(t *testing.T) {
		var count, check int64
		client, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			if dst == ":2" {
				atomic.AddInt64(&check, 1)
			}
			return &mockConn{DoFn: func(cmd Completed) RedisResult {
				if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
					return slotsMultiResp
				}
				if atomic.AddInt64(&count, 1) <= 3 {
					return newResult(RedisMessage{typ: '-', string: "MOVED 0 :2"}, nil)
				}
				return newResult(RedisMessage{typ: '+', string: "b"}, nil)
			}}

		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.Do(context.Background(), client.B().Get().Key("a").Build()).ToString(); err != nil || v != "b" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if atomic.LoadInt64(&check) == 0 {
			t.Fatalf("unexpected check value %v", check)
		}
	})

	t.Run("slot moved new (multi 1)", func(t *testing.T) {
		var count, check int64
		client, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			if dst == ":2" {
				atomic.AddInt64(&check, 1)
			}
			return &mockConn{DoFn: func(cmd Completed) RedisResult {
				return slotsMultiResp
			}, DoMultiFn: func(multi ...Completed) *redisresults {
				ret := make([]RedisResult, len(multi))
				if atomic.AddInt64(&count, 1) <= 3 {
					for i := range ret {
						ret[i] = newResult(RedisMessage{typ: '-', string: "MOVED 0 :2"}, nil)
					}
					return &redisresults{s: ret}
				}
				for i := range ret {
					ret[i] = newResult(RedisMessage{typ: '+', string: multi[i].Commands()[1]}, nil)
				}
				return &redisresults{s: ret}
			}}
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.DoMulti(context.Background(), client.B().Get().Key("a").Build())[0].ToString(); err != nil || v != "a" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if atomic.LoadInt64(&check) == 0 {
			t.Fatalf("unexpected check value %v", check)
		}
	})

	t.Run("slot moved new (multi 2)", func(t *testing.T) {
		var count, check int64
		client, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			if dst == ":2" {
				atomic.AddInt64(&check, 1)
			}
			return &mockConn{DoFn: func(cmd Completed) RedisResult {
				return slotsMultiResp
			}, DoMultiFn: func(multi ...Completed) *redisresults {
				ret := make([]RedisResult, len(multi))
				if atomic.AddInt64(&count, 1) <= 3 {
					for i := range ret {
						ret[i] = newResult(RedisMessage{typ: '-', string: "MOVED 0 :2"}, nil)
					}
					return &redisresults{s: ret}
				}
				for i := range ret {
					ret[i] = newResult(RedisMessage{typ: '+', string: multi[i].Commands()[1]}, nil)
				}
				return &redisresults{s: ret}
			}}
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		for i, resp := range client.DoMulti(context.Background(), client.B().Set().Key("a").Value("a").Build(), client.B().Get().Key("b").Build()) {
			if v, err := resp.ToString(); err != nil {
				t.Fatalf("unexpected resp %v %v", v, err)
			} else if i == 0 && v != "a" {
				t.Fatalf("unexpected resp %v %v", v, err)
			} else if i == 1 && v != "b" {
				t.Fatalf("unexpected resp %v %v", v, err)
			}
		}
		if atomic.LoadInt64(&check) == 0 {
			t.Fatalf("unexpected check value %v", check)
		}
	})

	t.Run("slot moved new (multi 2) TRYAGAIN", func(t *testing.T) {
		var count int64
		client, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return &mockConn{DoFn: func(cmd Completed) RedisResult {
				return slotsMultiResp
			}, DoMultiFn: func(multi ...Completed) *redisresults {
				ret := make([]RedisResult, len(multi))
				if atomic.AddInt64(&count, 1) <= 2 {
					for i := range ret {
						ret[i] = newResult(RedisMessage{typ: '-', string: "TRYAGAIN"}, nil)
					}
					return &redisresults{s: ret}
				}
				for i := range ret {
					ret[i] = newResult(RedisMessage{typ: '+', string: multi[i].Commands()[1]}, nil)
				}
				return &redisresults{s: ret}
			}}
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		for i, resp := range client.DoMulti(context.Background(), client.B().Set().Key("a").Value("a").Build(), client.B().Get().Key("b").Build()) {
			if v, err := resp.ToString(); err != nil && i != 0 && !strings.Contains(err.Error(), "TRYAGAIN") {
				t.Fatalf("unexpected resp %v %v", v, err)
			} else if i == 1 && v != "b" {
				t.Fatalf("unexpected resp %v %v", v, err)
			}
		}
	})

	t.Run("slot moved (cache)", func(t *testing.T) {
		var count int64
		client, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return &mockConn{
				DoFn: func(cmd Completed) RedisResult {
					return slotsMultiResp
				},
				DoCacheFn: func(cmd Cacheable, ttl time.Duration) RedisResult {
					if atomic.AddInt64(&count, 1) <= 3 {
						return newResult(RedisMessage{typ: '-', string: "MOVED 0 :1"}, nil)
					}
					return newResult(RedisMessage{typ: '+', string: "b"}, nil)
				},
			}
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.DoCache(context.Background(), client.B().Get().Key("a").Cache(), 100).ToString(); err != nil || v != "b" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot moved (cache multi 1)", func(t *testing.T) {
		var count int64
		client, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return &mockConn{DoFn: func(cmd Completed) RedisResult {
				return slotsMultiResp
			}, DoMultiCacheFn: func(multi ...CacheableTTL) *redisresults {
				ret := make([]RedisResult, len(multi))
				if atomic.AddInt64(&count, 1) <= 3 {
					for i := range ret {
						ret[i] = newResult(RedisMessage{typ: '-', string: "MOVED 0 :1"}, nil)
					}
					return &redisresults{s: ret}
				}
				for i := range ret {
					ret[i] = newResult(RedisMessage{typ: '+', string: multi[i].Cmd.Commands()[1]}, nil)
				}
				return &redisresults{s: ret}
			}}
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.DoMultiCache(context.Background(), CT(client.B().Get().Key("a").Cache(), 100))[0].ToString(); err != nil || v != "a" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot moved (cache multi 2)", func(t *testing.T) {
		var count int64
		client, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return &mockConn{DoFn: func(cmd Completed) RedisResult {
				return slotsMultiResp
			}, DoMultiCacheFn: func(multi ...CacheableTTL) *redisresults {
				ret := make([]RedisResult, len(multi))
				if atomic.AddInt64(&count, 1) <= 3 {
					for i := range ret {
						ret[i] = newResult(RedisMessage{typ: '-', string: "MOVED 0 :1"}, nil)
					}
					return &redisresults{s: ret}
				}
				for i := range ret {
					ret[i] = newResult(RedisMessage{typ: '+', string: multi[i].Cmd.Commands()[1]}, nil)
				}
				return &redisresults{s: ret}
			}}
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		for i, resp := range client.DoMultiCache(context.Background(), CT(client.B().Get().Key("a").Cache(), 100), CT(client.B().Get().Key("b").Cache(), 100)) {
			if v, err := resp.ToString(); err != nil {
				t.Fatalf("unexpected resp %v %v", v, err)
			} else if i == 0 && v != "a" {
				t.Fatalf("unexpected resp %v %v", v, err)
			} else if i == 1 && v != "b" {
				t.Fatalf("unexpected resp %v %v", v, err)
			}
		}
	})

	t.Run("slot asking", func(t *testing.T) {
		var count int64
		client, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return &mockConn{
				DoFn: func(cmd Completed) RedisResult {
					if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
						return slotsMultiResp
					}
					return newResult(RedisMessage{typ: '-', string: "ASK 0 :1"}, nil)
				},
				DoMultiFn: func(multi ...Completed) *redisresults {
					if atomic.AddInt64(&count, 1) <= 3 {
						return &redisresults{s: []RedisResult{{}, newResult(RedisMessage{typ: '-', string: "ASK 0 :1"}, nil)}}
					}
					return &redisresults{s: []RedisResult{{}, newResult(RedisMessage{typ: '+', string: "b"}, nil)}}
				},
			}
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.Do(context.Background(), client.B().Get().Key("a").Build()).ToString(); err != nil || v != "b" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot asking DoMulti (single)", func(t *testing.T) {
		var count int64
		client, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return &mockConn{
				DoFn: func(cmd Completed) RedisResult {
					return slotsMultiResp
				},
				DoMultiFn: func(multi ...Completed) *redisresults {
					ret := make([]RedisResult, len(multi))
					if atomic.AddInt64(&count, 1) <= 3 {
						for i := range ret {
							ret[i] = newResult(RedisMessage{typ: '-', string: "ASK 0 :1"}, nil)
						}
						return &redisresults{s: ret}
					}
					for i := 0; i < len(multi); i += 2 {
						ret[i] = newResult(RedisMessage{typ: '+', string: "OK"}, nil)
						ret[i+1] = newResult(RedisMessage{typ: '+', string: multi[i+1].Commands()[1]}, nil)
					}
					return &redisresults{s: ret}
				},
			}
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.DoMulti(context.Background(), client.B().Get().Key("a").Build())[0].ToString(); err != nil || v != "a" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot asking DoMulti (multi)", func(t *testing.T) {
		var count int64
		client, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return &mockConn{
				DoFn: func(cmd Completed) RedisResult {
					return slotsMultiResp
				},
				DoMultiFn: func(multi ...Completed) *redisresults {
					ret := make([]RedisResult, len(multi))
					if atomic.AddInt64(&count, 1) <= 3 {
						for i := range ret {
							ret[i] = newResult(RedisMessage{typ: '-', string: "ASK 0 :1"}, nil)
						}
						return &redisresults{s: ret}
					}
					for i := 0; i < len(multi); i += 2 {
						ret[i] = newResult(RedisMessage{typ: '+', string: "OK"}, nil)
						ret[i+1] = newResult(RedisMessage{typ: '+', string: multi[i+1].Commands()[1]}, nil)
					}
					return &redisresults{s: ret}
				},
			}
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		for i, resp := range client.DoMulti(context.Background(), client.B().Get().Key("a").Build(), client.B().Get().Key("b").Build()) {
			if v, err := resp.ToString(); err != nil {
				t.Fatalf("unexpected resp %v %v", v, err)
			} else if i == 0 && v != "a" {
				t.Fatalf("unexpected resp %v %v", v, err)
			} else if i == 1 && v != "b" {
				t.Fatalf("unexpected resp %v %v", v, err)
			}
		}
	})

	t.Run("slot asking (cache)", func(t *testing.T) {
		var count int64
		client, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return &mockConn{
				DoFn: func(cmd Completed) RedisResult {
					return slotsMultiResp
				},
				DoCacheFn: func(cmd Cacheable, ttl time.Duration) RedisResult {
					return newResult(RedisMessage{typ: '-', string: "ASK 0 :1"}, nil)
				},
				DoMultiFn: func(multi ...Completed) *redisresults {
					if atomic.AddInt64(&count, 1) <= 3 {
						return &redisresults{s: []RedisResult{{}, {}, {}, {}, {}, newResult(RedisMessage{typ: '-', string: "ASK 0 :1"}, nil)}}
					}
					return &redisresults{s: []RedisResult{{}, {}, {}, {}, {}, newResult(RedisMessage{typ: '*', values: []RedisMessage{{}, {typ: '+', string: "b"}}}, nil)}}
				},
			}
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.DoCache(context.Background(), client.B().Get().Key("a").Cache(), 100).ToString(); err != nil || v != "b" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot asking (cache multi 1)", func(t *testing.T) {
		var count int64
		client, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return &mockConn{
				DoFn: func(cmd Completed) RedisResult {
					return slotsMultiResp
				},
				DoMultiCacheFn: func(multi ...CacheableTTL) *redisresults {
					return &redisresults{s: []RedisResult{newResult(RedisMessage{typ: '-', string: "ASK 0 :1"}, nil)}}
				},
				DoMultiFn: func(multi ...Completed) *redisresults {
					if atomic.AddInt64(&count, 1) <= 3 {
						return &redisresults{s: []RedisResult{{}, {}, {}, {}, {}, newResult(RedisMessage{typ: '-', string: "ASK 0 :1"}, nil)}}
					}
					return &redisresults{s: []RedisResult{{}, {}, {}, {}, {}, newResult(RedisMessage{typ: '*', values: []RedisMessage{{}, {typ: '+', string: "b"}}}, nil)}}
				},
			}
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.DoMultiCache(context.Background(), CT(client.B().Get().Key("a").Cache(), 100))[0].ToString(); err != nil || v != "b" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot asking (cache multi 2)", func(t *testing.T) {
		var count int64
		client, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return &mockConn{
				DoFn: func(cmd Completed) RedisResult {
					return slotsMultiResp
				},
				DoMultiCacheFn: func(multi ...CacheableTTL) *redisresults {
					return &redisresults{s: []RedisResult{newResult(RedisMessage{typ: '-', string: "ASK 0 :1"}, nil)}}
				},
				DoMultiFn: func(multi ...Completed) *redisresults {
					if atomic.AddInt64(&count, 1) <= 3 {
						return &redisresults{s: []RedisResult{
							{}, {}, {}, {}, {}, newResult(RedisMessage{typ: '-', string: "ASK 0 :1"}, nil),
							{}, {}, {}, {}, {}, newResult(RedisMessage{typ: '-', string: "ASK 0 :1"}, nil),
						}}
					}
					return &redisresults{s: []RedisResult{
						{}, {}, {}, {}, {}, newResult(RedisMessage{typ: '*', values: []RedisMessage{{}, {}, {typ: '+', string: multi[4].Commands()[1]}}}, nil),
						{}, {}, {}, {}, {}, newResult(RedisMessage{typ: '*', values: []RedisMessage{{}, {}, {typ: '+', string: multi[10].Commands()[1]}}}, nil),
					}}
				},
			}
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		for i, resp := range client.DoMultiCache(context.Background(), CT(client.B().Get().Key("a").Cache(), 100), CT(client.B().Get().Key("b").Cache(), 100)) {
			if v, err := resp.ToString(); err != nil {
				t.Fatalf("unexpected resp %v %v", v, err)
			} else if i == 0 && v != "a" {
				t.Fatalf("unexpected resp %v %v", v, err)
			} else if i == 1 && v != "b" {
				t.Fatalf("unexpected resp %v %v", v, err)
			}
		}
	})

	t.Run("slot try again", func(t *testing.T) {
		var count int64
		client, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return &mockConn{DoFn: func(cmd Completed) RedisResult {
				if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
					return slotsMultiResp
				}
				if atomic.AddInt64(&count, 1) <= 3 {
					return newResult(RedisMessage{typ: '-', string: "TRYAGAIN"}, nil)
				}
				return newResult(RedisMessage{typ: '+', string: "b"}, nil)
			}}
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.Do(context.Background(), client.B().Get().Key("a").Build()).ToString(); err != nil || v != "b" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot try again DoMulti 1", func(t *testing.T) {
		var count int64
		client, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return &mockConn{DoFn: func(cmd Completed) RedisResult {
				return slotsMultiResp
			}, DoMultiFn: func(multi ...Completed) *redisresults {
				if atomic.AddInt64(&count, 1) <= 3 {
					return &redisresults{s: []RedisResult{newResult(RedisMessage{typ: '-', string: "TRYAGAIN"}, nil)}}
				}
				ret := make([]RedisResult, len(multi))
				ret[0] = newResult(RedisMessage{typ: '+', string: "b"}, nil)
				return &redisresults{s: ret}
			}}
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.DoMulti(context.Background(), client.B().Get().Key("a").Build())[0].ToString(); err != nil || v != "b" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot try again DoMulti 2", func(t *testing.T) {
		var count int64
		client, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return &mockConn{DoFn: func(cmd Completed) RedisResult {
				return slotsMultiResp
			}, DoMultiFn: func(multi ...Completed) *redisresults {
				if atomic.AddInt64(&count, 1) <= 3 {
					return &redisresults{s: []RedisResult{newResult(RedisMessage{typ: '-', string: "TRYAGAIN"}, nil)}}
				}
				ret := make([]RedisResult, len(multi))
				ret[0] = newResult(RedisMessage{typ: '+', string: multi[0].Commands()[1]}, nil)
				return &redisresults{s: ret}
			}}
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		for i, resp := range client.DoMulti(context.Background(), client.B().Get().Key("a").Build(), client.B().Get().Key("b").Build()) {
			if v, err := resp.ToString(); err != nil {
				t.Fatalf("unexpected resp %v %v", v, err)
			} else if i == 0 && v != "a" {
				t.Fatalf("unexpected resp %v %v", v, err)
			} else if i == 1 && v != "b" {
				t.Fatalf("unexpected resp %v %v", v, err)
			}
		}
	})

	t.Run("slot try again (cache)", func(t *testing.T) {
		var count int64
		client, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return &mockConn{
				DoFn: func(cmd Completed) RedisResult {
					return slotsMultiResp
				},
				DoCacheFn: func(cmd Cacheable, ttl time.Duration) RedisResult {
					if atomic.AddInt64(&count, 1) <= 3 {
						return newResult(RedisMessage{typ: '-', string: "TRYAGAIN"}, nil)
					}
					return newResult(RedisMessage{typ: '+', string: "b"}, nil)
				},
			}
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.DoCache(context.Background(), client.B().Get().Key("a").Cache(), 100).ToString(); err != nil || v != "b" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot try again (cache multi 1)", func(t *testing.T) {
		var count int64
		client, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return &mockConn{DoFn: func(cmd Completed) RedisResult {
				return slotsMultiResp
			}, DoMultiCacheFn: func(multi ...CacheableTTL) *redisresults {
				if atomic.AddInt64(&count, 1) <= 3 {
					return &redisresults{s: []RedisResult{newResult(RedisMessage{typ: '-', string: "TRYAGAIN"}, nil)}}
				}
				return &redisresults{s: []RedisResult{newResult(RedisMessage{typ: '+', string: "b"}, nil)}}
			}}
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.DoMultiCache(context.Background(), CT(client.B().Get().Key("a").Cache(), 100))[0].ToString(); err != nil || v != "b" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot try again (cache multi 2)", func(t *testing.T) {
		var count int64
		client, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return &mockConn{DoFn: func(cmd Completed) RedisResult {
				if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
					return slotsMultiResp
				}
				return shardsMultiResp
			}, DoMultiCacheFn: func(multi ...CacheableTTL) *redisresults {
				if atomic.AddInt64(&count, 1) <= 3 {
					return &redisresults{s: []RedisResult{newResult(RedisMessage{typ: '-', string: "TRYAGAIN"}, nil)}}
				}
				return &redisresults{s: []RedisResult{newResult(RedisMessage{typ: '+', string: multi[0].Cmd.Commands()[1]}, nil)}}
			}}
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		for i, resp := range client.DoMultiCache(context.Background(), CT(client.B().Get().Key("a").Cache(), 100), CT(client.B().Get().Key("b").Cache(), 100)) {
			if v, err := resp.ToString(); err != nil {
				t.Fatalf("unexpected resp %v %v", v, err)
			} else if i == 0 && v != "a" {
				t.Fatalf("unexpected resp %v %v", v, err)
			} else if i == 1 && v != "b" {
				t.Fatalf("unexpected resp %v %v", v, err)
			}
		}
	})
}

func TestClusterClientRetry(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	SetupClientRetry(t, func(m *mockConn) Client {
		m.DoOverride = map[string]func(cmd Completed) RedisResult{
			"CLUSTER SLOTS": func(cmd Completed) RedisResult { return slotsMultiResp },
		}
		c, err := newClusterClient(&ClientOption{InitAddress: []string{":0"}}, func(dst string, opt *ClientOption) conn {
			return m
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		return c
	})
}

func TestClusterClientReplicaOnly_PickReplica(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	m := &mockConn{
		DoFn: func(cmd Completed) RedisResult {
			if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
				return slotsMultiResp
			}
			return RedisResult{}
		},
	}

	client, err := newClusterClient(&ClientOption{InitAddress: []string{"127.0.0.1:0"}, ReplicaOnly: true}, func(dst string, opt *ClientOption) conn {
		copiedM := *m
		return &copiedM
	})
	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}
	t.Run("replicas should be picked", func(t *testing.T) {
		if client.pslots[0] != client.conns["127.0.1.1:1"].conn {
			t.Fatalf("unexpected replica node assigned to slot 0")
		}
		if client.pslots[8192] != client.conns["127.0.1.1:1"].conn {
			t.Fatalf("unexpected replica node assigned to slot 8192")
		}
		if client.pslots[8193] != client.conns["127.0.3.1:1"].conn {
			t.Fatalf("unexpected replica node assigned to slot 8193")
		}
		if client.pslots[16383] != client.conns["127.0.3.1:1"].conn {
			t.Fatalf("unexpected replica node assigned to slot 16383")
		}
	})
}

func TestClusterClientReplicaOnly_PickMasterIfNoReplica(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	t.Run("replicas should be picked", func(t *testing.T) {
		m := &mockConn{
			DoFn: func(cmd Completed) RedisResult {
				if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
					return slotsMultiResp
				}
				return RedisResult{}
			},
		}

		client, err := newClusterClient(&ClientOption{InitAddress: []string{"127.0.0.1:0"}, ReplicaOnly: true}, func(dst string, opt *ClientOption) conn {
			copiedM := *m
			return &copiedM
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}

		if client.pslots[0] != client.conns["127.0.1.1:1"].conn {
			t.Fatalf("unexpected replica node assigned to slot 0")
		}
		if client.pslots[8192] != client.conns["127.0.1.1:1"].conn {
			t.Fatalf("unexpected replica node assigned to slot 8192")
		}
		if client.pslots[8193] != client.conns["127.0.3.1:1"].conn {
			t.Fatalf("unexpected replica node assigned to slot 8193")
		}
		if client.pslots[16383] != client.conns["127.0.3.1:1"].conn {
			t.Fatalf("unexpected replica node assigned to slot 16383")
		}
	})

	t.Run("distributed to replicas", func(t *testing.T) {
		m := &mockConn{
			DoFn: func(cmd Completed) RedisResult {
				if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
					return slotsMultiRespWithMultiReplicas
				}
				return RedisResult{}
			},
		}

		client, err := newClusterClient(&ClientOption{InitAddress: []string{"127.0.0.1:0"}, ReplicaOnly: true}, func(dst string, opt *ClientOption) conn {
			copiedM := *m
			return &copiedM
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}

		for slot := 0; slot < 8193; slot++ {
			if client.pslots[slot] == client.conns["127.0.0.2:1"].conn {
				continue
			}
			if client.pslots[slot] == client.conns["127.0.0.3:2"].conn {
				continue
			}
			if client.pslots[slot] == client.conns["127.0.0.4:3"].conn {
				continue
			}

			t.Fatalf("unexpected replica node assigned to slot %d", slot)
		}

		for slot := 8193; slot < 16384; slot++ {
			if client.pslots[slot] == client.conns["127.0.1.2:1"].conn {
				continue
			}
			if client.pslots[slot] == client.conns["127.0.1.3:2"].conn {
				continue
			}
			if client.pslots[slot] == client.conns["127.0.1.4:3"].conn {
				continue
			}

			t.Fatalf("unexpected replica node assigned to slot %d", slot)
		}
	})
}

func TestClusterShardsParsing(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	t.Run("master selection", func(t *testing.T) {
		result := parseShards(shardsRespTls.val, "127.0.0.1:5", true)
		if len(result) != 1 {
			t.Fatalf("unexpected result %v", result)
		}
		if _, ok := result["127.0.1.1:1"]; !ok {
			t.Fatal("unexpected master node")
		}
	})

	t.Run("port selection", func(t *testing.T) {
		result := parseShards(shardsRespTls.val, "127.0.0.1:5", true)
		if len(result) != 1 {
			t.Fatalf("unexpected result %v", result)
		}
		for _, val := range result {
			nodes := val.nodes
			sort.Strings(nodes)
			if len(nodes) != 3 ||
				nodes[0] != "127.0.1.1:1" ||
				nodes[1] != "127.0.2.1:2" ||
				nodes[2] != "127.0.3.1:3" {
				t.Fatalf("unexpected nodes %v", nodes)
			}
		}

		result = parseShards(shardsRespTls.val, "127.0.0.1:5", false)
		if len(result) != 1 {
			t.Fatalf("unexpected result %v", result)
		}
		for _, val := range result {
			nodes := val.nodes
			sort.Strings(nodes)
			if len(nodes) != 3 ||
				nodes[0] != "127.0.1.1:0" ||
				nodes[1] != "127.0.2.1:0" ||
				nodes[2] != "127.0.3.1:3" {
				t.Fatalf("unexpected nodes %v", nodes)
			}
		}
	})

	t.Run("master position", func(t *testing.T) {
		result := parseShards(shardsRespTls.val, "127.0.0.1:5", true)
		if len(result) != 1 {
			t.Fatalf("unexpected result %v", result)
		}
		for master, group := range result {
			if len(group.nodes) == 0 || group.nodes[0] != master {
				t.Fatalf("unexpected first node %v", group)
			}
		}
	})
}
