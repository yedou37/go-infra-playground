package selectpatterns

import "sync"

// OrDone forwards values from in until either done is closed or in is closed.
//
// This pattern is useful when a caller wants to stop a pipeline stage early
// without leaking a goroutine that is still blocked on receive or send.
//
// TODO:
// - Return a receive-only channel.
// - Start a goroutine that forwards values.
// - Stop forwarding promptly when done is closed.
// - Close the output channel when forwarding stops.
func OrDone[T any](done <-chan struct{}, in <-chan T) <-chan T {
	panic("TODO: implement OrDone")
}

// FanIn merges multiple input channels into one output channel.
//
// This is a common select-based building block in controller-style code:
// many producers, one consumer, and a stop signal that should shut everything
// down cleanly.
//
// TODO:
// - Forward values from every input channel to one output channel.
// - Respect done and stop promptly when it is closed.
// - Close the output channel after all forwarding goroutines have exited.
// - Use a sync.WaitGroup to know when all forwarders are done.
func FanIn[T any](done <-chan struct{}, inputs ...<-chan T) <-chan T {
	panic("TODO: implement FanIn")
}

var _ sync.WaitGroup
