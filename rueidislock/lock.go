package rueidislock

import (
	"context"
	"encoding/binary"
	"errors"
	"math/rand"
	"reflect"
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

// LockerOption should be passed to NewLocker to construct a Locker
type LockerOption struct {
	// ClientOption is passed to rueidis.NewClient or LockerOption.ClientBuilder to build a rueidis.Client
	ClientOption rueidis.ClientOption
	// ClientBuilder can be used to modify rueidis.Client used by Locker
	ClientBuilder func(option rueidis.ClientOption) (rueidis.Client, error)
	// KeyPrefix is the prefix of redis key for locks. Default value is "rueidislock".
	KeyPrefix string
	// KeyValidity is the validity duration of locks and will be extended periodically by the ExtendInterval. Default value is 5s.
	KeyValidity time.Duration
	// ExtendInterval is the interval to extend KeyValidity. Default value is 1s.
	ExtendInterval time.Duration
	// TryNextAfter is the timeout duration before trying the next redis key for locks. Default value is 20ms.
	TryNextAfter time.Duration
	// KeyMajority is at least how many redis keys in a total of KeyMajority*2-1 should be acquired to be a valid lock.
	// Default value is 2.
	KeyMajority int32
}

// Locker is the interface of rueidislock
type Locker interface {
	// WithContext acquires a distributed redis lock by name by waiting for it. It may return ErrLockerClosed.
	WithContext(ctx context.Context, name string) (context.Context, context.CancelFunc, error)
	// TryWithContext tries to acquire a distributed redis lock by name without waiting. It may return ErrNotLocked.
	TryWithContext(ctx context.Context, name string) (context.Context, context.CancelFunc, error)
	// Close closes the underlying rueidis.Client
	Close()
}

// NewLocker creates the distributed Locker backed by redis client side caching
func NewLocker(option LockerOption) (Locker, error) {
	if option.KeyPrefix == "" {
		option.KeyPrefix = "rueidislock"
	}
	if option.KeyValidity <= 0 {
		option.KeyValidity = time.Second * 5
	}
	if option.ExtendInterval <= 0 {
		option.ExtendInterval = time.Second
	}
	if option.TryNextAfter <= 0 {
		option.TryNextAfter = time.Millisecond * 20
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

	var err error
	if option.ClientBuilder != nil {
		impl.client, err = option.ClientBuilder(option.ClientOption)
	} else {
		impl.client, err = rueidis.NewClient(option.ClientOption)
	}
	if err != nil {
		return nil, err
	}
	return impl, nil
}

type locker struct {
	client   rueidis.Client
	prefix   string
	validity time.Duration
	interval time.Duration
	timeout  time.Duration
	majority int32
	totalcnt int32

	mu    sync.RWMutex
	gates map[string]*gate
}

type gate struct {
	w   int
	ch  chan struct{}
	csc []chan struct{}
}

func makegate(size int32) *gate {
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

func keyname(prefix, name string, i int32) string {
	ia := strconv.Itoa(int(i))
	sb := strings.Builder{}
	sb.Grow(len(prefix) + len(name) + len(ia) + 2)
	sb.WriteString(prefix)
	sb.WriteByte(':')
	sb.WriteString(ia)
	sb.WriteByte(':')
	sb.WriteString(name)
	return sb.String()
}

func (m *locker) acquire(ctx context.Context, key, val string, deadline time.Time) (err error) {
	ctx, cancel := context.WithTimeout(ctx, m.timeout)
	resp := m.client.DoMulti(ctx,
		m.client.B().Set().Key(key).Value(val).Nx().PxatMillisecondsTimestamp(deadline.UnixMilli()).Build(),
		m.client.B().Get().Key(key).Build(),
	)
	cancel()
	if err = resp[0].Error(); rueidis.IsRedisNil(err) {
		return ErrNotLocked
	}
	return err
}

func (m *locker) script(ctx context.Context, script *rueidis.Lua, key, val string, deadline time.Time) error {
	ctx, cancel := context.WithDeadline(ctx, deadline)
	resp := script.Exec(ctx, m.client, []string{key}, []string{val, strconv.FormatInt(deadline.UnixMilli(), 10)})
	cancel()
	if v, err := resp.AsInt64(); err != nil || v == 1 {
		return err
	}
	return ErrNotLocked
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
		return nil, ErrLockerClosed
	}
}

func (m *locker) trygate(name string) (g *gate) {
	m.mu.Lock()
	g, ok := m.gates[name]
	if !ok {
		g = makegate(m.totalcnt)
		g.w++
		m.gates[name] = g
	}
	m.mu.Unlock()
	return g
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

func (m *locker) try(ctx context.Context, cancel context.CancelFunc, name string, g *gate) context.CancelFunc {
	var err error

	val := random()
	deadline := time.Now().Add(m.validity)
	cacneltm := time.AfterFunc(m.validity, cancel)
	released := int32(0)
	acquired := int32(0)
	failures := int32(0)

	done := make(chan struct{})
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
		if err != ErrNotLocked {
			_ = m.script(context.Background(), delkey, key, val, deadline)
		}
		if released := atomic.AddInt32(&released, 1); released >= m.majority {
			cancel()
			if released == m.totalcnt {
				if atomic.LoadInt32(&failures) >= m.majority {
					tm := time.NewTimer(m.interval)
					cases := make([]reflect.SelectCase, 0, m.totalcnt+1)
					cases = append(cases, reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(tm.C)})
					for i := int32(0); i < m.totalcnt; i++ {
						cases = append(cases, reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(g.csc[i])})
					}
					reflect.Select(cases)
					tm.Stop()
				}
				m.mu.Lock()
				if g.w--; g.w == 0 {
					delete(m.gates, name)
				} else {
					if _, ok := m.gates[name]; ok {
						g.ch <- struct{}{}
					}
				}
				m.mu.Unlock()
				close(done)
			}
		}
	}

	acquire := func(err error, key string, ch chan struct{}) error {
		select {
		case <-ch:
		default:
		}
		if err != ErrNotLocked {
			err = m.acquire(ctx, key, val, deadline)
		}
		go monitoring(err, key, deadline, ch)
		return err
	}

	var i int32
	for a, f := atomic.LoadInt32(&acquired), atomic.LoadInt32(&failures); a < m.majority && f < m.majority; i++ {
		if err = acquire(err, keyname(m.prefix, name, i), g.csc[i]); err == nil {
			a = atomic.AddInt32(&acquired, 1)
		} else {
			f = atomic.AddInt32(&failures, 1)
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
		return func() {
			cancel()
			<-done
		}
	}
	return nil
}

func (m *locker) TryWithContext(ctx context.Context, name string) (context.Context, context.CancelFunc, error) {
	ctx, cancel := context.WithCancel(ctx)
	if g := m.trygate(name); g != nil {
		if cancel := m.try(ctx, cancel, name, g); cancel != nil {
			return ctx, cancel, nil
		}
	}
	cancel()
	return ctx, cancel, ErrNotLocked
}

func (m *locker) WithContext(ctx context.Context, name string) (context.Context, context.CancelFunc, error) {
	for {
		ctx, cancel := context.WithCancel(ctx)
		g, err := m.waitgate(ctx, name)
		if g != nil {
			if cancel := m.try(ctx, cancel, name, g); cancel != nil {
				return ctx, cancel, nil
			}
		}
		if cancel(); err != nil {
			return ctx, cancel, err
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

// ErrNotLocked is returned from the Locker.TryWithContext when it fails
var ErrNotLocked = errors.New("not locked")

// ErrLockerClosed is returned from the Locker.WithContext when the Locker is closed
var ErrLockerClosed = errors.New("locker closed")
