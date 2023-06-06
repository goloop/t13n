[![Go Report Card](https://goreportcard.com/badge/github.com/goloop/t13n)](https://goreportcard.com/report/github.com/goloop/t13n) [![License](https://img.shields.io/badge/license-MIT-brightgreen)](https://github.com/goloop/t13n/blob/master/LICENSE) [![License](https://img.shields.io/badge/godoc-YES-green)](https://godoc.org/github.com/goloop/t13n)


# t13n

Package t13n (transliteration) implements methods for converting Unicode text to ASCII.


## Installation

Use go get to install this module.:

```
$ go get -u github.com/goloop/t13n
```

## Quick Start

Import t13n into your code as

```go
import "github.com/goloop/t13n"
```

You can use fast transliteration functions or create a conversion object of t13n.T13n type.


### Conversion functions

Just do it!

```go
package main

import (
	"fmt"

	"github.com/goloop/t13n"
)

func main() {
	// Original Lao and Japanese.
	fmt.Println(t13n.Make("ເຮືອຮົບລັດເຊຍໄປ くそ ຕົວທ່ານເອງ!"))

	// Output: eyy`obldesnyaip Ku So howthaane`ng!
}
```

#### Transliterate unicode character to ASCII

Use the `String` method to transliterate a single unicode character into the appropriate ASCII string.

```go
package main

import (
	"fmt"

	"github.com/goloop/t13n"
)

func main() {
	// Converts a unicode character to an ASCII string.
	// Without the use of linguistic regional properties.
	shi, jie := t13n.String('世'), t13n.String('界')
	fmt.Printf("Hello %s%s\n", shi, jie)

	// Output: Hello Shi Jie
}
```

The transliteration of a character can take into account the language rules of a particular region.

```go
package main

import (
	"fmt"

	"github.com/goloop/t13n"
	"github.com/goloop/t13n/lang"
)

func main() {
	// Converts a unicode character to an ASCII string.
	// Without the use of linguistic regional properties.
	y, o, r := t13n.String('й'), t13n.String('о'), t13n.String('р')
	fmt.Printf("%s%s%s\n", y, o, r)

	// Using linguistic regional transliteration rules.
	// For example, in the Ukrainian language the letter `й`
	// at the beginning of the word is translated as` y`.
	transRules := lang.Rules(lang.UK)
	if t, _, ok := transRules(lang.TransState{Curr: 'й', IsBegin: true}); ok {
		y = t
	}

	fmt.Printf("%s%s%s\n", y, o, r)

    // Output:
    //  ior
	//  yor
}
```



#### Transliterate unicode string to ASCII string.

Use the `Make` method to transliterate unicode string to ASCII string.

```go
package main

import (
	"fmt"

	"github.com/goloop/t13n"
)

func main() {
	// Simple string transliteration - without taking
	// into account regional peculiarities.
	fmt.Println(t13n.Make("こんにちは、みんな!"))

	// Output: Ko N Ni Chi Ha Mi N Na!
}
```

#### Linguistic features of transliteration

Use the "Trans" method to take advantage of regional transliteration.

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

	// Output: Dobroho vechora, my z Ukrainy!
}
```

#### Transliteration with custom rules

Use the `Render` method to set custom conversion rules. For example, create a slug generation method.

```go
package main

import (
	"fmt"
	"strings"

	"github.com/goloop/t13n"
	"github.com/goloop/t13n/lang"
)

// The slugRules sets custom transliteration rules
// but better to use the github.com/goloop/slug module.
//
// Returns an ASCII string of the converted character according
// to its custom rules. Offset index - which character to go to next.
// And true if the conversion was successful and you need to use this result.
//
// The method must be of the lang.TransRules type.
func slugRules(ts lang.TransState) (string, int, bool) {
	// Ignore ranges.
	// Important: edit maps ascending order only!
	var ignored = [][]int{
		{0, 47},
		{58, 64},
		{91, 96},
		{123, 141},
		{143, 152},
		{155, 155},
	}

	switch ts.Value {
	case " ", "~", "_", "\t", "\n":
		ts.Value = "-"
	case "@":
		ts.Value = "at"
	case "&":
		ts.Value = "and"
	case "#":
		ts.Value = "sharp"
	case "%":
		ts.Value = "pct"
	default:
		id := int(ts.Curr)
		for _, d := range ignored {
			// If the item isn't in the following ranges.
			if id < d[0] {
				break
			}

			// If the item is in the current range.
			if id >= d[0] && id <= d[1] {
				ts.Value = ""
				break
			}
		}
	}

	if len(ts.Value) > 1 && strings.HasSuffix(ts.Value, " ") {
		runes := []rune(ts.Value)
		if ts.Next != 0 {
			ts.Value = string(runes[:len(runes)-1]) + "-"
		}
	}

	return ts.Value, 0, true
}

func main() {
	// Create slug from text on different languages.
	host := "https://example.com/"
	slug := t13n.Render(lang.None, "你好世界", slugRules)
	fmt.Printf("%s%s\n", host, slug)

	// Output: https://example.com/Ni-Hao-Shi-Jie
}
```

#### Multithreaded transliteration

The transliteration of the line is done by dividing the line into several parts and performing their transliteration in separate goroutines.

By default, the number of threads is set as one. But you can set the number of threads use `Together` method.

Strings shorter than 256 characters are transliterated in one stream regardless of Together settings.

```go
package main

import (
	"fmt"
	"time"

	"github.com/goloop/t13n"
	"github.com/goloop/t13n/lang"
)

func main() {
	text := `
        Кілька днів назад ...
    ` // very long text here, 254101 characters.

	// Single-threaded.
	t13n.Together(1)

	start := time.Now()
	str := t13n.Trans(lang.UK, text)
	t1 := time.Since(start)

	// Multithreaded.
	t13n.Together(12)

	start = time.Now()
	_ = t13n.Trans(lang.UK, text)
	t2 := time.Since(start)

	// Print result.
	fmt.Printf("Single-threaded:\nLength: %d\nTime: %s\n\n", len(str), t1)
	fmt.Printf("Multithreaded:\nLength: %d\nTime: %s\n\n", len(str), t2)

	// Output:
	//   Single-threaded:
	//   Length: 254101
	//   Time: 4.327649979s
	//
	//   Multithreaded:
	//   Length: 254101
	//   Time: 372.701373ms
}
```


### As generator of transliteration

You can create a transliteration object also.

```go
package main

import (
	"fmt"

	"github.com/goloop/t13n"
	"github.com/goloop/t13n/lang"
)

func main() {
	text := `
        Отак подивишся здаля на москаля,
        Неначе й справді він людина,
        Та від Курил і до Кремля
        Воно було і є скотина.
    `

	uk := t13n.New(lang.UK)
	uk.Together(12) // multithreaded
	// uk := t13n.New().Lang(lang.UK).ParallelTasks(12)

	sl := t13n.New("")
	sl.Lang(lang.SL)
	sl.Together(1) // single-threaded
	// sl := t13n.New().Lang(lang.SL).ParallelTasks(1)

	// Print result.
	fmt.Println(uk.Make(text))
	fmt.Println(sl.Make(text))

	// Output:
	//   Otak podyvyshsia zdalia na moskalia,
	//   Nenache y spravdi vin liudyna,
	//   Ta vid Kuryl i do Kremlia
	//   Vono bulo i ye skotyna.
	//
	//   Otak podivishsia zdalia na moskalia,
	//   Nenache i spravdi vin liudina,
	//   Ta vid Kuril i do Kremlia
	//   Vono bulo i ie skotina.
}
```

## Functions

- **Make**(t string) string

  Make transliterates a unicode string to an ASCII string, it's doesn't take into account regional linguistic features of transliteration.

-- **Render**(l, t string, ctr lang.TransRules) (result string)

  Render transliterates a Unicode string into an ASCII string with taking into account regional linguistic features of the transliteration depending from the language.

  The third parameter can specify the function of custom transliteration rules or nil.

- **String**(c rune) string

  String returns string value by rune from the main lib, it's doesn't take into account regional linguistic features of transliteration.

- **Together**(pt int) int

  Together sets the number of parallel transliteration tasks.

- **Trans**(l, t string) string

  Trans transliterates a Unicode string into an ASCII string with taking into account regional linguistic features of the transliteration depending from the language.

- **Version**() string

  Version returns the version of the module it's has a format `"v{major_version}.{minor_version}.{patch_version}"`.

- **New**(l string) *T13n

  New retursn pointer to T13n.


## Method of T13n object

- **Lang**(l string)

  Lang sets the type of language features to use during transliteration.

- **Make**(text string) string

  Make transliterates a unicode string to an ASCII string. This method takes into account the selected language and apply regional transliteration settings.

- **Rules**(ctr lang.TransRules)

  Rules establishes a custom extensions method of language rules.

- **Together**(pt int) int

  Together sets the number of parallel transliteration tasks.
