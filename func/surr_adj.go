//
// surr_adj.go
//

package functions

import (
	. "../util"
)

// Find the points adjacent to k after expanding by j unit cells
//
func SurrAdj(k []int, j int, AdjCuml [][][]int) []int {
	if len(k) == 0 {
		return k
	}

	AdjCj := AdjCuml[j]
	boundary := make([]int, 0, len(AdjCj)*len(AdjCj[0]))

	for i := 0; i < len(k); i++ {
		ind := k[i]
		for j := 0; j < len(AdjCj[ind]); j++ {
			if AdjCj[ind][j] > 0 {
				boundary = append(boundary, j)
			}
		}
	}

	b2 := make([]int, 0, len(boundary))
	for i := 0; i < len(boundary); i++ {
		if !Member(boundary[i], k) {
			b2 = append(b2, boundary[i])
		}
	}

	return Unique(b2)
}
