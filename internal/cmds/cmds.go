package cmds

import "strings"

const (
	optInTag = uint16(1 << 15)
	blockTag = uint16(1 << 14)
	readonly = uint16(1 << 13)
	noRetTag = uint16(1<<12) | readonly // make noRetTag can also be retried
	InitSlot = uint16(1 << 15)
	NoSlot   = InitSlot + 1
)

var (
	OptInCmd = Completed{
		cs: &CommandSlice{s: []string{"CLIENT", "CACHING", "YES"}},
		cf: optInTag,
	}
	QuitCmd = Completed{
		cs: &CommandSlice{s: []string{"QUIT"}},
	}
	SlotCmd = Completed{
		cs: &CommandSlice{s: []string{"CLUSTER", "SLOTS"}},
	}
	AskingCmd = Completed{
		cs: &CommandSlice{s: []string{"ASKING"}},
	}
)

type Completed struct {
	cs *CommandSlice
	cf uint16
	ks uint16
}

func (c *Completed) IsEmpty() bool {
	return c.cs == nil || len(c.cs.s) == 0
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
	return c.cs.s
}

func (c *Completed) CommandSlice() *CommandSlice {
	return c.cs
}

func (c *Completed) Slot() uint16 {
	return c.ks
}

type Cacheable Completed

func (c *Cacheable) Slot() uint16 {
	return c.ks
}

func (c *Cacheable) Commands() []string {
	return c.cs.s
}

func (c *Cacheable) CommandSlice() *CommandSlice {
	return c.cs
}

func (c *Cacheable) CacheKey() (key, command string) {
	if len(c.cs.s) == 2 {
		return c.cs.s[1], c.cs.s[0]
	}

	length := 0
	for i, v := range c.cs.s {
		if i == 1 {
			continue
		}
		length += len(v)
	}
	sb := strings.Builder{}
	sb.Grow(length)
	for i, v := range c.cs.s {
		if i == 1 {
			key = v
		} else {
			sb.WriteString(v)
		}
	}
	return key, sb.String()
}

func NewCompleted(ss []string) Completed {
	return Completed{cs: &CommandSlice{s: ss}}
}

func NewBlockingCompleted(ss []string) Completed {
	return Completed{cs: &CommandSlice{s: ss}, cf: blockTag}
}

func NewReadOnlyCompleted(ss []string) Completed {
	return Completed{cs: &CommandSlice{s: ss}, cf: readonly}
}

func NewMultiCompleted(cs [][]string) []Completed {
	ret := make([]Completed, len(cs))
	for i, c := range cs {
		ret[i] = NewCompleted(c)
	}
	return ret
}

func check(prev, new uint16) uint16 {
	if prev == InitSlot || prev == new {
		return new
	}
	panic(multiKeySlotErr)
}

const multiKeySlotErr = "multi key command with different key slots are not allowed"
