package rueidis

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"

	"github.com/rueian/rueidis/internal/cmds"
	"github.com/rueian/rueidis/internal/proto"
)

var (
	ErrNoNodes = errors.New("no node to retrieve cluster slots")
	ErrNoSlot  = errors.New("slot not covered")
)

type clusterClient struct {
	cmd *cmds.Builder
	opt ClientOption

	mu     sync.RWMutex
	sc     call
	slots  [16384]conn
	conns  map[string]conn
	connFn connFn
}

func newClusterClient(opt ClientOption, connFn connFn) (client *clusterClient, err error) {
	if opt.ShuffleInit {
		rand.Shuffle(len(opt.InitAddress), func(i, j int) {
			opt.InitAddress[i], opt.InitAddress[j] = opt.InitAddress[j], opt.InitAddress[i]
		})
	}

	client = &clusterClient{
		cmd:    cmds.NewBuilder(cmds.InitSlot),
		opt:    opt,
		connFn: connFn,
		conns:  make(map[string]conn),
	}

	if _, err = client.init(); err != nil {
		return nil, err
	}

	if err = client.refresh(); err != nil {
		client.Close()
		return nil, err
	}

	opt.PubSubOption.installHook(client.cmd, func() (cc conn) {
		var err error
		for cc == nil && err != ErrConnClosing {
			cc, err = client.pick(cmds.InitSlot)
		}
		return cc
	})

	return client, nil
}

func (c *clusterClient) init() (cc conn, err error) {
	if len(c.opt.InitAddress) == 0 {
		return nil, ErrNoNodes
	}
	for _, addr := range c.opt.InitAddress {
		cc = c.connFn(addr, c.opt)
		if err = cc.Dial(); err == nil {
			c.mu.Lock()
			if prev, ok := c.conns[addr]; ok {
				go prev.Close()
			}
			c.conns[addr] = cc
			c.mu.Unlock()
			return cc, nil
		}
	}
	return nil, err
}

func (c *clusterClient) refresh() (err error) {
	return c.sc.Do(c._refresh)
}

