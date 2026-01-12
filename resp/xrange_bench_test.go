package resp

import (
	"bufio"
	"context"
	"strconv"
	"testing"
	"time"

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
	b.Run("DoWithReader", func(b *testing.B) {
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
						b.Fatalf("expected 2 fields, got %d", fieldCount)
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

func process(val []byte) {
	if len(val) == 0 {
		panic("empty value")
	}
}

func BenchmarkXRead1000Entries(b *testing.B) {
	ctx := context.Background()
	client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
	if err != nil {
		b.Fatal(err)
	}
	defer client.Close()

	// Setup: Create a stream with 1000 entries
	streamKey := "benchmark:xread:stream"

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

	// Benchmark 1: Using client.Do with AsXRead
	b.Run("Do_AsXRead", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()

		for b.Loop() {
			lastID := "0-0"
			totalEntries := 0

			// Read in batches of 10 until we've read all 1000 entries
			for totalEntries < 1000 {
				result := client.Do(ctx, client.B().Xread().
					Count(100).
					Streams().
					Key(streamKey).
					Id(lastID).
					Build())

				if result.Error() != nil {
					b.Fatal(result.Error())
				}

				streams, err := result.AsXRead()
				if err != nil {
					b.Fatal(err)
				}

				if len(streams) == 0 {
					break
				}

				entries := streams[streamKey]
				if len(entries) == 0 {
					break
				}

				totalEntries += len(entries)
				lastID = entries[len(entries)-1].ID
			}

			if totalEntries != 1000 {
				b.Fatalf("expected 1000 entries, got %d", totalEntries)
			}
		}
	})

	// Benchmark 2: Using client.DoWithReader with resp package
	b.Run("DoWithReader", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()

		for b.Loop() {
			lastID := "0-0"
			totalEntries := 0

			// Read in batches of 10 until we've read all 1000 entries
			for totalEntries < 1000 {
				var batchEntries int
				var newLastID string

				err := client.DoWithReader(ctx, client.B().Xread().
					Count(100).
					Streams().
					Key(streamKey).
					Id(lastID).
					Build(), func(reader *bufio.Reader) error {

					respReader := NewReader(reader)

					// Check if response is null (no data)
					if respReader.PeekKind() == KindNull {
						_ = respReader.ReadNull()
						return nil
					}

					var streamCount int64
					var err error
					var isMap bool

					// XREAD returns a map in RESP3, array in RESP2
					if respReader.PeekKind() == KindMap {
						// RESP3: Map of stream_key -> entries_array
						streamCount, err = respReader.ExpectMap()
						if err != nil {
							return err
						}
						isMap = true
					} else {
						// RESP2: Array of [stream_key, entries_array]
						streamCount, err = respReader.ExpectArray()
						if err != nil {
							return err
						}
						isMap = false
					}

					if streamCount == 0 {
						return nil
					}

					// Parse each stream
					for i := int64(0); i < streamCount; i++ {
						if isMap {
							// RESP3 map: key and value are separate
							// Read and discard stream key (string)
							if _, err := respReader.ReadStringBytes(); err != nil {
								return err
							}
						} else {
							// RESP2 array: each element is [stream_key, entries_array]
							if err := respReader.ExpectArrayWithLen(2); err != nil {
								return err
							}
							// Read and discard stream key
							if _, err := respReader.ReadStringBytes(); err != nil {
								return err
							}
						}

						// Read entries array
						entryCount, err := respReader.ExpectArray()
						if err != nil {
							return err
						}

						batchEntries = int(entryCount)

						// Parse each entry
						for j := int64(0); j < entryCount; j++ {
							// [id, [field, value, ...]]
							if err := respReader.ExpectArrayWithLen(2); err != nil {
								return err
							}

							// Read ID
							id, err := respReader.ReadStringBytes()
							if err != nil {
								return err
							}

							// Save the last ID for next iteration
							if j == entryCount-1 {
								// Copy the ID bytes to be safe
								idCopy := make([]byte, len(id))
								copy(idCopy, id)
								newLastID = string(idCopy)
							}

							// Read field-value pairs
							fieldCount, err := respReader.ExpectArray()
							if err != nil {
								return err
							}

							// Read and process field-value pairs
							for k := int64(0); k < fieldCount; k += 2 {
								// Read field name
								if _, err := respReader.ReadStringBytes(); err != nil {
									return err
								}
								// Read field value
								buf, err := respReader.ReadStringBytes()
								if err != nil {
									return err
								}
								safeBuf := make([]byte, len(buf))
								copy(safeBuf, buf)
								process(safeBuf)
							}
						}
					}

					return nil
				})

				if err != nil {
					b.Fatal(err)
				}

				if batchEntries == 0 {
					break
				}

				totalEntries += batchEntries
				lastID = newLastID
			}

			if totalEntries != 1000 {
				b.Fatalf("expected 1000 entries, got %d", totalEntries)
			}
		}
	})

	// Cleanup
	client.Do(ctx, client.B().Del().Key(streamKey).Build())
}

func BenchmarkXReadStreaming(b *testing.B) {
	ctx := context.Background()
	client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
	if err != nil {
		b.Fatal(err)
	}
	defer client.Close()

	streamKey := "benchmark:xread:streaming"

	// Clean up any existing stream
	client.Do(ctx, client.B().Del().Key(streamKey).Build())

	// Benchmark 1: Using client.Do with AsXRead
	b.Run("Do_AsXRead", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()

		for b.Loop() {
			// Clean up before each iteration
			client.Do(ctx, client.B().Del().Key(streamKey).Build())

			// Start producer goroutine that adds entries
			done := make(chan struct{})
			go func() {
				i := 0
				for {
					select {
					case <-done:
						return
					default:
					}
					time.Sleep(time.Millisecond)
					err := client.Do(ctx, client.B().Xadd().
						Key(streamKey).
						Id("*").
						FieldValue().
						FieldValue("field1", "value"+strconv.Itoa(i)).
						Build()).Error()
					if err != nil {
						b.Error(err)
						return
					}
				}
			}()

			lastID := "$" // Start from new entries only
			totalEntries := 0

			// Consumer: read entries as they arrive using blocking XREAD
			for totalEntries < 100 {
				result := client.Do(ctx, client.B().Xread().
					Block(1000).
					Streams().
					Key(streamKey).
					Id(lastID).
					Build())

				if result.Error() != nil {
					b.Fatal(result.Error())
				}

				streams, err := result.AsXRead()
				if err != nil {
					b.Fatal(err)
				}

				if len(streams) == 0 {
					continue // Timeout, try again
				}

				entries := streams[streamKey]
				if len(entries) > 0 {
					totalEntries += len(entries)
					lastID = entries[len(entries)-1].ID
				}
			}

			close(done)

			if totalEntries != 100 {
				b.Fatalf("expected 100 entries, got %d", totalEntries)
			}
		}
	})

	// Benchmark 2: Using client.DoWithReader with resp package
	b.Run("DoWithReader", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()

		for b.Loop() {
			// Clean up before each iteration
			client.Do(ctx, client.B().Del().Key(streamKey).Build())

			// Start producer goroutine that adds entries
			done := make(chan struct{})
			go func() {
				i := 0
				for {
					select {
					case <-done:
						return
					default:
					}
					time.Sleep(time.Millisecond)
					err := client.Do(ctx, client.B().Xadd().
						Key(streamKey).
						Id("*").
						FieldValue().
						FieldValue("field1", "value"+strconv.Itoa(i)).
						Build()).Error()
					if err != nil {
						b.Error(err)
						return
					}
				}
			}()

			lastID := "$" // Start from new entries only
			totalEntries := 0

			// Consumer: read entries as they arrive using blocking XREAD
			for totalEntries < 100 {
				var batchEntries int
				var newLastID string

				err := client.DoWithReader(ctx, client.B().Xread().
					Block(1000). // Block for up to 1 second
					Streams().
					Key(streamKey).
					Id(lastID).
					Build(), func(reader *bufio.Reader) error {

					respReader := NewReader(reader)

					// Check if response is null (timeout, no new data)
					if respReader.PeekKind() == KindNull {
						_ = respReader.ReadNull()
						return nil
					}

					var streamCount int64
					var err error
					var isMap bool

					// XREAD returns a map in RESP3, array in RESP2
					if respReader.PeekKind() == KindMap {
						streamCount, err = respReader.ExpectMap()
						isMap = true
					} else {
						streamCount, err = respReader.ExpectArray()
						isMap = false
					}
					if err != nil {
						return err
					}

					if streamCount == 0 {
						return nil
					}

					// Parse each stream
					for i := int64(0); i < streamCount; i++ {
						if isMap {
							// RESP3 map: key and value are separate
							if _, err := respReader.ReadStringBytes(); err != nil {
								return err
							}
						} else {
							// RESP2 array: each element is [stream_key, entries_array]
							if err := respReader.ExpectArrayWithLen(2); err != nil {
								return err
							}
							if _, err := respReader.ReadStringBytes(); err != nil {
								return err
							}
						}

						// Read entries array
						entryCount, err := respReader.ExpectArray()
						if err != nil {
							return err
						}

						batchEntries = int(entryCount)

						// Parse each entry
						for j := int64(0); j < entryCount; j++ {
							// [id, [field, value, ...]]
							if err := respReader.ExpectArrayWithLen(2); err != nil {
								return err
							}

							// Read ID
							id, err := respReader.ReadStringBytes()
							if err != nil {
								return err
							}

							// Save the last ID for next iteration
							if j == entryCount-1 {
								idCopy := make([]byte, len(id))
								copy(idCopy, id)
								newLastID = string(idCopy)
							}

							// Read field-value pairs
							fieldCount, err := respReader.ExpectArray()
							if err != nil {
								return err
							}

							// Read and process field-value pairs
							for k := int64(0); k < fieldCount; k += 2 {
								// Read field name
								if _, err := respReader.ReadStringBytes(); err != nil {
									return err
								}
								// Read field value
								buf, err := respReader.ReadStringBytes()
								if err != nil {
									return err
								}
								safeBuf := make([]byte, len(buf))
								copy(safeBuf, buf)
								process(safeBuf)
							}
						}
					}

					return nil
				})

				if err != nil {
					b.Fatal(err)
				}

				if batchEntries > 0 {
					totalEntries += batchEntries
					lastID = newLastID
				}
			}

			close(done)

			if totalEntries != 100 {
				b.Fatalf("expected 100 entries, got %d", totalEntries)
			}
		}
	})

	// Cleanup
	client.Do(ctx, client.B().Del().Key(streamKey).Build())
}
