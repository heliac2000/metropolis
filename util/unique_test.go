package util_test

import (
	"reflect"
	"testing"
)
import . "./"

func TestUnique(t *testing.T) {
	testCases := [][]([]int){
		{{3, 2, 3, 1, 0, 2, 10}, {3, 2, 1, 0, 10}},
		{{1, 1, 1}, {1}},
		{{10, 12, 9, 1, 20, 15}, {10, 12, 9, 1, 20, 15}},
	}

	for _, c := range testCases {
		actual, expected := Unique(c[0]), c[1]
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("\ngot  %v\nwant %v", actual, expected)
			return
		}
	}
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
