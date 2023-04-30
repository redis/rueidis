package om

import (
	"context"
	"strconv"
	"testing"

	"github.com/redis/rueidis"
)

type Book struct {
	Key   string `json:"key"   redis:",key"`
	Loc   string `json:"loc"   redis:",loc"`
	Ver   int64  `json:"ver"   redis:",ver"`
	ID    int64  `json:"id"    redis:",id"`
	Count int64  `json:"count" redis:",count"`
}

func TestAggregateCursor(t *testing.T) {
	ctx := context.Background()

	client := setup(t)
	client.Do(ctx, client.B().Flushall().Build())
	defer client.Close()

	jsonRepo := NewJSONRepository("book", Book{}, client)
	hashRepo := NewHashRepository("book", Book{}, client)

	if err := jsonRepo.CreateIndex(ctx, func(schema FtCreateSchema) rueidis.Completed {
		return schema.
			FieldName("$.id").As("id").Numeric().
			FieldName("$.loc").As("loc").Tag().
			FieldName("$.count").As("count").Numeric().Sortable().Build()
	}); err != nil {
		t.Fatal(err)
	}

	if err := hashRepo.CreateIndex(ctx, func(schema FtCreateSchema) rueidis.Completed {
		return schema.
			FieldName("id").As("id").Numeric().
			FieldName("loc").As("loc").Tag().
			FieldName("count").As("count").Numeric().Sortable().Build()
	}); err != nil {
		t.Fatal(err)
	}

	t.Run("Hash", func(t *testing.T) {
		testRepo(t, hashRepo)
	})

	t.Run("Json", func(t *testing.T) {
		testRepo(t, jsonRepo)
	})
}

func testRepo(t *testing.T, repo Repository[Book]) {
	for i := 0; i < 10; i++ {
		book := repo.NewEntity()
		book.ID = int64(i / 2)
		book.Count = int64(i)
		book.Loc = "1"
		if err := repo.Save(context.Background(), book); err != nil {
			panic(err)
		}
	}

	t.Run("Deadline exceed", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		_, err := repo.Aggregate(ctx, func(search FtAggregateIndex) rueidis.Completed {
			return search.Query("any").Build()
		})
		if err != context.Canceled {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("Without cursor", func(t *testing.T) {
		cursor, err := repo.Aggregate(context.Background(), func(search FtAggregateIndex) rueidis.Completed {
			return search.Query("@loc:{1}").
				Groupby(1).Property("@id").Reduce("MIN").Nargs(1).Arg("@count").As("minCount").
				Sortby(2).Property("@minCount").Asc().Build()
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		result, err := cursor.Read(context.Background())
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if len(result) != 5 || int64(len(result)) != cursor.Total() {
			t.Fatalf("unexpected total %v %v", len(result), cursor.Total())
		}
		if _, err = cursor.Read(context.Background()); err != EndOfCursor {
			t.Fatalf("unexpected err %v", err)
		}
		if err = cursor.Del(context.Background()); err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		for i, record := range result {
			if record["id"] != strconv.Itoa(i) {
				t.Fatalf("unexpected value %v", record["id"])
			}
			if record["minCount"] != strconv.Itoa(i*2) {
				t.Fatalf("unexpected value %v", record["minCount"])
			}
		}
	})

	t.Run("With cursor", func(t *testing.T) {
		cursor, err := repo.Aggregate(context.Background(), func(search FtAggregateIndex) rueidis.Completed {
			return search.Query("@loc:{1}").
				Groupby(1).Property("@id").Reduce("MIN").Nargs(1).Arg("@count").As("minCount").
				Sortby(2).Property("@minCount").Asc().Withcursor().Count(2).Build()
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if cursor.Total() != 5 {
			t.Fatalf("unexpected total %v", cursor.Total())
		}
		var result []map[string]string
		for {
			partial, err := cursor.Read(context.Background())
			if err == EndOfCursor {
				break
			}
			if len(partial) > 2 {
				t.Fatalf("unexpected partial len %v", len(partial))
			}
			result = append(result, partial...)
		}
		if err = cursor.Del(context.Background()); err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if len(result) != 5 {
			t.Fatalf("unexpected total %v", len(result))
		}
		for i, record := range result {
			if record["id"] != strconv.Itoa(i) {
				t.Fatalf("unexpected value %v", record["id"])
			}
			if record["minCount"] != strconv.Itoa(i*2) {
				t.Fatalf("unexpected value %v", record["minCount"])
			}
		}
	})

	t.Run("Read deadline", func(t *testing.T) {
		cursor, err := repo.Aggregate(context.Background(), func(search FtAggregateIndex) rueidis.Completed {
			return search.Query("@loc:{1}").
				Groupby(1).Property("@id").Reduce("MIN").Nargs(1).Arg("@count").As("minCount").
				Sortby(2).Property("@minCount").Asc().Withcursor().Count(2).Build()
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		if _, err = cursor.Read(ctx); err != nil { // initial result is not affected by ctx
			t.Fatalf("unexpected err %v", err)
		}
		if _, err = cursor.Read(ctx); err != context.Canceled {
			t.Fatalf("unexpected err %v", err)
		}
	})

	t.Run("Del cursor", func(t *testing.T) {
		cursor, err := repo.Aggregate(context.Background(), func(search FtAggregateIndex) rueidis.Completed {
			return search.Query("@loc:{1}").
				Groupby(1).Property("@id").Reduce("MIN").Nargs(1).Arg("@count").As("minCount").
				Sortby(2).Property("@minCount").Asc().Withcursor().Count(2).Build()
		})
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		if err := cursor.Del(context.Background()); err != nil {
			t.Fatalf("unexpected err %v", err)
		}
	})
}
