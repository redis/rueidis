package rueidis

import (
	"bufio"
	"context"
	"errors"
	"io"
	"net"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/rueian/rueidis/internal/cmds"
)

type redisExpect struct {
	*redisMock
	err error
}

type redisMock struct {
	t    *testing.T
	buf  *bufio.Reader
	conn net.Conn
}

func (r *redisMock) ReadMessage() (RedisMessage, error) {
	m, err := readNextMessage(r.buf)
	if err != nil {
		return RedisMessage{}, err
	}
	return m, nil
}

func (r *redisMock) Expect(expected ...string) *redisExpect {
	if len(expected) == 0 {
		return &redisExpect{redisMock: r}
	}
	m, err := r.ReadMessage()
	if err != nil {
		return &redisExpect{redisMock: r, err: err}
	}
	if len(expected) != len(m.values) {
		r.t.Fatalf("redismock receive unexpected command length: expected %v, got : %v", len(expected), m.values)
	}
	for i, expected := range expected {
		if m.values[i].string != expected {
			r.t.Fatalf("redismock receive unexpected command: expected %v, got : %v", expected, m.values[i])
		}
	}
	return &redisExpect{redisMock: r}
}

func (r *redisExpect) ReplyString(replies ...string) *redisExpect {
	for _, reply := range replies {
		if r.err == nil {
			r.Reply(RedisMessage{typ: '+', string: reply})
		}
	}
	return r
}

func (r *redisExpect) ReplyInteger(replies ...int64) *redisExpect {
	for _, reply := range replies {
		if r.err == nil {
			r.Reply(RedisMessage{typ: ':', integer: reply})
		}
	}
	return r
}

func (r *redisExpect) Reply(replies ...RedisMessage) *redisExpect {
	for _, reply := range replies {
		if r.err == nil {
			r.err = write(r.conn, reply)
		}
	}
	return r
}

func (r *redisMock) Close() {
	r.conn.Close()
}

func write(o io.Writer, m RedisMessage) (err error) {
	_, err = o.Write([]byte{m.typ})
	switch m.typ {
	case '+', '-', '_':
		_, err = o.Write(append([]byte(m.string), '\r', '\n'))
	case ':':
		_, err = o.Write(append([]byte(strconv.FormatInt(m.integer, 10)), '\r', '\n'))
	case '%', '>', '*':
		size := int64(len(m.values))
		if m.typ == '%' {
			if size%2 != 0 {
				panic("map message with wrong value length")
			}
			size /= 2
		}
		_, err = o.Write(append([]byte(strconv.FormatInt(size, 10)), '\r', '\n'))
		for _, v := range m.values {
			err = write(o, v)
		}
	default:
		panic("unimplemented write type")
	}
	return err
}

func setup(t *testing.T, option ClientOption) (*pipe, *redisMock, func(), func()) {
	n1, n2 := net.Pipe()
	mock := &redisMock{
		t:    t,
		buf:  bufio.NewReader(n2),
		conn: n2,
	}
	go func() {
		mock.Expect("HELLO", "3").
			Reply(RedisMessage{
				typ: '%',
				values: []RedisMessage{
					{typ: '+', string: "key"},
					{typ: '+', string: "value"},
				},
			})
		mock.Expect("CLIENT", "TRACKING", "ON", "OPTIN").
			ReplyString("OK")
	}()
	p, err := newPipe(n1, &option)
	if err != nil {
		t.Fatalf("pipe setup failed: %v", err)
	}
	if info := p.Info(); info["key"].string != "value" {
		t.Fatalf("pipe setup failed, unexpected hello response: %v", p.Info())
	}
	return p, mock, func() {
			go func() { mock.Expect("QUIT").ReplyString("OK") }()
			p.Close()
			mock.Close()
		}, func() {
			n1.Close()
			n2.Close()
		}
}

func ExpectOK(t *testing.T, result RedisResult) {
	val, err := result.ToMessage()
	if err != nil {
		t.Errorf("unexpected error result: %v", err)
	}
	if str, _ := val.ToString(); str != "OK" {
		t.Errorf("unexpected result: %v", str)
	}
}

