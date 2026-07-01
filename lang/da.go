package lang

// Danish (DA).
var danish = map[int]string{
	197: "Aa", // 197, U+00C5, 'Å', "A"
	216: "Oe", // 216, U+00D8, 'Ø', "O"
	229: "aa", // 229, U+00E5, 'å', "a"
	248: "oe", // 248, U+00F8, 'ø', "o"
}

// daRules implements the rules of transliteration.
var daRules = mapRules(danish)
