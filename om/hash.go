package om

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/rueian/rueidis"
)

func NewHashRepository(prefix string, schema interface{}, client rueidis.Client) Repository {
	repo := &HashRepository{
		prefix: prefix,
		idx:    "hashidx:" + prefix,
		typ:    reflect.TypeOf(schema),
		client: client,
	}
	repo.schema = newSchema(repo.typ)
	repo.factory = newHashConvFactory(repo.typ, repo.schema)
	return repo
}

var _ Repository = (*HashRepository)(nil)

type HashRepository struct {
	prefix  string
	idx     string
	typ     reflect.Type
	schema  schema
	factory *hashConvFactory
	client  rueidis.Client
}

func (r *HashRepository) NewEntity() (entity interface{}) {
	v := reflect.New(r.typ)
	v.Elem().Field(r.schema.key.idx).Set(reflect.ValueOf(id()))
	return v.Interface()
}

func (r *HashRepository) Fetch(ctx context.Context, id string) (v interface{}, err error) {
	record, err := r.client.Do(ctx, r.client.B().Hgetall().Key(key(r.prefix, id)).Build()).ToMap()
	if err != nil {
		return nil, err
	}
	val, err := r.fromHash(record)
	if err != nil {
		return nil, err
	}
	return val.Interface(), nil
}

func (r *HashRepository) FetchCache(ctx context.Context, id string, ttl time.Duration) (v interface{}, err error) {
	record, err := r.client.DoCache(ctx, r.client.B().Hgetall().Key(key(r.prefix, id)).Cache(), ttl).ToMap()
	if err != nil {
		return nil, err
	}
	val, err := r.fromHash(record)
	if err != nil {
		return nil, err
	}
	return val.Interface(), nil
}

func (r *HashRepository) Save(ctx context.Context, entity interface{}) (err error) {
	val, ok := ptrValueOf(entity, r.typ)
	if !ok {
		panic(fmt.Sprintf("input entity should be a pointer to %v", r.typ))
	}

	fields := r.factory.NewConverter(val).ToHash()

	keyVal := fields[r.schema.key.name]
	verVal := fields[r.schema.ver.name]

	args := make([]string, 0, len(fields)*2)
	args = append(args, r.schema.ver.name, verVal) // keep the ver field be the first pair for the hashSaveScript
	delete(fields, r.schema.ver.name)
	for k, v := range fields {
		args = append(args, k, v)
	}

	str, err := hashSaveScript.Exec(ctx, r.client, []string{key(r.prefix, keyVal)}, args).ToString()
	if rueidis.IsRedisNil(err) {
		return ErrVersionMismatch
	}
	if err != nil {
		return err
	}
	ver, _ := strconv.ParseInt(str, 10, 64)
	val.Field(r.schema.ver.idx).SetInt(ver)
	return nil
}

func (r *HashRepository) Remove(ctx context.Context, id string) error {
	return r.client.Do(ctx, r.client.B().Del().Key(key(r.prefix, id)).Build()).Error()
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

	n, _ := resp[0].ToInt64()
	s := reflect.MakeSlice(reflect.SliceOf(reflect.PtrTo(r.typ)), 0, len(resp[1:])/2)
	for i := 2; i < len(resp); i += 2 {
		kv, _ := resp[i].ToArray()
		v, err := r.fromArray(kv)
		if err != nil {
			return 0, nil, err
		}
		s = reflect.Append(s, v)
	}
	return n, s.Interface(), nil
}

func (r *HashRepository) fromHash(record map[string]rueidis.RedisMessage) (v reflect.Value, err error) {
	if len(record) == 0 {
		return reflect.Value{}, ErrEmptyHashRecord
	}
	fields := make(map[string]string, len(record))
	for k, v := range record {
		if s, err := v.ToString(); err == nil {
			fields[k] = s
		}
	}

	v = reflect.New(r.typ)
	if err := r.factory.NewConverter(v.Elem()).FromHash(fields); err != nil {
		return reflect.Value{}, err
	}
	return v, nil
}

func (r *HashRepository) fromArray(record []rueidis.RedisMessage) (v reflect.Value, err error) {
	fields := make(map[string]string, len(record)/2)
	for i := 0; i < len(record); i += 2 {
		k, _ := record[i].ToString()
		if s, err := record[i+1].ToString(); err == nil {
			fields[k] = s
		}
	}
	v = reflect.New(r.typ)
	if err := r.factory.NewConverter(v.Elem()).FromHash(fields); err != nil {
		return reflect.Value{}, err
	}
	return v, nil
}

var hashSaveScript = rueidis.NewLuaScript(`
local v = redis.call('HGET',KEYS[1],ARGV[1])
if (not v or v == ARGV[2])
then
  ARGV[2] = tostring(tonumber(ARGV[2])+1)
  if redis.call('HSET',KEYS[1],unpack(ARGV)) then return ARGV[2] end
end
return nil
`)
