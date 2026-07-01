package lang

// Russian (RU)
var russian = map[int]string{
	1025: "Jo",  // 1025, U+0401, 'Ё'
	1046: "Zh",  // 1046, U+0416, 'Ж'
	1049: "J",   // 1049, U+0419, 'Й'
	1063: "Ch",  // 1063, U+0427, 'Ч'
	1064: "Sh",  // 1064, U+0428, 'Ш'
	1065: "Shh", // 1065, U+0429, 'Щ'
	1069: "Eh",  // 1069, U+042D, 'Э'
	1070: "Ju",  // 1070, U+042E, 'Ю'
	1071: "Ja",  // 1071, U+042F, 'Я'
	1078: "zh",  // 1078, U+0436, 'ж'
	1081: "j",   // 1081, U+0439, 'й'
	1095: "ch",  // 1095, U+0447, 'ч'
	1096: "sh",  // 1096, U+0448, 'ш'
	1097: "shh", // 1097, U+0449, 'щ'
	1101: "eh",  // 1101, U+044D, 'э'
	1102: "ju",  // 1102, U+044E, 'ю'
	1103: "ja",  // 1103, U+044F, 'я'
	1105: "jo",  // 1105, U+0451, 'ё'
}

// ruRules implements the rules of transliteration into Russian.
func ruRules(ts TransState) (string, int, bool) {
	cid, nid := int(ts.Curr), int(ts.Next)

	switch {
	case cid == 1059 && (nid == 1049 || nid == 1081): // У + Й/й
		return "Uy", 1, true
	case cid == 1091 && (nid == 1049 || nid == 1081): // у + Й/й
		return "uy", 1, true
	default:
		if v, ok := russian[cid]; ok {
			return v, 0, true
		}
	}

	return "", 0, false
}
