//
// qabba_block.go
//

package functions

// 1. Compute the probability q_{ab,ba}^{k-->j} that CaB is reduced to
// give CbB, and CbB is extended to give CaB.
//
func QabbaBlock(pcab, ccab []int, ocab []float64,
	pcbb, ccbb []int, ocbb []float64, preda, creda [][]int, oreda [][]float64,
	canon [][]int, extb [][][]int, lb, i1 int) float64 {

	// Turn to true if Cc is in the reduction set of Ca, and Cd is in
	// the extension set of Cb
	//
	if isReductionSet(preda, creda, oreda, pcbb, ccbb, ocbb) &&
		InExt(pcab, ccab, ocab, pcbb, ccbb, ocbb, extb) {
		return computeContribution(pcbb, pcab, canon, i1, lb, len(preda))
	}

	return 0.0
}

// 1. Compute the probability q_{ab,0a}^{k-->j} that CaB is reduced to
// give zero island, and CbB is extended to give Cda.
//
func Qab0aBlock(pcab, ccab []int, ocab []float64,
	pcbb, ccbb []int, ocbb []float64, preda, creda [][]int, oreda [][]float64,
	canon [][]int, extb [][][]int, lb, i1 int) float64 {

	// Turn to true if Cc is in the reduction set of Ca, and Cd is in
	// the extension set of Cb
	//
	if isReductionSet(preda, creda, oreda, []int{0}, []int{0}, []float64{0}) &&
		InExt(pcab, ccab, ocab, pcbb, ccbb, ocbb, extb) {
		return computeContribution(pcbb, pcab, canon, i1, lb, len(preda))
	}

	return 0.0
}
