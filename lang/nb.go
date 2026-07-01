package lang

// Norwegian (NB)
var norwegian = map[int]string{
	216: "Oe", // 216, U+00D8, 'Ø', "O"
	248: "oe", // 248, U+00F8, 'ø', "o"
}

// nbRules implements the rules of transliteration.
var nbRules = mapRules(norwegian)