func TestNewPipe(t *testing.T) {
	t.Run("Auth", func(t *testing.T) {
		n1, n2 := net.Pipe()
		mock := &redisMock{buf: bufio.NewReader(n2), conn: n2}
		go func() {
			mock.Expect("HELLO", "3", "AUTH", "un", "pa", "SETNAME", "cn").
				Reply(RedisMessage{
					typ:    '%',
					values: []RedisMessage{{typ: '+', string: "key"}, {typ: '+', string: "value"}},
				})
			mock.Expect("CLIENT", "TRACKING", "ON", "OPTIN").
				ReplyString("OK")
			mock.Expect("SELECT", "1").
				ReplyString("OK")
		}()
		p, err := newPipe(n1, &ClientOption{
			SelectDB:   1,
			Username:   "un",
			Password:   "pa",
			ClientName: "cn",
		})
		if err != nil {
			t.Fatalf("pipe setup failed: %v", err)
		}
		go func() { mock.Expect("QUIT").ReplyString("OK") }()
		p.Close()
		mock.Close()
		n1.Close()
		n2.Close()
	})
	t.Run("Network Error", func(t *testing.T) {
		n1, n2 := net.Pipe()
		n1.Close()
		n2.Close()
		if _, err := newPipe(n1, &ClientOption{}); err != io.ErrClosedPipe {
			t.Fatalf("pipe setup should failed with io.ErrClosedPipe, but got %v", err)
		}
	})
}

func TestWriteSingleFlush(t *testing.T) {
	p, mock, cancel, _ := setup(t, ClientOption{})
	defer cancel()
	go func() { mock.Expect("PING").ReplyString("OK") }()
	ExpectOK(t, p.Do(context.Background(), cmds.NewCompleted([]string{"PING"})))
}

func TestIgnoreOutOfBandDataDuringSyncMode(t *testing.T) {
	p, mock, cancel, _ := setup(t, ClientOption{})
	defer cancel()
	go func() {
		mock.Expect("PING").Reply(RedisMessage{typ: '>', string: "This should be ignore"}).ReplyString("OK")
	}()
	ExpectOK(t, p.Do(context.Background(), cmds.NewCompleted([]string{"PING"})))
}

func TestWriteSinglePipelineFlush(t *testing.T) {
	p, mock, cancel, _ := setup(t, ClientOption{})
	defer cancel()
	times := 2000
	wg := sync.WaitGroup{}
	wg.Add(times)

	for i := 0; i < times; i++ {
		go func() {
			ExpectOK(t, p.Do(context.Background(), cmds.NewCompleted([]string{"PING"})))
		}()
	}
	for i := 0; i < times; i++ {
		mock.Expect("PING").ReplyString("OK")
	}
}

func TestWriteMultiFlush(t *testing.T) {
	p, mock, cancel, _ := setup(t, ClientOption{})
	defer cancel()
	go func() {
		mock.Expect("PING").Expect("PING").ReplyString("OK").ReplyString("OK")
	}()
	for _, resp := range p.DoMulti(context.Background(), cmds.NewCompleted([]string{"PING"}), cmds.NewCompleted([]string{"PING"})) {
		ExpectOK(t, resp)
	}
}

func TestWriteMultiPipelineFlush(t *testing.T) {
	p, mock, cancel, _ := setup(t, ClientOption{})
	defer cancel()
	times := 2000
	wg := sync.WaitGroup{}
	wg.Add(times)

	for i := 0; i < times; i++ {
		go func() {
			for _, resp := range p.DoMulti(context.Background(), cmds.NewCompleted([]string{"PING"}), cmds.NewCompleted([]string{"PING"})) {
				ExpectOK(t, resp)
			}
		}()
	}

	for i := 0; i < times; i++ {
		mock.Expect("PING").Expect("PING").ReplyString("OK").ReplyString("OK")
	}
}

func TestPanicOnProtocolBug(t *testing.T) {
	p, mock, _, _ := setup(t, ClientOption{})

	go func() {
		mock.Expect().ReplyString("cause panic")
	}()

	defer func() {
		if v := recover(); v != protocolbug {
			t.Fatalf("should panic on protocolbug")
		}
	}()

	p._backgroundRead()
}

