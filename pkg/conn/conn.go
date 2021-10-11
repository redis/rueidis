package conn

import (
	"bufio"
	"errors"
	"net"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/rueian/rueidis/internal/cache"
	"github.com/rueian/rueidis/internal/proto"
	"github.com/rueian/rueidis/internal/queue"
)

const DefaultCacheBytes = 128 * (1 << 20)

var (
	OptInCmd       = []string{"CLIENT", "CACHING", "YES"}
	ErrConnClosing = errors.New("connection is closing")
)

type Conn struct {
	waits int32
	state int32

	conn  net.Conn
	queue queue.Queue
	cache cache.Cache
	error atomic.Value

	r *bufio.Reader
	w *bufio.Writer

	info proto.Message
}

type Option struct {
	CacheSize  int
	SelectDB   int
	Username   string
	Password   string
	ClientName string
}

func newConn(conn net.Conn, option Option) (*Conn, error) {
	if option.CacheSize <= 0 {
		option.CacheSize = DefaultCacheBytes
	}

	c := &Conn{
		conn:  conn,
		queue: queue.NewRing(),
		cache: cache.NewLRU(option.CacheSize),
		r:     bufio.NewReader(conn),
		w:     bufio.NewWriter(conn),
	}
	go c.reading()

	helloCmd := []string{"HELLO", "3"}
	if option.Username != "" {
		helloCmd = append(helloCmd, "AUTH", option.Username, option.Password)
	}
	if option.ClientName != "" {
		helloCmd = append(helloCmd, "SETNAME", option.ClientName)
	}

	init := [][]string{helloCmd, {"CLIENT", "TRACKING", "ON", "OPTIN"}}
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

	return c, nil
}

func (c *Conn) reading() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	exit := func() {
		_ = c.conn.Close() // force both read & write goroutine to exit
		wg.Done()
	}
	go func() { // write goroutine
		defer exit()

		var err error
		for atomic.LoadInt32(&c.state) != 2 {
			if one, multi := c.queue.NextCmd(); one != nil {
				err = proto.WriteCmd(c.w, one)
			} else if multi != nil {
				for _, cmd := range multi {
					err = proto.WriteCmd(c.w, cmd)
				}
			} else {
				err = c.w.Flush()
				runtime.Gosched()
			}
			if err != nil {
				c.error.CompareAndSwap(nil, err)
				return
			}
		}
	}()
	go func() { // read goroutine
		defer exit()
		for atomic.LoadInt32(&c.state) != 2 {
			msg, err := proto.ReadNextMessage(c.r)
			if err != nil {
				c.error.CompareAndSwap(nil, err)
				return
			}
			if msg.Type == '>' {
				c.handlePush(msg.Values)
				continue
			}
			_, multi, ch := c.queue.NextResultCh()
			if ch == nil {
				panic("receive unexpected out of band message")
			}
			ch <- proto.Result{Val: msg, Err: err}

			if multi == nil {
				continue
			}

			// in current implementation, the cache opt-in will only be in the first cmd in batch

			// TODO: handle opt-in cache for MULTI, LUA commands
			opted := cmdEqual(multi[0], OptInCmd)

			// read the rest cmd responses
			for i := 1; i < len(multi); {
				msg, err = proto.ReadNextMessage(c.r)
				if err != nil {
					c.error.CompareAndSwap(nil, err)
					for ; i < len(multi); i++ {
						ch <- proto.Result{Val: msg, Err: err}
					}
					return
				}
				if msg.Type == '>' {
					c.handlePush(msg.Values)
					continue
				}
				if msg.Type != '-' && msg.Type != '!' && opted {
					c.cache.Update(multi[i][1], msg)
				}
				ch <- proto.Result{Val: msg, Err: err}
				i++
			}
		}
	}()
	wg.Wait()
	atomic.CompareAndSwapInt32(&c.state, 0, 1)

	err, ok := c.error.Load().(error)
	if !ok {
		err = ErrConnClosing
	}

	// clean up write queue and read queue
	for atomic.LoadInt32(&c.waits) != 0 {
		c.queue.NextCmd()
		if one, multi, ch := c.queue.NextResultCh(); one != nil {
			ch <- proto.Result{Err: err}
		} else {
			for i := 0; i < len(multi); i++ {
				ch <- proto.Result{Err: err}
			}
		}
	}

	atomic.CompareAndSwapInt32(&c.state, 1, 2)
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

func (c *Conn) Info() proto.Message {
	return c.info
}

func (c *Conn) Do(cmd []string) (res proto.Result) {
	atomic.AddInt32(&c.waits, 1)
	if atomic.LoadInt32(&c.state) == 0 {
		res = <-c.queue.PutOne(cmd)
	} else {
		res.Err = c.error.Load().(error)
	}
	atomic.AddInt32(&c.waits, -1)
	return res
}

func (c *Conn) DoMulti(cmds ...[]string) []proto.Result {
	res := make([]proto.Result, len(cmds))
	atomic.AddInt32(&c.waits, 1)
	if atomic.LoadInt32(&c.state) == 0 {
		ch := c.queue.PutMulti(cmds)
		for i := range res {
			res[i] = <-ch
		}
	} else {
		err := c.error.Load().(error)
		for i := 0; i < len(res); i++ {
			res[i].Err = err
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
	return c.DoMulti(OptInCmd, cmd)[1]
}

func (c *Conn) Close() {
	c.error.CompareAndSwap(nil, ErrConnClosing)
	atomic.CompareAndSwapInt32(&c.state, 0, 1)
	for atomic.LoadInt32(&c.waits) != 0 {
		runtime.Gosched()
	}
	atomic.CompareAndSwapInt32(&c.state, 1, 2)
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
