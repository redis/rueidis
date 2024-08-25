package rueidislock

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/redis/rueidis"
	"github.com/redis/rueidis/internal/util"
)

// LockerOption should be passed to NewLocker to construct a Locker
type LockerOption struct {
	// ClientBuilder can be used to modify rueidis.Client used by Locker
	ClientBuilder func(option rueidis.ClientOption) (rueidis.Client, error)
	// KeyPrefix is the prefix of redis key for locks. Default value is "rueidislock".
	KeyPrefix string
	// ClientOption is passed to rueidis.NewClient or LockerOption.ClientBuilder to build a rueidis.Client
	ClientOption rueidis.ClientOption
	// KeyValidity is the validity duration of locks and will be extended periodically by the ExtendInterval. Default value is 5s.
	KeyValidity time.Duration
	// ExtendInterval is the interval to extend KeyValidity. Default value is 1s.
	ExtendInterval time.Duration
	// TryNextAfter is the timeout duration before trying the next redis key for locks. Default value is 20ms.
	TryNextAfter time.Duration
	// KeyMajority is at least how many redis keys in a total of KeyMajority*2-1 should be acquired to be a valid lock.
	// Default value is 2.
	KeyMajority int32
	// NoLoopTracking will use NOLOOP in the CLIENT TRACKING command to avoid unnecessary notifications and thus have better performance.
	// This can only be enabled if all your redis nodes >= 7.0.5. (https://github.com/redis/redis/pull/11052)
	NoLoopTracking bool
	// Use SET PX instead of SET PXAT when acquiring locks to be compatible with Redis < 6.2
	FallbackSETPX bool
}

// Locker is the interface of rueidislock
type Locker interface {
	// WithContext acquires a distributed redis lock by name by waiting for it. It may return ErrLockerClosed.
	WithContext(ctx context.Context, name string) (context.Context, context.CancelFunc, error)
	// TryWithContext tries to acquire a distributed redis lock by name without waiting. It may return ErrNotLocked.
	TryWithContext(ctx context.Context, name string) (context.Context, context.CancelFunc, error)
	// ForceWithContext takes over a distributed redis lock by canceling the original holder. It may return ErrNotLocked.
	ForceWithContext(ctx context.Context, name string) (context.Context, context.CancelFunc, error)
	// Client exports the underlying rueidis.Client
	Client() rueidis.Client
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
		option.ExtendInterval = option.KeyValidity / 2
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
		noloop:   option.NoLoopTracking,
		setpx:    option.FallbackSETPX,
	}

	if option.ClientOption.DisableCache {
		impl.noloop = true
		impl.nocsc = true
	} else {
		if option.NoLoopTracking {
			option.ClientOption.ClientTrackingOptions = []string{"OPTOUT", "NOLOOP"}
		} else {
			option.ClientOption.ClientTrackingOptions = []string{"OPTOUT"}
		}
		option.ClientOption.OnInvalidations = impl.onInvalidations
	}
	option.ClientOption.PipelineMultiplex = -1 // this ensures the CSC goes to the same connection.

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
	gates    map[string]*gate
	prefix   string
	validity time.Duration
	interval time.Duration
	timeout  time.Duration
	mu       sync.RWMutex
	majority int32
	totalcnt int32
	noloop   bool
	nocsc    bool
	setpx    bool
}

type gate struct {
	ch  chan struct{}
	csc []chan struct{}
	w   int
}

func makegate(size int32) *gate {
	csc := make([]chan struct{}, size)
	for i := 0; i < len(csc); i++ {
		csc[i] = make(chan struct{}, 1)
	}
	return &gate{ch: make(chan struct{}, 1), csc: csc}
}

