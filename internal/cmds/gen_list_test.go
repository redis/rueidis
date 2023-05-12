// Code generated DO NOT EDIT

package cmds

import "testing"

func list0(s Builder) {
	s.Blmove().Source("1").Destination("1").Left().Left().Timeout(1).Build()
	s.Blmove().Source("1").Destination("1").Left().Right().Timeout(1).Build()
	s.Blmove().Source("1").Destination("1").Right().Left().Timeout(1).Build()
	s.Blmove().Source("1").Destination("1").Right().Right().Timeout(1).Build()
	s.Blmpop().Timeout(1).Numkeys(1).Key("1").Key("1").Left().Count(1).Build()
	s.Blmpop().Timeout(1).Numkeys(1).Key("1").Key("1").Left().Build()
	s.Blmpop().Timeout(1).Numkeys(1).Key("1").Key("1").Right().Count(1).Build()
	s.Blmpop().Timeout(1).Numkeys(1).Key("1").Key("1").Right().Build()
	s.Blpop().Key("1").Key("1").Timeout(1).Build()
	s.Brpop().Key("1").Key("1").Timeout(1).Build()
	s.Brpoplpush().Source("1").Destination("1").Timeout(1).Build()
	s.Lindex().Key("1").Index(1).Build()
	s.Lindex().Key("1").Index(1).Cache()
	s.Linsert().Key("1").Before().Pivot("1").Element("1").Build()
	s.Linsert().Key("1").After().Pivot("1").Element("1").Build()
	s.Llen().Key("1").Build()
	s.Llen().Key("1").Cache()
	s.Lmove().Source("1").Destination("1").Left().Left().Build()
	s.Lmove().Source("1").Destination("1").Left().Right().Build()
	s.Lmove().Source("1").Destination("1").Right().Left().Build()
	s.Lmove().Source("1").Destination("1").Right().Right().Build()
	s.Lmpop().Numkeys(1).Key("1").Key("1").Left().Count(1).Build()
	s.Lmpop().Numkeys(1).Key("1").Key("1").Left().Build()
	s.Lmpop().Numkeys(1).Key("1").Key("1").Right().Count(1).Build()
	s.Lmpop().Numkeys(1).Key("1").Key("1").Right().Build()
	s.Lpop().Key("1").Count(1).Build()
	s.Lpop().Key("1").Build()
	s.Lpos().Key("1").Element("1").Rank(1).Count(1).Maxlen(1).Build()
	s.Lpos().Key("1").Element("1").Rank(1).Count(1).Maxlen(1).Cache()
	s.Lpos().Key("1").Element("1").Rank(1).Count(1).Build()
	s.Lpos().Key("1").Element("1").Rank(1).Count(1).Cache()
	s.Lpos().Key("1").Element("1").Rank(1).Maxlen(1).Build()
	s.Lpos().Key("1").Element("1").Rank(1).Maxlen(1).Cache()
	s.Lpos().Key("1").Element("1").Rank(1).Build()
	s.Lpos().Key("1").Element("1").Rank(1).Cache()
	s.Lpos().Key("1").Element("1").Count(1).Maxlen(1).Build()
	s.Lpos().Key("1").Element("1").Count(1).Maxlen(1).Cache()
	s.Lpos().Key("1").Element("1").Count(1).Build()
	s.Lpos().Key("1").Element("1").Count(1).Cache()
	s.Lpos().Key("1").Element("1").Maxlen(1).Build()
	s.Lpos().Key("1").Element("1").Maxlen(1).Cache()
	s.Lpos().Key("1").Element("1").Build()
	s.Lpos().Key("1").Element("1").Cache()
	s.Lpush().Key("1").Element("1").Element("1").Build()
	s.Lpushx().Key("1").Element("1").Element("1").Build()
	s.Lrange().Key("1").Start(1).Stop(1).Build()
	s.Lrange().Key("1").Start(1).Stop(1).Cache()
	s.Lrem().Key("1").Count(1).Element("1").Build()
	s.Lset().Key("1").Index(1).Element("1").Build()
	s.Ltrim().Key("1").Start(1).Stop(1).Build()
	s.Rpop().Key("1").Count(1).Build()
	s.Rpop().Key("1").Build()
	s.Rpoplpush().Source("1").Destination("1").Build()
	s.Rpush().Key("1").Element("1").Element("1").Build()
	s.Rpushx().Key("1").Element("1").Element("1").Build()
}

func TestCommand_InitSlot_list(t *testing.T) {
	var s = NewBuilder(InitSlot)
	list0(s)
}

func TestCommand_NoSlot_list(t *testing.T) {
	var s = NewBuilder(NoSlot)
	list0(s)
}
