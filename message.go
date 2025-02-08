package rueidis

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"github.com/redis/rueidis/internal/util"
)

const messageStructSize = int(unsafe.Sizeof(RedisMessage{}))

// Nil represents a Redis Nil message
var Nil = &RedisError{typ: typeNull}

// ErrParse is a parse error that occurs when a Redis message cannot be parsed correctly.
var errParse = errors.New("rueidis: parse error")

// IsRedisNil is a handy method to check if error is a redis nil response.
// All redis nil response returns as an error.
func IsRedisNil(err error) bool {
	return err == Nil
}

// IsParseErr checks if the error is a parse error
func IsParseErr(err error) bool {
	return errors.Is(err, errParse)
}

// IsRedisBusyGroup checks if it is a redis BUSYGROUP message.
func IsRedisBusyGroup(err error) bool {
	if ret, yes := IsRedisErr(err); yes {
		return ret.IsBusyGroup()
	}
	return false
}

// IsRedisErr is a handy method to check if error is a redis ERR response.
func IsRedisErr(err error) (ret *RedisError, ok bool) {
	ret, ok = err.(*RedisError)
	return ret, ok && ret != Nil
}

// RedisError is an error response or a nil message from redis instance
type RedisError RedisMessage

func (r *RedisError) Error() string {
	if r.IsNil() {
		return "redis nil message"
	}
	return r.string
}

// IsNil checks if it is a redis nil message.
func (r *RedisError) IsNil() bool {
	return r.typ == typeNull
}

// IsMoved checks if it is a redis MOVED message and returns moved address.
func (r *RedisError) IsMoved() (addr string, ok bool) {
	if ok = strings.HasPrefix(r.string, "MOVED"); ok {
		addr = fixIPv6HostPort(strings.Split(r.string, " ")[2])
	}
	return
}

// IsAsk checks if it is a redis ASK message and returns ask address.
func (r *RedisError) IsAsk() (addr string, ok bool) {
	if ok = strings.HasPrefix(r.string, "ASK"); ok {
		addr = fixIPv6HostPort(strings.Split(r.string, " ")[2])
	}
	return
}

func fixIPv6HostPort(addr string) string {
	if strings.IndexByte(addr, '.') < 0 && len(addr) > 0 && addr[0] != '[' { // skip ipv4 and enclosed ipv6
		if i := strings.LastIndexByte(addr, ':'); i >= 0 {
			return net.JoinHostPort(addr[:i], addr[i+1:])
		}
	}
	return addr
}

// IsTryAgain checks if it is a redis TRYAGAIN message and returns ask address.
func (r *RedisError) IsTryAgain() bool {
	return strings.HasPrefix(r.string, "TRYAGAIN")
}

// IsLoading checks if it is a redis LOADING message
func (r *RedisError) IsLoading() bool {
	return strings.HasPrefix(r.string, "LOADING")
}

// IsClusterDown checks if it is a redis CLUSTERDOWN message and returns ask address.
func (r *RedisError) IsClusterDown() bool {
	return strings.HasPrefix(r.string, "CLUSTERDOWN")
}

// IsNoScript checks if it is a redis NOSCRIPT message.
func (r *RedisError) IsNoScript() bool {
	return strings.HasPrefix(r.string, "NOSCRIPT")
}

// IsBusyGroup checks if it is a redis BUSYGROUP message.
func (r *RedisError) IsBusyGroup() bool {
	return strings.HasPrefix(r.string, "BUSYGROUP")
}

func newResult(val RedisMessage, err error) RedisResult {
	return RedisResult{val: val, err: err}
}

func newErrResult(err error) RedisResult {
	return RedisResult{err: err}
}

// RedisResult is the return struct from Client.Do or Client.DoCache
// it contains either a redis response or an underlying error (ex. network timeout).
type RedisResult struct {
	err error
	val RedisMessage
}

// NonRedisError can be used to check if there is an underlying error (ex. network timeout).
func (r RedisResult) NonRedisError() error {
	return r.err
}

// Error returns either underlying error or redis error or nil
func (r RedisResult) Error() (err error) {
	if r.err != nil {
		err = r.err
	} else {
		err = r.val.Error()
	}
	return
}

// ToMessage retrieves the RedisMessage
func (r RedisResult) ToMessage() (v RedisMessage, err error) {
	if r.err != nil {
		err = r.err
	} else {
		err = r.val.Error()
	}
	return r.val, err
}

// ToInt64 delegates to RedisMessage.ToInt64
func (r RedisResult) ToInt64() (v int64, err error) {
	if r.err != nil {
		err = r.err
	} else {
		v, err = r.val.ToInt64()
	}
	return
}

// ToBool delegates to RedisMessage.ToBool
func (r RedisResult) ToBool() (v bool, err error) {
	if r.err != nil {
		err = r.err
	} else {
		v, err = r.val.ToBool()
	}
	return
}

// ToFloat64 delegates to RedisMessage.ToFloat64
func (r RedisResult) ToFloat64() (v float64, err error) {
	if r.err != nil {
		err = r.err
	} else {
		v, err = r.val.ToFloat64()
	}
	return
}

// ToString delegates to RedisMessage.ToString
func (r RedisResult) ToString() (v string, err error) {
	if r.err != nil {
		err = r.err
	} else {
		v, err = r.val.ToString()
	}
	return
}

// AsReader delegates to RedisMessage.AsReader
func (r RedisResult) AsReader() (v io.Reader, err error) {
	if r.err != nil {
		err = r.err
	} else {
		v, err = r.val.AsReader()
	}
	return
}

// AsBytes delegates to RedisMessage.AsBytes
func (r RedisResult) AsBytes() (v []byte, err error) {
	if r.err != nil {
		err = r.err
	} else {
		v, err = r.val.AsBytes()
	}
	return
}

// DecodeJSON delegates to RedisMessage.DecodeJSON
func (r RedisResult) DecodeJSON(v any) (err error) {
	if r.err != nil {
		err = r.err
	} else {
		err = r.val.DecodeJSON(v)
	}
	return
}

// AsInt64 delegates to RedisMessage.AsInt64
func (r RedisResult) AsInt64() (v int64, err error) {
	if r.err != nil {
		err = r.err
	} else {
		v, err = r.val.AsInt64()
	}
	return
}

// AsUint64 delegates to RedisMessage.AsUint64
func (r RedisResult) AsUint64() (v uint64, err error) {
	if r.err != nil {
		err = r.err
	} else {
		v, err = r.val.AsUint64()
	}
	return
}

// AsBool delegates to RedisMessage.AsBool
func (r RedisResult) AsBool() (v bool, err error) {
	if r.err != nil {
		err = r.err
	} else {
		v, err = r.val.AsBool()
	}
	return
}

// AsFloat64 delegates to RedisMessage.AsFloat64
func (r RedisResult) AsFloat64() (v float64, err error) {
	if r.err != nil {
		err = r.err
	} else {
		v, err = r.val.AsFloat64()
	}
	return
}

// ToArray delegates to RedisMessage.ToArray
func (r RedisResult) ToArray() (v []RedisMessage, err error) {
	if r.err != nil {
		err = r.err
	} else {
		v, err = r.val.ToArray()
	}
	return
}

