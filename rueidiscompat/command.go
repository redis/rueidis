package rueidiscompat

import (
	"fmt"
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

type BoolSliceCmd struct {
	res rueidis.RedisResult
	val []bool
	err error
}

func newBoolSliceCmd(res rueidis.RedisResult) *BoolSliceCmd {
	ints, err := res.AsIntSlice()
	if err != nil {
		return &BoolSliceCmd{res: res, err: err}
	}
	val := make([]bool, 0, len(ints))
	for i := 0; i < len(ints); i++ {
		val = append(val, i == 1)
	}
	return &BoolSliceCmd{res: res, val: val, err: err}
}

func (cmd *BoolSliceCmd) SetVal(val []bool) {
	cmd.val = val
}

func (cmd *BoolSliceCmd) Val() []bool {
	return cmd.val
}

func (cmd *BoolSliceCmd) Result() ([]bool, error) {
	return cmd.val, cmd.err
}

func (cmd *BoolSliceCmd) String() (string, error) {
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
	keys   []string
	err    error
}

func newScanCmd(res rueidis.RedisResult) *ScanCmd {
	cursorSlice, err := res.ToArray()
	if err != nil {
		return &ScanCmd{res: res, err: err}
	}
	rawCursor, rawPage := cursorSlice[0], cursorSlice[1]
	cursor, err := rawCursor.ToInt64()
	if err != nil {
		return &ScanCmd{res: res, err: err}
	}
	page, err := rawPage.AsStrSlice()
	return &ScanCmd{res: res, cursor: uint64(cursor), keys: page, err: err}
}

func (cmd *ScanCmd) SetVal(page []string, cursor uint64) {
	cmd.keys = page
	cmd.cursor = cursor
}

func (cmd *ScanCmd) Val() (keys []string, cursor uint64) {
	return cmd.keys, cmd.cursor
}

func (cmd *ScanCmd) Result() (keys []string, cursor uint64, err error) {
	return cmd.keys, cmd.cursor, cmd.err
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

type StringStructMapCmd struct {
	res rueidis.RedisResult
	val map[string]struct{}
	err error
}

func newStringStructMapCmd(res rueidis.RedisResult) *StringStructMapCmd {
	strSlice, err := res.AsStrSlice()
	if err != nil {
		return &StringStructMapCmd{res: res, err: err}
	}
	val := make(map[string]struct{}, len(strSlice))
	for _, v := range strSlice {
		val[v] = struct{}{}
	}
	return &StringStructMapCmd{res: res, val: val, err: err}
}

func (cmd *StringStructMapCmd) SetVal(val map[string]struct{}) {
	cmd.val = val
}

func (cmd *StringStructMapCmd) Val() map[string]struct{} {
	return cmd.val
}

func (cmd *StringStructMapCmd) Result() (map[string]struct{}, error) {
	return cmd.val, cmd.err
}

func (cmd *StringStructMapCmd) String() (string, error) {
	return cmd.res.ToString()
}

type XMessageSliceCmd struct {
	res rueidis.RedisResult
	val []rueidis.XMessage
	err error
}

func newXMessageSliceCmd(res rueidis.RedisResult) *XMessageSliceCmd {
	val, err := res.AsXMessageSlice()
	return &XMessageSliceCmd{res: res, val: val, err: err}
}

func (cmd *XMessageSliceCmd) SetVal(val []rueidis.XMessage) {
	cmd.val = val
}

func (cmd *XMessageSliceCmd) Val() []rueidis.XMessage {
	return cmd.val
}

func (cmd *XMessageSliceCmd) Result() ([]rueidis.XMessage, error) {
	return cmd.val, cmd.err
}

func (cmd *XMessageSliceCmd) String() (string, error) {
	return cmd.res.ToString()
}

type XStream struct {
	Stream   string
	Messages []rueidis.XMessage
}

type XStreamSliceCmd struct {
	res rueidis.RedisResult
	val []XStream
	err error
}

func newXStreamSliceCmd(res rueidis.RedisResult) *XStreamSliceCmd {
	arrs, err := res.ToArray()
	if err != nil {
		return &XStreamSliceCmd{res: res, err: err}
	}
	val := make([]XStream, 0, len(arrs))
	for _, v := range arrs {
		arr, err := v.ToArray()
		if err != nil {
			return &XStreamSliceCmd{res: res, err: err}
		}
		if len(arr) != 2 {
			return &XStreamSliceCmd{res: res, err: fmt.Errorf("got %d, wanted 2", len(arr))}
		}
		stream, err := arr[0].ToString()
		if err != nil {
			return &XStreamSliceCmd{res: res, err: err}
		}
		msgs, err := arr[1].AsXMessageSlice()
		if err != nil {
			return &XStreamSliceCmd{res: res, err: err}
		}
		val = append(val, XStream{Stream: stream, Messages: msgs})
	}
	return &XStreamSliceCmd{res: res, val: val, err: err}
}

func (cmd *XStreamSliceCmd) SetVal(val []XStream) {
	cmd.val = val
}

func (cmd *XStreamSliceCmd) Val() []XStream {
	return cmd.val
}

func (cmd *XStreamSliceCmd) Result() ([]XStream, error) {
	return cmd.val, cmd.err
}

func (cmd *XStreamSliceCmd) String() (string, error) {
	return cmd.res.ToString()
}

type XPending struct {
	Count     int64
	Lower     string
	Higher    string
	Consumers map[string]int64
}

type XPendingCmd struct {
	res rueidis.RedisResult
	val XPending
	err error
}

func newXPendingCmd(res rueidis.RedisResult) *XPendingCmd {
	arr, err := res.ToArray()
	if err != nil {
		return &XPendingCmd{res: res, err: err}
	}
	if len(arr) != 4 {
		return &XPendingCmd{res: res, err: fmt.Errorf("got %d, wanted 4", len(arr))}
	}
	count, err := arr[0].ToInt64()
	if err != nil {
		return &XPendingCmd{res: res, err: err}
	}
	lower, err := arr[1].ToString()
	if err != nil {
		return &XPendingCmd{res: res, err: err}
	}
	higher, err := arr[2].ToString()
	if err != nil {
		return &XPendingCmd{res: res, err: err}
	}
	val := XPending{
		Count:  count,
		Lower:  lower,
		Higher: higher,
	}
	consumerArr, err := arr[3].ToArray()
	if err != nil {
		return &XPendingCmd{res: res, err: err}
	}
	for _, v := range consumerArr {
		consumer, err := v.ToArray()
		if err != nil {
			return &XPendingCmd{res: res, err: err}
		}
		if len(arr) != 2 {
			return &XPendingCmd{res: res, err: fmt.Errorf("got %d, wanted 2", len(arr))}
		}
		consumerName, err := consumer[0].ToString()
		if err != nil {
			return &XPendingCmd{res: res, err: err}
		}
		consumerPending, err := consumer[1].AsInt64()
		if err != nil {
			return &XPendingCmd{res: res, err: err}
		}
		if val.Consumers == nil {
			val.Consumers = make(map[string]int64)
		}
		val.Consumers[consumerName] = consumerPending
	}
	return &XPendingCmd{res: res, val: val, err: err}
}

func (cmd *XPendingCmd) SetVal(val XPending) {
	cmd.val = val
}

func (cmd *XPendingCmd) Val() XPending {
	return cmd.val
}

func (cmd *XPendingCmd) Result() (XPending, error) {
	return cmd.val, cmd.err
}

func (cmd *XPendingCmd) String() (string, error) {
	return cmd.res.ToString()
}

type XPendingExt struct {
	ID         string
	Consumer   string
	Idle       time.Duration
	RetryCount int64
}

type XPendingExtCmd struct {
	res rueidis.RedisResult
	val []XPendingExt
	err error
}

func newXPendingExtCmd(res rueidis.RedisResult) *XPendingExtCmd {
	arrs, err := res.ToArray()
	if err != nil {
		return &XPendingExtCmd{res: res, err: err}
	}
	val := make([]XPendingExt, 0, len(arrs))
	for _, v := range arrs {
		arr, err := v.ToArray()
		if err != nil {
			return &XPendingExtCmd{res: res, err: err}
		}
		if len(arr) != 4 {
			return &XPendingExtCmd{res: res, err: fmt.Errorf("got %d, wanted 4", len(arr))}
		}
		id, err := arr[0].ToString()
		if err != nil {
			return &XPendingExtCmd{res: res, err: err}
		}
		consumer, err := arr[1].ToString()
		if err != nil {
			return &XPendingExtCmd{res: res, err: err}
		}
		idle, err := arr[2].ToInt64()
		if err != nil {
			return &XPendingExtCmd{res: res, err: err}
		}
		retryCount, err := arr[3].ToInt64()
		if err != nil {
			return &XPendingExtCmd{res: res, err: err}
		}
		val = append(val, XPendingExt{
			ID:         id,
			Consumer:   consumer,
			Idle:       time.Duration(idle) * time.Millisecond,
			RetryCount: retryCount,
		})
	}
	return &XPendingExtCmd{res: res, val: val, err: err}
}

func (cmd *XPendingExtCmd) SetVal(val []XPendingExt) {
	cmd.val = val
}

func (cmd *XPendingExtCmd) Val() []XPendingExt {
	return cmd.val
}

func (cmd *XPendingExtCmd) Result() ([]XPendingExt, error) {
	return cmd.val, cmd.err
}

func (cmd *XPendingExtCmd) String() (string, error) {
	return cmd.res.ToString()
}

type XAutoClaimCmd struct {
	res   rueidis.RedisResult
	start string
	val   []rueidis.XMessage
	err   error
}

func newXAutoClaimCmd(res rueidis.RedisResult) *XAutoClaimCmd {
	arr, err := res.ToArray()
	if err != nil {
		return &XAutoClaimCmd{res: res, err: err}
	}
	if len(arr) != 2 {
		return &XAutoClaimCmd{res: res, err: fmt.Errorf("got %d, wanted 2", len(arr))}
	}
	start, err := arr[0].ToString()
	if err != nil {
		return &XAutoClaimCmd{res: res, err: err}
	}
	val, err := arr[1].AsXMessageSlice()
	if err != nil {
		return &XAutoClaimCmd{res: res, err: err}
	}
	return &XAutoClaimCmd{res: res, val: val, start: start, err: err}
}

func (cmd *XAutoClaimCmd) SetVal(val []rueidis.XMessage, start string) {
	cmd.val = val
	cmd.start = start
}

func (cmd *XAutoClaimCmd) Val() (messages []rueidis.XMessage, start string) {
	return cmd.val, cmd.start
}

func (cmd *XAutoClaimCmd) Result() (messages []rueidis.XMessage, start string, err error) {
	return cmd.val, cmd.start, cmd.err
}

func (cmd *XAutoClaimCmd) String() (string, error) {
	return cmd.res.ToString()
}

type XAutoClaimJustIDCmd struct {
	res   rueidis.RedisResult
	start string
	val   []string
	err   error
}

func newXAutoClaimJustIDCmd(res rueidis.RedisResult) *XAutoClaimJustIDCmd {
	arr, err := res.ToArray()
	if err != nil {
		return &XAutoClaimJustIDCmd{res: res, err: err}
	}
	if len(arr) != 2 {
		return &XAutoClaimJustIDCmd{res: res, err: fmt.Errorf("got %d, wanted 2", len(arr))}
	}
	start, err := arr[0].ToString()
	if err != nil {
		return &XAutoClaimJustIDCmd{res: res, err: err}
	}
	val, err := arr[1].AsStrSlice()
	if err != nil {
		return &XAutoClaimJustIDCmd{res: res, err: err}
	}
	return &XAutoClaimJustIDCmd{res: res, val: val, start: start, err: err}
}

func (cmd *XAutoClaimJustIDCmd) SetVal(val []string, start string) {
	cmd.val = val
	cmd.start = start
}

func (cmd *XAutoClaimJustIDCmd) Val() (ids []string, start string) {
	return cmd.val, cmd.start
}

func (cmd *XAutoClaimJustIDCmd) Result() (ids []string, start string, err error) {
	return cmd.val, cmd.start, cmd.err
}

func (cmd *XAutoClaimJustIDCmd) String() (string, error) {
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

// Note: len(Fields) and len(Values) must be the same.
// MaxLen/MaxLenApprox and MinID are in conflict, only one of them can be used.
type XAddArgs struct {
	Stream     string
	NoMkStream bool
	MaxLen     int64 // MAXLEN N

	MinID string
	// Approx causes MaxLen and MinID to use "~" matcher (instead of "=").
	Approx bool
	Limit  int64
	ID     string
	Fields []string
	Values []string
}

// Note: len(Streams) and len(IDs) must be the same.
type XReadArgs struct {
	Streams []string // list of streams
	IDs     []string // list of ids
	Count   int64
	Block   time.Duration
}

// Note: len(Streams) and len(IDs) must be the same.
type XReadGroupArgs struct {
	Group    string
	Consumer string
	Streams  []string // list of streams
	IDs      []string // list of ids
	Count    int64
	Block    time.Duration
	NoAck    bool
}

type XPendingExtArgs struct {
	Stream   string
	Group    string
	Idle     time.Duration
	Start    string
	End      string
	Count    int64
	Consumer string
}

type XClaimArgs struct {
	Stream   string
	Group    string
	Consumer string
	MinIdle  time.Duration
	Messages []string
}

type XAutoClaimArgs struct {
	Stream   string
	Group    string
	MinIdle  time.Duration
	Start    string
	Count    int64
	Consumer string
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
