//go:build go1.18

package om

import (
	"context"
	"time"
)

type GenericRepository[T any] interface {
	NewEntity() (entity *T)
	Fetch(ctx context.Context, id string) (*T, error)
	FetchCache(ctx context.Context, id string, ttl time.Duration) (v *T, err error)
	Search(ctx context.Context, cmdFn func(search FtSearchIndex) Completed) (int64, []*T, error)
	Save(ctx context.Context, entity *T) (err error)
	Remove(ctx context.Context, id string) error
	CreateIndex(ctx context.Context, cmdFn func(schema FtCreateSchema) Completed) error
	DropIndex(ctx context.Context) error
}
