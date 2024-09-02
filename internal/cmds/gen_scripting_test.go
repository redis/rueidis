// Code generated by go generate; DO NOT EDIT

package cmds

import "testing"

func scripting0(s Builder) {
	s.Eval().Script("1").Numkeys(1).Key("1").Key("1").Arg("1").Arg("1").Build()
	s.Eval().Script("1").Numkeys(1).Key("1").Key("1").Build()
	s.Eval().Script("1").Numkeys(1).Arg("1").Arg("1").Build()
	s.Eval().Script("1").Numkeys(1).Build()
	s.EvalRo().Script("1").Numkeys(1).Key("1").Key("1").Arg("1").Arg("1").Build()
	s.EvalRo().Script("1").Numkeys(1).Key("1").Key("1").Arg("1").Arg("1").Cache()
	s.EvalRo().Script("1").Numkeys(1).Key("1").Key("1").Build()
	s.EvalRo().Script("1").Numkeys(1).Key("1").Key("1").Cache()
	s.EvalRo().Script("1").Numkeys(1).Arg("1").Arg("1").Build()
	s.EvalRo().Script("1").Numkeys(1).Arg("1").Arg("1").Cache()
	s.EvalRo().Script("1").Numkeys(1).Build()
	s.EvalRo().Script("1").Numkeys(1).Cache()
	s.Evalsha().Sha1("1").Numkeys(1).Key("1").Key("1").Arg("1").Arg("1").Build()
	s.Evalsha().Sha1("1").Numkeys(1).Key("1").Key("1").Build()
	s.Evalsha().Sha1("1").Numkeys(1).Arg("1").Arg("1").Build()
	s.Evalsha().Sha1("1").Numkeys(1).Build()
	s.EvalshaRo().Sha1("1").Numkeys(1).Key("1").Key("1").Arg("1").Arg("1").Build()
	s.EvalshaRo().Sha1("1").Numkeys(1).Key("1").Key("1").Arg("1").Arg("1").Cache()
	s.EvalshaRo().Sha1("1").Numkeys(1).Key("1").Key("1").Build()
	s.EvalshaRo().Sha1("1").Numkeys(1).Key("1").Key("1").Cache()
	s.EvalshaRo().Sha1("1").Numkeys(1).Arg("1").Arg("1").Build()
	s.EvalshaRo().Sha1("1").Numkeys(1).Arg("1").Arg("1").Cache()
	s.EvalshaRo().Sha1("1").Numkeys(1).Build()
	s.EvalshaRo().Sha1("1").Numkeys(1).Cache()
	s.Fcall().Function("1").Numkeys(1).Key("1").Key("1").Arg("1").Arg("1").Build()
	s.Fcall().Function("1").Numkeys(1).Key("1").Key("1").Build()
	s.Fcall().Function("1").Numkeys(1).Arg("1").Arg("1").Build()
	s.Fcall().Function("1").Numkeys(1).Build()
	s.FcallRo().Function("1").Numkeys(1).Key("1").Key("1").Arg("1").Arg("1").Build()
	s.FcallRo().Function("1").Numkeys(1).Key("1").Key("1").Arg("1").Arg("1").Cache()
	s.FcallRo().Function("1").Numkeys(1).Key("1").Key("1").Build()
	s.FcallRo().Function("1").Numkeys(1).Key("1").Key("1").Cache()
	s.FcallRo().Function("1").Numkeys(1).Arg("1").Arg("1").Build()
	s.FcallRo().Function("1").Numkeys(1).Arg("1").Arg("1").Cache()
	s.FcallRo().Function("1").Numkeys(1).Build()
	s.FcallRo().Function("1").Numkeys(1).Cache()
	s.FunctionDelete().LibraryName("1").Build()
	s.FunctionDump().Build()
	s.FunctionFlush().Async().Build()
	s.FunctionFlush().Sync().Build()
	s.FunctionFlush().Build()
	s.FunctionHelp().Build()
	s.FunctionKill().Build()
	s.FunctionList().Libraryname("1").Withcode().Build()
	s.FunctionList().Libraryname("1").Build()
	s.FunctionList().Withcode().Build()
	s.FunctionList().Build()
	s.FunctionLoad().Replace().FunctionCode("1").Build()
	s.FunctionLoad().FunctionCode("1").Build()
	s.FunctionRestore().SerializedValue("1").Flush().Build()
	s.FunctionRestore().SerializedValue("1").Append().Build()
	s.FunctionRestore().SerializedValue("1").Replace().Build()
	s.FunctionRestore().SerializedValue("1").Build()
	s.FunctionStats().Build()
	s.ScriptDebug().Yes().Build()
	s.ScriptDebug().Sync().Build()
	s.ScriptDebug().No().Build()
	s.ScriptExists().Sha1("1").Sha1("1").Build()
	s.ScriptFlush().Async().Build()
	s.ScriptFlush().Sync().Build()
	s.ScriptFlush().Build()
	s.ScriptKill().Build()
	s.ScriptLoad().Script("1").Build()
	s.ScriptShow().Sha1("1").Build()
}

func TestCommand_InitSlot_scripting(t *testing.T) {
	var s = NewBuilder(InitSlot)
	t.Run("0", func(t *testing.T) { scripting0(s) })
}

func TestCommand_NoSlot_scripting(t *testing.T) {
	var s = NewBuilder(NoSlot)
	t.Run("0", func(t *testing.T) { scripting0(s) })
}
