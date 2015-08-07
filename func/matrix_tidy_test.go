package functions_test

import (
	"reflect"
	"testing"
)
import . "./"

type testCasesMatrixTidy struct {
	z, z2   [][]float64
	removed []int
}

func TestMatrixTidy(t *testing.T) {
	testCases := []testCasesMatrixTidy{
		{
			// [R] MatrixTidy(t(array(c(1,2,5,7,1,2), c(2,3))))
			//
			z:       [][]float64{{1, 2}, {5, 7}, {1, 2}},
			z2:      [][]float64{{1, 2}, {5, 7}},
			removed: []int{0, 0, 1},
		},
		{
			// [R] MatrixTidy(t(array(c(1,2,5,7,1,2,5,7.4,5,7.9), c(2,5))))
			//
			z:       [][]float64{{1, 2}, {5, 7}, {1, 2}, {5, 7.4}, {5, 7.9}},
			z2:      [][]float64{{1, 2}, {5, 7}, {5, 7.9}},
			removed: []int{0, 0, 1, 1, 0},
		},
	}

	var z2 [][]float64
	var r []int
	for _, z := range testCases {
		z2, r = MatrixTidy(z.z)
		if !reflect.DeepEqual(z2, z.z2) || !reflect.DeepEqual(r, z.removed) {
			t.Errorf("\ngot  %v, %v\nwant %v, %v", z2, r, z.z2, z.removed)
			return
		}
	}
}

// Local Variables:
// compile-command: (concat "go test -gcflags='-B' -v " (file-name-nondirectory buffer-file-name))
// End:
