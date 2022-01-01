package rueidis

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"math/rand"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/rueian/rueidis/internal/cmds"
	"github.com/rueian/rueidis/internal/proto"
)

func TestNewLuaScript(t *testing.T) {
	body := strconv.Itoa(rand.Int())
	sum := sha1.Sum([]byte(body))
	sha := hex.EncodeToString(sum[:])

	k := []string{"1", "2"}
	a := []string{"3", "4"}

	eval := false

	c := &client{
		BFn: func() *cmds.Builder {
			return cmds.NewBuilder(cmds.NoSlot)
		},
		DoFn: func(ctx context.Context, cmd cmds.Completed) (resp proto.Result) {
			if reflect.DeepEqual(cmd.Commands(), []string{"EVALSHA", sha, "2", "1", "2", "3", "4"}) {
				eval = true
				return proto.NewResult(proto.Message{Type: '-', String: "NOSCRIPT"}, nil)
			}
			if eval && reflect.DeepEqual(cmd.Commands(), []string{"EVAL", body, "2", "1", "2", "3", "4"}) {
				return proto.NewResult(proto.Message{Type: '_'}, nil)
			}
			return proto.NewResult(proto.Message{Type: '+', String: "unexpected"}, nil)
		},
	}

	script := NewLuaScript(body)

	if !script.Exec(context.Background(), c, k, a).RedisError().IsNil() {
		t.Fatalf("ret mistmatch")
	}
}

func TestNewLuaScriptReadOnly(t *testing.T) {
	body := strconv.Itoa(rand.Int())
	sum := sha1.Sum([]byte(body))
	sha := hex.EncodeToString(sum[:])

	k := []string{"1", "2"}
	a := []string{"3", "4"}

	eval := false

	c := &client{
		BFn: func() *cmds.Builder {
			return cmds.NewBuilder(cmds.NoSlot)
		},
		DoFn: func(ctx context.Context, cmd cmds.Completed) (resp proto.Result) {
			if reflect.DeepEqual(cmd.Commands(), []string{"EVALSHA_RO", sha, "2", "1", "2", "3", "4"}) {
				eval = true
				return proto.NewResult(proto.Message{Type: '-', String: "NOSCRIPT"}, nil)
			}
			if eval && reflect.DeepEqual(cmd.Commands(), []string{"EVAL_RO", body, "2", "1", "2", "3", "4"}) {
				return proto.NewResult(proto.Message{Type: '_'}, nil)
			}
			return proto.NewResult(proto.Message{Type: '+', String: "unexpected"}, nil)
		},
	}

	script := NewLuaScriptReadOnly(body)

	if !script.Exec(context.Background(), c, k, a).RedisError().IsNil() {
		t.Fatalf("ret mistmatch")
	}
}

type client struct {
	BFn         func() *cmds.Builder
	DoFn        func(ctx context.Context, cmd cmds.Completed) (resp proto.Result)
	DoCacheFn   func(ctx context.Context, cmd cmds.Cacheable, ttl time.Duration) (resp proto.Result)
	DedicatedFn func(fn func(DedicatedClient) error) (err error)
	CloseFn     func()
}

func (c *client) B() *cmds.Builder {
	if c.BFn != nil {
		return c.BFn()
	}
	return nil
}

func (c *client) Do(ctx context.Context, cmd cmds.Completed) (resp proto.Result) {
	if c.DoFn != nil {
		return c.DoFn(ctx, cmd)
	}
	return proto.Result{}
}

func (c *client) DoCache(ctx context.Context, cmd cmds.Cacheable, ttl time.Duration) (resp proto.Result) {
	if c.DoCacheFn != nil {
		return c.DoCacheFn(ctx, cmd, ttl)
	}
	return proto.Result{}
}

func (c *client) Dedicated(fn func(DedicatedClient) error) (err error) {
	if c.DedicatedFn != nil {
		return c.DedicatedFn(fn)
	}
	return nil
}

func (c *client) Close() {
	if c.CloseFn != nil {
		c.CloseFn()
	}
}

func ExampleLua_exec() {
	client, err := NewClient(ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
	if err != nil {
		panic(err)
	}
	defer client.Close()

	ctx := context.Background()

	script := NewLuaScript("return {KEYS[1],KEYS[2],ARGV[1],ARGV[2]}")

	script.Exec(ctx, client, []string{"k1", "k2"}, []string{"a1", "a2"}).ToArray()
}
