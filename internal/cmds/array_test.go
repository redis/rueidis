package cmds

import (
	"reflect"
	"testing"
)

func TestArgrepBuildsMultiplePredicates(t *testing.T) {
	cmd := NewBuilder(NoSlot).Argrep().Key("array").
		Start("-").End("+").
		Exact().String("foo").
		And().Match().String("bar").
		Limit(10).Withvalues().Nocase().
		Build()

	want := []string{
		"ARGREP", "array", "-", "+",
		"EXACT", "foo",
		"AND", "MATCH", "bar",
		"LIMIT", "10", "WITHVALUES", "NOCASE",
	}
	if got := cmd.Commands(); !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected command: got %v want %v", got, want)
	}
	if !cmd.IsReadOnly() {
		t.Fatalf("ARGREP should be readonly")
	}
}
