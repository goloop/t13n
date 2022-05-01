package lang

// Esperanto (EO)
var esperanto = map[int]string{
	24:  "Cx", // 24, U+0018, '', ""
	25:  "cx", // 25, U+0019, '', ""
	284: "Gx", // 284, U+011C, 'Ĝ', "G"
	285: "gx", // 285, U+011D, 'ĝ', "g"
	292: "Hx", // 292, U+0124, 'Ĥ', "H"
	293: "hx", // 293, U+0125, 'ĥ', "h"
	308: "Jx", // 308, U+0134, 'Ĵ', "J"
	309: "jx", // 309, U+0135, 'ĵ', "j"
	348: "Sx", // 348, U+015C, 'Ŝ', "S"
	349: "sx", // 349, U+015D, 'ŝ', "s"
	364: "Ux", // 364, U+016C, 'Ŭ', "U"
	365: "ux", // 365, U+016D, 'ŭ', "u"
}

// The eoRules implements the rules of transliteration into Esperanto.
func eoRules(ts TransState) (string, int, bool) {
	result, id, offset, changed := "", int(ts.Curr), 0, false
	if v, ok := esperanto[id]; ok {
		result = v
		changed = true
	}

	return result, offset, changed
}
