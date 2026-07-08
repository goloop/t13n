package t13n

import (
	"strings"
	"testing"
	"time"

	"github.com/goloop/t13n/v2/lang"
)

// TestMake tests Make against a broad range of scripts.
func TestMake(t *testing.T) {
	tests := []struct {
		value    string
		expected string
	}{
		{"Bzia zbaşa", "Bzia zbasa"},
		{"Фэсапщы", "Fesapshchy"},
		{"Salam əleyküm", "Salam aleykum"},
		{"Qysh je", "Qysh je"},
		{"Ç’kemi", "C'kemi"},
		{"ሰላም።", "salaame."},
		{"السلام عليكم", "lslm `lykm"},
		{"Héébee", "Heebee"},
		{"নমস্কাৰ", "nmskaar'"},
		{"сәләм", "salam"},
		{"Helô", "Helo"},
		{"Chào chị", "Chao chi"},
		{"Γειά σου", "Geia sou"},
		{"Bouônjour", "Bouonjour"},
		{"Güata Tàg", "Guata Tag"},
		{"Híu!", "Hiu!"},
		{"ᐊᐃᓐᖓᐃ", "ainngai"},
		{"Sæll", "Saell"},
		{"ສະບາຍດີ", "sabaanydii"},
		{"Hello 世界", "Hello Shi Jie"},
		{"こんにちは、みんな", "Ko N Ni Chi Ha Mi N Na"},
		{"", ""},
	}

	for _, test := range tests {
		if v := Make(test.value); v != test.expected {
			t.Errorf("Make(%q): got %q want %q", test.value, v, test.expected)
		}
	}
}

// TestRenderSlug tests Render with a custom rule that turns text into a slug.
func TestRenderSlug(t *testing.T) {
	tests := []struct {
		value    string
		expected string
	}{
		{"Bzia zbaşa", "bzia-zbasa"},
		{"Фэсапщы", "fesapshchy"},
		{"Salam əleyküm", "salam-aleykum"},
		{"Qysh je", "qysh-je"},
		{"Ç’kemi", "ckemi"},
		{"ሰላም።", "salaame"},
		{"السلام عليكم", "lslm-lykm"},
		{"Héébee", "heebee"},
		{"নমস্কাৰ", "nmskaar'"},
		{"сәләм", "salam"},
		{"Helô", "helo"},
		{"Chào chị", "chao-chi"},
		{"Γειά σου", "geia-sou"},
		{"Bouônjour", "bouonjour"},
		{"Güata Tàg", "guata-tag"},
		{"Híu!", "hiu"},
		{"ᐊᐃᓐᖓᐃ", "ainngai"},
		{"Sæll", "saell"},
		{"ສະບາຍດີ", "sabaanydii"},
		{"Hello 世界", "hello-shi-jie"},
	}

	slug := func(ts lang.TransState) (string, int, bool) {
		switch ts.Value {
		case " ", "~", "_":
			return "-", 0, true
		case "!", "@", "#", "$", "%", "`", "\"", "'", ".":
			fallthrough
		case "^", "&", "*", "(", ")", "+", "<", ">", "?":
			return "", 0, true
		}

		if strings.HasSuffix(ts.Value, " ") {
			runes := []rune(ts.Value)
			if ts.Next != 0 {
				ts.Value = string(runes[:len(runes)-1]) + "-"
			}
		}

		return strings.ToLower(ts.Value), 0, true
	}

	for _, test := range tests {
		if v := Render(lang.None, test.value, slug); v != test.expected {
			t.Errorf("Render(slug, %q): got %q want %q",
				test.value, v, test.expected)
		}
	}
}

