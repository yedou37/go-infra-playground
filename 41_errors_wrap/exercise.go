// Package errwrap is about Go's modern error model.
//
// Things this exercise pins down:
//
//   - fmt.Errorf("...: %w", err) wraps an error so that errors.Is/As can
//     walk the chain. Plain "%v" or "%s" does NOT wrap.
//   - errors.Is is for SENTINEL errors (predefined values like io.EOF,
//     ErrNotFound). It walks the wrap chain.
//   - errors.As is for TYPED errors: it tries to extract a specific
//     concrete (or interface) type from the chain.
//   - Custom error types should usually implement Error() and, if they
//     wrap something, Unwrap() *too*. This is the modern Go convention.
//
// In production Go, almost every error coming out of a non-trivial
// function is wrapped at least once on its way up. Knowing how to
// distinguish "what kind of error is this?" without doing string match
// is one of the most-used Go skills there is.
package errwrap

import (
	"errors"
	"fmt"
)

// ErrNotFound is a sentinel error that callers compare against using
// errors.Is.
var ErrNotFound = errors.New("not found")

// QuotaExceeded is a typed error carrying structured information about
// the violation. Callers use errors.As to extract it.
type QuotaExceeded struct {
	Resource string
	Limit    int
	Got      int
}

// Error implements the error interface for QuotaExceeded.
//
// TODO:
//   - Return a human-readable message that includes Resource, Limit and Got.
//   - Format is not asserted by tests; any sensible string is fine.
func (e *QuotaExceeded) Error() string {
	return fmt.Sprintf(
		"quota exceeded: resource=%s limit=%d got=%d",
		e.Resource, e.Limit, e.Got,
	)
}

// LookupAndAnnotate returns ErrNotFound wrapped with extra context, so that
// errors.Is(err, ErrNotFound) is still true.
//
// TODO:
//   - If `key` is empty, return errors.New("empty key") (NOT wrapping
//     ErrNotFound; an empty key is a different failure).
//   - Otherwise return an error of the form
//     fmt.Errorf("lookup %q: %w", key, ErrNotFound).
func LookupAndAnnotate(key string) error {
	if key == "" {
		return errors.New("empty key")
	}
	return fmt.Errorf("lookup %q: %w", key, ErrNotFound)
}

// CheckQuota returns a wrapped *QuotaExceeded when got > limit.
//
// TODO:
//   - If got <= limit, return nil.
//   - Else build a *QuotaExceeded and return it wrapped via fmt.Errorf
//     so that errors.As(err, &target) can still extract the concrete
//     *QuotaExceeded.
//   - The wrapping context should mention the resource, e.g.
//     fmt.Errorf("admission for %s: %w", resource, qe).
func CheckQuota(resource string, limit, got int) error {
	if got <= limit {
		return nil
	}
	err := &QuotaExceeded{
		Resource: resource,
		Limit:    limit,
		Got:      got,
	}
	return fmt.Errorf("admission for %s: %w", resource, err)
}

// IsNotFound reports whether err (or anything in its wrap chain) is
// ErrNotFound. Implement using errors.Is — DO NOT compare strings.
func IsNotFound(err error) bool {
	return errors.Is(err, ErrNotFound)
}

// AsQuotaExceeded extracts a *QuotaExceeded from err's chain.
// Returns (nil, false) if err is nil or no such error is present.
// Implement using errors.As — DO NOT type-switch directly.
func AsQuotaExceeded(err error) (*QuotaExceeded, bool) {
	if err == nil {
		return nil, false
	}
	var target *QuotaExceeded
	if !errors.As(err, &target) {
		return nil, false
	}
	return target, true
}
