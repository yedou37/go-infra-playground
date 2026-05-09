package iorw

import (
	"bytes"
	"errors"
	"io"
	"testing"
)

func TestSliceReaderEmptyBuffer(t *testing.T) {
	r := NewSliceReader([]byte("hello"))
	n, err := r.Read(nil)
	if n != 0 || err != nil {
		t.Fatalf("Read(nil) want (0, nil), got (%d, %v)", n, err)
	}
}

func TestSliceReaderFullCopyViaIoReadAll(t *testing.T) {
	r := NewSliceReader([]byte("hello world"))
	got, err := io.ReadAll(r)
	if err != nil {
		t.Fatalf("ReadAll err: %v", err)
	}
	if string(got) != "hello world" {
		t.Fatalf("got %q", string(got))
	}
}

func TestSliceReaderEOFOnExhaustion(t *testing.T) {
	r := NewSliceReader([]byte("ab"))
	buf := make([]byte, 1)
	if n, err := r.Read(buf); n != 1 || err != nil {
		t.Fatalf("first read: (%d,%v); want (1,nil)", n, err)
	}
	if n, err := r.Read(buf); n != 1 || err != nil {
		t.Fatalf("second read: (%d,%v); want (1,nil)", n, err)
	}
	if n, err := r.Read(buf); !(n == 0 && errors.Is(err, io.EOF)) {
		t.Fatalf("third read: (%d,%v); want (0, io.EOF)", n, err)
	}
}

func TestSliceReaderDefensiveCopy(t *testing.T) {
	src := []byte("abc")
	r := NewSliceReader(src)
	src[0] = 'Z'
	got, _ := io.ReadAll(r)
	if string(got) != "abc" {
		t.Fatalf("reader must hold a defensive copy; got %q", string(got))
	}
}

func TestCountingWriterTracksBytes(t *testing.T) {
	var buf bytes.Buffer
	w := &CountingWriter{W: &buf}
	if _, err := w.Write([]byte("hello")); err != nil {
		t.Fatalf("write err: %v", err)
	}
	if _, err := w.Write([]byte(" world")); err != nil {
		t.Fatalf("write err: %v", err)
	}
	if buf.String() != "hello world" {
		t.Fatalf("underlying got %q", buf.String())
	}
	if w.N != int64(len("hello world")) {
		t.Fatalf("N=%d, want %d", w.N, len("hello world"))
	}
}

// failingWriter writes the first `okN` bytes and then returns an error.
type failingWriter struct {
	okN int
}

func (f *failingWriter) Write(p []byte) (int, error) {
	if len(p) <= f.okN {
		f.okN -= len(p)
		return len(p), nil
	}
	wrote := f.okN
	f.okN = 0
	return wrote, errors.New("disk full")
}

func TestCountingWriterPartialWrite(t *testing.T) {
	w := &CountingWriter{W: &failingWriter{okN: 4}}
	n, err := w.Write([]byte("hello"))
	if err == nil {
		t.Fatal("expected partial-write error")
	}
	if n != 4 || w.N != 4 {
		t.Fatalf("partial write tracking wrong; n=%d N=%d", n, w.N)
	}
}

func TestComposesWithIoCopy(t *testing.T) {
	r := NewSliceReader([]byte("payload"))
	var buf bytes.Buffer
	w := &CountingWriter{W: &buf}
	n, err := io.Copy(w, r)
	if err != nil {
		t.Fatalf("io.Copy err: %v", err)
	}
	if n != int64(len("payload")) || w.N != n {
		t.Fatalf("Copy n=%d, w.N=%d, want %d", n, w.N, len("payload"))
	}
	if buf.String() != "payload" {
		t.Fatalf("buf=%q", buf.String())
	}
}
