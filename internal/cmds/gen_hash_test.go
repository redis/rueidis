// Code generated DO NOT EDIT

package cmds

import "testing"

func hash0(s Builder) {
	s.Hdel().Key("1").Field("1").Field("1").Build()
	s.Hexists().Key("1").Field("1").Build()
	s.Hexists().Key("1").Field("1").Cache()
	s.Hget().Key("1").Field("1").Build()
	s.Hget().Key("1").Field("1").Cache()
	s.Hgetall().Key("1").Build()
	s.Hgetall().Key("1").Cache()
	s.Hincrby().Key("1").Field("1").Increment(1).Build()
	s.Hincrbyfloat().Key("1").Field("1").Increment(1).Build()
	s.Hkeys().Key("1").Build()
	s.Hkeys().Key("1").Cache()
	s.Hlen().Key("1").Build()
	s.Hlen().Key("1").Cache()
	s.Hmget().Key("1").Field("1").Field("1").Build()
	s.Hmget().Key("1").Field("1").Field("1").Cache()
	s.Hmset().Key("1").FieldValue().FieldValue("1", "1").FieldValue("1", "1").Build()
	s.Hrandfield().Key("1").Count(1).Withvalues().Build()
	s.Hrandfield().Key("1").Count(1).Build()
	s.Hrandfield().Key("1").Build()
	s.Hscan().Key("1").Cursor(1).Match("1").Count(1).Build()
	s.Hscan().Key("1").Cursor(1).Match("1").Build()
	s.Hscan().Key("1").Cursor(1).Count(1).Build()
	s.Hscan().Key("1").Cursor(1).Build()
	s.Hset().Key("1").FieldValue().FieldValue("1", "1").FieldValue("1", "1").Build()
	s.Hsetnx().Key("1").Field("1").Value("1").Build()
	s.Hstrlen().Key("1").Field("1").Build()
	s.Hstrlen().Key("1").Field("1").Cache()
	s.Hvals().Key("1").Build()
	s.Hvals().Key("1").Cache()
}

func TestCommand_InitSlot_hash(t *testing.T) {
	var s = NewBuilder(InitSlot)
	hash0(s)
}

func TestCommand_NoSlot_hash(t *testing.T) {
	var s = NewBuilder(NoSlot)
	hash0(s)
}
