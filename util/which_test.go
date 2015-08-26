package util_test

import (
	"reflect"
	"testing"
)
import . "./"

// WhichOverZero
//
// func TestWhichOverZero(t *testing.T) {
// 	hop := [][]int{{0, 1, 2, 3, -1, 10, 15, -20}}
// 	actual := WhichOverZero(0, hop)
// 	expected := []int{1, 2, 3, 5, 6}
//
// 	if !reflect.DeepEqual(actual, expected) {
// 		t.Errorf("\ngot  %v\nwant %v", actual, expected)
// 		return
// 	}
// }

// Which
//
type testCasesWhich struct {
	slice    []int
	a        int
	expected []int
}

func TestWhich(t *testing.T) {
	testCases := []testCasesWhich{
		{
			slice:    []int{-1, 1, 0, 3, -1, 2},
			a:        -1,
			expected: []int{0, 4},
		},
	}

	for _, tc := range testCases {
		actual := Which(tc.slice, tc.a)
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("\ngot  %v\nwant %v", actual, tc.expected)
			return
		}
	}
}

// WhichIn
//
type testCasesWhichIn struct {
	arr_a, arr_b []int
	expected     bool
}

func TestWhichIn(t *testing.T) {
	testCases := []testCasesWhichIn{
		{
			arr_a:    []int{1, 2, 3},
			arr_b:    []int{3, 4, 5, 6},
			expected: true,
		},
		{
			arr_a:    []int{1, 2, 3},
			arr_b:    []int{4, 5, 6},
			expected: false,
		},
	}

	var actual bool
	for _, tc := range testCases {
		actual = WhichIn(tc.arr_a, tc.arr_b)
		if actual != tc.expected {
			t.Errorf("\ngot  %v\nwant %v", actual, tc.expected)
			return
		}
	}
}

// WhichNotIn
//
type testCasesWhichNotIn struct {
	arr_a, arr_b, expected []int
}

func TestWhichNotIn(t *testing.T) {
	testCases := []testCasesWhichNotIn{
		{
			arr_a:    []int{1, 2, 3},
			arr_b:    []int{3, 4, 5, 6},
			expected: []int{1, 2},
		},
		{
			arr_a:    []int{1, 2, 3},
			arr_b:    []int{4, 5, 6},
			expected: []int{1, 2, 3},
		},
	}

	var actual []int
	for _, tc := range testCases {
		actual = WhichNotIn(tc.arr_a, tc.arr_b)
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("\ngot  %v\nwant %v", actual, tc.expected)
			return
		}
	}
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
