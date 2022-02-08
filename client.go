package rueidis

import (
	"context"
	"time"

	"github.com/rueian/rueidis/internal/cmds"
)

type singleClient struct {
	conn conn
	cmd  cmds.Builder
}

func newSingleClient(opt *ClientOption, prev conn, connFn connFn) (*singleClient, error) {
	if len(opt.InitAddress) == 0 {
		return nil, ErrNoAddr
	}

	conn := connFn(opt.InitAddress[0], opt)
	conn.Override(prev)
	if err := conn.Dial(); err != nil {
		return nil, err
	}

	return &singleClient{cmd: cmds.NewBuilder(cmds.NoSlot), conn: conn}, nil
}

func (c *singleClient) B() cmds.Builder {
	return c.cmd
}

func (c *singleClient) Do(ctx context.Context, cmd cmds.Completed) (resp RedisResult) {
	resp = c.conn.Do(ctx, cmd)
	cmds.Put(cmd.CommandSlice())
	return resp
}

func (c *singleClient) DoCache(ctx context.Context, cmd cmds.Cacheable, ttl time.Duration) (resp RedisResult) {
	resp = c.conn.DoCache(ctx, cmd, ttl)
	cmds.Put(cmd.CommandSlice())
	return resp
}

func (c *singleClient) Dedicated(fn func(DedicatedClient) error) (err error) {
	wire := c.conn.Acquire()
	err = fn(&dedicatedSingleClient{cmd: c.cmd, wire: wire})
	c.conn.Store(wire)
	return err
}

func (c *singleClient) Close() {
	c.conn.Close()
}

type dedicatedSingleClient struct {
	wire wire
	cmd  cmds.Builder
}

func (c *dedicatedSingleClient) B() cmds.Builder {
	return c.cmd
}

func (c *dedicatedSingleClient) Do(ctx context.Context, cmd cmds.Completed) (resp RedisResult) {
	resp = c.wire.Do(ctx, cmd)
	cmds.Put(cmd.CommandSlice())
	return resp
}

func (c *dedicatedSingleClient) DoMulti(ctx context.Context, multi ...cmds.Completed) (resp []RedisResult) {
	if len(multi) == 0 {
		return nil
	}
	resp = c.wire.DoMulti(ctx, multi...)
	for _, cmd := range multi {
		cmds.Put(cmd.CommandSlice())
	}
	return resp
}
