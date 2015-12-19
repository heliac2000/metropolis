//
// island_symmetry_block.go
//

package functions

import (
	. "../util"
)

// Compute the symmetry of block CaB
//
func IslandSymmetryBlock(cab []int) float64 {
	if cab[0] == 0 {
		return 2
	}

	l := len(cab)
	xtest := make([][]float64, 0, l)
	for i := 0; i < l; i++ {
		xtest = append(xtest, CopyVectorFloat(Inp.UnitCellCoords[cab[i]]))
	}

	xtestr := make([][]float64, 0, l)
	for _, v := range RotationBlock(cab) {
		xtestr = append(xtestr, CopyVectorFloat(Inp.UnitCellCoords[v]))
	}

	rmn, cnt := make([]int, len(xtestr)), 0
	for k := 0; k < len(xtest); k++ {
		for j := 0; j < len(xtestr); j++ {
			if rmn[j] == 0 &&
				xtest[k][0] == xtestr[j][0] && xtest[k][1] == xtestr[j][1] {
				rmn[j], cnt = 1, cnt+1
				break
			}
		}
	}

	if cnt == len(xtestr) || len(xtestr) == 1 {
		return 2
	}

	return 1
}
