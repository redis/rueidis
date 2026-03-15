package om

import (
	"context"
	"reflect"
	"strconv"
	"time"

	"github.com/oklog/ulid/v2"

	"github.com/redis/rueidis"
	"github.com/redis/rueidis/internal/cmds"
)

// NewHashRepository creates a HashRepository.
// The prefix parameter is used as redis key prefix. The entity stored by the repository will be named in the form of `{prefix}:{id}`
// The schema parameter should be a struct with fields tagged with `redis:",key"`. The `redis:",ver"` tag is optional for optimistic locking.
func NewHashRepository[T any](prefix string, schema T, client rueidis.Client, opts ...RepositoryOption) Repository[T] {
	repo := &HashRepository[T]{
		prefix: prefix,
		idx:    "hashidx:" + prefix,
		typ:    reflect.TypeOf(schema),
		client: client,
	}
	repo.schema = newSchema(repo.typ)
	repo.factory = newHashConvFactory(repo.typ, repo.schema)
	for _, opt := range opts {
		opt((*HashRepository[any])(repo))
	}
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
	reflect.ValueOf(&v).Elem().Field(r.schema.key.idx).Set(reflect.ValueOf(ulid.Make().String()))
	return &v
}

// Fetch an entity whose name is `{prefix}:{id}`
func (r *HashRepository[T]) Fetch(ctx context.Context, id string) (v *T, err error) {
	record, err := r.client.Do(ctx, r.client.B().Hgetall().Key(key(r.prefix, id)).Build()).AsStrMap()
	if err == nil {
		v, err = r.fromHash(record)
	}
	return v, err
}

// FetchCache is like Fetch, but it uses the client side caching mechanism.
func (r *HashRepository[T]) FetchCache(ctx context.Context, id string, ttl time.Duration) (v *T, err error) {
	record, err := r.client.DoCache(ctx, r.client.B().Hgetall().Key(key(r.prefix, id)).Cache(), ttl).AsStrMap()
	if err == nil {
		v, err = r.fromHash(record)
	}
	return v, err
}

func (r *HashRepository[T]) toExec(entity *T) (verf reflect.Value, exec rueidis.LuaExec) {
	val := reflect.ValueOf(entity).Elem()
	if !r.schema.verless {
		verf = val.Field(r.schema.ver.idx)
	} else {
		verf = reflect.ValueOf(int64(0)) // verless, set verf to a dummy value
	}
	fields := r.factory.NewConverter(val).ToHash()
	keyVal := fields[r.schema.key.name]
	var verVal string
	if !r.schema.verless {
		verVal = fields[r.schema.ver.name]
	}
	extVal := int64(0)
	if r.schema.ext != nil {
		if ext, ok := val.Field(r.schema.ext.idx).Interface().(time.Time); ok && !ext.IsZero() {
			extVal = ext.UnixMilli()
		}
	}
	exec.Keys = []string{key(r.prefix, keyVal)}
	if extVal != 0 {
		exec.Args = make([]string, 0, len(fields)*2+1)
	} else {
		exec.Args = make([]string, 0, len(fields)*2)
	}
	exec.Args = append(exec.Args, r.schema.ver.name, verVal) // keep the ver field be the first pair for the hashSaveScript
	delete(fields, r.schema.ver.name)
	for k, v := range fields {
		exec.Args = append(exec.Args, k, v)
	}
	if extVal != 0 {
		exec.Args = append(exec.Args, strconv.FormatInt(extVal, 10))
	}
	return
}

// Save the entity under the redis key of `{prefix}:{id}`.
// If the entity has a `redis:",ver"` field, it uses optimistic locking to prevent lost updates.
func (r *HashRepository[T]) Save(ctx context.Context, entity *T) (err error) {
	verf, exec := r.toExec(entity)
	str, err := hashSaveScript.Exec(ctx, r.client, exec.Keys, exec.Args).ToString()
	if rueidis.IsRedisNil(err) {
		return ErrVersionMismatch
	}
	if err == nil && !r.schema.verless {
		ver, _ := strconv.ParseInt(str, 10, 64)
		verf.SetInt(ver)
	}
	return err
}

// SaveMulti batches multiple HashRepository.Save at once
func (r *HashRepository[T]) SaveMulti(ctx context.Context, entities ...*T) []error {
	errs := make([]error, len(entities))
	verf := make([]reflect.Value, len(entities))
	exec := make([]rueidis.LuaExec, len(entities))
	for i, entity := range entities {
		verf[i], exec[i] = r.toExec(entity)
	}
	for i, resp := range hashSaveScript.ExecMulti(ctx, r.client, exec...) {
		str, err := resp.ToString()
		if rueidis.IsRedisNil(err) {
			errs[i] = ErrVersionMismatch
			continue
		}
		if err == nil && !r.schema.verless {
			ver, _ := strconv.ParseInt(str, 10, 64)
			verf[i].SetInt(ver)
		} else if err != nil {
			errs[i] = err
		}
	}
	return errs
}

