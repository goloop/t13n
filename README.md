[//]: # (!!!Don't modify the README.md, use `make readme` to generate it!!!)


[![Go Report Card](https://goreportcard.com/badge/github.com/goloop/t13n)](https://goreportcard.com/report/github.com/goloop/t13n) [![License](https://img.shields.io/badge/license-BSD-blue)](https://github.com/goloop/t13n/blob/master/LICENSE) [![License](https://img.shields.io/badge/godoc-YES-green)](https://godoc.org/github.com/goloop/t13n)

*Version: v1.0.0*

# t13n

Package t13n (transliteration) implements methods for converting text in unicode
format to ASCII format.


## Installation

To install this module use `go get` as:

    $ go get -u github.com/goloop/t13n

## Quick Start

To use this module import it as: `github.com/goloop/t13n`


### Conversion functions

#### Rune to ASCII character as string.

Use the `String` method to convert a rune to an ASCII character.

```go
package main

import (
	"fmt"

	"github.com/goloop/t13n"
)

func main() {
	// Converts a unicode character to an ASCII string.
	a, b := t13n.String('世'), t13n.String('界')
	fmt.Printf("Hello %s%s\n", a, b)

	// Output:
	//      Hello shijie
}
```

#### Unicode string to ASCII string.

Use the `Make` method to convert unicode string to ASCII string.

```go
package main

import (
	"fmt"

	"github.com/goloop/t13n"
)

func main() {
	// Simple string transliteration - without taking
	// into account regional peculiarities.
	str := t13n.Make("Bonjour à tous!")
	fmt.Println(str)

	// Output:
	//      Bonjour a tous!
}
```

#### Conversion with dialect features.

Use the `Trans` method to take advantage of regional transliteration features. 

```go
package main

import (
	"fmt"

	"github.com/goloop/t13n"
	"github.com/goloop/t13n/lang"
)

func main() {
	// You can specify the desired language as a constant,
	// for example: lang.UK, lang.DE, lang.SL, ... etc. or as a string,
	// for example: "uk", "de", "sl", ... etc..
	// Use lang.None or "" to ignore regional rules.
	str := t13n.Trans(lang.UK, "Доброго вечора, ми з України!")
	fmt.Println(str)

	// Output:
	//      Dobroho vechora, my z Ukrainy!
}
```

#### Converts with custom rules.

Use the Render method to set custom conversion rules. For example, create a slug generation method.

```go
package main

import (
	"fmt"
	"strings"

	"github.com/goloop/t13n"
	"github.com/goloop/t13n/lang"
)

// The slugRules sets the rules for converting strings to slug.
// Returns an ASCII string of the converted character according
// to its custom rules. Offset index - which character to go to next.
// And true if the conversion was successful and you need to use this result.
//
// The method must be of the lang.TransRules type.
func slugRules(ts lang.TransState) (string, int, bool) {
	switch ts.Value {
	case " ", "~", "_":
		return "_", 0, true
	case "@", "#", "$", "%", "`", "\"", "'", ".":
		fallthrough
	case "^", "&", "*", "(", ")", "+", "<", ">", "?":
		return "", 0, true
	}

	if strings.HasSuffix(ts.Value, " ") {
		runes := []rune(ts.Value)
		if ts.Next != 0 {
			ts.Value = string(runes[:len(runes)-1]) + "-"
		}
	}

	return strings.ToLower(ts.Value), 0, true
}

func main() {
	// Create slug from text on different languages.
	host := "https://example.com/"
	slug := t13n.Render(lang.None, "你好世界", slugRules)
	fmt.Printf("%s%s\n", host, slug)

	// Output:
	//      https://example.com/ni-hao-shi-jie
}
```


### As t13n generator

You can create a conversion object also.

```go
package main

import (
	"fmt"

	"github.com/goloop/t13n"
	"github.com/goloop/t13n/lang"
)

func main() {
	text := `
        Зіроньки на небі сяють,
        наші хлопці не гуляють.
        Бьють паскуду москаля,
        аж тремтить моя земля.
    `

	t := t13n.New()
	t.Lang(lang.UK)
	fmt.Println(t.Make(text))
    
	// Output:
	//      Zironky na nebi siaiut,
	//      nashi khloptsi ne huliaiut.
	//      Biut paskudu moskalia,
	//      azh tremtyt moia zemlia.
}
```

## Usage

#### func  Make

    func Make(t string) string

Make converts a unicode string to an ASCII string.

#### func  Render

    func Render(l, t string, fn lang.TransRules) (result string)

Render converts a unicode string to an ASCII string with the rules of the
selected language and with custom rules.

#### func  String

    func String(c rune) string

String returns string value by rune from the main dictionary.

#### func  Trans

    func Trans(l, t string) string

Trans converts a unicode string to an ASCII string with the rules of the
selected language.

#### func  Version

    func Version() string

Version returns the version of the module.

#### type T13n

    type T13n struct {}


T13n the t13n constructor.

#### func  New

    func New() *T13n

New retursn pointer to T13n.

#### func (*T13n) Lang

    func (t *T13n) Lang(l string) *T13n

Lang sets the type of language features to use during transliteration.

#### func (*T13n) Make

    func (t *T13n) Make(text string) string

Make converts a unicode string to an ASCII string with the rules of the selected
language.

#### func (*T13n) Rules

    func (t *T13n) Rules(fn lang.TransRules) *T13n

Rules establishes a custom extensions method of language rules.
