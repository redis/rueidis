package rueidis

import (
	"bufio"
	"errors"
	"io"
	"math"
	"strconv"
	"strings"
)

var errChunked = errors.New("unbounded redis message")
var errOldNull = errors.New("RESP2 null")

const (
	typeBlobString     = byte('$')
	typeSimpleString   = byte('+')
	typeSimpleErr      = byte('-')
	typeInteger        = byte(':')
	typeNull           = byte('_')
	typeEnd            = byte('.')
	typeFloat          = byte(',')
	typeBool           = byte('#')
	typeBlobErr        = byte('!')
	typeVerbatimString = byte('=')
	typeBigNumber      = byte('(')
	typeArray          = byte('*')
	typeMap            = byte('%')
	typeSet            = byte('~')
	typeAttribute      = byte('|')
	typePush           = byte('>')
)

var typeNames = make(map[byte]string, 16)

type reader func(i *bufio.Reader) (RedisMessage, error)

var readers = [256]reader{}

func init() {
	readers[typeBlobString] = readBlobString
	readers[typeSimpleString] = readSimpleString
	readers[typeSimpleErr] = readSimpleString
	readers[typeInteger] = readInteger
	readers[typeNull] = readNull
	readers[typeFloat] = readSimpleString
	readers[typeBool] = readBoolean
	readers[typeBlobErr] = readBlobString
	readers[typeVerbatimString] = readBlobString
	readers[typeBigNumber] = readSimpleString
	readers[typeArray] = readArray
	readers[typeMap] = readMap
	readers[typeSet] = readArray
	readers[typeAttribute] = readMap
	readers[typePush] = readArray
	readers[typeEnd] = readNull

	typeNames[typeBlobString] = "blob string"
	typeNames[typeSimpleString] = "simple string"
	typeNames[typeSimpleErr] = "simple error"
	typeNames[typeInteger] = "int64"
	typeNames[typeNull] = "null"
	typeNames[typeFloat] = "float64"
	typeNames[typeBool] = "boolean"
	typeNames[typeBlobErr] = "blob error"
	typeNames[typeVerbatimString] = "verbatim string"
	typeNames[typeBigNumber] = "big number"
	typeNames[typeArray] = "array"
	typeNames[typeMap] = "map"
	typeNames[typeSet] = "set"
	typeNames[typeAttribute] = "attribute"
	typeNames[typePush] = "push"
	typeNames[typeEnd] = "null"
}

func readSimpleString(i *bufio.Reader) (m RedisMessage, err error) {
	m.string, err = readS(i)
	return
}

func readBlobString(i *bufio.Reader) (m RedisMessage, err error) {
	m.string, err = readB(i)
	if err == errChunked {
		sb := strings.Builder{}
		for {
			if _, err = i.Discard(1); err != nil { // discard the ';'
				return RedisMessage{}, err
			}
			length, err := readI(i)
			if err != nil {
				return RedisMessage{}, err
			}
			if length == 0 {
				return RedisMessage{string: sb.String()}, nil
			}
			sb.Grow(int(length))
			if _, err = io.CopyN(&sb, i, length); err != nil {
				return RedisMessage{}, err
			}
			if _, err = i.Discard(2); err != nil {
				return RedisMessage{}, err
			}
		}
	}
	return
}

func readInteger(i *bufio.Reader) (m RedisMessage, err error) {
	m.integer, err = readI(i)
	return
}

func readBoolean(i *bufio.Reader) (m RedisMessage, err error) {
	b, err := i.ReadByte()
	if err != nil {
		return RedisMessage{}, err
	}
	if b == 't' {
		m.integer = 1
	}
	_, err = i.Discard(2)
	return
}

func readNull(i *bufio.Reader) (m RedisMessage, err error) {
	_, err = i.Discard(2)
	return
}

func readArray(i *bufio.Reader) (m RedisMessage, err error) {
	length, err := readI(i)
	if err == nil {
		if length == -1 {
			return m, errOldNull
		}
		m.values, err = readA(i, length)
	} else if err == errChunked {
		m.values, err = readE(i)
	}
	return m, err
}

func readMap(i *bufio.Reader) (m RedisMessage, err error) {
	length, err := readI(i)
	if err == nil {
		m.values, err = readA(i, length*2)
	} else if err == errChunked {
		m.values, err = readE(i)
	}
	return m, err
}

const ok = "OK"
const okrn = "OK\r\n"

func readS(i *bufio.Reader) (string, error) {
	if peek, _ := i.Peek(2); string(peek) == ok {
		if peek, _ = i.Peek(4); string(peek) == okrn {
			_, _ = i.Discard(4)
			return ok, nil
		}
	}
	bs, err := i.ReadBytes('\n')
	if err != nil {
		return "", err
	}
	if trim := len(bs) - 2; trim < 0 {
		return "", errors.New(unexpectedNoCRLF)
	} else {
		bs = bs[:trim]
	}
	return BinaryString(bs), nil
}

func readI(i *bufio.Reader) (v int64, err error) {
	bs, err := i.ReadSlice('\n')
	if err != nil {
		return 0, err
	}
	if len(bs) < 3 {
		return 0, errors.New(unexpectedNoCRLF)
	}
	if bs[0] == '?' {
		return 0, errChunked
	}
	var s = int64(1)
	if bs[0] == '-' {
		s = -1
		bs = bs[1:]
	}
	for _, c := range bs[:len(bs)-2] {
		if d := int64(c - '0'); d >= 0 && d <= 9 {
			v = v*10 + d
		} else {
			return 0, errors.New(unexpectedNumByte + strconv.Itoa(int(c)))
		}
	}
	return v * s, nil
}

