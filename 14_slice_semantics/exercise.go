package slicesemantics

// Packet intentionally mixes a pure value field and a slice field so you can
// observe which mutations affect only the local copy and which reach shared
// backing storage.
type Packet struct {
	Version int
	Data    []byte
}

// Sum returns the sum of xs without modifying the input.
//
// TODO:
// - Accept a slice directly.
// - Return 0 for nil or empty input.
// - Do not mutate xs.
func Sum(xs []int) int {
	panic("TODO: implement Sum")
}

// ZeroInPlace sets every element in xs to zero.
//
// TODO:
// - Modify the shared backing array through the slice.
// - Do not allocate a new slice.
func ZeroInPlace(xs []int) {
	panic("TODO: implement ZeroInPlace")
}

// AppendValue appends v and returns the resulting slice.
//
// TODO:
//   - Use append.
//   - Return the new slice.
//   - Callers should assign the returned value back if they want to observe any
//     header change.
func AppendValue(xs []int, v int) []int {
	panic("TODO: implement AppendValue")
}

// Clone returns a copy of xs with independent backing storage.
//
// TODO:
// - Preserve nil.
// - Allocate a new slice and copy the contents.
func Clone(xs []int) []int {
	panic("TODO: implement Clone")
}

// MutatePacket demonstrates mixed struct field behavior.
//
// TODO:
// - Increment Version on the local copy only.
// - If Data is non-empty, increment Data[0].
// - Return the local Version value after incrementing.
func MutatePacket(p Packet) int {
	panic("TODO: implement MutatePacket")
}
