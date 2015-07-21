//
// rotation_block.go
//

package functions

import (
	"math"

	. "../util"
)

// Rotational isomorphism check
//
// Perform a TWO-FOLD rotation of an block about center of
// mass. Island is object of type output by RandomIslandUnitCell
//
// islandP: Island block positions(index)
// islandC: Island block characters(index)
// islandO: Island block orientations(angle)
func RotationBlock(islandP []int) []int {
	// Fetch the island coordinates
	r, c := len(islandP), len(Inp.UnitCellCoords[0])
	xtestC := Create2DimArray(float64(0), r, c).([][]float64)
	xcom := Create2DimArray(float64(0), c, r).([][]float64)
	for i, p := range islandP {
		copy(xtestC[i], Inp.UnitCellCoords[p])
		for j := 0; j < c; j++ {
			xcom[j][i] = Inp.UnitCellCoords[p][j]
		}
	}

	// Island center of mass
	xcom = [][]float64{{Average(xcom[0]), Average(xcom[1])}}

	// Find block closest to center of mass
	_, ind := GetKnnx(xtestC, xcom, 1)
	xcom2 := ind[0][0]

	xtestR := Create2DimArray(float64(0), r, c).([][]float64)
	for i := 0; i < r; i++ {
		xtestR[i][0] = xtestC[i][0]*math.Cos(math.Pi) - xtestC[i][1]*math.Sin(math.Pi)
		xtestR[i][1] = xtestC[i][0]*math.Sin(math.Pi) + xtestC[i][1]*math.Cos(math.Pi)
	}

	// Assuming this is unchanged on rotation
	deltaX, deltaY := xtestC[xcom2][0]-xtestR[xcom2][0], xtestC[xcom2][1]-xtestR[xcom2][1]
	for i := 0; i < r; i++ {
		xtestR[i][0] += deltaX
		xtestR[i][1] += deltaY
	}

	_, ind = GetKnnx(Inp.UnitCellCoords, xtestR, 1)
	islandPR := make([]int, len(ind))
	for i := 0; i < len(ind); i++ {
		islandPR[i] = ind[i][0]
	}

	return islandPR
}
