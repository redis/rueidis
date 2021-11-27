package om

import (
	"context"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/rueian/rueidis/internal/proto"
)

const (
	PKField      = "_"
	VersionField = "_v"
	SliceSepTag  = "sep"
)

type ObjectSaver func(key string, fields map[string]string) (ver int64, err error)
type ObjectFetcher func(key string) (map[string]proto.Message, error)
type ObjectCacheFetcher func(key string, ttl time.Duration) (map[string]proto.Message, error)

func NewHashRepository(prefix string, schema interface{}, saver ObjectSaver, fetcher ObjectFetcher, cache ObjectCacheFetcher) *HashRepository {
	repo := &HashRepository{
		prefix:  prefix,
		typ:     reflect.TypeOf(schema),
		saver:   saver,
		cache:   cache,
		fetcher: fetcher,
	}
	if _, ok := schema.(HashConverter); !ok {
		repo.factory = newHashConvFactory(repo.typ)
	}
	return repo
}

func parseStructTag(tag reflect.StructTag) (name string, options map[string]string, ok bool) {
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

type HashRepository struct {
	prefix  string
	typ     reflect.Type
	factory *hashConvFactory

	saver   ObjectSaver
	cache   ObjectCacheFetcher
	fetcher ObjectFetcher
}

func (r *HashRepository) key(id string) (key string) {
	sb := strings.Builder{}
	sb.Grow(len(r.prefix) + len(id) + 1)
	sb.WriteString(r.prefix)
	sb.WriteString(":")
	sb.WriteString(id)
	return sb.String()
}

func (r *HashRepository) conv(v reflect.Value) (entity interface{}, conv HashConverter) {
	if r.factory != nil {
		return v.Interface(), r.factory.NewConverter(v)
	}
	entity = v.Interface()
	return entity, entity.(HashConverter)
}

func (r *HashRepository) Make() (entity interface{}) {
	v := reflect.New(r.typ)
	entity, conv := r.conv(v)
	_ = conv.FromHash(id(), nil)
	return entity
}

func (r *HashRepository) fromHash(id string, record map[string]proto.Message) (v interface{}, err error) {
	fields := make(map[string]string, len(record))
	for k, v := range record {
		if !v.IsNil() {
			fields[k] = v.String
		}
	}

	ret := reflect.New(r.typ)

	v, conv := r.conv(ret)
	if err := conv.FromHash(id, fields); err != nil {
		return nil, err
	}
	return v, nil
}

func (r *HashRepository) Fetch(ctx context.Context, id string) (v interface{}, err error) {
	record, err := r.fetcher(r.key(id))
	if err != nil {
		return nil, err
	}
	return r.fromHash(id, record)
}

func (r *HashRepository) FetchCache(ctx context.Context, id string, ttl time.Duration) (v interface{}, err error) {
	record, err := r.cache(r.key(id), ttl)
	if err != nil {
		return nil, err
	}
	return r.fromHash(id, record)
}

func (r *HashRepository) Save(ctx context.Context, entity interface{}) (err error) {
	var conv HashConverter

	if r.factory != nil {
		conv = r.factory.NewConverter(reflect.ValueOf(entity))
	} else {
		conv = entity.(HashConverter)
	}

	id, fields := conv.ToHash()
	if ver, err := r.saver(r.key(id), fields); err != nil {
		return err
	} else {
		fields[VersionField] = strconv.FormatInt(ver, 10)
	}
	return conv.FromHash(id, fields)
}
