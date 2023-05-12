// Code generated DO NOT EDIT

package cmds

import "testing"

func cms0(s Builder) {
	s.CmsIncrby().Key("1").Item("1").Increment(1).Build()
	s.CmsInfo().Key("1").Build()
	s.CmsInfo().Key("1").Cache()
	s.CmsInitbydim().Key("1").Width(1).Depth(1).Build()
	s.CmsInitbyprob().Key("1").Error(1).Probability(1).Build()
	s.CmsMerge().Destination("1").Numkeys(1).Source("1").Source("1").Weights().Weight(1).Weight(1).Build()
	s.CmsMerge().Destination("1").Numkeys(1).Source("1").Source("1").Build()
	s.CmsQuery().Key("1").Item("1").Item("1").Build()
	s.CmsQuery().Key("1").Item("1").Item("1").Cache()
}

func TestCommand_InitSlot_cms(t *testing.T) {
	var s = NewBuilder(InitSlot)
	cms0(s)
}

func TestCommand_NoSlot_cms(t *testing.T) {
	var s = NewBuilder(NoSlot)
	cms0(s)
}
