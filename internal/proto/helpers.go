package proto

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
	"unsafe"
)

var chunked = errors.New("unbounded redis message")

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
	m.String, err = ReadS(i)
	return
}

func readBlobString(i *bufio.Reader) (m Message, err error) {
	m.String, err = ReadB(i)
	if err == chunked {
		sb := strings.Builder{}
		for {
			if _, err = i.Discard(1); err != nil { // discard the ';'
				return Message{}, err
			}
			length, err := ReadI(i)
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
	m.Integer, err = ReadI(i)
	return
}

func readDouble(i *bufio.Reader) (m Message, err error) {
	str, err := ReadS(i)
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
	length, err := ReadI(i)
	if err == chunked {
		m.Values, err = ReadE(i)
	}
	if err != nil {
		return Message{}, err
	}
	m.Values, err = ReadA(i, int(length))
	return
}

func readMap(i *bufio.Reader) (m Message, err error) {
	length, err := ReadI(i)
	if err == chunked {
		m.Values, err = ReadE(i)
	}
	if err != nil {
		return Message{}, err
	}
	m.Values, err = ReadA(i, int(length*2))
	return
}

func ReadS(i *bufio.Reader) (string, error) {
	bs, err := i.ReadBytes('\n')
	if err != nil {
		return "", err
	}
	if trim := len(bs) - 2; trim < 0 {
		panic("received unexpected simple string message ending without CRLF")
	} else {
		bs = bs[:trim]
	}
	return *(*string)(unsafe.Pointer(&bs)), nil
}

func ReadI(i *bufio.Reader) (int64, error) {
	var v int64
	var neg bool
	for {
		c, err := i.ReadByte()
		if err != nil {
			return 0, err
		}
		switch {
		case '0' <= c && c <= '9':
			v = v*10 + int64(c-'0')
		case '\r' == c:
			_, err = i.Discard(1)
			if neg {
				return v * -1, err
			}
			return v, err
		case '-' == c:
			neg = true
		case '?' == c:
			return 0, chunked
		default:
			panic("received unexpected number byte: " + string(c))
		}
	}
}

func ReadB(i *bufio.Reader) (string, error) {
	length, err := ReadI(i)
	if err != nil {
		return "", err
	}
	bs := make([]byte, length)
	if _, err = io.ReadFull(i, bs); err != nil {
		return "", err
	}
	if _, err = i.Discard(2); err != nil {
		return "", err
	}
	return *(*string)(unsafe.Pointer(&bs)), nil
}

func ReadE(i *bufio.Reader) ([]Message, error) {
	v := make([]Message, 0)
	for {
		n, err := ReadNextMessage(i)
		if err != nil {
			return nil, err
		}
		if n.Type == '.' {
			return v, err
		}
		v = append(v, n)
	}
}

func ReadA(i *bufio.Reader, length int) (v []Message, err error) {
	v = make([]Message, length)
	for n := 0; n < length; n++ {
		if v[n], err = ReadNextMessage(i); err != nil {
			return nil, err
		}
	}
	return v, nil
}

func WriteB(o *bufio.Writer, id byte, str string) (err error) {
	_ = WriteS(o, id, strconv.Itoa(len(str)))
	_, _ = o.WriteString(str)
	_, err = o.WriteString("\r\n")
	return err
}

func WriteS(o *bufio.Writer, id byte, str string) (err error) {
	_ = o.WriteByte(id)
	_, _ = o.WriteString(str)
	_, err = o.WriteString("\r\n")
	return err
}
