package rueidis

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"net"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/redis/rueidis/internal/cmds"
)

func setupMux(wires []*mockWire) (conn *mux, checkClean func(t *testing.T)) {
	return setupMuxWithOption(wires, &ClientOption{})
}

func setupMuxWithOption(wires []*mockWire, option *ClientOption) (conn *mux, checkClean func(t *testing.T)) {
	var mu sync.Mutex
	var count = -1
	wfn := func(_ context.Context) wire {
		mu.Lock()
		defer mu.Unlock()
		count++
		return wires[count]
	}
	if option.BlockingPipeline == 0 {
		option.BlockingPipeline = DefaultBlockingPipeline
	}
	return newMux("", option, (*mockWire)(nil), (*mockWire)(nil), wfn, wfn), func(t *testing.T) {
		if count != len(wires)-1 {
			t.Fatalf("there is %d remaining unused wires", len(wires)-count-1)
		}
	}
}

func TestNewMuxDailErr(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	c := 0
	e := errors.New("any")
	m := makeMux("", &ClientOption{}, func(ctx context.Context, dst string, opt *ClientOption) (net.Conn, error) {
		timer := time.NewTimer(time.Millisecond * 10) // delay time
		defer timer.Stop()
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-timer.C:
			// noop
		}
		c++
		return nil, e
	})
	if err := m.Dial(); err != e {
		t.Fatalf("unexpected return %v", err)
	}
	ctx1, cancel1 := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel1()
	if _, err := m._pipe(ctx1, 0); err != context.DeadlineExceeded {
		t.Fatalf("unexpected return %v", err)
	}
	if c != 1 {
		t.Fatalf("dialFn not called")
	}
	if w := m.pipe(context.Background(), 0); w != m.dead { // c = 2
		t.Fatalf("unexpected wire %v", w)
	}
	ctx2, cancel2 := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel2()
	if w := m.pipe(ctx2, 0); w != m.dead {
		t.Fatalf("unexpected wire %v", w)
	}
	if err := m.Dial(); err != e { // c = 3
		t.Fatalf("unexpected return %v", err)
	}
	if w := m.Acquire(context.Background()); w != m.dead {
		t.Fatalf("unexpected wire %v", w)
	}
	ctx3, cancel3 := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel3()
	if w := m.Acquire(ctx3); w != m.dead {
		t.Fatalf("unexpected wire %v", w)
	}
	ctx4, cancel4 := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel4()
	if w := m.Acquire(ctx4); w != m.dead {
		t.Fatalf("unexpected wire %v", w)
	}
	if c != 5 {
		t.Fatalf("dialFn not called %v", c)
	}
}

func TestNewMux(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	n1, n2 := net.Pipe()
	mock := &redisMock{t: t, buf: bufio.NewReader(n2), conn: n2}
	go func() {
		mock.Expect("HELLO", "3").
			Reply(slicemsg(
				'%',
				[]RedisMessage{
					strmsg('+', "proto"),
					{typ: ':', intlen: 3},
				},
			))
		mock.Expect("CLIENT", "TRACKING", "ON", "OPTIN").
			ReplyString("OK")
		mock.Expect("CLIENT", "SETINFO", "LIB-NAME", LibName).
			ReplyError("UNKNOWN COMMAND")
		mock.Expect("CLIENT", "SETINFO", "LIB-VER", LibVer).
			ReplyError("UNKNOWN COMMAND")
		mock.Expect("PING").ReplyString("OK")
		mock.Close()
	}()
	m := makeMux("", &ClientOption{}, func(_ context.Context, dst string, opt *ClientOption) (net.Conn, error) {
		return n1, nil
	})
	if err := m.Dial(); err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	t.Run("Override with previous mux", func(t *testing.T) {
		m2 := makeMux("", &ClientOption{}, func(_ context.Context, dst string, opt *ClientOption) (net.Conn, error) {
			return n1, nil
		})
		m2.Override(m)
		if err := m2.Dial(); err != nil {
			t.Fatalf("unexpected error %v", err)
		}
		m2.Close()
	})
}

func TestNewMuxPipelineMultiplex(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	for _, v := range []int{-1, 0, 1, 2} {
		m := makeMux("", &ClientOption{PipelineMultiplex: v}, func(_ context.Context, dst string, opt *ClientOption) (net.Conn, error) { return nil, nil })
		if (v < 0 && len(m.muxwires) != 1) || (v >= 0 && len(m.muxwires) != 1<<v) {
			t.Fatalf("unexpected len(m.muxwires): %v", len(m.muxwires))
		}
	}
}

func TestMuxAddr(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	m := makeMux("dst1", &ClientOption{}, nil)
	if m.Addr() != "dst1" {
		t.Fatalf("unexpected m.Addr != dst1")
	}
}

func TestMuxOptInCmd(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())

	if m := makeMux("dst1", &ClientOption{
		ClientTrackingOptions: []string{"OPTOUT"},
	}, nil); m.OptInCmd() != cmds.OptInNopCmd {
		t.Fatalf("unexpected OptInCmd")
	}
	if m := makeMux("dst1", &ClientOption{
		ClientTrackingOptions: []string{"PREFIX", "a", "BCAST"},
	}, nil); m.OptInCmd() != cmds.OptInNopCmd {
		t.Fatalf("unexpected OptInCmd")
	}
	if m := makeMux("dst1", &ClientOption{
		ClientTrackingOptions: nil,
	}, nil); m.OptInCmd() != cmds.OptInCmd {
		t.Fatalf("unexpected OptInCmd")
	}
}

