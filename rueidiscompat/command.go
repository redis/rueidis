// Copyright (c) 2013 The github.com/go-redis/redis Authors.
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
// * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
// * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package rueidiscompat

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/redis/rueidis"
	"github.com/redis/rueidis/internal/util"
)

type baseCmd[T any] struct {
	err error
	val T
}

func (cmd *baseCmd[T]) SetVal(val T) {
	cmd.val = val
}

func (cmd *baseCmd[T]) Val() T {
	return cmd.val
}

func (cmd *baseCmd[T]) SetErr(err error) {
	cmd.err = err
}

func (cmd *baseCmd[T]) Err() error {
	return cmd.err
}

func (cmd *baseCmd[T]) Result() (T, error) {
	return cmd.Val(), cmd.Err()
}

type Cmd struct {
	baseCmd[any]
}

func newCmd(res rueidis.RedisResult) *Cmd {
	cmd := &Cmd{}
	val, err := res.ToAny()
	if err != nil {
		cmd.err = err
		return cmd
	}
	cmd.SetVal(val)
	return cmd
}

func (cmd *Cmd) Text() (string, error) {
	if cmd.err != nil {
		return "", cmd.err
	}
	return toString(cmd.val)
}

func toString(val any) (string, error) {
	switch val := val.(type) {
	case string:
		return val, nil
	default:
		err := fmt.Errorf("redis: unexpected type=%T for String", val)
		return "", err
	}
}

func (cmd *Cmd) Int() (int, error) {
	if cmd.err != nil {
		return 0, cmd.err
	}
	switch val := cmd.val.(type) {
	case int64:
		return int(val), nil
	case string:
		return strconv.Atoi(val)
	default:
		err := fmt.Errorf("redis: unexpected type=%T for Int", val)
		return 0, err
	}
}

func (cmd *Cmd) Int64() (int64, error) {
	if cmd.err != nil {
		return 0, cmd.err
	}
	return toInt64(cmd.val)
}

func toInt64(val any) (int64, error) {
	switch val := val.(type) {
	case int64:
		return val, nil
	case string:
		return strconv.ParseInt(val, 10, 64)
	default:
		err := fmt.Errorf("redis: unexpected type=%T for Int64", val)
		return 0, err
	}
}

func (cmd *Cmd) Uint64() (uint64, error) {
	if cmd.err != nil {
		return 0, cmd.err
	}
	return toUint64(cmd.val)
}

func toUint64(val any) (uint64, error) {
	switch val := val.(type) {
	case int64:
		return uint64(val), nil
	case string:
		return strconv.ParseUint(val, 10, 64)
	default:
		err := fmt.Errorf("redis: unexpected type=%T for Uint64", val)
		return 0, err
	}
}

func (cmd *Cmd) Float32() (float32, error) {
	if cmd.err != nil {
		return 0, cmd.err
	}
	return toFloat32(cmd.val)
}

func toFloat32(val any) (float32, error) {
	switch val := val.(type) {
	case int64:
		return float32(val), nil
	case string:
		f, err := util.ToFloat32(val)
		if err != nil {
			return 0, err
		}
		return f, nil
	default:
		err := fmt.Errorf("redis: unexpected type=%T for Float32", val)
		return 0, err
	}
}

func (cmd *Cmd) Float64() (float64, error) {
	if cmd.err != nil {
		return 0, cmd.err
	}
	return toFloat64(cmd.val)
}

func toFloat64(val any) (float64, error) {
	switch val := val.(type) {
	case int64:
		return float64(val), nil
	case string:
		return util.ToFloat64(val)
	default:
		err := fmt.Errorf("redis: unexpected type=%T for Float64", val)
		return 0, err
	}
}

func (cmd *Cmd) Bool() (bool, error) {
	if cmd.err != nil {
		return false, cmd.err
	}
	return toBool(cmd.val)
}

func toBool(val any) (bool, error) {
	switch val := val.(type) {
	case int64:
		return val != 0, nil
	case string:
		return strconv.ParseBool(val)
	default:
		err := fmt.Errorf("redis: unexpected type=%T for Bool", val)
		return false, err
	}
}

func (cmd *Cmd) Slice() ([]any, error) {
	if cmd.err != nil {
		return nil, cmd.err
	}
	switch val := cmd.val.(type) {
	case []any:
		return val, nil
	default:
		return nil, fmt.Errorf("redis: unexpected type=%T for Slice", val)
	}
}

func (cmd *Cmd) StringSlice() ([]string, error) {
	slice, err := cmd.Slice()
	if err != nil {
		return nil, err
	}

	ss := make([]string, len(slice))
	for i, iface := range slice {
		val, err := toString(iface)
		if err != nil {
			return nil, err
		}
		ss[i] = val
	}
	return ss, nil
}

func (cmd *Cmd) Int64Slice() ([]int64, error) {
	slice, err := cmd.Slice()
	if err != nil {
		return nil, err
	}

	nums := make([]int64, len(slice))
	for i, iface := range slice {
		val, err := toInt64(iface)
		if err != nil {
			return nil, err
		}
		nums[i] = val
	}
	return nums, nil
}

func (cmd *Cmd) Uint64Slice() ([]uint64, error) {
	slice, err := cmd.Slice()
	if err != nil {
		return nil, err
	}

	nums := make([]uint64, len(slice))
	for i, iface := range slice {
		val, err := toUint64(iface)
		if err != nil {
			return nil, err
		}
		nums[i] = val
	}
	return nums, nil
}

func (cmd *Cmd) Float32Slice() ([]float32, error) {
	slice, err := cmd.Slice()
	if err != nil {
		return nil, err
	}

	floats := make([]float32, len(slice))
	for i, iface := range slice {
		val, err := toFloat32(iface)
		if err != nil {
			return nil, err
		}
		floats[i] = val
	}
	return floats, nil
}

func (cmd *Cmd) Float64Slice() ([]float64, error) {
	slice, err := cmd.Slice()
	if err != nil {
		return nil, err
	}

	floats := make([]float64, len(slice))
	for i, iface := range slice {
		val, err := toFloat64(iface)
		if err != nil {
			return nil, err
		}
		floats[i] = val
	}
	return floats, nil
}

func (cmd *Cmd) BoolSlice() ([]bool, error) {
	slice, err := cmd.Slice()
	if err != nil {
		return nil, err
	}

	bools := make([]bool, len(slice))
	for i, iface := range slice {
		val, err := toBool(iface)
		if err != nil {
			return nil, err
		}
		bools[i] = val
	}
	return bools, nil
}

type StringCmd struct {
	baseCmd[string]
}

func newStringCmd(res rueidis.RedisResult) *StringCmd {
	cmd := &StringCmd{}
	val, err := res.ToString()
	cmd.SetErr(err)
	cmd.SetVal(val)
	return cmd
}

func (cmd *StringCmd) Bytes() ([]byte, error) {
	return []byte(cmd.val), cmd.err
}

func (cmd *StringCmd) Bool() (bool, error) {
	return cmd.val != "", cmd.err
}

func (cmd *StringCmd) Int() (int, error) {
	if cmd.err != nil {
		return 0, cmd.err
	}
	return strconv.Atoi(cmd.Val())
}

func (cmd *StringCmd) Int64() (int64, error) {
	if cmd.err != nil {
		return 0, cmd.err
	}
	return strconv.ParseInt(cmd.Val(), 10, 64)
}

func (cmd *StringCmd) Uint64() (uint64, error) {
	if cmd.err != nil {
		return 0, cmd.err
	}
	return strconv.ParseUint(cmd.Val(), 10, 64)
}

func (cmd *StringCmd) Float32() (float32, error) {
	if cmd.err != nil {
		return 0, cmd.err
	}
	v, err := util.ToFloat32(cmd.Val())
	if err != nil {
		return 0, err
	}
	return v, nil
}

func (cmd *StringCmd) Float64() (float64, error) {
	if cmd.err != nil {
		return 0, cmd.err
	}
	return util.ToFloat64(cmd.Val())
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
	baseCmd[bool]
}

func newBoolCmd(res rueidis.RedisResult) *BoolCmd {
	cmd := &BoolCmd{}
	val, err := res.AsBool()
	if rueidis.IsRedisNil(err) {
		val = false
		err = nil
	}
	cmd.SetVal(val)
	cmd.SetErr(err)
	return cmd
}

type IntCmd struct {
	baseCmd[int64]
}

func newIntCmd(res rueidis.RedisResult) *IntCmd {
	cmd := &IntCmd{}
	val, err := res.AsInt64()
	cmd.SetErr(err)
	cmd.SetVal(val)
	return cmd
}

func (cmd *IntCmd) Uint64() (uint64, error) {
	return uint64(cmd.val), cmd.err
}

type DurationCmd struct {
	baseCmd[time.Duration]
}

func newDurationCmd(res rueidis.RedisResult, precision time.Duration) *DurationCmd {
	cmd := &DurationCmd{}
	val, err := res.AsInt64()
	cmd.SetErr(err)
	if val > 0 {
		cmd.SetVal(time.Duration(val) * precision)
		return cmd
	}
	cmd.SetVal(time.Duration(val))
	return cmd
}

type StatusCmd = StringCmd

func newStatusCmd(res rueidis.RedisResult) *StatusCmd {
	cmd := &StatusCmd{}
	val, err := res.ToString()
	cmd.SetErr(err)
	cmd.SetVal(val)
	return cmd
}

type SliceCmd struct {
	baseCmd[[]any]
	keys []string
}

