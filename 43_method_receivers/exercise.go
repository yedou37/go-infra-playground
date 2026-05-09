// Package receivers is about value vs pointer receivers, which is one of
// the most-confused Go topics for people coming from other languages.
//
// What this exercise pins down:
//
//   - A method with a VALUE receiver gets a COPY of the struct. Mutations
//     in the method are invisible to the caller.
//   - A method with a POINTER receiver mutates the struct in place.
//   - The METHOD SET of `T` is methods with value receivers.
//     The METHOD SET of `*T` is methods with both value AND pointer
//     receivers. So an interface that requires a pointer-receiver method
//     can only be satisfied by *T, not T.
//   - This is why `var s Stringer = T{}` sometimes compiles and sometimes
//     does not, depending on whether String() has a value or pointer
//     receiver.
//
// You will implement a Counter both ways and a fmt.Stringer.
package receivers

// Counter holds an integer count.
type Counter struct {
	N int
}

// IncByValue tries to increment via a VALUE receiver. By Go's semantics
// this CANNOT actually mutate the caller's Counter; the test verifies
// that fact. The point of the exercise is to write the obvious-looking
// (but wrong) code so you can see why it's wrong.
//
// TODO:
//   - Implement IncByValue so that it does `c.N++` inside the method.
//   - Tests will assert that the caller's Counter is unchanged.
func (c Counter) IncByValue() {
	panic("TODO: implement IncByValue with a VALUE receiver")
}

// IncByPointer increments via a POINTER receiver. Mutations are visible.
//
// TODO:
//   - Implement so that c.N++ persists.
func (c *Counter) IncByPointer() {
	panic("TODO: implement IncByPointer with a POINTER receiver")
}

// String implements fmt.Stringer with a POINTER receiver.
//
// TODO:
//   - Return a string of the form "Counter(<N>)".
//   - Because the receiver is *Counter, only *Counter (not Counter)
//     satisfies fmt.Stringer. The test will verify this with an
//     interface-assignment check.
func (c *Counter) String() string {
	panic("TODO: implement *Counter.String")
}

// NewCounter returns *Counter so callers don't accidentally lose
// mutations by assigning to a value variable.
//
// TODO:
//   - Allocate and return &Counter{N: start}.
func NewCounter(start int) *Counter {
	panic("TODO: implement NewCounter")
}
