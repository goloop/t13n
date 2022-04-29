package lang

// Slovenian (SL),
var slovenian = map[int]string{
	272: "Dj", // 272, U+0110, 'Đ', "D"
	273: "dj", // 273, U+0111, 'đ', "d"
}

// The slRules implements the rules of transliteration into Slovenian.
func slRules(ts TransState) (string, int, bool) {
	result, id, seek, changed := "", int(ts.Curr), 0, false
	if v, ok := slovenian[id]; ok {
		result = v
		changed = true
	}

	return result, seek, changed
}
