package proto

// This is a minimum implementation of proto message using golang interface
// that showing the performance difference with using simple golang struct

import (
	"bufio"
	"strconv"

	"github.com/rueian/rueidis/internal/proto"
)

type reader func(i *bufio.Reader) (Message, error)

var readers = [128]reader{}

func init() {
	readers['$'] = readBlobString
	readers['*'] = readArray
}

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
	return proto.WriteB(o, '$', s.Val)
}

type Array struct {
	Attributes
	Val []Message
}

func (a *Array) WriteTo(o *bufio.Writer) (err error) {
	err = proto.WriteS(o, '*', strconv.Itoa(len(a.Val)))
	for _, m := range a.Val {
		err = m.WriteTo(o)
	}
	return err
}

type Attributes struct {
	Key []Message
	Val []Message
}

func (a *Attributes) WriteTo(o *bufio.Writer) (err error) {
	return nil
}

func readBlobString(i *bufio.Reader) (Message, error) {
	v, err := proto.ReadB(i)
	if err != nil {
		return nil, err
	}
	return &String{Val: v}, nil
}

func readArray(i *bufio.Reader) (Message, error) {
	length, err := proto.ReadI(i)
	if err != nil {
		return nil, err
	}
	v := make([]Message, length)
	for n := int64(0); n < length; n++ {
		if v[n], err = ReadNextInterfaceMessage(i); err != nil {
			return nil, err
		}
	}
	return &Array{Val: v}, nil
}

func ReadNextInterfaceMessage(i *bufio.Reader) (Message, error) {
	t, err := i.ReadByte()
	if err != nil {
		return nil, err
	}
	fn := readers[t]
	if fn == nil {
		panic("received unknown message type: " + string(t))
	}
	msg, err := fn(i)
	if err != nil {
		return nil, err
	}
	return msg, nil
}
