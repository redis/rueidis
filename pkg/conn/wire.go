package conn

import (
	"bufio"
	"net"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/rueian/rueidis/internal/cache"
	"github.com/rueian/rueidis/internal/cmds"
	"github.com/rueian/rueidis/internal/proto"
	"github.com/rueian/rueidis/internal/queue"
)

type errWrap struct {
	error
}

type Wire interface {
	Do(cmd cmds.Completed) proto.Result
	DoCache(cmd cmds.Cacheable, ttl time.Duration) proto.Result
	DoMulti(multi ...cmds.Completed) []proto.Result
	Info() map[string]proto.Message
	Error() error
	Close()
}

type wire struct {
	waits int32
	state int32

	conn  net.Conn
	queue queue.Queue
	cache cache.Cache
	error atomic.Value

	r *bufio.Reader
	w *bufio.Writer

	info map[string]proto.Message

	psHandlers PubSubHandlers
}

type PubSubHandlers struct {
	OnMessage      func(channel, message string)
	OnPMessage     func(pattern, channel, message string)
	OnSubscribed   func(channel string, active int64)
	OnUnSubscribed func(channel string, active int64)
}

func newWire(conn net.Conn, option Option) (c *wire, err error) {
	if option.CacheSizeEachConn <= 0 {
		option.CacheSizeEachConn = DefaultCacheBytes
	}

	c = &wire{
		conn:  conn,
		queue: queue.NewRing(),
		cache: cache.NewLRU(option.CacheSizeEachConn),
		r:     bufio.NewReader(conn),
		w:     bufio.NewWriter(conn),

		psHandlers: option.PubSubHandlers,
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

	for i, r := range c.DoMulti(cmds.NewMultiCompleted(init)...) {
		if i == 0 {
			c.info, err = r.ToMap()
		} else {
			err = r.Error()
		}
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}

func (c *wire) reading() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	exit := func() {
		_ = c.conn.Close() // force both read & write goroutine to exit
		wg.Done()
	}
	go func() { // write goroutine
		defer exit()

		var (
			err   error
			ones  = make([]cmds.Completed, 1)
			multi []cmds.Completed
			ch    chan proto.Result
		)

		for atomic.LoadInt32(&c.state) != 2 {
			if ones[0], multi, ch = c.queue.NextWriteCmd(); multi == nil {
				if !ones[0].IsEmpty() {
					multi = ones
				} else {
					if c.w.Buffered() == 0 {
						runtime.Gosched()
					} else {
						err = c.w.Flush()
					}
				}
			}
			for _, cmd := range multi {
				if err = proto.WriteCmd(c.w, cmd.Commands()); cmd.NoReply() {
					ch <- proto.NewErrResult(err)
				}
			}
			if err != nil {
				c.error.CompareAndSwap(nil, &errWrap{error: err})
				return
			}
		}
	}()
	go func() { // read goroutine
		defer exit()

		var (
			err   error
			msg   proto.Message
			tmp   proto.Message
			ones  = make([]cmds.Completed, 1)
			multi []cmds.Completed
			ch    chan proto.Result
			ff    int // fulfilled count
		)

		for {
			if msg, err = proto.ReadNextMessage(c.r); err != nil {
				c.error.CompareAndSwap(nil, &errWrap{error: err})
				return
			}
			if msg.Type == '>' {
				c.handlePush(msg.Values)
				continue
			}
			// if unfulfilled multi commands are lead by opt-in and get success response
			if ff != len(multi) && len(multi) == 3 && multi[0].IsOptIn() {
				if ff == 1 {
					tmp = msg
				} else if ff == 2 {
					cacheable := cmds.Cacheable(multi[ff-1])
					ck, cc := cacheable.CacheKey()
					c.cache.Update(ck, cc, tmp, msg.Integer)
					tmp = proto.Message{}
				}
			}
		nextCMD:
			if ff == len(multi) {
				ff = 0
				ones[0], multi, ch = c.queue.NextResultCh() // ch should not be nil, otherwise it must be a protocol bug
			}
			if multi == nil {
				multi = ones
			}
			if multi[ff].NoReply() {
				ff++
				goto nextCMD
			} else {
				ff++
				ch <- proto.NewResult(msg, err)
			}
		}
	}()
	wg.Wait()

	var (
		ones  = make([]cmds.Completed, 1)
		multi []cmds.Completed
		ch    chan proto.Result
	)

	// clean up write queue and read queue
	atomic.CompareAndSwapInt32(&c.state, 0, 1)
	// clean up cache and free pending calls
	c.cache.FreeAndClose(proto.Message{Type: '-', String: ErrConnClosing.Error()})
	for atomic.LoadInt32(&c.waits) != 0 {
		c.queue.NextWriteCmd()
		if ones[0], multi, ch = c.queue.NextResultCh(); ch == nil {
			runtime.Gosched()
			continue
		}
		if multi == nil {
			multi = ones
		}
		for i := 0; i < len(multi); i++ {
			ch <- proto.NewErrResult(c.Error())
		}
	}
	atomic.CompareAndSwapInt32(&c.state, 1, 2)
}

func (c *wire) handlePush(values []proto.Message) {
	if len(values) < 2 {
		return
	}
	// TODO: handle other push data
	// tracking-redir-broken
	// server-cpu-usage
	switch values[0].String {
	case "invalidate":
		c.cache.Delete(values[1].Values)
	case "message":
		if c.psHandlers.OnMessage != nil {
			c.psHandlers.OnMessage(values[1].String, values[2].String)
		}
	case "pmessage":
		if c.psHandlers.OnPMessage != nil {
			c.psHandlers.OnPMessage(values[1].String, values[2].String, values[3].String)
		}
	case "subscribe", "psubscribe":
		if c.psHandlers.OnSubscribed != nil {
			c.psHandlers.OnSubscribed(values[1].String, values[2].Integer)
		}
	case "unsubscribe", "punsubscribe":
		if c.psHandlers.OnUnSubscribed != nil {
			c.psHandlers.OnUnSubscribed(values[1].String, values[2].Integer)
		}
	}
}

func (c *wire) Info() map[string]proto.Message {
	return c.info
}

func (c *wire) Do(cmd cmds.Completed) (resp proto.Result) {
	atomic.AddInt32(&c.waits, 1)
	if atomic.LoadInt32(&c.state) == 0 {
		resp = <-c.queue.PutOne(cmd)
	} else {
		resp = proto.NewErrResult(c.Error())
	}
	atomic.AddInt32(&c.waits, -1)
	return resp
}

func (c *wire) DoMulti(multi ...cmds.Completed) []proto.Result {
	resp := make([]proto.Result, len(multi))
	atomic.AddInt32(&c.waits, 1)
	if atomic.LoadInt32(&c.state) == 0 {
		ch := c.queue.PutMulti(multi)
		for i := range resp {
			resp[i] = <-ch
		}
	} else {
		err := c.Error()
		for i := 0; i < len(resp); i++ {
			resp[i] = proto.NewErrResult(err)
		}
	}
	atomic.AddInt32(&c.waits, -1)
	return resp
}

func (c *wire) DoCache(cmd cmds.Cacheable, ttl time.Duration) proto.Result {
	ck, cc := cmd.CacheKey()
	if v, entry := c.cache.GetOrPrepare(ck, cc, ttl); v.Type != 0 {
		return proto.NewResult(v, nil)
	} else if entry != nil {
		return proto.NewResult(entry.Wait(), nil)
	}
	return c.DoMulti(cmds.OptInCmd, cmds.Completed(cmd), cmds.NewCompleted([]string{"PTTL", ck}))[1]
}

func (c *wire) Error() error {
	if err, ok := c.error.Load().(*errWrap); ok {
		return err.error
	}
	return nil
}

func (c *wire) Close() {
	swapped := c.error.CompareAndSwap(nil, &errWrap{error: ErrConnClosing})
	atomic.CompareAndSwapInt32(&c.state, 0, 1)
	for atomic.LoadInt32(&c.waits) != 0 {
		runtime.Gosched()
	}
	if swapped {
		<-c.queue.PutOne(cmds.QuitCmd)
	}
	atomic.CompareAndSwapInt32(&c.state, 1, 2)
}
