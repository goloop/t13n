package t13n

import "testing"

// TestTable verifies structural invariants of the decoded base table: it
// covers exactly the Basic Multilingual Plane and every replacement is pure
// 7-bit ASCII (that is the whole point of transliteration).
func TestTable(t *testing.T) {
	tbl := table()
	if len(tbl) != tableSize {
		t.Fatalf("table size: got %d want %d", len(tbl), tableSize)
	}

	for id, item := range tbl {
		for _, c := range item {
			if c > 127 {
				t.Errorf("non-ASCII value %q at code point %d", item, id)
			}
		}
	}
}
