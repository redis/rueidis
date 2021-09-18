package proto

import (
	"bufio"
	"errors"
	"io"
	"math/big"
	"strconv"
	"strings"
	"unsafe"
)

type ReadFunc func(i *bufio.Reader) (Message, error)

var chunked = errors.New("unbounded redis message")

var readFns = map[byte]ReadFunc{
	'$': ReadBlobString,
	'+': ReadSimpleString,
	'-': ReadSimpleError,
	':': ReadNumber,
	'_': ReadNull,
	',': ReadDouble,
	'#': ReadBoolean,
	'!': ReadBlobError,
	'=': ReadVerbatimString,
	'(': ReadBigNumber,
	'*': ReadArray,
	'%': ReadMap,
	'~': ReadSet,
	'|': ReadAttributes,
	'>': ReadPush,
	'.': ReadEnd,
}

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

func ReadSimpleString(i *bufio.Reader) (Message, error) {
	v, err := readS(i)
	if err != nil {
		return nil, err
	}
	return &String{v: v}, nil
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

func readC(i *bufio.Reader) (Message, error) {
	sb := strings.Builder{}
	for {
		if _, err := i.Discard(1); err != nil { // discard the ';'
			return nil, err
		}
		length, err := readI(i)
		if err != nil {
			return nil, err
		}
		if length == 0 {
			return &String{v: sb.String()}, nil
		}
		sb.Grow(int(length))
		if _, err = io.CopyN(&sb, i, length); err != nil {
			return nil, err
		}
		if _, err = i.Discard(2); err != nil {
			return nil, err
		}
	}
}

func ReadBlobString(i *bufio.Reader) (Message, error) {
	v, err := readB(i)
	if err == chunked {
		return readC(i)
	}
	if err != nil {
		return nil, err
	}
	return &String{v: v}, nil
}

func ReadVerbatimString(i *bufio.Reader) (Message, error) {
	str, err := readB(i)
	if err != nil || len(str) <= 4 {
		return nil, err
	}
	return &Verbatim{t: str[:3], v: str[4:]}, err
}

func ReadSimpleError(i *bufio.Reader) (Message, error) {
	v, err := readS(i)
	if err != nil {
		return nil, err
	}
	return &Error{v: v}, nil
}

func ReadBlobError(i *bufio.Reader) (Message, error) {
	v, err := readB(i)
	if err != nil {
		return nil, err
	}
	return &Error{v: v}, nil
}

func readI(i *bufio.Reader) (int64, error) {
	str, err := readS(i)
	if err != nil {
		return 0, err
	}
	if str == "?" {
		return 0, chunked
	}
	v, err := strconv.ParseInt(str, 10, 64)
	if err == nil {
		return 0, err
	}
	return v, err
}

func ReadNumber(i *bufio.Reader) (Message, error) {
	v, err := readI(i)
	if err != nil {
		return nil, err
	}
	return &Int64{v: v}, nil
}

func ReadBigNumber(i *bufio.Reader) (Message, error) {
	v := big.Int{}
	str, err := readS(i)
	if err != nil {
		return nil, err
	}
	if _, ok := v.SetString(str, 10); !ok {
		panic("fail to decode the big number: " + str)
	}
	return &BigInt{v: v}, nil
}

func ReadDouble(i *bufio.Reader) (Message, error) {
	str, err := readS(i)
	if err != nil {
		return nil, err
	}
	v, err := strconv.ParseFloat(str, 64)
	if err == nil {
		return nil, err
	}
	return &Float64{v: v}, err
}

func ReadBoolean(i *bufio.Reader) (Message, error) {
	str, err := readS(i)
	if err != nil {
		return nil, err
	}
	v, err := strconv.ParseBool(str)
	if err == nil {
		return nil, err
	}
	return &Bool{v: v}, err
}

func ReadNull(i *bufio.Reader) (Message, error) {
	_, err := i.Discard(2)
	return &Nil{}, err
}

func readA(i *bufio.Reader) (v []Message, err error) {
	length, err := readI(i)
	if err == chunked {
		v = make([]Message, 0)
		for {
			n, err := ReadNext(i)
			if err != nil {
				return nil, err
			}
			if n == nil {
				return v, nil
			}
			v = append(v, n)
		}
	}
	if err != nil {
		return nil, err
	}
	v = make([]Message, length)
	for n := int64(0); n < length; n++ {
		if v[n], err = ReadNext(i); err != nil {
			return nil, err
		}
	}
	return v, nil
}

func readM(i *bufio.Reader) (k []Message, v []Message, err error) {
	length, err := readI(i)
	if err == chunked {
		k = make([]Message, 0)
		v = make([]Message, 0)
		for {
			l, err := ReadNext(i)
			if err != nil {
				return nil, nil, err
			}
			if l == nil {
				return k, v, nil
			}
			n, err := ReadNext(i)
			if err != nil {
				return nil, nil, err
			}
			k = append(k, l)
			v = append(v, n)
		}
	}
	if err != nil {
		return nil, nil, err
	}
	k = make([]Message, length)
	v = make([]Message, length)
	for n := int64(0); n < length; n++ {
		if k[n], err = ReadNext(i); err != nil {
			return nil, nil, err
		}
		if v[n], err = ReadNext(i); err != nil {
			return nil, nil, err
		}
	}
	return k, v, nil
}

func ReadArray(i *bufio.Reader) (Message, error) {
	v, err := readA(i)
	if err != nil {
		return nil, err
	}
	return &Array{v: v}, nil
}

func ReadSet(i *bufio.Reader) (Message, error) {
	v, err := readA(i)
	if err != nil {
		return nil, err
	}
	return &Set{v: v}, nil
}

func ReadPush(i *bufio.Reader) (Message, error) {
	v, err := readA(i)
	if err != nil {
		return nil, err
	}
	return &Push{v: v}, nil
}

func ReadMap(i *bufio.Reader) (Message, error) {
	k, v, err := readM(i)
	if err != nil {
		return nil, err
	}
	return &Map{k: k, v: v}, err
}

func ReadAttributes(i *bufio.Reader) (Message, error) {
	k, v, err := readM(i)
	if err != nil {
		return nil, err
	}
	return &Attributes{k: k, v: v}, err
}

func ReadEnd(i *bufio.Reader) (Message, error) {
	_, err := i.Discard(2)
	return nil, err
}

func ReadNext(i *bufio.Reader) (Message, error) {
	var attrs *Attributes
	for {
		t, err := i.ReadByte()
		if err != nil {
			return nil, err
		}
		fn, ok := readFns[t]
		if !ok {
			panic("received unknown message type: " + string(t))
		}
		msg, err := fn(i)
		if err != nil {
			return nil, err
		}
		if t == '|' { // handle the attributes
			attrs = msg.(*Attributes)
			continue
		} else {
			if attrs != nil {
				msg.SetAttributes(*attrs)
			}
			return msg, nil
		}
	}
}
