package rueidis

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	"github.com/rueian/rueidis/internal/cmds"
)

// ErrNoSlot indicates that there is no redis node owns the key slot.
var ErrNoSlot = errors.New("the slot has no redis node")

type clusterClient struct {
	slots  [16384]conn
	opt    *ClientOption
	conns  map[string]conn
	connFn connFn
	sc     call
	mu     sync.RWMutex
	stop   uint32
	cmd    cmds.Builder
}

func newClusterClient(opt *ClientOption, connFn connFn) (client *clusterClient, err error) {
	client = &clusterClient{
		cmd:    cmds.NewBuilder(cmds.InitSlot),
		opt:    opt,
		connFn: connFn,
		conns:  make(map[string]conn),
	}

	if err = client.init(); err != nil {
		return nil, err
	}

	if err = client.refresh(); err != nil {
		return client, err
	}

	return client, nil
}

func (c *clusterClient) init() error {
	if len(c.opt.InitAddress) == 0 {
		return ErrNoAddr
	}
	results := make(chan error, len(c.opt.InitAddress))
	for _, addr := range c.opt.InitAddress {
		cc := c.connFn(addr, c.opt)
		go func(addr string, cc conn) {
			if err := cc.Dial(); err == nil {
				c.mu.Lock()
				if prev, ok := c.conns[addr]; ok {
					go prev.Close()
				}
				c.conns[addr] = cc
				c.mu.Unlock()
				results <- nil
			} else {
				results <- err
			}
		}(addr, cc)
	}
	es := make([]error, cap(results))
	for i := 0; i < cap(results); i++ {
		if err := <-results; err == nil {
			return nil
		} else {
			es[i] = err
		}
	}
	return es[0]
}

func (c *clusterClient) refresh() (err error) {
	return c.sc.Do(c._refresh)
}

func (c *clusterClient) _refresh() (err error) {
	var reply RedisMessage

retry:
	c.mu.RLock()
	results := make(chan RedisResult, len(c.conns))
	for _, cc := range c.conns {
		go func(c conn) { results <- c.Do(context.Background(), cmds.SlotCmd) }(cc)
	}
	c.mu.RUnlock()

	for i := 0; i < cap(results); i++ {
		if reply, err = (<-results).ToMessage(); len(reply.values) != 0 {
			break
		}
	}

	if err != nil {
		return err
	}

	if len(reply.values) == 0 {
		if err = c.init(); err != nil {
			return err
		}
		goto retry
	}

	groups := parseSlots(reply)

	// TODO support read from replicas
	conns := make(map[string]conn, len(groups))
	for _, g := range groups {
		for _, addr := range g.nodes {
			conns[addr] = c.connFn(addr, c.opt)
		}
	}
	// make sure InitAddress always be present
	for _, addr := range c.opt.InitAddress {
		if _, ok := conns[addr]; !ok {
			conns[addr] = c.connFn(addr, c.opt)
		}
	}

	var removes []conn

	c.mu.RLock()
	for addr, cc := range c.conns {
		if _, ok := conns[addr]; ok {
			conns[addr] = cc
		} else {
			removes = append(removes, cc)
		}
	}
	c.mu.RUnlock()

	slots := [16384]conn{}
	for master, g := range groups {
		cc := conns[master]
		for _, slot := range g.slots {
			for i := slot[0]; i <= slot[1]; i++ {
				slots[i] = cc
			}
		}
	}

	c.mu.Lock()
	c.slots = slots
	c.conns = conns
	c.mu.Unlock()

	for _, cc := range removes {
		go cc.Close()
	}

	return nil
}

func (c *clusterClient) single() conn {
	return c._pick(cmds.InitSlot)
}

func (c *clusterClient) nodes() []string {
	c.mu.RLock()
	nodes := make([]string, 0, len(c.conns))
	for addr := range c.conns {
		nodes = append(nodes, addr)
	}
	c.mu.RUnlock()
	return nodes
}

type group struct {
	nodes []string
	slots [][2]int64
}

