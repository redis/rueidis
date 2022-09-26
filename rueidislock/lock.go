package rueidislock

import (
	"context"
	"encoding/binary"
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/rueian/rueidis"
)

var sources sync.Pool

func init() {
	sources = sync.Pool{New: func() any { return rand.New(rand.NewSource(time.Now().UnixNano())) }}
}

type LockerOption struct {
	ClientOption   rueidis.ClientOption
	KeyPrefix      string
	KeyMajority    int
	KeyValidity    time.Duration
	ExtendInterval time.Duration
	TryNextAfter   time.Duration
}

type Locker interface {
	WithContext(ctx context.Context, name string) (context.Context, context.CancelFunc, error)
	Close()
}

func NewLocker(option LockerOption) (Locker, error) {
	if option.KeyPrefix == "" {
		option.KeyPrefix = "rueidislock"
	}
	if option.KeyValidity <= 0 {
		option.KeyValidity = time.Second * 5
	}
	if option.ExtendInterval <= 0 {
		option.ExtendInterval = time.Millisecond * 500
	}
	if option.TryNextAfter <= 0 {
		option.TryNextAfter = time.Millisecond * 5
	}
	if option.KeyMajority <= 0 {
		option.KeyMajority = 2
	}
	impl := &locker{
		prefix:   option.KeyPrefix,
		validity: option.KeyValidity,
		interval: option.ExtendInterval,
		timeout:  option.TryNextAfter,
		majority: option.KeyMajority,
		totalcnt: option.KeyMajority*2 - 1,
		gates:    make(map[string]*gate),
	}
	option.ClientOption.DisableCache = false
	option.ClientOption.ClientTrackingOptions = []string{"OPTOUT"}
	option.ClientOption.OnInvalidations = impl.onInvalidations
	client, err := rueidis.NewClient(option.ClientOption)
	if err != nil {
		return nil, err
	}
	impl.client = client
	return impl, nil
}

type locker struct {
	client   rueidis.Client
	prefix   string
	validity time.Duration
	interval time.Duration
	timeout  time.Duration
	majority int
	totalcnt int

	mu    sync.RWMutex
	gates map[string]*gate
}

type gate struct {
	w   int
	ch  chan struct{}
	csc []chan struct{}
}

func makegate(size int) *gate {
	csc := make([]chan struct{}, size)
	for i := 0; i < len(csc); i++ {
		csc[i] = make(chan struct{}, 1)
	}
	return &gate{ch: make(chan struct{}, 1), csc: csc}
}

func random() string {
	val := make([]byte, 24)
	src := sources.Get().(rand.Source64)
	binary.BigEndian.PutUint64(val[0:8], src.Uint64())
	binary.BigEndian.PutUint64(val[8:16], src.Uint64())
	binary.BigEndian.PutUint64(val[16:24], src.Uint64())
	sources.Put(src)
	return rueidis.BinaryString(val)
}

func keyname(prefix, name string, i int) string {
	ia := strconv.Itoa(i)
	sb := strings.Builder{}
	sb.Grow(len(prefix) + len(name) + len(ia) + 2)
	sb.WriteString(prefix)
	sb.WriteByte(':')
	sb.WriteString(ia)
	sb.WriteByte(':')
	sb.WriteString(name)
	return sb.String()
}

func (m *locker) acquire(ctx context.Context, key, val string, deadline time.Time) error {
	ctx, cancel := context.WithTimeout(ctx, m.timeout)
	resp := m.client.DoMulti(ctx,
		m.client.B().Set().Key(key).Value(val).Nx().PxatMillisecondsTimestamp(deadline.UnixMilli()).Build(),
		m.client.B().Get().Key(key).Build(),
	)
	cancel()
	if v, _ := resp[1].ToString(); v != val {
		return errNotLock
	}
	return nil
}

func (m *locker) script(ctx context.Context, script *rueidis.Lua, key, val string, deadline time.Time) error {
	ctx, cancel := context.WithDeadline(ctx, deadline)
	resp := script.Exec(ctx, m.client, []string{key}, []string{val, strconv.FormatInt(deadline.UnixMilli(), 10)})
	cancel()
	if v, err := resp.AsInt64(); err != nil || v == 1 {
		return err
	}
	return errNotLock
}

