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
// CoutX: Positions
// CoutC: Characters
// CoutO: Orientations
func CoordsIsland(CoutX, CoutC, CoutO []int, UnitCell2, UnitCellCoords [][]float64,
	MoleculeCoords *MoleculeCoordinates) [][]float64 {
	l := len(CoutX)
	coords1 := Create2DimArray(float64(0), l, 3).([][]float64)
	for i := 0; i < l; i++ {
		coords1[i][0] = UnitCellCoords[CoutX[i]][0] - (UnitCell2[3][1] - UnitCell2[CoutC[i]][1])
		coords1[i][1] = UnitCellCoords[CoutX[i]][1] - (UnitCell2[3][1] - UnitCell2[CoutC[i]][2])
	}

	IslandTemp := make([][]float64, 0, l*len(MoleculeCoords.All))
	for j := 0; j < l; j++ {
		M1 := ShiftMCpos(MoleculeCoords, coords1[j])
		IslandTemp = append(IslandTemp, RotateZ(M1, float64(CoutO[j])*math.Pi/180.0)...)
	}

	return IslandTemp
}
