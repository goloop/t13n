package lang

// Croatian (HR)
var croatian = map[int]string{
	272: "Dj", // 272, U+0110, 'Đ', "D"
	273: "dj", // 273, U+0111, 'đ', "d"
}

// hrRules implements the rules of transliteration.
var hrRules = mapRules(croatian)
