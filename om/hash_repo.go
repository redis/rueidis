package om

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/rueian/rueidis"
	"github.com/rueian/rueidis/internal/cmds"
)

type FtCreateSchema = cmds.FtCreateSchema
type FtSearchIndex = cmds.FtSearchIndex
type Completed = cmds.Completed

var ErrVersionMismatch = errors.New("object version mismatched, please retry")

func NewHashRepository(prefix string, schema interface{}, client rueidis.Client) *HashRepository {
	repo := &HashRepository{
		prefix: prefix,
		idx:    "idx:" + prefix,
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
	idx     string
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

func (r *HashRepository) converter(v reflect.Value) (conv HashConverter) {
	if r.factory != nil {
		return r.factory.NewConverter(v)
	}
	return v.Interface().(HashConverter)
}

func (r *HashRepository) NewEntity() (entity interface{}) {
	v := reflect.New(r.typ)
	_ = r.converter(v).FromHash(id(), nil)
	return v.Interface()
}

func (r *HashRepository) fromHash(id string, record map[string]rueidis.RedisMessage) (v reflect.Value, err error) {
	fields := make(map[string]string, len(record))
	for k, v := range record {
		if s, err := v.ToString(); err == nil {
			fields[k] = s
		}
	}

	v = reflect.New(r.typ)
	if err := r.converter(v).FromHash(id, fields); err != nil {
		return reflect.Value{}, err
	}
	return v, nil
}

func (r *HashRepository) fromArray(id string, record []rueidis.RedisMessage) (v reflect.Value, err error) {
	fields := make(map[string]string, len(record)/2)
	for i := 0; i < len(record); i += 2 {
		k, _ := record[i].ToString()
		if s, err := record[i+1].ToString(); err == nil {
			fields[k] = s
		}
	}

	v = reflect.New(r.typ)
	if err := r.converter(v).FromHash(id, fields); err != nil {
		return reflect.Value{}, err
	}
	return v, nil
}

func (r *HashRepository) Fetch(ctx context.Context, id string) (v interface{}, err error) {
	record, err := r.client.Do(ctx, r.client.B().Hgetall().Key(r.key(id)).Build()).ToMap()
	if err != nil {
		return nil, err
	}
	val, err := r.fromHash(id, record)
	if err != nil {
		return nil, err
	}
	return val.Interface(), nil
}

func (r *HashRepository) FetchCache(ctx context.Context, id string, ttl time.Duration) (v interface{}, err error) {
	record, err := r.client.DoCache(ctx, r.client.B().Hgetall().Key(r.key(id)).Cache(), ttl).ToMap()
	if err != nil {
		return nil, err
	}
	val, err := r.fromHash(id, record)
	if err != nil {
		return nil, err
	}
	return val.Interface(), nil
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
		if rueidis.IsRedisNil(err) {
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

func (r *HashRepository) CreateIndex(ctx context.Context, cmdFn func(schema FtCreateSchema) Completed) error {
	return r.client.Do(ctx, cmdFn(r.client.B().FtCreate().Index(r.idx).OnHash().Prefix(1).Prefix(r.prefix+":").Schema())).Error()
}

func (r *HashRepository) DropIndex(ctx context.Context) error {
	return r.client.Do(ctx, r.client.B().FtDropindex().Index(r.idx).Build()).Error()
}

func (r *HashRepository) Search(ctx context.Context, cmdFn func(search FtSearchIndex) Completed) (int64, interface{}, error) {
	resp, err := r.client.Do(ctx, cmdFn(r.client.B().FtSearch().Index(r.idx))).ToArray()
	if err != nil {
		return 0, nil, err
	}
	prefix := r.prefix + ":"

	n, _ := resp[0].ToInt64()
	s := reflect.MakeSlice(reflect.SliceOf(reflect.PtrTo(r.typ)), 0, len(resp[1:])/2)
	for i := 1; i < len(resp); i += 2 {
		id, _ := resp[i].ToString()
		kv, _ := resp[i+1].ToArray()

		v, err := r.fromArray(strings.TrimPrefix(id, prefix), kv)
		if err != nil {
			return 0, nil, err
		}
		s = reflect.Append(s, v)
	}
	return n, s.Interface(), nil
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
