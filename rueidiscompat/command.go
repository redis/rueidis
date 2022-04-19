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

type StatusCmd struct {
	baseCmd
	val string
}

func NewStatusCmd(ctx context.Context, result rueidis.RedisResult) *StatusCmd {
	v, err := result.ToString()
	return &StatusCmd{
		baseCmd: baseCmd{
			RedisResult: result,
			ctx:         ctx,
			err:         err,
		},
		val: v,
	}
}

func (cmd *StatusCmd) SetVal(val string) {
	cmd.val = val
}

func (cmd *StatusCmd) Val() string {
	return cmd.val
}

func (cmd *StatusCmd) Result() (string, error) {
	return cmd.val, cmd.err
}

func (cmd *StatusCmd) String() (string, error) {
	return cmd.ToString()
}

type SliceCmd struct {
	baseCmd
	val []rueidis.RedisMessage
}

func NewSliceCmd(ctx context.Context, result rueidis.RedisResult) *SliceCmd {
	v, err := result.ToArray()
	return &SliceCmd{
		baseCmd: baseCmd{
			RedisResult: result,
			ctx:         ctx,
			err:         err,
		},
		val: v,
	}
}

func (cmd *SliceCmd) SetVal(val []rueidis.RedisMessage) {
	cmd.val = val
}

func (cmd *SliceCmd) Val() []rueidis.RedisMessage {
	return cmd.val
}

func (cmd *SliceCmd) Result() ([]rueidis.RedisMessage, error) {
	return cmd.val, cmd.err
}

func (cmd *SliceCmd) String() (string, error) {
	return cmd.ToString()
}

type StringSliceCmd struct {
	baseCmd
	val []string
}

func NewStringSliceCmd(ctx context.Context, result rueidis.RedisResult) *StringSliceCmd {
	v, err := result.AsStrSlice()
	return &StringSliceCmd{
		baseCmd: baseCmd{
			RedisResult: result,
			ctx:         ctx,
			err:         err,
		},
		val: v,
	}
}

func (cmd *StringSliceCmd) SetVal(val []string) {
	cmd.val = val
}

func (cmd *StringSliceCmd) Val() []string {
	return cmd.val
}

func (cmd *StringSliceCmd) Result() ([]string, error) {
	return cmd.val, cmd.err
}

func (cmd *StringSliceCmd) String() (string, error) {
	return cmd.ToString()
}

type FloatCmd struct {
	baseCmd
	val float64
}

func NewFloatCmd(ctx context.Context, result rueidis.RedisResult) *FloatCmd {
	v, err := result.ToFloat64()
	return &FloatCmd{
		baseCmd: baseCmd{
			RedisResult: result,
			ctx:         ctx,
			err:         err,
		},
		val: v,
	}
}

func (cmd *FloatCmd) SetVal(val float64) {
	cmd.val = val
}

func (cmd *FloatCmd) Val() float64 {
	return cmd.val
}

func (cmd *FloatCmd) Result() (float64, error) {
	return cmd.val, cmd.err
}

func (cmd *FloatCmd) String() (string, error) {
	cmd.ToArray()
	return cmd.ToString()
}

type Sort struct {
	By            string
	Offset, Count int64
	Get           []string
	Order         string
	Alpha         bool
}

// SetArgs provides arguments for the SetArgs function.
type SetArgs struct {
	// Mode can be `NX` or `XX` or empty.
	Mode string

	// Zero `TTL` or `Expiration` means that the key has no expiration time.
	TTL      time.Duration
	ExpireAt time.Time

	// When Get is true, the command returns the old value stored at key, or nil when key did not exist.
	Get bool

	// KeepTTL is a Redis KEEPTTL option to keep existing TTL, it requires your redis-server version >= 6.0,
	// otherwise you will receive an error: (error) ERR syntax error.
	KeepTTL bool
}

type BitCount struct {
	Start, End int64
}

func usePrecise(dur time.Duration) bool {
	return dur < time.Second || dur%time.Second != 0
}

func formatMs(dur time.Duration) int64 {
	if dur > 0 && dur < time.Millisecond {
		// too small, truncate too 1ms
		return 1
	}
	return int64(dur / time.Millisecond)
}

func formatSec(dur time.Duration) int64 {
	if dur > 0 && dur < time.Second {
		// too small ,truncate too 1s
		return 1
	}
	return int64(dur / time.Second)
}
