package rueidis

import (
	"bufio"
	"bytes"
	"context"
	"errors"
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

	"github.com/redis/rueidis/internal/cmds"
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

func (r *redisExpect) ReplyBlobString(replies ...string) *redisExpect {
	for _, reply := range replies {
		if r.err == nil {
			r.Reply(RedisMessage{typ: '$', string: reply})
		}
	}
	return r
}

func (r *redisExpect) ReplyError(replies ...string) *redisExpect {
	for _, reply := range replies {
		if r.err == nil {
			r.Reply(RedisMessage{typ: '-', string: reply})
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
	case '$':
		_, _ = o.Write(append([]byte(strconv.Itoa(len(m.string))), '\r', '\n'))
		_, err = o.Write(append([]byte(m.string), '\r', '\n'))
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
	if option.CacheSizeEachConn <= 0 {
		option.CacheSizeEachConn = DefaultCacheBytes
	}
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
					{typ: '+', string: "version"},
					{typ: '+', string: "6.0.0"},
					{typ: '+', string: "proto"},
					{typ: ':', integer: 3},
				},
			})
		if !option.DisableCache {
			mock.Expect("CLIENT", "TRACKING", "ON", "OPTIN").
				ReplyString("OK")
		}
		mock.Expect("CLIENT", "SETINFO", "LIB-NAME", LibName).
			ReplyError("UNKNOWN COMMAND")
		mock.Expect("CLIENT", "SETINFO", "LIB-VER", LibVer).
			ReplyError("UNKNOWN COMMAND")
	}()
	p, err := newPipe(func() (net.Conn, error) { return n1, nil }, &option)
	if err != nil {
		t.Fatalf("pipe setup failed: %v", err)
	}
	if info := p.Info(); info["version"].string != "6.0.0" {
		t.Fatalf("pipe setup failed, unexpected hello response: %v", p.Info())
	}
	if version := p.Version(); version != 6 {
		t.Fatalf("pipe setup failed, unexpected version: %v", p.Version())
	}
	return p, mock, func() {
			go func() { mock.Expect("PING").ReplyString("OK") }()
			p.Close()
			mock.Close()
			for atomic.LoadInt32(&p.state) != 4 {
				t.Log("wait the pipe to be closed")
				time.Sleep(time.Millisecond * 100)
			}
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
	defer ShouldNotLeaked(SetupLeakDetection())
	t.Run("Auth without Username", func(t *testing.T) {
		n1, n2 := net.Pipe()
		mock := &redisMock{buf: bufio.NewReader(n2), conn: n2, t: t}
		go func() {
			mock.Expect("HELLO", "3", "AUTH", "default", "pa", "SETNAME", "cn").
				Reply(RedisMessage{
					typ: '%',
					values: []RedisMessage{
						{typ: '+', string: "proto"},
						{typ: ':', integer: 3},
					},
				})
			mock.Expect("CLIENT", "TRACKING", "ON", "OPTIN").
				ReplyString("OK")
			mock.Expect("SELECT", "1").
				ReplyString("OK")
			mock.Expect("CLIENT", "NO-TOUCH", "ON").
				ReplyString("OK")
			mock.Expect("CLIENT", "NO-EVICT", "ON").
				ReplyString("OK")
			mock.Expect("CLIENT", "SETINFO", "LIB-NAME", "libname").
				ReplyString("OK")
			mock.Expect("CLIENT", "SETINFO", "LIB-VER", "1").
				ReplyString("OK")
		}()
		p, err := newPipe(func() (net.Conn, error) { return n1, nil }, &ClientOption{
			SelectDB:      1,
			Password:      "pa",
			ClientName:    "cn",
			ClientNoEvict: true,
			ClientSetInfo: []string{"libname", "1"},
			ClientNoTouch: true,
		})
		if err != nil {
			t.Fatalf("pipe setup failed: %v", err)
		}
		go func() { mock.Expect("PING").ReplyString("OK") }()
		p.Close()
		mock.Close()
		n1.Close()
		n2.Close()
	})
	t.Run("AlwaysRESP2", func(t *testing.T) {
		n1, n2 := net.Pipe()
		mock := &redisMock{buf: bufio.NewReader(n2), conn: n2, t: t}
		go func() {
			mock.Expect("AUTH", "pa").
				ReplyString("OK")
			mock.Expect("CLIENT", "SETNAME", "cn").
				ReplyString("OK")
			mock.Expect("SELECT", "1").
				ReplyString("OK")
			mock.Expect("CLIENT", "NO-TOUCH", "ON").
				ReplyString("OK")
			mock.Expect("CLIENT", "NO-EVICT", "ON").
				ReplyString("OK")
			mock.Expect("CLIENT", "SETINFO", "LIB-NAME", "libname").
				ReplyString("OK")
			mock.Expect("CLIENT", "SETINFO", "LIB-VER", "1").
				ReplyString("OK")
		}()
		p, err := newPipe(func() (net.Conn, error) { return n1, nil }, &ClientOption{
			SelectDB:      1,
			Password:      "pa",
			ClientName:    "cn",
			ClientNoEvict: true,
			ClientSetInfo: []string{"libname", "1"},
			ClientNoTouch: true,
			AlwaysRESP2:   true,
			DisableCache:  true,
		})
		if err != nil {
			t.Fatalf("pipe setup failed: %v", err)
		}
		go func() { mock.Expect("PING").ReplyString("OK") }()
		p.Close()
		mock.Close()
		n1.Close()
		n2.Close()
	})
	t.Run("Auth with Username", func(t *testing.T) {
		n1, n2 := net.Pipe()
		mock := &redisMock{buf: bufio.NewReader(n2), conn: n2, t: t}
		go func() {
			mock.Expect("HELLO", "3", "AUTH", "ua", "pa", "SETNAME", "cn").
				Reply(RedisMessage{
					typ: '%',
					values: []RedisMessage{
						{typ: '+', string: "proto"},
						{typ: ':', integer: 3},
					},
				})
			mock.Expect("CLIENT", "TRACKING", "ON", "OPTIN").
				ReplyString("OK")
			mock.Expect("SELECT", "1").
				ReplyString("OK")
			mock.Expect("CLIENT", "SETINFO", "LIB-NAME", LibName).
				ReplyError("UNKNOWN COMMAND")
			mock.Expect("CLIENT", "SETINFO", "LIB-VER", LibVer).
				ReplyError("UNKNOWN COMMAND")
		}()
		p, err := newPipe(func() (net.Conn, error) { return n1, nil }, &ClientOption{
			SelectDB:   1,
			Username:   "ua",
			Password:   "pa",
			ClientName: "cn",
		})
		if err != nil {
			t.Fatalf("pipe setup failed: %v", err)
		}
		go func() { mock.Expect("PING").ReplyString("OK") }()
		p.Close()
		mock.Close()
		n1.Close()
		n2.Close()
	})
	t.Run("Auth with Credentials Function", func(t *testing.T) {
		n1, n2 := net.Pipe()
		mock := &redisMock{buf: bufio.NewReader(n2), conn: n2, t: t}
		go func() {
			mock.Expect("HELLO", "3", "AUTH", "ua", "pa", "SETNAME", "cn").
				Reply(RedisMessage{
					typ: '%',
					values: []RedisMessage{
						{typ: '+', string: "proto"},
						{typ: ':', integer: 3},
					},
				})
			mock.Expect("CLIENT", "TRACKING", "ON", "OPTIN").
				ReplyString("OK")
			mock.Expect("SELECT", "1").
				ReplyString("OK")
			mock.Expect("CLIENT", "SETINFO", "LIB-NAME", LibName).
				ReplyError("UNKNOWN COMMAND")
			mock.Expect("CLIENT", "SETINFO", "LIB-VER", LibVer).
				ReplyError("UNKNOWN COMMAND")
		}()
		p, err := newPipe(func() (net.Conn, error) { return n1, nil }, &ClientOption{
			SelectDB: 1,
			AuthCredentialsFn: func(context AuthCredentialsContext) (AuthCredentials, error) {
				return AuthCredentials{
					Username: "ua",
					Password: "pa",
				}, nil
			},
			ClientName: "cn",
		})
		if err != nil {
			t.Fatalf("pipe setup failed: %v", err)
		}
		go func() { mock.Expect("PING").ReplyString("OK") }()
		p.Close()
		mock.Close()
		n1.Close()
		n2.Close()
	})
	t.Run("With ClientSideTrackingOptions", func(t *testing.T) {
		n1, n2 := net.Pipe()
		mock := &redisMock{buf: bufio.NewReader(n2), conn: n2, t: t}
		go func() {
			mock.Expect("HELLO", "3").
				Reply(RedisMessage{
					typ: '%',
					values: []RedisMessage{
						{typ: '+', string: "proto"},
						{typ: ':', integer: 3},
					},
				})
			mock.Expect("CLIENT", "TRACKING", "ON", "OPTIN", "NOLOOP").
				ReplyString("OK")
			mock.Expect("CLIENT", "SETINFO", "LIB-NAME", LibName).
				ReplyError("UNKNOWN COMMAND")
			mock.Expect("CLIENT", "SETINFO", "LIB-VER", LibVer).
				ReplyError("UNKNOWN COMMAND")
		}()
		p, err := newPipe(func() (net.Conn, error) { return n1, nil }, &ClientOption{
			ClientTrackingOptions: []string{"OPTIN", "NOLOOP"},
		})
		if err != nil {
			t.Fatalf("pipe setup failed: %v", err)
		}
		go func() { mock.Expect("PING").ReplyString("OK") }()
		p.Close()
		mock.Close()
		n1.Close()
		n2.Close()
	})
	t.Run("Init with ReplicaOnly", func(t *testing.T) {
		n1, n2 := net.Pipe()
		mock := &redisMock{buf: bufio.NewReader(n2), conn: n2, t: t}
		go func() {
			mock.Expect("HELLO", "3", "AUTH", "ua", "pa", "SETNAME", "cn").
				Reply(RedisMessage{
					typ: '%',
					values: []RedisMessage{
						{typ: '+', string: "proto"},
						{typ: ':', integer: 3},
					},
				})
			mock.Expect("CLIENT", "TRACKING", "ON", "OPTIN").
				ReplyString("OK")
			mock.Expect("SELECT", "1").
				ReplyString("OK")
			mock.Expect("READONLY").
				ReplyString("OK")
			mock.Expect("CLIENT", "SETINFO", "LIB-NAME", LibName).
				ReplyError("UNKNOWN COMMAND")
			mock.Expect("CLIENT", "SETINFO", "LIB-VER", LibVer).
				ReplyError("UNKNOWN COMMAND")
		}()
		p, err := newPipe(func() (net.Conn, error) { return n1, nil }, &ClientOption{
			SelectDB:    1,
			Username:    "ua",
			Password:    "pa",
			ClientName:  "cn",
			ReplicaOnly: true,
		})
		if err != nil {
			t.Fatalf("pipe setup failed: %v", err)
		}
		go func() { mock.Expect("PING").ReplyString("OK") }()
		p.Close()
		mock.Close()
		n1.Close()
		n2.Close()
	})
	t.Run("Init with ReplicaOnly ignores READONLY Error", func(t *testing.T) {
		n1, n2 := net.Pipe()
		mock := &redisMock{buf: bufio.NewReader(n2), conn: n2, t: t}
		go func() {
			mock.Expect("HELLO", "3", "AUTH", "ua", "pa", "SETNAME", "cn").
				Reply(RedisMessage{
					typ: '%',
					values: []RedisMessage{
						{typ: '+', string: "proto"},
						{typ: ':', integer: 3},
					},
				})
			mock.Expect("CLIENT", "TRACKING", "ON", "OPTIN").
				ReplyString("OK")
			mock.Expect("SELECT", "1").
				ReplyString("OK")
			mock.Expect("READONLY").
				ReplyError("This instance has cluster support disabled")
			mock.Expect("CLIENT", "SETINFO", "LIB-NAME", LibName).
				ReplyError("UNKNOWN COMMAND")
			mock.Expect("CLIENT", "SETINFO", "LIB-VER", LibVer).
				ReplyError("UNKNOWN COMMAND")
		}()
		p, err := newPipe(func() (net.Conn, error) { return n1, nil }, &ClientOption{
			SelectDB:    1,
			Username:    "ua",
			Password:    "pa",
			ClientName:  "cn",
			ReplicaOnly: true,
		})
		if err != nil {
			t.Fatalf("pipe setup failed: %v", err)
		}
		go func() { mock.Expect("PING").ReplyString("OK") }()
		p.Close()
		mock.Close()
		n1.Close()
		n2.Close()
	})
	t.Run("Network Error", func(t *testing.T) {
		n1, n2 := net.Pipe()
		n1.Close()
		n2.Close()
		if _, err := newPipe(func() (net.Conn, error) { return n1, nil }, &ClientOption{}); err != io.ErrClosedPipe {
			t.Fatalf("pipe setup should failed with io.ErrClosedPipe, but got %v", err)
		}
	})
	t.Run("Auth Credentials Function Error", func(t *testing.T) {
		n1, n2 := net.Pipe()
		mock := &redisMock{buf: bufio.NewReader(n2), conn: n2, t: t}
		go func() { mock.Expect("PING").ReplyString("OK") }()
		_, err := newPipe(func() (net.Conn, error) { return n1, nil }, &ClientOption{
			SelectDB: 1,
			AuthCredentialsFn: func(context AuthCredentialsContext) (AuthCredentials, error) {
				return AuthCredentials{}, fmt.Errorf("auth credential failure")
			},
			ClientName: "cn",
		})
		if err.Error() != "auth credential failure" {
			t.Fatalf("pipe setup failed: %v", err)
		}
		mock.Close()
		n1.Close()
		n2.Close()
	})
}

