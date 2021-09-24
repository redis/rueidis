package queue

type Task struct {
	W chan interface{}
	C []string
}

type Queue interface {
	Put(m *Task)
	Next1(try bool) *Task
	Next2() *Task
}

// roundUp takes an uint64 greater than 0 and rounds it up to the next
// power of 2.
func roundUp(v uint64) uint64 {
	v--
	v |= v >> 1
	v |= v >> 2
	v |= v >> 4
	v |= v >> 8
	v |= v >> 16
	v |= v >> 32
	v++
	return v
}
