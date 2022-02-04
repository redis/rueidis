package cmds

import "sync"

var pool = &sync.Pool{New: newCommandSlice}

// CommandSlice is the command container managed by the sync.Pool
type CommandSlice struct {
	s []string
}

func newCommandSlice() interface{} {
	return &CommandSlice{s: make([]string, 0, 2)}
}

// NewBuilder creates a Builder and initializes the internal sync.Pool
func NewBuilder(initSlot uint16) Builder {
	return Builder{ks: initSlot}
}

// Builder builds commands by reusing CommandSlice from the sync.Pool
type Builder struct {
	ks uint16
}

func get() *CommandSlice {
	return pool.Get().(*CommandSlice)
}

// Put recycles the CommandSlice
func Put(cs *CommandSlice) {
	cs.s = cs.s[:0]
	pool.Put(cs)
}
