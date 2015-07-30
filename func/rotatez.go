//
// rotatez.go
//

package functions

import (
	"math"

	. "../util"
)

// Rotate molecule about z axis
//
func RotateZ(MoleculeCoords [][]float64, theta float64) [][]float64 {
	rot := Copy2DimArray(MoleculeCoords).([][]float64)
	r := len(rot)
	c1, c2 := make([]float64, r), make([]float64, r)
	for i := 0; i < r; i++ {
		c1[i], c2[i] = rot[i][0], rot[i][1]
	}
	rcom := []float64{Average(c1), Average(c2)}

	for i := 0; i < r; i++ {
		rot[i][0] = rot[i][0]*math.Cos(theta) - rot[i][1]*math.Sin(theta)
		rot[i][1] = rot[i][0]*math.Sin(theta) + rot[i][1]*math.Cos(theta)
	}

	for i := 0; i < r; i++ {
		c1[i], c2[i] = rot[i][0], rot[i][1]
	}
	rcom2 := []float64{Average(c1) - rcom[0], Average(c2) - rcom[1]}

	for i := 0; i < r; i++ {
		rot[i][0] -= rcom2[0]
		rot[i][1] -= rcom2[1]
	}

	return rot
}
