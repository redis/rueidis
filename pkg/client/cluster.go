package client

import (
	"errors"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
	"unsafe"

	"github.com/rueian/rueidis/internal/cmds"
	"github.com/rueian/rueidis/internal/proto"
	"github.com/rueian/rueidis/pkg/conn"
	"github.com/rueian/rueidis/pkg/om"
	"github.com/rueian/rueidis/pkg/script"
	"github.com/rueian/rueidis/pkg/singleflight"
)

var (
	ErrNoNodes = errors.New("no node to retrieve cluster slots")
	ErrNoSlot  = errors.New("slot not covered")
)

type ClusterClientOption struct {
	InitAddress []string
	ShuffleInit bool
	ConnOption  conn.Option
}

type ClusterClient struct {
	Cmd    *cmds.SBuilder
	opt    ClusterClientOption
	connFn conn.ConnFn

	mu    sync.RWMutex
	sg    singleflight.Call
	slots [16384]conn.Conn
	conns map[string]conn.Conn
}

func NewClusterClient(option ClusterClientOption, connFn conn.ConnFn) (client *ClusterClient, err error) {
	if option.ShuffleInit {
		rand.Shuffle(len(option.InitAddress), func(i, j int) {
			option.InitAddress[i], option.InitAddress[j] = option.InitAddress[j], option.InitAddress[i]
		})
	}

	client = &ClusterClient{
		Cmd:    cmds.NewSBuilder(),
		opt:    option,
		connFn: connFn,
		conns:  make(map[string]conn.Conn),
	}

	if _, err = client.initConn(); err != nil {
		return nil, err
	}

	if err = client.refreshSlots(); err != nil {
		return nil, err
	}

	return client, nil
}

