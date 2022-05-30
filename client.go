package rueidis

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/rueian/rueidis/internal/cmds"
)

type singleClient struct {
	conn conn
	stop uint32
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
retry:
	resp = c.conn.Do(ctx, cmd)
	if cmd.IsReadOnly() && c.isRetryable(resp.NonRedisError(), ctx) {
		goto retry
	}
	cmds.Put(cmd.CommandSlice())
	return resp
}

func (c *singleClient) DoCache(ctx context.Context, cmd cmds.Cacheable, ttl time.Duration) (resp RedisResult) {
retry:
	resp = c.conn.DoCache(ctx, cmd, ttl)
	if c.isRetryable(resp.NonRedisError(), ctx) {
		goto retry
	}
	cmds.Put(cmd.CommandSlice())
	return resp
}

func (c *singleClient) Receive(ctx context.Context, subscribe cmds.Completed, fn func(msg PubSubMessage)) (err error) {
retry:
	err = c.conn.Receive(ctx, subscribe, fn)
	if c.isRetryable(err, ctx) {
		goto retry
	}
	cmds.Put(subscribe.CommandSlice())
	return err
}

func (c *singleClient) Dedicated(fn func(DedicatedClient) error) (err error) {
	wire := c.conn.Acquire()
	err = fn(&dedicatedSingleClient{cmd: c.cmd, wire: wire})
	c.conn.Store(wire)
	return err
}

func (c *singleClient) Dedicate() (DedicatedClient, func()) {
	wire := c.conn.Acquire()
	return &dedicatedSingleClient{cmd: c.cmd, wire: wire}, func() { c.conn.Store(wire) }
}

func (c *singleClient) Close() {
	atomic.StoreUint32(&c.stop, 1)
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
retry:
	resp = c.wire.Do(ctx, cmd)
	if cmd.IsReadOnly() && isRetryable(resp.NonRedisError(), c.wire, ctx) {
		goto retry
	}
	cmds.Put(cmd.CommandSlice())
	return resp
}

func (c *dedicatedSingleClient) DoMulti(ctx context.Context, multi ...cmds.Completed) (resp []RedisResult) {
	if len(multi) == 0 {
		return nil
	}
	readonly := allReadOnly(multi)
retry:
	resp = c.wire.DoMulti(ctx, multi...)
	if readonly && anyRetryable(resp, c.wire, ctx) {
		goto retry
	}
	for _, cmd := range multi {
		cmds.Put(cmd.CommandSlice())
	}
	return resp
}

func (c *dedicatedSingleClient) Receive(ctx context.Context, subscribe cmds.Completed, fn func(msg PubSubMessage)) (err error) {
retry:
	err = c.wire.Receive(ctx, subscribe, fn)
	if isRetryable(err, c.wire, ctx) {
		goto retry
	}
	cmds.Put(subscribe.CommandSlice())
	return err
}

func (c *dedicatedSingleClient) SetPubSubHooks(hooks PubSubHooks) <-chan error {
	return c.wire.SetPubSubHooks(hooks)
}

func (c *dedicatedSingleClient) Close() {
	c.wire.Close()
}

func (c *singleClient) isRetryable(err error, ctx context.Context) bool {
	return err != nil && atomic.LoadUint32(&c.stop) == 0 && ctx.Err() == nil
}

func isRetryable(err error, w wire, ctx context.Context) bool {
	return err != nil && w.Error() == nil && ctx.Err() == nil
}

func anyRetryable(resp []RedisResult, w wire, ctx context.Context) bool {
	for _, r := range resp {
		if isRetryable(r.NonRedisError(), w, ctx) {
			return true
		}
	}
	return false
}

func allReadOnly(multi []cmds.Completed) bool {
	for _, cmd := range multi {
		if cmd.IsWrite() {
			return false
		}
	}
	return true
}

func allSameSlot(multi []cmds.Completed) bool {
	for i := 1; i < len(multi); i++ {
		if multi[0].Slot() != multi[i].Slot() {
			return false
		}
	}
	return true
}
