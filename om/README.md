# Generic Object Mapping

The `om.NewHashRepository` and `om.NewJSONRepository` creates an OM repository backed by redis hash or RedisJSON.

```golang
package main

import (
    "context"
    "fmt"
    "time"

    "github.com/redis/rueidis"
    "github.com/redis/rueidis/om"
)

type Example struct {
    Key string `json:"key" redis:",key"` // the redis:",key" is required to indicate which field is the ULID key
    Ver int64  `json:"ver" redis:",ver"` // the redis:",ver" is required to do optimistic locking to prevent lost update
    Str string `json:"myStr"`            // both NewHashRepository and NewJSONRepository use json tag as field name
}

func main() {
    ctx := context.Background()
    c, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
    if err != nil {
        panic(err)
    }
    // create the repo with NewHashRepository or NewJSONRepository
    repo := om.NewHashRepository("my_prefix", Example{}, c)

    exp := repo.NewEntity()
    exp.Str = "mystr"
    fmt.Println(exp.Key) // output 01FNH4FCXV9JTB9WTVFAAKGSYB
    repo.Save(ctx, exp) // success

    // lookup "my_prefix:01FNH4FCXV9JTB9WTVFAAKGSYB" through client side caching
    exp2, _ := repo.FetchCache(ctx, exp.Key, time.Second*5)
    fmt.Println(exp2.Str) // output "mystr", which equals to exp.Str

    exp2.Ver = 0         // if someone changes the version during your GET then SET operation,
    repo.Save(ctx, exp2) // the save will fail with ErrVersionMismatch.
}

```

### Object Mapping + RediSearch

If you have RediSearch, you can create and search the repository against the index.

```golang

if _, ok := repo.(*om.HashRepository[Example]); ok {
    repo.CreateIndex(ctx, func(schema om.FtCreateSchema) rueidis.Completed {
        return schema.FieldName("myStr").Text().Build() // Note that the Example.Str field is mapped to myStr on redis by its json tag
    })
}

if _, ok := repo.(*om.JSONRepository[Example]); ok {
    repo.CreateIndex(ctx, func(schema om.FtCreateSchema) rueidis.Completed {
        return schema.FieldName("$.myStr").Text().Build() // the field name of json index should be a json path syntax
    })
}

exp := repo.NewEntity()
exp.Str = "foo"
repo.Save(ctx, exp)

n, records, _ := repo.Search(ctx, func(search om.FtSearchIndex) rueidis.Completed {
    return search.Query("foo").Build() // you have full query capability by building the command from om.FtSearchIndex
})

fmt.Println("total", n) // n is total number of results matched in redis, which is >= len(records)

for _, v := range records {
    fmt.Println(v.Str) // print "foo"
}
```

### Change Search Index Name

The default index name for `HashRepository` and `JSONRepository` is `hashidx:{prefix}` and `jsonidx:{prefix}` respectively.

They can be changed by `WithIndexName` option to allow searching difference indexes:

```golang
repo1 := om.NewHashRepository("my_prefix", Example{}, c, om.WithIndexName("my_index1"))
repo2 := om.NewHashRepository("my_prefix", Example{}, c, om.WithIndexName("my_index2"))
```

### Object Mapping Limitation

`NewHashRepository` only accepts these field types:
* `string`, `*string`
* `int64`, `*int64`
* `bool`, `*bool`
* `[]byte`, `json.RawMessage`
* `[]float32`, `[]float64` for vector search

Field projection by RediSearch is not supported.