func TestNewRESP2Pipe(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	t.Run("Without DisableCache", func(t *testing.T) {
		n1, n2 := net.Pipe()
		mock := &redisMock{buf: bufio.NewReader(n2), conn: n2, t: t}
		go func() {
			mock.Expect("HELLO", "3").
				ReplyError("ERR unknown command `HELLO`")
			mock.Expect("CLIENT", "TRACKING", "ON", "OPTIN").
				ReplyError("ERR unknown subcommand or wrong number of arguments for 'TRACKING'. Try CLIENT HELP")
			mock.Expect("CLIENT", "SETINFO", "LIB-NAME", LibName).
				ReplyError("UNKNOWN COMMAND")
			mock.Expect("CLIENT", "SETINFO", "LIB-VER", LibVer).
				ReplyError("UNKNOWN COMMAND")
			mock.Expect("PING").ReplyString("OK")
		}()
		if _, err := newPipe(func() (net.Conn, error) { return n1, nil }, &ClientOption{}); !errors.Is(err, ErrNoCache) {
			t.Fatalf("unexpected err: %v", err)
		}
		mock.Close()
		n1.Close()
		n2.Close()
	})
	t.Run("Without DisableCache 2", func(t *testing.T) {
		n1, n2 := net.Pipe()
		mock := &redisMock{buf: bufio.NewReader(n2), conn: n2, t: t}
		go func() {
			mock.Expect("HELLO", "3").
				ReplyError("ERR unknown command `HELLO`")
			mock.Expect("CLIENT", "TRACKING", "ON", "OPTIN").
				ReplyString("OK")
			mock.Expect("CLIENT", "SETINFO", "LIB-NAME", LibName).
				ReplyError("UNKNOWN COMMAND")
			mock.Expect("CLIENT", "SETINFO", "LIB-VER", LibVer).
				ReplyError("UNKNOWN COMMAND")
			mock.Expect("PING").ReplyString("OK")
		}()
		if _, err := newPipe(func() (net.Conn, error) { return n1, nil }, &ClientOption{}); !errors.Is(err, ErrNoCache) {
			t.Fatalf("unexpected err: %v", err)
		}
		mock.Close()
		n1.Close()
		n2.Close()
	})
	t.Run("With Hello Proto 2", func(t *testing.T) { // kvrocks version 2.2.0
		n1, n2 := net.Pipe()
		mock := &redisMock{buf: bufio.NewReader(n2), conn: n2, t: t}
		go func() {
			mock.Expect("HELLO", "3").
				Reply(RedisMessage{typ: '*', values: []RedisMessage{
					{typ: '+', string: "server"},
					{typ: '+', string: "redis"},
					{typ: '+', string: "proto"},
					{typ: ':', integer: 2},
				}})
			mock.Expect("CLIENT", "SETINFO", "LIB-NAME", LibName).
				ReplyError("UNKNOWN COMMAND")
			mock.Expect("CLIENT", "SETINFO", "LIB-VER", LibVer).
				ReplyError("UNKNOWN COMMAND")
		}()
		p, err := newPipe(func() (net.Conn, error) { return n1, nil }, &ClientOption{
			DisableCache: true,
		})
		if err != nil {
			t.Fatalf("pipe setup failed: %v", err)
		}
		if p.version >= 6 {
			t.Fatalf("unexpected p.version: %v", p.version)
		}
		go func() { mock.Expect("PING").ReplyString("OK") }()
		p.Close()
		mock.Close()
		n1.Close()
		n2.Close()
	})
	t.Run("Auth without Username", func(t *testing.T) {
		n1, n2 := net.Pipe()
		mock := &redisMock{buf: bufio.NewReader(n2), conn: n2, t: t}
		go func() {
			mock.Expect("HELLO", "3", "AUTH", "default", "pa", "SETNAME", "cn").
				ReplyError("ERR unknown command `HELLO`")
			mock.Expect("SELECT", "1").
				ReplyError("ERR ACL")
			mock.Expect("CLIENT", "SETINFO", "LIB-NAME", LibName).
				ReplyError("UNKNOWN COMMAND")
			mock.Expect("CLIENT", "SETINFO", "LIB-VER", LibVer).
				ReplyError("UNKNOWN COMMAND")
			mock.Expect("AUTH", "pa").
				ReplyString("OK")
			mock.Expect("CLIENT", "SETNAME", "cn").
				ReplyString("OK")
			mock.Expect("SELECT", "1").
				ReplyString("OK")
			mock.Expect("CLIENT", "SETINFO", "LIB-NAME", LibName).
				ReplyError("UNKNOWN COMMAND")
			mock.Expect("CLIENT", "SETINFO", "LIB-VER", LibVer).
				ReplyError("UNKNOWN COMMAND")
		}()
		p, err := newPipe(func() (net.Conn, error) { return n1, nil }, &ClientOption{
			SelectDB:     1,
			Password:     "pa",
			ClientName:   "cn",
			DisableCache: true,
		})
		if err != nil {
			t.Fatalf("pipe setup failed: %v", err)
		}
		if p.version >= 6 {
			t.Fatalf("unexpected p.version: %v", p.version)
		}
		go func() { mock.Expect("PING").ReplyString("OK") }()
		p.Close()
		mock.Close()
		n1.Close()
		n2.Close()
	})
	t.Run("Auth with Username", func(t *testing.T) {
		n1, n2 := net.Pipe()
		mock := &redisMock{buf: bufio.NewReader(n2), conn: n2, t: t}
		go func() {
			mock.Expect("HELLO", "3", "AUTH", "ua", "pa", "SETNAME", "cn").
				ReplyError("ERR unknown command `HELLO`")
			mock.Expect("SELECT", "1").
				ReplyError("ERR ACL")
			mock.Expect("CLIENT", "SETINFO", "LIB-NAME", LibName).
				ReplyError("UNKNOWN COMMAND")
			mock.Expect("CLIENT", "SETINFO", "LIB-VER", LibVer).
				ReplyError("UNKNOWN COMMAND")
			mock.Expect("AUTH", "ua", "pa").
				ReplyString("OK")
			mock.Expect("CLIENT", "SETNAME", "cn").
				ReplyString("OK")
			mock.Expect("SELECT", "1").
				ReplyString("OK")
			mock.Expect("CLIENT", "SETINFO", "LIB-NAME", LibName).
				ReplyError("UNKNOWN COMMAND")
			mock.Expect("CLIENT", "SETINFO", "LIB-VER", LibVer).
				ReplyError("UNKNOWN COMMAND")
		}()
		p, err := newPipe(func() (net.Conn, error) { return n1, nil }, &ClientOption{
			SelectDB:     1,
			Username:     "ua",
			Password:     "pa",
			ClientName:   "cn",
			DisableCache: true,
		})
		if err != nil {
			t.Fatalf("pipe setup failed: %v", err)
		}
		if p.version >= 6 {
			t.Fatalf("unexpected p.version: %v", p.version)
		}
		go func() { mock.Expect("PING").ReplyString("OK") }()
		p.Close()
		mock.Close()
		n1.Close()
		n2.Close()
	})
	t.Run("Init with ReplicaOnly", func(t *testing.T) {
		n1, n2 := net.Pipe()
		mock := &redisMock{buf: bufio.NewReader(n2), conn: n2, t: t}
		go func() {
			mock.Expect("HELLO", "3", "AUTH", "default", "pa", "SETNAME", "cn").
				ReplyError("ERR unknown command `HELLO`")
			mock.Expect("SELECT", "1").
				ReplyError("ERR ACL")
			mock.Expect("READONLY").
				ReplyError("ERR ACL")
			mock.Expect("CLIENT", "SETINFO", "LIB-NAME", LibName).
				ReplyError("UNKNOWN COMMAND")
			mock.Expect("CLIENT", "SETINFO", "LIB-VER", LibVer).
				ReplyError("UNKNOWN COMMAND")
			mock.Expect("AUTH", "pa").
				ReplyString("OK")
			mock.Expect("CLIENT", "SETNAME", "cn").
				ReplyString("OK")
			mock.Expect("SELECT", "1").
				ReplyString("OK")
			mock.Expect("READONLY").
				ReplyString("OK")
			mock.Expect("CLIENT", "SETINFO", "LIB-NAME", LibName).
				ReplyError("UNKNOWN COMMAND")
			mock.Expect("CLIENT", "SETINFO", "LIB-VER", LibVer).
				ReplyError("UNKNOWN COMMAND")
		}()
		p, err := newPipe(func() (net.Conn, error) { return n1, nil }, &ClientOption{
			SelectDB:     1,
			Password:     "pa",
			ClientName:   "cn",
			DisableCache: true,
			ReplicaOnly:  true,
		})
		if err != nil {
			t.Fatalf("pipe setup failed: %v", err)
		}
		if p.version >= 6 {
			t.Fatalf("unexpected p.version: %v", p.version)
		}
		go func() { mock.Expect("PING").ReplyString("OK") }()
		p.Close()
		mock.Close()
		n1.Close()
		n2.Close()
	})
	t.Run("Init with ReplicaOnly ignores READONLY error", func(t *testing.T) {
		n1, n2 := net.Pipe()
		mock := &redisMock{buf: bufio.NewReader(n2), conn: n2, t: t}
		go func() {
			mock.Expect("HELLO", "3", "AUTH", "default", "pa", "SETNAME", "cn").
				ReplyError("ERR unknown command `HELLO`")
			mock.Expect("SELECT", "1").
				ReplyError("ERR ACL")
			mock.Expect("READONLY").
				ReplyError("ERR ACL")
			mock.Expect("CLIENT", "SETINFO", "LIB-NAME", LibName).
				ReplyError("UNKNOWN COMMAND")
			mock.Expect("CLIENT", "SETINFO", "LIB-VER", LibVer).
				ReplyError("UNKNOWN COMMAND")
			mock.Expect("AUTH", "pa").
				ReplyString("OK")
			mock.Expect("CLIENT", "SETNAME", "cn").
				ReplyString("OK")
			mock.Expect("SELECT", "1").
				ReplyString("OK")
			mock.Expect("READONLY").
				ReplyError("This instance has cluster support disabled")
			mock.Expect("CLIENT", "SETINFO", "LIB-NAME", LibName).
				ReplyError("UNKNOWN COMMAND")
			mock.Expect("CLIENT", "SETINFO", "LIB-VER", LibVer).
				ReplyError("UNKNOWN COMMAND")
		}()
		p, err := newPipe(func() (net.Conn, error) { return n1, nil }, &ClientOption{
			SelectDB:     1,
			Password:     "pa",
			ClientName:   "cn",
			DisableCache: true,
			ReplicaOnly:  true,
		})
		if err != nil {
			t.Fatalf("pipe setup failed: %v", err)
		}
		if p.version >= 6 {
			t.Fatalf("unexpected p.version: %v", p.version)
		}
		go func() { mock.Expect("PING").ReplyString("OK") }()
		p.Close()
		mock.Close()
		n1.Close()
		n2.Close()
	})
	t.Run("Network Error", func(t *testing.T) {
		n1, n2 := net.Pipe()
		mock := &redisMock{buf: bufio.NewReader(n2), conn: n2, t: t}
		go func() {
			mock.Expect("HELLO", "3", "AUTH", "ua", "pa", "SETNAME", "cn").
				ReplyError("ERR unknown command `HELLO`")
			mock.Expect("SELECT", "1").
				ReplyError("ERR ACL")
			mock.Expect("CLIENT", "SETINFO", "LIB-NAME", LibName).
				ReplyError("UNKNOWN COMMAND")
			mock.Expect("CLIENT", "SETINFO", "LIB-VER", LibVer).
				ReplyError("UNKNOWN COMMAND")
			n1.Close()
			n2.Close()
		}()
		_, err := newPipe(func() (net.Conn, error) { return n1, nil }, &ClientOption{
			SelectDB:     1,
			Username:     "ua",
			Password:     "pa",
			ClientName:   "cn",
			DisableCache: true,
		})
		if err != io.ErrClosedPipe {
			t.Fatalf("pipe setup should failed with io.ErrClosedPipe, but got %v", err)
		}
	})
}

func TestWriteSingleFlush(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, mock, cancel, _ := setup(t, ClientOption{})
	defer cancel()
	go func() { mock.Expect("PING").ReplyString("OK") }()
	ExpectOK(t, p.Do(context.Background(), cmds.NewCompleted([]string{"PING"})))
}

func TestIgnoreOutOfBandDataDuringSyncMode(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, mock, cancel, _ := setup(t, ClientOption{})
	defer cancel()
	go func() {
		mock.Expect("PING").Reply(RedisMessage{typ: '>', string: "This should be ignore"}).ReplyString("OK")
	}()
	ExpectOK(t, p.Do(context.Background(), cmds.NewCompleted([]string{"PING"})))
}