// newSliceCmd returns SliceCmd according to input arguments, if the caller is JSONObjKeys,
// set isJSONObjKeys to true.
func newSliceCmd(res rueidis.RedisResult, isJSONObjKeys bool, keys ...string) *SliceCmd {
	cmd := &SliceCmd{keys: keys}
	arr, err := res.ToArray()
	if err != nil {
		cmd.SetErr(err)
		return cmd
	}
	vals := make([]any, len(arr))
	if isJSONObjKeys {
		for i, v := range arr {
			// for JSON.OBJKEYS
			if v.IsNil() {
				continue
			}
			// convert to any which underlying type is []any
			arr, err := v.ToAny()
			if err != nil {
				cmd.SetErr(err)
				return cmd
			}
			vals[i] = arr
		}
		cmd.SetVal(vals)
		return cmd
	}
	for i, v := range arr {
		// keep the old behavior the same as before (don't handle error while parsing v as string)
		if s, err := v.ToString(); err == nil {
			vals[i] = s
		}
	}
	cmd.SetVal(vals)
	return cmd
}

// Scan scans the results from the map into a destination struct. The map keys
// are matched in the Redis struct fields by the `redis:"field"` tag.
// NOTE: result from JSON.ObjKeys should not call this.
func (cmd *SliceCmd) Scan(dst any) error {
	if cmd.err != nil {
		return cmd.err
	}
	return Scan(dst, cmd.keys, cmd.val)
}

type StringSliceCmd struct {
	baseCmd[[]string]
}

func newStringSliceCmd(res rueidis.RedisResult) *StringSliceCmd {
	cmd := &StringSliceCmd{}
	val, err := res.AsStrSlice()
	cmd.SetVal(val)
	cmd.SetErr(err)
	return cmd
}

type IntSliceCmd struct {
	err error
	val []int64
}

func newIntSliceCmd(res rueidis.RedisResult) *IntSliceCmd {
	val, err := res.AsIntSlice()
	return &IntSliceCmd{val: val, err: err}
}

func (cmd *IntSliceCmd) SetVal(val []int64) {
	cmd.val = val
}

func (cmd *IntSliceCmd) SetErr(err error) {
	cmd.err = err
}

func (cmd *IntSliceCmd) Val() []int64 {
	return cmd.val
}

func (cmd *IntSliceCmd) Err() error {
	return cmd.err
}

func (cmd *IntSliceCmd) Result() ([]int64, error) {
	return cmd.val, cmd.err
}

type BoolSliceCmd struct {
	baseCmd[[]bool]
}

func newBoolSliceCmd(res rueidis.RedisResult) *BoolSliceCmd {
	cmd := &BoolSliceCmd{}
	ints, err := res.AsIntSlice()
	if err != nil {
		cmd.SetErr(err)
		return cmd
	}
	val := make([]bool, 0, len(ints))
	for _, i := range ints {
		val = append(val, i == 1)
	}
	cmd.SetVal(val)
	return cmd
}

type FloatSliceCmd struct {
	baseCmd[[]float64]
}

func newFloatSliceCmd(res rueidis.RedisResult) *FloatSliceCmd {
	cmd := &FloatSliceCmd{}
	val, err := res.AsFloatSlice()
	cmd.SetErr(err)
	cmd.SetVal(val)
	return cmd
}

type ZSliceCmd struct {
	baseCmd[[]Z]
}

func newZSliceCmd(res rueidis.RedisResult) *ZSliceCmd {
	cmd := &ZSliceCmd{}
	scores, err := res.AsZScores()
	if err != nil {
		cmd.SetErr(err)
		return cmd
	}
	val := make([]Z, 0, len(scores))
	for _, s := range scores {
		val = append(val, Z{Member: s.Member, Score: s.Score})
	}
	cmd.SetVal(val)
	return cmd
}

func newZSliceSingleCmd(res rueidis.RedisResult) *ZSliceCmd {
	cmd := &ZSliceCmd{}
	s, err := res.AsZScore()
	if err != nil {
		cmd.SetErr(err)
		return cmd
	}
	cmd.SetVal([]Z{{Member: s.Member, Score: s.Score}})
	return cmd
}

type FloatCmd struct {
	baseCmd[float64]
}

func newFloatCmd(res rueidis.RedisResult) *FloatCmd {
	cmd := &FloatCmd{}
	val, err := res.AsFloat64()
	cmd.SetErr(err)
	cmd.SetVal(val)
	return cmd
}

type ScanCmd struct {
	err    error
	keys   []string
	cursor uint64
}

func newScanCmd(res rueidis.RedisResult) *ScanCmd {
	e, err := res.AsScanEntry()
	return &ScanCmd{cursor: e.Cursor, keys: e.Elements, err: err}
}

func (cmd *ScanCmd) SetVal(keys []string, cursor uint64) {
	cmd.keys = keys
	cmd.cursor = cursor
}

func (cmd *ScanCmd) Val() (keys []string, cursor uint64) {
	return cmd.keys, cmd.cursor
}

func (cmd *ScanCmd) Err() error {
	return cmd.err
}

func (cmd *ScanCmd) Result() (keys []string, cursor uint64, err error) {
	return cmd.keys, cmd.cursor, cmd.err
}

type KeyValue struct {
	Key   string
	Value string
}

type KeyValueSliceCmd struct {
	baseCmd[[]KeyValue]
}

func newKeyValueSliceCmd(res rueidis.RedisResult) *KeyValueSliceCmd {
	cmd := &KeyValueSliceCmd{}
	arr, err := res.ToArray()
	for _, a := range arr {
		kv, _ := a.AsStrSlice()
		for i := 0; i < len(kv); i += 2 {
			cmd.val = append(cmd.val, KeyValue{Key: kv[i], Value: kv[i+1]})
		}
	}
	cmd.SetErr(err)
	return cmd
}

type KeyValuesCmd struct {
	err error
	val rueidis.KeyValues
}

func newKeyValuesCmd(res rueidis.RedisResult) *KeyValuesCmd {
	ret := &KeyValuesCmd{}
	ret.val, ret.err = res.AsLMPop()
	return ret
}

func (cmd *KeyValuesCmd) SetVal(key string, val []string) {
	cmd.val.Key = key
	cmd.val.Values = val
}

func (cmd *KeyValuesCmd) SetErr(err error) {
	cmd.err = err
}

func (cmd *KeyValuesCmd) Val() (string, []string) {
	return cmd.val.Key, cmd.val.Values
}

func (cmd *KeyValuesCmd) Err() error {
	return cmd.err
}

func (cmd *KeyValuesCmd) Result() (string, []string, error) {
	return cmd.val.Key, cmd.val.Values, cmd.err
}

type KeyFlags struct {
	Key   string
	Flags []string
}

type KeyFlagsCmd struct {
	baseCmd[[]KeyFlags]
}

func newKeyFlagsCmd(res rueidis.RedisResult) *KeyFlagsCmd {
	ret := &KeyFlagsCmd{}
	if ret.err = res.Error(); ret.err == nil {
		kfs, _ := res.ToArray()
		ret.val = make([]KeyFlags, len(kfs))
		for i := 0; i < len(kfs); i++ {
			if kf, _ := kfs[i].ToArray(); len(kf) >= 2 {
				ret.val[i].Key, _ = kf[0].ToString()
				ret.val[i].Flags, _ = kf[1].AsStrSlice()
			}
		}
	}
	return ret
}

type ZSliceWithKeyCmd struct {
	err error
	key string
	val []Z
}

func newZSliceWithKeyCmd(res rueidis.RedisResult) *ZSliceWithKeyCmd {
	v, err := res.AsZMPop()
	if err != nil {
		return &ZSliceWithKeyCmd{err: err}
	}
	val := make([]Z, 0, len(v.Values))
	for _, s := range v.Values {
		val = append(val, Z{Member: s.Member, Score: s.Score})
	}
	return &ZSliceWithKeyCmd{key: v.Key, val: val}
}

func (cmd *ZSliceWithKeyCmd) SetVal(key string, val []Z) {
	cmd.key = key
	cmd.val = val
}

func (cmd *ZSliceWithKeyCmd) SetErr(err error) {
	cmd.err = err
}

func (cmd *ZSliceWithKeyCmd) Val() (string, []Z) {
	return cmd.key, cmd.val
}

func (cmd *ZSliceWithKeyCmd) Err() error {
	return cmd.err
}

func (cmd *ZSliceWithKeyCmd) Result() (string, []Z, error) {
	return cmd.key, cmd.val, cmd.err
}

type StringStringMapCmd struct {
	baseCmd[map[string]string]
}

func newStringStringMapCmd(res rueidis.RedisResult) *StringStringMapCmd {
	cmd := &StringStringMapCmd{}
	val, err := res.AsStrMap()
	cmd.SetErr(err)
	cmd.SetVal(val)
	return cmd
}

// Scan scans the results from the map into a destination struct. The map keys
// are matched in the Redis struct fields by the `redis:"field"` tag.
func (cmd *StringStringMapCmd) Scan(dest interface{}) error {
	if cmd.Err() != nil {
		return cmd.Err()
	}

	strct, err := Struct(dest)
	if err != nil {
		return err
	}

	for k, v := range cmd.val {
		if err := strct.Scan(k, v); err != nil {
			return err
		}
	}

	return nil
}

type StringIntMapCmd struct {
	baseCmd[map[string]int64]
}

func newStringIntMapCmd(res rueidis.RedisResult) *StringIntMapCmd {
	cmd := &StringIntMapCmd{}
	val, err := res.AsIntMap()
	cmd.SetErr(err)
	cmd.SetVal(val)
	return cmd
}

type StringStructMapCmd struct {
	baseCmd[map[string]struct{}]
}

func newStringStructMapCmd(res rueidis.RedisResult) *StringStructMapCmd {
	cmd := &StringStructMapCmd{}
	strSlice, err := res.AsStrSlice()
	if err != nil {
		cmd.SetErr(err)
		return cmd
	}
	val := make(map[string]struct{}, len(strSlice))
	for _, v := range strSlice {
		val[v] = struct{}{}
	}
	cmd.SetVal(val)
	return cmd
}

