package resp

import (
	"bufio"
	"context"
	"strconv"
	"strings"
	"testing"
	"time"
	"unsafe"

	"github.com/redis/rueidis"
)

func BenchmarkXRange(b *testing.B) {
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

	totalEntries := 1000

	// Add 1000 entries to the stream
	for i := 0; i < totalEntries; i++ {
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

			if len(entries) != totalEntries {
				b.Fatalf("expected %d entries, got %d", totalEntries, len(entries))
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

			if len(results) != totalEntries {
				b.Fatalf("expected %d entries, got %d", totalEntries, entryCount)
			}
		}
	})

	// Cleanup
	client.Do(ctx, client.B().Del().Key(streamKey).Build())
}

var testVal []byte

func process(val []byte) {
	if len(val) == 0 {
		panic("empty value")
	}
	testVal = val
}

func BenchmarkXRead(b *testing.B) {
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

	totalEntries := 1000
	var batchSize int64 = 1000

	// Add entries to the stream
	for i := 0; i < totalEntries; i++ {
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
			consumedEntries := 0

			// Read in batches until we've read all entries
			for consumedEntries < totalEntries {
				result := client.Do(ctx, client.B().Xread().
					Count(batchSize).
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

				consumedEntries += len(entries)
				lastID = entries[len(entries)-1].ID
			}

			if consumedEntries != totalEntries {
				b.Fatalf("expected %d entries, got %d", totalEntries, consumedEntries)
			}
		}
	})

	// Benchmark 2: Using client.DoWithReader with resp package
	b.Run("DoWithReader", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()

		for b.Loop() {
			lastID := "0-0"
			consumedEntries := 0

			// Read in batches until we've read all entries
			for consumedEntries < totalEntries {
				var batchEntries int
				var newLastID string

				err := client.DoWithReader(ctx, client.B().Xread().
					Count(batchSize).
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

				consumedEntries += batchEntries
				lastID = newLastID
			}

			if consumedEntries != totalEntries {
				b.Fatalf("expected %d entries, got %d", totalEntries, consumedEntries)
			}
		}
	})

	// Cleanup
	client.Do(ctx, client.B().Del().Key(streamKey).Build())
}

func BenchmarkBlockingXRead(b *testing.B) {
	ctx := context.Background()
	client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
	if err != nil {
		b.Fatal(err)
	}
	defer client.Close()

	streamKey := "benchmark:xread:streaming"

	// Clean up any existing stream
	client.Do(ctx, client.B().Del().Key(streamKey).Build())

	totalEntries := 100
	var batchSize int64 = 1

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
			consumedEntries := 0

			// Consumer: read entries as they arrive using blocking XREAD
			for consumedEntries < totalEntries {
				result := client.Do(ctx, client.B().Xread().
					Count(batchSize).
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
					consumedEntries += len(entries)
					lastID = entries[len(entries)-1].ID
				}
			}

			close(done)

			if consumedEntries != totalEntries {
				b.Fatalf("expected %d entries, got %d", totalEntries, consumedEntries)
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
			consumedEntries := 0

			// Consumer: read entries as they arrive using blocking XREAD
			for consumedEntries < totalEntries {
				var batchEntries int
				var newLastID string

				err := client.DoWithReader(ctx, client.B().Xread().
					Count(batchSize).
					Block(1000).
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
					consumedEntries += batchEntries
					lastID = newLastID
				}
			}

			close(done)

			if consumedEntries != totalEntries {
				b.Fatalf("expected %d entries, got %d", totalEntries, consumedEntries)
			}
		}
	})

	// Cleanup
	client.Do(ctx, client.B().Del().Key(streamKey).Build())
}

// Simple publication struct for benchmark
type BenchPub struct {
	Offset uint64
	Data   []byte
}

var testPubs []BenchPub

