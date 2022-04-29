package lang

// Bulgarian (BG).
var bulgarian = map[int]string{
	1049: "Y",   // 1049, U+0419, 'Й', "I"
	1081: "y",   // 1081, U+0439, 'й', "i"
	1065: "Sht", // 1065, U+0429, 'Щ', "Shch"
	1097: "sht", // 1097, U+0449, 'щ', "shch"
	1066: "A",   // 1066, U+042A, 'Ъ', "'"
	1098: "a",   // 1098, U+044A, 'ъ', "'"
	1068: "Y",   // 1068, U+042C, 'Ь', "'"
	1100: "y",   // 1100, U+044C, 'ь', "'"
	1070: "Yu",  // 1070, U+042E, 'Ю', "Iu"
	1102: "yu",  // 1102, U+044E, 'ю', "iu"
	1071: "ya",  // 1071, U+042F, 'Я', "Ia"
	1103: "Ya",  // 1103, U+044F, 'я', "ia"
}

// The bgRules implements the rules of transliteration into Bulgarian.
func bgRules(ts TransState) (string, int, bool) {
	result, id, seek, changed := "", int(ts.Curr), 0, false
	if v, ok := bulgarian[id]; ok {
		result = v
		changed = true
	}

	return result, seek, changed
}
