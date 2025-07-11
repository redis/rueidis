package om

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/redis/rueidis"
)

// NewJSONRepository creates a JSONRepository.
// The prefix parameter is used as redis key prefix. The entity stored by the repository will be named in the form of `{prefix}:{id}`
// The schema parameter should be a struct with fields tagged with `redis:",key"` and `redis:",ver"`
func NewJSONRepository[T any](prefix string, schema T, client rueidis.Client, opts ...RepositoryOption) Repository[T] {
	repo := &JSONRepository[T]{
		prefix: prefix,
		idx:    "jsonidx:" + prefix,
		typ:    reflect.TypeOf(schema),
		client: client,
	}
	repo.schema = newSchema(repo.typ)
	for _, opt := range opts {
		opt((*JSONRepository[any])(repo))
	}
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
	reflect.ValueOf(&v).Elem().Field(r.schema.key.idx).Set(reflect.ValueOf(ulid.Make().String()))
	return &v
}

// Fetch an entity whose name is `{prefix}:{id}`
func (r *JSONRepository[T]) Fetch(ctx context.Context, id string) (v *T, err error) {
	record, err := r.client.Do(ctx, r.client.B().JsonGet().Key(key(r.prefix, id)).Path(".").Build()).ToString()
	if err == nil {
		v, err = r.decode(record)
	}
	return v, err
}

// FetchCache is like Fetch, but it uses the client side caching mechanism.
func (r *JSONRepository[T]) FetchCache(ctx context.Context, id string, ttl time.Duration) (v *T, err error) {
	record, err := r.client.DoCache(ctx, r.client.B().JsonGet().Key(key(r.prefix, id)).Path(".").Cache(), ttl).ToString()
	if err == nil {
		v, err = r.decode(record)
	}
	return v, err
}

func (r *JSONRepository[T]) decode(record string) (*T, error) {
	var v T
	if err := json.Unmarshal([]byte(record), &v); err != nil {
		return nil, err
	}
	return &v, nil
}

func (r *JSONRepository[T]) toExec(entity *T) (verf reflect.Value, exec rueidis.LuaExec) {
	val := reflect.ValueOf(entity).Elem()
	verf = val.Field(r.schema.ver.idx)
	extVal := int64(0)
	if r.schema.ext != nil {
		if ext, ok := val.Field(r.schema.ext.idx).Interface().(time.Time); ok && !ext.IsZero() {
			extVal = ext.UnixMilli()
		}
	}
	exec.Keys = []string{key(r.prefix, val.Field(r.schema.key.idx).String())}
	if extVal != 0 {
		exec.Args = []string{r.schema.ver.name, strconv.FormatInt(verf.Int(), 10), rueidis.JSON(entity), strconv.FormatInt(extVal, 10)}
	} else {
		exec.Args = []string{r.schema.ver.name, strconv.FormatInt(verf.Int(), 10), rueidis.JSON(entity)}
	}
	return
}

// Save the entity under the redis key of `{prefix}:{id}`.
// It also uses the `redis:",ver"` field and lua script to perform optimistic locking and prevent lost update.
func (r *JSONRepository[T]) Save(ctx context.Context, entity *T) (err error) {
	valf, exec := r.toExec(entity)
	str, err := jsonSaveScript.Exec(ctx, r.client, exec.Keys, exec.Args).ToString()
	if rueidis.IsRedisNil(err) {
		return ErrVersionMismatch
	}
	if err == nil {
		ver, _ := strconv.ParseInt(str, 10, 64)
		valf.SetInt(ver)
	}
	return err
}

// SaveMulti batches multiple HashRepository.Save at once
func (r *JSONRepository[T]) SaveMulti(ctx context.Context, entities ...*T) []error {
	errs := make([]error, len(entities))
	valf := make([]reflect.Value, len(entities))
	exec := make([]rueidis.LuaExec, len(entities))
	for i, entity := range entities {
		valf[i], exec[i] = r.toExec(entity)
	}
	for i, resp := range jsonSaveScript.ExecMulti(ctx, r.client, exec...) {
		if str, err := resp.ToString(); err != nil {
			if errs[i] = err; rueidis.IsRedisNil(err) {
				errs[i] = ErrVersionMismatch
			}
		} else {
			ver, _ := strconv.ParseInt(str, 10, 64)
			valf[i].SetInt(ver)
		}
	}
	return errs
}

// Remove the entity under the redis key of `{prefix}:{id}`.
func (r *JSONRepository[T]) Remove(ctx context.Context, id string) error {
	return r.client.Do(ctx, r.client.B().Del().Key(key(r.prefix, id)).Build()).Error()
}

// AlterIndex uses FT.ALTER from the RediSearch module to alter index under the name `jsonidx:{prefix}`
// You can use the cmdFn parameter to mutate the index alter command.
func (r *JSONRepository[T]) AlterIndex(ctx context.Context, cmdFn func(alter FtAlterIndex) rueidis.Completed) error {
	return r.client.Do(ctx, cmdFn(r.client.B().FtAlter().Index(r.idx))).Error()
}

