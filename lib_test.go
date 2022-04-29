package t13n

import (
	"testing"
)

// TestLib tests lib data.
func TestLib(t *testing.T) {
	// Size.
	if len(lib) != 65534 {
		t.Errorf("expected %d but %d", 65534, len(lib))
	}

	// Max value of ASCII is 256.
	for id, item := range lib {
		for _, c := range item {
			if int(c) > 256 {
				t.Errorf("incorrect value \"%s\" in %d position", item, id)
			}
		}
	}
}
