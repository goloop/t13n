package lang

// Ukrainian (UK)
var ukrainian = map[int]string{
	// [MIN VALUE: 1028]
	1028: "Ye", // 1028, U+0404, 'Є', "Ie" -> Ie
	1031: "Yi", // 1031, U+0407, 'Ї', "Yi" -> I
	1043: "H",  // 1043, U+0413, 'Г', "G"
	1048: "Y",  // 1048, U+0418, 'И', "I"
	1049: "Y",  // 1049, U+0419, 'Й', "I" -> I
	1066: "",   // 1066, U+042A, 'Ъ', "'"
	1068: "",   // 1068, U+042C, 'Ь', "'"
	1070: "Yu", // 1070, U+042E, 'Ю', "Iu" -> Iu
	1071: "Ya", // 1071, U+042F, 'Я', "Ia" -> Ia
	1075: "h",  // 1075, U+0433, 'г', "g"
	1080: "y",  // 1080, U+0438, 'и', "i"
	1081: "y",  // 1081, U+0439, 'й', "i" -> i
	1098: "",   // 1098, U+044A, 'ъ', "'"
	1100: "",   // 1100, U+044C, 'ь', "'"
	1102: "yu", // 1102, U+044E, 'ю', "iu" -> iu
	1103: "ya", // 1103, U+044F, 'я', "ia" -> ia
	1108: "ye", // 1108, U+0454, 'є', "ie" -> ie
	1111: "yi", // 1111, U+0457, 'ї', "yi" -> i
	1168: "G",  // 1168, U+0490, 'Ґ', "G'"
	1169: "g",  // 1169, U+0491, 'ґ', "g'"
	// [MAX VALUE: 1169]
}

// The ukRules implements the rules of transliteration into Ukrainian.
func ukRules(ts TransState) (string, int, bool) {
	var internal = map[int]string{
		1028: "Ie", // 1028, U+0404, 'Є', "Ye"
		1031: "I",  // 1031, U+0407, 'Ї', "Yi"
		1049: "I",  // 1049, U+0419, 'Й', "Y"
		1070: "Iu", // 1070, U+042E, 'Ю', "Yu"
		1071: "Ia", // 1071, U+042F, 'Я', "Ya"
		1081: "i",  // 1081, U+0439, 'й', "y"
		1102: "iu", // 1102, U+044E, 'ю', "yu"
		1103: "ia", // 1103, U+044F, 'я', "ya"
		1108: "ie", // 1108, U+0454, 'є', "ye"
		1111: "i",  // 1111, U+0457, 'ї', "yi"
	}

	if ts.IsApostrophe {
		return "", 0, true
	}

	result, cid, nid, offset, changed := "", int(ts.Curr), int(ts.Next), 0, false
	switch {
	case cid < 1028 && cid > 1169: // not ukrainian
		return result, offset, changed
	case cid == 1047 && (nid == 1043 || nid == 1075): // ЗГ || Зг
		changed = true
		result = "Zgh"
		offset = 1
	case cid == 1079 && (nid == 1043 || nid == 1075): // зГ || зг
		changed = true
		result = "zgh"
		offset = 1
	case cid == 1059 && (nid == 1049 || nid == 1081): // УЙ || Уй
		changed = true
		result = "Uy"
		offset = 1
	case cid == 1091 && (nid == 1049 || nid == 1081): // уЙ || уй
		changed = true
		result = "uy"
		offset = 1
	default:
		if v, ok := ukrainian[cid]; ok {
			if !ts.IsBegin {
				if w, ok := internal[cid]; ok {
					v = w
				}
			}

			result = v
			changed = true
		}
	}

	return result, offset, changed
}
