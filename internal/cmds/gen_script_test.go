// Code generated DO NOT EDIT

package cmds

import "testing"

func script0(s Builder) {
	s.AiScriptdel().Key("1").Build()
	s.AiScriptget().Key("1").Meta().Source().Build()
	s.AiScriptget().Key("1").Meta().Source().Cache()
	s.AiScriptget().Key("1").Meta().Build()
	s.AiScriptget().Key("1").Meta().Cache()
	s.AiScriptget().Key("1").Source().Build()
	s.AiScriptget().Key("1").Source().Cache()
	s.AiScriptget().Key("1").Build()
	s.AiScriptget().Key("1").Cache()
	s.AiScriptstore().Key("1").Cpu().Tag("1").EntryPoints(1).EntryPoint("1").EntryPoint("1").Build()
	s.AiScriptstore().Key("1").Cpu().EntryPoints(1).EntryPoint("1").EntryPoint("1").Build()
	s.AiScriptstore().Key("1").Gpu().Tag("1").EntryPoints(1).EntryPoint("1").EntryPoint("1").Build()
	s.AiScriptstore().Key("1").Gpu().EntryPoints(1).EntryPoint("1").EntryPoint("1").Build()
}

func TestCommand_InitSlot_script(t *testing.T) {
	var s = NewBuilder(InitSlot)
	script0(s)
}

func TestCommand_NoSlot_script(t *testing.T) {
	var s = NewBuilder(NoSlot)
	script0(s)
}
