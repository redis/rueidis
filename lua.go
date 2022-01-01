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

// Lua represents a redis lua script. It should be created from the NewLuaScript() or NewLuaScriptReadOnly()
type Lua struct {
	script   string
	sha1     string
	readonly bool
}

// Exec the script to the given Client.
// It will first try with the EVALSHA/EVALSHA_RO and then EVAL/EVAL_RO if first try failed.
// Cross slot keys are prohibited if the Client is a cluster client.
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
