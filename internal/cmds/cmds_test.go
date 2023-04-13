package cmds

import (
	"reflect"
	"strings"
	"testing"
)

func TestCacheable_CacheKey(t *testing.T) {
	key, cmd := CacheKey(Cacheable{cs: newCommandSlice([]string{"GET", "A"})})
	if key != "A" || cmd != "GET" {
		t.Fatalf("unexpected ret %v %v", key, cmd)
	}

	key, cmd = CacheKey(Cacheable{cs: newCommandSlice([]string{"HMGET", "A", "B", "C"})})
	if key != "A" || cmd != "HMGETBC" {
		t.Fatalf("unexpected ret %v %v", key, cmd)
	}
}

func TestCacheable_IsMGet(t *testing.T) {
	if cmd := Cacheable(NewMGetCompleted([]string{"MGET", "K"})); !cmd.IsMGet() {
		t.Fatalf("should be mget")
	}
}

func TestCacheable_MGetCacheKey(t *testing.T) {
	if cmd := Cacheable(NewMGetCompleted([]string{"MGET", "K"})); MGetCacheKey(cmd, 0) != "K" {
		t.Fatalf("should be K")
	}
	if cmd := Cacheable(NewMGetCompleted([]string{"JSON.MGET", "K"})); MGetCacheKey(cmd, 0) != "K" {
		t.Fatalf("should be K")
	}
}

func TestCacheable_MGetCacheCmd(t *testing.T) {
	if cmd := Cacheable(NewMGetCompleted([]string{"MGET", "K"})); MGetCacheCmd(cmd) != "GET" {
		t.Fatalf("should be GET")
	}
	if cmd := Cacheable(NewMGetCompleted([]string{"JSON.MGET", "K", "$"})); MGetCacheCmd(cmd) != "JSON.GET$" {
		t.Fatalf("should be JSON.GET$")
	}
}

func TestCompleted_IsEmpty(t *testing.T) {
	if cmd := NewCompleted([]string{}); !cmd.IsEmpty() {
		t.Fatalf("should be empty")
	}
	if cmd := NewCompleted([]string{"a", "b"}); cmd.IsEmpty() {
		t.Fatalf("should not be empty")
	}
}

func TestCompleted_IsBlock(t *testing.T) {
	if cmd := NewCompleted([]string{"a", "b"}); cmd.IsBlock() {
		t.Fatalf("should not be block command")
	}
	if cmd := NewBlockingCompleted([]string{"a", "b"}); !cmd.IsBlock() {
		t.Fatalf("should be block command")
	}
}

func TestCompleted_ToBlock(t *testing.T) {
	cmd := NewCompleted([]string{"a", "b"})
	if cmd.IsBlock() {
		t.Fatalf("should not be block command")
	}
	if cmd = ToBlock(cmd); !cmd.IsBlock() {
		t.Fatalf("should be block command")
	}
}

func TestCompleted_IsOptIn(t *testing.T) {
	if cmd := NewCompleted([]string{"a", "b"}); cmd.IsOptIn() {
		t.Fatalf("should not be opt in command")
	}
	if cmd := OptInCmd; !cmd.IsOptIn() {
		t.Fatalf("should be opt in command")
	}
}

func TestCompleted_NoReply(t *testing.T) {
	if cmd := NewCompleted([]string{"a", "b"}); cmd.NoReply() {
		t.Fatalf("should not be no reply command")
	}
	builder := NewBuilder(InitSlot)
	if cmd := builder.Subscribe().Channel("").Build(); !cmd.NoReply() {
		t.Fatalf("should be no reply command")
	}
	if cmd := builder.Unsubscribe().Channel("").Build(); !cmd.NoReply() {
		t.Fatalf("should be no reply command")
	}
	if cmd := builder.Psubscribe().Pattern("").Build(); !cmd.NoReply() {
		t.Fatalf("should be no reply command")
	}
	if cmd := builder.Punsubscribe().Pattern("").Build(); !cmd.NoReply() {
		t.Fatalf("should be no reply command")
	}
}

func TestComplete_IsReadOnly(t *testing.T) {
	if cmd := NewCompleted([]string{"a", "b"}); cmd.IsReadOnly() {
		t.Fatalf("should not be no readonly command")
	}
	if cmd := NewCompleted([]string{"a", "b"}); !cmd.IsWrite() {
		t.Fatalf("should be write command")
	}
	if cmd := NewReadOnlyCompleted([]string{"a", "b"}); !cmd.IsReadOnly() {
		t.Fatalf("should be readonly command")
	}
	if cmd := NewReadOnlyCompleted([]string{"a", "b"}); cmd.IsWrite() {
		t.Fatalf("should not be write command")
	}
}

func TestNewMultiCompleted(t *testing.T) {
	multi := NewMultiCompleted([][]string{{"a", "b"}, {"c", "d"}})
	if strings.Join(multi[0].Commands(), " ") != "a b" {
		t.Fatalf("unexpected command %v", multi[0].Commands())
	}
	if strings.Join(multi[1].Commands(), " ") != "c d" {
		t.Fatalf("unexpected command %v", multi[1].Commands())
	}
}

func TestCompleted_PanicCrossSlot(t *testing.T) {
	defer func() {
		if !strings.Contains(recover().(string), multiKeySlotErr) {
			t.Fatal("cross key slot not panic as expected")
		}
	}()

	builder := NewBuilder(InitSlot)
	builder.Mget().Key("a").Key("b")
}

