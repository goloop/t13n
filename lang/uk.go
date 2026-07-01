package lang

// Ukrainian (UK)
var ukrainian = map[int]string{
	// [MIN VALUE: 1028]
	1028: "Ye", // 1028, U+0404, 'Є'
	1031: "Yi", // 1031, U+0407, 'Ї'
	1043: "H",  // 1043, U+0413, 'Г'
	1048: "Y",  // 1048, U+0418, 'И'
	1049: "Y",  // 1049, U+0419, 'Й'
	1066: "",   // 1066, U+042A, 'Ъ'
	1068: "",   // 1068, U+042C, 'Ь'
	1070: "Yu", // 1070, U+042E, 'Ю'
	1071: "Ya", // 1071, U+042F, 'Я'
	1075: "h",  // 1075, U+0433, 'г'
	1080: "y",  // 1080, U+0438, 'и'
	1081: "y",  // 1081, U+0439, 'й'
	1098: "",   // 1098, U+044A, 'ъ'
	1100: "",   // 1100, U+044C, 'ь'
	1102: "yu", // 1102, U+044E, 'ю'
	1103: "ya", // 1103, U+044F, 'я'
	1108: "ye", // 1108, U+0454, 'є'
	1111: "yi", // 1111, U+0457, 'ї'
	1168: "G",  // 1168, U+0490, 'Ґ'
	1169: "g",  // 1169, U+0491, 'ґ'
	// [MAX VALUE: 1169]
}

// ukrainianInternal holds the in-word (non-initial) forms for the Ukrainian
// letters whose official romanization differs at the start of a word (for
// example, initial 'Є' is "Ye", but "ie" mid-word).
var ukrainianInternal = map[int]string{
	1028: "Ie", // 1028, U+0404, 'Є'
	1031: "I",  // 1031, U+0407, 'Ї'
	1049: "I",  // 1049, U+0419, 'Й'
	1070: "Iu", // 1070, U+042E, 'Ю'
	1071: "Ia", // 1071, U+042F, 'Я'
	1081: "i",  // 1081, U+0439, 'й'
	1102: "iu", // 1102, U+044E, 'ю'
	1103: "ia", // 1103, U+044F, 'я'
	1108: "ie", // 1108, U+0454, 'є'
	1111: "i",  // 1111, U+0457, 'ї'
}

// ukRules implements the rules of transliteration into Ukrainian.
func ukRules(ts TransState) (string, int, bool) {
	// An in-word apostrophe is dropped in Ukrainian romanization.
	if ts.IsApostrophe {
		return "", 0, true
	}

	cid, nid := int(ts.Curr), int(ts.Next)

	switch {
	case cid == 1047 && (nid == 1043 || nid == 1075): // З + Г/г -> Zgh
		return "Zgh", 1, true
	case cid == 1079 && (nid == 1043 || nid == 1075): // з + Г/г -> zgh
		return "zgh", 1, true
	case cid == 1059 && (nid == 1049 || nid == 1081): // У + Й/й -> Uy
		return "Uy", 1, true
	case cid == 1091 && (nid == 1049 || nid == 1081): // у + Й/й -> uy
		return "uy", 1, true
	default:
		if v, ok := ukrainian[cid]; ok {
			if !ts.IsBegin {
				if w, ok := ukrainianInternal[cid]; ok {
					v = w
				}
			}
			return v, 0, true
		}
	}

	return "", 0, false
}
