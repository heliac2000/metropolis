//
// coords_island.go
//

package functions

import (
	"math"
)

// Output coordinates of island
//
// CoutX: Positions(index)
// CoutC: Characters(index)
// CoutO: Orientations(angle)
func CoordsIsland(coutX, coutC []int, coutO []float64) [][]float64 {
	l := len(coutX)
	islandTemp := make([][]float64, 0, l*len(Inp.MoleculeCoordinates.All))
	coords := make([]float64, 2)

	for i := 0; i < l; i++ {
		coords[0] = Inp.UnitCellCoords[coutX[i]][0] - (Inp.UnitCell2[3][1] - Inp.UnitCell2[coutC[i]][1])
		coords[1] = Inp.UnitCellCoords[coutX[i]][1] - (Inp.UnitCell2[3][1] - Inp.UnitCell2[coutC[i]][2])
		islandTemp = append(islandTemp,
			RotateZ(ShiftMCpos(Inp.MoleculeCoordinates, coords), coutO[i]*math.Pi/180.0)...)
	}

	return islandTemp
}
