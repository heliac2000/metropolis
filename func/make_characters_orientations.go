//
// make_characters_rientations.go
//

package functions

import (
	"sync"

	. "../util"
)

// Check for overlap in CharactersOrientations when adding unit cell
// Xadd
//
func MakeCharactersOrientations(zcoords [][]float64, xadd []int) [][]int {
	// Make a list of rows in CharactersOrientations to keep. Zcoords is
	// the coordinates of the atoms in Island
	l := len(Inp.CharactersOrientations)
	addO := Copy2DimArrayInt(Inp.CharactersOrientations)
	for i := 0; i < l; i++ {
		addO[i][0] = xadd[i%len(xadd)]
	}

	var wg sync.WaitGroup
	wg.Add(l)
	keep := make([]bool, l)
	for k := 0; k < l; k++ {
		go func(k int) {
			defer wg.Done()
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
			if min > Mcut {
				keep[k] = true
			}
		}(k)
	}
	wg.Wait()

	addOK := make([][]int, 0, l)
	for i, k := range keep {
		if k {
			addOK = append(addOK, addO[i])
		}
	}

	return addOK
}
