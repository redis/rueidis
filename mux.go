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
type wireFn func(onDisconnected func(err error)) (wire, error)

type singleconnect struct {
	w wire
	e error
	g sync.WaitGroup
}

type conn interface {
	wire
	Dial() error
	Override(conn)
	Acquire() wire
	Store(w wire)
	OnDisconnected(func(err error))
}

var _ conn = (*mux)(nil)

type mux struct {
	init   wire
	dead   wire
	wire   atomic.Value
	doneFn atomic.Value
	sc     *singleconnect
	pool   *pool
	wireFn wireFn
	dst    string
	mu     sync.Mutex
}

func makeMux(dst string, option *ClientOption, dialFn dialFn, retryOnRefuse bool) *mux {
	return newMux(dst, option, (*pipe)(nil), dead, func(onDisconnected func(err error)) (w wire, err error) {
		conn, err := dialFn(dst, option)
		if err == nil {
			w, err = newPipe(conn, option, onDisconnected)
		} else if !retryOnRefuse {
			if e, ok := err.(net.Error); ok && !e.Timeout() && !e.Temporary() {
				return dead, err
			}
		}
		return w, err
	})
}

func newMux(dst string, option *ClientOption, init, dead wire, wireFn wireFn) *mux {
	m := &mux{dst: dst, init: init, dead: dead, wireFn: wireFn}
	m.wire.Store(init)
	m.pool = newPool(option.BlockingPoolSize, dead, m._newPooledWire)
	return m
}

func (m *mux) _newPooledWire() wire {
retry:
	if wire, err := m.wireFn(nil); err == nil || wire == m.dead {
		return wire
	}
	goto retry
}

func (m *mux) Override(cc conn) {
	if m2, ok := cc.(*mux); ok {
		m.wire.CompareAndSwap(m.init, m2.wire.Load())
	}
}

func (m *mux) pipe() wire {
retry:
	if wire, err := m._pipe(); err == nil || wire == m.dead {
		return wire
	}
	goto retry
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
		if w, err = m.wireFn(m.disconnected); err == nil || w == m.dead {
			m.wire.Store(w)
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

func (m *mux) disconnected(err error) {
	if fn := m.doneFn.Load(); fn != nil {
		fn.(func(err error))(err)
	}
}

func (m *mux) OnDisconnected(fn func(err error)) {
	m.doneFn.CompareAndSwap(nil, fn)
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
retry:
	if cmd.IsBlock() {
		resp = m.blocking(ctx, cmd)
	} else {
		resp = m.pipeline(ctx, cmd)
	}
	if cmd.IsReadOnly() && isNetworkErr(resp.NonRedisError()) {
		goto retry
	}
	return resp
}

func (m *mux) DoMulti(ctx context.Context, multi ...cmds.Completed) (resp []RedisResult) {
	var block, write bool
	for _, cmd := range multi {
		block = block || cmd.IsBlock()
		write = write || cmd.IsWrite()
	}
retry:
	if block {
		resp = m.blockingMulti(ctx, multi)
	} else {
		resp = m.pipelineMulti(ctx, multi)
	}
	if !write {
		for _, r := range resp {
			if isNetworkErr(r.NonRedisError()) {
				goto retry
			}
		}
	}
	return resp
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
	if resp = wire.Do(ctx, cmd); isNetworkErr(resp.NonRedisError()) {
		m.wire.CompareAndSwap(wire, m.init)
	}
	return resp
}

func (m *mux) pipelineMulti(ctx context.Context, cmd []cmds.Completed) (resp []RedisResult) {
	wire := m.pipe()
	resp = wire.DoMulti(ctx, cmd...)
	for _, r := range resp {
		if isNetworkErr(r.NonRedisError()) {
			m.wire.CompareAndSwap(wire, m.init)
			return resp
		}
	}
	return resp
}

func (m *mux) DoCache(ctx context.Context, cmd cmds.Cacheable, ttl time.Duration) RedisResult {
retry:
	wire := m.pipe()
	resp := wire.DoCache(ctx, cmd, ttl)
	if isNetworkErr(resp.NonRedisError()) {
		m.wire.CompareAndSwap(wire, m.init)
		goto retry
	}
	return resp
}

func (m *mux) Acquire() wire {
	return m.pool.Acquire()
}

func (m *mux) Store(w wire) {
	m.pool.Store(w)
}

func (m *mux) Close() {
	if prev := m.wire.Swap(m.dead).(wire); prev != m.init {
		prev.Close()
	}
	m.pool.Close()
}

func isNetworkErr(err error) bool {
	return err != nil &&
		err != ErrClosing &&
		err != context.Canceled &&
		err != context.DeadlineExceeded
}
