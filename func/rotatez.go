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
func RotateZ(moleculeCoords [][]float64, theta float64) [][]float64 {
	var rot [][]float64

	Copy2DimArray(&rot, moleculeCoords)
	l, sumX, sumY := len(rot), 0.0, 0.0
	for i := 0; i < l; i++ {
		sumX += rot[i][0]
		sumY += rot[i][1]
	}
	rcomX, rcomY := sumX/float64(l), sumY/float64(l)

	sumX, sumY = 0.0, 0.0
	for i := 0; i < l; i++ {
		rot[i][0], rot[i][1] =
			rot[i][0]*math.Cos(theta)-rot[i][1]*math.Sin(theta),
			rot[i][0]*math.Sin(theta)+rot[i][1]*math.Cos(theta)
		sumX += rot[i][0]
		sumY += rot[i][1]
	}
	rcomX, rcomY = sumX/float64(l)-rcomX, sumY/float64(l)-rcomY

	for i := 0; i < l; i++ {
		rot[i][0] -= rcomX
		rot[i][1] -= rcomY
	}

	return rot
}
