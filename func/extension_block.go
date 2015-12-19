//
// extension_block.go
//

package functions

import (
	. "../util"
)

// Generate the extension of Island. Island is in form of (unit cell
// labels, characters, labels)
//
func ExtensionBlock(xtest []int, zcoords [][]float64) ([][][]int, int) {
	// Count the number of elements in the extension set
	var xadd [][]int

	lx := 0
	xtestAppend := make([][][]int, 0, len(Inp.AdjCuml[Npower-1]))
	if len(xtest) == 1 && xtest[0] == 0 {
		Copy2DimArray(&xadd, Inp.CharactersOrientations)
		for i := 0; i < len(xadd); i++ {
			xadd[i][0] = UCcenter
		}
		xtestAppend = append(xtestAppend, xadd)
		lx = len(Inp.CharactersOrientations)
	} else {
		xsurr := Unique(SurrAdj(xtest, Inp.AdjCuml[Npower-1]))
		// Is there a way to get rid of this loop?
		for k := 0; k < len(xsurr); k++ {
			// Make the CharacersOrientations based checking for overlap
			// condition. Must make this faster
			xadd := MakeCharactersOrientations(zcoords, []int{xsurr[k]})
			if len(xadd) > 0 && len(xadd)*len(xadd[0]) > 0 {
				lx += len(xadd)
				for i := 0; i < len(xadd); i++ {
					xadd[i][0] = xsurr[k]
				}
				xtestAppend = append(xtestAppend, xadd)
			}
		}
	}

	return xtestAppend, lx
}
