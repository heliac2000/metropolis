//
// surr_adj.go
//

package functions

import (
	"sort"

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
	b2 = Unique(b2)
	sort.Ints(b2)

	return b2
}

// Find the points adjacent to k after expanding by j unit cells and
// exclusing first q cells
//
func SurrAdjEx(k []int, j, q int, AdjCuml [][][]int) []int {
	if len(k) == 0 {
		return k
	}

	boundary := SurrAdj(k, j, AdjCuml)
	boundaryq := SurrAdj(k, q, AdjCuml)

	b2 := make([]int, 0, len(boundary))
	for i := 0; i < len(boundary); i++ {
		if !Member(boundary[i], boundaryq) {
			b2 = append(b2, boundary[i])
		}
	}

	return b2
}
