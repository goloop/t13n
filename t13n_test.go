package t13n

import (
	"testing"

	"github.com/goloop/t13n/lang"
)

// TestNew tests New function.
func TestNew(t *testing.T) {
	expected := "Dobroho vechora, mu z Ukrainu!"

	ctr := func(ts lang.TransState) (string, int, bool) {
		if ts.Value == "y" {
			return "u", 0, true
		}

		return "", 0, false
	}

	trans := New(lang.UK)
	trans.Rules(ctr)
	trans.Together(1)
	if v := trans.Make("Доброго вечора, ми з України!"); v != expected {
		t.Errorf("expected %s but %s", expected, v)
	}
}
