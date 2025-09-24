package rueidis

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

var slotsResp = newResult(slicemsg('*', []RedisMessage{
	slicemsg('*', []RedisMessage{
		{typ: ':', intlen: 0},
		{typ: ':', intlen: 16383},
		slicemsg('*', []RedisMessage{ // master
			strmsg('+', "127.0.0.1"),
			{typ: ':', intlen: 0},
			strmsg('+', ""),
		}),
		slicemsg('*', []RedisMessage{ // replica
			strmsg('+', "127.0.1.1"),
			{typ: ':', intlen: 1},
			strmsg('+', ""),
		}),
	}),
}), nil)

var slotsMultiResp = newResult(slicemsg('*', []RedisMessage{
	slicemsg('*', []RedisMessage{
		{typ: ':', intlen: 0},
		{typ: ':', intlen: 8192},
		slicemsg('*', []RedisMessage{ // master
			strmsg('+', "127.0.0.1"),
			{typ: ':', intlen: 0},
			strmsg('+', ""),
		}),
		slicemsg('*', []RedisMessage{ // replica
			strmsg('+', "127.0.1.1"),
			{typ: ':', intlen: 1},
			strmsg('+', ""),
		}),
	}),
	slicemsg('*', []RedisMessage{
		{typ: ':', intlen: 8193},
		{typ: ':', intlen: 16383},
		slicemsg('*', []RedisMessage{ // master
			strmsg('+', "127.0.2.1"),
			{typ: ':', intlen: 0},
			strmsg('+', ""),
		}),
		slicemsg('*', []RedisMessage{ // replica
			strmsg('+', "127.0.3.1"),
			{typ: ':', intlen: 1},
			strmsg('+', ""),
		}),
	}),
}), nil)

var slotsMultiRespWithoutReplicas = newResult(slicemsg('*', []RedisMessage{
	slicemsg('*', []RedisMessage{
		{typ: ':', intlen: 0},
		{typ: ':', intlen: 8192},
		slicemsg('*', []RedisMessage{ // master
			strmsg('+', "127.0.0.1"),
			{typ: ':', intlen: 0},
			strmsg('+', ""),
		}),
	}),
	slicemsg('*', []RedisMessage{
		{typ: ':', intlen: 8193},
		{typ: ':', intlen: 16383},
		slicemsg('*', []RedisMessage{ // master
			strmsg('+', "127.0.1.1"),
			{typ: ':', intlen: 0},
			strmsg('+', ""),
		}),
	}),
}), nil)

var slotsMultiRespWithMultiReplicas = newResult(slicemsg('*', []RedisMessage{
	slicemsg('*', []RedisMessage{
		{typ: ':', intlen: 0},
		{typ: ':', intlen: 8192},
		slicemsg('*', []RedisMessage{ // master
			strmsg('+', "127.0.0.1"),
			{typ: ':', intlen: 0},
			strmsg('+', ""),
		}),
		slicemsg('*', []RedisMessage{ // replica1
			strmsg('+', "127.0.0.2"),
			{typ: ':', intlen: 1},
			strmsg('+', ""),
		}),
		slicemsg('*', []RedisMessage{ // replica2
			strmsg('+', "127.0.0.3"),
			{typ: ':', intlen: 2},
			strmsg('+', ""),
		}),
		slicemsg('*', []RedisMessage{ // replica3
			strmsg('+', "127.0.0.4"),
			{typ: ':', intlen: 3},
			strmsg('+', ""),
		}),
	}),
	slicemsg('*', []RedisMessage{
		{typ: ':', intlen: 8193},
		{typ: ':', intlen: 16383},
		slicemsg('*', []RedisMessage{ // master
			strmsg('+', "127.0.1.1"),
			{typ: ':', intlen: 0},
			strmsg('+', ""),
		}),
		slicemsg('*', []RedisMessage{ // replica1
			strmsg('+', "127.0.1.2"),
			{typ: ':', intlen: 1},
			strmsg('+', ""),
		}),
		slicemsg('*', []RedisMessage{ // replica2
			strmsg('+', "127.0.1.3"),
			{typ: ':', intlen: 2},
			strmsg('+', ""),
		}),
		slicemsg('*', []RedisMessage{ // replica3
			strmsg('+', "127.0.1.4"),
			{typ: ':', intlen: 3},
			strmsg('+', ""),
		}),
	}),
}), nil)

var singleSlotResp = newResult(slicemsg('*', []RedisMessage{
	slicemsg('*', []RedisMessage{
		{typ: ':', intlen: 0},
		{typ: ':', intlen: 0},
		slicemsg('*', []RedisMessage{ // master
			strmsg('+', "127.0.0.1"),
			{typ: ':', intlen: 0},
			strmsg('+', ""),
		}),
	}),
}), nil)

var singleSlotResp2 = newResult(slicemsg('*', []RedisMessage{
	slicemsg('*', []RedisMessage{
		{typ: ':', intlen: 0},
		{typ: ':', intlen: 0},
		slicemsg('*', []RedisMessage{ // master
			strmsg('+', "127.0.3.1"),
			{typ: ':', intlen: 3},
			strmsg('+', ""),
		}),
	}),
}), nil)

var singleSlotWithoutIP = newResult(slicemsg('*', []RedisMessage{
	slicemsg('*', []RedisMessage{
		{typ: ':', intlen: 0},
		{typ: ':', intlen: 0},
		slicemsg('*', []RedisMessage{ // master
			strmsg('+', ""),
			{typ: ':', intlen: 4},
			strmsg('+', ""),
		}),
		slicemsg('*', []RedisMessage{ // replica
			strmsg('+', "?"),
			{typ: ':', intlen: 1},
			strmsg('+', ""),
		}),
	}),
	slicemsg('*', []RedisMessage{
		{typ: ':', intlen: 0},
		{typ: ':', intlen: 0},
		slicemsg('*', []RedisMessage{ // master
			strmsg('+', "?"),
			{typ: ':', intlen: 4},
			strmsg('+', ""),
		}),
	}),
}), nil)

var shardsResp = newResult(slicemsg(typeArray, []RedisMessage{
	slicemsg(typeMap, []RedisMessage{
		strmsg(typeBlobString, "slots"),
		slicemsg(typeArray, []RedisMessage{
			strmsg(typeBlobString, "0"),
			strmsg(typeBlobString, "16383"),
		}),
		strmsg(typeBlobString, "nodes"),
		slicemsg(typeArray, []RedisMessage{
			slicemsg(typeMap, []RedisMessage{ // failed master
				strmsg(typeBlobString, "id"),
				strmsg(typeBlobString, ""),

				strmsg(typeBlobString, "port"),
				{typ: typeInteger, intlen: 0},

				strmsg(typeBlobString, "ip"),
				strmsg(typeBlobString, "127.0.0.99"),

				strmsg(typeBlobString, "endpoint"),
				strmsg(typeBlobString, "127.0.0.99"),

				strmsg(typeBlobString, "role"),
				strmsg(typeBlobString, "master"),

				strmsg(typeBlobString, "replication-offset"),
				{typ: typeInteger, intlen: 72156},

				strmsg(typeBlobString, "health"),
				strmsg(typeBlobString, "fail"),
			}),
			slicemsg(typeMap, []RedisMessage{ // master
				strmsg(typeBlobString, "id"),
				strmsg(typeBlobString, ""),

				strmsg(typeBlobString, "port"),
				{typ: typeInteger, intlen: 0},

				strmsg(typeBlobString, "ip"),
				strmsg(typeBlobString, "127.0.0.1"),

				strmsg(typeBlobString, "endpoint"),
				strmsg(typeBlobString, "127.0.0.1"),

				strmsg(typeBlobString, "role"),
				strmsg(typeBlobString, "master"),

				strmsg(typeBlobString, "replication-offset"),
				{typ: typeInteger, intlen: 72156},

				strmsg(typeBlobString, "health"),
				strmsg(typeBlobString, "online"),
			}),
			slicemsg(typeMap, []RedisMessage{ // replica
				strmsg(typeBlobString, "id"),
				strmsg(typeBlobString, ""),

				strmsg(typeBlobString, "port"),
				{typ: typeInteger, intlen: 1},

				strmsg(typeBlobString, "ip"),
				strmsg(typeBlobString, "127.0.1.1"),

				strmsg(typeBlobString, "endpoint"),
				strmsg(typeBlobString, "127.0.1.1"),

				strmsg(typeBlobString, "role"),
				strmsg(typeBlobString, "replica"),

				strmsg(typeBlobString, "replication-offset"),
				{typ: typeInteger, intlen: 72156},

				strmsg(typeBlobString, "health"),
				strmsg(typeBlobString, "online"),
			}),
		}),
	}),
}), nil)

var shardsRespTls = newResult(slicemsg(typeArray, []RedisMessage{
	slicemsg(typeMap, []RedisMessage{
		strmsg(typeBlobString, "slots"),
		slicemsg(typeArray, []RedisMessage{
			strmsg(typeBlobString, "0"),
			strmsg(typeBlobString, "16383"),
		}),
		strmsg(typeBlobString, "nodes"),
		slicemsg(typeArray, []RedisMessage{
			slicemsg(typeMap, []RedisMessage{ // replica, tls
				strmsg(typeBlobString, "id"),
				strmsg(typeBlobString, ""),

				strmsg(typeBlobString, "tls-port"),
				{typ: typeInteger, intlen: 2},

				strmsg(typeBlobString, "ip"),
				strmsg(typeBlobString, "127.0.2.1"),

				strmsg(typeBlobString, "endpoint"),
				strmsg(typeBlobString, "127.0.2.1"),

				strmsg(typeBlobString, "role"),
				strmsg(typeBlobString, "replica"),

				strmsg(typeBlobString, "replication-offset"),
				{typ: typeInteger, intlen: 72156},

				strmsg(typeBlobString, "health"),
				strmsg(typeBlobString, "online"),
			}),
			slicemsg(typeMap, []RedisMessage{ // failed master, tls + port
				strmsg(typeBlobString, "id"),
				strmsg(typeBlobString, ""),

				strmsg(typeBlobString, "port"),
				{typ: typeInteger, intlen: 0},

				strmsg(typeBlobString, "tls-port"),
				{typ: typeInteger, intlen: 1},

				strmsg(typeBlobString, "ip"),
				strmsg(typeBlobString, "127.0.1.99"),

				strmsg(typeBlobString, "endpoint"),
				strmsg(typeBlobString, "127.0.1.99"),

				strmsg(typeBlobString, "role"),
				strmsg(typeBlobString, "master"),

				strmsg(typeBlobString, "replication-offset"),
				{typ: typeInteger, intlen: 72156},

				strmsg(typeBlobString, "health"),
				strmsg(typeBlobString, "fail"),
			}),
			slicemsg(typeMap, []RedisMessage{ // master, tls + port
				strmsg(typeBlobString, "id"),
				strmsg(typeBlobString, ""),

				strmsg(typeBlobString, "port"),
				{typ: typeInteger, intlen: 0},

				strmsg(typeBlobString, "tls-port"),
				{typ: typeInteger, intlen: 1},

				strmsg(typeBlobString, "ip"),
				strmsg(typeBlobString, "127.0.1.1"),

				strmsg(typeBlobString, "endpoint"),
				strmsg(typeBlobString, "127.0.1.1"),

				strmsg(typeBlobString, "role"),
				strmsg(typeBlobString, "master"),

				strmsg(typeBlobString, "replication-offset"),
				{typ: typeInteger, intlen: 72156},

				strmsg(typeBlobString, "health"),
				strmsg(typeBlobString, "online"),
			}),
			slicemsg(typeMap, []RedisMessage{ // replica, port
				strmsg(typeBlobString, "id"),
				strmsg(typeBlobString, ""),

				strmsg(typeBlobString, "port"),
				{typ: typeInteger, intlen: 3},

				strmsg(typeBlobString, "ip"),
				strmsg(typeBlobString, "127.0.3.1"),

				strmsg(typeBlobString, "endpoint"),
				strmsg(typeBlobString, "127.0.3.1"),

				strmsg(typeBlobString, "role"),
				strmsg(typeBlobString, "replica"),

				strmsg(typeBlobString, "replication-offset"),
				{typ: typeInteger, intlen: 72156},

				strmsg(typeBlobString, "health"),
				strmsg(typeBlobString, "online"),
			}),
		}),
	}),
}), nil)

var shardsMultiResp = newResult(slicemsg('*', []RedisMessage{
	slicemsg(typeMap, []RedisMessage{
		strmsg(typeBlobString, "slots"),
		slicemsg(typeArray, []RedisMessage{
			strmsg(typeBlobString, "0"),
			strmsg(typeBlobString, "8192"),
		}),
		strmsg(typeBlobString, "nodes"),
		slicemsg(typeArray, []RedisMessage{
			slicemsg(typeMap, []RedisMessage{ // failed master
				strmsg(typeBlobString, "id"),
				strmsg(typeBlobString, ""),

				strmsg(typeBlobString, "port"),
				{typ: typeInteger, intlen: 0},

				strmsg(typeBlobString, "ip"),
				strmsg(typeBlobString, "127.0.0.99"),

				strmsg(typeBlobString, "endpoint"),
				strmsg(typeBlobString, "127.0.0.99"),

				strmsg(typeBlobString, "role"),
				strmsg(typeBlobString, "master"),

				strmsg(typeBlobString, "replication-offset"),
				{typ: typeInteger, intlen: 72156},

				strmsg(typeBlobString, "health"),
				strmsg(typeBlobString, "fail"),
			}),
			slicemsg(typeMap, []RedisMessage{ // master
				strmsg(typeBlobString, "id"),
				strmsg(typeBlobString, ""),

				strmsg(typeBlobString, "port"),
				{typ: typeInteger, intlen: 0},

				strmsg(typeBlobString, "ip"),
				strmsg(typeBlobString, "127.0.0.1"),

				strmsg(typeBlobString, "endpoint"),
				strmsg(typeBlobString, "127.0.0.1"),

				strmsg(typeBlobString, "role"),
				strmsg(typeBlobString, "master"),

				strmsg(typeBlobString, "replication-offset"),
				{typ: typeInteger, intlen: 72156},

				strmsg(typeBlobString, "health"),
				strmsg(typeBlobString, "online"),
			}),
			slicemsg(typeMap, []RedisMessage{ // replica
				strmsg(typeBlobString, "id"),
				strmsg(typeBlobString, ""),

				strmsg(typeBlobString, "port"),
				{typ: typeInteger, intlen: 1},

				strmsg(typeBlobString, "ip"),
				strmsg(typeBlobString, "127.0.1.1"),

				strmsg(typeBlobString, "endpoint"),
				strmsg(typeBlobString, "127.0.1.1"),

				strmsg(typeBlobString, "role"),
				strmsg(typeBlobString, "replica"),

				strmsg(typeBlobString, "replication-offset"),
				{typ: typeInteger, intlen: 72156},

				strmsg(typeBlobString, "health"),
				strmsg(typeBlobString, "online"),
			}),
		}),
	}),
	slicemsg(typeMap, []RedisMessage{
		strmsg(typeBlobString, "slots"),
		slicemsg(typeArray, []RedisMessage{
			strmsg(typeBlobString, "8193"),
			strmsg(typeBlobString, "16383"),
		}),
		strmsg(typeBlobString, "nodes"),
		slicemsg(typeArray, []RedisMessage{
			slicemsg(typeMap, []RedisMessage{ // failed master
				strmsg(typeBlobString, "id"),
				strmsg(typeBlobString, ""),

				strmsg(typeBlobString, "port"),
				{typ: typeInteger, intlen: 0},

				strmsg(typeBlobString, "ip"),
				strmsg(typeBlobString, "127.0.2.99"),

				strmsg(typeBlobString, "endpoint"),
				strmsg(typeBlobString, "127.0.2.99"),

				strmsg(typeBlobString, "role"),
				strmsg(typeBlobString, "master"),

				strmsg(typeBlobString, "replication-offset"),
				{typ: typeInteger, intlen: 72156},

				strmsg(typeBlobString, "health"),
				strmsg(typeBlobString, "fail"),
			}),
			slicemsg(typeMap, []RedisMessage{ // master
				strmsg(typeBlobString, "id"),
				strmsg(typeBlobString, ""),

				strmsg(typeBlobString, "port"),
				{typ: typeInteger, intlen: 0},

				strmsg(typeBlobString, "ip"),
				strmsg(typeBlobString, "127.0.2.1"),

				strmsg(typeBlobString, "endpoint"),
				strmsg(typeBlobString, "127.0.2.1"),

				strmsg(typeBlobString, "role"),
				strmsg(typeBlobString, "master"),

				strmsg(typeBlobString, "replication-offset"),
				{typ: typeInteger, intlen: 72156},

				strmsg(typeBlobString, "health"),
				strmsg(typeBlobString, "online"),
			}),
			slicemsg(typeMap, []RedisMessage{ // replica
				strmsg(typeBlobString, "id"),
				strmsg(typeBlobString, ""),

				strmsg(typeBlobString, "port"),
				{typ: typeInteger, intlen: 1},

				strmsg(typeBlobString, "ip"),
				strmsg(typeBlobString, "127.0.3.1"),

				strmsg(typeBlobString, "endpoint"),
				strmsg(typeBlobString, "127.0.3.1"),

				strmsg(typeBlobString, "role"),
				strmsg(typeBlobString, "replica"),

				strmsg(typeBlobString, "replication-offset"),
				{typ: typeInteger, intlen: 72156},

				strmsg(typeBlobString, "health"),
				strmsg(typeBlobString, "online"),
			}),
		}),
	}),
}), nil)

var singleShardResp2 = newResult(slicemsg('*', []RedisMessage{
	slicemsg(typeMap, []RedisMessage{
		strmsg(typeBlobString, "slots"),
		slicemsg(typeArray, []RedisMessage{
			strmsg(typeBlobString, "0"),
			strmsg(typeBlobString, "0"),
		}),
		strmsg(typeBlobString, "nodes"),
		slicemsg(typeArray, []RedisMessage{
			slicemsg(typeMap, []RedisMessage{ // failed master
				strmsg(typeBlobString, "id"),
				strmsg(typeBlobString, ""),

				strmsg(typeBlobString, "port"),
				{typ: typeInteger, intlen: 3},

				strmsg(typeBlobString, "ip"),
				strmsg(typeBlobString, "127.0.3.99"),

				strmsg(typeBlobString, "endpoint"),
				strmsg(typeBlobString, "127.0.3.99"),

				strmsg(typeBlobString, "role"),
				strmsg(typeBlobString, "master"),

				strmsg(typeBlobString, "replication-offset"),
				{typ: typeInteger, intlen: 72156},

				strmsg(typeBlobString, "health"),
				strmsg(typeBlobString, "fail"),
			}),
			slicemsg(typeMap, []RedisMessage{ // master
				strmsg(typeBlobString, "id"),
				strmsg(typeBlobString, ""),

				strmsg(typeBlobString, "port"),
				{typ: typeInteger, intlen: 3},

				strmsg(typeBlobString, "ip"),
				strmsg(typeBlobString, "127.0.3.1"),

				strmsg(typeBlobString, "endpoint"),
				strmsg(typeBlobString, "127.0.3.1"),

				strmsg(typeBlobString, "role"),
				strmsg(typeBlobString, "master"),

				strmsg(typeBlobString, "replication-offset"),
				{typ: typeInteger, intlen: 72156},

				strmsg(typeBlobString, "health"),
				strmsg(typeBlobString, "online"),
			}),
		}),
	}),
}), nil)

var singleShardWithoutIP = newResult(slicemsg(typeArray, []RedisMessage{
	slicemsg(typeMap, []RedisMessage{
		strmsg(typeBlobString, "slots"),
		slicemsg(typeArray, []RedisMessage{
			strmsg(typeBlobString, "0"),
			strmsg(typeBlobString, "0"),
		}),
		strmsg(typeBlobString, "nodes"),
		slicemsg(typeArray, []RedisMessage{
			slicemsg(typeMap, []RedisMessage{ // failed master
				strmsg(typeBlobString, "id"),
				strmsg(typeBlobString, ""),

				strmsg(typeBlobString, "port"),
				{typ: typeInteger, intlen: 4},

				strmsg(typeBlobString, "ip"),
				strmsg(typeBlobString, ""),

				strmsg(typeBlobString, "endpoint"),
				strmsg(typeBlobString, ""),

				strmsg(typeBlobString, "role"),
				strmsg(typeBlobString, "master"),

				strmsg(typeBlobString, "replication-offset"),
				{typ: typeInteger, intlen: 72156},

				strmsg(typeBlobString, "health"),
				strmsg(typeBlobString, "fail"),
			}),
			slicemsg(typeMap, []RedisMessage{ // master
				strmsg(typeBlobString, "id"),
				strmsg(typeBlobString, ""),

				strmsg(typeBlobString, "port"),
				{typ: typeInteger, intlen: 4},

				strmsg(typeBlobString, "ip"),
				strmsg(typeBlobString, ""),

				strmsg(typeBlobString, "endpoint"),
				strmsg(typeBlobString, ""),

				strmsg(typeBlobString, "role"),
				strmsg(typeBlobString, "master"),

				strmsg(typeBlobString, "replication-offset"),
				{typ: typeInteger, intlen: 72156},

				strmsg(typeBlobString, "health"),
				strmsg(typeBlobString, "online"),
			}),
			slicemsg(typeMap, []RedisMessage{ // replica
				strmsg(typeBlobString, "id"),
				strmsg(typeBlobString, ""),

				strmsg(typeBlobString, "port"),
				{typ: typeInteger, intlen: 1},

				strmsg(typeBlobString, "ip"),
				strmsg(typeBlobString, "?"),

				strmsg(typeBlobString, "endpoint"),
				strmsg(typeBlobString, "?"),

				strmsg(typeBlobString, "role"),
				strmsg(typeBlobString, "replica"),

				strmsg(typeBlobString, "replication-offset"),
				{typ: typeInteger, intlen: 72156},

				strmsg(typeBlobString, "health"),
				strmsg(typeBlobString, "online"),
			}),
		}),
	}),
	slicemsg(typeMap, []RedisMessage{
		strmsg(typeBlobString, "slots"),
		slicemsg(typeArray, []RedisMessage{
			strmsg(typeBlobString, "0"),
			strmsg(typeBlobString, "0"),
		}),
		strmsg(typeBlobString, "nodes"),
		slicemsg(typeArray, []RedisMessage{
			slicemsg(typeMap, []RedisMessage{ // master
				strmsg(typeBlobString, "id"),
				strmsg(typeBlobString, ""),

				strmsg(typeBlobString, "port"),
				{typ: typeInteger, intlen: 4},

				strmsg(typeBlobString, "ip"),
				strmsg(typeBlobString, "?"),

				strmsg(typeBlobString, "endpoint"),
				strmsg(typeBlobString, "?"),

				strmsg(typeBlobString, "role"),
				strmsg(typeBlobString, "master"),

				strmsg(typeBlobString, "replication-offset"),
				{typ: typeInteger, intlen: 72156},

				strmsg(typeBlobString, "health"),
				strmsg(typeBlobString, "online"),
			}),
		}),
	}),
}), nil)

