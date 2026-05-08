package bufferreuse

// Encoder reuses one internal buffer across calls.
//
// This is common in high-throughput code, but it is only safe if returned
// results do not expose that reusable buffer directly.
type Encoder struct {
	buf []byte
}

// EncodeInt writes the decimal form of n into the reusable buffer and returns
// the encoded bytes.
//
// TODO:
//   - Reuse e.buf across calls by resetting its length to zero.
//   - Support n in the range [0, 99].
//   - Return a copy of the encoded bytes so a later call cannot mutate a previous
//     result.
func (e *Encoder) EncodeInt(n int) []byte {
	panic("TODO: implement EncodeInt")
}
