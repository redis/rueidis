package rueidis

import (
	"crypto/sha1"
	"encoding/hex"
	"math/rand"
	"reflect"
	"strconv"
	"testing"

	"github.com/rueian/rueidis/internal/proto"
)

func TestNewLuaScript(t *testing.T) {
	body := strconv.Itoa(rand.Int())
	sum := sha1.Sum([]byte(body))
	sha := hex.EncodeToString(sum[:])

	k := []string{"1", "2"}
	a := []string{"3", "4"}

	eval := false
	evalSha := false

	script := newLuaScript(body, func(in string, keys []string, args []string) proto.Result {
		eval = true
		if in != body {
			t.Fatalf("input should be %q", body)
		}
		if !reflect.DeepEqual(k, keys) || !reflect.DeepEqual(a, args) {
			t.Fatalf("parameter mistmatch")
		}
		return proto.NewResult(proto.Message{Type: '_'}, nil)
	}, func(in string, keys []string, args []string) proto.Result {
		evalSha = true
		if in != sha {
			t.Fatalf("input should be %q", sha)
		}
		if !reflect.DeepEqual(k, keys) || !reflect.DeepEqual(a, args) {
			t.Fatalf("parameter mistmatch")
		}
		return proto.NewResult(proto.Message{Type: '-', String: "NOSCRIPT"}, nil)
	})

	if !script.Exec(k, a).RedisError().IsNil() {
		t.Fatalf("ret mistmatch")
	}
	if !eval {
		t.Fatalf("eval fn not called")
	}
	if !evalSha {
		t.Fatalf("evalSha fn not called")
	}
}
