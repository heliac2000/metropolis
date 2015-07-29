//
// make_characters_rientations.go
//

package functions

import (
	. "../util"
)

// Check for overlap in CharactersOrientations when adding unit cell
// Xadd
//
func MakeCharactersRientations(zcoords [][]float64, xadd []int) [][]int {
	// Make a list of rows in CharactersOrientations to keep. Zcoords is
	// the coordinates of the atoms in Island
	l := len(Inp.CharactersOrientations)
	addO := Copy2DimArray(Inp.CharactersOrientations).([][]int)
	for i := 0; i < l; i++ {
		addO[i][0] = xadd[i%len(xadd)]
	}

	mcut, keep := float64(0), make([]int, 0, l)
	for k := 0; k < l; k++ {
		distKnnx, _ := GetKnnx(zcoords,
			CoordsIsland(addO[k][0:1], addO[k][1:2], []float64{float64(addO[k][2])}), 1)

		min := MinFloat(distKnnx[0]...)
		if len(distKnnx) > 1 {
			for i := 1; i < len(distKnnx); i++ {
				if m := MinFloat(distKnnx[i]...); m < min {
					min = m
				}
			}
		}
		if min > mcut {
			keep = append(keep, k)
		}
	}

	addOK := make([][]int, 0, len(keep))
	for _, k := range keep {
		addOK = append(addOK, addO[k])
	}

	return addOK
}