func TestWriteSinglePipelineFlush(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
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

func TestWriteWithMaxFlushDelay(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, mock, cancel, _ := setup(t, ClientOption{
		AlwaysPipelining: true,
		MaxFlushDelay:    20 * time.Microsecond,
	})
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

func TestBlockWriteWithNoMaxFlushDelay(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, mock, cancel, _ := setup(t, ClientOption{
		AlwaysPipelining: true,
		MaxFlushDelay:    20 * time.Microsecond,
	})
	defer cancel()
	times := 2000
	wg := sync.WaitGroup{}
	wg.Add(times)

	for i := 0; i < times; i++ {
		go func() {
			for _, resp := range p.DoMulti(context.Background(),
				cmds.NewBlockingCompleted([]string{"PING"}),
				cmds.NewBlockingCompleted([]string{"PING"})).s {
				ExpectOK(t, resp)
			}
		}()
	}
	for i := 0; i < times; i++ {
		mock.Expect("PING").ReplyString("OK").Expect("PING").ReplyString("OK")
	}
}

func TestWriteMultiFlush(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, mock, cancel, _ := setup(t, ClientOption{})
	defer cancel()
	go func() {
		mock.Expect("PING").Expect("PING").ReplyString("OK").ReplyString("OK")
	}()
	for _, resp := range p.DoMulti(context.Background(), cmds.NewCompleted([]string{"PING"}), cmds.NewCompleted([]string{"PING"})).s {
		ExpectOK(t, resp)
	}
}

func TestWriteMultiPipelineFlush(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, mock, cancel, _ := setup(t, ClientOption{})
	defer cancel()
	times := 2000
	wg := sync.WaitGroup{}
	wg.Add(times)

	for i := 0; i < times; i++ {
		go func() {
			for _, resp := range p.DoMulti(context.Background(), cmds.NewCompleted([]string{"PING"}), cmds.NewCompleted([]string{"PING"})).s {
				ExpectOK(t, resp)
			}
		}()
	}

	for i := 0; i < times; i++ {
		mock.Expect("PING").Expect("PING").ReplyString("OK").ReplyString("OK")
	}
}

func TestDoStreamAutoPipelinePanic(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, _, _, shutdown := setup(t, ClientOption{})
	p.background()
	defer func() {
		if msg := recover(); !strings.Contains(msg.(string), "bug") {
			t.Fatal("should panic")
		}
		shutdown()
	}()
	p.DoStream(context.Background(), nil, cmds.NewCompleted([]string{"PING"}))
}

func TestDoMultiStreamAutoPipelinePanic(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, _, _, shutdown := setup(t, ClientOption{})
	p.background()
	defer func() {
		if msg := recover(); !strings.Contains(msg.(string), "bug") {
			t.Fatal("should panic")
		}
		shutdown()
	}()
	p.DoMultiStream(context.Background(), nil, cmds.NewCompleted([]string{"PING"}))
}

func TestDoStreamConcurrentPanic(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, mock, _, shutdown := setup(t, ClientOption{})
	defer func() {
		if msg := recover(); !strings.Contains(msg.(string), "bug") {
			t.Fatal("should panic")
		}
		shutdown()
	}()
	go func() {
		mock.Expect("PING")
	}()
	p.DoStream(context.Background(), nil, cmds.NewCompleted([]string{"PING"}))
	p.DoStream(context.Background(), nil, cmds.NewCompleted([]string{"PING"}))
}

func TestDoMultiStreamConcurrentPanic(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, mock, _, shutdown := setup(t, ClientOption{})
	defer func() {
		if msg := recover(); !strings.Contains(msg.(string), "bug") {
			t.Fatal("should panic")
		}
		shutdown()
	}()
	go func() {
		mock.Expect("PING")
	}()
	p.DoMultiStream(context.Background(), nil, cmds.NewCompleted([]string{"PING"}))
	p.DoMultiStream(context.Background(), nil, cmds.NewCompleted([]string{"PING"}))
}

func TestDoStreamRecycle(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, mock, cancel, _ := setup(t, ClientOption{})
	defer cancel()
	go func() {
		mock.Expect("PING").ReplyString("OK")
	}()
	conns := newPool(1, nil, nil)
	s := p.DoStream(context.Background(), conns, cmds.NewCompleted([]string{"PING"}))
	buf := bytes.NewBuffer(nil)
	if err := s.Error(); err != nil {
		t.Errorf("unexpected err %v\n", err)
	}
	for s.HasNext() {
		n, err := s.WriteTo(buf)
		if err != nil {
			t.Errorf("unexpected err %v\n", err)
		}
		if n != 2 {
			t.Errorf("unexpected n %v\n", n)
		}
	}
	if buf.String() != "OK" {
		t.Errorf("unexpected result %v\n", buf.String())
	}
	if err := s.Error(); err != io.EOF {
		t.Errorf("unexpected err %v\n", err)
	}
	if w := conns.Acquire(); w != p {
		t.Errorf("pipe is not recycled\n")
	}
}

type limitedbuffer struct {
	buf []byte
}

func (b *limitedbuffer) String() string {
	return string(b.buf)
}

func (b *limitedbuffer) Write(buf []byte) (n int, err error) {
	return n, err
}

func (b *limitedbuffer) ReadFrom(r io.Reader) (n int64, err error) {
	if n, err := r.Read(b.buf); err == nil {
		return int64(n), io.EOF
	} else {
		return int64(n), err
	}
}

func TestDoStreamRecycleDestinationFull(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, mock, cancel, _ := setup(t, ClientOption{})
	defer cancel()
	go func() {
		mock.Expect("PING").ReplyBlobString("OK")
	}()
	conns := newPool(1, nil, nil)
	s := p.DoStream(context.Background(), conns, cmds.NewCompleted([]string{"PING"}))
	buf := &limitedbuffer{buf: make([]byte, 1)}
	if err := s.Error(); err != nil {
		t.Errorf("unexpected err %v\n", err)
	}
	for s.HasNext() {
		n, err := s.WriteTo(buf)
		if err != io.EOF {
			t.Errorf("unexpected err %v\n", err)
		}
		if n != 1 {
			t.Errorf("unexpected n %v\n", n)
		}
	}
	if buf.String() != "O" {
		t.Errorf("unexpected result %v\n", buf.String())
	}
	if err := s.Error(); err != io.EOF {
		t.Errorf("unexpected err %v\n", err)
	}
	if w := conns.Acquire(); w != p {
		t.Errorf("pipe is not recycled\n")
	}
}

func TestDoMultiStreamRecycle(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, mock, cancel, _ := setup(t, ClientOption{})
	defer cancel()
	go func() {
		mock.Expect("PING").Expect("PING").ReplyString("OK").ReplyString("OK")
	}()
	conns := newPool(1, nil, nil)
	s := p.DoMultiStream(context.Background(), conns, cmds.NewCompleted([]string{"PING"}), cmds.NewCompleted([]string{"PING"}))
	buf := bytes.NewBuffer(nil)
	if err := s.Error(); err != nil {
		t.Errorf("unexpected err %v\n", err)
	}
	for s.HasNext() {
		n, err := s.WriteTo(buf)
		if err != nil {
			t.Errorf("unexpected err %v\n", err)
		}
		if n != 2 {
			t.Errorf("unexpected n %v\n", n)
		}
	}
	if buf.String() != "OKOK" {
		t.Errorf("unexpected result %v\n", buf.String())
	}
	if err := s.Error(); err != io.EOF {
		t.Errorf("unexpected err %v\n", err)
	}
	if w := conns.Acquire(); w != p {
		t.Errorf("pipe is not recycled\n")
	}
}

func TestDoMultiStreamRecycleDestinationFull(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, mock, cancel, _ := setup(t, ClientOption{})
	defer cancel()
	go func() {
		mock.Expect("PING").Expect("PING").ReplyBlobString("OK").ReplyBlobString("OK")
	}()
	conns := newPool(1, nil, nil)
	s := p.DoMultiStream(context.Background(), conns, cmds.NewCompleted([]string{"PING"}), cmds.NewCompleted([]string{"PING"}))
	buf := &limitedbuffer{buf: make([]byte, 1)}
	if err := s.Error(); err != nil {
		t.Errorf("unexpected err %v\n", err)
	}
	for s.HasNext() {
		n, err := s.WriteTo(buf)
		if err != io.EOF {
			t.Errorf("unexpected err %v\n", err)
		}
		if n != 1 {
			t.Errorf("unexpected n %v\n", n)
		}
	}
	if buf.String() != "O" {
		t.Errorf("unexpected result %v\n", buf.String())
	}
	if err := s.Error(); err != io.EOF {
		t.Errorf("unexpected err %v\n", err)
	}
	if w := conns.Acquire(); w != p {
		t.Errorf("pipe is not recycled\n")
	}
}

func TestNoReplyExceedRingSize(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, mock, cancel, _ := setup(t, ClientOption{})
	defer cancel()

	times := (2 << (DefaultRingScale - 1)) * 3
	wait := make(chan struct{})
	go func() {
		for i := 0; i < times; i++ {
			if err := p.Do(context.Background(), cmds.UnsubscribeCmd).Error(); err != nil {
				t.Errorf("unexpected err %v", err)
			}
		}
		close(wait)
	}()

	for i := 0; i < times; i++ {
		mock.Expect("UNSUBSCRIBE").Reply(RedisMessage{typ: '>', values: []RedisMessage{
			{typ: '+', string: "unsubscribe"},
			{typ: '+', string: "1"},
			{typ: ':', integer: 0},
		}})
	}
	<-wait
}

func TestPanicOnProtocolBug(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
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
	defer ShouldNotLeaked(SetupLeakDetection())
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
	defer ShouldNotLeaked(SetupLeakDetection())
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
			v, _ := p.DoCache(context.Background(), Cacheable(cmds.NewCompleted([]string{"GET", "a"})), 10*time.Second).ToMessage()
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
		if v, _ := p.DoCache(context.Background(), Cacheable(cmds.NewCompleted([]string{"GET", "a"})), time.Second).ToMessage(); v.string == "2" {
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
		if v, _ := p.DoCache(context.Background(), Cacheable(cmds.NewCompleted([]string{"GET", "a"})), time.Second).ToMessage(); v.string == "3" {
			break
		}
		t.Logf("waiting for invalidating")
	}
}

func TestClientSideCachingExecAbort(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
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

	v, err := p.DoCache(context.Background(), Cacheable(cmds.NewCompleted([]string{"GET", "a"})), 10*time.Second).ToMessage()
	if err != ErrDoCacheAborted {
		t.Errorf("unexpected err, got %v", err)
	}
	if v.IsCacheHit() {
		t.Errorf("unexpected cache hit")
	}
	if v, entry := p.cache.Flight("a", "GET", time.Second, time.Now()); v.typ != 0 || entry != nil {
		t.Errorf("unexpected cache value and entry %v %v", v, entry)
	}
}

func TestClientSideCachingWithNonRedisError(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, _, _, closeConn := setup(t, ClientOption{})
	closeConn()

	v, err := p.DoCache(context.Background(), Cacheable(cmds.NewCompleted([]string{"GET", "a"})), 10*time.Second).ToMessage()
	if err != io.EOF && !strings.HasPrefix(err.Error(), "io:") {
		t.Errorf("unexpected err, got %v", err)
	}
	if v.IsCacheHit() {
		t.Errorf("unexpected cache hit")
	}
	if v, entry := p.cache.Flight("a", "GET", time.Second, time.Now()); v.typ != 0 || entry != nil {
		t.Errorf("unexpected cache value and entry %v %v", v, entry)
	}
}

func TestClientSideCachingMGet(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, mock, cancel, _ := setup(t, ClientOption{})
	defer cancel()

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
		mock.Expect("CLIENT", "CACHING", "YES").
			Expect("MULTI").
			Expect("PTTL", "a1").
			Expect("PTTL", "a2").
			Expect("PTTL", "a3").
			Expect("MGET", "a1", "a2", "a3").
			Expect("EXEC").
			ReplyString("OK").
			ReplyString("OK").
			ReplyString("OK").
			ReplyString("OK").
			ReplyString("OK").
			ReplyString("OK").
			Reply(RedisMessage{typ: '*', values: []RedisMessage{
				{typ: ':', integer: 1000},
				{typ: ':', integer: 2000},
				{typ: ':', integer: 3000},
				{typ: '*', values: []RedisMessage{
					{typ: ':', integer: 1},
					{typ: ':', integer: 2},
					{typ: ':', integer: 3},
				}},
			}})
	}()
	// single flight
	miss := uint64(0)
	hits := uint64(0)
	for i := 0; i < 2; i++ {
		v, _ := p.DoCache(context.Background(), Cacheable(cmds.NewMGetCompleted([]string{"MGET", "a1", "a2", "a3"})), 10*time.Second).ToMessage()
		arr, _ := v.ToArray()
		if len(arr) != 3 {
			t.Errorf("unexpected cached mget length, expected 3, got %v", len(arr))
		}
		for i, v := range arr {
			if v.integer != int64(i+1) {
				t.Errorf("unexpected cached mget response, expected %v, got %v", i+1, v.integer)
			}
		}
		if ttl := p.cache.(*lru).GetTTL("a1", "GET"); !roughly(ttl, time.Second) {
			t.Errorf("unexpected ttl %v", ttl)
		}
		if ttl := p.cache.(*lru).GetTTL("a2", "GET"); !roughly(ttl, time.Second*2) {
			t.Errorf("unexpected ttl %v", ttl)
		}
		if ttl := p.cache.(*lru).GetTTL("a3", "GET"); !roughly(ttl, time.Second*3) {
			t.Errorf("unexpected ttl %v", ttl)
		}
		if v.IsCacheHit() {
			atomic.AddUint64(&hits, 1)
		} else {
			atomic.AddUint64(&miss, 1)
		}
	}

	if v := atomic.LoadUint64(&miss); v != 1 {
		t.Fatalf("unexpected cache miss count %v", v)
	}

	if v := atomic.LoadUint64(&hits); v != uint64(1) {
		t.Fatalf("unexpected cache hits count %v", v)
	}

	// partial cache invalidation
	invalidateCSC(RedisMessage{typ: '*', values: []RedisMessage{{typ: '+', string: "a1"}, {typ: '+', string: "a3"}}})
	go func() {
		mock.Expect("CLIENT", "CACHING", "YES").
			Expect("MULTI").
			Expect("PTTL", "a1").
			Expect("PTTL", "a3").
			Expect("MGET", "a1", "a3").
			Expect("EXEC").
			ReplyString("OK").
			ReplyString("OK").
			ReplyString("OK").
			ReplyString("OK").
			ReplyString("OK").
			Reply(RedisMessage{typ: '*', values: []RedisMessage{
				{typ: ':', integer: 10000},
				{typ: ':', integer: 30000},
				{typ: '*', values: []RedisMessage{
					{typ: ':', integer: 10},
					{typ: ':', integer: 30},
				}},
			}})
	}()

	for {
		if p.cache.(*lru).GetTTL("a1", "GET") == -2 && p.cache.(*lru).GetTTL("a3", "GET") == -2 {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}

	v, _ := p.DoCache(context.Background(), Cacheable(cmds.NewMGetCompleted([]string{"MGET", "a1", "a2", "a3"})), 10*time.Second).ToMessage()
	arr, _ := v.ToArray()
	if len(arr) != 3 {
		t.Errorf("unexpected cached mget length, expected 3, got %v", len(arr))
	}
	if arr[1].integer != 2 {
		t.Errorf("unexpected cached mget response, expected %v, got %v", 2, arr[1].integer)
	}
	if arr[0].integer != 10 {
		t.Errorf("unexpected cached mget response, expected %v, got %v", 10, arr[0].integer)
	}
	if arr[2].integer != 30 {
		t.Errorf("unexpected cached mget response, expected %v, got %v", 30, arr[2].integer)
	}
	if ttl := p.cache.(*lru).GetTTL("a1", "GET"); !roughly(ttl, time.Second*10) {
		t.Errorf("unexpected ttl %v", ttl)
	}
	if ttl := p.cache.(*lru).GetTTL("a2", "GET"); !roughly(ttl, time.Second*2) {
		t.Errorf("unexpected ttl %v", ttl)
	}
	if ttl := p.cache.(*lru).GetTTL("a3", "GET"); !roughly(ttl, time.Second*30) {
		t.Errorf("unexpected ttl %v", ttl)
	}
}

func TestClientSideCachingJSONMGet(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, mock, cancel, _ := setup(t, ClientOption{})
	defer cancel()

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
		mock.Expect("CLIENT", "CACHING", "YES").
			Expect("MULTI").
			Expect("PTTL", "a1").
			Expect("PTTL", "a2").
			Expect("PTTL", "a3").
			Expect("JSON.MGET", "a1", "a2", "a3", "$").
			Expect("EXEC").
			ReplyString("OK").
			ReplyString("OK").
			ReplyString("OK").
			ReplyString("OK").
			ReplyString("OK").
			ReplyString("OK").
			Reply(RedisMessage{typ: '*', values: []RedisMessage{
				{typ: ':', integer: 1000},
				{typ: ':', integer: 2000},
				{typ: ':', integer: 3000},
				{typ: '*', values: []RedisMessage{
					{typ: ':', integer: 1},
					{typ: ':', integer: 2},
					{typ: ':', integer: 3},
				}},
			}})
	}()
	// single flight
	miss := uint64(0)
	hits := uint64(0)
	for i := 0; i < 2; i++ {
		v, _ := p.DoCache(context.Background(), Cacheable(cmds.NewMGetCompleted([]string{"JSON.MGET", "a1", "a2", "a3", "$"})), 10*time.Second).ToMessage()
		arr, _ := v.ToArray()
		if len(arr) != 3 {
			t.Errorf("unexpected cached mget length, expected 3, got %v", len(arr))
		}
		for i, v := range arr {
			if v.integer != int64(i+1) {
				t.Errorf("unexpected cached mget response, expected %v, got %v", i+1, v.integer)
			}
		}
		if ttl := p.cache.(*lru).GetTTL("a1", "JSON.GET$"); !roughly(ttl, time.Second) {
			t.Errorf("unexpected ttl %v", ttl)
		}
		if ttl := p.cache.(*lru).GetTTL("a2", "JSON.GET$"); !roughly(ttl, time.Second*2) {
			t.Errorf("unexpected ttl %v", ttl)
		}
		if ttl := p.cache.(*lru).GetTTL("a3", "JSON.GET$"); !roughly(ttl, time.Second*3) {
			t.Errorf("unexpected ttl %v", ttl)
		}
		if v.IsCacheHit() {
			atomic.AddUint64(&hits, 1)
		} else {
			atomic.AddUint64(&miss, 1)
		}
	}

	if v := atomic.LoadUint64(&miss); v != 1 {
		t.Fatalf("unexpected cache miss count %v", v)
	}

	if v := atomic.LoadUint64(&hits); v != uint64(1) {
		t.Fatalf("unexpected cache hits count %v", v)
	}

	// partial cache invalidation
	invalidateCSC(RedisMessage{typ: '*', values: []RedisMessage{{typ: '+', string: "a1"}, {typ: '+', string: "a3"}}})
	go func() {
		mock.Expect("CLIENT", "CACHING", "YES").
			Expect("MULTI").
			Expect("PTTL", "a1").
			Expect("PTTL", "a3").
			Expect("JSON.MGET", "a1", "a3", "$").
			Expect("EXEC").
			ReplyString("OK").
			ReplyString("OK").
			ReplyString("OK").
			ReplyString("OK").
			ReplyString("OK").
			Reply(RedisMessage{typ: '*', values: []RedisMessage{
				{typ: ':', integer: 10000},
				{typ: ':', integer: 30000},
				{typ: '*', values: []RedisMessage{
					{typ: ':', integer: 10},
					{typ: ':', integer: 30},
				}},
			}})
	}()

	for {
		if p.cache.(*lru).GetTTL("a1", "JSON.GET$") == -2 && p.cache.(*lru).GetTTL("a3", "JSON.GET$") == -2 {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}

	v, _ := p.DoCache(context.Background(), Cacheable(cmds.NewMGetCompleted([]string{"JSON.MGET", "a1", "a2", "a3", "$"})), 10*time.Second).ToMessage()
	arr, _ := v.ToArray()
	if len(arr) != 3 {
		t.Errorf("unexpected cached mget length, expected 3, got %v", len(arr))
	}
	if arr[1].integer != 2 {
		t.Errorf("unexpected cached mget response, expected %v, got %v", 2, arr[1].integer)
	}
	if arr[0].integer != 10 {
		t.Errorf("unexpected cached mget response, expected %v, got %v", 10, arr[0].integer)
	}
	if arr[2].integer != 30 {
		t.Errorf("unexpected cached mget response, expected %v, got %v", 30, arr[2].integer)
	}
	if ttl := p.cache.(*lru).GetTTL("a1", "JSON.GET$"); !roughly(ttl, time.Second*10) {
		t.Errorf("unexpected ttl %v", ttl)
	}
	if ttl := p.cache.(*lru).GetTTL("a2", "JSON.GET$"); !roughly(ttl, time.Second*2) {
		t.Errorf("unexpected ttl %v", ttl)
	}
	if ttl := p.cache.(*lru).GetTTL("a3", "JSON.GET$"); !roughly(ttl, time.Second*30) {
		t.Errorf("unexpected ttl %v", ttl)
	}
}

func TestClientSideCachingExecAbortMGet(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, mock, cancel, _ := setup(t, ClientOption{})
	defer cancel()

	go func() {
		mock.Expect("CLIENT", "CACHING", "YES").
			Expect("MULTI").
			Expect("PTTL", "a1").
			Expect("PTTL", "a2").
			Expect("MGET", "a1", "a2").
			Expect("EXEC").
			ReplyString("OK").
			ReplyString("OK").
			ReplyString("OK").
			ReplyString("OK").
			ReplyString("OK").
			Reply(RedisMessage{typ: '_'})
	}()

	v, err := p.DoCache(context.Background(), Cacheable(cmds.NewMGetCompleted([]string{"MGET", "a1", "a2"})), 10*time.Second).ToMessage()
	if err != ErrDoCacheAborted {
		t.Errorf("unexpected err, got %v", err)
	}
	if v.IsCacheHit() {
		t.Errorf("unexpected cache hit")
	}
	if v, entry := p.cache.Flight("a1", "GET", time.Second, time.Now()); v.typ != 0 || entry != nil {
		t.Errorf("unexpected cache value and entry %v %v", v, entry)
	}
	if v, entry := p.cache.Flight("a2", "GET", time.Second, time.Now()); v.typ != 0 || entry != nil {
		t.Errorf("unexpected cache value and entry %v %v", v, entry)
	}
}

func TestClientSideCachingWithNonRedisErrorMGet(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, _, _, closeConn := setup(t, ClientOption{})
	closeConn()

	v, err := p.DoCache(context.Background(), Cacheable(cmds.NewMGetCompleted([]string{"MGET", "a1", "a2"})), 10*time.Second).ToMessage()
	if err != io.EOF && !strings.HasPrefix(err.Error(), "io:") {
		t.Errorf("unexpected err, got %v", err)
	}
	if v.IsCacheHit() {
		t.Errorf("unexpected cache hit")
	}
	if v, entry := p.cache.Flight("a1", "GET", time.Second, time.Now()); v.typ != 0 || entry != nil {
		t.Errorf("unexpected cache value and entry %v %v", v, entry)
	}
	if v, entry := p.cache.Flight("a2", "GET", time.Second, time.Now()); v.typ != 0 || entry != nil {
		t.Errorf("unexpected cache value and entry %v %v", v, entry)
	}
}

func TestClientSideCachingWithSideChannelMGet(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, _, _, closeConn := setup(t, ClientOption{})
	closeConn()

	p.cache.Flight("a1", "GET", 10*time.Second, time.Now())
	go func() {
		time.Sleep(100 * time.Millisecond)
		m := RedisMessage{typ: '+', string: "OK"}
		m.setExpireAt(time.Now().Add(10 * time.Millisecond).UnixMilli())
		p.cache.Update("a1", "GET", m)
	}()

	v, _ := p.DoCache(context.Background(), Cacheable(cmds.NewMGetCompleted([]string{"MGET", "a1"})), 10*time.Second).AsStrSlice()
	if v[0] != "OK" {
		t.Errorf("unexpected value, got %v", v)
	}
}

func TestClientSideCachingWithSideChannelErrorMGet(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, _, _, closeConn := setup(t, ClientOption{})
	closeConn()

	p.cache.Flight("a1", "GET", 10*time.Second, time.Now())
	go func() {
		time.Sleep(100 * time.Millisecond)
		p.cache.Cancel("a1", "GET", io.EOF)
	}()

	_, err := p.DoCache(context.Background(), Cacheable(cmds.NewMGetCompleted([]string{"MGET", "a1"})), 10*time.Second).ToMessage()
	if err != io.EOF {
		t.Errorf("unexpected err, got %v", err)
	}
}

func TestClientSideCachingDoMultiCacheMGet(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, _, cancel, _ := setup(t, ClientOption{})
	defer cancel()

	defer func() {
		if !strings.Contains(recover().(string), panicmgetcsc) {
			t.Fatal("should panic")
		}
	}()
	p.DoMultiCache(context.Background(), []CacheableTTL{
		CT(Cacheable(cmds.NewMGetCompleted([]string{"MGET", "a1"})), time.Second*10),
	}...)
}

func TestClientSideCachingDoMultiCache(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	testfn := func(t *testing.T, option ClientOption) {
		p, mock, cancel, _ := setup(t, option)
		defer cancel()

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
			mock.Expect("CLIENT", "CACHING", "YES").
				Expect("MULTI").
				Expect("PTTL", "a1").
				Expect("GET", "a1").
				Expect("EXEC").
				Expect("CLIENT", "CACHING", "YES").
				Expect("MULTI").
				Expect("PTTL", "a2").
				Expect("GET", "a2").
				Expect("EXEC").
				Expect("CLIENT", "CACHING", "YES").
				Expect("MULTI").
				Expect("PTTL", "a3").
				Expect("GET", "a3").
				Expect("EXEC").
				ReplyString("OK").
				ReplyString("OK").
				ReplyString("OK").
				ReplyString("OK").
				Reply(RedisMessage{typ: '*', values: []RedisMessage{
					{typ: ':', integer: 1000},
					{typ: ':', integer: 1},
				}}).
				ReplyString("OK").
				ReplyString("OK").
				ReplyString("OK").
				ReplyString("OK").
				Reply(RedisMessage{typ: '*', values: []RedisMessage{
					{typ: ':', integer: 2000},
					{typ: ':', integer: 2},
				}}).
				ReplyString("OK").
				ReplyString("OK").
				ReplyString("OK").
				ReplyString("OK").
				Reply(RedisMessage{typ: '*', values: []RedisMessage{
					{typ: ':', integer: 3000},
					{typ: ':', integer: 3},
				}})
		}()
		// single flight
		miss := uint64(0)
		hits := uint64(0)
		for i := 0; i < 2; i++ {
			arr := p.DoMultiCache(context.Background(), []CacheableTTL{
				CT(Cacheable(cmds.NewCompleted([]string{"GET", "a1"})), time.Second*10),
				CT(Cacheable(cmds.NewCompleted([]string{"GET", "a2"})), time.Second*10),
				CT(Cacheable(cmds.NewCompleted([]string{"GET", "a3"})), time.Second*10),
			}...).s
			if len(arr) != 3 {
				t.Errorf("unexpected cached mget length, expected 3, got %v", len(arr))
			}
			for i, v := range arr {
				if v.val.integer != int64(i+1) {
					t.Errorf("unexpected cached mget response, expected %v, got %v", i+1, v.val.integer)
				}
				if v.val.IsCacheHit() {
					atomic.AddUint64(&hits, 1)
				} else {
					atomic.AddUint64(&miss, 1)
				}
			}
			if ttl := time.Duration(arr[0].CachePTTL()) * time.Millisecond; !roughly(ttl, time.Second) {
				t.Errorf("unexpected ttl %v", ttl)
			}
			if ttl := time.Duration(arr[1].CachePTTL()) * time.Millisecond; !roughly(ttl, time.Second*2) {
				t.Errorf("unexpected ttl %v", ttl)
			}
			if ttl := time.Duration(arr[2].CachePTTL()) * time.Millisecond; !roughly(ttl, time.Second*3) {
				t.Errorf("unexpected ttl %v", ttl)
			}
		}

		if v := atomic.LoadUint64(&miss); v != 3 {
			t.Fatalf("unexpected cache miss count %v", v)
		}

		if v := atomic.LoadUint64(&hits); v != 3 {
			t.Fatalf("unexpected cache hits count %v", v)
		}

		// partial cache invalidation
		invalidateCSC(RedisMessage{typ: '*', values: []RedisMessage{{typ: '+', string: "a1"}, {typ: '+', string: "a3"}}})
		go func() {
			mock.Expect("CLIENT", "CACHING", "YES").
				Expect("MULTI").
				Expect("PTTL", "a1").
				Expect("GET", "a1").
				Expect("EXEC").
				Expect("CLIENT", "CACHING", "YES").
				Expect("MULTI").
				Expect("PTTL", "a3").
				Expect("GET", "a3").
				Expect("EXEC").
				ReplyString("OK").
				ReplyString("OK").
				ReplyString("OK").
				ReplyString("OK").
				Reply(RedisMessage{typ: '*', values: []RedisMessage{
					{typ: ':', integer: 10000},
					{typ: ':', integer: 10},
				}}).
				ReplyString("OK").
				ReplyString("OK").
				ReplyString("OK").
				ReplyString("OK").
				Reply(RedisMessage{typ: '*', values: []RedisMessage{
					{typ: ':', integer: 30000},
					{typ: ':', integer: 30},
				}})
		}()

		if cache, ok := p.cache.(*lru); ok {
			for {
				if cache.GetTTL("a1", "GET") == -2 && p.cache.(*lru).GetTTL("a3", "GET") == -2 {
					break
				}
				time.Sleep(10 * time.Millisecond)
			}
		} else {
			time.Sleep(time.Second)
		}

		arr := p.DoMultiCache(context.Background(), []CacheableTTL{
			CT(Cacheable(cmds.NewCompleted([]string{"GET", "a1"})), time.Second*10),
			CT(Cacheable(cmds.NewCompleted([]string{"GET", "a2"})), time.Second*10),
			CT(Cacheable(cmds.NewCompleted([]string{"GET", "a3"})), time.Second*10),
		}...).s
		if len(arr) != 3 {
			t.Errorf("unexpected cached mget length, expected 3, got %v", len(arr))
		}
		if arr[1].val.integer != 2 {
			t.Errorf("unexpected cached mget response, expected %v, got %v", 2, arr[1].val.integer)
		}
		if arr[0].val.integer != 10 {
			t.Errorf("unexpected cached mget response, expected %v, got %v", 10, arr[0].val.integer)
		}
		if arr[2].val.integer != 30 {
			t.Errorf("unexpected cached mget response, expected %v, got %v", 30, arr[2].val.integer)
		}
		if ttl := time.Duration(arr[0].CachePTTL()) * time.Millisecond; !roughly(ttl, time.Second*10) {
			t.Errorf("unexpected ttl %v", ttl)
		}
		if ttl := time.Duration(arr[1].CachePTTL()) * time.Millisecond; !roughly(ttl, time.Second*2) {
			t.Errorf("unexpected ttl %v", ttl)
		}
		if ttl := time.Duration(arr[2].CachePTTL()) * time.Millisecond; !roughly(ttl, time.Second*30) {
			t.Errorf("unexpected ttl %v", ttl)
		}
	}
	t.Run("LRU", func(t *testing.T) {
		testfn(t, ClientOption{})
	})
	t.Run("Simple", func(t *testing.T) {
		testfn(t, ClientOption{
			NewCacheStoreFn: func(option CacheStoreOption) CacheStore {
				return NewSimpleCacheAdapter(&simple{store: map[string]RedisMessage{}})
			},
		})
	})
}

func TestClientSideCachingExecAbortDoMultiCache(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	testfn := func(t *testing.T, option ClientOption) {
		p, mock, cancel, _ := setup(t, option)
		defer cancel()

		go func() {
			mock.Expect("CLIENT", "CACHING", "YES").
				Expect("MULTI").
				Expect("PTTL", "a1").
				Expect("GET", "a1").
				Expect("EXEC").
				Expect("CLIENT", "CACHING", "YES").
				Expect("MULTI").
				Expect("PTTL", "a2").
				Expect("GET", "a2").
				Expect("EXEC").
				ReplyString("OK").
				ReplyString("OK").
				ReplyString("OK").
				ReplyString("OK").
				Reply(RedisMessage{typ: '*', values: []RedisMessage{
					{typ: ':', integer: 1000},
					{typ: ':', integer: 1},
				}}).
				ReplyString("OK").
				ReplyString("OK").
				ReplyString("OK").
				ReplyString("OK").
				Reply(RedisMessage{typ: '_'})
		}()

		arr := p.DoMultiCache(context.Background(), []CacheableTTL{
			CT(Cacheable(cmds.NewCompleted([]string{"GET", "a1"})), time.Second*10),
			CT(Cacheable(cmds.NewCompleted([]string{"GET", "a2"})), time.Second*10),
		}...).s
		for i, resp := range arr {
			v, err := resp.ToMessage()
			if i == 0 {
				if v.integer != 1 {
					t.Errorf("unexpected cached response, expected %v, got %v", 1, v.integer)
				}
			} else {
				if err != ErrDoCacheAborted {
					t.Errorf("unexpected err, got %v", err)
				}
				if v.IsCacheHit() {
					t.Errorf("unexpected cache hit")
				}
			}
		}
		if v, entry := p.cache.Flight("a1", "GET", time.Second, time.Now()); v.integer != 1 {
			t.Errorf("unexpected cache value and entry %v %v", v.integer, entry)
		}
		if ttl := time.Duration(arr[0].CachePTTL()) * time.Millisecond; !roughly(ttl, time.Second) {
			t.Errorf("unexpected ttl %v", ttl)
		}
		if v, entry := p.cache.Flight("a2", "GET", time.Second, time.Now()); v.typ != 0 || entry != nil {
			t.Errorf("unexpected cache value and entry %v %v", v, entry)
		}
	}
	t.Run("LRU", func(t *testing.T) {
		testfn(t, ClientOption{})
	})
	t.Run("Simple", func(t *testing.T) {
		testfn(t, ClientOption{
			NewCacheStoreFn: func(option CacheStoreOption) CacheStore {
				return NewSimpleCacheAdapter(&simple{store: map[string]RedisMessage{}})
			},
		})
	})
}

func TestClientSideCachingWithNonRedisErrorDoMultiCache(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	testfn := func(t *testing.T, option ClientOption) {
		p, _, _, closeConn := setup(t, option)
		closeConn()

		arr := p.DoMultiCache(context.Background(), []CacheableTTL{
			CT(Cacheable(cmds.NewCompleted([]string{"GET", "a1"})), time.Second*10),
			CT(Cacheable(cmds.NewCompleted([]string{"GET", "a2"})), time.Second*10),
		}...).s
		for _, resp := range arr {
			v, err := resp.ToMessage()
			if err != io.EOF && !strings.HasPrefix(err.Error(), "io:") {
				t.Errorf("unexpected err, got %v", err)
			}
			if v.IsCacheHit() {
				t.Errorf("unexpected cache hit")
			}
		}
		if v, entry := p.cache.Flight("a1", "GET", time.Second, time.Now()); v.typ != 0 || entry != nil {
			t.Errorf("unexpected cache value and entry %v %v", v, entry)
		}
		if v, entry := p.cache.Flight("a2", "GET", time.Second, time.Now()); v.typ != 0 || entry != nil {
			t.Errorf("unexpected cache value and entry %v %v", v, entry)
		}
	}
	t.Run("LRU", func(t *testing.T) {
		testfn(t, ClientOption{})
	})
	t.Run("Simple", func(t *testing.T) {
		testfn(t, ClientOption{
			NewCacheStoreFn: func(option CacheStoreOption) CacheStore {
				return NewSimpleCacheAdapter(&simple{store: map[string]RedisMessage{}})
			},
		})
	})
}

func TestClientSideCachingWithSideChannelDoMultiCache(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	testfn := func(t *testing.T, option ClientOption) {
		p, _, _, closeConn := setup(t, option)
		closeConn()

		p.cache.Flight("a1", "GET", 10*time.Second, time.Now())
		go func() {
			time.Sleep(100 * time.Millisecond)
			m := RedisMessage{typ: '+', string: "OK"}
			m.setExpireAt(time.Now().Add(10 * time.Millisecond).UnixMilli())
			p.cache.Update("a1", "GET", m)
		}()

		arr := p.DoMultiCache(context.Background(), []CacheableTTL{
			CT(Cacheable(cmds.NewCompleted([]string{"GET", "a1"})), time.Second*10),
		}...).s
		if arr[0].val.string != "OK" {
			t.Errorf("unexpected value, got %v", arr[0].val.string)
		}
	}
	t.Run("LRU", func(t *testing.T) {
		testfn(t, ClientOption{})
	})
	t.Run("Simple", func(t *testing.T) {
		testfn(t, ClientOption{
			NewCacheStoreFn: func(option CacheStoreOption) CacheStore {
				return NewSimpleCacheAdapter(&simple{store: map[string]RedisMessage{}})
			},
		})
	})
}

func TestClientSideCachingWithSideChannelErrorDoMultiCache(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	testfn := func(t *testing.T, option ClientOption) {
		p, _, _, closeConn := setup(t, option)
		closeConn()
		p.cache.Flight("a1", "GET", 10*time.Second, time.Now())
		go func() {
			time.Sleep(100 * time.Millisecond)
			p.cache.Cancel("a1", "GET", io.EOF)
		}()

		arr := p.DoMultiCache(context.Background(), []CacheableTTL{
			CT(Cacheable(cmds.NewCompleted([]string{"GET", "a1"})), time.Second*10),
		}...).s
		if arr[0].err != io.EOF {
			t.Errorf("unexpected err, got %v", arr[0].err)
		}
	}
	t.Run("LRU", func(t *testing.T) {
		testfn(t, ClientOption{})
	})
	t.Run("Simple", func(t *testing.T) {
		testfn(t, ClientOption{
			NewCacheStoreFn: func(option CacheStoreOption) CacheStore {
				return NewSimpleCacheAdapter(&simple{store: map[string]RedisMessage{}})
			},
		})
	})
}

func TestClientSideCachingMissCacheTTL(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	testfn := func(t *testing.T, option ClientOption) {
		t.Run("DoCache GET", func(t *testing.T) {
			p, mock, cancel, _ := setup(t, option)
			defer cancel()
			expectCSC := func(pttl int64, key string) {
				mock.Expect("CLIENT", "CACHING", "YES").
					Expect("MULTI").
					Expect("PTTL", key).
					Expect("GET", key).
					Expect("EXEC").
					ReplyString("OK").
					ReplyString("OK").
					ReplyString("OK").
					ReplyString("OK").
					Reply(RedisMessage{typ: '*', values: []RedisMessage{
						{typ: ':', integer: pttl},
						{typ: '+', string: key},
					}})
			}
			go func() {
				expectCSC(-1, "a")
				expectCSC(1000, "b")
				expectCSC(20000, "c")
			}()
			v, _ := p.DoCache(context.Background(), Cacheable(cmds.NewCompleted([]string{"GET", "a"})), 10*time.Second).ToMessage()
			if ttl := v.CacheTTL(); ttl != 10 {
				t.Errorf("unexpected cached ttl, expected %v, got %v", 10, ttl)
			}
			v, _ = p.DoCache(context.Background(), Cacheable(cmds.NewCompleted([]string{"GET", "b"})), 10*time.Second).ToMessage()
			if ttl := v.CacheTTL(); ttl != 1 {
				t.Errorf("unexpected cached ttl, expected %v, got %v", 1, ttl)
			}
			v, _ = p.DoCache(context.Background(), Cacheable(cmds.NewCompleted([]string{"GET", "c"})), 10*time.Second).ToMessage()
			if ttl := v.CacheTTL(); ttl != 10 {
				t.Errorf("unexpected cached ttl, expected %v, got %v", 10, ttl)
			}
		})
		t.Run("DoCache MGET", func(t *testing.T) {
			p, mock, cancel, _ := setup(t, option)
			defer cancel()
			go func() {
				mock.Expect("CLIENT", "CACHING", "YES").
					Expect("MULTI").
					Expect("PTTL", "a").
					Expect("PTTL", "b").
					Expect("PTTL", "c").
					Expect("MGET", "a", "b", "c").
					Expect("EXEC").
					ReplyString("OK").
					ReplyString("OK").
					ReplyString("OK").
					ReplyString("OK").
					ReplyString("OK").
					ReplyString("OK").
					Reply(RedisMessage{typ: '*', values: []RedisMessage{
						{typ: ':', integer: -1},
						{typ: ':', integer: 1000},
						{typ: ':', integer: 20000},
						{typ: '*', values: []RedisMessage{
							{typ: '+', string: "a"},
							{typ: '+', string: "b"},
							{typ: '+', string: "c"},
						}},
					}})
			}()
			v, _ := p.DoCache(context.Background(), Cacheable(cmds.NewMGetCompleted([]string{"MGET", "a", "b", "c"})), 10*time.Second).ToArray()
			if ttl := v[0].CacheTTL(); ttl != 10 {
				t.Errorf("unexpected cached ttl, expected %v, got %v", 10, ttl)
			}
			if ttl := v[1].CacheTTL(); ttl != 1 {
				t.Errorf("unexpected cached ttl, expected %v, got %v", 1, ttl)
			}
			if ttl := v[2].CacheTTL(); ttl != 10 {
				t.Errorf("unexpected cached ttl, expected %v, got %v", 10, ttl)
			}
		})
		t.Run("DoMultiCache", func(t *testing.T) {
			p, mock, cancel, _ := setup(t, option)
			defer cancel()
			go func() {
				mock.Expect("CLIENT", "CACHING", "YES").
					Expect("MULTI").
					Expect("PTTL", "a1").
					Expect("GET", "a1").
					Expect("EXEC").
					Expect("CLIENT", "CACHING", "YES").
					Expect("MULTI").
					Expect("PTTL", "a2").
					Expect("GET", "a2").
					Expect("EXEC").
					Expect("CLIENT", "CACHING", "YES").
					Expect("MULTI").
					Expect("PTTL", "a3").
					Expect("GET", "a3").
					Expect("EXEC").
					ReplyString("OK").
					ReplyString("OK").
					ReplyString("OK").
					ReplyString("OK").
					Reply(RedisMessage{typ: '*', values: []RedisMessage{
						{typ: ':', integer: -1},
						{typ: ':', integer: 1},
					}}).
					ReplyString("OK").
					ReplyString("OK").
					ReplyString("OK").
					ReplyString("OK").
					Reply(RedisMessage{typ: '*', values: []RedisMessage{
						{typ: ':', integer: 1000},
						{typ: ':', integer: 2},
					}}).
					ReplyString("OK").
					ReplyString("OK").
					ReplyString("OK").
					ReplyString("OK").
					Reply(RedisMessage{typ: '*', values: []RedisMessage{
						{typ: ':', integer: 20000},
						{typ: ':', integer: 3},
					}})
			}()
			arr := p.DoMultiCache(context.Background(), []CacheableTTL{
				CT(Cacheable(cmds.NewCompleted([]string{"GET", "a1"})), time.Second*10),
				CT(Cacheable(cmds.NewCompleted([]string{"GET", "a2"})), time.Second*10),
				CT(Cacheable(cmds.NewCompleted([]string{"GET", "a3"})), time.Second*10),
			}...).s
			if ttl := arr[0].CacheTTL(); ttl != 10 {
				t.Errorf("unexpected cached ttl, expected %v, got %v", 10, ttl)
			}
			if ttl := arr[1].CacheTTL(); ttl != 1 {
				t.Errorf("unexpected cached ttl, expected %v, got %v", 1, ttl)
			}
			if ttl := arr[2].CacheTTL(); ttl != 10 {
				t.Errorf("unexpected cached ttl, expected %v, got %v", 10, ttl)
			}
		})
	}
	t.Run("LRU", func(t *testing.T) {
		testfn(t, ClientOption{})
	})
	t.Run("Simple", func(t *testing.T) {
		testfn(t, ClientOption{
			NewCacheStoreFn: func(option CacheStoreOption) CacheStore {
				return NewSimpleCacheAdapter(&simple{store: map[string]RedisMessage{}})
			},
		})
	})
}

// https://github.com/redis/redis/issues/8935
func TestClientSideCachingRedis6InvalidationBug1(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, mock, cancel, _ := setup(t, ClientOption{})
	defer cancel()

	expectCSC := func() {
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
				{
					typ: '>',
					values: []RedisMessage{
						{typ: '+', string: "invalidate"},
						{typ: '*', values: []RedisMessage{{typ: '+', string: "a"}}},
					},
				},
				{typ: ':', integer: -2},
			}}).Reply(RedisMessage{typ: '_'})
	}

	go func() {
		expectCSC()
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
			v, _ := p.DoCache(context.Background(), Cacheable(cmds.NewCompleted([]string{"GET", "a"})), 10*time.Second).ToMessage()
			if v.typ != '_' {
				t.Errorf("unexpected cached result, expected null, got %v", v.string)
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
}

// https://github.com/redis/redis/issues/8935
func TestClientSideCachingRedis6InvalidationBug2(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, mock, cancel, _ := setup(t, ClientOption{})
	defer cancel()

	expectCSC := func() {
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
				{typ: ':', integer: -2},
				{
					typ: '>',
					values: []RedisMessage{
						{typ: '+', string: "invalidate"},
						{typ: '*', values: []RedisMessage{{typ: '+', string: "a"}}},
					},
				},
			}}).Reply(RedisMessage{typ: '_'})
	}

	go func() {
		expectCSC()
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
			v, _ := p.DoCache(context.Background(), Cacheable(cmds.NewCompleted([]string{"GET", "a"})), 10*time.Second).ToMessage()
			if v.typ != '_' {
				t.Errorf("unexpected cached result, expected null, got %v", v.string)
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
}

// https://github.com/redis/redis/issues/8935
func TestClientSideCachingRedis6InvalidationBugErr(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, mock, _, closeConn := setup(t, ClientOption{})

	expectCSC := func() {
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
				{typ: ':', integer: -2},
				{
					typ: '>',
					values: []RedisMessage{
						{typ: '+', string: "invalidate"},
						{typ: '*', values: []RedisMessage{{typ: '+', string: "a"}}},
					},
				},
			}})
	}

	go func() {
		expectCSC()
		closeConn()
	}()

	err := p.DoCache(context.Background(), Cacheable(cmds.NewCompleted([]string{"GET", "a"})), 10*time.Second).Error()
	if err != io.EOF && !strings.HasPrefix(err.Error(), "io:") {
		t.Errorf("unexpected err %v", err)
	}
}

