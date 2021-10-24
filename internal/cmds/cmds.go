package cmds

import "strings"

var (
	optInTag = uint32(1 << 31)
	OptInCmd = Completed{
		cs: []string{"CLIENT", "CACHING", "YES"},
		cf: optInTag,
	}
)

type Completed struct {
	cs []string
	cf uint32
}

func (c *Completed) IsEmpty() bool {
	return len(c.cs) == 0
}

func (c *Completed) IsOptIn() bool {
	return c.cf&optInTag == optInTag
}

func (c *Completed) Commands() []string {
	return c.cs
}

type Cacheable Completed

func (c *Cacheable) CacheKey() (key, command string) {
	if len(c.cs) == 2 {
		return c.cs[1], c.cs[0]
	}

	length := 0
	for i, v := range c.cs {
		if i == 1 {
			continue
		}
		length += len(v)
	}
	sb := strings.Builder{}
	sb.Grow(length)
	for i, v := range c.cs {
		if i == 1 {
			key = v
		} else {
			sb.WriteString(v)
		}
	}
	return key, sb.String()
}

func NewCompleted(cs []string) Completed {
	return Completed{cs: cs}
}

func NewMultiCompleted(cs [][]string) []Completed {
	ret := make([]Completed, len(cs))
	for i, c := range cs {
		ret[i] = NewCompleted(c)
	}
	return ret
}
