package functions_test

import (
	"reflect"
	"testing"

	. "./"
)

type testCasesUniqueOrientations struct {
	pos      []int
	same     []int
	expected []int
}

func TestUniqueOrientations(t *testing.T) {
	SetInitData()

	// NOTICE: R is 1-base index, golang is 0-base.
	testCases := []testCasesUniqueOrientations{
		{
			// [R] uniqueOrientations(list(c(5, 40, 200), c(1), c(1)))
			//     [[1]][1] 1 2
			//     [[2]][[2]][[1]] 50 40 31
			//
			pos:      []int{4, 39, 199},
			same:     []int{1, 2},
			expected: []int{49, 39, 30},
		},
		{
			// [R] uniqueOrientations(list(c(1), c(1), c(1)))
			pos:      []int{0},
			same:     []int{1},
			expected: []int{0},
		},
		{
			// [R] uniqueOrientations(list(c(1, 10, 100), c(1), c(1)))
			pos:      []int{0, 9, 99},
			same:     []int{1, 2},
			expected: []int{18, 9, 750},
		},
	}

	for _, tc := range testCases {
		actual, same := UniqueOrientations(tc.pos)
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
