package rueidis

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"
)

const iteration = 1000

var generators = map[byte]func(i int64, f float64, str string) string{}

//gocyclo:ignore
func init() {
	rand.Seed(time.Now().UnixNano())

	generators['$'] = func(i int64, f float64, str string) string {
		return strconv.Itoa(len(str)) + "\r\n" + str + "\r\n"
	}
	generators['!'] = generators['$']
	generators['='] = generators['$']
	generators['+'] = func(i int64, f float64, str string) string {
		return str + "\r\n"
	}
	generators['-'] = generators['+']
	generators['('] = generators['+'] // big number as string
	generators[':'] = func(i int64, f float64, str string) string {
		return strconv.FormatInt(i, 10) + "\r\n"
	}
	generators['_'] = func(i int64, f float64, str string) string {
		return "\r\n"
	}
	generators[','] = func(i int64, f float64, str string) string {
		return strconv.FormatFloat(f, 'f', -1, 64) + "\r\n"
	}
	generators['#'] = func(i int64, f float64, str string) string {
		if i%2 == 1 {
			return "t\r\n"
		}
		return "f\r\n"
	}
	generators['*'] = func(i int64, f float64, str string) string {
		l := i%20 + 1
		if l == 0 {
			l = 1
		}
		if l < 0 {
			l *= -1
		}
		sb := strings.Builder{}
		sb.WriteString(strconv.FormatInt(l, 10))
		sb.WriteString("\r\n")
		for {
			for k, g := range generators {
				if k == '*' || k == '%' || k == '~' || k == '>' {
					continue
				}
				sb.WriteByte(k)
				sb.WriteString(g(i, f, random(k == '+' || k == '-' || k == '(')))
				l--
				if l == 0 {
					return sb.String()
				}
			}
		}
	}
	generators['%'] = func(i int64, f float64, str string) string {
		l := i % 20
		if l == 0 {
			l = 1
		}
		if l < 0 {
			l *= -1
		}
		sb := strings.Builder{}
		sb.WriteString(strconv.FormatInt(l, 10))
		sb.WriteString("\r\n")
		l *= 2
		for {
			for k, g := range generators {
				if k == '*' || k == '%' || k == '~' || k == '>' {
					continue
				}
				sb.WriteByte(k)
				sb.WriteString(g(i, f, random(k == '+' || k == '-' || k == '(')))
				l--
				if l == 0 {
					return sb.String()
				}
			}
		}
	}
	generators['~'] = generators['*']
	generators['>'] = generators['*']
}

func TestReadNextMessage(t *testing.T) {
	b := bytes.NewBuffer(nil)
	r := bufio.NewReader(b)

	for i := 0; i < iteration; i++ {
		for k, g := range generators {
			b.WriteByte(k)
			b.WriteString(g(rand.Int63(), rand.Float64(), random(k == '+' || k == '-' || k == '(')))
			msg, err := readNextMessage(r)
			if err != nil {
				t.Fatalf("unexpected error %v", err)
			}
			if msg.typ != k {
				t.Fatalf("unexpected msg type, expected %v, got %v", k, msg.typ)
			}
			// TODO test msg value
		}
	}
}

func TestWriteCmdAndRead(t *testing.T) {
	for i := 0; i < iteration; i++ {
		b := bytes.NewBuffer(nil)
		o := bufio.NewWriter(b)
		cmd := make([]string, randN(20))
		for i := range cmd {
			cmd[i] = random(false)
		}
		if err := writeCmd(o, cmd); err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		_ = o.Flush()
		if m, err := readNextMessage(bufio.NewReader(b)); err != nil {
			t.Fatalf("unexpected err %v", err)
		} else if m.typ != '*' {
			t.Fatalf("unexpected m.typ: expected *, got %v", m.typ)
		} else if len(m.values) != len(cmd) {
			t.Fatalf("unexpected m.values: expected %v, got %v", len(cmd), len(m.values))
		} else {
			for i, v := range m.values {
				if v.typ != '$' {
					t.Fatalf("unexpected v.values: expected $, got %v", v.typ)
				}
				if v.string != cmd[i] {
					t.Fatalf("unexpected v.string\n expected %v \n got %v", cmd[i], v.string)
				}
			}
		}
	}
}

