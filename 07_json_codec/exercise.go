package jsoncodec

import "encoding/json"

// LeaseSpec is a small config-like payload.
//
// The field names intentionally use the lowerCamelCase shape that is common in
// API payloads and Kubernetes-style JSON.
type LeaseSpec struct {
	Holder     string            `json:"holder"`
	TTLSeconds int               `json:"ttlSeconds"`
	Labels     map[string]string `json:"labels,omitempty"`
}

// Encode converts spec to JSON.
//
// TODO:
// - Use the standard library JSON encoder.
// - Respect the struct tags above.
// - Do not add extra spaces or indentation.
func Encode(spec LeaseSpec) ([]byte, error) {
	panic("TODO: implement Encode")
}

// Decode parses one JSON object into LeaseSpec.
//
// In production config loading and API decoding, "best effort" parsing is often
// too dangerous. Strict decoding catches typos early.
//
// TODO:
// - Decode exactly one JSON value into LeaseSpec.
// - Reject unknown fields.
// - Reject trailing non-whitespace data after the first JSON value.
// - Return an error for malformed or empty input.
func Decode(data []byte) (LeaseSpec, error) {
	panic("TODO: implement Decode")
}

var _ = json.Unmarshal
