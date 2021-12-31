package om

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/rueian/rueidis"
	"github.com/rueian/rueidis/internal/proto"
)

var ErrVersionMismatch = errors.New("object version mismatched, please retry")

func NewHashRepository(prefix string, schema interface{}, client rueidis.Client) *HashRepository {
	repo := &HashRepository{
		prefix: prefix,
		typ:    reflect.TypeOf(schema),
		client: client,
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
	client  rueidis.Client
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
	record, err := r.client.Do(ctx, r.client.B().Hgetall().Key(r.key(id)).Build()).ToMap()
	if err != nil {
		return nil, err
	}
	return r.fromHash(id, record)
}

func (r *HashRepository) FetchCache(ctx context.Context, id string, ttl time.Duration) (v interface{}, err error) {
	record, err := r.client.DoCache(ctx, r.client.B().Hgetall().Key(r.key(id)).Cache(), ttl).ToMap()
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
		fields[VersionField], err = saveScript.Exec(ctx, r.client, []string{r.key(id)}, args).ToString()
		if proto.IsRedisNil(err) {
			return ErrVersionMismatch
		}
		if err != nil {
			return err
		}
		return conv.FromHash(id, fields)
	}
	cmd := r.client.B().Hset().Key(r.key(id)).FieldValue()
	for f, v := range fields {
		cmd = cmd.FieldValue(f, v)
	}
	return r.client.Do(ctx, cmd.Build()).Error()
}

func (r *HashRepository) Remove(ctx context.Context, id string) error {
	return r.client.Do(ctx, r.client.B().Del().Key(r.key(id)).Build()).Error()
}

var saveScript = rueidis.NewLuaScript(fmt.Sprintf(`
local v = redis.call('HGET',KEYS[1],'%s')
if (not v or v == ARGV[2])
then
  ARGV[2] = tostring(tonumber(ARGV[2])+1)
  if redis.call('HSET',KEYS[1],unpack(ARGV)) then return ARGV[2] end
end
return nil
`, VersionField))
