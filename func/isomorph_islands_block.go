//
// isomorph_islands_block.go
//

package functions

import (
	. "../util"
)

// Check if islands Xtest1 and Xtest2 are rotational isomorphs
//
func IsomorphIslandsBlock(xtest1, ctest1 []int, otest1 []float64,
	xtest2, ctest2 []int, otest2 []float64) bool {
	isomorph := true
	l1, l2 := len(xtest1), len(xtest2)

	if l1 != l2 || (l1 == 1 && xtest1[0] == 0) || (l2 == 1 && xtest2[0] == 0) {
		isomorph = false
	}

	// Compare the characters and orientations of the points in Xtest1
	// and Xtest2
	if isomorph {
		cnt, rmv := 0, make([]int, len(ctest2))
		for i := 0; i < len(ctest1); i++ {
			for j := 0; j < len(ctest2); j++ {
				if rmv[j] == 0 && ctest1[i] == ctest2[j] && otest1[i] == otest2[j] {
					cnt, rmv[j] = cnt+1, 1
					break
				}
			}
		}
		if cnt != len(ctest1) {
			isomorph = false
		}
	}

	if isomorph && len(xtest1) > 1 {
		xtest1C := make([][]float64, l1)
		xtest1Ccol1, xtest1Ccol2 := make([]float64, l1), make([]float64, l1)
		for i := 0; i < l1; i++ {
			xtest1C[i] = Inp.UnitCellCoords[xtest1[i]]
			xtest1Ccol1[i], xtest1Ccol2[i] = xtest1C[i][0], xtest1C[i][1]
		}
		xtest2C := make([][]float64, l2)
		xtest2Ccol1, xtest2Ccol2 := make([]float64, l2), make([]float64, l2)
		for i := 0; i < l2; i++ {
			xtest2C[i] = Inp.UnitCellCoords[xtest2[i]]
			xtest2Ccol1[i], xtest2Ccol2[i] = xtest2C[i][0], xtest2C[i][1]
		}
		comX1 := [][]float64{{Average(xtest1Ccol1), Average(xtest1Ccol2)}}
		comX2 := [][]float64{{Average(xtest2Ccol1), Average(xtest2Ccol2)}}
		comX1pt := func(_ [][]float64, ind [][]int) int { return ind[0][0] }(GetKnnx(xtest1C, comX1, 1))
		comX2pt := func(_ [][]float64, ind [][]int) int { return ind[0][0] }(GetKnnx(xtest2C, comX2, 1))
		deltaX := xtest1C[comX1pt][0] - xtest2C[comX2pt][0]
		deltaY := xtest1C[comX1pt][1] - xtest2C[comX2pt][1]

		// Move Island1 so that its center of mass is over the center of
		// mass of Island2
		xtest1CMove := Copy2DimArrayFloat(xtest1C)
		for i := 0; i < len(xtest1CMove); i++ {
			xtest1CMove[i][0] -= deltaX
			xtest1CMove[i][1] -= deltaY
		}
		_, ind := GetKnnx(Inp.UnitCellCoords, xtest1CMove, 1)
		xtest1MoveTemp := make([]int, len(ind))
		for i := 0; i < len(ind); i++ {
			xtest1MoveTemp[i] = ind[i][0]
		}

		cnt, rmv := 0, make([]int, len(xtest1MoveTemp))
		for i := 0; i < len(xtest1MoveTemp); i++ {
			for j := 0; j < l2; j++ {
				if rmv[j] == 0 && xtest1MoveTemp[i] == xtest2[j] &&
					ctest1[i] == ctest2[j] && otest1[i] == otest2[j] {
					cnt, rmv[j] = cnt+1, 1
					break
				}
			}
		}

		// Now try the rotation about the center of mass
		if cnt < l2 {
			xtest1Rotate := RotationBlock(xtest1MoveTemp)
			cnt, rmv := 0, make([]int, len(xtest1Rotate))
			for i := 0; i < len(xtest1Rotate); i++ {
				for j := 0; j < l2; j++ {
					if rmv[j] == 0 && xtest1Rotate[i] == xtest2[j] &&
						ctest1[i] == ctest2[j] && otest1[i] == otest2[j] {
						cnt, rmv[j] = cnt+1, 1
						break
					}
				}
			}
			if cnt < l2 {
				isomorph = false
			}
		}
	}

	if l1 == 1 && l2 == 1 &&
		((xtest1[0] == 0 && xtest2[0] == 0) ||
			(ctest1[0] == ctest2[0] && otest1[0] == otest2[0])) {
		return true
	}

	return isomorph
}
