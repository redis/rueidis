package rueidis

import (
	"bufio"
	"context"
	"net"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/rueian/rueidis/internal/cmds"
)

type wire interface {
	Do(ctx context.Context, cmd cmds.Completed) RedisResult
	DoCache(ctx context.Context, cmd cmds.Cacheable, ttl time.Duration) RedisResult
	DoMulti(ctx context.Context, multi ...cmds.Completed) []RedisResult
	Info() map[string]RedisMessage
	Error() error
	Close()
}

var _ wire = (*pipe)(nil)

type pipe struct {
	waits int32
	state int32
	sleep int32

	once  sync.Once
	cond  sync.Cond
	conn  net.Conn
	queue queue
	cache cache
	error atomic.Value

	r *bufio.Reader
	w *bufio.Writer

	info map[string]RedisMessage

	cbs PubSubOption

	doneFn func(err error)
}

func newPipe(conn net.Conn, option *ClientOption, onDisconnected func(err error)) (p *pipe, err error) {
	if option.CacheSizeEachConn <= 0 {
		option.CacheSizeEachConn = DefaultCacheBytes
	}

	p = &pipe{
		conn:  conn,
		cond:  sync.Cond{L: noLock{}},
		queue: newRing(),
		cache: newLRU(option.CacheSizeEachConn),
		r:     bufio.NewReader(conn),
		w:     bufio.NewWriter(conn),

		cbs:    option.PubSubOption,
		doneFn: onDisconnected,
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

	for i, r := range p.DoMulti(context.Background(), cmds.NewMultiCompleted(init)...) {
		if i == 0 {
			p.info, err = r.ToMap()
		} else {
			err = r.Error()
		}
		if err != nil {
			p.Close()
			return nil, err
		}
	}
	return p, nil
}

func (p *pipe) _sleep() (slept bool) {
	atomic.AddInt32(&p.sleep, 1) // create barrier
	if slept = atomic.LoadInt32(&p.waits) == 0 && atomic.LoadInt32(&p.state) == 1; slept {
		p.cond.Wait()
	}
	atomic.AddInt32(&p.sleep, -1)
	return slept
}

func (p *pipe) _awake() {
	for atomic.LoadInt32(&p.sleep) != 0 {
		p.cond.Broadcast()
		runtime.Gosched()
	}
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
	go func() {
		p._backgroundWrite()
		exit()
	}()
	go func() {
		p._backgroundRead()
		exit()
		p._awake()
	}()
	wg.Wait()

	if p.doneFn != nil {
		go p.doneFn(p.Error())
	}

	var (
		ones  = make([]cmds.Completed, 1)
		multi []cmds.Completed
		ch    chan RedisResult
	)

	// clean up cache and free pending calls
	p.cache.FreeAndClose(RedisMessage{typ: '-', string: ErrClosing.Error()})
	for atomic.LoadInt32(&p.waits) != 0 {
		if ones[0], multi, ch = p.queue.NextWriteCmd(); ch != nil {
			if multi == nil {
				multi = ones
			}
			for _, one := range multi {
				if one.NoReply() {
					ch <- newErrResult(p.Error())
				}
			}
		}
		if ones[0], multi, ch = p.queue.NextResultCh(); ch != nil {
			if multi == nil {
				multi = ones
			}
			for _, one := range multi {
				if !one.NoReply() {
					ch <- newErrResult(p.Error())
				}
			}
		} else {
			runtime.Gosched()
		}
	}
	atomic.CompareAndSwapInt32(&p.state, 2, 3)
}

func (p *pipe) _backgroundWrite() {
	var (
		err   error
		ones  = make([]cmds.Completed, 1)
		multi []cmds.Completed
		ch    chan RedisResult
	)

	for atomic.LoadInt32(&p.state) != 3 {
		if ones[0], multi, ch = p.queue.NextWriteCmd(); ch == nil {
			if p.w.Buffered() == 0 {
				err = p.Error()
			} else {
				err = p.w.Flush()
			}
		} else if multi == nil {
			multi = ones
		}
		for _, cmd := range multi {
			if err = writeCmd(p.w, cmd.Commands()); cmd.NoReply() {
				err = p.w.Flush()
				ch <- newErrResult(err)
			}
		}
		if err != nil {
			if err != ErrClosing { // ignore ErrClosing to allow final QUIT command to be sent
				p.error.CompareAndSwap(nil, &errs{error: err})
				return
			}
			runtime.Gosched()
		} else if ch == nil && !p._sleep() {
			runtime.Gosched()
		}
	}
}

func (p *pipe) _backgroundRead() {
	var (
		err   error
		msg   RedisMessage
		ones  = make([]cmds.Completed, 1)
		multi []cmds.Completed
		ch    chan RedisResult
		ff    int // fulfilled count
	)

	for {
		if msg, err = readNextMessage(p.r); err != nil {
			p.error.CompareAndSwap(nil, &errs{error: err})
			return
		}
		if msg.typ == '>' {
			p.handlePush(msg.values)
			continue
		}
		// if unfulfilled multi commands are lead by opt-in and get success response
		if ff != len(multi) && len(multi) == 5 && multi[0].IsOptIn() {
			if ff == 4 {
				cacheable := cmds.Cacheable(multi[3])
				ck, cc := cacheable.CacheKey()
				if len(msg.values) != 2 { // EXEC aborted
					p.cache.Update(ck, cc, msg, 0)
				} else {
					cp := msg.values[1]
					cp.attrs = cacheMark
					p.cache.Update(ck, cc, cp, msg.values[0].integer)
				}
			}
		}
	nextCMD:
		if ff == len(multi) {
			ff = 0
			ones[0], multi, ch = p.queue.NextResultCh() // ch should not be nil, otherwise it must be a protocol bug
			if ch == nil {
				panic(protocolbug)
			}
		}
		if multi == nil {
			multi = ones
		}
		if multi[ff].NoReply() {
			ff++
			goto nextCMD
		} else {
			ff++
			ch <- newResult(msg, err)
		}
	}
}

func (p *pipe) handlePush(values []RedisMessage) {
	if len(values) < 2 {
		return
	}
	// TODO: handle other push data
	// tracking-redir-broken
	// server-cpu-usage
	switch values[0].string {
	case "invalidate":
		if values[1].IsNil() {
			p.cache.Delete(nil)
		} else {
			p.cache.Delete(values[1].values)
		}
	case "message":
		if p.cbs.onMessage != nil {
			p.cbs.onMessage(values[1].string, values[2].string)
		}
	case "pmessage":
		if p.cbs.onPMessage != nil {
			p.cbs.onPMessage(values[1].string, values[2].string, values[3].string)
		}
	case "subscribe", "psubscribe":
		if p.cbs.onSubscribed != nil {
			p.cbs.onSubscribed(values[1].string, values[2].integer)
		}
	case "unsubscribe", "punsubscribe":
		if p.cbs.onUnSubscribed != nil {
			p.cbs.onUnSubscribed(values[1].string, values[2].integer)
		}
	}
}

func (p *pipe) Info() map[string]RedisMessage {
	return p.info
}

func (p *pipe) Do(ctx context.Context, cmd cmds.Completed) (resp RedisResult) {
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
		resp = newErrResult(p.Error())
	}
	if left := atomic.AddInt32(&p.waits, -1); state == 0 && waits == 1 && left != 0 {
		p.background()
	}
	return resp

queue:
	ch := p.queue.PutOne(cmd)
	if waits == 1 {
		p._awake()
	}
	select {
	case <-ctx.Done():
		resp = newErrResult(ctx.Err())
		go func() {
			<-ch
			atomic.AddInt32(&p.waits, -1)
		}()
	case resp = <-ch:
		atomic.AddInt32(&p.waits, -1)
	}
	return resp
}

