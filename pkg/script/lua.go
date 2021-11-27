package script

import (
	"crypto/sha1"
	"encoding/hex"

	"github.com/rueian/rueidis/internal/proto"
)

type EvalFn func(script string, keys []string, args []string) proto.Result

type Lua struct {
	script string
	sha1   string

	eval    EvalFn
	evalSha EvalFn
}

func (s *Lua) Exec(keys, args []string) proto.Result {
	r := s.evalSha(s.sha1, keys, args)
	if err := r.RedisError(); err != nil && err.IsNoScript() {
		r = s.eval(s.script, keys, args)
	}
	return r
}

func NewLuaScript(body string, eval, evalSha EvalFn) *Lua {
	sum := sha1.Sum([]byte(body))
	return &Lua{
		script:  body,
		sha1:    hex.EncodeToString(sum[:]),
		eval:    eval,
		evalSha: evalSha,
	}
}
