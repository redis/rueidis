package om

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/rueian/rueidis/internal/proto"
	"github.com/rueian/rueidis/pkg/script"
)

var ErrVersionMismatch = errors.New("object version mismatched, please retry")

type ObjectSaver func(key string, fields map[string]string) (err error)
type ObjectScripter func(script string) *script.Lua
type ObjectFetcher func(key string) (map[string]proto.Message, error)
type ObjectCacheFetcher func(key string, ttl time.Duration) (map[string]proto.Message, error)

func NewHashRepository(prefix string, schema interface{}, saver ObjectSaver, fetcher ObjectFetcher, cache ObjectCacheFetcher, scripter ObjectScripter) *HashRepository {
	repo := &HashRepository{
		prefix:  prefix,
		typ:     reflect.TypeOf(schema),
		saver:   saver,
		cache:   cache,
		fetcher: fetcher,
		script:  scripter(saveScript),
	}
	if _, ok := schema.(HashConverter); !ok {
		repo.factory = newHashConvFactory(repo.typ)
	}
	return repo
}

type HashRepository struct {
	prefix  string
	typ     reflect.Type
	factory *hashConvFactory

	saver   ObjectSaver
	cache   ObjectCacheFetcher
	fetcher ObjectFetcher
	script  *script.Lua
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

	v, conv := r.conv(reflect.New(r.typ))
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
	if ver, ok := fields[VersionField]; ok {
		args := make([]string, 0, len(fields)*2)
		args = append(args, VersionField, ver)
		for f, v := range fields {
			if f == VersionField {
				continue
			}
			args = append(args, f, v)
		}
		fields[VersionField], err = r.script.Exec([]string{r.key(id)}, args).ToString()
		if proto.IsRedisNil(err) {
			return ErrVersionMismatch
		}
		if err != nil {
			return err
		}
		return conv.FromHash(id, fields)
	}
	return r.saver(r.key(id), fields)
}

var saveScript = fmt.Sprintf(`
local v = redis.call('HGET',KEYS[1],'%s')
if (not v or v == ARGV[2])
then
  ARGV[2] = tostring(tonumber(ARGV[2])+1)
  if redis.call('HSET',KEYS[1],unpack(ARGV)) then return ARGV[2] end
end
return nil
`, VersionField)