func (c *ClusterClient) initConn() (cc conn.Conn, err error) {
	if len(c.opt.InitAddress) == 0 {
		return nil, ErrNoNodes
	}
	for _, addr := range c.opt.InitAddress {
		cc = c.connFn(addr, c.opt.ConnOption)
		if err = cc.Dialable(); err == nil {
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

func (c *ClusterClient) refreshSlots() (err error) {
	return c.sg.Do(c._refreshSlots)
}

func (c *ClusterClient) _refreshSlots() (err error) {
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
		if _, err = c.initConn(); err != nil {
			return err
		}
		goto retry
	}

	groups := parseSlots(reply)

	// TODO support read from replicas
	masters := make(map[string]conn.Conn, len(groups))
	for addr := range groups {
		masters[addr] = c.connFn(addr, c.opt.ConnOption)
	}

	var removes []conn.Conn

	c.mu.RLock()
	for addr, cc := range c.conns {
		if _, ok := masters[addr]; ok {
			masters[addr] = cc
		} else {
			removes = append(removes, cc)
		}
	}
	c.mu.RUnlock()

	slots := [16384]conn.Conn{}
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

func (c *ClusterClient) nodes() []string {
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

func (c *ClusterClient) pick(slot uint16) (p conn.Conn) {
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

func (c *ClusterClient) pickConn(slot uint16) (p conn.Conn, err error) {
	if p = c.pick(slot); p == nil {
		if err := c.refreshSlots(); err != nil {
			return nil, err
		}
		if p = c.pick(slot); p == nil {
			return nil, ErrNoSlot
		}
	}
	return p, nil
}

func (c *ClusterClient) pickOrNewConn(addr string) (p conn.Conn) {
	c.mu.RLock()
	p = c.conns[addr]
	c.mu.RUnlock()
	if p != nil {
		return p
	}
	c.mu.Lock()
	if p = c.conns[addr]; p == nil {
		p = c.connFn(addr, c.opt.ConnOption)
		c.conns[addr] = p
	}
	c.mu.Unlock()
	return p
}

func (c *ClusterClient) Do(cmd cmds.SCompleted) (resp proto.Result) {
retry:
	cc, err := c.pickConn(cmd.Slot())
	if err != nil {
		resp = proto.NewErrResult(err)
		goto ret
	}
	resp = cc.Do(cmds.Completed(cmd))
process:
	if err := resp.RedisError(); err != nil {
		if addr, ok := err.IsMoved(); ok {
			go c.refreshSlots()
			resp = c.pickOrNewConn(addr).Do(cmds.Completed(cmd))
			goto process
		} else if addr, ok = err.IsAsk(); ok {
			resp = c.pickOrNewConn(addr).DoMulti(cmds.AskingCmd, cmds.Completed(cmd))[1]
			goto process
		} else if err.IsTryAgain() {
			runtime.Gosched()
			goto retry
		}
	}
ret:
	c.Cmd.Put(cmd.Commands())
	return resp
}

func (c *ClusterClient) DoCache(cmd cmds.SCacheable, ttl time.Duration) (resp proto.Result) {
retry:
	cc, err := c.pickConn(cmd.Slot())
	if err != nil {
		resp = proto.NewErrResult(err)
		goto ret
	}
	resp = cc.DoCache(cmds.Cacheable(cmd), ttl)
process:
	if err := resp.RedisError(); err != nil {
		if addr, ok := err.IsMoved(); ok {
			go c.refreshSlots()
			resp = c.pickOrNewConn(addr).DoCache(cmds.Cacheable(cmd), ttl)
			goto process
		} else if addr, ok = err.IsAsk(); ok {
			// TODO ASKING OPT-IN Caching
			resp = c.pickOrNewConn(addr).DoMulti(cmds.AskingCmd, cmds.Completed(cmd))[1]
			goto process
		} else if err.IsTryAgain() {
			runtime.Gosched()
			goto retry
		}
	}
ret:
	c.Cmd.Put(cmd.Commands())
	return resp
}

func (c *ClusterClient) DedicatedWire(fn func(*DedicatedClusterClient) error) (err error) {
	dcc := &DedicatedClusterClient{client: c, slot: cmds.InitSlot}
	err = fn(dcc)
	dcc.release()
	return err
}

func (c *ClusterClient) NewLuaScript(body string) *script.Lua {
	return script.NewLuaScript(body, c.eval, c.evalSha)
}

func (c *ClusterClient) NewLuaScriptReadOnly(body string) *script.Lua {
	return script.NewLuaScript(body, c.evalRo, c.evalShaRo)
}

func (c *ClusterClient) eval(body string, keys, args []string) proto.Result {
	return c.Do(c.Cmd.Eval().Script(body).Numkeys(int64(len(keys))).Key(keys...).Arg(args...).Build())
}

func (c *ClusterClient) evalSha(sha string, keys, args []string) proto.Result {
	return c.Do(c.Cmd.Evalsha().Sha1(sha).Numkeys(int64(len(keys))).Key(keys...).Arg(args...).Build())
}

func (c *ClusterClient) evalRo(body string, keys, args []string) proto.Result {
	return c.Do(c.Cmd.EvalRo().Script(body).Numkeys(int64(len(keys))).Key(keys...).Arg(args...).Build())
}

func (c *ClusterClient) evalShaRo(sha string, keys, args []string) proto.Result {
	return c.Do(c.Cmd.EvalshaRo().Sha1(sha).Numkeys(int64(len(keys))).Key(keys...).Arg(args...).Build())
}

func (c *ClusterClient) NewHashRepository(prefix string, schema interface{}) *om.HashRepository {
	return om.NewHashRepository(prefix, schema, &HashObjectClusterClientAdapter{c: c}, c.NewLuaScript)
}

func (c *ClusterClient) Close() {
	c.mu.RLock()
	for _, cc := range c.conns {
		go cc.Close()
	}
	c.mu.RUnlock()
}

type DedicatedClusterClient struct {
	client *ClusterClient
	conn   conn.Conn
	wire   conn.Wire
	slot   uint16
}

func (c *DedicatedClusterClient) check(slot uint16) {
	if slot == cmds.InitSlot {
		return
	}
	if c.slot == cmds.InitSlot {
		c.slot = slot
	} else if c.slot != slot {
		panic("cross slot command in DedicatedWire is prohibited")
	}
}

func (c *DedicatedClusterClient) acquire() (err error) {
	if c.wire != nil {
		return nil
	}
	if c.slot == cmds.InitSlot {
		panic("the first command in DedicatedClusterClient should contain the slot key")
	}
	if c.conn, err = c.client.pickConn(c.slot); err != nil {
		return err
	}
	c.wire = c.conn.Acquire()
	return nil
}

func (c *DedicatedClusterClient) release() {
	if c.wire != nil {
		c.conn.Store(c.wire)
	}
}

func (c *DedicatedClusterClient) Do(cmd cmds.SCompleted) (resp proto.Result) {
	c.check(cmd.Slot())
	if err := c.acquire(); err != nil {
		return proto.NewErrResult(err)
	} else {
		resp = c.wire.Do(cmds.Completed(cmd))
	}
	c.client.Cmd.Put(cmd.Commands())
	return resp
}

func (c *DedicatedClusterClient) DoMulti(multi ...cmds.SCompleted) (resp []proto.Result) {
	if len(multi) == 0 {
		return nil
	}
	for _, cmd := range multi {
		c.check(cmd.Slot())
	}
	if err := c.acquire(); err == nil {
		resp = c.wire.DoMulti(unsafe.Slice((*cmds.Completed)(&multi[0]), len(multi))...)
	} else {
		resp = make([]proto.Result, len(multi))
		for i := range resp {
			resp[i] = proto.NewErrResult(err)
		}
	}
	for _, cmd := range multi {
		c.client.Cmd.Put(cmd.Commands())
	}
	return resp
}

type HashObjectClusterClientAdapter struct {
	c *ClusterClient
}

func (h *HashObjectClusterClientAdapter) Save(key string, fields map[string]string) error {
	cmd := h.c.Cmd.Hset().Key(key).FieldValue()
	for f, v := range fields {
		cmd = cmd.FieldValue(f, v)
	}
	return h.c.Do(cmd.Build()).Error()
}

func (h *HashObjectClusterClientAdapter) Fetch(key string) (map[string]proto.Message, error) {
	return h.c.Do(h.c.Cmd.Hgetall().Key(key).Build()).ToMap()
}

func (h *HashObjectClusterClientAdapter) FetchCache(key string, ttl time.Duration) (map[string]proto.Message, error) {
	return h.c.DoCache(h.c.Cmd.Hgetall().Key(key).Cache(), ttl).ToMap()
}

func (h *HashObjectClusterClientAdapter) Remove(key string) error {
	return h.c.Do(h.c.Cmd.Del().Key(key).Build()).Error()
}
