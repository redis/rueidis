//go:build !unix

package rueidis

import "time"

func processCPUTime() (user, sys time.Duration, ok bool) {
	return 0, 0, false
}
