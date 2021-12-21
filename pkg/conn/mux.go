package conn

import (
	"crypto/tls"
	"errors"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/rueian/rueidis/internal/cmds"
	"github.com/rueian/rueidis/internal/proto"
)

var ErrConnClosing = errors.New("connection is closing")

// DefaultCacheBytes = 128 MiB.
const DefaultCacheBytes = 128 * (1 << 20)

// DefaultPoolSize = 1000
const DefaultPoolSize = 1000

type Option struct {
	// CacheSizeEachConn is redis client side cache size that bind to each TCP connection to a single redis instance.
	// The default is DefaultCacheBytes.
	CacheSizeEachConn int

	// BlockingPoolSize is the size of the connection pool shared by blocking commands (ex BLPOP, XREAD with BLOCK).
	// The default is DefaultPoolSize.
	BlockingPoolSize int

	// Redis AUTH parameters
	Username   string
	Password   string
	ClientName string
	SelectDB   int

	// TCP & TLS
	DialTimeout time.Duration
	TLSConfig   *tls.Config

	// Redis PubSub callbacks
	PubSubHandlers PubSubHandlers
}

type ConnFn func(dst string, opt Option) Conn
type DialFn func(dst string, opt Option) (net.Conn, error)
type wireFn func(conn net.Conn, opt Option) (Wire, error)

type singleconnect struct {
	w Wire
	e error
	g sync.WaitGroup
}

type Conn interface {
	Wire
	Dial() error
	Acquire() Wire
	Store(w Wire)
}

type mux struct {
	dst  string
	opt  Option
	pool *pool
	dead Wire
	wire atomic.Value
	mu   sync.Mutex
	sc   *singleconnect

	dialFn DialFn
	wireFn wireFn
}

func NewMux(dst string, option Option, dialFn DialFn) *mux {
	return newMux(dst, option, (*wire)(nil), dialFn, func(conn net.Conn, opt Option) (Wire, error) {
		return newWire(conn, opt)
	})
}

func newMux(dst string, option Option, dead Wire, dialFn DialFn, wireFn wireFn) *mux {
	conn := &mux{dst: dst, opt: option, dead: dead, dialFn: dialFn, wireFn: wireFn}
	conn.wire.Store(dead)
	conn.pool = newPool(option.BlockingPoolSize, conn.dialRetry)
	return conn
}

func (m *mux) connect() (w Wire, err error) {
	if w = m.wire.Load().(Wire); w != m.dead {
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

	if w = m.wire.Load().(Wire); w == m.dead {
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

func (m *mux) dial() (w Wire, err error) {
	conn, err := m.dialFn(m.dst, m.opt)
	if err == nil {
		w, err = m.wireFn(conn, m.opt)
	}
	return w, err
}

func (m *mux) dialRetry() Wire {
retry:
	if wire, err := m.dial(); err == nil {
		return wire
	}
	goto retry
}

func (m *mux) acquire() Wire {
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

func (m *mux) Acquire() Wire {
	return m.pool.Acquire()
}

func (m *mux) Store(w Wire) {
	m.pool.Store(w)
}

func (m *mux) Close() {
	m.acquire().Close()
	m.pool.Close()
}

func isNetworkErr(err error) bool {
	return err != nil && err != ErrConnClosing
}