func random() string {
	return rueidis.BinaryString(util.RandomBytes())
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

func (m *locker) acquire(ctx context.Context, key, val string, deadline time.Time, force bool) (err error) {
	ctx, cancel := context.WithTimeout(ctx, m.timeout)
	var resp rueidis.RedisResult
	if force {
		if m.setpx {
			resp = fcqms.Exec(ctx, m.client, []string{key}, []string{val, strconv.FormatInt(m.validity.Milliseconds(), 10)})
		} else {
			resp = fcqat.Exec(ctx, m.client, []string{key}, []string{val, strconv.FormatInt(deadline.UnixMilli(), 10)})
		}
	} else {
		if m.setpx {
			resp = acqms.Exec(ctx, m.client, []string{key}, []string{val, strconv.FormatInt(m.validity.Milliseconds(), 10)})
		} else {
			resp = acqat.Exec(ctx, m.client, []string{key}, []string{val, strconv.FormatInt(deadline.UnixMilli(), 10)})
		}
	}
	cancel()
	if err = resp.Error(); rueidis.IsRedisNil(err) {
		return fmt.Errorf("%w: key %s is held by others", ErrNotLocked, key)
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
		if m.gates == nil {
			m.mu.Unlock()
			return nil, ErrLockerClosed
		}
		g = makegate(m.totalcnt)
		g.w++
		m.gates[name] = g
		m.mu.Unlock()
		return g, nil
	} else {
		g.w++
		m.mu.Unlock()
	}
	var timeout <-chan time.Time
	if m.nocsc {
		timeout = time.After(m.timeout)
	}
	select {
	case <-ctx.Done():
		m.removegate(g, name)
		return nil, ctx.Err()
	case _, ok = <-g.ch:
		if ok {
			return g, nil
		}
		return nil, ErrLockerClosed
	case <-timeout:
		return g, nil
	}
}

func (m *locker) trygate(name string) (g *gate) {
	m.mu.Lock()
	if _, ok := m.gates[name]; !ok && m.gates != nil {
		g = makegate(m.totalcnt)
		g.w++
		m.gates[name] = g
	}
	m.mu.Unlock()
	return g
}

func (m *locker) forcegate(name string) (g *gate) {
	m.mu.Lock()
	if g = m.gates[name]; g == nil && m.gates != nil {
		g = makegate(m.totalcnt)
		m.gates[name] = g
	}
	if g != nil {
		g.w++
	}
	m.mu.Unlock()
	return g
}

func (m *locker) removegate(g *gate, name string) {
	m.mu.Lock()
	if g.w--; g.w == 0 && m.gates[name] == g {
		delete(m.gates, name)
	}
	m.mu.Unlock()
}

func (m *locker) onInvalidations(messages []rueidis.RedisMessage) {
	if messages == nil {
		m.mu.RLock()
		for _, g := range m.gates {
			for _, ch := range g.csc {
				select {
				case ch <- struct{}{}:
				default:
				}
			}
			select {
			case g.ch <- struct{}{}:
			default:
			}
		}
		m.mu.RUnlock()
	}
	for _, msg := range messages {
		k, _ := msg.ToString()
		if strings.HasPrefix(k, m.prefix) {
			if ks := strings.SplitN(k[len(m.prefix)+1:], ":", 2); len(ks) == 2 {
				m.mu.RLock()
				g, ok := m.gates[ks[1]]
				if ok {
					n, _ := strconv.Atoi(ks[0])
					select {
					case g.csc[n] <- struct{}{}:
					default:
					}
					select {
					case g.ch <- struct{}{}:
					default:
					}
				}
				m.mu.RUnlock()
			}
		}
	}
}

func (m *locker) try(ctx context.Context, cancel context.CancelFunc, name string, g *gate, force bool) (context.CancelFunc, error) {
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
						if !m.noloop {
							<-csc
						}
					}
				case _, ok := <-csc:
					if !ok {
						err = ErrLockerClosed
					} else {
						if err = m.script(ctx, extend, key, val, deadline); err == nil {
							if !m.noloop {
								<-csc
							}
						}
					}
				}
			}
		}
		if !errors.Is(err, ErrNotLocked) {
			_ = m.script(context.Background(), delkey, key, val, deadline)
		}
		if released := atomic.AddInt32(&released, 1); released >= m.majority {
			cancel()
			if released == m.totalcnt && atomic.LoadInt32(&failures) < m.majority {
				m.mu.Lock()
				if g.w--; g.w == 0 {
					if m.gates[name] == g {
						delete(m.gates, name)
					}
				} else if m.gates != nil {
					select {
					case g.ch <- struct{}{}:
					default:
					}
				}
				m.mu.Unlock()
			}
			if released == m.totalcnt {
				close(done)
			}
		}
	}

	acquire := func(err error, key string, ch chan struct{}, force bool) error {
		select {
		case <-ch:
		default:
		}
		if !errors.Is(err, ErrNotLocked) {
			if err = m.acquire(ctx, key, val, deadline, force); force && err == nil {
				m.mu.RLock()
				if m.gates != nil {
					select {
					case ch <- struct{}{}:
					default:
					}
				}
				m.mu.RUnlock()
			}
		}
		go monitoring(err, key, deadline, ch)
		return err
	}

	var i int32
	for ; atomic.LoadInt32(&acquired) < m.majority && atomic.LoadInt32(&failures) < m.majority; i++ {
		if err = acquire(err, keyname(m.prefix, name, i), g.csc[i], force); err == nil {
			atomic.AddInt32(&acquired, 1)
		} else {
			atomic.AddInt32(&failures, 1)
		}
	}
	if i < m.totalcnt {
		go func(i int32, err error) {
			for ; i < m.totalcnt; i++ {
				err = acquire(err, keyname(m.prefix, name, i), g.csc[i], force)
			}
		}(i, err)
	}
	if cacneltm.Stop() && atomic.LoadInt32(&failures) < m.majority {
		return func() {
			cancel()
			<-done
		}, nil
	}
	<-done
	if err == nil {
		err = fmt.Errorf("%w: failed to acquire the majority of keys (%d/%d)", ErrNotLocked, atomic.LoadInt32(&acquired), m.totalcnt)
	}
	return cancel, err
}