//gocyclo:ignore
func TestClusterClientInit(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	t.Run("Init no nodes", func(t *testing.T) {
		if _, err := newClusterClient(
			&ClientOption{InitAddress: []string{}},
			func(dst string, opt *ClientOption) conn { return nil },
			newRetryer(defaultRetryDelayFn),
		); err != ErrNoAddr {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("Init no dialable", func(t *testing.T) {
		v := errors.New("dial err")
		if _, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn {
				return &mockConn{DialFn: func() error { return v }}
			},
			newRetryer(defaultRetryDelayFn),
		); err != v {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("Init ReadNodeSelector and ReplicaSelector both should not be set", func(t *testing.T) {
		client, err := newClusterClient(
			&ClientOption{
				InitAddress: []string{":0"},
				SendToReplicas: func(cmd Completed) bool {
					return true
				},
				ReadNodeSelector: func(slot uint16, nodes []NodeInfo) int {
					return 0
				}},
			func(dst string, opt *ClientOption) conn {
				return &mockConn{DialFn: func() error { return nil }}
			},
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if client.opt.ReplicaSelector != nil {
			t.Fatal("unexpected ReplicaSelector set")
		}
	})

	t.Run("Refresh err", func(t *testing.T) {
		v := errors.New("refresh err")
		if _, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn {
				return &mockConn{DoFn: func(cmd Completed) RedisResult { return newErrResult(v) }}
			},
			newRetryer(defaultRetryDelayFn),
		); err != v {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("Refresh skip zero slots", func(t *testing.T) {
		var first int64
		if _, err := newClusterClient(
			&ClientOption{InitAddress: []string{"127.0.0.1:0", "127.0.1.1:1"}},
			func(dst string, opt *ClientOption) conn {
				return &mockConn{
					DoFn: func(cmd Completed) RedisResult {
						if atomic.AddInt64(&first, 1) == 1 {
							return newResult(slicemsg('*', []RedisMessage{}), nil)
						}
						return slotsResp
					},
				}
			},
			newRetryer(defaultRetryDelayFn),
		); err != nil || atomic.AddInt64(&first, 1) < 2 {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("Refresh skip zero shards", func(t *testing.T) {
		var first int64
		if _, err := newClusterClient(
			&ClientOption{InitAddress: []string{"127.0.0.1:0", "127.0.1.1:1"}},
			func(dst string, opt *ClientOption) conn {
				return &mockConn{
					DoFn: func(cmd Completed) RedisResult {
						if atomic.AddInt64(&first, 1) == 1 {
							return newResult(slicemsg('*', []RedisMessage{}), nil)
						}
						return shardsResp
					},
					VersionFn: func() int { return 8 },
				}
			},
			newRetryer(defaultRetryDelayFn),
		); err != nil || atomic.AddInt64(&first, 1) < 2 {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("Refresh no slots cluster", func(t *testing.T) {
		if _, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn {
				return &mockConn{
					DoFn: func(cmd Completed) RedisResult {
						return newResult(slicemsg('*', []RedisMessage{}), nil)
					},
				}
			},
			newRetryer(defaultRetryDelayFn),
		); err != nil {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("Refresh no shards cluster", func(t *testing.T) {
		if _, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn {
				return &mockConn{
					DoFn: func(cmd Completed) RedisResult {
						return newResult(slicemsg('*', []RedisMessage{}), nil)
					},
					VersionFn: func() int { return 8 },
				}
			},
			newRetryer(defaultRetryDelayFn),
		); err != nil {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("Refresh cluster of 1 node without knowing its own ip", func(t *testing.T) {
		getClient := func(version int) (client *clusterClient, err error) {
			return newClusterClient(
				&ClientOption{InitAddress: []string{"127.0.4.1:4"}},
				func(dst string, opt *ClientOption) conn {
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
				},
				newRetryer(defaultRetryDelayFn),
			)
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
			client, err := getClient(8)
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
			client, err := newClusterClient(
				&ClientOption{InitAddress: []string{"127.0.1.1:1", "127.0.2.1:2"}},
				func(dst string, opt *ClientOption) conn {
					return &mockConn{
						DoFn: func(cmd Completed) RedisResult {
							if atomic.LoadInt64(&first) == 1 {
								return singleSlotResp2
							}
							return slotsResp
						},
					}
				},
				newRetryer(defaultRetryDelayFn),
			)
			if err != nil {
				t.Fatalf("unexpected err %v", err)
			}
			testFunc(t, client, &first)
		})

		t.Run("shards", func(t *testing.T) {
			var first int64
			client, err := newClusterClient(
				&ClientOption{InitAddress: []string{"127.0.1.1:1", "127.0.2.1:2"}},
				func(dst string, opt *ClientOption) conn {
					return &mockConn{
						DoFn: func(cmd Completed) RedisResult {
							if atomic.LoadInt64(&first) == 1 {
								return singleShardResp2
							}
							return shardsResp
						},
						VersionFn: func() int { return 8 },
					}
				},
				newRetryer(defaultRetryDelayFn),
			)
			if err != nil {
				t.Fatalf("unexpected err %v", err)
			}
			testFunc(t, client, &first)
		})
	})

	t.Run("Refresh InitAddress which is not in CLUSTER SLOTS / CLUSTER SHARDS should be hidden", func(t *testing.T) {
		testFunc := func(t *testing.T, client *clusterClient, num *int64) {
			nodesWithHidden := client.nodes()
			sort.Strings(nodesWithHidden)
			if len(nodesWithHidden) != 4 ||
				nodesWithHidden[0] != "127.0.0.1:0" ||
				nodesWithHidden[1] != "127.0.1.1:1" ||
				nodesWithHidden[2] != "127.0.2.1:2" ||
				nodesWithHidden[3] != "redis.example.com" {
				t.Fatalf("unexpected nodes %v", nodesWithHidden)
			}

			nodes := client.Nodes()
			_, ok := nodes["127.0.0.1:0"]
			_, ok2 := nodes["127.0.1.1:1"]
			if len(nodes) != 2 || !ok || !ok2 {
				t.Fatalf("unexpected nodes %v", nodes)
			}

			atomic.AddInt64(num, 1)

			if err := client.refresh(context.Background()); err != nil {
				t.Fatalf("unexpected err %v", err)
			}

			nodesWithHidden = client.nodes()
			sort.Strings(nodesWithHidden)
			if len(nodesWithHidden) != 4 ||
				nodesWithHidden[0] != "127.0.1.1:1" ||
				nodesWithHidden[1] != "127.0.2.1:2" ||
				nodesWithHidden[2] != "127.0.3.1:3" ||
				nodesWithHidden[3] != "redis.example.com" {
				t.Fatalf("unexpected nodes %v", nodesWithHidden)
			}

			nodes = client.Nodes()
			_, ok = nodes["127.0.3.1:3"]
			if len(nodes) != 1 || !ok {
				t.Fatalf("unexpected nodes %v", nodes)
			}
		}

		t.Run("slots", func(t *testing.T) {
			var first int64
			client, err := newClusterClient(
				&ClientOption{InitAddress: []string{"127.0.1.1:1", "127.0.2.1:2", "redis.example.com"}},
				func(dst string, opt *ClientOption) conn {
					return &mockConn{
						DoFn: func(cmd Completed) RedisResult {
							if atomic.LoadInt64(&first) == 1 {
								return singleSlotResp2
							}
							return slotsResp
						},
					}
				},
				newRetryer(defaultRetryDelayFn),
			)
			if err != nil {
				t.Fatalf("unexpected err %v", err)
			}
			testFunc(t, client, &first)
		})

		t.Run("shards", func(t *testing.T) {
			var first int64
			client, err := newClusterClient(
				&ClientOption{InitAddress: []string{"127.0.1.1:1", "127.0.2.1:2", "redis.example.com"}},
				func(dst string, opt *ClientOption) conn {
					return &mockConn{
						DoFn: func(cmd Completed) RedisResult {
							if atomic.LoadInt64(&first) == 1 {
								return singleShardResp2
							}
							return shardsResp
						},
						VersionFn: func() int { return 8 },
					}
				},
				newRetryer(defaultRetryDelayFn),
			)
			if err != nil {
				t.Fatalf("unexpected err %v", err)
			}
			testFunc(t, client, &first)
		})
	})

	t.Run("Shards tls", func(t *testing.T) {
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{"127.0.0.1:0"}, TLSConfig: &tls.Config{}},
			func(dst string, opt *ClientOption) conn {
				return &mockConn{
					DoFn: func(cmd Completed) RedisResult {
						return shardsRespTls
					},
					VersionFn: func() int { return 8 },
				}
			},
			newRetryer(defaultRetryDelayFn),
		)
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
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}

		if client.wslots[0] != client.conns["127.0.0.1:0"].conn {
			t.Fatalf("unexpected node assigned to pslot 0")
		}
		if client.wslots[8192] != client.conns["127.0.0.1:0"].conn {
			t.Fatalf("unexpected node assigned to pslot 8192")
		}
		if client.wslots[8193] != client.conns["127.0.1.1:0"].conn {
			t.Fatalf("unexpected node assigned to pslot 8193")
		}
		if client.wslots[16383] != client.conns["127.0.1.1:0"].conn {
			t.Fatalf("unexpected node assigned to pslot 16383")
		}
		if client.rslots[0][0].conn != client.conns["127.0.0.1:0"].conn {
			t.Fatalf("unexpected node assigned to rslot 0")
		}
		if client.rslots[8192][0].conn != client.conns["127.0.0.1:0"].conn {
			t.Fatalf("unexpected node assigned to rslot 8192")
		}
		if client.rslots[8193][0].conn != client.conns["127.0.1.1:0"].conn {
			t.Fatalf("unexpected node assigned to rslot 8193")
		}
		if client.rslots[16383][0].conn != client.conns["127.0.1.1:0"].conn {
			t.Fatalf("unexpected node assigned to rslot 16383")
		}
	})

	t.Run("Refresh cluster which has multi nodes per shard with SendToReplica option", func(t *testing.T) {
		primaryNodeConn := &mockConn{
			DoFn: func(cmd Completed) RedisResult {
				if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
					return slotsMultiResp
				}
				return RedisResult{
					err: errors.New("unexpected call"),
				}
			},
		}
		replicaNodeConn := &mockConn{
			DoFn: func(cmd Completed) RedisResult {
				return RedisResult{
					err: errors.New("unexpected call"),
				}
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
				if dst == "127.0.0.1:0" || dst == "127.0.2.1:0" {
					if opt.ReplicaOnly {
						t.Fatalf("unexpected replicaOnly option in primary node")
					}
					return primaryNodeConn
				} else {
					if !opt.ReplicaOnly {
						t.Fatalf("unexpected replicaOnly option in replica node")
					}
					return replicaNodeConn
				}
			},
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}

		if client.wslots[0] != client.conns["127.0.0.1:0"].conn {
			t.Fatalf("unexpected node assigned to pslot 0")
		}
		if client.wslots[8192] != client.conns["127.0.0.1:0"].conn {
			t.Fatalf("unexpected node assigned to pslot 8192")
		}
		if client.wslots[8193] != client.conns["127.0.2.1:0"].conn {
			t.Fatalf("unexpected node assigned to pslot 8193")
		}
		if client.wslots[16383] != client.conns["127.0.2.1:0"].conn {
			t.Fatalf("unexpected node assigned to pslot 16383")
		}
		if client.rslots[0][0].conn != client.conns["127.0.1.1:1"].conn {
			t.Fatalf("unexpected node assigned to rslot 0")
		}
		if client.rslots[8192][0].conn != client.conns["127.0.1.1:1"].conn {
			t.Fatalf("unexpected node assigned to rslot 8192")
		}
		if client.rslots[8193][0].conn != client.conns["127.0.3.1:1"].conn {
			t.Fatalf("unexpected node assigned to rslot 8193")
		}
		if client.rslots[16383][0].conn != client.conns["127.0.3.1:1"].conn {
			t.Fatalf("unexpected node assigned to rslot 16383")
		}
	})

	t.Run("Negative ShardRefreshInterval", func(t *testing.T) {
		_, err := newClusterClient(
			&ClientOption{
				InitAddress: []string{"127.0.0.1:0"},
				ClusterOption: ClusterOption{
					ShardsRefreshInterval: -1 * time.Millisecond,
				},
			},
			func(dst string, opt *ClientOption) conn {
				return &mockConn{
					DoFn: func(cmd Completed) RedisResult {
						return singleSlotResp
					},
				}
			},
			newRetryer(defaultRetryDelayFn),
		)
		if !errors.Is(err, ErrInvalidShardsRefreshInterval) {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("Refresh cluster which has only primary node per shard with ReplicaSelector option", func(t *testing.T) {
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
				ReplicaSelector: func(slot uint16, replicas []NodeInfo) int {
					return 0
				},
			},
			func(dst string, opt *ClientOption) conn {
				copiedM := *m
				return &copiedM
			},
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}

		if client.wslots[0] != client.conns["127.0.0.1:0"].conn {
			t.Fatalf("unexpected node assigned to pslot 0")
		}
		if client.wslots[8192] != client.conns["127.0.0.1:0"].conn {
			t.Fatalf("unexpected node assigned to pslot 8192")
		}
		if client.wslots[8193] != client.conns["127.0.1.1:0"].conn {
			t.Fatalf("unexpected node assigned to pslot 8193")
		}
		if client.wslots[16383] != client.conns["127.0.1.1:0"].conn {
			t.Fatalf("unexpected node assigned to pslot 16383")
		}
		if client.rslots[0][0].conn != client.conns["127.0.0.1:0"].conn {
			t.Fatalf("unexpected node assigned to rslot 0")
		}
		if client.rslots[8192][0].conn != client.conns["127.0.0.1:0"].conn {
			t.Fatalf("unexpected node assigned to rslot 8192")
		}
		if client.rslots[8193][0].conn != client.conns["127.0.1.1:0"].conn {
			t.Fatalf("unexpected node assigned to rslot 8193")
		}
		if client.rslots[16383][0].conn != client.conns["127.0.1.1:0"].conn {
			t.Fatalf("unexpected node assigned to rslot 16383")
		}
	})

	t.Run("Refresh cluster which has multi replicas per shard with ReplicaSelector option. Returned index is within range", func(t *testing.T) {
		primaryNodeConn := &mockConn{
			DoFn: func(cmd Completed) RedisResult {
				if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
					return slotsMultiRespWithMultiReplicas
				}
				return RedisResult{
					err: errors.New("unexpected call"),
				}
			},
		}
		replicaNodeConn1 := &mockConn{
			DoFn: func(cmd Completed) RedisResult {
				return RedisResult{
					err: errors.New("unexpected call"),
				}
			},
		}
		replicaNodeConn2 := &mockConn{
			DoFn: func(cmd Completed) RedisResult {
				return RedisResult{
					err: errors.New("unexpected call"),
				}
			},
		}
		replicaNodeConn3 := &mockConn{
			DoFn: func(cmd Completed) RedisResult {
				return RedisResult{
					err: errors.New("unexpected call"),
				}
			},
		}

		client, err := newClusterClient(
			&ClientOption{
				InitAddress: []string{"127.0.0.1:0"},
				SendToReplicas: func(cmd Completed) bool {
					return true
				},
				ReplicaSelector: func(slot uint16, replicas []NodeInfo) int {
					return 1
				},
			},
			func(dst string, opt *ClientOption) conn {
				switch {
				case dst == "127.0.0.2:1" || dst == "127.0.1.2:1":
					return replicaNodeConn1
				case dst == "127.0.0.3:2" || dst == "127.0.1.3:2":
					return replicaNodeConn2
				case dst == "127.0.0.4:3" || dst == "127.0.1.4:3":
					return replicaNodeConn3
				default:
					return primaryNodeConn
				}
			},
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}

		if client.wslots[0] != primaryNodeConn {
			t.Fatalf("unexpected node assigned to pslot 0")
		}
		if client.wslots[8192] != primaryNodeConn {
			t.Fatalf("unexpected node assigned to pslot 8192")
		}
		if client.wslots[8193] != primaryNodeConn {
			t.Fatalf("unexpected node assigned to pslot 8193")
		}
		if client.wslots[16383] != primaryNodeConn {
			t.Fatalf("unexpected node assigned to pslot 16383")
		}
		if client.rslots[0][0].conn != replicaNodeConn2 {
			t.Fatalf("unexpected node assigned to rslot 0")
		}
		if client.rslots[8192][0].conn != replicaNodeConn2 {
			t.Fatalf("unexpected node assigned to rslot 8192")
		}
		if client.rslots[8193][0].conn != replicaNodeConn2 {
			t.Fatalf("unexpected node assigned to rslot 8193")
		}
		if client.rslots[16383][0].conn != replicaNodeConn2 {
			t.Fatalf("unexpected node assigned to rslot 16383")
		}
	})

	t.Run("Refresh cluster which has multi replicas per shard with ReplicaSelector option. Returned index is out of range", func(t *testing.T) {
		primaryNodeConn := &mockConn{
			DoFn: func(cmd Completed) RedisResult {
				if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
					return slotsMultiRespWithMultiReplicas
				}
				return RedisResult{
					err: errors.New("unexpected call"),
				}
			},
		}
		replicaNodeConn1 := &mockConn{
			DoFn: func(cmd Completed) RedisResult {
				return RedisResult{
					err: errors.New("unexpected call"),
				}
			},
		}
		replicaNodeConn2 := &mockConn{
			DoFn: func(cmd Completed) RedisResult {
				return RedisResult{
					err: errors.New("unexpected call"),
				}
			},
		}
		replicaNodeConn3 := &mockConn{
			DoFn: func(cmd Completed) RedisResult {
				return RedisResult{
					err: errors.New("unexpected call"),
				}
			},
		}

		client, err := newClusterClient(
			&ClientOption{
				InitAddress: []string{"127.0.0.1:0"},
				SendToReplicas: func(cmd Completed) bool {
					return true
				},
				ReplicaSelector: func(slot uint16, replicas []NodeInfo) int {
					return -1
				},
			},
			func(dst string, opt *ClientOption) conn {
				switch {
				case dst == "127.0.0.2:1" || dst == "127.0.1.2:1":
					return replicaNodeConn1
				case dst == "127.0.0.3:2" || dst == "127.0.1.3:2":
					return replicaNodeConn2
				case dst == "127.0.0.4:3" || dst == "127.0.1.4:3":
					return replicaNodeConn3
				default:
					return primaryNodeConn
				}
			},
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}

		if client.wslots[0] != primaryNodeConn {
			t.Fatalf("unexpected node assigned to pslot 0")
		}
		if client.wslots[8192] != primaryNodeConn {
			t.Fatalf("unexpected node assigned to pslot 8192")
		}
		if client.wslots[8193] != primaryNodeConn {
			t.Fatalf("unexpected node assigned to pslot 8193")
		}
		if client.wslots[16383] != primaryNodeConn {
			t.Fatalf("unexpected node assigned to pslot 16383")
		}
		if client.rslots[0][0].conn != primaryNodeConn {
			t.Fatalf("unexpected node assigned to rslot 0")
		}
		if client.rslots[8192][0].conn != primaryNodeConn {
			t.Fatalf("unexpected node assigned to rslot 8192")
		}
		if client.rslots[8193][0].conn != primaryNodeConn {
			t.Fatalf("unexpected node assigned to rslot 8193")
		}
		if client.rslots[16383][0].conn != primaryNodeConn {
			t.Fatalf("unexpected node assigned to rslot 16383")
		}
	})

	t.Run("Refresh cluster which has multi replicas with az", func(t *testing.T) {
		primaryNodeConn := &mockConn{
			DoFn: func(cmd Completed) RedisResult {
				if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
					return slotsMultiRespWithMultiReplicas
				}
				return RedisResult{
					err: errors.New("unexpected call"),
				}
			},
			AZFn: func() string {
				return "us-west-1a"
			},
		}
		replicaNodeConn1 := &mockConn{
			DoFn: func(cmd Completed) RedisResult {
				return RedisResult{
					err: errors.New("unexpected call"),
				}
			},
			AZFn: func() string {
				return "us-west-1a"
			},
		}
		replicaNodeConn2 := &mockConn{
			DoFn: func(cmd Completed) RedisResult {
				return RedisResult{
					err: errors.New("unexpected call"),
				}
			},
			AZFn: func() string {
				return "us-west-1b"
			},
		}
		replicaNodeConn3 := &mockConn{
			DoFn: func(cmd Completed) RedisResult {
				return RedisResult{
					err: errors.New("unexpected call"),
				}
			},
			AZFn: func() string {
				return "us-west-1c"
			},
		}

		client, err := newClusterClient(
			&ClientOption{
				InitAddress: []string{"127.0.0.1:0"},
				SendToReplicas: func(cmd Completed) bool {
					return true
				},
				ReplicaSelector: func(slot uint16, replicas []NodeInfo) int {
					for i, replica := range replicas {
						if replica.AZ == "us-west-1b" {
							return i
						}
					}
					return -1
				},
				EnableReplicaAZInfo: true,
			},
			func(dst string, opt *ClientOption) conn {
				switch {
				case dst == "127.0.0.2:1" || dst == "127.0.1.2:1":
					return replicaNodeConn1
				case dst == "127.0.0.3:2" || dst == "127.0.1.3:2":
					return replicaNodeConn2
				case dst == "127.0.0.4:3" || dst == "127.0.1.4:3":
					return replicaNodeConn3
				default:
					return primaryNodeConn
				}
			},
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}

		if client.wslots[0] != primaryNodeConn {
			t.Fatalf("unexpected node assigned to pslot 0")
		}
		if client.wslots[8192] != primaryNodeConn {
			t.Fatalf("unexpected node assigned to pslot 8192")
		}
		if client.wslots[8193] != primaryNodeConn {
			t.Fatalf("unexpected node assigned to pslot 8193")
		}
		if client.wslots[16383] != primaryNodeConn {
			t.Fatalf("unexpected node assigned to pslot 16383")
		}
		if client.rslots[0][0].conn != replicaNodeConn2 {
			t.Fatalf("unexpected node assigned to rslot 0")
		}
		if client.rslots[8192][0].conn != replicaNodeConn2 {
			t.Fatalf("unexpected node assigned to rslot 8192")
		}
		if client.rslots[8193][0].conn != replicaNodeConn2 {
			t.Fatalf("unexpected node assigned to rslot 8193")
		}
		if client.rslots[16383][0].conn != replicaNodeConn2 {
			t.Fatalf("unexpected node assigned to rslot 16383")
		}
	})

	t.Run("Refresh cluster which has multi replicas with ReadNodeSelector option", func(t *testing.T) {
		primaryNodeConn := &mockConn{
			DoFn: func(cmd Completed) RedisResult {
				if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
					return slotsMultiRespWithMultiReplicas
				}
				return RedisResult{
					err: errors.New("unexpected call"),
				}
			},
		}
		replicaNodeConn1 := &mockConn{
			DoFn: func(cmd Completed) RedisResult {
				return RedisResult{
					err: errors.New("unexpected call"),
				}
			},
		}
		replicaNodeConn2 := &mockConn{
			DoFn: func(cmd Completed) RedisResult {
				return RedisResult{
					err: errors.New("unexpected call"),
				}
			},
		}
		replicaNodeConn3 := &mockConn{
			DoFn: func(cmd Completed) RedisResult {
				return RedisResult{
					err: errors.New("unexpected call"),
				}
			},
		}

		client, err := newClusterClient(
			&ClientOption{
				InitAddress: []string{"127.0.0.1:0"},
				SendToReplicas: func(cmd Completed) bool {
					return true
				},
				ReadNodeSelector: func(slot uint16, nodes []NodeInfo) int {
					return 0
				},
			},
			func(dst string, opt *ClientOption) conn {
				switch {
				case dst == "127.0.0.2:1" || dst == "127.0.1.2:1":
					return replicaNodeConn1
				case dst == "127.0.0.3:2" || dst == "127.0.1.3:2":
					return replicaNodeConn2
				case dst == "127.0.0.4:3" || dst == "127.0.1.4:3":
					return replicaNodeConn3
				default:
					return primaryNodeConn
				}
			},
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}

		for i := 0; i < 16384; i++ {
			if client.wslots[i] != primaryNodeConn {
				t.Fatalf("unexpected node assigned to pslot %d", i)
			}
		}

		for i := 0; i < 16384; i++ {
			if len(client.rslots[i]) != 4 {
				t.Fatalf("unexpected number of replicas for rslot %d, expected 4, got %d", i, len(client.rslots[i]))
			}
		}

		for i := 0; i < 16384; i++ {
			for j := 0; j < 4; j++ {
				if j == 0 && client.rslots[i][j].conn != primaryNodeConn {
					t.Fatalf("unexpected node assigned to rslot %d at index %d", i, j)
				} else if j == 1 && client.rslots[i][j].conn != replicaNodeConn1 {
					t.Fatalf("unexpected node assigned to rslot %d at index %d", i, j)
				} else if j == 2 && client.rslots[i][j].conn != replicaNodeConn2 {
					t.Fatalf("unexpected node assigned to rslot %d at index %d", i, j)
				} else if j == 3 && client.rslots[i][j].conn != replicaNodeConn3 {
					t.Fatalf("unexpected node assigned to rslot %d at index %d", i, j)
				}
			}
		}
	})
}

//gocyclo:ignore
func TestClusterClient(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	m := &mockConn{
		DoFn: func(cmd Completed) RedisResult {
			if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
				return slotsMultiResp
			}
			return RedisResult{}
		},
		DoStreamFn: func(cmd Completed) RedisResultStream {
			return RedisResultStream{e: errors.New(cmd.Commands()[1])}
		},
		DoMultiFn: func(multi ...Completed) *redisresults {
			resps := make([]RedisResult, len(multi))
			for i, cmd := range multi {
				resps[i] = newResult(strmsg('+', strings.Join(cmd.Commands(), " ")), nil)
			}
			return &redisresults{s: resps}
		},
		DoMultiStreamFn: func(cmd ...Completed) MultiRedisResultStream {
			return MultiRedisResultStream{e: errors.New(cmd[0].Commands()[1])}
		},
		DoMultiCacheFn: func(multi ...CacheableTTL) *redisresults {
			resps := make([]RedisResult, len(multi))
			for i, cmd := range multi {
				resps[i] = newResult(strmsg('+', strings.Join(cmd.Cmd.Commands(), " ")), nil)
			}
			return &redisresults{s: resps}
		},
		DoOverride: map[string]func(cmd Completed) RedisResult{
			"GET Do": func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "Do"), nil)
			},
			"INFO": func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "Info"), nil)
			},
		},
		DoCacheOverride: map[string]func(cmd Cacheable, ttl time.Duration) RedisResult{
			"GET DoCache": func(cmd Cacheable, ttl time.Duration) RedisResult {
				return newResult(strmsg('+', "DoCache"), nil)
			},
		},
	}

	client, err := newClusterClient(
		&ClientOption{InitAddress: []string{"127.0.0.1:0"}},
		func(dst string, opt *ClientOption) conn {
			return m
		},
		newRetryer(defaultRetryDelayFn),
	)
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

	t.Run("Mode", func(t *testing.T) {
		if client.Mode() != ClientModeCluster {
			t.Fatalf("unexpected mode %v", client.Mode())
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

	t.Run("Delegate DoStream", func(t *testing.T) {
		c := client.B().Get().Key("Do").Build()
		if s := client.DoStream(context.Background(), c); s.Error().Error() != "Do" {
			t.Fatalf("unexpected response %v", s.Error())
		}
	})

	t.Run("Delegate DoMulti Empty", func(t *testing.T) {
		if resps := client.DoMulti(context.Background()); resps != nil {
			t.Fatalf("unexpected response %v", resps)
		}
	})

	t.Run("Delegate DoMultiStream Empty", func(t *testing.T) {
		if s := client.DoMultiStream(context.Background()); s.Error() != io.EOF {
			t.Fatalf("unexpected response %v", err)
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

	t.Run("Delegate DoMultiStream Single Slot", func(t *testing.T) {
		c1 := client.B().Get().Key("K1{a}").Build()
		c2 := client.B().Get().Key("K2{a}").Build()
		if s := client.DoMultiStream(context.Background(), c1, c2); s.Error().Error() != "K1{a}" {
			t.Fatalf("unexpected response %v", s.Error())
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

	t.Run("Delegate DoMultiStream Single Slot + Init Slot", func(t *testing.T) {
		c1 := client.B().Info().Section("ANY").Build()
		c2 := client.B().Get().Key("K1{a}").Build()
		if s := client.DoMultiStream(context.Background(), c1, c2); s.Error().Error() != "ANY" {
			t.Fatalf("unexpected response %v", s.Error())
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

	t.Run("Delegate DoMultiStream Cross Slot + Init Slot", func(t *testing.T) {
		defer func() {
			if err := recover(); !strings.Contains(err.(string), "across multiple slots") {
				t.Errorf("DoMulti should panic if Cross Slot + Init Slot")
			}
		}()
		c1 := client.B().Get().Key("K1{a}").Build()
		c2 := client.B().Get().Key("K1{b}").Build()
		c3 := client.B().Info().Build()
		client.DoMultiStream(context.Background(), c1, c2, c3)
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
		select {
		case _, ok := <-client.stopCh:
			if ok {
				t.Fatalf("stopCh should be closed")
			}
		}
	})

	t.Run("Dedicated Err, but no retry", func(t *testing.T) {
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
						newResult(strmsg('+', "OK"), nil),
						newResult(strmsg('+', "OK"), nil),
						newResult(slicemsg('*', []RedisMessage{strmsg('+', "a")}), nil),
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
				return newResult(strmsg('+', "Delegate"), nil)
			},
			DoMultiFn: func(cmd ...Completed) *redisresults {
				if len(cmd) == 4 {
					return &redisresults{s: []RedisResult{
						newResult(strmsg('+', "OK"), nil),
						newResult(strmsg('+', "OK"), nil),
						newResult(strmsg('+', "OK"), nil),
						newResult(slicemsg('*', []RedisMessage{
							strmsg('+', "Delegate0"),
							strmsg('+', "Delegate1"),
						}), nil),
					}}
				}
				return &redisresults{s: []RedisResult{
					newResult(strmsg('+', "Delegate0"), nil),
					newResult(strmsg('+', "Delegate1"), nil),
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
			)[3].val.values() {
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
			t.Fatalf("Dedicated doesn't put back the wire")
		}
		if !closed {
			t.Fatalf("Dedicated doesn't delegate Close")
		}
	})

	t.Run("Dedicated Delegate", func(t *testing.T) {
		closed := false
		w := &mockWire{
			DoFn: func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "Delegate"), nil)
			},
			DoMultiFn: func(cmd ...Completed) *redisresults {
				if len(cmd) == 4 {
					return &redisresults{s: []RedisResult{
						newResult(strmsg('+', "OK"), nil),
						newResult(strmsg('+', "OK"), nil),
						newResult(strmsg('+', "OK"), nil),
						newResult(slicemsg('*', []RedisMessage{
							strmsg('+', "Delegate0"),
							strmsg('+', "Delegate1"),
						}), nil),
					}}
				}
				return &redisresults{s: []RedisResult{
					newResult(strmsg('+', "Delegate0"), nil),
					newResult(strmsg('+', "Delegate1"), nil),
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
		)[3].val.values() {
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
			t.Fatalf("Dedicated doesn't put back the wire")
		}
		if !closed {
			t.Fatalf("Dedicated doesn't delegate Close")
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

	t.Run("Dedicate ErrDedicatedClientRecycled after released", func(t *testing.T) {
		check := func(err error) {
			if !errors.Is(err, ErrDedicatedClientRecycled) {
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
					resp := c.Do(context.Background(), c.B().Get().Key("k").Build())
					check(resp.Error())
				},
				func() {
					resp := c.DoMulti(context.Background(), c.B().Get().Key("k").Build())
					for _, r := range resp {
						check(r.Error())
					}
				},
				func() {
					err := c.Receive(context.Background(), c.B().Subscribe().Channel("k").Build(), func(msg PubSubMessage) {})
					check(err)
				},
				func() {
					ch := c.SetPubSubHooks(PubSubHooks{})
					check(<-ch)
				},
			} {
				fn()
			}
		}
	})
}

//gocyclo:ignore
func TestClusterClient_SendToOnlyPrimaryNodes(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())

	primaryNodeConn := &mockConn{
		DoOverride: map[string]func(cmd Completed) RedisResult{
			"CLUSTER SLOTS": func(cmd Completed) RedisResult {
				return slotsMultiResp
			},
			"GET Do": func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "GET Do"), nil)
			},
			"GET K1{a}": func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "GET K1{a}"), nil)
			},
			"GET K2{a}": func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "GET K2{a}"), nil)
			},
			"INFO": func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "INFO"), nil)
			},
		},
		DoMultiFn: func(multi ...Completed) *redisresults {
			resps := make([]RedisResult, len(multi))
			for i, cmd := range multi {
				resps[i] = newResult(strmsg('+', strings.Join(cmd.Commands(), " ")), nil)
			}
			return &redisresults{s: resps}
		},
		DoCacheOverride: map[string]func(cmd Cacheable, ttl time.Duration) RedisResult{
			"GET DoCache": func(cmd Cacheable, ttl time.Duration) RedisResult {
				return newResult(strmsg('+', "GET DoCache"), nil)
			},
			"GET K1{a}": func(cmd Cacheable, ttl time.Duration) RedisResult {
				return newResult(strmsg('+', "GET K1{a}"), nil)
			},
			"GET K2{a}": func(cmd Cacheable, ttl time.Duration) RedisResult {
				return newResult(strmsg('+', "GET K2{a}"), nil)
			},
		},
		DoMultiCacheFn: func(multi ...CacheableTTL) *redisresults {
			resps := make([]RedisResult, len(multi))
			for i, cmd := range multi {
				resps[i] = newResult(strmsg('+', strings.Join(cmd.Cmd.Commands(), " ")), nil)
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
		newRetryer(defaultRetryDelayFn),
	)
	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}

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
						newResult(strmsg('+', "OK"), nil),
						newResult(strmsg('+', "OK"), nil),
						newResult(slicemsg('*', []RedisMessage{strmsg('+', "a")}), nil),
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
				return newResult(strmsg('+', "Delegate"), nil)
			},
			DoMultiFn: func(cmd ...Completed) *redisresults {
				if len(cmd) == 4 {
					return &redisresults{s: []RedisResult{
						newResult(strmsg('+', "OK"), nil),
						newResult(strmsg('+', "OK"), nil),
						newResult(strmsg('+', "OK"), nil),
						newResult(slicemsg('*', []RedisMessage{
							strmsg('+', "Delegate0"),
							strmsg('+', "Delegate1"),
						}), nil),
					}}
				}
				return &redisresults{s: []RedisResult{
					newResult(strmsg('+', "Delegate0"), nil),
					newResult(strmsg('+', "Delegate1"), nil),
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
			)[3].val.values() {
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
			t.Fatalf("Dedicated doesn't put back the wire")
		}
		if !closed {
			t.Fatalf("Dedicated doesn't delegate Close")
		}
	})

	t.Run("Dedicate", func(t *testing.T) {
		closed := false
		w := &mockWire{
			DoFn: func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "Delegate"), nil)
			},
			DoMultiFn: func(cmd ...Completed) *redisresults {
				if len(cmd) == 4 {
					return &redisresults{s: []RedisResult{
						newResult(strmsg('+', "OK"), nil),
						newResult(strmsg('+', "OK"), nil),
						newResult(strmsg('+', "OK"), nil),
						newResult(slicemsg('*', []RedisMessage{
							strmsg('+', "Delegate0"),
							strmsg('+', "Delegate1"),
						}), nil),
					}}
				}
				return &redisresults{s: []RedisResult{
					newResult(strmsg('+', "Delegate0"), nil),
					newResult(strmsg('+', "Delegate1"), nil),
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
		)[3].val.values() {
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
			t.Fatalf("Dedicated doesn't put back the wire")
		}
		if !closed {
			t.Fatalf("Dedicated doesn't delegate Close")
		}
	})
}

//gocyclo:ignore
func TestClusterClient_SendToOnlyReplicaNodes(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())

	primaryNodeConn := &mockConn{
		DoOverride: map[string]func(cmd Completed) RedisResult{
			"CLUSTER SLOTS": func(cmd Completed) RedisResult {
				return slotsMultiResp
			},
			"INFO": func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "INFO"), nil)
			},
			"GET K1{a}": func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "GET K1{a}"), nil)
			},
		},
	}
	replicaNodeConn := &mockConn{
		DoMultiFn: func(multi ...Completed) *redisresults {
			resps := make([]RedisResult, len(multi))
			for i, cmd := range multi {
				resps[i] = newResult(strmsg('+', strings.Join(cmd.Commands(), " ")), nil)
			}
			return &redisresults{s: resps}
		},
		DoMultiCacheFn: func(multi ...CacheableTTL) *redisresults {
			resps := make([]RedisResult, len(multi))
			for i, cmd := range multi {
				resps[i] = newResult(strmsg('+', strings.Join(cmd.Cmd.Commands(), " ")), nil)
			}
			return &redisresults{s: resps}
		},
		DoOverride: map[string]func(cmd Completed) RedisResult{
			"GET Do": func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "GET Do"), nil)
			},
			"GET K1{a}": func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "GET K1{a}"), nil)
			},
			"GET K2{a}": func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "GET K2{a}"), nil)
			},
		},
		DoCacheOverride: map[string]func(cmd Cacheable, ttl time.Duration) RedisResult{
			"GET DoCache": func(cmd Cacheable, ttl time.Duration) RedisResult {
				return newResult(strmsg('+', "GET DoCache"), nil)
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
		newRetryer(defaultRetryDelayFn),
	)
	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}

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
						newResult(strmsg('+', "OK"), nil),
						newResult(strmsg('+', "OK"), nil),
						newResult(slicemsg('*', []RedisMessage{strmsg('+', "a")}), nil),
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
		replicaNodeConn.AcquireFn = func() wire {
			return w
		} // Subscribe can work on replicas
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
				return newResult(strmsg('+', "Delegate"), nil)
			},
			DoMultiFn: func(cmd ...Completed) *redisresults {
				if len(cmd) == 4 {
					return &redisresults{s: []RedisResult{
						newResult(strmsg('+', "OK"), nil),
						newResult(strmsg('+', "OK"), nil),
						newResult(strmsg('+', "OK"), nil),
						newResult(slicemsg('*', []RedisMessage{
							strmsg('+', "Delegate0"),
							strmsg('+', "Delegate1"),
						}), nil),
					}}
				}
				return &redisresults{s: []RedisResult{
					newResult(strmsg('+', "Delegate0"), nil),
					newResult(strmsg('+', "Delegate1"), nil),
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
			)[3].val.values() {
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
			t.Fatalf("Dedicated doesn't put back the wire")
		}
		if !closed {
			t.Fatalf("Dedicated doesn't delegate Close")
		}
	})

	t.Run("Dedicate", func(t *testing.T) {
		closed := false
		w := &mockWire{
			DoFn: func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "Delegate"), nil)
			},
			DoMultiFn: func(cmd ...Completed) *redisresults {
				if len(cmd) == 4 {
					return &redisresults{s: []RedisResult{
						newResult(strmsg('+', "OK"), nil),
						newResult(strmsg('+', "OK"), nil),
						newResult(strmsg('+', "OK"), nil),
						newResult(slicemsg('*', []RedisMessage{
							strmsg('+', "Delegate0"),
							strmsg('+', "Delegate1"),
						}), nil),
					}}
				}
				return &redisresults{s: []RedisResult{
					newResult(strmsg('+', "Delegate0"), nil),
					newResult(strmsg('+', "Delegate1"), nil),
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
		)[3].val.values() {
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
			t.Fatalf("Dedicated doesn't put back the wire")
		}
		if !closed {
			t.Fatalf("Dedicated doesn't delegate Close")
		}
	})
}

