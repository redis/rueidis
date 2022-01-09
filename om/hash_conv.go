package om

import (
	"fmt"
	"reflect"
	"strconv"
	"unsafe"
)

func newHashConvFactory(t reflect.Type, schema schema) *hashConvFactory {
	factory := &hashConvFactory{converters: make(map[string]conv, len(schema.fields))}
	for name, f := range schema.fields {
		var converter converter
		var ok bool

		switch f.typ.Kind() {
		case reflect.Ptr:
			converter, ok = converters.ptr[f.typ.Elem().Kind()]
		case reflect.Slice:
			converter, ok = converters.slice[f.typ.Elem().Kind()]
		default:
			converter, ok = converters.val[f.typ.Kind()]
		}
		if !ok {
			panic(fmt.Sprintf("schema %q should not contain unsupported field type %s.", t, f.typ.Kind()))
		}
		factory.converters[name] = conv{conv: converter, idx: f.idx}
	}
	return factory
}

type hashConvFactory struct {
	converters map[string]conv
}

type conv struct {
	idx  int
	conv converter
}

func (f hashConvFactory) NewConverter(entity reflect.Value) hashConv {
	return hashConv{factory: f, entity: entity}
}

type hashConv struct {
	factory hashConvFactory
	entity  reflect.Value
}

func (r hashConv) ToHash() (fields map[string]string) {
	fields = make(map[string]string, len(r.factory.converters))
	for f, converter := range r.factory.converters {
		ref := r.entity.Field(converter.idx)
		if v, ok := converter.conv.ValueToString(ref); ok {
			fields[f] = v
		}
	}
	return fields
}

func (r hashConv) FromHash(fields map[string]string) error {
	for f, field := range r.factory.converters {
		v, ok := fields[f]
		if !ok {
			continue
		}
		val, err := field.conv.StringToValue(v)
		if err != nil {
			return err
		}
		r.entity.Field(field.idx).Set(val)
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
				return *(*string)(unsafe.Pointer(&buf)), true
			},
			StringToValue: func(value string) (reflect.Value, error) {
				buf := []byte(value)
				return reflect.ValueOf(buf), nil
			},
		},
	},
}
