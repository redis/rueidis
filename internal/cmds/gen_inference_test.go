// Code generated DO NOT EDIT

package cmds

import "testing"

func inference0(s Builder) {
	s.AiModelexecute().Key("1").Inputs(1).Input("1").Input("1").Outputs(1).Output("1").Output("1").Timeout(1).Build()
	s.AiModelexecute().Key("1").Inputs(1).Input("1").Input("1").Outputs(1).Output("1").Output("1").Timeout(1).Cache()
	s.AiModelexecute().Key("1").Inputs(1).Input("1").Input("1").Outputs(1).Output("1").Output("1").Build()
	s.AiModelexecute().Key("1").Inputs(1).Input("1").Input("1").Outputs(1).Output("1").Output("1").Cache()
	s.AiScriptexecute().Key("1").Function("1").Keys(1).Key("1").Key("1").Inputs(1).Input("1").Input("1").Timeout(1).Build()
	s.AiScriptexecute().Key("1").Function("1").Keys(1).Key("1").Key("1").Inputs(1).Input("1").Input("1").Build()
	s.AiScriptexecute().Key("1").Function("1").Keys(1).Key("1").Key("1").Args(1).Arg("1").Arg("1").Timeout(1).Build()
	s.AiScriptexecute().Key("1").Function("1").Keys(1).Key("1").Key("1").Args(1).Arg("1").Arg("1").Build()
	s.AiScriptexecute().Key("1").Function("1").Keys(1).Key("1").Key("1").Outputs(1).Output("1").Output("1").Timeout(1).Build()
	s.AiScriptexecute().Key("1").Function("1").Keys(1).Key("1").Key("1").Outputs(1).Output("1").Output("1").Build()
	s.AiScriptexecute().Key("1").Function("1").Keys(1).Key("1").Key("1").Timeout(1).Build()
	s.AiScriptexecute().Key("1").Function("1").Keys(1).Key("1").Key("1").Build()
	s.AiScriptexecute().Key("1").Function("1").Inputs(1).Input("1").Input("1").Args(1).Arg("1").Arg("1").Timeout(1).Build()
	s.AiScriptexecute().Key("1").Function("1").Inputs(1).Input("1").Input("1").Args(1).Arg("1").Arg("1").Build()
	s.AiScriptexecute().Key("1").Function("1").Inputs(1).Input("1").Input("1").Outputs(1).Output("1").Output("1").Timeout(1).Build()
	s.AiScriptexecute().Key("1").Function("1").Inputs(1).Input("1").Input("1").Outputs(1).Output("1").Output("1").Build()
	s.AiScriptexecute().Key("1").Function("1").Inputs(1).Input("1").Input("1").Timeout(1).Build()
	s.AiScriptexecute().Key("1").Function("1").Inputs(1).Input("1").Input("1").Build()
	s.AiScriptexecute().Key("1").Function("1").Args(1).Arg("1").Arg("1").Outputs(1).Output("1").Output("1").Timeout(1).Build()
	s.AiScriptexecute().Key("1").Function("1").Args(1).Arg("1").Arg("1").Outputs(1).Output("1").Output("1").Build()
	s.AiScriptexecute().Key("1").Function("1").Args(1).Arg("1").Arg("1").Timeout(1).Build()
	s.AiScriptexecute().Key("1").Function("1").Args(1).Arg("1").Arg("1").Build()
	s.AiScriptexecute().Key("1").Function("1").Outputs(1).Output("1").Output("1").Timeout(1).Build()
	s.AiScriptexecute().Key("1").Function("1").Outputs(1).Output("1").Output("1").Build()
	s.AiScriptexecute().Key("1").Function("1").Timeout(1).Build()
	s.AiScriptexecute().Key("1").Function("1").Build()
}

func TestCommand_InitSlot_inference(t *testing.T) {
	var s = NewBuilder(InitSlot)
	inference0(s)
}

func TestCommand_NoSlot_inference(t *testing.T) {
	var s = NewBuilder(NoSlot)
	inference0(s)
}
