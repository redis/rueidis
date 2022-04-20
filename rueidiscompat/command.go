package rueidiscompat

import (
	"strconv"
	"time"

	"github.com/rueian/rueidis"
)

type StringCmd struct {
	res rueidis.RedisResult
	val string
	err error
}

func newStringCmd(res rueidis.RedisResult) *StringCmd {
	val, err := res.ToString()
	return &StringCmd{res: res, val: val, err: err}
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
	return cmd.res.ToBool()
}

func (cmd *StringCmd) Int() (int, error) {
	i, err := cmd.res.ToInt64()
	return int(i), err
}

func (cmd *StringCmd) Int64() (int64, error) {
	return cmd.res.ToInt64()
}

func (cmd *StringCmd) Uint64() (uint64, error) {
	if cmd.err != nil {
		return 0, cmd.err
	}
	return strconv.ParseUint(cmd.Val(), 10, 64)
}

func (cmd *StringCmd) Float32() (float32, error) {
	f, err := cmd.res.ToFloat64()
	if err != nil {
		return 0, err
	}
	return float32(f), nil
}

func (cmd *StringCmd) Float64() (float64, error) {
	return cmd.res.ToFloat64()
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
	res rueidis.RedisResult
	val bool
	err error
}

func newBoolCmd(res rueidis.RedisResult) *BoolCmd {
	val, err := res.AsBool()
	if rueidis.IsRedisNil(err) {
		val = false
		err = nil
	}
	return &BoolCmd{res: res, val: val, err: err}
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
	return cmd.res.ToString()
}

type IntCmd struct {
	res rueidis.RedisResult
	val int64
	err error
}

func newIntCmd(res rueidis.RedisResult) *IntCmd {
	val, err := res.AsInt64()
	return &IntCmd{res: res, val: val, err: err}
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
	return cmd.res.ToString()
}

type StatusCmd struct {
	res rueidis.RedisResult
	val string
	err error
}

func newStatusCmd(res rueidis.RedisResult) *StatusCmd {
	val, err := res.ToString()
	return &StatusCmd{res: res, val: val, err: err}
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
	return cmd.res.ToString()
}

type SliceCmd struct {
	res rueidis.RedisResult
	val []rueidis.RedisMessage
	err error
}

func newSliceCmd(res rueidis.RedisResult) *SliceCmd {
	val, err := res.ToArray()
	return &SliceCmd{res: res, val: val, err: err}
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
	return cmd.res.ToString()
}

type StringSliceCmd struct {
	res rueidis.RedisResult
	val []string
	err error
}

func newStringSliceCmd(res rueidis.RedisResult) *StringSliceCmd {
	val, err := res.AsStrSlice()
	return &StringSliceCmd{res: res, val: val, err: err}
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
	return cmd.res.ToString()
}

type IntSliceCmd struct {
	res rueidis.RedisResult
	val []int64
	err error
}

func newIntSliceCmd(res rueidis.RedisResult) *IntSliceCmd {
	val, err := res.AsIntSlice()
	return &IntSliceCmd{res: res, val: val, err: err}
}

func (cmd *IntSliceCmd) SetVal(val []int64) {
	cmd.val = val
}

func (cmd *IntSliceCmd) Val() []int64 {
	return cmd.val
}

func (cmd *IntSliceCmd) Result() ([]int64, error) {
	return cmd.val, cmd.err
}

func (cmd *IntSliceCmd) String() (string, error) {
	return cmd.res.ToString()
}

type FloatCmd struct {
	res rueidis.RedisResult
	val float64
	err error
}

func newFloatCmd(res rueidis.RedisResult) *FloatCmd {
	val, err := res.ToFloat64()
	return &FloatCmd{res: res, val: val, err: err}
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
	return cmd.res.ToString()
}

type ScanCmd struct {
	res    rueidis.RedisResult
	cursor uint64
	page   []string
	err    error
}

func newScanCmd(res rueidis.RedisResult) *ScanCmd {
	cursor_list, err := res.ToArray()
	if err != nil {
		return &ScanCmd{res: res, err: err}
	}
	raw_cursor, raw_page := cursor_list[0], cursor_list[1]
	cursor, err := raw_cursor.ToInt64()
	if err != nil {
		return &ScanCmd{res: res, err: err}
	}
	page, err := raw_page.AsStrSlice()
	return &ScanCmd{res: res, cursor: uint64(cursor), page: page, err: err}
}

func (cmd *ScanCmd) SetVal(page []string, cursor uint64) {
	cmd.page = page
	cmd.cursor = cursor
}

func (cmd *ScanCmd) Val() (keys []string, cursor uint64) {
	return cmd.page, cmd.cursor
}

func (cmd *ScanCmd) Result() (keys []string, cursor uint64, err error) {
	return cmd.page, cmd.cursor, cmd.err
}

func (cmd *ScanCmd) String() (string, error) {
	return cmd.res.ToString()
}

type StringStringMapCmd struct {
	res rueidis.RedisResult
	val map[string]string
	err error
}

func newStringStringMapCmd(res rueidis.RedisResult) *StringStringMapCmd {
	val, err := res.AsStrMap()
	return &StringStringMapCmd{res: res, val: val, err: err}
}

func (cmd *StringStringMapCmd) SetVal(val map[string]string) {
	cmd.val = val
}

func (cmd *StringStringMapCmd) Val() map[string]string {
	return cmd.val
}

func (cmd *StringStringMapCmd) Result() (map[string]string, error) {
	return cmd.val, cmd.err
}

func (cmd *StringStringMapCmd) String() (string, error) {
	return cmd.res.ToString()
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

type BitPos struct {
	BitCount
	Byte bool
}

type BitFieldArg struct {
	Encoding string
	Offset   int64
}

type BitField struct {
	Get       *BitFieldArg
	Set       *BitFieldArg
	IncrBy    *BitFieldArg
	Increment int64
	Overflow  string
}

type LPosArgs struct {
	Rank, MaxLen int64
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