func parseSlots(slots RedisMessage) map[string]group {
	groups := make(map[string]group, len(slots.values))
	for _, v := range slots.values {
		master := fmt.Sprintf("%s:%d", v.values[2].values[0].string, v.values[2].values[1].integer)
		g, ok := groups[master]
		if !ok {
			g.slots = make([][2]int64, 0)
			g.nodes = make([]string, 0, len(v.values)-2)
			for i := 2; i < len(v.values); i++ {
				dst := fmt.Sprintf("%s:%d", v.values[i].values[0].string, v.values[i].values[1].integer)
				g.nodes = append(g.nodes, dst)
			}
		}
		g.slots = append(g.slots, [2]int64{v.values[0].integer, v.values[1].integer})
		groups[master] = g
	}
	return groups
}

func (c *clusterClient) _pick(slot uint16) (p conn) {
	c.mu.RLock()
	if slot == cmds.InitSlot {
		for _, cc := range c.conns {
			p = cc
			break
		}
	} else {
		p = c.slots[slot]
	}
	c.mu.RUnlock()
	return p
}

func (c *clusterClient) pick(slot uint16) (p conn, err error) {
	if p = c._pick(slot); p == nil {
		if err := c.refresh(); err != nil {
			return nil, err
		}
		if p = c._pick(slot); p == nil {
			return nil, ErrNoSlot
		}
	}
	return p, nil
}

func (c *clusterClient) redirectOrNew(addr string) (p conn) {
	c.mu.RLock()
	p = c.conns[addr]
	c.mu.RUnlock()
	if p != nil && !p.Is(addr) {
		return p
	}
	c.mu.Lock()
	if p = c.conns[addr]; p == nil {
		p = c.connFn(addr, c.opt)
		c.conns[addr] = p
	} else if p.Is(addr) {
		// try reconnection if the MOVED redirects to the same host,
		// because the same hostname may actually be resolved into another destination
		// depending on the fail-over implementation. ex: AWS MemoryDB's resize process.
		go p.Close()
		p = c.connFn(addr, c.opt)
		c.conns[addr] = p
	}
	c.mu.Unlock()
	return p
}

func (c *clusterClient) B() cmds.Builder {
	return c.cmd
}

func (c *clusterClient) Do(ctx context.Context, cmd cmds.Completed) (resp RedisResult) {
retry:
	cc, err := c.pick(cmd.Slot())
	if err != nil {
		resp = newErrResult(err)
		goto ret
	}
	resp = cc.Do(ctx, cmd)
process:
	if c.shouldRefreshRetry(resp.NonRedisError(), ctx) && cmd.IsReadOnly() {
		goto retry
	}
	if err := resp.RedisError(); err != nil {
		if addr, ok := err.IsMoved(); ok {
			go c.refresh()
			resp = c.redirectOrNew(addr).Do(ctx, cmd)
			goto process
		} else if addr, ok = err.IsAsk(); ok {
			resp = c.redirectOrNew(addr).DoMulti(ctx, cmds.AskingCmd, cmd)[1]
			goto process
		} else if err.IsTryAgain() {
			runtime.Gosched()
			goto retry
		}
	}
ret:
	cmds.Put(cmd.CommandSlice())
	return resp
}

func (c *clusterClient) DoCache(ctx context.Context, cmd cmds.Cacheable, ttl time.Duration) (resp RedisResult) {
retry:
	cc, err := c.pick(cmd.Slot())
	if err != nil {
		resp = newErrResult(err)
		goto ret
	}
	resp = cc.DoCache(ctx, cmd, ttl)
process:
	if c.shouldRefreshRetry(resp.NonRedisError(), ctx) {
		goto retry
	}
	if err := resp.RedisError(); err != nil {
		if addr, ok := err.IsMoved(); ok {
			go c.refresh()
			resp = c.redirectOrNew(addr).DoCache(ctx, cmd, ttl)
			goto process
		} else if addr, ok = err.IsAsk(); ok {
			// TODO ASKING OPT-IN Caching
			resp = c.redirectOrNew(addr).DoMulti(ctx, cmds.AskingCmd, cmds.Completed(cmd))[1]
			goto process
		} else if err.IsTryAgain() {
			runtime.Gosched()
			goto retry
		}
	}
ret:
	cmds.Put(cmd.CommandSlice())
	return resp
}