func TestDisableClientSideCaching(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, mock, cancel, _ := setup(t, ClientOption{DisableCache: true})
	defer cancel()
	p.background()

	go func() {
		mock.Expect().Reply(RedisMessage{
			typ: '>',
			values: []RedisMessage{
				{typ: '+', string: "invalidate"},
				{typ: '*', values: []RedisMessage{{typ: '+', string: "a"}}},
			},
		})
		mock.Expect("GET", "a").ReplyString("1").
			Expect("GET", "b").
			Expect("GET", "c").
			ReplyString("2").
			ReplyString("3")
	}()

	v, _ := p.DoCache(context.Background(), Cacheable(cmds.NewCompleted([]string{"GET", "a"})), 10*time.Second).ToMessage()
	if v.string != "1" {
		t.Errorf("unexpected cached result, expected %v, got %v", "1", v.string)
	}

	vs := p.DoMultiCache(context.Background(),
		CT(Cacheable(cmds.NewCompleted([]string{"GET", "b"})), 10*time.Second),
		CT(Cacheable(cmds.NewCompleted([]string{"GET", "c"})), 10*time.Second)).s
	if vs[0].val.string != "2" {
		t.Errorf("unexpected cached result, expected %v, got %v", "1", v.string)
	}
	if vs[1].val.string != "3" {
		t.Errorf("unexpected cached result, expected %v, got %v", "1", v.string)
	}
}

