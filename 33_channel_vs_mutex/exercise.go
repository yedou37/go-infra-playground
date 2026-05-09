package channelvsmutex

import "sync"

// MutexCounter counts using a mutex.
type MutexCounter struct {
	mu sync.Mutex
	n  int
}

// Inc increments the counter by one using the mutex.
func (c *MutexCounter) Inc() {
	panic("TODO: implement Inc")
}

// Load returns the current value using the mutex.
func (c *MutexCounter) Load() int {
	panic("TODO: implement Load")
}

// ChannelCounter counts by sending increment requests to one internal goroutine.
type ChannelCounter struct {
	inc  chan struct{}
	get  chan chan int
	stop chan struct{}
}

// NewChannelCounter starts the internal goroutine.
//
// TODO:
// - Allocate the channels.
// - Start a goroutine that owns the integer state.
// - On `inc`, increment the state.
// - On `get`, send the current state back.
// - On `stop`, exit the goroutine.
func NewChannelCounter() *ChannelCounter {
	panic("TODO: implement NewChannelCounter")
}

// Inc increments the counter by one through the internal goroutine.
func (c *ChannelCounter) Inc() {
	panic("TODO: implement Inc")
}

// Load returns the current value through the internal goroutine.
func (c *ChannelCounter) Load() int {
	panic("TODO: implement Load")
}

// Close stops the internal goroutine.
func (c *ChannelCounter) Close() {
	panic("TODO: implement Close")
}
