package cmds

import (
	"reflect"
	"testing"
)

func TestPut(t *testing.T) {
retry:
	cs1 := get()
	cs1.s = append(cs1.s, "1", "1", "1", "1", "1")
	Put(cs1)
	cs2 := get()
	if cs1 != cs2 {
		goto retry
	}
	if len(cs2.s) != 0 {
		t.Fatalf("Put doesn't clean the CommandSlice")
	}
}

func TestArbitraryIsZero(t *testing.T) {
	builder := NewBuilder(NoSlot)
	if cmd := builder.Arbitrary("any", "cmd"); cmd.IsZero() {
		t.Fatalf("arbitrary failed")
	}
	var cmd Arbitrary
	if !cmd.IsZero() {
		t.Fatalf("arbitrary failed")
	}
}

func TestArbitrary(t *testing.T) {
	builder := NewBuilder(NoSlot)
	cmd := builder.Arbitrary("any", "cmd").Keys("k1", "k2").Args("a1", "a2")
	if c := cmd.Build(); !reflect.DeepEqual(c.Commands(), []string{"any", "cmd", "k1", "k2", "a1", "a2"}) {
		t.Fatalf("arbitrary failed")
	}
	if c := cmd.Blocking(); !c.IsBlock() {
		t.Fatalf("arbitrary failed")
	}
	if c := cmd.ReadOnly(); !c.IsReadOnly() {
		t.Fatalf("arbitrary failed")
	}

	builder2 := NewBuilder(InitSlot)

	defer func() {
		if e := recover(); e != multiKeySlotErr {
			t.Errorf("arbitrary not check slots")
		}
	}()

	builder2.Arbitrary().Keys("k1", "k2")
}

func TestEmptyArbitrary(t *testing.T) {
	builder := NewBuilder(NoSlot)
	defer func() {
		if e := recover(); e != arbitraryNoCommand {
			t.Errorf("arbitrary not check empty")
		}
	}()
	builder.Arbitrary().Build()
}

func TestEmptySubscribe(t *testing.T) {
	builder := NewBuilder(NoSlot)
	defer func() {
		if e := recover(); e != arbitrarySubscribe {
			t.Errorf("arbitrary not check subscribe command")
		}
	}()
	builder.Arbitrary("SUBSCRIBE").Build()
}

func TestEmptyArbitraryMultiGet(t *testing.T) {
	builder := NewBuilder(NoSlot)
	defer func() {
		if e := recover(); e != arbitraryNoCommand {
			t.Errorf("arbitrary not check empty")
		}
	}()
	builder.Arbitrary().MultiGet()
}

func TestArbitraryMultiGet(t *testing.T) {
	builder := NewBuilder(NoSlot)
	cacheable := Cacheable(builder.Arbitrary("MGET").Args("KKK").MultiGet())
	if !cacheable.IsMGet() {
		t.Fatalf("arbitrary failed")
	}
}

func TestArbitraryMultiGetPanic(t *testing.T) {
	builder := NewBuilder(NoSlot)
	defer func() {
		if e := recover(); e != arbitraryMultiGet {
			t.Errorf("arbitrary not check MGET command")
		}
	}()
	builder.Arbitrary("SUBSCRIBE").MultiGet()
}
