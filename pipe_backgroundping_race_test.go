package rueidis

import (
	"sync"
	"testing"
	"time"
)

// TestPipe_BackgroundPing_NoDataRace tests that backgroundPing doesn't have data races
// when timer callbacks fire repeatedly. This is the simplest test to catch the race.
// This test will fail with -race if the data race exists in backgroundPing.
func TestPipe_BackgroundPing_NoDataRace(t *testing.T) {
	t.Parallel()

	if testing.Short() {
		t.Skip("Skipping race detection test in short mode")
	}

	// Use the existing setup function with a short keep-alive interval
	option := ClientOption{
		ConnWriteTimeout: 100 * time.Millisecond,
	}
	option.Dialer.KeepAlive = 20 * time.Millisecond // Short interval for rapid timer firings

	p, mock, cancel, closeConn := setup(t, option)
	defer cancel()
	_ = closeConn
	_ = mock
	_ = p

	// Simply let the pipe exist with background ping active
	// The background ping timer will fire multiple times
	// If there's a race in the prev variable or timer access, -race will catch it
	time.Sleep(300 * time.Millisecond) // Let timer fire ~15 times

	// The race detector will report any concurrent access to 'prev' variable
}

// TestPipe_BackgroundPing_ConcurrentClients tests backgroundPing with multiple
// concurrent clients, each with their own background ping timers.
func TestPipe_BackgroundPing_ConcurrentClients(t *testing.T) {
	t.Parallel()

	if testing.Short() {
		t.Skip("Skipping concurrent race test in short mode")
	}

	numClients := 10
	var wg sync.WaitGroup

	for i := 0; i < numClients; i++ {
		wg.Add(1)
		go func(clientID int) {
			defer wg.Done()

			// Use the existing setup function with a short keep-alive interval
			option := ClientOption{
				ConnWriteTimeout: 100 * time.Millisecond,
			}
			option.Dialer.KeepAlive = 30 * time.Millisecond

			_, _, cancel, _ := setup(t, option)
			defer cancel()

			// Let the pipe exist with background ping active
			time.Sleep(200 * time.Millisecond) // Let timer fire multiple times
		}(i)
	}

	wg.Wait()
}

// TestPipe_BackgroundPing_RapidConnectDisconnect reproduces the scenario
// where pipes are created and destroyed rapidly (like crash recovery).
func TestPipe_BackgroundPing_RapidConnectDisconnect(t *testing.T) {
	t.Parallel()

	if testing.Short() {
		t.Skip("Skipping rapid connect/disconnect test in short mode")
	}

	for iteration := 0; iteration < 20; iteration++ {
		// Use the existing setup function with a short keep-alive interval
		option := ClientOption{
			ConnWriteTimeout: 100 * time.Millisecond,
		}
		option.Dialer.KeepAlive = 40 * time.Millisecond

		p, _, cancel, _ := setup(t, option)

		// Brief operation period
		time.Sleep(50 * time.Millisecond)

		// Immediate close (like consumer crash)
		cancel()

		// Brief pause to let background goroutines settle
		time.Sleep(5 * time.Millisecond)

		_ = p
	}
}
