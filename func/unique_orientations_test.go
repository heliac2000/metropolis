package functions_test

import (
	"reflect"
	"testing"

	. "./"
)

type testCasesUniqueOrientations struct {
	pos, chr []int
	ori      []float64
	same     []int
	expected []int
}

func TestUniqueOrientations(t *testing.T) {
	SetInitData("./data")

	// NOTICE: R is 1-base index, golang is 0-base.
	testCases := []testCasesUniqueOrientations{
		{
			// [R] uniqueOrientations(list(c(1227, 1181), c(6, 2), c(0, 60)))
			//
			pos: []int{1226, 1180}, chr: []int{5, 1}, ori: []float64{0, 60},
			same:     []int{0, 1},
			expected: []int{1134, 1180},
		},
	}

	for _, tc := range testCases {
		same, actual := UniqueOrientations(tc.pos, tc.chr, tc.ori)
		if !reflect.DeepEqual(same, tc.same) {
			t.Errorf("\ngot  %v\nwant %v", same, tc.same)
			return
		}
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("\ngot  %v\nwant %v", actual, tc.expected)
			return
		}
	}

	return
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
