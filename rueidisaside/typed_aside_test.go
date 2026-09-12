package rueidisaside

import (
	"context"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/redis/rueidis"
)

type testStruct struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func TestTypedCacheAsideClient_Get(t *testing.T) {
	baseClient := makeClient(t, addr)
	t.Cleanup(baseClient.Close)

	serializer := func(v *testStruct) (string, error) {
		if v == nil {
			return "nilTestStruct", nil
		}
		b, err := json.Marshal(v)
		return string(b), err
	}
	deserializer := func(s string) (*testStruct, error) {
		if s == "nilTestStruct" {
			return nil, nil
		}
		var v testStruct
		err := json.Unmarshal([]byte(s), &v)
		return &v, err
	}

	client := NewTypedCacheAsideClient[testStruct](baseClient, serializer, deserializer)

	t.Run("successful get and cache", func(t *testing.T) {
		expected := &testStruct{ID: 1, Name: "test"}
		key := randStr()
		val, err := client.Get(context.Background(), time.Second, key, func(ctx context.Context, key string) (*testStruct, error) {
			return expected, nil
		})
		if err != nil {
			t.Fatal(err)
		}
		if val.ID != expected.ID || val.Name != expected.Name {
			t.Fatalf("expected %v, got %v", expected, val)
		}

		// Test cached value
		val2, err := client.Get(context.Background(), time.Second, key, nil)
		if err != nil {
			t.Fatal(err)
		}
		if val.ID != expected.ID || val.Name != expected.Name {
			t.Fatalf("cached value mismatch: expected %v, got %v", expected, val2)
		}
	})

	t.Run("serialization error", func(t *testing.T) {
		badSerializer := func(v *testStruct) (string, error) {
			return "", errors.New("serialization error")
		}
		clientWithBadSerializer := NewTypedCacheAsideClient[testStruct](baseClient, badSerializer, deserializer)

		key := randStr()
		_, err := clientWithBadSerializer.Get(context.Background(), time.Second, key, func(ctx context.Context, key string) (*testStruct, error) {
			return &testStruct{ID: 1, Name: "test"}, nil
		})

		if err == nil {
			t.Fatal("expected serialization error")
		}
	})

	t.Run("deserialization error", func(t *testing.T) {
		badDeserializer := func(s string) (*testStruct, error) {
			return nil, errors.New("deserialization error")
		}
		clientWithBadDeserializer := NewTypedCacheAsideClient[testStruct](baseClient, serializer, badDeserializer)

		key := randStr()
		_, err := clientWithBadDeserializer.Get(context.Background(), time.Second, key, func(ctx context.Context, key string) (*testStruct, error) {
			return &testStruct{ID: 1, Name: "test"}, nil
		})

		if err == nil {
			t.Fatal("expected deserialization error")
		}
	})

	t.Run("nil value handling", func(t *testing.T) {
		key := randStr()
		val, err := client.Get(context.Background(), time.Second, key, func(ctx context.Context, key string) (*testStruct, error) {
			return nil, nil
		})
		if err != nil {
			t.Fatalf("nil value should not return error: %v", err)
		}
		if val != nil {
			t.Fatalf("expected nil value, got %v", val)
		}
	})
}

func TestTypedCacheAsideClient_Del(t *testing.T) {
	baseClient := makeClient(t, addr)
	t.Cleanup(baseClient.Close)

	serializer := func(v *testStruct) (string, error) {
		b, err := json.Marshal(v)
		return string(b), err
	}

	deserializer := func(s string) (*testStruct, error) {
		var v testStruct
		err := json.Unmarshal([]byte(s), &v)
		return &v, err
	}

	client := NewTypedCacheAsideClient[testStruct](baseClient, serializer, deserializer)

	// Set a value first
	key := randStr()
	testVal := &testStruct{ID: 1, Name: "test"}
	_, err := client.Get(context.Background(), time.Second, key, func(ctx context.Context, key string) (*testStruct, error) {
		return testVal, nil
	})
	if err != nil {
		t.Fatal(err)
	}

	// Verify it's cached
	_, err = client.Get(context.Background(), time.Second, key, func(ctx context.Context, key string) (*testStruct, error) {
		t.Fatal("this function should not be called because the value should be cached")
		return testVal, nil
	})
	if err != nil {
		t.Fatal(err)
	}

	// Delete the value
	err = client.Del(context.Background(), key)
	if err != nil {
		t.Fatal(err)
	}

	// Verify it's deleted
	called := false
	_, err = client.Get(context.Background(), time.Second, key, func(ctx context.Context, key string) (val *testStruct, err error) {
		called = true
		return testVal, nil
	})
	if err != nil && !rueidis.IsRedisNil(err) {
		t.Fatal("expected error for deleted key")
	}
	if !called {
		t.Fatal("expected function to be called because the value should be deleted")
	}
}

