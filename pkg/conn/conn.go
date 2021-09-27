package conn

import (
	"bufio"
	"errors"
	"net"
	"runtime"
	"sync/atomic"
	"time"

	"github.com/rueian/rueidis/internal/cache"
	"github.com/rueian/rueidis/internal/proto"
	"github.com/rueian/rueidis/internal/queue"
)

var (
	PingCmd        = []string{"PING"}
	OptInCmd       = []string{"CLIENT", "CACHING", "yes"}
	TrackingCmd    = []string{"CLIENT", "TRACKING", "on", "OPTIN"}
	ErrConnClosing = errors.New("conn is closing")
)

type Conn struct {
	waits int32
	state int32

	conn net.Conn
	err  atomic.Value

	r *bufio.Reader
	w *bufio.Writer
	q *queue.Ring

	cache *cache.LRU

	info proto.Message
}

type Option struct {
	CacheSize  int
	Username   string
	Password   string
	ClientName string
}

func NewConn(conn net.Conn, option Option) (*Conn, error) {
	c := &Conn{
		conn:  conn,
		cache: cache.NewLRU(option.CacheSize),
		r:     bufio.NewReader(conn),
		w:     bufio.NewWriter(conn),
		q:     queue.NewRing(),
	}
	c.reading()

	helloCmd := []string{"HELLO", "3"}
	if option.Username != "" {
		helloCmd = append(helloCmd, "AUTH", option.Username, option.Password)
	}
	if option.ClientName != "" {
		helloCmd = append(helloCmd, "SETNAME", option.ClientName)
	}

	res := c.WriteMulti([][]string{helloCmd, TrackingCmd})
	for _, r := range res {
		if r.Err != nil {
			return nil, r.Err
		}
	}

	c.info = res[0].Val

	// make client side caching be as healthy as possible
	c.pinging()

	return c, nil
}

func (c *Conn) pinging() {
	go func() {
		for atomic.LoadInt32(&c.state) == 0 {
			time.Sleep(time.Second)
			c.WriteOne(PingCmd) // if the ping fail, the client side caching will be cleared
		}
	}()
}

func (c *Conn) reading() {
	go func() {
		var err error
		defer func() {
			// if any error, make sure to clear the cache
			c.cache.DeleteAll()
			if err != nil {
				c.err.Store(err)
			}
			_ = c.conn.Close()
			c.Close()
		}()
		for atomic.LoadInt32(&c.state) != 2 {
			cmd := c.q.TryNextCmd()
			if cmd == nil {
				if err = c.w.Flush(); err != nil {
					return
				}
				cmd = c.q.NextCmd()
			}
			if err = proto.WriteCmd(c.w, cmd); err != nil {
				return
			}
		}
	}()
	go func() {
		var opted bool
		for atomic.LoadInt32(&c.state) != 2 {
			msg, err := proto.ReadNextMessage(c.r)
			if msg.Type == '>' {
				// TODO: handle other push data
				// tracking-redir-broken
				// message
				// pmessage
				// subscribe
				// unsubscribe
				// psubscribe
				// punsubscribe
				// server-cpu-usage
				if len(msg.Values) < 2 {
					continue
				}
				if msg.Values[0].String == "invalidate" {
					c.cache.Delete(msg.Values[1].Values)
				}
				continue
			}
			cmd, ch := c.q.NextResultCh()
			if err == nil && msg.Type != '-' && msg.Type != '!' && opted {
				c.cache.Update(cmd[1], msg)
			}
			opted = cmdEqual(cmd, OptInCmd)
			ch <- proto.Result{Val: msg, Err: err}
		}
	}()
}

func (c *Conn) Into() proto.Message {
	return c.info
}

func (c *Conn) WriteOne(cmd []string) (res proto.Result) {
	atomic.AddInt32(&c.waits, 1)
	if atomic.LoadInt32(&c.state) == 0 {
		res = <-c.q.PutOne(cmd)
	} else {
		res.Err = ErrConnClosing
	}
	atomic.AddInt32(&c.waits, -1)
	return res
}

func (c *Conn) WriteMulti(cmd [][]string) []proto.Result {
	res := make([]proto.Result, len(cmd))
	atomic.AddInt32(&c.waits, 1)
	if atomic.LoadInt32(&c.state) == 0 {
		for i, ch := range c.q.PutMulti(cmd) {
			res[i] = <-ch
		}
	} else {
		for i := 0; i < len(res); i++ {
			res[i].Err = ErrConnClosing
		}
	}
	atomic.AddInt32(&c.waits, -1)
	return res
}

func (c *Conn) WriteOptInCache(cmd []string, ttl time.Duration) proto.Result {
	if v := c.cache.GetOrPrepare(cmd[1], ttl); v.Type != 0 {
		return proto.Result{Val: v}
	}
	return c.WriteMulti([][]string{OptInCmd, cmd})[1]
}

func (c *Conn) Close() {
	atomic.CompareAndSwapInt32(&c.state, 0, 1)
	for atomic.LoadInt32(&c.waits) != 0 {
		runtime.Gosched()
	}
	atomic.CompareAndSwapInt32(&c.state, 1, 2)
}

func (c *Conn) Err() error {
	v := c.err.Load()
	if v == nil {
		return nil
	}
	return v.(error)
}

func cmdEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, s := range a {
		if s != b[i] {
			return false
		}
	}
	return true
}
