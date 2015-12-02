package util_test

import (
	"testing"

	. "./"
)

// File exists ?
//
func TestFileExists(t *testing.T) {
	testCases := map[string]bool{
		"file.go":      true,
		"/etc":         true,
		"unknown_file": false,
		".":            true,
		"unknown_dir/": false,
		"../.git":      true,
	}

	var actual bool
	for f, expected := range testCases {
		actual = FileExists(f)
		if actual != expected {
			t.Errorf("\ngot  %v(file = %s)\nwant %v", actual, f, expected)
			return
		}
	}
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
