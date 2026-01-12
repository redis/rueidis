package resp

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"unsafe"

	"github.com/redis/rueidis"
)

type Kind byte

const (
	KindString Kind = iota
	KindInt
	KindArray
	KindMap
	KindBool
	KindNull
)

type Reader struct {
	r *bufio.Reader
}

func NewReader(r *bufio.Reader) *Reader {
	return &Reader{r: r}
}

func (r *Reader) PeekKind() Kind {
	b, _ := r.r.Peek(1)
	switch b[0] {
	case '$', '=', '+':
		return KindString
	case ':':
		return KindInt
	case '*', '~':
		return KindArray
	case '%':
		return KindMap
	case '#':
		return KindBool
	case '_':
		return KindNull
	default:
		return KindNull
	}
}

func (r *Reader) ExpectArray() (int64, error) {
	if err := r.expectArrayType(); err != nil {
		return 0, err
	}
	return r.readArrayLen()
}

func (r *Reader) expectArrayType() error {
	b, err := r.r.ReadByte()
	if err != nil {
		return err
	}
	if b != '*' && b != '~' {
		return fmt.Errorf("expected array, got %q", b)
	}
	return nil
}

func (r *Reader) readArrayLen() (int64, error) {
	line, err := r.readLine()
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(unsafe.String(&line[0], len(line)), 10, 64)
}

// ExpectArrayWithLen reads an array header from r and checks that its length matches `expected`.
// expected must be >= 0. Returns an error if the type is not array or the length mismatches.
func (r *Reader) ExpectArrayWithLen(expected int64) error {
	if expected < 0 {
		panic("ExpectArrayWithLen must be used only with fixed-length arrays")
	}
	count, err := r.ExpectArray()
	if err != nil {
		return err
	}
	if count != expected {
		return fmt.Errorf("expected array of length %d, got %d", expected, count)
	}
	return nil
}

func (r *Reader) ReadInt64() (int64, error) {
	if b, err := r.r.ReadByte(); err != nil {
		return 0, err
	} else if b != ':' {
		return 0, fmt.Errorf("expected int, got %q", b)
	}

	line, err := r.readLine()
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(unsafe.String(&line[0], len(line)), 10, 64)
}

func (r *Reader) ReadStringBytes() ([]byte, error) {
	b, err := r.r.ReadByte()
	if err != nil {
		return nil, err
	}

	switch b {
	case '$', '=', '!': // Blob strings / verbatim / blob error
		length, err := rueidis.ReadInt(r.r)
		if err != nil {
			return nil, err
		}
		return readBlobZeroAlloc(r.r, length)
	case '+': // Simple string
		line, err := r.readLine()
		if err != nil {
			return nil, err
		}
		return line, nil
	case '_': // Null
		if err := r.ReadNull(); err != nil {
			return nil, err
		}
		return nil, rueidis.Nil

	default:
		return nil, fmt.Errorf("unexpected RESP type %q", b)
	}
}

// readBlobZeroAlloc reads a RESP blob payload of the given length.
// It returns a slice backed by bufio.Reader internal memory.
// The slice is valid only until the next read/discard.
//
// It consumes: <payload>\r\n
//
// Intended for use inside DoWithReader callbacks.
func readBlobZeroAlloc(r *bufio.Reader, length int64) ([]byte, error) {
	if length < 0 {
		return nil, rueidis.Nil
	}

	if length == 0 {
		_, err := r.Discard(2) // \r\n
		return nil, err
	}

	b, err := r.Peek(int(length))
	if err != nil {
		return nil, err
	}

	// discard payload + CRLF
	if _, err := r.Discard(int(length) + 2); err != nil {
		return nil, err
	}

	return b, nil
}

func (r *Reader) SkipValue() error {
	typ, err := r.r.ReadByte()
	if err != nil {
		return err
	}

	switch typ {

	case '+', '-', ':', ',', '(', '#':
		// Simple string, error, integer, double, big number, boolean
		for {
			b, err := r.r.ReadByte()
			if err != nil {
				return err
			}
			if b == '\n' {
				return nil
			}
		}

	case '$', '!', '=':
		// Blob string / blob error / verbatim
		length, err := rueidis.ReadInt(r.r)
		if err != nil {
			return err
		}
		if length < 0 {
			return nil
		}
		_, err = r.r.Discard(int(length) + 2)
		return err

	case '*', '~':
		// Array / set
		count, err := rueidis.ReadInt(r.r)
		if err != nil {
			return err
		}
		if count < 0 {
			return nil
		}
		for i := int64(0); i < count; i++ {
			if err := r.SkipValue(); err != nil {
				return err
			}
		}
		return nil

	case '%':
		// Map: key + value
		count, err := rueidis.ReadInt(r.r)
		if err != nil {
			return err
		}
		if count < 0 {
			return nil
		}
		for i := int64(0); i < count*2; i++ {
			if err := r.SkipValue(); err != nil {
				return err
			}
		}
		return nil

	case '_':
		// Null
		_, err := r.readLine()
		return err

	default:
		return fmt.Errorf("unexpected RESP type: %q", typ)
	}
}

var ErrNotNull = errors.New("resp: value is not null")

func (r *Reader) ReadNull() error {
	b, err := r.r.ReadByte()
	if err != nil {
		return err
	}

	switch b {

	// RESP3 null
	case '_':
		// consume CRLF
		if err := r.discardLine(); err != nil {
			return err
		}
		return nil

	// RESP2 null bulk / null array
	case '$', '*':
		n, err := rueidis.ReadInt(r.r)
		if err != nil {
			return err
		}
		if n != -1 {
			return ErrNotNull
		}
		return nil

	default:
		return ErrNotNull
	}
}

func (r *Reader) discardLine() error {
	for {
		b, err := r.r.ReadByte()
		if err != nil {
			return err
		}
		if b == '\n' {
			return nil
		}
	}
}

func (r *Reader) readLine() ([]byte, error) {
	b, err := r.r.ReadSlice('\n')
	if err != nil {
		return nil, err
	}
	if len(b) < 2 || b[len(b)-2] != '\r' {
		return nil, errors.New("invalid RESP line ending")
	}
	return b[:len(b)-2], nil // strip CRLF
}