func TestOnInvalidations(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	ch := make(chan []RedisMessage)
	_, mock, cancel, _ := setup(t, ClientOption{
		OnInvalidations: func(messages []RedisMessage) {
			ch <- messages
		},
	})

	go func() {
		mock.Expect().Reply(RedisMessage{
			typ: '>',
			values: []RedisMessage{
				{typ: '+', string: "invalidate"},
				{typ: '*', values: []RedisMessage{{typ: '+', string: "a"}}},
			},
		})
	}()

	if messages := <-ch; messages[0].string != "a" {
		t.Fatalf("unexpected invlidation %v", messages)
	}

	go func() {
		mock.Expect().Reply(RedisMessage{
			typ: '>',
			values: []RedisMessage{
				{typ: '+', string: "invalidate"},
				{typ: '_'},
			},
		})
	}()

	if messages := <-ch; messages != nil {
		t.Fatalf("unexpected invlidation %v", messages)
	}

	go cancel()

	if messages := <-ch; messages != nil {
		t.Fatalf("unexpected invlidation %v", messages)
	}
}

func TestMultiHalfErr(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, mock, _, closeConn := setup(t, ClientOption{})

	expectCSC := func() {
		mock.Expect("CLIENT", "CACHING", "YES").
			Expect("MULTI").
			Expect("PTTL", "a").
			Expect("GET", "a").
			Expect("EXEC").
			ReplyString("OK").
			ReplyString("OK")
	}

	go func() {
		expectCSC()
		closeConn()
	}()

	err := p.DoCache(context.Background(), Cacheable(cmds.NewCompleted([]string{"GET", "a"})), 10*time.Second).Error()
	if err != io.EOF && !strings.HasPrefix(err.Error(), "io:") {
		t.Errorf("unexpected err %v", err)
	}
}

