package t13n

import (
	"runtime"

	"github.com/goloop/t13n/lang"
)

// ParallelTasks the number of parallel transliteration tasks.
var ParallelTasks = 1

func init() {
	ParallelTasks = runtime.NumCPU()
}

// New retursn pointer to T13n.
func New() *T13n {
	return &T13n{lang: lang.None}
}

// T13n the t13n constructor.
type T13n struct {
	lang string
	fn   lang.TransRules
}

// Make converts a unicode string to an ASCII string
// with the rules of the selected language.
func (t *T13n) Make(text string) string {
	return Render(t.lang, text, t.fn)
}

// Rules establishes a custom extensions method of language rules.
func (t *T13n) Rules(fn lang.TransRules) *T13n {
	t.fn = fn
	return t
}

// Lang sets the type of language features to use during transliteration.
func (t *T13n) Lang(l string) *T13n {
	t.lang = l
	return t
}
