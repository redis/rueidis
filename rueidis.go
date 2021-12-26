package rueidis

import (
	"crypto/tls"
	"errors"
	"net"
	"time"

	"github.com/rueian/rueidis/internal/proto"
)

const (
	DefaultCacheBytes = 128 * (1 << 20) // 128 MiB
	DefaultPoolSize   = 1000
)

var ErrConnClosing = errors.New("connection is closing")

type ConnOption struct {
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
	DialTimeout time.Duration
	TLSConfig   *tls.Config

	// Redis PubSub callbacks
	PubSubHandlers PubSubHandlers
}

func NewClusterClient(option ClusterClientOption) (*ClusterClient, error) {
	return newClusterClient(option, makeConn)
}

func NewSingleClient(option SingleClientOption) (*SingleClient, error) {
	return newSingleClient(option, makeConn)
}

func IsRedisNil(err error) bool {
	return proto.IsRedisNil(err)
}

func makeConn(dst string, opt ConnOption) conn {
	return makeMux(dst, opt, dial)
}

func dial(dst string, opt ConnOption) (conn net.Conn, err error) {
	dialer := &net.Dialer{Timeout: opt.DialTimeout, KeepAlive: time.Second}
	if opt.TLSConfig != nil {
		conn, err = tls.DialWithDialer(dialer, "tcp", dst, opt.TLSConfig)
	} else {
		conn, err = dialer.Dial("tcp", dst)
	}
	return conn, err
}
