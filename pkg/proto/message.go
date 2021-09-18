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
	v string
}

func (s *String) WriteTo(o *bufio.Writer) error {
	return blob(o, '$', s.v)
}

type Verbatim struct {
	Attributes
	t string
	v string
}

func (v *Verbatim) WriteTo(o *bufio.Writer) error {
	return blob(o, '=', v.t+":"+v.v)
}

type Error struct {
	Attributes
	v string
}

func (e *Error) WriteTo(o *bufio.Writer) error {
	return blob(o, '!', e.v)
}

type Int64 struct {
	Attributes
	v int64
}

func (i *Int64) WriteTo(o *bufio.Writer) error {
	return write(o, ':', strconv.FormatInt(i.v, 10))
}

type BigInt struct {
	Attributes
	v big.Int
}

func (i *BigInt) WriteTo(o *bufio.Writer) error {
	return write(o, '(', i.v.String())
}

type Float64 struct {
	Attributes
	v float64
}

func (f *Float64) WriteTo(o *bufio.Writer) error {
	return write(o, ',', strconv.FormatFloat(f.v, 'f', -1, 64))
}

type Bool struct {
	Attributes
	v bool
}

func (b *Bool) WriteTo(o *bufio.Writer) error {
	if b.v {
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
	v []Message
}

func (a *Array) WriteTo(o *bufio.Writer) (err error) {
	return writeA(o, '*', a.v)
}

type Set struct {
	Attributes
	v []Message
}

func (s *Set) WriteTo(o *bufio.Writer) (err error) {
	return writeA(o, '~', s.v)
}

type Map struct {
	Attributes
	k []Message
	v []Message
}

func (s *Map) WriteTo(o *bufio.Writer) (err error) {
	return writeM(o, '%', s.k, s.v)
}

type Attributes struct {
	k []Message
	v []Message
}

func (s *Attributes) WriteTo(o *bufio.Writer) (err error) {
	return writeM(o, '|', s.k, s.v)
}

type Push struct {
	Attributes
	v []Message
}

func (s *Push) WriteTo(o *bufio.Writer) (err error) {
	return writeA(o, '>', s.v)
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
