package cmds

import "testing"

func TestCacheable_CacheKey(t *testing.T) {
	key, cmd := (&Cacheable{cs: &CommandSlice{s: []string{"GET", "A"}}}).CacheKey()
	if key != "A" || cmd != "GET" {
		t.Fatalf("unexpected ret %v %v", key, cmd)
	}

	key, cmd = (&Cacheable{cs: &CommandSlice{s: []string{"HMGET", "A", "B", "C"}}}).CacheKey()
	if key != "A" || cmd != "HMGETBC" {
		t.Fatalf("unexpected ret %v %v", key, cmd)
	}
}
