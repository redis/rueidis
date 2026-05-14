package rueidiscompatmock

import (
	"context"
	"encoding"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/redis/rueidis"
	"github.com/redis/rueidis/mock"
	"go.uber.org/mock/gomock"
)

type ClientMock struct {
	raw *mock.Client

	mu        sync.Mutex
	queue     []*expectation
	unmatched []error
	ordered   bool
}

type expectation struct {
	matcher gomock.Matcher
	result  rueidis.RedisResult
}

func NewAdapter(m *mock.Client) *ClientMock {
	cm := &ClientMock{raw: m, ordered: true}
	cm.wire()
	return cm
}

func (m *ClientMock) MatchExpectationsInOrder(b bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.ordered = b
}

func (m *ClientMock) ClearExpect() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.queue = nil
	m.unmatched = nil
}

func (m *ClientMock) ExpectationsWereMet() error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if len(m.unmatched) > 0 {
		return m.unmatched[0]
	}
	if len(m.queue) > 0 {
		cmds := make([]string, 0, len(m.queue))
		for _, e := range m.queue {
			cmds = append(cmds, e.matcher.String())
		}
		return fmt.Errorf("rueidiscompatmock: there are remaining expectations: %v", cmds)
	}
	return nil
}

func (m *ClientMock) wire() {
	m.raw.EXPECT().
		Do(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, cmd rueidis.Completed) rueidis.RedisResult {
			return m.consume(1, []rueidis.Completed{cmd})[0]
		}).
		AnyTimes()
	m.raw.EXPECT().
		DoMulti(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, cmds ...rueidis.Completed) []rueidis.RedisResult {
			return m.consume(len(cmds), cmds)
		}).
		AnyTimes()
}

func (m *ClientMock) consume(n int, cmds []rueidis.Completed) []rueidis.RedisResult {
	m.mu.Lock()
	defer m.mu.Unlock()
	out := make([]rueidis.RedisResult, n)
	for i, cmd := range cmds {
		idx, ok := m.matchLocked(cmd)
		if !ok {
			err := fmt.Errorf("rueidiscompatmock: no expectation for command %v", cmd.Commands())
			m.unmatched = append(m.unmatched, err)
			out[i] = mock.ErrorResult(err)
			continue
		}
		out[i] = m.queue[idx].result
		m.queue = append(m.queue[:idx], m.queue[idx+1:]...)
	}
	return out
}

func (m *ClientMock) matchLocked(cmd rueidis.Completed) (int, bool) {
	if m.ordered {
		if len(m.queue) == 0 {
			return 0, false
		}
		if !m.queue[0].matcher.Matches(cmd) {
			return 0, false
		}
		return 0, true
	}
	for i, e := range m.queue {
		if e.matcher.Matches(cmd) {
			return i, true
		}
	}
	return 0, false
}

func (m *ClientMock) push(matcher gomock.Matcher, defaultResult rueidis.RedisResult) *expectation {
	m.mu.Lock()
	defer m.mu.Unlock()
	e := &expectation{matcher: matcher, result: defaultResult}
	m.queue = append(m.queue, e)
	return e
}

const keepTTL = -1

func usePrecise(d time.Duration) bool {
	return d < time.Second || d%time.Second != 0
}

func formatMs(d time.Duration) int64 { return int64(d / time.Millisecond) }

func formatSec(d time.Duration) int64 { return int64(d / time.Second) }

func str(arg any) string {
	if arg == nil {
		return ""
	}
	switch v := arg.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	case bool:
		if v {
			return "1"
		}
		return "0"
	case time.Time:
		return v.Format(time.RFC3339Nano)
	case time.Duration:
		return strconv.FormatInt(v.Nanoseconds(), 10)
	case encoding.BinaryMarshaler:
		if data, err := v.MarshalBinary(); err == nil {
			return rueidis.BinaryString(data)
		}
	}
	return fmt.Sprint(arg)
}

func (m *ClientMock) ExpectGet(key string) *ExpectedString {
	e := m.push(mock.Match("GET", key), mock.Result(mock.RedisNil()))
	return &ExpectedString{exp: e}
}

func (m *ClientMock) ExpectSet(key string, value any, expiration time.Duration) *ExpectedStatus {
	e := m.push(setMatcher(key, value, expiration), mock.Result(mock.RedisString("OK")))
	return &ExpectedStatus{exp: e}
}

func (m *ClientMock) ExpectSetNX(key string, value any, expiration time.Duration) *ExpectedBool {
	e := m.push(setNXMatcher(key, value, expiration), mock.Result(mock.RedisBool(true)))
	return &ExpectedBool{exp: e}
}

func (m *ClientMock) ExpectGetSet(key string, value any) *ExpectedString {
	e := m.push(mock.Match("GETSET", key, str(value)), mock.Result(mock.RedisNil()))
	return &ExpectedString{exp: e}
}