func TestMuxDialSuppress(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	var wires, waits, done int64
	blocking := make(chan struct{})
	m := newMux("", &ClientOption{}, (*mockWire)(nil), (*mockWire)(nil), func(_ context.Context) wire {
		atomic.AddInt64(&wires, 1)
		<-blocking
		return &mockWire{}
	}, func(_ context.Context) wire {
		return &mockWire{}
	})
	for i := 0; i < 1000; i++ {
		go func() {
			atomic.AddInt64(&waits, 1)
			m.Info()
			atomic.AddInt64(&done, 1)
		}()
	}
	for atomic.LoadInt64(&waits) != 1000 {
		runtime.Gosched()
	}
	close(blocking)
	for atomic.LoadInt64(&done) != 1000 {
		runtime.Gosched()
	}
	if atomic.LoadInt64(&wires) != 1 {
		t.Fatalf("wireFn is not suppressed")
	}
}

//gocyclo:ignore
func TestMuxReuseWire(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	t.Run("reuse wire if no error", func(t *testing.T) {
		m, checkClean := setupMux([]*mockWire{
			{
				DoFn: func(cmd Completed) RedisResult {
					return newResult(strmsg('+', "PONG"), nil)
				},
			},
		})
		defer checkClean(t)
		defer m.Close()
		for i := 0; i < 2; i++ {
			if err := m.Do(context.Background(), cmds.NewCompleted([]string{"PING"})).Error(); err != nil {
				t.Fatalf("unexpected error %v", err)
			}
		}
		m.Close()
	})

	t.Run("reuse blocking (dpool) pool", func(t *testing.T) {
		blocking := make(chan struct{})
		response := make(chan RedisResult)
		m, checkClean := setupMux([]*mockWire{
			{
				// leave first wire for pipeline calls
			},
			{
				DoFn: func(cmd Completed) RedisResult {
					return newResult(strmsg('+', "ACQUIRED"), nil)
				},
			},
			{
				DoFn: func(cmd Completed) RedisResult {
					blocking <- struct{}{}
					return <-response
				},
			},
		})
		defer checkClean(t)
		defer m.Close()
		if err := m.Dial(); err != nil {
			t.Fatalf("unexpected dial error %v", err)
		}

		wire1 := m.dpool.Acquire(context.Background())

		go func() {
			// this should use the second wire
			if val, err := m.Do(context.Background(), cmds.NewBlockingCompleted([]string{"PING"})).ToString(); err != nil {
				t.Errorf("unexpected error %v", err)
			} else if val != "BLOCK_RESPONSE" {
				t.Errorf("unexpected response %v", val)
			}
			close(blocking)
		}()
		<-blocking

		m.dpool.Store(wire1)
		// this should use the first wire
		if val, err := m.Do(context.Background(), cmds.NewBlockingCompleted([]string{"PING"})).ToString(); err != nil {
			t.Fatalf("unexpected error %v", err)
		} else if val != "ACQUIRED" {
			t.Fatalf("unexpected response %v", val)
		}

		response <- newResult(strmsg('+', "BLOCK_RESPONSE"), nil)
		<-blocking
	})

	t.Run("reuse blocking (spool) pool", func(t *testing.T) {
		blocking := make(chan struct{})
		response := make(chan RedisResult)
		m, checkClean := setupMux([]*mockWire{
			{
				// leave first wire for pipeline calls
				DoFn: func(cmd Completed) RedisResult {
					return newResult(strmsg('+', "PIPELINED"), nil)
				},
			},
			{
				DoFn: func(cmd Completed) RedisResult {
					return newResult(strmsg('+', "ACQUIRED"), nil)
				},
			},
			{
				DoFn: func(cmd Completed) RedisResult {
					blocking <- struct{}{}
					return <-response
				},
			},
		})
		m.usePool = true // switch to spool
		defer checkClean(t)
		defer m.Close()
		if err := m.Dial(); err != nil {
			t.Fatalf("unexpected dial error %v", err)
		}

		wire1 := m.spool.Acquire(context.Background())

		go func() {
			// this should use the second wire
			if val, err := m.Do(context.Background(), cmds.NewBlockingCompleted([]string{"PING"})).ToString(); err != nil {
				t.Errorf("unexpected error %v", err)
			} else if val != "BLOCK_RESPONSE" {
				t.Errorf("unexpected response %v", val)
			}
			close(blocking)
		}()
		<-blocking

		m.spool.Store(wire1)
		// this should use the first wire
		if val, err := m.Do(context.Background(), cmds.NewBlockingCompleted([]string{"PING"})).ToString(); err != nil {
			t.Fatalf("unexpected error %v", err)
		} else if val != "ACQUIRED" {
			t.Fatalf("unexpected response %v", val)
		}

		// this should use auto pipeline
		if val, err := m.Do(context.Background(), cmds.NewCompleted([]string{"PING"}).ToPipe()).ToString(); err != nil {
			t.Fatalf("unexpected error %v", err)
		} else if val != "PIPELINED" {
			t.Fatalf("unexpected response %v", val)
		}

		response <- newResult(strmsg('+', "BLOCK_RESPONSE"), nil)
		<-blocking
	})

	t.Run("reuse blocking (dpool) pool DoMulti", func(t *testing.T) {
		blocking := make(chan struct{})
		response := make(chan RedisResult)
		m, checkClean := setupMux([]*mockWire{
			{
				// leave first wire for pipeline calls
				DoMultiFn: func(cmd ...Completed) *redisresults {
					return &redisresults{s: []RedisResult{newResult(strmsg('+', "PIPELINED"), nil)}}
				},
			},
			{
				DoMultiFn: func(cmd ...Completed) *redisresults {
					return &redisresults{s: []RedisResult{newResult(strmsg('+', "ACQUIRED"), nil)}}
				},
			},
			{
				DoMultiFn: func(cmd ...Completed) *redisresults {
					blocking <- struct{}{}
					return &redisresults{s: []RedisResult{<-response}}
				},
			},
		})
		m.usePool = true // switch to spool
		defer checkClean(t)
		defer m.Close()
		if err := m.Dial(); err != nil {
			t.Fatalf("unexpected dial error %v", err)
		}

		wire1 := m.spool.Acquire(context.Background())

		go func() {
			// this should use the second wire
			if val, err := m.DoMulti(context.Background(), cmds.NewBlockingCompleted([]string{"PING"})).s[0].ToString(); err != nil {
				t.Errorf("unexpected error %v", err)
			} else if val != "BLOCK_RESPONSE" {
				t.Errorf("unexpected response %v", val)
			}
			close(blocking)
		}()
		<-blocking

		m.spool.Store(wire1)
		// this should use the first wire
		if val, err := m.DoMulti(context.Background(), cmds.NewBlockingCompleted([]string{"PING"})).s[0].ToString(); err != nil {
			t.Fatalf("unexpected error %v", err)
		} else if val != "ACQUIRED" {
			t.Fatalf("unexpected response %v", val)
		}

		// this should use auto pipeline
		if val, err := m.DoMulti(context.Background(), cmds.NewCompleted([]string{"PING"}).ToPipe()).s[0].ToString(); err != nil {
			t.Fatalf("unexpected error %v", err)
		} else if val != "PIPELINED" {
			t.Fatalf("unexpected response %v", val)
		}

		response <- newResult(strmsg('+', "BLOCK_RESPONSE"), nil)
		<-blocking
	})

	t.Run("reuse blocking (spool) pool DoMulti", func(t *testing.T) {
		blocking := make(chan struct{})
		response := make(chan RedisResult)
		m, checkClean := setupMux([]*mockWire{
			{
				// leave first wire for pipeline calls
			},
			{
				DoMultiFn: func(cmd ...Completed) *redisresults {
					return &redisresults{s: []RedisResult{newResult(strmsg('+', "ACQUIRED"), nil)}}
				},
			},
			{
				DoMultiFn: func(cmd ...Completed) *redisresults {
					blocking <- struct{}{}
					return &redisresults{s: []RedisResult{<-response}}
				},
			},
		})
		defer checkClean(t)
		defer m.Close()
		if err := m.Dial(); err != nil {
			t.Fatalf("unexpected dial error %v", err)
		}

		wire1 := m.dpool.Acquire(context.Background())

		go func() {
			// this should use the second wire
			if val, err := m.DoMulti(context.Background(), cmds.NewBlockingCompleted([]string{"PING"})).s[0].ToString(); err != nil {
				t.Errorf("unexpected error %v", err)
			} else if val != "BLOCK_RESPONSE" {
				t.Errorf("unexpected response %v", val)
			}
			close(blocking)
		}()
		<-blocking

		m.dpool.Store(wire1)
		// this should use the first wire
		if val, err := m.DoMulti(context.Background(), cmds.NewBlockingCompleted([]string{"PING"})).s[0].ToString(); err != nil {
			t.Fatalf("unexpected error %v", err)
		} else if val != "ACQUIRED" {
			t.Fatalf("unexpected response %v", val)
		}

		response <- newResult(strmsg('+', "BLOCK_RESPONSE"), nil)
		<-blocking
	})

	t.Run("unsubscribe blocking pool", func(t *testing.T) {
		cleaned := false
		m, checkClean := setupMux([]*mockWire{
			{
				// leave first wire for pipeline calls
			},
			{
				CleanSubscriptionsFn: func() {
					cleaned = true
				},
			},
		})
		defer checkClean(t)
		defer m.Close()

		if err := m.Dial(); err != nil {
			t.Fatalf("unexpected dial error %v", err)
		}

		wire1 := m.Acquire(context.Background())
		m.Store(wire1)

		if !cleaned {
			t.Fatalf("CleanSubscriptions not called")
		}
	})
}

