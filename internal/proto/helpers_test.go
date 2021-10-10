package proto

import (
	"bufio"
	"bytes"
	"math/rand"
	"strconv"
	"strings"
	"testing"
	"time"
)

const iteration = 2000

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestReadI(t *testing.T) {
	for i := 0; i < iteration; i++ {
		int1 := rand.Int63() - rand.Int63()
		int2, err := ReadI(source(strconv.FormatInt(int1, 10)))
		if err != nil {
			t.Fatalf("unexpected err: %v", err)
		}
		if int1 != int2 {
			t.Fatalf("ReadI fail to read the int: \n expected: %v \n got: %v", int1, int2)
		}
	}
}

func TestWriteBReadB(t *testing.T) {
	TWriterAndReader(t, WriteB, ReadB, false)
}

func TestWriteSReadS(t *testing.T) {
	TWriterAndReader(t, WriteS, ReadS, true)
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

func source(str string) *bufio.Reader {
	return bufio.NewReader(bytes.NewReader(append([]byte(str), '\r', '\n')))
}

func random(trim bool) string {
	bs := make([]byte, randN(5000))
	if _, err := rand.Read(bs); err != nil {
		panic(err)
	}
	if trim {
		return strings.NewReplacer("\r", "", "\n", "").Replace(string(bs))
	}
	return string(bs)
}

func randN(n int) (v int) {
	for v == 0 {
		v = rand.Intn(n)
	}
	return
}
