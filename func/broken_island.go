//
// broken_island.go
//

package functions

import (
	. "../util"
)

// Check that the island Xtest is not broken, i.e., all points are
// contained in neighborhood of another point
//
func BrokenIslandUnitCell(Xtest [][]int, AdjCuml [][][]int) bool {
	l := len(Xtest) * len(Xtest[0])
	if l > 1 {
		xt, r, c := make([]int, 0, l), len(Xtest), len(Xtest[0])
		for i := 0; i < c; i++ {
			for j := 0; j < r; j++ {
				xt = append(xt, Xtest[j][i])
			}
		}

		AdjX := Create2DimArray(int(0), l, l).([][]int)
		for k := 0; k < l-1; k++ {
			for j := k + 1; j < l; j++ {
				if Member(xt[k], SurrAdj([]int{xt[j]}, AdjCuml[Npower-1])) {
					AdjX[k][j], AdjX[j][k] = 1, 1
				}
			}
		}

		_, _, no := GraphClusters(GraphAdjacency(AdjX, ADJ_UNDIRECTED))
		if no > 1 {
			return true
		}
	}

	return false
}
