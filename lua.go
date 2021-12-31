package rueidis

import (
	"context"
	"crypto/sha1"
	"encoding/hex"

	"github.com/rueian/rueidis/internal/proto"
)

func NewLuaScript(script string) *Lua {
	sum := sha1.Sum([]byte(script))
	return &Lua{script: script, sha1: hex.EncodeToString(sum[:])}
}

func NewLuaScriptReadOnly(script string) *Lua {
	lua := NewLuaScript(script)
	lua.readonly = true
	return lua
}

type Lua struct {
	script   string
	sha1     string
	readonly bool
}

func (s *Lua) Exec(ctx context.Context, c Client, keys, args []string) (resp proto.Result) {
	if s.readonly {
		resp = c.Do(ctx, c.B().EvalshaRo().Sha1(s.sha1).Numkeys(int64(len(keys))).Key(keys...).Arg(args...).Build())
	} else {
		resp = c.Do(ctx, c.B().Evalsha().Sha1(s.sha1).Numkeys(int64(len(keys))).Key(keys...).Arg(args...).Build())
	}
	if err := resp.RedisError(); err != nil && err.IsNoScript() {
		if s.readonly {
			resp = c.Do(ctx, c.B().EvalRo().Script(s.script).Numkeys(int64(len(keys))).Key(keys...).Arg(args...).Build())
		} else {
			resp = c.Do(ctx, c.B().Eval().Script(s.script).Numkeys(int64(len(keys))).Key(keys...).Arg(args...).Build())
		}
	}
	return resp
}
