package queue

type Chan struct {
	ch1 chan *Task
	ch2 chan *Task
}

func NewChan(size uint64) *Chan {
	return &Chan{ch1: make(chan *Task, roundUp(size)), ch2: make(chan *Task, roundUp(size))}
}

func (r *Chan) Put(t *Task) {
	r.ch1 <- t
}

func (r *Chan) Next1() *Task {
	t := <-r.ch1
	r.ch2 <- t
	return t
}

func (r *Chan) Next2() *Task {
	return <-r.ch2
}
