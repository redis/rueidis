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

func TestNewTypedCacheAsideClient(t *testing.T) {
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

	client, err := NewTypedCacheAsideClient[testStruct](baseClient, serializer, deserializer)
	if err != nil {
		t.Fatal(err)
	}

	if client == nil {
		t.Fatal("client should not be nil")
	}
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

	client, _ := NewTypedCacheAsideClient[testStruct](baseClient, serializer, deserializer)

	t.Run("successful get and cache", func(t *testing.T) {
		expected := &testStruct{ID: 1, Name: "test"}
		val, err := client.Get(context.Background(), time.Second, "test-key", func(ctx context.Context, key string) (*testStruct, error) {
			return expected, nil
		})

		if err != nil {
			t.Fatal(err)
		}
		if val != expected {
			t.Fatalf("expected %v, got %v", expected, val)
		}

		// Test cached value
		val2, err := client.Get(context.Background(), time.Second, "test-key", nil)
		if err != nil {
			t.Fatal(err)
		}
		if val2 != expected {
			t.Fatalf("cached value mismatch: expected %v, got %v", expected, val2)
		}
	})

	t.Run("serialization error", func(t *testing.T) {
		badSerializer := func(v *testStruct) (string, error) {
			return "", errors.New("serialization error")
		}
		clientWithBadSerializer, _ := NewTypedCacheAsideClient[testStruct](baseClient, badSerializer, deserializer)

		_, err := clientWithBadSerializer.Get(context.Background(), time.Second, "test-key", func(ctx context.Context, key string) (*testStruct, error) {
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
		clientWithBadDeserializer, _ := NewTypedCacheAsideClient[testStruct](baseClient, serializer, badDeserializer)

		_, err := clientWithBadDeserializer.Get(context.Background(), time.Second, "test-key", func(ctx context.Context, key string) (*testStruct, error) {
			return &testStruct{ID: 1, Name: "test"}, nil
		})

		if err == nil {
			t.Fatal("expected deserialization error")
		}
	})

	t.Run("nil value handling", func(t *testing.T) {
		val, err := client.Get(context.Background(), time.Second, "test-key", func(ctx context.Context, key string) (*testStruct, error) {
			return nil, nil
		})

		if err != nil {
			t.Fatalf("nil valud should not return error: %v", err)
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

	client, _ := NewTypedCacheAsideClient[testStruct](baseClient, serializer, deserializer)

	// Set a value first
	testVal := &testStruct{ID: 1, Name: "test"}
	_, err := client.Get(context.Background(), time.Second, "test-key", func(ctx context.Context, key string) (*testStruct, error) {
		return testVal, nil
	})
	if err != nil {
		t.Fatal(err)
	}

	// Delete the value
	err = client.Del(context.Background(), "test-key")
	if err != nil {
		t.Fatal(err)
	}

	// Verify it's deleted
	_, err = client.Get(context.Background(), time.Second, "test-key", nil)
	if err != nil && !rueidis.IsRedisNil(err) {
		t.Fatal("expected error for deleted key")
	}
}
