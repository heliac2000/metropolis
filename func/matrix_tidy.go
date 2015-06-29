//
// matrix_tidy.go
//

package functions

import "math"

// Remove duplicated rows from the matrix X, or coordinates within
// epsilon
//
func MatrixTidy(Z [][]float64) ([][]float64, []int) {
	remove := make([]int, len(Z))
	for k := 0; k < len(Z)-1; k++ {
		for j := k + 1; j < len(Z); j++ {
			if remove[j] == 0 {
				if math.Sqrt((Z[k][0]-Z[j][0])*(Z[k][0]-Z[j][0])+
					(Z[k][1]-Z[j][1])*(Z[k][1]-Z[j][1])) < Epsilon {
					remove[j] = 1
				}
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
