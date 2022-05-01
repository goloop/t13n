package t13n

import (
	"math"
	"strings"
	"unicode"

	"github.com/goloop/t13n/lang"
)

// The transChunkState the state of transliteration of a single chunk.
type transChunkState struct {
	id     int
	value  string
	offset int
}

// The isDelimiter returns true if char is characters delimiter.
//
// Importantly:
// - the number is also a delimiter of characters;
// - the apostrophe is also a delimiter.
func isDelimiter(c rune) bool {
	// Separator ranges.
	// Important: edit maps ascending order only!
	var separators = [][]int{
		{9, 9},
		{32, 64},
		{91, 96},
		{123, 126},
		{8216, 8217},
	}

	// The zero character is also a separator.
	id := int(c)
	if id == 0 {
		return true
	}

	for _, interval := range separators {
		// Abort the check if there are no actual ranges in the list below.
		if id < interval[0] {
			break
		}

		// Char is separator if the character belongs to a range.
		if id >= interval[0] && id <= interval[1] {
			return true
		}
	}

	return false
}

// The isHieroglyph returns true if char is hieroglyph.
func isHieroglyph(c rune) bool {
	var hieroglyphs = [][]int{
		{11904, 11929},
		{11931, 12019},
		{12032, 12245},
		{12293, 12295},
		{12321, 12329},
		{12344, 12347},

		{12353, 12438},
		{12445, 12447},

		{13312, 19903},
		{19968, 40956},
		{63744, 64109},
		{64112, 64217},

		// Todo: find all ranges of characters.
		// See example: unicode.Is(unicode.Han, 'ん')
	}

	id := int(c)
	for _, item := range hieroglyphs {
		if id < item[0] {
			break
		}

		if id >= item[0] && id <= item[1] {
			return true
		}
	}

	return false
}

// The isApostrophe returns true if char is apostrophe.
// The apostrophe is quote between the characters.
func isApostrophe(ts lang.TransState) bool {
	// The apostrophe is a symbol '`' between two liters.
	// If there is no previous or next character,
	// it is exactly an apostrophe.
	if ts.Prev == 0 || ts.Next == 0 {
		return false
	}

	switch ts.Curr {
	case 39, 96, 8216, 8217:
		if !isDelimiter(ts.Prev) && !isDelimiter(ts.Next) {
			return true
		}
	}

	return false
}

// The toChunks splits the slice of characters (chars) into chunks
// of the specified number (n). The number of chunks is indicated as
// the recommended number. If the number of characters in the slice
// is less than the specified recommended number of chunks -
// the number of chunks will change.
//
// Returns a chunk slice (slice of character slices) and
// the actual number of chunks.
//
// Note:
// A simpler algorithm is to split the runes by the nearest separator.
// This will prevent line breaks between characters that form a single
// sound (this is important for some language regional rules).
//
// But this approach requires an additional subcycle of search
// for a separator, very inefficient for the generated custom text
// where there are no separators, doesn't allow to allocate memory
// in advance to save chunks (this should be done in the process of
// chunk formation, which takes some CPU time).
func toChunks(chars []rune, n int) ([][]rune, int) {
	var (
		step   int      // the size of the chunk
		count  int      // the actual number of chunks
		result [][]rune // the slice of chunks

		total = len(chars)
	)

	// The number of chunks cannot be more than
	// the actual number of characters in the slice
	// and cannot be less than one (if chars isn't empty).
	switch {
	case total < n:
		n = total
	case n < 1 || len(chars) < 256:
		n = 1
	}

	// Split slice of characters into chunks.
	step = int(math.Ceil(float64(total) / float64(n)))
	count = int(math.Ceil(float64(total) / float64(step)))
	result = make([][]rune, 0, count)

	for i := 0; i < step*count; i += step {
		end := i + step
		if end > total {
			end = total
		}

		result = append(result, chars[i:end])
	}

	return result, count
}

