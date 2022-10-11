package om

import (
	"context"
	"encoding/json"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/rueian/rueidis"
)

// NewJSONRepository creates an JSONRepository.
// The prefix parameter is used as redis key prefix. The entity stored by the repository will be named in the form of `{prefix}:{id}`
// The schema parameter should be a struct with fields tagged with `redis:",key"` and `redis:",ver"`
func NewJSONRepository[T any](prefix string, schema T, client rueidis.Client) Repository[T] {
	repo := &JSONRepository[T]{
		prefix: prefix,
		idx:    "jsonidx:" + prefix,
		typ:    reflect.TypeOf(schema),
		client: client,
	}
	repo.schema = newSchema(repo.typ)
	return repo
}

var _ Repository[any] = (*JSONRepository[any])(nil)

// JSONRepository is an OM repository backed by RedisJSON.
type JSONRepository[T any] struct {
	schema schema
	typ    reflect.Type
	client rueidis.Client
	prefix string
	idx    string
}

// NewEntity returns an empty entity and will have the `redis:",key"` field be set with ULID automatically.
func (r *JSONRepository[T]) NewEntity() *T {
	var v T
	reflect.ValueOf(&v).Elem().Field(r.schema.key.idx).Set(reflect.ValueOf(id()))
	return &v
}

// Fetch an entity whose name is `{prefix}:{id}`
func (r *JSONRepository[T]) Fetch(ctx context.Context, id string) (v *T, err error) {
	record, err := r.client.Do(ctx, r.client.B().JsonGet().Key(key(r.prefix, id)).Build()).ToString()
	if err == nil {
		v, err = r.decode(record)
	}
	return v, err
}

// FetchCache is like Fetch, but it uses client side caching mechanism.
func (r *JSONRepository[T]) FetchCache(ctx context.Context, id string, ttl time.Duration) (v *T, err error) {
	record, err := r.client.DoCache(ctx, r.client.B().JsonGet().Key(key(r.prefix, id)).Cache(), ttl).ToString()
	if err == nil {
		v, err = r.decode(record)
	}
	return v, err
}

func (r *JSONRepository[T]) decode(record string) (*T, error) {
	var v T
	if err := json.NewDecoder(strings.NewReader(record)).Decode(&v); err != nil {
		return nil, err
	}
	return &v, nil
}

// Save the entity under the redis key of `{prefix}:{id}`.
// It also uses the `redis:",ver"` field and lua script to perform optimistic locking and prevent lost update.
func (r *JSONRepository[T]) Save(ctx context.Context, entity *T) (err error) {
	val := reflect.ValueOf(entity).Elem()

	keyField := val.Field(r.schema.key.idx)
	verField := val.Field(r.schema.ver.idx)

	str, err := jsonSaveScript.Exec(ctx, r.client, []string{key(r.prefix, keyField.String())}, []string{
		r.schema.ver.name, strconv.FormatInt(verField.Int(), 10), rueidis.JSON(entity),
	}).ToString()
	if rueidis.IsRedisNil(err) {
		return ErrVersionMismatch
	}
	if err == nil {
		ver, _ := strconv.ParseInt(str, 10, 64)
		verField.SetInt(ver)
	}
	return err
}

// Remove the entity under the redis key of `{prefix}:{id}`.
func (r *JSONRepository[T]) Remove(ctx context.Context, id string) error {
	return r.client.Do(ctx, r.client.B().Del().Key(key(r.prefix, id)).Build()).Error()
}

// CreateIndex uses FT.CREATE from the RediSearch module to create inverted index under the name `jsonidx:{prefix}`
// You can use the cmdFn parameter to mutate the index construction command,
// and note that the field name should be specified with JSON path syntax, otherwise the index may not work as expected.
func (r *JSONRepository[T]) CreateIndex(ctx context.Context, cmdFn func(schema FtCreateSchema) Completed) error {
	return r.client.Do(ctx, cmdFn(r.client.B().FtCreate().Index(r.idx).OnJson().Prefix(1).Prefix(r.prefix+":").Schema())).Error()
}

// DropIndex uses FT.DROPINDEX from the RediSearch module to drop index whose name is `jsonidx:{prefix}`
func (r *JSONRepository[T]) DropIndex(ctx context.Context) error {
	return r.client.Do(ctx, r.client.B().FtDropindex().Index(r.idx).Build()).Error()
}

// Search uses FT.SEARCH from the RediSearch module to search the index whose name is `jsonidx:{prefix}`
// It returns three values:
// 1. total count of match results inside the redis, and note that it might be larger than returned search result.
// 2. search result, and note that its length might smaller than the first return value.
// 3. error if any
// You can use the cmdFn parameter to mutate the search command.
func (r *JSONRepository[T]) Search(ctx context.Context, cmdFn func(search FtSearchIndex) Completed) (n int64, s []*T, err error) {
	resp, err := r.client.Do(ctx, cmdFn(r.client.B().FtSearch().Index(r.idx))).ToArray()
	if err == nil {
		n, _ = resp[0].ToInt64()
		s = make([]*T, 0, len(resp[1:])/2)
		for i := 2; i < len(resp); i += 2 {
			if kv, _ := resp[i].ToArray(); len(kv) >= 2 {
				for j := len(kv) - 2; j >= 0; i -= 2 {
					if k, _ := kv[j].ToString(); k == "$" {
						record, _ := kv[j+1].ToString()
						v, err := r.decode(record)
						if err != nil {
							return 0, nil, err
						}
						s = append(s, v)
						break
					}
				}
			}
		}
	}
	return n, s, err
}

// Aggregate performs the FT.AGGREGATE and returns a *AggregateCursor for accessing the results
func (r *JSONRepository[T]) Aggregate(ctx context.Context, cmdFn func(search FtAggregateIndex) Completed) (cursor *AggregateCursor, err error) {
	resp, err := r.client.Do(ctx, cmdFn(r.client.B().FtAggregate().Index(r.idx))).ToArray()
	if err != nil {
		return nil, err
	}
	return newAggregateCursor(r.idx, r.client, resp), nil
}

// IndexName returns the index name used in the FT.CREATE
func (r *JSONRepository[T]) IndexName() string {
	return r.idx
}

var jsonSaveScript = rueidis.NewLuaScript(`
local v = redis.call('JSON.GET',KEYS[1],ARGV[1])
if (not v or v == ARGV[2])
then
  redis.call('JSON.SET',KEYS[1],'$',ARGV[3])
  return redis.call('JSON.NUMINCRBY',KEYS[1],ARGV[1],1)
end
return nil
`)
