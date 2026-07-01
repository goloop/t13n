# t13n — reference

The full reference for the `t13n` package: the mental model, single-character
and whole-text transliteration, language rules, custom rules, the reusable
`T13n`, iteration and practical recipes.

Ukrainian version: **[DOC.UK.md](DOC.UK.md)**.

## Contents

- [Mental model](#mental-model)
- [Whole-text functions](#whole-text-functions)
- [Single characters](#single-characters)
- [Language rules](#language-rules)
- [Custom rules](#custom-rules)
- [The reusable T13n](#the-reusable-t13n)
- [Iteration](#iteration)
- [Version](#version)
- [Recipes and tips](#recipes-and-tips)

## Mental model

`t13n` (transliteration) converts Unicode text to ASCII. Three guarantees hold
for every function:

1. **The output is always pure 7-bit ASCII.** Whatever goes in, what comes out
   is safe for URLs, filenames and ASCII-only systems.
2. **Nothing panics.** Any input is accepted, including invalid UTF-8; unmapped
   code points simply produce no output.
3. **The base table is lazy and compact.** It is stored as an embedded resource
   and decoded on first use, so importing the package is cheap.

Transliteration runs in a single, allocation-light pass. There are two ways in:
package-level functions for one-off calls, and a reusable `T13n` configured with
functional options for repeated use.

```go
import "github.com/goloop/t13n/v2"
```

## Whole-text functions

```go
func Make(text string) string
func Trans(l, text string) string
func Render(l, text string, ctr lang.TransRules) string
```

- `Make` transliterates without regional rules.
- `Trans` applies the rules of language `l`.
- `Render` is `Trans` plus a custom rule function (or `nil`).

```go
t13n.Make("こんにちは、みんな!")       // "Ko N Ni Chi Ha Mi N Na!"
t13n.Make("世界")                     // "Shi Jie"
```

## Single characters

```go
func String(c rune) string
func Rune(c rune) (string, bool)
```

`String` transliterates one code point (returns `""` when unmapped). `Rune` also
reports whether a mapping exists — `ok` is false for unmapped code points such
as emoji or anything outside the Basic Multilingual Plane:

```go
s, ok := t13n.Rune('界') // "Jie ", true
_, ok = t13n.Rune('😀')  // "", false
```

## Language rules

`Trans`/`Render` take a language `l`. Pass a `lang` constant (for example
`lang.UK`, `lang.DE`, `lang.SL`) or the equivalent string (`"uk"`, `"de"`,
`"sl"`); use `lang.None` or `""` to skip regional rules:

```go
import "github.com/goloop/t13n/v2/lang"

t13n.Trans(lang.UK, "Доброго вечора, ми з України!")
// "Dobroho vechora, my z Ukrainy!"
```

Regional rules matter because the same character transliterates differently
across languages — for instance Ukrainian and Slovenian render the same Cyrillic
text differently.

## Custom rules

A custom rule runs **after** the language rules, letting you post-process the
output — for example to build a slug. A rule is a `lang.TransRules`:

```go
type TransRules func(TransState) (string, int, bool)
```

It returns the replacement string, an offset (how many following runes it
consumed, for digraphs), and whether it applied. `TransState` gives the
surrounding context:

| Field | Meaning |
|-------|---------|
| `Prev`, `Curr`, `Next` | the previous / current / next rune (`rune(0)` at edges) |
| `Value`        | the proposed translation of `Curr` |
| `IsBegin`      | `Curr` is at the start of a word |
| `IsApostrophe` | `Curr` is an apostrophe between two non-delimiters |

```go
func slug(ts lang.TransState) (string, int, bool) {
    switch ts.Value {
    case " ", "_", "~":
        return "-", 0, true
    }
    return strings.ToLower(ts.Value), 0, true
}

t13n.Render(lang.UK, "Доброго вечора", slug) // "dobroho-vechora"
```

> For real slugs, prefer the `github.com/goloop/slug` module — this example only
> illustrates custom rules.

## The reusable T13n

```go
func New(opts ...Option) *T13n

func WithLang(l string) Option
func WithRules(r lang.TransRules) Option
func WithFallback(f func(rune) string) Option
func WithStrictASCII() Option
```

`New` builds a configurable, reusable transliterator. It is safe for concurrent
use once configured, so build it once and share it:

```go
uk := t13n.New(t13n.WithLang(lang.UK))
sl := t13n.New(t13n.WithLang(lang.SL))

text := "Отак подивишся здаля на москаля"
uk.Make(text) // "Otak podyvyshsia zdalia na moskalia"
sl.Make(text) // "Otak podivishsia zdalia na moskalia"
```

The setters are chainable and interchangeable with the options:

```go
func (t *T13n) Make(text string) string
func (t *T13n) Lang(l string) *T13n
func (t *T13n) Rules(r lang.TransRules) *T13n
func (t *T13n) Fallback(f func(rune) string) *T13n
func (t *T13n) StrictASCII() *T13n
```

```go
tr := t13n.New().Lang(lang.UK).Rules(slug)
```

`Fallback` sets a function to handle runes the base table doesn't map (instead
of dropping them); `StrictASCII` guarantees the output contains only ASCII by
discarding anything that would not be.

## Iteration

```go
func RunesSeq(l, text string) iter.Seq2[rune, string]
```

`RunesSeq` yields `(rune, transliteration)` pairs for `text` under language `l`,
so you can drive your own rendering without materialising the whole result:

```go
for r, s := range t13n.RunesSeq(lang.UK, "Київ") {
    fmt.Printf("%c -> %q\n", r, s)
}
```

## Version

```go
func Version() string
```

Returns the module version as `"v{major}.{minor}.{patch}"`.

## Recipes and tips

**Build one transliterator per language.** `t13n.New(t13n.WithLang(...))` is
immutable-once-configured and concurrency-safe — construct it at startup and
reuse it, rather than calling `Trans` with the language on every call.

**Detect unmapped code points.** Use `Rune` (not `String`) when you need to know
whether a character had a mapping, e.g. to substitute a placeholder for emoji.

**Guarantee ASCII output.** Add `WithStrictASCII()` (or call `StrictASCII()`)
when the result feeds a strictly ASCII sink and you would rather drop than pass
through an unexpected character.

**Post-process with a rule.** Reach for `Render`/`Rules` to lower-case, swap
separators or collapse characters in the same pass — but use the `slug` module
when you actually want slugs.

**Stream large input.** `RunesSeq` lets you consume the transliteration lazily,
pair by pair, instead of building one large string.
