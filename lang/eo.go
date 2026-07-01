package lang

// Esperanto (EO), x-system.
var esperanto = map[int]string{
	264: "Cx", // 264, U+0108, 'Ĉ'
	265: "cx", // 265, U+0109, 'ĉ'
	284: "Gx", // 284, U+011C, 'Ĝ'
	285: "gx", // 285, U+011D, 'ĝ'
	292: "Hx", // 292, U+0124, 'Ĥ'
	293: "hx", // 293, U+0125, 'ĥ'
	308: "Jx", // 308, U+0134, 'Ĵ'
	309: "jx", // 309, U+0135, 'ĵ'
	348: "Sx", // 348, U+015C, 'Ŝ'
	349: "sx", // 349, U+015D, 'ŝ'
	364: "Ux", // 364, U+016C, 'Ŭ'
	365: "ux", // 365, U+016D, 'ŭ'
}

// eoRules implements the rules of transliteration into Esperanto.
var eoRules = mapRules(esperanto)
