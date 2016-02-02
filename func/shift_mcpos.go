//
// shift_mcpos.go
//

package functions

import . "../util"

// Put COM of molecule at coordinate X (1 x 2 vector)
//
func ShiftMCpos(moleculeCoords *MoleculeCoordinates, x []float64) [][]float64 {
	shift := Copy2DimArrayFloat(moleculeCoords.All)

	r, sumX, sumY := len(shift), 0.0, 0.0
	for i := 0; i < r; i++ {
		sumX += shift[i][0]
		sumY += shift[i][1]
	}

	y1, y2 := x[0]-sumX/float64(r), x[1]-sumY/float64(r)
	for i := 0; i < r; i++ {
		shift[i][0] += y1
		shift[i][1] += y2
	}

	return shift
}
