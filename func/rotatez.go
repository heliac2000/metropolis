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
	c1, c2, c3 := make([]float64, r), make([]float64, r), make([]float64, r)
	for i := 0; i < r; i++ {
		c1[i], c2[i], c3[i] = rot[i][0], rot[i][1], rot[i][2]
	}
	rcom := []float64{Average(c1), Average(c2), Average(c3)}

	for i := 0; i < len(rot); i++ {
		rot[i][0], rot[i][1] =
			rot[i][0]*math.Cos(theta)-rot[i][1]*math.Sin(theta),
			rot[i][0]*math.Sin(theta)+rot[i][1]*math.Cos(theta)
	}

	for i := 0; i < r; i++ {
		c1[i], c2[i], c3[i] = rot[i][0], rot[i][1], rot[i][2]
	}
	rcom2 := []float64{Average(c1) - rcom[0], Average(c2) - rcom[1], Average(c3) - rcom[2]}

	center := Copy2DimArray(rot).([][]float64)
	for i := 0; i < len(rot); i++ {
		center[i][0] -= rcom2[0]
		center[i][1] -= rcom2[1]
		center[i][2] -= rcom2[2]
	}

	return center
}
