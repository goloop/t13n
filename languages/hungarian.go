package languages

// Hungarian (HU)
var hungarian = map[int]string{
	196: "Ae", // 196, U+00C4, 'Ä', "A"
	214: "Oe", // 214, U+00D6, 'Ö', "O"
	220: "Ue", // 220, U+00DC, 'Ü', "U"
	228: "ae", // 228, U+00E4, 'ä', "a"
	246: "oe", // 246, U+00F6, 'ö', "o"
	252: "ue", // 252, U+00FC, 'ü', "u"
	272: "Dz", // 272, U+0110, 'Đ', "D"
	273: "dz", // 273, U+0111, 'đ', "d"
	336: "Oe", // 336, U+0150, 'Ő', "O"
	337: "oe", // 337, U+0151, 'ő', "o"
	368: "Ue", // 368, U+0170, 'Ű', "U"
	369: "ue", // 369, U+0171, 'ű', "u"
}

// The hungarianRules implements the rules of transliteration into Hungarian.
func hungarianRules(p, c, n rune, b bool) (string, int, bool) {
	result, id, seek, changed := "", int(c), 0, false
	if v, ok := hungarian[id]; ok {
		result = v
		changed = true
	}

	return result, seek, changed
}
