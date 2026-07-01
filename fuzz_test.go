package t13n

import (
	"testing"

	"github.com/goloop/t13n/v2/lang"
)

// FuzzMake asserts the invariant that must hold for every possible input:
// Make never panics and always yields pure 7-bit ASCII.
func FuzzMake(f *testing.F) {
	seeds := []string{
		"", "Hello", "Доброго вечора", "世界", "😀", "\xff\xfe",
		"Ĉu?", "Müller", "a'b", "з’їхати",
	}
	for _, s := range seeds {
		f.Add(s)
	}

	f.Fuzz(func(t *testing.T, s string) {
		out := Make(s)
		if !isASCII(out) {
			t.Fatalf("Make(%q) produced non-ASCII %q", s, out)
		}
	})
}

// FuzzTrans asserts the same invariant across every language with regional
// rules, so that a rule (including digraph offsets) can never emit non-ASCII
// or run off the end of the input.
func FuzzTrans(f *testing.F) {
	langs := []string{
		lang.UK, lang.RU, lang.BG, lang.DE, lang.EO, lang.SR, lang.MK,
	}

	seeds := []string{
		"Згурський", "Хуйло", "Яя", "Ĉapelo", "Straße", "ЂЊЋ", "з’я",
	}
	for _, s := range seeds {
		f.Add(s)
	}

	f.Fuzz(func(t *testing.T, s string) {
		for _, l := range langs {
			out := Trans(l, s)
			if !isASCII(out) {
				t.Fatalf("Trans(%q, %q) produced non-ASCII %q", l, s, out)
			}
		}
	})
}

// FuzzInstanceMatchesPackage checks that the object API never diverges from the
// package-level shortcut for the same configuration, on arbitrary input.
func FuzzInstanceMatchesPackage(f *testing.F) {
	for _, s := range []string{"", "Київ", "世界", "Süß", "\xff"} {
		f.Add(s)
	}

	tr := New(WithLang(lang.UK))
	f.Fuzz(func(t *testing.T, s string) {
		if got, want := tr.Make(s), Trans(lang.UK, s); got != want {
			t.Fatalf("instance %q != package %q for %q", got, want, s)
		}
	})
}