// AsStrSlice delegates to RedisMessage.AsStrSlice
func (r RedisResult) AsStrSlice() (v []string, err error) {
	if r.err != nil {
		err = r.err
	} else {
		v, err = r.val.AsStrSlice()
	}
	return
}

// AsIntSlice delegates to RedisMessage.AsIntSlice
func (r RedisResult) AsIntSlice() (v []int64, err error) {
	if r.err != nil {
		err = r.err
	} else {
		v, err = r.val.AsIntSlice()
	}
	return
}

// AsFloatSlice delegates to RedisMessage.AsFloatSlice
func (r RedisResult) AsFloatSlice() (v []float64, err error) {
	if r.err != nil {
		err = r.err
	} else {
		v, err = r.val.AsFloatSlice()
	}
	return
}

// AsBoolSlice delegates to RedisMessage.AsBoolSlice
func (r RedisResult) AsBoolSlice() (v []bool, err error) {
	if r.err != nil {
		err = r.err
	} else {
		v, err = r.val.AsBoolSlice()
	}
	return
}

// AsXRangeEntry delegates to RedisMessage.AsXRangeEntry
func (r RedisResult) AsXRangeEntry() (v XRangeEntry, err error) {
	if r.err != nil {
		err = r.err
	} else {
		v, err = r.val.AsXRangeEntry()
	}
	return
}

// AsXRange delegates to RedisMessage.AsXRange
func (r RedisResult) AsXRange() (v []XRangeEntry, err error) {
	if r.err != nil {
		err = r.err
	} else {
		v, err = r.val.AsXRange()
	}
	return
}

// AsZScore delegates to RedisMessage.AsZScore
func (r RedisResult) AsZScore() (v ZScore, err error) {
	if r.err != nil {
		err = r.err
	} else {
		v, err = r.val.AsZScore()
	}
	return
}

// AsZScores delegates to RedisMessage.AsZScores
func (r RedisResult) AsZScores() (v []ZScore, err error) {
	if r.err != nil {
		err = r.err
	} else {
		v, err = r.val.AsZScores()
	}
	return
}

// AsXRead delegates to RedisMessage.AsXRead
func (r RedisResult) AsXRead() (v map[string][]XRangeEntry, err error) {
	if r.err != nil {
		err = r.err
	} else {
		v, err = r.val.AsXRead()
	}
	return
}

func (r RedisResult) AsLMPop() (v KeyValues, err error) {
	if r.err != nil {
		err = r.err
	} else {
		v, err = r.val.AsLMPop()
	}
	return
}

func (r RedisResult) AsZMPop() (v KeyZScores, err error) {
	if r.err != nil {
		err = r.err
	} else {
		v, err = r.val.AsZMPop()
	}
	return
}

func (r RedisResult) AsFtSearch() (total int64, docs []FtSearchDoc, err error) {
	if r.err != nil {
		err = r.err
	} else {
		total, docs, err = r.val.AsFtSearch()
	}
	return
}

func (r RedisResult) AsFtAggregate() (total int64, docs []map[string]string, err error) {
	if r.err != nil {
		err = r.err
	} else {
		total, docs, err = r.val.AsFtAggregate()
	}
	return
}

func (r RedisResult) AsFtAggregateCursor() (cursor, total int64, docs []map[string]string, err error) {
	if r.err != nil {
		err = r.err
	} else {
		cursor, total, docs, err = r.val.AsFtAggregateCursor()
	}
	return
}

func (r RedisResult) AsGeosearch() (locations []GeoLocation, err error) {
	if r.err != nil {
		err = r.err
	} else {
		locations, err = r.val.AsGeosearch()
	}
	return
}

// AsMap delegates to RedisMessage.AsMap
func (r RedisResult) AsMap() (v map[string]RedisMessage, err error) {
	if r.err != nil {
		err = r.err
	} else {
		v, err = r.val.AsMap()
	}
	return
}

// AsStrMap delegates to RedisMessage.AsStrMap
func (r RedisResult) AsStrMap() (v map[string]string, err error) {
	if r.err != nil {
		err = r.err
	} else {
		v, err = r.val.AsStrMap()
	}
	return
}

// AsIntMap delegates to RedisMessage.AsIntMap
func (r RedisResult) AsIntMap() (v map[string]int64, err error) {
	if r.err != nil {
		err = r.err
	} else {
		v, err = r.val.AsIntMap()
	}
	return
}

// AsScanEntry delegates to RedisMessage.AsScanEntry.
func (r RedisResult) AsScanEntry() (v ScanEntry, err error) {
	if r.err != nil {
		err = r.err
	} else {
		v, err = r.val.AsScanEntry()
	}
	return
}

// ToMap delegates to RedisMessage.ToMap
func (r RedisResult) ToMap() (v map[string]RedisMessage, err error) {
	if r.err != nil {
		err = r.err
	} else {
		v, err = r.val.ToMap()
	}
	return
}

// ToAny delegates to RedisMessage.ToAny
func (r RedisResult) ToAny() (v any, err error) {
	if r.err != nil {
		err = r.err
	} else {
		v, err = r.val.ToAny()
	}
	return
}

func (r RedisResult) AsClientInfo() (v *ClientInfo, err error) {
	if r.err != nil {
		err = r.err
	} else {
		v, err = r.val.AsClientInfo()
	}
	return
}

func (r RedisResult) AsACLLogEntry() (v []*ACLLogEntry, err error) {
	if r.err != nil {
		err = r.err
	} else {
		v, err = r.val.AsACLLogEntry()
	}
	return
}

// IsCacheHit delegates to RedisMessage.IsCacheHit
func (r RedisResult) IsCacheHit() bool {
	return r.val.IsCacheHit()
}

// CacheTTL delegates to RedisMessage.CacheTTL
func (r RedisResult) CacheTTL() int64 {
	return r.val.CacheTTL()
}

// CachePTTL delegates to RedisMessage.CachePTTL
func (r RedisResult) CachePTTL() int64 {
	return r.val.CachePTTL()
}

// CachePXAT delegates to RedisMessage.CachePXAT
func (r RedisResult) CachePXAT() int64 {
	return r.val.CachePXAT()
}

// String returns human-readable representation of RedisResult
func (r *RedisResult) String() string {
	v, _ := (*prettyRedisResult)(r).MarshalJSON()
	return string(v)
}

type prettyRedisResult RedisResult

// MarshalJSON implements json.Marshaler interface
func (r *prettyRedisResult) MarshalJSON() ([]byte, error) {
	type PrettyRedisResult struct {
		Message *prettyRedisMessage `json:"Message,omitempty"`
		Error   string              `json:"Error,omitempty"`
	}
	obj := PrettyRedisResult{}
	if r.err != nil {
		obj.Error = r.err.Error()
	} else {
		obj.Message = (*prettyRedisMessage)(&r.val)
	}
	return json.Marshal(obj)
}

// RedisMessage is a redis response message, it may be a nil response
type RedisMessage struct {
	attrs   *RedisMessage
	string  string
	values  []RedisMessage
	integer int64
	typ     byte
	ttl     [7]byte
}

func (m *RedisMessage) cachesize() int {
	n := 9 // typ (1) + length (8) TODO: can we use VarInt instead of fixed 8 bytes for length?
	switch m.typ {
	case typeInteger, typeNull, typeBool:
	case typeArray, typeMap, typeSet:
		for _, val := range m.values {
			n += val.cachesize()
		}
	default:
		n += len(m.string)
	}
	return n
}

