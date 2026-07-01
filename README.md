[![Go Report Card](https://goreportcard.com/badge/github.com/goloop/t13n)](https://goreportcard.com/report/github.com/goloop/t13n) [![License](https://img.shields.io/badge/license-MIT-brightgreen)](https://github.com/goloop/t13n/blob/master/LICENSE) [![License](https://img.shields.io/badge/godoc-YES-green)](https://godoc.org/github.com/goloop/t13n) [![Stay with Ukraine](https://img.shields.io/static/v1?label=Stay%20with&message=Ukraine%20♥&color=ffD700&labelColor=0057B8&style=flat)](https://u24.gov.ua/)


# t13n

Package t13n (transliteration) converts Unicode text to ASCII.

The output is always pure 7-bit ASCII, no function panics on any input
(including invalid UTF-8), and the base table is stored as a compact embedded
resource that is decoded lazily on first use.


## Installation

```bash
go get -u github.com/goloop/t13n/v2
```

```go
import "github.com/goloop/t13n/v2"
```

## Quick start

```go
package main

import (
	"fmt"

	"github.com/goloop/t13n/v2"
)

func main() {
	// Plain transliteration, no regional rules.
	fmt.Println(t13n.Make("こんにちは、みんな!"))

	// Output: Ko N Ni Chi Ha Mi N Na!
}
```

### A single character

Use `String` to transliterate one code point, or `Rune` when you also need to
know whether a mapping exists (`ok` is false for unmapped code points, such as
emoji or anything outside the Basic Multilingual Plane):

```go
fmt.Println(t13n.Make("世界")) // "Shi Jie"

s, ok := t13n.Rune('界') // "Jie ", true
_, ok = t13n.Rune('😀')  // "", false
```

### Language-specific transliteration

Use `Trans` to apply the regional rules of a language. Pass a `lang` constant
(for example `lang.UK`, `lang.DE`, `lang.SL`) or the equivalent string
(`"uk"`, `"de"`, `"sl"`); use `lang.None` or `""` to skip regional rules.

```go
package main

import (
	"fmt"

	"github.com/goloop/t13n/v2"
	"github.com/goloop/t13n/v2/lang"
)

func main() {
	fmt.Println(t13n.Trans(lang.UK, "Доброго вечора, ми з України!"))

	// Output: Dobroho vechora, my z Ukrainy!
}
```

### Custom rules

Use `Render` to add a custom rule function that runs after the language rules,
for example to build a slug. (For real slugs, prefer the
`github.com/goloop/slug` module.)

```go
package main

import (
	"fmt"
	"strings"

	"github.com/goloop/t13n/v2"
	"github.com/goloop/t13n/v2/lang"
)

func slug(ts lang.TransState) (string, int, bool) {
	switch ts.Value {
	case " ", "_", "~":
		return "-", 0, true
	}
	return strings.ToLower(ts.Value), 0, true
}

func main() {
	fmt.Println(t13n.Render(lang.UK, "Доброго вечора", slug))

	// Output: dobroho-vechora
}
```

A rule returns the replacement string, an offset (how many following runes it
consumed, for digraphs), and whether it applied.

### Reusable transliterator

`New` builds a configurable, reusable transliterator using functional options.
It is safe for concurrent use once configured.

```go
package main

import (
	"fmt"

	"github.com/goloop/t13n/v2"
	"github.com/goloop/t13n/v2/lang"
)

func main() {
	uk := t13n.New(t13n.WithLang(lang.UK))
	sl := t13n.New(t13n.WithLang(lang.SL))

	text := "Отак подивишся здаля на москаля"
	fmt.Println(uk.Make(text)) // Otak podyvyshsia zdalia na moskalia
	fmt.Println(sl.Make(text)) // Otak podivishsia zdalia na moskalia
}
```

The setters are chainable and interchangeable with options:

```go
tr := t13n.New().Lang(lang.UK).Rules(slug)
```

## API

Package-level functions:

- **Make**(text string) string — transliterate without regional rules.
- **Trans**(l, text string) string — transliterate with the language `l` rules.
- **Render**(l, text string, ctr lang.TransRules) string — as `Trans`, plus a
  custom rule function (or nil).
- **String**(c rune) string — base transliteration of one code point ("" if
  unmapped).
- **Rune**(c rune) (string, bool) — base transliteration of one code point and
  whether a mapping exists.
- **Version**() string — module version, `"v{major}.{minor}.{patch}"`.
- **New**(opts ...Option) *T13n — build a reusable transliterator.
- **WithLang**(l string) Option, **WithRules**(r lang.TransRules) Option —
  configuration options for `New`.

`*T13n` methods:

- **Make**(text string) string — transliterate using the configured language
  and rules.
- **Lang**(l string) *T13n — set the language (chainable).
- **Rules**(r lang.TransRules) *T13n — set a custom rule function (chainable).

## Migrating from v1

- The import path is now `github.com/goloop/t13n/v2`.
- `New` takes functional options: `New(WithLang(lang.UK))` instead of
  `New(lang.UK)`. `Lang`/`Rules` are chainable and return `*T13n`.
- Parallel processing and the `Together` function/method were removed;
  transliteration now runs in a single, allocation-light pass.
- New `Rune(c) (string, bool)` reports whether a code point has a mapping.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
