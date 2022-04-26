package t13n

import (
	"testing"
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
			if isCharDelimiter(c) {
				total += 1
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
			p, c, n := rune(0), runes[i], rune(0)
			if i > 0 {
				p = runes[i-1]
			}

			if i < len(runes)-1 {
				n = runes[i+1]
			}

			if isApostrophe(p, c, n) {
				total += 1
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

// TestValueByIndex tests valueByIndex function.
func TestValueByIndex(t *testing.T) {
	// Overflow.
	overflow := func() (result bool) {
		// Testing an inaccessible index will cause panic!
		defer func() {
			if r := recover(); r != nil {
				result = false
				return
			}
		}()

		a := valueByIndex(len(dictionary))
		if a != "" {
			return
		}

		return true
	}

	if o := overflow(); !o {
		t.Error("expected true but false")
	}
}
