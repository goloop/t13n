package lang

// Slovenian (SL),
var slovenian = map[int]string{
	272: "Dj", // 272, U+0110, 'Đ', "D"
	273: "dj", // 273, U+0111, 'đ', "d"
}

// slRules implements the rules of transliteration.
var slRules = mapRules(slovenian)