// TestRenderSlugDigraph pins that a custom rule returning offset 0 does not
// undo the consumption of a regional digraph: the rune consumed by the digraph
// must not be transliterated a second time.
func TestRenderSlugDigraph(t *testing.T) {
	slug := func(ts lang.TransState) (string, int, bool) {
		if ts.Value == " " {
			return "-", 0, true
		}
		return strings.ToLower(ts.Value), 0, true
	}

	tests := []struct {
		value, expected string
	}{
		{"Згадка", "zghadka"},
		{"Гуйва", "huyva"},
		{"зг уй", "zgh-uy"},
	}
	for _, test := range tests {
		if v := Render(lang.UK, test.value, slug); v != test.expected {
			t.Errorf("Render(UK, slug, %q): got %q want %q",
				test.value, v, test.expected)
		}
	}

	// A custom rule may extend consumption beyond the regional rule by
	// returning a larger offset: here 'x' swallows the following rune.
	consume := func(ts lang.TransState) (string, int, bool) {
		if ts.Curr == 'x' {
			return "X", 1, true
		}
		return ts.Value, 0, true
	}
	if v := Render(lang.None, "xyz", consume); v != "Xz" {
		t.Errorf("custom rule extending offset: got %q want %q", v, "Xz")
	}
}

// TestWhitespaceWordBoundary checks that line breaks and a non-breaking space
// act as word separators (so word-initial regional forms are chosen) and that
// a non-breaking space is rendered as a plain space instead of vanishing.
func TestWhitespaceWordBoundary(t *testing.T) {
	tests := []struct {
		value, expected string
	}{
		{"хата\nЄвропа", "khata\nYevropa"},
		{"хата\rЄвропа", "khata\rYevropa"},
		{"хата Європа", "khata Yevropa"}, // NBSP between the words
		{"хата Європа", "khata Yevropa"}, // plain space, unchanged
	}
	for _, test := range tests {
		if v := Trans(lang.UK, test.value); v != test.expected {
			t.Errorf("Trans(UK, %q): got %q want %q",
				test.value, v, test.expected)
		}
	}
}

// TestNegativeOffsetNoHang makes sure a custom rule that returns a negative
// offset cannot rewind the walk into an endless loop.
func TestNegativeOffsetNoHang(t *testing.T) {
	bad := func(lang.TransState) (string, int, bool) { return "x", -1, true }

	done := make(chan string, 1)
	go func() { done <- Render(lang.None, "abc", bad) }()
	select {
	case got := <-done:
		if got != "xxx" {
			t.Errorf("negative offset: got %q, want %q", got, "xxx")
		}
	case <-time.After(5 * time.Second):
		t.Fatal("Render with a negative custom offset did not terminate")
	}
}

// TestBosnian tests Trans for Bosnian, which has no regional rules and so
// relies purely on the base table.
func TestBosnian(t *testing.T) {
	tests := []struct {
		value    string
		expected string
	}{
		{"Bihać", "Bihac"},
		{"Čapljina", "Capljina"},
		{"Goražde", "Gorazde"},
		{"Široki Brijeg", "Siroki Brijeg"},
		{"Živinice", "Zivinice"},
		// Adjacent uppercase multi-letter expansions stay title-cased so
		// letter boundaries remain unambiguous ("DjNj", not "DJNJ").
		{"ЂЊЋŽĆČŠ", "DjNjTshZCCS"},
		{
			"Ђ Е Ж З И Ј К Л Љ М Н Њ О П Р С Т Ћ У Ф Х Ц Ч Џ Ш",
			"Dj E Zh Z I J K L Lj M N Nj O P R S T Tsh U F Kh Ts Ch Dzh Sh",
		},
		{
			"ђ е ж з и ј к л љ м н њ о п р с т ћ у ф х ц ч џ ш",
			"dj e zh z i j k l lj m n nj o p r s t tsh u f kh ts ch dzh sh",
		},
		{
			"A D Đ E Ž Z I J T Ć U F H C Č Š",
			"A D D E Z Z I J T C U F H C C S",
		},
		{
			"a d đ e ž z i t ć u f h c č š",
			"a d d e z z i t c u f h c c s",
		},
	}

	for _, test := range tests {
		if v := Trans(lang.BS, test.value); v != test.expected {
			t.Errorf("Trans(BS, %q): got %q want %q",
				test.value, v, test.expected)
		}
	}
}