func TestResponseSequenceWithPushMessageInjected(t *testing.T) {
	p, mock, cancel, _ := setup(t, ClientOption{})
	defer cancel()

	times := 2000
	wg := sync.WaitGroup{}
	wg.Add(times)
	for i := 0; i < times; i++ {
		go func(i int) {
			defer wg.Done()
			v := strconv.Itoa(i)
			if val, _ := p.Do(context.Background(), cmds.NewCompleted([]string{"GET", v})).ToMessage(); val.string != v {
				t.Errorf("out of order response, expected %v, got %v", v, val.string)
			}
		}(i)
	}
	for i := 0; i < times; i++ {
		m, _ := mock.ReadMessage()
		mock.Expect().ReplyString(m.values[1].string).
			Reply(RedisMessage{typ: '>', values: []RedisMessage{{typ: '+', string: "should be ignore"}}})
	}
	wg.Wait()
}

func TestClientSideCaching(t *testing.T) {
	p, mock, cancel, _ := setup(t, ClientOption{})
	defer cancel()

	expectCSC := func(ttl int64, resp string) {
		mock.Expect("CLIENT", "CACHING", "YES").
			Expect("MULTI").
			Expect("PTTL", "a").
			Expect("GET", "a").
			Expect("EXEC").
			ReplyString("OK").
			ReplyString("OK").
			ReplyString("OK").
			ReplyString("OK").
			Reply(RedisMessage{typ: '*', values: []RedisMessage{
				{typ: ':', integer: ttl},
				{typ: '+', string: resp},
			}})
	}
	invalidateCSC := func(keys RedisMessage) {
		mock.Expect().Reply(RedisMessage{
			typ: '>',
			values: []RedisMessage{
				{typ: '+', string: "invalidate"},
				keys,
			},
		})
	}

	go func() {
		expectCSC(-1, "1")
	}()
	// single flight
	miss := uint64(0)
	hits := uint64(0)
	times := 2000
	wg := sync.WaitGroup{}
	wg.Add(times)
	for i := 0; i < times; i++ {
		go func() {
			defer wg.Done()
			v, _ := p.DoCache(context.Background(), cmds.Cacheable(cmds.NewCompleted([]string{"GET", "a"})), 10*time.Second).ToMessage()
			if v.string != "1" {
				t.Errorf("unexpected cached result, expected %v, got %v", "1", v.string)
			}
			if v.IsCacheHit() {
				atomic.AddUint64(&hits, 1)
			} else {
				atomic.AddUint64(&miss, 1)
			}
		}()
	}
	wg.Wait()

	if v := atomic.LoadUint64(&miss); v != 1 {
		t.Fatalf("unexpected cache miss count %v", v)
	}

	if v := atomic.LoadUint64(&hits); v != uint64(times-1) {
		t.Fatalf("unexpected cache hits count %v", v)
	}

	// cache invalidation
	invalidateCSC(RedisMessage{typ: '*', values: []RedisMessage{{typ: '+', string: "a"}}})
	go func() {
		expectCSC(-1, "2")
	}()

	for {
		if v, _ := p.DoCache(context.Background(), cmds.Cacheable(cmds.NewCompleted([]string{"GET", "a"})), time.Second).ToMessage(); v.string == "2" {
			break
		}
		t.Logf("waiting for invalidating")
	}

	// cache flush invalidation
	invalidateCSC(RedisMessage{typ: '_'})
	go func() {
		expectCSC(-1, "3")
	}()

	for {
		if v, _ := p.DoCache(context.Background(), cmds.Cacheable(cmds.NewCompleted([]string{"GET", "a"})), time.Second).ToMessage(); v.string == "3" {
			break
		}
		t.Logf("waiting for invalidating")
	}
}

func TestClientSideCachingExecAbort(t *testing.T) {
	p, mock, cancel, _ := setup(t, ClientOption{})
	defer cancel()

	go func() {
		mock.Expect("CLIENT", "CACHING", "YES").
			Expect("MULTI").
			Expect("PTTL", "a").
			Expect("GET", "a").
			Expect("EXEC").
			ReplyString("OK").
			ReplyString("OK").
			ReplyString("OK").
			ReplyString("OK").
			Reply(RedisMessage{typ: '_'})
	}()

	v, err := p.DoCache(context.Background(), cmds.Cacheable(cmds.NewCompleted([]string{"GET", "a"})), 10*time.Second).ToMessage()
	if !IsRedisNil(err) {
		t.Errorf("unexpected err, got %v", err)
	}
	if v.IsCacheHit() {
		t.Errorf("unexpected cache hit")
	}
}

