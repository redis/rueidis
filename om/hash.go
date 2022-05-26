package om

import (
	"context"
	"reflect"
	"strconv"
	"time"

	"github.com/rueian/rueidis"
)

// NewHashRepository creates an HashRepository.
// The prefix parameter is used as redis key prefix. The entity stored by the repository will be named in the form of `{prefix}:{id}`
// The schema parameter should be a struct with fields tagged with `redis:",key"` and `redis:",ver"`
func NewHashRepository[T any](prefix string, schema T, client rueidis.Client) Repository[T] {
	repo := &HashRepository[T]{
		prefix: prefix,
		idx:    "hashidx:" + prefix,
		typ:    reflect.TypeOf(schema),
		client: client,
	}
	repo.schema = newSchema(repo.typ)
	repo.factory = newHashConvFactory(repo.typ, repo.schema)
	return repo
}

var _ Repository[any] = (*HashRepository[any])(nil)

// HashRepository is an OM repository backed by redis hash.
type HashRepository[T any] struct {
	schema  schema
	typ     reflect.Type
	client  rueidis.Client
	factory *hashConvFactory
	prefix  string
	idx     string
}

// NewEntity returns an empty entity and will have the `redis:",key"` field be set with ULID automatically.
func (r *HashRepository[T]) NewEntity() (entity *T) {
	var v T
	reflect.ValueOf(&v).Elem().Field(r.schema.key.idx).Set(reflect.ValueOf(id()))
	return &v
}

// Fetch an entity whose name is `{prefix}:{id}`
func (r *HashRepository[T]) Fetch(ctx context.Context, id string) (v *T, err error) {
	record, err := r.client.Do(ctx, r.client.B().Hgetall().Key(key(r.prefix, id)).Build()).ToMap()
	if err == nil {
		v, err = r.fromHash(record)
	}
	return v, err
}

// FetchCache is like Fetch, but it uses client side caching mechanism.
func (r *HashRepository[T]) FetchCache(ctx context.Context, id string, ttl time.Duration) (v *T, err error) {
	record, err := r.client.DoCache(ctx, r.client.B().Hgetall().Key(key(r.prefix, id)).Cache(), ttl).ToMap()
	if err == nil {
		v, err = r.fromHash(record)
	}
	return v, err
}

// Save the entity under the redis key of `{prefix}:{id}`.
// It also uses the `redis:",ver"` field and lua script to perform optimistic locking and prevent lost update.
func (r *HashRepository[T]) Save(ctx context.Context, entity *T) (err error) {
	val := reflect.ValueOf(entity).Elem()

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
	if err == nil {
		ver, _ := strconv.ParseInt(str, 10, 64)
		val.Field(r.schema.ver.idx).SetInt(ver)
	}
	return err
}

// Remove the entity under the redis key of `{prefix}:{id}`.
func (r *HashRepository[T]) Remove(ctx context.Context, id string) error {
	return r.client.Do(ctx, r.client.B().Del().Key(key(r.prefix, id)).Build()).Error()
}

// CreateIndex uses FT.CREATE from the RediSearch module to create inverted index under the name `hashidx:{prefix}`
// You can use the cmdFn parameter to mutate the index construction command.
func (r *HashRepository[T]) CreateIndex(ctx context.Context, cmdFn func(schema FtCreateSchema) Completed) error {
	return r.client.Do(ctx, cmdFn(r.client.B().FtCreate().Index(r.idx).OnHash().Prefix(1).Prefix(r.prefix+":").Schema())).Error()
}

// DropIndex uses FT.DROPINDEX from the RediSearch module to drop index whose name is `hashidx:{prefix}`
func (r *HashRepository[T]) DropIndex(ctx context.Context) error {
	return r.client.Do(ctx, r.client.B().FtDropindex().Index(r.idx).Build()).Error()
}

// Search uses FT.SEARCH from the RediSearch module to search the index whose name is `hashidx:{prefix}`
// It returns three values:
// 1. total count of match results inside the redis, and note that it might be larger than returned search result.
// 2. search result, and note that its length might smaller than the first return value.
// 3. error if any
// You can use the cmdFn parameter to mutate the search command.
func (r *HashRepository[T]) Search(ctx context.Context, cmdFn func(search FtSearchIndex) Completed) (n int64, s []*T, err error) {
	resp, err := r.client.Do(ctx, cmdFn(r.client.B().FtSearch().Index(r.idx))).ToArray()
	if err == nil {
		n, _ = resp[0].ToInt64()
		s = make([]*T, 0, len(resp[1:])/2)
		for i := 2; i < len(resp); i += 2 {
			kv, _ := resp[i].ToArray()
			v, err := r.fromArray(kv)
			if err != nil {
				return 0, nil, err
			}
			s = append(s, v)
		}
	}
	return n, s, err
}

func (r *HashRepository[T]) fromHash(record map[string]rueidis.RedisMessage) (*T, error) {
	if len(record) == 0 {
		return nil, ErrEmptyHashRecord
	}
	fields := make(map[string]string, len(record))
	for k, v := range record {
		if s, err := v.ToString(); err == nil {
			fields[k] = s
		}
	}
	return r.fromFields(fields)
}

func (r *HashRepository[T]) fromArray(record []rueidis.RedisMessage) (*T, error) {
	fields := make(map[string]string, len(record)/2)
	for i := 0; i < len(record); i += 2 {
		k, _ := record[i].ToString()
		if s, err := record[i+1].ToString(); err == nil {
			fields[k] = s
		}
	}
	return r.fromFields(fields)
}

func (r *HashRepository[T]) fromFields(fields map[string]string) (*T, error) {
	var v T
	if err := r.factory.NewConverter(reflect.ValueOf(&v).Elem()).FromHash(fields); err != nil {
		return nil, err
	}
	return &v, nil
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
