//
// reduction_block.go
//

package functions

import (
	. "../util"
)

// Generate the reduction of Island. Island is in form of (unit cell
// labels, characters, labels). nlevel is the level of the boundary to
// consider.
//
func ReductionBlock(xtest, ctest, otest []int) [][][]int {
	if len(xtest) == 1 {
		return [][][]int{{{0, 0, 0}}}
	}

	l := len(xtest)
	xtestReduction := make([][][]int, 0, l)

	for k := 0; k < l; k++ {
		xout := make([]int, 0, l)
		for i := 0; i < l; i++ {
			if xtest[i] != xtest[k] {
				xout = append(xout, xtest[i])
			}
		}
		if BrokenIslandUnitCell(xout) {
			continue
		}

		cind := make([]int, 0, len(ctest))
		for i := 0; i < len(ctest); i++ {
			if i != k {
				cind = append(cind, i)
			}
		}
		oind := make([]int, 0, len(otest))
		for i := 0; i < len(otest); i++ {
			if i != k {
				oind = append(oind, i)
			}
		}
		xremOut := Create2DimArray(int(0), l-1, 3).([][]int)
		for i := 0; i < l-1; i++ {
			xremOut[i][0] = xout[i%len(xout)]
			xremOut[i][1] = ctest[cind[i%len(cind)]]
			xremOut[i][2] = otest[oind[i%len(oind)]]
		}
		xtestReduction = append(xtestReduction, xremOut)
	}

	return xtestReduction
}
