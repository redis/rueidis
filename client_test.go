package rueidis

import (
	"context"
	"errors"
	"io"
	"net"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/redis/rueidis/internal/cmds"
)

type mockConn struct {
	DoFn            func(cmd Completed) RedisResult
	DoCacheFn       func(cmd Cacheable, ttl time.Duration) RedisResult
	DoMultiFn       func(multi ...Completed) *redisresults
	DoMultiCacheFn  func(multi ...CacheableTTL) *redisresults
	ReceiveFn       func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error
	DoStreamFn      func(cmd Completed) RedisResultStream
	DoMultiStreamFn func(cmd ...Completed) MultiRedisResultStream
	InfoFn          func() map[string]RedisMessage
	VersionFn       func() int
	AZFn            func() string
	ErrorFn         func() error
	CloseFn         func()
	DialFn          func() error
	AcquireFn       func() wire
	StoreFn         func(w wire)
	OverrideFn      func(c conn)
	AddrFn          func() string

	DoOverride      map[string]func(cmd Completed) RedisResult
	DoCacheOverride map[string]func(cmd Cacheable, ttl time.Duration) RedisResult
	ReceiveOverride map[string]func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error
}

func (m *mockConn) Override(c conn) {
	if m.OverrideFn != nil {
		m.OverrideFn(c)
	}
}

func (m *mockConn) Dial() error {
	if m.DialFn != nil {
		return m.DialFn()
	}
	return nil
}

func (m *mockConn) Acquire(ctx context.Context) wire {
	if m.AcquireFn != nil {
		return m.AcquireFn()
	}
	return nil
}

func (m *mockConn) Store(w wire) {
	if m.StoreFn != nil {
		m.StoreFn(w)
	}
}

func (m *mockConn) Do(ctx context.Context, cmd Completed) RedisResult {
	if fn := m.DoOverride[strings.Join(cmd.Commands(), " ")]; fn != nil {
		return fn(cmd)
	}
	if m.DoFn != nil {
		return m.DoFn(cmd)
	}
	return RedisResult{}
}

func (m *mockConn) DoCache(ctx context.Context, cmd Cacheable, ttl time.Duration) RedisResult {
	if fn := m.DoCacheOverride[strings.Join(cmd.Commands(), " ")]; fn != nil {
		return fn(cmd, ttl)
	}
	if m.DoCacheFn != nil {
		return m.DoCacheFn(cmd, ttl)
	}
	return RedisResult{}
}

func (m *mockConn) DoMultiCache(ctx context.Context, multi ...CacheableTTL) *redisresults {
	overrides := make([]RedisResult, 0, len(multi))
	for _, cmd := range multi {
		if fn := m.DoCacheOverride[strings.Join(cmd.Cmd.Commands(), " ")]; fn != nil {
			overrides = append(overrides, fn(cmd.Cmd, cmd.TTL))
		}
	}
	if len(overrides) == len(multi) {
		return &redisresults{s: overrides}
	}
	if m.DoMultiCacheFn != nil {
		return m.DoMultiCacheFn(multi...)
	}
	return nil
}

func (m *mockConn) DoMulti(ctx context.Context, multi ...Completed) *redisresults {
	overrides := make([]RedisResult, 0, len(multi))
	for _, cmd := range multi {
		if fn := m.DoOverride[strings.Join(cmd.Commands(), " ")]; fn != nil {
			overrides = append(overrides, fn(cmd))
		}
	}
	if len(overrides) == len(multi) {
		return &redisresults{s: overrides}
	}
	if m.DoMultiFn != nil {
		return m.DoMultiFn(multi...)
	}
	return nil
}

func (m *mockConn) Receive(ctx context.Context, subscribe Completed, hdl func(message PubSubMessage)) error {
	if fn := m.ReceiveOverride[strings.Join(subscribe.Commands(), " ")]; fn != nil {
		return fn(ctx, subscribe, hdl)
	}
	if m.ReceiveFn != nil {
		return m.ReceiveFn(ctx, subscribe, hdl)
	}
	return nil
}

func (m *mockConn) DoStream(ctx context.Context, cmd Completed) RedisResultStream {
	if m.DoStreamFn != nil {
		return m.DoStreamFn(cmd)
	}
	return RedisResultStream{}
}

func (m *mockConn) DoMultiStream(ctx context.Context, cmd ...Completed) MultiRedisResultStream {
	if m.DoMultiStreamFn != nil {
		return m.DoMultiStreamFn(cmd...)
	}
	return MultiRedisResultStream{}
}

func (m *mockConn) CleanSubscriptions() {
	panic("not implemented")
}

func (m *mockConn) SetPubSubHooks(_ PubSubHooks) <-chan error {
	panic("not implemented")
}

func (m *mockConn) SetOnCloseHook(func(error)) {

}

func (m *mockConn) Info() map[string]RedisMessage {
	if m.InfoFn != nil {
		return m.InfoFn()
	}
	return nil
}

func (m *mockConn) Version() int {
	if m.VersionFn != nil {
		return m.VersionFn()
	}
	return 0
}

func (m *mockConn) AZ() string {
	if m.AZFn != nil {
		return m.AZFn()
	}
	return ""
}

func (m *mockConn) Error() error {
	if m.ErrorFn != nil {
		return m.ErrorFn()
	}
	return nil
}

func (m *mockConn) Close() {
	if m.CloseFn != nil {
		m.CloseFn()
	}
}

func (m *mockConn) Addr() string {
	if m.AddrFn != nil {
		return m.AddrFn()
	}
	return ""
}

func (m *mockConn) OptInCmd() cmds.Completed {
	return cmds.OptInCmd
}

func TestNewSingleClientNoNode(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	if _, err := newSingleClient(
		&ClientOption{}, nil, func(dst string, opt *ClientOption) conn {
			return nil
		}, newRetryer(defaultRetryDelayFn),
	); err != ErrNoAddr {
		t.Fatalf("unexpected err %v", err)
	}
}

func TestNewSingleClientReplicaOnlyNotSupported(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	if _, err := newSingleClient(
		&ClientOption{ReplicaOnly: true, InitAddress: []string{"localhost"}}, nil, func(dst string, opt *ClientOption) conn { return nil }, newRetryer(defaultRetryDelayFn),
	); err != ErrReplicaOnlyNotSupported {
		t.Fatalf("unexpected err %v", err)
	}
}

func TestNewSingleClientError(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	v := errors.New("dial err")
	if _, err := newSingleClient(
		&ClientOption{InitAddress: []string{""}}, nil, func(dst string, opt *ClientOption) conn { return &mockConn{DialFn: func() error { return v }} }, newRetryer(defaultRetryDelayFn),
	); err != v {
		t.Fatalf("unexpected err %v", err)
	}
}

func TestNewSingleClientOverride(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	m1 := &mockConn{}
	var m2 conn
	if _, err := newSingleClient(
		&ClientOption{InitAddress: []string{""}},
		m1,
		func(dst string, opt *ClientOption) conn {
			return &mockConn{OverrideFn: func(c conn) { m2 = c }}
		},
		newRetryer(defaultRetryDelayFn),
	); err != nil {
		t.Fatalf("unexpected err %v", err)
	}
	if m2.(*mockConn) != m1 {
		t.Fatalf("unexpected m2 %v", m2)
	}
}

