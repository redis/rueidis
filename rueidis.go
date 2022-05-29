// Package rueidis is a fast Golang Redis RESP3 client that does auto pipelining and supports client side caching.
package rueidis

import (
	"context"
	"crypto/tls"
	"errors"
	"math/rand"
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
	ErrClosing = errors.New("rueidis client is closing or unable to connect redis")
	// ErrNoAddr means the ClientOption.InitAddress is empty
	ErrNoAddr = errors.New("no alive address in InitAddress")
)

// ClientOption should be passed to NewClient to construct a Client
type ClientOption struct {
	// TCP & TLS
	// Dialer can be used to customized how rueidis connect to a redis instance via TCP, including:
	// - Timeout, the default is DefaultDialTimeout
	// - KeepAlive, the default is DefaultTCPKeepAlive
	// The Dialer.KeepAlive interval is used to detect an unresponsive idle tcp connection.
	// OS takes at least (tcp_keepalive_probes+1)*Dialer.KeepAlive time to conclude an idle connection to be unresponsive.
	// For example: DefaultTCPKeepAlive = 1s and the default of tcp_keepalive_probes on Linux is 9.
	// Therefore, it takes at least 10s to kill an idle and unresponsive tcp connection on Linux by default.
	Dialer    net.Dialer
	TLSConfig *tls.Config

	// Sentinel options, including MasterSet and Auth options
	Sentinel SentinelOption

	// Redis AUTH parameters
	Username   string
	Password   string
	ClientName string

	// InitAddress point to redis nodes.
	// Rueidis will connect to them one by one and issue CLUSTER SLOT command to initialize the cluster client until success.
	// If len(InitAddress) == 1 and the address is not running in cluster mode, rueidis will fall back to the single client mode.
	// If ClientOption.Sentinel.MasterSet is set, then InitAddress will be used to connect sentinels
	InitAddress []string

	SelectDB int

	// CacheSizeEachConn is redis client side cache size that bind to each TCP connection to a single redis instance.
	// The default is DefaultCacheBytes.
	CacheSizeEachConn int

	// BlockingPoolSize is the size of the connection pool shared by blocking commands (ex BLPOP, XREAD with BLOCK).
	// The default is DefaultPoolSize.
	BlockingPoolSize int

	// ConnWriteTimeout is applied net.Conn.SetWriteDeadline and periodic PING to redis
	// Since the Dialer.KeepAlive will not be triggered if there is data in the outgoing buffer,
	// ConnWriteTimeout should be set in order to detect local congestion or unresponsive redis server.
	// This default is ClientOption.Dialer.KeepAlive * (9+1), where 9 is the default of tcp_keepalive_probes on Linux.
	ConnWriteTimeout time.Duration

	// ShuffleInit is a handy flag that shuffles the InitAddress after passing to the NewClient() if it is true
	ShuffleInit bool
}

// SentinelOption contains MasterSet,
type SentinelOption struct {
	// TCP & TLS, same as ClientOption but for connecting sentinel
	Dialer    net.Dialer
	TLSConfig *tls.Config

	// MasterSet is the redis master set name monitored by sentinel cluster.
	// If this field is set, then ClientOption.InitAddress will be used to connect to sentinel cluster.
	MasterSet string

	// Redis AUTH parameters for sentinel
	Username   string
	Password   string
	ClientName string
}

// Client is the redis client interface for both single redis instance and redis cluster. It should be created from the NewClient()
type Client interface {
	// B is the getter function to the command builder for the client
	// If the client is a cluster client, the command builder also prohibits cross key slots in one command.
	B() cmds.Builder
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

	// Receive accepts SUBSCRIBE, SSUBSCRIBE, PSUBSCRIBE command and a message handler.
	// Receive will block and then return value only when the following cases:
	//   1. nil, when received any unsubscribe/punsubscribe message related to the provided `subscribe` command.
	//   2. ErrClosing, when the client is closed manually.
	//   3. ctx.Err(), when the deadline of `ctx` is exceeded.
	Receive(ctx context.Context, subscribe cmds.Completed, fn func(msg PubSubMessage)) error

	// Dedicated acquire a connection from the blocking connection pool, no one else can use the connection
	// during Dedicated. The main usage of Dedicated is CAS operation, which is WATCH + MULTI + EXEC.
	// However, one should try to avoid CAS operation but use Lua script instead, because occupying a connection
	// is not good for performance.
	Dedicated(fn func(DedicatedClient) error) (err error)

	// Dedicate does the same as Dedicated, but it exposes DedicatedClient directly
	// and requires user to invoke cancel() manually to put connection back to the pool.
	Dedicate() (client DedicatedClient, cancel func())

	// Close will make further calls to the client be rejected with ErrClosing,
	// and Close will wait until all pending calls finished.
	Close()
}

