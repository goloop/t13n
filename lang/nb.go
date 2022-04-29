package lang

// Norwegian (NB)
var norwegian = map[int]string{
	216: "Oe", // 216, U+00D8, 'Ø', "O"
	248: "oe", // 248, U+00F8, 'ø', "o"
}

// The nbRules implements the rules of transliteration into Norwegian.
func nbRules(ts TransState) (string, int, bool) {
	result, id, seek, changed := "", int(ts.Curr), 0, false
	if v, ok := norwegian[id]; ok {
		result = v
		changed = true
	}

	return result, seek, changed
}
