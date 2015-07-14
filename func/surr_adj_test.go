package functions_test

import (
	"reflect"
	"testing"

	. "../util"
	. "./"
)

type testCasesSurrAdj struct {
	k        []int
	j        int
	expected []int
}

var AdjCuml [][][]int = LoadFromCsvFileList("./data/AdjCuml.csv")

func TestSurrAdj(t *testing.T) {
	testCases := []testCasesSurrAdj{
		{
			// R> SurrAdj(2, 3)
			//   [1] 1 3 4 5 26 27 28 29 51 52 53 77
			//
			// R は 1-base, golang は 0-base
			k: []int{1}, j: 2,
			expected: []int{0, 2, 3, 4, 25, 26, 27, 28, 50, 51, 52, 76},
		},
		// R> SurrAdj(c(5, 21, 124), 4)
		//    [1]   1   2   3   4   6   7   8   9  17  18  19  20  22  23  24  25  27  28  29
		//   [20]  30  31  32  33  43  44  45  46  47  48  49  50  53  54  55  56  57  69  70
		//   [39]  71  72  73  74  75  79  80  81  95  96  97  98  99 100 105 120 121 122 123
		//   [58] 125 146 147 148 149 150 172 173 174 175 198 199 200 224
		{
			k: []int{4, 20, 123}, j: 3,
			expected: []int{
				0, 1, 2, 3, 5, 6, 7, 8, 16, 17, 18, 19, 21, 22, 23, 24, 26, 27, 28,
				29, 30, 31, 32, 42, 43, 44, 45, 46, 47, 48, 49, 52, 53, 54, 55, 56,
				68, 69, 70, 71, 72, 73, 74, 78, 79, 80, 94, 95, 96, 97, 98, 99, 104,
				119, 120, 121, 122, 124, 145, 146, 147, 148, 149, 171, 172, 173, 174,
				197, 198, 199, 223},
		},
	}

	for _, tc := range testCases {
		adj := SurrAdj(tc.k, tc.j, AdjCuml)
		if !reflect.DeepEqual(adj, tc.expected) {
			t.Errorf("\ngot  %v\nwant %v", adj, tc.expected)
			return
		}
	}
}

type testCasesSurrAdjEx struct {
	k        []int
	j, q     int
	expected []int
}

func TestSurrAdjEx(t *testing.T) {
	testCases := []testCasesSurrAdjEx{
		// R> SurrAdjEx(c(5, 21, 124), 4, 3)
		//    [1]   1   9  17  25  27  33  43  50  53  57  69  79  81  95 105 120 146 172 198
		//   [20] 200 224
		//
		// R は 1-base, golang は 0-base
		{
			k: []int{4, 20, 123}, j: 3, q: 2,
			expected: []int{
				0, 8, 16, 24, 26, 32, 42, 49, 52, 56, 68, 78,
				80, 94, 104, 119, 145, 171, 197, 199, 223,
			},
		},
		{
			// R> SurrAdjEx(2, 3, 1)
			//   [1] 4 5 26 28 29 51 52 53 77
			k: []int{1}, j: 2, q: 0,
			expected: []int{3, 4, 25, 27, 28, 50, 51, 52, 76},
		},
	}

	for _, tc := range testCases {
		adj := SurrAdjEx(tc.k, tc.j, tc.q, AdjCuml)
		if !reflect.DeepEqual(adj, tc.expected) {
			t.Errorf("\ngot  %v\nwant %v", adj, tc.expected)
			return
		}
	}
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
