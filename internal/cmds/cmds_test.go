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

func TestCacheable_Scripting(t *testing.T) {
	key, cmd := CacheKey(Cacheable{cs: newCommandSlice([]string{"EVALSHA_RO", "sha1", "1", "OOO", "XXX"}), cf: scrRoTag})
	if key != "OOO" || cmd != "EVALSHA_ROsha11XXX" {
		t.Fatalf("unexpected ret %v %v", key, cmd)
	}
	defer func() {
		if err := recover().(string); err != multiKeyCacheErr {
			t.Fatalf("not panic as expected")
		}
	}()
	CacheKey(Cacheable{cs: newCommandSlice([]string{"EVALSHA_RO", "sha1", "2", "OOO", "XXX"}), cf: scrRoTag})
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
	if ToBlock(&cmd); !cmd.IsBlock() {
		t.Fatalf("should be block command")
	}
}

func TestCompleted_IsOptIn(t *testing.T) {
	if cmd := NewCompleted([]string{"a", "b"}); cmd.IsOptIn() {
		t.Fatalf("should not be opt-in command")
	}
	if cmd := OptInCmd; !cmd.IsOptIn() {
		t.Fatalf("should be opt-in command")
	}
	if cmd := OptInNopCmd; !cmd.IsOptIn() {
		t.Fatalf("should be opt-in command")
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
	if cmd := builder.Ssubscribe().Channel("").Build(); !cmd.NoReply() {
		t.Fatalf("should be no reply command")
	}
	if cmd := builder.Sunsubscribe().Channel("").Build(); !cmd.NoReply() {
		t.Fatalf("should be no reply command")
	}
}

func TestCompleted_IsUnsub(t *testing.T) {
	if cmd := NewCompleted([]string{"a", "b"}); cmd.IsUnsub() {
		t.Fatalf("should not be no reply command")
	}
	builder := NewBuilder(InitSlot)
	if cmd := builder.Subscribe().Channel("").Build(); cmd.IsUnsub() {
		t.Fatalf("should be not be unsub command")
	}
	if cmd := builder.Unsubscribe().Channel("").Build(); !cmd.IsUnsub() {
		t.Fatalf("should be unsub command")
	}
	if cmd := builder.Psubscribe().Pattern("").Build(); cmd.IsUnsub() {
		t.Fatalf("should be not be unsub command")
	}
	if cmd := builder.Punsubscribe().Pattern("").Build(); !cmd.IsUnsub() {
		t.Fatalf("should be unsub command")
	}
	if cmd := builder.Ssubscribe().Channel("").Build(); cmd.IsUnsub() {
		t.Fatalf("should be not be unsub command")
	}
	if cmd := builder.Sunsubscribe().Channel("").Build(); !cmd.IsUnsub() {
		t.Fatalf("should be unsub command")
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
		t.Fatalf("unexpected diffs")
	}
	cacheable := Cacheable(completed)
	if !reflect.DeepEqual(cacheable.cs.s, cs) || !reflect.DeepEqual(cacheable.Commands(), cs) {
		t.Fatalf("unexpected diffs")
	}
}

func TestCompleted_Slots(t *testing.T) {
	for _, init := range []uint16{InitSlot, NoSlot} {
		builder := NewBuilder(init)
		c1 := builder.Get().Key("a").Build()
		c2 := builder.Get().Key("b").Build()
		c3 := Cacheable(c1)
		c4 := Cacheable(c2)
		c5 := c1.SetSlot("c")
		c6 := c2.SetSlot("c")
		if c1.Slot() == c2.Slot() {
			t.Fatalf("unexpected same slot")
		}
		if c3.Slot() == c4.Slot() {
			t.Fatalf("unexpected same slot")
		}
		if c5.Slot() != c6.Slot() {
			t.Fatalf("unexpected different slot")
		}
	}
}

func TestCompleted_Append(t *testing.T) {
	builder := NewBuilder(InitSlot)
	c := builder.Get().Key("a").Build()
	AppendCompleted(c, "b")
	if c.Commands()[len(c.Commands())-1] != "b" {
		t.Fatalf("unexpected command %v", c.Commands())
	}
	CompletedCS(c).Verify()
}

func TestMGets(t *testing.T) {
	keys := []string{"{1}", "{2}", "{3}", "{1}", "{2}", "{3}"}
	ret := MGets(keys)
	for _, key := range keys {
		ks := slot(key)
		cp := ret[slot(key)]
		cp.cs.Verify()
		if cp.ks != ks {
			t.Fatalf("ks mismatch %v %v", cp.ks, ks)
		}
		if cp.cf != mtGetTag {
			t.Fatalf("cf should be mtGetTag")
		}
		if !reflect.DeepEqual(cp.cs.s, []string{"MGET", key, key}) {
			t.Fatalf("cs mismatch %v %v", cp.cs.s, []string{"MGET", key, key})
		}
	}
}

func TestMDels(t *testing.T) {
	keys := []string{"{1}", "{2}", "{3}", "{1}", "{2}", "{3}"}
	ret := MDels(keys)
	for _, key := range keys {
		ks := slot(key)
		cp := ret[slot(key)]
		cp.cs.Verify()
		if cp.ks != ks {
			t.Fatalf("ks mismatch %v %v", cp.ks, ks)
		}
		if cp.cf == mtGetTag {
			t.Fatalf("cf should not be mtGetTag")
		}
		if !reflect.DeepEqual(cp.cs.s, []string{"DEL", key, key}) {
			t.Fatalf("cs mismatch %v %v", cp.cs.s, []string{"DEL", key, key})
		}
	}
}

func TestJsonMGets(t *testing.T) {
	keys := []string{"{1}", "{2}", "{3}", "{1}", "{2}", "{3}"}
	ret := JsonMGets(keys, "$")
	for _, key := range keys {
		ks := slot(key)
		cp := ret[slot(key)]
		cp.cs.Verify()
		if cp.ks != ks {
			t.Fatalf("ks mismatch %v %v", cp.ks, ks)
		}
		if cp.cf != mtGetTag {
			t.Fatalf("cf should be mtGetTag")
		}
		if !reflect.DeepEqual(cp.cs.s, []string{"JSON.MGET", key, key, "$"}) {
			t.Fatalf("cs mismatch %v %v", cp.cs.s, []string{"JSON.MGET", key, key, "$"})
		}
	}
}

func TestJsonMSets(t *testing.T) {
	keys := map[string]string{"{1}": "{1}", "{2}": "{2}", "{3}": "{3}", "{1}a": "{1}a", "{2}a": "{2}a", "{3}a": "{3}a"}
	ret := JsonMSets(keys, "$")
	for _, key := range keys {
		ks := slot(key)
		cp := ret[slot(key)]
		cp.cs.Verify()
		if cp.ks != ks {
			t.Fatalf("ks mismatch %v %v", cp.ks, ks)
		}
		if cp.IsReadOnly() {
			t.Fatalf("cf should not be readonly")
		}
		key := strings.TrimSuffix(key, "a")
		key2 := key + "a"
		if !reflect.DeepEqual(cp.cs.s, []string{"JSON.MSET", key, "$", key, key2, "$", key2}) &&
			!reflect.DeepEqual(cp.cs.s, []string{"JSON.MSET", key2, "$", key2, key, "$", key}) {
			t.Fatalf("cs mismatch %v", cp.cs.s)
		}
	}
}

func TestMSets(t *testing.T) {
	keys := map[string]string{"{1}": "{1}", "{2}": "{2}", "{3}": "{3}", "{1}a": "{1}a", "{2}a": "{2}a", "{3}a": "{3}a"}
	ret := MSets(keys)
	for _, key := range keys {
		ks := slot(key)
		cp := ret[slot(key)]
		cp.cs.Verify()
		if cp.ks != ks {
			t.Fatalf("ks mismatch %v %v", cp.ks, ks)
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
		cp.cs.Verify()
		if cp.ks != ks {
			t.Fatalf("ks mismatch %v %v", cp.ks, ks)
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
	if c1.Pin(); c1.cs.r == 0 {
		t.Fail()
	}
	cc := Cacheable(NewMGetCompleted([]string{"MGET", "K"}))
	if cc.Pin(); c1.cs.r == 0 {
		t.Fail()
	}
}

func TestCompletedCS(t *testing.T) {
	if c1 := NewMGetCompleted([]string{"MGET", "K"}); CompletedCS(c1) != c1.cs {
		t.Fail()
	}
}

func TestCacheableCS(t *testing.T) {
	if c1 := Cacheable(NewMGetCompleted([]string{"MGET", "K"})); CacheableCS(c1) != c1.cs {
		t.Fail()
	}
}
