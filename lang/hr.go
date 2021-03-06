package lang

// Croatian (HR)
var croatian = map[int]string{
	272: "Dj", // 272, U+0110, 'Đ', "D"
	273: "dj", // 273, U+0111, 'đ', "d"
}

// The hrRules implements the rules of transliteration into Croatian.
func hrRules(ts TransState) (string, int, bool) {
	result, id, offset, changed := "", int(ts.Curr), 0, false
	if v, ok := croatian[id]; ok {
		result = v
		changed = true
	}

	return result, offset, changed
}
