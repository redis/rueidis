// Code generated DO NOT EDIT

package cmds

import "testing"

func suggestion0(s Builder) {
	s.FtSugadd().Key("1").String("1").Score(1).Incr().Payload("1").Build()
	s.FtSugadd().Key("1").String("1").Score(1).Incr().Build()
	s.FtSugadd().Key("1").String("1").Score(1).Payload("1").Build()
	s.FtSugadd().Key("1").String("1").Score(1).Build()
	s.FtSugdel().Key("1").String("1").Build()
	s.FtSugget().Key("1").Prefix("1").Fuzzy().Withscores().Withpayloads().Max(1).Build()
	s.FtSugget().Key("1").Prefix("1").Fuzzy().Withscores().Withpayloads().Build()
	s.FtSugget().Key("1").Prefix("1").Fuzzy().Withscores().Max(1).Build()
	s.FtSugget().Key("1").Prefix("1").Fuzzy().Withscores().Build()
	s.FtSugget().Key("1").Prefix("1").Fuzzy().Withpayloads().Max(1).Build()
	s.FtSugget().Key("1").Prefix("1").Fuzzy().Withpayloads().Build()
	s.FtSugget().Key("1").Prefix("1").Fuzzy().Max(1).Build()
	s.FtSugget().Key("1").Prefix("1").Fuzzy().Build()
	s.FtSugget().Key("1").Prefix("1").Withscores().Withpayloads().Max(1).Build()
	s.FtSugget().Key("1").Prefix("1").Withscores().Withpayloads().Build()
	s.FtSugget().Key("1").Prefix("1").Withscores().Max(1).Build()
	s.FtSugget().Key("1").Prefix("1").Withscores().Build()
	s.FtSugget().Key("1").Prefix("1").Withpayloads().Max(1).Build()
	s.FtSugget().Key("1").Prefix("1").Withpayloads().Build()
	s.FtSugget().Key("1").Prefix("1").Max(1).Build()
	s.FtSugget().Key("1").Prefix("1").Build()
	s.FtSuglen().Key("1").Build()
}

func TestCommand_InitSlot_suggestion(t *testing.T) {
	var s = NewBuilder(InitSlot)
	suggestion0(s)
}

func TestCommand_NoSlot_suggestion(t *testing.T) {
	var s = NewBuilder(NoSlot)
	suggestion0(s)
}