func (p *pipe) DoMulti(ctx context.Context, multi ...cmds.Completed) []RedisResult {
	isOptIn := multi[0].IsOptIn() // len(multi) > 0 should have already been checked by upper layer
	noReply := multi[0].NoReply()

	for _, cmd := range multi[1:] {
		if noReply != cmd.NoReply() {
			panic(prohibitmix)
		}
	}

	waits := atomic.AddInt32(&p.waits, 1) // if this is 1, and background worker is not started, no need to queue
	state := atomic.LoadInt32(&p.state)
	resp := make([]RedisResult, len(multi))

	if state == 1 {
		goto queue
	}

	if state == 0 {
		if waits != 1 {
			goto queue
		}
		if isOptIn || noReply {
			p.background()
			goto queue
		}
		resp = p.syncDoMulti(resp, multi)
	} else {
		err := newErrResult(p.Error())
		for i := 0; i < len(resp); i++ {
			resp[i] = err
		}
	}
	if left := atomic.AddInt32(&p.waits, -1); state == 0 && waits == 1 && left != 0 {
		p.background()
	}
	return resp

queue:
	ch := p.queue.PutMulti(multi)
	if waits == 1 {
		p._awake()
	}
	var i int
	for ; i < len(resp); i++ {
		select {
		case <-ctx.Done():
			goto abort
		case resp[i] = <-ch:
		}
	}
	atomic.AddInt32(&p.waits, -1)
	return resp
abort:
	go func(i int) {
		for ; i < len(resp); i++ {
			<-ch
		}
		atomic.AddInt32(&p.waits, -1)
	}(i)
	for err := newErrResult(ctx.Err()); i < len(resp); i++ {
		resp[i] = err
	}
	return resp
}

