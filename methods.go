package t13n

import (
	"runtime"

	"github.com/goloop/t13n/lang"
)

// String returns string value by rune from the main lib, it's doesn't
// take into account regional linguistic features of transliteration.
func String(c rune) string {
	if id := int(c); id < len(lib) {
		return lib[id]
	}

	return ""
}

// Make transliterates a unicode string to an ASCII string, it's doesn't
// take into account regional linguistic features of transliteration.
func Make(t string) string {
	return Render(lang.None, t, nil)
}

// Trans transliterates a Unicode string into an ASCII string
// with taking into account regional linguistic features of
// the transliteration depending from the language.
func Trans(l, t string) string {
	return Render(l, t, nil)
}

// Render transliterates a Unicode string into an ASCII string
// with taking into account regional linguistic features of
// the transliteration depending from the language.
//
// The third parameter can specify the function of
// custom transliteration rules or nil.
func Render(l, t string, ctr lang.TransRules) (result string) {
	return renderString(l, t, ctr, parallelTasks)
}

// Together sets the number of parallel transliteration tasks.
func Together(pt int) int {
	if pt > 0 && pt < runtime.NumCPU()*3 {
		parallelTasks = pt
	}

	return parallelTasks
}
