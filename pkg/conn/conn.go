package conn

import (
	"bufio"
	"errors"
	"net"
	"runtime"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/rueian/rueidis/internal/cache"
	"github.com/rueian/rueidis/internal/cmds"
	"github.com/rueian/rueidis/internal/proto"
	"github.com/rueian/rueidis/internal/queue"
)

const DefaultCacheBytes = 128 * (1 << 20)

var (
	PingCmd        = []string{"PING"}
	OptInCmd       = []string{"CLIENT", "CACHING", "yes"}
	TrackingCmd    = []string{"CLIENT", "TRACKING", "on", "OPTIN"}
	ErrConnClosing = errors.New("conn is closing")
)

type Conn struct {
	Cmd *cmds.Builder

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
	SelectDB   int
	Username   string
	Password   string
	ClientName string
}

func NewConn(dst string, option Option) (*Conn, error) {
	tcp, err := net.Dial("tcp", dst)
	if err != nil {
		return nil, err
	}
	return newConn(tcp, option)
}

func newConn(conn net.Conn, option Option) (*Conn, error) {
	if option.CacheSize <= 0 {
		option.CacheSize = DefaultCacheBytes
	}

	c := &Conn{
		Cmd: cmds.NewBuilder(),

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

	init := [][]string{helloCmd, TrackingCmd}
	if option.SelectDB != 0 {
		init = append(init, []string{"SELECT", strconv.Itoa(option.SelectDB)})
	}

	res := c.DoMulti(init...)
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
			c.Do(PingCmd) // if the ping fail, the client side caching will be cleared
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
			cmds := c.q.TryNextCmd()
			if cmds == nil {
				if err = c.w.Flush(); err != nil {
					return
				}
				cmds = c.q.NextCmd()
			}
			for _, cmd := range cmds {
				if err = proto.WriteCmd(c.w, cmd); err != nil {
					return
				}
				c.Cmd.Put(cmd)
			}
		}
	}()
	go func() {
		for atomic.LoadInt32(&c.state) != 2 {
			msg, err := proto.ReadNextMessage(c.r)
			if msg.Type == '>' {
				c.handlePush(msg.Values)
				continue
			}
			cmds, ch := c.q.NextResultCh()
			ch <- proto.Result{Val: msg, Err: err}

			// in current implementation, the cache opt-in will only be in the first cmd in batch
			// TODO: handle opt-in cache for MULTI, LUA commands
			opted := cmdEqual(cmds[0], OptInCmd)

			// read the rest cmd responses
			for i := 1; i < len(cmds); {
				msg, err = proto.ReadNextMessage(c.r)
				if msg.Type == '>' {
					c.handlePush(msg.Values)
					continue
				}
				if err == nil && msg.Type != '-' && msg.Type != '!' && opted {
					c.cache.Update(cmds[i][1], msg)
				}
				ch <- proto.Result{Val: msg, Err: err}
				i++
			}
		}
	}()
}

func (c *Conn) handlePush(values []proto.Message) {
	// TODO: handle other push data
	// tracking-redir-broken
	// message
	// pmessage
	// subscribe
	// unsubscribe
	// psubscribe
	// punsubscribe
	// server-cpu-usage
	if len(values) < 2 {
		return
	}
	if values[0].String == "invalidate" {
		c.cache.Delete(values[1].Values)
	}
}

func (c *Conn) Into() proto.Message {
	return c.info
}

func (c *Conn) Do(cmd []string) (res proto.Result) {
	atomic.AddInt32(&c.waits, 1)
	if atomic.LoadInt32(&c.state) == 0 {
		res = <-c.q.PutOne(cmd)
	} else {
		res.Err = ErrConnClosing
	}
	atomic.AddInt32(&c.waits, -1)
	return res
}

func (c *Conn) DoMulti(cmd ...[]string) []proto.Result {
	res := make([]proto.Result, len(cmd))
	atomic.AddInt32(&c.waits, 1)
	if atomic.LoadInt32(&c.state) == 0 {
		ch := c.q.PutMulti(cmd)
		for i := range res {
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

func (c *Conn) DoCache(cmd []string, ttl time.Duration) proto.Result {
retry:
	if v, ch := c.cache.GetOrPrepare(cmd[1], ttl); v.Type != 0 {
		return proto.Result{Val: v}
	} else if ch != nil {
		<-ch
		goto retry
	}
	return c.DoMulti(c.Cmd.ClientCaching().Yes().Build(), cmd)[1]
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
