package t13n

// New retursn pointer to T13n.
func New(lang string) *T13n {
	return &T13n{lang: lang}
}

// T13n the t13n constructor.
type T13n struct {
	lang string
}

// Render converts a unicode string to an ASCII string with
// the rules of the selected language.
func (t *T13n) Render(text string) string {
	return Render(t.lang, text)
}
