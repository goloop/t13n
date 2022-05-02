package t13n

import (
	"runtime"

	"github.com/goloop/t13n/lang"
)

var (
	// ParallelTasks the number of parallel transliteration tasks.
	// By default, the number of threads is set as the number of CPU cores.
	ParallelTasks = 1

	// The minimum number of characters to parallelize the transliteration.
	singleThreadedLen = 256
)

// Module initialization.
func init() {
	ParallelTasks = runtime.NumCPU()
}

// New retursn pointer to T13n.
func New() *T13n {
	return &T13n{
		lang: lang.None,
		ctr:  nil,
		pt:   ParallelTasks,
	}
}

// T13n the transliteration constructor.
type T13n struct {
	// Language code whose regional rules are
	// to be used during transliteration.
	lang string

	// Custom translation rules.
	ctr lang.TransRules

	// Number of parallel transliteration tasks.
	pt int
}

// Make transliterates a unicode string to an ASCII string.
// This method takes into account the selected language and
// apply regional transliteration settings.
func (t *T13n) Make(text string) string {
	return renderString(t.lang, text, t.ctr, t.pt)
}

// Rules establishes a custom extensions method of language rules.
func (t *T13n) Rules(ctr lang.TransRules) *T13n {
	t.ctr = ctr
	return t
}

// Lang sets the type of language features to use during transliteration.
func (t *T13n) Lang(l string) *T13n {
	t.lang = l
	return t
}

// ParallelTasks sets the number of parallel transliteration tasks.
func (t *T13n) ParallelTasks(pt int) *T13n {
	t.pt = pt
	return t
}