//gocyclo:ignore
func TestMuxDelegation(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	t.Run("wire info", func(t *testing.T) {
		m, checkClean := setupMux([]*mockWire{
			{
				InfoFn: func() map[string]RedisMessage {
					return map[string]RedisMessage{"key": strmsg('+', "value")}
				},
			},
		})
		defer checkClean(t)
		defer m.Close()
		if info := m.Info(); info == nil {
			t.Fatalf("unexpected info %v", info)
		} else if infoKey := info["key"]; infoKey.string() != "value" {
			t.Fatalf("unexpected info %v", info)
		}
	})

	t.Run("wire version", func(t *testing.T) {
		m, checkClean := setupMux([]*mockWire{
			{
				VersionFn: func() int {
					return 7
				},
			},
		})
		defer checkClean(t)
		defer m.Close()
		if version := m.Version(); version != 7 {
			t.Fatalf("unexpected version %v", version)
		}
	})

	t.Run("wire az", func(t *testing.T) {
		m, checkClean := setupMux([]*mockWire{
			{
				AZFn: func() string {
					return "az"
				},
			},
		})
		defer checkClean(t)
		defer m.Close()
		if az := m.AZ(); az != "az" {
			t.Fatalf("unexpected az %v", az)
		}
	})

	t.Run("wire err", func(t *testing.T) {
		e := errors.New("err")
		m, checkClean := setupMux([]*mockWire{
			{
				ErrorFn: func() error {
					return e
				},
			},
		})
		defer checkClean(t)
		defer m.Close()
		if err := m.Error(); err != e {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("wire do", func(t *testing.T) {
		m, checkClean := setupMux([]*mockWire{
			{
				DoFn: func(cmd Completed) RedisResult {
					return newErrResult(context.DeadlineExceeded)
				},
				ErrorFn: func() error {
					return context.DeadlineExceeded
				},
			},
			{
				DoFn: func(cmd Completed) RedisResult {
					if cmd.Commands()[0] != "READONLY_COMMAND" {
						t.Fatalf("command should be READONLY_COMMAND")
					}
					return newResult(strmsg('+', "READONLY_COMMAND_RESPONSE"), nil)
				},
			},
		})
		defer checkClean(t)
		defer m.Close()
		if err := m.Do(context.Background(), cmds.NewReadOnlyCompleted([]string{"READONLY_COMMAND"})).Error(); !errors.Is(err, context.DeadlineExceeded) {
			t.Fatalf("unexpected error %v", err)
		}
		if val, err := m.Do(context.Background(), cmds.NewReadOnlyCompleted([]string{"READONLY_COMMAND"})).ToString(); err != nil {
			t.Fatalf("unexpected error %v", err)
		} else if val != "READONLY_COMMAND_RESPONSE" {
			t.Fatalf("unexpected response %v", val)
		}
	})

	t.Run("wire do stream", func(t *testing.T) {
		m, checkClean := setupMux([]*mockWire{
			{
				DoStreamFn: func(pool *pool, cmd Completed) RedisResultStream {
					return RedisResultStream{e: errors.New(cmd.Commands()[0])}
				},
			},
		})
		defer checkClean(t)
		defer m.Close()
		if s := m.DoStream(context.Background(), cmds.NewReadOnlyCompleted([]string{"READONLY_COMMAND"})); s.Error().Error() != "READONLY_COMMAND" {
			t.Fatalf("unexpected error %v", s.Error())
		}
	})

	t.Run("wire do multi", func(t *testing.T) {
		m, checkClean := setupMux([]*mockWire{
			{
				DoMultiFn: func(multi ...Completed) *redisresults {
					return &redisresults{s: []RedisResult{newErrResult(context.DeadlineExceeded)}}
				},
				ErrorFn: func() error {
					return context.DeadlineExceeded
				},
			},
			{
				DoMultiFn: func(multi ...Completed) *redisresults {
					return &redisresults{s: []RedisResult{newResult(strmsg('+', "MULTI_COMMANDS_RESPONSE"), nil)}}
				},
			},
		})
		defer checkClean(t)
		defer m.Close()
		if err := m.DoMulti(context.Background(), cmds.NewReadOnlyCompleted([]string{"READONLY_COMMAND"})).s[0].Error(); !errors.Is(err, context.DeadlineExceeded) {
			t.Fatalf("unexpected error %v", err)
		}
		if val, err := m.DoMulti(context.Background(), cmds.NewReadOnlyCompleted([]string{"READONLY_COMMAND"})).s[0].ToString(); err != nil {
			t.Fatalf("unexpected error %v", err)
		} else if val != "MULTI_COMMANDS_RESPONSE" {
			t.Fatalf("unexpected response %v", val)
		}
	})

	t.Run("wire do multi stream", func(t *testing.T) {
		m, checkClean := setupMux([]*mockWire{
			{
				DoMultiStreamFn: func(pool *pool, cmd ...Completed) MultiRedisResultStream {
					return MultiRedisResultStream{e: errors.New(cmd[0].Commands()[0])}
				},
			},
		})
		defer checkClean(t)
		defer m.Close()
		if s := m.DoMultiStream(context.Background(), cmds.NewReadOnlyCompleted([]string{"READONLY_COMMAND"})); s.Error().Error() != "READONLY_COMMAND" {
			t.Fatalf("unexpected error %v", s.Error())
		}
	})

	t.Run("wire do cache", func(t *testing.T) {
		m, checkClean := setupMux([]*mockWire{
			{
				DoCacheFn: func(cmd Cacheable, ttl time.Duration) RedisResult {
					return newErrResult(context.DeadlineExceeded)
				},
				ErrorFn: func() error {
					return context.DeadlineExceeded
				},
			},
			{
				DoCacheFn: func(cmd Cacheable, ttl time.Duration) RedisResult {
					return newResult(strmsg('+', "READONLY_COMMAND_RESPONSE"), nil)
				},
			},
		})
		defer checkClean(t)
		defer m.Close()
		if err := m.DoCache(context.Background(), Cacheable(cmds.NewReadOnlyCompleted([]string{"READONLY_COMMAND"})), time.Second).Error(); !errors.Is(err, context.DeadlineExceeded) {
			t.Fatalf("unexpected error %v", err)
		}
		if val, err := m.DoCache(context.Background(), Cacheable(cmds.NewReadOnlyCompleted([]string{"READONLY_COMMAND"})), time.Second).ToString(); err != nil {
			t.Fatalf("unexpected error %v", err)
		} else if val != "READONLY_COMMAND_RESPONSE" {
			t.Fatalf("unexpected response %v", val)
		}
	})

	t.Run("wire do multi cache", func(t *testing.T) {
		m, checkClean := setupMux([]*mockWire{
			{
				DoMultiCacheFn: func(multi ...CacheableTTL) *redisresults {
					return &redisresults{s: []RedisResult{newErrResult(context.DeadlineExceeded)}}
				},
				ErrorFn: func() error {
					return context.DeadlineExceeded
				},
			},
			{
				DoMultiCacheFn: func(multi ...CacheableTTL) *redisresults {
					return &redisresults{s: []RedisResult{newResult(strmsg('+', "MULTI_COMMANDS_RESPONSE"), nil)}}
				},
			},
		})
		defer checkClean(t)
		defer m.Close()
		if err := m.DoMultiCache(context.Background(), CT(Cacheable(cmds.NewReadOnlyCompleted([]string{"READONLY_COMMAND"})), time.Second)).s[0].Error(); !errors.Is(err, context.DeadlineExceeded) {
			t.Fatalf("unexpected error %v", err)
		}
		if val, err := m.DoMultiCache(context.Background(), CT(Cacheable(cmds.NewReadOnlyCompleted([]string{"READONLY_COMMAND"})), time.Second)).s[0].ToString(); err != nil {
			t.Fatalf("unexpected error %v", err)
		} else if val != "MULTI_COMMANDS_RESPONSE" {
			t.Fatalf("unexpected response %v", val)
		}
	})

	t.Run("wire do multi cache multiple slots", func(t *testing.T) {
		multiplex := 1
		wires := make([]*mockWire, 1<<multiplex)
		for i := range wires {
			idx := uint16(i)
			wires[i] = &mockWire{
				DoMultiCacheFn: func(multi ...CacheableTTL) *redisresults {
					result := make([]RedisResult, len(multi))
					for j, cmd := range multi {
						if s := cmd.Cmd.Slot() & uint16(len(wires)-1); s != idx {
							result[j] = newErrResult(fmt.Errorf("wrong slot %v %v", s, idx))
						} else {
							result[j] = newResult(strmsg('+', cmd.Cmd.Commands()[1]), nil)
						}
					}
					return &redisresults{s: result}
				},
			}
		}
		m, checkClean := setupMuxWithOption(wires, &ClientOption{PipelineMultiplex: multiplex})
		defer checkClean(t)
		defer m.Close()

		for i := range wires {
			m._pipe(context.Background(), uint16(i))
		}

		builder := cmds.NewBuilder(cmds.NoSlot)

		for count := 1; count <= 3; count++ {
			commands := make([]CacheableTTL, count)
			for c := 0; c < count; c++ {
				commands[c] = CT(builder.Get().Key(strconv.Itoa(c)).Cache(), time.Second)
			}
			for i, resp := range m.DoMultiCache(context.Background(), commands...).s {
				if v, err := resp.ToString(); err != nil || v != strconv.Itoa(i) {
					t.Fatalf("unexpected resp %v %v", v, err)
				}
			}
		}
	})

	t.Run("wire do multi cache multiple slots fail", func(t *testing.T) {
		multiplex := 1
		wires := make([]*mockWire, 1<<multiplex)
		for i := range wires {
			idx := uint16(i)
			wires[i] = &mockWire{
				DoMultiCacheFn: func(multi ...CacheableTTL) *redisresults {
					for _, cmd := range multi {
						if s := cmd.Cmd.Slot() & uint16(len(wires)-1); s != idx {
							return &redisresults{s: []RedisResult{newErrResult(fmt.Errorf("wrong slot %v %v", s, idx))}}
						}
					}
					return &redisresults{s: []RedisResult{newErrResult(context.DeadlineExceeded)}}
				},
				ErrorFn: func() error {
					return context.DeadlineExceeded
				},
			}
		}
		m, checkClean := setupMuxWithOption(wires, &ClientOption{PipelineMultiplex: multiplex})
		defer checkClean(t)
		defer m.Close()

		for i := range wires {
			m._pipe(context.Background(), uint16(i))
		}

		builder := cmds.NewBuilder(cmds.NoSlot)
		commands := make([]CacheableTTL, 4)
		for c := 0; c < len(commands); c++ {
			commands[c] = CT(builder.Get().Key(strconv.Itoa(c)).Cache(), time.Second)
		}
		if err := m.DoMultiCache(context.Background(), commands...).s[0].Error(); !errors.Is(err, context.DeadlineExceeded) {
			t.Fatalf("unexpected error %v", err)
		}
	})

	t.Run("wire receive", func(t *testing.T) {
		m, checkClean := setupMux([]*mockWire{
			{
				ReceiveFn: func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
					return context.DeadlineExceeded
				},
				ErrorFn: func() error {
					return context.DeadlineExceeded
				},
			},
			{
				ReceiveFn: func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
					if subscribe.Commands()[0] != "SUBSCRIBE" {
						t.Fatalf("command should be SUBSCRIBE")
					}
					return nil
				},
			},
		})
		defer checkClean(t)
		defer m.Close()
		if err := m.Receive(context.Background(), cmds.NewCompleted([]string{"SUBSCRIBE"}), func(message PubSubMessage) {}); !errors.Is(err, context.DeadlineExceeded) {
			t.Fatalf("unexpected error %v", err)
		}
		if err := m.Receive(context.Background(), cmds.NewCompleted([]string{"SUBSCRIBE"}), func(message PubSubMessage) {}); err != nil {
			t.Fatalf("unexpected error %v", err)
		}
	})

	t.Run("single blocking", func(t *testing.T) {
		blocked := make(chan struct{})
		responses := make(chan RedisResult)

		m, checkClean := setupMux([]*mockWire{
			{
				// leave first wire for pipeline calls
			},
			{
				DoFn: func(cmd Completed) RedisResult {
					blocked <- struct{}{}
					return <-responses
				},
			},
			{
				DoFn: func(cmd Completed) RedisResult {
					blocked <- struct{}{}
					return <-responses
				},
			},
		})
		defer checkClean(t)
		defer m.Close()
		if err := m.Dial(); err != nil {
			t.Fatalf("unexpected dial error %v", err)
		}

		wg := sync.WaitGroup{}
		wg.Add(2)
		for i := 0; i < 2; i++ {
			go func() {
				if val, err := m.Do(context.Background(), cmds.NewBlockingCompleted([]string{"BLOCK"})).ToString(); err != nil {
					t.Errorf("unexpected error %v", err)
				} else if val != "BLOCK_COMMANDS_RESPONSE" {
					t.Errorf("unexpected response %v", val)
				} else {
					wg.Done()
				}
			}()
		}
		for i := 0; i < 2; i++ {
			<-blocked
		}
		for i := 0; i < 2; i++ {
			responses <- newResult(strmsg('+', "BLOCK_COMMANDS_RESPONSE"), nil)
		}
		wg.Wait()
	})

	t.Run("single blocking no recycle the wire if err", func(t *testing.T) {
		closed := false
		m, checkClean := setupMux([]*mockWire{
			{
				// leave first wire for pipeline calls
			},
			{
				DoFn: func(cmd Completed) RedisResult {
					return newErrResult(context.DeadlineExceeded)
				},
				ErrorFn: func() error {
					return context.DeadlineExceeded
				},
				CloseFn: func() {
					closed = true
				},
			},
			{
				DoFn: func(cmd Completed) RedisResult {
					return newResult(strmsg('+', "OK"), nil)
				},
			},
		})
		defer checkClean(t)
		defer m.Close()
		if err := m.Dial(); err != nil {
			t.Fatalf("unexpected dial error %v", err)
		}
		if err := m.Do(context.Background(), cmds.NewBlockingCompleted([]string{"BLOCK"})).Error(); err != context.DeadlineExceeded {
			t.Errorf("unexpected error %v", err)
		}
		if val, err := m.Do(context.Background(), cmds.NewBlockingCompleted([]string{"BLOCK"})).ToString(); err != nil || val != "OK" {
			t.Errorf("unexpected response %v %v", err, val)
		}
		if !closed {
			t.Errorf("wire not closed")
		}
	})

	t.Run("multiple blocking", func(t *testing.T) {
		blocked := make(chan struct{})
		responses := make(chan RedisResult)

		m, checkClean := setupMux([]*mockWire{
			{
				// leave first wire for pipeline calls
			},
			{
				DoMultiFn: func(cmd ...Completed) *redisresults {
					blocked <- struct{}{}
					return &redisresults{s: []RedisResult{<-responses}}
				},
			},
			{
				DoMultiFn: func(cmd ...Completed) *redisresults {
					blocked <- struct{}{}
					return &redisresults{s: []RedisResult{<-responses}}
				},
			},
		})
		defer checkClean(t)
		defer m.Close()
		if err := m.Dial(); err != nil {
			t.Fatalf("unexpected dial error %v", err)
		}

		wg := sync.WaitGroup{}
		wg.Add(2)
		for i := 0; i < 2; i++ {
			go func() {
				if val, err := m.DoMulti(
					context.Background(),
					cmds.NewReadOnlyCompleted([]string{"READONLY"}),
					cmds.NewBlockingCompleted([]string{"BLOCK"}),
				).s[0].ToString(); err != nil {
					t.Errorf("unexpected error %v", err)
				} else if val != "BLOCK_COMMANDS_RESPONSE" {
					t.Errorf("unexpected response %v", val)
				} else {
					wg.Done()
				}
			}()
		}
		for i := 0; i < 2; i++ {
			<-blocked
		}
		for i := 0; i < 2; i++ {
			responses <- newResult(strmsg('+', "BLOCK_COMMANDS_RESPONSE"), nil)
		}
		wg.Wait()
	})

	t.Run("multiple long pipeline", func(t *testing.T) {
		blocked := make(chan struct{})
		responses := make(chan RedisResult)

		m, checkClean := setupMux([]*mockWire{
			{
				// leave first wire for pipeline calls
			},
			{
				DoMultiFn: func(cmd ...Completed) *redisresults {
					blocked <- struct{}{}
					return &redisresults{s: []RedisResult{<-responses}}
				},
			},
			{
				DoMultiFn: func(cmd ...Completed) *redisresults {
					blocked <- struct{}{}
					return &redisresults{s: []RedisResult{<-responses}}
				},
			},
		})
		defer checkClean(t)
		defer m.Close()
		if err := m.Dial(); err != nil {
			t.Fatalf("unexpected dial error %v", err)
		}

		wg := sync.WaitGroup{}
		wg.Add(2)
		for i := 0; i < 2; i++ {
			go func() {
				pipeline := make(Commands, DefaultBlockingPipeline)
				for i := 0; i < len(pipeline); i++ {
					pipeline[i] = cmds.NewCompleted([]string{"SET"})
				}
				if val, err := m.DoMulti(context.Background(), pipeline...).s[0].ToString(); err != nil {
					t.Errorf("unexpected error %v", err)
				} else if val != "BLOCK_COMMANDS_RESPONSE" {
					t.Errorf("unexpected response %v", val)
				} else {
					wg.Done()
				}
			}()
		}
		for i := 0; i < 2; i++ {
			<-blocked
		}
		for i := 0; i < 2; i++ {
			responses <- newResult(strmsg('+', "BLOCK_COMMANDS_RESPONSE"), nil)
		}
		wg.Wait()
	})

	t.Run("multi blocking no recycle the wire if err", func(t *testing.T) {
		closed := false
		m, checkClean := setupMux([]*mockWire{
			{
				// leave first wire for pipeline calls
			},
			{
				DoMultiFn: func(cmd ...Completed) *redisresults {
					return &redisresults{s: []RedisResult{newErrResult(context.DeadlineExceeded)}}
				},
				ErrorFn: func() error {
					return context.DeadlineExceeded
				},
				CloseFn: func() {
					closed = true
				},
			},
			{
				DoFn: func(cmd Completed) RedisResult {
					return newResult(strmsg('+', "OK"), nil)
				},
			},
		})
		defer checkClean(t)
		defer m.Close()
		if err := m.Dial(); err != nil {
			t.Fatalf("unexpected dial error %v", err)
		}
		if err := m.DoMulti(
			context.Background(),
			cmds.NewReadOnlyCompleted([]string{"READONLY"}),
			cmds.NewBlockingCompleted([]string{"BLOCK"}),
		).s[0].Error(); err != context.DeadlineExceeded {
			t.Errorf("unexpected error %v", err)
		}
		if val, err := m.Do(context.Background(), cmds.NewBlockingCompleted([]string{"BLOCK"})).ToString(); err != nil || val != "OK" {
			t.Errorf("unexpected response %v %v", err, val)
		}
		if !closed {
			t.Errorf("wire not closed")
		}
	})
}

//gocyclo:ignore
func TestMuxRegisterCloseHook(t *testing.T) {
	defer ShouldNotLeak(SetupLeakDetection())
	t.Run("trigger hook with unexpected error", func(t *testing.T) {
		var hook atomic.Value
		m, checkClean := setupMux([]*mockWire{
			{
				DoFn: func(cmd Completed) RedisResult {
					return newResult(strmsg('+', "PONG1"), nil)
				},
				SetOnCloseHookFn: func(fn func(error)) {
					hook.Store(fn)
				},
			},
			{
				DoFn: func(cmd Completed) RedisResult {
					return newResult(strmsg('+', "PONG2"), nil)
				},
			},
		})
		defer checkClean(t)
		defer m.Close()
		if resp, _ := m.Do(context.Background(), cmds.NewCompleted([]string{"PING"})).ToString(); resp != "PONG1" {
			t.Fatalf("unexpected response %v", resp)
		}
		hook.Load().(func(error))(errors.New("any")) // invoke the hook, this should cause the first wire be discarded
		if resp, _ := m.Do(context.Background(), cmds.NewCompleted([]string{"PING"})).ToString(); resp != "PONG2" {
			t.Fatalf("unexpected response %v", resp)
		}
	})
	t.Run("not trigger hook with ErrClosing", func(t *testing.T) {
		var hook atomic.Value
		m, checkClean := setupMux([]*mockWire{
			{
				DoFn: func(cmd Completed) RedisResult {
					return newResult(strmsg('+', "PONG1"), nil)
				},
				SetOnCloseHookFn: func(fn func(error)) {
					hook.Store(fn)
				},
			},
		})
		defer checkClean(t)
		defer m.Close()
		if resp, _ := m.Do(context.Background(), cmds.NewCompleted([]string{"PING"})).ToString(); resp != "PONG1" {
			t.Fatalf("unexpected response %v", resp)
		}
		hook.Load().(func(error))(ErrClosing) // invoke the hook, this should cause the first wire be discarded
		if resp, _ := m.Do(context.Background(), cmds.NewCompleted([]string{"PING"})).ToString(); resp != "PONG1" {
			t.Fatalf("unexpected response %v", resp)
		}
	})
}

func BenchmarkClientSideCaching(b *testing.B) {
	setup := func(b *testing.B) *mux {
		c := makeMux("127.0.0.1:6379", &ClientOption{CacheSizeEachConn: DefaultCacheBytes}, func(_ context.Context, dst string, opt *ClientOption) (conn net.Conn, err error) {
			return net.Dial("tcp", dst)
		})
		if err := c.Dial(); err != nil {
			panic(err)
		}
		b.SetParallelism(100)
		b.ResetTimer()
		return c
	}
	b.Run("Do", func(b *testing.B) {
		m := setup(b)
		cmd := cmds.NewCompleted([]string{"GET", "a"})
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				m.Do(context.Background(), cmd)
			}
		})
	})
	b.Run("DoCache", func(b *testing.B) {
		m := setup(b)
		cmd := Cacheable(cmds.NewCompleted([]string{"GET", "a"}))
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				m.DoCache(context.Background(), cmd, time.Second*5)
			}
		})
	})
}

