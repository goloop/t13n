package t13n

import (
	"testing"

	"github.com/goloop/t13n/v2/lang"
)

// TestNewOptions checks that New applies functional options and that the
// language rules plus a custom rule compose correctly.
func TestNewOptions(t *testing.T) {
	// The custom rule turns every "y" produced by the Ukrainian rules
	// into "u", so "ми" -> "my" -> "mu".
	ctr := func(ts lang.TransState) (string, int, bool) {
		if ts.Value == "y" {
			return "u", 0, true
		}

		return "", 0, false
	}

	tr := New(WithLang(lang.UK), WithRules(ctr))
	got := tr.Make("Доброго вечора, ми з України!")
	want := "Dobroho vechora, mu z Ukrainu!"
	if got != want {
		t.Errorf("Make: got %q want %q", got, want)
	}
}

// TestNewDefault checks that a zero-option New behaves exactly like Make.
func TestNewDefault(t *testing.T) {
	const in = "Привіт, 世界"
	if got, want := New().Make(in), Make(in); got != want {
		t.Errorf("New().Make = %q, want Make = %q", got, want)
	}
}

// TestChaining checks that the chainable setters mutate and return the
// receiver, so they can be used fluently and interchangeably with options.
func TestChaining(t *testing.T) {
	const in = "Ґорґани"

	built := New().Lang(lang.UK)
	if built.Make(in) != Trans(lang.UK, in) {
		t.Errorf("Lang chaining: got %q want %q",
			built.Make(in), Trans(lang.UK, in))
	}

	// Lang and Rules return the same pointer they were called on.
	if built.Lang(lang.RU) != built {
		t.Error("Lang did not return the receiver")
	}
	if built.Rules(nil) != built {
		t.Error("Rules did not return the receiver")
	}
}

// TestInstanceMatchesPackage ensures the object API and the package-level
// shortcuts agree for the same configuration.
func TestInstanceMatchesPackage(t *testing.T) {
	inputs := []string{"", "Hello", "Доброго ранку", "こんにちは", "Süß"}
	langs := []string{lang.None, lang.UK, lang.RU, lang.DE}

	for _, l := range langs {
		tr := New(WithLang(l))
		for _, in := range inputs {
			if got, want := tr.Make(in), Trans(l, in); got != want {
				t.Errorf("lang %q, %q: instance %q != package %q",
					l, in, got, want)
			}
		}
	}
}
