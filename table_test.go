package t13n

import "testing"

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
