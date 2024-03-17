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
    
    bf := rueidisprob.NewBloomFilter(client, "bloom_filter", 1000, 0.01)
    
    err := bf.Add("hello")
    if err != nil {
        panic(err)
    }
    
    err := bf.Add("world")
    if err != nil {
        panic(err)
    }
    
    exists, err := bf.Exists("hello")
    if err != nil {
        panic(err)
    }
    fmt.Println(exists) // true
    
    exists, err = bf.Exists("world")
    if err != nil {
        panic(err)
    }
    fmt.Println(exists) // true
}
```
