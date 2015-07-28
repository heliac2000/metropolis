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
			ctemp:  [][]int{{1}},
			chtemp: [][]int{{2}},
			otemp:  [][]float64{{3}},
			csort:  [][]int{{1}, {0}},
			chsort: [][]int{{2}, {0}},
			osort:  [][]float64{{3}, {0}},
		},
		{
			// [R] cg = Canonical.Gen(); Canonical.Order(cg)
			//
			// cg: {ctemp, chtemp, otemp}
			//
			ctemp:  [][]int{{313}, {313}, {313}, {313}, {313}, {313}, {313}, {313}, {313}, {313}},
			chtemp: [][]int{{5}, {6}, {4}, {4}, {5}, {7}, {4}, {6}, {5}, {6}},
			otemp:  [][]float64{{60}, {0}, {0}, {0}, {120}, {0}, {120}, {120}, {120}, {0}},
			csort:  [][]int{{313}, {313}, {313}, {313}, {313}, {313}, {313}, {313}, {313}, {313}, {0}},
			chsort: [][]int{{5}, {6}, {4}, {4}, {5}, {7}, {4}, {6}, {5}, {6}, {0}},
			osort:  [][]float64{{60}, {0}, {0}, {0}, {120}, {0}, {120}, {120}, {120}, {0}, {0}},
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