// Remove the entity under the redis key of `{prefix}:{id}`.
func (r *HashRepository[T]) Remove(ctx context.Context, id string) error {
	return r.client.Do(ctx, r.client.B().Del().Key(key(r.prefix, id)).Build()).Error()
}

// AlterIndex uses FT.ALTER from the RediSearch module to alter index under the name `hashidx:{prefix}`
// You can use the cmdFn parameter to mutate the index alter command.
func (r *HashRepository[T]) AlterIndex(ctx context.Context, cmdFn func(alter FtAlterIndex) rueidis.Completed) error {
	return r.client.Do(ctx, cmdFn(r.client.B().FtAlter().Index(r.idx))).Error()
}

// CreateIndex uses FT.CREATE from the RediSearch module to create an inverted index under the name `hashidx:{prefix}`
// You can use the cmdFn parameter to mutate the index construction command.
func (r *HashRepository[T]) CreateIndex(ctx context.Context, cmdFn func(schema FtCreateSchema) rueidis.Completed) error {
	return r.client.Do(ctx, cmdFn(r.client.B().FtCreate().Index(r.idx).OnHash().Prefix(1).Prefix(r.prefix+":").Schema())).Error()
}

// CreateAndAliasIndex creates a new index, aliases it, and drops the old index if needed.
func (r *HashRepository[T]) CreateAndAliasIndex(ctx context.Context, cmdFn func(schema FtCreateSchema) rueidis.Completed) error {
	return createAndAliasIndex(ctx, r.idx, r.client, func(idx string) cmds.FtCreatePrefixPrefix {
		return r.client.B().FtCreate().
			Index(idx).
			OnHash().
			Prefix(1).
			Prefix(r.prefix + ":")
	}, cmdFn)
}

// DropIndex uses FT.DROPINDEX from the RediSearch module to drop the index whose name is `hashidx:{prefix}`
func (r *HashRepository[T]) DropIndex(ctx context.Context) error {
	return r.client.Do(ctx, r.client.B().FtDropindex().Index(r.idx).Build()).Error()
}

// Search uses FT.SEARCH from the RediSearch module to search the index whose name is `hashidx:{prefix}`
// It returns three values:
// 1. total count of match results inside the redis, and note that it might be larger than the returned search result.
// 2. the search result, and note that its length might be smaller than the first return value.
// 3. error if any
// You can use the cmdFn parameter to mutate the search command.
func (r *HashRepository[T]) Search(ctx context.Context, cmdFn func(search FtSearchIndex) rueidis.Completed) (n int64, s []*T, err error) {
	n, resp, err := r.client.Do(ctx, cmdFn(r.client.B().FtSearch().Index(r.idx))).AsFtSearch()
	if err == nil {
		s = make([]*T, len(resp))
		for i, v := range resp {
			if s[i], err = r.fromFields(v.Doc); err != nil {
				return 0, nil, err
			}
		}
	}
	return n, s, err
}

// Aggregate performs the FT.AGGREGATE and returns a *AggregateCursor for accessing the results
func (r *HashRepository[T]) Aggregate(ctx context.Context, cmdFn func(agg FtAggregateIndex) rueidis.Completed) (cursor *AggregateCursor, err error) {
	cid, total, resp, err := r.client.Do(ctx, cmdFn(r.client.B().FtAggregate().Index(r.idx))).AsFtAggregateCursor()
	if err != nil {
		return nil, err
	}
	return newAggregateCursor(r.idx, r.client, resp, cid, total), nil
}

// IndexName returns the index name used in the FT.CREATE
func (r *HashRepository[T]) IndexName() string {
	return r.idx
}

func (r *HashRepository[T]) fromHash(record map[string]string) (*T, error) {
	if len(record) == 0 {
		return nil, ErrEmptyHashRecord
	}
	return r.fromFields(record)
}

func (r *HashRepository[T]) fromFields(fields map[string]string) (*T, error) {
	var v T
	if err := r.factory.NewConverter(reflect.ValueOf(&v).Elem()).FromHash(fields); err != nil {
		return nil, err
	}
	return &v, nil
}

var hashSaveScript = rueidis.NewLuaScript(`
if (ARGV[1] == '')
then
  local e = (#ARGV % 2 == 1) and table.remove(ARGV) or nil
  if redis.call('HSET',KEYS[1],unpack(ARGV))
  then
    if e then redis.call('PEXPIREAT',KEYS[1],e) end
  end
  return ARGV[2]
end
local v = redis.call('HGET',KEYS[1],ARGV[1])
if (not v or v == ARGV[2])
then
  ARGV[2] = tostring(tonumber(ARGV[2])+1)
  local e = (#ARGV % 2 == 1) and table.remove(ARGV) or nil
  if redis.call('HSET',KEYS[1],unpack(ARGV))
  then
    if e then redis.call('PEXPIREAT',KEYS[1],e) end
    return ARGV[2]
  end
end
return nil
`)
