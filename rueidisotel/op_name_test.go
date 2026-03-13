package rueidisotel

import (
	"testing"
	"time"

	"github.com/redis/rueidis"
	"github.com/redis/rueidis/internal/cmds"
)

func TestDefaultOpNameResolver_OpName(t *testing.T) {
	resolver := DefaultOpNameResolver{}
	cmd := cmds.NewBuilder(cmds.NoSlot).Get().Key("one").Build()

	got := resolver.OpName(t.Context(), cmd)
	if got != "GET" {
		t.Errorf("OpName() = %v, want %v", got, "GET")
	}
}

func TestDefaultOpNameResolver_MultiOpName(t *testing.T) {
	builder := cmds.NewBuilder(cmds.NoSlot)

	tests := [...]struct {
		name  string
		limit int
		cmds  rueidis.Commands
		want  string
	}{
		{
			name: "empty commands",
		},
		{
			name: "single command",
			cmds: rueidis.Commands{
				builder.Get().Key("one").Build(),
			},
			want: "GET",
		},
		{
			name: "many commands without limit",
			cmds: rueidis.Commands{
				builder.Get().Key("one").Build(),
				builder.Mget().Key("two").Build(),
				builder.Strlen().Key("three").Build(),
			},
			want: "GET MGET STRLEN",
		},
		{
			name:  "many commands with limit",
			limit: 2,
			cmds: rueidis.Commands{
				builder.Get().Key("one").Build(),
				builder.Mget().Key("two").Build(),
				builder.Strlen().Key("three").Build(),
			},
			want: "GET MGET",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resolver := DefaultOpNameResolver{Limit: tt.limit}
			got := resolver.MultiOpName(t.Context(), tt.cmds)
			if got != tt.want {
				t.Errorf("MultiOpName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultOpNameResolver_MultiCacheableOpName(t *testing.T) {
	builder := cmds.NewBuilder(cmds.NoSlot)

	tests := [...]struct {
		name  string
		limit int
		cmds  []rueidis.CacheableTTL
		want  string
	}{
		{
			name: "empty commands",
		},
		{
			name: "single command",
			cmds: []rueidis.CacheableTTL{
				rueidis.CT(builder.Get().Key("one").Cache(), time.Minute),
			},
			want: "GET",
		},
		{
			name: "many commands without limit",
			cmds: []rueidis.CacheableTTL{
				rueidis.CT(builder.Get().Key("one").Cache(), time.Minute),
				rueidis.CT(builder.Mget().Key("two").Cache(), time.Minute),
				rueidis.CT(builder.Strlen().Key("three").Cache(), time.Minute),
			},
			want: "GET MGET STRLEN",
		},
		{
			name:  "many commands with limit",
			limit: 2,
			cmds: []rueidis.CacheableTTL{
				rueidis.CT(builder.Get().Key("one").Cache(), time.Minute),
				rueidis.CT(builder.Mget().Key("two").Cache(), time.Minute),
				rueidis.CT(builder.Strlen().Key("three").Cache(), time.Minute),
			},
			want: "GET MGET",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resolver := DefaultOpNameResolver{Limit: tt.limit}
			got := resolver.MultiCacheableOpName(t.Context(), tt.cmds)
			if got != tt.want {
				t.Errorf("MultiOpName() = %v, want %v", got, tt.want)
			}
		})
	}
}
