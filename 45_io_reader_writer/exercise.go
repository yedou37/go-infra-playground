// Package iorw is about implementing io.Reader / io.Writer, the most
// pervasive interfaces in the standard library.
//
// What this exercise pins down:
//
//   - The Read method's contract:
//       Read(p []byte) (n int, err error)
//     - It MAY return 0 < n < len(p), even if more data is available.
//     - When the source is exhausted, return n=0 (or the last bytes) and
//       err == io.EOF. Idiomatically you can return (n>0, io.EOF) on the
//       very last call, OR (n>0, nil) followed by (0, io.EOF) on the next
//       call. Both are valid; tests only require behavior not specific
//       edge encoding.
//   - Once an error has been returned, the same error should keep being
//     returned on subsequent calls.
//   - The Write contract:
//       Write(p []byte) (n int, err error)
//     - Must write all of p unless an error occurred; n < len(p) implies err != nil.
//
// You will implement two tiny types and verify they compose with the
// standard library (io.Copy).
package iorw

import "io"

// SliceReader streams a fixed []byte to readers, like bytes.NewReader
// but stripped down. It MUST satisfy io.Reader.
type SliceReader struct {
	// TODO: store the source bytes and the current read offset.
}

// NewSliceReader returns a SliceReader over a defensive COPY of src,
// so that later mutations to src don't affect the reader.
func NewSliceReader(src []byte) *SliceReader {
	panic("TODO: implement NewSliceReader; defensive copy src")
}

// Read implements io.Reader.
//
// Required behavior:
//   - Copy from the internal buffer at the current offset into p.
//   - Advance the offset by the number of bytes copied.
//   - When at EOF, return (0, io.EOF) on subsequent calls.
//   - len(p) == 0 must return (0, nil) without advancing.
func (r *SliceReader) Read(p []byte) (int, error) {
	panic("TODO: implement SliceReader.Read")
}

// CountingWriter wraps an underlying io.Writer and counts how many bytes
// have been written successfully. It MUST satisfy io.Writer.
type CountingWriter struct {
	W io.Writer
	N int64
}

// Write implements io.Writer.
//
// Required behavior:
//   - Forward p to the underlying writer.
//   - Add the number of successfully written bytes to N.
//   - On a partial-write error, N must reflect what actually got through.
func (c *CountingWriter) Write(p []byte) (int, error) {
	panic("TODO: implement CountingWriter.Write")
}
