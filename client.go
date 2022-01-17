package rueidis

import (
	"context"
	"time"

	"github.com/rueian/rueidis/internal/cmds"
)

type singleClient struct {
	cmd  *cmds.Builder
	conn conn
}

func newSingleClient(opt ClientOption, connFn connFn) (*singleClient, error) {
	if len(opt.InitAddress) == 0 {
		return nil, ErrNoAddr
	}

	client := &singleClient{cmd: cmds.NewBuilder(cmds.NoSlot), conn: connFn(opt.InitAddress[0], opt)}

	if err := client.conn.Dial(); err != nil {
		return nil, err
	}

	if opt.PubSubOption.onConnected != nil {
		var install func(error)
		install = func(prev error) {
			if prev != ErrClosing {
				dcc := &dedicatedSingleClient{cmd: client.cmd, wire: client.conn}
				client.conn.OnDisconnected(install)
				opt.PubSubOption.onConnected(prev, dcc)
			}
		}
		install(nil)
	}

	return client, nil
}

func (c *singleClient) B() *cmds.Builder {
	return c.cmd
}

func (c *singleClient) Do(ctx context.Context, cmd cmds.Completed) (resp RedisResult) {
	resp = c.conn.Do(cmd)
	c.cmd.Put(cmd.CommandSlice())
	return resp
}

func (c *singleClient) DoCache(ctx context.Context, cmd cmds.Cacheable, ttl time.Duration) (resp RedisResult) {
	resp = c.conn.DoCache(cmd, ttl)
	c.cmd.Put(cmd.CommandSlice())
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
	cmd  *cmds.Builder
	wire wire
}

func (c *dedicatedSingleClient) B() *cmds.Builder {
	return c.cmd
}

func (c *dedicatedSingleClient) Do(ctx context.Context, cmd cmds.Completed) (resp RedisResult) {
	resp = c.wire.Do(cmd)
	c.cmd.Put(cmd.CommandSlice())
	return resp
}

func (c *dedicatedSingleClient) DoMulti(ctx context.Context, multi ...cmds.Completed) (resp []RedisResult) {
	if len(multi) == 0 {
		return nil
	}
	resp = c.wire.DoMulti(multi...)
	for _, cmd := range multi {
		c.cmd.Put(cmd.CommandSlice())
	}
	return resp
}
