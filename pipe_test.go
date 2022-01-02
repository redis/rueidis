package rueidis

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/rueian/rueidis/internal/cmds"
	"github.com/rueian/rueidis/internal/proto"
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

func (r *redisMock) ReadMessage() (proto.Message, error) {
	m, err := proto.ReadNextMessage(r.buf)
	if err != nil {
		return proto.Message{}, err
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
	if len(expected) != len(m.Values) {
		r.t.Fatalf("redismock receive unexpected command length: expected %v, got : %v", len(expected), m.Values)
	}
	for i, expected := range expected {
		if m.Values[i].String != expected {
			r.t.Fatalf("redismock receive unexpected command: expected %v, got : %v", expected, m.Values[i])
		}
	}
	return &redisExpect{redisMock: r}
}

func (r *redisExpect) ReplyString(replies ...string) *redisExpect {
	for _, reply := range replies {
		if r.err == nil {
			r.Reply(proto.Message{Type: '+', String: reply})
		}
	}
	return r
}

func (r *redisExpect) ReplyInteger(replies ...int64) *redisExpect {
	for _, reply := range replies {
		if r.err == nil {
			r.Reply(proto.Message{Type: ':', Integer: reply})
		}
	}
	return r
}

func (r *redisExpect) Reply(replies ...proto.Message) *redisExpect {
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

func write(o io.Writer, m proto.Message) (err error) {
	_, err = o.Write([]byte{m.Type})
	switch m.Type {
	case '+', '-':
		_, err = o.Write(append([]byte(m.String), '\r', '\n'))
	case ':':
		_, err = o.Write(append([]byte(strconv.FormatInt(m.Integer, 10)), '\r', '\n'))
	case '%', '>', '*':
		size := int64(len(m.Values))
		if m.Type == '%' {
			if size%2 != 0 {
				panic("map message with wrong value length")
			}
			size /= 2
		}
		_, err = o.Write(append([]byte(strconv.FormatInt(size, 10)), '\r', '\n'))
		for _, v := range m.Values {
			err = write(o, v)
		}
	default:
		panic("unimplemented write type")
	}
	return err
}

func setup(t *testing.T, option ClientOption) (*pipe, *redisMock, func(), func()) {
	return setupWithDisconnectedFn(t, option, nil)
}

func setupWithDisconnectedFn(t *testing.T, option ClientOption, onDisconnected func(err error)) (*pipe, *redisMock, func(), func()) {
	n1, n2 := net.Pipe()
	mock := &redisMock{
		t:    t,
		buf:  bufio.NewReader(n2),
		conn: n2,
	}
	go func() {
		mock.Expect("HELLO", "3").
			Reply(proto.Message{
				Type: '%',
				Values: []proto.Message{
					{Type: '+', String: "key"},
					{Type: '+', String: "value"},
				},
			})
		mock.Expect("CLIENT", "TRACKING", "ON", "OPTIN").
			ReplyString("OK")
	}()
	p, err := newPipe(n1, option, onDisconnected)
	if err != nil {
		t.Fatalf("pipe setup failed: %v", err)
	}
	if info := p.Info(); info["key"].String != "value" {
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

func ExpectOK(t *testing.T, result proto.Result) {
	val, err := result.Value()
	if err != nil {
		t.Errorf("unexpected error result: %v", err)
	}
	if str, _ := val.ToString(); str != "OK" {
		t.Errorf("unexpected result: %v", fmt.Sprintf("%s", str))
	}
}

func TestNewPipe(t *testing.T) {
	t.Run("Auth", func(t *testing.T) {
		n1, n2 := net.Pipe()
		mock := &redisMock{buf: bufio.NewReader(n2), conn: n2}
		go func() {
			mock.Expect("HELLO", "3", "AUTH", "un", "pa", "SETNAME", "cn").
				Reply(proto.Message{
					Type:   '%',
					Values: []proto.Message{{Type: '+', String: "key"}, {Type: '+', String: "value"}},
				})
			mock.Expect("CLIENT", "TRACKING", "ON", "OPTIN").
				ReplyString("OK")
			mock.Expect("SELECT", "1").
				ReplyString("OK")
		}()
		p, err := newPipe(n1, ClientOption{
			SelectDB:   1,
			Username:   "un",
			Password:   "pa",
			ClientName: "cn",
		}, nil)
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
		if _, err := newPipe(n1, ClientOption{}, nil); err != io.ErrClosedPipe {
			t.Fatalf("pipe setup should failed with io.ErrClosedPipe, but got %v", err)
		}
	})
}

func TestOnDisconnectedHook(t *testing.T) {
	done := make(chan struct{})
	p, _, _, closeConn := setupWithDisconnectedFn(t, ClientOption{}, func(err error) {
		if err != io.EOF && !strings.HasPrefix(err.Error(), "io:") {
			t.Errorf("unexpected err %v", err)
		}
		close(done)
	})
	closeConn()
	if err := p.Do(cmds.NewCompleted([]string{"PING"})).Error(); err != io.EOF && !strings.HasPrefix(err.Error(), "io:") {
		t.Errorf("unexpected err %v", err)
	}
	<-done
}

func TestWriteSingleFlush(t *testing.T) {
	p, mock, cancel, _ := setup(t, ClientOption{})
	defer cancel()
	go func() { mock.Expect("PING").ReplyString("OK") }()
	ExpectOK(t, p.Do(cmds.NewCompleted([]string{"PING"})))
}

func TestIgnoreOutOfBandDataDuringSyncMode(t *testing.T) {
	p, mock, cancel, _ := setup(t, ClientOption{})
	defer cancel()
	go func() {
		mock.Expect("PING").Reply(proto.Message{Type: '>', String: "This should be ignore"}).ReplyString("OK")
	}()
	ExpectOK(t, p.Do(cmds.NewCompleted([]string{"PING"})))
}

func TestWriteSinglePipelineFlush(t *testing.T) {
	p, mock, cancel, _ := setup(t, ClientOption{})
	defer cancel()
	times := 5000
	wg := sync.WaitGroup{}
	wg.Add(times)

	for i := 0; i < times; i++ {
		go func() {
			ExpectOK(t, p.Do(cmds.NewCompleted([]string{"PING"})))
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
	for _, resp := range p.DoMulti(cmds.NewCompleted([]string{"PING"}), cmds.NewCompleted([]string{"PING"})) {
		ExpectOK(t, resp)
	}
}

func TestWriteMultiPipelineFlush(t *testing.T) {
	p, mock, cancel, _ := setup(t, ClientOption{})
	defer cancel()
	times := 5000
	wg := sync.WaitGroup{}
	wg.Add(times)

	for i := 0; i < times; i++ {
		go func() {
			for _, resp := range p.DoMulti(cmds.NewCompleted([]string{"PING"}), cmds.NewCompleted([]string{"PING"})) {
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

	times := 5000
	wg := sync.WaitGroup{}
	wg.Add(times)
	for i := 0; i < times; i++ {
		go func(i int) {
			defer wg.Done()
			v := strconv.Itoa(i)
			if val, _ := p.Do(cmds.NewCompleted([]string{"GET", v})).Value(); val.String != v {
				t.Errorf("out of order response, expected %v, got %v", v, val.String)
			}
		}(i)
	}
	for i := 0; i < times; i++ {
		m, _ := mock.ReadMessage()
		mock.Expect().ReplyString(m.Values[1].String).
			Reply(proto.Message{Type: '>', Values: []proto.Message{{Type: '+', String: "should be ignore"}}})
	}
	wg.Wait()
}

func TestClientSideCaching(t *testing.T) {
	p, mock, cancel, _ := setup(t, ClientOption{})
	defer cancel()

	go func() {
		mock.Expect("CLIENT", "CACHING", "YES").
			Expect("GET", "a").
			Expect("PTTL", "a").
			ReplyString("OK").
			ReplyString("1").
			ReplyInteger(-1)
	}()

	// single flight
	wg := sync.WaitGroup{}
	wg.Add(5000)
	for i := 0; i < 5000; i++ {
		go func() {
			defer wg.Done()
			if v, _ := p.DoCache(cmds.Cacheable(cmds.NewCompleted([]string{"GET", "a"})), 10*time.Second).Value(); v.String != "1" {
				t.Errorf("unexpected cached result, expected %v, got %v", "1", v.String)
			}
		}()
	}
	wg.Wait()

	// cache invalidation
	mock.Expect().Reply(proto.Message{
		Type: '>',
		Values: []proto.Message{
			{Type: '+', String: "invalidate"},
			{Type: '*', Values: []proto.Message{{Type: '+', String: "a"}}},
		},
	})
	go func() {
		mock.Expect("CLIENT", "CACHING", "YES").
			Expect("GET", "a").
			Expect("PTTL", "a").
			ReplyString("OK").
			ReplyString("2").
			ReplyInteger(-1)
	}()

	for {
		if v, _ := p.DoCache(cmds.Cacheable(cmds.NewCompleted([]string{"GET", "a"})), time.Second).Value(); v.String != "2" {
			t.Logf("waiting for invalidating")
			continue
		}
		break
	}
}

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
		}

		go func() {
			for _, c := range commands {
				mock.Expect(c.Commands()...)
				mock.Expect("GET", "k").ReplyString("v")
			}
		}()

		for _, c := range commands {
			p.Do(c)
			if v, _ := p.Do(builder.Get().Key("k").Build()).Value(); v.String != "v" {
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

		p.DoMulti(commands...)
		if v, _ := p.Do(builder.Get().Key("k").Build()).Value(); v.String != "v" {
			t.Fatalf("no-reply commands should not affect nornal commands")
		}
	})

	t.Run("PubSub Push Message", func(t *testing.T) {
		count := make([]int32, 4)
		p, mock, cancel, _ := setup(t, ClientOption{
			PubSubOption: PubSubOption{
				onMessage: func(channel, message string) {
					if channel != "1" || message != "2" {
						t.Fatalf("unexpected onMessage")
					}
					count[0]++
				},
				onPMessage: func(pattern, channel, message string) {
					if pattern != "3" || channel != "4" || message != "5" {
						t.Fatalf("unexpected onPMessage")
					}
					count[1]++
				},
				onSubscribed: func(channel string, active int64) {
					if channel != "6" || active != 7 {
						t.Fatalf("unexpected onSubscribed")
					}
					count[2]++
				},
				onUnSubscribed: func(channel string, active int64) {
					if channel != "8" || active != 9 {
						t.Fatalf("unexpected onUnSubscribed")
					}
					count[3]++
				},
			},
		})
		activate := builder.Subscribe().Channel("a").Build()
		go p.Do(activate)
		mock.Expect(activate.Commands()...).Reply(
			proto.Message{Type: '>', Values: []proto.Message{
				{Type: '+', String: "message"},
				{Type: '+', String: "1"},
				{Type: '+', String: "2"},
			}},
			proto.Message{Type: '>', Values: []proto.Message{
				{Type: '+', String: "pmessage"},
				{Type: '+', String: "3"},
				{Type: '+', String: "4"},
				{Type: '+', String: "5"},
			}},
			proto.Message{Type: '>', Values: []proto.Message{
				{Type: '+', String: "subscribe"},
				{Type: '+', String: "6"},
				{Type: ':', Integer: 7},
			}},
			proto.Message{Type: '>', Values: []proto.Message{
				{Type: '+', String: "psubscribe"},
				{Type: '+', String: "6"},
				{Type: ':', Integer: 7},
			}},
			proto.Message{Type: '>', Values: []proto.Message{
				{Type: '+', String: "unsubscribe"},
				{Type: '+', String: "8"},
				{Type: ':', Integer: 9},
			}},
			proto.Message{Type: '>', Values: []proto.Message{
				{Type: '+', String: "punsubscribe"},
				{Type: '+', String: "8"},
				{Type: ':', Integer: 9},
			}},
		)
		cancel()
		if count[0] != 1 {
			t.Fatalf("unexpected onMessage count")
		}
		if count[1] != 1 {
			t.Fatalf("unexpected onPMessage count")
		}
		if count[2] != 2 {
			t.Fatalf("unexpected onSubscribed count")
		}
		if count[3] != 2 {
			t.Fatalf("unexpected onUnSubscribed count")
		}
	})
}

func TestExitOnWriteError(t *testing.T) {
	p, _, _, closeConn := setup(t, ClientOption{})

	closeConn()

	for i := 0; i < 2; i++ {
		if err := p.Do(cmds.NewCompleted([]string{"GET", "a"})).NonRedisError(); err != io.EOF && !strings.HasPrefix(err.Error(), "io:") {
			t.Errorf("unexpected cached result, expected io err, got %v", err)
		}
	}
}

func TestExitOnPubSubSubscribeWriteError(t *testing.T) {
	p, _, _, closeConn := setup(t, ClientOption{})

	activate := cmds.NewBuilder(cmds.NoSlot).Subscribe().Channel("a").Build()

	count := int64(0)
	wg := sync.WaitGroup{}
	wg.Add(5000)
	for i := 0; i < 5000; i++ {
		go func() {
			defer wg.Done()
			atomic.AddInt64(&count, 1)
			if err := p.Do(activate).NonRedisError(); err != io.EOF && !strings.HasPrefix(err.Error(), "io:") {
				t.Errorf("unexpected result, expected io err, got %v", err)
			}
		}()
	}
	for atomic.LoadInt64(&count) != 5000 {
		runtime.Gosched()
	}
	closeConn()
	wg.Wait()
}

func TestExitOnWriteMultiError(t *testing.T) {
	p, _, _, closeConn := setup(t, ClientOption{})

	closeConn()

	for i := 0; i < 2; i++ {
		if err := p.DoMulti(cmds.NewCompleted([]string{"GET", "a"}))[0].NonRedisError(); err != io.EOF && !strings.HasPrefix(err.Error(), "io:") {
			t.Errorf("unexpected result, expected io err, got %v", err)
		}
	}
}

func TestExitAllGoroutineOnWriteError(t *testing.T) {
	conn, mock, _, closeConn := setup(t, ClientOption{})

	// start the background worker
	activate := cmds.NewBuilder(cmds.NoSlot).Subscribe().Channel("a").Build()
	go conn.Do(activate)
	mock.Expect(activate.Commands()...)

	closeConn()
	wg := sync.WaitGroup{}
	wg.Add(5000)
	for i := 0; i < 5000; i++ {
		go func() {
			defer wg.Done()
			if err := conn.Do(cmds.NewCompleted([]string{"GET", "a"})).NonRedisError(); err != io.EOF && !strings.HasPrefix(err.Error(), "io:") {
				t.Errorf("unexpected result, expected io err, got %v", err)
			}
			if err := conn.DoMulti(cmds.NewCompleted([]string{"GET", "a"}))[0].NonRedisError(); err != io.EOF && !strings.HasPrefix(err.Error(), "io:") {
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
		if err := p.Do(cmds.NewCompleted([]string{"GET", "a"})).NonRedisError(); err != io.EOF && !strings.HasPrefix(err.Error(), "io:") {
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
		if err := p.DoMulti(cmds.NewCompleted([]string{"GET", "a"}))[0].NonRedisError(); err != io.EOF && !strings.HasPrefix(err.Error(), "io:") {
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
	wg.Add(5000)
	for i := 0; i < 5000; i++ {
		go func() {
			defer wg.Done()
			if err := p.Do(cmds.NewCompleted([]string{"GET", "a"})).NonRedisError(); err != io.EOF && !strings.HasPrefix(err.Error(), "io:") {
				t.Errorf("unexpected result, expected io err, got %v", err)
			}
			if err := p.DoMulti(cmds.NewCompleted([]string{"GET", "a"}))[0].NonRedisError(); err != io.EOF && !strings.HasPrefix(err.Error(), "io:") {
				t.Errorf("unexpected result, expected io err, got %v", err)
			}
		}()
	}
	wg.Wait()
}

func TestCloseAndWaitPendingCMDs(t *testing.T) {
	p, mock, _, _ := setup(t, ClientOption{})

	var (
		loop = 5000
		wg   sync.WaitGroup
	)

	wg.Add(loop)
	for i := 0; i < loop; i++ {
		go func() {
			defer wg.Done()

			if v, _ := p.Do(cmds.NewCompleted([]string{"GET", "a"})).Value(); v.String != "b" {
				t.Errorf("unexpected GET result %v", v.String)
			}
		}()
	}
	for i := 0; i < loop; i++ {
		r := mock.Expect("GET", "a")
		if i == loop-1 {
			go p.Close()
		}
		r.ReplyString("b")
	}
	mock.Expect("QUIT").ReplyString("OK")
	wg.Wait()
}
