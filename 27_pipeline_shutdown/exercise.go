package pipelineshutdown

import "context"

// StartPipeline starts a two-stage pipeline:
// stage 1 adds 1, and stage 2 multiplies by 2.
//
// This mirrors the shape of real pipelines in infrastructure code:
// each stage runs in its own goroutine, all stages stop on cancellation, and
// the final output channel is closed exactly once after the pipeline exits.
//
// TODO:
// - Start two goroutines connected by an internal channel.
// - Stop both stages promptly when ctx is done.
// - Close the internal channel when stage 1 exits.
// - Close the output channel when stage 2 exits.
// - For each input value v, emit (v + 1) * 2.
func StartPipeline(ctx context.Context, in <-chan int) <-chan int {
	panic("TODO: implement StartPipeline")
}