func (m *ClientMock) ExpectAppend(key, value string) *ExpectedInt {
	e := m.push(mock.Match("APPEND", key, value), mock.Result(mock.RedisInt64(0)))
	return &ExpectedInt{exp: e}
}

func (m *ClientMock) ExpectStrLen(key string) *ExpectedInt {
	e := m.push(mock.Match("STRLEN", key), mock.Result(mock.RedisInt64(0)))
	return &ExpectedInt{exp: e}
}

func (m *ClientMock) ExpectDel(keys ...string) *ExpectedInt {
	e := m.push(delMatcher(keys...), mock.Result(mock.RedisInt64(0)))
	return &ExpectedInt{exp: e}
}

func (m *ClientMock) ExpectExists(keys ...string) *ExpectedInt {
	args := append([]string{"EXISTS"}, keys...)
	e := m.push(mock.Match(args...), mock.Result(mock.RedisInt64(0)))
	return &ExpectedInt{exp: e}
}

func (m *ClientMock) ExpectType(key string) *ExpectedStatus {
	e := m.push(mock.Match("TYPE", key), mock.Result(mock.RedisString("none")))
	return &ExpectedStatus{exp: e}
}

func (m *ClientMock) ExpectTTL(key string) *ExpectedDuration {
	e := m.push(mock.Match("TTL", key), mock.Result(mock.RedisInt64(-2)))
	return &ExpectedDuration{exp: e, precision: time.Second}
}

func (m *ClientMock) ExpectExpire(key string, expiration time.Duration) *ExpectedBool {
	e := m.push(mock.Match("EXPIRE", key, strconv.FormatInt(formatSec(expiration), 10)), mock.Result(mock.RedisBool(false)))
	return &ExpectedBool{exp: e}
}

func (m *ClientMock) ExpectPing() *ExpectedStatus {
	e := m.push(mock.Match("PING"), mock.Result(mock.RedisString("PONG")))
	return &ExpectedStatus{exp: e}
}

func (m *ClientMock) ExpectEcho(message any) *ExpectedString {
	e := m.push(mock.Match("ECHO", str(message)), mock.Result(mock.RedisString("")))
	return &ExpectedString{exp: e}
}

func (m *ClientMock) ExpectIncr(key string) *ExpectedInt {
	e := m.push(mock.Match("INCR", key), mock.Result(mock.RedisInt64(0)))
	return &ExpectedInt{exp: e}
}

func (m *ClientMock) ExpectIncrBy(key string, value int64) *ExpectedInt {
	e := m.push(mock.Match("INCRBY", key, strconv.FormatInt(value, 10)), mock.Result(mock.RedisInt64(0)))
	return &ExpectedInt{exp: e}
}

func (m *ClientMock) ExpectDecr(key string) *ExpectedInt {
	e := m.push(mock.Match("DECR", key), mock.Result(mock.RedisInt64(0)))
	return &ExpectedInt{exp: e}
}

func (m *ClientMock) ExpectDecrBy(key string, value int64) *ExpectedInt {
	e := m.push(mock.Match("DECRBY", key, strconv.FormatInt(value, 10)), mock.Result(mock.RedisInt64(0)))
	return &ExpectedInt{exp: e}
}

func (m *ClientMock) ExpectMGet(keys ...string) *ExpectedSlice {
	args := append([]string{"MGET"}, keys...)
	e := m.push(mock.Match(args...), mock.Result(mock.RedisArray()))
	return &ExpectedSlice{exp: e}
}

func (m *ClientMock) ExpectMSet(values ...any) *ExpectedStatus {
	args := []string{"MSET"}
	for _, v := range values {
		args = append(args, str(v))
	}
	e := m.push(mock.Match(args...), mock.Result(mock.RedisString("OK")))
	return &ExpectedStatus{exp: e}
}

func (m *ClientMock) ExpectHGet(key, field string) *ExpectedString {
	e := m.push(mock.Match("HGET", key, field), mock.Result(mock.RedisNil()))
	return &ExpectedString{exp: e}
}

func (m *ClientMock) ExpectHSet(key string, values ...any) *ExpectedInt {
	args := append([]string{"HSET", key}, hsetArgsToSlice(values)...)
	e := m.push(mock.Match(args...), mock.Result(mock.RedisInt64(0)))
	return &ExpectedInt{exp: e}
}

func (m *ClientMock) ExpectHDel(key string, fields ...string) *ExpectedInt {
	args := append([]string{"HDEL", key}, fields...)
	e := m.push(mock.Match(args...), mock.Result(mock.RedisInt64(0)))
	return &ExpectedInt{exp: e}
}

func (m *ClientMock) ExpectHGetAll(key string) *ExpectedStringStringMap {
	e := m.push(mock.Match("HGETALL", key), mock.Result(mock.RedisMap(map[string]rueidis.RedisMessage{})))
	return &ExpectedStringStringMap{exp: e}
}

