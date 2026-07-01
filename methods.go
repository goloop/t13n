package t13n

import (
	"iter"

	"github.com/goloop/t13n/v2/lang"
)

// version is the module version, reported by Version.
const version = "v2.1.0"

// Version returns the module version in the form
// "v{major}.{minor}.{patch}".
func Version() string {
	return version
}

// String returns the base ASCII transliteration of a single code point,
// ignoring any regional linguistic rules. It returns an empty string when
// the rune has no mapping (including runes outside the Basic Multilingual
// Plane and negative runes); use [Rune] to tell "no mapping" apart from a
// mapping that is intentionally empty.
func String(c rune) string {
	if id := int(c); id >= 0 && id < tableSize {
		return table()[id]
	}

	return ""
}

// Rune returns the base ASCII transliteration of a single code point and
// reports whether a mapping exists. ok is false for runes outside the table
// (negative or above U+FFFD) and for code points with no replacement, which
// lets callers detect characters that would otherwise silently vanish.
func Rune(c rune) (string, bool) {
	if id := int(c); id >= 0 && id < tableSize {
		if s := table()[id]; s != "" {
			return s, true
		}
	}

	return "", false
}

// Make transliterates a Unicode string to ASCII without applying any
// regional linguistic rules.
func Make(text string) string {
	return render(lang.None, text, nil, nil, false)
}

// Trans transliterates a Unicode string to ASCII, applying the regional
// rules of the given language (see the lang package for language codes).
func Trans(l, text string) string {
	return render(l, text, nil, nil, false)
}

// Render transliterates a Unicode string to ASCII, applying the regional
// rules of the given language and, when ctr is non-nil, a custom rule
// function applied last (for example, to build slugs).
func Render(l, text string, ctr lang.TransRules) string {
	return render(l, text, ctr, nil, false)
}

// RunesSeq returns an iterator over the transliteration of text with the
// regional rules of the given language. Each step yields the source rune that
// produced a unit of output together with that unit's ASCII value; runes
// consumed by a digraph rule are folded into the preceding step. It lets a
// caller stream the result without building the whole string at once.
func RunesSeq(l, text string) iter.Seq2[rune, string] {
	return func(yield func(rune, string) bool) {
		walk(l, text, nil, nil, yield)
	}
}
