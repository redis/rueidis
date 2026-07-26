package rueidisaside

type flightKey struct {
	key string
	id  string // client id generation
}

type flight struct {
	done chan struct{}
}

func (c *Client) beginFlight(key, id string) (f *flight, leader bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.id != id {
		return nil, false
	}
	fk := flightKey{key: key, id: id}
	if f = c.flights[fk]; f != nil {
		return f, false
	}
	f = &flight{done: make(chan struct{})}
	c.flights[fk] = f
	return f, true
}

func (c *Client) finishFlight(key, id string, f *flight) {
	c.mu.Lock()
	defer c.mu.Unlock()

	fk := flightKey{key: key, id: id}
	if c.flights[fk] == f {
		delete(c.flights, fk)
		close(f.done)
	}
}

func (c *Client) resetFlightsLocked() {
	for _, f := range c.flights {
		close(f.done)
	}
	c.flights = make(map[flightKey]*flight)
}
