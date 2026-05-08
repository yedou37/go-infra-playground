package interfaces

import "context"

// Event is a small domain value passed to handlers.
type Event struct {
	Name string
}

// Handler is intentionally small.
//
// In production Go code, tiny interfaces are often easier to test, easier to
// mock, and easier to evolve than large "god interfaces".
type Handler interface {
	Handle(context.Context, Event) error
}

// HandlerFunc adapts a function to the Handler interface.
type HandlerFunc func(context.Context, Event) error

// Handle lets HandlerFunc satisfy Handler.
//
// TODO:
// - Call f with the provided arguments.
func (f HandlerFunc) Handle(ctx context.Context, event Event) error {
	panic("TODO: implement Handle")
}

// Dispatch calls handlers in order with the same event.
//
// TODO:
// - Stop on the first error and return it.
// - If all handlers succeed, return nil.
// - Keep the signature simple: ctx first, then the domain value, then the
//   dependencies.
func Dispatch(ctx context.Context, event Event, handlers ...Handler) error {
	panic("TODO: implement Dispatch")
}
