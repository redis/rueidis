package cmds

import "strings"

var (
	optInTag = uint16(1 << 15)
	blockTag = uint16(1 << 14)
	noRetTag = uint16(1 << 13)
	initSlot = uint16(1 << 15)
	OptInCmd = Completed{
		cs: []string{"CLIENT", "CACHING", "YES"},
		cf: optInTag,
	}
	PingCmd = Completed{
		cs: []string{"PING"},
	}
	QuitCmd = Completed{
		cs: []string{"QUIT"},
	}
)

type Completed struct {
	cs []string
	cf uint16
	ks uint16
}

func (c *Completed) IsEmpty() bool {
	return len(c.cs) == 0
}

func (c *Completed) IsOptIn() bool {
	return c.cf&optInTag == optInTag
}

func (c *Completed) IsBlock() bool {
	return c.cf&blockTag == blockTag
}

func (c *Completed) NoReply() bool {
	return c.cf&noRetTag == noRetTag
}

func (c *Completed) Commands() []string {
	return c.cs
}

type Cacheable Completed
type SCompleted Completed
type SCacheable Completed

func (c *Cacheable) Commands() []string {
	return c.cs
}

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

var multiKeySlotErr = "multi key command with different key slots are not allowed"