//gocyclo:ignore
func TestClusterClient_SendReadOperationToReplicaNodesWriteOperationToPrimaryNodes(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())

	primaryNodeConn := &mockConn{
		DoOverride: map[string]func(cmd Completed) RedisResult{
			"CLUSTER SLOTS": func(cmd Completed) RedisResult {
				return slotsMultiResp
			},
			"INFO": func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "INFO"), nil)
			},
			"SET Do V": func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "SET Do V"), nil)
			},
			"SET K2{a} V2{a}": func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "SET K2{a} V2{a}"), nil)
			},
		},
		DoMultiFn: func(multi ...Completed) *redisresults {
			resps := make([]RedisResult, len(multi))
			for i, cmd := range multi {
				if strings.HasPrefix(strings.Join(cmd.Commands(), " "), "SET K1") {
					resps[i] = newResult(strmsg('+', strings.Join(cmd.Commands(), " ")), nil)
					continue
				}
				if strings.HasPrefix(strings.Join(cmd.Commands(), " "), "SET K2") {
					resps[i] = newResult(strmsg('+', strings.Join(cmd.Commands(), " ")), nil)
					continue
				}
				if strings.HasPrefix(strings.Join(cmd.Commands(), " "), "MULTI") {
					resps[i] = newResult(strmsg('+', "MULTI"), nil)
					continue
				}
				if strings.HasPrefix(strings.Join(cmd.Commands(), " "), "EXEC") {
					resps[i] = newResult(strmsg('+', "EXEC"), nil)
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
				return newResult(strmsg('+', "GET Do"), nil)
			},
			"GET K1{a}": func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "GET K1{a}"), nil)
			},
			"GET K2{a}": func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "GET K2{a}"), nil)
			},
		},
		DoMultiFn: func(multi ...Completed) *redisresults {
			resps := make([]RedisResult, len(multi))
			for i, cmd := range multi {
				if strings.HasPrefix(strings.Join(cmd.Commands(), " "), "GET K1") {
					resps[i] = newResult(strmsg('+', strings.Join(cmd.Commands(), " ")), nil)
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
				return newResult(strmsg('+', "GET DoCache"), nil)
			},
			"GET K1{a}": func(cmd Cacheable, ttl time.Duration) RedisResult {
				return newResult(strmsg('+', "GET K1{a}"), nil)
			},
			"GET K2{a}": func(cmd Cacheable, ttl time.Duration) RedisResult {
				return newResult(strmsg('+', "GET K2{a}"), nil)
			},
		},
		DoMultiCacheFn: func(multi ...CacheableTTL) *redisresults {
			resps := make([]RedisResult, len(multi))
			for i, cmd := range multi {
				if strings.HasPrefix(strings.Join(cmd.Cmd.Commands(), " "), "GET K1") {
					resps[i] = newResult(strmsg('+', strings.Join(cmd.Cmd.Commands(), " ")), nil)
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
				return cmd.IsReadOnly()
			},
		},
		func(dst string, opt *ClientOption) conn {
			if dst == "127.0.0.1:0" || dst == "127.0.2.1:0" { // primary nodes
				return primaryNodeConn
			} else { // replica nodes
				return replicaNodeConn
			}
		},
		newRetryer(defaultRetryDelayFn),
	)
	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}

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
						newResult(strmsg('+', "OK"), nil),
						newResult(strmsg('+', "OK"), nil),
						newResult(slicemsg('*', []RedisMessage{strmsg('+', "a")}), nil),
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
		replicaNodeConn.AcquireFn = func() wire {
			return w
		} // Subscribe can work on replicas
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
				return newResult(strmsg('+', "Delegate"), nil)
			},
			DoMultiFn: func(cmd ...Completed) *redisresults {
				if len(cmd) == 4 {
					return &redisresults{s: []RedisResult{
						newResult(strmsg('+', "OK"), nil),
						newResult(strmsg('+', "OK"), nil),
						newResult(strmsg('+', "OK"), nil),
						newResult(slicemsg('*', []RedisMessage{
							strmsg('+', "Delegate0"),
							strmsg('+', "Delegate1"),
						}), nil),
					}}
				}
				return &redisresults{s: []RedisResult{
					newResult(strmsg('+', "Delegate0"), nil),
					newResult(strmsg('+', "Delegate1"), nil),
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
			)[3].val.values() {
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
			t.Fatalf("Dedicated doesn't put back the wire")
		}
		if !closed {
			t.Fatalf("Dedicated doesn't delegate Close")
		}
	})

	t.Run("Dedicate", func(t *testing.T) {
		closed := false
		w := &mockWire{
			DoFn: func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "Delegate"), nil)
			},
			DoMultiFn: func(cmd ...Completed) *redisresults {
				if len(cmd) == 4 {
					return &redisresults{s: []RedisResult{
						newResult(strmsg('+', "OK"), nil),
						newResult(strmsg('+', "OK"), nil),
						newResult(strmsg('+', "OK"), nil),
						newResult(slicemsg('*', []RedisMessage{
							strmsg('+', "Delegate0"),
							strmsg('+', "Delegate1"),
						}), nil),
					}}
				}
				return &redisresults{s: []RedisResult{
					newResult(strmsg('+', "Delegate0"), nil),
					newResult(strmsg('+', "Delegate1"), nil),
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
		)[3].val.values() {
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
			t.Fatalf("Dedicated doesn't put back the wire")
		}
		if !closed {
			t.Fatalf("Dedicated doesn't delegate Close")
		}
	})
}

