package lang

// Swedish (SV).
var swedish = map[int]string{
	196: "Ae", // 196, U+00C4, 'Ä', "A"
	214: "Oe", // 214, U+00D6, 'Ö', "O"
	228: "ae", // 228, U+00E4, 'ä', "a"
	246: "oe", // 246, U+00F6, 'ö', "o"
}

// The svRules implements the rules of transliteration into Swedish.
func svRules(ts TransState) (string, int, bool) {
	result, id, offset, changed := "", int(ts.Curr), 0, false
	if v, ok := swedish[id]; ok {
		result = v
		changed = true
	}

	return result, offset, changed
}
