package cache

import (
	"time"

	"github.com/rueian/rueidis/internal/proto"
)

type Cache interface {
	GetOrPrepare(key, cmd string, ttl time.Duration) (v proto.Message, entry *Entry)
	Update(key, cmd string, value proto.Message, pttl int64)
	Delete(keys []proto.Message)
	FreeAndClose(notice proto.Message)
}
