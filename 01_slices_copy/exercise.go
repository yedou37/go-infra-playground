package slicescopy

// CloneInts returns a new slice with the same contents as src.
//
// TODO:
// - Return nil if src is nil.
// - Allocate a new backing array.
// - Use copy instead of append.
func CloneInts(src []int) []int {
	panic("TODO: implement CloneInts")
}

// Window returns a copy of src[start:end].
//
// TODO:
// - Validate the bounds.
// - Return a copied slice, not a view into the original backing array.
// - Return ok=false on invalid bounds.
func Window(src []int, start, end int) (_ []int, ok bool) {
	panic("TODO: implement Window")
}

// Push appends value to dst and returns the result.
//
// TODO:
// - Use append.
// - Do not preallocate manually unless you can explain why.
func Push(dst []int, value int) []int {
	panic("TODO: implement Push")
}
