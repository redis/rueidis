package om

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/rueian/rueidis"
)

type JSONTestStruct struct {
	Key    string `redis:",key"`
	Nested struct{ F1 string }
	Val    []byte
	Ver    int64 `redis:",ver"`
}

//gocyclo:ignore
func TestNewJSONRepository(t *testing.T) {
	ctx := context.Background()

	client := setup(t)
	client.Do(ctx, client.B().Flushall().Build())
	defer client.Close()

	repo := NewJSONRepository("json", JSONTestStruct{}, client)

	t.Run("IndexName", func(t *testing.T) {
		if name := repo.IndexName(); name != "jsonidx:json" {
			t.Fatal("unexpected value")
		}
	})

	t.Run("NewEntity", func(t *testing.T) {
		e := repo.NewEntity()
		ulid.MustParse(e.Key)
	})

	t.Run("Save", func(t *testing.T) {
		e := repo.NewEntity()

		// test save
		e.Val = []byte("any")
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
			if e == ei {
				t.Fatalf("e's address should not be the same as ee's")
			}
			if !reflect.DeepEqual(e, ei) {
				t.Fatalf("e should be the same as ee")
			}
		})

		t.Run("FetchCache", func(t *testing.T) {
			ei, err := repo.FetchCache(ctx, e.Key, time.Minute)
			if err != nil {
				t.Fatal(err)
			}
			ee := ei
			if e == ee {
				t.Fatalf("e's address should not be the same as ee's")
			}
			if !reflect.DeepEqual(e, ee) {
				t.Fatalf("ee should be the same as e")
			}
		})

		t.Run("Search", func(t *testing.T) {
			err := repo.CreateIndex(ctx, func(schema FtCreateSchema) rueidis.Completed {
				return schema.FieldName("$.Val").Text().Build()
			})
			if err != nil {
				t.Fatal(err)
			}
			time.Sleep(time.Second)
			n, records, err := repo.Search(ctx, func(search FtSearchIndex) rueidis.Completed {
				return search.Query("*").Build()
			})
			if err != nil {
				t.Fatal(err)
			}
			if n != 1 {
				t.Fatalf("unexpected total count %v", n)
			}
			if len(records) != 1 {
				t.Fatalf("unexpected return count %v", n)
			}
			if !reflect.DeepEqual(e, records[0]) {
				t.Fatalf("items[0] should be the same as e")
			}
			if err = repo.DropIndex(ctx); err != nil {
				t.Fatal(err)
			}
		})

		t.Run("Search Sort", func(t *testing.T) {
			err := repo.CreateIndex(ctx, func(schema FtCreateSchema) rueidis.Completed {
				return schema.FieldName("$.Val").Text().Sortable().Build()
			})
			if err != nil {
				t.Fatal(err)
			}
			time.Sleep(time.Second)
			n, records, err := repo.Search(ctx, func(search FtSearchIndex) rueidis.Completed {
				return search.Query("*").Sortby("$.Val").Build()
			})
			if err != nil {
				t.Fatal(err)
			}
			if n != 1 {
				t.Fatalf("unexpected total count %v", n)
			}
			if len(records) != 1 {
				t.Fatalf("unexpected return count %v", n)
			}
			if !reflect.DeepEqual(e, records[0]) {
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
			_, err = repo.FetchCache(ctx, e.Key, time.Minute)
			if !IsRecordNotFound(err) {
				t.Fatalf("should not be found, but got %v", e)
			}
		})
	})

	t.Run("SaveMulti", func(t *testing.T) {
		entities := []*JSONTestStruct{
			repo.NewEntity(),
			repo.NewEntity(),
			repo.NewEntity(),
		}

		for _, e := range entities {
			e.Val = []byte("any")
		}

		for i, err := range repo.SaveMulti(context.Background(), entities...) {
			if err != nil {
				t.Fatal(err)
			}
			if entities[i].Ver != 1 {
				t.Fatalf("unexpected ver %d", entities[i].Ver)
			}
		}

		entities[len(entities)-1].Ver = 0

		for i, err := range repo.SaveMulti(context.Background(), entities...) {
			if i == len(entities)-1 {
				if err != ErrVersionMismatch {
					t.Fatalf("unexpected err %v", err)
				}
			} else {
				if err != nil {
					t.Fatal(err)
				}
				if entities[i].Ver != 2 {
					t.Fatalf("unexpected ver %d", entities[i].Ver)
				}
			}
		}
	})
}
