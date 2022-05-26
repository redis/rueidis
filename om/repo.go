package om

import (
	"context"
	"errors"
	"time"

	"github.com/rueian/rueidis"
	"github.com/rueian/rueidis/internal/cmds"
)

type (
	// FtCreateSchema is the FT.CREATE command builder
	FtCreateSchema = cmds.FtCreateSchema
	// FtSearchIndex is the FT.SEARCH command builder
	FtSearchIndex = cmds.FtSearchIndex
	// Completed is the command builder result, should be created from the Build() of command builder
	Completed = cmds.Completed
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
	Search(ctx context.Context, cmdFn func(search FtSearchIndex) Completed) (int64, []*T, error)
	Save(ctx context.Context, entity *T) (err error)
	Remove(ctx context.Context, id string) error
	CreateIndex(ctx context.Context, cmdFn func(schema FtCreateSchema) Completed) error
	DropIndex(ctx context.Context) error
}
