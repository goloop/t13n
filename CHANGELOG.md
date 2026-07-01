# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [2.1.0]

### Added
- `WithFallback(func(rune) string)` option (and chainable `Fallback`) supplies a
  replacement for runes with no mapping, including code points outside the Basic
  Multilingual Plane, so they need not be dropped.
- `WithStrictASCII()` option (and chainable `StrictASCII`) strips any non-ASCII
  produced by a custom rule or fallback, keeping the output pure 7-bit ASCII.
- `RunesSeq(l, text string) iter.Seq2[rune, string]` iterates the
  transliteration one unit at a time without building the whole string.

### Fixed
- Decoding the embedded table now fails with a clear panic on a truncated
  entry instead of looping forever.
- Macedonian: comments in the rule table no longer contradict their values.

### Documentation
- The ASCII guarantee is clarified to cover the built-in tables and language
  rules; custom rules own their output (see `WithStrictASCII`).

## [2.0.0]

### Changed
- Import path is now `github.com/goloop/t13n/v2`.
- `New` takes functional options (`WithLang`, `WithRules`); `Lang` and `Rules`
  are chainable and return `*T13n`.
- Transliteration runs in a single, allocation-light pass; the base table is
  stored as a compact embedded resource decoded lazily on first use, which
  removes about a megabyte from importing binaries.

### Added
- `Rune(c rune) (string, bool)` reports whether a code point has a mapping.
- `Version() string`.

### Removed
- Parallel processing and the `Together` function/method, along with the
  package-level mutable configuration.

### Fixed
- `String` no longer panics on negative runes.
- Bulgarian: capital `Я` and small `я` are no longer swapped.
- Esperanto: `Ĉ`/`ĉ` now transliterate to `Cx`/`cx` (their entries had been
  attached to control code points and never fired).
- Removed dead code and an unreachable condition in the Russian and Ukrainian
  rules.
- All-caps diacritic expansions are title-cased (`ÄÖÜ` → `AeOeUe`) so letter
  boundaries stay unambiguous.
