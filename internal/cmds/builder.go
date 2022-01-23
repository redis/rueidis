package cmds

import "sync"

// CommandSlice is the command container managed by the sync.Pool
type CommandSlice struct {
	s []string
}

// NewBuilder creates a Builder and initializes the internal sync.Pool
func NewBuilder(initSlot uint16) *Builder {
	return &Builder{ks: initSlot, sp: sync.Pool{New: func() interface{} {
		return &CommandSlice{s: make([]string, 0, 2)}
	}}}
}

// Builder builds commands by reusing CommandSlice from the sync.Pool
type Builder struct {
	sp sync.Pool
	ks uint16
}

func (b *Builder) get() *CommandSlice {
	return b.sp.Get().(*CommandSlice)
}

// Put recycles the CommandSlice
func (b *Builder) Put(cs *CommandSlice) {
	cs.s = cs.s[:0]
	b.sp.Put(cs)
}
