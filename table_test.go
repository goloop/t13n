package t13n

import (
	"bytes"
	"compress/flate"
	"testing"
)

// deflate flate-compresses b so tests can build valid compressed payloads with
// arbitrary decoded contents.
func deflate(t *testing.T, b []byte) []byte {
	t.Helper()

	var buf bytes.Buffer
	w, err := flate.NewWriter(&buf, flate.DefaultCompression)
	if err != nil {
		t.Fatalf("flate.NewWriter: %v", err)
	}
	if _, err := w.Write(b); err != nil {
		t.Fatalf("write: %v", err)
	}
	if err := w.Close(); err != nil {
		t.Fatalf("close: %v", err)
	}

	return buf.Bytes()
}

// TestDecodeRoundTrip checks that the lazily decoded table is non-empty and
// self-consistent with the String accessor for a few well-known code points.
func TestDecodeRoundTrip(t *testing.T) {
	tbl := table()
	spot := map[rune]string{
		'A':  "A",
		'世':  "Shi ",
		'й':  "i",
		0x18: "", // a control code point maps to nothing
	}
	for c, want := range spot {
		if got := tbl[c]; got != want {
			t.Errorf("table[%d] = %q, want %q", c, got, want)
		}
	}
}

// TestDecodeCorrupt checks that a corrupt resource fails loudly instead of
// silently yielding a half-built table.
func TestDecodeCorrupt(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("decode did not panic on corrupt data")
		}
	}()

	decode([]byte{0xff, 0x00, 0x13, 0x37, 0xde, 0xad})
}

// TestDecodeTruncatedVarint checks that a valid flate stream whose decoded
// payload ends on a dangling varint continuation byte fails loudly with a
// panic rather than spinning in an endless loop.
func TestDecodeTruncatedVarint(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("decode did not panic on truncated varint")
		}
	}()

	// A lone 0x80 is a varint with the continuation bit set but no following
	// byte, so binary.Uvarint reports n == 0.
	decode(deflate(t, []byte{0x80}))
}

// TestDecodeTruncatedLength checks that a truncated varint in the length field
// (after a well-formed code-point delta) also panics rather than loops.
func TestDecodeTruncatedLength(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("decode did not panic on truncated length")
		}
	}()

	// 0x01 is a complete delta varint; the trailing 0x80 is a dangling
	// continuation byte in the length position.
	decode(deflate(t, []byte{0x01, 0x80}))
}
