package mock

import (
	"reflect"
	"strings"
	"testing"

	"github.com/redis/rueidis"
	"github.com/redis/rueidis/internal/cmds"
	"go.uber.org/mock/gomock"
)

func TestMatch_Completed(t *testing.T) {
	cmd := cmds.NewBuilder(cmds.NoSlot).Get().Key("k").Build()
	if m := Match("GET", "k"); !m.Matches(cmd) {
		t.Fatalf("not matched %s", m.String())
	}
}

func TestMatchFn_Completed(t *testing.T) {
	cmd := cmds.NewBuilder(cmds.NoSlot).Get().Key("k").Build()
	if m := MatchFn(func(cmd []string) bool {
		return reflect.DeepEqual(cmd, []string{"GET", "k"})
	}, "GET k"); !m.Matches(cmd) {
		t.Fatalf("not matched %s", m.String())
	}
}

func TestMatch_Cacheable(t *testing.T) {
	cmd := cmds.NewBuilder(cmds.NoSlot).Get().Key("k").Cache()
	if m := Match("GET", "k"); !m.Matches(cmd) {
		t.Fatalf("not matched %s", m.String())
	}
}

func TestMatchFn_Cacheable(t *testing.T) {
	cmd := cmds.NewBuilder(cmds.NoSlot).Get().Key("k").Cache()
	if m := MatchFn(func(cmd []string) bool {
		return reflect.DeepEqual(cmd, []string{"GET", "k"})
	}, "GET k"); !m.Matches(cmd) {
		t.Fatalf("not matched %s", m.String())
	}
}

func TestMatch_CacheableTTL(t *testing.T) {
	cmd := cmds.NewBuilder(cmds.NoSlot).Get().Key("k").Cache()
	if m := Match("GET", "k"); !m.Matches(rueidis.CacheableTTL{Cmd: cmd}) {
		t.Fatalf("not matched %s", m.String())
	}
}

func TestMatchFn_CacheableTTL(t *testing.T) {
	cmd := cmds.NewBuilder(cmds.NoSlot).Get().Key("k").Cache()
	if m := MatchFn(func(cmd []string) bool {
		return reflect.DeepEqual(cmd, []string{"GET", "k"})
	}, "GET k"); !m.Matches(cmd) {
		t.Fatalf("not matched %s", m.String())
	}
}

func TestMatch_Other(t *testing.T) {
	if m := Match("GET", "k"); m.Matches(1) {
		t.Fatalf("unexpected matched %s", m.String())
	}
	if m := Match("GET", "k"); m.Matches([]rueidis.Completed{
		cmds.NewBuilder(cmds.NoSlot).Get().Key("k").Build(), // https://github.com/redis/rueidis/issues/120
	}) {
		t.Fatalf("unexpected matched %s", m.String())
	}
}

func TestMatchFn_Other(t *testing.T) {
	if m := MatchFn(func(cmd []string) bool {
		return reflect.DeepEqual(cmd, []string{"GET", "k"})
	}, "GET k"); m.Matches(1) {
		t.Fatalf("unexpected matched %s", m.String())
	}
	if m := MatchFn(func(cmd []string) bool {
		return reflect.DeepEqual(cmd, []string{"GET", "k"})
	}, "GET k"); m.Matches([]rueidis.Completed{
		cmds.NewBuilder(cmds.NoSlot).Get().Key("k").Build(), // https://github.com/redis/rueidis/issues/120
	}) {
		t.Fatalf("unexpected matched %s", m.String())
	}
}

func TestMatch_Format(t *testing.T) {
	matcher := Match("GET", "t")
	if !strings.Contains(matcher.String(), "GET t") {
		t.Fatalf("unexpected format %v", matcher.String())
	}
	if !strings.Contains(matcher.(gomock.GotFormatter).Got(cmds.NewBuilder(cmds.NoSlot).Get().Key("k").Build()), "GET k") {
		t.Fatalf("unexpected format %v", matcher.String())
	}
}

func TestMatchFn_Format(t *testing.T) {
	matcher := MatchFn(func(cmd []string) bool {
		return reflect.DeepEqual(cmd, []string{"GET", "t"})
	}, "GET t")
	if !strings.Contains(matcher.String(), "GET t") {
		t.Fatalf("unexpected format %v", matcher.String())
	}
	if !strings.Contains(matcher.(gomock.GotFormatter).Got(cmds.NewBuilder(cmds.NoSlot).Get().Key("k").Build()), "GET k") {
		t.Fatalf("unexpected format %v", matcher.String())
	}
}
