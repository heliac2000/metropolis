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
			//   [1] 1 3 4 5 51 52 53 54 101 102 103 152
			//
			// R は 1-base, golang は 0-base
			k: []int{1}, j: 2,
			expected: []int{0, 2, 3, 4, 50, 51, 52, 53, 100, 101, 102, 151},
		},
		// R> SurrAdj(c(5, 21, 124), 4)
		//  [1]  1   2   3   4   6   7   8   9  17  18  19  20  22  23  24  25  26  52  53
		// [20]  54  55  56  57  58  68  69  70  71  72  73  74  75  76  77 103 104 105 106
		// [39] 107 119 120 121 122 123 125 126 127 128 154 155 156 170 171 172 173 174 175
		// [58] 176 177 205 221 222 223 224 225 226 273 274 275 324
		{
			k: []int{4, 20, 123}, j: 3,
			expected: []int{
				0, 1, 2, 3, 5, 6, 7, 8, 16, 17, 18, 19, 21, 22, 23, 24, 25, 51, 52,
				53, 54, 55, 56, 57, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 102, 103, 104, 105,
				106, 118, 119, 120, 121, 122, 124, 125, 126, 127, 153, 154, 155, 169, 170, 171, 172, 173, 174,
				175, 176, 204, 220, 221, 222, 223, 224, 225, 272, 273, 274, 323,
			},
		},
	}

	for _, tc := range testCases {
		adj := SurrAdj(tc.k, AdjCuml[tc.j])
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
		//    [1]   1   9  17  26  52  58  68  77 103 107 119 128 154 156 170 177 205 221 222
		//   [20] 226 273 275 324
		//
		// R は 1-base, golang は 0-base
		{
			k: []int{4, 20, 123}, j: 3, q: 2,
			expected: []int{
				0, 8, 16, 25, 51, 57, 67, 76, 102, 106, 118, 127,
				153, 155, 169, 176, 204, 220, 221, 225, 272, 274, 323,
			},
		},
		{
			// R> SurrAdjEx(2, 3, 1)
			//   [1] 4 5 51 53 54 101 102 103 152
			k: []int{1}, j: 2, q: 0,
			expected: []int{3, 4, 50, 52, 53, 100, 101, 102, 151},
		},
	}

	for _, tc := range testCases {
		adj := SurrAdjEx(tc.k, AdjCuml[tc.j], AdjCuml[tc.q])
		if !reflect.DeepEqual(adj, tc.expected) {
			t.Errorf("\ngot  %v\nwant %v", adj, tc.expected)
			return
		}
	}
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
