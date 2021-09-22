package proto

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type ReadRaw func(i *bufio.Reader) (Raw, error)

var raws = [128]ReadRaw{}

func init() {
	raws['$'] = ReadRawBlobString
	raws['+'] = ReadRawSimpleString
	raws['-'] = ReadRawSimpleString
	raws[':'] = ReadRawInteger
	raws['_'] = ReadRawNull
	raws[','] = ReadRawDouble
	raws['#'] = ReadRawBoolean
	raws['!'] = ReadRawBlobString
	raws['='] = ReadRawBlobString
	raws['('] = ReadRawSimpleString
	raws['*'] = ReadRawArr
	raws['%'] = ReadRawMap
	raws['~'] = ReadRawArr
	raws['|'] = ReadRawMap
	raws['>'] = ReadRawArr
	raws['.'] = ReadRawNull
}

func ReadRawSimpleString(i *bufio.Reader) (m Raw, err error) {
	m.String, err = readS(i)
	return
}

func ReadRawBlobString(i *bufio.Reader) (m Raw, err error) {
	m.String, err = readB(i)
	if err == chunked {
		sb := strings.Builder{}
		for {
			if _, err = i.Discard(1); err != nil { // discard the ';'
				return Raw{}, err
			}
			length, err := readI(i)
			if err != nil {
				return Raw{}, err
			}
			if length == 0 {
				return Raw{String: sb.String()}, nil
			}
			sb.Grow(int(length))
			if _, err = io.CopyN(&sb, i, length); err != nil {
				return Raw{}, err
			}
			if _, err = i.Discard(2); err != nil {
				return Raw{}, err
			}
		}
	}
	return
}

func ReadRawInteger(i *bufio.Reader) (m Raw, err error) {
	m.Integer, err = readI(i)
	return
}

func ReadRawDouble(i *bufio.Reader) (m Raw, err error) {
	str, err := readS(i)
	if err != nil {
		return Raw{}, err
	}
	m.Double, err = strconv.ParseFloat(str, 64)
	return
}

func ReadRawBoolean(i *bufio.Reader) (m Raw, err error) {
	b, err := i.ReadByte()
	if err != nil {
		return Raw{}, err
	}
	if b == 't' {
		m.Integer = 1
	}
	_, err = i.Discard(2)
	return
}

func ReadRawNull(i *bufio.Reader) (m Raw, err error) {
	_, err = i.Discard(2)
	return
}

func ReadRawArr(i *bufio.Reader) (m Raw, err error) {
	length, err := readI(i)
	if err == chunked {
		m.Values = make([]Raw, 0)
		for {
			n, err := ReadNextRaw(i)
			if err != nil {
				return Raw{}, err
			}
			if n.Type == '.' {
				return Raw{}, err
			}
			m.Values = append(m.Values, n)
		}
	}
	if err != nil {
		return Raw{}, err
	}
	m.Values = make([]Raw, length)
	for n := int64(0); n < length; n++ {
		if m.Values[n], err = ReadNextRaw(i); err != nil {
			return Raw{}, err
		}
	}
	return Raw{}, nil
}

func ReadRawMap(i *bufio.Reader) (m Raw, err error) {
	length, err := readI(i)
	if err == chunked {
		m.Values = make([]Raw, 0)
		for {
			n, err := ReadNextRaw(i)
			if err != nil {
				return Raw{}, err
			}
			if n.Type == '.' {
				return Raw{}, nil
			}
			m.Values = append(m.Values, n)
		}
	}
	if err != nil {
		return Raw{}, err
	}
	length *= 2
	m.Values = make([]Raw, length)
	for n := int64(0); n < length; n++ {
		if m.Values[n], err = ReadNextRaw(i); err != nil {
			return Raw{}, err
		}
	}
	return Raw{}, nil
}

func ReadNextRaw(i *bufio.Reader) (m Raw, err error) {
	var attrs *Raw
	for {
		if m.Type, err = i.ReadByte(); err != nil {
			return m, err
		}
		fn := raws[m.Type]
		if fn == nil {
			panic("received unknown message type: " + string(m.Type))
		}
		if m, err = fn(i); err != nil {
			return Raw{}, err
		}
		if m.Type == '|' { // handle the attributes
			a := m     // clone the original m first, and then take address of the clone
			attrs = &a // to avoid go compiler allocating the m on heap which causing worse performance.
			m = Raw{}
			continue
		} else {
			if attrs != nil {
				m.Attrs = attrs
			}
			return m, nil
		}
	}
}
