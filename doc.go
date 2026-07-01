// Package t13n converts Unicode text to ASCII (transliteration).
//
// It provides simple package-level functions and a reusable, configurable
// transliterator. Transliteration is driven by a base table covering the
// Basic Multilingual Plane and, optionally, language-specific rules and a
// custom rule function.
//
// Basic usage:
//
//	// Plain transliteration, no regional rules.
//	ascii := t13n.Make("こんにちは") // "Ko N Ni Chi Ha"
//
//	// Language-specific transliteration.
//	uk := t13n.Trans(lang.UK, "Доброго вечора") // "Dobroho vechora"
//
//	// A reusable transliterator configured with functional options.
//	tr := t13n.New(t13n.WithLang(lang.UK))
//	result := tr.Make("Доброго вечора")
//
// A custom rule function runs after the language rules and can be used, for
// example, to build slugs:
//
//	slug := func(ts lang.TransState) (string, int, bool) {
//	    if ts.Value == " " {
//	        return "-", 0, true
//	    }
//	    return strings.ToLower(ts.Value), 0, true
//	}
//	s := t13n.Render(lang.UK, "Доброго вечора", slug) // "dobroho-vechora"
//
// Guarantees:
//   - The output of the built-in tables and language rules is always pure
//     7-bit ASCII. A custom rule (see [Render], [WithRules]) or fallback (see
//     [WithFallback]) is responsible for its own output; enable
//     [WithStrictASCII] to strip any non-ASCII it produces.
//   - Characters with no mapping (including code points outside the Basic
//     Multilingual Plane) are dropped; use [Rune] to detect them or
//     [WithFallback] to supply a replacement.
//   - No function panics on any input, including invalid UTF-8.
//
// For language-specific transliteration, use the constants in the lang package
// (for example lang.UK, lang.DE, lang.RU).
//
// The package-level functions ([String], [Rune], [Make], [Trans], [Render])
// are stateless and safe for concurrent use. A [T13n] value is safe for
// concurrent use once configured; its setters are intended for setup before
// the value is shared.
package t13n
