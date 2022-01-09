package om

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/rueian/rueidis/internal/cmds"
)

const ignoreField = "-"

type FtCreateSchema = cmds.FtCreateSchema
type FtSearchIndex = cmds.FtSearchIndex
type Completed = cmds.Completed

var ErrVersionMismatch = errors.New("object version mismatched, please retry")

type schema struct {
	keyField *field
	verField *field
	fields   map[string]*field
}

type field struct {
	name       string
	idx        int
	typ        reflect.Type
	isKeyField bool
	isVerField bool
}

func newSchema(t reflect.Type) schema {
	if t.Kind() != reflect.Struct {
		panic(fmt.Sprintf("schema %q should be a struct", t))
	}

	schema := schema{fields: make(map[string]*field, t.NumField())}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if !f.IsExported() {
			continue
		}
		field := parse(f)
		if field.name == ignoreField {
			continue
		}
		field.idx = i
		schema.fields[field.name] = &field

		if field.isKeyField {
			if f.Type.Kind() != reflect.String {
				panic(fmt.Sprintf("field with tag `redis:\",key\"` in schema %q should be a string", t))
			}
			schema.keyField = &field
		}
		if field.isVerField {
			if f.Type.Kind() != reflect.Int64 {
				panic(fmt.Sprintf("field with tag `redis:\",ver\"` in schema %q should be a int64", t))
			}
			schema.verField = &field
		}
	}

	if schema.keyField == nil {
		panic(fmt.Sprintf("schema %q should have one field with `redis:\",key\"` tag", t))
	}
	if schema.verField == nil {
		panic(fmt.Sprintf("schema %q should have one field with `redis:\",ver\"` tag", t))
	}

	return schema
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
	field.isKeyField = strings.Contains(v, ",key")
	field.isVerField = strings.Contains(v, ",ver")
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

func ptrValueOf(entity interface{}, typ reflect.Type) (reflect.Value, bool) {
	val := reflect.ValueOf(entity)
	if val.Kind() != reflect.Ptr {
		return reflect.Value{}, false
	}
	val = val.Elem()
	return val, val.Type() == typ
}