//gocyclo:ignore
func TestPubSub(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	builder := cmds.NewBuilder(cmds.NoSlot)
	t.Run("NoReply Commands In Do", func(t *testing.T) {
		p, mock, cancel, _ := setup(t, ClientOption{})
		defer cancel()

		commands := []Completed{
			builder.Subscribe().Channel("a").Build(),
			builder.Psubscribe().Pattern("b").Build(),
			builder.Unsubscribe().Channel("c").Build(),
			builder.Punsubscribe().Pattern("d").Build(),
			builder.Ssubscribe().Channel("c").Build(),
			builder.Sunsubscribe().Channel("d").Build(),
		}

		go func() {
			for _, c := range commands {
				mock.Expect(c.Commands()...).Reply(RedisMessage{typ: '>', values: []RedisMessage{
					{typ: '+', string: strings.ToLower(c.Commands()[0])},
					{typ: '+', string: strings.ToLower(c.Commands()[1])},
				}})
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

		commands := []Completed{
			builder.Subscribe().Channel("a").Build(),
			builder.Psubscribe().Pattern("b").Build(),
			builder.Unsubscribe().Channel("c").Build(),
			builder.Punsubscribe().Pattern("d").Build(),
		}

		go func() {
			for _, c := range commands {
				mock.Expect(c.Commands()...).Reply(RedisMessage{typ: '>', values: []RedisMessage{
					{typ: '+', string: strings.ToLower(c.Commands()[0])},
					{typ: '+', string: strings.ToLower(c.Commands()[1])},
				}})
			}
			mock.Expect("GET", "k").ReplyString("v")
		}()

		p.DoMulti(context.Background(), commands...)
		if v, _ := p.Do(context.Background(), builder.Get().Key("k").Build()).ToMessage(); v.string != "v" {
			t.Fatalf("no-reply commands should not affect nornal commands")
		}
	})

	t.Run("PubSub Subscribe RedisMessage", func(t *testing.T) {
		ctx := context.Background()
		p, mock, cancel, _ := setup(t, ClientOption{})

		activate := builder.Subscribe().Channel("1").Build()
		deactivate := builder.Unsubscribe().Channel("1").Build()
		go func() {
			mock.Expect(activate.Commands()...).Reply(
				RedisMessage{typ: '>', values: []RedisMessage{
					{typ: '+', string: "subscribe"},
					{typ: '+', string: "1"},
					{typ: ':', integer: 1},
				}},
				RedisMessage{typ: '>', values: []RedisMessage{
					{typ: '+', string: "message"},
					{typ: '+', string: "1"},
					{typ: '+', string: "2"},
				}},
			)
			mock.Expect(deactivate.Commands()...).Reply(
				RedisMessage{typ: '>', values: []RedisMessage{
					{typ: '+', string: "unsubscribe"},
					{typ: '+', string: "1"},
					{typ: ':', integer: 0},
				}},
			)
		}()

		if err := p.Receive(ctx, activate, func(msg PubSubMessage) {
			if msg.Channel == "1" && msg.Message == "2" {
				if err := p.Do(ctx, deactivate).Error(); err != nil {
					t.Fatalf("unexpected err %v", err)
				}
			}
		}); err != nil {
			t.Fatalf("unexpected err %v", err)
		}

		cancel()
	})

	t.Run("PubSub SSubscribe RedisMessage", func(t *testing.T) {
		ctx := context.Background()
		p, mock, cancel, _ := setup(t, ClientOption{})

		activate := builder.Ssubscribe().Channel("1").Build()
		deactivate := builder.Sunsubscribe().Channel("1").Build()
		go func() {
			mock.Expect(activate.Commands()...).Reply(
				RedisMessage{typ: '>', values: []RedisMessage{
					{typ: '+', string: "ssubscribe"},
					{typ: '+', string: "1"},
					{typ: ':', integer: 1},
				}},
				RedisMessage{typ: '>', values: []RedisMessage{
					{typ: '+', string: "smessage"},
					{typ: '+', string: "1"},
					{typ: '+', string: "2"},
				}},
			)
			mock.Expect(deactivate.Commands()...).Reply(
				RedisMessage{typ: '>', values: []RedisMessage{
					{typ: '+', string: "sunsubscribe"},
					{typ: '+', string: "1"},
					{typ: ':', integer: 0},
				}},
			)
		}()

		if err := p.Receive(ctx, activate, func(msg PubSubMessage) {
			if msg.Channel == "1" && msg.Message == "2" {
				if err := p.Do(ctx, deactivate).Error(); err != nil {
					t.Fatalf("unexpected err %v", err)
				}
			}
		}); err != nil {
			t.Fatalf("unexpected err %v", err)
		}

		cancel()
	})

	t.Run("PubSub PSubscribe RedisMessage", func(t *testing.T) {
		ctx := context.Background()
		p, mock, cancel, _ := setup(t, ClientOption{})

		activate := builder.Psubscribe().Pattern("1").Build()
		deactivate := builder.Punsubscribe().Pattern("1").Build()
		go func() {
			mock.Expect(activate.Commands()...).Reply(
				RedisMessage{typ: '>', values: []RedisMessage{
					{typ: '+', string: "psubscribe"},
					{typ: '+', string: "1"},
					{typ: ':', integer: 1},
				}},
				RedisMessage{typ: '>', values: []RedisMessage{
					{typ: '+', string: "pmessage"},
					{typ: '+', string: "1"},
					{typ: '+', string: "2"},
					{typ: '+', string: "3"},
				}},
			)
			mock.Expect(deactivate.Commands()...).Reply(
				RedisMessage{typ: '>', values: []RedisMessage{
					{typ: '+', string: "punsubscribe"},
					{typ: '+', string: "1"},
					{typ: ':', integer: 0},
				}},
			)
		}()

		if err := p.Receive(ctx, activate, func(msg PubSubMessage) {
			if msg.Pattern == "1" && msg.Channel == "2" && msg.Message == "3" {
				if err := p.Do(ctx, deactivate).Error(); err != nil {
					t.Fatalf("unexpected err %v", err)
				}
			}
		}); err != nil {
			t.Fatalf("unexpected err %v", err)
		}

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
					{typ: '+', string: "subscribe"},
					{typ: '+', string: "1"},
					{typ: ':', integer: 1},
				}},
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
		}); !errors.Is(err, context.Canceled) {
			t.Fatalf("unexpected err %v", err)
		}

		cancel()
	})

	t.Run("PubSub Subscribe Redis Error", func(t *testing.T) {
		ctx := context.Background()
		p, mock, cancel, _ := setup(t, ClientOption{})

		commands := []Completed{
			builder.Subscribe().Channel("1").Build(),
			builder.Psubscribe().Pattern("1").Build(),
			builder.Ssubscribe().Channel("1").Build(),
			builder.Unsubscribe().Channel("1").Build(),
			builder.Punsubscribe().Pattern("1").Build(),
			builder.Sunsubscribe().Channel("1").Build(),
		}
		go func() {
			for _, cmd := range commands {
				mock.Expect(cmd.Commands()...).Reply(RedisMessage{typ: '-', string: cmd.Commands()[0]})
			}
		}()
		for _, cmd := range commands {
			if err := p.Do(ctx, cmd).Error(); err == nil || !strings.Contains(err.Error(), cmd.Commands()[0]) {
				t.Fatalf("unexpected err %v", err)
			}
		}

		cancel()
	})

	t.Run("PubSub Subscribe Response", func(t *testing.T) {
		ctx := context.Background()
		p, mock, cancel, _ := setup(t, ClientOption{})

		commands := []Completed{
			builder.Subscribe().Channel("a", "b", "c").Build(),
			builder.Psubscribe().Pattern("a", "b", "c").Build(),
			builder.Ssubscribe().Channel("a", "b", "c").Build(),
			builder.Unsubscribe().Channel("a", "b", "c").Build(),
			builder.Punsubscribe().Pattern("a", "b", "c").Build(),
			builder.Sunsubscribe().Channel("a", "b", "c").Build(),
		}

		for i, cmd1 := range commands {
			cmd2 := builder.Get().Key(strconv.Itoa(i)).Build()
			go func() {
				mock.Expect(cmd1.Commands()...).Reply(
					RedisMessage{typ: '>', values: []RedisMessage{
						{typ: '+', string: "subscribe"},
						{typ: '+', string: "a"},
						{typ: ':', integer: 1},
					}},
					RedisMessage{typ: '>', values: []RedisMessage{ // skip
						{typ: '+', string: "subscribe"},
						{typ: '+', string: "b"},
						{typ: ':', integer: 1},
					}},
					RedisMessage{typ: '>', values: []RedisMessage{ // skip
						{typ: '+', string: "subscribe"},
						{typ: '+', string: "c"},
						{typ: ':', integer: 1},
					}},
				).Expect(cmd2.Commands()...).ReplyString(strconv.Itoa(i))
			}()

			if err := p.Do(ctx, cmd1).Error(); err != nil {
				t.Fatalf("unexpected err %v", err)
			}
			if v, err := p.Do(ctx, cmd2).ToString(); err != nil || v != strconv.Itoa(i) {
				t.Fatalf("unexpected val %v %v", v, err)
			}
		}
		cancel()
	})

	t.Run("PubSub Wildcard Unsubscribe Response", func(t *testing.T) {
		ctx := context.Background()
		p, mock, cancel, _ := setup(t, ClientOption{})

		commands := []Completed{
			builder.Unsubscribe().Build(),
			builder.Punsubscribe().Build(),
			builder.Sunsubscribe().Build(),
		}

		replies := [][]RedisMessage{{
			{
				typ: '>',
				values: []RedisMessage{
					{typ: '+', string: "unsubscribe"},
					{typ: '_'},
					{typ: ':', integer: 0},
				},
			},
		}, {
			{
				typ: '>',
				values: []RedisMessage{
					{typ: '+', string: "punsubscribe"},
					{typ: '+', string: "1"},
					{typ: ':', integer: 0},
				},
			},
		}, {
			{
				typ: '>',
				values: []RedisMessage{
					{typ: '+', string: "sunsubscribe"},
					{typ: '+', string: "2"},
					{typ: ':', integer: 0},
				},
			},
			{
				typ: '>',
				values: []RedisMessage{
					{typ: '+', string: "sunsubscribe"},
					{typ: '+', string: "3"},
					{typ: ':', integer: 0},
				},
			},
		}}

		for i, cmd1 := range commands {
			cmd2 := builder.Get().Key(strconv.Itoa(i)).Build()
			go func() {
				mock.Expect(cmd1.Commands()...).Reply(replies[i]...).Expect(cmd2.Commands()...).ReplyString(strconv.Itoa(i))
			}()

			if err := p.Do(ctx, cmd1).Error(); err != nil {
				t.Fatalf("unexpected err %v", err)
			}
			if v, err := p.Do(ctx, cmd2).ToString(); err != nil || v != strconv.Itoa(i) {
				t.Fatalf("unexpected val %v %v", v, err)
			}
		}
		cancel()
	})

	t.Run("PubSub Proactive UNSUBSCRIBE/PUNSUBSCRIBE/SUNSCRIBE", func(t *testing.T) {
		for _, command := range []string{
			"unsubscribe",
			"punsubscribe",
			"sunsubscribe",
		} {
			command := command
			t.Run(command, func(t *testing.T) {
				ctx := context.Background()
				p, mock, cancel, _ := setup(t, ClientOption{})

				commands := []Completed{
					builder.Sunsubscribe().Build(),
					builder.Ssubscribe().Channel("3").Build(),
				}

				replies := [][]RedisMessage{
					{
						{ // proactive unsubscribe before user unsubscribe
							typ: '>',
							values: []RedisMessage{
								{typ: '+', string: command},
								{typ: '+', string: "1"},
								{typ: ':', integer: 0},
							},
						},
						{ // proactive unsubscribe before user unsubscribe
							typ: '>',
							values: []RedisMessage{
								{typ: '+', string: command},
								{typ: '+', string: "2"},
								{typ: ':', integer: 0},
							},
						},
						{ // user unsubscribe
							typ: '>',
							values: []RedisMessage{
								{typ: '+', string: command},
								{typ: '_'},
								{typ: ':', integer: 0},
							},
						},
						{ // proactive unsubscribe after user unsubscribe
							typ: '>',
							values: []RedisMessage{
								{typ: '+', string: command},
								{typ: '_'},
								{typ: ':', integer: 0},
							},
						},
					},
					{
						{ // user ssubscribe
							typ: '>',
							values: []RedisMessage{
								{typ: '+', string: "ssubscribe"},
								{typ: '+', string: "3"},
								{typ: ':', integer: 0},
							},
						},
						{ // proactive unsubscribe after user ssubscribe
							typ: '>',
							values: []RedisMessage{
								{typ: '+', string: command},
								{typ: '+', string: "3"},
								{typ: ':', integer: 0},
							},
						},
					},
				}

				p.background()

				// proactive unsubscribe before other commands
				mock.Expect().Reply(RedisMessage{ // proactive unsubscribe before user unsubscribe
					typ: '>',
					values: []RedisMessage{
						{typ: '+', string: command},
						{typ: '+', string: "0"},
						{typ: ':', integer: 0},
					},
				})

				time.Sleep(time.Millisecond * 100)

				for i, cmd1 := range commands {
					cmd2 := builder.Get().Key(strconv.Itoa(i)).Build()
					go func() {
						mock.Expect(cmd1.Commands()...).Reply(replies[i]...).Expect(cmd2.Commands()...).ReplyString(strconv.Itoa(i))
					}()
					if err := p.Do(ctx, cmd1).Error(); err != nil {
						t.Fatalf("unexpected err %v", err)
					}
					if v, err := p.Do(ctx, cmd2).ToString(); err != nil || v != strconv.Itoa(i) {
						t.Fatalf("unexpected val %v %v", v, err)
					}
				}
				cancel()
			})
		}
	})

	t.Run("PubSub Unexpected Subscribe", func(t *testing.T) {
		shouldPanic := func(push string) (pass bool) {
			defer func() { pass = recover() == protocolbug }()

			p, mock, _, _ := setup(t, ClientOption{})
			atomic.StoreInt32(&p.state, 1)
			p.queue.PutOne(builder.Get().Key("a").Build())
			p.queue.NextWriteCmd()
			go func() {
				mock.Expect().Reply(RedisMessage{
					typ: '>', values: []RedisMessage{
						{typ: '+', string: push},
						{typ: '+', string: ""},
					},
				})
			}()
			p._backgroundRead()
			return
		}
		for _, push := range []string{
			"subscribe",
			"psubscribe",
			"ssubscribe",
		} {
			if !shouldPanic(push) {
				t.Fatalf("should panic on protocolbug")
			}
		}
	})

	t.Run("PubSub MULTI/EXEC Subscribe", func(t *testing.T) {
		shouldPanic := func(cmd Completed) (pass bool) {
			defer func() { pass = recover() == multiexecsub }()

			p, mock, _, _ := setup(t, ClientOption{})
			atomic.StoreInt32(&p.state, 1)
			p.queue.PutOne(cmd)
			p.queue.NextWriteCmd()
			go func() {
				mock.Expect().Reply(RedisMessage{typ: '+', string: "QUEUED"})
			}()
			p._backgroundRead()
			return
		}
		for _, push := range []Completed{
			builder.Subscribe().Channel("ch1").Build(),
			builder.Psubscribe().Pattern("ch1").Build(),
			builder.Ssubscribe().Channel("ch1").Build(),
			builder.Unsubscribe().Channel("ch1").Build(),
			builder.Punsubscribe().Pattern("ch1").Build(),
			builder.Sunsubscribe().Channel("ch1").Build(),
		} {
			if !shouldPanic(push) {
				t.Fatalf("should panic on protocolbug")
			}
		}
	})

	t.Run("RESP2 pubsub mixed", func(t *testing.T) {
		p, _, cancel, _ := setup(t, ClientOption{})
		p.version = 5
		defer cancel()

		commands := []Completed{
			builder.Subscribe().Channel("a").Build(),
			builder.Psubscribe().Pattern("b").Build(),
			builder.Get().Key("c").Build(),
		}
		for _, resp := range p.DoMulti(context.Background(), commands...).s {
			if e := resp.Error(); e != ErrRESP2PubSubMixed {
				t.Fatalf("unexpected err %v", e)
			}
		}
	})

	t.Run("RESP2 pubsub connect error", func(t *testing.T) {
		p, _, cancel, _ := setup(t, ClientOption{})
		p.version = 5
		e := errors.New("any")
		p.r2psFn = func() (p *pipe, err error) {
			return nil, e
		}
		defer cancel()

		if err := p.Receive(context.Background(), builder.Subscribe().Channel("a").Build(), nil); err != e {
			t.Fatalf("unexpected err %v", err)
		}

		if err := p.Do(context.Background(), builder.Subscribe().Channel("a").Build()).Error(); err != e {
			t.Fatalf("unexpected err %v", err)
		}

		if err := p.DoMulti(context.Background(), builder.Subscribe().Channel("a").Build()).s[0].Error(); err != e {
			t.Fatalf("unexpected err %v", err)
		}
	})
}

//gocyclo:ignore
func TestPubSubHooks(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	builder := cmds.NewBuilder(cmds.NoSlot)

	t.Run("Empty Hooks", func(t *testing.T) {
		p, _, cancel, _ := setup(t, ClientOption{})
		defer cancel()
		if ch := p.SetPubSubHooks(PubSubHooks{}); ch != nil {
			t.Fatalf("unexpected ch %v", ch)
		}
	})

	t.Run("Close on error", func(t *testing.T) {
		p, _, cancel, closeConn := setup(t, ClientOption{})
		defer cancel()
		ch := p.SetPubSubHooks(PubSubHooks{
			OnMessage: func(m PubSubMessage) {},
		})
		closeConn()
		if err := <-ch; err != io.EOF && !strings.HasPrefix(err.Error(), "io:") {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("Swap Hooks", func(t *testing.T) {
		p, _, cancel, _ := setup(t, ClientOption{})
		defer cancel()
		ch1 := p.SetPubSubHooks(PubSubHooks{
			OnMessage: func(m PubSubMessage) {},
		})
		ch2 := p.SetPubSubHooks(PubSubHooks{
			OnSubscription: func(s PubSubSubscription) {},
		})
		if err := <-ch1; err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		ch3 := p.SetPubSubHooks(PubSubHooks{})
		if err := <-ch2; err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if ch3 != nil {
			t.Fatalf("unexpected ch %v", ch3)
		}
	})

	t.Run("PubSubHooks OnSubscription", func(t *testing.T) {
		ctx := context.Background()
		p, mock, cancel, _ := setup(t, ClientOption{})

		var s1, s2, u1, u2 bool

		ch := p.SetPubSubHooks(PubSubHooks{
			OnSubscription: func(s PubSubSubscription) {
				if s.Kind == "subscribe" && s.Channel == "1" && s.Count == 1 {
					s1 = true
				}
				if s.Kind == "psubscribe" && s.Channel == "2" && s.Count == 2 {
					s2 = true
				}
				if s.Kind == "unsubscribe" && s.Channel == "1" && s.Count == 1 {
					u1 = true
				}
				if s.Kind == "punsubscribe" && s.Channel == "2" && s.Count == 2 {
					u2 = true
				}
			},
		})

		activate1 := builder.Subscribe().Channel("1").Build()
		activate2 := builder.Psubscribe().Pattern("2").Build()
		deactivate1 := builder.Unsubscribe().Channel("1").Build()
		deactivate2 := builder.Punsubscribe().Pattern("2").Build()
		go func() {
			mock.Expect(activate1.Commands()...).Expect(activate2.Commands()...).Reply(
				RedisMessage{typ: '>', values: []RedisMessage{
					{typ: '+', string: "subscribe"},
					{typ: '+', string: "1"},
					{typ: ':', integer: 1},
				}},
				RedisMessage{typ: '>', values: []RedisMessage{
					{typ: '+', string: "psubscribe"},
					{typ: '+', string: "2"},
					{typ: ':', integer: 2},
				}},
				RedisMessage{typ: '>', values: []RedisMessage{
					{typ: '+', string: "message"},
					{typ: '+', string: "1"},
					{typ: '+', string: "11"},
				}},
				RedisMessage{typ: '>', values: []RedisMessage{
					{typ: '+', string: "pmessage"},
					{typ: '+', string: "2"},
					{typ: '+', string: "22"},
					{typ: '+', string: "222"},
				}},
			)
			mock.Expect(deactivate1.Commands()...).Expect(deactivate2.Commands()...).Reply(
				RedisMessage{typ: '>', values: []RedisMessage{
					{typ: '+', string: "unsubscribe"},
					{typ: '+', string: "1"},
					{typ: ':', integer: 1},
				}},
				RedisMessage{typ: '>', values: []RedisMessage{
					{typ: '+', string: "punsubscribe"},
					{typ: '+', string: "2"},
					{typ: ':', integer: 2},
				}},
			)
			cancel()
		}()

		for _, r := range p.DoMulti(ctx, activate1, activate2).s {
			if err := r.Error(); err != nil {
				t.Fatalf("unexpected err %v", err)
			}
		}
		for _, r := range p.DoMulti(ctx, deactivate1, deactivate2).s {
			if err := r.Error(); err != nil {
				t.Fatalf("unexpected err %v", err)
			}
		}
		if err := <-ch; err != ErrClosing {
			t.Fatalf("unexpected err %v", err)
		}
		if !s1 {
			t.Fatalf("unexpecetd s1")
		}
		if !s2 {
			t.Fatalf("unexpecetd s2")
		}
		if !u1 {
			t.Fatalf("unexpecetd u1")
		}
		if !u2 {
			t.Fatalf("unexpecetd u2")
		}
	})

	t.Run("PubSubHooks OnMessage", func(t *testing.T) {
		ctx := context.Background()
		p, mock, cancel, _ := setup(t, ClientOption{})

		var m1, m2 bool

		ch := p.SetPubSubHooks(PubSubHooks{
			OnMessage: func(m PubSubMessage) {
				if m.Channel == "1" && m.Message == "11" {
					m1 = true
				}
				if m.Pattern == "2" && m.Channel == "22" && m.Message == "222" {
					m2 = true
				}
			},
		})

		activate1 := builder.Subscribe().Channel("1").Build()
		activate2 := builder.Psubscribe().Pattern("2").Build()
		deactivate1 := builder.Unsubscribe().Channel("1").Build()
		deactivate2 := builder.Punsubscribe().Pattern("2").Build()
		go func() {
			mock.Expect(activate1.Commands()...).Expect(activate2.Commands()...).Reply(
				RedisMessage{typ: '>', values: []RedisMessage{
					{typ: '+', string: "subscribe"},
					{typ: '+', string: "1"},
					{typ: ':', integer: 1},
				}},
				RedisMessage{typ: '>', values: []RedisMessage{
					{typ: '+', string: "psubscribe"},
					{typ: '+', string: "2"},
					{typ: ':', integer: 2},
				}},
				RedisMessage{typ: '>', values: []RedisMessage{
					{typ: '+', string: "message"},
					{typ: '+', string: "1"},
					{typ: '+', string: "11"},
				}},
				RedisMessage{typ: '>', values: []RedisMessage{
					{typ: '+', string: "pmessage"},
					{typ: '+', string: "2"},
					{typ: '+', string: "22"},
					{typ: '+', string: "222"},
				}},
			)
			mock.Expect(deactivate1.Commands()...).Expect(deactivate2.Commands()...).Reply(
				RedisMessage{typ: '>', values: []RedisMessage{
					{typ: '+', string: "unsubscribe"},
					{typ: '+', string: "1"},
					{typ: ':', integer: 1},
				}},
				RedisMessage{typ: '>', values: []RedisMessage{
					{typ: '+', string: "punsubscribe"},
					{typ: '+', string: "2"},
					{typ: ':', integer: 2},
				}},
			)
			cancel()
		}()

		for _, r := range p.DoMulti(ctx, activate1, activate2).s {
			if err := r.Error(); err != nil {
				t.Fatalf("unexpected err %v", err)
			}
		}
		for _, r := range p.DoMulti(ctx, deactivate1, deactivate2).s {
			if err := r.Error(); err != nil {
				t.Fatalf("unexpected err %v", err)
			}
		}
		if err := <-ch; err != ErrClosing {
			t.Fatalf("unexpected err %v", err)
		}
		if !m1 {
			t.Fatalf("unexpecetd m1")
		}
		if !m2 {
			t.Fatalf("unexpecetd m2")
		}
	})
}

func TestExitOnWriteError(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, _, _, closeConn := setup(t, ClientOption{})

	closeConn()

	for i := 0; i < 2; i++ {
		if err := p.Do(context.Background(), cmds.NewCompleted([]string{"GET", "a"})).NonRedisError(); err != io.EOF && !strings.HasPrefix(err.Error(), "io:") {
			t.Errorf("unexpected cached result, expected io err, got %v", err)
		}
	}
}

func TestExitOnPubSubSubscribeWriteError(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
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
	defer ShouldNotLeaked(SetupLeakDetection())
	p, _, _, closeConn := setup(t, ClientOption{})

	closeConn()

	for i := 0; i < 2; i++ {
		if err := p.DoMulti(context.Background(), cmds.NewCompleted([]string{"GET", "a"})).s[0].NonRedisError(); err != io.EOF && !strings.HasPrefix(err.Error(), "io:") {
			t.Errorf("unexpected result, expected io err, got %v", err)
		}
	}
}

func TestExitOnRingFullAndConnError(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, mock, _, closeConn := setup(t, ClientOption{
		RingScaleEachConn: 1,
	})
	p.background()

	// fill the ring
	for i := 0; i < len(p.queue.(*ring).store); i++ {
		go func() {
			if err := p.Do(context.Background(), cmds.NewCompleted([]string{"GET", "a"})).Error(); err != io.EOF && !strings.HasPrefix(err.Error(), "io:") {
				t.Errorf("unexpected result, expected io err, got %v", err)
			}
		}()
	}
	// let writer loop over the ring
	for i := 0; i < len(p.queue.(*ring).store); i++ {
		mock.Expect("GET", "a")
	}

	time.Sleep(time.Second) // make sure the writer is waiting for the next write
	closeConn()

	if err := p.Do(context.Background(), cmds.NewCompleted([]string{"GET", "a"})).Error(); err != io.EOF && !strings.HasPrefix(err.Error(), "io:") {
		t.Errorf("unexpected result, expected io err, got %v", err)
	}
}

func TestExitOnRingFullAndPingTimout(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, mock, _, _ := setup(t, ClientOption{
		RingScaleEachConn: 1,
		ConnWriteTimeout:  500 * time.Millisecond,
		Dialer:            net.Dialer{KeepAlive: 500 * time.Millisecond},
	})
	p.background()

	// fill the ring
	for i := 0; i < len(p.queue.(*ring).store); i++ {
		go func() {
			if err := p.Do(context.Background(), cmds.NewCompleted([]string{"GET", "a"})).Error(); err != context.DeadlineExceeded {
				t.Errorf("unexpected result, expected context.DeadlineExceeded, got %v", err)
			}
		}()
	}
	// let writer loop over the ring
	for i := 0; i < len(p.queue.(*ring).store); i++ {
		mock.Expect("GET", "a")
	}

	if err := p.Do(context.Background(), cmds.NewCompleted([]string{"GET", "a"})).Error(); err != context.DeadlineExceeded {
		t.Errorf("unexpected result, expected context.DeadlineExceeded, got %v", err)
	}
}

func TestExitAllGoroutineOnWriteError(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
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
			if err := conn.DoMulti(context.Background(), cmds.NewCompleted([]string{"GET", "a"})).s[0].NonRedisError(); err != io.EOF && !strings.HasPrefix(err.Error(), "io:") {
				t.Errorf("unexpected result, expected io err, got %v", err)
			}
		}()
	}
	wg.Wait()
}

func TestExitOnReadError(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
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
	defer ShouldNotLeaked(SetupLeakDetection())
	p, mock, _, closeConn := setup(t, ClientOption{})

	go func() {
		mock.Expect("GET", "a")
		closeConn()
	}()

	for i := 0; i < 2; i++ {
		if err := p.DoMulti(context.Background(), cmds.NewCompleted([]string{"GET", "a"})).s[0].NonRedisError(); err != io.EOF && !strings.HasPrefix(err.Error(), "io:") {
			t.Errorf("unexpected result, expected io err, got %v", err)
		}
	}
}

func TestExitAllGoroutineOnReadError(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
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
			if err := p.DoMulti(context.Background(), cmds.NewCompleted([]string{"GET", "a"})).s[0].NonRedisError(); err != io.EOF && !strings.HasPrefix(err.Error(), "io:") {
				t.Errorf("unexpected result, expected io err, got %v", err)
			}
		}()
	}
	wg.Wait()
}

