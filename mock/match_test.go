package mock

import (
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/rueian/rueidis"
	"github.com/rueian/rueidis/internal/cmds"
)

func TestMatch_Completed(t *testing.T) {
	cmd := cmds.NewBuilder(cmds.NoSlot).Get().Key("k").Build()
	if m := Match("GET", "k"); !m.Matches(cmd) {
		t.Fatalf("not matched %s", m.String())
	}
}

func TestMatch_Cacheable(t *testing.T) {
	cmd := cmds.NewBuilder(cmds.NoSlot).Get().Key("k").Cache()
	if m := Match("GET", "k"); !m.Matches(cmd) {
		t.Fatalf("not matched %s", m.String())
	}
}

func TestMatch_CacheableTTL(t *testing.T) {
	cmd := cmds.NewBuilder(cmds.NoSlot).Get().Key("k").Cache()
	if m := Match("GET", "k"); !m.Matches(rueidis.CacheableTTL{Cmd: cmd}) {
		t.Fatalf("not matched %s", m.String())
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
