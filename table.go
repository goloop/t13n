package t13n

import (
	"bytes"
	"compress/flate"
	_ "embed"
	"encoding/binary"
	"io"
	"sync"
)

// tableSize is the number of Unicode code points covered by the base
// transliteration table: the Basic Multilingual Plane, U+0000..U+FFFD.
const tableSize = 65534

// libBin holds the base transliteration table as a compact, flate-compressed
// resource. Each non-empty entry is encoded as varint(delta code point) +
// varint(length) + bytes, which keeps the data an order of magnitude smaller
// than the equivalent array of string literals and out of the executable's
// static image.
//
//go:embed lib.bin
var libBin []byte

// table decodes the embedded resource into the base table on first use and
// caches the result. A program that imports the package but never
// transliterates pays no memory cost for the table at all.
var table = sync.OnceValue(decodeTable)

// decodeTable inflates the embedded libBin resource into the base table. It
// is the source for the package-wide lazy table.
func decodeTable() *[tableSize]string {
	return decode(libBin)
}

// decode inflates the flate-compressed, varint-encoded table data into the
// base transliteration table, where the index is a Unicode code point and the
// value is its ASCII replacement (an empty string means "no mapping").
//
// The data is embedded at build time and immutable at runtime, so any decode
// failure means the binary itself is corrupt: decode fails loudly rather than
// silently returning a half-built table.
func decode(data []byte) *[tableSize]string {
	raw, err := io.ReadAll(flate.NewReader(bytes.NewReader(data)))
	if err != nil {
		panic("t13n: cannot decode embedded transliteration table: " + err.Error())
	}

	arr := new([tableSize]string)
	cp := 0
	for i := 0; i < len(raw); {
		delta, n := binary.Uvarint(raw[i:])
		if n <= 0 {
			panic("t13n: corrupt transliteration table (truncated code point)")
		}
		i += n
		cp += int(delta)

		length, n := binary.Uvarint(raw[i:])
		if n <= 0 {
			panic("t13n: corrupt transliteration table (truncated length)")
		}
		i += n
		arr[cp] = string(raw[i : i+int(length)])
		i += int(length)
	}

	return arr
}
