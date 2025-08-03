//go:build go1.23

package cmds

import (
	"maps"
	"testing"
)

func iter0(s Builder) {
	s.Hmset().Key("1").FieldValue().FieldValueIter(maps.All(map[string]string{"1": "1"})).Build()
	s.Hset().Key("1").FieldValue().FieldValueIter(maps.All(map[string]string{"1": "1"})).Build()
	s.Hsetex().Key("1").Fields().Numfields(1).FieldValue().FieldValueIter(maps.All(map[string]string{"1": "1"})).Build()
	s.Xadd().Key("1").Id("*").FieldValue().FieldValueIter(maps.All(map[string]string{"1": "1"})).Build()
	s.Zadd().Key("1").ScoreMember().ScoreMemberIter(maps.All(map[string]float64{"1": float64(1)})).Build()
}

func TestIter(t *testing.T) {
	var s = NewBuilder(InitSlot)
	t.Run("0", func(t *testing.T) { iter0(s) })
}
