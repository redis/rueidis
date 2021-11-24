package rueidis

import (
	"github.com/rueian/rueidis/pkg/client"
	"github.com/rueian/rueidis/pkg/conn"
)

var (
	ErrNoSlot      = client.ErrNoSlot
	ErrNoNodes     = client.ErrNoNodes
	ErrConnClosing = conn.ErrConnClosing
)

type SingleClientOption client.SingleClientOption

type ClusterClientOption client.ClusterClientOption

func NewClusterClient(option ClusterClientOption) (*client.ClusterClient, error) {
	return client.NewClusterClient(client.ClusterClientOption(option))
}

func NewSingleClient(option SingleClientOption) (*client.SingleClient, error) {
	return client.NewSingleClient(client.SingleClientOption(option))
}