func (m *RedisMessage) serialize(o *bytes.Buffer) {
	var buf [8]byte // TODO: can we use VarInt instead of fixed 8 bytes for length?
	o.WriteByte(m.typ)
	switch m.typ {
	case typeInteger, typeNull, typeBool:
		binary.BigEndian.PutUint64(buf[:], uint64(m.integer))
		o.Write(buf[:])
	case typeArray, typeMap, typeSet:
		binary.BigEndian.PutUint64(buf[:], uint64(len(m.values)))
		o.Write(buf[:])
		for _, val := range m.values {
			val.serialize(o)
		}
	default:
		binary.BigEndian.PutUint64(buf[:], uint64(len(m.string)))
		o.Write(buf[:])
		o.WriteString(m.string)
	}
}

var ErrCacheUnmarshal = errors.New("cache unmarshal error")

func (m *RedisMessage) unmarshalView(c int64, buf []byte) (int64, error) {
	var err error
	if int64(len(buf)) < c+9 {
		return 0, ErrCacheUnmarshal
	}
	m.typ = buf[c]
	c += 1
	m.integer = int64(binary.BigEndian.Uint64(buf[c : c+8]))
	c += 8 // TODO: can we use VarInt instead of fixed 8 bytes for length?
	switch m.typ {
	case typeInteger, typeNull, typeBool:
	case typeArray, typeMap, typeSet:
		m.values = make([]RedisMessage, m.integer)
		m.integer = 0
		for i := range m.values {
			if c, err = m.values[i].unmarshalView(c, buf); err != nil {
				break
			}
		}
	default:
		if int64(len(buf)) < c+m.integer {
			return 0, ErrCacheUnmarshal
		}
		m.string = BinaryString(buf[c : c+m.integer])
		c += m.integer
		m.integer = 0
	}
	return c, err
}

// CacheSize returns the buffer size needed by the CacheMarshal.
func (m *RedisMessage) CacheSize() int {
	return m.cachesize() + 7 // 7 for ttl
}

// CacheMarshal writes serialized RedisMessage to the provided buffer.
// If the provided buffer is nil, CacheMarshal will allocate one.
// Note that output format is not compatible with different client versions.
func (m *RedisMessage) CacheMarshal(buf []byte) []byte {
	if buf == nil {
		buf = make([]byte, 0, m.CacheSize())
	}
	o := bytes.NewBuffer(buf)
	o.Write(m.ttl[:7])
	m.serialize(o)
	return o.Bytes()
}

// CacheUnmarshalView construct the RedisMessage from the buffer produced by CacheMarshal.
// Note that the buffer can't be reused after CacheUnmarshalView since it uses unsafe.String on top of the buffer.
func (m *RedisMessage) CacheUnmarshalView(buf []byte) error {
	if len(buf) < 7 {
		return ErrCacheUnmarshal
	}
	copy(m.ttl[:7], buf[:7])
	if _, err := m.unmarshalView(7, buf); err != nil {
		return err
	}
	m.attrs = cacheMark
	return nil
}

// IsNil check if message is a redis nil response
func (m *RedisMessage) IsNil() bool {
	return m.typ == typeNull
}

// IsInt64 check if message is a redis RESP3 int response
func (m *RedisMessage) IsInt64() bool {
	return m.typ == typeInteger
}

// IsFloat64 check if message is a redis RESP3 double response
func (m *RedisMessage) IsFloat64() bool {
	return m.typ == typeFloat
}

// IsString check if message is a redis string response
func (m *RedisMessage) IsString() bool {
	return m.typ == typeBlobString || m.typ == typeSimpleString
}

// IsBool check if message is a redis RESP3 bool response
func (m *RedisMessage) IsBool() bool {
	return m.typ == typeBool
}

// IsArray check if message is a redis array response
func (m *RedisMessage) IsArray() bool {
	return m.typ == typeArray || m.typ == typeSet
}

// IsMap check if message is a redis RESP3 map response
func (m *RedisMessage) IsMap() bool {
	return m.typ == typeMap
}

// Error check if message is a redis error response, including nil response
func (m *RedisMessage) Error() error {
	if m.typ == typeNull {
		return Nil
	}
	if m.typ == typeSimpleErr || m.typ == typeBlobErr {
		// kvrocks: https://github.com/redis/rueidis/issues/152#issuecomment-1333923750
		mm := *m
		mm.string = strings.TrimPrefix(m.string, "ERR ")
		return (*RedisError)(&mm)
	}
	return nil
}

// ToString check if message is a redis string response, and return it
func (m *RedisMessage) ToString() (val string, err error) {
	if m.IsString() {
		return m.string, nil
	}
	if m.IsInt64() || m.values != nil {
		typ := m.typ
		return "", fmt.Errorf("%w: redis message type %s is not a string", errParse, typeNames[typ])
	}
	return m.string, m.Error()
}

// AsReader check if message is a redis string response and wrap it with the strings.NewReader
func (m *RedisMessage) AsReader() (reader io.Reader, err error) {
	str, err := m.ToString()
	if err != nil {
		return nil, err
	}
	return strings.NewReader(str), nil
}

// AsBytes check if message is a redis string response and return it as an immutable []byte
func (m *RedisMessage) AsBytes() (bs []byte, err error) {
	str, err := m.ToString()
	if err != nil {
		return nil, err
	}
	return unsafe.Slice(unsafe.StringData(str), len(str)), nil
}

// DecodeJSON check if message is a redis string response and treat it as json, then unmarshal it into provided value
func (m *RedisMessage) DecodeJSON(v any) (err error) {
	str, err := m.ToString()
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(strings.NewReader(str))
	return decoder.Decode(v)
}

// AsInt64 check if message is a redis string response, and parse it as int64
func (m *RedisMessage) AsInt64() (val int64, err error) {
	if m.IsInt64() {
		return m.integer, nil
	}
	v, err := m.ToString()
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(v, 10, 64)
}

// AsUint64 check if message is a redis string response, and parse it as uint64
func (m *RedisMessage) AsUint64() (val uint64, err error) {
	if m.IsInt64() {
		return uint64(m.integer), nil
	}
	v, err := m.ToString()
	if err != nil {
		return 0, err
	}
	return strconv.ParseUint(v, 10, 64)
}

// AsBool checks if message is non-nil redis response, and parses it as bool
func (m *RedisMessage) AsBool() (val bool, err error) {
	if err = m.Error(); err != nil {
		return
	}
	switch m.typ {
	case typeBlobString, typeSimpleString:
		val = m.string == "OK"
		return
	case typeInteger:
		val = m.integer != 0
		return
	case typeBool:
		val = m.integer == 1
		return
	default:
		typ := m.typ
		return false, fmt.Errorf("%w: redis message type %s is not a int, string or bool", errParse, typeNames[typ])
	}
}

// AsFloat64 check if message is a redis string response, and parse it as float64
func (m *RedisMessage) AsFloat64() (val float64, err error) {
	if m.IsFloat64() {
		return util.ToFloat64(m.string)
	}
	v, err := m.ToString()
	if err != nil {
		return 0, err
	}
	return util.ToFloat64(v)
}

