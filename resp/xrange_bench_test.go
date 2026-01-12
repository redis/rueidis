package resp

import (
	"bufio"
	"context"
	"strconv"
	"testing"

	"github.com/redis/rueidis"
)

func BenchmarkXRange1000Entries(b *testing.B) {
	ctx := context.Background()
	client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
	if err != nil {
		b.Fatal(err)
	}
	defer client.Close()

	// Setup: Create a stream with 1000 entries
	streamKey := "benchmark:xrange:stream"

	// Clean up any existing stream
	client.Do(ctx, client.B().Del().Key(streamKey).Build())

	// Add 1000 entries to the stream
	for i := 0; i < 1000; i++ {
		err := client.Do(ctx, client.B().Xadd().
			Key(streamKey).
			Id("*").
			FieldValue().
			FieldValue("field1", "value"+strconv.Itoa(i)).
			Build()).Error()
		if err != nil {
			b.Fatal(err)
		}
	}

	// Benchmark 1: Using client.Do with AsXRangeSlices
	b.Run("Do_AsXRangeSlices", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()

		for b.Loop() {
			result := client.Do(ctx, client.B().Xrange().
				Key(streamKey).
				Start("-").
				End("+").
				Build())

			if result.Error() != nil {
				b.Fatal(result.Error())
			}

			entries, err := result.AsXRangeSlices()
			if err != nil {
				b.Fatal(err)
			}

			if len(entries) != 1000 {
				b.Fatalf("expected 1000 entries, got %d", len(entries))
			}
		}
	})

	// Benchmark 2: Using client.DoWithReader with resp package
	b.Run("DoWithReader_RespParse", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()

		for b.Loop() {
			var entryCount int

			var results [][]byte
			err := client.DoWithReader(ctx, client.B().Xrange().
				Key(streamKey).
				Start("-").
				End("+").
				Build(), func(reader *bufio.Reader) error {

				respReader := NewReader(reader)

				// Expect array of entries
				count, err := respReader.ExpectArray()
				if err != nil {
					return err
				}

				entryCount = int(count)
				results = make([][]byte, entryCount)

				// Parse each entry
				for j := int64(0); j < count; j++ {
					// Each entry is a 2-element array [id, fields]
					if err := respReader.ExpectArrayWithLen(2); err != nil {
						return err
					}

					// Read and discard ID
					if _, err := respReader.ReadStringBytes(); err != nil {
						return err
					}

					// Read field-value pairs
					fieldCount, err := respReader.ExpectArray()
					if err != nil {
						return err
					}

					if fieldCount != 2 {
						panic("expected 2 fields, got " + strconv.Itoa(int(fieldCount)))
					}

					_, err = respReader.ReadStringBytes() // Key.
					if err != nil {
						return err
					}
					buf, err := respReader.ReadStringBytes() // Value.
					if err != nil {
						return err
					}
					// copy buf.
					safeBuf := make([]byte, len(buf))
					copy(safeBuf, buf)
					results[j] = safeBuf
				}

				return nil
			})

			if err != nil {
				b.Fatal(err)
			}

			if len(results) != 1000 {
				b.Fatalf("expected 1000 entries, got %d", entryCount)
			}
		}
	})

	// Cleanup
	client.Do(ctx, client.B().Del().Key(streamKey).Build())
}
