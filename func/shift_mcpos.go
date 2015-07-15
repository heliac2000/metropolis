//
// shift_mcpos.go
//

package functions

import . "../util"

// Put COM of molecule at coordinate X (1 x 2 vector)
//
func ShiftMCpos(MoleculeCoords *MoleculeCoordinates, X []float64) [][]float64 {
	MoleculeCoordsShift := Copy2DimArray(MoleculeCoords.All).([][]float64)
	r := len(MoleculeCoords.All)
	c1, c2 := make([]float64, r), make([]float64, r)
	for i := 0; i < r; i++ {
		c1[i], c2[i] = MoleculeCoordsShift[i][0], MoleculeCoordsShift[i][1]
	}

	y1, y2 := X[0]-Average(c1), X[1]-Average(c2)
	for i := 0; i < r; i++ {
		MoleculeCoordsShift[i][0] += y1
		MoleculeCoordsShift[i][1] += y2
	}

	return MoleculeCoordsShift
}