// TestBulgarian tests Trans for Bulgarian, including the capital/small Я
// mapping that used to be inverted.
func TestBulgarian(t *testing.T) {
	tests := []struct {
		value    string
		expected string
	}{
		{"Č Ć É F Ô Š Ž", "C C E F O S Z"},
		{"č ć é f ô š ž", "c c e f o s z"},
		{"Ќ Ѣ Џ Њ Ъ", "Kj E Dzh Nj A"},
		{"ќ ѣ џ њ ъ", "kj e dzh nj a"},
		// Capital Я -> "Ya", small я -> "ya" (must not be swapped).
		{"Я", "Ya"},
		{"я", "ya"},
		{"Яя", "Yaya"},
	}

	for _, test := range tests {
		if v := Trans(lang.BG, test.value); v != test.expected {
			t.Errorf("Trans(BG, %q): got %q want %q",
				test.value, v, test.expected)
		}
	}
}

// TestCatalan tests Trans for Catalan.
func TestCatalan(t *testing.T) {
	value := "À à É é È è Í í Ï ï Ó ó Ò ò Ú ú Ü ü Ç ç"
	want := "A a E e E e I i I i O o O o U u U u C c"
	if v := Trans(lang.CA, value); v != want {
		t.Errorf("Trans(CA): got %q want %q", v, want)
	}
}

// TestCroatian tests Trans for Croatian.
func TestCroatian(t *testing.T) {
	tests := []struct {
		value    string
		expected string
	}{
		{"Č Ć DŽ Đ LJ NJ Š Ž", "C C DZ Dj LJ NJ S Z"},
		{"č ć dž đ lj nj š ž", "c c dz dj lj nj s z"},
	}

	for _, test := range tests {
		if v := Trans(lang.HR, test.value); v != test.expected {
			t.Errorf("Trans(HR, %q): got %q want %q",
				test.value, v, test.expected)
		}
	}
}

// TestDanish tests Trans for Danish.
func TestDanish(t *testing.T) {
	value, want := "Æ Ø Å æ ø å", "AE Oe Aa ae oe aa"
	if v := Trans(lang.DA, value); v != want {
		t.Errorf("Trans(DA): got %q want %q", v, want)
	}
}

// TestEsperanto tests Trans for Esperanto: every accented letter, including Ĉ,
// must use the x-system (Ĉ -> "Cx"), which the dead 24/25 entries broke before.
func TestEsperanto(t *testing.T) {
	tests := []struct {
		value    string
		expected string
	}{
		{"Ĉ Ĝ Ĥ Ĵ Ŝ Ŭ", "Cx Gx Hx Jx Sx Ux"},
		{"ĉ ĝ ĥ ĵ ŝ ŭ", "cx gx hx jx sx ux"},
	}

	for _, test := range tests {
		if v := Trans(lang.EO, test.value); v != test.expected {
			t.Errorf("Trans(EO, %q): got %q want %q",
				test.value, v, test.expected)
		}
	}
}

// TestGerman tests Trans for German, including the all-caps title-case
// behaviour ("ÄÖÜ" -> "AeOeUe", not "AEOEUE").
func TestGerman(t *testing.T) {
	tests := []struct {
		value    string
		expected string
	}{
		{"Ä Ö Ü ẞ", "Ae Oe Ue Ss"},
		{"ä ö ü ß", "ae oe ue ss"},
		{"ÄÖÜ", "AeOeUe"},
		{"Müller", "Mueller"},
	}

	for _, test := range tests {
		if v := Trans(lang.DE, test.value); v != test.expected {
			t.Errorf("Trans(DE, %q): got %q want %q",
				test.value, v, test.expected)
		}
	}
}

