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

	var adjX [][]int
	Create2DimArray(&adjX, l, l)
	for k := 0; k < l-1; k++ {
		for j := k + 1; j < l; j++ {
			if Member(xtest[k], SurrAdj([]int{xtest[j]}, Inp.AdjCuml[Npower-1])) {
				adjX[k][j], adjX[j][k] = 1, 1
			}
		}
	}

	_, _, no := GraphClusters(GraphAdjacency(adjX, ADJ_UNDIRECTED))
	if no > 1 {
		return true
	}

	return false
}
