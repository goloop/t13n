package lang

import "strings"

const (
	// None to ignore any language rules.
	None = ""
	// GB is British English language.
	GB = "gb"
	// US is American English language.
	US = "us"
	// BG is Bulgarian language.
	BG = "bg"
	// HR is Croatian language.
	HR = "hr"
	// DA is Danish language.
	DA = "da"
	// EO is Esperanto language.
	EO = "eo"
	// DE is German language.
	DE = "de"
	// HU is Hungarian language.
	HU = "hu"
	// MK is Macedonian language.
	MK = "mk"
	// NB is Norwegian language.
	NB = "nb"
	// RU is Russian language.
	RU = "ru"
	// SR is Serbian language.
	SR = "sr"
	// SL is Slovenian language.
	SL = "sl"
	// SV is Swedish language.
	SV = "sv"
	// UK is Ukrainian language.
	UK = "uk"
	// AB is Abkhazian language.
	AB = "ab"
	// AA is Afar language.
	AA = "aa"
	// AF is Afrikaans language.
	AF = "af"
	// AK is Akan language.
	AK = "ak"
	// SQ is Albanian language.
	SQ = "sq"
	// AM is Amharic language.
	AM = "am"
	// AR is Arabic language.
	AR = "ar"
	// AN is Aragonese language.
	AN = "an"
	// HY is Armenian language.
	HY = "hy"
	// AS is Assamese language.
	AS = "as"
	// AV is Avaric language.
	AV = "av"
	// AE is Avestan language.
	AE = "ae"
	// AY is Aymara language.
	AY = "ay"
	// AZ is Azerbaijani language.
	AZ = "az"
	// BM is Bambara language.
	BM = "bm"
	// BA is Bashkir language.
	BA = "ba"
	// EU is Basque language.
	EU = "eu"
	// BE is Belarusian language.
	BE = "be"
	// BN is Bengali language.
	BN = "bn"
	// BI is Bislama language.
	BI = "bi"
	// BS is Bosnian language.
	BS = "bs"
	// BR is Breton language.
	BR = "br"
	// MY is Burmese language.
	MY = "my"
	// CH is Chamorro language.
	CH = "ch"
	// CE is Chechen language.
	CE = "ce"
	// NY is Nyanja language.
	NY = "ny"
	// CA is Catalan language.
	CA = "ca"
	// ZH is Chinese language.
	ZH = "zh"
	// CU is Church Slavonic language.
	CU = "cu"
	// CV is Chuvash language.
	CV = "cv"
	// KW is Cornish language.
	KW = "kw"
	// CO is Corsican language.
	CO = "co"
	// CR is Cree language.
	CR = "cr"
	// CS is Czech language.
	CS = "cs"
	// DV is Divehi, Dhivehi, Maldivian  languages.
	DV = "dv"
	// NL is Dutch, Flemish  languages.
	NL = "nl"
	// DZ is Dzongkha language.
	DZ = "dz"
	// EN is English language.
	EN = "en"
	// ET is Estonian language.
	ET = "et"
	// EE is Ewe language.
	EE = "ee"
	// FO is Faroese language.
	FO = "fo"
	// FJ is Fijian language.
	FJ = "fj"
	// FI is Finnish language.
	FI = "fi"
	// FR is French language.
	FR = "fr"
	// FY is Western Frisian language.
	FY = "fy"
	// FF is Fulah language.
	FF = "ff"
	// GD is Gaelic language.
	GD = "gd"
	// GL is Galician language.
	GL = "gl"
	// LG is Ganda language.
	LG = "lg"
	// KA is Georgian language.
	KA = "ka"
	// EL is Greek language.
	EL = "el"
	// KL is Kalaallisut language.
	KL = "kl"
	// GN is Guarani language.
	GN = "gn"
	// GU is Gujarati language.
	GU = "gu"
	// HT is Haitian language.
	HT = "ht"
	// HA is Hausa language.
	HA = "ha"
	// HE is Hebrew language.
	HE = "he"
	// HZ is Herero language.
	HZ = "hz"
	// HI is Hindi language.
	HI = "hi"
	// HO is Hiri language.
	HO = "ho"
	// IS is Icelandic language.
	IS = "is"
	// IO is Ido language.
	IO = "io"
	// IG is Igbo language.
	IG = "ig"
	// ID is Indonesian language.
	ID = "id"
	// IA is Interlingua language.
	IA = "ia"
	// IE is Interlingue language.
	IE = "ie"
	// IU is Inuktitut language.
	IU = "iu"
	// IK is Inupiaq language.
	IK = "ik"
	// GA is Irish language.
	GA = "ga"
	// IT is Italian language.
	IT = "it"
	// JA is Japanese language.
	JA = "ja"
	// JV is Javanese language.
	JV = "jv"
	// KN is Kannada language.
	KN = "kn"
	// KR is Kanuri language.
	KR = "kr"
	// KS is Kashmiri language.
	KS = "ks"
	// KK is Kazakh language.
	KK = "kk"
	// KM is Khmer language.
	KM = "km"
	// KI is Kikuyu language.
	KI = "ki"
	// RW is Kinyarwanda language.
	RW = "rw"
	// KY is Kirghiz language.
	KY = "ky"
	// KV is Komi language.
	KV = "kv"
	// KG is Kongo language.
	KG = "kg"
	// KO is Korean language.
	KO = "ko"
	// KJ is Kuanyama language.
	KJ = "kj"
	// KU is Kurdish language.
	KU = "ku"
	// LO is Lao language.
	LO = "lo"
	// LA is Latin language.
	LA = "la"
	// LV is Latvian language.
	LV = "lv"
	// LI is Limburgan language.
	LI = "li"
	// LN is Lingala language.
	LN = "ln"
	// LT is Lithuanian language.
	LT = "lt"
	// LU is Luba Katanga language.
	LU = "lu"
	// LB is Luxembourgish language.
	LB = "lb"
	// MG is Malagasy language.
	MG = "mg"
	// MS is Malay language.
	MS = "ms"
	// ML is Malayalam language.
	ML = "ml"
	// MT is Maltese language.
	MT = "mt"
	// GV is Manx language.
	GV = "gv"
	// MI is Maori language.
	MI = "mi"
	// MR is Marathi language.
	MR = "mr"
	// MH is Marshallese language.
	MH = "mh"
	// MN is Mongolian language.
	MN = "mn"
	// NA is Nauru language.
	NA = "na"
	// NV is Navajo language.
	NV = "nv"
	// ND is NorthNdebele language.
	ND = "nd"
	// NR is South Ndebele language.
	NR = "nr"
	// NG is Ndonga language.
	NG = "ng"
	// NE is Nepali language.
	NE = "ne"
	// NN is Norwegian Nynorsk language.
	NN = "nn"
	// II is Sichuan Yi language.
	II = "ii"
	// OC is Occitan language.
	OC = "oc"
	// OJ is Ojibwa language.
	OJ = "oj"
	// OR is Oriya language.
	OR = "or"
	// OM is Oromo language.
	OM = "om"
	// OS is Ossetian language.
	OS = "os"
	// PI is Pali language.
	PI = "pi"
	// PS is Pashto language.
	PS = "ps"
	// FA is Persian language.
	FA = "fa"
	// PL is Polish language.
	PL = "pl"
	// PT is Portuguese language.
	PT = "pt"
	// PA is Punjabi language.
	PA = "pa"
	// QU is Quechua language.
	QU = "qu"
	// RO is Romanian, Moldavian, Moldovan languages.
	RO = "ro"
	// RM is Romansh language.
	RM = "rm"
	// RN is Rundi language.
	RN = "rn"
	// SE is Northern Sami language.
	SE = "se"
	// SM is Samoan language.
	SM = "sm"
	// SG is Sango language.
	SG = "sg"
	// SA is Sanskrit language.
	SA = "sa"
	// SC is Sardinian language.
	SC = "sc"
	// SN is Shona language.
	SN = "sn"
	// SD is Sindhi language.
	SD = "sd"
	// SI is Sinhala language.
	SI = "si"
	// SK is Slovak language.
	SK = "sk"
	// SO is Somali language.
	SO = "so"
	// ST is Southern Sotho language.
	ST = "st"
	// ES is Spanish, Castilian language.
	ES = "es"
	// SU is Sundanese language.
	SU = "su"
	// SW is Swahili language.
	SW = "sw"
	// SS is Swati language.
	SS = "ss"
	// TL is Tagalog language.
	TL = "tl"
	// TY is Tahitian language.
	TY = "ty"
	// TG is Tajik language.
	TG = "tg"
	// TA is Tamil language.
	TA = "ta"
	// TT is Tatar language.
	TT = "tt"
	// TE is Telugu language.
	TE = "te"
	// TH is Thai language.
	TH = "th"
	// BO is Tibetan language.
	BO = "bo"
	// TI is Tigrinya language.
	TI = "ti"
	// TO is Tonga language.
	TO = "to"
	// TS is Tsonga language.
	TS = "ts"
	// TN is Tswana language.
	TN = "tn"
	// TR is Turkish language.
	TR = "tr"
	// TK is Turkmen language.
	TK = "tk"
	// TW is Twi language.
	TW = "tw"
	// UG is Uighur, Uyghur languages.
	UG = "ug"
	// UR is Urdu language.
	UR = "ur"
	// UZ is Uzbek language.
	UZ = "uz"
	// VE is Venda language.
	VE = "ve"
	// VI is Vietnamese language.
	VI = "vi"
	// VO is Volapuk language.
	VO = "vo"
	// WA is Walloon language.
	WA = "wa"
	// CY is Welsh language.
	CY = "cy"
	// WO is Wolof language.
	WO = "wo"
	// XH is Xhosa language.
	XH = "xh"
	// YI is Yiddish language.
	YI = "yi"
	// YO is Yoruba language.
	YO = "yo"
	// ZS is Zhuang language.
	ZS = "zs"
	// ZA is Chuang language.
	ZA = "za"
	// ZU is Zulu language.
	ZU = "zu"
)

