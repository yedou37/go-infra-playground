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
// This is the classic "functional options" pattern.
//
// The key idea is that constructor parameters are represented as small
// functions that modify a config object:
//
//	WithName("scheduler")        -> returns func(*Config)
//	WithTimeout(2 * time.Second) -> returns func(*Config)
//
// NewConfig first creates a Config with sane defaults, then applies the
// provided Option values from left to right.
//
// Why use this instead of a long constructor like:
//
//	NewConfig(name string, timeout time.Duration, maxInFlight int)
//
// Because industrial code often grows more configuration over time. Functional
// options make that growth easier to manage:
//   - Call sites stay readable because each argument is named by the helper.
//   - Most fields can have defaults, so callers only override what they care
//     about.
//   - Adding a new option usually does not require changing every existing call.
//   - Option application order is explicit, so later options can intentionally
//     override earlier ones.
//
// Another subtle point:
// WithName("default") does not change anything by itself. It only creates an
// Option function. The mutation happens only when that returned function is
// applied to a Config, for example:
//
//	c := Config{}
//	opt := WithName("scheduler")
//	opt(&c)
//
// In this exercise, Option is just a function type, but the pattern is common
// in clients, servers, controllers, and libraries where constructors would
// otherwise become long and brittle.
type Option func(*Config)

// WithName returns an Option that sets the Name field.
func WithName(name string) Option {
	return func(c *Config) {
		c.Name = name
	}
}

// WithTimeout returns an Option that sets Timeout if d > 0.
func WithTimeout(d time.Duration) Option {
	return func(c *Config) {
		if d > 0 {
			c.Timeout = d
		}
	}
}

// WithMaxInFlight returns an Option that sets MaxInFlight if n > 0.
func WithMaxInFlight(n int) Option {
	return func(c *Config) {
		if n > 0 {
			c.MaxInFlight = n
		}
	}
}

// NewConfig builds a Config with defaults and then applies opts in order.
//
// The constructor has two phases:
// 1. Create a base Config with defaults.
// 2. Apply each Option to that base Config from left to right.
//
// The left-to-right rule matters. For example:
//
//	NewConfig(WithName("apiserver"), WithName("scheduler"))
//
// should end with Name == "scheduler" because the later option wins.
//
// This "defaults first, overrides second" shape is a very common Go design for
// constructors that need both readability and future extensibility.
//
// TODO:
//   - Start from these defaults:
//     Name: "default"
//     Timeout: 5 * time.Second
//     MaxInFlight: 1
//   - Apply options from left to right.
func NewConfig(opts ...Option) Config {
	c := Config{
		Name:        "default",
		Timeout:     5 * time.Second,
		MaxInFlight: 1,
	}
	for _, opt := range opts {
		opt(&c)
	}
	return c
}
