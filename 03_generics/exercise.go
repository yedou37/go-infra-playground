package generics

// Pair stores two values of the same type.
//
// `T any` means T is a type parameter, and `any` is the constraint.
// In this case, any type is allowed.
type Pair[T any] struct {
	First  T
	Second T
}

// Swap returns a Pair with First and Second exchanged.
//
// TODO:
// - Preserve the same type parameter T.
func Swap[T any](p Pair[T]) Pair[T] {
	panic("TODO: implement Swap")
}

// MapSlice applies fn to each element in src and returns the mapped result.
//
// TODO:
// - Preserve nil: if src is nil, return nil.
// - Allocate the result with the correct length.
// - This function uses two type parameters: one for input and one for output.
func MapSlice[T any, U any](src []T, fn func(T) U) []U {
	panic("TODO: implement MapSlice")
}

// Last returns the last element from src.
//
// TODO:
// - Return the zero value of T and false if src is empty.
func Last[T any](src []T) (T, bool) {
	panic("TODO: implement Last")
}
