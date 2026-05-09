// Package embed is about Go's struct embedding, the idiomatic alternative
// to inheritance.
//
// What this exercise pins down:
//
//   - Embedding promotes fields and methods of the embedded type to the
//     outer type. You can call them as if they were defined on the outer.
//   - You can OVERRIDE a promoted method by defining a method of the same
//     name on the outer type. Inside the override you can still reach the
//     inner one via the embedded field name.
//   - Promoted methods make the outer type satisfy interfaces the inner
//     type satisfies — without writing forwarding code.
//   - Embedding != inheritance. There is no "super"; you call inner
//     methods explicitly through the embedded field.
//
// You will build a small Logger and a NoopLogger, plus a "tagged" logger
// that wraps another Logger and prefixes every line.
package embed

import "strings"

// Logger is the interface every concrete logger satisfies.
type Logger interface {
	Log(msg string)
	Lines() []string
}

// BaseLogger captures every Log call into a slice.
type BaseLogger struct {
	captured []string
}

// Log appends msg to captured.
//
// TODO:
//   - Implement so that subsequent Lines() returns the captured slice in
//     insertion order.
func (b *BaseLogger) Log(msg string) {
	panic("TODO: implement BaseLogger.Log")
}

// Lines returns a defensive copy of the captured lines.
//
// TODO:
//   - Return a freshly allocated []string with the same contents.
//   - Returning the internal slice is wrong because callers could
//     mutate it; ownership exercises in this repo cover that point.
func (b *BaseLogger) Lines() []string {
	panic("TODO: implement BaseLogger.Lines")
}

// PrefixLogger embeds *BaseLogger and prepends a fixed Prefix to every
// message before forwarding.
//
// You should NOT redeclare Lines() — it is automatically promoted from
// *BaseLogger because of embedding.
type PrefixLogger struct {
	*BaseLogger
	Prefix string
}

// NewPrefixLogger returns a PrefixLogger that owns its own BaseLogger.
//
// TODO:
//   - Allocate the inner *BaseLogger so callers don't have to.
//   - Set the prefix.
func NewPrefixLogger(prefix string) *PrefixLogger {
	panic("TODO: implement NewPrefixLogger")
}

// Log overrides the promoted Log method.
//
// TODO:
//   - Build "<Prefix>: <msg>" (use strings.Builder or simple concatenation).
//   - Delegate to the embedded BaseLogger via p.BaseLogger.Log(...).
//   - Do NOT call p.Log(...) here, that would recurse.
func (p *PrefixLogger) Log(msg string) {
	panic("TODO: implement PrefixLogger.Log; remember to call p.BaseLogger.Log")
}

// Helper used in tests; not part of the exercise per se.
func joinLines(lines []string) string {
	return strings.Join(lines, "\n")
}
