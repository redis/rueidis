package mock

import (
	"fmt"
	"strings"

	"github.com/redis/rueidis"
	"go.uber.org/mock/gomock"
)

func Match(cmd ...string) gomock.Matcher {
	return gomock.GotFormatterAdapter(
		gomock.GotFormatterFunc(func(i any) string {
			return format(i)
		}),
		&cmdMatcher{expect: cmd},
	)
}

type cmdMatcher struct {
	expect []string
}

func (c *cmdMatcher) Matches(x any) bool {
	return gomock.Eq(commands(x)).Matches(c.expect)
}

func (c *cmdMatcher) String() string {
	return fmt.Sprintf("redis command %v", c.expect)
}

func MatchFn(fn func(cmd []string) bool, description ...string) gomock.Matcher {
	return gomock.GotFormatterAdapter(
		gomock.GotFormatterFunc(func(i any) string {
			return format(i)
		}),
		&fnMatcher{matcher: fn, description: description},
	)
}

type fnMatcher struct {
	matcher     func(cmd []string) bool
	description []string
}

func (c *fnMatcher) Matches(x any) bool {
	if cmd, ok := commands(x).([]string); ok {
		return c.matcher(cmd)
	}
	return false
}

func (c *fnMatcher) String() string {
	return fmt.Sprintf("matches %v", strings.Join(c.description, " "))
}

func format(v any) string {
	if _, ok := v.([]any); !ok {
		v = []any{v}
	}
	sb := &strings.Builder{}
	sb.WriteString("\n")
	for i, c := range v.([]any) {
		fmt.Fprintf(sb, "index %d redis command %v\n", i+1, commands(c))
	}
	return sb.String()
}

func commands(x any) any {
	if cmd, ok := x.(rueidis.Completed); ok {
		return cmd.Commands()
	}
	if cmd, ok := x.(rueidis.Cacheable); ok {
		return cmd.Commands()
	}
	if cmd, ok := x.(rueidis.CacheableTTL); ok {
		return cmd.Cmd.Commands()
	}
	return x
}
