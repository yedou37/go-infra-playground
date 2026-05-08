package subsliceleak

// StableWindow returns a detached copy of src[start:end].
//
// This exercise targets a subtle memory bug:
// returning a small subslice of a large buffer can accidentally keep the entire
// backing array alive. Copying avoids that retention and also prevents aliasing.
//
// TODO:
// - Validate bounds and return ok=false on invalid input.
// - Return a copied window, not a direct subslice.
// - Ensure the returned slice has no extra capacity beyond its length.
func StableWindow(src []byte, start, end int) (_ []byte, ok bool) {
	panic("TODO: implement StableWindow")
}
