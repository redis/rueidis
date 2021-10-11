package cmds

import "sync"

func NewBuilder() *Builder {
	return &Builder{sp: sync.Pool{New: func() interface{} {
		return make([]string, 0, 2)
	}}}
}

type Builder struct {
	sp sync.Pool
}

func (b *Builder) get() []string {
	return b.sp.Get().([]string)
}

func (b *Builder) Put(s []string) {
	b.sp.Put(s[:0])
}
