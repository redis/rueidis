package om

import (
	"context"
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/oklog/ulid/v2"
)

type HashTestStruct struct {
	Key string `redis:",key"`
	Ver int64  `redis:",ver"`
	Val []byte
	F1  bool
	F2  *bool
	F3  *string
	F4  *int64
}

func TestNewHashRepository(t *testing.T) {
	ctx := context.Background()

	client := setup(t)
	client.Do(ctx, client.B().Flushall().Build())
	defer client.Close()

	repo := NewHashRepository("hash", HashTestStruct{}, client)

	t.Run("NewEntity", func(t *testing.T) {
		e := repo.NewEntity().(*HashTestStruct)
		ulid.MustParse(e.Key)
	})

	t.Run("Save", func(t *testing.T) {
		f4 := rand.Int63()
		e := repo.NewEntity().(*HashTestStruct)

		// test save
		e.Val = []byte("any")
		e.F1 = true
		e.F2 = &e.F1
		e.F3 = &e.Key
		e.F4 = &f4
		if err := repo.Save(ctx, e); err != nil {
			t.Fatal(err)
		}
		if e.Ver != 1 {
			t.Fatalf("ver should be increment")
		}

		// test ErrVersionMismatch
		e.Ver = 0
		if err := repo.Save(ctx, e); err != ErrVersionMismatch {
			t.Fatalf("save should fail if ErrVersionMismatch, got: %v", err)
		}
		e.Ver = 1 // restore

		t.Run("Fetch", func(t *testing.T) {
			ei, err := repo.Fetch(ctx, e.Key)
			if err != nil {
				t.Fatal(err)
			}
			ee := ei.(*HashTestStruct)
			if e == ee {
				t.Fatalf("e's address should not be the same as ee's")
			}
			if !reflect.DeepEqual(e, ee) {
				t.Fatalf("e should be the same as ee")
			}
		})

		t.Run("FetchCache", func(t *testing.T) {
			ei, err := repo.FetchCache(ctx, e.Key, time.Minute)
			if err != nil {
				t.Fatal(err)
			}
			ee := ei.(*HashTestStruct)
			if e == ee {
				t.Fatalf("e's address should not be the same as ee's")
			}
			if !reflect.DeepEqual(e, ee) {
				t.Fatalf("ee should be the same as e")
			}
		})

		t.Run("Search", func(t *testing.T) {
			err := repo.CreateIndex(ctx, func(schema FtCreateSchema) Completed {
				return schema.FieldName("Val").Text().Build()
			})
			time.Sleep(time.Second)
			if err != nil {
				t.Fatal(err)
			}
			n, records, err := repo.Search(ctx, func(search FtSearchIndex) Completed {
				return search.Query("*").Build()
			})
			if err != nil {
				t.Fatal(err)
			}
			if n != 1 {
				t.Fatalf("unexpected total count %v", n)
			}
			items := records.([]*HashTestStruct)
			if len(items) != 1 {
				t.Fatalf("unexpected return count %v", n)
			}
			if !reflect.DeepEqual(e, items[0]) {
				t.Fatalf("items[0] should be the same as e")
			}
			if err = repo.DropIndex(ctx); err != nil {
				t.Fatal(err)
			}
		})

		t.Run("Delete", func(t *testing.T) {
			if err := repo.Remove(ctx, e.Key); err != nil {
				t.Fatal(err)
			}
			ei, err := repo.Fetch(ctx, e.Key)
			if !IsRecordNotFound(err) {
				t.Fatalf("should not be found, but got %v", ei)
			}
			ei, err = repo.FetchCache(ctx, e.Key, time.Minute)
			if !IsRecordNotFound(err) {
				t.Fatalf("should not be found, but got %v", e)
			}
		})
	})
}
