package functions_test

import (
	"testing"

	. "./"
)

type testCasesERPReaction struct {
	pcanon, ccanon         [][]int
	ocanon                 [][]float64
	pcanon_out, ccanon_out [][]int
	ocanon_out             [][]float64
	expected               float64
}

func TestExtensionReductionProbabilityReaction(t *testing.T) {
	SetInitData(
		"./data/PrecursorUnitCell.csv", "./data/UnitCell2.csv", "./data/PrecursorUnitCellAxes.csv",
		"./data/kernelregS_Rep_log.json", "./data/kernelregS_Att.json", "./data/svm_model.json")

	// NOTICE: R is 1-base index, golang is 0-base.
	//   ccdb and ccbb are characters, 1-base in R but 0-base in Golang
	//
	testCases := []testCasesERPReaction{
		{
			// [R]
			// canon = Canonical.Order(Canonical.Gen())
			// canon_out = Canonical.Order(ExtensionReductionBlock(canon)[[1]])
			// =>
			// canon = list(
			//           list(313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 0),
			//           list(4, 4, 6, 6, 5, 5, 5, 6, 4, 5, 0),
			//           list(90, 0, 60, 120, 30, 120, 60, 0, 90, 30, 0))
			//
			// canon_out = list(
			//           list(313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 0),
			//           list(4, 4, 6, 6, 5, 5, 5, 6, 4, 5, 0),
			//           list(90, 0, 60, 120, 30, 120, 60, 0, 90, 30, 0))
			//
			// format(ExtensionReductionProbabilityReaction(canon, canon_out), digits=22, trim=T)
			// [1] "0.01754385964912280604366"
			//
			pcanon:     [][]int{{313}, {313}, {313}, {313}, {313}, {313}, {313}, {313}, {313}, {313}, {0}},
			ccanon:     [][]int{{3}, {3}, {5}, {5}, {4}, {4}, {4}, {5}, {3}, {4}, {0}},
			ocanon:     [][]float64{{90}, {0}, {60}, {120}, {30}, {120}, {60}, {0}, {90}, {30}, {0}},
			pcanon_out: [][]int{{313}, {313}, {313}, {313}, {313}, {313}, {313}, {313}, {313}, {313}, {0}},
			ccanon_out: [][]int{{3}, {3}, {5}, {5}, {4}, {4}, {4}, {5}, {3}, {4}, {0}},
			ocanon_out: [][]float64{{90}, {0}, {60}, {120}, {30}, {120}, {60}, {0}, {90}, {30}, {0}},
			expected:   0.01754385964912280604366,
		},
	}

	for _, tc := range testCases {
		actual := ExtensionReductionProbabilityReaction(
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
