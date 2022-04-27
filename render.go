package t13n

import (
	"strings"
	"unicode"

	language "github.com/goloop/t13n/lang"
)

// Render converts a unicode string to an ASCII string with
// the rules of the selected language.
func Render(lang, text string) string {
	var languageRules = language.Rules(lang)

	result, begin, runes, wasupper := "", true, []rune(text), false
	for i := 0; i < len(runes); i++ {
		p, c, n := rune(0), runes[i], rune(0)

		if i > 0 {
			p = runes[i-1]
		}

		if i < len(runes)-1 {
			n = runes[i+1]
		}

		// Convert char to t13n and set changes of current language.
		v := valueByIndex(int(c))
		if t, m, ok := languageRules(p, c, n, begin); ok {
			i += m
			v = t
		}

		// If t13n is long value like: Th, ae etc. - write all symbols
		// of t13n upper if next char is upper to.
		cupper, nupper := unicode.IsUpper(c), unicode.IsUpper(n)
		switch {
		case wasupper && cupper:
			fallthrough
		case n != 0 && cupper && nupper:
			v = strings.ToUpper(v)
			wasupper = true
		default:
			wasupper = cupper
		}

		// Add a title style of all hieroglyphs or delete apostrophes.
		apostrophe := false
		if unicode.Is(unicode.Han, c) {
			v = strings.Title(v)
		} else {
			apostrophe = isApostrophe(p, c, n)
			if apostrophe {
				v = ""
			}
		}

		result += v
		begin = isCharDelimiter(c) && !apostrophe
		if begin {
			wasupper = false
		}
	}

	return result
}
