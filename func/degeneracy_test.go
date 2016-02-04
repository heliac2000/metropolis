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
	SetInitData("./data")

	testCases := []testCasesDegeneracy{
		{
			// [R]
			// coutp <- list()
			// for(k in 1:Nparallel){
			//   coutp[[k]] <- list()
			//   coutp[[k]][[1]] = Canonical.Order(Canonical.Gen())
			// }
			// cout = coutp[[1]]
			// cout_temp = Canonical.Order(ExtensionReductionBlock(cout[[1]])[[1]])
			// format(degeneracy(cout_temp), digit=22)
			//
			pos:      [][]int{{1226, 1180}, {1226}, {1226}, {1226}, {1226}, {1226}, {1226}, {1226}, {1226}, {0}},
			chr:      [][]int{{5, 1}, {6}, {2}, {10}, {9}, {1}, {2}, {5}, {7}, {0}},
			ori:      [][]float64{{0, 60}, {0}, {120}, {120}, {60}, {60}, {120}, {120}, {60}, {0}},
			expected: 3.760098224563903029162e+30,
		},
		{
			pos:      [][]int{{1625, 1475, 1275, 1825, 1125, 1975, 925, 725, 575, 2175}, {0}},
			chr:      [][]int{{6, 4, 6, 4, 4, 7, 6, 7, 4, 6}, {0}},
			ori:      [][]float64{{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, {0}},
			expected: 5000,
		},
	}

	for _, tc := range testCases {
		actual := Degeneracy(tc.pos, tc.chr, tc.ori)
		if round(actual) != round(tc.expected) {
			n := math.Floor(math.Log10(actual))
			if math.Abs(actual-tc.expected)/math.Pow(10, n) > 1.0E-12 {
				t.Errorf("\ngot  %v\nwant %v", actual, tc.expected)
			}
		}
	}
}

func round(f float64) float64 {
	return math.Floor(f + .5)
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
