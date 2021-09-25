package proto

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

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

type StringArray []string

type reader func(i *bufio.Reader) (Message, error)

var readers = [128]reader{}

func init() {
	readers['$'] = readBlobString
	readers['+'] = readSimpleString
	readers['-'] = readSimpleString
	readers[':'] = readInteger
	readers['_'] = readNull
	readers[','] = readDouble
	readers['#'] = readBoolean
	readers['!'] = readBlobString
	readers['='] = readBlobString
	readers['('] = readSimpleString
	readers['*'] = readArray
	readers['%'] = readMap
	readers['~'] = readArray
	readers['|'] = readMap
	readers['>'] = readArray
	readers['.'] = readNull
}

func readSimpleString(i *bufio.Reader) (m Message, err error) {
	m.String, err = readS(i)
	return
}

func readBlobString(i *bufio.Reader) (m Message, err error) {
	m.String, err = readB(i)
	if err == chunked {
		sb := strings.Builder{}
		for {
			if _, err = i.Discard(1); err != nil { // discard the ';'
				return Message{}, err
			}
			length, err := readI(i)
			if err != nil {
				return Message{}, err
			}
			if length == 0 {
				return Message{String: sb.String()}, nil
			}
			sb.Grow(int(length))
			if _, err = io.CopyN(&sb, i, length); err != nil {
				return Message{}, err
			}
			if _, err = i.Discard(2); err != nil {
				return Message{}, err
			}
		}
	}
	return
}

func readInteger(i *bufio.Reader) (m Message, err error) {
	m.Integer, err = readI(i)
	return
}

func readDouble(i *bufio.Reader) (m Message, err error) {
	str, err := readS(i)
	if err != nil {
		return Message{}, err
	}
	m.Double, err = strconv.ParseFloat(str, 64)
	return
}

func readBoolean(i *bufio.Reader) (m Message, err error) {
	b, err := i.ReadByte()
	if err != nil {
		return Message{}, err
	}
	if b == 't' {
		m.Integer = 1
	}
	_, err = i.Discard(2)
	return
}

func readNull(i *bufio.Reader) (m Message, err error) {
	_, err = i.Discard(2)
	return
}

func readArray(i *bufio.Reader) (m Message, err error) {
	length, err := readI(i)
	if err == chunked {
		m.Values, err = readE(i)
	}
	if err != nil {
		return Message{}, err
	}
	m.Values, err = readA(i, int(length))
	return
}

func readMap(i *bufio.Reader) (m Message, err error) {
	length, err := readI(i)
	if err == chunked {
		m.Values, err = readE(i)
	}
	if err != nil {
		return Message{}, err
	}
	m.Values, err = readA(i, int(length*2))
	return
}

func ReadNextMessage(i *bufio.Reader) (m Message, err error) {
	var attrs *Message
	for {
		if m.Type, err = i.ReadByte(); err != nil {
			return m, err
		}
		fn := readers[m.Type]
		if fn == nil {
			panic("received unknown message type: " + string(m.Type))
		}
		if m, err = fn(i); err != nil {
			return Message{}, err
		}
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
	err = write(o, '*', strconv.Itoa(len(cmd)))
	for _, m := range cmd {
		err = writeB(o, '$', m)
	}
	return err
}
