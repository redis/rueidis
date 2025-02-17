package cache

type node[V any] struct {
	key  string
	next *node[V]
	val  V
}
type chain[V any] struct {
	node[V]
}

func (h *chain[V]) find(key string) (val V, ok bool) {
	if h.node.key == key {
		return h.node.val, true
	}
	for curr := h.node.next; curr != nil; curr = curr.next {
		if curr.key == key {
			return curr.val, true
		}
	}
	return val, ok
}

func (h *chain[V]) insert(key string, val V) {
	if h.node.key == "" {
		h.node.key = key
		h.node.val = val
	} else if h.node.key == key {
		h.node.val = val
	} else {
		n := &node[V]{key: key, val: val}
		n.next = h.node.next
		h.node.next = n
	}
}

func (h *chain[V]) empty() bool {
	return h.node.next == nil && h.node.key == ""
}

func (h *chain[V]) delete(key string) bool {
	var zero V
	if h.node.key == key {
		h.node.key = ""
		h.node.val = zero
		return h.node.next == nil
	}

	if h.node.next == nil {
		return h.node.key == ""
	}

	if h.node.next.key == key {
		h.node.next.key = ""
		h.node.next.val = zero
		h.node.next, h.node.next.next = h.node.next.next, nil
		return h.empty()
	}

	prev := h.node.next
	curr := h.node.next.next
	for curr != nil {
		if curr.key == key {
			curr.key = ""
			curr.val = zero
			prev.next, curr.next = curr.next, nil
			break
		}
		prev, curr = curr, curr.next
	}
	return h.empty()
}