type XMessageSliceCmd struct {
	baseCmd[[]XMessage]
}

func newXMessageSliceCmd(res rueidis.RedisResult) *XMessageSliceCmd {
	cmd := &XMessageSliceCmd{}
	val, err := res.AsXRange()
	cmd.SetErr(err)
	cmd.val = make([]XMessage, len(val))
	for i, r := range val {
		cmd.val[i] = newXMessage(r)
	}
	return cmd
}

func newXMessage(r rueidis.XRangeEntry) XMessage {
	if r.FieldValues == nil {
		return XMessage{ID: r.ID, Values: nil}
	}
	m := XMessage{ID: r.ID, Values: make(map[string]any, len(r.FieldValues))}
	for k, v := range r.FieldValues {
		m.Values[k] = v
	}
	return m
}

type XStream struct {
	Stream   string
	Messages []XMessage
}

type XStreamSliceCmd struct {
	baseCmd[[]XStream]
}

func newXStreamSliceCmd(res rueidis.RedisResult) *XStreamSliceCmd {
	cmd := &XStreamSliceCmd{}
	streams, err := res.AsXRead()
	if err != nil {
		cmd.SetErr(err)
		return cmd
	}
	val := make([]XStream, 0, len(streams))
	for name, messages := range streams {
		msgs := make([]XMessage, 0, len(messages))
		for _, r := range messages {
			msgs = append(msgs, newXMessage(r))
		}
		val = append(val, XStream{Stream: name, Messages: msgs})
	}
	cmd.SetVal(val)
	return cmd
}

type XPending struct {
	Consumers map[string]int64
	Lower     string
	Higher    string
	Count     int64
}

type XPendingCmd struct {
	baseCmd[XPending]
}

func newXPendingCmd(res rueidis.RedisResult) *XPendingCmd {
	cmd := &XPendingCmd{}
	arr, err := res.ToArray()
	if err != nil {
		cmd.SetErr(err)
		return cmd
	}
	if len(arr) < 4 {
		cmd.SetErr(fmt.Errorf("got %d, wanted 4", len(arr)))
		return cmd
	}
	count, err := arr[0].AsInt64()
	if err != nil {
		cmd.SetErr(err)
		return cmd
	}
	lower, err := arr[1].ToString()
	if err != nil {
		cmd.SetErr(err)
		return cmd
	}
	higher, err := arr[2].ToString()
	if err != nil {
		cmd.SetErr(err)
		return cmd
	}
	val := XPending{
		Count:  count,
		Lower:  lower,
		Higher: higher,
	}
	consumerArr, err := arr[3].ToArray()
	if err != nil {
		cmd.SetErr(err)
		return cmd
	}
	for _, v := range consumerArr {
		consumer, err := v.ToArray()
		if err != nil {
			cmd.SetErr(err)
			return cmd
		}
		if len(consumer) < 2 {
			cmd.SetErr(fmt.Errorf("got %d, wanted 2", len(arr)))
			return cmd
		}
		consumerName, err := consumer[0].ToString()
		if err != nil {
			cmd.SetErr(err)
			return cmd
		}
		consumerPending, err := consumer[1].AsInt64()
		if err != nil {
			cmd.SetErr(err)
			return cmd
		}
		if val.Consumers == nil {
			val.Consumers = make(map[string]int64)
		}
		val.Consumers[consumerName] = consumerPending
	}
	cmd.SetVal(val)
	return cmd
}

type XPendingExt struct {
	ID         string
	Consumer   string
	Idle       time.Duration
	RetryCount int64
}

type XPendingExtCmd struct {
	baseCmd[[]XPendingExt]
}

func newXPendingExtCmd(res rueidis.RedisResult) *XPendingExtCmd {
	cmd := &XPendingExtCmd{}
	arrs, err := res.ToArray()
	if err != nil {
		cmd.SetErr(err)
		return cmd
	}
	val := make([]XPendingExt, 0, len(arrs))
	for _, v := range arrs {
		arr, err := v.ToArray()
		if err != nil {
			cmd.SetErr(err)
			return cmd
		}
		if len(arr) < 4 {
			cmd.SetErr(fmt.Errorf("got %d, wanted 4", len(arr)))
			return cmd
		}
		id, err := arr[0].ToString()
		if err != nil {
			cmd.SetErr(err)
			return cmd
		}
		consumer, err := arr[1].ToString()
		if err != nil {
			cmd.SetErr(err)
			return cmd
		}
		idle, err := arr[2].AsInt64()
		if err != nil {
			cmd.SetErr(err)
			return cmd
		}
		retryCount, err := arr[3].AsInt64()
		if err != nil {
			cmd.SetErr(err)
			return cmd
		}
		val = append(val, XPendingExt{
			ID:         id,
			Consumer:   consumer,
			Idle:       time.Duration(idle) * time.Millisecond,
			RetryCount: retryCount,
		})
	}
	cmd.SetVal(val)
	return cmd
}

type XAutoClaimCmd struct {
	err   error
	start string
	val   []XMessage
}

func newXAutoClaimCmd(res rueidis.RedisResult) *XAutoClaimCmd {
	arr, err := res.ToArray()
	if err != nil {
		return &XAutoClaimCmd{err: err}
	}
	if len(arr) < 2 {
		return &XAutoClaimCmd{err: fmt.Errorf("got %d, wanted 2", len(arr))}
	}
	start, err := arr[0].ToString()
	if err != nil {
		return &XAutoClaimCmd{err: err}
	}
	ranges, err := arr[1].AsXRange()
	if err != nil {
		return &XAutoClaimCmd{err: err}
	}
	val := make([]XMessage, 0, len(ranges))
	for _, r := range ranges {
		val = append(val, newXMessage(r))
	}
	return &XAutoClaimCmd{val: val, start: start, err: err}
}

func (cmd *XAutoClaimCmd) SetVal(val []XMessage, start string) {
	cmd.val = val
	cmd.start = start
}

func (cmd *XAutoClaimCmd) SetErr(err error) {
	cmd.err = err
}

func (cmd *XAutoClaimCmd) Val() (messages []XMessage, start string) {
	return cmd.val, cmd.start
}

func (cmd *XAutoClaimCmd) Err() error {
	return cmd.err
}

func (cmd *XAutoClaimCmd) Result() (messages []XMessage, start string, err error) {
	return cmd.val, cmd.start, cmd.err
}

type XAutoClaimJustIDCmd struct {
	err   error
	start string
	val   []string
}

func newXAutoClaimJustIDCmd(res rueidis.RedisResult) *XAutoClaimJustIDCmd {
	arr, err := res.ToArray()
	if err != nil {
		return &XAutoClaimJustIDCmd{err: err}
	}
	if len(arr) < 2 {
		return &XAutoClaimJustIDCmd{err: fmt.Errorf("got %d, wanted 2", len(arr))}
	}
	start, err := arr[0].ToString()
	if err != nil {
		return &XAutoClaimJustIDCmd{err: err}
	}
	val, err := arr[1].AsStrSlice()
	if err != nil {
		return &XAutoClaimJustIDCmd{err: err}
	}
	return &XAutoClaimJustIDCmd{val: val, start: start, err: err}
}

func (cmd *XAutoClaimJustIDCmd) SetVal(val []string, start string) {
	cmd.val = val
	cmd.start = start
}

func (cmd *XAutoClaimJustIDCmd) SetErr(err error) {
	cmd.err = err
}

func (cmd *XAutoClaimJustIDCmd) Val() (ids []string, start string) {
	return cmd.val, cmd.start
}

func (cmd *XAutoClaimJustIDCmd) Err() error {
	return cmd.err
}

func (cmd *XAutoClaimJustIDCmd) Result() (ids []string, start string, err error) {
	return cmd.val, cmd.start, cmd.err
}

type XInfoGroup struct {
	Name            string
	LastDeliveredID string
	Consumers       int64
	Pending         int64
	EntriesRead     int64
	Lag             int64
}

type XInfoGroupsCmd struct {
	baseCmd[[]XInfoGroup]
}

func newXInfoGroupsCmd(res rueidis.RedisResult) *XInfoGroupsCmd {
	cmd := &XInfoGroupsCmd{}
	arr, err := res.ToArray()
	if err != nil {
		cmd.SetErr(err)
		return cmd
	}
	groupInfos := make([]XInfoGroup, 0, len(arr))
	for _, v := range arr {
		info, err := v.AsMap()
		if err != nil {
			cmd.SetErr(err)
			return cmd
		}
		var group XInfoGroup
		if attr, ok := info["name"]; ok {
			group.Name, _ = attr.ToString()
		}
		if attr, ok := info["consumers"]; ok {
			group.Consumers, _ = attr.AsInt64()
		}
		if attr, ok := info["pending"]; ok {
			group.Pending, _ = attr.AsInt64()
		}
		if attr, ok := info["entries-read"]; ok {
			group.EntriesRead, _ = attr.AsInt64()
		}
		if attr, ok := info["lag"]; ok {
			group.Lag, _ = attr.AsInt64()
		}
		if attr, ok := info["last-delivered-id"]; ok {
			group.LastDeliveredID, _ = attr.ToString()
		}
		groupInfos = append(groupInfos, group)
	}
	cmd.SetVal(groupInfos)
	return cmd
}

type XInfoStream struct {
	FirstEntry           XMessage
	LastEntry            XMessage
	LastGeneratedID      string
	MaxDeletedEntryID    string
	RecordedFirstEntryID string
	Length               int64
	RadixTreeKeys        int64
	RadixTreeNodes       int64
	Groups               int64
	EntriesAdded         int64
}
type XInfoStreamCmd struct {
	baseCmd[XInfoStream]
}

