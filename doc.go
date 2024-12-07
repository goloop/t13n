// Package t13n implements methods for converting Unicode text
// to ASCII (transliteration).
//
// The package provides both simple conversion functions and an
// object-oriented interface for transliterating text. It supports
// language-specific transliteration rules and custom transliteration
// extensions.
//
// Basic usage:
//
//	// Simple conversion.
//	ascii := t13n.Make("こんにちは")
//
//	// Language-specific conversion.
//	ukText := t13n.Trans(lang.UK, "Доброго вечора")
//
//	// Object-oriented usage with custom settings.
//	t := t13n.New(lang.UK)
//	t.Together(12) // enable parallel processing
//	result := t.Make("Доброго вечора")
//
// Features:
//   - Single character and string transliteration.
//   - Language-specific conversion rules.
//   - Custom transliteration rules support.
//   - Parallel processing for large texts.
//   - Object-oriented interface.
//   - Preservation of case sensitivity.
//   - Special handling for apostrophes and hieroglyphs.
//
// The package automatically switches to parallel processing for texts longer
// than 256 characters when parallel tasks are enabled. The number of parallel
// tasks can be controlled using the Together() function or method.
//
// For language-specific transliteration, use the lang package constants:
//   - lang.None - No language-specific rules
//   - lang.UK - Ukrainian
//   - lang.DE - German
//   - lang.EN - English
//
// The package is thread-safe and can be used concurrently.
package t13n
