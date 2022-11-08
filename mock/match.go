package mock

import (
	"fmt"
	"strings"

	"github.com/golang/mock/gomock"
	"github.com/rueian/rueidis"
	"github.com/rueian/rueidis/internal/cmds"
)

func Match(cmd ...string) gomock.Matcher {
	return gomock.GotFormatterAdapter(
		gomock.GotFormatterFunc(func(i interface{}) string {
			return format(i)
		}),
		&cmdMatcher{expect: cmd},
	)
}

type cmdMatcher struct {
	expect []string
}

func (c *cmdMatcher) Matches(x interface{}) bool {
	return gomock.Eq(commands(x)).Matches(c.expect)
}

func (c *cmdMatcher) String() string {
	return fmt.Sprintf("redis command %v", c.expect)
}

func format(v interface{}) string {
	if _, ok := v.([]interface{}); !ok {
		v = []interface{}{v}
	}
	sb := &strings.Builder{}
	sb.WriteString("\n")
	for i, c := range v.([]interface{}) {
		fmt.Fprintf(sb, "index %d redis command %v\n", i+1, commands(c))
	}
	return sb.String()
}

func commands(x interface{}) interface{} {
	if cmd, ok := x.(cmds.Completed); ok {
		return cmd.Commands()
	}
	if cmd, ok := x.(cmds.Cacheable); ok {
		return cmd.Commands()
	}
	if cmd, ok := x.(rueidis.CacheableTTL); ok {
		return cmd.Cmd.Commands()
	}
	return x
}
