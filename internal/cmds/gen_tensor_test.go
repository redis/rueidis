// Code generated DO NOT EDIT

package cmds

import "testing"

func tensor0(s Builder) {
	s.AiTensorget().Key("1").Meta().Blob().Build()
	s.AiTensorget().Key("1").Meta().Blob().Cache()
	s.AiTensorget().Key("1").Meta().Values().Build()
	s.AiTensorget().Key("1").Meta().Values().Cache()
	s.AiTensorget().Key("1").Meta().Build()
	s.AiTensorget().Key("1").Meta().Cache()
	s.AiTensorset().Key("1").Float().Shape(1).Shape(1).Blob("1").Values("1").Values("1").Build()
	s.AiTensorset().Key("1").Float().Shape(1).Shape(1).Blob("1").Build()
	s.AiTensorset().Key("1").Float().Shape(1).Shape(1).Values("1").Values("1").Build()
	s.AiTensorset().Key("1").Float().Shape(1).Shape(1).Build()
	s.AiTensorset().Key("1").Double().Shape(1).Shape(1).Blob("1").Values("1").Values("1").Build()
	s.AiTensorset().Key("1").Double().Shape(1).Shape(1).Blob("1").Build()
	s.AiTensorset().Key("1").Double().Shape(1).Shape(1).Values("1").Values("1").Build()
	s.AiTensorset().Key("1").Double().Shape(1).Shape(1).Build()
	s.AiTensorset().Key("1").Int8().Shape(1).Shape(1).Blob("1").Values("1").Values("1").Build()
	s.AiTensorset().Key("1").Int8().Shape(1).Shape(1).Blob("1").Build()
	s.AiTensorset().Key("1").Int8().Shape(1).Shape(1).Values("1").Values("1").Build()
	s.AiTensorset().Key("1").Int8().Shape(1).Shape(1).Build()
	s.AiTensorset().Key("1").Int16().Shape(1).Shape(1).Blob("1").Values("1").Values("1").Build()
	s.AiTensorset().Key("1").Int16().Shape(1).Shape(1).Blob("1").Build()
	s.AiTensorset().Key("1").Int16().Shape(1).Shape(1).Values("1").Values("1").Build()
	s.AiTensorset().Key("1").Int16().Shape(1).Shape(1).Build()
	s.AiTensorset().Key("1").Int32().Shape(1).Shape(1).Blob("1").Values("1").Values("1").Build()
	s.AiTensorset().Key("1").Int32().Shape(1).Shape(1).Blob("1").Build()
	s.AiTensorset().Key("1").Int32().Shape(1).Shape(1).Values("1").Values("1").Build()
	s.AiTensorset().Key("1").Int32().Shape(1).Shape(1).Build()
	s.AiTensorset().Key("1").Int64().Shape(1).Shape(1).Blob("1").Values("1").Values("1").Build()
	s.AiTensorset().Key("1").Int64().Shape(1).Shape(1).Blob("1").Build()
	s.AiTensorset().Key("1").Int64().Shape(1).Shape(1).Values("1").Values("1").Build()
	s.AiTensorset().Key("1").Int64().Shape(1).Shape(1).Build()
	s.AiTensorset().Key("1").Uint8().Shape(1).Shape(1).Blob("1").Values("1").Values("1").Build()
	s.AiTensorset().Key("1").Uint8().Shape(1).Shape(1).Blob("1").Build()
	s.AiTensorset().Key("1").Uint8().Shape(1).Shape(1).Values("1").Values("1").Build()
	s.AiTensorset().Key("1").Uint8().Shape(1).Shape(1).Build()
	s.AiTensorset().Key("1").Uint16().Shape(1).Shape(1).Blob("1").Values("1").Values("1").Build()
	s.AiTensorset().Key("1").Uint16().Shape(1).Shape(1).Blob("1").Build()
	s.AiTensorset().Key("1").Uint16().Shape(1).Shape(1).Values("1").Values("1").Build()
	s.AiTensorset().Key("1").Uint16().Shape(1).Shape(1).Build()
	s.AiTensorset().Key("1").String().Shape(1).Shape(1).Blob("1").Values("1").Values("1").Build()
	s.AiTensorset().Key("1").String().Shape(1).Shape(1).Blob("1").Build()
	s.AiTensorset().Key("1").String().Shape(1).Shape(1).Values("1").Values("1").Build()
	s.AiTensorset().Key("1").String().Shape(1).Shape(1).Build()
	s.AiTensorset().Key("1").Bool().Shape(1).Shape(1).Blob("1").Values("1").Values("1").Build()
	s.AiTensorset().Key("1").Bool().Shape(1).Shape(1).Blob("1").Build()
	s.AiTensorset().Key("1").Bool().Shape(1).Shape(1).Values("1").Values("1").Build()
	s.AiTensorset().Key("1").Bool().Shape(1).Shape(1).Build()
}

func TestCommand_InitSlot_tensor(t *testing.T) {
	var s = NewBuilder(InitSlot)
	tensor0(s)
}

func TestCommand_NoSlot_tensor(t *testing.T) {
	var s = NewBuilder(NoSlot)
	tensor0(s)
}
