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

	"github.com/rueian/rueidis/internal/proto"
)

type redisMock struct {
	t    *testing.T
	buf  *bufio.Reader
	conn net.Conn
}

func (r *redisMock) ReadMessage() proto.Message {
	m, err := proto.ReadNextMessage(r.buf)
	if err != nil {
		r.t.Fatalf("redismock receive unexpected err while ReadNextMessage: %v", err)
	}
	return m
}

func (r *redisMock) Expect(expected []string) *redisMock {
	m := r.ReadMessage()
	if len(expected) != len(m.Values) {
		r.t.Fatalf("redismock receive unexpected command length: expected %v, got : %v", len(expected), len(m.Values))
	}
	for i, expected := range expected {
		if m.Values[i].String != expected {
			r.t.Fatalf("redismock receive unexpected command: expected %v, got : %v", expected, m.Values[i])
		}
	}
	return r
}

func (r *redisMock) ReplyString(replies ...string) *redisMock {
	for _, reply := range replies {
		r.Reply(proto.Message{Type: '+', String: reply})
	}
	return r
}

func (r *redisMock) Reply(replies ...proto.Message) *redisMock {
	for _, reply := range replies {
		if err := write(r.conn, reply); err != nil {
			r.t.Fatalf("unexpected error while setting mock replies: %v", err)
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

func setup(t *testing.T, option Option) (*Conn, *redisMock, func(), func()) {
	n1, n2 := net.Pipe()
	mock := &redisMock{
		t:    t,
		buf:  bufio.NewReader(n2),
		conn: n2,
	}
	go func() {
		mock.Expect([]string{"HELLO", "3"}).
			Reply(proto.Message{
				Type: '%',
				Values: []proto.Message{
					{Type: '+', String: "key"},
					{Type: '+', String: "value"},
				},
			})
		mock.Expect([]string{"CLIENT", "TRACKING", "ON", "OPTIN"}).
			ReplyString("OK")
	}()
	c, err := newConn(n1, option)
	if err != nil {
		t.Fatalf("conn setup failed: %v", err)
	}
	if c.Info().Values[0].String != "key" || c.Info().Values[1].String != "value" {
		t.Fatalf("conn setup failed, unexpected hello response: %v", c.Info())
	}
	return c, mock, func() {
			c.Close()
			mock.Close()
		}, func() {
			n1.Close()
			n2.Close()
		}
}

func ExpectOK(t *testing.T, result proto.Result) {
	if result.Err != nil {
		t.Fatalf("unexpected error result: %v", result.Err)
	}
	if result.Val.Type != '+' || result.Val.String != "OK" {
		t.Fatalf("unexpected result: %v", fmt.Sprintf("%c%s", result.Val.Type, result.Val.String))
	}
}

func TestWriteSingleFlush(t *testing.T) {
	conn, mock, cancel, _ := setup(t, Option{})
	defer cancel()
	go func() { mock.Expect([]string{"PING"}).ReplyString("OK") }()
	ExpectOK(t, conn.Do([]string{"PING"}))
}

func TestWriteMultiFlush(t *testing.T) {
	conn, mock, cancel, _ := setup(t, Option{})
	defer cancel()
	go func() {
		mock.Expect([]string{"PING"}).Expect([]string{"PING"})
		mock.ReplyString("OK").ReplyString("OK")
	}()
	for _, resp := range conn.DoMulti([]string{"PING"}, []string{"PING"}) {
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
			if resp := conn.Do([]string{"GET", v}).Val.String; resp != v {
				t.Errorf("out of order response, expected %v, got %v", v, resp)
			}
		}(i)
	}
	for i := 0; i < times; i++ {
		m := mock.ReadMessage()
		mock.ReplyString(m.Values[1].String)
		mock.Reply(proto.Message{Type: '>', Values: []proto.Message{{Type: '+', String: "should be ignore"}}})
	}
	wg.Wait()
}

func TestClientSideCaching(t *testing.T) {
	conn, mock, cancel, _ := setup(t, Option{})
	defer cancel()

	go func() {
		mock.Expect([]string{"CLIENT", "CACHING", "YES"}).
			Expect([]string{"GET", "a"}).
			ReplyString("OK").
			ReplyString("1")
	}()

	// single flight
	wg := sync.WaitGroup{}
	wg.Add(5000)
	for i := 0; i < 5000; i++ {
		go func() {
			defer wg.Done()
			if v := conn.DoCache([]string{"GET", "a"}, time.Second).Val.String; v != "1" {
				t.Errorf("unexpected cached result, expected %v, got %v", "1", v)
			}
		}()
	}
	wg.Wait()

	// cache invalidation
	mock.Reply(proto.Message{
		Type: '>',
		Values: []proto.Message{
			{Type: '+', String: "invalidate"},
			{Type: '*', Values: []proto.Message{{Type: '+', String: "a"}}},
		},
	})
	go func() {
		mock.Expect([]string{"CLIENT", "CACHING", "YES"}).
			Expect([]string{"GET", "a"}).
			ReplyString("OK").
			ReplyString("2")
	}()

	// single flight
	wg.Add(5000)
	for i := 0; i < 5000; i++ {
		go func() {
			defer wg.Done()
			if v := conn.DoCache([]string{"GET", "a"}, time.Second).Val.String; v != "2" {
				t.Errorf("unexpected non cached result, expected %v, got %v", "2", v)
			}
		}()
	}
	wg.Wait()
}

func TestExitAllGoroutineOnWriteError(t *testing.T) {
	conn, _, cancel, closePipe := setup(t, Option{})
	defer cancel()

	closePipe()

	wg := sync.WaitGroup{}
	wg.Add(5000)
	for i := 0; i < 5000; i++ {
		go func() {
			defer wg.Done()
			if v := conn.Do([]string{"GET", "a"}); v.Err != io.ErrClosedPipe && v.Err != ErrConnClosing {
				t.Errorf("unexpected cached result, expected io.ErrClosedPipe or ErrConnClosing, got %v", v.Err)
			}
		}()
	}
	wg.Wait()
}

func TestExitAllGoroutineOnReadError(t *testing.T) {
	conn, mock, cancel, closePipe := setup(t, Option{})
	defer cancel()

	go func() {
		mock.Expect([]string{"GET", "a"})
		closePipe()
	}()

	wg := sync.WaitGroup{}
	wg.Add(5000)
	for i := 0; i < 5000; i++ {
		go func() {
			defer wg.Done()
			if v := conn.Do([]string{"GET", "a"}); v.Err != io.ErrClosedPipe && v.Err != ErrConnClosing {
				t.Errorf("unexpected cached result, expected io.ErrClosedPipe or ErrConnClosing, got %v", v.Err)
			}
		}()
	}
	wg.Wait()
}