// TestHungarian tests Trans for Hungarian.
func TestHungarian(t *testing.T) {
	tests := []struct {
		value    string
		expected string
	}{
		{"Á É Í Ó Ö Ő Ú Ü Ű", "A E I O Oe Oe U Ue Ue"},
		{"á é í ó ö ő ú ü ű", "a e i o oe oe u ue ue"},
	}

	for _, test := range tests {
		if v := Trans(lang.HU, test.value); v != test.expected {
			t.Errorf("Trans(HU, %q): got %q want %q",
				test.value, v, test.expected)
		}
	}
}

// TestMacedonian tests Trans for Macedonian.
func TestMacedonian(t *testing.T) {
	tests := []struct {
		value    string
		expected string
	}{
		{"Ѓ Ќ Џ Љ Њ", "Gj Kj Dj Lj Nj"},
		{"ѓ ќ џ љ њ", "gj kj dj lj nj"},
	}

	for _, test := range tests {
		if v := Trans(lang.MK, test.value); v != test.expected {
			t.Errorf("Trans(MK, %q): got %q want %q",
				test.value, v, test.expected)
		}
	}
}

// TestNorwegian tests Trans for Norwegian.
func TestNorwegian(t *testing.T) {
	tests := []struct {
		value    string
		expected string
	}{
		{"Æ Ø Å Ð Ô Ê Å Ä Ö", "AE Oe A D O E A A O"},
		{"æ ø å ð ô ê å ä ö", "ae oe a d o e a a o"},
	}

	for _, test := range tests {
		if v := Trans(lang.NB, test.value); v != test.expected {
			t.Errorf("Trans(NB, %q): got %q want %q",
				test.value, v, test.expected)
		}
	}
}

// TestRussian tests Trans for Russian.
func TestRussian(t *testing.T) {
	tests := []struct {
		value    string
		expected string
	}{
		{"Ъ Ы Ћ Ѱ Ѳ Ѵ", "' Y Tsh Ps F Y"},
		{"ъ ы ћ ѱ ѳ ѵ", "' y tsh ps f y"},
		{"Путин - Хуйло", "Putin - Khuylo"},
	}

	for _, test := range tests {
		if v := Trans(lang.RU, test.value); v != test.expected {
			t.Errorf("Trans(RU, %q): got %q want %q",
				test.value, v, test.expected)
		}
	}
}

// TestSerbian tests Trans for Serbian.
func TestSerbian(t *testing.T) {
	value, want := "Đ đ Ђ ђ", "Dj dj Dje dje"
	if v := Trans(lang.SR, value); v != want {
		t.Errorf("Trans(SR): got %q want %q", v, want)
	}
}

// TestSlovenian tests Trans for Slovenian.
func TestSlovenian(t *testing.T) {
	tests := []struct {
		value    string
		expected string
	}{
		{"Đ đ Ђ ђ", "Dj dj Dj dj"},
		{
			"Ä, Å, Æ, Ç, Ë, Ï, Ń, Ö, SS, Ş, Ü",
			"A, A, AE, C, E, I, N, O, SS, S, U",
		},
		{
			"ä, å, æ, ç, ë, ï, ń, ö, ß, ş, ü",
			"a, a, ae, c, e, i, n, o, ss, s, u",
		},
	}

	for _, test := range tests {
		if v := Trans(lang.SL, test.value); v != test.expected {
			t.Errorf("Trans(SL, %q): got %q want %q",
				test.value, v, test.expected)
		}
	}
}

// TestSwedish tests Trans for Swedish.
func TestSwedish(t *testing.T) {
	tests := []struct {
		value    string
		expected string
	}{
		{"Å Ä Ö", "A Ae Oe"},
		{"å ä ö", "a ae oe"},
	}

	for _, test := range tests {
		if v := Trans(lang.SV, test.value); v != test.expected {
			t.Errorf("Trans(SV, %q): got %q want %q",
				test.value, v, test.expected)
		}
	}
}