func newXInfoStreamCmd(res rueidis.RedisResult) *XInfoStreamCmd {
	cmd := &XInfoStreamCmd{}
	kv, err := res.AsMap()
	if err != nil {
		cmd.SetErr(err)
		return cmd
	}
	var val XInfoStream
	if v, ok := kv["length"]; ok {
		val.Length, _ = v.AsInt64()
	}
	if v, ok := kv["radix-tree-keys"]; ok {
		val.RadixTreeKeys, _ = v.AsInt64()
	}
	if v, ok := kv["radix-tree-nodes"]; ok {
		val.RadixTreeNodes, _ = v.AsInt64()
	}
	if v, ok := kv["groups"]; ok {
		val.Groups, _ = v.AsInt64()
	}
	if v, ok := kv["last-generated-id"]; ok {
		val.LastGeneratedID, _ = v.ToString()
	}
	if v, ok := kv["max-deleted-entry-id"]; ok {
		val.MaxDeletedEntryID, _ = v.ToString()
	}
	if v, ok := kv["recorded-first-entry-id"]; ok {
		val.RecordedFirstEntryID, _ = v.ToString()
	}
	if v, ok := kv["entries-added"]; ok {
		val.EntriesAdded, _ = v.AsInt64()
	}
	if v, ok := kv["first-entry"]; ok {
		if r, err := v.AsXRangeEntry(); err == nil {
			val.FirstEntry = newXMessage(r)
		}
	}
	if v, ok := kv["last-entry"]; ok {
		if r, err := v.AsXRangeEntry(); err == nil {
			val.LastEntry = newXMessage(r)
		}
	}
	cmd.SetVal(val)
	return cmd
}

type XInfoStreamConsumerPending struct {
	DeliveryTime  time.Time
	ID            string
	DeliveryCount int64
}

type XInfoStreamGroupPending struct {
	DeliveryTime  time.Time
	ID            string
	Consumer      string
	DeliveryCount int64
}

type XInfoStreamConsumer struct {
	SeenTime time.Time
	Name     string
	Pending  []XInfoStreamConsumerPending
	PelCount int64
}

type XInfoStreamGroup struct {
	Name            string
	LastDeliveredID string
	Pending         []XInfoStreamGroupPending
	Consumers       []XInfoStreamConsumer
	EntriesRead     int64
	Lag             int64
	PelCount        int64
}

type XInfoStreamFull struct {
	LastGeneratedID      string
	MaxDeletedEntryID    string
	RecordedFirstEntryID string
	Entries              []XMessage
	Groups               []XInfoStreamGroup
	Length               int64
	RadixTreeKeys        int64
	RadixTreeNodes       int64
	EntriesAdded         int64
}

type XInfoStreamFullCmd struct {
	baseCmd[XInfoStreamFull]
}

func newXInfoStreamFullCmd(res rueidis.RedisResult) *XInfoStreamFullCmd {
	cmd := &XInfoStreamFullCmd{}
	kv, err := res.AsMap()
	if err != nil {
		cmd.SetErr(err)
		return cmd
	}
	var val XInfoStreamFull
	if v, ok := kv["length"]; ok {
		val.Length, _ = v.AsInt64()
	}
	if v, ok := kv["radix-tree-keys"]; ok {
		val.RadixTreeKeys, _ = v.AsInt64()
	}
	if v, ok := kv["radix-tree-nodes"]; ok {
		val.RadixTreeNodes, _ = v.AsInt64()
	}
	if v, ok := kv["last-generated-id"]; ok {
		val.LastGeneratedID, _ = v.ToString()
	}
	if v, ok := kv["entries-added"]; ok {
		val.EntriesAdded, _ = v.AsInt64()
	}
	if v, ok := kv["max-deleted-entry-id"]; ok {
		val.MaxDeletedEntryID, _ = v.ToString()
	}
	if v, ok := kv["recorded-first-entry-id"]; ok {
		val.RecordedFirstEntryID, _ = v.ToString()
	}
	if v, ok := kv["groups"]; ok {
		val.Groups, err = readStreamGroups(v)
		if err != nil {
			cmd.SetErr(err)
			return cmd
		}
	}
	if v, ok := kv["entries"]; ok {
		ranges, err := v.AsXRange()
		if err != nil {
			cmd.SetErr(err)
			return cmd
		}
		val.Entries = make([]XMessage, 0, len(ranges))
		for _, r := range ranges {
			val.Entries = append(val.Entries, newXMessage(r))
		}
	}
	cmd.SetVal(val)
	return cmd
}

func readStreamGroups(res rueidis.RedisMessage) ([]XInfoStreamGroup, error) {
	arr, err := res.ToArray()
	if err != nil {
		return nil, err
	}
	groups := make([]XInfoStreamGroup, 0, len(arr))
	for _, v := range arr {
		info, err := v.AsMap()
		if err != nil {
			return nil, err
		}
		var group XInfoStreamGroup
		if attr, ok := info["name"]; ok {
			group.Name, _ = attr.ToString()
		}
		if attr, ok := info["last-delivered-id"]; ok {
			group.LastDeliveredID, _ = attr.ToString()
		}
		if attr, ok := info["entries-read"]; ok {
			group.EntriesRead, _ = attr.AsInt64()
		}
		if attr, ok := info["lag"]; ok {
			group.Lag, _ = attr.AsInt64()
		}
		if attr, ok := info["pel-count"]; ok {
			group.PelCount, _ = attr.AsInt64()
		}
		if attr, ok := info["pending"]; ok {
			group.Pending, err = readXInfoStreamGroupPending(attr)
			if err != nil {
				return nil, err
			}
		}
		if attr, ok := info["consumers"]; ok {
			group.Consumers, err = readXInfoStreamConsumers(attr)
			if err != nil {
				return nil, err
			}
		}
		groups = append(groups, group)
	}
	return groups, nil
}

func readXInfoStreamGroupPending(res rueidis.RedisMessage) ([]XInfoStreamGroupPending, error) {
	arr, err := res.ToArray()
	if err != nil {
		return nil, err
	}
	pending := make([]XInfoStreamGroupPending, 0, len(arr))
	for _, v := range arr {
		info, err := v.ToArray()
		if err != nil {
			return nil, err
		}
		if len(info) < 4 {
			return nil, fmt.Errorf("got %d, wanted 4", len(arr))
		}
		var p XInfoStreamGroupPending
		p.ID, err = info[0].ToString()
		if err != nil {
			return nil, err
		}
		p.Consumer, err = info[1].ToString()
		if err != nil {
			return nil, err
		}
		delivery, err := info[2].AsInt64()
		if err != nil {
			return nil, err
		}
		p.DeliveryTime = time.Unix(delivery/1000, delivery%1000*int64(time.Millisecond))
		p.DeliveryCount, err = info[3].AsInt64()
		if err != nil {
			return nil, err
		}
		pending = append(pending, p)
	}
	return pending, nil
}

func readXInfoStreamConsumers(res rueidis.RedisMessage) ([]XInfoStreamConsumer, error) {
	arr, err := res.ToArray()
	if err != nil {
		return nil, err
	}
	consumer := make([]XInfoStreamConsumer, 0, len(arr))
	for _, v := range arr {
		info, err := v.AsMap()
		if err != nil {
			return nil, err
		}
		var c XInfoStreamConsumer
		if attr, ok := info["name"]; ok {
			c.Name, _ = attr.ToString()
		}
		if attr, ok := info["seen-time"]; ok {
			seen, _ := attr.AsInt64()
			c.SeenTime = time.Unix(seen/1000, seen%1000*int64(time.Millisecond))
		}
		if attr, ok := info["pel-count"]; ok {
			c.PelCount, _ = attr.AsInt64()
		}
		if attr, ok := info["pending"]; ok {
			pending, err := attr.ToArray()
			if err != nil {
				return nil, err
			}
			c.Pending = make([]XInfoStreamConsumerPending, 0, len(pending))
			for _, v := range pending {
				pendingInfo, err := v.ToArray()
				if err != nil {
					return nil, err
				}
				if len(pendingInfo) < 3 {
					return nil, fmt.Errorf("got %d, wanted 3", len(arr))
				}
				var p XInfoStreamConsumerPending
				p.ID, err = pendingInfo[0].ToString()
				if err != nil {
					return nil, err
				}
				delivery, err := pendingInfo[1].AsInt64()
				if err != nil {
					return nil, err
				}
				p.DeliveryTime = time.Unix(delivery/1000, delivery%1000*int64(time.Millisecond))
				p.DeliveryCount, err = pendingInfo[2].AsInt64()
				if err != nil {
					return nil, err
				}
				c.Pending = append(c.Pending, p)
			}
		}
		consumer = append(consumer, c)
	}
	return consumer, nil
}

type XInfoConsumer struct {
	Name    string
	Pending int64
	Idle    time.Duration
}
type XInfoConsumersCmd struct {
	baseCmd[[]XInfoConsumer]
}

func newXInfoConsumersCmd(res rueidis.RedisResult) *XInfoConsumersCmd {
	cmd := &XInfoConsumersCmd{}
	arr, err := res.ToArray()
	if err != nil {
		cmd.SetErr(err)
		return cmd
	}
	val := make([]XInfoConsumer, 0, len(arr))
	for _, v := range arr {
		info, err := v.AsMap()
		if err != nil {
			cmd.SetErr(err)
			return cmd
		}
		var consumer XInfoConsumer
		if attr, ok := info["name"]; ok {
			consumer.Name, _ = attr.ToString()
		}
		if attr, ok := info["pending"]; ok {
			consumer.Pending, _ = attr.AsInt64()
		}
		if attr, ok := info["idle"]; ok {
			idle, _ := attr.AsInt64()
			consumer.Idle = time.Duration(idle) * time.Millisecond
		}
		val = append(val, consumer)
	}
	cmd.SetVal(val)
	return cmd
}