func TestReadI(t *testing.T) {
	for i := 0; i < iteration; i++ {
		int1 := rand.Int63() - rand.Int63()
		int2, err := readI(source(strconv.FormatInt(int1, 10)))
		if err != nil {
			t.Fatalf("unexpected err: %v", err)
		}
		if int1 != int2 {
			t.Fatalf("readI fail to read the int: \n expected: %v \n got: %v", int1, int2)
		}
	}
}

func TestReadChunkedString(t *testing.T) {
	m, err := readNextMessage(bufio.NewReader(strings.NewReader("$?\r\n;4\r\nHell\r\n;5\r\no wor\r\n;1\r\nd\r\n;0\r\n")))
	if err != nil {
		t.Fatal(err)
	}
	if m.typ != '$' {
		t.Fatalf("unexpected msg type %v", m.typ)
	}
	if m.string != "Hello word" {
		t.Fatalf("unexpected msg string %v", m.string)
	}
}

func TestReadChunkedArray(t *testing.T) {
	data := "*?\r\n:1\r\n:2\r\n:3\r\n.\r\n"
	data += data
	in := bufio.NewReader(strings.NewReader(data))

	for i := 0; i < 2; i++ {
		m, err := readNextMessage(in)
		if err != nil {
			t.Fatal(err)
		}
		if m.typ != '*' {
			t.Fatalf("unexpected msg type %v", m.typ)
		}
		if len(m.values) != 3 {
			t.Fatalf("unexpected msg values length %v", len(m.values))
		}
		for i, v := range m.values {
			if v.typ != ':' || v.integer != int64(i+1) {
				t.Fatalf("unexpected msg values %v", m.values)
			}
		}
	}
}

func TestReadChunkedMap(t *testing.T) {
	data := "%?\r\n:1\r\n:2\r\n:3\r\n:4\r\n.\r\n"
	data += data
	in := bufio.NewReader(strings.NewReader(data))

	for i := 0; i < 2; i++ {
		m, err := readNextMessage(in)
		if err != nil {
			t.Fatal(err)
		}
		if m.typ != '%' {
			t.Fatalf("unexpected msg type %v", m.typ)
		}
		if len(m.values) != 4 {
			t.Fatalf("unexpected msg values length %v", len(m.values))
		}
		for i, v := range m.values {
			if v.typ != ':' || v.integer != int64(i+1) {
				t.Fatalf("unexpected msg values %v", m.values)
			}
		}
	}
}

// https://github.com/redis/redis-specifications/blob/master/protocol/RESP3.md#attribute-type
func TestReadAttr(t *testing.T) {
	data := "|1\r\n+key-popularity\r\n%2\r\n$1\r\na\r\n,0.1923\r\n$1\r\nb\r\n,0.0012\r\n*2\r\n:2039123\r\n:9543892\r\n"
	data += data
	in := bufio.NewReader(strings.NewReader(data))

	for i := 0; i < 2; i++ {
		m, err := readNextMessage(in)
		if err != nil {
			t.Fatal(err)
		}
		if m.typ != '*' {
			t.Fatalf("unexpected msg type %v", m.typ)
		}
		if m.values[0].integer != 2039123 {
			t.Fatalf("unexpected msg values[0] %v", m.values[0])
		}
		if m.values[1].integer != 9543892 {
			t.Fatalf("unexpected msg values[0] %v", m.values[1])
		}
		if !reflect.DeepEqual(*m.attrs, RedisMessage{typ: '|', values: []RedisMessage{
			{typ: '+', string: "key-popularity"},
			{typ: '%', values: []RedisMessage{
				{typ: '$', string: "a"},
				{typ: ',', string: "0.1923"},
				{typ: '$', string: "b"},
				{typ: ',', string: "0.0012"},
			}},
		}}) {
			t.Fatalf("unexpected msg attr %v", m.attrs)
		}
	}
}