func TestCloseAndWaitPendingCMDs(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
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
	mock.Expect("PING").ReplyString("OK")
	mock.Close()
	wg.Wait()
}

func TestCloseWithGracefulPeriodExceeded(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, mock, _, _ := setup(t, ClientOption{})
	go func() {
		p.Close()
	}()
	mock.Expect("PING")
	<-p.close
}

func TestCloseWithPipeliningAndGracefulPeriodExceeded(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, mock, _, _ := setup(t, ClientOption{AlwaysPipelining: true})
	go func() {
		p.Close()
	}()
	mock.Expect("PING")
	<-p.close
}

func TestAlreadyCanceledContext(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, _, close, closeConn := setup(t, ClientOption{})
	defer closeConn()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	if err := p.Do(ctx, cmds.NewCompleted([]string{"GET", "a"})).NonRedisError(); !errors.Is(err, context.Canceled) {
		t.Fatalf("unexpected err %v", err)
	}
	if err := p.DoMulti(ctx, cmds.NewCompleted([]string{"GET", "a"})).s[0].NonRedisError(); !errors.Is(err, context.Canceled) {
		t.Fatalf("unexpected err %v", err)
	}

	cp := newPool(1, nil, nil)
	if s := p.DoStream(ctx, cp, cmds.NewCompleted([]string{"GET", "a"})); !errors.Is(s.Error(), context.Canceled) {
		t.Fatalf("unexpected err %v", s.Error())
	}
	if s := p.DoMultiStream(ctx, cp, cmds.NewCompleted([]string{"GET", "a"})); !errors.Is(s.Error(), context.Canceled) {
		t.Fatalf("unexpected err %v", s.Error())
	}

	ctx, cancel = context.WithDeadline(context.Background(), time.Now().Add(-1*time.Second))
	cancel()

	if err := p.Do(ctx, cmds.NewCompleted([]string{"GET", "a"})).NonRedisError(); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("unexpected err %v", err)
	}
	if err := p.DoMulti(ctx, cmds.NewCompleted([]string{"GET", "a"})).s[0].NonRedisError(); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("unexpected err %v", err)
	}
	close()
}

func TestCancelContext_Do(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, mock, shutdown, _ := setup(t, ClientOption{})

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		mock.Expect("GET", "a")
		cancel()
		mock.Expect().ReplyString("OK")
	}()

	if err := p.Do(ctx, cmds.NewCompleted([]string{"GET", "a"})).NonRedisError(); !errors.Is(err, context.Canceled) {
		t.Fatalf("unexpected err %v", err)
	}
	shutdown()
}

func TestCancelContext_DoStream(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, _, _, _ := setup(t, ClientOption{})

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	defer cancel()

	cp := newPool(1, nil, nil)
	s := p.DoStream(ctx, cp, cmds.NewCompleted([]string{"GET", "a"}))
	if err := s.Error(); err != io.EOF && !strings.Contains(err.Error(), "i/o") {
		t.Fatalf("unexpected err %v", err)
	}
	if len(cp.list) != 0 {
		t.Fatalf("unexpected pool length %v", len(cp.list))
	}
}

func TestWriteDeadlineIsShorterThanContextDeadline_DoStream(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, _, _, _ := setup(t, ClientOption{ConnWriteTimeout: 100 * time.Millisecond})

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	cp := newPool(1, nil, nil)
	startTime := time.Now()
	s := p.DoStream(ctx, cp, cmds.NewCompleted([]string{"GET", "a"}))
	if err := s.Error(); err != io.EOF && !strings.Contains(err.Error(), "i/o") {
		t.Fatalf("unexpected err %v", err)
	}
	if time.Since(startTime) >= time.Second {
		t.Fatalf("unexpected time %v", time.Since(startTime))
	}
	if len(cp.list) != 0 {
		t.Fatalf("unexpected pool length %v", len(cp.list))
	}
}

func TestCancelContext_Do_Block(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, mock, shutdown, _ := setup(t, ClientOption{})

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		mock.Expect("GET", "a")
		cancel()
		mock.Expect().ReplyString("OK")
	}()

	if err := p.Do(ctx, cmds.NewBlockingCompleted([]string{"GET", "a"})).NonRedisError(); !errors.Is(err, context.Canceled) {
		t.Fatalf("unexpected err %v", err)
	}
	shutdown()
}

func TestCancelContext_DoMulti(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, mock, shutdown, _ := setup(t, ClientOption{})

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		mock.Expect("GET", "a")
		cancel()
		mock.Expect().ReplyString("OK")
	}()

	if err := p.DoMulti(ctx, cmds.NewCompleted([]string{"GET", "a"})).s[0].NonRedisError(); !errors.Is(err, context.Canceled) {
		t.Fatalf("unexpected err %v", err)
	}
	shutdown()
}

func TestCancelContext_DoMulti_Block(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, mock, shutdown, _ := setup(t, ClientOption{})

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		mock.Expect("GET", "a")
		cancel()
		mock.Expect().ReplyString("OK")
	}()

	if err := p.DoMulti(ctx, cmds.NewBlockingCompleted([]string{"GET", "a"})).s[0].NonRedisError(); !errors.Is(err, context.Canceled) {
		t.Fatalf("unexpected err %v", err)
	}
	shutdown()
}

func TestCancelContext_DoMultiStream(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, _, _, _ := setup(t, ClientOption{})

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	defer cancel()

	cp := newPool(1, nil, nil)
	s := p.DoMultiStream(ctx, cp, cmds.NewCompleted([]string{"GET", "a"}))
	if err := s.Error(); err != io.EOF && !strings.Contains(err.Error(), "i/o") {
		t.Fatalf("unexpected err %v", err)
	}
	if len(cp.list) != 0 {
		t.Fatalf("unexpected pool length %v", len(cp.list))
	}
}

func TestWriteDeadlineIsShorterThanContextDeadline_DoMultiStream(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, _, _, _ := setup(t, ClientOption{ConnWriteTimeout: 100 * time.Millisecond})

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	cp := newPool(1, nil, nil)
	startTime := time.Now()
	s := p.DoMultiStream(ctx, cp, cmds.NewCompleted([]string{"GET", "a"}))
	if err := s.Error(); err != io.EOF && !strings.Contains(err.Error(), "i/o") {
		t.Fatalf("unexpected err %v", err)
	}
	if time.Since(startTime) >= time.Second {
		t.Fatalf("unexpected time %v", time.Since(startTime))
	}
	if len(cp.list) != 0 {
		t.Fatalf("unexpected pool length %v", len(cp.list))
	}
}

func TestForceClose_Do_Block(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, mock, _, _ := setup(t, ClientOption{})

	go func() {
		mock.Expect("GET", "a")
		p.Close()
	}()

	if err := p.Do(context.Background(), cmds.NewBlockingCompleted([]string{"GET", "a"})).NonRedisError(); err != io.EOF && !strings.HasPrefix(err.Error(), "io:") {
		t.Fatalf("unexpected err %v", err)
	}
}

func TestTimeout_DoStream(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, _, _, _ := setup(t, ClientOption{ConnWriteTimeout: time.Millisecond * 30})

	cp := newPool(1, nil, nil)

	s := p.DoStream(context.Background(), cp, cmds.NewCompleted([]string{"GET", "a"}))
	if err := s.Error(); err != io.EOF && !strings.Contains(err.Error(), "i/o") {
		t.Fatalf("unexpected err %v", s.Error())
	}
	if len(cp.list) != 0 {
		t.Fatalf("unexpected pool length %v", len(cp.list))
	}
}

func TestForceClose_DoStream_Block(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, mock, _, _ := setup(t, ClientOption{ConnWriteTimeout: time.Second})

	go func() {
		mock.Expect("GET", "a")
		p.Close()
	}()

	cp := newPool(1, nil, nil)

	s := p.DoStream(context.Background(), cp, cmds.NewBlockingCompleted([]string{"GET", "a"}))
	if s.Error() != nil {
		t.Fatalf("unexpected err %v", s.Error())
	}
	buf := bytes.NewBuffer(nil)
	for s.HasNext() {
		n, err := s.WriteTo(buf)
		if err != io.EOF && !strings.HasPrefix(err.Error(), "io:") {
			t.Errorf("unexpected err %v\n", err)
		}
		if n != 0 {
			t.Errorf("unexpected n %v\n", n)
		}
	}
	if len(cp.list) != 0 {
		t.Fatalf("unexpected pool length %v", len(cp.list))
	}
}

func TestForceClose_Do_Canceled_Block(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, mock, _, _ := setup(t, ClientOption{})

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		mock.Expect("GET", "a")
		cancel()
		mock.Expect().ReplyString("OK")
	}()

	if err := p.Do(ctx, cmds.NewBlockingCompleted([]string{"GET", "a"})).NonRedisError(); !errors.Is(err, context.Canceled) {
		t.Fatalf("unexpected err %v", err)
	}
	p.Close()
}

func TestForceClose_DoMulti_Block(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, mock, _, _ := setup(t, ClientOption{})

	go func() {
		mock.Expect("GET", "a")
		p.Close()
	}()

	if err := p.DoMulti(context.Background(), cmds.NewBlockingCompleted([]string{"GET", "a"})).s[0].NonRedisError(); err != io.EOF && !strings.HasPrefix(err.Error(), "io:") {
		t.Fatalf("unexpected err %v", err)
	}
}

func TestTimeout_DoMultiStream(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, _, _, _ := setup(t, ClientOption{ConnWriteTimeout: time.Millisecond * 30})

	cp := newPool(1, nil, nil)

	s := p.DoMultiStream(context.Background(), cp, cmds.NewCompleted([]string{"GET", "a"}))
	if err := s.Error(); err != io.EOF && !strings.Contains(err.Error(), "i/o") {
		t.Fatalf("unexpected err %v", s.Error())
	}
	if len(cp.list) != 0 {
		t.Fatalf("unexpected pool length %v", len(cp.list))
	}
}

func TestForceClose_DoMultiStream_Block(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, mock, _, _ := setup(t, ClientOption{ConnWriteTimeout: time.Second})

	go func() {
		mock.Expect("GET", "a")
		p.Close()
	}()

	cp := newPool(1, nil, nil)

	s := p.DoMultiStream(context.Background(), cp, cmds.NewBlockingCompleted([]string{"GET", "a"}))
	if s.Error() != nil {
		t.Fatalf("unexpected err %v", s.Error())
	}
	buf := bytes.NewBuffer(nil)
	for s.HasNext() {
		n, err := s.WriteTo(buf)
		if err != io.EOF && !strings.HasPrefix(err.Error(), "io:") {
			t.Errorf("unexpected err %v\n", err)
		}
		if n != 0 {
			t.Errorf("unexpected n %v\n", n)
		}
	}
	if len(cp.list) != 0 {
		t.Fatalf("unexpected pool length %v", len(cp.list))
	}
}

func TestForceClose_DoMulti_Canceled_Block(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, mock, _, _ := setup(t, ClientOption{})

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		mock.Expect("GET", "a")
		cancel()
		mock.Expect().ReplyString("OK")
	}()

	if err := p.DoMulti(ctx, cmds.NewBlockingCompleted([]string{"GET", "a"})).s[0].NonRedisError(); !errors.Is(err, context.Canceled) {
		t.Fatalf("unexpected err %v", err)
	}
	p.Close()
}

func TestSyncModeSwitchingWithDeadlineExceed_Do(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, mock, _, closeConn := setup(t, ClientOption{})
	defer closeConn()

	ctx, cancel := context.WithTimeout(context.Background(), time.Microsecond*100)
	defer cancel()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			if err := p.Do(ctx, cmds.NewCompleted([]string{"GET", "a"})).NonRedisError(); !errors.Is(err, context.DeadlineExceeded) {
				t.Errorf("unexpected err %v", err)
			}
			wg.Done()
		}()
	}

	mock.Expect("GET", "a")
	time.Sleep(time.Second / 2)
	mock.Expect().ReplyString("OK")
	wg.Wait()
	p.Close()
}

func TestSyncModeSwitchingWithDeadlineExceed_DoMulti(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, mock, _, closeConn := setup(t, ClientOption{})
	defer closeConn()

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			if err := p.DoMulti(ctx, cmds.NewCompleted([]string{"GET", "a"})).s[0].NonRedisError(); !errors.Is(err, context.DeadlineExceeded) {
				t.Errorf("unexpected err %v", err)
			}
			wg.Done()
		}()
	}

	mock.Expect("GET", "a")
	time.Sleep(time.Second / 2)
	mock.Expect().ReplyString("OK")
	wg.Wait()
	p.Close()
}

func TestOngoingDeadlineContextInSyncMode_Do(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, _, _, closeConn := setup(t, ClientOption{})
	defer closeConn()

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(1*time.Second/2))
	defer cancel()

	if err := p.Do(ctx, cmds.NewCompleted([]string{"GET", "a"})).NonRedisError(); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("unexpected err %v", err)
	}
	p.Close()
}

func TestWriteDeadlineInSyncMode_Do(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, _, _, closeConn := setup(t, ClientOption{ConnWriteTimeout: 1 * time.Second / 2, Dialer: net.Dialer{KeepAlive: time.Second / 3}})
	defer closeConn()

	if err := p.Do(context.Background(), cmds.NewCompleted([]string{"GET", "a"})).NonRedisError(); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("unexpected err %v", err)
	}
	p.Close()
}

func TestWriteDeadlineIsShorterThanContextDeadlineInSyncMode_Do(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, _, _, closeConn := setup(t, ClientOption{ConnWriteTimeout: 100 * time.Millisecond, Dialer: net.Dialer{KeepAlive: time.Second}})
	defer closeConn()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	startTime := time.Now()
	if err := p.Do(ctx, cmds.NewCompleted([]string{"GET", "a"})).NonRedisError(); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("unexpected err %v", err)
	}

	if time.Since(startTime) >= time.Second {
		t.Fatalf("unexpected time %v", time.Since(startTime))
	}

	p.Close()
}

