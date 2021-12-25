package rueidis

import (
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/rueian/rueidis/internal/cmds"
	"github.com/rueian/rueidis/internal/proto"
)

type connFn func(dst string, opt ConnOption) conn
type dialFn func(dst string, opt ConnOption) (net.Conn, error)
type wireFn func(conn net.Conn, opt ConnOption, onDisconnected func(err error)) (wire, error)

type singleconnect struct {
	w wire
	e error
	g sync.WaitGroup
}

type conn interface {
	wire
	Dial() error
	Acquire() wire
	Store(w wire)
	OnDisconnected(func(err error))
}

var _ conn = (*mux)(nil)

type mux struct {
	dst  string
	opt  ConnOption
	pool *pool
	dead wire
	wire atomic.Value
	mu   sync.Mutex
	sc   *singleconnect

	dialFn dialFn
	wireFn wireFn

	onDisconnected atomic.Value
}

func makeMux(dst string, option ConnOption, dialFn dialFn) *mux {
	return newMux(dst, option, (*pipe)(nil), dialFn, func(conn net.Conn, opt ConnOption, onDisconnected func(err error)) (wire, error) {
		return newPipe(conn, opt, onDisconnected)
	})
}

func newMux(dst string, option ConnOption, dead wire, dialFn dialFn, wireFn wireFn) *mux {
	m := &mux{dst: dst, opt: option, dead: dead, dialFn: dialFn, wireFn: wireFn}
	m.wire.Store(dead)
	m.pool = newPool(option.BlockingPoolSize, m.dialRetry)
	return m
}

func (m *mux) connect() (w wire, err error) {
	if w = m.wire.Load().(wire); w != m.dead {
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

	if w = m.wire.Load().(wire); w == m.dead {
		if w, err = m.dial(); err == nil {
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

func (m *mux) dial() (w wire, err error) {
	conn, err := m.dialFn(m.dst, m.opt)
	if err == nil {
		w, err = m.wireFn(conn, m.opt, m.disconnected)
	}
	return w, err
}

func (m *mux) disconnected(err error) {
	if fn := m.onDisconnected.Load(); fn != nil {
		fn.(func(err error))(err)
	}
}

func (m *mux) OnDisconnected(fn func(err error)) {
	m.onDisconnected.CompareAndSwap(nil, fn)
}

func (m *mux) dialRetry() wire {
retry:
	if wire, err := m.dial(); err == nil {
		return wire
	}
	goto retry
}

func (m *mux) acquire() wire {
retry:
	if wire, err := m.connect(); err == nil {
		return wire
	}
	goto retry
}

func (m *mux) Dial() error { // no retry
	_, err := m.connect()
	return err
}

func (m *mux) Info() map[string]proto.Message {
	return m.acquire().Info()
}

func (m *mux) Error() error {
	return m.acquire().Error()
}

func (m *mux) Do(cmd cmds.Completed) (resp proto.Result) {
retry:
	if cmd.IsBlock() {
		resp = m.blocking(cmd)
	} else {
		resp = m.pipeline(cmd)
	}
	if cmd.IsReadOnly() && isNetworkErr(resp.NonRedisError()) {
		goto retry
	}
	return resp
}

func (m *mux) DoMulti(multi ...cmds.Completed) (resp []proto.Result) {
	var block, write bool
	for _, cmd := range multi {
		block = block || cmd.IsBlock()
		write = write || cmd.IsWrite()
	}
retry:
	if block {
		resp = m.blockingMulti(multi)
	} else {
		resp = m.pipelineMulti(multi)
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

func (m *mux) blocking(cmd cmds.Completed) (resp proto.Result) {
	wire := m.pool.Acquire()
	resp = wire.Do(cmd)
	m.pool.Store(wire)
	return resp
}

func (m *mux) blockingMulti(cmd []cmds.Completed) (resp []proto.Result) {
	wire := m.pool.Acquire()
	resp = wire.DoMulti(cmd...)
	m.pool.Store(wire)
	return resp
}

func (m *mux) pipeline(cmd cmds.Completed) (resp proto.Result) {
	wire := m.acquire()
	if resp = wire.Do(cmd); isNetworkErr(resp.NonRedisError()) {
		m.wire.CompareAndSwap(wire, m.dead)
	}
	return resp
}

func (m *mux) pipelineMulti(cmd []cmds.Completed) (resp []proto.Result) {
	wire := m.acquire()
	resp = wire.DoMulti(cmd...)
	for _, r := range resp {
		if isNetworkErr(r.NonRedisError()) {
			m.wire.CompareAndSwap(wire, m.dead)
			return resp
		}
	}
	return resp
}

func (m *mux) DoCache(cmd cmds.Cacheable, ttl time.Duration) proto.Result {
retry:
	wire := m.acquire()
	resp := wire.DoCache(cmd, ttl)
	if isNetworkErr(resp.NonRedisError()) {
		m.wire.CompareAndSwap(wire, m.dead)
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
	m.acquire().Close()
	m.pool.Close()
}

func isNetworkErr(err error) bool {
	return err != nil && err != ErrConnClosing
}
