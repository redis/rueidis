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

func newSentinelClient(opt *ClientOption, connFn connFn) (client *sentinelClient, err error) {

	client = &sentinelClient{
		cmd:       cmds.NewBuilder(cmds.NoSlot),
		mOpt:      opt,
		sOpt:      newSentinelOpt(opt),
		connFn:    connFn,
		sentinels: list.New(),
	}

	for _, sentinel := range opt.InitAddress {
		client.sentinels.PushBack(sentinel)
	}

	if err = client.refresh(); err != nil {
		client.Close()
		return nil, err
	}

	return client, nil
}

type sentinelClient struct {
	mConn     atomic.Value
	sConn     conn
	mOpt      *ClientOption
	sOpt      *ClientOption
	connFn    connFn
	sentinels *list.List
	mAddr     string
	sAddr     string
	sc        call
	mu        sync.Mutex
	closed    uint32
	cmd       cmds.Builder
}

func (c *sentinelClient) B() cmds.Builder {
	return c.cmd
}

func (c *sentinelClient) Do(ctx context.Context, cmd cmds.Completed) (resp RedisResult) {
retry:
	resp = c.mConn.Load().(conn).Do(ctx, cmd)
	if c.shouldRetry(resp.NonRedisError()) {
		goto retry
	}
	cmds.Put(cmd.CommandSlice())
	return resp
}

func (c *sentinelClient) DoCache(ctx context.Context, cmd cmds.Cacheable, ttl time.Duration) (resp RedisResult) {
retry:
	resp = c.mConn.Load().(conn).DoCache(ctx, cmd, ttl)
	if c.shouldRetry(resp.NonRedisError()) {
		goto retry
	}
	cmds.Put(cmd.CommandSlice())
	return resp
}

func (c *sentinelClient) Receive(ctx context.Context, subscribe cmds.Completed, fn func(msg PubSubMessage)) (err error) {
retry:
	if err = c.mConn.Load().(conn).Receive(ctx, subscribe, fn); c.shouldRetry(err) {
		goto retry
	}
	cmds.Put(subscribe.CommandSlice())
	return err
}

func (c *sentinelClient) Dedicated(fn func(DedicatedClient) error) (err error) {
	master := c.mConn.Load().(conn)
	wire := master.Acquire()
	err = fn(&dedicatedSingleClient{cmd: c.cmd, wire: wire})
	master.Store(wire)
	return err
}

func (c *sentinelClient) Close() {
	atomic.StoreUint32(&c.closed, 1)
	c.mu.Lock()
	if c.sConn != nil {
		c.sConn.Close()
	}
	if master := c.mConn.Load(); master != nil {
		master.(conn).Close()
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
	if c.mAddr == addr {
		master = c.mConn.Load().(conn)
		if master.Error() != nil {
			master = nil
		}
	}
	if master == nil {
		master = c.connFn(addr, c.mOpt)
		if err = master.Dial(); err != nil {
			return err
		}
	}
	if resp, err := master.Do(context.Background(), cmds.RoleCmd).ToArray(); err != nil {
		master.Close()
		return err
	} else if resp[0].string != "master" {
		master.Close()
		return errNotMaster
	}
	c.mAddr = addr
	if old := c.mConn.Swap(master); old != nil {
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

	c.mu.Lock()
	head := c.sentinels.Front()
	for e := head; e != nil; {
		if atomic.LoadUint32(&c.closed) == 1 {
			c.mu.Unlock()
			return nil
		}
		addr := e.Value.(string)

		if c.sAddr != addr || c.sConn == nil || c.sConn.Error() != nil {
			if c.sConn != nil {
				c.sConn.Close()
			}
			c.sAddr = addr
			c.sConn = c.connFn(addr, c.sOpt)
			err = c.sConn.Dial()
		}
		if err == nil {
			if master, sentinels, err = c.listWatch(c.sConn); err == nil {
				for _, sentinel := range sentinels {
					c._addSentinel(sentinel)
				}
				if err = c._switchMaster(master); err == nil {
					break
				}
			}
			c.sConn.Close()
		}
		c.sentinels.MoveToBack(e)
		if e = c.sentinels.Front(); e == head {
			break
		}
	}
	c.mu.Unlock()

	if err == nil {
		if master := c.mConn.Load(); master == nil {
			err = ErrNoAddr
		} else {
			err = master.(conn).Error()
		}
	}
	return err
}

func (c *sentinelClient) listWatch(cc conn) (master string, sentinels []string, err error) {
	ctx := context.Background()
	sentinelsCMD := c.cmd.SentinelSentinels().Master(c.mOpt.Sentinel.MasterSet).Build()
	getMasterCMD := c.cmd.SentinelGetMasterAddrByName().Master(c.mOpt.Sentinel.MasterSet).Build()
	defer func() {
		cmds.Put(sentinelsCMD.CommandSlice())
		cmds.Put(getMasterCMD.CommandSlice())
	}()

	go func() {
		if err := cc.Receive(ctx, cmds.SentinelSubscribe, func(event PubSubMessage) {
			switch event.Channel {
			case "+sentinel":
				m := strings.SplitN(event.Message, " ", 4)
				c.addSentinel(fmt.Sprintf("%s:%s", m[2], m[3]))
			case "+switch-master":
				m := strings.SplitN(event.Message, " ", 5)
				if m[0] == c.sOpt.Sentinel.MasterSet {
					c.switchMasterRetry(fmt.Sprintf("%s:%s", m[3], m[4]))
				}
			case "+reboot":
				m := strings.SplitN(event.Message, " ", 4)
				if m[0] == "master" && m[1] == c.sOpt.Sentinel.MasterSet {
					c.switchMasterRetry(fmt.Sprintf("%s:%s", m[2], m[3]))
				}
			}
		}); err != nil && err != ErrClosing {
			c.refreshRetry()
		}
	}()

	resp := cc.DoMulti(ctx, sentinelsCMD, getMasterCMD)
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

func newSentinelOpt(opt *ClientOption) *ClientOption {
	o := *opt
	o.Username = o.Sentinel.Username
	o.Password = o.Sentinel.Password
	o.ClientName = o.Sentinel.ClientName
	o.Dialer = o.Sentinel.Dialer
	o.TLSConfig = o.Sentinel.TLSConfig
	return &o
}

var errNotMaster = errors.New("the redis is not master")
