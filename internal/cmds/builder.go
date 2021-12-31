package cmds

import "sync"

type CommandSlice struct {
	s []string
}

func NewBuilder(initSlot uint16) *Builder {
	return &Builder{ks: initSlot, sp: sync.Pool{New: func() interface{} {
		return &CommandSlice{s: make([]string, 0, 2)}
	}}}
}

type Builder struct {
	sp sync.Pool
	ks uint16
}

func (b *Builder) get() *CommandSlice {
	return b.sp.Get().(*CommandSlice)
}

func (b *Builder) Put(cs *CommandSlice) {
	cs.s = cs.s[:0]
	b.sp.Put(cs)
}
