![Go Report Card](https://img.shields.io/badge/go%20report-retired-lightgrey) [![License](https://img.shields.io/badge/license-MIT-brightgreen)](https://github.com/goloop/t13n/blob/master/LICENSE) [![License](https://img.shields.io/badge/godoc-YES-green)](https://pkg.go.dev/github.com/goloop/t13n/v2) [![Stay with Ukraine](https://img.shields.io/static/v1?label=Stay%20with&message=Ukraine%20♥&color=ffD700&labelColor=0057B8&style=flat)](https://u24.gov.ua/)

# t13n

`t13n` (transliteration) converts Unicode text to ASCII. The output is always
pure 7-bit ASCII, no function panics on any input (including invalid UTF-8), and
the base table is stored as a compact embedded resource decoded lazily on first
use.

Use the package-level functions for one-off calls, or build a reusable,
concurrency-safe `T13n` with functional options for repeated use.

## Features

- Pure ASCII output from any Unicode text; never panics.
- Language-specific rules (`lang.UK`, `lang.DE`, `lang.SL`, …) or plain mode.
- Custom post-processing rules, single-character lookups, and a `RunesSeq`
  iterator.
- Reusable transliterator via `New` and chainable setters.
- Compact, lazily-decoded embedded table.

## Installation

```bash
go get -u github.com/goloop/t13n/v2
```

```go
import "github.com/goloop/t13n/v2"
```

Requires Go 1.24 or newer.

## Quick start

```go
package main

import (
    "fmt"

    "github.com/goloop/t13n/v2"
    "github.com/goloop/t13n/v2/lang"
)

func main() {
    // Plain transliteration, no regional rules.
    fmt.Println(t13n.Make("こんにちは、みんな!")) // Ko N Ni Chi Ha Mi N Na!
    fmt.Println(t13n.Make("世界"))               // Shi Jie

    // Language-specific rules.
    fmt.Println(t13n.Trans(lang.UK, "Доброго вечора, ми з України!"))
    // Dobroho vechora, my z Ukrainy!

    // A single character, with a "was it mapped?" flag.
    s, ok := t13n.Rune('界') // "Jie ", true
    _, no := t13n.Rune('😀') // "", false
    fmt.Println(s, ok, no)

    // A reusable, concurrency-safe transliterator.
    uk := t13n.New(t13n.WithLang(lang.UK))
    fmt.Println(uk.Make("Отак подивишся здаля на москаля"))
    // Otak podyvyshsia zdalia na moskalia
}
```

## Documentation

- Full reference and recipes: [DOC.md](DOC.md) · [DOC.UK.md](DOC.UK.md)
- Package API: [pkg.go.dev/github.com/goloop/t13n/v2](https://pkg.go.dev/github.com/goloop/t13n/v2)
- Changes between versions: [CHANGELOG.md](CHANGELOG.md)

## Contributing

Contributions are welcome. Please run `go test ./...`, `go vet ./...` and
`gofmt -l .` before submitting a pull request.

## License

`t13n` is released under the MIT License. See [LICENSE](LICENSE).