// DedicatedClient is obtained from Client.Dedicated() and it will be bound to single redis connection and
// no other commands can be pipelined in to this connection during Client.Dedicated().
// If the DedicatedClient is obtained from cluster client, the first command to it must have a Key() to identify the redis node.
type DedicatedClient interface {
	// B is inherited from the Client
	B() cmds.Builder
	// Do is the same as Client's
	// The cmd parameter is recycled after passing into Do() and should not be reused.
	Do(ctx context.Context, cmd cmds.Completed) (resp RedisResult)
	// DoMulti takes multiple redis commands and sends them together, reducing RTT from the user code.
	// The multi parameters are recycled after passing into DoMulti() and should not be reused.
	DoMulti(ctx context.Context, multi ...cmds.Completed) (resp []RedisResult)
	// Receive is the same as Client's
	Receive(ctx context.Context, subscribe cmds.Completed, fn func(msg PubSubMessage)) error
	// SetPubSubHooks is an alternative way to processing Pub/Sub messages instead of using Receive.
	// SetPubSubHooks is non-blocking and allows users to subscribe/unsubscribe channels later.
	// Note that the hooks will be called sequentially but in another goroutine.
	// The return value will be either:
	//   1. an error channel, if the hooks passed in is not zero, or
	//   2. nil, if the hooks passed in is zero. (used for reset hooks)
	// In the former case, the error channel is guaranteed to be close when the hooks will not be called anymore,
	// and has at most one error describing the reason why the hooks will not be called anymore.
	// Users can use the error channel to detect disconnection.
	SetPubSubHooks(hooks PubSubHooks) <-chan error
}

// NewClient uses ClientOption to initialize the Client for both cluster client and single client.
// It will first try to connect as cluster client. If the len(ClientOption.InitAddress) == 1 and
// the address does not enable cluster mode, the NewClient() will use single client instead.
func NewClient(option ClientOption) (client Client, err error) {
	if option.CacheSizeEachConn <= 0 {
		option.CacheSizeEachConn = DefaultCacheBytes
	}
	if option.Dialer.Timeout == 0 {
		option.Dialer.Timeout = DefaultDialTimeout
	}
	if option.Dialer.KeepAlive == 0 {
		option.Dialer.KeepAlive = DefaultTCPKeepAlive
	}
	if option.ConnWriteTimeout == 0 {
		option.ConnWriteTimeout = option.Dialer.KeepAlive * 10
	}
	if option.ShuffleInit {
		rand.Shuffle(len(option.InitAddress), func(i, j int) {
			option.InitAddress[i], option.InitAddress[j] = option.InitAddress[j], option.InitAddress[i]
		})
	}
	if option.Sentinel.MasterSet != "" {
		return newSentinelClient(&option, makeConn)
	}
	if client, err = newClusterClient(&option, makeConn); err != nil {
		if len(option.InitAddress) == 1 && err.Error() == redisErrMsgClusterDisabled {
			client, err = newSingleClient(&option, client.(*clusterClient).single(), makeConn)
		} else if client != (*clusterClient)(nil) {
			client.Close()
			return nil, err
		}
	}
	return client, err
}

func makeConn(dst string, opt *ClientOption) conn {
	return makeMux(dst, opt, dial)
}

func dial(dst string, opt *ClientOption) (conn net.Conn, err error) {
	if opt.TLSConfig != nil {
		conn, err = tls.DialWithDialer(&opt.Dialer, "tcp", dst, opt.TLSConfig)
	} else {
		conn, err = opt.Dialer.Dial("tcp", dst)
	}
	return conn, err
}

const redisErrMsgClusterDisabled = "ERR This instance has cluster support disabled"
