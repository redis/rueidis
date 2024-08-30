//go:build go1.23

package cmds

import (
	"maps"
)

func iter0(s Builder) {
	s.Hmset().Key("1").FieldValue().FieldValues(maps.All(map[string]string{"1": "1"})).Build()
	s.Hset().Key("1").FieldValue().FieldValues(maps.All(map[string]string{"1": "1"})).Build()
	s.Xadd().Key("1").Id("*").FieldValue().FieldValues(maps.All(map[string]string{"1": "1"})).Build()
}
