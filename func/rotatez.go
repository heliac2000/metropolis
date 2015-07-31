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
	rot := Copy2DimArray(moleculeCoords).([][]float64)
	r := len(rot)

	sumX, sumY := 0.0, 0.0
	for i := 0; i < r; i++ {
		sumX += rot[i][0]
		sumY += rot[i][1]
	}
	rcomX, rcomY := sumX/float64(r), sumY/float64(r)

	sumX, sumY = 0.0, 0.0
	for i := 0; i < r; i++ {
		rot[i][0], rot[i][1] =
			rot[i][0]*math.Cos(theta)-rot[i][1]*math.Sin(theta),
			rot[i][0]*math.Sin(theta)+rot[i][1]*math.Cos(theta)
		sumX += rot[i][0]
		sumY += rot[i][1]
	}
	rcomX, rcomY = sumX/float64(r)-rcomX, sumY/float64(r)-rcomY

	for i := 0; i < r; i++ {
		rot[i][0] -= rcomX
		rot[i][1] -= rcomY
	}

	return rot
}