func TestWriteBReadB(t *testing.T) {
	TWriterAndReader(t, writeB, readB, false)
}

func TestWriteSReadS(t *testing.T) {
	TWriterAndReader(t, writeS, readS, true)
}

func TWriterAndReader(t *testing.T, writer func(*bufio.Writer, byte, string) error, reader func(*bufio.Reader) (string, error), trim bool) {
	for i := 0; i < iteration; i++ {
		b := bytes.NewBuffer(nil)
		o := bufio.NewWriter(b)
		str1 := random(trim)
		if err := writer(o, str1[0], str1); err != nil {
			t.Fatalf("unexpected err: %v", err)
		}
		_ = o.Flush()
		r := bufio.NewReader(b)
		if id, err := r.ReadByte(); err != nil {
			t.Fatalf("unexpected err: %v", err)
		} else if id != str1[0] {
			t.Fatalf("unexpected id: expected %v, got %v", str1[0], id)
		}
		if str2, err := reader(r); err != nil {
			t.Fatalf("unexpected err: %v", err)
		} else if str1 != str2 {
			t.Fatalf("fail to read the string: \n expected: %v \n got: %v", str1, str2)
		}
	}
}

func TestRand(t *testing.T) {
	read := func(in *bufio.Reader) (m RedisMessage, err error) {
		defer func() {
			if err == nil {
				rec := recover()
				msg, ok := rec.(string)
				if ok && !strings.HasPrefix(msg, unexpectedNoCRLF) &&
					!strings.HasPrefix(msg, unexpectedNumByte) &&
					!strings.HasPrefix(msg, unknownMessageType) {
					t.Fatalf("unexpected panic %v", msg)
				}
				err = errors.New("panic as expected")
			}
		}()
		m, err = readNextMessage(in)
		return
	}
	for i := 0; i < iteration; i++ {
		if _, err := read(bufio.NewReader(strings.NewReader(random(false)))); err != nil {
			if err != io.EOF && err.Error() != "panic as expected" && err.Error() != "unbounded redis message" {
				t.Fatalf("unexpected err %v", err)
			}
		}
	}
}

func TestChunkedStringRand(t *testing.T) {
	chunkedPrefix := "$?\n;"

	read := func(in *bufio.Reader) (m RedisMessage, err error) {
		defer func() {
			if err == nil {
				rec := recover()
				msg, ok := rec.(string)
				if ok && !strings.HasPrefix(msg, unexpectedNoCRLF) &&
					!strings.HasPrefix(msg, unexpectedNumByte) &&
					!strings.HasPrefix(msg, unknownMessageType) {
					t.Fatalf("unexpected panic %v", msg)
				}
				err = errors.New("panic as expected")
			}
		}()
		m, err = readNextMessage(in)
		return
	}

	for i := 0; i < iteration; i++ {
		if m, err := read(bufio.NewReader(strings.NewReader(chunkedPrefix + random(false)))); err == nil {
			t.Fatalf("unexpected no err %v", m)
		} else if err != io.EOF && err != errChunked && err.Error() != "panic as expected" {
			t.Fatalf("unexpected err %v", err)
		}
	}
}

func source(str string) *bufio.Reader {
	return bufio.NewReader(bytes.NewReader(append([]byte(str), '\r', '\n')))
}

func random(trim bool) string {
retry:
	bs := make([]byte, randN(5000))
	if _, err := rand.Read(bs); err != nil {
		panic(err)
	}
	if trim {
		if v := strings.NewReplacer("\r", "", "\n", "").Replace(string(bs)); len(v) != 0 {
			return v
		}
		goto retry
	}
	return string(bs)
}

func randN(n int) (v int) {
	for v == 0 {
		v = rand.Intn(n)
	}
	return
}