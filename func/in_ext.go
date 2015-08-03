//
// in_ext.go
//

package functions

// Rapid check if Cd is contained in the extension set of Cb (Extb)
//
func InExt(pcdb, ccdb []int, ocdb []float64,
	pcbb, ccbb []int, ocbb []float64, extb [][][]int) bool {

	// Check that character and orientation of points in CbB are subset
	// of those in CdB, and output the remaining
	l := len(ccdb)
	remv := make([]int, l)
	for k := 0; k < len(ccbb); k++ {
		for j := 0; j < l; j++ {
			if remv[j] == 0 && ccbb[k] == ccdb[j] && ocbb[k] == ocdb[j] {
				remv[j] = 1
				break
			}
		}
	}

	keep := make([]int, 0, l)
	for i := 0; i < l; i++ {
		if remv[i] == 0 {
			keep = append(keep, i)
		}
	}
	if len(keep) != 1 {
		return false
	}

	ptest, ctest, otest :=
		make([]int, 0, len(pcbb)+1), make([]int, 0, len(ccbb)+1), make([]float64, 0, len(ocbb)+1)
	if pcbb[0] != 0 {
		ptest, ctest, otest =
			append(ptest, pcbb...), append(ctest, ccbb...), append(otest, ocbb...)
	}
	ctest, otest = append(ctest, ccdb[keep[0]]), append(otest, ocdb[keep[0]])

	ptest1 := make([]int, 0, len(ptest)+1)
	for k := 0; k < len(extb); k++ {
		ptest1 = append(ptest, extb[k][0][0])
		if IsomorphIslandsBlock(pcdb, ccdb, ocdb, ptest1, ctest, otest) {
			return true
		}
		ptest1 = nil
	}

	return false
}
