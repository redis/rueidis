package om

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

const ignoreField = "-"

type schema struct {
	key    *field
	ver    *field
	ext    *field
	fields map[string]*field
}

type field struct {
	typ   reflect.Type
	name  string
	idx   int
	isKey bool
	isVer bool
	isExt bool
}

func newSchema(t reflect.Type) schema {
	if t.Kind() != reflect.Struct {
		panic(fmt.Sprintf("schema %q should be a struct", t))
	}

	s := schema{fields: make(map[string]*field, t.NumField())}

	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		if !sf.IsExported() {
			continue
		}
		f := parse(sf)
		if f.name == ignoreField {
			continue
		}
		f.idx = i
		s.fields[f.name] = &f

		if f.isKey {
			if sf.Type.Kind() != reflect.String {
				panic(fmt.Sprintf("field with tag `redis:\",key\"` in schema %q should be a string", t))
			}
			s.key = &f
		}
		if f.isVer {
			if sf.Type.Kind() != reflect.Int64 {
				panic(fmt.Sprintf("field with tag `redis:\",ver\"` in schema %q should be a int64", t))
			}
			s.ver = &f
		}
		if f.isExt {
			if sf.Type != reflect.TypeOf(time.Time{}) {
				panic(fmt.Sprintf("field with tag `redis:\",exat\"` in schema %q should be a time.Time", t))
			}
			s.ext = &f
		}
	}

	if s.key == nil {
		panic(fmt.Sprintf("schema %q should have one field with `redis:\",key\"` tag", t))
	}
	if s.ver == nil {
		panic(fmt.Sprintf("schema %q should have one field with `redis:\",ver\"` tag", t))
	}

	return s
}

func parse(f reflect.StructField) (field field) {
	v, _ := f.Tag.Lookup("json")
	vs := strings.SplitN(v, ",", 1)
	if vs[0] == "" {
		field.name = f.Name
	} else {
		field.name = vs[0]
	}

	v, _ = f.Tag.Lookup("redis")
	field.isKey = strings.Contains(v, ",key")
	field.isVer = strings.Contains(v, ",ver")
	field.isExt = strings.Contains(v, ",exat")
	field.typ = f.Type
	return field
}

func key(prefix, id string) (key string) {
	sb := strings.Builder{}
	sb.Grow(len(prefix) + len(id) + 1)
	sb.WriteString(prefix)
	sb.WriteString(":")
	sb.WriteString(id)
	return sb.String()
}
