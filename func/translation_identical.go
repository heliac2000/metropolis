//
// translation_identical.go
//

package functions

import (
	"reflect"

	. "../util"
)

// Determine if two islands are identical upon translation
//
func TranslationIdentical(pos1, chr1 []int, ori1 []float64, pos2, chr2 []int, ori2 []float64) bool {
	// Find center of masses of Ir[[1]] and Ir[[2]], superimpose COMs,
	// then check if thare are the same.
	if len(pos1) == 1 &&
		(!reflect.DeepEqual(chr1, chr2) || !reflect.DeepEqual(ori1, ori2)) {
		return false
	}

	if len(pos1) > 1 {
		x1 := make([][]float64, len(pos1))
		for i, v := range pos1 {
			x1[i] = CopyVectorFloat(Inp.UnitCellCoords[v])
		}
		x2 := make([][]float64, len(pos2))
		for i, v := range pos2 {
			x2[i] = CopyVectorFloat(Inp.UnitCellCoords[v])
		}

		x11, x12 := 0.0, 0.0
		for _, v := range x1 {
			x11 += v[0]
			x12 += v[1]
		}
		comX1 := [][]float64{{x11 / float64(len(x1)), x12 / float64(len(x1))}}

		x21, x22 := 0.0, 0.0
		for _, v := range x2 {
			x21 += v[0]
			x22 += v[1]
		}
		comX2 := [][]float64{{x21 / float64(len(x2)), x22 / float64(len(x2))}}

		_, idx1 := GetKnnx(x1, comX1, 1)
		_, idx2 := GetKnnx(x2, comX2, 1)
		deltaX := x1[idx1[0][0]][0] - x2[idx2[0][0]][0]
		deltaY := x1[idx1[0][0]][1] - x2[idx2[0][0]][1]

		// Move Island1 so that its center of mass is over the center of
		// mass of Island2
		for _, v := range x1 {
			v[0] -= deltaX
			v[1] -= deltaY
		}
		_, idx := GetKnnx(Inp.UnitCellCoords, x1, 1)
		x1MoveTemp := make([]int, len(idx))
		for i, v := range idx {
			x1MoveTemp[i] = v[0]
		}

		// Sort the islands in order of unit cel index
		ind1, ind2 := SortIndexReturn(x1MoveTemp), SortIndexReturn(pos2)

		for k := 0; k < len(ind1); k++ {
			if x1MoveTemp[ind1[k]] != pos2[ind2[k]] ||
				chr1[ind1[k]] != chr2[ind2[k]] || ori1[ind1[k]] != ori2[ind2[k]] {
				return false
			}
		}
	}

	return true
}
