package rueidis

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"
	"unsafe"
)

const messageStructSize = int(unsafe.Sizeof(RedisMessage{}))

// IsRedisNil is a handy method to check if error is redis nil response.
// All redis nil response returns as an error.
func IsRedisNil(err error) bool {
	if e, ok := err.(*RedisError); ok {
		return e.IsNil()
	}
	return false
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
	return r.typ == '_'
}

// IsMoved checks if it is a redis MOVED message and returns moved address.
func (r *RedisError) IsMoved() (addr string, ok bool) {
	if ok = strings.HasPrefix(r.string, "MOVED"); ok {
		addr = strings.Split(r.string, " ")[2]
	}
	return
}

// IsAsk checks if it is a redis ASK message and returns ask address.
func (r *RedisError) IsAsk() (addr string, ok bool) {
	if ok = strings.HasPrefix(r.string, "ASK"); ok {
		addr = strings.Split(r.string, " ")[2]
	}
	return
}

// IsTryAgain checks if it is a redis TRYAGAIN message and returns ask address.
func (r *RedisError) IsTryAgain() bool {
	return strings.HasPrefix(r.string, "TRYAGAIN")
}

// IsNoScript checks if it is a redis NOSCRIPT message.
func (r *RedisError) IsNoScript() bool {
	return strings.HasPrefix(r.string, "NOSCRIPT")
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

// RedisError can be used to check if the redis response is an error message.
func (r RedisResult) RedisError() *RedisError {
	if err := r.val.Error(); err != nil {
		return err.(*RedisError)
	}
	return nil
}

// NonRedisError can be used to check if there is an underlying error (ex. network timeout).
func (r RedisResult) NonRedisError() error {
	return r.err
}

// Error returns either underlying error or redis error or nil
func (r RedisResult) Error() error {
	if r.err != nil {
		return r.err
	}
	if err := r.val.Error(); err != nil {
		return err
	}
	return nil
}

// ToMessage retrieves the RedisMessage
func (r RedisResult) ToMessage() (RedisMessage, error) {
	return r.val, r.Error()
}

// ToInt64 delegates to RedisMessage.ToInt64
func (r RedisResult) ToInt64() (int64, error) {
	if err := r.Error(); err != nil {
		return 0, err
	}
	return r.val.ToInt64()
}

// ToBool delegates to RedisMessage.ToBool
func (r RedisResult) ToBool() (bool, error) {
	if err := r.Error(); err != nil {
		return false, err
	}
	return r.val.ToBool()
}

// ToFloat64 delegates to RedisMessage.ToFloat64
func (r RedisResult) ToFloat64() (float64, error) {
	if err := r.Error(); err != nil {
		return 0, err
	}
	return r.val.ToFloat64()
}

// ToString delegates to RedisMessage.ToString
func (r RedisResult) ToString() (string, error) {
	if err := r.Error(); err != nil {
		return "", err
	}
	return r.val.ToString()
}

// AsReader delegates to RedisMessage.AsReader
func (r RedisResult) AsReader() (reader io.Reader, err error) {
	if err := r.Error(); err != nil {
		return nil, err
	}
	return r.val.AsReader()
}

// DecodeJSON delegates to RedisMessage.DecodeJSON
func (r RedisResult) DecodeJSON(v interface{}) error {
	if err := r.Error(); err != nil {
		return err
	}
	return r.val.DecodeJSON(v)
}

// AsInt64 delegates to RedisMessage.AsInt64
func (r RedisResult) AsInt64() (int64, error) {
	if err := r.Error(); err != nil {
		return 0, err
	}
	return r.val.AsInt64()
}

// AsBool delegates to RedisMessage.AsBool
func (r RedisResult) AsBool() (bool, error) {
	if err := r.Error(); err != nil {
		return false, err
	}
	return r.val.AsBool()
}

// AsFloat64 delegates to RedisMessage.AsFloat64
func (r RedisResult) AsFloat64() (float64, error) {
	if err := r.Error(); err != nil {
		return 0, err
	}
	return r.val.AsFloat64()
}

// ToArray delegates to RedisMessage.ToArray
func (r RedisResult) ToArray() ([]RedisMessage, error) {
	if err := r.Error(); err != nil {
		return nil, err
	}
	return r.val.ToArray()
}

// AsStrSlice delegates to RedisMessage.AsStrSlice
func (r RedisResult) AsStrSlice() ([]string, error) {
	if err := r.Error(); err != nil {
		return nil, err
	}
	return r.val.AsStrSlice()
}

// AsMap delegates to RedisMessage.AsMap
func (r RedisResult) AsMap() (map[string]RedisMessage, error) {
	if err := r.Error(); err != nil {
		return nil, err
	}
	return r.val.AsMap()
}

// AsStrMap delegates to RedisMessage.AsStrMap
func (r RedisResult) AsStrMap() (map[string]string, error) {
	if err := r.Error(); err != nil {
		return nil, err
	}
	return r.val.AsStrMap()
}

// ToMap delegates to RedisMessage.ToMap
func (r RedisResult) ToMap() (map[string]RedisMessage, error) {
	if err := r.Error(); err != nil {
		return nil, err
	}
	return r.val.ToMap()
}

// IsCacheHit delegates to RedisMessage.IsCacheHit
func (r RedisResult) IsCacheHit() bool {
	return r.val.IsCacheHit()
}

// RedisMessage is a redis response message, it may be a nil response
type RedisMessage struct {
	typ     byte
	string  string
	values  []RedisMessage
	attrs   *RedisMessage
	integer int64
}

// IsNil check if message is a redis nil response
func (m *RedisMessage) IsNil() bool {
	return m.typ == '_'
}

// Error check if message is a redis error response, including nil response
func (m *RedisMessage) Error() error {
	if m.typ == '-' || m.typ == '_' || m.typ == '!' {
		return (*RedisError)(m)
	}
	return nil
}

// ToString check if message is a redis string response, and return it
func (m *RedisMessage) ToString() (val string, err error) {
	if m.typ == '$' || m.typ == '+' {
		return m.string, nil
	}
	if m.typ == ':' || m.values != nil {
		typ := m.typ
		panic(fmt.Sprintf("redis message type %c is not a string", typ))
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

// DecodeJSON check if message is a redis string response and treat it as json, then unmarshal it into provided value
func (m *RedisMessage) DecodeJSON(v interface{}) (err error) {
	str, err := m.ToString()
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(strings.NewReader(str))
	return decoder.Decode(v)
}

// AsInt64 check if message is a redis string response, and parse it as int64
func (m *RedisMessage) AsInt64() (val int64, err error) {
	v, err := m.ToString()
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(v, 10, 64)
}

// AsBool checks if message is non-nil redis response, and parses it as bool
func (m *RedisMessage) AsBool() (val bool, err error) {
	if err = m.Error(); err != nil {
		return
	}
	switch m.typ {
	case '$', '+':
		val = m.string == "OK"
		return
	case ':':
		val = m.integer != 0
		return
	case '#':
		val = m.integer == 1
		return
	default:
		typ := m.typ
		panic(fmt.Sprintf("redis message type %c is not a int, string or bool", typ))
	}
}

// AsFloat64 check if message is a redis string response, and parse it as float64
func (m *RedisMessage) AsFloat64() (val float64, err error) {
	v, err := m.ToString()
	if err != nil {
		return 0, err
	}
	return strconv.ParseFloat(v, 64)
}

// ToInt64 check if message is a redis int response, and return it
func (m *RedisMessage) ToInt64() (val int64, err error) {
	if m.typ == ':' {
		return m.integer, nil
	}
	if err = m.Error(); err != nil {
		return 0, err
	}
	typ := m.typ
	panic(fmt.Sprintf("redis message type %c is not a int64", typ))
}

// ToBool check if message is a redis bool response, and return it
func (m *RedisMessage) ToBool() (val bool, err error) {
	if m.typ == '#' {
		return m.integer == 1, nil
	}
	if err = m.Error(); err != nil {
		return false, err
	}
	typ := m.typ
	panic(fmt.Sprintf("redis message type %c is not a bool", typ))
}

// ToFloat64 check if message is a redis double response, and return it
func (m *RedisMessage) ToFloat64() (val float64, err error) {
	if m.typ == ',' {
		return strconv.ParseFloat(m.string, 64)
	}
	if err = m.Error(); err != nil {
		return 0, err
	}
	typ := m.typ
	panic(fmt.Sprintf("redis message type %c is not a float64", typ))
}

// ToArray check if message is a redis array/set response, and return it
func (m *RedisMessage) ToArray() ([]RedisMessage, error) {
	if m.typ == '*' || m.typ == '~' {
		return m.values, nil
	}
	if err := m.Error(); err != nil {
		return nil, err
	}
	typ := m.typ
	panic(fmt.Sprintf("redis message type %c is not a array", typ))
}

// AsStrSlice check if message is a redis array/set response, and convert to []string.
// Non string value will be ignored, including nil value.
func (m *RedisMessage) AsStrSlice() ([]string, error) {
	values, err := m.ToArray()
	if err != nil {
		return nil, err
	}
	s := make([]string, 0, len(values))
	for _, v := range values {
		if v.typ == '$' || v.typ == '+' || len(v.string) != 0 {
			s = append(s, v.string)
		}
	}
	return s, nil
}

// AsMap check if message is a redis array/set response, and convert to map[string]RedisMessage
func (m *RedisMessage) AsMap() (map[string]RedisMessage, error) {
	values, err := m.ToArray()
	if err != nil {
		return nil, err
	}
	return toMap(values), nil
}

// AsStrMap check if message is a redis map/array/set response, and convert to map[string]string.
// Non string value will be ignored, including nil value.
func (m *RedisMessage) AsStrMap() (map[string]string, error) {
	if err := m.Error(); err != nil {
		return nil, err
	}
	if m.typ == '%' || m.typ == '*' || m.typ == '~' {
		r := make(map[string]string, len(m.values)/2)
		for i := 0; i < len(m.values); i += 2 {
			k := m.values[i]
			v := m.values[i+1]
			if (k.typ == '$' || k.typ == '+') && (v.typ == '$' || v.typ == '+' || len(v.string) != 0) {
				r[k.string] = v.string
			}
		}
		return r, nil
	}
	typ := m.typ
	panic(fmt.Sprintf("redis message type %c is not a map/array/set", typ))
}

// ToMap check if message is a redis map response, and return it
func (m *RedisMessage) ToMap() (map[string]RedisMessage, error) {
	if m.typ == '%' {
		return toMap(m.values), nil
	}
	if err := m.Error(); err != nil {
		return nil, err
	}
	typ := m.typ
	panic(fmt.Sprintf("redis message type %c is not a map", typ))
}

// IsCacheHit check if message is from client side cache
func (m *RedisMessage) IsCacheHit() bool {
	return m.attrs == cacheMark
}

func toMap(values []RedisMessage) map[string]RedisMessage {
	r := make(map[string]RedisMessage, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		if values[i].typ == '$' || values[i].typ == '+' {
			r[values[i].string] = values[i+1]
			continue
		}
		typ := values[i].typ
		panic(fmt.Sprintf("redis message type %c as map key is not supported", typ))
	}
	return r
}

func (m *RedisMessage) approximateSize() (s int) {
	s += messageStructSize
	s += len(m.string)
	for _, v := range m.values {
		s += v.approximateSize()
	}
	return
}