func TestCompleted_CommandSlice(t *testing.T) {
	cs := []string{"a", "b", "c"}
	completed := NewCompleted(cs)
	if !reflect.DeepEqual(completed.cs.s, cs) || !reflect.DeepEqual(completed.Commands(), cs) {
		t.Fatalf("unexpecetd diffs")
	}
	cacheable := Cacheable(completed)
	if !reflect.DeepEqual(cacheable.cs.s, cs) || !reflect.DeepEqual(cacheable.Commands(), cs) {
		t.Fatalf("unexpecetd diffs")
	}
}

func TestCompleted_Slots(t *testing.T) {
	builder := NewBuilder(InitSlot)
	c1 := builder.Get().Key("a").Build()
	c2 := builder.Get().Key("b").Build()
	c3 := Cacheable(c1)
	c4 := Cacheable(c2)
	if c1.Slot() == c2.Slot() {
		t.Fatalf("unexpected same slot")
	}
	if c3.Slot() == c4.Slot() {
		t.Fatalf("unexpected same slot")
	}
}

func TestMGets(t *testing.T) {
	keys := []string{"{1}", "{2}", "{3}", "{1}", "{2}", "{3}"}
	ret := MGets(keys)
	for _, key := range keys {
		ks := slot(key)
		cp := ret[slot(key)]
		if cp.ks != ks {
			t.Fatalf("ks mistmatch %v %v", cp.ks, ks)
		}
		if cp.cf != mtGetTag {
			t.Fatalf("cf should be mtGetTag")
		}
		if !reflect.DeepEqual(cp.cs.s, []string{"MGET", key, key}) {
			t.Fatalf("cs mismatch %v %v", cp.cs.s, []string{"MGET", key, key})
		}
	}
}

func TestJsonMGets(t *testing.T) {
	keys := []string{"{1}", "{2}", "{3}", "{1}", "{2}", "{3}"}
	ret := JsonMGets(keys, "$")
	for _, key := range keys {
		ks := slot(key)
		cp := ret[slot(key)]
		if cp.ks != ks {
			t.Fatalf("ks mistmatch %v %v", cp.ks, ks)
		}
		if cp.cf != mtGetTag {
			t.Fatalf("cf should be mtGetTag")
		}
		if !reflect.DeepEqual(cp.cs.s, []string{"JSON.MGET", key, key, "$"}) {
			t.Fatalf("cs mismatch %v %v", cp.cs.s, []string{"JSON.MGET", key, key, "$"})
		}
	}
}

func TestMSets(t *testing.T) {
	keys := map[string]string{"{1}": "{1}", "{2}": "{2}", "{3}": "{3}", "{1}a": "{1}a", "{2}a": "{2}a", "{3}a": "{3}a"}
	ret := MSets(keys)
	for _, key := range keys {
		ks := slot(key)
		cp := ret[slot(key)]
		if cp.ks != ks {
			t.Fatalf("ks mistmatch %v %v", cp.ks, ks)
		}
		if cp.IsReadOnly() {
			t.Fatalf("cf should not be readonly")
		}
		key := strings.TrimSuffix(key, "a")
		key2 := key + "a"
		if !reflect.DeepEqual(cp.cs.s, []string{"MSET", key, key, key2, key2}) &&
			!reflect.DeepEqual(cp.cs.s, []string{"MSET", key2, key2, key, key}) {
			t.Fatalf("cs mismatch %v", cp.cs.s)
		}
	}
}

func TestMSetNXs(t *testing.T) {
	keys := map[string]string{"{1}": "{1}", "{2}": "{2}", "{3}": "{3}", "{1}a": "{1}a", "{2}a": "{2}a", "{3}a": "{3}a"}
	ret := MSetNXs(keys)
	for _, key := range keys {
		ks := slot(key)
		cp := ret[slot(key)]
		if cp.ks != ks {
			t.Fatalf("ks mistmatch %v %v", cp.ks, ks)
		}
		if cp.IsReadOnly() {
			t.Fatalf("cf should not be readonly")
		}
		key := strings.TrimSuffix(key, "a")
		key2 := key + "a"
		if !reflect.DeepEqual(cp.cs.s, []string{"MSETNX", key, key, key2, key2}) &&
			!reflect.DeepEqual(cp.cs.s, []string{"MSETNX", key2, key2, key, key}) {
			t.Fatalf("cs mismatch %v", cp.cs.s)
		}
	}
}

func TestCmdPin(t *testing.T) {
	c1 := NewMGetCompleted([]string{"MGET", "K"})
	if c1.Pin(); c1.cs.Unpin() != 0 {
		t.Fail()
	}
	cc := Cacheable(NewMGetCompleted([]string{"MGET", "K"}))
	if cc.Pin(); cc.cs.Unpin() != 0 {
		t.Fail()
	}
}

func TestCompletedCommandSlice(t *testing.T) {
	if c1 := NewMGetCompleted([]string{"MGET", "K"}); CompletedCommandSlice(c1) != c1.cs {
		t.Fail()
	}
}

func TestCacheableCommandSlice(t *testing.T) {
	if c1 := Cacheable(NewMGetCompleted([]string{"MGET", "K"})); CacheableCommandSlice(c1) != c1.cs {
		t.Fail()
	}
}
