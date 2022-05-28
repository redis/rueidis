package rueidis

import (
	"context"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/rueian/rueidis/internal/cmds"
)

type connFn func(dst string, opt *ClientOption) conn
type dialFn func(dst string, opt *ClientOption) (net.Conn, error)
type wireFn func() wire

type singleconnect struct {
	w wire
	e error
	g sync.WaitGroup
}

type conn interface {
	Do(ctx context.Context, cmd cmds.Completed) RedisResult
	DoCache(ctx context.Context, cmd cmds.Cacheable, ttl time.Duration) RedisResult
	DoMulti(ctx context.Context, multi ...cmds.Completed) []RedisResult
	Receive(ctx context.Context, subscribe cmds.Completed, fn func(message PubSubMessage)) error
	Info() map[string]RedisMessage
	Error() error
	Close()
	Dial() error
	Override(conn)
	Acquire() wire
	Store(w wire)
	Is(addr string) bool
}

var _ conn = (*mux)(nil)

type mux struct {
	init   wire
	dead   wire
	wire   atomic.Value
	sc     *singleconnect
	pool   *pool
	wireFn wireFn
	dst    string
	mu     sync.Mutex
}

func makeMux(dst string, option *ClientOption, dialFn dialFn) *mux {
	dead := deadFn()
	return newMux(dst, option, (*pipe)(nil), dead, func() (w wire) {
		conn, err := dialFn(dst, option)
		if err == nil {
			w, err = newPipe(conn, option)
		}
		if err != nil {
			dead.error.Store(&errs{error: err})
			w = dead
		}
		return w
	})
}

func newMux(dst string, option *ClientOption, init, dead wire, wireFn wireFn) *mux {
	m := &mux{dst: dst, init: init, dead: dead, wireFn: wireFn}
	m.wire.Store(init)
	m.pool = newPool(option.BlockingPoolSize, dead, wireFn)
	return m
}

func (m *mux) Override(cc conn) {
	if m2, ok := cc.(*mux); ok {
		m.wire.CompareAndSwap(m.init, m2.wire.Load())
	}
}

func (m *mux) pipe() wire {
	w, _ := m._pipe()
	return w
}

func (m *mux) _pipe() (w wire, err error) {
	if w = m.wire.Load().(wire); w != m.init {
		return w, nil
	}

	m.mu.Lock()
	sc := m.sc
	if m.sc == nil {
		m.sc = &singleconnect{}
		m.sc.g.Add(1)
	}
	m.mu.Unlock()

	if sc != nil {
		sc.g.Wait()
		return sc.w, sc.e
	}

	if w = m.wire.Load().(wire); w == m.init {
		if w = m.wireFn(); w != m.dead {
			m.wire.Store(w)
		} else {
			err = w.Error()
		}
	}

	m.mu.Lock()
	sc = m.sc
	m.sc = nil
	m.mu.Unlock()

	sc.w = w
	sc.e = err
	sc.g.Done()

	return w, err
}

func (m *mux) Dial() error { // no retry
	_, err := m._pipe()
	return err
}

func (m *mux) Info() map[string]RedisMessage {
	return m.pipe().Info()
}

func (m *mux) Error() error {
	return m.pipe().Error()
}

func (m *mux) Do(ctx context.Context, cmd cmds.Completed) (resp RedisResult) {
	if cmd.IsBlock() {
		resp = m.blocking(ctx, cmd)
	} else {
		resp = m.pipeline(ctx, cmd)
	}
	return resp
}

func (m *mux) DoMulti(ctx context.Context, multi ...cmds.Completed) (resp []RedisResult) {
	for _, cmd := range multi {
		if cmd.IsBlock() {
			goto block
		}
	}
	return m.pipelineMulti(ctx, multi)
block:
	return m.blockingMulti(ctx, multi)
}

func (m *mux) blocking(ctx context.Context, cmd cmds.Completed) (resp RedisResult) {
	wire := m.pool.Acquire()
	resp = wire.Do(ctx, cmd)
	m.pool.Store(wire)
	return resp
}

func (m *mux) blockingMulti(ctx context.Context, cmd []cmds.Completed) (resp []RedisResult) {
	wire := m.pool.Acquire()
	resp = wire.DoMulti(ctx, cmd...)
	m.pool.Store(wire)
	return resp
}

func (m *mux) pipeline(ctx context.Context, cmd cmds.Completed) (resp RedisResult) {
	wire := m.pipe()
	if resp = wire.Do(ctx, cmd); isBroken(resp.NonRedisError(), wire) {
		m.wire.CompareAndSwap(wire, m.init)
	}
	return resp
}

func (m *mux) pipelineMulti(ctx context.Context, cmd []cmds.Completed) (resp []RedisResult) {
	wire := m.pipe()
	resp = wire.DoMulti(ctx, cmd...)
	for _, r := range resp {
		if isBroken(r.NonRedisError(), wire) {
			m.wire.CompareAndSwap(wire, m.init)
			return resp
		}
	}
	return resp
}

func (m *mux) DoCache(ctx context.Context, cmd cmds.Cacheable, ttl time.Duration) RedisResult {
	wire := m.pipe()
	resp := wire.DoCache(ctx, cmd, ttl)
	if isBroken(resp.NonRedisError(), wire) {
		m.wire.CompareAndSwap(wire, m.init)
	}
	return resp
}

func (m *mux) Receive(ctx context.Context, subscribe cmds.Completed, fn func(message PubSubMessage)) error {
	wire := m.pipe()
	err := wire.Receive(ctx, subscribe, fn)
	if isBroken(err, wire) {
		m.wire.CompareAndSwap(wire, m.init)
	}
	return err
}

func (m *mux) Acquire() wire {
	return m.pool.Acquire()
}

func (m *mux) Store(w wire) {
	w.SetPubSubHooks(PubSubHooks{})
	if w.Pipelining() {
		w.DoMulti(context.Background(), cmds.UnsubscribeCmd, cmds.PUnsubscribeCmd, cmds.SUnsubscribeCmd)
	}
	m.pool.Store(w)
}

func (m *mux) Close() {
	if prev := m.wire.Swap(m.dead).(wire); prev != m.init && prev != m.dead {
		prev.Close()
	}
	m.pool.Close()
}

func (m *mux) Is(addr string) bool {
	return m.dst == addr
}

func isBroken(err error, w wire) bool {
	return err != nil && err != ErrClosing && w.Error() != nil
}
