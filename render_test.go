package t13n

import (
	"testing"

	"github.com/goloop/t13n/lang"
)

// TestRender tests Render function.
func TestRender(t *testing.T) {
	var tests = []struct {
		value    string
		expected string
	}{
		{"Bzia zbaşa", "Bzia zbasa"},
		{"Фэсапщы", "Fesapshchy"},
		{"Salam əleyküm", "Salam aleykum"},
		{"Qysh je", "Qysh je"},
		{"Ç’kemi", "Ckemi"},
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
		{"Hello 世界", "Hello ShiJie"},
	}

	for _, test := range tests {
		if v := Render(lang.None, test.value); v != test.expected {
			t.Errorf("expected %s but %s", test.expected, v)
		}
	}

}

// TestBosnian tests Render function for Bosnian language.
func TestBosnian(t *testing.T) {
	var tests = []struct {
		value    string
		expected string
	}{
		{"Bihać", "Bihac"},
		{"Čapljina", "Capljina"},
		{"Goražde", "Gorazde"},
		{"Široki Brijeg", "Siroki Brijeg"},
		{"Živinice", "Zivinice"},
		{"ЂЊЋŽĆČŠ", "DJNJTSHZCCS"},
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
		if v := Render(lang.Bs, test.value); v != test.expected {
			t.Errorf("expected %s but %s", test.expected, v)
		}
	}

}

// TestBulgarian tests Render function for Bulgarian language.
func TestBulgarian(t *testing.T) {
	var tests = []struct {
		value    string
		expected string
	}{
		{"Č Ć É F Ô Š Ž", "C C E F O S Z"},
		{"č ć é f ô š ž", "c c e f o s z"},
		{"Ќ Ѣ Џ Њ Ъ", "Kj E Dzh Nj A"},
		{"ќ ѣ џ њ ъ", "kj e dzh nj a"},
	}

	for _, test := range tests {
		if v := Render(lang.Bg, test.value); v != test.expected {
			t.Errorf("expected %s but %s", test.expected, v)
		}
	}

}

// TestCatalan tests Render function for Catalan language.
func TestCatalan(t *testing.T) {
	var tests = []struct {
		value    string
		expected string
	}{
		{
			"À à É é È è Í í Ï ï Ó ó Ò ò Ú ú Ü ü Ç ç",
			"A a E e E e I i I i O o O o U u U u C c",
		},
	}

	for _, test := range tests {
		if v := Render(lang.Ca, test.value); v != test.expected {
			t.Errorf("expected %s but %s", test.expected, v)
		}
	}
}

// TestCroatian tests Render function for Croatian language.
func TestCroatian(t *testing.T) {
	var tests = []struct {
		value    string
		expected string
	}{
		{"Č Ć DŽ Đ LJ NJ Š Ž", "C C DZ Dj LJ NJ S Z"},
		{"č ć dž đ lj nj š ž", "c c dz dj lj nj s z"},
	}

	for _, test := range tests {
		if v := Render(lang.Hr, test.value); v != test.expected {
			t.Errorf("expected %s but %s", test.expected, v)
		}
	}
}

// TestDanish tests Render function for Danish language.
func TestDanish(t *testing.T) {
	var tests = []struct {
		value    string
		expected string
	}{
		{"Æ Ø Å æ ø å", "AE Oe Aa ae oe aa"},
	}

	for _, test := range tests {
		if v := Render(lang.Da, test.value); v != test.expected {
			t.Errorf("expected %s but %s", test.expected, v)
		}
	}
}

// TestEsperanto tests Render function for Esperanto language.
func TestEsperanto(t *testing.T) {
	var tests = []struct {
		value    string
		expected string
	}{
		{"Ĉ Ĝ Ĥ Ĵ Ŝ Ŭ", "C Gx Hx Jx Sx Ux"},
		{"ĉ ĝ ĥ ĵ ŝ ŭ", "c gx hx jx sx ux"},
	}

	for _, test := range tests {
		if v := Render(lang.Eo, test.value); v != test.expected {
			t.Errorf("expected %s but %s", test.expected, v)
		}
	}
}

