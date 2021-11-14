package client

import (
	"errors"
	"fmt"
	"math/rand"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/rueian/rueidis/internal/cmds"
	"github.com/rueian/rueidis/internal/proto"
	"github.com/rueian/rueidis/pkg/conn"
	"github.com/rueian/rueidis/pkg/singleflight"
)

var (
	ErrNoNodes = errors.New("no node to retrieve cluster slots")
	ErrNoSlot  = errors.New("slot not covered")
)

type ClusterClientOption struct {
	InitAddress []string
	ShuffleInit bool

	CacheSizeEachConn int
	Username          string
	Password          string
	ClientName        string
	DialTimeout       time.Duration

	PubSubHandlers conn.PubSubHandlers
}

func (option ClusterClientOption) connOption() conn.Option {
	return conn.Option{
		CacheSize:      option.CacheSizeEachConn,
		Username:       option.Username,
		Password:       option.Password,
		ClientName:     option.ClientName,
		DialTimeout:    option.DialTimeout,
		PubSubHandlers: option.PubSubHandlers,
	}
}

type ClusterClient struct {
	Cmd *cmds.SBuilder
	opt ClusterClientOption

	mu    sync.RWMutex
	sg    singleflight.Call
	slots [16384]*conn.Conn
	conns map[string]*conn.Conn
}

func NewClusterClient(option ClusterClientOption) (client *ClusterClient, err error) {
	if option.ShuffleInit {
		rand.Shuffle(len(option.InitAddress), func(i, j int) {
			option.InitAddress[i], option.InitAddress[j] = option.InitAddress[j], option.InitAddress[i]
		})
	}

	client = &ClusterClient{
		Cmd:   cmds.NewSBuilder(),
		opt:   option,
		conns: make(map[string]*conn.Conn),
	}

	if _, err = client.initConn(); err != nil {
		return nil, err
	}

	if err = client.refreshSlots(); err != nil {
		return nil, err
	}

	return client, nil
}

