package functions_test

import (
	"reflect"
	"testing"

	. "./"
)

type testCanonicalOrder struct {
	ctemp, chtemp [][]int
	otemp         [][]float64
	csort, chsort [][]int
	osort         [][]float64
}

func TestCanonicalOrder(t *testing.T) {
	SetInitData("./data/PrecursorUnitCell.csv", "./data/PrecursorUnitCellAxes.csv")

	testCases := []testCanonicalOrder{
		{
			ctemp:  [][]int{{0}},
			chtemp: [][]int{{0}},
			otemp:  [][]float64{{0}},
			csort:  [][]int{{0}},
			chsort: [][]int{{0}},
			osort:  [][]float64{{0}},
		},
		{
			// [R] cg = Canonical.Gen(); Canonical.Order(cg)
			// cg: {ctemp, chtemp, otemp}
			//
			// NOTICE: R is 1-base index, golang is 0-base.
			//   chtemp and chsort are characters, 1-base in R but 0-base in Golang
			//
			ctemp:  [][]int{{313, 212, 390, 233}, {313}, {313, 285, 259}, {313}, {0}, {313}, {0}, {0}, {0}, {0}},
			chtemp: [][]int{{3, 6, 3, 4}, {5}, {6, 6, 4}, {6}, {0}, {4}, {0}, {0}, {0}, {0}},
			otemp:  [][]float64{{0, 0, 90, 60}, {60}, {0, 0, 30}, {0}, {0}, {120}, {0}, {0}, {0}, {0}},
			csort:  [][]int{{313, 212, 390, 233}, {313, 285, 259}, {313}, {313}, {313}, {0}},
			chsort: [][]int{{3, 6, 3, 4}, {6, 6, 4}, {5}, {6}, {4}, {0}},
			osort:  [][]float64{{0, 0, 90, 60}, {0, 0, 30}, {60}, {0}, {120}, {0}},
		},
	}

	for _, tc := range testCases {
		csort, chsort, osort := CanonicalOrder(tc.ctemp, tc.chtemp, tc.otemp)
		if !reflect.DeepEqual(csort, tc.csort) {
			t.Errorf("\ngot  %v\nwant %v", csort, tc.csort)
			return
		} else if !reflect.DeepEqual(chsort, tc.chsort) {
			t.Errorf("\ngot  %v\nwant %v", chsort, tc.chsort)
			return
		} else if !reflect.DeepEqual(osort, tc.osort) {
			t.Errorf("\ngot  %v\nwant %v", osort, tc.osort)
			return
		}
	}
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
