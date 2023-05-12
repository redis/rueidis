// Code generated DO NOT EDIT

package cmds

import "testing"

func bitmap0(s Builder) {
	s.Bitcount().Key("1").Start(1).End(1).Byte().Build()
	s.Bitcount().Key("1").Start(1).End(1).Byte().Cache()
	s.Bitcount().Key("1").Start(1).End(1).Bit().Build()
	s.Bitcount().Key("1").Start(1).End(1).Bit().Cache()
	s.Bitcount().Key("1").Start(1).End(1).Build()
	s.Bitcount().Key("1").Start(1).End(1).Cache()
	s.Bitcount().Key("1").Build()
	s.Bitcount().Key("1").Cache()
	s.Bitfield().Key("1").Get("1", 1).OverflowWrap().Set("1", 1, 1).Incrby("1", 1, 1).Build()
	s.Bitfield().Key("1").Get("1", 1).OverflowWrap().Set("1", 1, 1).Build()
	s.Bitfield().Key("1").Get("1", 1).OverflowWrap().Incrby("1", 1, 1).Build()
	s.Bitfield().Key("1").Get("1", 1).OverflowSat().Set("1", 1, 1).Incrby("1", 1, 1).Build()
	s.Bitfield().Key("1").Get("1", 1).OverflowSat().Set("1", 1, 1).Build()
	s.Bitfield().Key("1").Get("1", 1).OverflowSat().Incrby("1", 1, 1).Build()
	s.Bitfield().Key("1").Get("1", 1).OverflowFail().Set("1", 1, 1).Incrby("1", 1, 1).Build()
	s.Bitfield().Key("1").Get("1", 1).OverflowFail().Set("1", 1, 1).Build()
	s.Bitfield().Key("1").Get("1", 1).OverflowFail().Incrby("1", 1, 1).Build()
	s.Bitfield().Key("1").Get("1", 1).Set("1", 1, 1).Incrby("1", 1, 1).Build()
	s.Bitfield().Key("1").Get("1", 1).Set("1", 1, 1).Build()
	s.Bitfield().Key("1").Get("1", 1).Incrby("1", 1, 1).Build()
	s.Bitfield().Key("1").OverflowWrap().Set("1", 1, 1).Incrby("1", 1, 1).Build()
	s.Bitfield().Key("1").OverflowWrap().Set("1", 1, 1).Build()
	s.Bitfield().Key("1").OverflowWrap().Incrby("1", 1, 1).Build()
	s.Bitfield().Key("1").OverflowSat().Set("1", 1, 1).Incrby("1", 1, 1).Build()
	s.Bitfield().Key("1").OverflowSat().Set("1", 1, 1).Build()
	s.Bitfield().Key("1").OverflowSat().Incrby("1", 1, 1).Build()
	s.Bitfield().Key("1").OverflowFail().Set("1", 1, 1).Incrby("1", 1, 1).Build()
	s.Bitfield().Key("1").OverflowFail().Set("1", 1, 1).Build()
	s.Bitfield().Key("1").OverflowFail().Incrby("1", 1, 1).Build()
	s.Bitfield().Key("1").Set("1", 1, 1).Incrby("1", 1, 1).Build()
	s.Bitfield().Key("1").Set("1", 1, 1).Build()
	s.Bitfield().Key("1").Incrby("1", 1, 1).Build()
	s.Bitfield().Key("1").Build()
	s.BitfieldRo().Key("1").Get().Get("1", 1).Get("1", 1).Build()
	s.BitfieldRo().Key("1").Get().Get("1", 1).Get("1", 1).Cache()
	s.BitfieldRo().Key("1").Build()
	s.BitfieldRo().Key("1").Cache()
	s.Bitop().And().Destkey("1").Key("1").Key("1").Build()
	s.Bitop().Or().Destkey("1").Key("1").Key("1").Build()
	s.Bitop().Xor().Destkey("1").Key("1").Key("1").Build()
	s.Bitop().Not().Destkey("1").Key("1").Key("1").Build()
	s.Bitpos().Key("1").Bit(1).Start(1).End(1).Byte().Build()
	s.Bitpos().Key("1").Bit(1).Start(1).End(1).Byte().Cache()
	s.Bitpos().Key("1").Bit(1).Start(1).End(1).Bit().Build()
	s.Bitpos().Key("1").Bit(1).Start(1).End(1).Bit().Cache()
	s.Bitpos().Key("1").Bit(1).Start(1).End(1).Build()
	s.Bitpos().Key("1").Bit(1).Start(1).End(1).Cache()
	s.Bitpos().Key("1").Bit(1).Start(1).Build()
	s.Bitpos().Key("1").Bit(1).Start(1).Cache()
	s.Bitpos().Key("1").Bit(1).Build()
	s.Bitpos().Key("1").Bit(1).Cache()
	s.Getbit().Key("1").Offset(1).Build()
	s.Getbit().Key("1").Offset(1).Cache()
	s.Setbit().Key("1").Offset(1).Value(1).Build()
}

func TestCommand_InitSlot_bitmap(t *testing.T) {
	var s = NewBuilder(InitSlot)
	bitmap0(s)
}

func TestCommand_NoSlot_bitmap(t *testing.T) {
	var s = NewBuilder(NoSlot)
	bitmap0(s)
}
