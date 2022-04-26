package languages

// Slovenian (SL),
var slovenian = map[int]string{
	272: "Dj", // 272, U+0110, 'Đ', "D"
	273: "dj", // 273, U+0111, 'đ', "d"
}

// The slovenianRules implements the rules of transliteration into Slovenian.
func slovenianRules(p, c, n rune, b bool) (string, int, bool) {
	result, id, seek, changed := "", int(c), 0, false
	if v, ok := slovenian[id]; ok {
		result = v
		changed = true
	}

	return result, seek, changed
}