func (p *pipe) syncDo(cmd cmds.Completed) (resp RedisResult) {
	var msg RedisMessage
	err := writeCmd(p.w, cmd.Commands())
	if err == nil {
		if err = p.w.Flush(); err == nil {
			msg, err = syncRead(p.r)
		}
	}
	if err != nil {
		p.error.CompareAndSwap(nil, &errs{error: err})
		atomic.CompareAndSwapInt32(&p.state, 1, 3) // stopping the worker and let it do the cleaning
		p.background()                             // start the background worker
	}
	return newResult(msg, err)
}

func (p *pipe) syncDoMulti(resp []RedisResult, multi []cmds.Completed) []RedisResult {
	var err error
	var msg RedisMessage

	for _, cmd := range multi {
		_ = writeCmd(p.w, cmd.Commands())
	}
	if err = p.w.Flush(); err != nil {
		goto abort
	}
	for i := 0; i < len(resp); i++ {
		if msg, err = syncRead(p.r); err != nil {
			goto abort
		}
		resp[i] = newResult(msg, err)
	}
	return resp
abort:
	p.error.CompareAndSwap(nil, &errs{error: err})
	atomic.CompareAndSwapInt32(&p.state, 1, 3) // stopping the worker and let it do the cleaning
	p.background()                             // start the background worker
	for i := 0; i < len(resp); i++ {
		resp[i] = newErrResult(err)
	}
	return resp
}

func syncRead(r *bufio.Reader) (m RedisMessage, err error) {
next:
	if m, err = readNextMessage(r); err != nil {
		return m, err
	}
	if m.typ == '>' {
		goto next
	}
	return m, nil
}

func (p *pipe) DoCache(ctx context.Context, cmd cmds.Cacheable, ttl time.Duration) RedisResult {
	ck, cc := cmd.CacheKey()
	if v, entry := p.cache.GetOrPrepare(ck, cc, ttl); v.typ != 0 {
		return newResult(v, nil)
	} else if entry != nil {
		return newResult(entry.Wait(), nil)
	}
	exec, err := p.DoMulti(ctx, cmds.OptInCmd, cmds.MultiCmd, cmds.NewCompleted([]string{"PTTL", ck}), cmds.Completed(cmd), cmds.ExecCmd)[4].ToArray()
	if err != nil {
		return newErrResult(err)
	}
	return newResult(exec[1], nil)
}

func (p *pipe) Error() error {
	if err, ok := p.error.Load().(*errs); ok {
		return err.error
	}
	return nil
}

func (p *pipe) Close() {
	swapped := p.error.CompareAndSwap(nil, errClosing)
	atomic.CompareAndSwapInt32(&p.state, 0, 2)
	atomic.CompareAndSwapInt32(&p.state, 1, 2)
	p._awake()
	for atomic.LoadInt32(&p.waits) != 0 {
		runtime.Gosched()
	}
	if swapped {
		p.background()
		<-p.queue.PutOne(cmds.QuitCmd)
	}
	atomic.CompareAndSwapInt32(&p.state, 2, 3)
}

var dead *pipe

func init() {
	dead = &pipe{state: 3}
	dead.error.Store(errClosing)
}

const (
	protocolbug = "protocol bug, message handled out of order"
	prohibitmix = "mixing SUBSCRIBE, PSUBSCRIBE, UNSUBSCRIBE, PUNSUBSCRIBE with other commands in DoMulti is prohibited"
)

var cacheMark = &(RedisMessage{})
var errClosing = &errs{error: ErrClosing}

type errs struct{ error }

type noLock struct{}

func (n noLock) Lock() {}

func (n noLock) Unlock() {}
