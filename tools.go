package t13n

import (
	"strings"

	"github.com/goloop/t13n/v2/lang"
)

// delimiterRanges lists the code point ranges treated as character
// delimiters, in ascending order. A delimiter separates words: besides
// obvious whitespace and punctuation, digits and apostrophe-like quotes
// count as delimiters. The ranges are a package-level value so the hot
// path does not rebuild them on every call.
var delimiterRanges = [...][2]rune{
	{9, 9},       // tab
	{32, 64},     // space, punctuation, digits, ':' ';' '<' '=' '>' '?' '@'
	{91, 96},     // '[' '\' ']' '^' '_' '`'
	{123, 126},   // '{' '|' '}' '~'
	{8216, 8217}, // left/right single quotation marks
}

// hieroglyphRanges lists the code point ranges treated as hieroglyphs
// (CJK ideographs and kana), in ascending order. See delimiterRanges for
// why this is a package-level value.
var hieroglyphRanges = [...][2]rune{
	{11904, 11929},
	{11931, 12019},
	{12032, 12245},
	{12293, 12295},
	{12321, 12329},
	{12344, 12347},
	{12353, 12438},
	{12445, 12447},
	{13312, 19903},
	{19968, 40956},
	{63744, 64109},
	{64112, 64217},
}

// inRanges reports whether c falls inside one of the given ascending,
// non-overlapping ranges.
func inRanges(c rune, ranges [][2]rune) bool {
	for _, r := range ranges {
		if c < r[0] {
			return false
		}
		if c <= r[1] {
			return true
		}
	}

	return false
}

// isDelimiter reports whether c separates characters. The NUL rune and every
// range in delimiterRanges (whitespace, punctuation, digits, apostrophes)
// count as delimiters.
func isDelimiter(c rune) bool {
	if c == 0 {
		return true
	}

	return inRanges(c, delimiterRanges[:])
}

// isHieroglyph reports whether c is a CJK ideograph or kana character.
func isHieroglyph(c rune) bool {
	return inRanges(c, hieroglyphRanges[:])
}

// isApostrophe reports whether the current character is an apostrophe: a
// quote-like rune standing between two non-delimiter characters (as in
// "d'Artagnan"), as opposed to an opening or closing quotation mark.
func isApostrophe(ts lang.TransState) bool {
	// With no neighbour on either side the rune cannot be an in-word
	// apostrophe.
	if ts.Prev == 0 || ts.Next == 0 {
		return false
	}

	switch ts.Curr {
	case '\'', '`', 8216, 8217:
		return !isDelimiter(ts.Prev) && !isDelimiter(ts.Next)
	}

	return false
}

// walk transliterates text to ASCII in a single left-to-right pass, applying
// the base table, then the language rules ltr, an optional fallback fb for
// runes that would otherwise vanish, and finally the custom rules ctr. For
// each emitted unit it calls emit with the source rune that produced the value
// and the value itself; emit returns false to stop the walk early (used by the
// RunesSeq iterator). A rule may consume trailing runes by returning a positive
// offset (for digraphs such as "зг" -> "zgh"). The pass is deliberately
// sequential: transliteration is CPU-light, and a single pass avoids both the
// quadratic concatenation of the previous implementation and the complexity
// (and data races) of stitching parallel chunks back together.
func walk(
	l, text string,
	ctr lang.TransRules,
	fb func(rune) string,
	emit func(rune, string) bool,
) {
	if text == "" {
		return
	}

	runes := []rune(text)
	ltr := lang.Rules(l)
	tbl := table()

	isBegin := true
	for i := 0; i < len(runes); i++ {
		ts := lang.TransState{Curr: runes[i], IsBegin: isBegin}
		if i > 0 {
			ts.Prev = runes[i-1]
		}
		if i < len(runes)-1 {
			ts.Next = runes[i+1]
		}

		// Base transliteration from the table. mapped tracks whether any
		// stage produced a real mapping, so the fallback fires only for
		// runes that would otherwise silently vanish.
		mapped := false
		if id := int(ts.Curr); id >= 0 && id < tableSize {
			ts.Value = tbl[id]
			mapped = ts.Value != ""
		}
		ts.IsApostrophe = isApostrophe(ts)

		// Regional language rules may override the value and consume
		// following runes.
		offset := 0
		if ltr != nil {
			if v, m, ok := ltr(ts); ok {
				ts.Value = v
				offset = m
				mapped = true
			}
		}

		// A hieroglyph carries a trailing space so that "世界" reads as
		// "Shi Jie"; drop it when the next character continues the word.
		if isHieroglyph(ts.Curr) {
			if ts.Next == 0 || isDelimiter(ts.Next) {
				ts.Value = strings.TrimRight(ts.Value, " ")
			}
		}

		// Fallback fills in a value for a rune with no mapping (unknown or
		// outside the Basic Multilingual Plane) so it need not vanish.
		if !mapped && fb != nil {
			ts.Value = fb(ts.Curr)
		}

		// Custom rules run last and win over the language rules.
		if ctr != nil {
			if v, m, ok := ctr(ts); ok {
				ts.Value = v
				offset = m
			}
		}

		if !emit(ts.Curr, ts.Value) {
			return
		}
		i += offset

		// The word boundary is defined by the character we just handled,
		// ignoring in-word apostrophes.
		isBegin = isDelimiter(ts.Curr) && !ts.IsApostrophe
	}
}

// render runs walk and assembles the transliterated string. When strict is
// true, any non-ASCII produced by a custom rule or fallback is stripped, so
// the "pure 7-bit ASCII" guarantee holds even for custom output.
func render(l, text string, ctr lang.TransRules, fb func(rune) string, strict bool) string {
	if text == "" {
		return ""
	}

	var b strings.Builder
	b.Grow(len(text))

	if strict {
		walk(l, text, ctr, fb, func(_ rune, v string) bool {
			b.WriteString(asciiOnly(v))
			return true
		})
	} else {
		walk(l, text, ctr, fb, func(_ rune, v string) bool {
			b.WriteString(v)
			return true
		})
	}

	return b.String()
}

// asciiOnly returns s with every non-ASCII byte removed. The common case, an
// already-ASCII string, is returned unchanged without allocating.
func asciiOnly(s string) string {
	ascii := true
	for i := 0; i < len(s); i++ {
		if s[i] >= 0x80 {
			ascii = false
			break
		}
	}
	if ascii {
		return s
	}

	b := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] < 0x80 {
			b = append(b, s[i])
		}
	}

	return string(b)
}
