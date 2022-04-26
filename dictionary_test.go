package t13n

import (
	"testing"
)

// TestDictionary tests dictionary data.
func TestDictionary(t *testing.T) {
	// Size.
	if len(dictionary) != 65535 {
		t.Errorf("expected %d but %d", 65535, len(dictionary))
	}

	// Max value of ASCII is 256.
	for id, item := range dictionary {
		for _, c := range item {
			if int(c) > 256 {
				t.Errorf("incorrect value \"%s\" in %d position", item, id)
			}
		}
	}
}