// ToInt64 check if message is a redis RESP3 int response, and return it
func (m *RedisMessage) ToInt64() (val int64, err error) {
	if m.IsInt64() {
		return m.integer, nil
	}
	if err = m.Error(); err != nil {
		return 0, err
	}
	typ := m.typ
	return 0, fmt.Errorf("%w: redis message type %s is not a RESP3 int64", errParse, typeNames[typ])
}

// ToBool check if message is a redis RESP3 bool response, and return it
func (m *RedisMessage) ToBool() (val bool, err error) {
	if m.IsBool() {
		return m.integer == 1, nil
	}
	if err = m.Error(); err != nil {
		return false, err
	}
	typ := m.typ
	return false, fmt.Errorf("%w: redis message type %s is not a RESP3 bool", errParse, typeNames[typ])
}

// ToFloat64 check if message is a redis RESP3 double response, and return it
func (m *RedisMessage) ToFloat64() (val float64, err error) {
	if m.IsFloat64() {
		return util.ToFloat64(m.string)
	}
	if err = m.Error(); err != nil {
		return 0, err
	}
	typ := m.typ
	return 0, fmt.Errorf("%w: redis message type %s is not a RESP3 float64", errParse, typeNames[typ])
}

// ToArray check if message is a redis array/set response, and return it
func (m *RedisMessage) ToArray() ([]RedisMessage, error) {
	if m.IsArray() {
		return m.values, nil
	}
	if err := m.Error(); err != nil {
		return nil, err
	}
	typ := m.typ
	return nil, fmt.Errorf("%w: redis message type %s is not a array", errParse, typeNames[typ])
}

// AsStrSlice check if message is a redis array/set response, and convert to []string.
// redis nil element and other non string element will be present as zero.
func (m *RedisMessage) AsStrSlice() ([]string, error) {
	values, err := m.ToArray()
	if err != nil {
		return nil, err
	}
	s := make([]string, 0, len(values))
	for _, v := range values {
		s = append(s, v.string)
	}
	return s, nil
}

// AsIntSlice check if message is a redis array/set response, and convert to []int64.
// redis nil element and other non integer element will be present as zero.
func (m *RedisMessage) AsIntSlice() ([]int64, error) {
	values, err := m.ToArray()
	if err != nil {
		return nil, err
	}
	s := make([]int64, len(values))
	for i, v := range values {
		if len(v.string) != 0 {
			if s[i], err = strconv.ParseInt(v.string, 10, 64); err != nil {
				return nil, err
			}
		} else {
			s[i] = v.integer
		}
	}
	return s, nil
}

// AsFloatSlice check if message is a redis array/set response, and convert to []float64.
// redis nil element and other non float element will be present as zero.
func (m *RedisMessage) AsFloatSlice() ([]float64, error) {
	values, err := m.ToArray()
	if err != nil {
		return nil, err
	}
	s := make([]float64, len(values))
	for i, v := range values {
		if len(v.string) != 0 {
			if s[i], err = util.ToFloat64(v.string); err != nil {
				return nil, err
			}
		} else {
			s[i] = float64(v.integer)
		}
	}
	return s, nil
}

// AsBoolSlice checks if message is a redis array/set response, and converts it to []bool.
// Redis nil elements and other non-boolean elements will be represented as false.
func (m *RedisMessage) AsBoolSlice() ([]bool, error) {
	values, err := m.ToArray()
	if err != nil {
		return nil, err
	}
	s := make([]bool, len(values))
	for i, v := range values {
		s[i], _ = v.AsBool() // Ignore error, non-boolean values will be false
	}
	return s, nil
}

// XRangeEntry is the element type of both XRANGE and XREVRANGE command response array
type XRangeEntry struct {
	FieldValues map[string]string
	ID          string
}

// AsXRangeEntry check if message is a redis array/set response of length 2, and convert to XRangeEntry
func (m *RedisMessage) AsXRangeEntry() (XRangeEntry, error) {
	values, err := m.ToArray()
	if err != nil {
		return XRangeEntry{}, err
	}
	if len(values) != 2 {
		return XRangeEntry{}, fmt.Errorf("got %d, wanted 2", len(values))
	}
	id, err := values[0].ToString()
	if err != nil {
		return XRangeEntry{}, err
	}
	fieldValues, err := values[1].AsStrMap()
	if err != nil {
		if IsRedisNil(err) {
			return XRangeEntry{ID: id, FieldValues: nil}, nil
		}
		return XRangeEntry{}, err
	}
	return XRangeEntry{
		ID:          id,
		FieldValues: fieldValues,
	}, nil
}

// AsXRange check if message is a redis array/set response, and convert to []XRangeEntry
func (m *RedisMessage) AsXRange() ([]XRangeEntry, error) {
	values, err := m.ToArray()
	if err != nil {
		return nil, err
	}
	msgs := make([]XRangeEntry, 0, len(values))
	for _, v := range values {
		msg, err := v.AsXRangeEntry()
		if err != nil {
			return nil, err
		}
		msgs = append(msgs, msg)
	}
	return msgs, nil
}

// AsXRead converts XREAD/XREADGRUOP response to map[string][]XRangeEntry
func (m *RedisMessage) AsXRead() (ret map[string][]XRangeEntry, err error) {
	if err = m.Error(); err != nil {
		return nil, err
	}
	if m.IsMap() {
		ret = make(map[string][]XRangeEntry, len(m.values)/2)
		for i := 0; i < len(m.values); i += 2 {
			if ret[m.values[i].string], err = m.values[i+1].AsXRange(); err != nil {
				return nil, err
			}
		}
		return ret, nil
	}
	if m.IsArray() {
		ret = make(map[string][]XRangeEntry, len(m.values))
		for _, v := range m.values {
			if !v.IsArray() || len(v.values) != 2 {
				return nil, fmt.Errorf("got %d, wanted 2", len(v.values))
			}
			if ret[v.values[0].string], err = v.values[1].AsXRange(); err != nil {
				return nil, err
			}
		}
		return ret, nil
	}
	typ := m.typ
	return nil, fmt.Errorf("%w: redis message type %s is not a map/array/set", errParse, typeNames[typ])
}

// ZScore is the element type of ZRANGE WITHSCORES, ZDIFF WITHSCORES and ZPOPMAX command response
type ZScore struct {
	Member string
	Score  float64
}

func toZScore(values []RedisMessage) (s ZScore, err error) {
	if len(values) == 2 {
		if s.Member, err = values[0].ToString(); err == nil {
			s.Score, err = values[1].AsFloat64()
		}
		return s, err
	}
	return ZScore{}, fmt.Errorf("redis message is not a map/array/set or its length is not 2")
}

// AsZScore converts ZPOPMAX and ZPOPMIN command with count 1 response to a single ZScore
func (m *RedisMessage) AsZScore() (s ZScore, err error) {
	arr, err := m.ToArray()
	if err != nil {
		return s, err
	}
	return toZScore(arr)
}