//gocyclo:ignore
func TestClusterClient_SendPrimaryNodeOnlyButOneSlotAssigned(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())

	primaryNodeConn := &mockConn{
		DoOverride: map[string]func(cmd Completed) RedisResult{
			"CLUSTER SLOTS": func(cmd Completed) RedisResult {
				return singleSlotResp
			},
			"INFO": func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "INFO"), nil)
			},
		},
		DoMultiFn: func(multi ...Completed) *redisresults {
			resps := make([]RedisResult, len(multi))
			for i, cmd := range multi {
				if strings.HasPrefix(strings.Join(cmd.Commands(), " "), "MULTI") {
					resps[i] = newResult(strmsg('+', "MULTI"), nil)
					continue
				}
				if strings.HasPrefix(strings.Join(cmd.Commands(), " "), "EXEC") {
					resps[i] = newResult(strmsg('+', "EXEC"), nil)
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
		newRetryer(defaultRetryDelayFn),
	)
	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}

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
	defer ShouldNotLeak(SetupLeakDetection())

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
			DoStreamFn: func(cmd Completed) RedisResultStream {
				return RedisResultStream{e: v}
			},
			DoMultiFn: func(multi ...Completed) *redisresults {
				res := make([]RedisResult, len(multi))
				for i := range res {
					res[i] = newErrResult(v)
				}
				return &redisresults{s: res}
			},
			DoMultiStreamFn: func(cmd ...Completed) MultiRedisResultStream {
				return MultiRedisResultStream{e: v}
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
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn { return m },
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if err := client.Do(ctx, client.B().Get().Key("a").Build()).Error(); err != v {
			t.Fatalf("unexpected err %v", err)
		}
		if s := client.DoStream(ctx, client.B().Get().Key("a").Build()); s.Error() != v {
			t.Fatalf("unexpected err %v", s.Error())
		}
		if err := client.DoMulti(ctx, client.B().Get().Key("a").Build())[0].Error(); err != v {
			t.Fatalf("unexpected err %v", err)
		}
		if s := client.DoMultiStream(ctx, client.B().Get().Key("a").Build()); s.Error() != v {
			t.Fatalf("unexpected err %v", s.Error())
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
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn { return m },
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if err := client.Do(context.Background(), client.B().Get().Key("a").Build()).Error(); err != v {
			t.Fatalf("unexpected err %v", err)
		}
		if s := client.DoStream(context.Background(), client.B().Get().Key("a").Build()); s.Error() != v {
			t.Fatalf("unexpected err %v", s.Error())
		}
		if err := client.DoMulti(context.Background(), client.B().Get().Key("a").Build())[0].Error(); err != v {
			t.Fatalf("unexpected err %v", err)
		}
		if s := client.DoMultiStream(context.Background(), client.B().Get().Key("a").Build()); s.Error() != v {
			t.Fatalf("unexpected err %v", s.Error())
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
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn { return m },
			newRetryer(defaultRetryDelayFn),
		)
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
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn { return m },
			newRetryer(defaultRetryDelayFn),
		)
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
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn { return m },
			newRetryer(defaultRetryDelayFn),
		)
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
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn {
				atomic.AddInt64(&check, 1)
				return &mockConn{DoFn: func(cmd Completed) RedisResult {
					if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
						return slotsMultiResp
					}
					if atomic.AddInt64(&count, 1) <= 3 {
						return newResult(strmsg('-', "MOVED 0 :0"), nil)
					}
					return newResult(strmsg('+', "b"), nil)
				}}
			},
			newRetryer(defaultRetryDelayFn),
		)
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
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn {
				return &mockConn{DoFn: func(cmd Completed) RedisResult {
					if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
						return slotsMultiResp
					}
					if atomic.AddInt64(&count, 1) <= 3 {
						return newResult(strmsg('-', "MOVED 0 :1"), nil)
					}
					return newResult(strmsg('+', "b"), nil)
				}}
			},
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.Do(context.Background(), client.B().Get().Key("a").Build()).ToString(); err != nil || v != "b" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot moved redirect once", func(t *testing.T) {
		var count int64
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn {
				return &mockConn{
					DoFn: func(cmd Completed) RedisResult {
						if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
							return slotsMultiResp
						}

						if strings.Contains(dst, ":0") {
							atomic.AddInt64(&count, 1)
							return newResult(strmsg('-', "MOVED 0 :2"), nil)
						}

						return newResult(strmsg('+', "b"), nil)
					},
				}
			},
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}

		for i := 0; i < 10; i++ {
			if v, err := client.Do(context.Background(), client.B().Get().Key("a").Build()).ToString(); err != nil || v != "b" {
				t.Fatalf("unexpected resp %v %v", v, err)
			}
		}

		if atomic.LoadInt64(&count) != 1 {
			t.Fatalf("unexpected count %v", count)
		}
	})

	t.Run("slot moved DoMulti (single)", func(t *testing.T) {
		var count int64
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn {
				return &mockConn{DoFn: func(cmd Completed) RedisResult {
					return slotsMultiResp
				}, DoMultiFn: func(multi ...Completed) *redisresults {
					ret := make([]RedisResult, len(multi))
					if atomic.AddInt64(&count, 1) <= 3 {
						for i := range ret {
							ret[i] = newResult(strmsg('-', "MOVED 0 :1"), nil)
						}
						return &redisresults{s: ret}
					}
					for i := range ret {
						ret[i] = newResult(strmsg('+', multi[i].Commands()[1]), nil)
					}
					return &redisresults{s: ret}
				}}
			},
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.DoMulti(context.Background(), client.B().Get().Key("a").Build())[0].ToString(); err != nil || v != "a" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot moved DoMulti (single) redirect once", func(t *testing.T) {
		var count int64
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn {
				return &mockConn{
					DoFn: func(cmd Completed) RedisResult {
						return slotsMultiResp
					},
					DoMultiFn: func(multi ...Completed) *redisresults {
						ret := make([]RedisResult, len(multi))

						if strings.Contains(dst, ":0") {
							atomic.AddInt64(&count, 1)
							for i := range ret {
								ret[i] = newResult(strmsg('-', "MOVED 0 :2"), nil)
							}
							return &redisresults{s: ret}
						}

						for i := range ret {
							ret[i] = newResult(strmsg('+', multi[i].Commands()[1]), nil)
						}
						return &redisresults{s: ret}
					},
				}
			},
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}

		for i := 0; i < 10; i++ {
			if v, err := client.DoMulti(context.Background(), client.B().Get().Key("a").Build())[0].ToString(); err != nil || v != "a" {
				t.Fatalf("unexpected resp %v %v", v, err)
			}
		}

		if atomic.LoadInt64(&count) != 1 {
			t.Fatalf("unexpected count %v", count)
		}
	})

	t.Run("slot moved DoMulti transactions", func(t *testing.T) {
		var count int64
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn {
				return &mockConn{DoFn: func(cmd Completed) RedisResult {
					return slotsMultiResp
				}, DoMultiFn: func(multi ...Completed) *redisresults {
					switch atomic.AddInt64(&count, 1) {
					case 1:
						return &redisresults{s: []RedisResult{
							newResult(strmsg('+', "1"), nil),
							newResult(strmsg('+', "OK"), nil),
							newResult(strmsg('+', "QUEUED"), nil),
							newResult(strmsg('-', "MOVED 0 :1"), nil),
							newResult(strmsg('-', "EXECABORT"), nil),
							newResult(strmsg('+', "4"), nil),
							newResult(strmsg('+', "OK"), nil),
							newResult(strmsg('-', "MOVED 0 :1"), nil),
							newResult(strmsg('+', "QUEUED"), nil),
							newResult(strmsg('-', "EXECABORT"), nil),
							newResult(strmsg('+', "7"), nil),
						}}
					case 2:
						return &redisresults{s: []RedisResult{
							newResult(strmsg('+', "OK"), nil),
							newResult(strmsg('+', "QUEUED"), nil),
							newResult(strmsg('+', "QUEUED"), nil),
							newResult(slicemsg('*', []RedisMessage{
								strmsg('+', "2"),
								strmsg('+', "3"),
							}), nil),
							newResult(strmsg('+', "OK"), nil),
							newResult(strmsg('+', "QUEUED"), nil),
							newResult(strmsg('+', "QUEUED"), nil),
							newResult(slicemsg('*', []RedisMessage{
								strmsg('+', "5"),
								strmsg('+', "6"),
							}), nil),
						}}
					}
					return nil
				}}
			},
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		resps := client.DoMulti(
			context.Background(),
			client.B().Get().Key("1{t}").Build(),
			client.B().Multi().Build(),
			client.B().Get().Key("2{t}").Build(),
			client.B().Get().Key("3{t}").Build(),
			client.B().Exec().Build(),
			client.B().Get().Key("4{t}").Build(),
			client.B().Multi().Build(),
			client.B().Get().Key("5{t}").Build(),
			client.B().Get().Key("6{t}").Build(),
			client.B().Exec().Build(),
			client.B().Get().Key("7{t}").Build(),
		)
		if v, err := resps[0].ToString(); err != nil || v != "1" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[1].ToString(); err != nil || v != "OK" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[2].ToString(); err != nil || v != "QUEUED" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[3].ToString(); err != nil || v != "QUEUED" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[4].AsStrSlice(); err != nil || !reflect.DeepEqual(v, []string{"2", "3"}) {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[5].ToString(); err != nil || v != "4" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[6].ToString(); err != nil || v != "OK" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[7].ToString(); err != nil || v != "QUEUED" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[8].ToString(); err != nil || v != "QUEUED" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[9].AsStrSlice(); err != nil || !reflect.DeepEqual(v, []string{"5", "6"}) {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[10].ToString(); err != nil || v != "7" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot moved DoMulti transactions ASKING", func(t *testing.T) {
		var count int64
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn {
				return &mockConn{DoFn: func(cmd Completed) RedisResult {
					return slotsMultiResp
				}, DoMultiFn: func(multi ...Completed) *redisresults {
					switch atomic.AddInt64(&count, 1) {
					case 1:
						return &redisresults{s: []RedisResult{
							newResult(strmsg('+', "1"), nil),
							newResult(strmsg('+', "OK"), nil),
							newResult(strmsg('+', "QUEUED"), nil),
							newResult(strmsg('-', "ASK 0 :1"), nil),
							newResult(strmsg('-', "EXECABORT"), nil),
							newResult(strmsg('+', "4"), nil),
							newResult(strmsg('+', "OK"), nil),
							newResult(strmsg('-', "ASK 0 :1"), nil),
							newResult(strmsg('+', "QUEUED"), nil),
							newResult(strmsg('-', "EXECABORT"), nil),
							newResult(strmsg('+', "7"), nil),
						}}
					case 2:
						return &redisresults{s: []RedisResult{
							newResult(strmsg('+', "OK"), nil),
							newResult(strmsg('+', "OK"), nil),
							newResult(strmsg('+', "QUEUED"), nil),
							newResult(strmsg('+', "QUEUED"), nil),
							newResult(slicemsg('*', []RedisMessage{
								strmsg('+', "2"),
								strmsg('+', "3"),
							}), nil),
							newResult(strmsg('+', "OK"), nil),
							newResult(strmsg('+', "OK"), nil),
							newResult(strmsg('+', "QUEUED"), nil),
							newResult(strmsg('+', "QUEUED"), nil),
							newResult(slicemsg('*', []RedisMessage{
								strmsg('+', "5"),
								strmsg('+', "6"),
							}), nil),
						}}
					}
					return nil
				}}
			},
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		resps := client.DoMulti(
			context.Background(),
			client.B().Get().Key("1{t}").Build(),
			client.B().Multi().Build(),
			client.B().Get().Key("2{t}").Build(),
			client.B().Get().Key("3{t}").Build(),
			client.B().Exec().Build(),
			client.B().Get().Key("4{t}").Build(),
			client.B().Multi().Build(),
			client.B().Get().Key("5{t}").Build(),
			client.B().Get().Key("6{t}").Build(),
			client.B().Exec().Build(),
			client.B().Get().Key("7{t}").Build(),
		)
		if v, err := resps[0].ToString(); err != nil || v != "1" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[1].ToString(); err != nil || v != "OK" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[2].ToString(); err != nil || v != "QUEUED" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[3].ToString(); err != nil || v != "QUEUED" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[4].AsStrSlice(); err != nil || !reflect.DeepEqual(v, []string{"2", "3"}) {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[5].ToString(); err != nil || v != "4" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[6].ToString(); err != nil || v != "OK" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[7].ToString(); err != nil || v != "QUEUED" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[8].ToString(); err != nil || v != "QUEUED" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[9].AsStrSlice(); err != nil || !reflect.DeepEqual(v, []string{"5", "6"}) {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[10].ToString(); err != nil || v != "7" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot moved DoMulti except transactions", func(t *testing.T) {
		var count int64
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn {
				return &mockConn{DoFn: func(cmd Completed) RedisResult {
					return slotsMultiResp
				}, DoMultiFn: func(multi ...Completed) *redisresults {
					switch atomic.AddInt64(&count, 1) {
					case 1:
						return &redisresults{s: []RedisResult{
							newResult(strmsg('-', "MOVED 0 :1"), nil),
							newResult(strmsg('+', "OK"), nil),
							newResult(strmsg('+', "QUEUED"), nil),
							newResult(strmsg('+', "QUEUED"), nil),
							newResult(slicemsg('*', []RedisMessage{
								strmsg('+', "2"),
								strmsg('+', "3"),
							}), nil),
							newResult(strmsg('-', "MOVED 0 :1"), nil),
							newResult(strmsg('+', "OK"), nil),
							newResult(strmsg('+', "QUEUED"), nil),
							newResult(strmsg('+', "QUEUED"), nil),
							newResult(slicemsg('*', []RedisMessage{
								strmsg('+', "5"),
								strmsg('+', "6"),
							}), nil),
							newResult(strmsg('-', "MOVED 0 :1"), nil),
						}}
					case 2:
						return &redisresults{s: []RedisResult{
							newResult(strmsg('+', "1"), nil),
							newResult(strmsg('+', "4"), nil),
							newResult(strmsg('+', "7"), nil),
						}}
					}
					return nil
				}}
			},
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		resps := client.DoMulti(
			context.Background(),
			client.B().Get().Key("1{t}").Build(),
			client.B().Multi().Build(),
			client.B().Get().Key("2{t}").Build(),
			client.B().Get().Key("3{t}").Build(),
			client.B().Exec().Build(),
			client.B().Get().Key("4{t}").Build(),
			client.B().Multi().Build(),
			client.B().Get().Key("5{t}").Build(),
			client.B().Get().Key("6{t}").Build(),
			client.B().Exec().Build(),
			client.B().Get().Key("7{t}").Build(),
		)
		if v, err := resps[0].ToString(); err != nil || v != "1" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[1].ToString(); err != nil || v != "OK" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[2].ToString(); err != nil || v != "QUEUED" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[3].ToString(); err != nil || v != "QUEUED" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[4].AsStrSlice(); err != nil || !reflect.DeepEqual(v, []string{"2", "3"}) {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[5].ToString(); err != nil || v != "4" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[6].ToString(); err != nil || v != "OK" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[7].ToString(); err != nil || v != "QUEUED" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[8].ToString(); err != nil || v != "QUEUED" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[9].AsStrSlice(); err != nil || !reflect.DeepEqual(v, []string{"5", "6"}) {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[10].ToString(); err != nil || v != "7" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot moved DoMulti except transactions ASKING", func(t *testing.T) {
		var count int64
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn {
				return &mockConn{DoFn: func(cmd Completed) RedisResult {
					return slotsMultiResp
				}, DoMultiFn: func(multi ...Completed) *redisresults {
					switch atomic.AddInt64(&count, 1) {
					case 1:
						return &redisresults{s: []RedisResult{
							newResult(strmsg('-', "ASK 0 :1"), nil),
							newResult(strmsg('+', "OK"), nil),
							newResult(strmsg('+', "QUEUED"), nil),
							newResult(strmsg('+', "QUEUED"), nil),
							newResult(slicemsg('*', []RedisMessage{
								strmsg('+', "2"),
								strmsg('+', "3"),
							}), nil),
							newResult(strmsg('-', "ASK 0 :1"), nil),
							newResult(strmsg('+', "OK"), nil),
							newResult(strmsg('+', "QUEUED"), nil),
							newResult(strmsg('+', "QUEUED"), nil),
							newResult(slicemsg('*', []RedisMessage{
								strmsg('+', "5"),
								strmsg('+', "6"),
							}), nil),
							newResult(strmsg('-', "ASK 0 :1"), nil),
						}}
					case 2:
						return &redisresults{s: []RedisResult{
							newResult(strmsg('+', "OK"), nil),
							newResult(strmsg('+', "1"), nil),
							newResult(strmsg('+', "OK"), nil),
							newResult(strmsg('+', "4"), nil),
							newResult(strmsg('+', "OK"), nil),
							newResult(strmsg('+', "7"), nil),
						}}
					}
					return nil
				}}
			},
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		resps := client.DoMulti(
			context.Background(),
			client.B().Get().Key("1{t}").Build(),
			client.B().Multi().Build(),
			client.B().Get().Key("2{t}").Build(),
			client.B().Get().Key("3{t}").Build(),
			client.B().Exec().Build(),
			client.B().Get().Key("4{t}").Build(),
			client.B().Multi().Build(),
			client.B().Get().Key("5{t}").Build(),
			client.B().Get().Key("6{t}").Build(),
			client.B().Exec().Build(),
			client.B().Get().Key("7{t}").Build(),
		)
		if v, err := resps[0].ToString(); err != nil || v != "1" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[1].ToString(); err != nil || v != "OK" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[2].ToString(); err != nil || v != "QUEUED" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[3].ToString(); err != nil || v != "QUEUED" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[4].AsStrSlice(); err != nil || !reflect.DeepEqual(v, []string{"2", "3"}) {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[5].ToString(); err != nil || v != "4" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[6].ToString(); err != nil || v != "OK" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[7].ToString(); err != nil || v != "QUEUED" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[8].ToString(); err != nil || v != "QUEUED" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[9].AsStrSlice(); err != nil || !reflect.DeepEqual(v, []string{"5", "6"}) {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[10].ToString(); err != nil || v != "7" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot moved DoMulti transactions mixed", func(t *testing.T) {
		var count int64
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn {
				return &mockConn{DoFn: func(cmd Completed) RedisResult {
					return slotsMultiResp
				}, DoMultiFn: func(multi ...Completed) *redisresults {
					switch atomic.AddInt64(&count, 1) {
					case 1:
						return &redisresults{s: []RedisResult{
							newResult(strmsg('+', "1"), nil),
							newResult(strmsg('+', "OK"), nil),
							newResult(strmsg('-', "MOVED 0 :1"), nil),
							newResult(strmsg('-', "MOVED 0 :1"), nil),
							newResult(strmsg('-', "EXECABORT"), nil),
							newResult(strmsg('-', "MOVED 0 :1"), nil),
							newResult(strmsg('+', "OK"), nil),
							newResult(strmsg('-', "MOVED 0 :1"), nil),
							newResult(strmsg('-', "MOVED 0 :1"), nil),
							newResult(strmsg('-', "EXECABORT"), nil),
							newResult(strmsg('+', "7"), nil),
						}}
					case 2:
						return &redisresults{s: []RedisResult{
							newResult(strmsg('+', "OK"), nil),
							newResult(strmsg('+', "QUEUED"), nil),
							newResult(strmsg('+', "QUEUED"), nil),
							newResult(slicemsg('*', []RedisMessage{
								strmsg('+', "2"),
								strmsg('+', "3"),
							}), nil),
							newResult(strmsg('+', "4"), nil),
							newResult(strmsg('+', "OK"), nil),
							newResult(strmsg('+', "QUEUED"), nil),
							newResult(strmsg('+', "QUEUED"), nil),
							newResult(slicemsg('*', []RedisMessage{
								strmsg('+', "5"),
								strmsg('+', "6"),
							}), nil),
						}}
					}
					return nil
				}}
			},
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		resps := client.DoMulti(
			context.Background(),
			client.B().Get().Key("1{t}").Build(),
			client.B().Multi().Build(),
			client.B().Get().Key("2{t}").Build(),
			client.B().Get().Key("3{t}").Build(),
			client.B().Exec().Build(),
			client.B().Get().Key("4{t}").Build(),
			client.B().Multi().Build(),
			client.B().Get().Key("5{t}").Build(),
			client.B().Get().Key("6{t}").Build(),
			client.B().Exec().Build(),
			client.B().Get().Key("7{t}").Build(),
		)
		if v, err := resps[0].ToString(); err != nil || v != "1" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[1].ToString(); err != nil || v != "OK" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[2].ToString(); err != nil || v != "QUEUED" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[3].ToString(); err != nil || v != "QUEUED" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[4].AsStrSlice(); err != nil || !reflect.DeepEqual(v, []string{"2", "3"}) {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[5].ToString(); err != nil || v != "4" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[6].ToString(); err != nil || v != "OK" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[7].ToString(); err != nil || v != "QUEUED" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[8].ToString(); err != nil || v != "QUEUED" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[9].AsStrSlice(); err != nil || !reflect.DeepEqual(v, []string{"5", "6"}) {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[10].ToString(); err != nil || v != "7" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot moved DoMulti transactions mixed ASKING", func(t *testing.T) {
		var count int64
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn {
				return &mockConn{DoFn: func(cmd Completed) RedisResult {
					return slotsMultiResp
				}, DoMultiFn: func(multi ...Completed) *redisresults {
					switch atomic.AddInt64(&count, 1) {
					case 1:
						return &redisresults{s: []RedisResult{
							newResult(strmsg('+', "1"), nil),
							newResult(strmsg('+', "OK"), nil),
							newResult(strmsg('-', "ASK 0 :1"), nil),
							newResult(strmsg('-', "ASK 0 :1"), nil),
							newResult(strmsg('-', "EXECABORT"), nil),
							newResult(strmsg('-', "ASK 0 :1"), nil),
							newResult(strmsg('+', "OK"), nil),
							newResult(strmsg('-', "ASK 0 :1"), nil),
							newResult(strmsg('-', "ASK 0 :1"), nil),
							newResult(strmsg('-', "EXECABORT"), nil),
							newResult(strmsg('+', "7"), nil),
						}}
					case 2:
						return &redisresults{s: []RedisResult{
							newResult(strmsg('+', "OK"), nil),
							newResult(strmsg('+', "OK"), nil),
							newResult(strmsg('+', "QUEUED"), nil),
							newResult(strmsg('+', "QUEUED"), nil),
							newResult(slicemsg('*', []RedisMessage{
								strmsg('+', "2"),
								strmsg('+', "3"),
							}), nil),
							newResult(strmsg('+', "OK"), nil),
							newResult(strmsg('+', "4"), nil),
							newResult(strmsg('+', "OK"), nil),
							newResult(strmsg('+', "OK"), nil),
							newResult(strmsg('+', "QUEUED"), nil),
							newResult(strmsg('+', "QUEUED"), nil),
							newResult(slicemsg('*', []RedisMessage{
								strmsg('+', "5"),
								strmsg('+', "6"),
							}), nil),
						}}
					}
					return nil
				}}
			},
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		resps := client.DoMulti(
			context.Background(),
			client.B().Get().Key("1{t}").Build(),
			client.B().Multi().Build(),
			client.B().Get().Key("2{t}").Build(),
			client.B().Get().Key("3{t}").Build(),
			client.B().Exec().Build(),
			client.B().Get().Key("4{t}").Build(),
			client.B().Multi().Build(),
			client.B().Get().Key("5{t}").Build(),
			client.B().Get().Key("6{t}").Build(),
			client.B().Exec().Build(),
			client.B().Get().Key("7{t}").Build(),
		)
		if v, err := resps[0].ToString(); err != nil || v != "1" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[1].ToString(); err != nil || v != "OK" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[2].ToString(); err != nil || v != "QUEUED" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[3].ToString(); err != nil || v != "QUEUED" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[4].AsStrSlice(); err != nil || !reflect.DeepEqual(v, []string{"2", "3"}) {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[5].ToString(); err != nil || v != "4" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[6].ToString(); err != nil || v != "OK" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[7].ToString(); err != nil || v != "QUEUED" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[8].ToString(); err != nil || v != "QUEUED" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[9].AsStrSlice(); err != nil || !reflect.DeepEqual(v, []string{"5", "6"}) {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[10].ToString(); err != nil || v != "7" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot moved DoMulti transactions edge cases 1", func(t *testing.T) {
		var count int64
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn {
				return &mockConn{DoFn: func(cmd Completed) RedisResult {
					return slotsMultiResp
				}, DoMultiFn: func(multi ...Completed) *redisresults {
					switch atomic.AddInt64(&count, 1) {
					case 1:
						return &redisresults{s: []RedisResult{
							newResult(strmsg('+', "1"), nil),
							newResult(strmsg('+', "OK"), nil),
							newResult(strmsg('-', "ERR Command not allowed inside a transaction"), nil),
							newResult(strmsg('-', "MOVED 0 :1"), nil),
							newResult(strmsg('-', "MOVED 0 :1"), nil),
							newResult(strmsg('-', "EXECABORT"), nil),
							newResult(strmsg('-', "MOVED 0 :1"), nil),
						}}
					case 2:
						return &redisresults{s: []RedisResult{
							newResult(strmsg('+', "4"), nil),
						}}
					}
					return nil
				}}
			},
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		resps := client.DoMulti(
			context.Background(),
			client.B().Get().Key("1{t}").Build(),
			client.B().Multi().Build(),
			client.B().Multi().Build(), // nested multi
			client.B().Get().Key("2{t}").Build(),
			client.B().Get().Key("3{t}").Build(),
			client.B().Exec().Build(),
			client.B().Get().Key("4{t}").Build(),
		)
		if v, err := resps[0].ToString(); err != nil || v != "1" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[1].ToString(); err != nil || v != "OK" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if err := resps[2].Error(); err == nil || !strings.Contains(err.Error(), "Command not allowed inside a transaction") {
			t.Fatalf("unexpected err %v", err)
		}
		if err := resps[3].Error(); err == nil || !strings.Contains(err.Error(), "MOVED 0 :1") {
			t.Fatalf("unexpected err %v", err)
		}
		if err := resps[4].Error(); err == nil || !strings.Contains(err.Error(), "MOVED 0 :1") {
			t.Fatalf("unexpected err %v", err)
		}
		if err := resps[5].Error(); err == nil || !strings.Contains(err.Error(), "EXECABORT") {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := resps[6].ToString(); err != nil || v != "4" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot moved DoMulti transactions edge cases 2", func(t *testing.T) {
		var count int64
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn {
				return &mockConn{DoFn: func(cmd Completed) RedisResult {
					return slotsMultiResp
				}, DoMultiFn: func(multi ...Completed) *redisresults {
					switch atomic.AddInt64(&count, 1) {
					case 1:
						return &redisresults{s: []RedisResult{
							newResult(strmsg('+', "1"), nil),
							newResult(strmsg('+', "OK"), nil),
							newResult(strmsg('-', "MOVED 0 :1"), nil),
							newResult(strmsg('-', "ERR Command not allowed inside a transaction"), nil),
							newResult(strmsg('-', "MOVED 0 :1"), nil),
							newResult(strmsg('-', "EXECABORT"), nil),
							newResult(strmsg('-', "MOVED 0 :1"), nil),
						}}
					case 2:
						return &redisresults{s: []RedisResult{
							newResult(strmsg('+', "4"), nil),
						}}
					}
					return nil
				}}
			},
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		resps := client.DoMulti(
			context.Background(),
			client.B().Get().Key("1{t}").Build(),
			client.B().Multi().Build(),
			client.B().Get().Key("2{t}").Build(),
			client.B().Multi().Build(), // nested multi
			client.B().Get().Key("3{t}").Build(),
			client.B().Exec().Build(),
			client.B().Get().Key("4{t}").Build(),
		)
		if v, err := resps[0].ToString(); err != nil || v != "1" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[1].ToString(); err != nil || v != "OK" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if err := resps[2].Error(); err == nil || !strings.Contains(err.Error(), "MOVED 0 :1") {
			t.Fatalf("unexpected err %v", err)
		}
		if err := resps[3].Error(); err == nil || !strings.Contains(err.Error(), "Command not allowed inside a transaction") {
			t.Fatalf("unexpected err %v", err)
		}
		if err := resps[4].Error(); err == nil || !strings.Contains(err.Error(), "MOVED 0 :1") {
			t.Fatalf("unexpected err %v", err)
		}
		if err := resps[5].Error(); err == nil || !strings.Contains(err.Error(), "EXECABORT") {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := resps[6].ToString(); err != nil || v != "4" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot moved DoMulti transactions edge cases 3", func(t *testing.T) {
		var count int64
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn {
				return &mockConn{DoFn: func(cmd Completed) RedisResult {
					return slotsMultiResp
				}, DoMultiFn: func(multi ...Completed) *redisresults {
					switch atomic.AddInt64(&count, 1) {
					case 1:
						return &redisresults{s: []RedisResult{
							newResult(strmsg('+', "1"), nil),
							newResult(strmsg('+', "OK"), nil),
							newResult(strmsg('-', "MOVED 0 :1"), nil),
							newResult(strmsg('-', "ERR Command not allowed inside a transaction"), nil),
							newResult(strmsg('-', "MOVED 0 :1"), nil),
						}}
					}
					return nil
				}}
			},
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		resps := client.DoMulti(
			context.Background(),
			client.B().Get().Key("1{t}").Build(),
			client.B().Multi().Build(),
			client.B().Get().Key("2{t}").Build(),
			client.B().Multi().Build(), // nested multi
			client.B().Get().Key("3{t}").Build(),
		)
		if v, err := resps[0].ToString(); err != nil || v != "1" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if v, err := resps[1].ToString(); err != nil || v != "OK" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
		if err := resps[2].Error(); err == nil || !strings.Contains(err.Error(), "MOVED 0 :1") {
			t.Fatalf("unexpected err %v", err)
		}
		if err := resps[3].Error(); err == nil || !strings.Contains(err.Error(), "Command not allowed inside a transaction") {
			t.Fatalf("unexpected err %v", err)
		}
		if err := resps[4].Error(); err == nil || !strings.Contains(err.Error(), "MOVED 0 :1") {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("slot moved DoMulti (multi)", func(t *testing.T) {
		var count int64
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn {
				return &mockConn{DoFn: func(cmd Completed) RedisResult {
					return slotsMultiResp
				}, DoMultiFn: func(multi ...Completed) *redisresults {
					ret := make([]RedisResult, len(multi))
					if atomic.AddInt64(&count, 1) <= 3 {
						for i := range ret {
							ret[i] = newResult(strmsg('-', "MOVED 0 :1"), nil)
						}
						return &redisresults{s: ret}
					}
					for i := range ret {
						ret[i] = newResult(strmsg('+', multi[i].Commands()[1]), nil)
					}
					return &redisresults{s: ret}
				}}
			},
			newRetryer(defaultRetryDelayFn),
		)
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
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn {
				return &mockConn{DoFn: func(cmd Completed) RedisResult {
					return slotsMultiResp
				}, DoMultiFn: func(multi ...Completed) *redisresults {
					ret := make([]RedisResult, len(multi))
					if atomic.AddInt64(&count, 1) <= 2 {
						for i := range ret {
							ret[i] = newResult(strmsg('-', "TRYAGAIN"), nil)
						}
						return &redisresults{s: ret}
					}
					for i := range ret {
						ret[i] = newResult(strmsg('+', multi[i].Commands()[1]), nil)
					}
					return &redisresults{s: ret}
				}}
			},
			newRetryer(defaultRetryDelayFn),
		)
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
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn {
				if dst == ":2" {
					atomic.AddInt64(&check, 1)
				}
				return &mockConn{DoFn: func(cmd Completed) RedisResult {
					if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
						return slotsMultiResp
					}
					if atomic.AddInt64(&count, 1) <= 3 {
						return newResult(strmsg('-', "MOVED 0 :2"), nil)
					}
					return newResult(strmsg('+', "b"), nil)
				}}
			},
			newRetryer(defaultRetryDelayFn),
		)
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
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn {
				if dst == ":2" {
					atomic.AddInt64(&check, 1)
				}
				return &mockConn{DoFn: func(cmd Completed) RedisResult {
					return slotsMultiResp
				}, DoMultiFn: func(multi ...Completed) *redisresults {
					ret := make([]RedisResult, len(multi))
					if atomic.AddInt64(&count, 1) <= 3 {
						for i := range ret {
							ret[i] = newResult(strmsg('-', "MOVED 0 :2"), nil)
						}
						return &redisresults{s: ret}
					}
					for i := range ret {
						ret[i] = newResult(strmsg('+', multi[i].Commands()[1]), nil)
					}
					return &redisresults{s: ret}
				}}
			},
			newRetryer(defaultRetryDelayFn),
		)
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
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn {
				if dst == ":2" {
					atomic.AddInt64(&check, 1)
				}
				return &mockConn{DoFn: func(cmd Completed) RedisResult {
					return slotsMultiResp
				}, DoMultiFn: func(multi ...Completed) *redisresults {
					ret := make([]RedisResult, len(multi))
					if atomic.AddInt64(&count, 1) <= 3 {
						for i := range ret {
							ret[i] = newResult(strmsg('-', "MOVED 0 :2"), nil)
						}
						return &redisresults{s: ret}
					}
					for i := range ret {
						ret[i] = newResult(strmsg('+', multi[i].Commands()[1]), nil)
					}
					return &redisresults{s: ret}
				}}
			},
			newRetryer(defaultRetryDelayFn),
		)
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
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn {
				return &mockConn{DoFn: func(cmd Completed) RedisResult {
					return slotsMultiResp
				}, DoMultiFn: func(multi ...Completed) *redisresults {
					ret := make([]RedisResult, len(multi))
					if atomic.AddInt64(&count, 1) <= 2 {
						for i := range ret {
							ret[i] = newResult(strmsg('-', "TRYAGAIN"), nil)
						}
						return &redisresults{s: ret}
					}
					for i := range ret {
						ret[i] = newResult(strmsg('+', multi[i].Commands()[1]), nil)
					}
					return &redisresults{s: ret}
				}}
			},
			newRetryer(defaultRetryDelayFn),
		)
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
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn {
				return &mockConn{
					DoFn: func(cmd Completed) RedisResult {
						return slotsMultiResp
					},
					DoCacheFn: func(cmd Cacheable, ttl time.Duration) RedisResult {
						if atomic.AddInt64(&count, 1) <= 3 {
							return newResult(strmsg('-', "MOVED 0 :1"), nil)
						}
						return newResult(strmsg('+', "b"), nil)
					},
				}
			},
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.DoCache(context.Background(), client.B().Get().Key("a").Cache(), 100).ToString(); err != nil || v != "b" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot moved (cache) redirect once", func(t *testing.T) {
		var count int64
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn {
				return &mockConn{
					DoFn: func(cmd Completed) RedisResult {
						return slotsMultiResp
					},
					DoCacheFn: func(cmd Cacheable, ttl time.Duration) RedisResult {
						if strings.Contains(dst, ":0") {
							atomic.AddInt64(&count, 1)
							return newResult(strmsg('-', "MOVED 0 :2"), nil)
						}

						return newResult(strmsg('+', "b"), nil)
					},
				}
			},
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}

		for i := 0; i < 10; i++ {
			if v, err := client.DoCache(context.Background(), client.B().Get().Key("a").Cache(), 100).ToString(); err != nil || v != "b" {
				t.Fatalf("unexpected resp %v %v", v, err)
			}
		}

		if atomic.LoadInt64(&count) != 1 {
			t.Fatalf("unexpected count %v", count)
		}
	})

	t.Run("slot moved (cache multi 1)", func(t *testing.T) {
		var count int64
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn {
				return &mockConn{DoFn: func(cmd Completed) RedisResult {
					return slotsMultiResp
				}, DoMultiCacheFn: func(multi ...CacheableTTL) *redisresults {
					ret := make([]RedisResult, len(multi))
					if atomic.AddInt64(&count, 1) <= 3 {
						for i := range ret {
							ret[i] = newResult(strmsg('-', "MOVED 0 :1"), nil)
						}
						return &redisresults{s: ret}
					}
					for i := range ret {
						ret[i] = newResult(strmsg('+', multi[i].Cmd.Commands()[1]), nil)
					}
					return &redisresults{s: ret}
				}}
			},
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.DoMultiCache(context.Background(), CT(client.B().Get().Key("a").Cache(), 100))[0].ToString(); err != nil || v != "a" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot moved (cache multi 1) redirect once", func(t *testing.T) {
		var count int64
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn {
				return &mockConn{
					DoFn: func(cmd Completed) RedisResult {
						return slotsMultiResp
					},
					DoMultiCacheFn: func(multi ...CacheableTTL) *redisresults {
						ret := make([]RedisResult, len(multi))

						if strings.Contains(dst, ":0") {
							atomic.AddInt64(&count, 1)
							for i := range ret {
								ret[i] = newResult(strmsg('-', "MOVED 0 :2"), nil)
							}
							return &redisresults{s: ret}
						}

						for i := range ret {
							ret[i] = newResult(strmsg('+', multi[i].Cmd.Commands()[1]), nil)
						}
						return &redisresults{s: ret}
					},
				}
			},
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}

		for i := 0; i < 10; i++ {
			if v, err := client.DoMultiCache(context.Background(), CT(client.B().Get().Key("a").Cache(), 100))[0].ToString(); err != nil || v != "a" {
				t.Fatalf("unexpected resp %v %v", v, err)
			}
		}

		if atomic.LoadInt64(&count) != 1 {
			t.Fatalf("unexpected count %v", count)
		}
	})

	t.Run("slot moved (cache multi 2)", func(t *testing.T) {
		var count int64
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn {
				return &mockConn{DoFn: func(cmd Completed) RedisResult {
					return slotsMultiResp
				}, DoMultiCacheFn: func(multi ...CacheableTTL) *redisresults {
					ret := make([]RedisResult, len(multi))
					if atomic.AddInt64(&count, 1) <= 3 {
						for i := range ret {
							ret[i] = newResult(strmsg('-', "MOVED 0 :1"), nil)
						}
						return &redisresults{s: ret}
					}
					for i := range ret {
						ret[i] = newResult(strmsg('+', multi[i].Cmd.Commands()[1]), nil)
					}
					return &redisresults{s: ret}
				}}
			},
			newRetryer(defaultRetryDelayFn),
		)
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
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn {
				return &mockConn{
					DoFn: func(cmd Completed) RedisResult {
						if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
							return slotsMultiResp
						}
						return newResult(strmsg('-', "ASK 0 :1"), nil)
					},
					DoMultiFn: func(multi ...Completed) *redisresults {
						if atomic.AddInt64(&count, 1) <= 3 {
							return &redisresults{s: []RedisResult{{}, newResult(strmsg('-', "ASK 0 :1"), nil)}}
						}
						return &redisresults{s: []RedisResult{{}, newResult(strmsg('+', "b"), nil)}}
					},
				}
			},
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.Do(context.Background(), client.B().Get().Key("a").Build()).ToString(); err != nil || v != "b" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot asking DoMulti (single)", func(t *testing.T) {
		var count int64
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn {
				return &mockConn{
					DoFn: func(cmd Completed) RedisResult {
						return slotsMultiResp
					},
					DoMultiFn: func(multi ...Completed) *redisresults {
						ret := make([]RedisResult, len(multi))
						if atomic.AddInt64(&count, 1) <= 3 {
							for i := range ret {
								ret[i] = newResult(strmsg('-', "ASK 0 :1"), nil)
							}
							return &redisresults{s: ret}
						}
						for i := 0; i < len(multi); i += 2 {
							ret[i] = newResult(strmsg('+', "OK"), nil)
							ret[i+1] = newResult(strmsg('+', multi[i+1].Commands()[1]), nil)
						}
						return &redisresults{s: ret}
					},
				}
			},
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.DoMulti(context.Background(), client.B().Get().Key("a").Build())[0].ToString(); err != nil || v != "a" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot asking DoMulti (multi)", func(t *testing.T) {
		var count int64
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn {
				return &mockConn{
					DoFn: func(cmd Completed) RedisResult {
						return slotsMultiResp
					},
					DoMultiFn: func(multi ...Completed) *redisresults {
						ret := make([]RedisResult, len(multi))
						if atomic.AddInt64(&count, 1) <= 3 {
							for i := range ret {
								ret[i] = newResult(strmsg('-', "ASK 0 :1"), nil)
							}
							return &redisresults{s: ret}
						}
						for i := 0; i < len(multi); i += 2 {
							ret[i] = newResult(strmsg('+', "OK"), nil)
							ret[i+1] = newResult(strmsg('+', multi[i+1].Commands()[1]), nil)
						}
						return &redisresults{s: ret}
					},
				}
			},
			newRetryer(defaultRetryDelayFn),
		)
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
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn {
				return &mockConn{
					DoFn: func(cmd Completed) RedisResult {
						return slotsMultiResp
					},
					DoCacheFn: func(cmd Cacheable, ttl time.Duration) RedisResult {
						return newResult(strmsg('-', "ASK 0 :1"), nil)
					},
					DoMultiFn: func(multi ...Completed) *redisresults {
						if atomic.AddInt64(&count, 1) <= 3 {
							return &redisresults{s: []RedisResult{{}, {}, {}, {}, {}, newResult(strmsg('-', "ASK 0 :1"), nil)}}
						}
						return &redisresults{s: []RedisResult{{}, {}, {}, {}, {}, newResult(slicemsg('*', []RedisMessage{{}, strmsg('+', "b")}), nil)}}
					},
				}
			},
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.DoCache(context.Background(), client.B().Get().Key("a").Cache(), 100).ToString(); err != nil || v != "b" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot asking (cache multi 1)", func(t *testing.T) {
		var count int64
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn {
				return &mockConn{
					DoFn: func(cmd Completed) RedisResult {
						return slotsMultiResp
					},
					DoMultiCacheFn: func(multi ...CacheableTTL) *redisresults {
						return &redisresults{s: []RedisResult{newResult(strmsg('-', "ASK 0 :1"), nil)}}
					},
					DoMultiFn: func(multi ...Completed) *redisresults {
						if atomic.AddInt64(&count, 1) <= 3 {
							return &redisresults{s: []RedisResult{{}, {}, {}, {}, {}, newResult(strmsg('-', "ASK 0 :1"), nil)}}
						}
						return &redisresults{s: []RedisResult{{}, {}, {}, {}, {}, newResult(slicemsg('*', []RedisMessage{{}, strmsg('+', "b")}), nil)}}
					},
				}
			},
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.DoMultiCache(context.Background(), CT(client.B().Get().Key("a").Cache(), 100))[0].ToString(); err != nil || v != "b" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot asking (cache multi 2)", func(t *testing.T) {
		var count int64
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn {
				return &mockConn{
					DoFn: func(cmd Completed) RedisResult {
						return slotsMultiResp
					},
					DoMultiCacheFn: func(multi ...CacheableTTL) *redisresults {
						return &redisresults{s: []RedisResult{newResult(strmsg('-', "ASK 0 :1"), nil)}}
					},
					DoMultiFn: func(multi ...Completed) *redisresults {
						if atomic.AddInt64(&count, 1) <= 3 {
							return &redisresults{s: []RedisResult{
								{}, {}, {}, {}, {}, newResult(strmsg('-', "ASK 0 :1"), nil),
								{}, {}, {}, {}, {}, newResult(strmsg('-', "ASK 0 :1"), nil),
							}}
						}
						return &redisresults{s: []RedisResult{
							{}, {}, {}, {}, {}, newResult(slicemsg('*', []RedisMessage{{}, {}, strmsg('+', multi[4].Commands()[1])}), nil),
							{}, {}, {}, {}, {}, newResult(slicemsg('*', []RedisMessage{{}, {}, strmsg('+', multi[10].Commands()[1])}), nil),
						}}
					},
				}
			},
			newRetryer(defaultRetryDelayFn),
		)
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
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn {
				return &mockConn{DoFn: func(cmd Completed) RedisResult {
					if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
						return slotsMultiResp
					}
					if atomic.AddInt64(&count, 1) <= 3 {
						return newResult(strmsg('-', "TRYAGAIN"), nil)
					}
					return newResult(strmsg('+', "b"), nil)
				}}
			},
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.Do(context.Background(), client.B().Get().Key("a").Build()).ToString(); err != nil || v != "b" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot try again DoMulti 1", func(t *testing.T) {
		var count int64
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn {
				return &mockConn{DoFn: func(cmd Completed) RedisResult {
					return slotsMultiResp
				}, DoMultiFn: func(multi ...Completed) *redisresults {
					if atomic.AddInt64(&count, 1) <= 3 {
						return &redisresults{s: []RedisResult{newResult(strmsg('-', "TRYAGAIN"), nil)}}
					}
					ret := make([]RedisResult, len(multi))
					ret[0] = newResult(strmsg('+', "b"), nil)
					return &redisresults{s: ret}
				}}
			},
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.DoMulti(context.Background(), client.B().Get().Key("a").Build())[0].ToString(); err != nil || v != "b" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot try again DoMulti 2", func(t *testing.T) {
		var count int64
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn {
				return &mockConn{DoFn: func(cmd Completed) RedisResult {
					return slotsMultiResp
				}, DoMultiFn: func(multi ...Completed) *redisresults {
					if atomic.AddInt64(&count, 1) <= 3 {
						return &redisresults{s: []RedisResult{newResult(strmsg('-', "TRYAGAIN"), nil)}}
					}
					ret := make([]RedisResult, len(multi))
					ret[0] = newResult(strmsg('+', multi[0].Commands()[1]), nil)
					return &redisresults{s: ret}
				}}
			},
			newRetryer(defaultRetryDelayFn),
		)
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
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn {
				return &mockConn{
					DoFn: func(cmd Completed) RedisResult {
						return slotsMultiResp
					},
					DoCacheFn: func(cmd Cacheable, ttl time.Duration) RedisResult {
						if atomic.AddInt64(&count, 1) <= 3 {
							return newResult(strmsg('-', "TRYAGAIN"), nil)
						}
						return newResult(strmsg('+', "b"), nil)
					},
				}
			},
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.DoCache(context.Background(), client.B().Get().Key("a").Cache(), 100).ToString(); err != nil || v != "b" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot try again (cache multi 1)", func(t *testing.T) {
		var count int64
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn {
				return &mockConn{DoFn: func(cmd Completed) RedisResult {
					return slotsMultiResp
				}, DoMultiCacheFn: func(multi ...CacheableTTL) *redisresults {
					if atomic.AddInt64(&count, 1) <= 3 {
						return &redisresults{s: []RedisResult{newResult(strmsg('-', "TRYAGAIN"), nil)}}
					}
					return &redisresults{s: []RedisResult{newResult(strmsg('+', "b"), nil)}}
				}}
			},
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if v, err := client.DoMultiCache(context.Background(), CT(client.B().Get().Key("a").Cache(), 100))[0].ToString(); err != nil || v != "b" {
			t.Fatalf("unexpected resp %v %v", v, err)
		}
	})

	t.Run("slot try again (cache multi 2)", func(t *testing.T) {
		var count int64
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn {
				return &mockConn{DoFn: func(cmd Completed) RedisResult {
					if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
						return slotsMultiResp
					}
					return shardsMultiResp
				}, DoMultiCacheFn: func(multi ...CacheableTTL) *redisresults {
					if atomic.AddInt64(&count, 1) <= 3 {
						return &redisresults{s: []RedisResult{newResult(strmsg('-', "TRYAGAIN"), nil)}}
					}
					return &redisresults{s: []RedisResult{newResult(strmsg('+', multi[0].Cmd.Commands()[1]), nil)}}
				}}
			},
			newRetryer(defaultRetryDelayFn),
		)
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
	defer ShouldNotLeak(SetupLeakDetection())
	SetupClientRetry(t, func(m *mockConn) Client {
		m.DoOverride = map[string]func(cmd Completed) RedisResult{
			"CLUSTER SLOTS": func(cmd Completed) RedisResult { return slotsMultiResp },
		}
		c, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn { return m },
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		return c
	})
}

func TestClusterClientReplicaOnly_PickReplica(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	m := &mockConn{
		DoFn: func(cmd Completed) RedisResult {
			if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
				return slotsMultiResp
			}
			return RedisResult{}
		},
	}

	client, err := newClusterClient(
		&ClientOption{InitAddress: []string{"127.0.0.1:0"}, ReplicaOnly: true},
		func(dst string, opt *ClientOption) conn {
			copiedM := *m
			return &copiedM
		},
		newRetryer(defaultRetryDelayFn),
	)
	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}
	t.Run("replicas should be picked", func(t *testing.T) {
		if client.wslots[0] != client.conns["127.0.1.1:1"].conn {
			t.Fatalf("unexpected replica node assigned to slot 0")
		}
		if client.wslots[8192] != client.conns["127.0.1.1:1"].conn {
			t.Fatalf("unexpected replica node assigned to slot 8192")
		}
		if client.wslots[8193] != client.conns["127.0.3.1:1"].conn {
			t.Fatalf("unexpected replica node assigned to slot 8193")
		}
		if client.wslots[16383] != client.conns["127.0.3.1:1"].conn {
			t.Fatalf("unexpected replica node assigned to slot 16383")
		}
	})
}

