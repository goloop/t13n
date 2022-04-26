package t13n

// The isCharDelimiter returns true if char is characters delimiter.
// Importantly:
// - the number is also a delimiter of characters;
// - the apostrophe is also a delimiter.
func isCharDelimiter(c rune) bool {
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
func isApostrophe(p, c, n rune) bool {
	var id = int(c)

	if p == 0 || n == 0 {
		return false
	}

	if id == 39 || id == 96 || id == 8217 || id == 8216 {
		if !isCharDelimiter(p) && !isCharDelimiter(n) {
			return true
		}
	}

	return false
}

// The valueByIndex returns value by index from main dictionary.
func valueByIndex(id int) string {
	switch {
	case id < 127:
		return string(rune(id))
	case id < 0 || id > len(dictionary)-1:
		return ""
	}

	return dictionary[id]
}
