package rueidisaside

import "testing"

func TestBeginFlightRejectsStaleGeneration(t *testing.T) {
	c := &Client{
		id:      "new-generation",
		flights: make(map[flightKey]*flight),
	}
	if f, _ := c.beginFlight("key", "old-generation"); f != nil {
		t.Fatal("stale generation created a flight")
	}
	if f, leader := c.beginFlight("key", "new-generation"); f == nil || !leader {
		t.Fatal("current generation did not create a flight")
	}
}