func (m *locker) waitgate(ctx context.Context, name string) (g *gate, err error) {
	m.mu.Lock()
	g, ok := m.gates[name]
	if !ok {
		g = makegate(m.totalcnt)
		g.w++
		m.gates[name] = g
		m.mu.Unlock()
		return g, nil
	} else {
		g.w++
		m.mu.Unlock()
	}
	select {
	case <-ctx.Done():
		m.mu.Lock()
		if g.w--; g.w == 0 {
			delete(m.gates, name)
		}
		m.mu.Unlock()
		return nil, ctx.Err()
	case _, ok = <-g.ch:
		if ok {
			return g, nil
		}
		m.mu.Lock()
		if g.w--; g.w == 0 {
			delete(m.gates, name)
		}
		m.mu.Unlock()
		return nil, ErrLockerClosed
	}
}

func (m *locker) onInvalidations(messages []rueidis.RedisMessage) {
	if messages == nil {
		m.mu.Lock()
		for _, g := range m.gates {
			close(g.ch)
			for _, ch := range g.csc {
				close(ch)
			}
		}
		m.gates = make(map[string]*gate)
		m.mu.Unlock()
	}
	for _, msg := range messages {
		k, _ := msg.ToString()
		if ks := strings.SplitN(k, ":", 3); len(ks) == 3 {
			m.mu.RLock()
			g, ok := m.gates[ks[2]]
			if ok {
				n, _ := strconv.Atoi(ks[1])
				select {
				case g.csc[n] <- struct{}{}:
				default:
				}
			}
			m.mu.RUnlock()
		}
	}
}

func (m *locker) WithContext(ctx context.Context, name string) (context.Context, context.CancelFunc, error) {
	for {
		ctx, cancel := context.WithCancel(ctx)

		g, err := m.waitgate(ctx, name)
		if err != nil {
			return ctx, cancel, err
		}

		val := random()
		deadline := time.Now().Add(m.validity)
		cacneltm := time.AfterFunc(m.validity, cancel)
		released := int32(0)

		monitoring := func(err error, key string, deadline time.Time, csc chan struct{}) {
			if err == nil {
				for timer := time.NewTimer(m.interval); err == nil; {
					select {
					case <-ctx.Done():
						err = ctx.Err()
					case <-timer.C:
						deadline = deadline.Add(m.interval)
						if err = m.script(ctx, extend, key, val, deadline); err == nil {
							timer.Reset(m.interval)
							<-csc
						}
					case <-csc:
						if err = m.script(ctx, extend, key, val, deadline); err == nil {
							<-csc
						}
					}
				}
			}
			if err != errNotLock {
				err = m.script(context.Background(), delkey, key, val, deadline)
			}
			if released := int(atomic.AddInt32(&released, 1)); released >= m.majority {
				cancel()
				if released == m.totalcnt {
					m.mu.Lock()
					if g.w--; g.w == 0 {
						delete(m.gates, name)
						m.mu.Unlock()
					} else {
						m.mu.Unlock()
						g.ch <- struct{}{}
					}
				}
			}
		}

		acquire := func(err error, key string, ch chan struct{}) error {
			select {
			case <-ch:
			default:
			}
			if err != errNotLock {
				err = m.acquire(ctx, key, val, deadline)
			}
			go monitoring(err, key, deadline, ch)
			return err
		}

		var i int
		var failures int
		for acquired := 0; acquired < m.majority && failures < m.majority; i++ {
			if err = acquire(err, keyname(m.prefix, name, i), g.csc[i]); err == nil {
				acquired++
			} else {
				failures++
			}
		}
		if i != m.totalcnt {
			go func() {
				for ; i < m.totalcnt; i++ {
					err = acquire(err, keyname(m.prefix, name, i), g.csc[i])
				}
			}()
		}
		if cacneltm.Stop() && failures < m.majority {
			return ctx, cancel, nil
		}
	}
}

func (m *locker) Close() {
	m.client.Close()
}

var (
	delkey = rueidis.NewLuaScript(`if redis.call("GET",KEYS[1]) == ARGV[1] then return redis.call("DEL",KEYS[1]) else return 0 end`)
	extend = rueidis.NewLuaScript(`if redis.call("GET",KEYS[1]) == ARGV[1] then return redis.call("PEXPIREAT",KEYS[1],ARGV[2]) else return 0 end`)
)

var errNotLock = errors.New("not lock")
var ErrLockerClosed = errors.New("locker closed")
