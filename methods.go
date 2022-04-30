package t13n

import (
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
	return render(l, t, fn, ParallelTasks)
}
