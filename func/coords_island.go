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
func CoordsIsland(CoutX, CoutC, CoutO []int) [][]float64 {
	l := len(CoutX)
	coords1 := Create2DimArray(float64(0), l, 3).([][]float64)
	for i := 0; i < l; i++ {
		coords1[i][0] = Inp.UnitCellCoords[CoutX[i]][0] - (Inp.UnitCell2[3][1] - Inp.UnitCell2[CoutC[i]][1])
		coords1[i][1] = Inp.UnitCellCoords[CoutX[i]][1] - (Inp.UnitCell2[3][1] - Inp.UnitCell2[CoutC[i]][2])
	}

	IslandTemp := make([][]float64, 0, l*len(Inp.MoleculeCoordinates.All))
	for j := 0; j < l; j++ {
		M1 := ShiftMCpos(Inp.MoleculeCoordinates, coords1[j])
		IslandTemp = append(IslandTemp, RotateZ(M1, float64(CoutO[j])*math.Pi/180.0)...)
	}

	return IslandTemp
}
