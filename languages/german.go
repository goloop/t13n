package languages

// German (DE)
var german = map[int]string{
	196:  "Ae", // 196, U+00C4, 'Ä', "A"
	214:  "Oe", // 214, U+00D6, 'Ö', "O"
	220:  "Ue", // 220, U+00DC, 'Ü', "U"
	223:  "ss", // 223, U+00DF, 'ß', "ss"
	228:  "ae", // 228, U+00E4, 'ä', "a"
	246:  "oe", // 246, U+00F6, 'ö', "o"
	252:  "ue", // 252, U+00FC, 'ü', "u"
	7838: "Ss", // 7838, U+1E9E, 'ẞ', "Ss"
}

// The germanRules implements the rules of transliteration into German.
func germanRules(p, c, n rune, b bool) (string, int, bool) {
	result, id, seek, changed := "", int(c), 0, false
	if v, ok := german[id]; ok {
		result = v
		changed = true
	}

	return result, seek, changed
}
