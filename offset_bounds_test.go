package t13n

import (
	"github.com/goloop/t13n/v2/lang"
	"math"
	"strings"
	"testing"
)

func TestRuleOffsetCannotOverflow(t *testing.T) {
	got := Render(lang.None, "abc", func(s lang.TransState) (string, int, bool) { return s.Value, math.MaxInt, true })
	if got != "a" {
		t.Fatalf("got %q", got)
	}
}

var longResult string

func BenchmarkLongText(b *testing.B) {
	text := strings.Repeat("A long sentence for transliteration. ", 100)
	Make(text)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		longResult = Make(text)
	}
}
