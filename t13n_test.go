package t13n

import (
	"strings"
	"testing"

	"github.com/goloop/t13n/languages"
)

// TestVersion tests the package version.
// Note: each time you change the major version, you need to fix the tests.
func TestVersion(t *testing.T) {
	var expected = "v0." // change it for major version

	version := Version()
	if strings.HasPrefix(version, expected) != true {
		t.Error("incorrect version")
	}

	if len(strings.Split(version, ".")) != 3 {
		t.Error("version format should be as " +
			"v{major_version}.{minor_version}.{patch_version}")
	}
}

// TestNew tests New function.
func TestNew(t *testing.T) {
	var expected = "Dobroho vechora, my z Ukrainy!"

	trans := New(languages.Ukrainian)
	if v := trans.Render("Доброго вечора, ми з України!"); v != expected {
		t.Errorf("expected %s but %s", expected, v)
	}
}
