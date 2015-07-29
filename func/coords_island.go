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
	coords1 := Create2DimArray(float64(0), l, 3).([][]float64)
	for i := 0; i < l; i++ {
		coords1[i][0] = Inp.UnitCellCoords[coutX[i]][0] - (Inp.UnitCell2[3][1] - Inp.UnitCell2[coutC[i]][1])
		coords1[i][1] = Inp.UnitCellCoords[coutX[i]][1] - (Inp.UnitCell2[3][1] - Inp.UnitCell2[coutC[i]][2])
	}

	IslandTemp := make([][]float64, 0, l*len(Inp.MoleculeCoordinates.All))
	for j := 0; j < l; j++ {
		M1 := ShiftMCpos(Inp.MoleculeCoordinates, coords1[j])
		IslandTemp = append(IslandTemp, RotateZ(M1, coutO[j]*math.Pi/180.0)...)
	}

	return IslandTemp
}
