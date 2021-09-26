package proto

import (
	"bufio"
	"strconv"
	"unsafe"
)

const MessageStructSize = int(unsafe.Sizeof(Message{}))

type Result struct {
	Val Message
	Err error
}

type Message struct {
	String  string
	Integer int64
	Double  float64
	Values  []Message
	Attrs   *Message
	Type    byte
}

func (m *Message) Size() (s int) {
	s += MessageStructSize
	s += len(m.String)
	for _, v := range m.Values {
		s += v.Size()
	}
	if m.Attrs != nil {
		s += m.Attrs.Size()
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