// TransState the transliteration state object of the current character.
type TransState struct {
	// Prev is previous character or rune(0).
	Prev rune

	// Curr is current character.
	Curr rune

	// Next is next character or rune(0).
	Next rune

	// Value is proposed translation.
	Value string

	// IsBegin is true if the Curr is at the beginning of the word.
	IsBegin bool

	// IsApostrophe is true if the Curr is apostrophe:
	// is symbol ` between two not delimiter characters.
	IsApostrophe bool
}

// TransRules a type of special function that can correct
// transliteration according to the specifics of the specified language.
type TransRules func(TransState) (string, int, bool)

// Rules returns the corresponding transliteration correction function
// for the specified language. Or returns nil if there is no such function.
func Rules(lang string) TransRules {
	var rules TransRules

	switch strings.ToLower(lang) {
	case BG:
		rules = bgRules
	case HR:
		rules = hrRules
	case DA:
		rules = daRules
	case EO:
		rules = eoRules
	case DE:
		rules = deRules
	case HU:
		rules = huRules
	case MK:
		rules = mkRules
	case NB:
		rules = nbRules
	case RU:
		rules = ruRules
	case SR:
		rules = srRules
	case SL:
		rules = slRules
	case SV:
		rules = svRules
	case UK:
		rules = ukRules
	default:
		rules = nil
	}

	return rules
}
