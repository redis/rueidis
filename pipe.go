package rueidis

import (
	"bufio"
	"context"
	"errors"
	"net"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/rueian/rueidis/internal/cmds"
)

type wire interface {
	Do(ctx context.Context, cmd cmds.Completed) RedisResult
	DoCache(ctx context.Context, cmd cmds.Cacheable, ttl time.Duration) RedisResult
	DoMulti(ctx context.Context, multi ...cmds.Completed) []RedisResult
	Receive(ctx context.Context, subscribe cmds.Completed, fn func(message PubSubMessage)) error
	Info() map[string]RedisMessage
	Error() error
	Close()
}

var _ wire = (*pipe)(nil)

type pipe struct {
	waits   int32
	state   int32
	slept   int32
	version int32

	once  sync.Once
	cond  sync.Cond
	conn  net.Conn
	queue queue
	cache cache
	error atomic.Value

	r *bufio.Reader
	w *bufio.Writer

	info  map[string]RedisMessage
	subs  *subs
	psubs *subs
}

func newPipe(conn net.Conn, option *ClientOption) (p *pipe, err error) {
	if option.CacheSizeEachConn <= 0 {
		option.CacheSizeEachConn = DefaultCacheBytes
	}

	if option.ConnWriteTimeout > 0 {
		conn = &writeTimeoutConn{Conn: conn, timeout: option.ConnWriteTimeout}
	}

	p = &pipe{
		conn:  conn,
		cond:  sync.Cond{L: &sync.Mutex{}},
		queue: newRing(),
		cache: newLRU(option.CacheSizeEachConn),
		r:     bufio.NewReader(conn),
		w:     bufio.NewWriter(conn),

		subs:  newSubs(),
		psubs: newSubs(),
	}

	helloCmd := []string{"HELLO", "3"}
	if option.Password != "" && option.Username == "" {
		option.Username = "default"
	}
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

	if ver, ok := p.info["version"]; ok {
		if v := strings.Split(ver.string, "."); len(v) != 0 {
			vv, _ := strconv.ParseInt(v[0], 10, 32)
			p.version = int32(vv)
		}
	}

	return p, nil
}

func (p *pipe) _sleep() (slept bool) {
	if slept = atomic.LoadInt32(&p.waits) == 0 && atomic.LoadInt32(&p.state) == 1; slept {
		p.cond.L.Lock()
		if slept = atomic.LoadInt32(&p.waits) == 0 && atomic.LoadInt32(&p.state) == 1; slept {
			p.slept = 1
			p.cond.Wait()
			p.slept = 0
		}
		p.cond.L.Unlock()
	}
	return slept
}

