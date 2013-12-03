package scheduling

/*
Copyright (c) 2013 Brandscreen Pty Ltd

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

import (
	"sync/atomic"
	"time"
)

/*
RateLimiter provides the capability of limiting some activity using a given rate.

This could be a requests per second to a backend web service or a number of emails per day.
*/
type RateLimiter struct {
	Limit     uint64
	Remaining int64
}

// NewRateLimiter returns a new RateLimiter using the given rate (limit/duration)
func NewRateLimiter(limit uint64, d time.Duration) *RateLimiter {
	limiter := &RateLimiter{
		Limit:     limit,
		Remaining: int64(limit),
	}

	// The following goroutine renews the capacity at the given duration up to the specified limit
	go (func(r *RateLimiter) {
		for _ = range time.Tick(d) {
			atomic.StoreInt64(&r.Remaining, int64(r.Limit))
			if atomic.LoadUint64(&r.Limit) == 0 {
				return
			}
		}
	})(limiter)

	return limiter
}

// Schedule returns true if the given activity may be scheduled according to the
// specified rate limits.
func (r *RateLimiter) Schedule() bool {
	// If there is remaining capacity then the activity may be scheduled, and
	// we take one of the remaining activities
	remaining := atomic.LoadInt64(&r.Remaining)
	if remaining > 0 {
		value := atomic.AddInt64(&r.Remaining, -1)
		if value >= 0 {
			return true
		}

		atomic.CompareAndSwapInt64(&r.Remaining, value, 0)
	}

	// Otherwise, the activity may not be scheduled.
	return false
}

// Close closes the rate limiter and cleans up
func (r *RateLimiter) Close() error {
	// Setting the limit to 0 will stop the goroutine created in NewRateLimiter
	atomic.StoreUint64(&r.Limit, 0)
	return nil
}
