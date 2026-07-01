// Command gen encodes the human-editable base transliteration table (data.go)
// into the compact, flate-compressed lib.bin resource embedded by the package.
//
// The table in data.go is the source of truth: edit it there, then regenerate
// with:
//
//	go run ./_gen
//
// This directory is prefixed with "_" so the Go tool ignores it during normal
// builds and tests; the 5 MB source table therefore never reaches the compiled
// library or a consumer's binary.
package main

import (
	"bytes"
	"compress/flate"
	"encoding/binary"
	"fmt"
	"os"
)

func main() {
	var buf bytes.Buffer
	var tmp [binary.MaxVarintLen64]byte
	put := func(v uint64) {
		n := binary.PutUvarint(tmp[:], v)
		buf.Write(tmp[:n])
	}

	prev, nonEmpty := 0, 0
	for cp, s := range lib {
		if s == "" {
			continue
		}
		put(uint64(cp - prev))
		prev = cp
		put(uint64(len(s)))
		buf.WriteString(s)
		nonEmpty++
	}

	var comp bytes.Buffer
	fw, err := flate.NewWriter(&comp, flate.BestCompression)
	if err != nil {
		fatal(err)
	}
	if _, err := fw.Write(buf.Bytes()); err != nil {
		fatal(err)
	}
	if err := fw.Close(); err != nil {
		fatal(err)
	}

	if err := os.WriteFile("lib.bin", comp.Bytes(), 0o644); err != nil {
		fatal(err)
	}
	fmt.Printf("lib.bin: %d entries, raw=%d bytes, flate=%d bytes\n",
		nonEmpty, buf.Len(), comp.Len())
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, "gen:", err)
	os.Exit(1)
}
