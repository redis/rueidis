package cmds

import "strings"

const (
	optInTag = uint16(1 << 15)
	blockTag = uint16(1 << 14)
	noRetTag = uint16(1 << 13)
	readonly = uint16(1 << 12)
	InitSlot = uint16(1 << 15)
)

var (
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
	SlotCmd = Completed{
		cs: []string{"CLUSTER", "SLOTS"},
	}
	AskingCmd = Completed{
		cs: []string{"ASKING"},
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

func (c *Completed) IsReadOnly() bool {
	return c.cf&readonly == readonly
}

func (c *Completed) IsWrite() bool {
	return !c.IsReadOnly()
}

func (c *Completed) Commands() []string {
	return c.cs
}

type Cacheable Completed
type SCompleted Completed

func (c *SCompleted) Commands() []string {
	return c.cs
}

func (c *SCompleted) Slot() uint16 {
	return c.ks
}

type SCacheable Completed

func (c *SCacheable) Slot() uint16 {
	return c.ks
}

func (c *SCacheable) Commands() []string {
	return c.cs
}

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

func NewBlockingCompleted(cs []string) Completed {
	return Completed{cs: cs, cf: blockTag}
}

func NewReadOnlyCompleted(cs []string) Completed {
	return Completed{cs: cs, cf: readonly}
}

func NewMultiCompleted(cs [][]string) []Completed {
	ret := make([]Completed, len(cs))
	for i, c := range cs {
		ret[i] = NewCompleted(c)
	}
	return ret
}

func checkSlot(prev, new uint16) uint16 {
	if prev == InitSlot || prev == new {
		return new
	}
	panic(multiKeySlotErr)
}

const multiKeySlotErr = "multi key command with different key slots are not allowed"
