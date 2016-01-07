//
// coords_gen.go
//

package functions

import "math"

// Turn pair into molecule coordinates
//
func CoordsGen(k1, k2, ch1, ch2 int, o1, o2 float64) [][]float64 {
	coutX, coutC, coutO := []int{k1 - 1, k2 - 1}, []int{ch1, ch2}, []float64{o1, o2}
	l := len(coutX)
	islandTemp := make([][]float64, 0, l*len(Inp.MoleculeCoordinates.All))
	coords := make([]float64, 2)

	for i := 0; i < l; i++ {
		coords[0] = Inp.UnitCellCoords[coutX[i]][0] - (Inp.UnitCell[CentralPoint][1] - Inp.UnitCell[coutC[i]][1])
		coords[1] = Inp.UnitCellCoords[coutX[i]][1] - (Inp.UnitCell[CentralPoint][2] - Inp.UnitCell[coutC[i]][2])
		islandTemp = append(islandTemp,
			RotateZ(ShiftMCpos(Inp.MoleculeCoordinates, coords), coutO[i]*math.Pi/180.0)...)
	}

	return islandTemp
}
