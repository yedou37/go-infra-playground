package atomiccounter

import "sync/atomic"

// Counter stores one integer count using atomic operations.
type Counter struct {
	n atomic.Int64
}

// Add increments the counter by delta and returns the new value.
func (c *Counter) Add(delta int64) int64 {
	panic("TODO: implement Add")
}

// Load returns the current value.
func (c *Counter) Load() int64 {
	panic("TODO: implement Load")
}

// Reset sets the counter back to zero.
func (c *Counter) Reset() {
	panic("TODO: implement Reset")
}
