package sliceownership

// Store keeps a slice internally.
//
// This exercise is about ownership: callers should not be able to mutate the
// internal state accidentally through shared backing storage.
type Store struct {
	items []int
}

// Set replaces the contents of the store with src.
//
// TODO:
// - Do not keep src directly.
// - Copy src so later caller mutations do not affect the store.
// - Preserve nil if src is nil.
func (s *Store) Set(src []int) {
	panic("TODO: implement Set")
}

// Snapshot returns a copy of the stored items.
//
// TODO:
// - Return nil if the store is empty and s.items is nil.
// - Return a copy, not the internal slice itself.
func (s *Store) Snapshot() []int {
	panic("TODO: implement Snapshot")
}

// Append appends one value to the internal slice.
func (s *Store) Append(v int) {
	panic("TODO: implement Append")
}
