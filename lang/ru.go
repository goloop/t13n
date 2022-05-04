package lang

// Russian (RU)
var russian = map[int]string{
	1025: "Jo",  // 1025, U+0401, 'Ё', "Io"
	1046: "Zh",  // 1046, U+0416, 'Ж', "Zh"
	1049: "J",   // 1049, U+0419, 'Й', "I"
	1063: "Ch",  // 1063, U+0427, 'Ч', "Ch"
	1064: "Sh",  // 1064, U+0428, 'Ш', "Sh"
	1065: "Shh", // 1065, U+0429, 'Щ', "Shch"
	1069: "Eh",  // 1069, U+042D, 'Э', "E"
	1070: "Ju",  // 1070, U+042E, 'Ю', "Iu"
	1071: "Ja",  // 1071, U+042F, 'Я', "Ia"
	1078: "zh",  // 1078, U+0436, 'ж', "zh"
	1081: "j",   // 1081, U+0439, 'й', "i"
	1095: "ch",  // 1095, U+0447, 'ч', "ch"
	1096: "sh",  // 1096, U+0448, 'ш', "sh"
	1097: "shh", // 1097, U+0449, 'щ', "shch"
	1101: "eh",  // 1101, U+044D, 'э', "e"
	1102: "ju",  // 1102, U+044E, 'ю', "iu"
	1103: "ja",  // 1103, U+044F, 'я', "ia"
	1105: "jo",  // 1105, U+0451, 'ё', "io"
}

// The ruRules implements the rules of transliteration into Russian.
func ruRules(ts TransState) (string, int, bool) {
	result, id, offset, changed := "", int(ts.Curr), 0, false
	if v, ok := russian[id]; ok {
		result = v
		changed = true
	}

	result, cid, nid, offset, changed := "", int(ts.Curr), int(ts.Next), 0, false
	switch {
	case cid < 1025 && cid > 1105: // not ukrainian
		return result, offset, changed
	case cid == 1059 && (nid == 1049 || nid == 1081): // УЙ || Уй
		changed = true
		result = "Uy"
		offset = 1
	case cid == 1091 && (nid == 1049 || nid == 1081): // уЙ || уй
		changed = true
		result = "uy"
		offset = 1
	default:
		if v, ok := russian[id]; ok {
			result = v
			changed = true
		}
	}

	return result, offset, changed
}
