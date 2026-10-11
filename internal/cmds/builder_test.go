package cmds

import (
	"reflect"
	"testing"
)

func TestPutCompleted(t *testing.T) {
retry:
	cs1 := get()
	cs1.s = append(cs1.s, "1", "1", "1", "1", "1")
	PutCompleted(Completed{cs: cs1})
	cs2 := get()
	if cs1 != cs2 {
		goto retry
	}
	if len(cs2.s) != 0 {
		t.Fatalf("Put doesn't clean the CommandSlice")
	}
}

func TestPutCompletedForce(t *testing.T) {
retry:
	cs1 := get()
	cs1.s = append(cs1.s, "1", "1", "1", "1", "1")
	cs1.r = 1 // pin
	PutCompletedForce(Completed{cs: cs1})
	cs2 := get()
	if cs1 != cs2 {
		goto retry
	}
	if len(cs2.s) != 0 {
		t.Fatalf("PutCompletedForce doesn't clean the CommandSlice")
	}
}

func TestPutCacheableForce(t *testing.T) {
retry:
	cs1 := get()
	cs1.s = append(cs1.s, "1", "1", "1", "1", "1")
	cs1.r = 1 // pin
	PutCacheableForce(Cacheable{cs: cs1})
	cs2 := get()
	if cs1 != cs2 {
		goto retry
	}
	if len(cs2.s) != 0 {
		t.Fatalf("PutCacheableForce doesn't clean the CommandSlice")
	}
}

