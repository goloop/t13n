package t13n

import "github.com/goloop/t13n/v2/lang"

// T13n is a reusable transliterator configured with a language and, optionally,
// a custom rule function. A configured T13n is safe for concurrent use by
// multiple goroutines; the configuration methods ([T13n.Lang], [T13n.Rules])
// are meant to be called during setup, before the value is shared.
type T13n struct {
	// lang is the language code whose regional rules are applied.
	lang string

	// rules is an optional custom rule function applied after the
	// language rules.
	rules lang.TransRules
}

// Option configures a [T13n] created with [New].
type Option func(*T13n)

// WithLang sets the language whose regional rules are applied during
// transliteration (see the lang package for language codes).
func WithLang(l string) Option {
	return func(t *T13n) { t.lang = l }
}

// WithRules sets a custom rule function applied after the language rules,
// for example to build slugs.
func WithRules(r lang.TransRules) Option {
	return func(t *T13n) { t.rules = r }
}

// New returns a transliterator configured by the given options. With no
// options it transliterates without regional rules, exactly like [Make].
func New(opts ...Option) *T13n {
	t := &T13n{lang: lang.None}
	for _, opt := range opts {
		opt(t)
	}

	return t
}

// Make transliterates a Unicode string to ASCII using the configured
// language and custom rules.
func (t *T13n) Make(text string) string {
	return render(t.lang, text, t.rules)
}

// Lang sets the language whose regional rules are applied and returns the
// receiver, so calls can be chained.
func (t *T13n) Lang(l string) *T13n {
	t.lang = l
	return t
}

// Rules sets a custom rule function applied after the language rules and
// returns the receiver, so calls can be chained.
func (t *T13n) Rules(r lang.TransRules) *T13n {
	t.rules = r
	return t
}