func TestOngoingDeadlineContextInSyncMode_DoMulti(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, _, _, closeConn := setup(t, ClientOption{})
	defer closeConn()

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(1*time.Second/2))
	defer cancel()

	if err := p.DoMulti(ctx, cmds.NewCompleted([]string{"GET", "a"})).s[0].NonRedisError(); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("unexpected err %v", err)
	}
	p.Close()
}

func TestWriteDeadlineInSyncMode_DoMulti(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, _, _, closeConn := setup(t, ClientOption{ConnWriteTimeout: time.Second / 2, Dialer: net.Dialer{KeepAlive: time.Second / 3}})
	defer closeConn()

	if err := p.DoMulti(context.Background(), cmds.NewCompleted([]string{"GET", "a"})).s[0].NonRedisError(); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("unexpected err %v", err)
	}
	p.Close()
}

func TestWriteDeadlineIsShorterThanContextDeadlineInSyncMode_DoMulti(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, _, _, closeConn := setup(t, ClientOption{ConnWriteTimeout: 100 * time.Millisecond, Dialer: net.Dialer{KeepAlive: time.Second}})
	defer closeConn()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	startTime := time.Now()
	if err := p.DoMulti(ctx, cmds.NewCompleted([]string{"GET", "a"})).s[0].NonRedisError(); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("unexpected err %v", err)
	}

	if time.Since(startTime) >= time.Second {
		t.Fatalf("unexpected time %v", time.Since(startTime))
	}

	p.Close()
}

func TestOngoingCancelContextInPipelineMode_Do(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, mock, close, closeConn := setup(t, ClientOption{})
	defer closeConn()

	p.background()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	canceled := int32(0)

	for i := 0; i < 5; i++ {
		go func() {
			_, err := p.Do(ctx, cmds.NewCompleted([]string{"GET", "a"})).ToString()
			if errors.Is(err, context.Canceled) {
				atomic.AddInt32(&canceled, 1)
			} else {
				t.Errorf("unexpected err %v", err)
			}
		}()
	}

	for atomic.LoadInt32(&p.waits) != 5 {
		t.Logf("wait p.waits to be 5 %v", atomic.LoadInt32(&p.waits))
		time.Sleep(time.Millisecond * 100)
	}

	cancel()

	for atomic.LoadInt32(&canceled) != 5 {
		t.Logf("wait canceled count to be 5 %v", atomic.LoadInt32(&canceled))
		time.Sleep(time.Millisecond * 100)
	}
	// the rest command is still send
	for i := 0; i < 5; i++ {
		mock.Expect("GET", "a").ReplyString("OK")
	}
	close()
}

func TestOngoingWriteTimeoutInPipelineMode_Do(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, _, _, closeConn := setup(t, ClientOption{ConnWriteTimeout: time.Second / 2, Dialer: net.Dialer{KeepAlive: time.Second / 3}})
	defer closeConn()

	p.background()

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	timeout := int32(0)

	for i := 0; i < 5; i++ {
		go func() {
			_, err := p.Do(ctx, cmds.NewCompleted([]string{"GET", "a"})).ToString()
			if errors.Is(err, context.DeadlineExceeded) {
				atomic.AddInt32(&timeout, 1)
			} else {
				t.Errorf("unexpected err %v", err)
			}
		}()
	}
	for atomic.LoadInt32(&p.waits) != 5 {
		t.Logf("wait p.waits to be 5 %v", atomic.LoadInt32(&p.waits))
		time.Sleep(time.Millisecond * 100)
	}
	for atomic.LoadInt32(&timeout) != 5 {
		t.Logf("wait timeout count to be 5 %v", atomic.LoadInt32(&timeout))
		time.Sleep(time.Millisecond * 100)
	}
	p.Close()
}

func TestOngoingCancelContextInPipelineMode_DoMulti(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, mock, close, closeConn := setup(t, ClientOption{})
	defer closeConn()

	p.background()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	canceled := int32(0)

	for i := 0; i < 5; i++ {
		go func() {
			_, err := p.DoMulti(ctx, cmds.NewCompleted([]string{"GET", "a"})).s[0].ToString()
			if errors.Is(err, context.Canceled) {
				atomic.AddInt32(&canceled, 1)
			} else {
				t.Errorf("unexpected err %v", err)
			}
		}()
	}

	for atomic.LoadInt32(&p.waits) != 5 {
		t.Logf("wait p.waits to be 5 %v", atomic.LoadInt32(&p.waits))
		time.Sleep(time.Millisecond * 100)
	}

	cancel()

	for atomic.LoadInt32(&canceled) != 5 {
		t.Logf("wait canceled count to be 5 %v", atomic.LoadInt32(&canceled))
		time.Sleep(time.Millisecond * 100)
	}
	// the rest command is still send
	for i := 0; i < 5; i++ {
		mock.Expect("GET", "a").ReplyString("OK")
	}
	close()
}

func TestOngoingWriteTimeoutInPipelineMode_DoMulti(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, _, _, closeConn := setup(t, ClientOption{ConnWriteTimeout: time.Second / 2, Dialer: net.Dialer{KeepAlive: time.Second / 3}})
	defer closeConn()

	p.background()

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	timeout := int32(0)

	for i := 0; i < 5; i++ {
		go func() {
			_, err := p.DoMulti(ctx, cmds.NewCompleted([]string{"GET", "a"})).s[0].ToString()
			if errors.Is(err, context.DeadlineExceeded) {
				atomic.AddInt32(&timeout, 1)
			} else {
				t.Errorf("unexpecetd err %v", err)
			}
		}()
	}
	for atomic.LoadInt32(&p.waits) != 5 {
		t.Logf("wait p.waits to be 5 %v", atomic.LoadInt32(&p.waits))
		time.Sleep(time.Millisecond * 100)
	}
	for atomic.LoadInt32(&timeout) != 5 {
		t.Logf("wait timeout count to be 5 %v", atomic.LoadInt32(&timeout))
		time.Sleep(time.Millisecond * 100)
	}
	p.Close()
}

func TestPipe_CleanSubscriptions_6(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, mock, cancel, _ := setup(t, ClientOption{ConnWriteTimeout: time.Second / 2, Dialer: net.Dialer{KeepAlive: time.Second / 3}})
	defer cancel()
	p.background()
	go func() {
		p.CleanSubscriptions()
	}()
	mock.Expect("UNSUBSCRIBE").Expect("PUNSUBSCRIBE").Expect("DISCARD").Reply(
		RedisMessage{typ: '>', values: []RedisMessage{
			{typ: '+', string: "unsubscribe"},
			{typ: '_'},
			{typ: ':', integer: 1},
		}},
		RedisMessage{typ: '>', values: []RedisMessage{
			{typ: '+', string: "punsubscribe"},
			{typ: '_'},
			{typ: ':', integer: 2},
		}},
		RedisMessage{typ: '+', string: "OK"},
	)
}

func TestPipe_CleanSubscriptions_7(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	p, mock, cancel, _ := setup(t, ClientOption{ConnWriteTimeout: time.Second / 2, Dialer: net.Dialer{KeepAlive: time.Second / 3}})
	p.version = 7
	defer cancel()
	p.background()
	go func() {
		p.CleanSubscriptions()
	}()
	mock.Expect("UNSUBSCRIBE").Expect("PUNSUBSCRIBE").Expect("SUNSUBSCRIBE").Expect("DISCARD").Reply(
		RedisMessage{typ: '>', values: []RedisMessage{
			{typ: '+', string: "unsubscribe"},
			{typ: '_'},
			{typ: ':', integer: 1},
		}},
		RedisMessage{typ: '>', values: []RedisMessage{
			{typ: '+', string: "punsubscribe"},
			{typ: '_'},
			{typ: ':', integer: 2},
		}},
		RedisMessage{typ: '>', values: []RedisMessage{
			{typ: '+', string: "sunsubscribe"},
			{typ: '_'},
			{typ: ':', integer: 3},
		}},
		RedisMessage{typ: '+', string: "OK"},
	)
}

func TestPingOnConnError(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	t.Run("sync", func(t *testing.T) {
		p, mock, _, closeConn := setup(t, ClientOption{ConnWriteTimeout: 3 * time.Second, Dialer: net.Dialer{KeepAlive: time.Second / 3}})
		mock.Expect("PING")
		closeConn()
		time.Sleep(time.Second / 2)
		p.Close()
		if err := p.Error(); err != io.EOF && !strings.HasPrefix(err.Error(), "io:") {
			t.Fatalf("unexpect err %v", err)
		}
	})
	t.Run("pipelining", func(t *testing.T) {
		p, mock, _, closeConn := setup(t, ClientOption{ConnWriteTimeout: 3 * time.Second, Dialer: net.Dialer{KeepAlive: time.Second / 3}})
		p.background()
		mock.Expect("PING")
		closeConn()
		time.Sleep(time.Second / 2)
		p.Close()
		if err := p.Error(); err != io.EOF && !strings.HasPrefix(err.Error(), "io:") {
			t.Fatalf("unexpect err %v", err)
		}
	})
}

//gocyclo:ignore
func TestBlockingCommandNoDeadline(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	// blocking command should not apply timeout
	timeout := 100 * time.Millisecond
	t.Run("sync do", func(t *testing.T) {
		p, mock, cancel, _ := setup(t, ClientOption{ConnWriteTimeout: timeout})
		defer cancel()
		go func() {
			time.Sleep(2 * timeout)
			mock.Expect("BLOCK").ReplyString("OK")
		}()
		if val, err := p.Do(context.Background(), cmds.NewBlockingCompleted([]string{"BLOCK"})).ToString(); err != nil || val != "OK" {
			t.Fatalf("unexpect resp %v %v", err, val)
		}
	})
	t.Run("sync do multi", func(t *testing.T) {
		p, mock, cancel, _ := setup(t, ClientOption{ConnWriteTimeout: timeout})
		defer cancel()
		go func() {
			time.Sleep(3 * timeout)
			mock.Expect("READ").ReplyString("READ").
				Expect("BLOCK").ReplyString("OK")
		}()
		if val, err := p.DoMulti(context.Background(),
			cmds.NewReadOnlyCompleted([]string{"READ"}),
			cmds.NewBlockingCompleted([]string{"BLOCK"})).s[1].ToString(); err != nil || val != "OK" {
			t.Fatalf("unexpect resp %v %v", err, val)
		}
	})
	t.Run("pipeline do - no ping", func(t *testing.T) {
		p, mock, cancel, _ := setup(t, ClientOption{ConnWriteTimeout: timeout, Dialer: net.Dialer{KeepAlive: timeout}})
		defer cancel()
		p.background()
		go func() {
			time.Sleep(3 * timeout)
			mock.Expect("BLOCK").ReplyString("OK")
		}()
		if val, err := p.Do(context.Background(), cmds.NewBlockingCompleted([]string{"BLOCK"})).ToString(); err != nil || val != "OK" {
			t.Fatalf("unexpect resp %v %v", err, val)
		}
	})
	t.Run("pipeline do - ignore ping timeout", func(t *testing.T) {
		p, mock, cancel, _ := setup(t, ClientOption{ConnWriteTimeout: timeout, Dialer: net.Dialer{KeepAlive: timeout}})
		defer cancel()
		p.background()
		wait := make(chan struct{})
		go func() {
			mock.Expect("PING")
			close(wait)
			time.Sleep(2 * timeout)
			mock.Expect("BLOCK").ReplyString("OK").ReplyString("OK")
		}()
		<-wait
		if val, err := p.Do(context.Background(), cmds.NewBlockingCompleted([]string{"BLOCK"})).ToString(); err != nil || val != "OK" {
			t.Fatalf("unexpect resp %v %v", err, val)
		}
	})
	t.Run("pipeline do multi - no ping", func(t *testing.T) {
		p, mock, cancel, _ := setup(t, ClientOption{ConnWriteTimeout: timeout, Dialer: net.Dialer{KeepAlive: timeout}})
		defer cancel()
		p.background()
		go func() {
			time.Sleep(3 * timeout)
			mock.Expect("READ").ReplyString("READ").
				Expect("BLOCK").ReplyString("OK")
		}()
		if val, err := p.DoMulti(context.Background(),
			cmds.NewReadOnlyCompleted([]string{"READ"}),
			cmds.NewBlockingCompleted([]string{"BLOCK"})).s[1].ToString(); err != nil || val != "OK" {
			t.Fatalf("unexpect resp %v %v", err, val)
		}
	})
	t.Run("pipeline do multi - ignore ping timeout", func(t *testing.T) {
		p, mock, cancel, _ := setup(t, ClientOption{ConnWriteTimeout: timeout, Dialer: net.Dialer{KeepAlive: timeout}})
		defer cancel()
		p.background()
		wait := make(chan struct{})
		go func() {
			mock.Expect("PING")
			close(wait)
			time.Sleep(2 * timeout)
			mock.Expect("READ").Expect("BLOCK").ReplyString("OK").ReplyString("READ").ReplyString("OK")
		}()
		<-wait
		if val, err := p.DoMulti(context.Background(),
			cmds.NewReadOnlyCompleted([]string{"READ"}),
			cmds.NewBlockingCompleted([]string{"BLOCK"})).s[1].ToString(); err != nil || val != "OK" {
			t.Fatalf("unexpect resp %v %v", err, val)
		}
	})
}

func TestDeadPipe(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	ctx := context.Background()
	if err := deadFn().Error(); err != ErrClosing {
		t.Fatalf("unexpected err %v", err)
	}
	if err := deadFn().Do(ctx, cmds.NewCompleted(nil)).Error(); err != ErrClosing {
		t.Fatalf("unexpected err %v", err)
	}
	if err := deadFn().DoMulti(ctx, cmds.NewCompleted(nil)).s[0].Error(); err != ErrClosing {
		t.Fatalf("unexpected err %v", err)
	}
	if err := deadFn().DoCache(ctx, Cacheable(cmds.NewCompleted(nil)), time.Second).Error(); err != ErrClosing {
		t.Fatalf("unexpected err %v", err)
	}
	if err := deadFn().Receive(ctx, cmds.NewCompleted(nil), func(message PubSubMessage) {}); err != ErrClosing {
		t.Fatalf("unexpected err %v", err)
	}
	if err := <-deadFn().SetPubSubHooks(PubSubHooks{OnMessage: func(m PubSubMessage) {}}); err != ErrClosing {
		t.Fatalf("unexpected err %v", err)
	}
}

func TestErrorPipe(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	ctx := context.Background()
	target := errors.New("any")
	if err := epipeFn(target).Error(); err != target {
		t.Fatalf("unexpected err %v", err)
	}
	if err := epipeFn(target).Do(ctx, cmds.NewCompleted(nil)).Error(); err != target {
		t.Fatalf("unexpected err %v", err)
	}
	if err := epipeFn(target).DoMulti(ctx, cmds.NewCompleted(nil)).s[0].Error(); err != target {
		t.Fatalf("unexpected err %v", err)
	}
	if err := epipeFn(target).DoCache(ctx, Cacheable(cmds.NewCompleted(nil)), time.Second).Error(); err != target {
		t.Fatalf("unexpected err %v", err)
	}
	if err := epipeFn(target).Receive(ctx, cmds.NewCompleted(nil), func(message PubSubMessage) {}); err != target {
		t.Fatalf("unexpected err %v", err)
	}
	if err := <-epipeFn(target).SetPubSubHooks(PubSubHooks{OnMessage: func(m PubSubMessage) {}}); err != target {
		t.Fatalf("unexpected err %v", err)
	}
}

func TestCloseHook(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	t.Run("normal close", func(t *testing.T) {
		var flag int32
		p, _, cancel, _ := setup(t, ClientOption{})
		p.SetOnCloseHook(func(error) {
			atomic.StoreInt32(&flag, 1)
		})
		cancel()
		if atomic.LoadInt32(&flag) != 1 {
			t.Fatalf("hook not be invoked")
		}
	})
	t.Run("disconnect", func(t *testing.T) {
		var flag int32
		p, _, _, closeConn := setup(t, ClientOption{})
		p.SetOnCloseHook(func(error) {
			atomic.StoreInt32(&flag, 1)
		})
		p.background()
		closeConn()
		for atomic.LoadInt32(&flag) != 1 {
			time.Sleep(time.Millisecond * 100)
			t.Log("wait close hook to be invoked")
		}
	})
}

func TestNoHelloRegex(t *testing.T) {
	defer ShouldNotLeaked(SetupLeakDetection())
	tests := []struct {
		name  string
		match bool
		resp  string
	}{
		{
			name:  "lowercase hello",
			match: true,
			resp:  "unknown command hello",
		},
		{
			name:  "uppercase hello",
			match: true,
			resp:  "unknown command HELLO",
		},
		{
			name:  "not hello",
			match: false,
			resp:  "unknown command not hello",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if match := noHello.MatchString(tt.resp); match != tt.match {
				t.Fatalf("unexpected match %v", match)
			}
		})
	}
}
