package t13n

import (
	"strings"
	"unicode"

	"github.com/goloop/t13n/lang"
)

// String returns string value by rune from the main dictionary.
func String(c rune) string {
	if id := int(c); id < len(lib) {
		return lib[id]
	}

	return ""
}

// Make converts a unicode string to an ASCII string.
func Make(t string) string {
	return Render(lang.None, t, nil)
}

// Trans converts a unicode string to an ASCII string
// with the rules of the selected language.
func Trans(l, t string) string {
	return Render(l, t, nil)
}

// Render converts a unicode string to an ASCII string with
// the rules of the selected language and with custom rules.
func Render(l, t string, fn lang.TransRules) (result string) {
	transRules := lang.Rules(l)
	charList := []rune(t)
	wasUpper := false
	isBegin := true

	for i := 0; i < len(charList); i++ {
		seek := 0
		ts := lang.TransState{
			Prev:    rune(0),
			Curr:    charList[i],
			Next:    rune(0),
			IsBegin: isBegin,
		}

		if i > 0 {
			ts.Prev = charList[i-1]
		}

		if i < len(charList)-1 {
			ts.Next = charList[i+1]
		}

		// Convert char to t13n and set changes of current language.
		ts.Value = String(ts.Curr)
		ts.IsApostrophe = isApostrophe(ts)
		if transRules != nil {
			if t, m, ok := transRules(ts); ok {
				seek = m
				ts.Value = t
			}
		}

		// If t13n is long value like: Th, ae etc. - write all symbols
		// of t13n upper if next char is upper to.
		switch cu, nu := unicode.IsUpper(ts.Curr), unicode.IsUpper(ts.Next); {
		case wasUpper && cu:
			fallthrough
		case ts.Next != 0 && cu && nu:
			ts.Value = strings.ToUpper(ts.Value)
			wasUpper = true
		default:
			wasUpper = cu
		}

		// Add a space to the right and add title style
		// of all hieroglyphs or delete apostrophes.
		if unicode.Is(unicode.Han, ts.Curr) {
			ts.Value = strings.Title(ts.Value)
			if int(ts.Next) != 0 && !isDelimiter(ts.Next) {
				ts.Value += " "
			}
		}

		// Custom extensions.
		if fn != nil {
			if t, m, ok := fn(ts); ok {
				seek = m
				ts.Value = t
			}
		}

		i += seek
		result += ts.Value
		isBegin = isDelimiter(ts.Curr) && !ts.IsApostrophe
		if isBegin {
			wasUpper = false
		}
	}

	return result
}
