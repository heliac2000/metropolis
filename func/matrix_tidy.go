//
// matrix_tidy.go
//

package functions

import "math"

// Remove duplicated rows from the matrix X, or coordinates within
// epsilon
//
func withinEpsilon(zk, zj []float64) bool {
	return math.Sqrt(
		(zk[0]-zj[0])*(zk[0]-zj[0])+(zk[1]-zj[1])*(zk[1]-zj[1])) < Epsilon
}

func MatrixTidy(z [][]float64) ([][]float64, []int) {
	l := len(z)
	remove := make([]int, l)

	for k := 0; k < l-1; k++ {
		for j := k + 1; j < l; j++ {
			if remove[j] == 0 && withinEpsilon(z[k], z[j]) {
				remove[j] = 1
			}
		}
	}

	z2 := make([][]float64, 0, l)
	z2 = append(z2, z[0])
	for j := 1; j < l; j++ {
		if remove[j] == 0 {
			z2 = append(z2, z[j])
		}
	}

	return z2, remove
}
