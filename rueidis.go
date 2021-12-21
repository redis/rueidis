package rueidis

import (
	"crypto/tls"
	"net"
	"time"

	"github.com/rueian/rueidis/internal/proto"
)

var (
	IsRedisNil = proto.IsRedisNil
)

func NewClusterClient(option ClusterClientOption) (*ClusterClient, error) {
	return newClusterClient(option, makeConn)
}

func NewSingleClient(option SingleClientOption) (*SingleClient, error) {
	return newSingleClient(option, makeConn)
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
