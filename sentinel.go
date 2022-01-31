package rueidis

import (
	"container/list"
	"context"
	"errors"
	"fmt"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/rueian/rueidis/internal/cmds"
)

func newSentinelClient(opt ClientOption, connFn connFn) (client *sentinelClient, err error) {
	events := make(chan sentinelEvent, 1)

	client = &sentinelClient{
		cmd:       cmds.NewBuilder(cmds.NoSlot),
		opt:       opt,
		connFn:    connFn,
		events:    events,
		sentinels: list.New(),
	}

	for _, sentinel := range opt.InitAddress {
		client.sentinels.PushBack(sentinel)
	}

	go func() {
		for event := range events {
			switch event.channel {
			case "+sentinel":
				m := strings.SplitN(event.message, " ", 4)
				client.addSentinel(fmt.Sprintf("%s:%s", m[2], m[3]))
			case "+switch-master":
				m := strings.SplitN(event.message, " ", 5)
				if m[0] == opt.Sentinel.MasterSet {
					client.switchMasterRetry(fmt.Sprintf("%s:%s", m[3], m[4]))
				}
			case "+reboot":
				m := strings.SplitN(event.message, " ", 4)
				if m[0] == "master" && m[1] == opt.Sentinel.MasterSet {
					client.switchMasterRetry(fmt.Sprintf("%s:%s", m[2], m[3]))
				}
			}
		}
	}()

	if err = client.refresh(); err != nil {
		client.Close()
		return nil, err
	}

	return client, nil
}

type sentinelClient struct {
	cmd *cmds.Builder
	opt ClientOption

	connFn connFn
	events chan sentinelEvent

	masterAddr string
	masterConn atomic.Value

	mu        sync.Mutex
	sentinels *list.List
	conn      conn
	addr      string
	closed    uint32

	sc call
}

func (c *sentinelClient) B() *cmds.Builder {
	return c.cmd
}

func (c *sentinelClient) Do(ctx context.Context, cmd cmds.Completed) (resp RedisResult) {
retry:
	resp = c.masterConn.Load().(conn).Do(cmd)
	if c.shouldRetry(resp.NonRedisError()) {
		goto retry
	}
	c.cmd.Put(cmd.CommandSlice())
	return resp
}

func (c *sentinelClient) DoCache(ctx context.Context, cmd cmds.Cacheable, ttl time.Duration) (resp RedisResult) {
retry:
	resp = c.masterConn.Load().(conn).DoCache(cmd, ttl)
	if c.shouldRetry(resp.NonRedisError()) {
		goto retry
	}
	c.cmd.Put(cmd.CommandSlice())
	return resp
}

func (c *sentinelClient) Dedicated(fn func(DedicatedClient) error) (err error) {
	master := c.masterConn.Load().(conn)
	wire := master.Acquire()
	err = fn(&dedicatedSingleClient{cmd: c.cmd, wire: wire})
	master.Store(wire)
	return err
}

func (c *sentinelClient) Close() {
	atomic.StoreUint32(&c.closed, 1)
	c.mu.Lock()
	if c.conn != nil {
		c.conn.Close()
	}
	if master := c.masterConn.Load(); master != nil {
		master.(conn).Close()
	}
	if c.events != nil {
		close(c.events)
		c.events = nil
	}
	c.mu.Unlock()

}

func (c *sentinelClient) shouldRetry(err error) (should bool) {
	if should = err == ErrClosing && atomic.LoadUint32(&c.closed) == 0; should {
		runtime.Gosched()
	}
	return should
}

func (c *sentinelClient) addSentinel(addr string) {
	c.mu.Lock()
	c._addSentinel(addr)
	c.mu.Unlock()
}

func (c *sentinelClient) _addSentinel(addr string) {
	for e := c.sentinels.Front(); e != nil; e = e.Next() {
		if e.Value.(string) == addr {
			return
		}
	}
	c.sentinels.PushFront(addr)
}

func (c *sentinelClient) switchMasterRetry(addr string) {
	c.mu.Lock()
	err := c._switchMaster(addr)
	c.mu.Unlock()
	if err != nil {
		go c.refreshRetry()
	}
}

