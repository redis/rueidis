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

type placeholder struct {
	addr string
	conn *conn.Conn
}

type ClusterClient struct {
	Cmd *cmds.SBuilder
	opt ClusterClientOption

	mu    sync.RWMutex
	sg    singleflight.Call
	slots [16384]placeholder
	conns map[string]*conn.Conn
}

func NewClusterClient(option ClusterClientOption) (client *ClusterClient, err error) {
	if len(option.InitAddress) == 0 {
		return nil, ErrNoNodes
	}

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
	for _, addr := range c.opt.InitAddress {
		if cc, err = conn.NewConn(addr, c.opt.connOption()); err == nil {
			c.mu.Lock()
			c.conns[addr] = cc
			c.mu.Unlock()
			break
		}
	}
	return cc, err
}

func (c *ClusterClient) refreshSlots() (err error) {
	return c.sg.Do(c._refreshSlots)
}

func (c *ClusterClient) _refreshSlots() (err error) {
	var reply proto.Message
retry:
	c.mu.RLock()
	for addr, cc := range c.conns {
		if resp := cc.Do(cmds.SlotCmd); resp.Err != nil {
			err = resp.Err
			go cc.Close()
			delete(c.conns, addr)
		} else if resp.Val.Type == '-' || resp.Val.Type == '!' {
			err = errors.New(resp.Val.String)
			go cc.Close()
			delete(c.conns, addr)
		} else {
			reply = resp.Val
			break
		}
	}
	c.mu.RUnlock()

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
	for _, g := range groups {
		masters[g.nodes[0]] = nil
	}

	var pending []string
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
	for addr, cc := range masters {
		if cc == nil {
			pending = append(pending, addr)
		}
	}

	slots := [16384]placeholder{}
	for master, g := range groups {
		for _, slot := range g.slots {
			for i := slot[0]; i <= slot[1]; i++ {
				slots[i] = placeholder{addr: master}
			}
		}
	}

	c.mu.Lock()
	c.slots = slots
	c.conns = masters
	c.mu.Unlock()

	for _, addr := range pending {
		go func(addr string, slots [][2]int64) {
			var cc *conn.Conn
			var err error
			for {
				c.mu.RLock()
				cc = c.conns[addr]
				c.mu.RUnlock()
				if cc != nil {
					return
				}
				if cc, err = conn.NewConn(addr, c.opt.connOption()); err == nil {
					c.mu.Lock()
					c.conns[addr] = cc
					for _, slot := range slots {
						for i := slot[0]; i <= slot[1]; i++ {
							if c.slots[i].addr == addr {
								c.slots[i].conn = cc
							}
						}
					}
					c.mu.Unlock()
					return
				}
			}
		}(addr, groups[addr].slots)
	}

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

func (c *ClusterClient) pickConn(slot uint16) (*conn.Conn, error) {
	var p placeholder
retry:
	if slot == cmds.InitSlot {
		c.mu.RLock()
		for addr, cc := range c.conns {
			p = placeholder{addr: addr, conn: cc}
			break
		}
	} else {
		c.mu.RLock()
		p = c.slots[slot]
	}
	c.mu.RUnlock()

	if p.addr == "" {
		if err := c.refreshSlots(); err != nil {
			return nil, err
		}
		c.mu.RLock()
		p = c.slots[slot]
		c.mu.RUnlock()
		if p.addr == "" {
			return nil, ErrNoSlot
		}
	}
	if p.conn == nil {
		runtime.Gosched()
		goto retry
	}
	return p.conn, nil
}

func (c *ClusterClient) Do(cmd cmds.SCompleted) (resp proto.Result) {
retry:
	cc, err := c.pickConn(cmd.Slot())
	if err != nil {
		resp.Err = err
		return
	}
	resp = cc.Do(cmds.Completed(cmd))
process:
	if resp.Val.Type == '-' {
		if strings.HasPrefix(resp.Val.String, "MOVED") {
			if err := c.refreshSlots(); err != nil {
				resp.Err = err
				return
			}
			goto retry
		} else if strings.HasPrefix(resp.Val.String, "ASK") {
			addr := strings.Split(resp.Val.String, " ")[2]
			// TODO single flight
			c.mu.RLock()
			cc = c.conns[addr]
			c.mu.RUnlock()
			if cc == nil {
				runtime.Gosched()
				goto retry
			}
			resp = cc.DoMulti(cmds.AskingCmd, cmds.Completed(cmd))[1]
			goto process
		} else if strings.HasPrefix(resp.Val.String, "TRYAGAIN") {
			runtime.Gosched()
			goto retry
		}
	}
	c.Cmd.Put(cmd.Commands())
	return resp
}

func (c *ClusterClient) DoMulti(multi ...cmds.SCompleted) (resp []proto.Result) {
	ccmd := make([]cmds.Completed, len(multi))
	resp = make([]proto.Result, len(multi))

	slot := cmds.InitSlot
	for i, cmd := range multi {
		ccmd[i] = cmds.Completed(cmd)
		if cmd.Slot() == cmds.InitSlot {
			continue
		}
		if slot == cmds.InitSlot {
			slot = cmd.Slot()
		} else if slot != cmd.Slot() {
			panic("mixed slot commands are not allowed")
		}
	}

retry:
	cc, err := c.pickConn(slot)
	if err != nil {
		for i := range resp {
			resp[i].Err = err
		}
		return
	}
	resp = cc.DoMulti(ccmd...)
process:
	for _, r := range resp {
		if r.Val.Type == '-' {
			if strings.HasPrefix(r.Val.String, "MOVED") {
				if err := c.refreshSlots(); err != nil {
					r.Err = err
					return
				}
				goto retry
			} else if strings.HasPrefix(r.Val.String, "ASK") {
				addr := strings.Split(r.Val.String, " ")[2]
				// TODO single flight
				c.mu.RLock()
				cc = c.conns[addr]
				c.mu.RUnlock()
				if cc == nil {
					runtime.Gosched()
					goto retry
				}
				resp = cc.DoMulti(ccmd...)
				goto process
			} else if strings.HasPrefix(r.Val.String, "TRYAGAIN") {
				runtime.Gosched()
				goto retry
			}
		}
	}
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
		return
	}
	resp = cc.DoCache(cmds.Cacheable(cmd), ttl)
process:
	if resp.Val.Type == '-' {
		if strings.HasPrefix(resp.Val.String, "MOVED") {
			if err := c.refreshSlots(); err != nil {
				resp.Err = err
				return
			}
			goto retry
		} else if strings.HasPrefix(resp.Val.String, "ASK") {
			addr := strings.Split(resp.Val.String, " ")[2]
			// TODO single flight
			c.mu.RLock()
			cc = c.conns[addr]
			c.mu.RUnlock()
			if cc == nil {
				runtime.Gosched()
				goto retry
			}
			// TODO ASKING OPT-IN Caching
			resp = cc.DoMulti(cmds.AskingCmd, cmds.Completed(cmd))[1]
			goto process
		} else if strings.HasPrefix(resp.Val.String, "TRYAGAIN") {
			runtime.Gosched()
			goto retry
		}
	}
	c.Cmd.Put(cmd.Commands())
	return resp
}

func (c *ClusterClient) Close() {
	c.mu.Lock()
	c.slots = [16384]placeholder{}
	for _, cc := range c.conns {
		cc.Close()
	}
	c.conns = nil
	c.mu.Unlock()
}
