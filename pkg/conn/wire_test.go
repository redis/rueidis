package conn

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strconv"
	"sync"
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
		r.t.Fatalf("redismock receive unexpected command length: expected %v, got : %v", len(expected), len(m.Values))
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

func setup(t *testing.T, option Option) (*wire, *redisMock, func(), func()) {
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
	c, err := newWire(n1, option)
	if err != nil {
		t.Fatalf("wire setup failed: %v", err)
	}
	if info := c.Info(); info["key"].String != "value" {
		t.Fatalf("wire setup failed, unexpected hello response: %v", c.Info())
	}
	return c, mock, func() {
			go func() { mock.Expect("QUIT").ReplyString("OK") }()
			c.Close()
			mock.Close()
		}, func() {
			n1.Close()
			n2.Close()
		}
}

func ExpectOK(t *testing.T, result proto.Result) {
	val, err := result.Value()
	if err != nil {
		t.Fatalf("unexpected error result: %v", err)
	}
	if str, _ := val.ToString(); str != "OK" {
		t.Fatalf("unexpected result: %v", fmt.Sprintf("%s", str))
	}
}

func TestNewConn(t *testing.T) {
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
		c, err := newWire(n1, Option{
			SelectDB:   1,
			Username:   "un",
			Password:   "pa",
			ClientName: "cn",
		})
		if err != nil {
			t.Fatalf("wire setup failed: %v", err)
		}
		go func() { mock.Expect("QUIT").ReplyString("OK") }()
		c.Close()
		mock.Close()
		n1.Close()
		n2.Close()
	})
	t.Run("Network Error", func(t *testing.T) {
		n1, n2 := net.Pipe()
		n1.Close()
		n2.Close()
		if _, err := newWire(n1, Option{}); err != io.ErrClosedPipe {
			t.Fatalf("wire setup should failed with io.ErrClosedPipe, but got %v", err)
		}
	})
}

func TestWriteSingleFlush(t *testing.T) {
	conn, mock, cancel, _ := setup(t, Option{})
	defer cancel()
	go func() { mock.Expect("PING").ReplyString("OK") }()
	ExpectOK(t, conn.Do(cmds.NewCompleted([]string{"PING"})))
}

func TestWriteMultiFlush(t *testing.T) {
	conn, mock, cancel, _ := setup(t, Option{})
	defer cancel()
	go func() {
		mock.Expect("PING").Expect("PING").ReplyString("OK").ReplyString("OK")
	}()
	for _, resp := range conn.DoMulti(cmds.NewCompleted([]string{"PING"}), cmds.NewCompleted([]string{"PING"})) {
		ExpectOK(t, resp)
	}
}

