package rueidis

import (
	"bufio"
	"context"
	"errors"
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
				},
			})
		if !option.DisableCache {
			mock.Expect("CLIENT", "TRACKING", "ON", "OPTIN").
				ReplyString("OK")
		}
	}()
	p, err := newPipe(n1, &option)
	if err != nil {
		t.Fatalf("pipe setup failed: %v", err)
	}
	if info := p.Info(); info["version"].string != "6.0.0" {
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
	t.Run("Auth without Username", func(t *testing.T) {
		n1, n2 := net.Pipe()
		mock := &redisMock{buf: bufio.NewReader(n2), conn: n2}
		go func() {
			mock.Expect("HELLO", "3", "AUTH", "default", "pa", "SETNAME", "cn").
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
	t.Run("Auth with Username", func(t *testing.T) {
		n1, n2 := net.Pipe()
		mock := &redisMock{buf: bufio.NewReader(n2), conn: n2}
		go func() {
			mock.Expect("HELLO", "3", "AUTH", "ua", "pa", "SETNAME", "cn").
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
			Username:   "ua",
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

func TestNoReplyExceedRingSize(t *testing.T) {
	p, mock, cancel, _ := setup(t, ClientOption{})
	defer cancel()

	times := (2 << (DefaultRingScale - 1)) * 3
	wait := make(chan struct{})
	go func() {
		for i := 0; i < times; i++ {
			if err := p.Do(context.Background(), cmds.UnsubscribeCmd).Error(); err != nil {
				t.Fatalf("unexpected err %v", err)
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
	if v, entry := p.cache.GetOrPrepare("a", "GET", time.Second); v.typ != 0 || entry != nil {
		t.Errorf("unexpected cache value and entry %v %v", v, entry)
	}
}

func TestClientSideCachingExecAbortWithMoved(t *testing.T) {
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
			Reply(RedisMessage{typ: '-', string: "MOVED 0 :0"}).
			Reply(RedisMessage{typ: '-', string: "EXECABORT"})
	}()

	v, err := p.DoCache(context.Background(), cmds.Cacheable(cmds.NewCompleted([]string{"GET", "a"})), 10*time.Second).ToMessage()
	if addr, ok := err.(*RedisError).IsMoved(); !ok || addr != ":0" {
		t.Errorf("unexpected err, got %v", err)
	}
	if v.IsCacheHit() {
		t.Errorf("unexpected cache hit")
	}
	if v, entry := p.cache.GetOrPrepare("a", "GET", time.Second); v.typ != 0 || entry != nil {
		t.Errorf("unexpected cache value and entry %v %v", v, entry)
	}
}

func TestClientSideCachingWithNonRedisError(t *testing.T) {
	p, _, _, closeConn := setup(t, ClientOption{})
	closeConn()

	v, err := p.DoCache(context.Background(), cmds.Cacheable(cmds.NewCompleted([]string{"GET", "a"})), 10*time.Second).ToMessage()
	if err != io.EOF && !strings.HasPrefix(err.Error(), "io:") {
		t.Errorf("unexpected err, got %v", err)
	}
	if v.IsCacheHit() {
		t.Errorf("unexpected cache hit")
	}
	if v, entry := p.cache.GetOrPrepare("a", "GET", time.Second); v.typ != 0 || entry != nil {
		t.Errorf("unexpected cache value and entry %v %v", v, entry)
	}
}

func TestClientSideCachingMGet(t *testing.T) {
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
		v, _ := p.DoCache(context.Background(), cmds.Cacheable(cmds.NewMGetCompleted([]string{"MGET", "a1", "a2", "a3"})), 10*time.Second).ToMessage()
		arr, _ := v.ToArray()
		if len(arr) != 3 {
			t.Errorf("unexpected cached mget length, expected 3, got %v", len(arr))
		}
		for i, v := range arr {
			if v.integer != int64(i+1) {
				t.Errorf("unexpected cached mget response, expected %v, got %v", i+1, v.integer)
			}
		}
		if ttl := p.cache.GetTTL("a1"); !roughly(ttl, time.Second) {
			t.Errorf("unexpected ttl %v", ttl)
		}
		if ttl := p.cache.GetTTL("a2"); !roughly(ttl, time.Second*2) {
			t.Errorf("unexpected ttl %v", ttl)
		}
		if ttl := p.cache.GetTTL("a3"); !roughly(ttl, time.Second*3) {
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
		if p.cache.GetTTL("a1") == -2 && p.cache.GetTTL("a3") == -2 {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}

	v, _ := p.DoCache(context.Background(), cmds.Cacheable(cmds.NewMGetCompleted([]string{"MGET", "a1", "a2", "a3"})), 10*time.Second).ToMessage()
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
	if ttl := p.cache.GetTTL("a1"); !roughly(ttl, time.Second*10) {
		t.Errorf("unexpected ttl %v", ttl)
	}
	if ttl := p.cache.GetTTL("a2"); !roughly(ttl, time.Second*2) {
		t.Errorf("unexpected ttl %v", ttl)
	}
	if ttl := p.cache.GetTTL("a3"); !roughly(ttl, time.Second*30) {
		t.Errorf("unexpected ttl %v", ttl)
	}
}

func TestClientSideCachingJSONMGet(t *testing.T) {
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
		v, _ := p.DoCache(context.Background(), cmds.Cacheable(cmds.NewMGetCompleted([]string{"JSON.MGET", "a1", "a2", "a3", "$"})), 10*time.Second).ToMessage()
		arr, _ := v.ToArray()
		if len(arr) != 3 {
			t.Errorf("unexpected cached mget length, expected 3, got %v", len(arr))
		}
		for i, v := range arr {
			if v.integer != int64(i+1) {
				t.Errorf("unexpected cached mget response, expected %v, got %v", i+1, v.integer)
			}
		}
		if ttl := p.cache.GetTTL("a1"); !roughly(ttl, time.Second) {
			t.Errorf("unexpected ttl %v", ttl)
		}
		if ttl := p.cache.GetTTL("a2"); !roughly(ttl, time.Second*2) {
			t.Errorf("unexpected ttl %v", ttl)
		}
		if ttl := p.cache.GetTTL("a3"); !roughly(ttl, time.Second*3) {
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
		if p.cache.GetTTL("a1") == -2 && p.cache.GetTTL("a3") == -2 {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}

	v, _ := p.DoCache(context.Background(), cmds.Cacheable(cmds.NewMGetCompleted([]string{"JSON.MGET", "a1", "a2", "a3", "$"})), 10*time.Second).ToMessage()
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
	if ttl := p.cache.GetTTL("a1"); !roughly(ttl, time.Second*10) {
		t.Errorf("unexpected ttl %v", ttl)
	}
	if ttl := p.cache.GetTTL("a2"); !roughly(ttl, time.Second*2) {
		t.Errorf("unexpected ttl %v", ttl)
	}
	if ttl := p.cache.GetTTL("a3"); !roughly(ttl, time.Second*30) {
		t.Errorf("unexpected ttl %v", ttl)
	}
}

func TestClientSideCachingExecAbortMGet(t *testing.T) {
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

	v, err := p.DoCache(context.Background(), cmds.Cacheable(cmds.NewMGetCompleted([]string{"MGET", "a1", "a2"})), 10*time.Second).ToMessage()
	if !IsRedisNil(err) {
		t.Errorf("unexpected err, got %v", err)
	}
	if v.IsCacheHit() {
		t.Errorf("unexpected cache hit")
	}
	if v, entry := p.cache.GetOrPrepare("a1", "GET", time.Second); v.typ != 0 || entry != nil {
		t.Errorf("unexpected cache value and entry %v %v", v, entry)
	}
	if v, entry := p.cache.GetOrPrepare("a2", "GET", time.Second); v.typ != 0 || entry != nil {
		t.Errorf("unexpected cache value and entry %v %v", v, entry)
	}
}

func TestClientSideCachingExecAbortWithMovedMGet(t *testing.T) {
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
			Reply(RedisMessage{typ: '-', string: "MOVED 0 :0"}).
			Reply(RedisMessage{typ: '-', string: "EXECABORT"})
	}()

	v, err := p.DoCache(context.Background(), cmds.Cacheable(cmds.NewMGetCompleted([]string{"MGET", "a1", "a2"})), 10*time.Second).ToMessage()
	if addr, ok := err.(*RedisError).IsMoved(); !ok || addr != ":0" {
		t.Errorf("unexpected err, got %v", err)
	}
	if v.IsCacheHit() {
		t.Errorf("unexpected cache hit")
	}
	if v, entry := p.cache.GetOrPrepare("a1", "GET", time.Second); v.typ != 0 || entry != nil {
		t.Errorf("unexpected cache value and entry %v %v", v, entry)
	}
	if v, entry := p.cache.GetOrPrepare("a2", "GET", time.Second); v.typ != 0 || entry != nil {
		t.Errorf("unexpected cache value and entry %v %v", v, entry)
	}
}

func TestClientSideCachingWithNonRedisErrorMGet(t *testing.T) {
	p, _, _, closeConn := setup(t, ClientOption{})
	closeConn()

	v, err := p.DoCache(context.Background(), cmds.Cacheable(cmds.NewMGetCompleted([]string{"MGET", "a1", "a2"})), 10*time.Second).ToMessage()
	if err != io.EOF && !strings.HasPrefix(err.Error(), "io:") {
		t.Errorf("unexpected err, got %v", err)
	}
	if v.IsCacheHit() {
		t.Errorf("unexpected cache hit")
	}
	if v, entry := p.cache.GetOrPrepare("a1", "GET", time.Second); v.typ != 0 || entry != nil {
		t.Errorf("unexpected cache value and entry %v %v", v, entry)
	}
	if v, entry := p.cache.GetOrPrepare("a2", "GET", time.Second); v.typ != 0 || entry != nil {
		t.Errorf("unexpected cache value and entry %v %v", v, entry)
	}
}

func TestClientSideCachingWithSideChannelMGet(t *testing.T) {
	p, _, _, closeConn := setup(t, ClientOption{})
	closeConn()

	p.cache.GetOrPrepare("a1", "GET", 10*time.Second)
	go func() {
		time.Sleep(100 * time.Millisecond)
		p.cache.Update("a1", "GET", RedisMessage{typ: '+', string: "OK"}, 10)
	}()

	v, _ := p.DoCache(context.Background(), cmds.Cacheable(cmds.NewMGetCompleted([]string{"MGET", "a1"})), 10*time.Second).AsStrSlice()
	if v[0] != "OK" {
		t.Errorf("unexpected value, got %v", v)
	}
}

func TestClientSideCachingWithSideChannelErrorMGet(t *testing.T) {
	p, _, _, closeConn := setup(t, ClientOption{})
	closeConn()

	p.cache.GetOrPrepare("a1", "GET", 10*time.Second)
	go func() {
		time.Sleep(100 * time.Millisecond)
		p.cache.Cancel("a1", "GET", RedisMessage{}, io.EOF)
	}()

	_, err := p.DoCache(context.Background(), cmds.Cacheable(cmds.NewMGetCompleted([]string{"MGET", "a1"})), 10*time.Second).ToMessage()
	if err != io.EOF {
		t.Errorf("unexpected err, got %v", err)
	}
}

func TestClientSideCachingDoMultiCacheMGet(t *testing.T) {
	p, _, cancel, _ := setup(t, ClientOption{})
	defer cancel()

	defer func() {
		if !strings.Contains(recover().(string), panicmgetcsc) {
			t.Fatal("should panic")
		}
	}()
	p.DoMultiCache(context.Background(), []CacheableTTL{
		CT(cmds.Cacheable(cmds.NewMGetCompleted([]string{"MGET", "a1"})), time.Second*10),
	}...)
}

func TestClientSideCachingDoMultiCache(t *testing.T) {
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
			Expect("GET", "a1").
			Expect("PTTL", "a2").
			Expect("GET", "a2").
			Expect("PTTL", "a3").
			Expect("GET", "a3").
			Expect("EXEC").
			ReplyString("OK").
			ReplyString("OK").
			ReplyString("OK").
			ReplyString("OK").
			ReplyString("OK").
			ReplyString("OK").
			ReplyString("OK").
			ReplyString("OK").
			Reply(RedisMessage{typ: '*', values: []RedisMessage{
				{typ: ':', integer: 1000},
				{typ: ':', integer: 1},
				{typ: ':', integer: 2000},
				{typ: ':', integer: 2},
				{typ: ':', integer: 3000},
				{typ: ':', integer: 3},
			}})
	}()
	// single flight
	miss := uint64(0)
	hits := uint64(0)
	for i := 0; i < 2; i++ {
		arr := p.DoMultiCache(context.Background(), []CacheableTTL{
			CT(cmds.Cacheable(cmds.NewCompleted([]string{"GET", "a1"})), time.Second*10),
			CT(cmds.Cacheable(cmds.NewCompleted([]string{"GET", "a2"})), time.Second*10),
			CT(cmds.Cacheable(cmds.NewCompleted([]string{"GET", "a3"})), time.Second*10),
		}...)
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
		if ttl := p.cache.GetTTL("a1"); !roughly(ttl, time.Second) {
			t.Errorf("unexpected ttl %v", ttl)
		}
		if ttl := p.cache.GetTTL("a2"); !roughly(ttl, time.Second*2) {
			t.Errorf("unexpected ttl %v", ttl)
		}
		if ttl := p.cache.GetTTL("a3"); !roughly(ttl, time.Second*3) {
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
			Expect("PTTL", "a3").
			Expect("GET", "a3").
			Expect("EXEC").
			ReplyString("OK").
			ReplyString("OK").
			ReplyString("OK").
			ReplyString("OK").
			ReplyString("OK").
			ReplyString("OK").
			Reply(RedisMessage{typ: '*', values: []RedisMessage{
				{typ: ':', integer: 10000},
				{typ: ':', integer: 10},
				{typ: ':', integer: 30000},
				{typ: ':', integer: 30},
			}})
	}()

	for {
		if p.cache.GetTTL("a1") == -2 && p.cache.GetTTL("a3") == -2 {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}

	arr := p.DoMultiCache(context.Background(), []CacheableTTL{
		CT(cmds.Cacheable(cmds.NewCompleted([]string{"GET", "a1"})), time.Second*10),
		CT(cmds.Cacheable(cmds.NewCompleted([]string{"GET", "a2"})), time.Second*10),
		CT(cmds.Cacheable(cmds.NewCompleted([]string{"GET", "a3"})), time.Second*10),
	}...)
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
	if ttl := p.cache.GetTTL("a1"); !roughly(ttl, time.Second*10) {
		t.Errorf("unexpected ttl %v", ttl)
	}
	if ttl := p.cache.GetTTL("a2"); !roughly(ttl, time.Second*2) {
		t.Errorf("unexpected ttl %v", ttl)
	}
	if ttl := p.cache.GetTTL("a3"); !roughly(ttl, time.Second*30) {
		t.Errorf("unexpected ttl %v", ttl)
	}
}

func TestClientSideCachingExecAbortDoMultiCache(t *testing.T) {
	p, mock, cancel, _ := setup(t, ClientOption{})
	defer cancel()

	go func() {
		mock.Expect("CLIENT", "CACHING", "YES").
			Expect("MULTI").
			Expect("PTTL", "a1").
			Expect("GET", "a1").
			Expect("PTTL", "a2").
			Expect("GET", "a2").
			Expect("EXEC").
			ReplyString("OK").
			ReplyString("OK").
			ReplyString("OK").
			ReplyString("OK").
			ReplyString("OK").
			ReplyString("OK").
			Reply(RedisMessage{typ: '_'})
	}()

	arr := p.DoMultiCache(context.Background(), []CacheableTTL{
		CT(cmds.Cacheable(cmds.NewCompleted([]string{"GET", "a1"})), time.Second*10),
		CT(cmds.Cacheable(cmds.NewCompleted([]string{"GET", "a2"})), time.Second*10),
	}...)
	for _, resp := range arr {
		v, err := resp.ToMessage()
		if !IsRedisNil(err) {
			t.Errorf("unexpected err, got %v", err)
		}
		if v.IsCacheHit() {
			t.Errorf("unexpected cache hit")
		}
	}
	if v, entry := p.cache.GetOrPrepare("a1", "GET", time.Second); v.typ != 0 || entry != nil {
		t.Errorf("unexpected cache value and entry %v %v", v, entry)
	}
	if v, entry := p.cache.GetOrPrepare("a2", "GET", time.Second); v.typ != 0 || entry != nil {
		t.Errorf("unexpected cache value and entry %v %v", v, entry)
	}
}

func TestClientSideCachingExecAbortWithMovedDoMultiCache(t *testing.T) {
	p, mock, cancel, _ := setup(t, ClientOption{})
	defer cancel()

	go func() {
		mock.Expect("CLIENT", "CACHING", "YES").
			Expect("MULTI").
			Expect("PTTL", "a1").
			Expect("GET", "a1").
			Expect("PTTL", "a2").
			Expect("GET", "a2").
			Expect("EXEC").
			ReplyString("OK").
			ReplyString("OK").
			ReplyString("OK").
			ReplyString("OK").
			ReplyString("OK").
			Reply(RedisMessage{typ: '-', string: "MOVED 0 :0"}).
			Reply(RedisMessage{typ: '-', string: "EXECABORT"})
	}()

	arr := p.DoMultiCache(context.Background(), []CacheableTTL{
		CT(cmds.Cacheable(cmds.NewCompleted([]string{"GET", "a1"})), time.Second*10),
		CT(cmds.Cacheable(cmds.NewCompleted([]string{"GET", "a2"})), time.Second*10),
	}...)
	for _, resp := range arr {
		v, err := resp.ToMessage()
		if addr, ok := err.(*RedisError).IsMoved(); !ok || addr != ":0" {
			t.Errorf("unexpected err, got %v", err)
		}
		if v.IsCacheHit() {
			t.Errorf("unexpected cache hit")
		}
	}
	if v, entry := p.cache.GetOrPrepare("a1", "GET", time.Second); v.typ != 0 || entry != nil {
		t.Errorf("unexpected cache value and entry %v %v", v, entry)
	}
	if v, entry := p.cache.GetOrPrepare("a2", "GET", time.Second); v.typ != 0 || entry != nil {
		t.Errorf("unexpected cache value and entry %v %v", v, entry)
	}
}

func TestClientSideCachingWithNonRedisErrorDoMultiCache(t *testing.T) {
	p, _, _, closeConn := setup(t, ClientOption{})
	closeConn()

	arr := p.DoMultiCache(context.Background(), []CacheableTTL{
		CT(cmds.Cacheable(cmds.NewCompleted([]string{"GET", "a1"})), time.Second*10),
		CT(cmds.Cacheable(cmds.NewCompleted([]string{"GET", "a2"})), time.Second*10),
	}...)
	for _, resp := range arr {
		v, err := resp.ToMessage()
		if err != io.EOF && !strings.HasPrefix(err.Error(), "io:") {
			t.Errorf("unexpected err, got %v", err)
		}
		if v.IsCacheHit() {
			t.Errorf("unexpected cache hit")
		}
	}
	if v, entry := p.cache.GetOrPrepare("a1", "GET", time.Second); v.typ != 0 || entry != nil {
		t.Errorf("unexpected cache value and entry %v %v", v, entry)
	}
	if v, entry := p.cache.GetOrPrepare("a2", "GET", time.Second); v.typ != 0 || entry != nil {
		t.Errorf("unexpected cache value and entry %v %v", v, entry)
	}
}

func TestClientSideCachingWithSideChannelDoMultiCache(t *testing.T) {
	p, _, _, closeConn := setup(t, ClientOption{})
	closeConn()

	p.cache.GetOrPrepare("a1", "GET", 10*time.Second)
	go func() {
		time.Sleep(100 * time.Millisecond)
		p.cache.Update("a1", "GET", RedisMessage{typ: '+', string: "OK"}, 10)
	}()

	arr := p.DoMultiCache(context.Background(), []CacheableTTL{
		CT(cmds.Cacheable(cmds.NewCompleted([]string{"GET", "a1"})), time.Second*10),
	}...)
	if arr[0].val.string != "OK" {
		t.Errorf("unexpected value, got %v", arr[0].val.string)
	}
}

func TestClientSideCachingWithSideChannelErrorDoMultiCache(t *testing.T) {
	p, _, _, closeConn := setup(t, ClientOption{})
	closeConn()

	p.cache.GetOrPrepare("a1", "GET", 10*time.Second)
	go func() {
		time.Sleep(100 * time.Millisecond)
		p.cache.Cancel("a1", "GET", RedisMessage{}, io.EOF)
	}()

	arr := p.DoMultiCache(context.Background(), []CacheableTTL{
		CT(cmds.Cacheable(cmds.NewCompleted([]string{"GET", "a1"})), time.Second*10),
	}...)
	if arr[0].err != io.EOF {
		t.Errorf("unexpected err, got %v", arr[0].err)
	}
}

// https://github.com/redis/redis/issues/8935
func TestClientSideCachingRedis6InvalidationBug1(t *testing.T) {
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
			v, _ := p.DoCache(context.Background(), cmds.Cacheable(cmds.NewCompleted([]string{"GET", "a"})), 10*time.Second).ToMessage()
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
			v, _ := p.DoCache(context.Background(), cmds.Cacheable(cmds.NewCompleted([]string{"GET", "a"})), 10*time.Second).ToMessage()
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

	err := p.DoCache(context.Background(), cmds.Cacheable(cmds.NewCompleted([]string{"GET", "a"})), 10*time.Second).Error()
	if err != io.EOF && !strings.HasPrefix(err.Error(), "io:") {
		t.Errorf("unexpected err %v", err)
	}
}

func TestDisableClientSideCaching(t *testing.T) {
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

	v, _ := p.DoCache(context.Background(), cmds.Cacheable(cmds.NewCompleted([]string{"GET", "a"})), 10*time.Second).ToMessage()
	if v.string != "1" {
		t.Errorf("unexpected cached result, expected %v, got %v", "1", v.string)
	}

	vs := p.DoMultiCache(context.Background(),
		CT(cmds.Cacheable(cmds.NewCompleted([]string{"GET", "b"})), 10*time.Second),
		CT(cmds.Cacheable(cmds.NewCompleted([]string{"GET", "c"})), 10*time.Second))
	if vs[0].val.string != "2" {
		t.Errorf("unexpected cached result, expected %v, got %v", "1", v.string)
	}
	if vs[1].val.string != "3" {
		t.Errorf("unexpected cached result, expected %v, got %v", "1", v.string)
	}
}

func TestMultiHalfErr(t *testing.T) {
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

	err := p.DoCache(context.Background(), cmds.Cacheable(cmds.NewCompleted([]string{"GET", "a"})), 10*time.Second).Error()
	if err != io.EOF && !strings.HasPrefix(err.Error(), "io:") {
		t.Errorf("unexpected err %v", err)
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

		commands := []cmds.Completed{
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

		commands := []cmds.Completed{
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

		commands := []cmds.Completed{
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

		p.psubs.Confirm("1")
		p.ssubs.Confirm("2")
		p.ssubs.Confirm("3")

		commands := []cmds.Completed{
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

	t.Run("PubSub Proactive SSUNSCRIBE", func(t *testing.T) {
		ctx := context.Background()
		p, mock, cancel, _ := setup(t, ClientOption{})

		p.ssubs.Confirm("0")
		p.ssubs.Confirm("1")
		p.ssubs.Confirm("2")

		commands := []cmds.Completed{
			builder.Sunsubscribe().Build(),
			builder.Ssubscribe().Channel("3").Build(),
		}

		replies := [][]RedisMessage{
			{
				{ // proactive sunsubscribe before user sunsubscribe
					typ: '>',
					values: []RedisMessage{
						{typ: '+', string: "sunsubscribe"},
						{typ: '+', string: "1"},
						{typ: ':', integer: 0},
					},
				},
				{ // proactive sunsubscribe before user sunsubscribe
					typ: '>',
					values: []RedisMessage{
						{typ: '+', string: "sunsubscribe"},
						{typ: '+', string: "2"},
						{typ: ':', integer: 0},
					},
				},
				{ // user sunsubscribe
					typ: '>',
					values: []RedisMessage{
						{typ: '+', string: "sunsubscribe"},
						{typ: '_'},
						{typ: ':', integer: 0},
					},
				},
				{ // proactive sunsubscribe after user sunsubscribe
					typ: '>',
					values: []RedisMessage{
						{typ: '+', string: "sunsubscribe"},
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
				{ // proactive sunsubscribe after user ssubscribe
					typ: '>',
					values: []RedisMessage{
						{typ: '+', string: "sunsubscribe"},
						{typ: '+', string: "3"},
						{typ: ':', integer: 0},
					},
				},
			},
		}

		p.background()

		// proactive sunsubscribe before other commands
		mock.Expect().Reply(RedisMessage{ // proactive sunsubscribe before user sunsubscribe
			typ: '>',
			values: []RedisMessage{
				{typ: '+', string: "sunsubscribe"},
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

	t.Run("PubSub Unexpected Subscribe/Unsubscribe", func(t *testing.T) {
		var shouldPanic = func(push string) (pass bool) {
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
			"unsubscribe",
			"psubscribe",
			"punsubscribe",
			"ssubscribe",
			// no unexpected sunsubscribe, because redis will send sunsubscribe proactively when slot changes
		} {
			if !shouldPanic(push) {
				t.Fatalf("should panic on protocolbug")
			}
		}
	})
}

//gocyclo:ignore
func TestPubSubHooks(t *testing.T) {
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

		for _, r := range p.DoMulti(ctx, activate1, activate2) {
			if err := r.Error(); err != nil {
				t.Fatalf("unexpected err %v", err)
			}
		}
		for _, r := range p.DoMulti(ctx, deactivate1, deactivate2) {
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

		for _, r := range p.DoMulti(ctx, activate1, activate2) {
			if err := r.Error(); err != nil {
				t.Fatalf("unexpected err %v", err)
			}
		}
		for _, r := range p.DoMulti(ctx, deactivate1, deactivate2) {
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
	mock.Close()
	wg.Wait()
}

func TestAlreadyCanceledContext(t *testing.T) {
	p, _, close, closeConn := setup(t, ClientOption{})
	defer closeConn()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	if err := p.Do(ctx, cmds.NewCompleted([]string{"GET", "a"})).NonRedisError(); !errors.Is(err, context.Canceled) {
		t.Fatalf("unexpected err %v", err)
	}
	if err := p.DoMulti(ctx, cmds.NewCompleted([]string{"GET", "a"}))[0].NonRedisError(); !errors.Is(err, context.Canceled) {
		t.Fatalf("unexpected err %v", err)
	}

	ctx, cancel = context.WithDeadline(context.Background(), time.Now().Add(-1*time.Second))
	cancel()

	if err := p.Do(ctx, cmds.NewCompleted([]string{"GET", "a"})).NonRedisError(); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("unexpected err %v", err)
	}
	if err := p.DoMulti(ctx, cmds.NewCompleted([]string{"GET", "a"}))[0].NonRedisError(); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("unexpected err %v", err)
	}
	close()
}

func TestCancelContext_Do(t *testing.T) {
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

func TestCancelContext_Do_Block(t *testing.T) {
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
	p, mock, shutdown, _ := setup(t, ClientOption{})

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		mock.Expect("GET", "a")
		cancel()
		mock.Expect().ReplyString("OK")
	}()

	if err := p.DoMulti(ctx, cmds.NewCompleted([]string{"GET", "a"}))[0].NonRedisError(); !errors.Is(err, context.Canceled) {
		t.Fatalf("unexpected err %v", err)
	}
	shutdown()
}

func TestCancelContext_DoMulti_Block(t *testing.T) {
	p, mock, shutdown, _ := setup(t, ClientOption{})

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		mock.Expect("GET", "a")
		cancel()
		mock.Expect().ReplyString("OK")
	}()

	if err := p.DoMulti(ctx, cmds.NewBlockingCompleted([]string{"GET", "a"}))[0].NonRedisError(); !errors.Is(err, context.Canceled) {
		t.Fatalf("unexpected err %v", err)
	}
	shutdown()
}

func TestOngoingDeadlineContextInSyncMode_Do(t *testing.T) {
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
	p, _, _, closeConn := setup(t, ClientOption{ConnWriteTimeout: 1 * time.Second / 2, Dialer: net.Dialer{KeepAlive: time.Second / 3}})
	defer closeConn()

	if err := p.Do(context.Background(), cmds.NewCompleted([]string{"GET", "a"})).NonRedisError(); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("unexpected err %v", err)
	}
	p.Close()
}

func TestOngoingDeadlineContextInSyncMode_DoMulti(t *testing.T) {
	p, _, _, closeConn := setup(t, ClientOption{})
	defer closeConn()

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(1*time.Second/2))
	defer cancel()

	if err := p.DoMulti(ctx, cmds.NewCompleted([]string{"GET", "a"}))[0].NonRedisError(); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("unexpected err %v", err)
	}
	p.Close()
}

func TestWriteDeadlineInSyncMode_DoMulti(t *testing.T) {
	p, _, _, closeConn := setup(t, ClientOption{ConnWriteTimeout: time.Second / 2, Dialer: net.Dialer{KeepAlive: time.Second / 3}})
	defer closeConn()

	if err := p.DoMulti(context.Background(), cmds.NewCompleted([]string{"GET", "a"}))[0].NonRedisError(); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("unexpected err %v", err)
	}
	p.Close()
}

func TestOngoingCancelContextInPipelineMode_Do(t *testing.T) {
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
	p, mock, close, closeConn := setup(t, ClientOption{})
	defer closeConn()

	p.background()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	canceled := int32(0)

	for i := 0; i < 5; i++ {
		go func() {
			_, err := p.DoMulti(ctx, cmds.NewCompleted([]string{"GET", "a"}))[0].ToString()
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
	p, _, _, closeConn := setup(t, ClientOption{ConnWriteTimeout: time.Second / 2, Dialer: net.Dialer{KeepAlive: time.Second / 3}})
	defer closeConn()

	p.background()

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	timeout := int32(0)

	for i := 0; i < 5; i++ {
		go func() {
			_, err := p.DoMulti(ctx, cmds.NewCompleted([]string{"GET", "a"}))[0].ToString()
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
	p, mock, cancel, _ := setup(t, ClientOption{ConnWriteTimeout: time.Second / 2, Dialer: net.Dialer{KeepAlive: time.Second / 3}})
	defer cancel()
	p.background()
	go func() {
		p.CleanSubscriptions()
	}()
	mock.Expect("UNSUBSCRIBE").Expect("PUNSUBSCRIBE").Reply(
		RedisMessage{typ: '>', values: []RedisMessage{
			{typ: '+', string: "unsubscribe"},
			{typ: '_'},
			{typ: ':', integer: 1},
		}},
		RedisMessage{typ: '>', values: []RedisMessage{
			{typ: '+', string: "punsubscribe"},
			{typ: '_'},
			{typ: ':', integer: 2},
		}})
}

func TestPipe_CleanSubscriptions_7(t *testing.T) {
	p, mock, cancel, _ := setup(t, ClientOption{ConnWriteTimeout: time.Second / 2, Dialer: net.Dialer{KeepAlive: time.Second / 3}})
	p.version = 7
	defer cancel()
	p.background()
	go func() {
		p.CleanSubscriptions()
	}()
	mock.Expect("UNSUBSCRIBE").Expect("PUNSUBSCRIBE").Expect("SUNSUBSCRIBE").Reply(
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
			{typ: ':', integer: 2},
		}})
}

func TestPingOnConnError(t *testing.T) {
	p, mock, _, closeConn := setup(t, ClientOption{ConnWriteTimeout: 3 * time.Second, Dialer: net.Dialer{KeepAlive: time.Second / 3}})
	p.background()
	mock.Expect("PING")
	closeConn()
	time.Sleep(time.Second / 2)
	p.Close()
	if err := p.Error(); err != io.EOF && !strings.HasPrefix(err.Error(), "io:") {
		t.Fatalf("unexpect err %v", err)
	}
}

//gocyclo:ignore
func TestBlockingCommandNoDeadline(t *testing.T) {
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
			cmds.NewBlockingCompleted([]string{"BLOCK"}))[1].ToString(); err != nil || val != "OK" {
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
			cmds.NewBlockingCompleted([]string{"BLOCK"}))[1].ToString(); err != nil || val != "OK" {
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
			cmds.NewBlockingCompleted([]string{"BLOCK"}))[1].ToString(); err != nil || val != "OK" {
			t.Fatalf("unexpect resp %v %v", err, val)
		}
	})
}

func TestDeadPipe(t *testing.T) {
	ctx := context.Background()
	if err := deadFn().Error(); err != ErrClosing {
		t.Fatalf("unexpected err %v", err)
	}
	if err := deadFn().Do(ctx, cmds.NewCompleted(nil)).Error(); err != ErrClosing {
		t.Fatalf("unexpected err %v", err)
	}
	if err := deadFn().DoMulti(ctx, cmds.NewCompleted(nil))[0].Error(); err != ErrClosing {
		t.Fatalf("unexpected err %v", err)
	}
	if err := deadFn().DoCache(ctx, cmds.Cacheable(cmds.NewCompleted(nil)), time.Second).Error(); err != ErrClosing {
		t.Fatalf("unexpected err %v", err)
	}
	if err := deadFn().Receive(ctx, cmds.NewCompleted(nil), func(message PubSubMessage) {}); err != ErrClosing {
		t.Fatalf("unexpected err %v", err)
	}
	if err := <-deadFn().SetPubSubHooks(PubSubHooks{OnMessage: func(m PubSubMessage) {}}); err != ErrClosing {
		t.Fatalf("unexpected err %v", err)
	}
}