func TestTypedCacheAsideClient_OverrideTTL(t *testing.T) {
	baseClient := makeClient(t, addr)
	t.Cleanup(baseClient.Close)

	serializer := func(v *testStruct) (string, error) {
		if v == nil {
			return "nilTestStruct", nil
		}
		b, err := json.Marshal(v)
		return string(b), err
	}

	deserializer := func(s string) (*testStruct, error) {
		if s == "nilTestStruct" {
			return nil, nil
		}
		var v testStruct
		err := json.Unmarshal([]byte(s), &v)
		return &v, err
	}

	client := NewTypedCacheAsideClient[testStruct](baseClient, serializer, deserializer)

	t.Run("override ttl for negative caching", func(t *testing.T) {
		key := randStr()

		val, err := client.Get(context.Background(), time.Second*5, key, func(ctx context.Context, key string) (*testStruct, error) {
			OverrideCacheTTL(ctx, time.Millisecond*300)
			return nil, nil
		})
		if err != nil || val != nil {
			t.Fatalf("expected nil value, got %v, err: %v", val, err)
		}

		val, err = client.Get(context.Background(), time.Second*5, key, nil)
		if err != nil || val != nil {
			t.Fatalf("expected cached nil value, got %v, err: %v", val, err)
		}

		time.Sleep(time.Millisecond * 400)

		found := &testStruct{ID: 42, Name: "found"}
		val, err = client.Get(context.Background(), time.Second*5, key, func(ctx context.Context, key string) (*testStruct, error) {
			return found, nil
		})
		if err != nil {
			t.Fatal(err)
		}
		if val.ID != found.ID || val.Name != found.Name {
			t.Fatalf("expected %v, got %v", found, val)
		}
	})
}

func TestTypedCacheAsideClientWithCodec_Get(t *testing.T) {
	baseClient := makeClient(t, addr)
	t.Cleanup(baseClient.Close)

	type contextKey string
	const requestContextKey contextKey = "request"
	key := randStr()
	expected := &testStruct{ID: 1, Name: "test"}
	serializerCalls := 0
	deserializerCalls := 0

	codec := TypedCodec[testStruct]{
		Marshal: func(ctx context.Context, gotKey string, val *testStruct) (string, error) {
			serializerCalls++
			if got, want := ctx.Value(requestContextKey), "populate"; got != want {
				t.Errorf("serializer context value = %v, want %v", got, want)
			}
			if gotKey != key {
				t.Errorf("serializer key = %q, want %q", gotKey, key)
			}
			bytes, err := json.Marshal(val)
			return string(bytes), err
		},
		Unmarshal: func(ctx context.Context, gotKey, str string) (*testStruct, error) {
			deserializerCalls++
			wantContext := "populate"
			if deserializerCalls == 2 {
				wantContext = "cached"
			}
			if got := ctx.Value(requestContextKey); got != wantContext {
				t.Errorf("deserializer context value = %v, want %v", got, wantContext)
			}
			if gotKey != key {
				t.Errorf("deserializer key = %q, want %q", gotKey, key)
			}
			var val testStruct
			if err := json.Unmarshal([]byte(str), &val); err != nil {
				return nil, err
			}
			return &val, nil
		},
	}

	client := NewTypedCacheAsideClientWithCodec[testStruct](baseClient, codec)

	val, err := client.Get(context.WithValue(context.Background(), requestContextKey, "populate"), time.Second, key, func(ctx context.Context, gotKey string) (*testStruct, error) {
		if got := ctx.Value(requestContextKey); got != "populate" {
			t.Errorf("fetch context value = %v, want %v", got, "populate")
		}
		if gotKey != key {
			t.Errorf("fetch key = %q, want %q", gotKey, key)
		}
		return expected, nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if *val != *expected {
		t.Fatalf("value = %v, want %v", val, expected)
	}

	cached, err := client.Get(context.WithValue(context.Background(), requestContextKey, "cached"), time.Second, key, func(context.Context, string) (*testStruct, error) {
		t.Fatal("fetch function should not be called for a cached value")
		return nil, nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if *cached != *expected {
		t.Fatalf("cached value = %v, want %v", cached, expected)
	}
	if serializerCalls != 1 {
		t.Errorf("serializer calls = %d, want 1", serializerCalls)
	}
	if deserializerCalls != 2 {
		t.Errorf("deserializer calls = %d, want 2", deserializerCalls)
	}
}
