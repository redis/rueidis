package rueidis_test

import (
	"bufio"
	"context"
	"fmt"
	"io"

	"github.com/redis/rueidis"
)

// Example_doWithReader demonstrates zero-allocation parsing with DoWithReader
func Example_doWithReader() {
	client, err := rueidis.NewClient(rueidis.ClientOption{
		InitAddress: []string{"127.0.0.1:6379"},
	})
	if err != nil {
		panic(err)
	}
	defer client.Close()

	ctx := context.Background()

	// Example 1: Parse string value with zero allocations
	// IMPORTANT: All errors must be checked and returned!
	var result string
	err = client.DoWithReader(ctx, client.B().Get().Key("mykey").Build(),
		func(r *bufio.Reader, typ byte) error {
			if typ == '$' { // Blob string
				length, err := rueidis.ReadInt(r)
				if err != nil {
					return err
				}
				if length == -1 {
					return nil // NULL value
				}
				buf := make([]byte, length)
				if _, err := io.ReadFull(r, buf); err != nil {
					return err // MUST check error
				}
				if _, err := r.Discard(2); err != nil { // \r\n
					return err // MUST check error
				}
				result = string(buf)
			}
			return nil
		})
	if err != nil && err != rueidis.Nil {
		panic(err)
	}

	fmt.Printf("Got value: %s\n", result)

	// Example 2: Stream large value directly to writer (zero-copy)
	err = client.DoWithReader(ctx, client.B().Get().Key("large-value").Build(),
		func(r *bufio.Reader, typ byte) error {
			if typ == '$' {
				length, err := rueidis.ReadInt(r)
				if err != nil {
					return err
				}
				if length > 0 {
					// Stream directly to output without intermediate allocation
					if _, err := io.CopyN(io.Discard, r, length); err != nil {
						return err
					}
					if _, err := r.Discard(2); err != nil {
						return err
					}
				}
			}
			return nil
		})
	if err != nil && err != rueidis.Nil {
		panic(err)
	}

	// Example 3: Parse array without allocating intermediate slice
	var count int
	err = client.DoWithReader(ctx, client.B().Lrange().Key("mylist").Start(0).Stop(-1).Build(),
		func(r *bufio.Reader, typ byte) error {
			if typ == '*' { // Array
				arrayLen, err := rueidis.ReadInt(r)
				if err != nil {
					return err
				}
				for i := int64(0); i < arrayLen; i++ {
					elemTyp, err := r.ReadByte()
					if err != nil {
						return err
					}
					if elemTyp == '$' {
						length, err := rueidis.ReadInt(r)
						if err != nil {
							return err
						}
						if length > 0 {
							// Process each element directly without storing
							if _, err := r.Discard(int(length) + 2); err != nil { // Skip data + \r\n
								return err
							}
							count++
						}
					}
				}
			}
			return nil
		})
	if err != nil && err != rueidis.Nil {
		panic(err)
	}

	fmt.Printf("Processed %d items\n", count)
}