// The renderChunk renders one chunk, launches as a separate goroutine.
// Where is:
//  id - chunk sequence number (starts from zero);
//  chars - slice of characters to transliteration;
//  prevOver - previous character before the first character
//      in the current chunk (the last character in the previous chunk);
//  nextOver - next character after the last character
//      in the current chunk (taken as the first character in the next chunk);
//  isBegin - true if before the first character in the chunk
//      was a separator (is determined outside the goroutine,
//      as the separator is not always a divider, such as an apostrophe);
//  prevIsUpper - true if prev char was is upper;
//  ltr - language translation rules;
//  ctr - custom translation rules;
//  ch - the result of the function.
func renderChunk(
	id int,
	chars []rune,
	prevOver, nextOver rune,
	isBegin, prevIsUpper bool,
	ltr, ctr lang.TransRules,
	ch chan transChunkState,
) {
	var (
		value  string
		offset int
	)

	// Parse current chunk.
	for i := 0; i < len(chars); i++ {
		ts := lang.TransState{
			Prev:    prevOver,
			Curr:    chars[i],
			Next:    nextOver,
			IsBegin: isBegin,
		}

		// If there is a previous character or the next character
		// in the current chunk, we need to use them.
		if i > 0 {
			ts.Prev = chars[i-1]
		}

		if i < len(chars)-1 {
			ts.Next = chars[i+1]
		}

		// Convert char to t13n and set changes of current language.
		ts.Value = String(ts.Curr)
		ts.IsApostrophe = isApostrophe(ts)
		if ltr != nil {
			offset = 0
			if t, m, ok := ltr(ts); ok {
				offset = m
				ts.Value = t
			}
		}

		// If t13n is long value like: Th, ae etc. - write all symbols
		// of t13n upper if next char is upper to.
		switch cu, nu := unicode.IsUpper(ts.Curr), unicode.IsUpper(ts.Next); {
		case prevIsUpper && cu:
			fallthrough
		case ts.Next != 0 && cu && nu:
			ts.Value = strings.ToUpper(ts.Value)
			prevIsUpper = true
		default:
			prevIsUpper = cu
		}

		// Add a space to the right and add title style
		// of all hieroglyphs or delete apostrophes.
		if isHieroglyph(ts.Curr) {
			ts.Value = strings.Title(ts.Value)
			if int(ts.Next) != 0 && !isDelimiter(ts.Next) {
				ts.Value += " "
			}
		}

		// Custom extensions.
		if ctr != nil {
			if t, m, ok := ctr(ts); ok {
				offset = m
				ts.Value = t
			}
		}

		i += offset
		value += ts.Value
		isBegin = isDelimiter(ts.Curr) && !ts.IsApostrophe

		// If the loop ends not with the last character but due to the offset,
		// then the offset must be canceled (it has already worked).
		if i-offset < len(chars)-1 {
			offset = 0
		}

		// If the previous character was a break, it cannot be is upper.
		if isBegin {
			prevIsUpper = false
		}
	}

	ch <- transChunkState{id: id, value: value, offset: offset}
}

// The renderString converts a unicode string to an ASCII string with
// the rules of the selected language and with custom rules.
//
// Where is:
//  l - language code;
//  t - text to conversion;
//  ctr - custom translation rules;
//  nt - number of threads.
func renderString(l, t string, ctr lang.TransRules, nt int) string {
	var result string

	// Split the text into chunks of runes.
	runes := []rune(t)
	chunks, total := toChunks(runes, nt)

	// The result of the work of goroutines on the transliteration of chunks.
	ch := make(chan transChunkState, total)
	tmp := map[int]transChunkState{}

	// Run goroutines for chunk parsing .
	isBegin := true
	ltr := lang.Rules(l)
	for i, chunk := range chunks {
		// Determine the previous and next characters between which
		// the first character in the chunk will be located.
		prevOver, nextOver := rune(0), rune(0)
		if i > 0 {
			chunk := chunks[i-1]
			prevOver = chunk[len(chunk)-1]
		}

		if i < total-1 {
			chunk := chunks[i+1]
			nextOver = chunk[0]
		}

		go renderChunk(
			i,
			chunk,
			prevOver,
			nextOver,
			isBegin,
			unicode.IsUpper(prevOver),
			ltr,
			ctr,
			ch,
		)

		// Parameters of the first character for the next chunk.
		if i < total-1 {
			nextChunk := chunks[i+1]
			ts := lang.TransState{
				Prev: chunk[len(chunk)-1],
				Curr: nextChunk[0],
				Next: rune(0),
			}

			// If the chunk has more than one character,
			// specify the second character in the chunk.
			if len(nextChunk) > 1 {
				ts.Next = nextChunk[1]
			}

			isBegin = isDelimiter(ts.Curr) && !isApostrophe(ts)
		}
	}

	// As result need to perform concatenation of chunks, for this save the
	// transliteration result of all chunks in the map by the chunk sequence
	// number (don't use a slice to avoid its sorting cycle).
	for len(tmp) < total {
		v := <-ch
		tmp[v.id] = v
	}
	close(ch)

	// Concatenation of chunks.
	offset := 0
	for i := 0; i < len(tmp); i++ {
		v, _ := tmp[i]
		runes := []rune(v.value)
		result += string(runes[offset:])
		offset = v.offset
	}

	return result
}
