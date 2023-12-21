// Code generated by go generate; DO NOT EDIT

package cmds

import "testing"

func json0(s Builder) {
	s.JsonArrappend().Key("1").Path("1").Value("1").Value("1").Build()
	s.JsonArrappend().Key("1").Value("1").Value("1").Build()
	s.JsonArrindex().Key("1").Path("1").Value("1").Start(1).Stop(1).Build()
	s.JsonArrindex().Key("1").Path("1").Value("1").Start(1).Stop(1).Cache()
	s.JsonArrindex().Key("1").Path("1").Value("1").Start(1).Build()
	s.JsonArrindex().Key("1").Path("1").Value("1").Start(1).Cache()
	s.JsonArrindex().Key("1").Path("1").Value("1").Build()
	s.JsonArrindex().Key("1").Path("1").Value("1").Cache()
	s.JsonArrinsert().Key("1").Path("1").Index(1).Value("1").Value("1").Build()
	s.JsonArrlen().Key("1").Path("1").Build()
	s.JsonArrlen().Key("1").Path("1").Cache()
	s.JsonArrlen().Key("1").Build()
	s.JsonArrlen().Key("1").Cache()
	s.JsonArrpop().Key("1").Path("1").Index(1).Build()
	s.JsonArrpop().Key("1").Path("1").Build()
	s.JsonArrpop().Key("1").Build()
	s.JsonArrtrim().Key("1").Path("1").Start(1).Stop(1).Build()
	s.JsonClear().Key("1").Path("1").Build()
	s.JsonClear().Key("1").Build()
	s.JsonDebugHelp().Build()
	s.JsonDebugMemory().Key("1").Path("1").Build()
	s.JsonDebugMemory().Key("1").Build()
	s.JsonDel().Key("1").Path("1").Build()
	s.JsonDel().Key("1").Build()
	s.JsonForget().Key("1").Path("1").Build()
	s.JsonForget().Key("1").Build()
	s.JsonGet().Key("1").Indent("1").Newline("1").Space("1").Path("1").Path("1").Build()
	s.JsonGet().Key("1").Indent("1").Newline("1").Space("1").Path("1").Path("1").Cache()
	s.JsonGet().Key("1").Indent("1").Newline("1").Space("1").Build()
	s.JsonGet().Key("1").Indent("1").Newline("1").Space("1").Cache()
	s.JsonGet().Key("1").Indent("1").Newline("1").Path("1").Path("1").Build()
	s.JsonGet().Key("1").Indent("1").Newline("1").Path("1").Path("1").Cache()
	s.JsonGet().Key("1").Indent("1").Newline("1").Build()
	s.JsonGet().Key("1").Indent("1").Newline("1").Cache()
	s.JsonGet().Key("1").Indent("1").Space("1").Build()
	s.JsonGet().Key("1").Indent("1").Space("1").Cache()
	s.JsonGet().Key("1").Indent("1").Path("1").Path("1").Build()
	s.JsonGet().Key("1").Indent("1").Path("1").Path("1").Cache()
	s.JsonGet().Key("1").Indent("1").Build()
	s.JsonGet().Key("1").Indent("1").Cache()
	s.JsonGet().Key("1").Newline("1").Build()
	s.JsonGet().Key("1").Newline("1").Cache()
	s.JsonGet().Key("1").Space("1").Build()
	s.JsonGet().Key("1").Space("1").Cache()
	s.JsonGet().Key("1").Path("1").Path("1").Build()
	s.JsonGet().Key("1").Path("1").Path("1").Cache()
	s.JsonGet().Key("1").Build()
	s.JsonGet().Key("1").Cache()
	s.JsonMerge().Key("1").Path("1").Value("1").Build()
	s.JsonMget().Key("1").Key("1").Path("1").Build()
	s.JsonMget().Key("1").Key("1").Path("1").Cache()
	s.JsonMset().Key("1").Path("1").Value("1").Key("1").Path("1").Value("1").Build()
	s.JsonNumincrby().Key("1").Path("1").Value(1).Build()
	s.JsonNummultby().Key("1").Path("1").Value(1).Build()
	s.JsonObjkeys().Key("1").Path("1").Build()
	s.JsonObjkeys().Key("1").Path("1").Cache()
	s.JsonObjkeys().Key("1").Build()
	s.JsonObjkeys().Key("1").Cache()
	s.JsonObjlen().Key("1").Path("1").Build()
	s.JsonObjlen().Key("1").Path("1").Cache()
	s.JsonObjlen().Key("1").Build()
	s.JsonObjlen().Key("1").Cache()
	s.JsonResp().Key("1").Path("1").Build()
	s.JsonResp().Key("1").Path("1").Cache()
	s.JsonResp().Key("1").Build()
	s.JsonResp().Key("1").Cache()
	s.JsonSet().Key("1").Path("1").Value("1").Nx().Build()
	s.JsonSet().Key("1").Path("1").Value("1").Xx().Build()
	s.JsonSet().Key("1").Path("1").Value("1").Build()
	s.JsonStrappend().Key("1").Path("1").Value("1").Build()
	s.JsonStrappend().Key("1").Value("1").Build()
	s.JsonStrlen().Key("1").Path("1").Build()
	s.JsonStrlen().Key("1").Path("1").Cache()
	s.JsonStrlen().Key("1").Build()
	s.JsonStrlen().Key("1").Cache()
	s.JsonToggle().Key("1").Path("1").Build()
	s.JsonType().Key("1").Path("1").Build()
	s.JsonType().Key("1").Path("1").Cache()
	s.JsonType().Key("1").Build()
	s.JsonType().Key("1").Cache()
}

func TestCommand_InitSlot_json(t *testing.T) {
	var s = NewBuilder(InitSlot)
	t.Run("0", func(t *testing.T) { json0(s) })
}

func TestCommand_NoSlot_json(t *testing.T) {
	var s = NewBuilder(NoSlot)
	t.Run("0", func(t *testing.T) { json0(s) })
}
