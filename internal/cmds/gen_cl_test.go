// Code generated by go generate; DO NOT EDIT

package cmds

import "testing"

func cl0(s Builder) {
	s.ClThrottle().Key("1").MaxBurst(1).CountPerPeriod(1).Period(1).Quantity(1).Build()
	s.ClThrottle().Key("1").MaxBurst(1).CountPerPeriod(1).Period(1).Build()
}

func TestCommand_InitSlot_cl(t *testing.T) {
	var s = NewBuilder(InitSlot)
	t.Run("0", func(t *testing.T) { cl0(s) })
}

func TestCommand_NoSlot_cl(t *testing.T) {
	var s = NewBuilder(NoSlot)
	t.Run("0", func(t *testing.T) { cl0(s) })
}