//gocyclo:ignore
func TestPubSub(t *testing.T) {
	builder := cmds.NewBuilder(cmds.NoSlot)
	t.Run("NoReply Commands In Do", func(t *testing.T) {
		p, mock, cancel, _ := setup(t, ClientOption{})
		defer cancel()

		commands := []cmds.Completed{
			builder.Subscribe().Channel("a").Build(),
			builder.Psubscribe().Pattern("b").Build(),
			builder.Unsubscribe().Channel("c").Build(),
			builder.Punsubscribe().Pattern("d").Build(),
			builder.Ssubscribe().Channel("c").Build(),
			builder.Sunsubscribe().Channel("d").Build(),
		}

		go func() {
			for _, c := range commands {
				mock.Expect(c.Commands()...)
				mock.Expect("GET", "k").ReplyString("v")
			}
		}()

		for _, c := range commands {
			p.Do(context.Background(), c)
			if v, _ := p.Do(context.Background(), builder.Get().Key("k").Build()).ToMessage(); v.string != "v" {
				t.Fatalf("no-reply commands should not affect nornal commands")
			}
		}
	})

	t.Run("NoReply Commands In DoMulti", func(t *testing.T) {
		p, mock, cancel, _ := setup(t, ClientOption{})
		defer cancel()

		commands := []cmds.Completed{
			builder.Subscribe().Channel("a").Build(),
			builder.Psubscribe().Pattern("b").Build(),
			builder.Unsubscribe().Channel("c").Build(),
			builder.Punsubscribe().Pattern("d").Build(),
		}

		go func() {
			for _, c := range commands {
				mock.Expect(c.Commands()...)
			}
			mock.Expect("GET", "k").ReplyString("v")
		}()

		p.DoMulti(context.Background(), commands...)
		if v, _ := p.Do(context.Background(), builder.Get().Key("k").Build()).ToMessage(); v.string != "v" {
			t.Fatalf("no-reply commands should not affect nornal commands")
		}
	})

	t.Run("Prohibit NoReply with other commands In DoMulti", func(t *testing.T) {
		p, _, cancel, _ := setup(t, ClientOption{})
		defer cancel()

		commands := []cmds.Completed{
			builder.Subscribe().Channel("a").Build(),
			builder.Psubscribe().Pattern("b").Build(),
			builder.Unsubscribe().Channel("c").Build(),
			builder.Punsubscribe().Pattern("d").Build(),
		}

		for _, cmd := range commands {
			done := make(chan struct{})
			go func() {
				defer func() {
					if msg := recover(); msg != prohibitmix {
						t.Errorf("unexpected panic msg %s", msg)
					} else {
						close(done)
					}
				}()
				p.DoMulti(context.Background(), cmd, builder.Get().Key("any").Build())
			}()
			<-done
		}
	})

	t.Run("PubSub Subscribe RedisMessage", func(t *testing.T) {
		ctx := context.Background()
		p, mock, cancel, _ := setup(t, ClientOption{})

		activate := builder.Subscribe().Channel("1").Build()
		go func() {
			mock.Expect(activate.Commands()...).Reply(
				RedisMessage{typ: '>', values: []RedisMessage{
					{typ: '+', string: "message"},
					{typ: '+', string: "1"},
					{typ: '+', string: "2"},
				}},
				RedisMessage{typ: '>', values: []RedisMessage{
					{typ: '+', string: "unsubscribe"},
					{typ: '+', string: "1"},
					{typ: ':', integer: 0},
				}},
			)
		}()

		received := make(chan struct{})

		if err := p.Receive(ctx, activate, func(msg PubSubMessage) {
			if msg.Channel == "1" && msg.Message == "2" {
				close(received)
			}
		}); err != nil {
			t.Fatalf("unexpected err %v", err)
		}

		<-received

		cancel()
	})

	t.Run("PubSub SSubscribe RedisMessage", func(t *testing.T) {
		ctx := context.Background()
		p, mock, cancel, _ := setup(t, ClientOption{})

		activate := builder.Ssubscribe().Channel("1").Build()
		go func() {
			mock.Expect(activate.Commands()...).Reply(
				RedisMessage{typ: '>', values: []RedisMessage{
					{typ: '+', string: "message"},
					{typ: '+', string: "1"},
					{typ: '+', string: "2"},
				}},
				RedisMessage{typ: '>', values: []RedisMessage{
					{typ: '+', string: "unsubscribe"},
					{typ: '+', string: "1"},
					{typ: ':', integer: 0},
				}},
			)
		}()

		received := make(chan struct{})

		if err := p.Receive(ctx, activate, func(msg PubSubMessage) {
			if msg.Channel == "1" && msg.Message == "2" {
				close(received)
			}
		}); err != nil {
			t.Fatalf("unexpected err %v", err)
		}

		<-received

		cancel()
	})

	t.Run("PubSub PSubscribe RedisMessage", func(t *testing.T) {
		ctx := context.Background()
		p, mock, cancel, _ := setup(t, ClientOption{})

		activate := builder.Psubscribe().Pattern("1").Build()
		go func() {
			mock.Expect(activate.Commands()...).Reply(
				RedisMessage{typ: '>', values: []RedisMessage{
					{typ: '+', string: "pmessage"},
					{typ: '+', string: "1"},
					{typ: '+', string: "2"},
					{typ: '+', string: "3"},
				}},
				RedisMessage{typ: '>', values: []RedisMessage{
					{typ: '+', string: "punsubscribe"},
					{typ: '+', string: "1"},
					{typ: ':', integer: 0},
				}},
			)
		}()

		received := make(chan struct{})

		if err := p.Receive(ctx, activate, func(msg PubSubMessage) {
			if msg.Pattern == "1" && msg.Channel == "2" && msg.Message == "3" {
				close(received)
			}
		}); err != nil {
			t.Fatalf("unexpected err %v", err)
		}

		<-received

		cancel()
	})

	t.Run("PubSub Wrong Command RedisMessage", func(t *testing.T) {
		p, _, cancel, _ := setup(t, ClientOption{})
		defer cancel()

		defer func() {
			if !strings.Contains(recover().(string), wrongreceive) {
				t.Fatal("Receive not panic as expected")
			}
		}()

		_ = p.Receive(context.Background(), builder.Get().Key("wrong").Build(), func(msg PubSubMessage) {})
	})

	t.Run("PubSub Subscribe fail", func(t *testing.T) {
		ctx := context.Background()
		p, _, _, closePipe := setup(t, ClientOption{})
		closePipe()

		if err := p.Receive(ctx, builder.Psubscribe().Pattern("1").Build(), func(msg PubSubMessage) {}); err != io.EOF && !strings.HasPrefix(err.Error(), "io:") {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("PubSub Subscribe context cancel", func(t *testing.T) {
		ctx, ctxCancel := context.WithCancel(context.Background())
		p, mock, cancel, _ := setup(t, ClientOption{})

		activate := builder.Subscribe().Channel("1").Build()
		go func() {
			mock.Expect(activate.Commands()...).Reply(
				RedisMessage{typ: '>', values: []RedisMessage{
					{typ: '+', string: "message"},
					{typ: '+', string: "1"},
					{typ: '+', string: "2"},
				}},
			)
		}()

		if err := p.Receive(ctx, activate, func(msg PubSubMessage) {
			if msg.Channel == "1" && msg.Message == "2" {
				ctxCancel()
			}
		}); err != context.Canceled {
			t.Fatalf("unexpected err %v", err)
		}

		cancel()
	})
}

func TestExitOnWriteError(t *testing.T) {
	p, _, _, closeConn := setup(t, ClientOption{})

	closeConn()

	for i := 0; i < 2; i++ {
		if err := p.Do(context.Background(), cmds.NewCompleted([]string{"GET", "a"})).NonRedisError(); err != io.EOF && !strings.HasPrefix(err.Error(), "io:") {
			t.Errorf("unexpected cached result, expected io err, got %v", err)
		}
	}
}

func TestExitOnPubSubSubscribeWriteError(t *testing.T) {
	p, _, _, closeConn := setup(t, ClientOption{})

	activate := cmds.NewBuilder(cmds.NoSlot).Subscribe().Channel("a").Build()

	count := int64(0)
	wg := sync.WaitGroup{}
	times := 2000
	wg.Add(times)
	for i := 0; i < times; i++ {
		go func() {
			defer wg.Done()
			atomic.AddInt64(&count, 1)
			if err := p.Do(context.Background(), activate).NonRedisError(); err != io.EOF && !strings.HasPrefix(err.Error(), "io:") {
				t.Errorf("unexpected result, expected io err, got %v", err)
			}
		}()
	}
	for atomic.LoadInt64(&count) < 1000 {
		runtime.Gosched()
	}
	closeConn()
	wg.Wait()
}

func TestExitOnWriteMultiError(t *testing.T) {
	p, _, _, closeConn := setup(t, ClientOption{})

	closeConn()

	for i := 0; i < 2; i++ {
		if err := p.DoMulti(context.Background(), cmds.NewCompleted([]string{"GET", "a"}))[0].NonRedisError(); err != io.EOF && !strings.HasPrefix(err.Error(), "io:") {
			t.Errorf("unexpected result, expected io err, got %v", err)
		}
	}
}

func TestExitAllGoroutineOnWriteError(t *testing.T) {
	conn, mock, _, closeConn := setup(t, ClientOption{})

	// start the background worker
	activate := cmds.NewBuilder(cmds.NoSlot).Subscribe().Channel("a").Build()
	go conn.Do(context.Background(), activate)
	mock.Expect(activate.Commands()...)

	closeConn()
	wg := sync.WaitGroup{}
	times := 2000
	wg.Add(times)
	for i := 0; i < times; i++ {
		go func() {
			defer wg.Done()
			if err := conn.Do(context.Background(), cmds.NewCompleted([]string{"GET", "a"})).NonRedisError(); err != io.EOF && !strings.HasPrefix(err.Error(), "io:") {
				t.Errorf("unexpected result, expected io err, got %v", err)
			}
			if err := conn.DoMulti(context.Background(), cmds.NewCompleted([]string{"GET", "a"}))[0].NonRedisError(); err != io.EOF && !strings.HasPrefix(err.Error(), "io:") {
				t.Errorf("unexpected result, expected io err, got %v", err)
			}
		}()
	}
	wg.Wait()
}

func TestExitOnReadError(t *testing.T) {
	p, mock, _, closeConn := setup(t, ClientOption{})

	go func() {
		mock.Expect("GET", "a")
		closeConn()
	}()

	for i := 0; i < 2; i++ {
		if err := p.Do(context.Background(), cmds.NewCompleted([]string{"GET", "a"})).NonRedisError(); err != io.EOF && !strings.HasPrefix(err.Error(), "io:") {
			t.Errorf("unexpected result, expected io err, got %v", err)
		}
	}
}

func TestExitOnReadMultiError(t *testing.T) {
	p, mock, _, closeConn := setup(t, ClientOption{})

	go func() {
		mock.Expect("GET", "a")
		closeConn()
	}()

	for i := 0; i < 2; i++ {
		if err := p.DoMulti(context.Background(), cmds.NewCompleted([]string{"GET", "a"}))[0].NonRedisError(); err != io.EOF && !strings.HasPrefix(err.Error(), "io:") {
			t.Errorf("unexpected result, expected io err, got %v", err)
		}
	}
}

func TestExitAllGoroutineOnReadError(t *testing.T) {
	p, mock, _, closeConn := setup(t, ClientOption{})

	go func() {
		mock.Expect("GET", "a")
		closeConn()
	}()

	wg := sync.WaitGroup{}
	times := 2000
	wg.Add(times)
	for i := 0; i < times; i++ {
		go func() {
			defer wg.Done()
			if err := p.Do(context.Background(), cmds.NewCompleted([]string{"GET", "a"})).NonRedisError(); err != io.EOF && !strings.HasPrefix(err.Error(), "io:") {
				t.Errorf("unexpected result, expected io err, got %v", err)
			}
			if err := p.DoMulti(context.Background(), cmds.NewCompleted([]string{"GET", "a"}))[0].NonRedisError(); err != io.EOF && !strings.HasPrefix(err.Error(), "io:") {
				t.Errorf("unexpected result, expected io err, got %v", err)
			}
		}()
	}
	wg.Wait()
}

func TestCloseAndWaitPendingCMDs(t *testing.T) {
	p, mock, _, _ := setup(t, ClientOption{})

	var (
		loop = 2000
		wg   sync.WaitGroup
	)

	wg.Add(loop)
	for i := 0; i < loop; i++ {
		go func() {
			defer wg.Done()
			if v, _ := p.Do(context.Background(), cmds.NewCompleted([]string{"GET", "a"})).ToMessage(); v.string != "b" {
				t.Errorf("unexpected GET result %v", v.string)
			}
		}()
	}
	for i := 0; i < loop; i++ {
		r := mock.Expect("GET", "a")
		if i == loop-1 {
			go p.Close()
			time.Sleep(time.Second / 2)
		}
		r.ReplyString("b")
	}
	mock.Expect("QUIT").ReplyString("OK")
	wg.Wait()
}

func TestAlreadyCanceledContext(t *testing.T) {
	p, _, close, closeConn := setup(t, ClientOption{})
	defer closeConn()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	if err := p.Do(ctx, cmds.NewCompleted([]string{"GET", "a"})).NonRedisError(); err != context.Canceled {
		t.Fatalf("unexpected err %v", err)
	}
	if err := p.DoMulti(ctx, cmds.NewCompleted([]string{"GET", "a"}))[0].NonRedisError(); err != context.Canceled {
		t.Fatalf("unexpected err %v", err)
	}

	ctx, cancel = context.WithDeadline(context.Background(), time.Now().Add(-1*time.Second))
	cancel()

	if err := p.Do(ctx, cmds.NewCompleted([]string{"GET", "a"})).NonRedisError(); err != context.DeadlineExceeded {
		t.Fatalf("unexpected err %v", err)
	}
	if err := p.DoMulti(ctx, cmds.NewCompleted([]string{"GET", "a"}))[0].NonRedisError(); err != context.DeadlineExceeded {
		t.Fatalf("unexpected err %v", err)
	}
	close()
}

func TestOngoingDeadlineContextInSyncMode_Do(t *testing.T) {
	p, _, _, closeConn := setup(t, ClientOption{})
	defer closeConn()

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(1*time.Second/2))
	defer cancel()

	if err := p.Do(ctx, cmds.NewCompleted([]string{"GET", "a"})).NonRedisError(); err != context.DeadlineExceeded {
		t.Fatalf("unexpected err %v", err)
	}
	p.Close()
}

func TestOngoingDeadlineContextInSyncMode_DoMulti(t *testing.T) {
	p, _, _, closeConn := setup(t, ClientOption{})
	defer closeConn()

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(1*time.Second/2))
	defer cancel()

	if err := p.DoMulti(ctx, cmds.NewCompleted([]string{"GET", "a"}))[0].NonRedisError(); err != context.DeadlineExceeded {
		t.Fatalf("unexpected err %v", err)
	}
	p.Close()
}

func TestOngoingCancelContextInPipelineMode_Do(t *testing.T) {
	p, mock, close, closeConn := setup(t, ClientOption{})
	defer closeConn()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	success := int32(0)
	canceled := int32(0)

	for i := 0; i < 100; i++ {
		go func() {
			s, err := p.Do(ctx, cmds.NewCompleted([]string{"GET", "a"})).ToString()
			if s == "OK" {
				atomic.AddInt32(&success, 1)
			} else if err == context.Canceled {
				atomic.AddInt32(&canceled, 1)
			}
		}()
	}
	for i := 0; i < 50; i++ {
		mock.Expect("GET", "a").ReplyString("OK")
	}
	for atomic.LoadInt32(&success) != 50 {
		t.Logf("wait success count to be 50 %v", atomic.LoadInt32(&success))
		time.Sleep(time.Millisecond * 100)
	}
	cancel()
	if atomic.LoadInt32(&canceled) != 50 {
		t.Logf("wait canceled count to be 50 %v", atomic.LoadInt32(&canceled))
		time.Sleep(time.Millisecond * 100)
	}
	// the rest command is still send
	for i := 0; i < 50; i++ {
		mock.Expect("GET", "a").ReplyString("OK")
	}
	close()
}

func TestOngoingWriteTimeoutInPipelineMode_Do(t *testing.T) {
	p, mock, _, closeConn := setup(t, ClientOption{ConnWriteTimeout: time.Second / 2})
	defer closeConn()

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	success := int32(0)
	timeout := int32(0)

	for i := 0; i < 100; i++ {
		go func() {
			s, err := p.Do(ctx, cmds.NewCompleted([]string{"GET", "a"})).ToString()
			if s == "OK" {
				atomic.AddInt32(&success, 1)
			} else if errors.Is(err, os.ErrDeadlineExceeded) {
				atomic.AddInt32(&timeout, 1)
			}
		}()
	}
	mock.Expect("GET", "a").ReplyString("OK")
	for atomic.LoadInt32(&success) != 1 {
		t.Logf("wait success count to be 1 %v", atomic.LoadInt32(&success))
		time.Sleep(time.Millisecond * 100)
	}
	if atomic.LoadInt32(&timeout) != 99 {
		t.Logf("wait timeout count to be 99 %v", atomic.LoadInt32(&timeout))
		time.Sleep(time.Millisecond * 100)
	}
	p.Close()
}

func TestOngoingCancelContextInPipelineMode_DoMulti(t *testing.T) {
	p, mock, close, closeConn := setup(t, ClientOption{})
	defer closeConn()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	success := int32(0)
	canceled := int32(0)

	for i := 0; i < 100; i++ {
		go func() {
			s, err := p.DoMulti(ctx, cmds.NewCompleted([]string{"GET", "a"}))[0].ToString()
			if s == "OK" {
				atomic.AddInt32(&success, 1)
			} else if err == context.Canceled {
				atomic.AddInt32(&canceled, 1)
			}
		}()
	}
	for i := 0; i < 50; i++ {
		mock.Expect("GET", "a").ReplyString("OK")
	}
	for atomic.LoadInt32(&success) != 50 {
		t.Logf("wait success count to be 50 %v", atomic.LoadInt32(&success))
		time.Sleep(time.Millisecond * 100)
	}
	cancel()
	if atomic.LoadInt32(&canceled) != 50 {
		t.Logf("wait canceled count to be 50 %v", atomic.LoadInt32(&canceled))
		time.Sleep(time.Millisecond * 100)
	}
	// the rest command is still send
	for i := 0; i < 50; i++ {
		mock.Expect("GET", "a").ReplyString("OK")
	}
	close()
}

func TestOngoingWriteTimeoutInPipelineMode_DoMulti(t *testing.T) {
	p, mock, _, closeConn := setup(t, ClientOption{ConnWriteTimeout: time.Second / 2})
	defer closeConn()

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	success := int32(0)
	timeout := int32(0)

	for i := 0; i < 100; i++ {
		go func() {
			s, err := p.DoMulti(ctx, cmds.NewCompleted([]string{"GET", "a"}))[0].ToString()
			if s == "OK" {
				atomic.AddInt32(&success, 1)
			} else if errors.Is(err, os.ErrDeadlineExceeded) {
				atomic.AddInt32(&timeout, 1)
			}
		}()
	}
	mock.Expect("GET", "a").ReplyString("OK")
	for atomic.LoadInt32(&success) != 1 {
		t.Logf("wait success count to be 1 %v", atomic.LoadInt32(&success))
		time.Sleep(time.Millisecond * 100)
	}
	if atomic.LoadInt32(&timeout) != 99 {
		t.Logf("wait timeout count to be 99 %v", atomic.LoadInt32(&timeout))
		time.Sleep(time.Millisecond * 100)
	}
	p.Close()
}

func TestDeadPipe(t *testing.T) {
	ctx := context.Background()
	if err := dead.Error(); err != ErrClosing {
		t.Fatalf("unexpected err %v", err)
	}
	if err := dead.Do(ctx, cmds.NewCompleted(nil)).Error(); err != ErrClosing {
		t.Fatalf("unexpected err %v", err)
	}
	if err := dead.DoMulti(ctx, cmds.NewCompleted(nil))[0].Error(); err != ErrClosing {
		t.Fatalf("unexpected err %v", err)
	}
	if err := dead.DoCache(ctx, cmds.Cacheable(cmds.NewCompleted(nil)), time.Second).Error(); err != ErrClosing {
		t.Fatalf("unexpected err %v", err)
	}
	if err := dead.Receive(ctx, cmds.NewCompleted(nil), func(message PubSubMessage) {}); err != ErrClosing {
		t.Fatalf("unexpected err %v", err)
	}
}