// CreateIndex uses FT.CREATE from the RediSearch module to create an inverted index under the name `jsonidx:{prefix}`
// You can use the cmdFn parameter to mutate the index construction command,
// and note that the field name should be specified with JSON path syntax; otherwise, the index may not work as expected.
func (r *JSONRepository[T]) CreateIndex(ctx context.Context, cmdFn func(schema FtCreateSchema) rueidis.Completed) error {
	return r.client.Do(ctx, cmdFn(r.client.B().FtCreate().Index(r.idx).OnJson().Prefix(1).Prefix(r.prefix+":").Schema())).Error()
}

// CreateAndAliasIndex creates a new index, aliases it, and drops the old index if needed.
func (r *JSONRepository[T]) CreateAndAliasIndex(ctx context.Context, cmdFn func(schema FtCreateSchema) rueidis.Completed) error {
	alias := r.idx

	var currentIndex string
	aliasExists := false
	infoCmd := r.client.B().FtInfo().Index(alias).Build()
	infoResp, err := r.client.Do(ctx, infoCmd).ToMap()
	if err != nil {
		if strings.Contains(err.Error(), "Unknown index name") {
			aliasExists = false
		} else {
			return fmt.Errorf("failed to check if index exists: %w", err)
		}
	} else {
		aliasExists = true
	}

	if aliasExists {
		message, ok := infoResp["index_name"]
		if !ok {
			return fmt.Errorf("index_name not found in FT.INFO response")
		}
		currentIndex, err = message.ToString()
		if err != nil {
			return fmt.Errorf("failed to convert index_name to string: %w", err)
		}
	}

	// Compute new index version name
	newIndex := alias + "_v1"
	if aliasExists && currentIndex != "" {
		lastVersionIndex := strings.LastIndex(currentIndex, "_v")
		if lastVersionIndex != -1 && lastVersionIndex+2 < len(currentIndex) {
			versionStr := currentIndex[lastVersionIndex+2:]
			if version, err := strconv.Atoi(versionStr); err == nil {
				newIndex = fmt.Sprintf("%s_v%d", alias, version+1)
			}
		}
	}

	// Create the new index with schema
	createCmd := r.client.B().FtCreate().
		Index(newIndex).
		OnJson().
		Prefix(1).
		Prefix(r.prefix + ":")
	if err := r.client.Do(ctx, cmdFn(createCmd.Schema())).Error(); err != nil {
		return fmt.Errorf("failed to create index %s: %w", newIndex, err)
	}

	// Set alias to point to new index
	var aliasErr error
	if aliasExists {
		aliasErr = r.client.Do(ctx, r.client.B().FtAliasupdate().Alias(alias).Index(newIndex).Build()).Error()
	} else {
		aliasErr = r.client.Do(ctx, r.client.B().FtAliasadd().Alias(alias).Index(newIndex).Build()).Error()
	}

	if aliasErr != nil {
		return fmt.Errorf("failed to update alias: %w", aliasErr)
	}

	// Drop old index if it's different from the new one
	if aliasExists && currentIndex != "" && currentIndex != newIndex {
		if err := r.client.Do(ctx, r.client.B().FtDropindex().Index(currentIndex).Build()).Error(); err != nil {
			return fmt.Errorf("failed to drop old index: %w", err)
		}
	}

	return nil
}


// DropIndex uses FT.DROPINDEX from the RediSearch module to drop the index whose name is `jsonidx:{prefix}`
func (r *JSONRepository[T]) DropIndex(ctx context.Context) error {
	return r.client.Do(ctx, r.client.B().FtDropindex().Index(r.idx).Build()).Error()
}

// Search uses FT.SEARCH from the RediSearch module to search the index whose name is `jsonidx:{prefix}`
// It returns three values:
// 1. total count of match results inside the redis, and note that it might be larger than the returned search result.
// 2. the search result, and note that its length might be smaller than the first return value.
// 3. error if any
// You can use the cmdFn parameter to mutate the search command.
func (r *JSONRepository[T]) Search(ctx context.Context, cmdFn func(search FtSearchIndex) rueidis.Completed) (n int64, s []*T, err error) {
	n, resp, err := r.client.Do(ctx, cmdFn(r.client.B().FtSearch().Index(r.idx))).AsFtSearch()
	if err == nil {
		s = make([]*T, len(resp))
		for i, v := range resp {
			doc := v.Doc["$"]
			doc = strings.TrimPrefix(doc, "[") // supports dialect 3
			doc = strings.TrimSuffix(doc, "]")
			if s[i], err = r.decode(doc); err != nil {
				return 0, nil, err
			}
		}
	}
	return n, s, err
}

// Aggregate performs the FT.AGGREGATE and returns a *AggregateCursor for accessing the results
func (r *JSONRepository[T]) Aggregate(ctx context.Context, cmdFn func(agg FtAggregateIndex) rueidis.Completed) (cursor *AggregateCursor, err error) {
	cid, total, resp, err := r.client.Do(ctx, cmdFn(r.client.B().FtAggregate().Index(r.idx))).AsFtAggregateCursor()
	if err != nil {
		return nil, err
	}
	return newAggregateCursor(r.idx, r.client, resp, cid, total), nil
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
  local v = redis.call('JSON.NUMINCRBY',KEYS[1],ARGV[1],1)
  if #ARGV == 4 then redis.call('PEXPIREAT',KEYS[1],ARGV[4]) end
  return v
end
return nil
`)