func TestClusterClientReplicaOnly_PickMasterIfNoReplica(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	t.Run("replicas should be picked", func(t *testing.T) {
		m := &mockConn{
			DoFn: func(cmd Completed) RedisResult {
				if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
					return slotsMultiResp
				}
				return RedisResult{}
			},
		}

		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{"127.0.0.1:0"}, ReplicaOnly: true},
			func(dst string, opt *ClientOption) conn {
				copiedM := *m
				return &copiedM
			},
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}

		if client.wslots[0] != client.conns["127.0.1.1:1"].conn {
			t.Fatalf("unexpected replica node assigned to slot 0")
		}
		if client.wslots[8192] != client.conns["127.0.1.1:1"].conn {
			t.Fatalf("unexpected replica node assigned to slot 8192")
		}
		if client.wslots[8193] != client.conns["127.0.3.1:1"].conn {
			t.Fatalf("unexpected replica node assigned to slot 8193")
		}
		if client.wslots[16383] != client.conns["127.0.3.1:1"].conn {
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

		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{"127.0.0.1:0"}, ReplicaOnly: true},
			func(dst string, opt *ClientOption) conn {
				copiedM := *m
				return &copiedM
			},
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}

		for slot := 0; slot < 8193; slot++ {
			if client.wslots[slot] == client.conns["127.0.0.2:1"].conn {
				continue
			}
			if client.wslots[slot] == client.conns["127.0.0.3:2"].conn {
				continue
			}
			if client.wslots[slot] == client.conns["127.0.0.4:3"].conn {
				continue
			}

			t.Fatalf("unexpected replica node assigned to slot %d", slot)
		}

		for slot := 8193; slot < 16384; slot++ {
			if client.wslots[slot] == client.conns["127.0.1.2:1"].conn {
				continue
			}
			if client.wslots[slot] == client.conns["127.0.1.3:2"].conn {
				continue
			}
			if client.wslots[slot] == client.conns["127.0.1.4:3"].conn {
				continue
			}

			t.Fatalf("unexpected replica node assigned to slot %d", slot)
		}
	})
}

