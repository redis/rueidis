// Code generated DO NOT EDIT

package cmds

import "testing"

func graph0(s Builder) {
	s.GraphConfigGet().Name("1").Build()
	s.GraphConfigSet().Name("1").Value("1").Build()
	s.GraphConstraintCreate().Build()
	s.GraphConstraintDrop().Build()
	s.GraphDelete().Graph("1").Build()
	s.GraphExplain().Graph("1").Query("1").Build()
	s.GraphList().Build()
	s.GraphProfile().Graph("1").Query("1").Timeout(1).Build()
	s.GraphProfile().Graph("1").Query("1").Build()
	s.GraphQuery().Graph("1").Query("1").Timeout(1).Build()
	s.GraphQuery().Graph("1").Query("1").Build()
	s.GraphRoQuery().Graph("1").Query("1").Timeout(1).Build()
	s.GraphRoQuery().Graph("1").Query("1").Timeout(1).Cache()
	s.GraphRoQuery().Graph("1").Query("1").Build()
	s.GraphRoQuery().Graph("1").Query("1").Cache()
	s.GraphSlowlog().Graph("1").Build()
}

func TestCommand_InitSlot_graph(t *testing.T) {
	var s = NewBuilder(InitSlot)
	graph0(s)
}

func TestCommand_NoSlot_graph(t *testing.T) {
	var s = NewBuilder(NoSlot)
	graph0(s)
}