// AsZScores converts ZRANGE WITHSCROES, ZDIFF WITHSCROES and ZPOPMAX/ZPOPMIN command with count > 1 responses to []ZScore
func (m *RedisMessage) AsZScores() ([]ZScore, error) {
	arr, err := m.ToArray()
	if err != nil {
		return nil, err
	}
	if len(arr) > 0 && arr[0].IsArray() {
		scores := make([]ZScore, len(arr))
		for i, v := range arr {
			if scores[i], err = toZScore(v.values); err != nil {
				return nil, err
			}
		}
		return scores, nil
	}
	scores := make([]ZScore, len(arr)/2)
	for i := 0; i < len(scores); i++ {
		j := i * 2
		if scores[i], err = toZScore(arr[j : j+2]); err != nil {
			return nil, err
		}
	}
	return scores, nil
}

// ScanEntry is the element type of both SCAN, SSCAN, HSCAN and ZSCAN command response.
type ScanEntry struct {
	Elements []string
	Cursor   uint64
}

// AsScanEntry check if message is a redis array/set response of length 2, and convert to ScanEntry.
func (m *RedisMessage) AsScanEntry() (e ScanEntry, err error) {
	msgs, err := m.ToArray()
	if err != nil {
		return ScanEntry{}, err
	}
	if len(msgs) >= 2 {
		if e.Cursor, err = msgs[0].AsUint64(); err == nil {
			e.Elements, err = msgs[1].AsStrSlice()
		}
		return e, err
	}
	typ := m.typ
	return ScanEntry{}, fmt.Errorf("%w: redis message type %s is not a scan response or its length is not at least 2", errParse, typeNames[typ])
}

// AsMap check if message is a redis array/set response, and convert to map[string]RedisMessage
func (m *RedisMessage) AsMap() (map[string]RedisMessage, error) {
	if err := m.Error(); err != nil {
		return nil, err
	}
	if (m.IsMap() || m.IsArray()) && len(m.values)%2 == 0 {
		return toMap(m.values)
	}
	typ := m.typ
	return nil, fmt.Errorf("%w: redis message type %s is not a map/array/set or its length is not even", errParse, typeNames[typ])
}

// AsStrMap check if message is a redis map/array/set response, and convert to map[string]string.
// redis nil element and other non string element will be present as zero.
func (m *RedisMessage) AsStrMap() (map[string]string, error) {
	if err := m.Error(); err != nil {
		return nil, err
	}
	if (m.IsMap() || m.IsArray()) && len(m.values)%2 == 0 {
		r := make(map[string]string, len(m.values)/2)
		for i := 0; i < len(m.values); i += 2 {
			k := m.values[i]
			v := m.values[i+1]
			r[k.string] = v.string
		}
		return r, nil
	}
	typ := m.typ
	return nil, fmt.Errorf("%w: redis message type %s is not a map/array/set or its length is not even", errParse, typeNames[typ])
}

// AsIntMap check if message is a redis map/array/set response, and convert to map[string]int64.
// redis nil element and other non integer element will be present as zero.
func (m *RedisMessage) AsIntMap() (map[string]int64, error) {
	if err := m.Error(); err != nil {
		return nil, err
	}
	if (m.IsMap() || m.IsArray()) && len(m.values)%2 == 0 {
		var err error
		r := make(map[string]int64, len(m.values)/2)
		for i := 0; i < len(m.values); i += 2 {
			k := m.values[i]
			v := m.values[i+1]
			if k.typ == typeBlobString || k.typ == typeSimpleString {
				if len(v.string) != 0 {
					if r[k.string], err = strconv.ParseInt(v.string, 0, 64); err != nil {
						return nil, err
					}
				} else if v.typ == typeInteger || v.typ == typeNull {
					r[k.string] = v.integer
				}
			}
		}
		return r, nil
	}
	typ := m.typ
	return nil, fmt.Errorf("%w: redis message type %s is not a map/array/set or its length is not even", errParse, typeNames[typ])
}

type KeyValues struct {
	Key    string
	Values []string
}

func (m *RedisMessage) AsLMPop() (kvs KeyValues, err error) {
	if err = m.Error(); err != nil {
		return KeyValues{}, err
	}
	if len(m.values) >= 2 {
		kvs.Key = m.values[0].string
		kvs.Values, err = m.values[1].AsStrSlice()
		return
	}
	typ := m.typ
	return KeyValues{}, fmt.Errorf("%w: redis message type %s is not a LMPOP response", errParse, typeNames[typ])
}

type KeyZScores struct {
	Key    string
	Values []ZScore
}

func (m *RedisMessage) AsZMPop() (kvs KeyZScores, err error) {
	if err = m.Error(); err != nil {
		return KeyZScores{}, err
	}
	if len(m.values) >= 2 {
		kvs.Key = m.values[0].string
		kvs.Values, err = m.values[1].AsZScores()
		return
	}
	typ := m.typ
	return KeyZScores{}, fmt.Errorf("%w: redis message type %s is not a ZMPOP response", errParse, typeNames[typ])
}

type FtSearchDoc struct {
	Doc   map[string]string
	Key   string
	Score float64
}

func (m *RedisMessage) AsFtSearch() (total int64, docs []FtSearchDoc, err error) {
	if err = m.Error(); err != nil {
		return 0, nil, err
	}
	if m.IsMap() {
		for i := 0; i < len(m.values); i += 2 {
			switch m.values[i].string {
			case "total_results":
				total = m.values[i+1].integer
			case "results":
				records := m.values[i+1].values
				docs = make([]FtSearchDoc, len(records))
				for d, record := range records {
					for j := 0; j < len(record.values); j += 2 {
						switch record.values[j].string {
						case "id":
							docs[d].Key = record.values[j+1].string
						case "extra_attributes":
							docs[d].Doc, _ = record.values[j+1].AsStrMap()
						case "score":
							docs[d].Score, _ = strconv.ParseFloat(record.values[j+1].string, 64)
						}
					}
				}
			case "error":
				for _, e := range m.values[i+1].values {
					e := e
					return 0, nil, (*RedisError)(&e)
				}
			}
		}
		return
	}
	if len(m.values) > 0 {
		total = m.values[0].integer
		wscore := false
		wattrs := false
		offset := 1
		if len(m.values) > 2 {
			if m.values[2].string == "" {
				wattrs = true
				offset++
			} else {
				_, err1 := strconv.ParseFloat(m.values[1].string, 64)
				_, err2 := strconv.ParseFloat(m.values[2].string, 64)
				wscore = err1 != nil && err2 == nil
				offset++
			}
		}
		if len(m.values) > 3 && m.values[3].string == "" {
			wattrs = true
			offset++
		}
		docs = make([]FtSearchDoc, 0, (len(m.values)-1)/offset)
		for i := 1; i < len(m.values); i++ {
			doc := FtSearchDoc{Key: m.values[i].string}
			if wscore {
				i++
				doc.Score, _ = strconv.ParseFloat(m.values[i].string, 64)
			}
			if wattrs {
				i++
				doc.Doc, _ = m.values[i].AsStrMap()
			}
			docs = append(docs, doc)
		}
		return
	}
	typ := m.typ
	return 0, nil, fmt.Errorf("%w: redis message type %s is not a FT.SEARCH response", errParse, typeNames[typ])
}