func TestClusterShardsParsing(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
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
			_nodes := val.nodes
			sort.Slice(_nodes, func(i, j int) bool {
				return _nodes[i].Addr < _nodes[j].Addr
			})
			if len(_nodes) != 3 ||
				_nodes[0].Addr != "127.0.1.1:1" ||
				_nodes[1].Addr != "127.0.2.1:2" ||
				_nodes[2].Addr != "127.0.3.1:3" {
				t.Fatalf("unexpected nodes %v", _nodes)
			}
		}

		result = parseShards(shardsRespTls.val, "127.0.0.1:5", false)
		if len(result) != 1 {
			t.Fatalf("unexpected result %v", result)
		}
		for _, val := range result {
			_nodes := val.nodes
			sort.Slice(_nodes, func(i, j int) bool {
				return _nodes[i].Addr < _nodes[j].Addr
			})
			if len(_nodes) != 3 ||
				_nodes[0].Addr != "127.0.1.1:0" ||
				_nodes[1].Addr != "127.0.2.1:0" ||
				_nodes[2].Addr != "127.0.3.1:3" {
				t.Fatalf("unexpected nodes %v", _nodes)
			}
		}
	})

	t.Run("master position", func(t *testing.T) {
		result := parseShards(shardsRespTls.val, "127.0.0.1:5", true)
		if len(result) != 1 {
			t.Fatalf("unexpected result %v", result)
		}
		for master, group := range result {
			if len(group.nodes) == 0 || group.nodes[0].Addr != master {
				t.Fatalf("unexpected first node %v", group)
			}
		}
	})
}

// https://github.com/redis/rueidis/issues/543
func TestConnectToNonAvailableCluster(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 100; i++ {
				_, err := NewClient(ClientOption{
					InitAddress: []string{"127.0.0.1:3000", "127.0.0.1:3001", "127.0.0.1:3002"},
				})
				if err == nil {
					t.Errorf("expected connect error")
				}
			}
		}()
	}
	wg.Wait()
}

func TestClusterTopologyRefreshment(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())

	t.Run("no refreshment", func(t *testing.T) {
		var callCount int64
		cc, err := newClusterClient(
			&ClientOption{
				InitAddress: []string{"127.0.0.1:0"},
				ClusterOption: ClusterOption{
					ShardsRefreshInterval: 0,
				},
			},
			func(dst string, opt *ClientOption) conn {
				return &mockConn{
					DoFn: func(cmd Completed) RedisResult {
						// initial call
						if atomic.CompareAndSwapInt64(&callCount, 0, 1) {
							return singleSlotResp
						}

						t.Fatalf("unexpected call")
						return singleSlotResp
					},
				}
			},
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		defer cc.Close()

		time.Sleep(3 * time.Second) // verify that no refreshment is called

		if atomic.LoadInt64(&callCount) != 1 {
			t.Fatalf("unexpected call count %d", callCount)
		}
	})

	t.Run("nothing changed", func(t *testing.T) {
		var callCount int64
		refreshWaitCh := make(chan struct{})
		cli, err := newClusterClient(
			&ClientOption{
				InitAddress: []string{"127.0.0.1:0"},
				ClusterOption: ClusterOption{
					ShardsRefreshInterval: time.Second,
				},
			},
			func(dst string, opt *ClientOption) conn {
				return &mockConn{
					DoFn: func(cmd Completed) RedisResult {
						if atomic.AddInt64(&callCount, 1) >= 3 {
							defer func() { recover() }()
							defer close(refreshWaitCh)
							return singleSlotResp
						}
						return singleSlotResp
					},
				}
			},
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}

		select {
		case <-refreshWaitCh:
			cli.Close()

			cli.mu.Lock()
			conns := cli.conns
			cli.mu.Unlock()
			if len(conns) != 1 {
				t.Fatalf("unexpected conns %v", conns)
			}
			if _, ok := conns["127.0.0.1:0"]; !ok {
				t.Fatalf("unexpected conns %v", conns)
			}
		}
	})

	t.Run("replicas are changed", func(t *testing.T) {
		var callCount int64
		refreshWaitCh := make(chan struct{})
		cli, err := newClusterClient(
			&ClientOption{
				InitAddress: []string{"127.0.0.1:0"},
				ClusterOption: ClusterOption{
					ShardsRefreshInterval: time.Second,
				},
			},
			func(dst string, opt *ClientOption) conn {
				return &mockConn{
					DoFn: func(cmd Completed) RedisResult {
						if c := atomic.AddInt64(&callCount, 1); c >= 6 {
							defer func() { recover() }()
							defer close(refreshWaitCh)
							return slotsResp
						} else if c >= 3 {
							return slotsResp
						}
						return singleSlotResp
					},
				}
			},
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}

		select {
		case <-refreshWaitCh:
			cli.Close()

			cli.mu.Lock()
			conns := cli.conns
			cli.mu.Unlock()
			if len(conns) != 2 {
				t.Fatalf("unexpected conns %v", conns)
			}
			if _, ok := conns["127.0.0.1:0"]; !ok {
				t.Fatalf("unexpected conns %v", conns)
			}
			if _, ok := conns["127.0.1.1:1"]; !ok {
				t.Fatalf("unexpected conns %v", conns)
			}
		}
	})

	t.Run("shards are changed", func(t *testing.T) {
		var callCount int64
		refreshWaitCh := make(chan struct{})
		cli, err := newClusterClient(
			&ClientOption{
				InitAddress: []string{"127.0.0.1:0"},
				ClusterOption: ClusterOption{
					ShardsRefreshInterval: time.Second,
				},
			},
			func(dst string, opt *ClientOption) conn {
				return &mockConn{
					DoFn: func(cmd Completed) RedisResult {
						if c := atomic.AddInt64(&callCount, 1); c >= 6 {
							defer func() { recover() }()
							defer close(refreshWaitCh)
							return slotsMultiRespWithoutReplicas
						} else if c >= 3 {
							return slotsMultiRespWithoutReplicas
						}
						return singleSlotResp
					},
				}
			},
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}

		select {
		case <-refreshWaitCh:
			cli.Close()

			cli.mu.Lock()
			conns := cli.conns
			cli.mu.Unlock()
			if len(conns) != 2 {
				t.Fatalf("unexpected conns %v", conns)
			}
			if _, ok := conns["127.0.0.1:0"]; !ok {
				t.Fatalf("unexpected conns %v", conns)
			}
			if _, ok := conns["127.0.1.1:0"]; !ok {
				t.Fatalf("unexpected conns %v", conns)
			}
		}
	})
}

func TestClusterClientLoadingRetry(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())

	setup := func() (*clusterClient, *mockConn) {
		m := &mockConn{
			DoFn: func(cmd Completed) RedisResult {
				if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
					return slotsMultiResp
				}
				return RedisResult{}
			},
		}
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn { return m },
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		return client, m
	}

	t.Run("Do Retry on Loading", func(t *testing.T) {
		client, m := setup()
		attempts := 0
		m.DoFn = func(cmd Completed) RedisResult {
			if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
				return slotsMultiResp
			}
			attempts++
			if attempts == 1 {
				return newResult(strmsg('-', "LOADING Redis is loading the dataset in memory"), nil)
			}
			return newResult(strmsg('+', "OK"), nil)
		}

		if v, err := client.Do(context.Background(), client.B().Get().Key("test").Build()).ToString(); err != nil || v != "OK" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
		if attempts != 2 {
			t.Fatalf("expected 2 attempts, got %v", attempts)
		}
	})

	t.Run("Do not retry on non-loading errors", func(t *testing.T) {
		client, m := setup()
		attempts := 0
		m.DoFn = func(cmd Completed) RedisResult {
			if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
				return slotsMultiResp
			}
			attempts++
			if attempts == 1 {
				return newResult(strmsg('-', "ERR some other error"), nil)
			}
			return newResult(strmsg('+', "OK"), nil)
		}

		if err := client.Do(context.Background(), client.B().Get().Key("test").Build()).Error(); err == nil {
			t.Fatal("expected error but got nil")
		}
		if attempts != 1 {
			t.Fatalf("unexpected attempts %v, expected no retry", attempts)
		}
	})

	t.Run("DoMulti Retry on Loading", func(t *testing.T) {
		client, m := setup()
		attempts := 0
		m.DoMultiFn = func(multi ...Completed) *redisresults {
			attempts++
			if attempts == 1 {
				return &redisresults{s: []RedisResult{newResult(strmsg('-', "LOADING Redis is loading the dataset in memory"), nil)}}
			}
			return &redisresults{s: []RedisResult{newResult(strmsg('+', "OK"), nil)}}
		}

		cmd := client.B().Get().Key("test").Build()
		resps := client.DoMulti(context.Background(), cmd)
		if len(resps) != 1 {
			t.Fatalf("unexpected response length %v", len(resps))
		}
		if v, err := resps[0].ToString(); err != nil || v != "OK" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("DoCache Retry on Loading", func(t *testing.T) {
		client, m := setup()
		attempts := 0
		m.DoCacheFn = func(cmd Cacheable, ttl time.Duration) RedisResult {
			attempts++
			if attempts == 1 {
				return newResult(strmsg('-', "LOADING Redis is loading the dataset in memory"), nil)
			}
			return newResult(strmsg('+', "OK"), nil)
		}

		cmd := client.B().Get().Key("test").Cache()
		if v, err := client.DoCache(context.Background(), cmd, time.Minute).ToString(); err != nil || v != "OK" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("DoMultiCache Retry on Loading", func(t *testing.T) {
		client, m := setup()
		attempts := 0
		m.DoMultiCacheFn = func(multi ...CacheableTTL) *redisresults {
			attempts++
			if attempts == 1 {
				return &redisresults{s: []RedisResult{newResult(strmsg('-', "LOADING Redis is loading the dataset in memory"), nil)}}
			}
			return &redisresults{s: []RedisResult{newResult(strmsg('+', "OK"), nil)}}
		}

		cmd := client.B().Get().Key("test").Cache()
		resps := client.DoMultiCache(context.Background(), CT(cmd, time.Minute))
		if len(resps) != 1 {
			t.Fatalf("unexpected response length %v", len(resps))
		}
		if v, err := resps[0].ToString(); err != nil || v != "OK" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("Dedicated Do Retry on Loading", func(t *testing.T) {
		client, m := setup()
		attempts := 0
		m.DoFn = func(cmd Completed) RedisResult {
			if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
				return slotsMultiResp
			}
			attempts++
			if attempts == 1 {
				return newResult(strmsg('-', "LOADING Redis is loading the dataset in memory"), nil)
			}
			return newResult(strmsg('+', "OK"), nil)
		}
		m.AcquireFn = func() wire { return &mockWire{DoFn: m.DoFn} }

		err := client.Dedicated(func(c DedicatedClient) error {
			if v, err := c.Do(context.Background(), c.B().Get().Key("test").Build()).ToString(); err != nil || v != "OK" {
				t.Fatalf("unexpected response %v %v", v, err)
			}
			return nil
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("Dedicated DoMulti Retry on Loading", func(t *testing.T) {
		client, m := setup()
		attempts := 0
		m.DoMultiFn = func(multi ...Completed) *redisresults {
			attempts++
			if attempts == 1 {
				return &redisresults{s: []RedisResult{newResult(strmsg('-', "LOADING Redis is loading the dataset in memory"), nil)}}
			}
			return &redisresults{s: []RedisResult{newResult(strmsg('+', "OK"), nil)}}
		}
		m.AcquireFn = func() wire { return &mockWire{DoMultiFn: m.DoMultiFn} }

		err := client.Dedicated(func(c DedicatedClient) error {
			resps := c.DoMulti(context.Background(), c.B().Get().Key("test").Build())
			if len(resps) != 1 {
				t.Fatalf("unexpected response length %v", len(resps))
			}
			if v, err := resps[0].ToString(); err != nil || v != "OK" {
				t.Fatalf("unexpected response %v %v", v, err)
			}
			return nil
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
	})
}

func TestClusterClientMovedRetry(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())

	setup := func() (*clusterClient, *mockConn) {
		m := &mockConn{
			DoFn: func(cmd Completed) RedisResult {
				if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
					return slotsMultiResp
				}
				return RedisResult{}
			},
		}
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn { return m },
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		return client, m
	}

	t.Run("DoMulti Retry on MOVED", func(t *testing.T) {
		client, m := setup()

		attempts := 0
		m.DoMultiFn = func(multi ...Completed) *redisresults {
			attempts++
			if attempts == 1 {
				return &redisresults{s: []RedisResult{newResult(strmsg('-', "MOVED 0 127.0.0.1"), nil)}}
			}
			return &redisresults{s: []RedisResult{newResult(strmsg('+', "OK"), nil)}}
		}

		cmd := client.B().Set().Key("test").Value(`test`).Build()
		resps := client.DoMulti(context.Background(), cmd)
		if len(resps) != 1 {
			t.Fatalf("unexpected response length %v", len(resps))
		}
		if v, err := resps[0].ToString(); err != nil || v != "OK" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("DoMulti Retry on MOVED for EXEC", func(t *testing.T) {
		client, m := setup()

		attempts := 0
		m.DoMultiFn = func(multi ...Completed) *redisresults {
			attempts++

			results := make([]RedisResult, len(multi))
			for i, cmd := range multi {
				cmdName := cmd.Commands()[0]
				if cmdName == "MULTI" {
					results[i] = newResult(strmsg('+', "OK"), nil)
				} else if cmdName == "EXEC" {
					if attempts == 1 {
						// return MOVED error only for EXEC on first attempt
						// making sure we do not panic on wslots connection reassignment because EXEC's "slot" is 16384 (InitSlot)
						results[i] = newResult(strmsg('-', "MOVED 1 127.0.0.1"), nil)
					} else {
						// return a successful transaction result for the retry
						results[i] = newResult(slicemsg('*', []RedisMessage{strmsg('+', "some_value")}), nil)
					}
				} else {
					results[i] = newResult(strmsg('+', "QUEUED"), nil)
				}
			}
			return &redisresults{s: results}
		}

		cmds := []Completed{
			client.B().Multi().Build(),
			client.B().Get().Key("some_key").Build(),
			client.B().Exec().Build(),
		}
		resps := client.DoMulti(context.Background(), cmds...)

		if attempts != 2 {
			t.Fatalf("expected 2 attempts, got %d", attempts)
		}

		if len(resps) != 3 {
			t.Fatalf("unexpected response length %v", len(resps))
		}

		if vs, err := resps[2].AsStrSlice(); err != nil || vs[0] != "some_value" {
			t.Fatalf("unexpected response %v %v", vs, err)
		}
	})

	t.Run("DoMulti Retry on ASK", func(t *testing.T) {
		client, m := setup()

		attempts := 0
		m.DoMultiFn = func(multi ...Completed) *redisresults {
			attempts++
			if attempts == 1 {
				return &redisresults{s: []RedisResult{newResult(strmsg('-', "ASK 0 127.0.0.1"), nil)}}
			}
			return &redisresults{s: []RedisResult{newResult(strmsg('+', "OK"), nil), newResult(strmsg('+', "OK"), nil)}}
		}

		cmd := client.B().Set().Key("test").Value(`test`).Build()
		resps := client.DoMulti(context.Background(), cmd)
		if len(resps) != 1 {
			t.Fatalf("unexpected response length %v", len(resps))
		}
		if v, err := resps[0].ToString(); err != nil || v != "OK" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})
}

func TestClusterClientCacheASKRetry(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())

	setup := func() (*clusterClient, *mockConn) {
		m := &mockConn{
			DoFn: func(cmd Completed) RedisResult {
				if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
					return slotsMultiResp
				}
				return RedisResult{}
			},
		}
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}},
			func(dst string, opt *ClientOption) conn { return m },
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		return client, m
	}

	t.Run("DoCache Retry on ASK", func(t *testing.T) {
		client, m := setup()
		attempts := 0
		m.DoCacheFn = func(cmd Cacheable, ttl time.Duration) RedisResult {
			return newResult(strmsg('-', "ASK 0 :0"), nil)
		}
		m.DoMultiFn = func(multi ...Completed) *redisresults {
			attempts++
			if attempts == 1 {
				return &redisresults{s: []RedisResult{{}, {}, {}, {}, newResult(strmsg('-', "ASK 0 :0"), nil), newResult(RedisMessage{typ: '_'}, nil)}}
			}
			return &redisresults{s: []RedisResult{{}, {}, {}, {}, {}, newResult(slicemsg('*', []RedisMessage{{}, {}, {}, {}, {}, strmsg('+', "OK")}), nil)}}
		}
		resp := client.DoCache(context.Background(), client.B().Get().Key("a1").Cache(), 10*time.Second)
		if v, err := resp.ToString(); err != nil || v != "OK" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
		if attempts != 2 {
			t.Fatalf("expected 2 attempts, got %v", attempts)
		}
	})

	t.Run("DoMultiCache Retry on ASK", func(t *testing.T) {
		client, m := setup()

		attempts := 0
		m.DoMultiCacheFn = func(multi ...CacheableTTL) *redisresults {
			return &redisresults{s: []RedisResult{newResult(strmsg('-', "ASK 0 :0"), nil)}}
		}
		m.DoMultiFn = func(multi ...Completed) *redisresults {
			attempts++
			if attempts == 1 {
				return &redisresults{s: []RedisResult{{}, {}, {}, {}, newResult(strmsg('-', "ASK 0 :0"), nil), newResult(RedisMessage{typ: '_'}, nil)}}
			}
			return &redisresults{s: []RedisResult{{}, {}, {}, {}, {}, newResult(slicemsg('*', []RedisMessage{{}, {}, {}, {}, {}, strmsg('+', "OK")}), nil)}}
		}
		resps := client.DoMultiCache(context.Background(), CT(client.B().Get().Key("a1").Cache(), 10*time.Second))
		if v, err := resps[0].ToString(); err != nil || v != "OK" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
		if attempts != 2 {
			t.Fatalf("expected 2 attempts, got %v", attempts)
		}
	})
}

//gocyclo:ignore
func TestClusterClient_SendReadOperationToReplicaNodeWriteOperationToPrimaryNode(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())

	primaryNodeConn := &mockConn{
		DoOverride: map[string]func(cmd Completed) RedisResult{
			"CLUSTER SLOTS": func(cmd Completed) RedisResult {
				return slotsMultiResp
			},
			"INFO": func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "INFO"), nil)
			},
			"SET Do V": func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "SET Do V"), nil)
			},
			"SET K2{a} V2{a}": func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "SET K2{a} V2{a}"), nil)
			},
		},
		DoMultiFn: func(multi ...Completed) *redisresults {
			resps := make([]RedisResult, len(multi))
			for i, cmd := range multi {
				if strings.HasPrefix(strings.Join(cmd.Commands(), " "), "SET K1") {
					resps[i] = newResult(strmsg('+', strings.Join(cmd.Commands(), " ")), nil)
					continue
				}
				if strings.HasPrefix(strings.Join(cmd.Commands(), " "), "SET K2") {
					resps[i] = newResult(strmsg('+', strings.Join(cmd.Commands(), " ")), nil)
					continue
				}
				if strings.HasPrefix(strings.Join(cmd.Commands(), " "), "MULTI") {
					resps[i] = newResult(strmsg('+', "MULTI"), nil)
					continue
				}
				if strings.HasPrefix(strings.Join(cmd.Commands(), " "), "EXEC") {
					resps[i] = newResult(strmsg('+', "EXEC"), nil)
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
				return newResult(strmsg('+', "GET Do"), nil)
			},
			"GET K1{a}": func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "GET K1{a}"), nil)
			},
			"GET K2{a}": func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "GET K2{a}"), nil)
			},
		},
		DoMultiFn: func(multi ...Completed) *redisresults {
			resps := make([]RedisResult, len(multi))
			for i, cmd := range multi {
				if strings.HasPrefix(strings.Join(cmd.Commands(), " "), "GET K1") {
					resps[i] = newResult(strmsg('+', strings.Join(cmd.Commands(), " ")), nil)
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
				return newResult(strmsg('+', "GET DoCache"), nil)
			},
			"GET K1{a}": func(cmd Cacheable, ttl time.Duration) RedisResult {
				return newResult(strmsg('+', "GET K1{a}"), nil)
			},
			"GET K2{a}": func(cmd Cacheable, ttl time.Duration) RedisResult {
				return newResult(strmsg('+', "GET K2{a}"), nil)
			},
		},
		DoMultiCacheFn: func(multi ...CacheableTTL) *redisresults {
			resps := make([]RedisResult, len(multi))
			for i, cmd := range multi {
				if strings.HasPrefix(strings.Join(cmd.Cmd.Commands(), " "), "GET K1") {
					resps[i] = newResult(strmsg('+', strings.Join(cmd.Cmd.Commands(), " ")), nil)
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
				return cmd.IsReadOnly()
			},
			ReplicaSelector: func(slot uint16, replicas []NodeInfo) int {
				return 0
			},
		},
		func(dst string, opt *ClientOption) conn {
			if dst == "127.0.0.1:0" || dst == "127.0.2.1:0" { // primary node
				return primaryNodeConn
			} else { // replica node
				return replicaNodeConn
			}
		},
		newRetryer(defaultRetryDelayFn),
	)
	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}

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
						newResult(strmsg('+', "OK"), nil),
						newResult(strmsg('+', "OK"), nil),
						newResult(slicemsg('*', []RedisMessage{strmsg('+', "a")}), nil),
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
		primaryNodeConn.AcquireFn = func() wire { return w }
		replicaNodeConn.AcquireFn = func() wire { return w } // Subscribe can work on replicas
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
				return newResult(strmsg('+', "Delegate"), nil)
			},
			DoMultiFn: func(cmd ...Completed) *redisresults {
				if len(cmd) == 4 {
					return &redisresults{s: []RedisResult{
						newResult(strmsg('+', "OK"), nil),
						newResult(strmsg('+', "OK"), nil),
						newResult(strmsg('+', "OK"), nil),
						newResult(slicemsg('*', []RedisMessage{
							strmsg('+', "Delegate0"),
							strmsg('+', "Delegate1"),
						}), nil),
					}}
				}
				return &redisresults{s: []RedisResult{
					newResult(strmsg('+', "Delegate0"), nil),
					newResult(strmsg('+', "Delegate1"), nil),
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
			)[3].val.values() {
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
			t.Fatalf("Dedicated doesn't put back the wire")
		}
		if !closed {
			t.Fatalf("Dedicated doesn't delegate Close")
		}
	})

	t.Run("Dedicate", func(t *testing.T) {
		closed := false
		w := &mockWire{
			DoFn: func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "Delegate"), nil)
			},
			DoMultiFn: func(cmd ...Completed) *redisresults {
				if len(cmd) == 4 {
					return &redisresults{s: []RedisResult{
						newResult(strmsg('+', "OK"), nil),
						newResult(strmsg('+', "OK"), nil),
						newResult(strmsg('+', "OK"), nil),
						newResult(slicemsg('*', []RedisMessage{
							strmsg('+', "Delegate0"),
							strmsg('+', "Delegate1"),
						}), nil),
					}}
				}
				return &redisresults{s: []RedisResult{
					newResult(strmsg('+', "Delegate0"), nil),
					newResult(strmsg('+', "Delegate1"), nil),
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
		)[3].val.values() {
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
			t.Fatalf("Dedicated doesn't put back the wire")
		}
		if !closed {
			t.Fatalf("Dedicated doesn't delegate Close")
		}
	})
}

