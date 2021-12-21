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

type errs struct {
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

var _ Wire = (*pipe)(nil)

type pipe struct {
	waits int32
	state int32

	once  sync.Once
	conn  net.Conn
	queue queue.Queue
	cache cache.Cache
	error atomic.Value

	r *bufio.Reader
	w *bufio.Writer

	info map[string]proto.Message

	cbs PubSubHandlers
}

type PubSubHandlers struct {
	OnMessage      func(channel, message string)
	OnPMessage     func(pattern, channel, message string)
	OnSubscribed   func(channel string, active int64)
	OnUnSubscribed func(channel string, active int64)
}

func newPipe(conn net.Conn, option Option) (p *pipe, err error) {
	if option.CacheSizeEachConn <= 0 {
		option.CacheSizeEachConn = DefaultCacheBytes
	}

	p = &pipe{
		conn:  conn,
		queue: queue.NewRing(),
		cache: cache.NewLRU(option.CacheSizeEachConn),
		r:     bufio.NewReader(conn),
		w:     bufio.NewWriter(conn),

		cbs: option.PubSubHandlers,
	}

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

	for i, r := range p.DoMulti(cmds.NewMultiCompleted(init)...) {
		if i == 0 {
			p.info, err = r.ToMap()
		} else {
			err = r.Error()
		}
		if err != nil {
			return nil, err
		}
	}
	return p, nil
}

func (p *pipe) background() {
	atomic.CompareAndSwapInt32(&p.state, 0, 1)
	p.once.Do(func() { go p._background() })
}

func (p *pipe) _background() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	exit := func() {
		// stop accepting new requests
		atomic.CompareAndSwapInt32(&p.state, 1, 2)
		_ = p.conn.Close() // force both read & write goroutine to exit
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

		for atomic.LoadInt32(&p.state) != 3 {
			if ones[0], multi, ch = p.queue.NextWriteCmd(); multi == nil {
				if !ones[0].IsEmpty() {
					multi = ones
				} else {
					if p.w.Buffered() == 0 {
						runtime.Gosched()
						err = p.Error() // check err later
					} else {
						err = p.w.Flush()
					}
				}
			}
			for _, cmd := range multi {
				if err = proto.WriteCmd(p.w, cmd.Commands()); cmd.NoReply() {
					ch <- proto.NewErrResult(err)
				}
			}
			if err != nil && err != ErrConnClosing {
				p.error.CompareAndSwap(nil, &errs{error: err})
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
			if msg, err = proto.ReadNextMessage(p.r); err != nil {
				p.error.CompareAndSwap(nil, &errs{error: err})
				return
			}
			if msg.Type == '>' {
				p.handlePush(msg.Values)
				continue
			}
			// if unfulfilled multi commands are lead by opt-in and get success response
			if ff != len(multi) && len(multi) == 3 && multi[0].IsOptIn() {
				if ff == 1 {
					tmp = msg
				} else if ff == 2 {
					cacheable := cmds.Cacheable(multi[ff-1])
					ck, cc := cacheable.CacheKey()
					p.cache.Update(ck, cc, tmp, msg.Integer)
					tmp = proto.Message{}
				}
			}
		nextCMD:
			if ff == len(multi) {
				ff = 0
				ones[0], multi, ch = p.queue.NextResultCh() // ch should not be nil, otherwise it must be a protocol bug
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

	// clean up cache and free pending calls
	p.cache.FreeAndClose(proto.Message{Type: '-', String: ErrConnClosing.Error()})
	for atomic.LoadInt32(&p.waits) != 0 {
		p.queue.NextWriteCmd()
		if ones[0], multi, ch = p.queue.NextResultCh(); ch == nil {
			runtime.Gosched()
			continue
		}
		if multi == nil {
			multi = ones
		}
		for i := 0; i < len(multi); i++ {
			ch <- proto.NewErrResult(p.Error())
		}
	}
	atomic.CompareAndSwapInt32(&p.state, 2, 3)
}

func (p *pipe) handlePush(values []proto.Message) {
	if len(values) < 2 {
		return
	}
	// TODO: handle other push data
	// tracking-redir-broken
	// server-cpu-usage
	switch values[0].String {
	case "invalidate":
		p.cache.Delete(values[1].Values)
	case "message":
		if p.cbs.OnMessage != nil {
			p.cbs.OnMessage(values[1].String, values[2].String)
		}
	case "pmessage":
		if p.cbs.OnPMessage != nil {
			p.cbs.OnPMessage(values[1].String, values[2].String, values[3].String)
		}
	case "subscribe", "psubscribe":
		if p.cbs.OnSubscribed != nil {
			p.cbs.OnSubscribed(values[1].String, values[2].Integer)
		}
	case "unsubscribe", "punsubscribe":
		if p.cbs.OnUnSubscribed != nil {
			p.cbs.OnUnSubscribed(values[1].String, values[2].Integer)
		}
	}
}

func (p *pipe) Info() map[string]proto.Message {
	return p.info
}

func (p *pipe) Do(cmd cmds.Completed) (resp proto.Result) {
	waits := atomic.AddInt32(&p.waits, 1) // if this is 1, and background worker is not started, no need to queue
	state := atomic.LoadInt32(&p.state)

	if state == 1 {
		goto queue
	}

	if state == 0 {
		if waits != 1 {
			goto queue
		}
		if cmd.NoReply() {
			p.background()
			goto queue
		}
		resp = p.syncDo(cmd)
	} else {
		resp = proto.NewErrResult(p.Error())
	}
	if left := atomic.AddInt32(&p.waits, -1); state == 0 && waits == 1 && left != 0 {
		p.background()
	}
	return resp

queue:
	resp = <-p.queue.PutOne(cmd)
	atomic.AddInt32(&p.waits, -1)
	return resp
}

func (p *pipe) DoMulti(multi ...cmds.Completed) []proto.Result {
	waits := atomic.AddInt32(&p.waits, 1) // if this is 1, and background worker is not started, no need to queue
	state := atomic.LoadInt32(&p.state)
	resp := make([]proto.Result, len(multi))

	if state == 1 {
		goto queue
	}

	if state == 0 {
		if waits != 1 {
			goto queue
		}
		for _, cmd := range multi {
			if cmd.IsOptIn() || cmd.NoReply() {
				p.background()
				goto queue
			}
		}
		resp = p.syncDoMulti(resp, multi)
	} else {
		err := p.Error()
		for i := 0; i < len(resp); i++ {
			resp[i] = proto.NewErrResult(err)
		}
	}
	if left := atomic.AddInt32(&p.waits, -1); state == 0 && waits == 1 && left != 0 {
		p.background()
	}
	return resp

queue:
	ch := p.queue.PutMulti(multi)
	for i := range resp {
		resp[i] = <-ch
	}
	atomic.AddInt32(&p.waits, -1)
	return resp
}

func (p *pipe) syncDo(cmd cmds.Completed) (resp proto.Result) {
	var msg proto.Message
	err := proto.WriteCmd(p.w, cmd.Commands())
	if err == nil {
		if err = p.w.Flush(); err == nil {
			msg, err = proto.ReadNextMessage(p.r)
		}
	}
	if err != nil {
		p.error.CompareAndSwap(nil, &errs{error: err})
		p.background()                             // start the background worker
		atomic.CompareAndSwapInt32(&p.state, 1, 3) // stopping the worker and let it do the cleaning
	}
	return proto.NewResult(msg, err)
}

func (p *pipe) syncDoMulti(resp []proto.Result, multi []cmds.Completed) []proto.Result {
	var err error
	var msg proto.Message

	for _, cmd := range multi {
		err = proto.WriteCmd(p.w, cmd.Commands())
	}
	if err = p.w.Flush(); err != nil {
		goto abort
	}
	for i := 0; i < len(resp); i++ {
		if msg, err = proto.ReadNextMessage(p.r); err != nil {
			goto abort
		}
		resp[i] = proto.NewResult(msg, err)
	}
	return resp
abort:
	p.error.CompareAndSwap(nil, &errs{error: err})
	p.background()                             // start the background worker
	atomic.CompareAndSwapInt32(&p.state, 1, 3) // stopping the worker and let it do the cleaning
	for i := 0; i < len(resp); i++ {
		resp[i] = proto.NewErrResult(err)
	}
	return resp
}

func (p *pipe) DoCache(cmd cmds.Cacheable, ttl time.Duration) proto.Result {
	ck, cc := cmd.CacheKey()
	if v, entry := p.cache.GetOrPrepare(ck, cc, ttl); v.Type != 0 {
		return proto.NewResult(v, nil)
	} else if entry != nil {
		return proto.NewResult(entry.Wait(), nil)
	}
	return p.DoMulti(cmds.OptInCmd, cmds.Completed(cmd), cmds.NewCompleted([]string{"PTTL", ck}))[1]
}

func (p *pipe) Error() error {
	if err, ok := p.error.Load().(*errs); ok {
		return err.error
	}
	return nil
}

func (p *pipe) Close() {
	swapped := p.error.CompareAndSwap(nil, &errs{error: ErrConnClosing})
	atomic.CompareAndSwapInt32(&p.state, 0, 2)
	atomic.CompareAndSwapInt32(&p.state, 1, 2)
	for atomic.LoadInt32(&p.waits) != 0 {
		runtime.Gosched()
	}
	if swapped {
		p.background()
		<-p.queue.PutOne(cmds.QuitCmd)
	}
	atomic.CompareAndSwapInt32(&p.state, 2, 3)
}
