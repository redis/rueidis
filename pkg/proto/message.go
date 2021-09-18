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
	_ = o.WriteByte('$')
	_, _ = o.WriteString(strconv.Itoa(len(s.v)))
	_, _ = o.WriteString("\r\n")
	_, _ = o.WriteString(s.v)
	_, err := o.WriteString("\r\n")
	return err
}

type Verbatim struct {
	Attributes
	t string
	v string
}

func (v *Verbatim) WriteTo(o *bufio.Writer) error {
	_ = o.WriteByte('=')
	_, _ = o.WriteString(strconv.Itoa(len(v.v) + len(v.t) + 1))
	_, _ = o.WriteString("\r\n")
	_, _ = o.WriteString(v.t)
	_, _ = o.WriteString(":")
	_, _ = o.WriteString(v.v)
	_, err := o.WriteString("\r\n")
	return err
}

type Error struct {
	Attributes
	v string
}

func (e *Error) WriteTo(o *bufio.Writer) error {
	_ = o.WriteByte('!')
	_, _ = o.WriteString(strconv.Itoa(len(e.v)))
	_, _ = o.WriteString("\r\n")
	_, _ = o.WriteString(e.v)
	_, err := o.WriteString("\r\n")
	return err
}

type Int64 struct {
	Attributes
	v int64
}

func (i *Int64) WriteTo(o *bufio.Writer) error {
	_ = o.WriteByte(':')
	_, _ = o.WriteString(strconv.FormatInt(i.v, 10))
	_, err := o.WriteString("\r\n")
	return err
}

type BigInt struct {
	Attributes
	v big.Int
}

func (i *BigInt) WriteTo(o *bufio.Writer) error {
	_ = o.WriteByte(':')
	_, _ = o.WriteString(i.v.String())
	_, err := o.WriteString("\r\n")
	return err
}

type Float64 struct {
	Attributes
	v float64
}

func (f *Float64) WriteTo(o *bufio.Writer) error {
	_ = o.WriteByte(',')
	_, _ = o.WriteString(strconv.FormatFloat(f.v, 'f', -1, 64))
	_, err := o.WriteString("\r\n")
	return err
}

type Bool struct {
	Attributes
	v bool
}

func (b *Bool) WriteTo(o *bufio.Writer) error {
	_ = o.WriteByte('#')
	if b.v {
		_ = o.WriteByte('t')
	} else {
		_ = o.WriteByte('f')
	}
	_, err := o.WriteString("\r\n")
	return err
}

type Nil struct {
	Attributes
}

func (n *Nil) WriteTo(o *bufio.Writer) error {
	_ = o.WriteByte('_')
	_, err := o.WriteString("\r\n")
	return err
}

type Array struct {
	Attributes
	v []Message
}

func (a *Array) WriteTo(o *bufio.Writer) (err error) {
	_ = o.WriteByte('*')
	_, _ = o.WriteString(strconv.Itoa(len(a.v)))
	_, err = o.WriteString("\r\n")
	for _, m := range a.v {
		err = m.WriteTo(o)
	}
	return err
}

type Set struct {
	Attributes
	v []Message
}

func (s *Set) WriteTo(o *bufio.Writer) (err error) {
	_ = o.WriteByte('~')
	_, _ = o.WriteString(strconv.Itoa(len(s.v)))
	_, err = o.WriteString("\r\n")
	for _, m := range s.v {
		err = m.WriteTo(o)
	}
	return err
}

type Map struct {
	Attributes
	k []Message
	v []Message
}

func (s *Map) WriteTo(o *bufio.Writer) (err error) {
	_ = o.WriteByte('%')
	_, _ = o.WriteString(strconv.Itoa(len(s.k)))
	_, err = o.WriteString("\r\n")
	for i, m := range s.k {
		err = m.WriteTo(o)
		err = s.v[i].WriteTo(o)
	}
	return err
}

type Attributes struct {
	k []Message
	v []Message
}

func (s *Attributes) WriteTo(o *bufio.Writer) (err error) {
	_ = o.WriteByte('|')
	_, _ = o.WriteString(strconv.Itoa(len(s.k)))
	_, err = o.WriteString("\r\n")
	for i, m := range s.k {
		err = m.WriteTo(o)
		err = s.v[i].WriteTo(o)
	}
	return err
}

type Push struct {
	Attributes
	v []Message
}

func (s *Push) WriteTo(o *bufio.Writer) (err error) {
	_ = o.WriteByte('>')
	_, _ = o.WriteString(strconv.Itoa(len(s.v)))
	_, err = o.WriteString("\r\n")
	for _, m := range s.v {
		err = m.WriteTo(o)
	}
	return err
}
