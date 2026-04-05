package mock

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"

	"github.com/redis/rueidis"
)

func Result(val rueidis.RedisMessage) rueidis.RedisResult {
	r := result{val: val}
	return *(*rueidis.RedisResult)(unsafe.Pointer(&r))
}

func ErrorResult(err error) rueidis.RedisResult {
	r := result{err: err}
	return *(*rueidis.RedisResult)(unsafe.Pointer(&r))
}

func RedisString(v string) rueidis.RedisMessage {
	m := strmsg('+', v)
	return *(*rueidis.RedisMessage)(unsafe.Pointer(&m))
}

func RedisBlobString(v string) rueidis.RedisMessage {
	m := strmsg('$', v)
	return *(*rueidis.RedisMessage)(unsafe.Pointer(&m))
}

func RedisError(v string) rueidis.RedisMessage {
	m := strmsg('-', v)
	return *(*rueidis.RedisMessage)(unsafe.Pointer(&m))
}

func RedisInt64(v int64) rueidis.RedisMessage {
	m := message{typ: ':', integer: v}
	return *(*rueidis.RedisMessage)(unsafe.Pointer(&m))
}

func RedisFloat64(v float64) rueidis.RedisMessage {
	m := strmsg(',', strconv.FormatFloat(v, 'f', -1, 64))
	return *(*rueidis.RedisMessage)(unsafe.Pointer(&m))
}

func RedisBool(v bool) rueidis.RedisMessage {
	m := message{typ: '#'}
	if v {
		m.integer = 1
	}
	return *(*rueidis.RedisMessage)(unsafe.Pointer(&m))
}

func RedisNil() rueidis.RedisMessage {
	m := message{typ: '_'}
	return *(*rueidis.RedisMessage)(unsafe.Pointer(&m))
}

func RedisArray(values ...rueidis.RedisMessage) rueidis.RedisMessage {
	m := slicemsg('*', values)
	return *(*rueidis.RedisMessage)(unsafe.Pointer(&m))
}

func RedisMap(kv map[string]rueidis.RedisMessage) rueidis.RedisMessage {
	values := make([]rueidis.RedisMessage, 0, 2*len(kv))
	for k, v := range kv {
		values = append(values, RedisString(k))
		values = append(values, v)
	}
	m := slicemsg('%', values)
	return *(*rueidis.RedisMessage)(unsafe.Pointer(&m))
}

func serialize(m message, buf *bytes.Buffer) {
	switch m.typ {
	case '$', '!', '=':
		buf.WriteString(fmt.Sprintf("%s%d\r\n%s\r\n", string(m.typ), len(m.string()), m.string()))
	case '+', '-', ',', '(':
		buf.WriteString(fmt.Sprintf("%s%s\r\n", string(m.typ), m.string()))
	case ':', '#':
		buf.WriteString(fmt.Sprintf("%s%d\r\n", string(m.typ), m.integer))
	case '_':
		buf.WriteString(fmt.Sprintf("%s\r\n", string(m.typ)))
	case '*':
		buf.WriteString(fmt.Sprintf("%s%d\r\n", string(m.typ), len(m.values())))
		for _, v := range m.values() {
			pv := *(*message)(unsafe.Pointer(&v))
			serialize(pv, buf)
		}
	case '%':
		buf.WriteString(fmt.Sprintf("%s%d\r\n", string(m.typ), len(m.values())/2))
		for _, v := range m.values() {
			pv := *(*message)(unsafe.Pointer(&v))
			serialize(pv, buf)
		}
	}
}

func RedisResultStreamError(err error) rueidis.RedisResultStream {
	s := stream{e: err}
	return *(*rueidis.RedisResultStream)(unsafe.Pointer(&s))
}

func RedisResultStream(ms ...rueidis.RedisMessage) rueidis.RedisResultStream {
	buf := bytes.NewBuffer(nil)
	for _, m := range ms {
		pm := *(*message)(unsafe.Pointer(&m))
		serialize(pm, buf)
	}
	s := stream{n: len(ms), p: &pool{size: 1, cond: sync.NewCond(&sync.Mutex{})}, w: &pipe{r: bufio.NewReader(buf)}}
	return *(*rueidis.RedisResultStream)(unsafe.Pointer(&s))
}

func MultiRedisResultStream(ms ...rueidis.RedisMessage) rueidis.MultiRedisResultStream {
	return RedisResultStream(ms...)
}

func MultiRedisResultStreamError(err error) rueidis.RedisResultStream {
	return RedisResultStreamError(err)
}

type message struct {
	attrs   *rueidis.RedisMessage
	bytes   *byte
	array   *rueidis.RedisMessage
	integer int64
	typ     byte
	ttl     [7]byte
}

func (m *message) string() string {
	if m.bytes == nil {
		return ""
	}
	return unsafe.String(m.bytes, m.integer)
}

func (m *message) values() []rueidis.RedisMessage {
	if m.array == nil {
		return nil
	}
	return unsafe.Slice(m.array, m.integer)
}

func slicemsg(typ byte, values []rueidis.RedisMessage) message {
	return message{
		typ:     typ,
		array:   unsafe.SliceData(values),
		integer: int64(len(values)),
	}
}

func strmsg(typ byte, value string) message {
	return message{
		typ:     typ,
		bytes:   unsafe.StringData(value),
		integer: int64(len(value)),
	}
}

type result struct {
	err error
	val rueidis.RedisMessage
}

type pool struct {
	dead    any
	cond    *sync.Cond
	timer   *time.Timer
	make    func() any
	list    []any
	cleanup time.Duration
	size    int
	minSize int
	cap     int
	down    bool
	timerOn bool
}

type pipe struct {
	conn            net.Conn
	clhks           atomic.Value // closed hook, invoked after the conn is closed
	queue           any
	cache           any
	pshks           atomic.Pointer[pshks] // pubsub hook, registered by the SetPubSubHooks
	error           atomic.Pointer[errs]
	r               *bufio.Reader
	w               *bufio.Writer
	close           chan struct{}
	onInvalidations func([]rueidis.RedisMessage)
	ssubs           *any // pubsub smessage subscriptions
	nsubs           *any // pubsub  message subscriptions
	psubs           *any // pubsub pmessage subscriptions
	r2p             *any
	pingTimer       *time.Timer // timer for background ping
	lftmTimer       *time.Timer // lifetime timer
	info            map[string]rueidis.RedisMessage
	timeout         time.Duration
	pinggap         time.Duration
	maxFlushDelay   time.Duration
	lftm            time.Duration // lifetime
	wrCounter       atomic.Uint64
	version         int32
	blcksig         int32
	state           int32
	bgState         int32
	r2ps            bool // identify this pipe is used for resp2 pubsub or not
	noNoDelay       bool
	optIn           bool
}

type stream struct {
	p *pool
	w *pipe
	e error
	n int
}

type errs struct{ error }

type pshks struct {
	hooks rueidis.PubSubHooks
	close chan error
}
