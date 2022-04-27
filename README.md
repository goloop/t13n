[//]: # (!!!Don't modify the README.md, use `make readme` to generate it!!!)


[![Go Report Card](https://goreportcard.com/badge/github.com/goloop/t13n)](https://goreportcard.com/report/github.com/goloop/t13n) [![License](https://img.shields.io/badge/license-BSD-blue)](https://github.com/goloop/t13n/blob/master/LICENSE) [![License](https://img.shields.io/badge/godoc-YES-green)](https://godoc.org/github.com/goloop/t13n)

*Version: v0.2.0*


# t13n

Package t13n (transliteration) implements methods for converting text in unicode
format to ASCII format.


## Installation

To install this module use `go get` as:

    $ go get -u github.com/goloop/t13n

## Quick Start

To use this module import it as: `github.com/goloop/t13n`


### Conversion functions

You can use the quick Render method to convert text. As the first argument, the method gets the type of language according to the rules of which you want to do the conversion (it can be either a string value or a specified constant) and the second value - the text to be converted.

Example:

```go
    package main

    import (
        "fmt"

        "github.com/goloop/t13n"
        "github.com/goloop/t13n/lang"
    )

    func main() {
        // For default conversion:
        //   t13n.Render(lang.None, "...") or t13n.Render("", "...")
        // For conversion according to the rules of the Ukrainian language:
        //   t13n.Render(lang.Uk, "...") or t13n.Render("uk", "...")
        // You can choose different languages: Ab, Aa, Af, ..., Zs, Za, Zu.
        msg := t13n.Render(lang.Uk, "Доброго вечора, ми з України!")
        fmt.Println(msg)
        // Output: Dobroho vechora, my z Ukrainy!
    }
```

### As t13n generator

You can  create a conversion object also.

Example:

```go
    package main

    import (
        "fmt"

        "github.com/goloop/t13n"
        "github.com/goloop/t13n/lang"
    )

    func main() {
        // For default conversion:
        //   t13n.New(lang.None) or t13n.New("")
        // For conversion according to the rules of the Ukrainian language:
        //   t13n.New(lang.Uk) or t13n.New("uk")
        // You can choose different languages: Ab, Aa, Af, ..., Zs, Za, Zu.
        t := t13n.New(lang.Uk)
        msg := t.Render("Доброго вечора, ми з України!")
        fmt.Println(msg)
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
