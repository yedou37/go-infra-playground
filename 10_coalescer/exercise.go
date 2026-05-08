package coalescer

// Coalescer turns many notifications into a smaller number of wakeups.
//
// A common controller pattern is:
// multiple updates happen quickly, but the consumer only needs to know that
// "something changed" and can reconcile once.
type Coalescer struct {
	ch chan struct{}
}

// New returns a Coalescer whose notification channel is buffered with capacity
// 1.
//
// TODO:
// - Allocate the channel.
// - Return a usable zero-work object.
func New() *Coalescer {
	panic("TODO: implement New")
}

// Notify reports that work may be available.
//
// TODO:
// - Send one empty struct into the channel if it is currently empty.
// - Do not block.
// - If a notification is already pending, coalesce the new one into the
//   existing pending signal.
func (c *Coalescer) Notify() {
	panic("TODO: implement Notify")
}

// C returns the receive-only notification channel.
func (c *Coalescer) C() <-chan struct{} {
	panic("TODO: implement C")
}
