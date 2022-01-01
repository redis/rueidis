package rueidis

import (
	"context"
	"crypto/tls"
	"errors"
	"net"
	"time"

	"github.com/rueian/rueidis/internal/cmds"
	"github.com/rueian/rueidis/internal/proto"
)

const (
	DefaultCacheBytes   = 128 * (1 << 20) // 128 MiB
	DefaultPoolSize     = 1000
	DefaultDialTimeout  = 5 * time.Second
	DefaultTCPKeepAlive = 1 * time.Second
)

var ErrConnClosing = errors.New("connection is closing")

type ClientOption struct {
	// InitAddress point to redis nodes.
	// Rueidis will connect to them one by one and issue CLUSTER SLOT command to initialize the cluster client until success.
	// If len(InitAddress) == 1 and the address is not running in cluster mode, rueidis will fall back to the single client mode.
	InitAddress []string
	// ShuffleInit is a handy flag that shuffles the InitAddress after passing to NewClient
	ShuffleInit bool

	// CacheSizeEachConn is redis client side cache size that bind to each TCP connection to a single redis instance.
	// The default is DefaultCacheBytes.
	CacheSizeEachConn int

	// BlockingPoolSize is the size of the connection pool shared by blocking commands (ex BLPOP, XREAD with BLOCK).
	// The default is DefaultPoolSize.
	BlockingPoolSize int

	// Redis AUTH parameters
	Username   string
	Password   string
	ClientName string
	SelectDB   int

	// TCP & TLS
	// Dialer can be used to customized how rueidis connect to a redis instance via TCP, including:
	// - Timeout, the default is DefaultDialTimeout
	// - KeepAlive, the default is DefaultTCPKeepAlive
	Dialer    net.Dialer
	TLSConfig *tls.Config

	// Redis PubSub callbacks, should be created from NewPubSubOption
	PubSubOption PubSubOption
}

type Client interface {
	B() *cmds.Builder
	Do(ctx context.Context, cmd cmds.Completed) (resp proto.Result)
	DoCache(ctx context.Context, cmd cmds.Cacheable, ttl time.Duration) (resp proto.Result)
	Dedicated(fn func(DedicatedClient) error) (err error)
	Close()
}

func NewClient(option ClientOption) (client Client, err error) {
	client, err = newClusterClient(option, makeConn)
	if err != nil && len(option.InitAddress) == 1 && err.Error() == redisErrMsgClusterDisabled {
		client, err = newSingleClient(option, makeConn)
	}
	return client, err
}

type DedicatedClient interface {
	B() *cmds.Builder
	Do(ctx context.Context, cmd cmds.Completed) (resp proto.Result)
	DoMulti(ctx context.Context, multi ...cmds.Completed) (resp []proto.Result)
}

func IsRedisNil(err error) bool {
	return proto.IsRedisNil(err)
}

func makeConn(dst string, opt ClientOption) conn {
	return makeMux(dst, opt, dial)
}

func dial(dst string, opt ClientOption) (conn net.Conn, err error) {
	if opt.Dialer.Timeout == 0 {
		opt.Dialer.Timeout = DefaultDialTimeout
	}
	if opt.Dialer.KeepAlive == 0 {
		opt.Dialer.KeepAlive = DefaultTCPKeepAlive
	}
	if opt.TLSConfig != nil {
		conn, err = tls.DialWithDialer(&opt.Dialer, "tcp", dst, opt.TLSConfig)
	} else {
		conn, err = opt.Dialer.Dial("tcp", dst)
	}
	return conn, err
}

const redisErrMsgClusterDisabled = "ERR This instance has cluster support disabled"
