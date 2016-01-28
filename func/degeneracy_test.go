//
// degeneracy_test.go
//

package functions_test

import (
	"math"
	"testing"

	. "./"
)

type testCasesDegeneracy struct {
	pos, chr [][]int
	ori      [][]float64
	expected float64
}

func TestDegeneracy(t *testing.T) {
	SetInitData()

	testCases := []testCasesDegeneracy{
		{
			// [R]
			// format(degeneracy(list(Zi)), digit=22)
			//
			pos:      [][]int{{1625, 1475, 1275, 1825, 1125, 1975, 925, 725, 575, 2175}, {0}},
			chr:      [][]int{{6, 4, 6, 4, 4, 7, 6, 7, 4, 6}, {0}},
			ori:      [][]float64{{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, {0}},
			expected: 5000,
		},
	}

	for _, tc := range testCases {
		actual := Degeneracy(tc.pos, tc.chr, tc.ori)
		if round(actual) != round(tc.expected) {
			t.Logf("\ngot  %v\nwant %v", actual, tc.expected)
		}
	}
}

func round(f float64) float64 {
	return math.Floor(f + .5)
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
