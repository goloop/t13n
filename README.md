[//]: # (!!!Don't modify the README.md, use `make readme` to generate it!!!)


[![Go Report Card](https://goreportcard.com/badge/github.com/goloop/t13n)](https://goreportcard.com/report/github.com/goloop/t13n) [![License](https://img.shields.io/badge/license-BSD-blue)](https://github.com/goloop/t13n/blob/master/LICENSE) [![License](https://img.shields.io/badge/godoc-YES-green)](https://godoc.org/github.com/goloop/t13n)

*Version: v0.0.1-alpha*


# t13n

Package t13n (transliteration) ...

## Installation

To install this module use `go get` as:

    $ go get -u github.com/goloop/t13n

## Quick Start

To use this module import it as: `github.com/goloop/t13n`

### Conversion functions

Example:

```go
    package main

    import (
        "fmt"

        "github.com/goloop/t13n"
        "github.com/goloop/t13n/lang"
    )

    func main() {
        t := t13n.Render(lang.Uk, "Доброго вечора, ми з України!")
        fmt.Println(t)
        // Output: Dobroho vechora, my z Ukrainy!
    }
```

### Style objects

...
Example:

```go
    package main

    import (
        "fmt"

        "github.com/goloop/t13n"
        "github.com/goloop/t13n/lang"
    )

    func main() {
        t = t13n.New(lang.Uk)
        fmt.Println(t.Render("Доброго вечора, ми з України!"))
        // Output: Dobroho vechora, my z Ukrainy!
    }
```

## Usage

#### func  Render

    func Render(lang, text string) string

Render converts a unicode string to an ASCII string with the rules of the
selected language.

#### func  Version

    func Version() string

Version returns the version of the module.

#### type T13n

    type T13n struct {
    }


T13n the t13n constructor.

#### func  New

    func New(lang string) *T13n

New retursn pointer to T13n.

#### func (*T13n) Render

    func (t *T13n) Render(text string) string

Render converts a unicode string to an ASCII string with the rules of the
selected language.