func (m *RedisMessage) AsFtAggregate() (total int64, docs []map[string]string, err error) {
	if err = m.Error(); err != nil {
		return 0, nil, err
	}
	if m.IsMap() {
		for i := 0; i < len(m.values); i += 2 {
			switch m.values[i].string {
			case "total_results":
				total = m.values[i+1].integer
			case "results":
				records := m.values[i+1].values
				docs = make([]map[string]string, len(records))
				for d, record := range records {
					for j := 0; j < len(record.values); j += 2 {
						switch record.values[j].string {
						case "extra_attributes":
							docs[d], _ = record.values[j+1].AsStrMap()
						}
					}
				}
			case "error":
				for _, e := range m.values[i+1].values {
					e := e
					return 0, nil, (*RedisError)(&e)
				}
			}
		}
		return
	}
	if len(m.values) > 0 {
		total = m.values[0].integer
		docs = make([]map[string]string, len(m.values)-1)
		for d, record := range m.values[1:] {
			docs[d], _ = record.AsStrMap()
		}
		return
	}
	typ := m.typ
	return 0, nil, fmt.Errorf("%w: redis message type %s is not a FT.AGGREGATE response", errParse, typeNames[typ])
}

func (m *RedisMessage) AsFtAggregateCursor() (cursor, total int64, docs []map[string]string, err error) {
	if m.IsArray() && len(m.values) == 2 && (m.values[0].IsArray() || m.values[0].IsMap()) {
		total, docs, err = m.values[0].AsFtAggregate()
		cursor = m.values[1].integer
	} else {
		total, docs, err = m.AsFtAggregate()
	}
	return
}

type GeoLocation struct {
	Name                      string
	Longitude, Latitude, Dist float64
	GeoHash                   int64
}

func (m *RedisMessage) AsGeosearch() ([]GeoLocation, error) {
	arr, err := m.ToArray()
	if err != nil {
		return nil, err
	}
	geoLocations := make([]GeoLocation, 0, len(arr))
	for _, v := range arr {
		var loc GeoLocation
		if v.IsString() {
			loc.Name = v.string
		} else {
			info := v.values
			var i int

			//name
			loc.Name = info[i].string
			i++
			//distance
			if i < len(info) && info[i].string != "" {
				loc.Dist, err = util.ToFloat64(info[i].string)
				if err != nil {
					return nil, err
				}
				i++
			}
			//hash
			if i < len(info) && info[i].IsInt64() {
				loc.GeoHash = info[i].integer
				i++
			}
			//coordinates
			if i < len(info) && info[i].values != nil {
				cord := info[i].values
				if len(cord) < 2 {
					return nil, fmt.Errorf("got %d, expected 2", len(info))
				}
				loc.Longitude, _ = cord[0].AsFloat64()
				loc.Latitude, _ = cord[1].AsFloat64()
			}
		}
		geoLocations = append(geoLocations, loc)
	}
	return geoLocations, nil
}

// ClientFlags is redis-server client flags, copy from redis/src/server.h (redis 7.0)
type ClientFlags uint64

const (
	ClientSlave            ClientFlags = 1 << 0  /* This client is a replica */
	ClientMaster           ClientFlags = 1 << 1  /* This client is a master */
	ClientMonitor          ClientFlags = 1 << 2  /* This client is a slave monitor, see MONITOR */
	ClientMulti            ClientFlags = 1 << 3  /* This client is in a MULTI context */
	ClientBlocked          ClientFlags = 1 << 4  /* The client is waiting in a blocking operation */
	ClientDirtyCAS         ClientFlags = 1 << 5  /* Watched keys modified. EXEC will fail. */
	ClientCloseAfterReply  ClientFlags = 1 << 6  /* Close after writing entire reply. */
	ClientUnBlocked        ClientFlags = 1 << 7  /* This client was unblocked and is stored in server.unblocked_clients */
	ClientScript           ClientFlags = 1 << 8  /* This is a non-connected client used by Lua */
	ClientAsking           ClientFlags = 1 << 9  /* Client issued the ASKING command */
	ClientCloseASAP        ClientFlags = 1 << 10 /* Close this client ASAP */
	ClientUnixSocket       ClientFlags = 1 << 11 /* Client connected via Unix domain socket */
	ClientDirtyExec        ClientFlags = 1 << 12 /* EXEC will fail for errors while queueing */
	ClientMasterForceReply ClientFlags = 1 << 13 /* Queue replies even if is master */
	ClientForceAOF         ClientFlags = 1 << 14 /* Force AOF propagation of current cmd. */
	ClientForceRepl        ClientFlags = 1 << 15 /* Force replication of current cmd. */
	ClientPrePSync         ClientFlags = 1 << 16 /* Instance don't understand PSYNC. */
	ClientReadOnly         ClientFlags = 1 << 17 /* Cluster client is in read-only state. */
	ClientPubSub           ClientFlags = 1 << 18 /* Client is in Pub/Sub mode. */
	ClientPreventAOFProp   ClientFlags = 1 << 19 /* Don't propagate to AOF. */
	ClientPreventReplProp  ClientFlags = 1 << 20 /* Don't propagate to slaves. */
	ClientPreventProp      ClientFlags = ClientPreventAOFProp | ClientPreventReplProp
	ClientPendingWrite     ClientFlags = 1 << 21 /* Client has output to send but a-write handler is yet not installed. */
	ClientReplyOff         ClientFlags = 1 << 22 /* Don't send replies to client. */
	ClientReplySkipNext    ClientFlags = 1 << 23 /* Set ClientREPLY_SKIP for next cmd */
	ClientReplySkip        ClientFlags = 1 << 24 /* Don't send just this reply. */
	ClientLuaDebug         ClientFlags = 1 << 25 /* Run EVAL in debug mode. */
	ClientLuaDebugSync     ClientFlags = 1 << 26 /* EVAL debugging without fork() */
	ClientModule           ClientFlags = 1 << 27 /* Non connected client used by some module. */
	ClientProtected        ClientFlags = 1 << 28 /* Client should not be freed for now. */
	ClientExecutingCommand ClientFlags = 1 << 29 /* Indicates that the client is currently in the process of handling
	   a command. usually this will be marked only during call()
	   however, blocked clients might have this flag kept until they
	   will try to reprocess the command. */
	ClientPendingCommand      ClientFlags = 1 << 30 /* Indicates the client has a fully * parsed command ready for execution. */
	ClientTracking            ClientFlags = 1 << 31 /* Client enabled keys tracking in order to perform client side caching. */
	ClientTrackingBrokenRedir ClientFlags = 1 << 32 /* Target client is invalid. */
	ClientTrackingBCAST       ClientFlags = 1 << 33 /* Tracking in BCAST mode. */
	ClientTrackingOptIn       ClientFlags = 1 << 34 /* Tracking in opt-in mode. */
	ClientTrackingOptOut      ClientFlags = 1 << 35 /* Tracking in opt-out mode. */
	ClientTrackingCaching     ClientFlags = 1 << 36 /* CACHING yes/no was given, depending on optin/optout mode. */
	ClientTrackingNoLoop      ClientFlags = 1 << 37 /* Don't send invalidation messages about writes performed by myself.*/
	ClientInTimeoutTable      ClientFlags = 1 << 38 /* This client is in the timeout table. */
	ClientProtocolError       ClientFlags = 1 << 39 /* Protocol error chatting with it. */
	ClientCloseAfterCommand   ClientFlags = 1 << 40 /* Close after executing commands * and writing entire reply. */
	ClientDenyBlocking        ClientFlags = 1 << 41 /* Indicate that the client should not be blocked. currently, turned on inside MULTI, Lua, RM_Call, and AOF client */
	ClientReplRDBOnly         ClientFlags = 1 << 42 /* This client is a replica that only wants RDB without replication buffer. */
	ClientNoEvict             ClientFlags = 1 << 43 /* This client is protected against client memory eviction. */
	ClientAllowOOM            ClientFlags = 1 << 44 /* Client used by RM_Call is allowed to fully execute scripts even when in OOM */
	ClientNoTouch             ClientFlags = 1 << 45 /* This client will not touch LFU/LRU stats. */
	ClientPushing             ClientFlags = 1 << 46 /* This client is pushing notifications. */
)