func TestResponseSequenceWithPushMessageInjected(t *testing.T) {
	conn, mock, cancel, _ := setup(t, Option{})
	defer cancel()

	times := 5000
	wg := sync.WaitGroup{}
	wg.Add(times)
	for i := 0; i < times; i++ {
		go func(i int) {
			defer wg.Done()
			v := strconv.Itoa(i)
			if val, _ := conn.Do(cmds.NewCompleted([]string{"GET", v})).Value(); val.String != v {
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
	conn, mock, cancel, _ := setup(t, Option{})
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
			if v, _ := conn.DoCache(cmds.Cacheable(cmds.NewCompleted([]string{"GET", "a"})), time.Second).Value(); v.String != "1" {
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

	// single flight
	wg.Add(5000)
	for i := 0; i < 5000; i++ {
		go func() {
			defer wg.Done()
			if v, _ := conn.DoCache(cmds.Cacheable(cmds.NewCompleted([]string{"GET", "a"})), time.Second).Value(); v.String != "2" {
				t.Errorf("unexpected non cached result, expected %v, got %v", "2", v.String)
			}
		}()
	}
	wg.Wait()
}

func TestPubSub(t *testing.T) {
	builder := cmds.NewBuilder()
	t.Run("NoReply Commands", func(t *testing.T) {
		conn, mock, cancel, _ := setup(t, Option{})
		defer cancel()

		commands := []cmds.Completed{
			builder.Subscribe().Channel("a").Build(),
			builder.Psubscribe().Pattern("b").Build(),
			builder.Unsubscribe().Channel("c").Build(),
			builder.Punsubscribe().Pattern("d").Build(),
		}

		for _, c := range commands {
			conn.Do(c)
			mock.Expect(c.Commands()...)
			go func() { mock.Expect("GET", "k").ReplyString("v") }()
			if v, _ := conn.Do(builder.Get().Key("k").Build()).Value(); v.String != "v" {
				t.Fatalf("no-reply commands should not affect nornal commands")
			}
		}
	})

	t.Run("PubSub Push Message", func(t *testing.T) {
		count := make([]int32, 4)
		_, mock, cancel, _ := setup(t, Option{
			PubSubHandlers: PubSubHandlers{
				OnMessage: func(channel, message string) {
					if channel != "1" || message != "2" {
						t.Fatalf("unexpected OnMessage")
					}
					count[0]++
				},
				OnPMessage: func(pattern, channel, message string) {
					if pattern != "3" || channel != "4" || message != "5" {
						t.Fatalf("unexpected OnPMessage")
					}
					count[1]++
				},
				OnSubscribed: func(channel string, active int64) {
					if channel != "6" || active != 7 {
						t.Fatalf("unexpected OnSubscribed")
					}
					count[2]++
				},
				OnUnSubscribed: func(channel string, active int64) {
					if channel != "8" || active != 9 {
						t.Fatalf("unexpected OnUnSubscribed")
					}
					count[3]++
				},
			},
		})
		mock.Expect().Reply(
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
			t.Fatalf("unexpected OnMessage count")
		}
		if count[1] != 1 {
			t.Fatalf("unexpected OnPMessage count")
		}
		if count[2] != 2 {
			t.Fatalf("unexpected OnSubscribed count")
		}
		if count[3] != 2 {
			t.Fatalf("unexpected OnUnSubscribed count")
		}
	})
}

func TestExitAllGoroutineOnWriteError(t *testing.T) {
	conn, _, _, closePipe := setup(t, Option{})

	closePipe()

	wg := sync.WaitGroup{}
	wg.Add(5000)
	for i := 0; i < 5000; i++ {
		go func() {
			defer wg.Done()
			if err := conn.Do(cmds.NewCompleted([]string{"GET", "a"})).NonRedisError(); err != io.ErrClosedPipe && err != ErrConnClosing {
				t.Errorf("unexpected cached result, expected io.ErrClosedPipe or ErrConnClosing, got %v", err)
			}
			if err := conn.DoMulti(cmds.NewCompleted([]string{"GET", "a"}))[0].NonRedisError(); err != io.ErrClosedPipe && err != ErrConnClosing {
				t.Errorf("unexpected cached result, expected io.ErrClosedPipe or ErrConnClosing, got %v", err)
			}
		}()
	}
	wg.Wait()
}

func TestExitAllGoroutineOnReadError(t *testing.T) {
	conn, mock, _, closePipe := setup(t, Option{})

	go func() {
		mock.Expect("GET", "a")
		closePipe()
	}()

	wg := sync.WaitGroup{}
	wg.Add(5000)
	for i := 0; i < 5000; i++ {
		go func() {
			defer wg.Done()
			if err := conn.Do(cmds.NewCompleted([]string{"GET", "a"})).NonRedisError(); err != io.ErrClosedPipe && err != ErrConnClosing {
				t.Errorf("unexpected cached result, expected io.ErrClosedPipe or ErrConnClosing, got %v", err)
			}
			if err := conn.DoMulti(cmds.NewCompleted([]string{"GET", "a"}))[0].NonRedisError(); err != io.ErrClosedPipe && err != ErrConnClosing {
				t.Errorf("unexpected cached result, expected io.ErrClosedPipe or ErrConnClosing, got %v", err)
			}
		}()
	}
	wg.Wait()
}

func TestCloseAndWaitPendingCMDs(t *testing.T) {
	conn, mock, _, _ := setup(t, Option{})
	var wg1, wg2 sync.WaitGroup
	wg1.Add(5001)
	wg2.Add(5000)
	for i := 0; i < 5000; i++ {
		go func() {
			defer wg1.Done()
			go wg2.Done()
			if v, _ := conn.Do(cmds.NewCompleted([]string{"GET", "a"})).Value(); v.String != "b" {
				t.Errorf("unexpected GET result %v", v.String)
			}
		}()
	}
	wg2.Wait()
	go func() {
		conn.Close()
		wg1.Done()
	}()
	for i := 0; i < 5000; i++ {
		mock.Expect("GET", "a").ReplyString("b")
	}
	mock.Expect("QUIT").ReplyString("OK")
	wg1.Wait()
}