func (c *sentinelClient) _switchMaster(addr string) (err error) {
	var master conn
	if atomic.LoadUint32(&c.closed) == 1 {
		return nil
	}
	if c.masterAddr == addr {
		master = c.masterConn.Load().(conn)
		if master.Error() != nil {
			master = nil
		}
	}
	if master == nil {
		master = c.connFn(addr, c.opt)
		if err = setupSingleConn(c.cmd, master, c.opt); err != nil {
			return err
		}
	}
	if resp, err := master.Do(cmds.RoleCmd).ToArray(); err != nil {
		master.Close()
		return err
	} else if resp[0].string != "master" {
		master.Close()
		return errNotMaster
	}
	c.masterAddr = addr
	if old := c.masterConn.Swap(master); old != nil {
		if prev := old.(conn); prev != master {
			prev.Close()
		}
	}
	return nil
}

func (c *sentinelClient) refreshRetry() {
retry:
	if err := c.refresh(); err != nil {
		goto retry
	}
}

func (c *sentinelClient) refresh() (err error) {
	return c.sc.Do(c._refresh)
}

func (c *sentinelClient) _refresh() (err error) {
	var master string
	var sentinels []string
	var opt = c.sentinelOpt(c.events)

	c.mu.Lock()
	head := c.sentinels.Front()
	for e := head; e != nil; {
		if atomic.LoadUint32(&c.closed) == 1 {
			c.mu.Unlock()
			return nil
		}
		addr := e.Value.(string)

		if c.addr != addr || c.conn == nil || c.conn.Error() != nil {
			if c.conn != nil {
				c.conn.Close()
			}
			c.addr = addr
			c.conn = c.connFn(addr, opt)
			err = c.conn.Dial()
		}
		if err == nil {
			if master, sentinels, err = c.listWatch(c.conn); err == nil {
				c.conn.OnDisconnected(func(prev error) {
					if prev != ErrClosing {
						c.refreshRetry()
					}
				})
				for _, sentinel := range sentinels {
					c._addSentinel(sentinel)
				}
				if err = c._switchMaster(master); err == nil {
					break
				}
			}
			c.conn.Close()
		}
		c.sentinels.MoveToBack(e)
		if e = c.sentinels.Front(); e == head {
			break
		}
	}
	c.mu.Unlock()

	if err == nil {
		if master := c.masterConn.Load(); master == nil {
			err = ErrNoAddr
		} else {
			err = master.(conn).Error()
		}
	}
	return err
}

func (c *sentinelClient) sentinelOpt(ch chan sentinelEvent) (o ClientOption) {
	o = c.opt
	o.Username = o.Sentinel.Username
	o.Password = o.Sentinel.Password
	o.ClientName = o.Sentinel.ClientName
	o.Dialer = o.Sentinel.Dialer
	o.TLSConfig = o.Sentinel.TLSConfig
	o.PubSubOption.onMessage = func(channel, message string) {
		ch <- sentinelEvent{channel: channel, message: message}
	}
	return o
}

func (c *sentinelClient) listWatch(cc conn) (master string, sentinels []string, err error) {
	sentinelsCMD := c.cmd.SentinelSentinels().Master(c.opt.Sentinel.MasterSet).Build()
	getMasterCMD := c.cmd.SentinelGetMasterAddrByName().Master(c.opt.Sentinel.MasterSet).Build()
	defer func() {
		c.cmd.Put(sentinelsCMD.CommandSlice())
		c.cmd.Put(getMasterCMD.CommandSlice())
	}()

	if err = cc.Do(cmds.SentinelSubscribe).Error(); err != nil {
		return "", nil, err
	}
	resp := cc.DoMulti(sentinelsCMD, getMasterCMD)
	others, err := resp[0].ToArray()
	if err != nil {
		return "", nil, err
	}
	for _, other := range others {
		if m, err := other.AsStrMap(); err == nil {
			sentinels = append(sentinels, fmt.Sprintf("%s:%s", m["ip"], m["port"]))
		}
	}
	m, err := resp[1].AsStrSlice()
	if err != nil {
		return "", nil, err
	}
	return fmt.Sprintf("%s:%s", m[0], m[1]), sentinels, nil
}

type sentinelEvent struct {
	channel string
	message string
}

var errNotMaster = errors.New("the redis is not master")
