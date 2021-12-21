package om

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type HashConverter interface {
	ToHash() (id string, fields map[string]string)
	FromHash(id string, fields map[string]string) error
}

const (
	PKOption     = "pk"
	IgnoreField  = "-"
	VersionField = "_v"
	SliceSepTag  = "sep"
)

func newHashConvFactory(t reflect.Type) *hashConvFactory {
	if t.Kind() != reflect.Struct {
		panic(fmt.Sprintf("schema %q should be a struct", t))
	}

	v := reflect.New(t)

	fields := make(map[string]field, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		name, options, ok := parseTag(f.Tag)
		if !ok {
			continue
		}
		if name == "" {
			panic(fmt.Sprintf("schema %q should not contain fields with empty redis tag", t))
		}
		if _, ok = fields[name]; ok {
			panic(fmt.Sprintf("schema %q should not contain fields with duplicated redis tag", t))
		}
		if !v.Elem().Field(i).CanSet() {
			panic(fmt.Sprintf("schema %q should not contain private fields with redis tag", t))
		}
		if name == IgnoreField {
			if _, ok := options[PKOption]; !ok {
				panic(fmt.Sprintf("schema %q should non pk fields with redis %q tag", t, "-"))
			}
		}
		if name == VersionField {
			if f.Type.Kind() != reflect.Int64 {
				panic(fmt.Sprintf("field with tag `redis:%q` in schema %q should be a int64", VersionField, t))
			}
		}
		if _, ok := options[PKOption]; ok {
			if f.Type.Kind() != reflect.String {
				panic(fmt.Sprintf("field with tag `redis:\",pk\"` in schema %q should be a string", t))
			}
		}

		var conv converter
		switch f.Type.Kind() {
		case reflect.Ptr:
			conv, ok = converters.ptr[f.Type.Elem().Kind()]
		case reflect.Slice:
			if builder := converters.slice[f.Type.Elem().Kind()]; builder != nil {
				sep := options[SliceSepTag]
				if len(sep) == 0 {
					panic(fmt.Sprintf("string slice field should have separator in tag `redis:\"%s,sep=<xxx>\"` in schema %q", name, t))
				}
				conv, ok = builder(sep), true
			}
		default:
			conv, ok = converters.val[f.Type.Kind()]
		}
		if !ok {
			panic(fmt.Sprintf("schema %q should not contain unsupported field type %s.", t, f.Type.Kind()))
		}
		fields[name] = field{position: i, options: options, converter: conv}
	}

	factory := &hashConvFactory{fields: fields, pk: -1}
	for _, f := range fields {
		if _, ok := f.options[PKOption]; ok {
			if factory.pk != -1 {
				panic(fmt.Sprintf("schema %q should contain only one field with tag `redis:\",pk\"`", t))
			}
			factory.pk = f.position
		}
	}
	if factory.pk == -1 {
		panic(fmt.Sprintf("schema %q should contain a string field with tag `redis:\",pk\"` as primary key", t))
	}
	if _, ok := fields[VersionField]; !ok {
		panic(fmt.Sprintf("schema %q should contain a int64 field with tag `redis:%q` as version tag", VersionField, t))
	}
	delete(fields, IgnoreField)

	return factory
}

type hashConvFactory struct {
	pk     int
	fields map[string]field
}

type field struct {
	position  int
	converter converter
	options   map[string]string
}

func (f hashConvFactory) NewConverter(entity reflect.Value) hashConv {
	if entity.Kind() == reflect.Ptr {
		entity = entity.Elem()
	}
	return hashConv{
		factory: f,
		entity:  entity,
	}
}

type hashConv struct {
	factory hashConvFactory
	entity  reflect.Value
}

func (r hashConv) ToHash() (id string, fields map[string]string) {
	fields = make(map[string]string, len(r.factory.fields))
	for f, field := range r.factory.fields {
		ref := r.entity.Field(field.position)
		if v, ok := field.converter.ValueToString(ref); ok {
			fields[f] = v
		}
	}
	return r.entity.Field(r.factory.pk).String(), fields
}

func (r hashConv) FromHash(id string, fields map[string]string) error {
	r.entity.Field(r.factory.pk).Set(reflect.ValueOf(id))
	for f, field := range r.factory.fields {
		v, ok := fields[f]
		if !ok {
			continue
		}
		val, err := field.converter.StringToValue(v)
		if err != nil {
			return err
		}
		r.entity.Field(field.position).Set(val)
	}
	return nil
}

func parseTag(tag reflect.StructTag) (name string, options map[string]string, ok bool) {
	if name, ok = tag.Lookup("redis"); !ok {
		return "", nil, false
	}
	tokens := strings.Split(name, ",")
	options = make(map[string]string, len(tokens)-1)
	for _, token := range tokens[1:] {
		kv := strings.SplitN(token, "=", 2)
		if len(kv) == 2 {
			options[kv[0]] = kv[1]
		} else {
			options[kv[0]] = ""
		}
	}
	return tokens[0], options, true
}

type converter struct {
	ValueToString func(value reflect.Value) (string, bool)
	StringToValue func(value string) (reflect.Value, error)
}

var converters = struct {
	val   map[reflect.Kind]converter
	ptr   map[reflect.Kind]converter
	slice map[reflect.Kind]func(sep string) converter
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
	slice: map[reflect.Kind]func(sep string) converter{
		reflect.String: func(sep string) converter {
			return converter{
				ValueToString: func(value reflect.Value) (string, bool) {
					length := value.Len()
					if length == 0 {
						return "", false
					}
					sb := strings.Builder{}
					for i := 0; i < length; i++ {
						sb.WriteString(value.Index(i).String())
						if i != length-1 {
							sb.WriteString(sep)
						}
					}
					return sb.String(), true
				},
				StringToValue: func(value string) (reflect.Value, error) {
					s := strings.Split(value, sep)
					return reflect.ValueOf(s), nil
				},
			}
		},
	},
}
