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
	SetInitData(
		"./data/PrecursorUnitCell.csv", "./data/UnitCell2.csv", "./data/PrecursorUnitCellAxes.csv",
		"./data/kernelregS_Rep_log.json", "./data/kernelregS_Att.json", "./data/svm_model.json")

	// NOTICE: R is 1-base index, golang is 0-base.
	//   ccdb and ccbb are characters, 1-base in R but 0-base in Golang
	//
	testCases := []testCasesInExt{
		{
			// [R]
			// ctest = Canonical.Order(Canonical.Gen())
			// cdb = list(ctest[[1]][[1]], ctest[[2]][[1]], ctest[[3]][[1]])
			// cbb = list(ctest[[1]][[2]], ctest[[2]][[2]], ctest[[3]][[2]])
			// extb = Extension.Block(cbb, CoordsIsland(cbb))[[1]]
			// inExt(cdb, cbb, extb)
			//
			//  cdb  = list(c(313, 338, 415, 261, 337, 239, 282, 292, 253),
			//              c(7, 5, 5, 4, 4, 4, 7, 4, 4),
			//              c(0, 120, 60, 90, 0, 120, 150, 0, 90))
			//  cbb  = list(c(313), c(7), c(0))
			//  extb = Extension.Block(cbb, CoordsIsland(cbb))[[1]]
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
		{
			// [R]
			// ctest = Canonical.Order(Canonical.Gen())
			// cbb = list(ctest[[1]][[2]], ctest[[2]][[2]], ctest[[3]][[2]])
			// extb = Extension.Block(cbb, CoordsIsland(cbb))[[1]]
			// cdb  = cbb
			// eadd = extb[[1]][1,]
			// cdb[[1]] <- c(cdb[[1]], eadd[1])
			// cdb[[2]] <- c(cdb[[2]], eadd[2])
			// cdb[[3]] <- c(cdb[[3]], eadd[3])
			// inExt(cdb, cbb, extb)
			//
			// cdb  = list(c(313, 338, 415, 261, 337, 239, 282, 292, 253, 114),
			//             c(7, 5, 5, 4, 4, 4, 7, 4, 4, 7),
			//             c(0, 120, 60, 90, 0, 120, 150, 0, 90, 0))
			// cbb  = list(c(313, 338, 415, 261, 337, 239, 282, 292, 253),
			//             c(7, 5, 5, 4, 4, 4, 7, 4, 4),
			//             c(0, 120, 60, 90, 0, 120, 150, 0, 90))
			// extb = Extension.Block(cbb, CoordsIsland(cbb))[[1]]
			// inExt(cdb, cbb, extb)
			//
			pcdb:     []int{313, 338, 415, 261, 337, 239, 282, 292, 253, 114},
			ccdb:     []int{6, 4, 4, 3, 3, 3, 6, 3, 3, 6},
			ocdb:     []float64{0, 120, 60, 90, 0, 120, 150, 0, 90, 0},
			pcbb:     []int{313, 338, 415, 261, 337, 239, 282, 292, 253},
			ccbb:     []int{6, 4, 4, 3, 3, 3, 6, 3, 3},
			ocbb:     []float64{0, 120, 60, 90, 0, 120, 150, 0, 90},
			expected: true,
		},
	}

	for _, tc := range testCases {
		//extb, _ := ExtensionBlock(tc.pcdb, CoordsIsland(tc.pcdb, tc.ccdb, tc.ocdb))
		extb, _ := ExtensionBlock(tc.pcbb, CoordsIsland(tc.pcbb, tc.ccbb, tc.ocbb))
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
