package channels

// StartGenerator starts a goroutine that sends 0..n-1 to a buffered channel,
// then closes the channel.
//
// TODO:
// - Return a receive-only channel.
// - Use a buffered channel with capacity bufSize.
// - If n <= 0, still return a closed channel.
func StartGenerator(n, bufSize int) <-chan int {
	ch := make(chan int, bufSize)
	go func() {
		for i := range n {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

// Drain reads all remaining values from ch and returns them in order.
//
// TODO:
// - Accept a receive-only channel.
// - Range over the channel until it is closed.
func Drain(ch <-chan int) []int {
	var res []int
	for v := range ch {
		res = append(res, v)
	}
	return res
}

// TrySend attempts to send v into ch without blocking.
//
// TODO:
// - Accept a send-only channel.
// - Use select with default.
// - Return true if the send succeeds.
func TrySend(ch chan<- int, v int) bool {
	select {
	case ch <- v:
		return true
	default:
		return false
	}
}
