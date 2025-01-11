# rueidisprob

A Probabilistic Data Structures without Redis Stack.

## Features

### Bloom Filter

It is a space-efficient probabilistic data structure that is used to test whether an element is a member of a set.
False positive matches are possible, but false negatives are not. 
In other words, a query returns either "possibly in set" or "definitely not in set".
Elements can be added to the set, but not removed.

Example:

```go
package main

import (
	"context"
	"fmt"

	"github.com/redis/rueidis"
	"github.com/redis/rueidis/rueidisprob"
)

func main() {
	client, err := rueidis.NewClient(rueidis.ClientOption{
		InitAddress: []string{"localhost:6379"},
	})
	if err != nil {
		panic(err)
	}

	bf, err := rueidisprob.NewBloomFilter(client, "bloom_filter", 1000, 0.01)

	err = bf.Add(context.Background(), "hello")
	if err != nil {
		panic(err)
	}

	err = bf.Add(context.Background(), "world")
	if err != nil {
		panic(err)
	}

	exists, err := bf.Exists(context.Background(), "hello")
	if err != nil {
		panic(err)
	}
	fmt.Println(exists) // true

	exists, err = bf.Exists(context.Background(), "world")
	if err != nil {
		panic(err)
	}
	fmt.Println(exists) // true
}
```

### Counting Bloom Filter

It is a variation of the standard Bloom filter that adds a counting mechanism to each element.
This allows for the filter to count the number of times an element has been added to the filter.
And it allows for the removal of elements from the filter.

Example:

```go

package main

import (
    "context"
    "fmt"

    "github.com/redis/rueidis"
    "github.com/redis/rueidis/rueidisprob"
)

func main() {
    client, err := rueidis.NewClient(rueidis.ClientOption{
        InitAddress: []string{"localhost:6379"},
    })
    if err != nil {
        panic(err)
    }

    cbf, err := rueidisprob.NewCountingBloomFilter(client, "counting_bloom_filter", 1000, 0.01)

    err = cbf.Add(context.Background(), "hello")
    if err != nil {
        panic(err)
    }

    err = cbf.Add(context.Background(), "world")
    if err != nil {
        panic(err)
    }

    exists, err := cbf.Exists(context.Background(), "hello")
    if err != nil {
        panic(err)
    }
    fmt.Println(exists) // true

    exists, err = cbf.Exists(context.Background(), "world")
    if err != nil {
        panic(err)
    }
    fmt.Println(exists) // true

    count, err := cbf.Count(context.Background())
    if err != nil {
        panic(err)
    }
    fmt.Println(count) // 2

    err = cbf.Remove(context.Background(), "hello")
    if err != nil {
        panic(err)
    }

    exists, err = cbf.Exists(context.Background(), "hello")
    if err != nil {
        panic(err)
    }
    fmt.Println(exists) // false

    count, err = cbf.Count(context.Background())
    if err != nil {
        panic(err)
    }
    fmt.Println(count) // 1
}
```

### Sliding Window Bloom Filter

It is a variation of the standard Bloom filter that adds a sliding window mechanism. 
Useful for use cases where you need to keep track of items for a certain amount of time.

Example:

```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/rueidis"
	"github.com/redis/rueidis/rueidisprob"
)

func main() {
	client, err := rueidis.NewClient(rueidis.ClientOption{
		InitAddress: []string{"localhost:6379"},
	})
	if err != nil {
		panic(err)
	}

	sbf, err := NewSlidingBloomFilter(client, "sliding_bloom_filter", 1000, 0.01, time.Minute)

	err = sbf.Add(context.Background(), "hello")
	if err != nil {
		panic(err)
	}

	err = sbf.Add(context.Background(), "world")
	if err != nil {
		panic(err)
	}

	exists, err := sbf.Exists(context.Background(), "hello")
	if err != nil {
		panic(err)
	}
	fmt.Println(exists) // true

	exists, err = sbf.Exists(context.Background(), "world")
	if err != nil {
		panic(err)
	}
	fmt.Println(exists) // true

	count, err := sbf.Count(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(count) // 2
}
```