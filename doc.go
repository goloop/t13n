// Package t13n (transliteration) implements methods
// for converting Unicode text to ASCII.
package t13n

// Version of the module as {major_version}.{minor_version}.{patch_version}.
const version = "1.1.0"

// Version returns the version of the module
// it's has a format "v{major_version}.{minor_version}.{patch_version}".
func Version() string {
	return "v" + version // add a "v" to the version
}
