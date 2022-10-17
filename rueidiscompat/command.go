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
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/rueian/rueidis"
)

type Cmd struct {
	val interface{}
	err error
}

func newCmd(res rueidis.RedisResult) *Cmd {
	val, err := res.ToAny()
	return &Cmd{val: val, err: err}
}

func (cmd *Cmd) SetVal(val interface{}) {
	cmd.val = val
}

func (cmd *Cmd) SetErr(err error) {
	cmd.err = err
}

func (cmd *Cmd) Val() interface{} {
	return cmd.val
}

func (cmd *Cmd) Err() error {
	return cmd.err
}

func (cmd *Cmd) Result() (interface{}, error) {
	return cmd.val, cmd.err
}

func (cmd *Cmd) Text() (string, error) {
	if cmd.err != nil {
		return "", cmd.err
	}
	return toString(cmd.val)
}

func toString(val interface{}) (string, error) {
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

func toInt64(val interface{}) (int64, error) {
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

func toUint64(val interface{}) (uint64, error) {
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

func toFloat32(val interface{}) (float32, error) {
	switch val := val.(type) {
	case int64:
		return float32(val), nil
	case string:
		f, err := strconv.ParseFloat(val, 32)
		if err != nil {
			return 0, err
		}
		return float32(f), nil
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

func toFloat64(val interface{}) (float64, error) {
	switch val := val.(type) {
	case int64:
		return float64(val), nil
	case string:
		return strconv.ParseFloat(val, 64)
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

func toBool(val interface{}) (bool, error) {
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

func (cmd *Cmd) Slice() ([]interface{}, error) {
	if cmd.err != nil {
		return nil, cmd.err
	}
	switch val := cmd.val.(type) {
	case []interface{}:
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
	val string
	err error
}

func newStringCmd(res rueidis.RedisResult) *StringCmd {
	val, err := res.ToString()
	return &StringCmd{val: val, err: err}
}

func (cmd *StringCmd) SetVal(val string) {
	cmd.val = val
}

func (cmd *StringCmd) SetErr(err error) {
	cmd.err = err
}

func (cmd *StringCmd) Val() string {
	return cmd.val
}

func (cmd *StringCmd) Err() error {
	return cmd.err
}

func (cmd *StringCmd) Result() (string, error) {
	return cmd.val, cmd.err
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
	v, err := strconv.ParseFloat(cmd.Val(), 32)
	if err != nil {
		return 0, err
	}
	return float32(v), nil
}

func (cmd *StringCmd) Float64() (float64, error) {
	if cmd.err != nil {
		return 0, cmd.err
	}
	return strconv.ParseFloat(cmd.Val(), 64)
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
	val bool
	err error
}

func newBoolCmd(res rueidis.RedisResult) *BoolCmd {
	val, err := res.AsBool()
	if rueidis.IsRedisNil(err) {
		val = false
		err = nil
	}
	return &BoolCmd{val: val, err: err}
}

func (cmd *BoolCmd) SetVal(val bool) {
	cmd.val = val
}

func (cmd *BoolCmd) SetErr(err error) {
	cmd.err = err
}

func (cmd *BoolCmd) Val() bool {
	return cmd.val
}

func (cmd *BoolCmd) Err() error {
	return cmd.err
}

func (cmd *BoolCmd) Result() (bool, error) {
	return cmd.val, cmd.err
}

type IntCmd struct {
	val int64
	err error
}

func newIntCmd(res rueidis.RedisResult) *IntCmd {
	val, err := res.AsInt64()
	return &IntCmd{val: val, err: err}
}

func (cmd *IntCmd) SetVal(val int64) {
	cmd.val = val
}

func (cmd *IntCmd) SetErr(err error) {
	cmd.err = err
}

func (cmd *IntCmd) Val() int64 {
	return cmd.val
}

func (cmd *IntCmd) Err() error {
	return cmd.err
}

func (cmd *IntCmd) Result() (int64, error) {
	return cmd.val, cmd.err
}

func (cmd *IntCmd) Uint64() (uint64, error) {
	return uint64(cmd.val), cmd.err
}

type DurationCmd struct {
	val time.Duration
	err error
}

func newDurationCmd(res rueidis.RedisResult, precision time.Duration) *DurationCmd {
	val, err := res.AsInt64()
	if err != nil {
		return &DurationCmd{val: 0, err: err}
	}
	if val > 0 {
		return &DurationCmd{val: time.Duration(val) * precision, err: err}
	}
	return &DurationCmd{val: time.Duration(val), err: err}
}

func (cmd *DurationCmd) SetVal(val time.Duration) {
	cmd.val = val
}

func (cmd *DurationCmd) SetErr(err error) {
	cmd.err = err
}

func (cmd *DurationCmd) Val() time.Duration {
	return cmd.val
}

func (cmd *DurationCmd) Err() error {
	return cmd.err
}

func (cmd *DurationCmd) Result() (time.Duration, error) {
	return cmd.val, cmd.err
}

type StatusCmd struct {
	val string
	err error
}

func newStatusCmd(res rueidis.RedisResult) *StatusCmd {
	val, err := res.ToString()
	return &StatusCmd{val: val, err: err}
}

func (cmd *StatusCmd) SetVal(val string) {
	cmd.val = val
}

func (cmd *StatusCmd) SetErr(err error) {
	cmd.err = err
}

func (cmd *StatusCmd) Val() string {
	return cmd.val
}

func (cmd *StatusCmd) Err() error {
	return cmd.err
}

func (cmd *StatusCmd) Result() (string, error) {
	return cmd.val, cmd.err
}

type SliceCmd struct {
	val []interface{}
	err error
}

func newSliceCmd(res rueidis.RedisResult) *SliceCmd {
	val, err := res.ToArray()
	slice := &SliceCmd{val: make([]interface{}, len(val)), err: err}
	for i, v := range val {
		if s, err := v.ToString(); err == nil {
			slice.val[i] = s
		}
	}
	return slice
}

func newSliceCmdFromMap(res rueidis.RedisResult) *SliceCmd {
	val, err := res.AsStrMap()
	slice := &SliceCmd{val: make([]interface{}, 0, len(val)*2), err: err}
	for k, v := range val {
		slice.val = append(slice.val, k, v)
	}
	return slice
}

func (cmd *SliceCmd) SetVal(val []interface{}) {
	cmd.val = val
}

func (cmd *SliceCmd) SetErr(err error) {
	cmd.err = err
}

func (cmd *SliceCmd) Val() []interface{} {
	return cmd.val
}

func (cmd *SliceCmd) Err() error {
	return cmd.err
}

func (cmd *SliceCmd) Result() ([]interface{}, error) {
	return cmd.val, cmd.err
}

type StringSliceCmd struct {
	val []string
	err error
}

func newStringSliceCmd(res rueidis.RedisResult) *StringSliceCmd {
	val, err := res.AsStrSlice()
	return &StringSliceCmd{val: val, err: err}
}

func flattenStringSliceCmd(res rueidis.RedisResult) *StringSliceCmd {
	arr, err := res.ToArray()
	if err != nil {
		return &StringSliceCmd{err: err}
	}
	val := make([]string, 0, len(arr)*2)
	for _, v := range arr {
		s, err := v.AsStrSlice()
		if err != nil {
			return &StringSliceCmd{err: err}
		}
		val = append(val, s...)
	}
	return &StringSliceCmd{val: val, err: err}
}

func (cmd *StringSliceCmd) SetVal(val []string) {
	cmd.val = val
}

func (cmd *StringSliceCmd) SetErr(err error) {
	cmd.err = err
}

func (cmd *StringSliceCmd) Val() []string {
	return cmd.val
}

func (cmd *StringSliceCmd) Err() error {
	return cmd.err
}

func (cmd *StringSliceCmd) Result() ([]string, error) {
	return cmd.val, cmd.err
}

type IntSliceCmd struct {
	val []int64
	err error
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
	val []bool
	err error
}

func newBoolSliceCmd(res rueidis.RedisResult) *BoolSliceCmd {
	ints, err := res.AsIntSlice()
	if err != nil {
		return &BoolSliceCmd{err: err}
	}
	val := make([]bool, 0, len(ints))
	for _, i := range ints {
		val = append(val, i == 1)
	}
	return &BoolSliceCmd{val: val, err: err}
}

func (cmd *BoolSliceCmd) SetVal(val []bool) {
	cmd.val = val
}

func (cmd *BoolSliceCmd) SetErr(err error) {
	cmd.err = err
}

func (cmd *BoolSliceCmd) Val() []bool {
	return cmd.val
}

func (cmd *BoolSliceCmd) Err() error {
	return cmd.err
}

func (cmd *BoolSliceCmd) Result() ([]bool, error) {
	return cmd.val, cmd.err
}

type FloatSliceCmd struct {
	val []float64
	err error
}

func newFloatSliceCmd(res rueidis.RedisResult) *FloatSliceCmd {
	val, err := res.AsFloatSlice()
	return &FloatSliceCmd{val: val, err: err}
}

func (cmd *FloatSliceCmd) SetVal(val []float64) {
	cmd.val = val
}

func (cmd *FloatSliceCmd) SetErr(err error) {
	cmd.err = err
}

func (cmd *FloatSliceCmd) Val() []float64 {
	return cmd.val
}

func (cmd *FloatSliceCmd) Err() error {
	return cmd.err
}

func (cmd *FloatSliceCmd) Result() ([]float64, error) {
	return cmd.val, cmd.err
}

type ZSliceCmd struct {
	val []Z
	err error
}

func newZSliceCmd(res rueidis.RedisResult) *ZSliceCmd {
	scores, err := res.AsZScores()
	if err != nil {
		return &ZSliceCmd{err: err}
	}
	val := make([]Z, 0, len(scores))
	for _, s := range scores {
		val = append(val, Z{Member: s.Member, Score: s.Score})
	}
	return &ZSliceCmd{val: val}
}

func newZSliceSingleCmd(res rueidis.RedisResult) *ZSliceCmd {
	s, err := res.AsZScore()
	if err != nil {
		return &ZSliceCmd{err: err}
	}
	return &ZSliceCmd{val: []Z{{Member: s.Member, Score: s.Score}}, err: err}
}

func (cmd *ZSliceCmd) SetVal(val []Z) {
	cmd.val = val
}

func (cmd *ZSliceCmd) SetErr(err error) {
	cmd.err = err
}

func (cmd *ZSliceCmd) Val() []Z {
	return cmd.val
}

func (cmd *ZSliceCmd) Err() error {
	return cmd.err
}

func (cmd *ZSliceCmd) Result() ([]Z, error) {
	return cmd.val, cmd.err
}

type FloatCmd struct {
	val float64
	err error
}

func newFloatCmd(res rueidis.RedisResult) *FloatCmd {
	val, err := res.AsFloat64()
	return &FloatCmd{val: val, err: err}
}

func (cmd *FloatCmd) SetVal(val float64) {
	cmd.val = val
}

func (cmd *FloatCmd) SetErr(err error) {
	cmd.err = err
}

func (cmd *FloatCmd) Val() float64 {
	return cmd.val
}

func (cmd *FloatCmd) Err() error {
	return cmd.err
}

func (cmd *FloatCmd) Result() (float64, error) {
	return cmd.val, cmd.err
}

type ScanCmd struct {
	cursor uint64
	keys   []string
	err    error
}

func newScanCmd(res rueidis.RedisResult) *ScanCmd {
	ret, err := res.ToArray()
	if err != nil {
		return &ScanCmd{err: err}
	}
	cursor, err := ret[0].AsInt64()
	if err != nil {
		return &ScanCmd{err: err}
	}
	keys, err := ret[1].AsStrSlice()
	return &ScanCmd{cursor: uint64(cursor), keys: keys, err: err}
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

type StringStringMapCmd struct {
	val map[string]string
	err error
}

func newStringStringMapCmd(res rueidis.RedisResult) *StringStringMapCmd {
	val, err := res.AsStrMap()
	return &StringStringMapCmd{val: val, err: err}
}

func (cmd *StringStringMapCmd) SetVal(val map[string]string) {
	cmd.val = val
}

func (cmd *StringStringMapCmd) SetErr(err error) {
	cmd.err = err
}

func (cmd *StringStringMapCmd) Val() map[string]string {
	return cmd.val
}

func (cmd *StringStringMapCmd) Err() error {
	return cmd.err
}

func (cmd *StringStringMapCmd) Result() (map[string]string, error) {
	return cmd.val, cmd.err
}

type StringIntMapCmd struct {
	val map[string]int64
	err error
}

func newStringIntMapCmd(res rueidis.RedisResult) *StringIntMapCmd {
	val, err := res.AsIntMap()
	return &StringIntMapCmd{val: val, err: err}
}

func (cmd *StringIntMapCmd) SetVal(val map[string]int64) {
	cmd.val = val
}

func (cmd *StringIntMapCmd) SetErr(err error) {
	cmd.err = err
}

func (cmd *StringIntMapCmd) Val() map[string]int64 {
	return cmd.val
}

func (cmd *StringIntMapCmd) Err() error {
	return cmd.err
}

func (cmd *StringIntMapCmd) Result() (map[string]int64, error) {
	return cmd.val, cmd.err
}

type StringStructMapCmd struct {
	val map[string]struct{}
	err error
}

func newStringStructMapCmd(res rueidis.RedisResult) *StringStructMapCmd {
	strSlice, err := res.AsStrSlice()
	if err != nil {
		return &StringStructMapCmd{err: err}
	}
	val := make(map[string]struct{}, len(strSlice))
	for _, v := range strSlice {
		val[v] = struct{}{}
	}
	return &StringStructMapCmd{val: val, err: err}
}

func (cmd *StringStructMapCmd) SetVal(val map[string]struct{}) {
	cmd.val = val
}

func (cmd *StringStructMapCmd) SetErr(err error) {
	cmd.err = err
}

func (cmd *StringStructMapCmd) Val() map[string]struct{} {
	return cmd.val
}

func (cmd *StringStructMapCmd) Err() error {
	return cmd.err
}

func (cmd *StringStructMapCmd) Result() (map[string]struct{}, error) {
	return cmd.val, cmd.err
}

type XMessageSliceCmd struct {
	val []XMessage
	err error
}

func newXMessageSliceCmd(res rueidis.RedisResult) *XMessageSliceCmd {
	val, err := res.AsXRange()
	slice := &XMessageSliceCmd{val: make([]XMessage, len(val)), err: err}
	for i, r := range val {
		slice.val[i] = newXMessage(r)
	}
	return slice
}

func newXMessage(r rueidis.XRangeEntry) XMessage {
	if r.FieldValues == nil {
		return XMessage{ID: r.ID, Values: nil}
	}
	m := XMessage{ID: r.ID, Values: make(map[string]interface{}, len(r.FieldValues))}
	for k, v := range r.FieldValues {
		m.Values[k] = v
	}
	return m
}

func (cmd *XMessageSliceCmd) SetVal(val []XMessage) {
	cmd.val = val
}

func (cmd *XMessageSliceCmd) SetErr(err error) {
	cmd.err = err
}

func (cmd *XMessageSliceCmd) Val() []XMessage {
	return cmd.val
}

func (cmd *XMessageSliceCmd) Err() error {
	return cmd.err
}

func (cmd *XMessageSliceCmd) Result() ([]XMessage, error) {
	return cmd.val, cmd.err
}

type XStream struct {
	Stream   string
	Messages []XMessage
}

type XStreamSliceCmd struct {
	val []XStream
	err error
}

func newXStreamSliceCmd(res rueidis.RedisResult) *XStreamSliceCmd {
	streams, err := res.AsXRead()
	if err != nil {
		return &XStreamSliceCmd{err: err}
	}
	val := make([]XStream, 0, len(streams))
	for name, messages := range streams {
		msgs := make([]XMessage, 0, len(messages))
		for _, r := range messages {
			msgs = append(msgs, newXMessage(r))
		}
		val = append(val, XStream{Stream: name, Messages: msgs})
	}
	return &XStreamSliceCmd{val: val, err: err}
}

func (cmd *XStreamSliceCmd) SetVal(val []XStream) {
	cmd.val = val
}

func (cmd *XStreamSliceCmd) SetErr(err error) {
	cmd.err = err
}

func (cmd *XStreamSliceCmd) Val() []XStream {
	return cmd.val
}

func (cmd *XStreamSliceCmd) Err() error {
	return cmd.err
}

func (cmd *XStreamSliceCmd) Result() ([]XStream, error) {
	return cmd.val, cmd.err
}

type XPending struct {
	Count     int64
	Lower     string
	Higher    string
	Consumers map[string]int64
}

type XPendingCmd struct {
	val XPending
	err error
}

func newXPendingCmd(res rueidis.RedisResult) *XPendingCmd {
	arr, err := res.ToArray()
	if err != nil {
		return &XPendingCmd{err: err}
	}
	if len(arr) < 4 {
		return &XPendingCmd{err: fmt.Errorf("got %d, wanted 4", len(arr))}
	}
	count, err := arr[0].AsInt64()
	if err != nil {
		return &XPendingCmd{err: err}
	}
	lower, err := arr[1].ToString()
	if err != nil {
		return &XPendingCmd{err: err}
	}
	higher, err := arr[2].ToString()
	if err != nil {
		return &XPendingCmd{err: err}
	}
	val := XPending{
		Count:  count,
		Lower:  lower,
		Higher: higher,
	}
	consumerArr, err := arr[3].ToArray()
	if err != nil {
		return &XPendingCmd{err: err}
	}
	for _, v := range consumerArr {
		consumer, err := v.ToArray()
		if err != nil {
			return &XPendingCmd{err: err}
		}
		if len(consumer) < 2 {
			return &XPendingCmd{err: fmt.Errorf("got %d, wanted 2", len(arr))}
		}
		consumerName, err := consumer[0].ToString()
		if err != nil {
			return &XPendingCmd{err: err}
		}
		consumerPending, err := consumer[1].AsInt64()
		if err != nil {
			return &XPendingCmd{err: err}
		}
		if val.Consumers == nil {
			val.Consumers = make(map[string]int64)
		}
		val.Consumers[consumerName] = consumerPending
	}
	return &XPendingCmd{val: val, err: err}
}

func (cmd *XPendingCmd) SetVal(val XPending) {
	cmd.val = val
}

func (cmd *XPendingCmd) SetErr(err error) {
	cmd.err = err
}

func (cmd *XPendingCmd) Val() XPending {
	return cmd.val
}

func (cmd *XPendingCmd) Err() error {
	return cmd.err
}

func (cmd *XPendingCmd) Result() (XPending, error) {
	return cmd.val, cmd.err
}

type XPendingExt struct {
	ID         string
	Consumer   string
	Idle       time.Duration
	RetryCount int64
}

type XPendingExtCmd struct {
	val []XPendingExt
	err error
}

func newXPendingExtCmd(res rueidis.RedisResult) *XPendingExtCmd {
	arrs, err := res.ToArray()
	if err != nil {
		return &XPendingExtCmd{err: err}
	}
	val := make([]XPendingExt, 0, len(arrs))
	for _, v := range arrs {
		arr, err := v.ToArray()
		if err != nil {
			return &XPendingExtCmd{err: err}
		}
		if len(arr) < 4 {
			return &XPendingExtCmd{err: fmt.Errorf("got %d, wanted 4", len(arr))}
		}
		id, err := arr[0].ToString()
		if err != nil {
			return &XPendingExtCmd{err: err}
		}
		consumer, err := arr[1].ToString()
		if err != nil {
			return &XPendingExtCmd{err: err}
		}
		idle, err := arr[2].AsInt64()
		if err != nil {
			return &XPendingExtCmd{err: err}
		}
		retryCount, err := arr[3].AsInt64()
		if err != nil {
			return &XPendingExtCmd{err: err}
		}
		val = append(val, XPendingExt{
			ID:         id,
			Consumer:   consumer,
			Idle:       time.Duration(idle) * time.Millisecond,
			RetryCount: retryCount,
		})
	}
	return &XPendingExtCmd{val: val, err: err}
}

func (cmd *XPendingExtCmd) SetVal(val []XPendingExt) {
	cmd.val = val
}

func (cmd *XPendingExtCmd) SetErr(err error) {
	cmd.err = err
}

func (cmd *XPendingExtCmd) Val() []XPendingExt {
	return cmd.val
}

func (cmd *XPendingExtCmd) Err() error {
	return cmd.err
}

func (cmd *XPendingExtCmd) Result() ([]XPendingExt, error) {
	return cmd.val, cmd.err
}

type XAutoClaimCmd struct {
	start string
	val   []XMessage
	err   error
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
	start string
	val   []string
	err   error
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
	Consumers       int64
	Pending         int64
	EntriesRead     int64
	Lag             int64
	LastDeliveredID string
}

type XInfoGroupsCmd struct {
	val []XInfoGroup
	err error
}

func newXInfoGroupsCmd(res rueidis.RedisResult) *XInfoGroupsCmd {
	arr, err := res.ToArray()
	if err != nil {
		return &XInfoGroupsCmd{err: err}
	}
	groupInfos := make([]XInfoGroup, 0, len(arr))
	for _, v := range arr {
		info, err := v.AsMap()
		if err != nil {
			return &XInfoGroupsCmd{err: err}
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
	return &XInfoGroupsCmd{val: groupInfos, err: err}
}

func (cmd *XInfoGroupsCmd) SetVal(val []XInfoGroup) {
	cmd.val = val
}

func (cmd *XInfoGroupsCmd) SetErr(err error) {
	cmd.err = err
}

func (cmd *XInfoGroupsCmd) Val() []XInfoGroup {
	return cmd.val
}

func (cmd *XInfoGroupsCmd) Err() error {
	return cmd.err
}

func (cmd *XInfoGroupsCmd) Result() ([]XInfoGroup, error) {
	return cmd.val, cmd.err
}

type XInfoStream struct {
	Length               int64
	RadixTreeKeys        int64
	RadixTreeNodes       int64
	Groups               int64
	LastGeneratedID      string
	MaxDeletedEntryID    string
	EntriesAdded         int64
	FirstEntry           XMessage
	LastEntry            XMessage
	RecordedFirstEntryID string
}
type XInfoStreamCmd struct {
	val XInfoStream
	err error
}

func newXInfoStreamCmd(res rueidis.RedisResult) *XInfoStreamCmd {
	kv, err := res.AsMap()
	if err != nil {
		return &XInfoStreamCmd{err: err}
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
	return &XInfoStreamCmd{val: val, err: err}
}

func (cmd *XInfoStreamCmd) SetVal(val XInfoStream) {
	cmd.val = val
}

func (cmd *XInfoStreamCmd) SetErr(err error) {
	cmd.err = err
}

func (cmd *XInfoStreamCmd) Val() XInfoStream {
	return cmd.val
}

func (cmd *XInfoStreamCmd) Err() error {
	return cmd.err
}

func (cmd *XInfoStreamCmd) Result() (XInfoStream, error) {
	return cmd.val, cmd.err
}

type XInfoStreamConsumerPending struct {
	ID            string
	DeliveryTime  time.Time
	DeliveryCount int64
}

type XInfoStreamGroupPending struct {
	ID            string
	Consumer      string
	DeliveryTime  time.Time
	DeliveryCount int64
}

type XInfoStreamConsumer struct {
	Name     string
	SeenTime time.Time
	PelCount int64
	Pending  []XInfoStreamConsumerPending
}

type XInfoStreamGroup struct {
	Name            string
	LastDeliveredID string
	EntriesRead     int64
	Lag             int64
	PelCount        int64
	Pending         []XInfoStreamGroupPending
	Consumers       []XInfoStreamConsumer
}

type XInfoStreamFull struct {
	Length               int64
	RadixTreeKeys        int64
	RadixTreeNodes       int64
	LastGeneratedID      string
	MaxDeletedEntryID    string
	EntriesAdded         int64
	Entries              []XMessage
	Groups               []XInfoStreamGroup
	RecordedFirstEntryID string
}

type XInfoStreamFullCmd struct {
	val XInfoStreamFull
	err error
}

func newXInfoStreamFullCmd(res rueidis.RedisResult) *XInfoStreamFullCmd {
	kv, err := res.AsMap()
	if err != nil {
		return &XInfoStreamFullCmd{err: err}
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
			return &XInfoStreamFullCmd{err: err}
		}
	}
	if v, ok := kv["entries"]; ok {
		ranges, err := v.AsXRange()
		if err != nil {
			return &XInfoStreamFullCmd{err: err}
		}
		val.Entries = make([]XMessage, 0, len(ranges))
		for _, r := range ranges {
			val.Entries = append(val.Entries, newXMessage(r))
		}
	}
	return &XInfoStreamFullCmd{val: val, err: err}
}

func (cmd *XInfoStreamFullCmd) SetVal(val XInfoStreamFull) {
	cmd.val = val
}

func (cmd *XInfoStreamFullCmd) SetErr(err error) {
	cmd.err = err
}

func (cmd *XInfoStreamFullCmd) Val() XInfoStreamFull {
	return cmd.val
}

func (cmd *XInfoStreamFullCmd) Err() error {
	return cmd.err
}

func (cmd *XInfoStreamFullCmd) Result() (XInfoStreamFull, error) {
	return cmd.val, cmd.err
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
	val []XInfoConsumer
	err error
}

func newXInfoConsumersCmd(res rueidis.RedisResult) *XInfoConsumersCmd {
	arr, err := res.ToArray()
	if err != nil {
		return &XInfoConsumersCmd{err: err}
	}
	val := make([]XInfoConsumer, 0, len(arr))
	for _, v := range arr {
		info, err := v.AsMap()
		if err != nil {
			return &XInfoConsumersCmd{err: err}
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
	return &XInfoConsumersCmd{val: val, err: err}
}

func (cmd *XInfoConsumersCmd) SetVal(val []XInfoConsumer) {
	cmd.val = val
}

func (cmd *XInfoConsumersCmd) SetErr(err error) {
	cmd.err = err
}

func (cmd *XInfoConsumersCmd) Val() []XInfoConsumer {
	return cmd.val
}

func (cmd *XInfoConsumersCmd) Err() error {
	return cmd.err
}

func (cmd *XInfoConsumersCmd) Result() ([]XInfoConsumer, error) {
	return cmd.val, cmd.err
}

// Z represents sorted set member.
type Z struct {
	Score  float64
	Member interface{}
}

// ZWithKey represents sorted set member including the name of the key where it was popped.
type ZWithKey struct {
	Z
	Key string
}

// ZStore is used as an arg to ZInter/ZInterStore and ZUnion/ZUnionStore.
type ZStore struct {
	Keys    []string
	Weights []int64
	// Can be SUM, MIN or MAX.
	Aggregate string
}

type ZWithKeyCmd struct {
	val ZWithKey
	err error
}

func newZWithKeyCmd(res rueidis.RedisResult) *ZWithKeyCmd {
	arr, err := res.ToArray()
	if err != nil {
		return &ZWithKeyCmd{err: err}
	}
	if len(arr) < 3 {
		return &ZWithKeyCmd{err: fmt.Errorf("got %d, wanted 3", len(arr))}
	}
	val := ZWithKey{}
	val.Key, err = arr[0].ToString()
	if err != nil {
		return &ZWithKeyCmd{err: err}
	}
	val.Member, err = arr[1].ToString()
	if err != nil {
		return &ZWithKeyCmd{err: err}
	}
	val.Score, err = arr[2].AsFloat64()
	if err != nil {
		return &ZWithKeyCmd{err: err}
	}
	return &ZWithKeyCmd{val: val, err: err}
}

func (cmd *ZWithKeyCmd) SetVal(val ZWithKey) {
	cmd.val = val
}

func (cmd *ZWithKeyCmd) SetErr(err error) {
	cmd.err = err
}

func (cmd *ZWithKeyCmd) Val() ZWithKey {
	return cmd.val
}

func (cmd *ZWithKeyCmd) Err() error {
	return cmd.err
}

func (cmd *ZWithKeyCmd) Result() (ZWithKey, error) {
	return cmd.val, cmd.err
}

type TimeCmd struct {
	val time.Time
	err error
}

func newTimeCmd(res rueidis.RedisResult) *TimeCmd {
	arr, err := res.ToArray()
	if err != nil {
		return &TimeCmd{err: err}
	}
	if len(arr) < 2 {
		return &TimeCmd{err: fmt.Errorf("got %d, wanted 2", len(arr))}
	}
	sec, err := arr[0].AsInt64()
	if err != nil {
		return &TimeCmd{err: err}
	}
	microSec, err := arr[1].AsInt64()
	if err != nil {
		return &TimeCmd{err: err}
	}
	return &TimeCmd{val: time.Unix(sec, microSec*1000), err: err}
}

func (cmd *TimeCmd) SetVal(val time.Time) {
	cmd.val = val
}

func (cmd *TimeCmd) SetErr(err error) {
	cmd.err = err
}

func (cmd *TimeCmd) Val() time.Time {
	return cmd.val
}

func (cmd *TimeCmd) Err() error {
	return cmd.err
}

func (cmd *TimeCmd) Result() (time.Time, error) {
	return cmd.val, cmd.err
}

type ClusterNode struct {
	ID   string
	Addr string
}

type ClusterSlot struct {
	Start int64
	End   int64
	Nodes []ClusterNode
}

type ClusterSlotsCmd struct {
	val []ClusterSlot
	err error
}

func newClusterSlotsCmd(res rueidis.RedisResult) *ClusterSlotsCmd {
	arr, err := res.ToArray()
	if err != nil {
		return &ClusterSlotsCmd{err: err}
	}
	val := make([]ClusterSlot, 0, len(arr))
	for _, v := range arr {
		slot, err := v.ToArray()
		if err != nil {
			return &ClusterSlotsCmd{err: err}
		}
		if len(slot) < 2 {
			return &ClusterSlotsCmd{err: fmt.Errorf("got %d, excpected atleast 2", len(slot))}
		}
		start, err := slot[0].AsInt64()
		if err != nil {
			return &ClusterSlotsCmd{err: err}
		}
		end, err := slot[1].AsInt64()
		if err != nil {
			return &ClusterSlotsCmd{err: err}
		}
		nodes := make([]ClusterNode, len(slot)-2)
		for i, j := 2, 0; i < len(slot); i, j = i+1, j+1 {
			node, err := slot[i].ToArray()
			if err != nil {
				return &ClusterSlotsCmd{err: err}
			}
			if len(node) < 2 {
				return &ClusterSlotsCmd{err: fmt.Errorf("got %d, expected 2 or 3", len(node))}
			}
			ip, err := node[0].ToString()
			if err != nil {
				return &ClusterSlotsCmd{err: err}
			}
			port, err := node[1].AsInt64()
			if err != nil {
				return &ClusterSlotsCmd{err: err}
			}
			nodes[j].Addr = net.JoinHostPort(ip, strconv.FormatInt(port, 10))
			if len(node) > 2 {
				id, err := node[2].ToString()
				if err != nil {
					return &ClusterSlotsCmd{err: err}
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
	return &ClusterSlotsCmd{val: val, err: err}
}

func (cmd *ClusterSlotsCmd) SetVal(val []ClusterSlot) {
	cmd.val = val
}

func (cmd *ClusterSlotsCmd) SetErr(err error) {
	cmd.err = err
}

func (cmd *ClusterSlotsCmd) Val() []ClusterSlot {
	return cmd.val
}

func (cmd *ClusterSlotsCmd) Err() error {
	return cmd.err
}

func (cmd *ClusterSlotsCmd) Result() ([]ClusterSlot, error) {
	return cmd.val, cmd.err
}

type GeoPos struct {
	Longitude, Latitude float64
}

type GeoPosCmd struct {
	val []*GeoPos
	err error
}

func newGeoPosCmd(res rueidis.RedisResult) *GeoPosCmd {
	arr, err := res.ToArray()
	if err != nil {
		return &GeoPosCmd{err: err}
	}
	val := make([]*GeoPos, 0, len(arr))
	for _, v := range arr {
		loc, err := v.ToArray()
		if err != nil {
			if rueidis.IsRedisNil(err) {
				val = append(val, nil)
				continue
			}
			return &GeoPosCmd{err: err}
		}
		if len(loc) != 2 {
			return &GeoPosCmd{err: fmt.Errorf("got %d, expected 2", len(loc))}
		}
		long, err := loc[0].AsFloat64()
		if err != nil {
			return &GeoPosCmd{err: err}
		}
		lat, err := loc[1].AsFloat64()
		if err != nil {
			return &GeoPosCmd{err: err}
		}
		val = append(val, &GeoPos{
			Longitude: long,
			Latitude:  lat,
		})
	}
	return &GeoPosCmd{val: val, err: err}
}

func (cmd *GeoPosCmd) SetVal(val []*GeoPos) {
	cmd.val = val
}

func (cmd *GeoPosCmd) SetErr(err error) {
	cmd.err = err
}

func (cmd *GeoPosCmd) Val() []*GeoPos {
	return cmd.val
}

func (cmd *GeoPosCmd) Err() error {
	return cmd.err
}

func (cmd *GeoPosCmd) Result() ([]*GeoPos, error) {
	return cmd.val, cmd.err
}

type GeoLocationCmd struct {
	val []GeoLocation
	err error
}

func newGeoLocationCmd(res rueidis.RedisResult, withDist, withGeoHash, withCoord bool) *GeoLocationCmd {
	arr, err := res.ToArray()
	if err != nil {
		return &GeoLocationCmd{err: err}
	}
	val := make([]GeoLocation, 0, len(arr))
	if !withDist && !withGeoHash && !withCoord {
		for _, v := range arr {
			name, err := v.ToString()
			if err != nil {
				return &GeoLocationCmd{err: err}
			}
			val = append(val, GeoLocation{Name: name})
		}
		return &GeoLocationCmd{val: val, err: err}
	}
	for _, v := range arr {
		info, err := v.ToArray()
		if err != nil {
			return &GeoLocationCmd{err: err}
		}
		var loc GeoLocation
		var i int
		loc.Name, err = info[i].ToString()
		i++
		if err != nil {
			return &GeoLocationCmd{err: err}
		}
		if withDist {
			loc.Dist, err = info[i].AsFloat64()
			i++
			if err != nil {
				return &GeoLocationCmd{err: err}
			}
		}
		if withGeoHash {
			loc.GeoHash, err = info[i].AsInt64()
			i++
			if err != nil {
				return &GeoLocationCmd{err: err}
			}
		}
		if withCoord {
			cord, err := info[i].ToArray()
			if err != nil {
				return &GeoLocationCmd{err: err}
			}
			if len(cord) != 2 {
				return &GeoLocationCmd{err: fmt.Errorf("got %d, expected 2", len(info))}
			}
			loc.Longitude, err = cord[0].AsFloat64()
			if err != nil {
				return &GeoLocationCmd{err: err}
			}
			loc.Latitude, err = cord[1].AsFloat64()
			if err != nil {
				return &GeoLocationCmd{err: err}
			}
		}
		val = append(val, loc)
	}
	return &GeoLocationCmd{val: val, err: err}
}

func (cmd *GeoLocationCmd) SetVal(val []GeoLocation) {
	cmd.val = val
}

func (cmd *GeoLocationCmd) SetErr(err error) {
	cmd.err = err
}

func (cmd *GeoLocationCmd) Val() []GeoLocation {
	return cmd.val
}

func (cmd *GeoLocationCmd) Err() error {
	return cmd.err
}

func (cmd *GeoLocationCmd) Result() ([]GeoLocation, error) {
	return cmd.val, cmd.err
}

type CommandInfo struct {
	Flags       []string
	ACLFlags    []string
	Name        string
	Arity       int64
	FirstKeyPos int64
	LastKeyPos  int64
	StepCount   int64
	ReadOnly    bool
}

type CommandsInfoCmd struct {
	val map[string]CommandInfo
	err error
}

func newCommandsInfoCmd(res rueidis.RedisResult) *CommandsInfoCmd {
	arr, err := res.ToArray()
	if err != nil {
		return &CommandsInfoCmd{err: err}
	}
	val := make(map[string]CommandInfo, len(arr))
	for _, v := range arr {
		info, err := v.ToArray()
		if err != nil {
			return &CommandsInfoCmd{err: err}
		}
		if len(info) < 6 {
			return &CommandsInfoCmd{err: fmt.Errorf("got %d, wanted at least 6", len(info))}
		}
		var cmd CommandInfo
		cmd.Name, err = info[0].ToString()
		if err != nil {
			return &CommandsInfoCmd{err: err}
		}
		cmd.Arity, err = info[1].AsInt64()
		if err != nil {
			return &CommandsInfoCmd{err: err}
		}
		cmd.Flags, err = info[2].AsStrSlice()
		if err != nil {
			if rueidis.IsRedisNil(err) {
				cmd.Flags = []string{}
			} else {
				return &CommandsInfoCmd{err: err}
			}
		}
		cmd.FirstKeyPos, err = info[3].AsInt64()
		if err != nil {
			return &CommandsInfoCmd{err: err}
		}
		cmd.LastKeyPos, err = info[4].AsInt64()
		if err != nil {
			return &CommandsInfoCmd{err: err}
		}
		cmd.StepCount, err = info[5].AsInt64()
		if err != nil {
			return &CommandsInfoCmd{err: err}
		}
		for _, flag := range cmd.Flags {
			if flag == "readonly" {
				cmd.ReadOnly = true
				break
			}
		}
		if len(info) == 6 {
			val[cmd.Name] = cmd
			continue
		}
		cmd.ACLFlags, err = info[6].AsStrSlice()
		if err != nil {
			if rueidis.IsRedisNil(err) {
				cmd.ACLFlags = []string{}
			} else {
				return &CommandsInfoCmd{err: err}
			}
		}
		val[cmd.Name] = cmd
	}
	return &CommandsInfoCmd{val: val, err: err}
}

func (cmd *CommandsInfoCmd) SetVal(val map[string]CommandInfo) {
	cmd.val = val
}

func (cmd *CommandsInfoCmd) SetErr(err error) {
	cmd.err = err
}

func (cmd *CommandsInfoCmd) Val() map[string]CommandInfo {
	return cmd.val
}

func (cmd *CommandsInfoCmd) Err() error {
	return cmd.err
}

func (cmd *CommandsInfoCmd) Result() (map[string]CommandInfo, error) {
	return cmd.val, cmd.err
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
	Increment int64
	Overflow  string
}

type LPosArgs struct {
	Rank, MaxLen int64
}

// Note: MaxLen/MaxLenApprox and MinID are in conflict, only one of them can be used.
type XAddArgs struct {
	Stream     string
	NoMkStream bool
	MaxLen     int64 // MAXLEN N

	MinID string
	// Approx causes MaxLen and MinID to use "~" matcher (instead of "=").
	Approx bool
	Limit  int64
	ID     string
	Values interface{}
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

type XMessage struct {
	ID     string
	Values map[string]interface{}
}

// Note: The GT, LT and NX options are mutually exclusive.
type ZAddArgs struct {
	NX      bool
	XX      bool
	LT      bool
	GT      bool
	Ch      bool
	Members []Z
}

// ZRangeArgs is all the options of the ZRange command.
// In version> 6.2.0, you can replace the(cmd):
//		ZREVRANGE,
//		ZRANGEBYSCORE,
//		ZREVRANGEBYSCORE,
//		ZRANGEBYLEX,
//		ZREVRANGEBYLEX.
// Please pay attention to your redis-server version.
//
// Rev, ByScore, ByLex and Offset+Count options require redis-server 6.2.0 and higher.
type ZRangeArgs struct {
	Key string

	Start interface{}
	Stop  interface{}

	// The ByScore and ByLex options are mutually exclusive.
	ByScore bool
	ByLex   bool

	Rev bool

	// limit offset count.
	Offset int64
	Count  int64
}

type ZRangeBy struct {
	Min, Max      string
	Offset, Count int64
}

type GeoLocation struct {
	Name                      string
	Longitude, Latitude, Dist float64
	GeoHash                   int64
}

// GeoRadiusQuery is used with GeoRadius to query geospatial index.
type GeoRadiusQuery struct {
	Radius float64
	// Can be m, km, ft, or mi. Default is km.
	Unit        string
	WithCoord   bool
	WithDist    bool
	WithGeoHash bool
	Count       int64
	// Can be ASC or DESC. Default is no sort order.
	Sort      string
	Store     string
	StoreDist string
}

// GeoSearchQuery is used for GEOSearch/GEOSearchStore command query.
type GeoSearchQuery struct {
	Member string

	// Latitude and Longitude when using FromLonLat option.
	Longitude float64
	Latitude  float64

	// Distance and unit when using ByRadius option.
	// Can use m, km, ft, or mi. Default is km.
	Radius     float64
	RadiusUnit string

	// Height, width and unit when using ByBox option.
	// Can be m, km, ft, or mi. Default is km.
	BoxWidth  float64
	BoxHeight float64
	BoxUnit   string

	// Can be ASC or DESC. Default is no sort order.
	Sort     string
	Count    int64
	CountAny bool
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
