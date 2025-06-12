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
    Key  string    `json:"key" redis:",key"`   // the redis:",key" is required to indicate which field is the ULID key
    Ver  int64     `json:"ver" redis:",ver"`   // the redis:",ver" is required to do optimistic locking to prevent lost update
    ExAt time.Time `json:"exat" redis:",exat"` // the redis:",exat" is optional for setting record expiry with unix timestamp
    Str  string    `json:"str"`                // both NewHashRepository and NewJSONRepository use json tag as field name
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
    exp.ExAt = time.Now().Add(time.Hour)
    fmt.Println(exp.Key) // output 01FNH4FCXV9JTB9WTVFAAKGSYB
    repo.Save(ctx, exp) // success

    // lookup "my_prefix:01FNH4FCXV9JTB9WTVFAAKGSYB" through client side caching
    exp2, _ := repo.FetchCache(ctx, exp.Key, time.Second*5)
    fmt.Println(exp2.Str) // output "mystr", which is equal to exp.Str

    exp2.Ver = 0         // if someone changes the version during your GET then SET operation,
    repo.Save(ctx, exp2) // the save will fail with ErrVersionMismatch.
}

```

### Object Mapping + RediSearch

If you have RediSearch, you can create and search the repository against the index.

```golang
if _, ok := repo.(*om.HashRepository[Example]); ok {
    repo.CreateIndex(ctx, func(schema om.FtCreateSchema) rueidis.Completed {
        return schema.FieldName("str").Tag().Build() // Note that the Example.Str field is mapped to str on redis by its json tag
    })
}

if _, ok := repo.(*om.JSONRepository[Example]); ok {
    repo.CreateIndex(ctx, func(schema om.FtCreateSchema) rueidis.Completed {
        return schema.FieldName("$.str").As("str").Tag().Build() // the FieldName of a JSON index should be a JSON path syntax
    })
}

exp := repo.NewEntity()
exp.Str = "special_chars:[$.-]"
repo.Save(ctx, exp)

n, records, _ := repo.Search(ctx, func(search om.FtSearchIndex) rueidis.Completed {
    // Note that by using the named parameters with DIALECT >= 2, you won't have to escape chars for building queries.
    return search.Query("@str:{$v}").Params().Nargs(2).NameValue().NameValue("v", exp.Str).Dialect(2).Build()
})

fmt.Println("total", n) // n is the total number of results matched in redis, which is >= len(records)

for _, v := range records {
    fmt.Println(v.Str) // print "special_chars:[$.-]"
}
```

### Change Search Index Name

The default index name for `HashRepository` and `JSONRepository` is `hashidx:{prefix}` and `jsonidx:{prefix}` respectively.

They can be changed by `WithIndexName` option to allow searching difference indexes:

```golang
repo1 := om.NewHashRepository("my_prefix", Example{}, c, om.WithIndexName("my_index1"))
repo2 := om.NewHashRepository("my_prefix", Example{}, c, om.WithIndexName("my_index2"))
```

### Object Expiry Timestamp

Setting a `redis:",exat"` tag on a `time.Time` field will set `PEXPIREAT` on the record accordingly when calling `.Save()`.

If the `time.Time` is zero, then the expiry will be untouched when calling `.Save()`.

### Object Mapping Limitation

`NewHashRepository` only accepts these field types:
* `string`, `*string`
* `int64`, `*int64`
* `bool`, `*bool`
* `[]byte`, `json.RawMessage`
* `[]float32`, `[]float64` for vector search
* `json.Marshaler+json.Unmarshaler`

Field projection by RediSearch is not supported.