func (c *clusterClient) Receive(ctx context.Context, subscribe cmds.Completed, fn func(msg PubSubMessage)) (err error) {
retry:
	cc, err := c.pick(subscribe.Slot())
	if err != nil {
		goto ret
	}
	if err = cc.Receive(ctx, subscribe, fn); c.shouldRefreshRetry(err, ctx) {
		goto retry
	}
ret:
	cmds.Put(subscribe.CommandSlice())
	return err
}

func (c *clusterClient) Dedicated(fn func(DedicatedClient) error) (err error) {
	dcc := &dedicatedClusterClient{cmd: c.cmd, client: c, slot: cmds.NoSlot}
	err = fn(dcc)
	dcc.release()
	return err
}

func (c *clusterClient) Dedicate() (DedicatedClient, func()) {
	dcc := &dedicatedClusterClient{cmd: c.cmd, client: c, slot: cmds.NoSlot}
	return dcc, func() { dcc.release() }
}

func (c *clusterClient) Close() {
	atomic.StoreUint32(&c.stop, 1)
	c.mu.RLock()
	for _, cc := range c.conns {
		go cc.Close()
	}
	c.mu.RUnlock()
}

func (c *clusterClient) shouldRefreshRetry(err error, ctx context.Context) (should bool) {
	if should = err != nil && atomic.LoadUint32(&c.stop) == 0; should {
		go c.refresh()
	}
	return should && ctx.Err() == nil
}

type dedicatedClusterClient struct {
	client *clusterClient
	conn   conn
	wire   wire

	mu   sync.Mutex
	cmd  cmds.Builder
	slot uint16
}

func (c *dedicatedClusterClient) check(slot uint16) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.slot == cmds.NoSlot {
		c.slot = slot
	} else if c.slot != slot {
		panic(panicMsgCxSlot)
	}
}

func (c *dedicatedClusterClient) acquire() (wire wire, err error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.wire != nil {
		return c.wire, nil
	}
	if c.slot == cmds.NoSlot {
		panic(panicMsgNoSlot)
	}
	if c.conn, err = c.client.pick(c.slot); err != nil {
		return nil, err
	}
	c.wire = c.conn.Acquire()
	return c.wire, nil
}

func (c *dedicatedClusterClient) release() {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.wire != nil {
		c.conn.Store(c.wire)
	}
}

func (c *dedicatedClusterClient) B() cmds.Builder {
	return c.cmd
}

func (c *dedicatedClusterClient) Do(ctx context.Context, cmd cmds.Completed) (resp RedisResult) {
	c.check(cmd.Slot())
retry:
	if w, err := c.acquire(); err != nil {
		resp = newErrResult(err)
	} else {
		resp = w.Do(ctx, cmd)
		if c.client.shouldRefreshRetry(resp.NonRedisError(), ctx) && cmd.IsReadOnly() && w.Error() == nil {
			goto retry
		}
	}
	cmds.Put(cmd.CommandSlice())
	return resp
}

func (c *dedicatedClusterClient) DoMulti(ctx context.Context, multi ...cmds.Completed) (resp []RedisResult) {
	if len(multi) == 0 {
		return nil
	}
	for _, cmd := range multi {
		c.check(cmd.Slot())
	}
	readonly := allReadOnly(multi)
retry:
	if w, err := c.acquire(); err == nil {
		resp = w.DoMulti(ctx, multi...)
		for _, resp := range resp {
			if c.client.shouldRefreshRetry(resp.NonRedisError(), ctx) && readonly && w.Error() == nil {
				goto retry
			}
		}
	} else {
		resp = make([]RedisResult, len(multi))
		for i := range resp {
			resp[i] = newErrResult(err)
		}
	}
	for _, cmd := range multi {
		cmds.Put(cmd.CommandSlice())
	}
	return resp
}

func (c *dedicatedClusterClient) Receive(ctx context.Context, subscribe cmds.Completed, fn func(msg PubSubMessage)) (err error) {
	c.check(subscribe.Slot())
	var w wire
retry:
	if w, err = c.acquire(); err == nil {
		if err = w.Receive(ctx, subscribe, fn); c.client.shouldRefreshRetry(err, ctx) && w.Error() == nil {
			goto retry
		}
	}
	cmds.Put(subscribe.CommandSlice())
	return err
}

const (
	panicMsgCxSlot = "cross slot command in Dedicated is prohibited"
	panicMsgNoSlot = "the first command should contain the slot key"
)
