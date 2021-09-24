package proto

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"unsafe"
)

var chunked = errors.New("unbounded redis message")

func readS(i *bufio.Reader) (string, error) {
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

func readI(i *bufio.Reader) (int64, error) {
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

func readB(i *bufio.Reader) (string, error) {
	length, err := readI(i)
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

func readE(i *bufio.Reader) ([]Message, error) {
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

func readA(i *bufio.Reader, length int) (v []Message, err error) {
	v = make([]Message, length)
	for n := 0; n < length; n++ {
		if v[n], err = ReadNextMessage(i); err != nil {
			return nil, err
		}
	}
	return v, nil
}

func writeB(o *bufio.Writer, id byte, str string) (err error) {
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
