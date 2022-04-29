package lang

// Serbian (SR)
var serbian = map[int]string{
	272:  "Dj",  // 272, U+0110, 'Đ', "D"
	273:  "dj",  // 273, U+0111, 'đ', "d"
	1026: "Dje", // 1026, U+0402, 'Ђ', "Dje"
	1106: "dje", // 1106, U+0452, 'ђ', "dje"
}

// The srRules implements the rules of transliteration into Serbian.
func srRules(ts TransState) (string, int, bool) {
	result, id, seek, changed := "", int(ts.Curr), 0, false
	if v, ok := serbian[id]; ok {
		result = v
		changed = true
	}

	return result, seek, changed
}
