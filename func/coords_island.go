//
// coords_island.go
//

package functions

import (
	"math"

	. "../util"
)

// Output coordinates of island
//
// CoutX: Positions(index)
// CoutC: Characters(index)
// CoutO: Orientations(angle)
func CoordsIsland(coutX, coutC []int, coutO []float64) [][]float64 {
	l := len(coutX)
	coords := Create2DimArray(float64(0), l, 3).([][]float64)
	for i := 0; i < l; i++ {
		coords[i][0] = Inp.UnitCellCoords[coutX[i]][0] - (Inp.UnitCell2[3][1] - Inp.UnitCell2[coutC[i]][1])
		coords[i][1] = Inp.UnitCellCoords[coutX[i]][1] - (Inp.UnitCell2[3][1] - Inp.UnitCell2[coutC[i]][2])
	}

	islandTemp := make([][]float64, 0, l*len(Inp.MoleculeCoordinates.All))
	for j := 0; j < l; j++ {
		m := ShiftMCpos(Inp.MoleculeCoordinates, coords[j])
		islandTemp = append(islandTemp, RotateZ(m, coutO[j]*math.Pi/180.0)...)
	}

	return islandTemp
}