func (p *pipe) _awake() {
	p.cond.L.Lock()
	slept := p.slept
	p.cond.L.Unlock()
	if slept == 1 {
		p.cond.Broadcast()
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

	p.subs.Close()
	p.psubs.Close()

	var (
		ones  = make([]cmds.Completed, 1)
		multi []cmds.Completed
		ch    chan RedisResult
		cond  *sync.Cond
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
		if ones[0], multi, ch, cond = p.queue.NextResultCh(); ch != nil {
			if multi == nil {
				multi = ones
			}
			for _, one := range multi {
				if !one.NoReply() {
					ch <- newErrResult(p.Error())
				}
			}
			cond.L.Unlock()
			cond.Signal()
		} else {
			cond.L.Unlock()
			cond.Signal()
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
		cond  *sync.Cond
		ones  = make([]cmds.Completed, 1)
		multi []cmds.Completed
		ch    chan RedisResult
		ff    int // fulfilled count
		ver   = p.version
	)

	for {
		if msg, err = readNextMessage(p.r); err != nil {
			p.error.CompareAndSwap(nil, &errs{error: err})
			return
		}
		if msg.typ == '>' {
			p.handlePush(msg.values)
			continue
		} else if ver < 7 && len(msg.values) != 0 {
			// This is a workaround for Redis 6's broken invalidation protocol: https://github.com/redis/redis/issues/8935
			// When Redis 6 handles MULTI, MGET, or other multi-keys command,
			// it will send invalidation message immediately if it finds the keys are expired, thus causing the multi-keys command response to be broken.
			// We fix this by fetching the next message and patch it back to the response.
			i := 0
			for j, v := range msg.values {
				if v.typ == '>' {
					p.handlePush(v.values)
				} else {
					if i != j {
						msg.values[i] = v
					}
					i++
				}
			}
			for ; i < len(msg.values); i++ {
				if msg.values[i], err = readNextMessage(p.r); err != nil {
					p.error.CompareAndSwap(nil, &errs{error: err})
					return
				}
			}
		}
		// if unfulfilled multi commands are lead by opt-in and get success response
		if ff == 4 && len(multi) == 5 && multi[0].IsOptIn() {
			cacheable := cmds.Cacheable(multi[3])
			ck, cc := cacheable.CacheKey()
			if len(msg.values) == 2 {
				cp := msg.values[1]
				cp.attrs = cacheMark
				p.cache.Update(ck, cc, cp, msg.values[0].integer)
			}
		}
	nextCMD:
		if ff == len(multi) {
			ff = 0
			ones[0], multi, ch, cond = p.queue.NextResultCh() // ch should not be nil, otherwise it must be a protocol bug
			if ch == nil {
				panic(protocolbug)
			}
			if multi == nil {
				multi = ones
			}
		}
		if multi[ff].NoReply() {
			ff++
			if ff == len(multi) {
				cond.L.Unlock()
				cond.Signal()
			}
			goto nextCMD
		} else {
			ff++
			ch <- newResult(msg, err)
			if ff == len(multi) {
				cond.L.Unlock()
				cond.Signal()
			}
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
	// subscribe
	// psubscribe
	switch values[0].string {
	case "invalidate":
		if values[1].IsNil() {
			p.cache.Delete(nil)
		} else {
			p.cache.Delete(values[1].values)
		}
	case "message":
		if len(values) >= 3 {
			p.subs.Publish(values[1].string, PubSubMessage{Channel: values[1].string, Message: values[2].string})
		}
	case "pmessage":
		if len(values) >= 4 {
			p.psubs.Publish(values[1].string, PubSubMessage{Pattern: values[1].string, Channel: values[2].string, Message: values[3].string})
		}
	case "unsubscribe":
		p.subs.Unsubscribe(values[1].string)
	case "punsubscribe":
		p.psubs.Unsubscribe(values[1].string)
	}
}

func (p *pipe) Receive(ctx context.Context, subscribe cmds.Completed, fn func(message PubSubMessage)) error {
	if p.subs == nil || p.psubs == nil {
		return ErrClosing
	}

	var sb *subs
	cmd, args := subscribe.Commands()[0], subscribe.Commands()[1:]

	switch cmd {
	case "SUBSCRIBE", "SSUBSCRIBE":
		sb = p.subs
	case "PSUBSCRIBE":
		sb = p.psubs
	default:
		panic(wrongreceive)
	}

	if id, ch := sb.Subscribe(args); ch != nil {
		defer sb.Remove(id)

		if err := p.Do(ctx, subscribe).Error(); err != nil {
			return err
		}
		if ctxCh := ctx.Done(); ctxCh == nil {
			for msg := range ch {
				fn(msg)
			}
		} else {
		next:
			select {
			case msg, ok := <-ch:
				if ok {
					fn(msg)
					goto next
				}
			case <-ctx.Done():
				go func() {
					for range ch {
					}
				}()
				return ctx.Err()
			}
		}
	}
	return p.Error()
}

func (p *pipe) Info() map[string]RedisMessage {
	return p.info
}

func (p *pipe) Do(ctx context.Context, cmd cmds.Completed) (resp RedisResult) {
	if err := ctx.Err(); err != nil {
		return newErrResult(ctx.Err())
	}

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
		resp = p.syncDo(ctx, cmd)
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
	if ctxCh := ctx.Done(); ctxCh == nil {
		resp = <-ch
		atomic.AddInt32(&p.waits, -1)
	} else {
		select {
		case resp = <-ch:
			atomic.AddInt32(&p.waits, -1)
		case <-ctxCh:
			resp = newErrResult(ctx.Err())
			go func() {
				<-ch
				atomic.AddInt32(&p.waits, -1)
			}()
		}
	}
	return resp
}

func (p *pipe) DoMulti(ctx context.Context, multi ...cmds.Completed) []RedisResult {
	resp := make([]RedisResult, len(multi))
	if err := ctx.Err(); err != nil {
		for i := 0; i < len(resp); i++ {
			resp[i] = newErrResult(err)
		}
		return resp
	}

	isOptIn := multi[0].IsOptIn() // len(multi) > 0 should have already been checked by upper layer
	noReply := multi[0].NoReply()

	for _, cmd := range multi[1:] {
		if noReply != cmd.NoReply() {
			panic(prohibitmix)
		}
	}

	waits := atomic.AddInt32(&p.waits, 1) // if this is 1, and background worker is not started, no need to queue
	state := atomic.LoadInt32(&p.state)

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
		resp = p.syncDoMulti(ctx, resp, multi)
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
	if ctxCh := ctx.Done(); ctxCh == nil {
		for ; i < len(resp); i++ {
			resp[i] = <-ch
		}
	} else {
		for ; i < len(resp); i++ {
			select {
			case resp[i] = <-ch:
			case <-ctxCh:
				goto abort
			}
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
	err := newErrResult(ctx.Err())
	for ; i < len(resp); i++ {
		resp[i] = err
	}
	return resp
}

func (p *pipe) syncDo(ctx context.Context, cmd cmds.Completed) (resp RedisResult) {
	if dl, ok := ctx.Deadline(); ok {
		p.conn.SetDeadline(dl)
		defer p.conn.SetDeadline(time.Time{})
	}

	var msg RedisMessage
	err := writeCmd(p.w, cmd.Commands())
	if err == nil {
		if err = p.w.Flush(); err == nil {
			msg, err = syncRead(p.r)
		}
	}
	if err != nil {
		if errors.Is(err, os.ErrDeadlineExceeded) {
			err = context.DeadlineExceeded
		}
		p.error.CompareAndSwap(nil, &errs{error: err})
		atomic.CompareAndSwapInt32(&p.state, 1, 3) // stopping the worker and let it do the cleaning
		p.background()                             // start the background worker
	}
	return newResult(msg, err)
}

func (p *pipe) syncDoMulti(ctx context.Context, resp []RedisResult, multi []cmds.Completed) []RedisResult {
	if dl, ok := ctx.Deadline(); ok {
		p.conn.SetDeadline(dl)
		defer p.conn.SetDeadline(time.Time{})
	}

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
	if errors.Is(err, os.ErrDeadlineExceeded) {
		err = context.DeadlineExceeded
	}
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
	if p.cache == nil {
		return newErrResult(ErrClosing)
	}
	ck, cc := cmd.CacheKey()
	if v, entry := p.cache.GetOrPrepare(ck, cc, ttl); v.typ != 0 {
		return newResult(v, nil)
	} else if entry != nil {
		return newResult(entry.Wait())
	}
	resp := p.DoMulti(
		ctx,
		cmds.OptInCmd,
		cmds.MultiCmd,
		cmds.NewCompleted([]string{"PTTL", ck}),
		cmds.Completed(cmd),
		cmds.ExecCmd,
	)
	exec, err := resp[4].ToArray()
	if err != nil {
		if _, ok := err.(*RedisError); !ok {
			p.cache.Cancel(ck, cc, RedisMessage{}, err)
			return newErrResult(err)
		}
		// EXEC aborted, return err of the input cmd in MULTI block
		if resp[3].val.typ != '+' {
			p.cache.Cancel(ck, cc, resp[3].val, nil)
			return newResult(resp[3].val, nil)
		}
		p.cache.Cancel(ck, cc, resp[4].val, nil)
		return newResult(resp[4].val, nil)
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
	p.error.CompareAndSwap(nil, errClosing)
	atomic.CompareAndSwapInt32(&p.state, 0, 2)
	atomic.CompareAndSwapInt32(&p.state, 1, 2)
	atomic.AddInt32(&p.waits, 1)
	if p.queue != nil {
		p.background()
		p._awake()
		for atomic.LoadInt32(&p.waits) != 1 {
			runtime.Gosched()
		}
		<-p.queue.PutOne(cmds.QuitCmd)
	}
	atomic.AddInt32(&p.waits, -1)
	atomic.CompareAndSwapInt32(&p.state, 2, 3)
}

var dead *pipe

func init() {
	dead = &pipe{state: 3}
	dead.error.Store(errClosing)
}

const (
	protocolbug  = "protocol bug, message handled out of order"
	prohibitmix  = "mixing SUBSCRIBE, PSUBSCRIBE, UNSUBSCRIBE, PUNSUBSCRIBE with other commands in DoMulti is prohibited"
	wrongreceive = `only SUBSCRIBE, SSUBSCRIBE, or PSUBSCRIBE command are allowed in Receive`
)

var cacheMark = &(RedisMessage{})
var errClosing = &errs{error: ErrClosing}

type errs struct{ error }

type writeTimeoutConn struct {
	net.Conn
	current time.Time
	timeout time.Duration
}

// automatically apply the write-deadline in Write call only when necessary,
// since the net.Conn is used behind the bufio.Writer
func (c *writeTimeoutConn) Write(b []byte) (n int, err error) {
	if c.current.IsZero() {
		_ = c.Conn.SetWriteDeadline(time.Now().Add(c.timeout))
		n, err = c.Conn.Write(b)
		_ = c.Conn.SetWriteDeadline(time.Time{})
		return
	}
	return c.Conn.Write(b)
}

func (c *writeTimeoutConn) SetDeadline(t time.Time) error {
	c.current = t
	return c.Conn.SetDeadline(t)
}
