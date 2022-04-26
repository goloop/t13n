package languages

const (
	None             = ""
	Bulgarian        = "bg"
	Croatian         = "hr"
	Danish           = "da"
	Esperanto        = "eo"
	German           = "de"
	Hungarian        = "hu"
	Macedonian       = "mk"
	Norwegian        = "nb"
	Russian          = "ru"
	Serbian          = "sr"
	Slovenian        = "sl"
	Swedish          = "sv"
	Ukrainian        = "uk"
	Abkhazian        = "ab"
	Afar             = "aa"
	Afrikaans        = "af"
	Akan             = "ak"
	Albanian         = "sq"
	Amharic          = "am"
	Arabic           = "ar"
	Aragonese        = "an"
	Armenian         = "hy"
	Assamese         = "as"
	Avaric           = "av"
	Avestan          = "ae"
	Aymara           = "ay"
	Azerbaijani      = "az"
	Bambara          = "bm"
	Bashkir          = "ba"
	Basque           = "eu"
	Belarusian       = "be"
	Bengali          = "bn"
	Bislama          = "bi"
	Bosnian          = "bs"
	Breton           = "br"
	Burmese          = "my"
	Chamorro         = "ch"
	Chechen          = "ce"
	Nyanja           = "ny"
	Catalan          = "ca"
	Chinese          = "zh"
	ChurchSlavonic   = "cu"
	Chuvash          = "cv"
	Cornish          = "kw"
	Corsican         = "co"
	Cree             = "cr"
	Czech            = "cs"
	Divehi           = "dv"
	Dhivehi          = "dv"
	Maldivian        = "dv"
	Dutch            = "nl"
	Flemish          = "nl"
	Dzongkha         = "dz"
	English          = "en"
	Estonian         = "et"
	Ewe              = "ee"
	Faroese          = "fo"
	Fijian           = "fj"
	Finnish          = "fi"
	French           = "fr"
	WesternFrisian   = "fy"
	Fulah            = "ff"
	Gaelic           = "gd"
	Galician         = "gl"
	Ganda            = "lg"
	Georgian         = "ka"
	Greek            = "el"
	Kalaallisut      = "kl"
	Guarani          = "gn"
	Gujarati         = "gu"
	Haitian          = "ht"
	Hausa            = "ha"
	Hebrew           = "he"
	Herero           = "hz"
	Hindi            = "hi"
	Hiri             = "ho"
	Icelandic        = "is"
	Ido              = "io"
	Igbo             = "ig"
	Indonesian       = "id"
	Interlingua      = "ia"
	Interlingue      = "ie"
	Inuktitut        = "iu"
	Inupiaq          = "ik"
	Irish            = "ga"
	Italian          = "it"
	Japanese         = "ja"
	Javanese         = "jv"
	Kannada          = "kn"
	Kanuri           = "kr"
	Kashmiri         = "ks"
	Kazakh           = "kk"
	Khmer            = "km"
	Kikuyu           = "ki"
	Kinyarwanda      = "rw"
	Kirghiz          = "ky"
	Komi             = "kv"
	Kongo            = "kg"
	Korean           = "ko"
	Kuanyama         = "kj"
	Kurdish          = "ku"
	Lao              = "lo"
	Latin            = "la"
	Latvian          = "lv"
	Limburgan        = "li"
	Lingala          = "ln"
	Lithuanian       = "lt"
	LubaKatanga      = "lu"
	Luxembourgish    = "lb"
	Malagasy         = "mg"
	Malay            = "ms"
	Malayalam        = "ml"
	Maltese          = "mt"
	Manx             = "gv"
	Maori            = "mi"
	Marathi          = "mr"
	Marshallese      = "mh"
	Mongolian        = "mn"
	Nauru            = "na"
	Navajo           = "nv"
	NorthNdebele     = "nd"
	SouthNdebele     = "nr"
	Ndonga           = "ng"
	Nepali           = "ne"
	NorwegianBokmal  = "nb"
	NorwegianNynorsk = "nn"
	SichuanYi        = "ii"
	Occitan          = "oc"
	Ojibwa           = "oj"
	Oriya            = "or"
	Oromo            = "om"
	Ossetian         = "os"
	Pali             = "pi"
	Pashto           = "ps"
	Persian          = "fa"
	Polish           = "pl"
	Portuguese       = "pt"
	Punjabi          = "pa"
	Quechua          = "qu"
	Romanian         = "ro"
	Moldavian        = "ro"
	Moldovan         = "ro"
	Romansh          = "rm"
	Rundi            = "rn"
	NorthernSami     = "se"
	Samoan           = "sm"
	Sango            = "sg"
	Sanskrit         = "sa"
	Sardinian        = "sc"
	Shona            = "sn"
	Sindhi           = "sd"
	Sinhala          = "si"
	Slovak           = "sk"
	Somali           = "so"
	SouthernSotho    = "st"
	Spanish          = "es"
	Castilian        = "es"
	Sundanese        = "su"
	Swahili          = "sw"
	Swati            = "ss"
	Tagalog          = "tl"
	Tahitian         = "ty"
	Tajik            = "tg"
	Tamil            = "ta"
	Tatar            = "tt"
	Telugu           = "te"
	Thai             = "th"
	Tibetan          = "bo"
	Tigrinya         = "ti"
	Tonga            = "to"
	Tsonga           = "ts"
	Tswana           = "tn"
	Turkish          = "tr"
	Turkmen          = "tk"
	Twi              = "tw"
	Uighur           = "ug"
	Uyghur           = "ug"
	Urdu             = "ur"
	Uzbek            = "uz"
	Venda            = "ve"
	Vietnamese       = "vi"
	Volapük          = "vo"
	Walloon          = "wa"
	Welsh            = "cy"
	Wolof            = "wo"
	Xhosa            = "xh"
	Yiddish          = "yi"
	Yoruba           = "yo"
	Zhuang           = "zs"
	Chuang           = "za"
	Zulu             = "zu"
)

// LanguageRulesMethod a type of special function that can correct
// transliteration according to the specifics of the specified language.
type LanguageRulesMethod func(rune, rune, rune, bool) (string, int, bool)

// The defaultLanguageRules a function that uses general correction rules
// for the default language.
func defaultLanguageRules(p, c, n rune, b bool) (string, int, bool) {
	return "", 0, false
}

// LanguageRules returns the corresponding transliteration
// correction function for the specified language.
func LanguageRules(lang string) LanguageRulesMethod {
	var languageRules LanguageRulesMethod

	switch lang {
	case Bulgarian:
		languageRules = bulgarianRules
	case Croatian:
		languageRules = croatianRules
	case Danish:
		languageRules = danishRules
	case Esperanto:
		languageRules = esperantoRules
	case German:
		languageRules = germanRules
	case Hungarian:
		languageRules = hungarianRules
	case Macedonian:
		languageRules = macedonianRules
	case Norwegian:
		languageRules = norwegianRules
	case Russian:
		languageRules = russianRules
	case Serbian:
		languageRules = serbianRules
	case Slovenian:
		languageRules = slovenianRules
	case Swedish:
		languageRules = swedishRules
	case Ukrainian:
		languageRules = ukrainianRules
	default:
		languageRules = defaultLanguageRules
	}

	return languageRules
}
