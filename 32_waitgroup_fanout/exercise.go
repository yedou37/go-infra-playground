package waitgroupfanout

import "sync"

// FanOut applies fn to every input concurrently and returns the results in
// input order.
//
// TODO:
// - Start one goroutine per input.
// - Use sync.WaitGroup to wait for all goroutines.
// - Preserve result order.
func FanOut[T any, U any](inputs []T, fn func(T) U) []U {
	panic("TODO: implement FanOut")
}

var _ sync.WaitGroup
