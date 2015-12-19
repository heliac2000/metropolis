//
// reduction_block.go
//

package functions

// Generate the reduction of Island. Island is in form of (unit cell
// labels, characters, labels). nlevel is the level of the boundary to
// consider.
//
func ReductionBlock(xtest, ctest []int, otest []float64) (
	xremOut, cremOut [][]int, oremOut [][]float64) {
	if len(xtest) == 1 {
		return [][]int{{0}}, [][]int{{0}}, [][]float64{{0}}
	}

	l := len(xtest)
	xremOut, cremOut, oremOut =
		make([][]int, 0, l), make([][]int, 0, l), make([][]float64, 0, l)
	for k := 0; k < l; k++ {
		xout := make([]int, 0, l)
		for _, v := range xtest {
			if v != xtest[k] {
				xout = append(xout, v)
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
		if len(cind) == 0 {
			continue
		}

		oind := make([]int, 0, len(otest))
		for i := 0; i < len(otest); i++ {
			if i != k {
				oind = append(oind, i)
			}
		}
		if len(oind) == 0 {
			continue
		}

		xo, co, oo := make([]int, l-1), make([]int, l-1), make([]float64, l-1)
		for i := 0; i < l-1; i++ {
			xo[i] = xout[i%len(xout)]
			co[i] = ctest[cind[i%len(cind)]]
			oo[i] = otest[oind[i%len(oind)]]
		}
		xremOut = append(xremOut, xo)
		cremOut = append(cremOut, co)
		oremOut = append(oremOut, oo)
	}

	return
}
