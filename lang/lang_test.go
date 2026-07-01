package lang

import "testing"

// call is a helper that invokes a rule for a single current/next character.
func call(r TransRules, curr, next rune) (string, int, bool) {
	return r(TransState{Curr: curr, Next: next})
}

// TestRulesDispatch checks the language dispatcher: known codes return a rule,
// unknown codes return nil, and matching is case-insensitive.
func TestRulesDispatch(t *testing.T) {
	known := []string{BG, HR, DA, EO, DE, HU, MK, NB, RU, SR, SL, SV, UK}
	for _, code := range known {
		if Rules(code) == nil {
			t.Errorf("Rules(%q) = nil, want a rule", code)
		}
	}

	if Rules("UK") == nil {
		t.Error("Rules is not case-insensitive")
	}

	for _, code := range []string{None, "zz", "english"} {
		if Rules(code) != nil {
			t.Errorf("Rules(%q) != nil, want nil", code)
		}
	}
}

// TestMapRules checks the shared factory: a hit returns (value, 0, true) and a
// miss returns ("", 0, false).
func TestMapRules(t *testing.T) {
	r := mapRules(map[int]string{65: "X"})
	if v, off, ok := call(r, 'A', 0); v != "X" || off != 0 || !ok {
		t.Errorf("hit: got (%q,%d,%v)", v, off, ok)
	}
	if v, off, ok := call(r, 'B', 0); v != "" || off != 0 || ok {
		t.Errorf("miss: got (%q,%d,%v)", v, off, ok)
	}
}

// TestBulgarianCase locks the capital/small Я mapping (previously inverted).
func TestBulgarianCase(t *testing.T) {
	cases := []struct {
		c    rune
		want string
	}{
		{'Я', "Ya"}, {'я', "ya"}, {'Й', "Y"}, {'й', "y"}, {'Ъ', "A"},
	}
	for _, c := range cases {
		if v, _, ok := call(bgRules, c.c, 0); !ok || v != c.want {
			t.Errorf("bgRules(%q): got %q (ok=%v) want %q", c.c, v, ok, c.want)
		}
	}
}

// TestEsperanto locks the x-system, including Ĉ/ĉ, whose entries used to sit
// on control code points (24/25) and never fire.
func TestEsperanto(t *testing.T) {
	cases := []struct {
		c    rune
		want string
	}{
		{'Ĉ', "Cx"}, {'ĉ', "cx"}, {'Ĝ', "Gx"}, {'ŭ', "ux"},
	}
	for _, c := range cases {
		if v, _, ok := call(eoRules, c.c, 0); !ok || v != c.want {
			t.Errorf("eoRules(%q): got %q (ok=%v) want %q", c.c, v, ok, c.want)
		}
	}

	// The old control-character entries must not exist.
	if _, _, ok := call(eoRules, rune(24), 0); ok {
		t.Error("eoRules still maps control code point 24")
	}
}

// TestRussianRules covers the digraph УЙ path, a plain map hit, and a miss.
func TestRussianRules(t *testing.T) {
	if v, off, ok := call(ruRules, 'У', 'й'); v != "Uy" || off != 1 || !ok {
		t.Errorf("УЙ: got (%q,%d,%v) want (Uy,1,true)", v, off, ok)
	}
	if v, off, ok := call(ruRules, 'у', 'й'); v != "uy" || off != 1 || !ok {
		t.Errorf("уй: got (%q,%d,%v) want (uy,1,true)", v, off, ok)
	}
	if v, off, ok := call(ruRules, 'Ж', 0); v != "Zh" || off != 0 || !ok {
		t.Errorf("Ж: got (%q,%d,%v) want (Zh,0,true)", v, off, ok)
	}
	// У without the й trigger falls through to the plain lookup (no entry).
	if _, _, ok := call(ruRules, 'У', 'A'); ok {
		t.Error("У without й trigger should not match")
	}
	if _, _, ok := call(ruRules, 'A', 0); ok {
		t.Error("ruRules matched a non-Cyrillic character")
	}
}

// TestUkrainianRules covers apostrophe dropping, the зг digraph, the initial
// vs in-word forms, and a miss.
func TestUkrainianRules(t *testing.T) {
	// Apostrophe is dropped (empty replacement, changed=true).
	if v, off, ok := (ukRules(TransState{Curr: '\'', IsApostrophe: true})); v != "" || off != 0 || !ok {
		t.Errorf("apostrophe: got (%q,%d,%v) want (\"\",0,true)", v, off, ok)
	}

	// Digraphs, both cases, consuming the following rune.
	digraphs := []struct {
		curr, next rune
		want       string
	}{
		{'з', 'г', "zgh"}, {'З', 'Г', "Zgh"},
		{'у', 'й', "uy"}, {'У', 'Й', "Uy"},
	}
	for _, d := range digraphs {
		if v, off, ok := call(ukRules, d.curr, d.next); v != d.want || off != 1 || !ok {
			t.Errorf("%c%c: got (%q,%d,%v) want (%q,1,true)",
				d.curr, d.next, v, off, ok, d.want)
		}
	}

	// A letter that has no distinct in-word form keeps its default value
	// even mid-word (Г -> "H").
	if v, _, _ := ukRules(TransState{Curr: 'Г', IsBegin: false}); v != "H" {
		t.Errorf("in-word Г: got %q want H", v)
	}

	// Initial Є -> "Ye", in-word Є -> "Ie".
	if v, _, _ := ukRules(TransState{Curr: 'Є', IsBegin: true}); v != "Ye" {
		t.Errorf("initial Є: got %q want Ye", v)
	}
	if v, _, _ := ukRules(TransState{Curr: 'Є', IsBegin: false}); v != "Ie" {
		t.Errorf("in-word Є: got %q want Ie", v)
	}

	if _, _, ok := call(ukRules, 'Z', 0); ok {
		t.Error("ukRules matched a non-Cyrillic character")
	}
}
