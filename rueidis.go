package rueidis

import (
	"crypto/tls"
	"net"
	"time"

	"github.com/rueian/rueidis/internal/proto"
	"github.com/rueian/rueidis/pkg/client"
	"github.com/rueian/rueidis/pkg/conn"
)

var (
	ErrNoSlot      = client.ErrNoSlot
	ErrNoNodes     = client.ErrNoNodes
	ErrConnClosing = conn.ErrConnClosing

	IsRedisNil = proto.IsRedisNil
)

type SingleClientOption client.SingleClientOption

type ClusterClientOption client.ClusterClientOption

func NewClusterClient(option ClusterClientOption) (*client.ClusterClient, error) {
	return client.NewClusterClient(client.ClusterClientOption(option), dial)
}

func NewSingleClient(option SingleClientOption) (*client.SingleClient, error) {
	return client.NewSingleClient(client.SingleClientOption(option), dial)
}

func dial(dst string, opt conn.Option) (conn net.Conn, err error) {
	dialer := &net.Dialer{Timeout: opt.DialTimeout, KeepAlive: time.Second}
	if opt.TLSConfig != nil {
		conn, err = tls.DialWithDialer(dialer, "tcp", dst, opt.TLSConfig)
	} else {
		conn, err = dialer.Dial("tcp", dst)
	}
	return conn, err
}
