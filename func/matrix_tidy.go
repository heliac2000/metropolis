//
// matrix_tidy.go
//

package functions

import "math"

// Remove duplicated rows from the matrix X, or coordinates within
// epsilon
//
func withinEpsilon(Zk, Zj []float64) bool {
	return math.Sqrt(
		(Zk[0]-Zj[0])*(Zk[0]-Zj[0])+(Zk[1]-Zj[1])*(Zk[1]-Zj[1])) < Epsilon
}

func MatrixTidy(Z [][]float64) ([][]float64, []int) {
	remove := make([]int, len(Z))
	for k := 0; k < len(Z)-1; k++ {
		for j := k + 1; j < len(Z); j++ {
			if remove[j] == 0 && withinEpsilon(Z[k], Z[j]) {
				remove[j] = 1
			}
		}
	}

	Z2 := make([][]float64, 0, len(Z))
	Z2 = append(Z2, Z[0])
	for j := 1; j < len(Z); j++ {
		if remove[j] == 0 {
			Z2 = append(Z2, Z[j])
		}
	}

	return Z2, remove
}
