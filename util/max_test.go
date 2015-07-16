package util_test

import (
	"testing"

	. "./"
)

func TestMax(t *testing.T) {
	testCases := map[int][]int{
		5:  []int{2, 5, 0, 1, 4},
		10: []int{-4, 0, 10, 9, 8},
	}

	var actual int
	for expected, arr := range testCases {
		actual = Max(arr...)
		if actual != expected {
			t.Errorf("\ngot  %v\nwant %v", actual, expected)
			return
		}
	}
}

func TestMin(t *testing.T) {
	testCases := map[int][]int{
		0:  []int{2, 5, 0, 1, 4},
		-4: []int{-4, 0, 10, 9, 8},
	}

	var actual int
	for expected, arr := range testCases {
		actual = Min(arr...)
		if actual != expected {
			t.Errorf("\ngot  %v\nwant %v", actual, expected)
			return
		}
	}
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
