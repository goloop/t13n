package t13n

import "github.com/goloop/t13n/lang"

// The isDelimiter returns true if char is characters delimiter.
// Importantly:
// - the number is also a delimiter of characters;
// - the apostrophe is also a delimiter.
func isDelimiter(c rune) bool {
	var separators = [][]int{
		{9, 9},
		{32, 64},
		{91, 96},
		{123, 126},
		{8216, 8217},
	}

	id := int(c)
	for _, interval := range separators {
		if id >= interval[0] && id <= interval[1] {
			return true
		}
	}

	return false
}

// The isApostrophe returns true if char is apostrophe.
// The apostrophe is quote between the characters.
func isApostrophe(ts lang.TransState) bool {
	if ts.Prev == 0 || ts.Next == 0 {
		return false
	}

	id := int(ts.Curr)
	if id == 39 || id == 96 || id == 8217 || id == 8216 {
		if !isDelimiter(ts.Prev) && !isDelimiter(ts.Next) {
			return true
		}
	}

	return false
}
