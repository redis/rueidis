package om

import (
	"fmt"
	"reflect"
	"strconv"
	"unsafe"

	"github.com/redis/rueidis"
)

func newHashConvFactory(t reflect.Type, schema schema) *hashConvFactory {
	factory := &hashConvFactory{fields: make(map[string]fieldConv, len(schema.fields))}
	for name, f := range schema.fields {
		conv, ok := converters.val[f.typ.Kind()]
		switch f.typ.Kind() {
		case reflect.Ptr:
			conv, ok = converters.ptr[f.typ.Elem().Kind()]
		case reflect.Slice:
			conv, ok = converters.slice[f.typ.Elem().Kind()]
		}
		if !ok {
			k := f.typ.Kind()
			panic(fmt.Sprintf("schema %q should not contain unsupported field type %s.", t, k))
		}
		factory.fields[name] = fieldConv{conv: conv, idx: f.idx}
	}
	return factory
}

type hashConvFactory struct {
	fields map[string]fieldConv
}

type fieldConv struct {
	conv converter
	idx  int
}

func (f hashConvFactory) NewConverter(entity reflect.Value) hashConv {
	return hashConv{factory: f, entity: entity}
}

type hashConv struct {
	factory hashConvFactory
	entity  reflect.Value
}

func (r hashConv) ToHash() (fields map[string]string) {
	fields = make(map[string]string, len(r.factory.fields))
	for k, f := range r.factory.fields {
		ref := r.entity.Field(f.idx)
		if v, ok := f.conv.ValueToString(ref); ok {
			fields[k] = v
		}
	}
	return fields
}

func (r hashConv) FromHash(fields map[string]string) error {
	for k, f := range r.factory.fields {
		v, ok := fields[k]
		if !ok {
			continue
		}
		val, err := f.conv.StringToValue(v)
		if err != nil {
			return err
		}
		r.entity.Field(f.idx).Set(val)
	}
	return nil
}

type converter struct {
	ValueToString func(value reflect.Value) (string, bool)
	StringToValue func(value string) (reflect.Value, error)
}

var converters = struct {
	val   map[reflect.Kind]converter
	ptr   map[reflect.Kind]converter
	slice map[reflect.Kind]converter
}{
	ptr: map[reflect.Kind]converter{
		reflect.Int64: {
			ValueToString: func(value reflect.Value) (string, bool) {
				if value.IsNil() {
					return "", false
				}
				return strconv.FormatInt(value.Elem().Int(), 10), true
			},
			StringToValue: func(value string) (reflect.Value, error) {
				v, err := strconv.ParseInt(value, 10, 64)
				if err != nil {
					return reflect.Value{}, err
				}
				return reflect.ValueOf(&v), nil
			},
		},
		reflect.String: {
			ValueToString: func(value reflect.Value) (string, bool) {
				if value.IsNil() {
					return "", false
				}
				return value.Elem().String(), true
			},
			StringToValue: func(value string) (reflect.Value, error) {
				return reflect.ValueOf(&value), nil
			},
		},
		reflect.Bool: {
			ValueToString: func(value reflect.Value) (string, bool) {
				if value.IsNil() {
					return "", false
				}
				if value.Elem().Bool() {
					return "t", true
				}
				return "f", true
			},
			StringToValue: func(value string) (reflect.Value, error) {
				b := value == "t"
				return reflect.ValueOf(&b), nil
			},
		},
	},
	val: map[reflect.Kind]converter{
		reflect.Int64: {
			ValueToString: func(value reflect.Value) (string, bool) {
				return strconv.FormatInt(value.Int(), 10), true
			},
			StringToValue: func(value string) (reflect.Value, error) {
				v, err := strconv.ParseInt(value, 10, 64)
				if err != nil {
					return reflect.Value{}, err
				}
				return reflect.ValueOf(v), nil
			},
		},
		reflect.String: {
			ValueToString: func(value reflect.Value) (string, bool) {
				return value.String(), true
			},
			StringToValue: func(value string) (reflect.Value, error) {
				return reflect.ValueOf(value), nil
			},
		},
		reflect.Bool: {
			ValueToString: func(value reflect.Value) (string, bool) {
				if value.Bool() {
					return "t", true
				}
				return "f", true
			},
			StringToValue: func(value string) (reflect.Value, error) {
				b := value == "t"
				return reflect.ValueOf(b), nil
			},
		},
	},
	slice: map[reflect.Kind]converter{
		reflect.Uint8: {
			ValueToString: func(value reflect.Value) (string, bool) {
				buf, ok := value.Interface().([]byte)
				if !ok {
					return "", false
				}
				return rueidis.BinaryString(buf), true
			},
			StringToValue: func(value string) (reflect.Value, error) {
				buf := unsafe.Slice(unsafe.StringData(value), len(value))
				return reflect.ValueOf(buf), nil
			},
		},
		reflect.Float32: {
			ValueToString: func(value reflect.Value) (string, bool) {
				vs, ok := value.Interface().([]float32)
				return rueidis.VectorString32(vs), ok
			},
			StringToValue: func(value string) (reflect.Value, error) {
				return reflect.ValueOf(rueidis.ToVector32(value)), nil
			},
		},
		reflect.Float64: {
			ValueToString: func(value reflect.Value) (string, bool) {
				vs, ok := value.Interface().([]float64)
				return rueidis.VectorString64(vs), ok
			},
			StringToValue: func(value string) (reflect.Value, error) {
				return reflect.ValueOf(rueidis.ToVector64(value)), nil
			},
		},
	},
}