func BenchmarkLuaScript(b *testing.B) {
	ctx := context.Background()
	client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
	if err != nil {
		b.Fatal(err)
	}
	defer client.Close()

	streamKey := "benchmark:lua:stream"
	metaKey := "benchmark:lua:meta"

	// Clean up any existing data
	client.Do(ctx, client.B().Del().Key(streamKey).Build())
	client.Do(ctx, client.B().Del().Key(metaKey).Build())

	totalEntries := 100
	requestLimit := strconv.Itoa(totalEntries)

	// Add entries to the stream
	for i := 0; i < totalEntries; i++ {
		err := client.Do(ctx, client.B().Xadd().
			Key(streamKey).
			Id("*").
			FieldValue().
			FieldValue("d", "value"+strconv.Itoa(i)).
			Build()).Error()
		if err != nil {
			b.Fatal(err)
		}
	}

	// Set up metadata
	client.Do(ctx, client.B().Hset().Key(metaKey).
		FieldValue().
		FieldValue("e", "v1").
		FieldValue("o", "100").
		Build())

	// Lua script that returns {offset, epoch, pubs}
	luaScript := rueidis.NewLuaScript(`
local stream_key = KEYS[1]
local meta_key = KEYS[2]
local start = ARGV[1]
local limit = ARGV[2]

local stream_meta = redis.call("hmget", meta_key, "e", "o")
local epoch = stream_meta[1]
local offset = stream_meta[2]

if epoch == false then
    epoch = "v1"
    offset = 0
    redis.call("hset", meta_key, "e", epoch, "o", offset)
end

if offset == false then
    offset = 0
end

local pubs = redis.call("xrange", stream_key, start, "+", "COUNT", limit)

return { offset, epoch, pubs }
`)

	// Benchmark 1: Standard approach using Exec
	b.Run("Do_Standard", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()

		for b.Loop() {
			result := luaScript.Exec(ctx, client,
				[]string{streamKey, metaKey},
				[]string{"-", requestLimit})

			if result.Error() != nil {
				b.Fatal(result.Error())
			}

			replies, err := result.ToArray()
			if err != nil {
				b.Fatal(err)
			}

			if len(replies) < 3 {
				b.Fatalf("expected at least 3 elements, got %d", len(replies))
			}

			// Parse offset
			offset, err := replies[0].AsUint64()
			if err != nil {
				b.Fatal(err)
			}

			// Parse epoch
			epoch, err := replies[1].ToString()
			if err != nil {
				b.Fatal(err)
			}

			// Parse publications
			pubValues, err := replies[2].ToArray()
			if err != nil {
				b.Fatal(err)
			}

			pubs := make([]BenchPub, len(pubValues))
			for i, v := range pubValues {
				entry, err := v.ToArray()
				if err != nil {
					b.Fatal(err)
				}
				if len(entry) != 2 {
					b.Fatal("invalid entry format")
				}

				// Parse ID to get offset
				id, err := entry[0].ToString()
				if err != nil {
					b.Fatal(err)
				}

				hyphenIdx := strings.IndexByte(id, '-')
				if hyphenIdx <= 0 {
					b.Fatal("invalid stream id")
				}
				pubOffset, err := strconv.ParseUint(id[:hyphenIdx], 10, 64)
				if err != nil {
					b.Fatal(err)
				}

				// Parse fields
				fieldsArr, err := entry[1].ToArray()
				if err != nil {
					b.Fatal(err)
				}

				var data []byte
				for j := 0; j < len(fieldsArr)-1; j += 2 {
					key, _ := fieldsArr[j].ToString()
					if key != "d" {
						continue
					}
					val, _ := fieldsArr[j+1].ToString()
					data = []byte(val)
					break
				}

				pubs[i] = BenchPub{
					Offset: pubOffset,
					Data:   data,
				}
			}

			if offset == 0 || epoch == "" || len(pubs) != totalEntries {
				b.Fatal("invalid result")
			}
			testPubs = pubs
		}
	})

	// Benchmark 2: Optimized approach using ExecWithReader
	b.Run("DoWithReader", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()

		for b.Loop() {
			var offset uint64
			var epoch string
			var pubs []BenchPub

			err := luaScript.ExecWithReader(ctx, client,
				[]string{streamKey, metaKey},
				[]string{"-", requestLimit},
				func(reader *bufio.Reader) error {

					respReader := NewReader(reader)

					// Expect top-level array: [offset, epoch, pubs]
					n, err := respReader.ExpectArray()
					if err != nil {
						return err
					}
					if n < 3 {
						return err
					}

					// Parse offset (can be int or string)
					switch respReader.PeekKind() {
					case KindInt:
						v, err := respReader.ReadInt64()
						if err != nil {
							return err
						}
						offset = uint64(v)
					case KindString:
						buf, err := respReader.ReadStringBytes()
						if err != nil {
							return err
						}
						// Use unsafe.String to avoid allocation
						v, err := strconv.ParseUint(
							unsafe.String(&buf[0], len(buf)),
							10,
							64,
						)
						if err != nil {
							return err
						}
						offset = v
					default:
						return err
					}

					// Parse epoch (string)
					epochBuf, err := respReader.ReadStringBytes()
					if err != nil {
						return err
					}
					// Make a safe copy
					epochCopy := make([]byte, len(epochBuf))
					copy(epochCopy, epochBuf)
					epoch = rueidis.BinaryString(epochCopy)

					// Parse publications array
					pubCount, err := respReader.ExpectArray()
					if err != nil {
						return err
					}

					pubs = make([]BenchPub, pubCount)

					for i := int64(0); i < pubCount; i++ {
						// Each entry: [id, fields]
						if _, err := respReader.ExpectArray(); err != nil {
							return err
						}

						// Parse ID to extract offset
						idBuf, err := respReader.ReadStringBytes()
						if err != nil {
							return err
						}

						// Find hyphen in ID without allocation using unsafe.String
						idStr := unsafe.String(&idBuf[0], len(idBuf))
						hyphenIdx := strings.IndexByte(idStr, '-')
						if hyphenIdx <= 0 {
							return err
						}

						// Parse offset from ID without allocation
						pubOffset, err := strconv.ParseUint(
							idStr[:hyphenIdx],
							10,
							64,
						)
						if err != nil {
							return err
						}

						// Parse field-value pairs
						fieldCount, err := respReader.ExpectArray()
						if err != nil {
							return err
						}

						var data []byte

						for j := int64(0); j < fieldCount; j += 2 {
							// Read key
							key, err := respReader.ReadStringBytes()
							if err != nil {
								return err
							}

							// Check if this is the data field
							isData := len(key) == 1 && key[0] == 'd'

							if isData {
								// Read value and process on-the-fly
								vbuf, err := respReader.ReadStringBytes()
								if err != nil {
									return err
								}
								// Copy to our struct (simulating conversion)
								data = make([]byte, len(vbuf))
								copy(data, vbuf)
							} else {
								// Skip other fields
								if err := respReader.SkipValue(); err != nil {
									return err
								}
							}
						}

						pubs[i] = BenchPub{
							Offset: pubOffset,
							Data:   data,
						}
					}

					return nil
				})

			if err != nil {
				b.Fatal(err)
			}

			if offset == 0 || epoch == "" || len(pubs) != totalEntries {
				b.Fatal("invalid result")
			}
			testPubs = pubs
		}
	})

	// Cleanup
	client.Do(ctx, client.B().Del().Key(streamKey).Build())
	client.Do(ctx, client.B().Del().Key(metaKey).Build())
}
