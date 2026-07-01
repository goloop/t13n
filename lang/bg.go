package lang

// Bulgarian (BG).
var bulgarian = map[int]string{
	1049: "Y",   // 1049, U+0419, 'Й'
	1081: "y",   // 1081, U+0439, 'й'
	1065: "Sht", // 1065, U+0429, 'Щ'
	1097: "sht", // 1097, U+0449, 'щ'
	1066: "A",   // 1066, U+042A, 'Ъ'
	1098: "a",   // 1098, U+044A, 'ъ'
	1068: "Y",   // 1068, U+042C, 'Ь'
	1100: "y",   // 1100, U+044C, 'ь'
	1070: "Yu",  // 1070, U+042E, 'Ю'
	1102: "yu",  // 1102, U+044E, 'ю'
	1071: "Ya",  // 1071, U+042F, 'Я' (capital -> capital)
	1103: "ya",  // 1103, U+044F, 'я' (small -> small)
}

// bgRules implements the rules of transliteration into Bulgarian.
var bgRules = mapRules(bulgarian)