func (c *ClusterClient) initConn() (cc *conn.Conn, err error) {
	if len(c.opt.InitAddress) == 0 {
		return nil, ErrNoNodes
	}
	for _, addr := range c.opt.InitAddress {
		cc = conn.NewConn(addr, c.opt.connOption())
		if err = cc.Dialable(); err == nil {
			c.mu.Lock()
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
		if resp := cc.Do(cmds.SlotCmd); resp.Err != nil {
			err = resp.Err
			dead = append(dead, addr)
		} else if resp.Val.Type == '-' || resp.Val.Type == '!' {
			err = errors.New(resp.Val.String)
			dead = append(dead, addr)
		} else {
			reply = resp.Val
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
	masters := make(map[string]*conn.Conn, len(groups))
	for addr := range groups {
		masters[addr] = conn.NewConn(addr, c.opt.connOption())
	}

	var removes []*conn.Conn

	c.mu.RLock()
	for addr, cc := range c.conns {
		if _, ok := masters[addr]; ok {
			masters[addr] = cc
		} else {
			removes = append(removes, cc)
		}
	}
	c.mu.RUnlock()

	slots := [16384]*conn.Conn{}
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

func (c *ClusterClient) pick(slot uint16) (p *conn.Conn) {
	if slot == cmds.InitSlot {
		c.mu.RLock()
		for _, cc := range c.conns {
			p = cc
			break
		}
	} else {
		c.mu.RLock()
		p = c.slots[slot]
	}
	c.mu.RUnlock()
	return p
}

func (c *ClusterClient) pickConn(slot uint16) (p *conn.Conn, err error) {
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

func (c *ClusterClient) pickOrNewConn(addr string) (p *conn.Conn) {
	c.mu.RLock()
	p = c.conns[addr]
	c.mu.RUnlock()
	if p != nil {
		return p
	}
	c.mu.Lock()
	if p = c.conns[addr]; p == nil {
		p = conn.NewConn(addr, c.opt.connOption())
		c.conns[addr] = p
	}
	c.mu.Unlock()
	return p
}

func (c *ClusterClient) Do(cmd cmds.SCompleted) (resp proto.Result) {
retry:
	cc, err := c.pickConn(cmd.Slot())
	if err != nil {
		resp.Err = err
		goto ret
	}
	resp = cc.Do(cmds.Completed(cmd))
process:
	if resp.Val.Type == '-' {
		if strings.HasPrefix(resp.Val.String, "MOVED") {
			go c.refreshSlots()
			addr := strings.Split(resp.Val.String, " ")[2]
			resp = c.pickOrNewConn(addr).Do(cmds.Completed(cmd))
			goto process
		} else if strings.HasPrefix(resp.Val.String, "ASK") {
			addr := strings.Split(resp.Val.String, " ")[2]
			resp = c.pickOrNewConn(addr).DoMulti(cmds.AskingCmd, cmds.Completed(cmd))[1]
			goto process
		} else if strings.HasPrefix(resp.Val.String, "TRYAGAIN") {
			runtime.Gosched()
			goto retry
		}
	}
ret:
	c.Cmd.Put(cmd.Commands())
	return resp
}

func checkMultiSlot(multi []cmds.SCompleted) (slot uint16) {
	slot = cmds.InitSlot
	for _, cmd := range multi {
		if cmd.Slot() == cmds.InitSlot {
			continue
		}
		if slot == cmds.InitSlot {
			slot = cmd.Slot()
		} else if slot != cmd.Slot() {
			panic("mixed slot commands are not allowed")
		}
	}
	return
}

func (c *ClusterClient) DoMulti(multi ...cmds.SCompleted) (resp []proto.Result) {
	slot := checkMultiSlot(multi)

	resp = make([]proto.Result, len(multi))
	ccmd := make([]cmds.Completed, len(multi))

	for i, cmd := range multi {
		ccmd[i] = cmds.Completed(cmd)
	}

retry:
	cc, err := c.pickConn(slot)
	if err != nil {
		for i := range resp {
			resp[i].Err = err
		}
		goto ret
	}
	resp = cc.DoMulti(ccmd...)
process:
	for i, r := range resp {
		if r.Val.Type == '-' {
			if strings.HasPrefix(r.Val.String, "MOVED") {
				go c.refreshSlots()
				addr := strings.Split(r.Val.String, " ")[2]
				resp = c.pickOrNewConn(addr).DoMulti(ccmd...)
				goto process
			} else if strings.HasPrefix(r.Val.String, "ASK") {
				addr := strings.Split(r.Val.String, " ")[2]
				resp[i] = c.pickOrNewConn(addr).DoMulti(cmds.AskingCmd, ccmd[i])[1]
			} else if strings.HasPrefix(r.Val.String, "TRYAGAIN") {
				runtime.Gosched()
				goto retry
			}
		}
	}
ret:
	for _, cmd := range ccmd {
		c.Cmd.Put(cmd.Commands())
	}
	return resp
}

func (c *ClusterClient) DoCache(cmd cmds.SCacheable, ttl time.Duration) (resp proto.Result) {
retry:
	cc, err := c.pickConn(cmd.Slot())
	if err != nil {
		resp.Err = err
		goto ret
	}
	resp = cc.DoCache(cmds.Cacheable(cmd), ttl)
process:
	if resp.Val.Type == '-' {
		if strings.HasPrefix(resp.Val.String, "MOVED") {
			go c.refreshSlots()
			addr := strings.Split(resp.Val.String, " ")[2]
			resp = c.pickOrNewConn(addr).DoCache(cmds.Cacheable(cmd), ttl)
			goto process
		} else if strings.HasPrefix(resp.Val.String, "ASK") {
			addr := strings.Split(resp.Val.String, " ")[2]
			// TODO ASKING OPT-IN Caching
			resp = c.pickOrNewConn(addr).DoMulti(cmds.AskingCmd, cmds.Completed(cmd))[1]
			goto process
		} else if strings.HasPrefix(resp.Val.String, "TRYAGAIN") {
			runtime.Gosched()
			goto retry
		}
	}
ret:
	c.Cmd.Put(cmd.Commands())
	return resp
}

func (c *ClusterClient) Close() {
	c.mu.RLock()
	for _, cc := range c.conns {
		go cc.Close()
	}
	c.mu.RUnlock()
}
