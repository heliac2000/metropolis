package functions_test

import (
	"testing"

	. "./"
)

type testCasesInExt struct {
	pcdb, ccdb []int
	ocdb       []float64
	pcbb, ccbb []int
	ocbb       []float64
	expected   bool
}

func TestInExt(t *testing.T) {
	SetInitData("./data/PrecursorUnitCell.csv", "./data/PrecursorUnitCellAxes.csv")

	// NOTICE: R is 1-base index, golang is 0-base.
	//   ccdb and ccbb are characters, 1-base in R but 0-base in Golang
	//
	testCases := []testCasesInExt{
		{
			// [R]
			// ctest = Canonical.Order(Canonical.Gen())
			// cdb = list(ctest[[1]][[1]], ctest[[2]][[1]], ctest[[3]][[1]])
			// cbb = list(ctest[[1]][[2]], ctest[[2]][[2]], ctest[[3]][[2]])
			// extb = Extension.Block(cdb, CoordsIsland(cdb))[[1]]
			// inExt(cdb, cbb, extb)
			//
			pcdb:     []int{313, 338, 415, 261, 337, 239, 282, 292, 253},
			ccdb:     []int{6, 4, 4, 3, 3, 3, 6, 3, 3},
			ocdb:     []float64{0, 120, 60, 90, 0, 120, 150, 0, 90},
			pcbb:     []int{313},
			ccbb:     []int{6},
			ocbb:     []float64{0},
			expected: false,
		},
	}

	for _, tc := range testCases {
		extb, _ := ExtensionBlock(tc.pcdb, CoordsIsland(tc.pcdb, tc.ccdb, tc.ocdb))
		actual := InExt(tc.pcdb, tc.ccdb, tc.ocdb, tc.pcbb, tc.ccbb, tc.ocbb, extb)
		if actual != tc.expected {
			t.Errorf("\ngot  %v\nwant %v", actual, tc.expected)
			return
		}
	}

	return
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