// Z represents sorted set member.
type Z struct {
	Member string
	Score  float64
}

// ZWithKey represents sorted set member including the name of the key where it was popped.
type ZWithKey struct {
	Z
	Key string
}

// ZStore is used as an arg to ZInter/ZInterStore and ZUnion/ZUnionStore.
type ZStore struct {
	Aggregate string
	Keys      []string
	Weights   []int64
}

type ZWithKeyCmd struct {
	baseCmd[ZWithKey]
}

func newZWithKeyCmd(res rueidis.RedisResult) *ZWithKeyCmd {
	cmd := &ZWithKeyCmd{}
	arr, err := res.ToArray()
	if err != nil {
		cmd.SetErr(err)
		return cmd
	}
	if len(arr) < 3 {
		cmd.SetErr(fmt.Errorf("got %d, wanted 3", len(arr)))
		return cmd
	}
	val := ZWithKey{}
	val.Key, err = arr[0].ToString()
	if err != nil {
		cmd.SetErr(err)
		return cmd
	}
	val.Member, err = arr[1].ToString()
	if err != nil {
		cmd.SetErr(err)
		return cmd
	}
	val.Score, err = arr[2].AsFloat64()
	if err != nil {
		cmd.SetErr(err)
		return cmd
	}
	cmd.SetVal(val)
	return cmd
}

type RankScore struct {
	Rank  int64
	Score float64
}

type RankWithScoreCmd struct {
	baseCmd[RankScore]
}

func newRankWithScoreCmd(res rueidis.RedisResult) *RankWithScoreCmd {
	ret := &RankWithScoreCmd{}
	if ret.err = res.Error(); ret.err == nil {
		vs, _ := res.ToArray()
		if len(vs) >= 2 {
			ret.val.Rank, _ = vs[0].AsInt64()
			ret.val.Score, _ = vs[1].AsFloat64()
		}
	}
	return ret
}

type TimeCmd struct {
	baseCmd[time.Time]
}

func newTimeCmd(res rueidis.RedisResult) *TimeCmd {
	cmd := &TimeCmd{}
	arr, err := res.ToArray()
	if err != nil {
		cmd.SetErr(err)
		return cmd
	}
	if len(arr) < 2 {
		cmd.SetErr(fmt.Errorf("got %d, wanted 2", len(arr)))
		return cmd
	}
	sec, err := arr[0].AsInt64()
	if err != nil {
		cmd.SetErr(err)
		return cmd
	}
	microSec, err := arr[1].AsInt64()
	if err != nil {
		cmd.SetErr(err)
		return cmd
	}
	cmd.SetVal(time.Unix(sec, microSec*1000))
	return cmd
}

type ClusterNode struct {
	ID   string
	Addr string
}

type ClusterSlot struct {
	Nodes []ClusterNode
	Start int64
	End   int64
}

type ClusterSlotsCmd struct {
	baseCmd[[]ClusterSlot]
}

func newClusterSlotsCmd(res rueidis.RedisResult) *ClusterSlotsCmd {
	cmd := &ClusterSlotsCmd{}
	arr, err := res.ToArray()
	if err != nil {
		cmd.SetErr(err)
		return cmd
	}
	val := make([]ClusterSlot, 0, len(arr))
	for _, v := range arr {
		slot, err := v.ToArray()
		if err != nil {
			cmd.SetErr(err)
			return cmd
		}
		if len(slot) < 2 {
			cmd.SetErr(fmt.Errorf("got %d, excpected atleast 2", len(slot)))
			return cmd
		}
		start, err := slot[0].AsInt64()
		if err != nil {
			cmd.SetErr(err)
			return cmd
		}
		end, err := slot[1].AsInt64()
		if err != nil {
			cmd.SetErr(err)
			return cmd
		}
		nodes := make([]ClusterNode, len(slot)-2)
		for i, j := 2, 0; i < len(slot); i, j = i+1, j+1 {
			node, err := slot[i].ToArray()
			if err != nil {
				cmd.SetErr(err)
				return cmd
			}
			if len(node) < 2 {
				cmd.SetErr(fmt.Errorf("got %d, expected 2 or 3", len(node)))
				return cmd
			}
			ip, err := node[0].ToString()
			if err != nil {
				cmd.SetErr(err)
				return cmd
			}
			port, err := node[1].AsInt64()
			if err != nil {
				cmd.SetErr(err)
				return cmd
			}
			nodes[j].Addr = net.JoinHostPort(ip, strconv.FormatInt(port, 10))
			if len(node) > 2 {
				id, err := node[2].ToString()
				if err != nil {
					cmd.SetErr(err)
					return cmd
				}
				nodes[j].ID = id
			}
		}
		val = append(val, ClusterSlot{
			Start: start,
			End:   end,
			Nodes: nodes,
		})
	}
	cmd.SetVal(val)
	return cmd
}

func newClusterShardsCmd(res rueidis.RedisResult) *ClusterShardsCmd {
	cmd := &ClusterShardsCmd{}
	arr, err := res.ToArray()
	if err != nil {
		cmd.SetErr(err)
		return cmd
	}
	val := make([]ClusterShard, 0, len(arr))
	for _, v := range arr {
		dict, err := v.ToMap()
		if err != nil {
			cmd.SetErr(err)
			return cmd
		}
		shard := ClusterShard{}
		{
			slots := dict["slots"]
			arr, _ := slots.ToArray()
			for i := 0; i+1 < len(arr); i += 2 {
				start, _ := arr[i].AsInt64()
				end, _ := arr[i+1].AsInt64()
				shard.Slots = append(shard.Slots, SlotRange{Start: start, End: end})
			}
		}
		{
			nodes, ok := dict["nodes"]
			if !ok {
				cmd.SetErr(errors.New("nodes not found"))
				return cmd
			}
			arr, err := nodes.ToArray()
			if err != nil {
				cmd.SetErr(err)
				return cmd
			}
			shard.Nodes = make([]Node, len(arr))
			for i := 0; i < len(arr); i++ {
				nodeMap, err := arr[i].ToMap()
				if err != nil {
					cmd.SetErr(err)
					return cmd
				}
				for k, v := range nodeMap {
					switch k {
					case "id":
						shard.Nodes[i].ID, _ = v.ToString()
					case "endpoint":
						shard.Nodes[i].Endpoint, _ = v.ToString()
					case "ip":
						shard.Nodes[i].IP, _ = v.ToString()
					case "hostname":
						shard.Nodes[i].Hostname, _ = v.ToString()
					case "port":
						shard.Nodes[i].Port, _ = v.ToInt64()
					case "tls-port":
						shard.Nodes[i].TLSPort, _ = v.ToInt64()
					case "role":
						shard.Nodes[i].Role, _ = v.ToString()
					case "replication-offset":
						shard.Nodes[i].ReplicationOffset, _ = v.ToInt64()
					case "health":
						shard.Nodes[i].Health, _ = v.ToString()
					}
				}
			}
		}
		val = append(val, shard)
	}
	cmd.SetVal(val)
	return cmd
}

type SlotRange struct {
	Start int64
	End   int64
}
type Node struct {
	ID                string
	Endpoint          string
	IP                string
	Hostname          string
	Role              string
	Health            string
	Port              int64
	TLSPort           int64
	ReplicationOffset int64
}
type ClusterShard struct {
	Slots []SlotRange
	Nodes []Node
}

type ClusterShardsCmd struct {
	baseCmd[[]ClusterShard]
}

type GeoPos struct {
	Longitude, Latitude float64
}

type GeoPosCmd struct {
	baseCmd[[]*GeoPos]
}

func newGeoPosCmd(res rueidis.RedisResult) *GeoPosCmd {
	cmd := &GeoPosCmd{}
	arr, err := res.ToArray()
	if err != nil {
		cmd.SetErr(err)
		return cmd
	}
	val := make([]*GeoPos, 0, len(arr))
	for _, v := range arr {
		loc, err := v.ToArray()
		if err != nil {
			if rueidis.IsRedisNil(err) {
				val = append(val, nil)
				continue
			}
			cmd.SetErr(err)
			return cmd
		}
		if len(loc) != 2 {
			cmd.SetErr(fmt.Errorf("got %d, expected 2", len(loc)))
			return cmd
		}
		long, err := loc[0].AsFloat64()
		if err != nil {
			cmd.SetErr(err)
			return cmd
		}
		lat, err := loc[1].AsFloat64()
		if err != nil {
			cmd.SetErr(err)
			return cmd
		}
		val = append(val, &GeoPos{
			Longitude: long,
			Latitude:  lat,
		})
	}
	cmd.SetVal(val)
	return cmd
}

type GeoLocationCmd struct {
	baseCmd[[]rueidis.GeoLocation]
}

func newGeoLocationCmd(res rueidis.RedisResult) *GeoLocationCmd {
	ret := &GeoLocationCmd{}
	ret.val, ret.err = res.AsGeosearch()
	return ret
}

type CommandInfo struct {
	Name        string
	Flags       []string
	ACLFlags    []string
	Arity       int64
	FirstKeyPos int64
	LastKeyPos  int64
	StepCount   int64
	ReadOnly    bool
}

type CommandsInfoCmd struct {
	baseCmd[map[string]CommandInfo]
}