func (c *clusterClient) _refresh() (err error) {
	var reply proto.Message
	var dead []string

retry:
	c.mu.RLock()
	for addr, cc := range c.conns {
		if reply, err = cc.Do(cmds.SlotCmd).Value(); err != nil {
			dead = append(dead, addr)
		} else {
			break
		}
	}
	c.mu.RUnlock()

	if len(dead) != 0 {
		c.mu.Lock()
		for _, addr := range dead {
			if cc, ok := c.conns[addr]; ok {
				delete(c.conns, addr)
				go cc.Close()
			}
		}
		c.mu.Unlock()
		dead = nil
	}

	if err != nil {
		return err
	}

	if len(reply.Values) == 0 {
		if _, err = c.init(); err != nil {
			return err
		}
		goto retry
	}

	groups := parseSlots(reply)

	// TODO support read from replicas
	masters := make(map[string]conn, len(groups))
	for addr := range groups {
		masters[addr] = c.connFn(addr, c.opt)
	}

	var removes []conn

	c.mu.RLock()
	for addr, cc := range c.conns {
		if _, ok := masters[addr]; ok {
			masters[addr] = cc
		} else {
			removes = append(removes, cc)
		}
	}
	c.mu.RUnlock()

	slots := [16384]conn{}
	for addr, g := range groups {
		for _, slot := range g.slots {
			for i := slot[0]; i <= slot[1]; i++ {
				slots[i] = masters[addr]
			}
		}
	}

	c.mu.Lock()
	c.slots = slots
	c.conns = masters
	c.mu.Unlock()

	for _, cc := range removes {
		go cc.Close()
	}

	return nil
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

func parseSlots(slots proto.Message) map[string]group {
	groups := make(map[string]group, len(slots.Values))
	for _, v := range slots.Values {
		master := fmt.Sprintf("%s:%d", v.Values[2].Values[0].String, v.Values[2].Values[1].Integer)
		g, ok := groups[master]
		if !ok {
			g.slots = make([][2]int64, 0)
			g.nodes = make([]string, 0, len(v.Values)-2)
			for i := 2; i < len(v.Values); i++ {
				dst := fmt.Sprintf("%s:%d", v.Values[i].Values[0].String, v.Values[i].Values[1].Integer)
				g.nodes = append(g.nodes, dst)
			}
		}
		g.slots = append(g.slots, [2]int64{v.Values[0].Integer, v.Values[1].Integer})
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

func (c *clusterClient) pickOrNew(addr string) (p conn) {
	c.mu.RLock()
	p = c.conns[addr]
	c.mu.RUnlock()
	if p != nil {
		return p
	}
	c.mu.Lock()
	if p = c.conns[addr]; p == nil {
		p = c.connFn(addr, c.opt)
		c.conns[addr] = p
	}
	c.mu.Unlock()
	return p
}

func (c *clusterClient) B() *cmds.Builder {
	return c.cmd
}

func (c *clusterClient) Do(ctx context.Context, cmd cmds.Completed) (resp proto.Result) {
retry:
	cc, err := c.pick(cmd.Slot())
	if err != nil {
		resp = proto.NewErrResult(err)
		goto ret
	}
	resp = cc.Do(cmd)
process:
	if err := resp.RedisError(); err != nil {
		if addr, ok := err.IsMoved(); ok {
			go c.refresh()
			resp = c.pickOrNew(addr).Do(cmd)
			goto process
		} else if addr, ok = err.IsAsk(); ok {
			resp = c.pickOrNew(addr).DoMulti(cmds.AskingCmd, cmd)[1]
			goto process
		} else if err.IsTryAgain() {
			runtime.Gosched()
			goto retry
		}
	}
ret:
	c.cmd.Put(cmd.CommandSlice())
	return resp
}

func (c *clusterClient) DoCache(ctx context.Context, cmd cmds.Cacheable, ttl time.Duration) (resp proto.Result) {
retry:
	cc, err := c.pick(cmd.Slot())
	if err != nil {
		resp = proto.NewErrResult(err)
		goto ret
	}
	resp = cc.DoCache(cmd, ttl)
process:
	if err := resp.RedisError(); err != nil {
		if addr, ok := err.IsMoved(); ok {
			go c.refresh()
			resp = c.pickOrNew(addr).DoCache(cmd, ttl)
			goto process
		} else if addr, ok = err.IsAsk(); ok {
			// TODO ASKING OPT-IN Caching
			resp = c.pickOrNew(addr).DoMulti(cmds.AskingCmd, cmds.Completed(cmd))[1]
			goto process
		} else if err.IsTryAgain() {
			runtime.Gosched()
			goto retry
		}
	}
ret:
	c.cmd.Put(cmd.CommandSlice())
	return resp
}

func (c *clusterClient) Dedicated(fn func(DedicatedClient) error) (err error) {
	dcc := &dedicatedClusterClient{cmd: c.cmd, client: c, slot: cmds.InitSlot}
	err = fn(dcc)
	dcc.release()
	return err
}

func (c *clusterClient) Close() {
	c.mu.RLock()
	for _, cc := range c.conns {
		go cc.Close()
	}
	c.mu.RUnlock()
}

type dedicatedClusterClient struct {
	cmd    *cmds.Builder
	client *clusterClient
	conn   conn
	wire   wire
	slot   uint16
}

func (c *dedicatedClusterClient) check(slot uint16) {
	if slot == cmds.InitSlot {
		return
	}
	if c.slot == cmds.InitSlot {
		c.slot = slot
	} else if c.slot != slot {
		panic(panicMsgCxSlot)
	}
}

func (c *dedicatedClusterClient) acquire() (err error) {
	if c.wire != nil {
		return nil
	}
	if c.slot == cmds.InitSlot {
		panic(panicMsgNoSlot)
	}
	if c.conn, err = c.client.pick(c.slot); err != nil {
		return err
	}
	c.wire = c.conn.Acquire()
	return nil
}

func (c *dedicatedClusterClient) release() {
	if c.wire != nil {
		c.conn.Store(c.wire)
	}
}

func (c *dedicatedClusterClient) B() *cmds.Builder {
	return c.cmd
}

func (c *dedicatedClusterClient) Do(ctx context.Context, cmd cmds.Completed) (resp proto.Result) {
	c.check(cmd.Slot())
	if err := c.acquire(); err != nil {
		return proto.NewErrResult(err)
	} else {
		resp = c.wire.Do(cmd)
	}
	c.cmd.Put(cmd.CommandSlice())
	return resp
}

func (c *dedicatedClusterClient) DoMulti(ctx context.Context, multi ...cmds.Completed) (resp []proto.Result) {
	if len(multi) == 0 {
		return nil
	}
	for _, cmd := range multi {
		c.check(cmd.Slot())
	}
	if err := c.acquire(); err == nil {
		resp = c.wire.DoMulti(multi...)
	} else {
		resp = make([]proto.Result, len(multi))
		for i := range resp {
			resp[i] = proto.NewErrResult(err)
		}
	}
	for _, cmd := range multi {
		c.cmd.Put(cmd.CommandSlice())
	}
	return resp
}

const (
	panicMsgCxSlot = "cross slot command in Dedicated is prohibited"
	panicMsgNoSlot = "the first command should contain the slot key"
)
