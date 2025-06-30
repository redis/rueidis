package om

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/redis/rueidis"
)

type JSONTestStruct struct {
	Key    string `redis:",key"`
	Nested struct{ F1 string }
	Val    []byte
	Ver    int64 `redis:",ver"`
}

func TestNewJsonRepositoryMismatch(t *testing.T) {
	ctx := context.Background()

	client := setup(t)
	client.Do(ctx, client.B().Flushall().Build())
	defer client.Close()

	repo := NewJSONRepository("jsonmismatch", Mismatch{}, client)
	if err := repo.CreateIndex(ctx, func(schema FtCreateSchema) rueidis.Completed {
		return schema.FieldName("$.F1").Tag().Build()
	}); err != nil {
		t.Fatal(err)
	}

	t.Run("Mismatch", func(t *testing.T) {
		e := repo.NewEntity()
		if err := repo.Save(ctx, e); err != nil {
			t.Fatal(err)
		}
		if err := client.Do(ctx, client.B().Del().Key("jsonmismatch:"+e.Key).Build()).Error(); err != nil {
			t.Fatal(err)
		}
		if err := client.Do(ctx, client.B().JsonSet().Key("jsonmismatch:"+e.Key).Path("$").Value(rueidis.JSON("1")).Build()).Error(); err != nil {
			t.Fatal(err)
		}
		if _, err := repo.Fetch(ctx, e.Key); err == nil {
			t.Fatal("Fetch not failed as expected")
		}
		if _, _, err := repo.Search(ctx, func(search FtSearchIndex) rueidis.Completed {
			return search.Query("*").Build()
		}); err == nil {
			t.Fatal("Search not failed as expected")
		}
	})
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
			n, records, err = repo.Search(ctx, func(search FtSearchIndex) rueidis.Completed {
				return search.Query("*").Dialect(3).Build()
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

		t.Run("Alter Index", func(t *testing.T) {
			err := repo.CreateIndex(ctx, func(schema FtCreateSchema) rueidis.Completed {
				return schema.FieldName("$.Val").Text().Build()
			})
			if err != nil {
				t.Fatal(err)
			}
			time.Sleep(time.Second)
			var entities []*JSONTestStruct
			for i := 3; i >= 1; i-- {
				e := repo.NewEntity()
				e.Val = []byte("any")
				e.Nested = struct {
					F1 string
				}{
					F1: fmt.Sprintf("%d", i),
				}
				err = repo.Save(ctx, e)
				if err != nil {
					t.Fatal(err)
				}
				entities = append(entities, e)
			}
			time.Sleep(time.Second)
			n, records, err := repo.Search(ctx, func(search FtSearchIndex) rueidis.Completed {
				return search.Query("*").Sortby("$.Nested.F1").Build()
			})
			if err == nil {
				t.Fatalf("search by property neither loaded nor in schema")
			}
			err = repo.AlterIndex(ctx, func(alter FtAlterIndex) rueidis.Completed {
				return alter.
					Schema().Add().Field("$.Nested.F1").Options("TEXT", "SORTABLE").
					Build()
			})
			if err != nil {
				t.Fatal(err)
			}
			time.Sleep(time.Second)
			n, records, err = repo.Search(ctx, func(search FtSearchIndex) rueidis.Completed {
				return search.Query("*").Sortby("$.Nested.F1").Build()
			})
			if err != nil {
				t.Fatal(err)
			}
			if n != 3 {
				t.Fatalf("unexpected total count %v", n)
			}
			if len(records) != 3 {
				t.Fatalf("unexpected return count %v", n)
			}
			if !reflect.DeepEqual(entities[2], records[0]) {
				t.Fatalf("entities[0] should be the same as records[2]")
			}
			if err = repo.DropIndex(ctx); err != nil {
				t.Fatal(err)
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

type JSONTestTTLStruct struct {
	Key  string    `redis:",key"`
	Ver  int64     `redis:",ver"`
	Exat time.Time `redis:",exat"`
}

//gocyclo:ignore
func TestNewJSONTTLRepository(t *testing.T) {
	ctx := context.Background()

	client := setup(t)
	client.Do(ctx, client.B().Flushall().Build())
	defer client.Close()

	repo := NewJSONRepository("jsonttl", JSONTestTTLStruct{}, client)

	t.Run("NewEntity", func(t *testing.T) {
		e := repo.NewEntity()
		ulid.MustParse(e.Key)
	})

	t.Run("Save", func(t *testing.T) {
		e := repo.NewEntity()

		// test save
		e.Exat = time.Now().Add(time.Minute)
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

		t.Run("ExpireAt", func(t *testing.T) {
			exat, err := client.Do(ctx, client.B().Pexpiretime().Key("jsonttl:"+e.Key).Build()).AsInt64()
			if err != nil {
				t.Fatal(err)
			}
			if exat != e.Exat.UnixMilli() {
				t.Fatalf("wrong exat")
			}
		})

		t.Run("Fetch", func(t *testing.T) {
			ei, err := repo.Fetch(ctx, e.Key)
			if err != nil {
				t.Fatal(err)
			}
			if e == ei {
				t.Fatalf("e's address should not be the same as ee's")
			}
			e.Exat = e.Exat.Truncate(time.Nanosecond)
			ei.Exat = ei.Exat.Truncate(time.Nanosecond)
			if !e.Exat.Equal(ei.Exat) {
				t.Fatalf("e should be the same as ee %v, %v", e, ei)
			}
		})

		t.Run("FetchCache", func(t *testing.T) {
			ei, err := repo.FetchCache(ctx, e.Key, time.Minute)
			if err != nil {
				t.Fatal(err)
			}
			if e == ei {
				t.Fatalf("e's address should not be the same as ee's")
			}
			e.Exat = e.Exat.Truncate(time.Nanosecond)
			ei.Exat = ei.Exat.Truncate(time.Nanosecond)
			if !e.Exat.Equal(ei.Exat) {
				t.Fatalf("ee should be the same as e %v, %v", e, ei)
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
		entities := []*JSONTestTTLStruct{
			repo.NewEntity(),
			repo.NewEntity(),
			repo.NewEntity(),
		}

		for _, e := range entities {
			e.Exat = time.Now().Add(time.Minute)
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

		for _, e := range entities {
			exat, err := client.Do(ctx, client.B().Pexpiretime().Key("jsonttl:"+e.Key).Build()).AsInt64()
			if err != nil {
				t.Fatal(err)
			}
			if exat != e.Exat.UnixMilli() {
				t.Fatalf("wrong exat")
			}
		}
	})
}

func TestCreateAndAliasIndex_JSON(t *testing.T) {
	ctx := context.Background()

	client := setup(t)
	client.Do(ctx, client.B().Flushall().Build())
	defer client.Close()

	repo := NewJSONRepository("jsonalias", JSONTestStruct{}, client)

	t.Run("CreateAndAliasIndex_JSON", func(t *testing.T) {
		err := repo.CreateAndAliasIndex(ctx, func(schema FtCreateSchema) rueidis.Completed {
			return schema.FieldName("$.val").As("val").Text().Build()
		})
		if err != nil {
			t.Fatalf("failed to create and alias JSON index: %v", err)
		}

		verifyAliasTarget(t, ctx, client, repo.IndexName(), repo.IndexName()+"_v1")

		err = repo.CreateAndAliasIndex(ctx, func(schema FtCreateSchema) rueidis.Completed {
			return schema.FieldName("$.val").As("val").Text().Build()
		})
		if err != nil {
			t.Fatalf("failed to create and alias new JSON index version: %v", err)
		}

		verifyAliasTarget(t, ctx, client, repo.IndexName(), repo.IndexName()+"_v2")
	})
}