func (m *locker) ForceWithContext(ctx context.Context, name string) (context.Context, context.CancelFunc, error) {
	var err error
	ctx, cancel := context.WithCancel(ctx)
	if g := m.forcegate(name); g != nil {
		if cancel, err = m.try(ctx, cancel, name, g, true); err == nil {
			return ctx, cancel, nil
		}
		m.removegate(g, name)
	}
	cancel()
	if err == nil {
		err = ErrLockerClosed
	}
	return ctx, cancel, err
}

func (m *locker) TryWithContext(ctx context.Context, name string) (context.Context, context.CancelFunc, error) {
	var err error
	ctx, cancel := context.WithCancel(ctx)
	if g := m.trygate(name); g != nil {
		if cancel, err = m.try(ctx, cancel, name, g, false); err == nil {
			return ctx, cancel, nil
		}
		m.removegate(g, name)
	}
	cancel()
	if err == nil {
		err = fmt.Errorf("%w: the lock is held by others or the locker is closed", ErrNotLocked)
	}
	return ctx, cancel, err
}

func (m *locker) WithContext(src context.Context, name string) (context.Context, context.CancelFunc, error) {
	for {
		ctx, cancel := context.WithCancel(src)
		g, err := m.waitgate(ctx, name)
		if g != nil {
			if cancel, err := m.try(ctx, cancel, name, g, false); err == nil {
				return ctx, cancel, nil
			}
			m.mu.Lock()
			g.w-- // do not delete g from m.gates here.
			m.mu.Unlock()
		}
		if cancel(); err != nil {
			return ctx, cancel, err
		}
	}
}

func (m *locker) Client() rueidis.Client {
	return m.client
}

func (m *locker) Close() {
	m.mu.Lock()
	for _, g := range m.gates {
		for _, ch := range g.csc {
			close(ch)
		}
		close(g.ch)
	}
	m.gates = nil
	m.mu.Unlock()
	m.client.Close()
}

var (
	delkey = rueidis.NewLuaScript(`if redis.call("GET",KEYS[1]) == ARGV[1] then return redis.call("DEL",KEYS[1]) end;return 0`)
	extend = rueidis.NewLuaScript(`if redis.call("GET",KEYS[1]) == ARGV[1] then local r = redis.call("PEXPIREAT",KEYS[1],ARGV[2]);redis.call("GET",KEYS[1]);return r end;return 0`)
	acqms  = rueidis.NewLuaScript(`local r = redis.call("SET",KEYS[1],ARGV[1],"NX","PX",ARGV[2]);redis.call("GET",KEYS[1]);return r`)
	acqat  = rueidis.NewLuaScript(`local r = redis.call("SET",KEYS[1],ARGV[1],"NX","PXAT",ARGV[2]);redis.call("GET",KEYS[1]);return r`)
	fcqms  = rueidis.NewLuaScript(`local r = redis.call("SET",KEYS[1],ARGV[1],"PX",ARGV[2]);redis.call("GET",KEYS[1]);return r`)
	fcqat  = rueidis.NewLuaScript(`local r = redis.call("SET",KEYS[1],ARGV[1],"PXAT",ARGV[2]);redis.call("GET",KEYS[1]);return r`)
)

// ErrNotLocked is returned from the Locker.TryWithContext when it fails
var ErrNotLocked = errors.New("not locked")

// ErrLockerClosed is returned from the Locker.WithContext when the Locker is closed
var ErrLockerClosed = errors.New("locker closed")
