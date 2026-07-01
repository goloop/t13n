package t13n

import (
	"testing"

	"github.com/goloop/t13n/v2/lang"
)

// TestIsDelimiter tests isDelimiter.
func TestIsDelimiter(t *testing.T) {
	tests := []struct {
		value string
		total int
	}{
		{string(rune(0)), 1}, // NUL is a delimiter too
		{"\tHello world!", 3},
		{"Hello, how are you?", 5},
		{" !\"#$%&'()*+,-./0123456789:;<=>?@", 33},
		{"[\\]^_`", 6},
		{"{|}~", 4},
		{"it’s", 1},
		{"it‘s", 1},
	}

	for _, test := range tests {
		total := 0
		for _, c := range test.value {
			if isDelimiter(c) {
				total++
			}
		}

		if total != test.total {
			t.Errorf("for %q: got %d delimiters, want %d",
				test.value, total, test.total)
		}
	}
}

// TestInRangesBoundaries pins the inclusive/exclusive edges of the range
// scanner: exactly the endpoints are inside, one below/above is outside, and
// the gaps between ranges are excluded.
func TestInRangesBoundaries(t *testing.T) {
	// Two disjoint ranges with a gap: [10,12] and [20,20].
	ranges := [][2]rune{{10, 12}, {20, 20}}
	cases := []struct {
		c    rune
		want bool
	}{
		{9, false}, {10, true}, {11, true}, {12, true}, {13, false},
		{19, false}, {20, true}, {21, false},
	}
	for _, c := range cases {
		if got := inRanges(c.c, ranges); got != c.want {
			t.Errorf("inRanges(%d): got %v want %v", c.c, got, c.want)
		}
	}
}

// TestIsHieroglyphBoundaries checks the first and last code point of a couple
// of ranges plus the surrounding gaps.
func TestIsHieroglyphBoundaries(t *testing.T) {
	cases := []struct {
		c    rune
		want bool
	}{
		{11903, false}, {11904, true}, {11929, true}, {11930, false},
		{19968, true}, {40956, true}, {40957, false},
		{'A', false}, {'世', true}, {'あ', true},
	}
	for _, c := range cases {
		if got := isHieroglyph(c.c); got != c.want {
			t.Errorf("isHieroglyph(%d): got %v want %v", c.c, got, c.want)
		}
	}
}

// TestIsApostrophe tests isApostrophe: a quote between two letters is an
// apostrophe, a quote at a word edge is not.
func TestIsApostrophe(t *testing.T) {
	tests := []struct {
		value string
		total int
	}{
		{"п‘ять, торф'яний, здоров`я, м’ясо, зв’язок", 5},
		{"сім’я, бур'ян, кур’єр, подвір‘я, під’їхати", 5},
		{"з’єднати, з’їхати, роз‘яснити", 3},
		{"‘hello’ 'world'", 0},
	}

	for _, test := range tests {
		total := 0
		runes := []rune(test.value)
		for i := 0; i < len(runes); i++ {
			ts := lang.TransState{Prev: rune(0), Curr: runes[i], Next: rune(0)}
			if i > 0 {
				ts.Prev = runes[i-1]
			}
			if i < len(runes)-1 {
				ts.Next = runes[i+1]
			}

			if isApostrophe(ts) {
				total++
			}
		}

		if total != test.total {
			t.Errorf("for %q: got %d apostrophes, want %d",
				test.value, total, test.total)
		}
	}
}
