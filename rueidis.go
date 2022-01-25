// Package rueidis is a fast Golang Redis RESP3 client that does auto pipelining and supports client side caching.
package rueidis

import (
	"context"
	"crypto/tls"
	"errors"
	"net"
	"time"

	"github.com/rueian/rueidis/internal/cmds"
)

const (
	// DefaultCacheBytes is the default value of ClientOption.CacheSizeEachConn, which is 128 MiB
	DefaultCacheBytes = 128 * (1 << 20)
	// DefaultPoolSize is the default value of ClientOption.BlockingPoolSize
	DefaultPoolSize = 1000
	// DefaultDialTimeout is the default value of ClientOption.Dialer.Timeout
	DefaultDialTimeout = 5 * time.Second
	// DefaultTCPKeepAlive is the default value of ClientOption.Dialer.KeepAlive
	DefaultTCPKeepAlive = 1 * time.Second
)

var (
	// ErrClosing means the Client.Close had been called
	ErrClosing = errors.New("rueidis client is closing")
	// ErrNoAddr means the ClientOption.InitAddress is empty
	ErrNoAddr = errors.New("no address in InitAddress")
)

// ClientOption should be passed to NewClient to construct a Client
type ClientOption struct {
	// InitAddress point to redis nodes.
	// Rueidis will connect to them one by one and issue CLUSTER SLOT command to initialize the cluster client until success.
	// If len(InitAddress) == 1 and the address is not running in cluster mode, rueidis will fall back to the single client mode.
	InitAddress []string
	// ShuffleInit is a handy flag that shuffles the InitAddress after passing to the NewClient() if it is true
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

// Client is the redis client interface for both single redis instance and redis cluster. It should be created from the NewClient()
type Client interface {
	// B is the getter function to the command builder for the client
	// If the client is a cluster client, the command builder also prohibits cross key slots in one command.
	B() *cmds.Builder
	// Do is the method sending user's redis command building from the B() to a redis node.
	//  client.Do(ctx, client.B().Get().Key("k").Build()).ToString()
	// All concurrent non-blocking commands will be pipelined automatically and have better throughput.
	// Blocking commands will use another separated connection pool.
	// The cmd parameter is recycled after passing into Do() and should not be reused.
	Do(ctx context.Context, cmd cmds.Completed) (resp RedisResult)
	// DoCache is similar to Do, but it uses opt-in client side caching and requires a client side TTL.
	// The explicit client side TTL specifies the maximum TTL on the client side.
	// If the key's TTL on the server is smaller than the client side TTL, the client side TTL will be capped.
	//  client.Do(ctx, client.B().Get().Key("k").Cache(), time.Minute).ToString()
	// The above example will send the following command to redis if cache miss:
	//  CLIENT CACHING YES
	//  GET k
	//  PTTL k
	// The in-memory cache size is configured by ClientOption.CacheSizeEachConn.
	// The cmd parameter is recycled after passing into DoCache() and should not be reused.
	DoCache(ctx context.Context, cmd cmds.Cacheable, ttl time.Duration) (resp RedisResult)
	// Dedicated acquire a connection from the blocking connection pool, no one else can use the connection
	// during Dedicated. The main usage of Dedicated is CAS operation, which is WATCH + MULTI + EXEC.
	// However, one should try to avoid CAS operation but use Lua script instead, because occupying a connection
	// is not good for performance.
	Dedicated(fn func(DedicatedClient) error) (err error)
	// Close will make further calls to the client be rejected with ErrClosing,
	// and Close will wait until all pending calls finished.
	Close()
}

// DedicatedClient is obtained from Client.Dedicated() and it will be bound to single redis connection and
// no other commands can be pipelined in to this connection during Client.Dedicated().
// If the DedicatedClient is obtained from cluster client, the first command to it must have a Key() to identify the redis node.
type DedicatedClient interface {
	// B is inherited from the Client
	B() *cmds.Builder
	// Do is the same as Client
	// The cmd parameter is recycled after passing into Do() and should not be reused.
	Do(ctx context.Context, cmd cmds.Completed) (resp RedisResult)
	// DoMulti takes multiple redis commands and sends them together, reducing RTT from the user code.
	// The multi parameters are recycled after passing into DoMulti() and should not be reused.
	DoMulti(ctx context.Context, multi ...cmds.Completed) (resp []RedisResult)
}

// NewClient uses ClientOption to initialize the Client for both cluster client and single client.
// It will first try to connect as cluster client. If the len(ClientOption.InitAddress) == 1 and
// the address does not enable cluster mode, the NewClient() will use single client instead.
func NewClient(option ClientOption) (client Client, err error) {
	if client, err = newClusterClient(option, makeClusterConn); err != nil {
		if len(option.InitAddress) == 1 && err.Error() == redisErrMsgClusterDisabled {
			client, err = newSingleClient(option, client.(*clusterClient).single(), makeSingleConn)
		} else if client != nil {
			client.Close()
			return nil, err
		}
	}
	return client, err
}

func makeClusterConn(dst string, opt ClientOption) conn {
	return makeMux(dst, opt, dial, false)
}

func makeSingleConn(dst string, opt ClientOption) conn {
	return makeMux(dst, opt, dial, true)
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