// ClientInfo is redis-server ClientInfo
type ClientInfo struct {
	Addr               string        // address/port of the client
	LAddr              string        // address/port of local address client connected to (bind address)
	Name               string        // the name set by the client with CLIENT SETNAME
	Events             string        // file descriptor events (see below)
	LastCmd            string        // cmd, last command played
	User               string        // the authenticated username of the client
	LibName            string        // redis version 7.2, client library name
	LibVer             string        // redis version 7.2, client library version
	ID                 int64         // redis version 2.8.12, a unique 64-bit client ID
	FD                 int64         // file descriptor corresponding to the socket
	Age                time.Duration // total duration of the connection in seconds
	Idle               time.Duration // idle time of the connection in seconds
	Flags              ClientFlags   // client flags (see below)
	DB                 int           // current database ID
	Sub                int           // number of channel subscriptions
	PSub               int           // number of pattern matching subscriptions
	SSub               int           // redis version 7.0.3, number of shard channel subscriptions
	Multi              int           // number of commands in a MULTI/EXEC context
	Watch              int           // redis version 7.4 RC1, number of keys this client is currently watching.
	QueryBuf           int           // qbuf, query buffer length (0 means no query pending)
	QueryBufFree       int           // qbuf-free, free space of the query buffer (0 means the buffer is full)
	ArgvMem            int           // incomplete arguments for the next command (already extracted from query buffer)
	MultiMem           int           // redis version 7.0, memory is used up by buffered multi commands
	BufferSize         int           // rbs, usable size of buffer
	BufferPeak         int           // rbp, peak used size of buffer in last 5 sec interval
	OutputBufferLength int           // obl, output buffer length
	OutputListLength   int           // oll, output list length (replies are queued in this list when the buffer is full)
	OutputMemory       int           // omem, output buffer memory usage
	TotalMemory        int           // tot-mem, total memory consumed by this client in its various buffers
	Redir              int64         // client id of current client tracking redirection
	Resp               int           // redis version 7.0, client RESP protocol version
}

// AsClientInfo check if message is a redis RESP3 client info response, and return it
func (m *RedisMessage) AsClientInfo() (*ClientInfo, error) {
	txt, err := m.ToString()
	if err != nil {
		return nil, err
	}
	info := &ClientInfo{}

	for _, s := range strings.Split(strings.TrimPrefix(strings.TrimSpace(txt), "txt:"), " ") {
		kv := strings.Split(s, "=")
		if len(kv) != 2 {
			return nil, fmt.Errorf("redis: unexpected client info data (%s)", s)
		}
		key, val := kv[0], kv[1]

		switch key {
		case "id":
			info.ID, err = strconv.ParseInt(val, 10, 64)
		case "addr":
			info.Addr = val
		case "laddr":
			info.LAddr = val
		case "fd":
			info.FD, err = strconv.ParseInt(val, 10, 64)
		case "name":
			info.Name = val
		case "age":
			var age int
			if age, err = strconv.Atoi(val); err == nil {
				info.Age = time.Duration(age) * time.Second
			}
		case "idle":
			var idle int
			if idle, err = strconv.Atoi(val); err == nil {
				info.Idle = time.Duration(idle) * time.Second
			}
		case "flags":
			if val == "N" {
				break
			}

			for i := 0; i < len(val); i++ {
				switch val[i] {
				case 'S':
					info.Flags |= ClientSlave
				case 'O':
					info.Flags |= ClientSlave | ClientMonitor
				case 'M':
					info.Flags |= ClientMaster
				case 'P':
					info.Flags |= ClientPubSub
				case 'x':
					info.Flags |= ClientMulti
				case 'b':
					info.Flags |= ClientBlocked
				case 't':
					info.Flags |= ClientTracking
				case 'R':
					info.Flags |= ClientTrackingBrokenRedir
				case 'B':
					info.Flags |= ClientTrackingBCAST
				case 'd':
					info.Flags |= ClientDirtyCAS
				case 'c':
					info.Flags |= ClientCloseAfterCommand
				case 'u':
					info.Flags |= ClientUnBlocked
				case 'A':
					info.Flags |= ClientCloseASAP
				case 'U':
					info.Flags |= ClientUnixSocket
				case 'r':
					info.Flags |= ClientReadOnly
				case 'e':
					info.Flags |= ClientNoEvict
				case 'T':
					info.Flags |= ClientNoTouch
				default:
					return nil, fmt.Errorf("redis: unexpected client info flags(%s)", string(val[i]))
				}
			}
		case "db":

			info.DB, err = strconv.Atoi(val)
		case "sub":
			info.Sub, err = strconv.Atoi(val)
		case "psub":
			info.PSub, err = strconv.Atoi(val)
		case "ssub":
			info.SSub, err = strconv.Atoi(val)
		case "multi":
			info.Multi, err = strconv.Atoi(val)
		case "watch":
			info.Watch, err = strconv.Atoi(val)
		case "qbuf":
			info.QueryBuf, err = strconv.Atoi(val)
		case "qbuf-free":
			info.QueryBufFree, err = strconv.Atoi(val)
		case "argv-mem":
			info.ArgvMem, err = strconv.Atoi(val)
		case "multi-mem":
			info.MultiMem, err = strconv.Atoi(val)
		case "rbs":
			info.BufferSize, err = strconv.Atoi(val)
		case "rbp":
			info.BufferPeak, err = strconv.Atoi(val)
		case "obl":
			info.OutputBufferLength, err = strconv.Atoi(val)
		case "oll":
			info.OutputListLength, err = strconv.Atoi(val)
		case "omem":
			info.OutputMemory, err = strconv.Atoi(val)
		case "tot-mem":
			info.TotalMemory, err = strconv.Atoi(val)
		case "events":
			info.Events = val
		case "cmd":
			info.LastCmd = val
		case "user":
			info.User = val
		case "redir":
			info.Redir, err = strconv.ParseInt(val, 10, 64)
		case "resp":
			info.Resp, err = strconv.Atoi(val)
		case "lib-name":
			info.LibName = val
		case "lib-ver":
			info.LibVer = val
		}

		if err != nil {
			return nil, err
		}

	}

	return info, nil
}

type ACLLogEntry struct {
	Count                int64
	Reason               string
	Context              string
	Object               string
	Username             string
	AgeSeconds           float64
	ClientInfo           *ClientInfo
	EntryID              int64
	TimestampCreated     int64
	TimestampLastUpdated int64
}

