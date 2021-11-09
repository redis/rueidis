package client

import (
	"time"

	"github.com/rueian/rueidis/internal/cmds"
	"github.com/rueian/rueidis/internal/proto"
	"github.com/rueian/rueidis/pkg/conn"
)

type SingleClient struct {
	Cmd  *cmds.Builder
	conn *conn.Conn
}

func NewSingleClient(dst string, option conn.Option) (*SingleClient, error) {
	c := conn.NewConn(dst, option)
	if err := c.Init(); err != nil {
		return nil, err
	}
	return &SingleClient{
		Cmd:  cmds.NewBuilder(),
		conn: c,
	}, nil
}

func (c *SingleClient) Info() proto.Message {
	return c.conn.Info()
}

func (c *SingleClient) Do(cmd cmds.Completed) (resp proto.Result) {
	resp = c.conn.Do(cmd)
	c.Cmd.Put(cmd.Commands())
	return resp
}

func (c *SingleClient) DoMulti(multi ...cmds.Completed) (resp []proto.Result) {
	resp = c.conn.DoMulti(multi...)
	for _, cmd := range multi {
		c.Cmd.Put(cmd.Commands())
	}
	return resp
}

func (c *SingleClient) DoCache(cmd cmds.Cacheable, ttl time.Duration) (resp proto.Result) {
	resp = c.conn.DoCache(cmd, ttl)
	c.Cmd.Put(cmd.Commands())
	return resp
}

func (c *SingleClient) Close() {
	c.conn.Close()
}