func TestPutCacheable(t *testing.T) {
retry:
	cs1 := get()
	cs1.s = append(cs1.s, "1", "1", "1", "1", "1")
	PutCacheable(Cacheable{cs: cs1})
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
	if c := builder.Arbitrary("any").Blocking(); !c.IsBlock() {
		t.Fatalf("arbitrary failed")
	}
	if c := builder.Arbitrary("any").ReadOnly(); !c.IsReadOnly() {
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

func TestBuiltTwice(t *testing.T) {
	src := NewBuilder(NoSlot).Get()
	cmd1 := src.Key("a")
	cmd2 := src.Key("b")
	cmd1.Build()
	defer func() {
		if e := recover(); e != ErrBuiltTwice {
			t.Errorf("arbitrary not check MGET command")
		}
	}()
	cmd2.Build()
}

func TestVerify(t *testing.T) {
	src := NewBuilder(NoSlot).Get()
	cmd1 := src.Key("a").Build()
	cmd1.cs.Verify()
	src.Key("b")
	defer func() {
		if e := recover(); e != ErrUnfinished {
			t.Errorf("arbitrary not check MGET command")
		}
	}()
	cmd1.cs.Verify()
}

func TestZrangeStartStopMinMax(t *testing.T) {
	b := NewBuilder(NoSlot)

	testCases := []struct {
		name     string
		startCmd Completed
		minCmd   Completed
		expected []string
	}{
		{
			name:     "plain range",
			startCmd: b.Zrange().Key("k").Start("1").Stop("2").Build(),
			minCmd:   b.Zrange().Key("k").Min("1").Max("2").Build(),
			expected: []string{"ZRANGE", "k", "1", "2"},
		},
		{
			name:     "byscore rev",
			startCmd: b.Zrange().Key("k").Start("+inf").Stop("1251").Byscore().Rev().Build(),
			minCmd:   b.Zrange().Key("k").Min("+inf").Max("1251").Byscore().Rev().Build(),
			expected: []string{"ZRANGE", "k", "+inf", "1251", "BYSCORE", "REV"},
		},
		{
			name:     "bylex limit",
			startCmd: b.Zrange().Key("k").Start("[a").Stop("[z").Bylex().Limit(0, 10).Build(),
			minCmd:   b.Zrange().Key("k").Min("[a").Max("[z").Bylex().Limit(0, 10).Build(),
			expected: []string{"ZRANGE", "k", "[a", "[z", "BYLEX", "LIMIT", "0", "10"},
		},
		{
			name:     "withscores",
			startCmd: b.Zrange().Key("k").Start("0").Stop("-1").Withscores().Build(),
			minCmd:   b.Zrange().Key("k").Min("0").Max("-1").Withscores().Build(),
			expected: []string{"ZRANGE", "k", "0", "-1", "WITHSCORES"},
		},
		{
			name:     "byscore rev limit withscores",
			startCmd: b.Zrange().Key("k").Start("+inf").Stop("0").Byscore().Rev().Limit(1, 5).Withscores().Build(),
			minCmd:   b.Zrange().Key("k").Min("+inf").Max("0").Byscore().Rev().Limit(1, 5).Withscores().Build(),
			expected: []string{"ZRANGE", "k", "+inf", "0", "BYSCORE", "REV", "LIMIT", "1", "5", "WITHSCORES"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if !reflect.DeepEqual(tc.startCmd.Commands(), tc.expected) {
				t.Fatalf("Start/Stop produced %v, expected %v", tc.startCmd.Commands(), tc.expected)
			}
			if !reflect.DeepEqual(tc.minCmd.Commands(), tc.expected) {
				t.Fatalf("Min/Max produced %v, expected %v", tc.minCmd.Commands(), tc.expected)
			}
			if !reflect.DeepEqual(tc.startCmd.Commands(), tc.minCmd.Commands()) {
				t.Fatalf("Start/Stop %v != Min/Max %v", tc.startCmd.Commands(), tc.minCmd.Commands())
			}
		})
	}

	// Cacheable test
	startCache := b.Zrange().Key("k").Start("1").Stop("2").Cache()
	minCache := b.Zrange().Key("k").Min("1").Max("2").Cache()
	if !reflect.DeepEqual(startCache.Commands(), minCache.Commands()) {
		t.Fatalf("Start/Stop Cache %v != Min/Max Cache %v", startCache.Commands(), minCache.Commands())
	}
}

func TestZrangestoreStartStopMinMax(t *testing.T) {
	b := NewBuilder(NoSlot)

	testCases := []struct {
		name     string
		startCmd Completed
		minCmd   Completed
		expected []string
	}{
		{
			name:     "plain store",
			startCmd: b.Zrangestore().Dst("dst").Src("src").Start("1").Stop("2").Build(),
			minCmd:   b.Zrangestore().Dst("dst").Src("src").Min("1").Max("2").Build(),
			expected: []string{"ZRANGESTORE", "dst", "src", "1", "2"},
		},
		{
			name:     "byscore rev limit",
			startCmd: b.Zrangestore().Dst("dst").Src("src").Start("+inf").Stop("1251").Byscore().Rev().Limit(0, 10).Build(),
			minCmd:   b.Zrangestore().Dst("dst").Src("src").Min("+inf").Max("1251").Byscore().Rev().Limit(0, 10).Build(),
			expected: []string{"ZRANGESTORE", "dst", "src", "+inf", "1251", "BYSCORE", "REV", "LIMIT", "0", "10"},
		},
		{
			name:     "bylex rev",
			startCmd: b.Zrangestore().Dst("dst").Src("src").Start("[z").Stop("[a").Bylex().Rev().Build(),
			minCmd:   b.Zrangestore().Dst("dst").Src("src").Min("[z").Max("[a").Bylex().Rev().Build(),
			expected: []string{"ZRANGESTORE", "dst", "src", "[z", "[a", "BYLEX", "REV"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if !reflect.DeepEqual(tc.startCmd.Commands(), tc.expected) {
				t.Fatalf("Start/Stop produced %v, expected %v", tc.startCmd.Commands(), tc.expected)
			}
			if !reflect.DeepEqual(tc.minCmd.Commands(), tc.expected) {
				t.Fatalf("Min/Max produced %v, expected %v", tc.minCmd.Commands(), tc.expected)
			}
			if !reflect.DeepEqual(tc.startCmd.Commands(), tc.minCmd.Commands()) {
				t.Fatalf("Start/Stop %v != Min/Max %v", tc.startCmd.Commands(), tc.minCmd.Commands())
			}
		})
	}
}

