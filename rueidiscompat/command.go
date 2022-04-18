package rueidiscompat

import (
	"context"
	"strconv"
	"time"

	"github.com/rueian/rueidis"
)

type baseCmd struct {
	rueidis.RedisResult
	ctx context.Context
	err error
}

type StringCmd struct {
	baseCmd
	val string
}

func NewStringCmd(ctx context.Context, result rueidis.RedisResult) *StringCmd {
	val, err := result.ToString()
	return &StringCmd{
		baseCmd: baseCmd{
			RedisResult: result,
			ctx:         ctx,
			err:         err,
		},
		val: val,
	}
}

func (cmd *StringCmd) SetVal(val string) {
	cmd.val = val
}

func (cmd *StringCmd) Val() string {
	return cmd.val
}

func (cmd *StringCmd) Result() (string, error) {
	return cmd.val, cmd.err
}

func (cmd *StringCmd) Bytes() ([]byte, error) {
	return []byte(cmd.val), cmd.err
}

func (cmd *StringCmd) Bool() (bool, error) {
	return cmd.ToBool()
}

func (cmd *StringCmd) Int() (int, error) {
	i, err := cmd.ToInt64()
	return int(i), err
}

func (cmd *StringCmd) Int64() (int64, error) {
	return cmd.ToInt64()
}

func (cmd *StringCmd) Uint64() (uint64, error) {
	if cmd.err != nil {
		return 0, cmd.err
	}
	return strconv.ParseUint(cmd.Val(), 10, 64)
}

func (cmd *StringCmd) Float32() (float32, error) {
	f, err := cmd.ToFloat64()
	if err != nil {
		return 0, err
	}
	return float32(f), nil
}

func (cmd *StringCmd) Float64() (float64, error) {
	return cmd.ToFloat64()
}

func (cmd *StringCmd) Time() (time.Time, error) {
	if cmd.err != nil {
		return time.Time{}, cmd.err
	}
	return time.Parse(time.RFC3339Nano, cmd.Val())
}

func (cmd *StringCmd) String() string {
	return cmd.val
}

type BoolCmd struct {
	baseCmd
	val bool
}

func NewBoolCmd(ctx context.Context, result rueidis.RedisResult) *BoolCmd {
	v, err := result.AsBool()
	if rueidis.IsRedisNil(err) {
		v = false
		err = nil
	}
	return &BoolCmd{
		baseCmd: baseCmd{
			RedisResult: result,
			ctx:         ctx,
			err:         err,
		},
		val: v,
	}
}

func (cmd *BoolCmd) SetVal(val bool) {
	cmd.val = val
}

func (cmd *BoolCmd) Val() bool {
	return cmd.val
}

func (cmd *BoolCmd) Result() (bool, error) {
	return cmd.val, cmd.err
}

func (cmd *BoolCmd) String() (string, error) {
	return cmd.ToString()
}

type IntCmd struct {
	baseCmd
	val int64
}

func NewIntCmd(ctx context.Context, result rueidis.RedisResult) *IntCmd {
	v, err := result.AsInt64()
	return &IntCmd{
		baseCmd: baseCmd{
			RedisResult: result,
			ctx:         ctx,
			err:         err,
		},
		val: v,
	}
}

func (cmd *IntCmd) SetVal(val int64) {
	cmd.val = val
}

func (cmd *IntCmd) Val() int64 {
	return cmd.val
}

func (cmd *IntCmd) Result() (int64, error) {
	return cmd.val, cmd.err
}

func (cmd *IntCmd) Uint64() (uint64, error) {
	return uint64(cmd.val), cmd.err
}

func (cmd *IntCmd) String() (string, error) {
	return cmd.ToString()
}

type Sort struct {
	By            string
	Offset, Count int64
	Get           []string
	Order         string
	Alpha         bool
}
