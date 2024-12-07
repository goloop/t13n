package t13n

import (
	"testing"

	"github.com/goloop/t13n/lang"
)

var (
	shortText  = "Hello, 世界! くそ Доброго вечора"
	mediumText = `
		Кілька днів назад ми з України побачили світ таким,
		яким він є насправді. Люди тут справжні - не штучні,
		небо блакитне - не сіре, а сонце жовте - не чорне.
		こんばんは, ми з України!
	`
	// Simulate a very long text by repeating medium text
	longText = func() string {
		result := ""
		for i := 0; i < 100; i++ {
			result += mediumText
		}
		return result
	}()
)

// Simple conversion benchmarks
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

// Language-specific conversion benchmarks
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

// Parallel processing benchmarks
func BenchmarkParallel_SingleThread(b *testing.B) {
	Together(1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Trans(lang.UK, longText)
	}
}

func BenchmarkParallel_MultiThread(b *testing.B) {
	Together(12)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Trans(lang.UK, longText)
	}
}

// Object-oriented usage benchmarks
func BenchmarkT13n_SingleThread(b *testing.B) {
	t := New(lang.UK)
	t.Together(1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t.Make(longText)
	}
}

func BenchmarkT13n_MultiThread(b *testing.B) {
	t := New(lang.UK)
	t.Together(12)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t.Make(longText)
	}
}
