package om

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/redis/rueidis"
)

type HashTestStruct struct {
	Key   string `redis:",key"`
	F2    *bool
	F2N   *bool
	F3    *string
	F3N   *string
	F4    *int64
	F4N   *int64
	Val   []byte
	Ver   int64 `redis:",ver"`
	F1    bool
	F5    *bool
	Vec32 []float32
	Vec64 []float64
	JSON  json.RawMessage
}

type Unsupported struct {
	Key string `redis:",key"`
	Ver int64  `redis:",ver"`
	F1  int32
}

type Mismatch struct {
	Key string `redis:",key"`
	Ver int64  `redis:",ver"`
	F1  int64
	F2  *int64
}

func TestNewHashRepositoryPanic(t *testing.T) {
	if v := recovered(func() {
		NewHashRepository("hash", Unsupported{}, nil)
	}); !strings.Contains(v, "should not contain unsupported field type") {
		t.Fatalf("unexpected message %v", v)
	}
}

func TestNewHashRepositoryMismatch(t *testing.T) {
	ctx := context.Background()

	client := setup(t)
	client.Do(ctx, client.B().Flushall().Build())
	defer client.Close()

	repo := NewHashRepository("hashmismatch", Mismatch{}, client)
	if err := repo.CreateIndex(ctx, func(schema FtCreateSchema) rueidis.Completed {
		return schema.FieldName("F1").Tag().Build()
	}); err != nil {
		t.Fatal(err)
	}

	t.Run("Mismatch int64", func(t *testing.T) {
		e := repo.NewEntity()
		if err := repo.Save(ctx, e); err != nil {
			t.Fatal(err)
		}
		if err := client.Do(ctx, client.B().Hmset().Key("hashmismatch:"+e.Key).FieldValue().FieldValue("F1", "").Build()).Error(); err != nil {
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

	t.Run("Mismatch *int64", func(t *testing.T) {
		e := repo.NewEntity()
		if err := repo.Save(ctx, e); err != nil {
			t.Fatal(err)
		}
		if err := client.Do(ctx, client.B().Hmset().Key("hashmismatch:"+e.Key).FieldValue().FieldValue("F2", "").Build()).Error(); err != nil {
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
func TestNewHashRepository(t *testing.T) {
	ctx := context.Background()

	client := setup(t)
	client.Do(ctx, client.B().Flushall().Build())
	defer client.Close()

	repo := NewHashRepository("hash", HashTestStruct{}, client)

	t.Run("NewEntity", func(t *testing.T) {
		e := repo.NewEntity()
		ulid.MustParse(e.Key)
	})

	t.Run("Save", func(t *testing.T) {
		f4 := rand.Int63()
		e := repo.NewEntity()
		f := false
		// test save
		e.Val = []byte("any")
		e.F1 = true
		e.F2 = &e.F1
		e.F3 = &e.Key
		e.F4 = &f4
		e.F5 = &f
		e.Vec32 = []float32{3, 2, 1}
		e.Vec64 = []float64{1, 2, 3}
		e.JSON = []byte(`[1]`)
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
			if e == ei {
				t.Fatalf("e's address should not be the same as ee's")
			}
			if !reflect.DeepEqual(e, ei) {
				t.Fatalf("ee should be the same as e")
			}
		})

		t.Run("Search", func(t *testing.T) {
			err := repo.CreateIndex(ctx, func(schema FtCreateSchema) rueidis.Completed {
				return schema.FieldName("Val").Text().Build()
			})
			time.Sleep(time.Second)
			if err != nil {
				t.Fatal(err)
			}
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
				return schema.FieldName("Val").Text().Sortable().Build()
			})
			time.Sleep(time.Second)
			if err != nil {
				t.Fatal(err)
			}
			n, records, err := repo.Search(ctx, func(search FtSearchIndex) rueidis.Completed {
				return search.Query("*").Sortby("Val").Build()
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
				return schema.FieldName("Val").Text().Build()
			})
			if err != nil {
				t.Fatal(err)
			}
			time.Sleep(time.Second)
			var entities []*HashTestStruct
			for i := 3; i >= 1; i-- {
				e := repo.NewEntity()
				e.Val = []byte("any")
				e.Vec32 = []float32{3, 2, 1}
				e.Vec64 = []float64{1, 2, 3}
				e.JSON = []byte(fmt.Sprintf("[%d]", i))
				err = repo.Save(ctx, e)
				if err != nil {
					t.Fatal(err)
				}
				entities = append(entities, e)
			}
			time.Sleep(time.Second)
			_, _, err = repo.Search(ctx, func(search FtSearchIndex) rueidis.Completed {
				return search.Query("*").Sortby("JSON").Build()
			})
			if err == nil {
				t.Fatalf("search by property neither loaded nor in schema")
			}
			err = repo.AlterIndex(ctx, func(alter FtAlterIndex) rueidis.Completed {
				return alter.
					Schema().Add().Field("JSON").Options("TEXT", "SORTABLE").
					Build()
			})
			if err != nil {
				t.Fatal(err)
			}
			time.Sleep(time.Second)
			n, records, err := repo.Search(ctx, func(search FtSearchIndex) rueidis.Completed {
				return search.Query("*").Sortby("JSON").Build()
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
		entities := []*HashTestStruct{
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

type HashTestTTLStruct struct {
	Key  string    `redis:",key"`
	Ver  int64     `redis:",ver"`
	Exat time.Time `redis:",exat"`
}

//gocyclo:ignore
func TestNewHashRepositoryTTL(t *testing.T) {
	ctx := context.Background()

	client := setup(t)
	client.Do(ctx, client.B().Flushall().Build())
	defer client.Close()

	repo := NewHashRepository("hashttl", HashTestTTLStruct{}, client)

	t.Run("NewEntity", func(t *testing.T) {
		e := repo.NewEntity()
		ulid.MustParse(e.Key)
	})

	t.Run("Save", func(t *testing.T) {
		e := repo.NewEntity()
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
			exat, err := client.Do(ctx, client.B().Pexpiretime().Key("hashttl:"+e.Key).Build()).AsInt64()
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
			e.Exat = e.Exat.Truncate(time.Millisecond)
			ei.Exat = ei.Exat.Truncate(time.Millisecond)
			if !e.Exat.Equal(ei.Exat) {
				t.Fatalf("e should be the same as ee %v %v", e, ei)
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
			e.Exat = e.Exat.Truncate(time.Millisecond)
			ei.Exat = ei.Exat.Truncate(time.Millisecond)
			if !e.Exat.Equal(ei.Exat) {
				t.Fatalf("ee should be the same as e %v %v", e, ei)
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
		entities := []*HashTestTTLStruct{
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
			exat, err := client.Do(ctx, client.B().Pexpiretime().Key("hashttl:"+e.Key).Build()).AsInt64()
			if err != nil {
				t.Fatal(err)
			}
			if exat != e.Exat.UnixMilli() {
				t.Fatalf("wrong exat")
			}
		}
	})
}

// TestCreateAndAliasIndex tests the CreateAndAliasIndex method of HashRepository.
func TestCreateAndAliasIndex(t *testing.T) {
	ctx := context.Background()

	client := setup(t)
	client.Do(ctx, client.B().Flushall().Build())
	defer client.Close()

	repo := NewHashRepository("hashalias", HashTestStruct{}, client)

	t.Run("CreateAndAliasIndex", func(t *testing.T) {
		err := repo.CreateAndAliasIndex(ctx, func(schema FtCreateSchema) rueidis.Completed {
			return schema.FieldName("Val").Text().Build()
		})
		if err != nil {
			t.Fatalf("failed to create and alias index: %v", err)
		}

		verifyAliasTarget(t, ctx, client, repo.IndexName(), repo.IndexName()+"_v1")

		// Step 3: Create new index version and update alias
		err = repo.CreateAndAliasIndex(ctx, func(schema FtCreateSchema) rueidis.Completed {
			return schema.FieldName("Val").Text().Build()
		})
		if err != nil {
			t.Fatalf("failed to create and alias new index version: %v", err)
		}

		verifyAliasTarget(t, ctx, client, repo.IndexName(), repo.IndexName()+"_v2")
	})
}

// Helper to verify that alias points to the expected index name
func verifyAliasTarget(t *testing.T, ctx context.Context, client rueidis.Client, aliasName string, expectedIndex string) {
	t.Helper()

	infoCmd := client.B().FtInfo().Index(aliasName).Build()
	infoResp, err := client.Do(ctx, infoCmd).ToMap()
	if err != nil {
		t.Fatalf("failed to fetch index info: %v", err)
	}

	indexMsg, ok := infoResp["index_name"]
	if !ok {
		t.Fatalf("FT.INFO response missing index_name field")
	}

	actualIndex, err := (&indexMsg).ToString()
	if err != nil {
		t.Fatalf("failed to convert index_name to string: %v", err)
	}

	if actualIndex != expectedIndex {
		t.Fatalf("alias does not point to the expected index. expected=%s got=%s", expectedIndex, actualIndex)
	}
}
