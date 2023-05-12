// Code generated DO NOT EDIT

package cmds

import "testing"

func set0(s Builder) {
	s.Sadd().Key("1").Member("1").Member("1").Build()
	s.Scard().Key("1").Build()
	s.Scard().Key("1").Cache()
	s.Sdiff().Key("1").Key("1").Build()
	s.Sdiffstore().Destination("1").Key("1").Key("1").Build()
	s.Sinter().Key("1").Key("1").Build()
	s.Sintercard().Numkeys(1).Key("1").Key("1").Limit(1).Build()
	s.Sintercard().Numkeys(1).Key("1").Key("1").Build()
	s.Sinterstore().Destination("1").Key("1").Key("1").Build()
	s.Sismember().Key("1").Member("1").Build()
	s.Sismember().Key("1").Member("1").Cache()
	s.Smembers().Key("1").Build()
	s.Smembers().Key("1").Cache()
	s.Smismember().Key("1").Member("1").Member("1").Build()
	s.Smismember().Key("1").Member("1").Member("1").Cache()
	s.Smove().Source("1").Destination("1").Member("1").Build()
	s.Spop().Key("1").Count(1).Build()
	s.Spop().Key("1").Build()
	s.Srandmember().Key("1").Count(1).Build()
	s.Srandmember().Key("1").Build()
	s.Srem().Key("1").Member("1").Member("1").Build()
	s.Sscan().Key("1").Cursor(1).Match("1").Count(1).Build()
	s.Sscan().Key("1").Cursor(1).Match("1").Build()
	s.Sscan().Key("1").Cursor(1).Count(1).Build()
	s.Sscan().Key("1").Cursor(1).Build()
	s.Sunion().Key("1").Key("1").Build()
	s.Sunionstore().Destination("1").Key("1").Key("1").Build()
}

func TestCommand_InitSlot_set(t *testing.T) {
	var s = NewBuilder(InitSlot)
	set0(s)
}

func TestCommand_NoSlot_set(t *testing.T) {
	var s = NewBuilder(NoSlot)
	set0(s)
}