func newCommandsInfoCmd(res rueidis.RedisResult) *CommandsInfoCmd {
	cmd := &CommandsInfoCmd{}
	arr, err := res.ToArray()
	if err != nil {
		cmd.SetErr(err)
		return cmd
	}
	val := make(map[string]CommandInfo, len(arr))
	for _, v := range arr {
		info, err := v.ToArray()
		if err != nil {
			cmd.SetErr(err)
			return cmd
		}
		if len(info) < 6 {
			cmd.SetErr(fmt.Errorf("got %d, wanted at least 6", len(info)))
			return cmd
		}
		var _cmd CommandInfo
		_cmd.Name, err = info[0].ToString()
		if err != nil {
			cmd.SetErr(err)
			return cmd
		}
		_cmd.Arity, err = info[1].AsInt64()
		if err != nil {
			cmd.SetErr(err)
			return cmd
		}
		_cmd.Flags, err = info[2].AsStrSlice()
		if err != nil {
			if rueidis.IsRedisNil(err) {
				_cmd.Flags = []string{}
			} else {
				cmd.SetErr(err)
				return cmd
			}
		}
		_cmd.FirstKeyPos, err = info[3].AsInt64()
		if err != nil {
			cmd.SetErr(err)
			return cmd
		}
		_cmd.LastKeyPos, err = info[4].AsInt64()
		if err != nil {
			cmd.SetErr(err)
			return cmd
		}
		_cmd.StepCount, err = info[5].AsInt64()
		if err != nil {
			cmd.SetErr(err)
			return cmd
		}
		for _, flag := range _cmd.Flags {
			if flag == "readonly" {
				_cmd.ReadOnly = true
				break
			}
		}
		if len(info) == 6 {
			val[_cmd.Name] = _cmd
			continue
		}
		_cmd.ACLFlags, err = info[6].AsStrSlice()
		if err != nil {
			if rueidis.IsRedisNil(err) {
				_cmd.ACLFlags = []string{}
			} else {
				cmd.SetErr(err)
				return cmd
			}
		}
		val[_cmd.Name] = _cmd
	}
	cmd.SetVal(val)
	return cmd
}

type Sort struct {
	By     string
	Order  string
	Get    []string
	Offset int64
	Count  int64
	Alpha  bool
}

// SetArgs provides arguments for the SetArgs function.
type SetArgs struct {
	ExpireAt time.Time
	Mode     string
	TTL      time.Duration
	Get      bool
	KeepTTL  bool
}

type BitCount struct {
	Start, End int64
	Unit       string // Stores BIT or BYTE
}

//type BitPos struct {
//	BitCount
//	Byte bool
//}

type BitFieldArg struct {
	Encoding string
	Offset   int64
}

type BitField struct {
	Get       *BitFieldArg
	Set       *BitFieldArg
	IncrBy    *BitFieldArg
	Overflow  string
	Increment int64
}

type LPosArgs struct {
	Rank, MaxLen int64
}

// Note: MaxLen/MaxLenApprox and MinID are in conflict, only one of them can be used.
type XAddArgs struct {
	Values     any
	Stream     string
	MinID      string
	ID         string
	MaxLen     int64
	Limit      int64
	NoMkStream bool
	Approx     bool
}

type XReadArgs struct {
	Streams []string // list of streams
	Count   int64
	Block   time.Duration
}

type XReadGroupArgs struct {
	Group    string
	Consumer string
	Streams  []string // list of streams
	Count    int64
	Block    time.Duration
	NoAck    bool
}

type XPendingExtArgs struct {
	Stream   string
	Group    string
	Start    string
	End      string
	Consumer string
	Idle     time.Duration
	Count    int64
}

type XClaimArgs struct {
	Stream   string
	Group    string
	Consumer string
	Messages []string
	MinIdle  time.Duration
}

type XAutoClaimArgs struct {
	Stream   string
	Group    string
	Start    string
	Consumer string
	MinIdle  time.Duration
	Count    int64
}

type XMessage struct {
	Values map[string]any
	ID     string
}

// Note: The GT, LT and NX options are mutually exclusive.
type ZAddArgs struct {
	Members []Z
	NX      bool
	XX      bool
	LT      bool
	GT      bool
	Ch      bool
}

// ZRangeArgs is all the options of the ZRange command.
// In version> 6.2.0, you can replace the(cmd):
//
//	ZREVRANGE,
//	ZRANGEBYSCORE,
//	ZREVRANGEBYSCORE,
//	ZRANGEBYLEX,
//	ZREVRANGEBYLEX.
//
// Please pay attention to your redis-server version.
//
// Rev, ByScore, ByLex and Offset+Count options require redis-server 6.2.0 and higher.
type ZRangeArgs struct {
	Start   any
	Stop    any
	Key     string
	Offset  int64
	Count   int64
	ByScore bool
	ByLex   bool
	Rev     bool
}

type ZRangeBy struct {
	Min, Max      string
	Offset, Count int64
}

type GeoLocation = rueidis.GeoLocation

// GeoRadiusQuery is used with GeoRadius to query geospatial index.
type GeoRadiusQuery struct {
	Unit        string
	Sort        string
	Store       string
	StoreDist   string
	Radius      float64
	Count       int64
	WithCoord   bool
	WithDist    bool
	WithGeoHash bool
}

// GeoSearchQuery is used for GEOSearch/GEOSearchStore command query.
type GeoSearchQuery struct {
	Member     string
	RadiusUnit string
	BoxUnit    string
	Sort       string
	Longitude  float64
	Latitude   float64
	Radius     float64
	BoxWidth   float64
	BoxHeight  float64
	Count      int64
	CountAny   bool
}

type GeoSearchLocationQuery struct {
	GeoSearchQuery

	WithCoord bool
	WithDist  bool
	WithHash  bool
}

type GeoSearchStoreQuery struct {
	GeoSearchQuery

	// When using the StoreDist option, the command stores the items in a
	// sorted set populated with their distance from the center of the circle or box,
	// as a floating-point number, in the same unit specified for that shape.
	StoreDist bool
}

func (q *GeoRadiusQuery) args() []string {
	args := make([]string, 0, 2)
	args = append(args, strconv.FormatFloat(q.Radius, 'f', -1, 64))
	if q.Unit != "" {
		args = append(args, q.Unit)
	} else {
		args = append(args, "km")
	}
	if q.WithCoord {
		args = append(args, "WITHCOORD")
	}
	if q.WithDist {
		args = append(args, "WITHDIST")
	}
	if q.WithGeoHash {
		args = append(args, "WITHHASH")
	}
	if q.Count > 0 {
		args = append(args, "COUNT", strconv.FormatInt(q.Count, 10))
	}
	if q.Sort != "" {
		args = append(args, q.Sort)
	}
	if q.Store != "" {
		args = append(args, "STORE")
		args = append(args, q.Store)
	}
	if q.StoreDist != "" {
		args = append(args, "STOREDIST")
		args = append(args, q.StoreDist)
	}
	return args
}

func (q *GeoSearchQuery) args() []string {
	args := make([]string, 0, 2)
	if q.Member != "" {
		args = append(args, "FROMMEMBER", q.Member)
	} else {
		args = append(args, "FROMLONLAT", strconv.FormatFloat(q.Longitude, 'f', -1, 64), strconv.FormatFloat(q.Latitude, 'f', -1, 64))
	}
	if q.Radius > 0 {
		if q.RadiusUnit == "" {
			q.RadiusUnit = "KM"
		}
		args = append(args, "BYRADIUS", strconv.FormatFloat(q.Radius, 'f', -1, 64), q.RadiusUnit)
	} else {
		if q.BoxUnit == "" {
			q.BoxUnit = "KM"
		}
		args = append(args, "BYBOX", strconv.FormatFloat(q.BoxWidth, 'f', -1, 64), strconv.FormatFloat(q.BoxHeight, 'f', -1, 64), q.BoxUnit)
	}
	if q.Sort != "" {
		args = append(args, q.Sort)
	}
	if q.Count > 0 {
		args = append(args, "COUNT", strconv.FormatInt(q.Count, 10))
		if q.CountAny {
			args = append(args, "ANY")
		}
	}
	return args
}

func (q *GeoSearchLocationQuery) args() []string {
	args := q.GeoSearchQuery.args()
	if q.WithCoord {
		args = append(args, "WITHCOORD")
	}
	if q.WithDist {
		args = append(args, "WITHDIST")
	}
	if q.WithHash {
		args = append(args, "WITHHASH")
	}
	return args
}

type Function struct {
	Name        string
	Description string
	Flags       []string
}

type Library struct {
	Name      string
	Engine    string
	Code      string
	Functions []Function
}

type FunctionListQuery struct {
	LibraryNamePattern string
	WithCode           bool
}

type FunctionListCmd struct {
	baseCmd[[]Library]
}

func newFunctionListCmd(res rueidis.RedisResult) *FunctionListCmd {
	cmd := &FunctionListCmd{}
	arr, err := res.ToArray()
	if err != nil {
		cmd.SetErr(err)
		return cmd
	}
	val := make([]Library, len(arr))
	for i := 0; i < len(arr); i++ {
		kv, _ := arr[i].AsMap()
		for k, v := range kv {
			switch k {
			case "library_name":
				val[i].Name, _ = v.ToString()
			case "engine":
				val[i].Engine, _ = v.ToString()
			case "library_code":
				val[i].Code, _ = v.ToString()
			case "functions":
				fns, _ := v.ToArray()
				val[i].Functions = make([]Function, len(fns))
				for j := 0; j < len(fns); j++ {
					fkv, _ := fns[j].AsMap()
					for k, v := range fkv {
						switch k {
						case "name":
							val[i].Functions[j].Name, _ = v.ToString()
						case "description":
							val[i].Functions[j].Description, _ = v.ToString()
						case "flags":
							val[i].Functions[j].Flags, _ = v.AsStrSlice()
						}
					}
				}
			}
		}
	}
	cmd.SetVal(val)
	return cmd
}

func (cmd *FunctionListCmd) First() (*Library, error) {
	if cmd.err != nil {
		return nil, cmd.err
	}
	if len(cmd.val) > 0 {
		return &cmd.val[0], nil
	}
	return nil, rueidis.Nil
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
		// too small, truncate too 1s
		return 1
	}
	return int64(dur / time.Second)
}

