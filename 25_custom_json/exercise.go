package customjson

import (
	"encoding/json"
	"time"
)

// Level is a small enum encoded as a JSON string.
type Level string

const (
	LevelInfo  Level = "info"
	LevelWarn  Level = "warn"
	LevelError Level = "error"
)

// MarshalJSON encodes the enum as a JSON string.
//
// TODO:
// - Encode only the known enum values above.
// - Return an error for unknown values.
func (l Level) MarshalJSON() ([]byte, error) {
	panic("TODO: implement MarshalJSON")
}

// UnmarshalJSON decodes the enum from a JSON string.
//
// TODO:
// - Accept only the known enum values above.
// - Return an error for unknown values.
func (l *Level) UnmarshalJSON(data []byte) error {
	panic("TODO: implement UnmarshalJSON")
}

// MillisDuration encodes a duration as an integer millisecond count.
type MillisDuration time.Duration

// MarshalJSON encodes the duration as a JSON number of milliseconds.
//
// TODO:
// - Convert the duration to whole milliseconds.
func (d MillisDuration) MarshalJSON() ([]byte, error) {
	panic("TODO: implement MarshalJSON")
}

// UnmarshalJSON decodes a non-negative millisecond count.
//
// TODO:
// - Decode a JSON number of milliseconds.
// - Reject negative values.
func (d *MillisDuration) UnmarshalJSON(data []byte) error {
	panic("TODO: implement UnmarshalJSON")
}

// Policy is a small config payload that uses both custom JSON types.
type Policy struct {
	Level   Level          `json:"level"`
	Timeout MillisDuration `json:"timeoutMillis"`
}

var _ = json.Marshal
