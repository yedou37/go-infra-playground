package arrayvsslice

// BumpArray returns a modified copy of a.
//
// This exercise contrasts array value semantics with slice reference-like
// semantics.
//
// TODO:
// - Increment every element by 1.
// - Return the modified array value.
func BumpArray(a [3]int) [3]int {
	panic("TODO: implement BumpArray")
}

// BumpSlice increments every element in xs in place.
//
// TODO:
// - Modify the provided slice directly.
func BumpSlice(xs []int) {
	panic("TODO: implement BumpSlice")
}

// ToSlice returns a slice view over a.
//
// TODO:
//   - Return a slice that refers to the same backing array as a local copy of the
//     input array value.
func ToSlice(a [3]int) []int {
	panic("TODO: implement ToSlice")
}
