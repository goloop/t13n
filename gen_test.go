package t13n

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

// TestGeneratorDeterministic rebuilds the embedded lib.bin from its source
// (the _gen command) and checks it matches the committed resource byte for
// byte. It guards against an accidental change in the encoder or the source
// table that would leave lib.bin out of date. The generator is built into and
// run from a temporary directory so the real lib.bin is never touched.
func TestGeneratorDeterministic(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping generator rebuild in -short mode")
	}
	if _, err := exec.LookPath("go"); err != nil {
		t.Skip("go toolchain not available")
	}

	dir := t.TempDir()
	bin := filepath.Join(dir, "gen")
	if runtime.GOOS == "windows" {
		bin += ".exe"
	}

	build := exec.Command("go", "build", "-o", bin, "./_gen")
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("build generator: %v\n%s", err, out)
	}

	run := exec.Command(bin)
	run.Dir = dir // gen writes lib.bin into its working directory
	if out, err := run.CombinedOutput(); err != nil {
		t.Fatalf("run generator: %v\n%s", err, out)
	}

	got, err := os.ReadFile(filepath.Join(dir, "lib.bin"))
	if err != nil {
		t.Fatalf("read regenerated lib.bin: %v", err)
	}

	if !bytes.Equal(got, libBin) {
		t.Errorf("regenerated lib.bin differs from committed resource "+
			"(%d vs %d bytes); run `go run ./_gen`", len(got), len(libBin))
	}
}
