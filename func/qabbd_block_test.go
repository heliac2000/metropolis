package functions_test

import (
	"testing"

	. "./"
)

type testCasesQabbdBlock struct {
	pcab, ccab []int
	ocab       []float64
	pcbb, ccbb []int
	ocbb       []float64
	pcdb, ccdb []int
	ocdb       []float64
	i1         int
	canon      [][]int
	expected   float64
}

func TestQabbdBlock(t *testing.T) {
	SetInitData("./data/PrecursorUnitCell.csv", "./data/PrecursorUnitCellAxes.csv")

	// NOTICE: R is 1-base index, golang is 0-base.
	//   ccdb is characters, 1-base in R but 0-base in Golang
	//
	testCases := []testCasesQabbdBlock{
		{
			// [R]
			// canon = Canonical.Order(Canonical.Gen())
			// canon_out = Canonical.Order(ExtensionReductionBlock(canon)[[1]])
			//
			// [NOTICE]
			// ## i1 = 1/i2 = 3(not used in qabbd.Block)/CoutIND = [1, 3]
			// cab  = list(canon[[1]][[1]], canon[[2]][[1]], canon[[3]][[1]])
			// cbb  = list(canon[[1]][[3]], canon[[2]][[3]], canon[[3]][[3]])
			// cdb  = list(canon_out[[1]][[3]], canon_out[[2]][[3]], canon_out[[3]][[3]])
			// reda = Reduction.Block(cab)
			// extb = Extension.Block(cbb, CoordsIsland(cbb))
			// lb   = extb[[2]]
			// extb = extb[[1]]
			//
			// format(qabbd.Block(cab, cbb, cdb, 1, 1, canon, reda, extb, lb), digits=22, trim=T)
			//
			pcab:     []int{313, 361, 364, 408, 317, 264, 340},
			ccab:     []int{6, 4, 4, 6, 6, 4, 5},
			ocab:     []float64{0, 30, 120, 150, 150, 30, 120},
			pcbb:     []int{361, 364, 408, 317, 264, 340},
			ccbb:     []int{4, 4, 6, 6, 4, 5},
			ocbb:     []float64{30, 120, 150, 150, 30, 120},
			pcdb:     []int{313},
			ccdb:     []int{6},
			ocdb:     []float64{60},
			i1:       0,
			canon:    [][]int{{313, 361, 364, 408, 317, 264, 340}, {313, 263, 387}, {0}},
			expected: 0.003888888888888888343415,
		},
	}

	for _, tc := range testCases {
		preda, creda, oreda := ReductionBlock(tc.pcab, tc.ccab, tc.ocab)
		extb, lb := ExtensionBlock(tc.pcbb, CoordsIsland(tc.pcbb, tc.ccbb, tc.ocbb))

		actual := QabbdBlock(tc.pcab, tc.ccab, tc.ocab, tc.pcbb, tc.ccbb, tc.ocbb,
			tc.pcdb, tc.ccdb, tc.ocdb, preda, creda, oreda, tc.canon, extb, lb, tc.i1)

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
