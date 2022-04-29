package t13n

import (
	"testing"

	"github.com/goloop/t13n/lang"
)

// TestIsCharDelimiter tests isCharDelimiter function.
func TestIsCharDelimiter(t *testing.T) {
	var tests = []struct {
		value string
		total int
	}{
		{"\tHello world!", 3},
		{"Hello, how are you?", 5},
		{" !\"#$%&'()*+,-./0123456789:;<=>?@", 33},
		{"[\\]^_`", 6},
		{"{|}~", 4},
		{"it’s", 1},
		{"it‘s", 1},
	}

	for _, test := range tests {
		total := 0
		for _, c := range test.value {
			if isDelimiter(c) {
				total++
			}
		}

		if total != test.total {
			t.Errorf(
				"for %s expected %d separator(s) but %d",
				test.value,
				test.total,
				total,
			)
		}
	}
}

// TestIsApostrophe tests isApostrophe function.
func TestIsApostrophe(t *testing.T) {
	var tests = []struct {
		value string
		total int
	}{
		{"п‘ять, торф'яний, здоров`я, м’ясо, зв’язок", 5},
		{"сім’я, бур'ян, кур’єр, подвір‘я, під’їхати", 5},
		{"з’єднати, з’їхати, роз‘яснити", 3},
		{"‘hello’ 'world'", 0},
	}

	for _, test := range tests {
		total := 0
		runes := []rune(test.value)
		for i := 0; i < len(runes); i++ {
			ts := lang.TransState{Prev: rune(0), Curr: runes[i], Next: rune(0)}
			if i > 0 {
				ts.Prev = runes[i-1]
			}

			if i < len(runes)-1 {
				ts.Next = runes[i+1]
			}

			if isApostrophe(ts) {
				total++
			}
		}

		if total != test.total {
			t.Errorf(
				"for %s expected %d apostrophe(s) but %d",
				test.value,
				test.total,
				total,
			)
		}
	}
}
