// Code generated DO NOT EDIT

package cmds

import "testing"

func topk0(s Builder) {
	s.TopkAdd().Key("1").Items("1").Items("1").Build()
	s.TopkCount().Key("1").Item("1").Item("1").Build()
	s.TopkIncrby().Key("1").Item("1").Increment(1).Build()
	s.TopkInfo().Key("1").Build()
	s.TopkInfo().Key("1").Cache()
	s.TopkList().Key("1").Withcount().Build()
	s.TopkList().Key("1").Withcount().Cache()
	s.TopkList().Key("1").Build()
	s.TopkList().Key("1").Cache()
	s.TopkQuery().Key("1").Item("1").Item("1").Build()
	s.TopkQuery().Key("1").Item("1").Item("1").Cache()
	s.TopkReserve().Key("1").Topk(1).Width(1).Depth(1).Decay(1).Build()
	s.TopkReserve().Key("1").Topk(1).Build()
}

func TestCommand_InitSlot_topk(t *testing.T) {
	var s = NewBuilder(InitSlot)
	topk0(s)
}

func TestCommand_NoSlot_topk(t *testing.T) {
	var s = NewBuilder(NoSlot)
	topk0(s)
}