// TestUkrainian tests Trans for Ukrainian against the official romanization of
// a large set of place names, covering initial vs in-word forms, зг -> zgh,
// apostrophes and the soft sign.
func TestUkrainian(t *testing.T) {
	tests := []struct {
		value    string
		expected string
	}{
		{"Алушта", "Alushta"},
		{"Андрій", "Andrii"},
		{"Борщагівка", "Borshchahivka"},
		{"Борисенко", "Borysenko"},
		{"Вінниця", "Vinnytsia"},
		{"Гадяч", "Hadiach"},
		{"Богдан", "Bohdan"},
		{"Згурський", "Zghurskyi"},
		{"Ґалаґан", "Galagan"},
		{"Ґорґани", "Gorgany"},
		{"Донецьк", "Donetsk"},
		{"Дмитро", "Dmytro"},
		{"Рівне", "Rivne"},
		{"Олег", "Oleh"},
		{"Есмань", "Esman"},
		{"Єнакієве", "Yenakiieve"},
		{"Гаєвич", "Haievych"},
		{"Короп’є", "Koropie"},
		{"Житомир", "Zhytomyr"},
		{"Жанна", "Zhanna"},
		{"Жежелів", "Zhezheliv"},
		{"Закарпаття", "Zakarpattia"},
		{"Казимирчук", "Kazymyrchuk"},
		{"Медвин", "Medvyn"},
		{"Михайленко", "Mykhailenko"},
		{"Іванків", "Ivankiv"},
		{"Іващенко", "Ivashchenko"},
		{"Їжакевич", "Yizhakevych"},
		{"Кадиївка", "Kadyivka"},
		{"Мар’їне", "Marine"},
		{"Мар'їне", "Marine"},
		{"Йосипівка", "Yosypivka"},
		{"Олексій", "Oleksii"},
		{"Київ", "Kyiv"},
		{"Коваленко", "Kovalenko"},
		{"Лебедин", "Lebedyn"},
		{"Леонід", "Leonid"},
		{"Миколаїв", "Mykolaiv"},
		{"Маринич", "Marynych"},
		{"Ніжин", "Nizhyn"},
		{"Наталія", "Nataliia"},
		{"Одеса", "Odesa"},
		{"Полтава", "Poltava"},
		{"Пу́тін — хуйло́", "Putin - khuylo"},
		{"Решетилівка", "Reshetylivka"},
		{"Рибчинський", "Rybchynskyi"},
		{"Суми", "Sumy"},
		{"Соломія", "Solomiia"},
		{"Тернопіль", "Ternopil"},
		{"Троць", "Trots"},
		{"Ужгород", "Uzhhorod"},
		{"Уляна", "Uliana"},
		{"Фастів", "Fastiv"},
		{"Філіпчук", "Filipchuk"},
		{"Харків", "Kharkiv"},
		{"Христина", "Khrystyna"},
		{"Біла Церква", "Bila Tserkva"},
		{"Стеценко", "Stetsenko"},
		{"Чернівці", "Chernivtsi"},
		{"Шевченко", "Shevchenko"},
		{"Шостка", "Shostka"},
		{"Кишеньки", "Kyshenky"},
		{"Щербухи", "Shcherbukhy"},
		{"Гоща", "Hoshcha"},
		{"Гаращенко", "Harashchenko"},
		{"Юрій", "Yurii"},
		{"Корюківка", "Koriukivka"},
		{"Яготин", "Yahotyn"},
		{"Ярошенко", "Yaroshenko"},
		{"Костянтин", "Kostiantyn"},
		{"Знам’янка", "Znamianka"},
		{"Феодосія", "Feodosiia"},
	}

	for _, test := range tests {
		if v := Trans(lang.UK, test.value); v != test.expected {
			t.Errorf("Trans(UK, %q): got %q want %q",
				test.value, v, test.expected)
		}
	}
}