// https://github.com/redis/go-redis/blob/f994ff1cd96299a5c8029ae3403af7b17ef06e8a/gears_commands.go#L21C1-L35C2
type TFunctionLoadOptions struct {
	Config  string
	Replace bool
}

type TFunctionListOptions struct {
	Library  string
	Verbose  int
	Withcode bool
}

type TFCallOptions struct {
	Keys      []string
	Arguments []string
}

type MapStringInterfaceSliceCmd struct {
	baseCmd[[]map[string]any]
}

func newMapStringInterfaceSliceCmd(res rueidis.RedisResult) *MapStringInterfaceSliceCmd {
	cmd := &MapStringInterfaceSliceCmd{}
	arr, err := res.ToArray()
	if err != nil {
		cmd.SetErr(err)
		return cmd
	}
	cmd.val = make([]map[string]any, 0, len(arr))
	for _, ele := range arr {
		m, err := ele.AsMap()
		eleMap := make(map[string]any, len(m))
		if err != nil {
			cmd.SetErr(err)
			return cmd
		}
		for k, v := range m {
			var val any
			if !v.IsNil() {
				var err error
				val, err = v.ToAny()
				if err != nil {
					cmd.SetErr(err)
					return cmd
				}
			}
			eleMap[k] = val
		}
		cmd.val = append(cmd.val, eleMap)
	}
	return cmd
}

type BFInsertOptions struct {
	Capacity   int64
	Error      float64
	Expansion  int64
	NonScaling bool
	NoCreate   bool
}

type BFReserveOptions struct {
	Capacity   int64
	Error      float64
	Expansion  int64
	NonScaling bool
}

type CFReserveOptions struct {
	Capacity      int64
	BucketSize    int64
	MaxIterations int64
	Expansion     int64
}

type CFInsertOptions struct {
	Capacity int64
	NoCreate bool
}

type BFInfo struct {
	Capacity      int64 `redis:"Capacity"`
	Size          int64 `redis:"Size"`
	Filters       int64 `redis:"Number of filters"`
	ItemsInserted int64 `redis:"Number of items inserted"`
	ExpansionRate int64 `redis:"Expansion rate"`
}

type BFInfoCmd struct {
	baseCmd[BFInfo]
}

func newBFInfoCmd(res rueidis.RedisResult) *BFInfoCmd {
	cmd := &BFInfoCmd{}
	info := BFInfo{}
	if err := res.Error(); err != nil {
		cmd.err = err
		return cmd
	}
	m, err := res.AsIntMap()
	if err != nil {
		cmd.err = err
		return cmd
	}
	keys := make([]string, 0, len(m))
	values := make([]any, 0, len(m))
	for k, v := range m {
		keys = append(keys, k)
		values = append(values, strconv.FormatInt(v, 10))
	}
	if err := Scan(&info, keys, values); err != nil {
		cmd.err = err
		return cmd
	}
	cmd.SetVal(info)
	return cmd
}

type ScanDump struct {
	Data string
	Iter int64
}

type ScanDumpCmd struct {
	baseCmd[ScanDump]
}

func newScanDumpCmd(res rueidis.RedisResult) *ScanDumpCmd {
	cmd := &ScanDumpCmd{}
	scanDump := ScanDump{}
	if err := res.Error(); err != nil {
		cmd.err = err
		return cmd
	}
	arr, err := res.ToArray()
	if err != nil {
		cmd.err = err
		return cmd
	}
	if len(arr) != 2 {
		panic(fmt.Sprintf("wrong length of redis message, got %v, want %v", len(arr), 2))
	}
	iter, err := arr[0].AsInt64()
	if err != nil {
		cmd.err = err
		return cmd
	}
	data, err := arr[1].ToString()
	if err != nil {
		cmd.err = err
		return cmd
	}
	scanDump.Iter = iter
	scanDump.Data = data
	cmd.SetVal(scanDump)
	return cmd
}

type CFInfo struct {
	Size             int64 `redis:"Size"`
	NumBuckets       int64 `redis:"Number of buckets"`
	NumFilters       int64 `redis:"Number of filters"`
	NumItemsInserted int64 `redis:"Number of items inserted"`
	NumItemsDeleted  int64 `redis:"Number of items deleted"`
	BucketSize       int64 `redis:"Bucket size"`
	ExpansionRate    int64 `redis:"Expansion rate"`
	MaxIteration     int64 `redis:"Max iterations"`
}

type CFInfoCmd struct {
	baseCmd[CFInfo]
}

func newCFInfoCmd(res rueidis.RedisResult) *CFInfoCmd {
	cmd := &CFInfoCmd{}
	info := CFInfo{}
	m, err := res.AsMap()
	if err != nil {
		cmd.err = err
		return cmd
	}
	keys := make([]string, 0, len(m))
	values := make([]any, 0, len(m))
	for k, v := range m {
		keys = append(keys, k)
		val, err := v.AsInt64()
		if err != nil {
			cmd.err = err
			return cmd
		}
		values = append(values, strconv.FormatInt(val, 10))
	}
	if err := Scan(&info, keys, values); err != nil {
		cmd.err = err
		return cmd
	}
	cmd.SetVal(info)
	return cmd
}

type CMSInfo struct {
	Width int64 `redis:"width"`
	Depth int64 `redis:"depth"`
	Count int64 `redis:"count"`
}

type CMSInfoCmd struct {
	baseCmd[CMSInfo]
}

func newCMSInfoCmd(res rueidis.RedisResult) *CMSInfoCmd {
	cmd := &CMSInfoCmd{}
	info := CMSInfo{}
	m, err := res.AsIntMap()
	if err != nil {
		cmd.err = err
		return cmd
	}
	keys := make([]string, 0, len(m))
	values := make([]any, 0, len(m))
	for k, v := range m {
		keys = append(keys, k)
		values = append(values, strconv.FormatInt(v, 10))
	}
	if err := Scan(&info, keys, values); err != nil {
		cmd.err = err
		return cmd
	}
	cmd.SetVal(info)
	return cmd
}

type TopKInfo struct {
	K     int64   `redis:"k"`
	Width int64   `redis:"width"`
	Depth int64   `redis:"depth"`
	Decay float64 `redis:"decay"`
}

type TopKInfoCmd struct {
	baseCmd[TopKInfo]
}

func newTopKInfoCmd(res rueidis.RedisResult) *TopKInfoCmd {
	cmd := &TopKInfoCmd{}
	info := TopKInfo{}
	m, err := res.ToMap()
	if err != nil {
		cmd.err = err
		return cmd
	}
	keys := make([]string, 0, len(m))
	values := make([]any, 0, len(m))
	for k, v := range m {
		keys = append(keys, k)
		switch k {
		case "k", "width", "depth":
			intVal, err := v.AsInt64()
			if err != nil {
				cmd.err = err
				return cmd
			}
			values = append(values, strconv.FormatInt(intVal, 10))
		case "decay":
			decay, err := v.AsFloat64()
			if err != nil {
				cmd.err = err
				return cmd
			}
			// args of strconv.FormatFloat is copied from cmds.TopkReserveParamsDepth.Decay
			values = append(values, strconv.FormatFloat(decay, 'f', -1, 64))
		default:
			panic("unexpected key")
		}
	}
	if err := Scan(&info, keys, values); err != nil {
		cmd.err = err
		return cmd
	}
	cmd.SetVal(info)
	return cmd
}

type MapStringIntCmd struct {
	baseCmd[map[string]int64]
}

func newMapStringIntCmd(res rueidis.RedisResult) *MapStringIntCmd {
	cmd := &MapStringIntCmd{}
	m, err := res.AsIntMap()
	if err != nil {
		cmd.err = err
		return cmd
	}
	cmd.SetVal(m)
	return cmd
}

// Ref: https://redis.io/commands/tdigest.info/
type TDigestInfo struct {
	Compression       int64 `redis:"Compression"`
	Capacity          int64 `redis:"Capacity"`
	MergedNodes       int64 `redis:"Merged nodes"`
	UnmergedNodes     int64 `redis:"UnmergedNodes"`
	MergedWeight      int64 `redis:"MergedWeight"`
	UnmergedWeight    int64 `redis:"Unmerged weight"`
	Observations      int64 `redis:"Observations"`
	TotalCompressions int64 `redis:"Total compressions"`
	MemoryUsage       int64 `redis:"Memory usage"`
}

type TDigestInfoCmd struct {
	baseCmd[TDigestInfo]
}

func newTDigestInfoCmd(res rueidis.RedisResult) *TDigestInfoCmd {
	cmd := &TDigestInfoCmd{}
	info := TDigestInfo{}
	m, err := res.AsIntMap()
	if err != nil {
		cmd.err = err
		return cmd
	}
	keys := make([]string, 0, len(m))
	values := make([]any, 0, len(m))
	for k, v := range m {
		keys = append(keys, k)
		values = append(values, strconv.FormatInt(v, 10))
	}
	if err := Scan(&info, keys, values); err != nil {
		cmd.err = err
		return cmd
	}
	cmd.SetVal(info)
	return cmd
}

type TDigestMergeOptions struct {
	Compression int64
	Override    bool
}

type TSOptions struct {
	Retention       int
	ChunkSize       int
	Encoding        string
	DuplicatePolicy string
	Labels          map[string]string
}
type TSIncrDecrOptions struct {
	Timestamp    int64
	Retention    int
	ChunkSize    int
	Uncompressed bool
	Labels       map[string]string
}

type TSAlterOptions struct {
	Retention       int
	ChunkSize       int
	DuplicatePolicy string
	Labels          map[string]string
}

type TSCreateRuleOptions struct {
	alignTimestamp int64
}

type TSGetOptions struct {
	Latest bool
}

type TSInfoOptions struct {
	Debug bool
}
type Aggregator int

