package util_test

import (
	"reflect"
	"testing"
)
import . "./"

type testCasesAddMarix struct {
	Z1, Z2   [][]int
	expected [][]int
}

func TestAddMatrix(t *testing.T) {
	testCases := []testCasesAddMarix{
		{
			Z1: [][]int{{1, 2}, {3, 4}, {5, 6}},
			Z2: [][]int{{7, 8}, {9, 10}},
			expected: [][]int{
				{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10},
			},
		},
	}

	var actual [][]int
	for _, z := range testCases {
		actual = AddMatrix(z.Z1, z.Z2)
		if !reflect.DeepEqual(actual, z.expected) {
			t.Errorf("\ngot  %v\nwant %v", actual, z.expected)
			return
		}
	}
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
