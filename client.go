package rueidis

import (
	"context"
	"time"

	"github.com/rueian/rueidis/internal/cmds"
	"github.com/rueian/rueidis/internal/proto"
)

type SingleClientOption struct {
	Address    string
	ConnOption ConnOption
}

type SingleClient struct {
	cmd  *cmds.Builder
	conn conn
}

func newSingleClient(opt SingleClientOption, connFn connFn) (*SingleClient, error) {
	client := &SingleClient{cmd: cmds.NewBuilder(cmds.NoSlot), conn: connFn(opt.Address, opt.ConnOption)}

	if err := client.conn.Dial(); err != nil {
		return nil, err
	}

	opt.ConnOption.PubSubHandlers.installHook(client.cmd, func() conn { return client.conn })

	return client, nil
}

func (c *SingleClient) B() *cmds.Builder {
	return c.cmd
}

func (c *SingleClient) Do(ctx context.Context, cmd cmds.Completed) (resp proto.Result) {
	resp = c.conn.Do(cmd)
	c.cmd.Put(cmd.CommandSlice())
	return resp
}

func (c *SingleClient) DoCache(ctx context.Context, cmd cmds.Cacheable, ttl time.Duration) (resp proto.Result) {
	resp = c.conn.DoCache(cmd, ttl)
	c.cmd.Put(cmd.CommandSlice())
	return resp
}

func (c *SingleClient) Dedicated(fn func(DedicatedClient) error) (err error) {
	wire := c.conn.Acquire()
	err = fn(&dedicatedSingleClient{cmd: c.cmd, wire: wire})
	c.conn.Store(wire)
	return err
}

func (c *SingleClient) Close() {
	c.conn.Close()
}

type dedicatedSingleClient struct {
	cmd  *cmds.Builder
	wire wire
}

func (c *dedicatedSingleClient) B() *cmds.Builder {
	return c.cmd
}

func (c *dedicatedSingleClient) Do(ctx context.Context, cmd cmds.Completed) (resp proto.Result) {
	resp = c.wire.Do(cmd)
	c.cmd.Put(cmd.CommandSlice())
	return resp
}

func (c *dedicatedSingleClient) DoMulti(ctx context.Context, multi ...cmds.Completed) (resp []proto.Result) {
	if len(multi) == 0 {
		return nil
	}
	resp = c.wire.DoMulti(multi...)
	for _, cmd := range multi {
		c.cmd.Put(cmd.CommandSlice())
	}
	return resp
}
