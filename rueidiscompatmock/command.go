package rueidiscompatmock

import (
	"time"

	"github.com/redis/rueidis"
	"github.com/redis/rueidis/mock"
)

type ExpectedString struct{ exp *expectation }

func (e *ExpectedString) SetVal(v string) *ExpectedString {
	e.exp.result = mock.Result(mock.RedisString(v))
	return e
}

func (e *ExpectedString) SetErr(err error) *ExpectedString {
	e.exp.result = mock.ErrorResult(err)
	return e
}

func (e *ExpectedString) RedisNil() *ExpectedString {
	e.exp.result = mock.Result(mock.RedisNil())
	return e
}

type ExpectedStatus = ExpectedString

type ExpectedBool struct{ exp *expectation }

func (e *ExpectedBool) SetVal(v bool) *ExpectedBool {
	e.exp.result = mock.Result(mock.RedisBool(v))
	return e
}

func (e *ExpectedBool) SetErr(err error) *ExpectedBool {
	e.exp.result = mock.ErrorResult(err)
	return e
}

type ExpectedInt struct{ exp *expectation }

func (e *ExpectedInt) SetVal(v int64) *ExpectedInt {
	e.exp.result = mock.Result(mock.RedisInt64(v))
	return e
}

func (e *ExpectedInt) SetErr(err error) *ExpectedInt {
	e.exp.result = mock.ErrorResult(err)
	return e
}

type ExpectedDuration struct {
	exp       *expectation
	precision time.Duration
}

func (e *ExpectedDuration) SetVal(v time.Duration) *ExpectedDuration {
	e.exp.result = mock.Result(mock.RedisInt64(int64(v / e.precision)))
	return e
}

func (e *ExpectedDuration) SetErr(err error) *ExpectedDuration {
	e.exp.result = mock.ErrorResult(err)
	return e
}

type ExpectedStringSlice struct{ exp *expectation }

func (e *ExpectedStringSlice) SetVal(v []string) *ExpectedStringSlice {
	msgs := make([]rueidis.RedisMessage, 0, len(v))
	for _, s := range v {
		msgs = append(msgs, mock.RedisString(s))
	}
	e.exp.result = mock.Result(mock.RedisArray(msgs...))
	return e
}

func (e *ExpectedStringSlice) SetErr(err error) *ExpectedStringSlice {
	e.exp.result = mock.ErrorResult(err)
	return e
}

type ExpectedSlice struct{ exp *expectation }

func (e *ExpectedSlice) SetVal(v []any) *ExpectedSlice {
	msgs := make([]rueidis.RedisMessage, 0, len(v))
	for _, item := range v {
		switch x := item.(type) {
		case nil:
			msgs = append(msgs, mock.RedisNil())
		case string:
			msgs = append(msgs, mock.RedisString(x))
		case int64:
			msgs = append(msgs, mock.RedisInt64(x))
		case bool:
			msgs = append(msgs, mock.RedisBool(x))
		default:
			msgs = append(msgs, mock.RedisString(str(item)))
		}
	}
	e.exp.result = mock.Result(mock.RedisArray(msgs...))
	return e
}

func (e *ExpectedSlice) SetErr(err error) *ExpectedSlice {
	e.exp.result = mock.ErrorResult(err)
	return e
}

type ExpectedStringStringMap struct{ exp *expectation }

func (e *ExpectedStringStringMap) SetVal(v map[string]string) *ExpectedStringStringMap {
	kv := make(map[string]rueidis.RedisMessage, len(v))
	for k, val := range v {
		kv[k] = mock.RedisString(val)
	}
	e.exp.result = mock.Result(mock.RedisMap(kv))
	return e
}

func (e *ExpectedStringStringMap) SetErr(err error) *ExpectedStringStringMap {
	e.exp.result = mock.ErrorResult(err)
	return e
}

type ExpectedCmd struct{ exp *expectation }

func (e *ExpectedCmd) SetVal(v string) *ExpectedCmd {
	e.exp.result = mock.Result(mock.RedisString(v))
	return e
}

func (e *ExpectedCmd) SetValInt(v int64) *ExpectedCmd {
	e.exp.result = mock.Result(mock.RedisInt64(v))
	return e
}

func (e *ExpectedCmd) SetErr(err error) *ExpectedCmd {
	e.exp.result = mock.ErrorResult(err)
	return e
}

func (e *ExpectedCmd) RedisNil() *ExpectedCmd {
	e.exp.result = mock.Result(mock.RedisNil())
	return e
}
