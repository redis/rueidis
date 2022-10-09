package mock

import (
	"strconv"
	"unsafe"

	"github.com/rueian/rueidis"
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
	m := message{typ: '+', string: v}
	return *(*rueidis.RedisMessage)(unsafe.Pointer(&m))
}

func RedisError(v string) rueidis.RedisMessage {
	m := message{typ: '-', string: v}
	return *(*rueidis.RedisMessage)(unsafe.Pointer(&m))
}

func RedisInt64(v int64) rueidis.RedisMessage {
	m := message{typ: ':', integer: v}
	return *(*rueidis.RedisMessage)(unsafe.Pointer(&m))
}

func RedisFloat64(v float64) rueidis.RedisMessage {
	m := message{typ: ',', string: strconv.FormatFloat(v, 'f', -1, 64)}
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
	m := message{typ: '*', values: values}
	return *(*rueidis.RedisMessage)(unsafe.Pointer(&m))
}

func RedisMap(kv map[string]rueidis.RedisMessage) rueidis.RedisMessage {
	values := make([]rueidis.RedisMessage, 0, 2*len(kv))
	for k, v := range kv {
		values = append(values, RedisString(k))
		values = append(values, v)
	}
	m := message{typ: '*', values: values}
	return *(*rueidis.RedisMessage)(unsafe.Pointer(&m))
}

type message struct {
	typ     byte
	string  string
	values  []rueidis.RedisMessage
	attrs   *rueidis.RedisMessage
	integer int64
}
type result struct {
	err error
	val rueidis.RedisMessage
}