func (m *ClientMock) ExpectLPush(key string, elements ...any) *ExpectedInt {
	args := []string{"LPUSH", key}
	for _, el := range elements {
		args = append(args, str(el))
	}
	e := m.push(mock.Match(args...), mock.Result(mock.RedisInt64(0)))
	return &ExpectedInt{exp: e}
}

func (m *ClientMock) ExpectRPush(key string, elements ...any) *ExpectedInt {
	args := []string{"RPUSH", key}
	for _, el := range elements {
		args = append(args, str(el))
	}
	e := m.push(mock.Match(args...), mock.Result(mock.RedisInt64(0)))
	return &ExpectedInt{exp: e}
}

func (m *ClientMock) ExpectLPop(key string) *ExpectedString {
	e := m.push(mock.Match("LPOP", key), mock.Result(mock.RedisNil()))
	return &ExpectedString{exp: e}
}

func (m *ClientMock) ExpectRPop(key string) *ExpectedString {
	e := m.push(mock.Match("RPOP", key), mock.Result(mock.RedisNil()))
	return &ExpectedString{exp: e}
}

func (m *ClientMock) ExpectLLen(key string) *ExpectedInt {
	e := m.push(mock.Match("LLEN", key), mock.Result(mock.RedisInt64(0)))
	return &ExpectedInt{exp: e}
}

func (m *ClientMock) ExpectSAdd(key string, members ...any) *ExpectedInt {
	args := []string{"SADD", key}
	for _, mm := range members {
		args = append(args, str(mm))
	}
	e := m.push(mock.Match(args...), mock.Result(mock.RedisInt64(0)))
	return &ExpectedInt{exp: e}
}

func (m *ClientMock) ExpectSRem(key string, members ...any) *ExpectedInt {
	args := []string{"SREM", key}
	for _, mm := range members {
		args = append(args, str(mm))
	}
	e := m.push(mock.Match(args...), mock.Result(mock.RedisInt64(0)))
	return &ExpectedInt{exp: e}
}

func (m *ClientMock) ExpectSMembers(key string) *ExpectedStringSlice {
	e := m.push(mock.Match("SMEMBERS", key), mock.Result(mock.RedisArray()))
	return &ExpectedStringSlice{exp: e}
}

func (m *ClientMock) ExpectEval(script string, keys []string, args ...any) *ExpectedCmd {
	e := m.push(evalMatcher(script, keys, args...), mock.Result(mock.RedisNil()))
	return &ExpectedCmd{exp: e}
}

func setMatcher(key string, value any, expiration time.Duration) gomock.Matcher {
	if expiration > 0 {
		if usePrecise(expiration) {
			return mock.Match("SET", key, str(value), "PX", strconv.FormatInt(formatMs(expiration), 10))
		}
		return mock.Match("SET", key, str(value), "EX", strconv.FormatInt(formatSec(expiration), 10))
	}
	if expiration == keepTTL {
		return mock.Match("SET", key, str(value), "KEEPTTL")
	}
	return mock.Match("SET", key, str(value))
}

func setNXMatcher(key string, value any, expiration time.Duration) gomock.Matcher {
	switch expiration {
	case 0:
		return mock.Match("SETNX", key, str(value))
	case keepTTL:
		return mock.Match("SET", key, str(value), "NX", "KEEPTTL")
	}
	if usePrecise(expiration) {
		return mock.Match("SET", key, str(value), "NX", "PX", strconv.FormatInt(formatMs(expiration), 10))
	}
	return mock.Match("SET", key, str(value), "NX", "EX", strconv.FormatInt(formatSec(expiration), 10))
}

func delMatcher(keys ...string) gomock.Matcher {
	args := append([]string{"DEL"}, keys...)
	return mock.Match(args...)
}

func hsetArgsToSlice(values []any) []string {
	if len(values) == 1 {
		switch v := values[0].(type) {
		case []string:
			return v
		case []any:
			out := make([]string, 0, len(v))
			for _, x := range v {
				out = append(out, str(x))
			}
			return out
		case map[string]any:
			out := make([]string, 0, len(v)*2)
			for k, x := range v {
				out = append(out, k, str(x))
			}
			return out
		case map[string]string:
			out := make([]string, 0, len(v)*2)
			for k, x := range v {
				out = append(out, k, x)
			}
			return out
		}
	}
	out := make([]string, 0, len(values))
	for _, v := range values {
		out = append(out, str(v))
	}
	return out
}

func evalMatcher(script string, keys []string, args ...any) gomock.Matcher {
	parts := []string{"EVAL", script, strconv.Itoa(len(keys))}
	parts = append(parts, keys...)
	for _, a := range args {
		parts = append(parts, str(a))
	}
	return mock.Match(parts...)
}
