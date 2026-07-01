package t13n

import (
	"strings"
	"testing"

	"github.com/goloop/t13n/v2/lang"
)

// TestWithFallback checks that a fallback supplies a replacement for runes
// with no mapping (unknown or outside the BMP) while leaving mapped runes and
// runes handled by the language rules untouched.
func TestWithFallback(t *testing.T) {
	tr := New(WithLang(lang.UK), WithFallback(func(rune) string { return "?" }))

	cases := []struct {
		in   string
		want string
	}{
		{"a😀b", "a?b"},    // emoji outside the BMP -> fallback
		{"Київ", "Kyiv"},  // fully mapped -> no fallback
		{"x\x18y", "x?y"}, // control code maps to "" -> fallback
	}

	for _, c := range cases {
		if got := tr.Make(c.in); got != c.want {
			t.Errorf("Make(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}

// TestFallbackYieldsToCustomRule checks that the custom rule still wins over
// the fallback, since it runs last.
func TestFallbackYieldsToCustomRule(t *testing.T) {
	ctr := func(ts lang.TransState) (string, int, bool) {
		if ts.Value == "?" {
			return "!", 0, true
		}
		return "", 0, false
	}

	tr := New(WithFallback(func(rune) string { return "?" }), WithRules(ctr))
	if got, want := tr.Make("😀"), "!"; got != want {
		t.Errorf("Make = %q, want %q", got, want)
	}
}

// TestStrictASCII checks that non-ASCII produced by a custom rule or fallback
// is stripped when strict mode is on, and passes through when it is off
// (documenting the BUG-N02 boundary: the guarantee covers built-in output).
func TestStrictASCII(t *testing.T) {
	nonASCII := func(lang.TransState) (string, int, bool) {
		return "é", 0, true
	}

	// Without strict mode a custom rule may emit non-ASCII.
	if got := New(WithRules(nonASCII)).Make("a"); got != "é" {
		t.Errorf("non-strict: got %q, want %q", got, "é")
	}

	// With strict mode the non-ASCII is stripped.
	if got := New(WithRules(nonASCII), WithStrictASCII()).Make("a"); got != "" {
		t.Errorf("strict rule: got %q, want empty", got)
	}

	// Strict mode also folds a non-ASCII fallback.
	strictFB := New(WithFallback(func(rune) string { return "★" }), WithStrictASCII())
	if got := strictFB.Make("😀"); got != "" {
		t.Errorf("strict fallback: got %q, want empty", got)
	}

	// Strict mode keeps the ASCII bytes of a mixed custom result.
	mixed := func(lang.TransState) (string, int, bool) { return "a-é-b", 0, true }
	if got := New(WithRules(mixed), WithStrictASCII()).Make("x"); got != "a--b" {
		t.Errorf("strict mixed: got %q, want %q", got, "a--b")
	}

	// ASCII custom output passes through unchanged under strict mode.
	ascii := func(lang.TransState) (string, int, bool) { return "OK", 0, true }
	if got := New(WithRules(ascii), WithStrictASCII()).Make("a"); got != "OK" {
		t.Errorf("strict ascii: got %q, want %q", got, "OK")
	}
}

// TestFeatureChaining checks the chainable Fallback and StrictASCII setters
// mutate and return the receiver.
func TestFeatureChaining(t *testing.T) {
	tr := New()
	if tr.Fallback(func(rune) string { return "?" }) != tr {
		t.Error("Fallback did not return the receiver")
	}
	if tr.StrictASCII() != tr {
		t.Error("StrictASCII did not return the receiver")
	}
	if got, want := tr.Make("😀"), "?"; got != want {
		t.Errorf("chained Make = %q, want %q", got, want)
	}
}

// TestRunesSeq checks that the streaming iterator reproduces the string form
// (including digraphs folded into a single step) and that breaking early stops
// the walk cleanly.
func TestRunesSeq(t *testing.T) {
	inputs := []string{"", "Hello", "Згурський", "世界!", "з'їв"}

	for _, in := range inputs {
		var b strings.Builder
		for _, v := range RunesSeq(lang.UK, in) {
			b.WriteString(v)
		}
		if got, want := b.String(), Trans(lang.UK, in); got != want {
			t.Errorf("RunesSeq(%q) = %q, want %q", in, got, want)
		}
	}

	// Breaking out of the range must stop the iterator without panicking.
	steps := 0
	for r := range RunesSeq(lang.UK, "abcdef") {
		steps++
		_ = r
		break
	}
	if steps != 1 {
		t.Errorf("early break produced %d steps, want 1", steps)
	}
}
