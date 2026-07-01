package lang

// Macedonian (MK).
//
// The values follow the official Macedonian romanisation, which deliberately
// differs from the generic base table for a few letters: Х -> H (not "Kh"),
// Ц -> C (not "Ts") and Џ -> Dj. Those overrides are the point of this map;
// the comments give each letter's code point and glyph, not the base value.
var macedonian = map[int]string{
	1027: "Gj", // 1027, U+0403, 'Ѓ'
	1029: "Dz", // 1029, U+0405, 'Ѕ'
	1032: "J",  // 1032, U+0408, 'Ј'
	1033: "Lj", // 1033, U+0409, 'Љ'
	1034: "Nj", // 1034, U+040A, 'Њ'
	1036: "Kj", // 1036, U+040C, 'Ќ'
	1039: "Dj", // 1039, U+040F, 'Џ' (Macedonian: not the base "Dzh")
	1046: "Zh", // 1046, U+0416, 'Ж'
	1061: "H",  // 1061, U+0425, 'Х' (Macedonian: not the base "Kh")
	1062: "C",  // 1062, U+0426, 'Ц' (Macedonian: not the base "Ts")
	1063: "Ch", // 1063, U+0427, 'Ч'
	1064: "Sh", // 1064, U+0428, 'Ш'
	1093: "h",  // 1093, U+0445, 'х' (Macedonian: not the base "kh")
	1094: "c",  // 1094, U+0446, 'ц' (Macedonian: not the base "ts")
	1119: "dj", // 1119, U+045F, 'џ' (Macedonian: not the base "dzh")
}

// mkRules implements the rules of transliteration.
var mkRules = mapRules(macedonian)
