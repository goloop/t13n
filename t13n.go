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

	// fallback, when set, provides a replacement for a rune that has no
	// mapping (unknown or outside the Basic Multilingual Plane), so it
	// need not be silently dropped.
	fallback func(rune) string

	// strict, when true, strips any non-ASCII produced by the custom
	// rules or fallback, keeping the output pure 7-bit ASCII.
	strict bool
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

// WithFallback sets a function that supplies a replacement for any rune with
// no mapping (an unknown character or one outside the Basic Multilingual
// Plane). Without it such runes are dropped; with it they can be rendered as,
// for example, "?" or a decomposition of the caller's choosing. The fallback
// runs before the custom rules, so those still win.
func WithFallback(f func(rune) string) Option {
	return func(t *T13n) { t.fallback = f }
}

// WithStrictASCII strips any non-ASCII produced by a custom rule or fallback,
// so the output is guaranteed to be pure 7-bit ASCII even when custom code
// returns other bytes. The built-in tables and language rules are always ASCII
// regardless of this option.
func WithStrictASCII() Option {
	return func(t *T13n) { t.strict = true }
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
	return render(t.lang, text, t.rules, t.fallback, t.strict)
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

// Fallback sets a replacement function for runes with no mapping (see
// [WithFallback]) and returns the receiver, so calls can be chained.
func (t *T13n) Fallback(f func(rune) string) *T13n {
	t.fallback = f
	return t
}

// StrictASCII enables stripping of non-ASCII from custom output (see
// [WithStrictASCII]) and returns the receiver, so calls can be chained.
func (t *T13n) StrictASCII() *T13n {
	t.strict = true
	return t
}
