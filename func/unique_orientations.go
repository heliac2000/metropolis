//
// unique_orientations.go
//

package functions

import . "../util"

// Determine the unique orientations of island Xtest
//
func UniqueOrientations(pos, chr []int, ori []float64) ([]int, []int) {
	// Strictly assumes two-fold symmetry
	ir := RotationBlock(pos)

	//Delete indices that appear more than once

	// Find center of masses of Ir[[1]] and Ir[[2]], superimpose COMs,
	// then check if thare are the same.
	x1 := make([][]float64, len(pos))
	for i, v := range pos {
		x1[i] = CopyVectorFloat(Inp.UnitCellCoords[v])
	}
	x2 := make([][]float64, len(ir))
	for i, v := range ir {
		x2[i] = CopyVectorFloat(Inp.UnitCellCoords[v])
	}

	if len(pos) > 1 {
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
		ind1, ind2 := SortIndexReturn(x1MoveTemp), SortIndexReturn(ir)
		for k := 0; k < len(ind1); k++ {
			i1, i2 := ind1[k], ind2[k]
			if x1MoveTemp[i1] != ir[i1] ||
				chr[i1] != chr[i2] || ori[i1] != ori[i2] {
				return []int{0, 1}, ir
			}
		}
	}

	return []int{0}, ir
}
