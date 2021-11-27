package om

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type toStr func(value reflect.Value) (string, bool)
type fromStr func(value string) (reflect.Value, error)

func int64PtrToStr(value reflect.Value) (string, bool) {
	return strconv.FormatInt(value.Elem().Int(), 10), true
}

func int64PtrFromStr(value string) (reflect.Value, error) {
	v, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(&v), nil
}

func strPtrToStr(value reflect.Value) (string, bool) {
	return value.Elem().String(), true
}

func strPtrFromStr(value string) (reflect.Value, error) {
	return reflect.ValueOf(&value), nil
}

func boolPtrToStr(value reflect.Value) (string, bool) {
	if value.Elem().Bool() {
		return "t", true
	}
	return "f", true
}

func boolPtrFromStr(value string) (reflect.Value, error) {
	b := value == "t"
	return reflect.ValueOf(&b), nil
}

func sliceToStr(sep string) toStr {
	return func(value reflect.Value) (string, bool) {
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
	}
}

func sliceFromStr(sep string) fromStr {
	return func(value string) (reflect.Value, error) {
		s := strings.Split(value, sep)
		return reflect.ValueOf(s), nil
	}
}

type HashConverter interface {
	ToHash() (id string, fields map[string]string)
	FromHash(id string, fields map[string]string) error
}

func newHashConvFactory(t reflect.Type) *hashConvFactory {
	if t.Kind() != reflect.Struct {
		panic(fmt.Sprintf("schema %q should be a struct", t))
	}

	v := reflect.New(t)

	fields := make(map[string]field, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		name, options, ok := parseStructTag(f.Tag)
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

		if name == PKField {
			if f.Type.Kind() != reflect.String {
				panic(fmt.Sprintf("field with tag `redis:\"_\"` in schema %q should be a string", t))
			}
			fields[name] = field{position: i}
		} else if name == VersionField {
			if f.Type.Kind() != reflect.Ptr || f.Type.Elem().Kind() != reflect.Int64 {
				panic(fmt.Sprintf("field with tag `redis:\"_v\"` in schema %q should be a *int64", t))
			}
			fields[name] = field{position: i, marshaler: int64PtrToStr, unmarshaler: int64PtrFromStr}
		} else {
			if f.Type.Kind() == reflect.Slice && f.Type.Elem().Kind() == reflect.String {
				sep, ok := options[SliceSepTag]
				if !ok {
					panic(fmt.Sprintf("string slice field should have separator in tag `redis:\"%s,sep=<xxx>\"` in schema %q", name, t))
				}
				fields[name] = field{position: i, marshaler: sliceToStr(sep), unmarshaler: sliceFromStr(sep)}
				continue
			} else if f.Type.Kind() == reflect.Ptr {
				switch f.Type.Elem().Kind() {
				case reflect.String:
					fields[name] = field{position: i, marshaler: strPtrToStr, unmarshaler: strPtrFromStr}
					continue
				case reflect.Bool:
					fields[name] = field{position: i, marshaler: boolPtrToStr, unmarshaler: boolPtrFromStr}
					continue
				case reflect.Int64:
					fields[name] = field{position: i, marshaler: int64PtrToStr, unmarshaler: int64PtrFromStr}
					continue
				}
			}
			panic(fmt.Sprintf("schema %q should not contain unsupported field type. Only *string, *int64, *bool, []string are supported", t))
		}
	}

	if _, ok := fields[PKField]; !ok {
		panic(fmt.Sprintf("schema %q should contain a string field with tag `redis:\"_\"` as primary key", t))
	}
	if _, ok := fields[VersionField]; !ok {
		panic(fmt.Sprintf("schema %q should contain a int64 field with tag `redis:\"_v\"` as version tag", t))
	}

	factory := &hashConvFactory{fields: fields}
	factory.pk = fields[PKField].position
	delete(fields, PKField)

	return factory
}

type hashConvFactory struct {
	pk     int
	fields map[string]field
}

type field struct {
	position    int
	marshaler   toStr
	unmarshaler fromStr
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
		if ref.IsNil() {
			continue
		}
		if v, ok := field.marshaler(ref); ok {
			fields[f] = v
		}
	}
	return r.entity.Field(r.factory.pk).String(), fields
}

func (r hashConv) FromHash(id string, fields map[string]string) error {
	r.entity.Field(r.factory.pk).Set(reflect.ValueOf(id))
	for f, v := range fields {
		field, ok := r.factory.fields[f]
		if !ok {
			continue
		}
		val, err := field.unmarshaler(v)
		if err != nil {
			return err
		}
		r.entity.Field(field.position).Set(val)
	}
	return nil
}
