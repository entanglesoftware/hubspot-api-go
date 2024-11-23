package decorator

import (
	"errors"
	"sync"
	"time"
)

// LimiterDecorator manages rate limiting for decorated functions.
type LimiterDecorator struct {
	rateLimit      time.Duration
	maxConcurrent  int
	tokens         chan struct{}
	limiterEnabled bool
	mutex          sync.Mutex
}

// NewLimiterDecorator initializes a LimiterDecorator with specified options.
func NewLimiterDecorator(rateLimit time.Duration, maxConcurrent int) *LimiterDecorator {
	decorator := &LimiterDecorator{
		rateLimit:      rateLimit,
		maxConcurrent:  maxConcurrent,
		tokens:         make(chan struct{}, maxConcurrent),
		limiterEnabled: true,
	}

	// Fill tokens to manage concurrency.
	for i := 0; i < maxConcurrent; i++ {
		decorator.tokens <- struct{}{}
	}

	return decorator
}

// Decorate wraps the given method with rate limiting.
func (d *LimiterDecorator) Decorate(method func(args ...interface{}) (interface{}, error)) func(args ...interface{}) (interface{}, error) {
	return func(args ...interface{}) (interface{}, error) {
		if !d.isLimiterEnabled() {
			return method(args...)
		}

		select {
		case <-d.tokens:
			defer d.replenishToken()
			return method(args...)
		default:
			return nil, errors.New("rate limit exceeded")
		}
	}
}

// EnableLimiter activates the rate limiter.
func (d *LimiterDecorator) EnableLimiter() {
	d.setLimiterEnabled(true)
}

// DisableLimiter deactivates the rate limiter.
func (d *LimiterDecorator) DisableLimiter() {
	d.setLimiterEnabled(false)
}

// replenishToken replenishes a token after the rate limit duration.
func (d *LimiterDecorator) replenishToken() {
	time.Sleep(d.rateLimit)
	d.tokens <- struct{}{}
}

// isLimiterEnabled safely retrieves the limiter status.
func (d *LimiterDecorator) isLimiterEnabled() bool {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	return d.limiterEnabled
}

// setLimiterEnabled safely sets the limiter status.
func (d *LimiterDecorator) setLimiterEnabled(enabled bool) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	d.limiterEnabled = enabled
}
