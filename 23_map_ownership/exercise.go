package mapownership

// Store keeps labels internally.
type Store struct {
	labels map[string]string
}

// SetLabels replaces the internal labels with a copy of src.
//
// TODO:
// - Preserve nil.
// - Allocate a fresh map and copy every entry.
func (s *Store) SetLabels(src map[string]string) {
	panic("TODO: implement SetLabels")
}

// Labels returns a copy of the internal labels.
//
// TODO:
// - Preserve nil.
// - Return a copied map so callers cannot mutate internal state.
func (s *Store) Labels() map[string]string {
	panic("TODO: implement Labels")
}

// Set stores one key/value pair, allocating the internal map if needed.
func (s *Store) Set(key, value string) {
	panic("TODO: implement Set")
}
