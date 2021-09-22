package proto

import (
	"bufio"
	"math/big"
	"strconv"
)

type Message interface {
	SetAttributes(attrs Attributes)
	WriteTo(o *bufio.Writer) error
}

func (a *Attributes) SetAttributes(attrs Attributes) {
	*a = attrs
}

type String struct {
	Attributes
	Val string
}

func (s *String) WriteTo(o *bufio.Writer) error {
	return blob(o, '$', s.Val)
}

type Verbatim struct {
	Attributes
	Ver string
	Val string
}

func (v *Verbatim) WriteTo(o *bufio.Writer) error {
	return blob(o, '=', v.Ver+":"+v.Val)
}

type Error struct {
	Attributes
	Val string
}

func (e *Error) WriteTo(o *bufio.Writer) error {
	return blob(o, '!', e.Val)
}

type Int64 struct {
	Attributes
	Val int64
}

func (i *Int64) WriteTo(o *bufio.Writer) error {
	return write(o, ':', strconv.FormatInt(i.Val, 10))
}

type BigInt struct {
	Attributes
	Val big.Int
}

func (i *BigInt) WriteTo(o *bufio.Writer) error {
	return write(o, '(', i.Val.String())
}

type Float64 struct {
	Attributes
	Val float64
}

func (f *Float64) WriteTo(o *bufio.Writer) error {
	return write(o, ',', strconv.FormatFloat(f.Val, 'f', -1, 64))
}

type Bool struct {
	Attributes
	Val bool
}

func (b *Bool) WriteTo(o *bufio.Writer) error {
	if b.Val {
		return write(o, '#', "t")
	} else {
		return write(o, '#', "f")
	}
}

type Nil struct {
	Attributes
}

func (n *Nil) WriteTo(o *bufio.Writer) error {
	return write(o, '_', "")
}

type Array struct {
	Attributes
	Val []Message
}

func (a *Array) WriteTo(o *bufio.Writer) (err error) {
	return writeA(o, '*', a.Val)
}

type StringArray []string

func (a *StringArray) SetAttributes(attrs Attributes) {
}

func (a StringArray) WriteTo(o *bufio.Writer) (err error) {
	err = write(o, '*', strconv.Itoa(len(a)))
	for _, m := range a {
		err = blob(o, '$', m)
	}
	return err
}

type Set struct {
	Attributes
	Val []Message
}

func (s *Set) WriteTo(o *bufio.Writer) (err error) {
	return writeA(o, '~', s.Val)
}

type Map struct {
	Attributes
	Key []Message
	Val []Message
}

func (s *Map) WriteTo(o *bufio.Writer) (err error) {
	return writeM(o, '%', s.Key, s.Val)
}

type Attributes struct {
	Key []Message
	Val []Message
}

func (s *Attributes) WriteTo(o *bufio.Writer) (err error) {
	return writeM(o, '|', s.Key, s.Val)
}

type Push struct {
	Attributes
	Val []Message
}

func (s *Push) WriteTo(o *bufio.Writer) (err error) {
	return writeA(o, '>', s.Val)
}

func blob(o *bufio.Writer, id byte, str string) (err error) {
	_ = write(o, id, strconv.Itoa(len(str)))
	_, _ = o.WriteString(str)
	_, err = o.WriteString("\r\n")
	return err
}

func write(o *bufio.Writer, id byte, str string) (err error) {
	_ = o.WriteByte(id)
	_, _ = o.WriteString(str)
	_, err = o.WriteString("\r\n")
	return err
}

func writeA(o *bufio.Writer, id byte, v []Message) (err error) {
	err = write(o, id, strconv.Itoa(len(v)))
	for _, m := range v {
		err = m.WriteTo(o)
	}
	return err
}

func writeM(o *bufio.Writer, id byte, k, v []Message) (err error) {
	err = write(o, id, strconv.Itoa(len(k)))
	for i, m := range k {
		err = m.WriteTo(o)
		err = v[i].WriteTo(o)
	}
	return err
}

type Raw struct {
	String  string
	Integer int64
	Double  float64
	Values  []Raw
	Attrs   *Raw
	Type    byte
}
