//
// trans_den_test.go
//

package functions_test

import (
	"reflect"
	"testing"

	. "./"
)

type testCasesTransDen struct {
	pos, chr [][]int
	ori      [][]float64
	expected []int
}

func TestTransDen(t *testing.T) {

	// NOTICE: R is 1-base index, golang is 0-base.
	testCases := []testCasesTransDen{
		{
			// [R]
			// ctemp <- list()
			// ctemp[[1]] <- c(1626,1476, 1276, 1826, 1126, 1976,  926,  726,  576, 2176)
			// ctemp[[2]] <- c(7,5,7,5,5,8,7,8,5,7)
			// ctemp[[3]] <- c(0,0,0,0,0,0,0,0,0,0)
			// TransDen(list(ctemp))
			// [1] 1
			//
			pos:      [][]int{{1625, 1475, 1275, 1825, 1125, 1975, 925, 725, 575, 2175}},
			chr:      [][]int{{6, 4, 6, 4, 4, 7, 6, 7, 4, 6}},
			ori:      [][]float64{{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			expected: []int{1},
		},
	}

	for _, tc := range testCases {
		actual := TransDen(tc.pos, tc.chr, tc.ori)
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("\ngot  %v\nwant %v", actual, tc.expected)
			return
		}
	}
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
