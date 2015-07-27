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
func BrokenIslandUnitCell(xtest []int) bool {
	l := len(xtest)
	if l <= 1 {
		return false
	}

	AdjX := Create2DimArray(int(0), l, l).([][]int)
	for k := 0; k < l-1; k++ {
		for j := k + 1; j < l; j++ {
			if Member(xtest[k], SurrAdj([]int{xtest[j]}, Inp.AdjCuml[Npower-1])) {
				AdjX[k][j], AdjX[j][k] = 1, 1
			}
		}
	}

	_, _, no := GraphClusters(GraphAdjacency(AdjX, ADJ_UNDIRECTED))
	if no > 1 {
		return true
	}

	return false
}
