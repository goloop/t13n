package languages

// Norwegian (NB)
var norwegian = map[int]string{
	216: "Oe", // 216, U+00D8, 'Ø', "O"
	248: "oe", // 248, U+00F8, 'ø', "o"
}

// The norwegianRules implements the rules of transliteration into Norwegian.
func norwegianRules(p, c, n rune, b bool) (string, int, bool) {
	result, id, seek, changed := "", int(c), 0, false
	if v, ok := norwegian[id]; ok {
		result = v
		changed = true
	}

	return result, seek, changed
}
