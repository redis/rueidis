package om

import (
	"context"
	"errors"
	"time"

	"github.com/redis/rueidis"
	"github.com/redis/rueidis/internal/cmds"
)

type (
	// FtCreateSchema is the FT.CREATE command builder
	FtCreateSchema = cmds.FtCreateSchema
	// FtSearchIndex is the FT.SEARCH command builder
	FtSearchIndex = cmds.FtSearchIndex
	// FtAggregateIndex is the FT.AGGREGATE command builder
	FtAggregateIndex = cmds.FtAggregateIndex
	// FtAlterSchema is the FT.ALTERINDEX command builder
	FtAlterIndex = cmds.FtAlterIndex
	// Arbitrary is an alias to cmds.Arbitrary. This allows the user to build an arbitrary command in Repository.CreateIndex
	Arbitrary = cmds.Arbitrary
)

var (
	// ErrVersionMismatch indicates that the optimistic update failed. That is someone else had already changed the entity.
	ErrVersionMismatch = errors.New("object version mismatched, please retry")
	// ErrEmptyHashRecord indicates the requested hash entity is not found.
	ErrEmptyHashRecord = errors.New("hash object not found")
)

// IsRecordNotFound checks if the error is indicating the requested entity is not found.
func IsRecordNotFound(err error) bool {
	return rueidis.IsRedisNil(err) || err == ErrEmptyHashRecord
}

// Repository is backed by HashRepository or JSONRepository
type Repository[T any] interface {
	NewEntity() (entity *T)
	Fetch(ctx context.Context, id string) (*T, error)
	FetchCache(ctx context.Context, id string, ttl time.Duration) (v *T, err error)
	Search(ctx context.Context, cmdFn func(search FtSearchIndex) rueidis.Completed) (int64, []*T, error)
	Aggregate(ctx context.Context, cmdFn func(agg FtAggregateIndex) rueidis.Completed) (*AggregateCursor, error)
	Save(ctx context.Context, entity *T) (err error)
	SaveMulti(ctx context.Context, entity ...*T) (errs []error)
	Remove(ctx context.Context, id string) error
	CreateIndex(ctx context.Context, cmdFn func(schema FtCreateSchema) rueidis.Completed) error
	CreateAndAliasIndex(ctx context.Context, cmdFn func(schema FtCreateSchema) rueidis.Completed) error
	AlterIndex(ctx context.Context, cmdFn func(alter FtAlterIndex) rueidis.Completed) error
	DropIndex(ctx context.Context) error
	IndexName() string
}

type RepositoryOption func(Repository[any])

func WithIndexName(name string) RepositoryOption {
	return func(r Repository[any]) {
		switch repo := r.(type) {
		case *HashRepository[any]:
			repo.idx = name
		case *JSONRepository[any]:
			repo.idx = name
		}
	}
}