//gocyclo:ignore
func TestClusterClient_SendToOnlyPrimaryNodeWhenPrimaryNodeSelected(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())

	primaryNodeConn := &mockConn{
		DoOverride: map[string]func(cmd Completed) RedisResult{
			"CLUSTER SLOTS": func(cmd Completed) RedisResult {
				return slotsMultiResp
			},
			"GET Do": func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "GET Do"), nil)
			},
			"GET K1{a}": func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "GET K1{a}"), nil)
			},
			"GET K2{a}": func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "GET K2{a}"), nil)
			},
			"INFO": func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "INFO"), nil)
			},
		},
		DoMultiFn: func(multi ...Completed) *redisresults {
			resps := make([]RedisResult, len(multi))
			for i, cmd := range multi {
				resps[i] = newResult(strmsg('+', strings.Join(cmd.Commands(), " ")), nil)
			}
			return &redisresults{s: resps}
		},
		DoCacheOverride: map[string]func(cmd Cacheable, ttl time.Duration) RedisResult{
			"GET DoCache": func(cmd Cacheable, ttl time.Duration) RedisResult {
				return newResult(strmsg('+', "GET DoCache"), nil)
			},
			"GET K1{a}": func(cmd Cacheable, ttl time.Duration) RedisResult {
				return newResult(strmsg('+', "GET K1{a}"), nil)
			},
			"GET K2{a}": func(cmd Cacheable, ttl time.Duration) RedisResult {
				return newResult(strmsg('+', "GET K2{a}"), nil)
			},
		},
		DoMultiCacheFn: func(multi ...CacheableTTL) *redisresults {
			resps := make([]RedisResult, len(multi))
			for i, cmd := range multi {
				resps[i] = newResult(strmsg('+', strings.Join(cmd.Cmd.Commands(), " ")), nil)
			}
			return &redisresults{s: resps}
		},
	}
	replicaNodeConn := &mockConn{}

	client, err := newClusterClient(
		&ClientOption{
			InitAddress: []string{"127.0.0.1:0"},
			SendToReplicas: func(cmd Completed) bool {
				return true
			},
			ReplicaSelector: func(slot uint16, replicas []NodeInfo) int {
				return -1
			},
		},
		func(dst string, opt *ClientOption) conn {
			if dst == "127.0.0.1:0" || dst == "127.0.2.1:0" { // primary nodes
				return primaryNodeConn
			} else { // replica nodes
				return replicaNodeConn
			}
		},
		newRetryer(defaultRetryDelayFn),
	)
	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}

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
						newResult(strmsg('+', "OK"), nil),
						newResult(strmsg('+', "OK"), nil),
						newResult(slicemsg('*', []RedisMessage{strmsg('+', "a")}), nil),
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
				return newResult(strmsg('+', "Delegate"), nil)
			},
			DoMultiFn: func(cmd ...Completed) *redisresults {
				if len(cmd) == 4 {
					return &redisresults{s: []RedisResult{
						newResult(strmsg('+', "OK"), nil),
						newResult(strmsg('+', "OK"), nil),
						newResult(strmsg('+', "OK"), nil),
						newResult(slicemsg('*', []RedisMessage{
							strmsg('+', "Delegate0"),
							strmsg('+', "Delegate1"),
						}), nil),
					}}
				}
				return &redisresults{s: []RedisResult{
					newResult(strmsg('+', "Delegate0"), nil),
					newResult(strmsg('+', "Delegate1"), nil),
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
			)[3].val.values() {
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
			t.Fatalf("Dedicated doesn't put back the wire")
		}
		if !closed {
			t.Fatalf("Dedicated doesn't delegate Close")
		}
	})

	t.Run("Dedicate", func(t *testing.T) {
		closed := false
		w := &mockWire{
			DoFn: func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "Delegate"), nil)
			},
			DoMultiFn: func(cmd ...Completed) *redisresults {
				if len(cmd) == 4 {
					return &redisresults{s: []RedisResult{
						newResult(strmsg('+', "OK"), nil),
						newResult(strmsg('+', "OK"), nil),
						newResult(strmsg('+', "OK"), nil),
						newResult(slicemsg('*', []RedisMessage{
							strmsg('+', "Delegate0"),
							strmsg('+', "Delegate1"),
						}), nil),
					}}
				}
				return &redisresults{s: []RedisResult{
					newResult(strmsg('+', "Delegate0"), nil),
					newResult(strmsg('+', "Delegate1"), nil),
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
		)[3].val.values() {
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
			t.Fatalf("Dedicated doesn't put back the wire")
		}
		if !closed {
			t.Fatalf("Dedicated doesn't delegate Close")
		}
	})
}

func TestClusterClientConnLifetime(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())

	setup := func() (*clusterClient, *mockConn) {
		m := &mockConn{
			DoFn: func(cmd Completed) RedisResult {
				if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
					return slotsMultiResp
				}
				return newResult(strmsg('+', "OK"), nil)
			},
		}
		client, err := newClusterClient(
			&ClientOption{InitAddress: []string{":0"}, ConnLifetime: 5 * time.Second},
			func(dst string, opt *ClientOption) conn { return m },
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected error %v", err)
		}
		return client, m
	}

	t.Run("Do ConnLifetime", func(t *testing.T) {
		client, m := setup()
		var attempts int64
		m.DoFn = func(cmd Completed) RedisResult {
			if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
				return slotsMultiResp
			}
			if atomic.AddInt64(&attempts, 1) == 1 {
				return newErrResult(errConnExpired)
			}
			return newResult(strmsg('+', "OK"), nil)
		}
		c := client.B().Get().Key("Do").Build()
		if v, err := client.Do(context.Background(), c).ToString(); err != nil || v != "OK" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("Do ConnLifetime in MOVE", func(t *testing.T) {
		client, m := setup()
		var attempts int64
		m.DoFn = func(cmd Completed) RedisResult {
			if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
				return slotsMultiResp
			}
			switch atomic.AddInt64(&attempts, 1) {
			case 1:
				return newResult(strmsg('-', "MOVED 0 :1"), nil)
			case 2:
				return newErrResult(errConnExpired)
			}
			return newResult(strmsg('+', "OK"), nil)
		}
		c := client.B().Get().Key("Do").Build()
		if v, err := client.Do(context.Background(), c).ToString(); err != nil || v != "OK" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("Do ConnLifetime in ASK", func(t *testing.T) {
		client, m := setup()
		var attempts int64
		m.DoFn = func(cmd Completed) RedisResult {
			if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
				return slotsMultiResp
			}
			return newResult(strmsg('-', "ASK 0 :0"), nil)
		}
		m.DoMultiFn = func(multi ...Completed) *redisresults {
			if atomic.AddInt64(&attempts, 1) == 1 {
				return &redisresults{s: []RedisResult{newResult(strmsg('+', "OK"), nil), newErrResult(errConnExpired)}}
			}
			return &redisresults{s: []RedisResult{newResult(strmsg('+', "OK"), nil), newResult(strmsg('+', "OK"), nil)}}
		}
		c := client.B().Get().Key("Do").Build()
		if v, err := client.Do(context.Background(), c).ToString(); err != nil || v != "OK" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("DoCache ConnLifetime", func(t *testing.T) {
		client, m := setup()
		var attempts int64
		m.DoFn = func(cmd Completed) RedisResult { return slotsMultiResp }
		m.DoCacheFn = func(cmd Cacheable, ttl time.Duration) RedisResult {
			if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
				return slotsMultiResp
			}
			if atomic.AddInt64(&attempts, 1) == 1 {
				return newErrResult(errConnExpired)
			}
			return newResult(strmsg('+', "OK"), nil)
		}
		if v, err := client.DoCache(context.Background(), client.B().Get().Key("Do").Cache(), 0).ToString(); err != nil || v != "OK" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("DoCache ConnLifetime in MOVE", func(t *testing.T) {
		client, m := setup()
		var attempts int64
		m.DoFn = func(cmd Completed) RedisResult { return slotsMultiResp }
		m.DoCacheFn = func(cmd Cacheable, ttl time.Duration) RedisResult {
			if strings.Join(cmd.Commands(), " ") == "CLUSTER SLOTS" {
				return slotsMultiResp
			}
			switch atomic.AddInt64(&attempts, 1) {
			case 1:
				return newResult(strmsg('-', "MOVED 0 :1"), nil)
			case 2:
				return newErrResult(errConnExpired)
			}
			return newResult(strmsg('+', "OK"), nil)
		}
		if v, err := client.DoCache(context.Background(), client.B().Get().Key("Do").Cache(), 0).ToString(); err != nil || v != "OK" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("DoCache ConnLifetime in ASK", func(t *testing.T) {
		client, m := setup()
		var attempts int64
		m.DoFn = func(cmd Completed) RedisResult { return slotsMultiResp }
		m.DoCacheFn = func(cmd Cacheable, ttl time.Duration) RedisResult {
			return newResult(strmsg('-', "ASK 0 :0"), nil)
		}
		m.DoMultiFn = func(multi ...Completed) *redisresults {
			if atomic.AddInt64(&attempts, 1) == 1 {
				return &redisresults{s: []RedisResult{{}, newErrResult(errConnExpired), newErrResult(errConnExpired), newErrResult(errConnExpired), newErrResult(errConnExpired), newErrResult(errConnExpired)}}
			}
			return &redisresults{s: []RedisResult{{}, {}, {}, {}, {}, newResult(slicemsg('*', []RedisMessage{{}, strmsg('+', "OK")}), nil)}}
		}
		if v, err := client.DoCache(context.Background(), client.B().Get().Key("Do").Cache(), 0).ToString(); err != nil || v != "OK" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("DoMulti ConnLifetime - at the head of processing", func(t *testing.T) {
		client, m := setup()
		var attempts int64
		m.DoFn = func(cmd Completed) RedisResult { return slotsMultiResp }
		m.DoMultiFn = func(multi ...Completed) *redisresults {
			if atomic.AddInt64(&attempts, 1) == 1 {
				return &redisresults{s: []RedisResult{newErrResult(errConnExpired)}}
			}
			return &redisresults{s: []RedisResult{newResult(strmsg('+', "OK"), nil)}}
		}
		if v, err := client.DoMulti(context.Background(), client.B().Get().Key("Do").Build())[0].ToString(); err != nil || v != "OK" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("DoMulti ConnLifetime - in the middle of processing", func(t *testing.T) {
		client, m := setup()
		var attempts int64
		m.DoFn = func(cmd Completed) RedisResult { return slotsMultiResp }
		m.DoMultiFn = func(multi ...Completed) *redisresults {
			if atomic.AddInt64(&attempts, 1) == 1 {
				return &redisresults{s: []RedisResult{newResult(strmsg('+', "OK"), nil), newErrResult(errConnExpired)}}
			}
			// recover the failure of the first call
			return &redisresults{s: []RedisResult{newResult(strmsg('+', "OK"), nil)}}
		}
		resps := client.DoMulti(context.Background(), client.B().Get().Key("Do").Build(), client.B().Get().Key("Do").Build())
		if len(resps) != 2 {
			t.Errorf("unexpected response length %v", len(resps))
		}
		for _, resp := range resps {
			if v, err := resp.ToString(); err != nil || v != "OK" {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		}
	})

	t.Run("DoMulti ConnLifetime Transaction Block", func(t *testing.T) {
		client, m := setup()
		var (
			attempts int64
			orgMulti []Completed
		)
		m.DoFn = func(cmd Completed) RedisResult { return slotsMultiResp }
		m.DoMultiFn = func(multi ...Completed) *redisresults {
			switch atomic.AddInt64(&attempts, 1) {
			case 1: // errConnExpired at the head of processing
				orgMulti = multi
				return &redisresults{s: []RedisResult{newErrResult(errConnExpired), newErrResult(errConnExpired), newErrResult(errConnExpired), newErrResult(errConnExpired), newErrResult(errConnExpired), newErrResult(errConnExpired)}}
			case 2: // errConnExpired at Multi command
				if len(multi) != 6 || !reflect.DeepEqual(multi[0].Commands(), orgMulti[0].Commands()) {
					t.Fatalf("unexpected multi when errConnExpired occurred at the head of processing, %v", multi)
				}
				return &redisresults{s: []RedisResult{
					newResult(strmsg('+', "1"), nil),
					newErrResult(errConnExpired), newErrResult(errConnExpired), newErrResult(errConnExpired), newErrResult(errConnExpired), newErrResult(errConnExpired),
				}}
			case 3: // errConnExpired in the middle of transaction block
				if len(multi) != 5 || !reflect.DeepEqual(multi[0].Commands(), orgMulti[1].Commands()) {
					t.Fatalf("unexpected multi when errConnExpired occurred at Multi Command, %v", multi)
				}
				return &redisresults{s: []RedisResult{
					newResult(strmsg('+', "OK"), nil),
					newResult(strmsg('+', "QUEUE"), nil),
					newErrResult(errConnExpired), newErrResult(errConnExpired), newErrResult(errConnExpired),
				}}
			case 4: // errConnExpired at Exec Command
				if len(multi) != 5 || !reflect.DeepEqual(multi[0].Commands(), orgMulti[1].Commands()) {
					t.Fatalf("unexpected multi when errConnExpired occurred in the middle of transaction block, %v", multi)
				}
				return &redisresults{s: []RedisResult{
					newResult(strmsg('+', "OK"), nil),
					newResult(strmsg('+', "QUEUE"), nil),
					newResult(strmsg('+', "QUEUE"), nil),
					newErrResult(errConnExpired), newErrResult(errConnExpired),
				}}
			case 5: // errConnExpired at end of processing
				if len(multi) != 5 || !reflect.DeepEqual(multi[0].Commands(), orgMulti[1].Commands()) {
					t.Fatalf("unexpected multi when errConnExpired occurred at at Exec Command, %v", multi)
				}
				return &redisresults{s: []RedisResult{
					newResult(strmsg('+', "OK"), nil),
					newResult(strmsg('+', "QUEUE"), nil),
					newResult(strmsg('+', "QUEUE"), nil),
					newResult(slicemsg('*', []RedisMessage{
						strmsg('+', "2"),
						strmsg('+', "3"),
					}), nil),
					newErrResult(errConnExpired),
				}}
			case 6:
				if len(multi) != 1 || !reflect.DeepEqual(multi[0].Commands(), orgMulti[5].Commands()) {
					t.Fatalf("unexpected multi when errConnExpired occurred at end of processing, %v", multi)
				}
			}
			return &redisresults{s: []RedisResult{newResult(strmsg('+', "4"), nil)}}
		}
		multi := []Completed{
			client.B().Get().Key("1{t}").Build(),
			client.B().Multi().Build(),
			client.B().Incr().Key("2{t}").Build(),
			client.B().Incr().Key("3{t}").Build(),
			client.B().Exec().Build(),
			client.B().Get().Key("4{t}").Build(),
		}
		resps := client.DoMulti(context.Background(), multi...)
		if len(resps) != 6 {
			t.Fatalf("unexpected response length %v", len(resps))
		}
		if v, err := resps[0].ToString(); err != nil || v != "1" {
			t.Errorf("unexpected response %v %v", v, err)
		}
		if v, err := resps[1].ToString(); err != nil || v != "OK" {
			t.Errorf("unexpected response %v %v", v, err)
		}
		if v, err := resps[2].ToString(); err != nil || v != "QUEUE" {
			t.Errorf("unexpected response %v %v", v, err)
		}
		if v, err := resps[3].ToString(); err != nil || v != "QUEUE" {
			t.Errorf("unexpected response %v %v", v, err)
		}
		if v, err := resps[4].AsStrSlice(); err != nil || !reflect.DeepEqual(v, []string{"2", "3"}) {
			t.Errorf("unexpected response %v %v", v, err)
		}
		if v, err := resps[5].ToString(); err != nil || v != "4" {
			t.Errorf("unexpected response %v %v", v, err)
		}
	})

	t.Run("DoMulti ConnLifetime in ASK - in the middle of processing", func(t *testing.T) {
		client, m := setup()
		var attempts int64
		m.DoFn = func(cmd Completed) RedisResult { return slotsMultiResp }
		m.DoMultiFn = func(multi ...Completed) *redisresults {
			ret := make([]RedisResult, len(multi))
			switch atomic.AddInt64(&attempts, 1) {
			case 1:
				for i := range ret {
					ret[i] = newResult(strmsg('-', "ASK 0 :1"), nil)
				}
				return &redisresults{s: ret}
			case 2:
				return &redisresults{s: []RedisResult{newResult(strmsg('+', "OK"), nil), newErrResult(errConnExpired), newErrResult(errConnExpired), newErrResult(errConnExpired)}}
			}
			for i := 0; i < len(multi); i += 2 {
				ret[i] = newResult(strmsg('+', "OK"), nil)
				ret[i+1] = newResult(strmsg('+', multi[i+1].Commands()[1]), nil)
			}
			return &redisresults{s: ret}
		}
		resps := client.DoMulti(context.Background(), client.B().Get().Key("a").Build(), client.B().Get().Key("b").Build())
		if len(resps) != 2 {
			t.Errorf("unexpected response length %v", len(resps))
		}
		for i, resp := range resps {
			v, err := resp.ToString()
			if err != nil || i == 0 && v != "a" {
				t.Fatalf("unexpected response %v %v", v, err)
			}
			if err != nil || i == 1 && v != "b" {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		}
	})

	t.Run("DoMulti ConnLifetime Transaction Block in ASK", func(t *testing.T) {
		client, m := setup()
		var (
			attempts int64
			orgMulti []Completed
		)
		m.DoFn = func(cmd Completed) RedisResult { return slotsMultiResp }
		m.DoMultiFn = func(multi ...Completed) *redisresults {
			switch atomic.AddInt64(&attempts, 1) {
			case 1:
				ret := make([]RedisResult, len(multi))
				for i := range ret {
					if isMulti(multi[i]) || isExec(multi[i]) {
						ret[i] = newResult(strmsg('+', "OK"), nil)
						continue
					}
					ret[i] = newResult(strmsg('-', "ASK 0 :1"), nil)
				}
				return &redisresults{s: ret}
			case 2: // errConnExpired at the head of processing
				orgMulti = multi
				return &redisresults{s: []RedisResult{newErrResult(errConnExpired), newErrResult(errConnExpired), newErrResult(errConnExpired), newErrResult(errConnExpired), newErrResult(errConnExpired), newErrResult(errConnExpired), newErrResult(errConnExpired), newErrResult(errConnExpired), newErrResult(errConnExpired)}}
			case 3: // errConnExpired at Asking command before Multi command
				if len(multi) != 9 || !reflect.DeepEqual(multi[0].Commands(), orgMulti[0].Commands()) {
					t.Fatalf("unexpected multi when errConnExpired occurred at the head of processing, %v", multi)
				}
				return &redisresults{s: []RedisResult{
					newResult(strmsg('+', "OK"), nil),
					newResult(strmsg('+', "1"), nil),
					newErrResult(errConnExpired), newErrResult(errConnExpired), newErrResult(errConnExpired), newErrResult(errConnExpired), newErrResult(errConnExpired), newErrResult(errConnExpired), newErrResult(errConnExpired),
				}}
			case 4: // errConnExpired at Multi command
				if len(multi) != 7 || !reflect.DeepEqual(multi[0].Commands(), orgMulti[2].Commands()) {
					t.Fatalf("unexpected multi when errConnExpired occurred at Asking command before Multi command, %v", multi)
				}
				return &redisresults{s: []RedisResult{
					newResult(strmsg('+', "OK"), nil),
					newErrResult(errConnExpired), newErrResult(errConnExpired), newErrResult(errConnExpired), newErrResult(errConnExpired), newErrResult(errConnExpired), newErrResult(errConnExpired),
				}}
			case 5: // errConnExpired in the middle of transaction block
				if len(multi) != 7 || !reflect.DeepEqual(multi[0].Commands(), orgMulti[2].Commands()) {
					t.Fatalf("unexpected multi when errConnExpired occurred at Multi Command, %v", multi)
				}
				return &redisresults{s: []RedisResult{
					newResult(strmsg('+', "OK"), nil),
					newResult(strmsg('+', "OK"), nil),
					newResult(strmsg('+', "QUEUE"), nil),
					newErrResult(errConnExpired), newErrResult(errConnExpired), newErrResult(errConnExpired), newErrResult(errConnExpired),
				}}
			case 6: // errConnExpired at Exec Command
				if len(multi) != 7 || !reflect.DeepEqual(multi[0].Commands(), orgMulti[2].Commands()) {
					t.Fatalf("unexpected multi when errConnExpired occurred in the middle of transaction block, %v", multi)
				}
				return &redisresults{s: []RedisResult{
					newResult(strmsg('+', "OK"), nil),
					newResult(strmsg('+', "QUEUE"), nil),
					newResult(strmsg('+', "QUEUE"), nil),
					newErrResult(errConnExpired), newErrResult(errConnExpired), newErrResult(errConnExpired),
				}}
			case 7: // errConnExpired at end of processing
				if len(multi) != 7 || !reflect.DeepEqual(multi[0].Commands(), orgMulti[2].Commands()) {
					t.Fatalf("unexpected multi when errConnExpired occurred at at Exec Command, %v", multi)
				}
				return &redisresults{s: []RedisResult{
					newResult(strmsg('+', "OK"), nil),
					newResult(strmsg('+', "QUEUE"), nil),
					newResult(strmsg('+', "QUEUE"), nil),
					newResult(slicemsg('*', []RedisMessage{
						strmsg('+', "2"),
						strmsg('+', "3"),
					}), nil),
					newResult(strmsg('+', "OK"), nil),
					newErrResult(errConnExpired),
				}}
			case 8:
				if len(multi) != 2 || !reflect.DeepEqual(multi[0].Commands(), orgMulti[7].Commands()) {
					t.Fatalf("unexpected multi when errConnExpired occurred at end of processing, %v", multi)
				}
			}
			return &redisresults{s: []RedisResult{newResult(strmsg('+', "4"), nil)}}
		}
		multi := []Completed{
			client.B().Get().Key("1{t}").Build(),
			client.B().Multi().Build(),
			client.B().Incr().Key("2{t}").Build(),
			client.B().Incr().Key("3{t}").Build(),
			client.B().Exec().Build(),
			client.B().Get().Key("4{t}").Build(),
		}
		resps := client.DoMulti(context.Background(), multi...)
		if len(resps) != 6 {
			t.Fatalf("unexpected response length %v", len(resps))
		}
		if v, err := resps[0].ToString(); err != nil || v != "1" {
			t.Errorf("unexpected response %v %v", v, err)
		}
		if v, err := resps[1].ToString(); err != nil || v != "OK" {
			t.Errorf("unexpected response %v %v", v, err)
		}
		if v, err := resps[2].ToString(); err != nil || v != "QUEUE" {
			t.Errorf("unexpected response %v %v", v, err)
		}
		if v, err := resps[3].ToString(); err != nil || v != "QUEUE" {
			t.Errorf("unexpected response %v %v", v, err)
		}
		if v, err := resps[4].AsStrSlice(); err != nil || !reflect.DeepEqual(v, []string{"2", "3"}) {
			t.Errorf("unexpected response %v %v", v, err)
		}
		if v, err := resps[5].ToString(); err != nil || v != "4" {
			t.Errorf("unexpected response %v %v", v, err)
		}
	})

	t.Run("DoMultiCache ConnLifetime - at the head of processing", func(t *testing.T) {
		client, m := setup()
		var attempts int64
		m.DoMultiCacheFn = func(multi ...CacheableTTL) *redisresults {
			if atomic.AddInt64(&attempts, 1) == 1 {
				return &redisresults{s: []RedisResult{newErrResult(errConnExpired)}}
			}
			// recover the failure of the first call
			return &redisresults{s: []RedisResult{newResult(strmsg('+', "OK"), nil)}}
		}
		if v, err := client.DoMultiCache(context.Background(), CT(client.B().Get().Key("Do").Cache(), 0))[0].ToString(); err != nil || v != "OK" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("DoMultiCache ConnLifetime in the middle of processing", func(t *testing.T) {
		client, m := setup()
		var attempts int64
		m.DoMultiCacheFn = func(multi ...CacheableTTL) *redisresults {
			if atomic.AddInt64(&attempts, 1) == 1 {
				return &redisresults{s: []RedisResult{newResult(strmsg('+', "OK"), nil), newErrResult(errConnExpired)}}
			}
			// recover the failure of the first call
			return &redisresults{s: []RedisResult{newResult(strmsg('+', "OK"), nil)}}
		}
		resps := client.DoMultiCache(context.Background(), CT(client.B().Get().Key("Do").Cache(), 0), CT(client.B().Get().Key("Do").Cache(), 0))
		if len(resps) != 2 {
			t.Errorf("unexpected response length %v", len(resps))
		}
		for _, resp := range resps {
			if v, err := resp.ToString(); err != nil || v != "OK" {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		}
	})

	t.Run("Receive ConnLifetime", func(t *testing.T) {
		client, m := setup()
		c := client.B().Subscribe().Channel("ch").Build()
		hdl := func(message PubSubMessage) {}
		var attempts int64
		m.ReceiveFn = func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
			if atomic.AddInt64(&attempts, 1) == 1 {
				return errConnExpired
			}
			return nil
		}
		if err := client.Receive(context.Background(), c, hdl); err != nil {
			t.Fatalf("unexpected response %v", err)
		}
	})
}

//gocyclo:ignore
func TestClusterClient_SendToAlternatePrimaryAndReplicaNodes(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())

	primaryNodeConn := &mockConn{
		DoMultiFn: func(multi ...Completed) *redisresults {
			resps := make([]RedisResult, len(multi))
			for i, cmd := range multi {
				resps[i] = newResult(strmsg('+', strings.Join(cmd.Commands(), " ")), nil)
			}
			return &redisresults{s: resps}
		},
		DoMultiCacheFn: func(multi ...CacheableTTL) *redisresults {
			resps := make([]RedisResult, len(multi))
			for i, cmd := range multi {
				resps[i] = newResult(strmsg('+', strings.Join(cmd.Cmd.Commands(), " ")), nil)
			}
			return &redisresults{s: resps}
		},
		DoOverride: map[string]func(cmd Completed) RedisResult{
			"CLUSTER SLOTS": func(cmd Completed) RedisResult {
				return slotsMultiResp
			},
			"INFO": func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "INFO"), nil)
			},
			"GET K1{a}": func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "GET K1{a}"), nil)
			},
			"GET K2{b}": func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "GET K2{b}"), nil)
			},
		},
	}
	replicaNodeConn := &mockConn{
		DoMultiFn: func(multi ...Completed) *redisresults {
			resps := make([]RedisResult, len(multi))
			for i, cmd := range multi {
				resps[i] = newResult(strmsg('+', strings.Join(cmd.Commands(), " ")), nil)
			}
			return &redisresults{s: resps}
		},
		DoMultiCacheFn: func(multi ...CacheableTTL) *redisresults {
			resps := make([]RedisResult, len(multi))
			for i, cmd := range multi {
				resps[i] = newResult(strmsg('+', strings.Join(cmd.Cmd.Commands(), " ")), nil)
			}
			return &redisresults{s: resps}
		},
		DoOverride: map[string]func(cmd Completed) RedisResult{
			"GET Do": func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "GET Do"), nil)
			},
			"GET K1{a}": func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "GET K1{a}"), nil)
			},
			"GET K2{b}": func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "GET K2{b}"), nil)
			},
		},
		DoCacheOverride: map[string]func(cmd Cacheable, ttl time.Duration) RedisResult{
			"GET DoCache": func(cmd Cacheable, ttl time.Duration) RedisResult {
				return newResult(strmsg('+', "GET DoCache"), nil)
			},
		},
	}

	nextNode := -1
	client, err := newClusterClient(
		&ClientOption{
			InitAddress: []string{"127.0.0.1:0"},
			SendToReplicas: func(cmd Completed) bool {
				nextNode++
				return (nextNode/2)%2 == 0
			},
		},
		func(dst string, opt *ClientOption) conn {
			if dst == "127.0.0.1:0" || dst == "127.0.2.1:0" { // primary nodes
				return primaryNodeConn
			} else { // replica nodes
				return replicaNodeConn
			}
		},
		newRetryer(defaultRetryDelayFn),
	)
	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}

	t.Run("DoMulti Multi Slot", func(t *testing.T) {
		c1 := client.B().Get().Key("K1{a}").Build()
		c2 := client.B().Get().Key("K2{b}").Build()
		resps := client.DoMulti(context.Background(), c1, c2)
		if v, err := resps[0].ToString(); err != nil || v != "GET K1{a}" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
		if v, err := resps[1].ToString(); err != nil || v != "GET K2{b}" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("DoMulti Multi Slot Large", func(t *testing.T) {
		var cmds []Completed
		for i := 0; i < 500; i++ {
			cmds = append(cmds, client.B().Get().Key("K1{a}").Build())
		}
		resps := client.DoMulti(context.Background(), cmds...)
		for _, resp := range resps {
			if v, err := resp.ToString(); err != nil || v != "GET K1{a}" {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		}
	})

	t.Run("DoMultiCache Multi Slot", func(t *testing.T) {
		c1 := client.B().Get().Key("K1{a}").Cache()
		c2 := client.B().Get().Key("K2{b}").Cache()
		resps := client.DoMultiCache(context.Background(), CT(c1, time.Second), CT(c2, time.Second))
		if v, err := resps[0].ToString(); err != nil || v != "GET K1{a}" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
		if v, err := resps[1].ToString(); err != nil || v != "GET K2{b}" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("DoMultiCache Multi Slot Large", func(t *testing.T) {
		var cmds []CacheableTTL
		for i := 0; i < 500; i++ {
			cmds = append(cmds, CT(client.B().Get().Key("K1{a}").Cache(), time.Second))
		}
		resps := client.DoMultiCache(context.Background(), cmds...)
		for _, resp := range resps {
			if v, err := resp.ToString(); err != nil || v != "GET K1{a}" {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		}
	})
}

func TestClusterClient_ReadNodeSelector_SendToOnlyPrimaryNodeWhenPrimaryNodeSelected(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())

	primaryNodeConn := &mockConn{
		DoOverride: map[string]func(cmd Completed) RedisResult{
			"CLUSTER SLOTS": func(cmd Completed) RedisResult {
				return slotsMultiResp
			},
			"GET Do": func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "GET Do"), nil)
			},
			"GET K1{a}": func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "GET K1{a}"), nil)
			},
			"GET K2{a}": func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "GET K2{a}"), nil)
			},
			"INFO": func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "INFO"), nil)
			},
		},
		DoMultiFn: func(multi ...Completed) *redisresults {
			resps := make([]RedisResult, len(multi))
			for i, cmd := range multi {
				resps[i] = newResult(strmsg('+', strings.Join(cmd.Commands(), " ")), nil)
			}
			return &redisresults{s: resps}
		},
		DoCacheOverride: map[string]func(cmd Cacheable, ttl time.Duration) RedisResult{
			"GET DoCache": func(cmd Cacheable, ttl time.Duration) RedisResult {
				return newResult(strmsg('+', "GET DoCache"), nil)
			},
			"GET K1{a}": func(cmd Cacheable, ttl time.Duration) RedisResult {
				return newResult(strmsg('+', "GET K1{a}"), nil)
			},
			"GET K2{a}": func(cmd Cacheable, ttl time.Duration) RedisResult {
				return newResult(strmsg('+', "GET K2{a}"), nil)
			},
		},
		DoMultiCacheFn: func(multi ...CacheableTTL) *redisresults {
			resps := make([]RedisResult, len(multi))
			for i, cmd := range multi {
				resps[i] = newResult(strmsg('+', strings.Join(cmd.Cmd.Commands(), " ")), nil)
			}
			return &redisresults{s: resps}
		},
	}
	replicaNodeConn := &mockConn{}

	client, err := newClusterClient(
		&ClientOption{
			InitAddress: []string{"127.0.0.1:0"},
			SendToReplicas: func(cmd Completed) bool {
				return true
			},
			ReadNodeSelector: func(slot uint16, nodes []NodeInfo) int {
				return 0
			},
		},
		func(dst string, opt *ClientOption) conn {
			if dst == "127.0.0.1:0" || dst == "127.0.2.1:0" { // primary nodes
				return primaryNodeConn
			} else { // replica nodes
				return replicaNodeConn
			}
		},
		newRetryer(defaultRetryDelayFn),
	)
	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}

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
						newResult(strmsg('+', "OK"), nil),
						newResult(strmsg('+', "OK"), nil),
						newResult(slicemsg('*', []RedisMessage{strmsg('+', "a")}), nil),
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
				return newResult(strmsg('+', "Delegate"), nil)
			},
			DoMultiFn: func(cmd ...Completed) *redisresults {
				if len(cmd) == 4 {
					return &redisresults{s: []RedisResult{
						newResult(strmsg('+', "OK"), nil),
						newResult(strmsg('+', "OK"), nil),
						newResult(strmsg('+', "OK"), nil),
						newResult(slicemsg('*', []RedisMessage{
							strmsg('+', "Delegate0"),
							strmsg('+', "Delegate1"),
						}), nil),
					}}
				}
				return &redisresults{s: []RedisResult{
					newResult(strmsg('+', "Delegate0"), nil),
					newResult(strmsg('+', "Delegate1"), nil),
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
			)[3].val.values() {
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
			t.Fatalf("Dedicated doesn't put back the wire")
		}
		if !closed {
			t.Fatalf("Dedicated doesn't delegate Close")
		}
	})

	t.Run("Dedicate", func(t *testing.T) {
		closed := false
		w := &mockWire{
			DoFn: func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "Delegate"), nil)
			},
			DoMultiFn: func(cmd ...Completed) *redisresults {
				if len(cmd) == 4 {
					return &redisresults{s: []RedisResult{
						newResult(strmsg('+', "OK"), nil),
						newResult(strmsg('+', "OK"), nil),
						newResult(strmsg('+', "OK"), nil),
						newResult(slicemsg('*', []RedisMessage{
							strmsg('+', "Delegate0"),
							strmsg('+', "Delegate1"),
						}), nil),
					}}
				}
				return &redisresults{s: []RedisResult{
					newResult(strmsg('+', "Delegate0"), nil),
					newResult(strmsg('+', "Delegate1"), nil),
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
		)[3].val.values() {
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
			t.Fatalf("Dedicated doesn't put back the wire")
		}
		if !closed {
			t.Fatalf("Dedicated doesn't delegate Close")
		}
	})
}