func (m *RedisMessage) AsACLLogEntry() ([]*ACLLogEntry, error) {
	arr, err := m.ToArray()
	if err != nil {
		return nil, err

	}
	logEntries := make([]*ACLLogEntry, 0, len(arr))

	for _, msg := range arr {
		log, err := msg.AsMap()
		if err != nil {
			return nil, err
		}
		entry := ACLLogEntry{}

		if attr, ok := log["count"]; ok {
			entry.Count, err = attr.AsInt64()
		}
		if attr, ok := log["reason"]; ok {
			entry.Reason, err = attr.ToString()
		}
		if attr, ok := log["context"]; ok {
			entry.Context, err = attr.ToString()
		}
		if attr, ok := log["object"]; ok {
			entry.Object, err = attr.ToString()
		}
		if attr, ok := log["username"]; ok {
			entry.Username, err = attr.ToString()
		}
		if attr, ok := log["age-seconds"]; ok {
			entry.AgeSeconds, err = attr.AsFloat64()
		}
		if attr, ok := log["client-info"]; ok {
			entry.ClientInfo, err = attr.AsClientInfo()
		}
		if attr, ok := log["entry-id"]; ok {
			entry.EntryID, err = attr.AsInt64()
		}
		if attr, ok := log["timestamp-created"]; ok {
			entry.TimestampCreated, err = attr.AsInt64()
		}
		if attr, ok := log["timestamp-last-updated"]; ok {
			entry.TimestampLastUpdated, err = attr.AsInt64()
		}

		if err != nil {
			return nil, err
		}
		logEntries = append(logEntries, &entry)
	}
	return logEntries, nil
}

// ToMap check if message is a redis RESP3 map response, and return it
func (m *RedisMessage) ToMap() (map[string]RedisMessage, error) {
	if m.IsMap() {
		return toMap(m.values)
	}
	if err := m.Error(); err != nil {
		return nil, err
	}
	typ := m.typ
	return nil, fmt.Errorf("%w: redis message type %s is not a RESP3 map", errParse, typeNames[typ])
}

// ToAny turns message into go any value
func (m *RedisMessage) ToAny() (any, error) {
	if err := m.Error(); err != nil {
		return nil, err
	}
	switch m.typ {
	case typeFloat:
		return util.ToFloat64(m.string)
	case typeBlobString, typeSimpleString, typeVerbatimString, typeBigNumber:
		return m.string, nil
	case typeBool:
		return m.integer == 1, nil
	case typeInteger:
		return m.integer, nil
	case typeMap:
		vs := make(map[string]any, len(m.values)/2)
		for i := 0; i < len(m.values); i += 2 {
			if v, err := m.values[i+1].ToAny(); err != nil && !IsRedisNil(err) {
				vs[m.values[i].string] = err
			} else {
				vs[m.values[i].string] = v
			}
		}
		return vs, nil
	case typeSet, typeArray:
		vs := make([]any, len(m.values))
		for i := 0; i < len(m.values); i++ {
			if v, err := m.values[i].ToAny(); err != nil && !IsRedisNil(err) {
				vs[i] = err
			} else {
				vs[i] = v
			}
		}
		return vs, nil
	}
	typ := m.typ
	return nil, fmt.Errorf("%w: redis message type %s is not a supported in ToAny", errParse, typeNames[typ])
}

// IsCacheHit check if message is from client side cache
func (m *RedisMessage) IsCacheHit() bool {
	return m.attrs == cacheMark
}

// CacheTTL returns the remaining TTL in seconds of client side cache
func (m *RedisMessage) CacheTTL() (ttl int64) {
	milli := m.CachePTTL()
	if milli > 0 {
		if ttl = milli / 1000; milli > ttl*1000 {
			ttl++
		}
		return ttl
	}
	return milli
}

// CachePTTL returns the remaining PTTL in seconds of client side cache
func (m *RedisMessage) CachePTTL() int64 {
	milli := m.getExpireAt()
	if milli == 0 {
		return -1
	}
	if milli = milli - time.Now().UnixMilli(); milli < 0 {
		milli = 0
	}
	return milli
}

// CachePXAT returns the remaining PXAT in seconds of client side cache
func (m *RedisMessage) CachePXAT() int64 {
	milli := m.getExpireAt()
	if milli == 0 {
		return -1
	}
	return milli
}

func (m *RedisMessage) relativePTTL(now time.Time) int64 {
	return m.getExpireAt() - now.UnixMilli()
}

func (m *RedisMessage) getExpireAt() int64 {
	return int64(m.ttl[0]) | int64(m.ttl[1])<<8 | int64(m.ttl[2])<<16 | int64(m.ttl[3])<<24 |
		int64(m.ttl[4])<<32 | int64(m.ttl[5])<<40 | int64(m.ttl[6])<<48
}

func (m *RedisMessage) setExpireAt(pttl int64) {
	m.ttl[0] = byte(pttl)
	m.ttl[1] = byte(pttl >> 8)
	m.ttl[2] = byte(pttl >> 16)
	m.ttl[3] = byte(pttl >> 24)
	m.ttl[4] = byte(pttl >> 32)
	m.ttl[5] = byte(pttl >> 40)
	m.ttl[6] = byte(pttl >> 48)
}

func toMap(values []RedisMessage) (map[string]RedisMessage, error) {
	r := make(map[string]RedisMessage, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		if values[i].typ == typeBlobString || values[i].typ == typeSimpleString {
			r[values[i].string] = values[i+1]
			continue
		}
		typ := values[i].typ
		return nil, fmt.Errorf("%w: redis message type %s as map key is not supported", errParse, typeNames[typ])
	}
	return r, nil
}

func (m *RedisMessage) approximateSize() (s int) {
	s += messageStructSize
	s += len(m.string)
	for _, v := range m.values {
		s += v.approximateSize()
	}
	return
}

// String returns human-readable representation of RedisMessage
func (m *RedisMessage) String() string {
	v, _ := (*prettyRedisMessage)(m).MarshalJSON()
	return string(v)
}

type prettyRedisMessage RedisMessage

// MarshalJSON implements json.Marshaler interface
func (m *prettyRedisMessage) MarshalJSON() ([]byte, error) {
	type PrettyRedisMessage struct {
		Value any    `json:"Value,omitempty"`
		Type  string `json:"Type,omitempty"`
		Error string `json:"Error,omitempty"`
		Ttl   string `json:"TTL,omitempty"`
	}
	org := (*RedisMessage)(m)
	strType, ok := typeNames[m.typ]
	if !ok {
		strType = "unknown"
	}
	obj := PrettyRedisMessage{Type: strType}
	if m.ttl != [7]byte{} {
		obj.Ttl = time.UnixMilli(org.CachePXAT()).UTC().String()
	}
	if err := org.Error(); err != nil {
		obj.Error = err.Error()
	}
	switch m.typ {
	case typeFloat, typeBlobString, typeSimpleString, typeVerbatimString, typeBigNumber:
		obj.Value = m.string
	case typeBool:
		obj.Value = m.integer == 1
	case typeInteger:
		obj.Value = m.integer
	case typeMap, typeSet, typeArray:
		values := make([]prettyRedisMessage, len(m.values))
		for i, value := range m.values {
			values[i] = prettyRedisMessage(value)
		}
		obj.Value = values
	}
	return json.Marshal(obj)
}