// TestGerman tests Render function for German language.
func TestGerman(t *testing.T) {
	var tests = []struct {
		value    string
		expected string
	}{
		{"Ä Ö Ü ẞ", "Ae Oe Ue Ss"},
		{"ä ö ü ß", "ae oe ue ss"},
	}

	for _, test := range tests {
		if v := Render(lang.De, test.value); v != test.expected {
			t.Errorf("expected %s but %s", test.expected, v)
		}
	}
}

// TestHungarian tests Render function for Hungarian language.
func TestHungarian(t *testing.T) {
	var tests = []struct {
		value    string
		expected string
	}{
		{"Á É Í Ó Ö Ő Ú Ü Ű", "A E I O Oe Oe U Ue Ue"},
		{"á é í ó ö ő ú ü ű", "a e i o oe oe u ue ue"},
	}

	for _, test := range tests {
		if v := Render(lang.Hu, test.value); v != test.expected {
			t.Errorf("expected %s but %s", test.expected, v)
		}
	}
}

// TestMacedonian tests Render function for Macedonian language.
func TestMacedonian(t *testing.T) {
	var tests = []struct {
		value    string
		expected string
	}{
		{"Ѓ Ќ Џ Љ Њ", "Gj Kj Dj Lj Nj"},
		{"ѓ ќ џ љ њ", "gj kj dj lj nj"},
	}

	for _, test := range tests {
		if v := Render(lang.Mk, test.value); v != test.expected {
			t.Errorf("expected %s but %s", test.expected, v)
		}
	}
}

// TestNorwegian tests Render function for Norwegian language.
func TestNorwegian(t *testing.T) {
	var tests = []struct {
		value    string
		expected string
	}{
		{"Æ Ø Å Ð Ô Ê Å Ä Ö", "AE Oe A D O E A A O"},
		{"æ ø å ð ô ê å ä ö", "ae oe a d o e a a o"},
	}

	for _, test := range tests {
		if v := Render(lang.Nb, test.value); v != test.expected {
			t.Errorf("expected %s but %s", test.expected, v)
		}
	}
}

// TestRussian tests Render function for Russian language.
func TestRussian(t *testing.T) {
	var tests = []struct {
		value    string
		expected string
	}{
		{"Ъ Ы Ћ Ѱ Ѳ Ѵ", "' Y Tsh Ps F Y"},
		{"ъ ы ћ ѱ ѳ ѵ", "' y tsh ps f y"},
		{"Путин - Хуйло", "Putin - Khuylo"},
	}

	for _, test := range tests {
		if v := Render(lang.Ru, test.value); v != test.expected {
			t.Errorf("expected %s but %s", test.expected, v)
		}
	}
}

// TestSerbian tests Render function for Serbian language.
func TestSerbian(t *testing.T) {
	var tests = []struct {
		value    string
		expected string
	}{
		{"Đ đ Ђ ђ", "Dj dj Dje dje"},
	}

	for _, test := range tests {
		if v := Render(lang.Sr, test.value); v != test.expected {
			t.Errorf("expected %s but %s", test.expected, v)
		}
	}
}

// TestSlovenian tests Render function for Slovenian language.
func TestSlovenian(t *testing.T) {
	var tests = []struct {
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
		if v := Render(lang.Sl, test.value); v != test.expected {
			t.Errorf("expected %s but %s", test.expected, v)
		}
	}
}

// TestSwedish tests Render function for Swedish language.
func TestSwedish(t *testing.T) {
	var tests = []struct {
		value    string
		expected string
	}{
		{"Å Ä Ö", "A Ae Oe"},
		{"å ä ö", "a ae oe"},
	}

	for _, test := range tests {
		if v := Render(lang.Sv, test.value); v != test.expected {
			t.Errorf("expected %s but %s", test.expected, v)
		}
	}
}

// TestUkrainian tests Render function for Ukrainian language.
func TestUkrainian(t *testing.T) {
	var tests = []struct {
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
		if v := Render(lang.Uk, test.value); v != test.expected {
			t.Errorf("expected %s but %s", test.expected, v)
		}
	}
}
