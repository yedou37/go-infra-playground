package onceinit

import "sync"

// Loader lazily initializes one value.
type Loader struct {
	once sync.Once
	v    string
}

// Get returns the initialized value.
//
// TODO:
// - Use sync.Once so init runs at most once.
// - Store and return the initialized value.
func (l *Loader) Get(init func() string) string {
	panic("TODO: implement Get")
}