const (
	Invalid = Aggregator(iota)
	Avg
	Sum
	Min
	Max
	Range
	Count
	First
	Last
	StdP
	StdS
	VarP
	VarS
	Twa
)

func (a Aggregator) String() string {
	switch a {
	case Invalid:
		return ""
	case Avg:
		return "AVG"
	case Sum:
		return "SUM"
	case Min:
		return "MIN"
	case Max:
		return "MAX"
	case Range:
		return "RANGE"
	case Count:
		return "COUNT"
	case First:
		return "FIRST"
	case Last:
		return "LAST"
	case StdP:
		return "STD.P"
	case StdS:
		return "STD.S"
	case VarP:
		return "VAR.P"
	case VarS:
		return "VAR.S"
	case Twa:
		return "TWA"
	default:
		return ""
	}
}

type TSTimestampValue struct {
	Timestamp int64
	Value     float64
}
type TSTimestampValueCmd struct {
	baseCmd[TSTimestampValue]
}

func newTSTimestampValueCmd(res rueidis.RedisResult) *TSTimestampValueCmd {
	cmd := &TSTimestampValueCmd{}
	val := TSTimestampValue{}
	if err := res.Error(); err != nil {
		cmd.err = err
		return cmd
	}
	arr, err := res.ToArray()
	if err != nil {
		cmd.err = err
		return cmd
	}
	if len(arr) != 2 {
		panic(fmt.Sprintf("wrong len of array reply, should be 2, got %v", len(arr)))
	}
	val.Timestamp, err = arr[0].AsInt64()
	if err != nil {
		cmd.err = err
		return cmd
	}
	val.Value, err = arr[1].AsFloat64()
	if err != nil {
		cmd.err = err
		return cmd
	}
	cmd.SetVal(val)
	return cmd
}

type MapStringInterfaceCmd struct {
	baseCmd[map[string]any]
}

func newMapStringInterfaceCmd(res rueidis.RedisResult) *MapStringInterfaceCmd {
	cmd := &MapStringInterfaceCmd{}
	m, err := res.AsMap()
	if err != nil {
		cmd.err = err
		return cmd
	}
	strIntMap := make(map[string]any, len(m))
	for k, ele := range m {
		var v any
		var err error
		if !ele.IsNil() {
			v, err = ele.ToAny()
			if err != nil {
				cmd.err = err
				return cmd
			}
		}
		strIntMap[k] = v
	}
	cmd.SetVal(strIntMap)
	return cmd
}

type TSTimestampValueSliceCmd struct {
	baseCmd[[]TSTimestampValue]
}

func newTSTimestampValueSliceCmd(res rueidis.RedisResult) *TSTimestampValueSliceCmd {
	cmd := &TSTimestampValueSliceCmd{}
	msgSlice, err := res.ToArray()
	if err != nil {
		cmd.err = err
		return cmd
	}
	tsValSlice := make([]TSTimestampValue, 0, len(msgSlice))
	for i := 0; i < len(msgSlice); i++ {
		msgArray, err := msgSlice[i].ToArray()
		if err != nil {
			cmd.err = err
			return cmd
		}
		tstmp, err := msgArray[0].AsInt64()
		if err != nil {
			cmd.err = err
			return cmd
		}
		val, err := msgArray[1].AsFloat64()
		if err != nil {
			cmd.err = err
			return cmd
		}
		tsValSlice = append(tsValSlice, TSTimestampValue{Timestamp: tstmp, Value: val})
	}
	cmd.SetVal(tsValSlice)
	return cmd
}

type MapStringSliceInterfaceCmd struct {
	baseCmd[map[string][]any]
}

func newMapStringSliceInterfaceCmd(res rueidis.RedisResult) *MapStringSliceInterfaceCmd {
	cmd := &MapStringSliceInterfaceCmd{}
	m, err := res.ToMap()
	if err != nil {
		cmd.err = err
		return cmd
	}
	mapStrSliceInt := make(map[string][]any, len(m))
	for k, entry := range m {
		vals, err := entry.ToArray()
		if err != nil {
			cmd.err = err
			return cmd
		}
		anySlice := make([]any, 0, len(vals))
		for _, v := range vals {
			var err error
			ele, err := v.ToAny()
			if err != nil {
				cmd.err = err
				return cmd
			}
			anySlice = append(anySlice, ele)
		}
		mapStrSliceInt[k] = anySlice
	}
	cmd.SetVal(mapStrSliceInt)
	return cmd
}

type TSRangeOptions struct {
	Latest          bool
	FilterByTS      []int
	FilterByValue   []int
	Count           int
	Align           interface{}
	Aggregator      Aggregator
	BucketDuration  int
	BucketTimestamp interface{}
	Empty           bool
}

type TSRevRangeOptions struct {
	Latest          bool
	FilterByTS      []int
	FilterByValue   []int
	Count           int
	Align           interface{}
	Aggregator      Aggregator
	BucketDuration  int
	BucketTimestamp interface{}
	Empty           bool
}

type TSMRangeOptions struct {
	Latest          bool
	FilterByTS      []int
	FilterByValue   []int
	WithLabels      bool
	SelectedLabels  []interface{}
	Count           int
	Align           interface{}
	Aggregator      Aggregator
	BucketDuration  int
	BucketTimestamp interface{}
	Empty           bool
	GroupByLabel    interface{}
	Reducer         interface{}
}

type TSMRevRangeOptions struct {
	Latest          bool
	FilterByTS      []int
	FilterByValue   []int
	WithLabels      bool
	SelectedLabels  []interface{}
	Count           int
	Align           interface{}
	Aggregator      Aggregator
	BucketDuration  int
	BucketTimestamp interface{}
	Empty           bool
	GroupByLabel    interface{}
	Reducer         interface{}
}

type TSMGetOptions struct {
	Latest         bool
	WithLabels     bool
	SelectedLabels []interface{}
}

type JSONSetArgs struct {
	Key   string
	Path  string
	Value interface{}
}

type JSONArrIndexArgs struct {
	Start int
	Stop  *int
}

type JSONArrTrimArgs struct {
	Start int
	Stop  *int
}

type JSONCmd struct {
	baseCmd[string]
	expanded []any // expanded will be used at JSONCmd.Expanded
	typ      jsonCmdTyp
}

type jsonCmdTyp int

const (
	TYP_STRING jsonCmdTyp = iota
	TYP_ARRAY
)

// https://github.com/redis/go-redis/blob/v9.3.0/json.go#L86
func (cmd *JSONCmd) Val() string {
	return cmd.val
}

// https://github.com/redis/go-redis/blob/v9.3.0/json.go#L105

func (cmd *JSONCmd) Expanded() (any, error) {
	if cmd.typ == TYP_STRING {
		return cmd.Val(), nil
	}
	// TYP_ARRAY
	return cmd.expanded, nil
}

func (cmd *JSONCmd) Result() (string, error) {
	return cmd.Val(), nil
}

func newJSONCmd(res rueidis.RedisResult) *JSONCmd {
	cmd := &JSONCmd{}
	msg, err := res.ToMessage()
	if err != nil {
		cmd.SetErr(err)
		return cmd
	}
	switch {
	// JSON.GET
	case msg.IsString():
		cmd.typ = TYP_STRING
		str, err := res.ToString()
		if err != nil {
			cmd.SetErr(err)
			return cmd
		}
		cmd.SetVal(str)
	// JSON.NUMINCRBY
	case msg.IsArray():
		// we set marshaled string to cmd.val
		// which will be used at cmd.Val()
		// and also stored parsed result to cmd.expanded,
		// which will be used at cmd.Expanded()
		cmd.typ = TYP_ARRAY
		arr, err := res.ToArray()
		if err != nil {
			cmd.SetErr(err)
			return cmd
		}
		expanded := make([]any, len(arr))
		for i, e := range arr {
			anyE, err := e.ToAny()
			if err != nil {
				if err == rueidis.Nil {
					continue
				}
				cmd.SetErr(err)
				return cmd
			}
			expanded[i] = anyE
		}
		cmd.expanded = expanded
		val, err := json.Marshal(cmd.expanded)
		if err != nil {
			cmd.SetErr(err)
			return cmd
		}
		cmd.SetVal(string(val))
	default:
		panic("invalid type, expect array or string")
	}
	return cmd
}

type JSONGetArgs struct {
	Indent  string
	Newline string
	Space   string
}

type IntPointerSliceCmd struct {
	baseCmd[[]*int64]
}

// newIntPointerSliceCmd initialises an IntPointerSliceCmd
func newIntPointerSliceCmd(res rueidis.RedisResult) *IntPointerSliceCmd {
	cmd := &IntPointerSliceCmd{}
	arr, err := res.ToArray()
	if err != nil {
		cmd.SetErr(err)
		return cmd
	}
	intPtrSlice := make([]*int64, len(arr))
	for i, e := range arr {
		if e.IsNil() {
			continue
		}
		len, err := e.ToInt64()
		if err != nil {
			cmd.SetErr(err)
			return cmd
		}
		intPtrSlice[i] = &len
	}
	cmd.SetVal(intPtrSlice)
	return cmd
}

type JSONSliceCmd struct {
	baseCmd[[]any]
}

func newJSONSliceCmd(res rueidis.RedisResult) *JSONSliceCmd {
	cmd := &JSONSliceCmd{}
	arr, err := res.ToArray()
	if err != nil {
		cmd.SetErr(err)
		return cmd
	}
	anySlice := make([]any, len(arr))
	for i, e := range arr {
		if e.IsNil() {
			continue
		}
		anyE, err := e.ToAny()
		if err != nil {
			cmd.SetErr(err)
			return cmd
		}
		anySlice[i] = anyE
	}
	cmd.SetVal(anySlice)
	return cmd
}
