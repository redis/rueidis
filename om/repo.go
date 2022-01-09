package om

import (
	"context"
	"errors"
	"github.com/rueian/rueidis"
	"time"

	"github.com/rueian/rueidis/internal/cmds"
)

type (
	FtCreateSchema = cmds.FtCreateSchema
	FtSearchIndex  = cmds.FtSearchIndex
	Completed      = cmds.Completed
)

var (
	ErrVersionMismatch = errors.New("object version mismatched, please retry")
	ErrEmptyHashRecord = errors.New("hash object not found")
)

func IsRecordNotFound(err error) bool {
	return rueidis.IsRedisNil(err) || err == ErrEmptyHashRecord
}

type Repository interface {
	NewEntity() (entity interface{})
	Fetch(ctx context.Context, id string) (interface{}, error)
	FetchCache(ctx context.Context, id string, ttl time.Duration) (v interface{}, err error)
	Search(ctx context.Context, cmdFn func(search FtSearchIndex) Completed) (int64, interface{}, error)
	Save(ctx context.Context, entity interface{}) (err error)
	Remove(ctx context.Context, id string) error
	CreateIndex(ctx context.Context, cmdFn func(schema FtCreateSchema) Completed) error
	DropIndex(ctx context.Context) error
}