//gocyclo:ignore
func TestSingleClient(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	m := &mockConn{
		AddrFn: func() string { return "myaddr" },
	}
	client, err := newSingleClient(
		&ClientOption{InitAddress: []string{""}},
		m,
		func(dst string, opt *ClientOption) conn { return m },
		newRetryer(defaultRetryDelayFn),
	)
	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}

	t.Run("Nodes", func(t *testing.T) {
		if nodes := client.Nodes(); len(nodes) != 1 || nodes["myaddr"] != client {
			t.Fatalf("unexpected nodes")
		}
	})

	t.Run("Mode", func(t *testing.T) {
		if v := client.Mode(); v != ClientModeStandalone {
			t.Fatalf("unexpected mode %v", v)
		}
	})

	t.Run("Delegate Do", func(t *testing.T) {
		c := client.B().Get().Key("Do").Build()
		m.DoFn = func(cmd Completed) RedisResult {
			if !reflect.DeepEqual(cmd.Commands(), c.Commands()) {
				t.Fatalf("unexpected command %v", cmd)
			}
			return newResult(strmsg('+', "Do"), nil)
		}
		if v, err := client.Do(context.Background(), c).ToString(); err != nil || v != "Do" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("Delegate DoStream", func(t *testing.T) {
		c := client.B().Get().Key("Do").Build()
		m.DoStreamFn = func(cmd Completed) RedisResultStream {
			return RedisResultStream{e: errors.New(cmd.Commands()[1])}
		}
		if s := client.DoStream(context.Background(), c); s.Error().Error() != "Do" {
			t.Fatalf("unexpected response %v", s.Error())
		}
	})

	t.Run("Delegate DoMulti", func(t *testing.T) {
		c := client.B().Get().Key("Do").Build()
		m.DoMultiFn = func(cmd ...Completed) *redisresults {
			if !reflect.DeepEqual(cmd[0].Commands(), c.Commands()) {
				t.Fatalf("unexpected command %v", cmd)
			}
			return &redisresults{s: []RedisResult{newResult(strmsg('+', "Do"), nil)}}
		}
		if len(client.DoMulti(context.Background())) != 0 {
			t.Fatalf("unexpected response length")
		}
		if v, err := client.DoMulti(context.Background(), c)[0].ToString(); err != nil || v != "Do" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("Delegate DoMultiStream", func(t *testing.T) {
		c := client.B().Get().Key("Do").Build()
		m.DoMultiStreamFn = func(cmd ...Completed) MultiRedisResultStream {
			return MultiRedisResultStream{e: errors.New(cmd[0].Commands()[1])}
		}
		if s := client.DoMultiStream(context.Background()); s.Error() != io.EOF {
			t.Fatalf("unexpected response %v", err)
		}
		if s := client.DoMultiStream(context.Background(), c); s.Error().Error() != "Do" {
			t.Fatalf("unexpected response %v", s.Error())
		}
	})

	t.Run("Delegate DoCache", func(t *testing.T) {
		c := client.B().Get().Key("DoCache").Cache()
		m.DoCacheFn = func(cmd Cacheable, ttl time.Duration) RedisResult {
			if !reflect.DeepEqual(cmd.Commands(), c.Commands()) || ttl != 100 {
				t.Fatalf("unexpected command %v, %v", cmd, ttl)
			}
			return newResult(strmsg('+', "DoCache"), nil)
		}
		if v, err := client.DoCache(context.Background(), c, 100).ToString(); err != nil || v != "DoCache" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("Delegate DoMultiCache", func(t *testing.T) {
		c := client.B().Get().Key("DoCache").Cache()
		m.DoMultiCacheFn = func(multi ...CacheableTTL) *redisresults {
			if !reflect.DeepEqual(multi[0].Cmd.Commands(), c.Commands()) || multi[0].TTL != 100 {
				t.Fatalf("unexpected command %v, %v", multi[0].Cmd, multi[0].TTL)
			}
			return &redisresults{s: []RedisResult{newResult(strmsg('+', "DoCache"), nil)}}
		}
		if len(client.DoMultiCache(context.Background())) != 0 {
			t.Fatalf("unexpected response length")
		}
		if v, err := client.DoMultiCache(context.Background(), CT(c, 100))[0].ToString(); err != nil || v != "DoCache" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("Delegate Receive", func(t *testing.T) {
		c := client.B().Subscribe().Channel("ch").Build()
		hdl := func(message PubSubMessage) {}
		m.ReceiveFn = func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
			if !reflect.DeepEqual(subscribe.Commands(), c.Commands()) {
				t.Fatalf("unexpected command %v", subscribe)
			}
			return nil
		}
		if err := client.Receive(context.Background(), c, hdl); err != nil {
			t.Fatalf("unexpected response %v", err)
		}
	})

	t.Run("Delegate Receive Redis Err", func(t *testing.T) {
		c := client.B().Subscribe().Channel("ch").Build()
		e := &RedisError{}
		m.ReceiveFn = func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
			return e
		}
		if err := client.Receive(context.Background(), c, func(message PubSubMessage) {}); err != e {
			t.Fatalf("unexpected response %v", err)
		}
	})

	t.Run("Delegate Close", func(t *testing.T) {
		called := false
		m.CloseFn = func() { called = true }
		client.Close()
		if !called {
			t.Fatalf("Close is not delegated")
		}
	})

	t.Run("Dedicated Err", func(t *testing.T) {
		v := errors.New("fn err")
		if err := client.Dedicated(func(client DedicatedClient) error {
			return v
		}); err != v {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("Dedicated Delegate Receive Redis Err", func(t *testing.T) {
		e := &RedisError{}
		w := &mockWire{
			ReceiveFn: func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
				return e
			},
		}
		m.AcquireFn = func() wire {
			return w
		}
		if err := client.Dedicated(func(c DedicatedClient) error {
			return c.Receive(context.Background(), c.B().Subscribe().Channel("a").Build(), func(msg PubSubMessage) {})
		}); err != e {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("Dedicated Delegate", func(t *testing.T) {
		closed := false
		w := &mockWire{
			DoFn: func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "Delegate"), nil)
			},
			DoMultiFn: func(cmd ...Completed) *redisresults {
				return &redisresults{s: []RedisResult{newResult(strmsg('+', "Delegate"), nil)}}
			},
			ReceiveFn: func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
				return ErrClosing
			},
			SetPubSubHooksFn: func(hooks PubSubHooks) <-chan error {
				ch := make(chan error, 1)
				ch <- ErrClosing
				close(ch)
				return ch
			},
			ErrorFn: func() error {
				return ErrClosing
			},
			CloseFn: func() {
				closed = true
			},
		}
		m.AcquireFn = func() wire {
			return w
		}
		stored := false
		m.StoreFn = func(ww wire) {
			if ww != w {
				t.Fatalf("received unexpected wire %v", ww)
			}
			stored = true
		}
		if err := client.Dedicated(func(c DedicatedClient) error {
			if v, err := c.Do(context.Background(), c.B().Get().Key("a").Build()).ToString(); err != nil || v != "Delegate" {
				t.Fatalf("unexpected response %v %v", v, err)
			}
			if v := c.DoMulti(context.Background()); len(v) != 0 {
				t.Fatalf("received unexpected response %v", v)
			}
			for _, resp := range c.DoMulti(context.Background(), c.B().Get().Key("a").Build()) {
				if v, err := resp.ToString(); err != nil || v != "Delegate" {
					t.Fatalf("unexpected response %v %v", v, err)
				}
			}
			if err := c.Receive(context.Background(), c.B().Subscribe().Channel("a").Build(), func(msg PubSubMessage) {}); err != ErrClosing {
				t.Fatalf("unexpected ret %v", err)
			}
			if err := <-c.SetPubSubHooks(PubSubHooks{}); err != ErrClosing {
				t.Fatalf("unexpected ret %v", err)
			}
			c.Close()
			return nil
		}); err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if !stored {
			t.Fatalf("Dedicated doesn't put back the wire")
		}
		if !closed {
			t.Fatalf("Dedicated doesn't delegate Close")
		}
	})

	t.Run("Dedicate Delegate", func(t *testing.T) {
		closed := false
		w := &mockWire{
			DoFn: func(cmd Completed) RedisResult {
				return newResult(strmsg('+', "Delegate"), nil)
			},
			DoMultiFn: func(cmd ...Completed) *redisresults {
				return &redisresults{s: []RedisResult{newResult(strmsg('+', "Delegate"), nil)}}
			},
			ReceiveFn: func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
				return ErrClosing
			},
			SetPubSubHooksFn: func(hooks PubSubHooks) <-chan error {
				ch := make(chan error, 1)
				ch <- ErrClosing
				close(ch)
				return ch
			},
			ErrorFn: func() error {
				return ErrClosing
			},
			CloseFn: func() {
				closed = true
			},
		}
		m.AcquireFn = func() wire {
			return w
		}
		stored := false
		m.StoreFn = func(ww wire) {
			if ww != w {
				t.Fatalf("received unexpected wire %v", ww)
			}
			stored = true
		}
		c, cancel := client.Dedicate()

		if v, err := c.Do(context.Background(), c.B().Get().Key("a").Build()).ToString(); err != nil || v != "Delegate" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
		if v := c.DoMulti(context.Background()); len(v) != 0 {
			t.Fatalf("received unexpected response %v", v)
		}
		for _, resp := range c.DoMulti(context.Background(), c.B().Get().Key("a").Build()) {
			if v, err := resp.ToString(); err != nil || v != "Delegate" {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		}
		if err := c.Receive(context.Background(), c.B().Subscribe().Channel("a").Build(), func(msg PubSubMessage) {}); err != ErrClosing {
			t.Fatalf("unexpected ret %v", err)
		}
		if err := <-c.SetPubSubHooks(PubSubHooks{}); err != ErrClosing {
			t.Fatalf("unexpected ret %v", err)
		}
		c.Close()

		cancel()

		if !stored {
			t.Fatalf("Dedicated doesn't put back the wire")
		}
		if !closed {
			t.Fatalf("Dedicated doesn't delegate Close")
		}
	})

	t.Run("Dedicate Delegate Release On Close", func(t *testing.T) {
		stored := 0
		w := &mockWire{}
		m.AcquireFn = func() wire { return w }
		m.StoreFn = func(ww wire) { stored++ }
		c, _ := client.Dedicate()

		c.Close()

		if stored != 1 {
			t.Fatalf("unexpected stored count %v", stored)
		}
	})

	t.Run("Dedicate Delegate No Duplicate Release", func(t *testing.T) {
		stored := 0
		w := &mockWire{}
		m.AcquireFn = func() wire { return w }
		m.StoreFn = func(ww wire) { stored++ }
		c, cancel := client.Dedicate()

		c.Close()
		c.Close() // should have no effect
		cancel()  // should have no effect
		cancel()  // should have no effect

		if stored != 1 {
			t.Fatalf("unexpected stored count %v", stored)
		}
	})

	t.Run("Dedicate ErrDedicatedClientRecycled after released", func(t *testing.T) {
		m.AcquireFn = func() wire { return &mockWire{} }
		check := func(err error) {
			if !errors.Is(err, ErrDedicatedClientRecycled) {
				t.Fatalf("unexpected err %v", err)
			}
		}
		for _, closeFn := range []func(client DedicatedClient, cancel func()){
			func(client DedicatedClient, cancel func()) {
				client.Close()
			},
			func(client DedicatedClient, cancel func()) {
				cancel()
			},
		} {
			c, cancel := client.Dedicate()
			closeFn(c, cancel)
			for _, fn := range []func(){
				func() {
					resp := c.Do(context.Background(), c.B().Get().Key("k").Build())
					check(resp.Error())
				},
				func() {
					resp := c.DoMulti(context.Background(), c.B().Get().Key("k").Build())
					for _, r := range resp {
						check(r.Error())
					}
				},
				func() {
					err := c.Receive(context.Background(), c.B().Subscribe().Channel("k").Build(), func(msg PubSubMessage) {})
					check(err)
				},
				func() {
					ch := c.SetPubSubHooks(PubSubHooks{})
					check(<-ch)
				},
			} {
				fn()
			}
		}
	})
}

func TestSingleClientRetry(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	SetupClientRetry(t, func(m *mockConn) Client {
		c, err := newSingleClient(
			&ClientOption{InitAddress: []string{""}},
			m,
			func(dst string, opt *ClientOption) conn { return m },
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		return c
	})
}

//gocyclo:ignore
func SetupClientRetry(t *testing.T, fn func(mock *mockConn) Client) {
	setup := func() (Client, *mockConn) {
		m := &mockConn{}
		return fn(m), m
	}

	makeDoFn := func(results ...RedisResult) func(cmd Completed) RedisResult {
		count := -1
		return func(cmd Completed) RedisResult {
			count++
			return results[count]
		}
	}

	makeDoCacheFn := func(results ...RedisResult) func(cmd Cacheable, ttl time.Duration) RedisResult {
		count := -1
		return func(cmd Cacheable, ttl time.Duration) RedisResult {
			count++
			return results[count]
		}
	}

	makeReceiveFn := func(results ...error) func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
		count := -1
		return func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
			count++
			return results[count]
		}
	}

	makeDoMultiFn := func(results ...[]RedisResult) func(multi ...Completed) *redisresults {
		count := -1
		return func(multi ...Completed) *redisresults {
			count++
			return &redisresults{s: results[count]}
		}
	}

	makeDoMultiCacheFn := func(results ...[]RedisResult) func(multi ...CacheableTTL) *redisresults {
		count := -1
		return func(multi ...CacheableTTL) *redisresults {
			count++
			return &redisresults{s: results[count]}
		}
	}

	t.Run("Delegate Do ReadOnly Retry", func(t *testing.T) {
		c, m := setup()
		m.DoFn = makeDoFn(
			newErrResult(ErrClosing),
			newResult(strmsg('+', "Do"), nil),
		)
		if v, err := c.Do(context.Background(), c.B().Get().Key("Do").Build()).ToString(); err != nil || v != "Do" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("Delegate Do ReadOnly NoRetry - closed", func(t *testing.T) {
		c, m := setup()
		m.DoFn = makeDoFn(newErrResult(ErrClosing))
		c.Close()
		if v, err := c.Do(context.Background(), c.B().Get().Key("Do").Build()).ToString(); err != ErrClosing {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("Delegate Do ReadOnly NoRetry - ctx done", func(t *testing.T) {
		c, m := setup()
		m.DoFn = makeDoFn(newErrResult(ErrClosing))
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		if v, err := c.Do(ctx, c.B().Get().Key("Do").Build()).ToString(); err != ErrClosing {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("Delegate Do ReadOnly NoRetry - not retryable", func(t *testing.T) {
		c, m := setup()
		if cli, ok := c.(*sentinelClient); ok {
			cli.retryHandler = &mockRetryHandler{
				WaitOrSkipRetryFunc: func(ctx context.Context, attempts int, _ Completed, err error) bool {
					return false
				},
			}
		}
		if cli, ok := c.(*clusterClient); ok {
			cli.retryHandler = &mockRetryHandler{
				WaitOrSkipRetryFunc: func(ctx context.Context, attempts int, _ Completed, err error) bool {
					return false
				},
			}
		}
		if cli, ok := c.(*singleClient); ok {
			cli.retryHandler = &mockRetryHandler{
				WaitOrSkipRetryFunc: func(ctx context.Context, attempts int, _ Completed, err error) bool {
					return false
				},
			}
		}
		m.DoFn = makeDoFn(
			newErrResult(ErrClosing),
			newResult(strmsg('+', "Do"), nil),
		)
		if v, err := c.Do(context.Background(), c.B().Get().Key("Do").Build()).ToString(); !errors.Is(err, ErrClosing) {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("Delegate Do Write NoRetry", func(t *testing.T) {
		c, m := setup()
		m.DoFn = makeDoFn(newErrResult(ErrClosing))
		if v, err := c.Do(context.Background(), c.B().Set().Key("Do").Value("V").Build()).ToString(); err != ErrClosing {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("Delegate DoMulti ReadOnly Retry", func(t *testing.T) {
		c, m := setup()
		m.DoMultiFn = makeDoMultiFn(
			[]RedisResult{newErrResult(ErrClosing)},
			[]RedisResult{newResult(strmsg('+', "Do"), nil)},
		)
		if v, err := c.DoMulti(context.Background(), c.B().Get().Key("Do").Build())[0].ToString(); err != nil || v != "Do" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("Delegate DoMulti ReadOnly NoRetry - closed", func(t *testing.T) {
		c, m := setup()
		m.DoMultiFn = makeDoMultiFn([]RedisResult{newErrResult(ErrClosing)})
		c.Close()
		if v, err := c.DoMulti(context.Background(), c.B().Get().Key("Do").Build())[0].ToString(); err != ErrClosing {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("Delegate DoMulti ReadOnly NoRetry - ctx done", func(t *testing.T) {
		c, m := setup()
		m.DoMultiFn = makeDoMultiFn([]RedisResult{newErrResult(ErrClosing)})
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		if v, err := c.DoMulti(ctx, c.B().Get().Key("Do").Build())[0].ToString(); err != ErrClosing {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("Delegate DoMulti ReadOnly NoRetry - not retryable", func(t *testing.T) {
		c, m := setup()
		if cli, ok := c.(*sentinelClient); ok {
			cli.retryHandler = &mockRetryHandler{
				WaitOrSkipRetryFunc: func(ctx context.Context, attempts int, _ Completed, err error) bool {
					return false
				},
			}
		}
		if cli, ok := c.(*clusterClient); ok {
			cli.retryHandler = &mockRetryHandler{
				WaitOrSkipRetryFunc: func(ctx context.Context, attempts int, _ Completed, err error) bool {
					return false
				},
				RetryDelayFn: func(attempts int, _ Completed, err error) time.Duration {
					return -1
				},
				WaitForRetryFn: func(ctx context.Context, duration time.Duration) {
					if duration != -1 {
						t.Fatalf("unexpected duration %v", duration)
					}
				},
			}
		}
		if cli, ok := c.(*singleClient); ok {
			cli.retryHandler = &mockRetryHandler{
				WaitOrSkipRetryFunc: func(ctx context.Context, attempts int, _ Completed, err error) bool {
					return false
				},
			}
		}
		m.DoMultiFn = makeDoMultiFn(
			[]RedisResult{newErrResult(ErrClosing)},
			[]RedisResult{newResult(strmsg('+', "Do"), nil)},
		)
		if v, err := c.DoMulti(context.Background(), c.B().Get().Key("Do").Build())[0].ToString(); !errors.Is(err, ErrClosing) {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("Delegate DoMulti Write NoRetry", func(t *testing.T) {
		c, m := setup()
		m.DoMultiFn = makeDoMultiFn([]RedisResult{newErrResult(ErrClosing)})
		if v, err := c.DoMulti(context.Background(), c.B().Set().Key("Do").Value("V").Build())[0].ToString(); err != ErrClosing {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("Delegate DoCache Retry", func(t *testing.T) {
		c, m := setup()
		m.DoCacheFn = makeDoCacheFn(
			newErrResult(ErrClosing),
			newResult(strmsg('+', "Do"), nil),
		)
		if v, err := c.DoCache(context.Background(), c.B().Get().Key("Do").Cache(), 0).ToString(); err != nil || v != "Do" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("Delegate DoCache NoRetry - closed", func(t *testing.T) {
		c, m := setup()
		m.DoCacheFn = makeDoCacheFn(newErrResult(ErrClosing))
		c.Close()
		if v, err := c.DoCache(context.Background(), c.B().Get().Key("Do").Cache(), 0).ToString(); err != ErrClosing {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("Delegate DoCache ReadOnly NoRetry - ctx done", func(t *testing.T) {
		c, m := setup()
		m.DoCacheFn = makeDoCacheFn(newErrResult(ErrClosing))
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		if v, err := c.DoCache(ctx, c.B().Get().Key("Do").Cache(), 0).ToString(); err != ErrClosing {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("Delegate DoCache ReadOnly NoRetry - not retryable", func(t *testing.T) {
		c, m := setup()
		if cli, ok := c.(*sentinelClient); ok {
			cli.retryHandler = &mockRetryHandler{
				WaitOrSkipRetryFunc: func(ctx context.Context, attempts int, _ Completed, err error) bool {
					return false
				},
			}
		}
		if cli, ok := c.(*clusterClient); ok {
			cli.retryHandler = &mockRetryHandler{
				WaitOrSkipRetryFunc: func(ctx context.Context, attempts int, _ Completed, err error) bool {
					return false
				},
			}
		}
		if cli, ok := c.(*singleClient); ok {
			cli.retryHandler = &mockRetryHandler{
				WaitOrSkipRetryFunc: func(ctx context.Context, attempts int, _ Completed, err error) bool {
					return false
				},
			}
		}
		m.DoCacheFn = makeDoCacheFn(
			newErrResult(ErrClosing),
			newResult(strmsg('+', "Do"), nil),
		)
		if v, err := c.DoCache(context.Background(), c.B().Get().Key("Do").Cache(), 0).ToString(); !errors.Is(err, ErrClosing) {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("Delegate DoMultiCache Retry", func(t *testing.T) {
		c, m := setup()
		m.DoMultiCacheFn = makeDoMultiCacheFn(
			[]RedisResult{newErrResult(ErrClosing)},
			[]RedisResult{newResult(strmsg('+', "Do"), nil)},
		)
		if v, err := c.DoMultiCache(context.Background(), CT(c.B().Get().Key("Do").Cache(), 0))[0].ToString(); err != nil || v != "Do" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("Delegate DoMultiCache NoRetry - closed", func(t *testing.T) {
		c, m := setup()
		m.DoMultiCacheFn = makeDoMultiCacheFn([]RedisResult{newErrResult(ErrClosing)})
		c.Close()
		if v, err := c.DoMultiCache(context.Background(), CT(c.B().Get().Key("Do").Cache(), 0))[0].ToString(); err != ErrClosing {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("Delegate DoMultiCache ReadOnly NoRetry - ctx done", func(t *testing.T) {
		c, m := setup()
		m.DoMultiCacheFn = makeDoMultiCacheFn([]RedisResult{newErrResult(ErrClosing)})
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		if v, err := c.DoMultiCache(ctx, CT(c.B().Get().Key("Do").Cache(), 0))[0].ToString(); err != ErrClosing {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("Delegate DoMultiCache ReadOnly NoRetry - not retryable", func(t *testing.T) {
		c, m := setup()
		if cli, ok := c.(*sentinelClient); ok {
			cli.retryHandler = &mockRetryHandler{
				WaitOrSkipRetryFunc: func(ctx context.Context, attempts int, _ Completed, err error) bool {
					return false
				},
			}
		}
		if cli, ok := c.(*clusterClient); ok {
			cli.retryHandler = &mockRetryHandler{
				WaitOrSkipRetryFunc: func(ctx context.Context, attempts int, _ Completed, err error) bool {
					return false
				},
				RetryDelayFn: func(attempts int, _ Completed, err error) time.Duration {
					return -1
				},
				WaitForRetryFn: func(ctx context.Context, duration time.Duration) {
					if duration != -1 {
						t.Fatalf("unexpected duration %v", duration)
					}
				},
			}
		}
		if cli, ok := c.(*singleClient); ok {
			cli.retryHandler = &mockRetryHandler{
				WaitOrSkipRetryFunc: func(ctx context.Context, attempts int, _ Completed, err error) bool {
					return false
				},
			}
		}
		m.DoMultiCacheFn = makeDoMultiCacheFn(
			[]RedisResult{newErrResult(ErrClosing)},
			[]RedisResult{newResult(strmsg('+', "Do"), nil)},
		)
		if v, err := c.DoMultiCache(context.Background(), CT(c.B().Get().Key("Do").Cache(), 0))[0].ToString(); !errors.Is(err, ErrClosing) {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("Delegate Receive Retry", func(t *testing.T) {
		c, m := setup()
		m.ReceiveFn = makeReceiveFn(ErrClosing, nil)
		if err := c.Receive(context.Background(), c.B().Subscribe().Channel("ch").Build(), nil); err != nil {
			t.Fatalf("unexpected response %v", err)
		}
	})

	t.Run("Delegate Receive NoRetry - closed", func(t *testing.T) {
		c, m := setup()
		m.ReceiveFn = makeReceiveFn(ErrClosing)
		c.Close()
		if err := c.Receive(context.Background(), c.B().Subscribe().Channel("ch").Build(), nil); err != ErrClosing {
			t.Fatalf("unexpected response %v", err)
		}
	})

	t.Run("Delegate Receive NoRetry - ctx done", func(t *testing.T) {
		c, m := setup()
		m.ReceiveFn = makeReceiveFn(ErrClosing)
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		if err := c.Receive(ctx, c.B().Subscribe().Channel("ch").Build(), nil); err != ErrClosing {
			t.Fatalf("unexpected response %v", err)
		}
	})

	t.Run("Delegate Receive NoRetry - not retryable", func(t *testing.T) {
		c, m := setup()
		if cli, ok := c.(*sentinelClient); ok {
			cli.retryHandler = &mockRetryHandler{
				WaitOrSkipRetryFunc: func(ctx context.Context, attempts int, _ Completed, err error) bool {
					return false
				},
			}
		}
		if cli, ok := c.(*clusterClient); ok {
			cli.retryHandler = &mockRetryHandler{
				WaitOrSkipRetryFunc: func(ctx context.Context, attempts int, _ Completed, err error) bool {
					return false
				},
			}
		}
		if cli, ok := c.(*singleClient); ok {
			cli.retryHandler = &mockRetryHandler{
				WaitOrSkipRetryFunc: func(ctx context.Context, attempts int, _ Completed, err error) bool {
					return false
				},
			}
		}
		m.ReceiveFn = makeReceiveFn(ErrClosing, nil)
		if err := c.Receive(context.Background(), c.B().Subscribe().Channel("ch").Build(), nil); !errors.Is(err, ErrClosing) {
			t.Fatalf("unexpected response %v", err)
		}
	})

	t.Run("Dedicate Delegate Do ReadOnly Retry", func(t *testing.T) {
		c, m := setup()
		m.DoFn = makeDoFn(
			newErrResult(ErrClosing),
			newResult(strmsg('+', "Do"), nil),
		)
		m.AcquireFn = func() wire { return &mockWire{DoFn: m.DoFn} }
		if ret := c.Dedicated(func(cc DedicatedClient) error {
			if v, err := cc.Do(context.Background(), c.B().Get().Key("Do").Build()).ToString(); err != nil || v != "Do" {
				t.Fatalf("unexpected response %v %v", v, err)
			}
			return errors.New("done")
		}); ret.Error() != "done" {
			t.Fatalf("Dedicated not executed")
		}
	})

	t.Run("Dedicate Delegate Do ReadOnly NoRetry - broken", func(t *testing.T) {
		c, m := setup()
		m.DoFn = makeDoFn(newErrResult(ErrClosing))
		m.ErrorFn = func() error { return ErrClosing }
		m.AcquireFn = func() wire { return &mockWire{DoFn: m.DoFn, ErrorFn: m.ErrorFn} }
		if ret := c.Dedicated(func(cc DedicatedClient) error {
			return cc.Do(context.Background(), c.B().Get().Key("Do").Build()).Error()
		}); ret != ErrClosing {
			t.Fatalf("unexpected response %v", ret)
		}
	})

	t.Run("Dedicate Delegate Do ReadOnly NoRetry - ctx done", func(t *testing.T) {
		c, m := setup()
		m.DoFn = makeDoFn(newErrResult(ErrClosing))
		m.AcquireFn = func() wire { return &mockWire{DoFn: m.DoFn} }
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		if ret := c.Dedicated(func(cc DedicatedClient) error {
			return cc.Do(ctx, c.B().Get().Key("Do").Build()).Error()
		}); ret != ErrClosing {
			t.Fatalf("unexpected response %v", ret)
		}
	})

	t.Run("Dedicate Delegate Do ReadOnly NoRetry - not retryable", func(t *testing.T) {
		c, m := setup()
		m.DoFn = makeDoFn(
			newErrResult(ErrClosing),
			newResult(strmsg('+', "Do"), nil),
		)
		m.AcquireFn = func() wire { return &mockWire{DoFn: m.DoFn} }
		if ret := c.Dedicated(func(cc DedicatedClient) error {
			if cli, ok := cc.(*dedicatedClusterClient); ok {
				cli.retryHandler = &mockRetryHandler{
					WaitOrSkipRetryFunc: func(ctx context.Context, attempts int, _ Completed, err error) bool {
						return false
					},
				}
			}
			if cli, ok := cc.(*dedicatedSingleClient); ok {
				cli.retryHandler = &mockRetryHandler{
					WaitOrSkipRetryFunc: func(ctx context.Context, attempts int, _ Completed, err error) bool {
						return false
					},
				}
			}

			if v, err := cc.Do(context.Background(), c.B().Get().Key("Do").Build()).ToString(); !errors.Is(err, ErrClosing) {
				t.Fatalf("unexpected response %v %v", v, err)
			}
			return errors.New("done")
		}); ret.Error() != "done" {
			t.Fatalf("Dedicated not executed")
		}
	})

	t.Run("Dedicate Delegate Do Write NoRetry", func(t *testing.T) {
		c, m := setup()
		m.DoFn = makeDoFn(newErrResult(ErrClosing))
		m.AcquireFn = func() wire { return &mockWire{DoFn: m.DoFn} }
		if ret := c.Dedicated(func(cc DedicatedClient) error {
			return cc.Do(context.Background(), c.B().Set().Key("Do").Value("Do").Build()).Error()
		}); ret != ErrClosing {
			t.Fatalf("unexpected response %v", ret)
		}
	})

	t.Run("Dedicate Delegate DoMulti ReadOnly Retry", func(t *testing.T) {
		c, m := setup()
		m.DoMultiFn = makeDoMultiFn(
			[]RedisResult{newErrResult(ErrClosing)},
			[]RedisResult{newResult(strmsg('+', "Do"), nil)},
		)
		m.AcquireFn = func() wire { return &mockWire{DoMultiFn: m.DoMultiFn} }
		if ret := c.Dedicated(func(cc DedicatedClient) error {
			if v, err := cc.DoMulti(context.Background(), c.B().Get().Key("Do").Build())[0].ToString(); err != nil || v != "Do" {
				t.Fatalf("unexpected response %v %v", v, err)
			}
			return errors.New("done")
		}); ret.Error() != "done" {
			t.Fatalf("Dedicated not executed")
		}
	})

	t.Run("Dedicate Delegate DoMulti ReadOnly NoRetry - broken", func(t *testing.T) {
		c, m := setup()
		m.DoMultiFn = makeDoMultiFn([]RedisResult{newErrResult(ErrClosing)})
		m.ErrorFn = func() error { return ErrClosing }
		m.AcquireFn = func() wire { return &mockWire{DoMultiFn: m.DoMultiFn, ErrorFn: m.ErrorFn} }
		if ret := c.Dedicated(func(cc DedicatedClient) error {
			return cc.DoMulti(context.Background(), c.B().Get().Key("Do").Build())[0].Error()
		}); ret != ErrClosing {
			t.Fatalf("unexpected response %v", ret)
		}
	})

	t.Run("Dedicate Delegate DoMulti ReadOnly NoRetry - ctx done", func(t *testing.T) {
		c, m := setup()
		m.DoMultiFn = makeDoMultiFn([]RedisResult{newErrResult(ErrClosing)})
		m.AcquireFn = func() wire { return &mockWire{DoMultiFn: m.DoMultiFn} }
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		if ret := c.Dedicated(func(cc DedicatedClient) error {
			return cc.DoMulti(ctx, c.B().Get().Key("Do").Build())[0].Error()
		}); ret != ErrClosing {
			t.Fatalf("unexpected response %v", ret)
		}
	})

	t.Run("Dedicate Delegate DoMulti ReadOnly NoRetry - not retryable", func(t *testing.T) {
		c, m := setup()
		m.DoMultiFn = makeDoMultiFn(
			[]RedisResult{newErrResult(ErrClosing)},
			[]RedisResult{newResult(strmsg('+', "Do"), nil)},
		)
		m.AcquireFn = func() wire { return &mockWire{DoMultiFn: m.DoMultiFn} }
		if ret := c.Dedicated(func(cc DedicatedClient) error {
			if cli, ok := cc.(*dedicatedClusterClient); ok {
				cli.retryHandler = &mockRetryHandler{
					WaitOrSkipRetryFunc: func(ctx context.Context, attempts int, _ Completed, err error) bool {
						return false
					},
				}
			}
			if cli, ok := cc.(*dedicatedSingleClient); ok {
				cli.retryHandler = &mockRetryHandler{
					WaitOrSkipRetryFunc: func(ctx context.Context, attempts int, _ Completed, err error) bool {
						return false
					},
				}
			}

			if v, err := cc.DoMulti(context.Background(), c.B().Get().Key("Do").Build())[0].ToString(); !errors.Is(err, ErrClosing) {
				t.Fatalf("unexpected response %v %v", v, err)
			}
			return errors.New("done")
		}); ret.Error() != "done" {
			t.Fatalf("Dedicated not executed")
		}
	})

	t.Run("Dedicate Delegate DoMulti Write NoRetry", func(t *testing.T) {
		c, m := setup()
		m.DoMultiFn = makeDoMultiFn([]RedisResult{newErrResult(ErrClosing)})
		m.AcquireFn = func() wire { return &mockWire{DoMultiFn: m.DoMultiFn} }
		if ret := c.Dedicated(func(cc DedicatedClient) error {
			return cc.DoMulti(context.Background(), c.B().Set().Key("Do").Value("Do").Build())[0].Error()
		}); ret != ErrClosing {
			t.Fatalf("unexpected response %v", ret)
		}
	})

	t.Run("Delegate Receive Retry", func(t *testing.T) {
		c, m := setup()
		m.ReceiveFn = makeReceiveFn(ErrClosing, nil)
		m.AcquireFn = func() wire { return &mockWire{ReceiveFn: m.ReceiveFn} }
		if ret := c.Dedicated(func(cc DedicatedClient) error {
			if err := cc.Receive(context.Background(), c.B().Subscribe().Channel("Do").Build(), nil); err != nil {
				t.Fatalf("unexpected response %v", err)
			}
			return errors.New("done")
		}); ret.Error() != "done" {
			t.Fatalf("Dedicated not executed")
		}
	})

	t.Run("Delegate Receive NoRetry - broken", func(t *testing.T) {
		c, m := setup()
		m.ReceiveFn = makeReceiveFn(ErrClosing)
		m.ErrorFn = func() error { return ErrClosing }
		m.AcquireFn = func() wire { return &mockWire{ReceiveFn: m.ReceiveFn, ErrorFn: m.ErrorFn} }
		if ret := c.Dedicated(func(cc DedicatedClient) error {
			return cc.Receive(context.Background(), c.B().Subscribe().Channel("Do").Build(), nil)
		}); ret != ErrClosing {
			t.Fatalf("unexpected response %v", ret)
		}
	})

	t.Run("Delegate Receive NoRetry - ctx done", func(t *testing.T) {
		c, m := setup()
		m.ReceiveFn = makeReceiveFn(ErrClosing)
		m.AcquireFn = func() wire { return &mockWire{ReceiveFn: m.ReceiveFn} }
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		if ret := c.Dedicated(func(cc DedicatedClient) error {
			return cc.Receive(ctx, c.B().Subscribe().Channel("Do").Build(), nil)
		}); ret != ErrClosing {
			t.Fatalf("unexpected response %v", ret)
		}
	})

	t.Run("Delegate Receive NoRetry - not retryable", func(t *testing.T) {
		c, m := setup()
		m.ReceiveFn = makeReceiveFn(ErrClosing, nil)
		m.AcquireFn = func() wire { return &mockWire{ReceiveFn: m.ReceiveFn} }
		if ret := c.Dedicated(func(cc DedicatedClient) error {
			if cli, ok := cc.(*dedicatedClusterClient); ok {
				cli.retryHandler = &mockRetryHandler{
					WaitOrSkipRetryFunc: func(ctx context.Context, attempts int, _ Completed, err error) bool {
						return false
					},
				}
			}
			if cli, ok := cc.(*dedicatedSingleClient); ok {
				cli.retryHandler = &mockRetryHandler{
					WaitOrSkipRetryFunc: func(ctx context.Context, attempts int, _ Completed, err error) bool {
						return false
					},
				}
			}

			if err := cc.Receive(context.Background(), c.B().Subscribe().Channel("Do").Build(), nil); !errors.Is(err, ErrClosing) {
				t.Fatalf("unexpected response %v", err)
			}
			return errors.New("done")
		}); ret.Error() != "done" {
			t.Fatalf("Dedicated not executed")
		}
	})
}

func TestSingleClientLoadingRetry(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())

	setup := func() (*singleClient, *mockConn) {
		m := &mockConn{}
		client, err := newSingleClient(
			&ClientOption{InitAddress: []string{""}},
			m,
			func(dst string, opt *ClientOption) conn { return m },
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		return client, m
	}

	t.Run("Do Retry on Loading", func(t *testing.T) {
		client, m := setup()
		attempts := 0
		m.DoFn = func(cmd Completed) RedisResult {
			attempts++
			if attempts == 1 {
				return newResult(strmsg('-', "LOADING Redis is loading the dataset in memory"), nil)
			}
			return newResult(strmsg('+', "OK"), nil)
		}

		if v, err := client.Do(context.Background(), client.B().Get().Key("test").Build()).ToString(); err != nil || v != "OK" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
		if attempts != 2 {
			t.Fatalf("expected 2 attempts, got %v", attempts)
		}
	})

	t.Run("Do not retry on non-loading errors", func(t *testing.T) {
		client, m := setup()
		attempts := 0
		m.DoFn = func(cmd Completed) RedisResult {
			attempts++
			if attempts == 1 {
				return newResult(strmsg('-', "ERR some other error"), nil)
			}
			return newResult(strmsg('+', "OK"), nil)
		}

		if err := client.Do(context.Background(), client.B().Get().Key("test").Build()).Error(); err == nil {
			t.Fatal("expected error but got nil")
		}
		if attempts != 1 {
			t.Fatalf("unexpected attempts %v, expected no retry", attempts)
		}
	})

	t.Run("DoMulti Retry on Loading", func(t *testing.T) {
		client, m := setup()
		attempts := 0
		m.DoMultiFn = func(multi ...Completed) *redisresults {
			attempts++
			if attempts == 1 {
				return &redisresults{s: []RedisResult{newResult(strmsg('-', "LOADING Redis is loading the dataset in memory"), nil)}}
			}
			return &redisresults{s: []RedisResult{newResult(strmsg('+', "OK"), nil)}}
		}

		cmd := client.B().Get().Key("test").Build()
		resps := client.DoMulti(context.Background(), cmd)
		if len(resps) != 1 {
			t.Fatalf("unexpected response length %v", len(resps))
		}
		if v, err := resps[0].ToString(); err != nil || v != "OK" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("DoCache Retry on Loading", func(t *testing.T) {
		client, m := setup()
		attempts := 0
		m.DoCacheFn = func(cmd Cacheable, ttl time.Duration) RedisResult {
			attempts++
			if attempts == 1 {
				return newResult(strmsg('-', "LOADING Redis is loading the dataset in memory"), nil)
			}
			return newResult(strmsg('+', "OK"), nil)
		}

		cmd := client.B().Get().Key("test").Cache()
		if v, err := client.DoCache(context.Background(), cmd, time.Minute).ToString(); err != nil || v != "OK" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("DoMultiCache Retry on Loading", func(t *testing.T) {
		client, m := setup()
		attempts := 0
		m.DoMultiCacheFn = func(multi ...CacheableTTL) *redisresults {
			attempts++
			if attempts == 1 {
				return &redisresults{s: []RedisResult{newResult(strmsg('-', "LOADING Redis is loading the dataset in memory"), nil)}}
			}
			return &redisresults{s: []RedisResult{newResult(strmsg('+', "OK"), nil)}}
		}

		cmd := client.B().Get().Key("test").Cache()
		resps := client.DoMultiCache(context.Background(), CT(cmd, time.Minute))
		if len(resps) != 1 {
			t.Fatalf("unexpected response length %v", len(resps))
		}
		if v, err := resps[0].ToString(); err != nil || v != "OK" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("Dedicated Do Retry on Loading", func(t *testing.T) {
		client, m := setup()
		attempts := 0
		m.DoFn = func(cmd Completed) RedisResult {
			attempts++
			if attempts == 1 {
				return newResult(strmsg('-', "LOADING Redis is loading the dataset in memory"), nil)
			}
			return newResult(strmsg('+', "OK"), nil)
		}
		m.AcquireFn = func() wire { return &mockWire{DoFn: m.DoFn} }

		err := client.Dedicated(func(c DedicatedClient) error {
			if v, err := c.Do(context.Background(), c.B().Get().Key("test").Build()).ToString(); err != nil || v != "OK" {
				t.Fatalf("unexpected response %v %v", v, err)
			}
			return nil
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("Dedicated DoMulti Retry on Loading", func(t *testing.T) {
		client, m := setup()
		attempts := 0
		m.DoMultiFn = func(multi ...Completed) *redisresults {
			attempts++
			if attempts == 1 {
				return &redisresults{s: []RedisResult{newResult(strmsg('-', "LOADING Redis is loading the dataset in memory"), nil)}}
			}
			return &redisresults{s: []RedisResult{newResult(strmsg('+', "OK"), nil)}}
		}
		m.AcquireFn = func() wire { return &mockWire{DoMultiFn: m.DoMultiFn} }

		err := client.Dedicated(func(c DedicatedClient) error {
			resps := c.DoMulti(context.Background(), c.B().Get().Key("test").Build())
			if len(resps) != 1 {
				t.Fatalf("unexpected response length %v", len(resps))
			}
			if v, err := resps[0].ToString(); err != nil || v != "OK" {
				t.Fatalf("unexpected response %v %v", v, err)
			}
			return nil
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
	})
}

func TestSingleClientConnLifetime(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())

	setup := func() (*singleClient, *mockConn) {
		m := &mockConn{}
		client, err := newSingleClient(
			&ClientOption{InitAddress: []string{""}, ConnLifetime: 5 * time.Second},
			m,
			func(dst string, opt *ClientOption) conn { return m },
			newRetryer(defaultRetryDelayFn),
		)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		return client, m
	}

	t.Run("Do", func(t *testing.T) {
		client, m := setup()
		m.DoFn = func(cmd Completed) RedisResult {
			return newResult(strmsg('+', "OK"), nil)
		}
		if v, err := client.Do(context.Background(), client.B().Get().Key("Do").Build()).ToString(); err != nil || v != "OK" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("DoMulti", func(t *testing.T) {
		client, m := setup()
		m.DoMultiFn = func(multi ...Completed) *redisresults {
			return &redisresults{s: []RedisResult{newResult(strmsg('+', "OK"), nil)}}
		}
		if v, err := client.DoMulti(context.Background(), client.B().Get().Key("Do").Build())[0].ToString(); err != nil || v != "OK" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("DoMulti ConnLifetime - at the head of processing", func(t *testing.T) {
		client, m := setup()
		attempts := 0
		m.DoMultiFn = func(multi ...Completed) *redisresults {
			attempts++
			if attempts == 1 {
				return &redisresults{s: []RedisResult{newErrResult(errConnExpired)}}
			}
			return &redisresults{s: []RedisResult{newResult(strmsg('+', "OK"), nil)}}
		}
		if v, err := client.DoMulti(context.Background(), client.B().Get().Key("Do").Build())[0].ToString(); err != nil || v != "OK" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("DoMulti ConnLifetime in the middle of processing", func(t *testing.T) {
		client, m := setup()
		attempts := 0
		m.DoMultiFn = func(multi ...Completed) *redisresults {
			attempts++
			if attempts == 1 {
				return &redisresults{s: []RedisResult{newResult(strmsg('+', "OK"), nil), newErrResult(errConnExpired)}}
			}
			// recover the failure of the first call
			return &redisresults{s: []RedisResult{newResult(strmsg('+', "OK"), nil)}}
		}
		resps := client.DoMulti(context.Background(), client.B().Get().Key("Do").Build(), client.B().Get().Key("Do").Build())
		if len(resps) != 2 {
			t.Errorf("unexpected response length %v", len(resps))
		}
		for _, resp := range resps {
			if v, err := resp.ToString(); err != nil || v != "OK" {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		}
	})

	t.Run("DoMulti ConnLifetime Transaction Block", func(t *testing.T) {
		client, m := setup()
		attempts := 0
		var orgMulti []Completed
		m.DoMultiFn = func(multi ...Completed) *redisresults {
			attempts++
			switch attempts {
			case 1: // errConnExpired at the head of processing
				orgMulti = multi
				return &redisresults{s: []RedisResult{newErrResult(errConnExpired), newErrResult(errConnExpired), newErrResult(errConnExpired), newErrResult(errConnExpired), newErrResult(errConnExpired), newErrResult(errConnExpired)}}
			case 2: // errConnExpired at Multi Command
				if len(multi) != 6 || !reflect.DeepEqual(multi[0].Commands(), orgMulti[0].Commands()) {
					t.Fatalf("unexpected multi when errConnExpired occurred at the head of processing, %v", multi)
				}
				return &redisresults{s: []RedisResult{
					newResult(strmsg('+', "1"), nil),
					newErrResult(errConnExpired), newErrResult(errConnExpired), newErrResult(errConnExpired), newErrResult(errConnExpired), newErrResult(errConnExpired),
				}}
			case 3: // errConnExpired in the middle of transaction block
				if len(multi) != 5 || !reflect.DeepEqual(multi[0].Commands(), orgMulti[1].Commands()) {
					t.Fatalf("unexpected multi when errConnExpired occurred at Multi Command, %v", multi)
				}
				return &redisresults{s: []RedisResult{
					newResult(strmsg('+', "OK"), nil),
					newResult(strmsg('+', "QUEUE"), nil),
					newErrResult(errConnExpired), newErrResult(errConnExpired), newErrResult(errConnExpired),
				}}
			case 4: // errConnExpired at Exec Command
				if len(multi) != 5 || !reflect.DeepEqual(multi[0].Commands(), orgMulti[1].Commands()) {
					t.Fatalf("unexpected multi when errConnExpired occurred in the middle of transaction block, %v", multi)
				}
				return &redisresults{s: []RedisResult{
					newResult(strmsg('+', "OK"), nil),
					newResult(strmsg('+', "QUEUE"), nil),
					newResult(strmsg('+', "QUEUE"), nil),
					newErrResult(errConnExpired), newErrResult(errConnExpired)}}
			case 5: // errConnExpired at end of processing
				if len(multi) != 5 || !reflect.DeepEqual(multi[0].Commands(), orgMulti[1].Commands()) {
					t.Fatalf("unexpected multi when errConnExpired occurred at at Exec Command, %v", multi)
				}
				return &redisresults{s: []RedisResult{
					newResult(strmsg('+', "OK"), nil),
					newResult(strmsg('+', "QUEUE"), nil),
					newResult(strmsg('+', "QUEUE"), nil),
					newResult(slicemsg('*', []RedisMessage{
						strmsg('+', "2"),
						strmsg('+', "3"),
					}), nil),
					newErrResult(errConnExpired),
				}}
			case 6:
				if len(multi) != 1 || !reflect.DeepEqual(multi[0].Commands(), orgMulti[5].Commands()) {
					t.Fatalf("unexpected multi when errConnExpired occurred at end of processing, %v", multi)
				}
			}
			return &redisresults{s: []RedisResult{newResult(strmsg('+', "4"), nil)}}
		}
		multi := []Completed{
			client.B().Get().Key("1{t}").Build(),
			client.B().Multi().Build(),
			client.B().Get().Key("2{t}").Build(),
			client.B().Get().Key("3{t}").Build(),
			client.B().Exec().Build(),
			client.B().Get().Key("4{t}").Build(),
		}
		resps := client.DoMulti(context.Background(), multi...)
		if len(resps) != 6 {
			t.Fatalf("unexpected response length %v", len(resps))
		}
		if v, err := resps[0].ToString(); err != nil || v != "1" {
			t.Errorf("unexpected response %v %v", v, err)
		}
		if v, err := resps[1].ToString(); err != nil || v != "OK" {
			t.Errorf("unexpected response %v %v", v, err)
		}
		if v, err := resps[2].ToString(); err != nil || v != "QUEUE" {
			t.Errorf("unexpected response %v %v", v, err)
		}
		if v, err := resps[3].ToString(); err != nil || v != "QUEUE" {
			t.Errorf("unexpected response %v %v", v, err)
		}
		if v, err := resps[4].AsStrSlice(); err != nil || !reflect.DeepEqual(v, []string{"2", "3"}) {
			t.Errorf("unexpected response %v %v", v, err)
		}
		if v, err := resps[5].ToString(); err != nil || v != "4" {
			t.Errorf("unexpected response %v %v", v, err)
		}
	})

	t.Run("DoMultiCache", func(t *testing.T) {
		client, m := setup()
		m.DoMultiCacheFn = func(multi ...CacheableTTL) *redisresults {
			return &redisresults{s: []RedisResult{newResult(strmsg('+', "OK"), nil)}}
		}
		cmd := client.B().Get().Key("Do").Cache()
		if v, err := client.DoMultiCache(context.Background(), CT(cmd, 0))[0].ToString(); err != nil || v != "OK" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("DoMultiCache ConnLifetime - at the head of processing", func(t *testing.T) {
		client, m := setup()
		attempts := 0
		m.DoMultiCacheFn = func(multi ...CacheableTTL) *redisresults {
			attempts++
			if attempts == 1 {
				return &redisresults{s: []RedisResult{newErrResult(errConnExpired)}}
			}
			return &redisresults{s: []RedisResult{newResult(strmsg('+', "OK"), nil)}}
		}
		cmd := client.B().Get().Key("Do").Cache()
		if v, err := client.DoMultiCache(context.Background(), CT(cmd, 0))[0].ToString(); err != nil || v != "OK" {
			t.Fatalf("unexpected response %v %v", v, err)
		}
	})

	t.Run("DoMultiCache ConnLifetime in the middle of processing", func(t *testing.T) {
		client, m := setup()
		attempts := 0
		m.DoMultiCacheFn = func(multi ...CacheableTTL) *redisresults {
			attempts++
			if attempts == 1 {
				return &redisresults{s: []RedisResult{newResult(strmsg('+', "OK"), nil), newErrResult(errConnExpired)}}
			}
			// recover the failure of the first call
			return &redisresults{s: []RedisResult{newResult(strmsg('+', "OK"), nil)}}
		}
		resps := client.DoMultiCache(context.Background(), CT(client.B().Get().Key("Do").Cache(), 0), CT(client.B().Get().Key("Do").Cache(), 0))
		if len(resps) != 2 {
			t.Errorf("unexpected response length %v", len(resps))
		}
		for _, resp := range resps {
			if v, err := resp.ToString(); err != nil || v != "OK" {
				t.Fatalf("unexpected response %v %v", v, err)
			}
		}
	})
}

func BenchmarkSingleClient_DoCache(b *testing.B) {
	ctx := context.Background()
	client, err := NewClient(ClientOption{InitAddress: []string{"127.0.0.1:6379"}, Dialer: net.Dialer{KeepAlive: -1}})
	if err != nil {
		b.Fatal(err)
	}
	keys := make([]string, 10000)
	for i := 0; i < 10000; i++ {
		keys[i] = strconv.Itoa(i)
	}
	mset := client.B().Mset().KeyValue()
	for _, v := range keys {
		mset = mset.KeyValue(v, v)
	}
	if err := client.Do(ctx, mset.Build()).Error(); err != nil {
		b.Fatal(err)
	}
	b.Run("NoCache", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if err := client.Do(ctx, client.B().Mget().Key(keys...).Build()).Error(); err != nil {
					b.Errorf("unexpected %v", err)
				}
			}
		})
		b.StopTimer()
	})
	b.Run("DoCache", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if err := client.DoCache(ctx, client.B().Mget().Key(keys...).Cache(), time.Minute).Error(); err != nil {
					b.Errorf("unexpected %v", err)
				}
			}
		})
		b.StopTimer()
	})
	client.Close()
}
