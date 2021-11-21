package proto

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"unsafe"
)

const MessageStructSize = int(unsafe.Sizeof(Message{}))

type RedisError Message

func (r *RedisError) Error() string {
	if r.IsNil() {
		return "redis nil message"
	}
	return r.String
}

func (r *RedisError) IsNil() bool {
	return r.Type == '_'
}

func (r *RedisError) IsMoved() (addr string, ok bool) {
	if ok = strings.HasPrefix(r.String, "MOVED"); ok {
		addr = strings.Split(r.String, " ")[2]
	}
	return
}

func (r *RedisError) IsAsk() (addr string, ok bool) {
	if ok = strings.HasPrefix(r.String, "ASK"); ok {
		addr = strings.Split(r.String, " ")[2]
	}
	return
}

func (r *RedisError) IsTryAgain() bool {
	return strings.HasPrefix(r.String, "TRYAGAIN")
}

func NewResult(val Message, err error) Result {
	return Result{val: val, err: err}
}

func NewErrResult(err error) Result {
	return Result{err: err}
}

type Result struct {
	val Message
	err error
}

func (r Result) RedisError() *RedisError {
	if err := r.val.Error(); err != nil {
		return err.(*RedisError)
	}
	return nil
}

func (r Result) NonRedisError() error {
	return r.err
}

func (r Result) Error() error {
	if r.err != nil {
		return r.err
	}
	if err := r.val.Error(); err != nil {
		return err
	}
	return r.err
}

func (r Result) Value() (Message, error) {
	return r.val, r.Error()
}

func (r Result) ToInt64() (int64, error) {
	if err := r.Error(); err != nil {
		return 0, err
	}
	return r.val.ToInt64()
}

func (r Result) ToBool() (bool, error) {
	if err := r.Error(); err != nil {
		return false, err
	}
	return r.val.ToBool()
}

func (r Result) ToFloat64() (float64, error) {
	if err := r.Error(); err != nil {
		return 0, err
	}
	return r.val.ToFloat64()
}

func (r Result) ToString() (string, error) {
	if err := r.Error(); err != nil {
		return "", err
	}
	return r.val.ToString()
}

func (r Result) ToArray() ([]Message, error) {
	if err := r.Error(); err != nil {
		return nil, err
	}
	return r.val.ToArray()
}

func (r Result) ToMap() (map[string]Message, error) {
	if err := r.Error(); err != nil {
		return nil, err
	}
	return r.val.ToMap()
}

type Message struct {
	String  string
	Integer int64
	Values  []Message
	Attrs   *Message
	Type    byte
}

func (m *Message) Error() error {
	if m.Type == '-' || m.Type == '_' || m.Type == '!' {
		return (*RedisError)(m)
	}
	return nil
}

func (m *Message) ToString() (val string, err error) {
	if m.Type == '$' || m.Type == '+' {
		return m.String, nil
	}
	if m.Type == ':' || m.Values != nil {
		panic(fmt.Sprintf("redis message type %c is not a string", m.Type))
	}
	return m.String, m.Error()
}

func (m *Message) ToInt64() (val int64, err error) {
	if m.Type == ':' {
		return m.Integer, nil
	}
	if err = m.Error(); err != nil {
		return 0, err
	}
	panic(fmt.Sprintf("redis message type %c is not a int64", m.Type))
}

func (m *Message) ToBool() (val bool, err error) {
	if m.Type == '#' {
		return m.Integer == 1, nil
	}
	if err = m.Error(); err != nil {
		return false, err
	}
	panic(fmt.Sprintf("redis message type %c is not a bool", m.Type))
}

func (m *Message) ToFloat64() (val float64, err error) {
	if m.Type == ',' {
		return strconv.ParseFloat(m.String, 64)
	}
	if err = m.Error(); err != nil {
		return 0, err
	}
	panic(fmt.Sprintf("redis message type %c is not a float64", m.Type))
}

func (m *Message) ToArray() ([]Message, error) {
	if m.Type == '*' || m.Type == '~' {
		return m.Values, nil
	}
	if err := m.Error(); err != nil {
		return nil, err
	}
	panic(fmt.Sprintf("redis message type %c is not a array", m.Type))
}

func (m *Message) ToMap() (map[string]Message, error) {
	if m.Type == '%' {
		r := make(map[string]Message, len(m.Values)/2)
		for i := 0; i < len(m.Values); i += 2 {
			if m.Values[i].Type == '$' || m.Values[i].Type == '+' {
				r[m.Values[i].String] = m.Values[i+1]
				continue
			}
			panic(fmt.Sprintf("redis message type %c as map key is not supported by ToMap", m.Values[i].Type))
		}
		return r, nil
	}
	if err := m.Error(); err != nil {
		return nil, err
	}
	panic(fmt.Sprintf("redis message type %c is not a map", m.Type))
}

func (m *Message) ApproximateSize() (s int) {
	s += MessageStructSize
	s += len(m.String)
	for _, v := range m.Values {
		s += v.ApproximateSize()
	}
	if m.Attrs != nil {
		s += m.Attrs.ApproximateSize()
	}
	return
}

func ReadNextMessage(i *bufio.Reader) (m Message, err error) {
	var attrs *Message
	var typ byte
	for {
		if typ, err = i.ReadByte(); err != nil {
			return m, err
		}
		fn := readers[typ]
		if fn == nil {
			panic("received unknown message type: " + string(typ))
		}
		if m, err = fn(i); err != nil {
			return Message{}, err
		}
		m.Type = typ
		if m.Type == '|' { // handle the attributes
			a := m     // clone the original m first, and then take address of the clone
			attrs = &a // to avoid go compiler allocating the m on heap which causing worse performance.
			m = Message{}
			continue
		}
		m.Attrs = attrs
		return m, nil
	}
}

func WriteCmd(o *bufio.Writer, cmd []string) (err error) {
	err = WriteS(o, '*', strconv.Itoa(len(cmd)))
	for _, m := range cmd {
		err = WriteB(o, '$', m)
	}
	return err
}
