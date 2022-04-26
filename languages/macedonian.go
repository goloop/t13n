package languages

// Macedonian (MK).
var macedonian = map[int]string{
	1027: "Gj", // 1027, U+0403, 'Ѓ', "Gj"
	1029: "Dz", // 1029, U+0405, 'Ѕ', "Dz"
	1032: "J",  // 1032, U+0408, 'Ј', "J"
	1033: "Lj", // 1033, U+0409, 'Љ', "Lj"
	1034: "Nj", // 1034, U+040A, 'Њ', "Nj"
	1036: "Kj", // 1036, U+040C, 'Ќ', "Kj"
	1039: "Dj", // 1039, U+040F, 'Џ', "Dzh"
	1046: "Zh", // 1046, U+0416, 'Ж', "Zh"
	1061: "H",  // 1061, U+0425, 'Х', "Kh"
	1062: "C",  // 1062, U+0426, 'Ц', "Ts"
	1063: "Ch", // 1063, U+0427, 'Ч', "Ch"
	1064: "Sh", // 1064, U+0428, 'Ш', "Sh"
	1093: "h",  // 1093, U+0445, 'х', "kh"
	1094: "c",  // 1094, U+0446, 'ц', "ts"
	1119: "dj", // 1119, U+045F, 'џ', "dzh"
}

// The macedonianRules implements the rules of transliteration into Macedonian.
func macedonianRules(p, c, n rune, b bool) (string, int, bool) {
	result, id, seek, changed := "", int(c), 0, false
	if v, ok := macedonian[id]; ok {
		result = v
		changed = true
	}

	return result, seek, changed
}
