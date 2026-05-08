package options

import "time"

// Config models a small constructor-style configuration object.
type Config struct {
	Name        string
	Timeout     time.Duration
	MaxInFlight int
}

// Option mutates a Config.
//
// This is the classic "functional options" pattern: constructor behavior stays
// readable while remaining extensible over time.
type Option func(*Config)

// WithName returns an Option that sets the Name field.
func WithName(name string) Option {
	panic("TODO: implement WithName")
}

// WithTimeout returns an Option that sets Timeout if d > 0.
func WithTimeout(d time.Duration) Option {
	panic("TODO: implement WithTimeout")
}

// WithMaxInFlight returns an Option that sets MaxInFlight if n > 0.
func WithMaxInFlight(n int) Option {
	panic("TODO: implement WithMaxInFlight")
}

// NewConfig builds a Config with defaults and then applies opts in order.
//
// TODO:
// - Start from these defaults:
//   Name: "default"
//   Timeout: 5 * time.Second
//   MaxInFlight: 1
// - Apply options from left to right.
func NewConfig(opts ...Option) Config {
	panic("TODO: implement NewConfig")
}
