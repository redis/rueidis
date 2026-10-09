package rueidislimiter

import "time"

// Algorithm specifies the rate limiting algorithm to use.
type Algorithm int

const (
	// AlgorithmFixedWindow selects the legacy Fixed Window counter algorithm. Default.
	AlgorithmFixedWindow Algorithm = iota
	// AlgorithmGCRA selects the Generic Cell Rate Algorithm (leaky bucket).
	AlgorithmGCRA
)

type RateLimitOption struct {
	limit     int64
	window    time.Duration
	burst     int64
	algorithm Algorithm
	hasBurst  bool
	hasAlg    bool
}

// WithCustomRateLimit creates a RateLimitOption with custom limit and window.
// Burst defaults to limit.
func WithCustomRateLimit(limit int, window time.Duration) RateLimitOption {
	return RateLimitOption{
		limit:    int64(limit),
		window:   window,
		burst:    int64(limit),
		hasBurst: true,
	}
}

// WithCustomRateLimitAndBurst creates a RateLimitOption with custom limit, window, and burst.
func WithCustomRateLimitAndBurst(limit int, window time.Duration, burst int) RateLimitOption {
	return RateLimitOption{
		limit:    int64(limit),
		window:   window,
		burst:    int64(burst),
		hasBurst: true,
	}
}

// WithAlgorithm creates a RateLimitOption specifying the algorithm strategy.
func WithAlgorithm(alg Algorithm) RateLimitOption {
	return RateLimitOption{
		algorithm: alg,
		hasAlg:    true,
	}
}

// WithBurst creates a RateLimitOption specifying a custom burst capacity.
func WithBurst(burst int) RateLimitOption {
	return RateLimitOption{
		burst:    int64(burst),
		hasBurst: true,
	}
}