type mockWire struct {
	DoFn            func(cmd Completed) RedisResult
	DoCacheFn       func(cmd Cacheable, ttl time.Duration) RedisResult
	DoMultiFn       func(multi ...Completed) *redisresults
	DoMultiCacheFn  func(multi ...CacheableTTL) *redisresults
	ReceiveFn       func(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error
	DoStreamFn      func(pool *pool, cmd Completed) RedisResultStream
	DoMultiStreamFn func(pool *pool, cmd ...Completed) MultiRedisResultStream
	InfoFn          func() map[string]RedisMessage
	AZFn            func() string
	VersionFn       func() int
	ErrorFn         func() error
	CloseFn         func()
	StopTimerFn     func() bool
	ResetTimerFn    func() bool

	CleanSubscriptionsFn func()
	SetPubSubHooksFn     func(hooks PubSubHooks) <-chan error
	SetOnCloseHookFn     func(fn func(error))
}

func (m *mockWire) Do(ctx context.Context, cmd Completed) RedisResult {
	if m.DoFn != nil {
		return m.DoFn(cmd)
	}
	return RedisResult{}
}

func (m *mockWire) DoCache(ctx context.Context, cmd Cacheable, ttl time.Duration) RedisResult {
	if m.DoCacheFn != nil {
		return m.DoCacheFn(cmd, ttl)
	}
	return RedisResult{}
}

func (m *mockWire) DoMultiCache(ctx context.Context, multi ...CacheableTTL) *redisresults {
	if m.DoMultiCacheFn != nil {
		return m.DoMultiCacheFn(multi...)
	}
	return nil
}

func (m *mockWire) DoMulti(ctx context.Context, multi ...Completed) *redisresults {
	if m.DoMultiFn != nil {
		return m.DoMultiFn(multi...)
	}
	return nil
}

func (m *mockWire) Receive(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
	if m.ReceiveFn != nil {
		return m.ReceiveFn(ctx, subscribe, fn)
	}
	return nil
}

func (m *mockWire) DoStream(ctx context.Context, pool *pool, cmd Completed) RedisResultStream {
	if m.DoStreamFn != nil {
		return m.DoStreamFn(pool, cmd)
	}
	return RedisResultStream{}
}

func (m *mockWire) DoMultiStream(ctx context.Context, pool *pool, cmd ...Completed) MultiRedisResultStream {
	if m.DoMultiStreamFn != nil {
		return m.DoMultiStreamFn(pool, cmd...)
	}
	return MultiRedisResultStream{}
}

func (m *mockWire) CleanSubscriptions() {
	if m.CleanSubscriptionsFn != nil {
		m.CleanSubscriptionsFn()
	}
}

func (m *mockWire) SetPubSubHooks(hooks PubSubHooks) <-chan error {
	if m.SetPubSubHooksFn != nil {
		return m.SetPubSubHooksFn(hooks)
	}
	return nil
}

func (m *mockWire) SetOnCloseHook(fn func(error)) {
	if m.SetOnCloseHookFn != nil {
		m.SetOnCloseHookFn(fn)
	}
}

func (m *mockWire) StopTimer() bool {
	if m.StopTimerFn != nil {
		return m.StopTimerFn()
	}
	return true
}

func (m *mockWire) ResetTimer() bool {
	if m.ResetTimerFn != nil {
		return m.ResetTimerFn()
	}
	return true
}

func (m *mockWire) Info() map[string]RedisMessage {
	if m.InfoFn != nil {
		return m.InfoFn()
	}
	return nil
}

func (m *mockWire) Version() int {
	if m.VersionFn != nil {
		return m.VersionFn()
	}
	return 0
}

func (m *mockWire) AZ() string {
	if m.AZFn != nil {
		return m.AZFn()
	}
	return ""
}

func (m *mockWire) Error() error {
	if m == nil {
		return ErrClosing
	}
	if m.ErrorFn != nil {
		return m.ErrorFn()
	}
	return nil
}

func (m *mockWire) Close() {
	if m == nil {
		return
	}
	if m.CloseFn != nil {
		m.CloseFn()
	}
}
