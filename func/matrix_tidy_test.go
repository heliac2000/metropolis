package functions_test

import (
	"reflect"
	"testing"
)
import . "./"

type testCasesMatrixTidy struct {
	Z, Z2   [][]float64
	removed []int
}

func TestMatrixTidy(t *testing.T) {
	testCases := []testCasesMatrixTidy{
		{
			// [R] MatrixTidy(t(array(c(1,2,5,7,1,2), c(2,3))))
			//
			Z:       [][]float64{{1, 2}, {5, 7}, {1, 2}},
			Z2:      [][]float64{{1, 2}, {5, 7}},
			removed: []int{0, 0, 1},
		},
		{
			// [R] MatrixTidy(t(array(c(1,2,5,7,1,2,5,7.4,5,7.9), c(2,5))))
			//
			Z:       [][]float64{{1, 2}, {5, 7}, {1, 2}, {5, 7.4}, {5, 7.9}},
			Z2:      [][]float64{{1, 2}, {5, 7}, {5, 7.9}},
			removed: []int{0, 0, 1, 1, 0},
		},
	}

	var z2 [][]float64
	var r []int
	for _, z := range testCases {
		z2, r = MatrixTidy(z.Z)
		if !reflect.DeepEqual(z2, z.Z2) || !reflect.DeepEqual(r, z.removed) {
			t.Errorf("\ngot  %v, %v\nwant %v, %v", z2, r, z.Z2, z.removed)
			return
		}
	}
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
