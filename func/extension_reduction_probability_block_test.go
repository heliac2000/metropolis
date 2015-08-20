package functions_test

import (
	"testing"

	. "./"
)

type testCasesERPBlock struct {
	pcanon, ccanon         [][]int
	ocanon                 [][]float64
	pcanon_out, ccanon_out [][]int
	ocanon_out             [][]float64
	expected               float64
}

func TestExtensionReductionProbabilityBlock(t *testing.T) {
	SetInitData("./data/PrecursorUnitCell.csv", "./data/UnitCell2.csv", "./data/PrecursorUnitCellAxes.csv")

	// NOTICE: R is 1-base index, golang is 0-base.
	//   ccdb and ccbb are characters, 1-base in R but 0-base in Golang
	//
	testCases := []testCasesERPBlock{
		{
			// [R]
			// canon = Canonical.Order(Canonical.Gen())
			// canon_out = Canonical.Order(ExtensionReductionBlock(canon)[[1]])
			// =>
			// canon = list(
			//           list(c(313, 361, 364, 408, 317, 264, 340),
			//                c(313, 263, 387), c(0)),
			//           list(c(7, 5, 5, 7, 7, 5, 6),
			//                c(6, 5, 5),  c(0)),
			//           list(c(0, 30, 120, 150, 150, 30, 120),
			//                c(0, 120, 60), c(0)))
			//
			// canon_out = list(
			//           list(c(361, 364, 408, 317, 264, 340),
			//                c(313, 263, 387), c(313), c(0)),
			//           list(c(5, 5, 7, 7, 5, 6),
			//                c(6, 5, 5), c(7), c(0)),
			//           list(c(30, 120, 150, 150, 30, 120),
			//                c(0, 120, 60), c(60), c(0)))
			//
			// format(ExtensionReductionProbabilityBlock(canon, canon_out), digits=22, trim=T)
			// [1] "0.003888888888888888343415"
			//
			pcanon:     [][]int{{313, 361, 364, 408, 317, 264, 340}, {313, 263, 387}, {0}},
			ccanon:     [][]int{{6, 4, 4, 6, 6, 4, 5}, {5, 4, 4}, {0}},
			ocanon:     [][]float64{{0, 30, 120, 150, 150, 30, 120}, {0, 120, 60}, {0}},
			pcanon_out: [][]int{{361, 364, 408, 317, 264, 340}, {313, 263, 387}, {313}, {0}},
			ccanon_out: [][]int{{4, 4, 6, 6, 4, 5}, {5, 4, 4}, {6}, {0}},
			ocanon_out: [][]float64{{30, 120, 150, 150, 30, 120}, {0, 120, 60}, {60}, {0}},
			expected:   0.003888888888888888343415,
		},
	}

	for _, tc := range testCases {
		actual := ExtensionReductionProbabilityBlock(
			tc.pcanon, tc.ccanon, tc.ocanon, tc.pcanon_out, tc.ccanon_out, tc.ocanon_out)

		if actual != tc.expected {
			t.Errorf("\ngot  %v\nwant %v", actual, tc.expected)
			return
		}
	}

	return
}

// Local Variables:
// compile-command: (concat "go test -gcflags='-B' -v " (file-name-nondirectory buffer-file-name))
// End:
