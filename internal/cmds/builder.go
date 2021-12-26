package cmds

import "sync"

type CommandSlice struct {
	s []string
}

func NewBuilder() *Builder {
	return &Builder{sp: sync.Pool{New: func() interface{} {
		return &CommandSlice{s: make([]string, 0, 2)}
	}}}
}

func NewSBuilder() *SBuilder {
	return &SBuilder{sp: sync.Pool{New: func() interface{} {
		return &CommandSlice{s: make([]string, 0, 2)}
	}}}
}

type Builder struct {
	sp sync.Pool
}

func (b *Builder) get() *CommandSlice {
	return b.sp.Get().(*CommandSlice)
}

func (b *Builder) Put(cs *CommandSlice) {
	cs.s = cs.s[:0]
	b.sp.Put(cs)
}

type SBuilder struct {
	sp sync.Pool
}

func (b *SBuilder) get() *CommandSlice {
	return b.sp.Get().(*CommandSlice)
}

func (b *SBuilder) Put(cs *CommandSlice) {
	cs.s = cs.s[:0]
	b.sp.Put(cs)
}
