package proto

import (
	"bufio"
	"bytes"
	"math/rand"
	"strconv"
	"strings"
	"testing"
)

var generators = map[byte]func(i int64, f float64, str string) string{}

func init() {
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
			msg, err := ReadNextMessage(r)
			if err != nil {
				t.Fatalf("unexpected error %v", err)
			}
			if msg.Type != k {
				t.Fatalf("unexpected msg type, expected %v, got %v", k, msg.Type)
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
		if err := WriteCmd(o, cmd); err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		_ = o.Flush()
		if m, err := ReadNextMessage(bufio.NewReader(b)); err != nil {
			t.Fatalf("unexpected err %v", err)
		} else if m.Type != '*' {
			t.Fatalf("unexpected m.Type: expected *, got %v", m.Type)
		} else if len(m.Values) != len(cmd) {
			t.Fatalf("unexpected m.Values: expected %v, got %v", len(cmd), len(m.Values))
		} else {
			for i, v := range m.Values {
				if v.Type != '$' {
					t.Fatalf("unexpected v.Values: expected $, got %v", v.Type)
				}
				if v.String != cmd[i] {
					t.Fatalf("unexpected v.String\n expected %v \n got %v", cmd[i], v.String)
				}
			}
		}
	}
}
