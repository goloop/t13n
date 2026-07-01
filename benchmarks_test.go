package t13n

import (
	"testing"

	"github.com/goloop/t13n/v2/lang"
)

var (
	shortText  = "Hello, 世界! くそ Доброго вечора"
	mediumText = `
		Кілька днів назад ми з України побачили світ таким,
		яким він є насправді. Люди тут справжні - не штучні,
		небо блакитне - не сіре, а сонце жовте - не чорне.
		こんばんは, ми з України!
	`
	// Simulate a very long text by repeating the medium text.
	longText = func() string {
		result := ""
		for i := 0; i < 100; i++ {
			result += mediumText
		}
		return result
	}()
)

func BenchmarkString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		String('世')
	}
}

func BenchmarkMake_Short(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Make(shortText)
	}
}

func BenchmarkMake_Medium(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Make(mediumText)
	}
}

func BenchmarkMake_Long(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Make(longText)
	}
}

func BenchmarkTrans_UK_Short(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Trans(lang.UK, shortText)
	}
}

func BenchmarkTrans_UK_Medium(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Trans(lang.UK, mediumText)
	}
}

func BenchmarkTrans_UK_Long(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Trans(lang.UK, longText)
	}
}

func BenchmarkT13n_Long(b *testing.B) {
	tr := New(WithLang(lang.UK))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tr.Make(longText)
	}
}