func TestClusterClient_ReadNodeSelector_SendToOnlyPrimaryNodeWhenIndexIsOutOfRange(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())

	primaryNodeConn := &mockConn{
		DoOverride: map[string]func(cmd Completed) RedisResult{
			"CLUSTER SLOTS": func(cmd Completed) RedisResult {
				return slotsMultiResp
			},
			"GET Do": func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "GET Do"), nil)
			},
			"GET K1{a}": func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "GET K1{a}"), nil)
			},
			"GET K2{a}": func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "GET K2{a}"), nil)
			},
			"INFO": func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "INFO"), nil)
			},
		},
		DoMultiFn: func(multi ...Completed) *redisresults {
			resps := make([]RedisResult, len(multi))
			for i, cmd := range multi {
				resps[i] = newResult(strmsg('+', strings.Join(cmd.Commands(), " ")), nil)
			}
			return &redisresults{s: resps}
		},
		DoCacheOverride: map[string]func(cmd Cacheable, ttl time.Duration) RedisResult{
			"GET DoCache": func(cmd Cacheable, ttl time.Duration) RedisResult {
				return newResult(strmsg('+', "GET DoCache"), nil)
			},
			"GET K1{a}": func(cmd Cacheable, ttl time.Duration) RedisResult {
				return newResult(strmsg('+', "GET K1{a}"), nil)
			},
			"GET K2{a}": func(cmd Cacheable, ttl time.Duration) RedisResult {
				return newResult(strmsg('+', "GET K2{a}"), nil)
			},
		},
		DoMultiCacheFn: func(multi ...CacheableTTL) *redisresults {
			resps := make([]RedisResult, len(multi))
			for i, cmd := range multi {
				resps[i] = newResult(strmsg('+', strings.Join(cmd.Cmd.Commands(), " ")), nil)
			}
			return &redisresults{s: resps}
		},
	}
	replicaNodeConn := &mockConn{}

	client, err := newClusterClient(
		&ClientOption{
			InitAddress: []string{"127.0.0.1:0"},
			SendToReplicas: func(cmd Completed) bool {
				return true
			},
			ReadNodeSelector: func(slot uint16, nodes []NodeInfo) int {
				return -1
			},
		},
		func(dst string, opt *ClientOption) conn {
			if dst == "127.0.0.1:0" || dst == "127.0.2.1:0" { // primary nodes
				return primaryNodeConn
			} else { // replica nodes
				return replicaNodeConn
			}
		},
		newRetryer(defaultRetryDelayFn),
	)
	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}

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
						newResult(strmsg('+', "OK"), nil),
						newResult(strmsg('+', "OK"), nil),
						newResult(slicemsg('*', []RedisMessage{strmsg('+', "a")}), nil),
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
				return newResult(strmsg('+', "Delegate"), nil)
			},
			DoMultiFn: func(cmd ...Completed) *redisresults {
				if len(cmd) == 4 {
					return &redisresults{s: []RedisResult{
						newResult(strmsg('+', "OK"), nil),
						newResult(strmsg('+', "OK"), nil),
						newResult(strmsg('+', "OK"), nil),
						newResult(slicemsg('*', []RedisMessage{
							strmsg('+', "Delegate0"),
							strmsg('+', "Delegate1"),
						}), nil),
					}}
				}
				return &redisresults{s: []RedisResult{
					newResult(strmsg('+', "Delegate0"), nil),
					newResult(strmsg('+', "Delegate1"), nil),
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
			)[3].val.values() {
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
			t.Fatalf("Dedicated doesn't put back the wire")
		}
		if !closed {
			t.Fatalf("Dedicated doesn't delegate Close")
		}
	})

	t.Run("Dedicate", func(t *testing.T) {
		closed := false
		w := &mockWire{
			DoFn: func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "Delegate"), nil)
			},
			DoMultiFn: func(cmd ...Completed) *redisresults {
				if len(cmd) == 4 {
					return &redisresults{s: []RedisResult{
						newResult(strmsg('+', "OK"), nil),
						newResult(strmsg('+', "OK"), nil),
						newResult(strmsg('+', "OK"), nil),
						newResult(slicemsg('*', []RedisMessage{
							strmsg('+', "Delegate0"),
							strmsg('+', "Delegate1"),
						}), nil),
					}}
				}
				return &redisresults{s: []RedisResult{
					newResult(strmsg('+', "Delegate0"), nil),
					newResult(strmsg('+', "Delegate1"), nil),
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
		)[3].val.values() {
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
			t.Fatalf("Dedicated doesn't put back the wire")
		}
		if !closed {
			t.Fatalf("Dedicated doesn't delegate Close")
		}
	})
}

func TestClusterClient_ReadNodeSelector_SendReadOperationToReplicaNodeWriteOperationToPrimaryNode(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())

	primaryNodeConn := &mockConn{
		DoOverride: map[string]func(cmd Completed) RedisResult{
			"CLUSTER SLOTS": func(cmd Completed) RedisResult {
				return slotsMultiResp
			},
			"INFO": func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "INFO"), nil)
			},
			"SET Do V": func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "SET Do V"), nil)
			},
			"SET K2{a} V2{a}": func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "SET K2{a} V2{a}"), nil)
			},
		},
		DoMultiFn: func(multi ...Completed) *redisresults {
			resps := make([]RedisResult, len(multi))
			for i, cmd := range multi {
				if strings.HasPrefix(strings.Join(cmd.Commands(), " "), "SET K1") {
					resps[i] = newResult(strmsg('+', strings.Join(cmd.Commands(), " ")), nil)
					continue
				}
				if strings.HasPrefix(strings.Join(cmd.Commands(), " "), "SET K2") {
					resps[i] = newResult(strmsg('+', strings.Join(cmd.Commands(), " ")), nil)
					continue
				}
				if strings.HasPrefix(strings.Join(cmd.Commands(), " "), "MULTI") {
					resps[i] = newResult(strmsg('+', "MULTI"), nil)
					continue
				}
				if strings.HasPrefix(strings.Join(cmd.Commands(), " "), "EXEC") {
					resps[i] = newResult(strmsg('+', "EXEC"), nil)
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
				return newResult(strmsg('+', "GET Do"), nil)
			},
			"GET K1{a}": func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "GET K1{a}"), nil)
			},
			"GET K2{a}": func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "GET K2{a}"), nil)
			},
		},
		DoMultiFn: func(multi ...Completed) *redisresults {
			resps := make([]RedisResult, len(multi))
			for i, cmd := range multi {
				if strings.HasPrefix(strings.Join(cmd.Commands(), " "), "GET K1") {
					resps[i] = newResult(strmsg('+', strings.Join(cmd.Commands(), " ")), nil)
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
				return newResult(strmsg('+', "GET DoCache"), nil)
			},
			"GET K1{a}": func(cmd Cacheable, ttl time.Duration) RedisResult {
				return newResult(strmsg('+', "GET K1{a}"), nil)
			},
			"GET K2{a}": func(cmd Cacheable, ttl time.Duration) RedisResult {
				return newResult(strmsg('+', "GET K2{a}"), nil)
			},
		},
		DoMultiCacheFn: func(multi ...CacheableTTL) *redisresults {
			resps := make([]RedisResult, len(multi))
			for i, cmd := range multi {
				if strings.HasPrefix(strings.Join(cmd.Cmd.Commands(), " "), "GET K1") {
					resps[i] = newResult(strmsg('+', strings.Join(cmd.Cmd.Commands(), " ")), nil)
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
				return cmd.IsReadOnly()
			},
			ReadNodeSelector: func(slot uint16, nodes []NodeInfo) int {
				return 1
			},
		},
		func(dst string, opt *ClientOption) conn {
			if dst == "127.0.0.1:0" || dst == "127.0.2.1:0" { // primary node
				return primaryNodeConn
			} else { // replica node
				return replicaNodeConn
			}
		},
		newRetryer(defaultRetryDelayFn),
	)
	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}

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
						newResult(strmsg('+', "OK"), nil),
						newResult(strmsg('+', "OK"), nil),
						newResult(slicemsg('*', []RedisMessage{strmsg('+', "a")}), nil),
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
		primaryNodeConn.AcquireFn = func() wire { return w }
		replicaNodeConn.AcquireFn = func() wire { return w } // Subscribe can work on replicas
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
				return newResult(strmsg('+', "Delegate"), nil)
			},
			DoMultiFn: func(cmd ...Completed) *redisresults {
				if len(cmd) == 4 {
					return &redisresults{s: []RedisResult{
						newResult(strmsg('+', "OK"), nil),
						newResult(strmsg('+', "OK"), nil),
						newResult(strmsg('+', "OK"), nil),
						newResult(slicemsg('*', []RedisMessage{
							strmsg('+', "Delegate0"),
							strmsg('+', "Delegate1"),
						}), nil),
					}}
				}
				return &redisresults{s: []RedisResult{
					newResult(strmsg('+', "Delegate0"), nil),
					newResult(strmsg('+', "Delegate1"), nil),
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
			)[3].val.values() {
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
			t.Fatalf("Dedicated doesn't put back the wire")
		}
		if !closed {
			t.Fatalf("Dedicated doesn't delegate Close")
		}
	})

	t.Run("Dedicate", func(t *testing.T) {
		closed := false
		w := &mockWire{
			DoFn: func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "Delegate"), nil)
			},
			DoMultiFn: func(cmd ...Completed) *redisresults {
				if len(cmd) == 4 {
					return &redisresults{s: []RedisResult{
						newResult(strmsg('+', "OK"), nil),
						newResult(strmsg('+', "OK"), nil),
						newResult(strmsg('+', "OK"), nil),
						newResult(slicemsg('*', []RedisMessage{
							strmsg('+', "Delegate0"),
							strmsg('+', "Delegate1"),
						}), nil),
					}}
				}
				return &redisresults{s: []RedisResult{
					newResult(strmsg('+', "Delegate0"), nil),
					newResult(strmsg('+', "Delegate1"), nil),
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
		)[3].val.values() {
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
			t.Fatalf("Dedicated doesn't put back the wire")
		}
		if !closed {
			t.Fatalf("Dedicated doesn't delegate Close")
		}
	})
}

func TestClusterClient_ReadNodeSelector_SendToAlternatePrimaryAndReplicaNodes(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())

	primaryNodeConn := &mockConn{
		DoMultiFn: func(multi ...Completed) *redisresults {
			resps := make([]RedisResult, len(multi))
			for i, cmd := range multi {
				resps[i] = newResult(strmsg('+', strings.Join(cmd.Commands(), " ")), nil)
			}
			return &redisresults{s: resps}
		},
		DoMultiCacheFn: func(multi ...CacheableTTL) *redisresults {
			resps := make([]RedisResult, len(multi))
			for i, cmd := range multi {
				resps[i] = newResult(strmsg('+', strings.Join(cmd.Cmd.Commands(), " ")), nil)
			}
			return &redisresults{s: resps}
		},
		DoOverride: map[string]func(cmd Completed) RedisResult{
			"CLUSTER SLOTS": func(cmd Completed) RedisResult {
				return slotsMultiResp
			},
			"INFO": func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "INFO"), nil)
			},
			"GET K1{a}": func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "GET K1{a}"), nil)
			},
			"GET K2{b}": func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "GET K2{b}"), nil)
			},
		},
	}
	replicaNodeConn := &mockConn{
		DoMultiFn: func(multi ...Completed) *redisresults {
			resps := make([]RedisResult, len(multi))
			for i, cmd := range multi {
				resps[i] = newResult(strmsg('+', strings.Join(cmd.Commands(), " ")), nil)
			}
			return &redisresults{s: resps}
		},
		DoMultiCacheFn: func(multi ...CacheableTTL) *redisresults {
			resps := make([]RedisResult, len(multi))
			for i, cmd := range multi {
				resps[i] = newResult(strmsg('+', strings.Join(cmd.Cmd.Commands(), " ")), nil)
			}
			return &redisresults{s: resps}
		},
		DoOverride: map[string]func(cmd Completed) RedisResult{
			"GET Do": func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "GET Do"), nil)
			},
			"GET K1{a}": func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "GET K1{a}"), nil)
			},
			"GET K2{b}": func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "GET K2{b}"), nil)
			},
		},
		DoCacheOverride: map[string]func(cmd Cacheable, ttl time.Duration) RedisResult{
			"GET DoCache": func(cmd Cacheable, ttl time.Duration) RedisResult {
				return newResult(strmsg('+', "GET DoCache"), nil)
			},
		},
	}

	nextNode := -1
	client, err := newClusterClient(
		&ClientOption{
			InitAddress: []string{"127.0.0.1:0"},
			ReadNodeSelector: func(_ uint16, _ []NodeInfo) int {
				nextNode++
				return (nextNode / 2) % 2
			},
		},
		func(dst string, opt *ClientOption) conn {
			if dst == "127.0.0.1:0" || dst == "127.0.2.1:0" { // primary nodes
				return primaryNodeConn
			} else { // replica nodes
				return replicaNodeConn
			}
		},
		newRetryer(defaultRetryDelayFn),
	)
	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}

	t.Run("DoMulti Multi Slot", func(t *testing.T) {
		c1 := client.B().Get().Key("K1{a}").Build()
		c2 := client.B().Get().Key("K2{b}").Build()
		resps := client.DoMulti(context.Background(), c1, c2)
		if v, err := resps[0].ToString(); err != nil || v != "GET K1{a}" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
		if v, err := resps[1].ToString(); err != nil || v != "GET K2{b}" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("DoMulti Multi Slot Large", func(t *testing.T) {
		var cmds []Completed
		for i := 0; i < 500; i++ {
			cmds = append(cmds, client.B().Get().Key("K1{a}").Build())
		}
		resps := client.DoMulti(context.Background(), cmds...)
		for _, resp := range resps {
			if v, err := resp.ToString(); err != nil || v != "GET K1{a}" {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		}
	})

	t.Run("DoMultiCache Multi Slot", func(t *testing.T) {
		c1 := client.B().Get().Key("K1{a}").Cache()
		c2 := client.B().Get().Key("K2{b}").Cache()
		resps := client.DoMultiCache(context.Background(), CT(c1, time.Second), CT(c2, time.Second))
		if v, err := resps[0].ToString(); err != nil || v != "GET K1{a}" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
		if v, err := resps[1].ToString(); err != nil || v != "GET K2{b}" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("DoMultiCache Multi Slot Large", func(t *testing.T) {
		var cmds []CacheableTTL
		for i := 0; i < 500; i++ {
			cmds = append(cmds, CT(client.B().Get().Key("K1{a}").Cache(), time.Second))
		}
		resps := client.DoMultiCache(context.Background(), cmds...)
		for _, resp := range resps {
			if v, err := resp.ToString(); err != nil || v != "GET K1{a}" {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		}
	})
}