func readB(i *bufio.Reader) (string, error) {
	length, err := readI(i)
	if err != nil {
		return "", err
	}
	if length == -1 {
		return "", errOldNull
	}
	bs := make([]byte, length)
	if _, err = io.ReadFull(i, bs); err != nil {
		return "", err
	}
	if _, err = i.Discard(2); err != nil {
		return "", err
	}
	return BinaryString(bs), nil
}

func readE(i *bufio.Reader) ([]RedisMessage, error) {
	v := make([]RedisMessage, 0)
	for {
		n, err := readNextMessage(i)
		if err != nil {
			return nil, err
		}
		if n.typ == '.' {
			return v, err
		}
		v = append(v, n)
	}
}

func readA(i *bufio.Reader, length int64) (v []RedisMessage, err error) {
	v = make([]RedisMessage, length)
	for n := int64(0); n < length; n++ {
		if v[n], err = readNextMessage(i); err != nil {
			return nil, err
		}
	}
	return v, nil
}

func writeB(o *bufio.Writer, id byte, str string) (err error) {
	_ = writeN(o, id, len(str))
	_, _ = o.WriteString(str)
	_, err = o.WriteString("\r\n")
	return err
}

func writeS(o *bufio.Writer, id byte, str string) (err error) {
	_ = o.WriteByte(id)
	_, _ = o.WriteString(str)
	_, err = o.WriteString("\r\n")
	return err
}

func writeN(o *bufio.Writer, id byte, n int) (err error) {
	_ = o.WriteByte(id)
	if n < 10 {
		_ = o.WriteByte(byte('0' + n))
	} else {
		for d := int(math.Pow10(int(math.Log10(float64(n))))); d > 0; d /= 10 {
			_ = o.WriteByte(byte('0' + n/d))
			n = n % d
		}
	}
	_, err = o.WriteString("\r\n")
	return err
}

func readNextMessage(i *bufio.Reader) (m RedisMessage, err error) {
	var attrs *RedisMessage
	var typ byte
	for {
		if typ, err = i.ReadByte(); err != nil {
			return RedisMessage{}, err
		}
		fn := readers[typ]
		if fn == nil {
			return RedisMessage{}, errors.New(unknownMessageType + strconv.Itoa(int(typ)))
		}
		if m, err = fn(i); err != nil {
			if err == errOldNull {
				return RedisMessage{typ: typeNull}, nil
			}
			return RedisMessage{}, err
		}
		m.typ = typ
		if m.typ == typeAttribute { // handle the attributes
			a := m     // clone the original m first, and then take address of the clone
			attrs = &a // to avoid go compiler allocating the m on heap which causing worse performance.
			m = RedisMessage{}
			continue
		}
		m.attrs = attrs
		return m, nil
	}
}

type streamreader struct {
	i    *bufio.Reader
	n    int64
	done func()
	once bool
}

func (r *streamreader) eof() {
	if !r.once && r.done != nil {
		r.once = true
		r.done()
	}
}

func (r *streamreader) read(buf []byte) (n int, err error) {
	if int64(len(buf)) > r.n {
		buf = buf[0:r.n]
	}
	n, err = r.i.Read(buf)
	if r.n -= int64(n); r.n == 0 && err == nil {
		_, err = r.i.Discard(2)
	}
	return
}

type blobreader struct {
	streamreader
}

func (r *blobreader) Read(buf []byte) (n int, err error) {
	if r.n == 0 {
		r.eof()
		return 0, io.EOF
	}
	return r.read(buf)
}

type chunkreader struct {
	streamreader
}

func (r *chunkreader) Read(buf []byte) (n int, err error) {
	if r.n == 0 {
		if _, err = r.i.Discard(1); err != nil { // discard the ';'
			return 0, err
		}
		if r.n, err = readI(r.i); err != nil {
			return 0, err
		}
		if r.n == 0 {
			r.n = -1
		}
	}
	if r.n == -1 {
		r.eof()
		return 0, io.EOF
	}
	return r.read(buf)
}

func nextStringReader(i *bufio.Reader, done func()) (io.Reader, error) {
	var typ byte
	var err error
	for {
		if typ, err = i.ReadByte(); err != nil {
			done()
			return nil, err
		}
		switch typ {
		case typeBlobString, typeVerbatimString:
			length, err := readI(i)
			if err != nil {
				if err == errChunked {
					return &chunkreader{streamreader{i: i, done: done}}, nil
				}
				done()
				return nil, err
			}
			if length == -1 {
				done()
				return nil, &RedisError{typ: typeNull}
			}
			return &blobreader{streamreader{i: i, n: length, done: done}}, nil
		default:
			_ = i.UnreadByte()
			m, err := readNextMessage(i)
			if err != nil {
				done()
				return nil, err
			}
			switch m.typ {
			case typeSimpleString, typeFloat, typeBigNumber:
				done()
				return strings.NewReader(m.string), nil
			case typeSimpleErr, typeBlobErr, typeNull:
				done()
				return nil, (*RedisError)(&m)
			case typeInteger, typeBool:
				done()
				return strings.NewReader(strconv.FormatInt(m.integer, 10)), nil
			case typePush, typeAttribute:
				continue
			default:
				panic("")
			}
		}
	}
}

func writeCmd(o *bufio.Writer, cmd []string) (err error) {
	err = writeN(o, '*', len(cmd))
	for _, m := range cmd {
		err = writeB(o, '$', m)
		// TODO: Can we set cmd[i] = "" here to allow GC to eagerly recycle memory?
		// Related: https://github.com/redis/rueidis/issues/364
	}
	return err
}

const (
	unexpectedNoCRLF   = "received unexpected simple string message ending without CRLF"
	unexpectedNumByte  = "received unexpected number byte: "
	unknownMessageType = "received unknown message type: "
)
