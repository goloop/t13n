package t13n

import (
	"strings"
	"sync"
	"testing"

	"github.com/goloop/t13n/v2/lang"
)

// isASCII reports whether every byte of s is 7-bit ASCII.
func isASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > 127 {
			return false
		}
	}

	return true
}

// TestStringBounds pins the boundary behaviour of String, including the two
// out-of-range cases that used to be wrong: a negative rune (which panicked
// with an index out of range) and a rune past the end of the table.
func TestStringBounds(t *testing.T) {
	cases := []struct {
		name string
		c    rune
		want string
	}{
		{"ascii letter", 'A', "A"},
		{"negative rune", -1, ""},               // must not panic
		{"min int32", -2147483648, ""},          // must not panic
		{"table overflow", rune(tableSize), ""}, // first index past the end
		{"bmp edge", 0xFFFD, ""},
		{"supplementary emoji", 0x1F600, ""}, // outside the BMP
		{"cjk ext-b", 0x20000, ""},           // outside the BMP
		{"max rune", 0x10FFFF, ""},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := String(c.c); got != c.want {
				t.Errorf("String(%d): got %q want %q", c.c, got, c.want)
			}
		})
	}
}

// TestRune checks that Rune distinguishes a real mapping from "no mapping",
// including for runes that String reports only as an empty string.
func TestRune(t *testing.T) {
	cases := []struct {
		c    rune
		want string
		ok   bool
	}{
		{'A', "A", true},
		{'世', "Shi ", true},
		{0, "", false},       // NUL maps to empty -> not a mapping
		{-1, "", false},      // negative -> out of range
		{0x1F600, "", false}, // emoji -> out of range
		{rune(tableSize), "", false},
	}

	for _, c := range cases {
		got, ok := Rune(c.c)
		if got != c.want || ok != c.ok {
			t.Errorf("Rune(%d): got (%q, %v) want (%q, %v)",
				c.c, got, ok, c.want, c.ok)
		}
	}
}

// TestSupplementaryPlane checks that characters outside the BMP (emoji, CJK
// Ext-B, etc.) are dropped rather than crashing, and that the surrounding
// ASCII survives — the failure this documents used to lose data silently.
func TestSupplementaryPlane(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"a😀b", "ab"},
		{"𠀀", ""},          // CJK Ext-B U+20000
		{"x𝟘y", "xy"},      // MATHEMATICAL DIGIT ZERO U+1D7D8
		{"🇺🇦Київ", "Kyiv"}, // regional indicators dropped, rest kept
	}

	for _, c := range cases {
		if got := Trans(lang.UK, c.in); got != c.want {
			t.Errorf("Trans(UK, %q): got %q want %q", c.in, got, c.want)
		}
	}
}

// TestASCIIInvariant is the core guarantee: whatever the input, the output is
// pure ASCII. It also exercises deliberately hostile input (invalid UTF-8,
// lone surrogates written as bytes) to prove nothing panics.
func TestASCIIInvariant(t *testing.T) {
	inputs := []string{
		"",
		"Hello, World!",
		"Ĉu vi parolas Esperanton?",
		"Doброго 世界 вечора",
		"\xff\xfe\x00\x80",                // invalid UTF-8 bytes
		string([]rune{-1, 'A', 0x10FFFF}), // -1 folds to U+FFFD
		strings.Repeat("Щ", 1000),
	}
	langs := []string{lang.None, lang.UK, lang.RU, lang.DE, lang.BG, lang.EO}

	for _, l := range langs {
		for _, in := range inputs {
			out := Trans(l, in)
			if !isASCII(out) {
				t.Errorf("Trans(%q, %q) produced non-ASCII %q", l, in, out)
			}
		}
	}
}

// TestASCIIPassthrough checks that printable ASCII text is transliterated to
// itself (transliteration must not mangle text that is already ASCII).
func TestASCIIPassthrough(t *testing.T) {
	const s = "The quick brown fox jumps over 13 lazy dogs. (v2.0)"
	if got := Make(s); got != s {
		t.Errorf("Make(%q) = %q, want unchanged", s, got)
	}
}

// TestConcurrentUse runs the package functions and a shared instance from many
// goroutines at once and checks every result against a sequential reference.
// It is the regression test for the old global-state data race and also
// exercises the lazy table initialisation under concurrent first access.
// Run with -race to make the guarantee meaningful.
func TestConcurrentUse(t *testing.T) {
	inputs := []string{
		"Доброго вечора", "Згурський", "世界", "Müller", "Ĉapelo", "Яя",
	}

	ref := make([]string, len(inputs))
	for i, in := range inputs {
		ref[i] = Trans(lang.UK, in)
	}

	tr := New(WithLang(lang.UK))

	var wg sync.WaitGroup
	for g := 0; g < 64; g++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 300; i++ {
				idx := i % len(inputs)
				if got := Trans(lang.UK, inputs[idx]); got != ref[idx] {
					t.Errorf("package: got %q want %q", got, ref[idx])
					return
				}
				if got := tr.Make(inputs[idx]); got != ref[idx] {
					t.Errorf("instance: got %q want %q", got, ref[idx])
					return
				}
			}
		}()
	}
	wg.Wait()
}
