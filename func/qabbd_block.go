//
// qabbd_block.go
//

package functions

// Compute the probability q_{ab,bd}^{k-->j}. Inputs are islands
// (positions, characters, and orientations)
//
func QabbdBlock(pcab, ccab []int, ocab []float64,
	pcbb, ccbb []int, ocbb []float64, pcdb, ccdb []int, ocdb []float64,
	preda, creda [][]int, oreda [][]float64,
	canon [][]int, extb [][][]int, lb, i1 int) float64 {

	// 1. Compute the probability q_{ab,bd}^{k-->j} that CaB is reduced
	// to give CbB, and CbB is extended to give CdB.
	//
	// Turn to true if Cc is in the reduction set of Ca, and Cd is in
	// the extension set of Cb
	//
	// Compute the contribution to the probability
	//
	qtot := 0.0
	if isReductionSet(preda, creda, oreda, pcbb, ccbb, ocbb) &&
		InExt(pcdb, ccdb, ocdb, pcbb, ccbb, ocbb, extb) {
		qtot = computeContribution(pcbb, pcab, canon, i1, lb, len(preda))
	}

	// 2. Compute the probability q_{ab,da}^{k-->j} that Ca is reduced
	// to give Cd, and Cb is extended to give Ca.
	//
	// Turn to true if Cd is in the reduction set of Ca, and Cc is in
	// the extension set of Cb
	//
	// Compute the contribution to the probability
	//
	if isReductionSet(preda, creda, oreda, pcdb, ccdb, ocdb) &&
		InExt(pcab, ccab, ocab, pcbb, ccbb, ocbb, extb) {
		qtot += computeContribution(pcbb, pcab, canon, i1, lb, len(preda))
	}

	return qtot
}

// 1. Compute the probability q_{ab,0d}^{k-->j} that CaB is reduced to
// give zero island, and CbB is extended to give CdB.
//
func Qab0dBlock(pcab, ccab []int, ocab []float64,
	pcbb, ccbb []int, ocbb []float64, pcdb, ccdb []int, ocdb []float64,
	preda, creda [][]int, oreda [][]float64,
	canon [][]int, extb [][][]int, lb, i1 int) float64 {

	// Turn to true if Cc is in the reduction set of Ca, and Cd is in
	// the extension set of Cb
	//
	// Empty block
	//
	if isReductionSet(preda, creda, oreda, []int{0}, []int{0}, []float64{0}) &&
		InExt(pcdb, ccdb, ocdb, pcbb, ccbb, ocbb, extb) {
		return computeContribution(pcbb, pcab, canon, i1, lb, len(preda))
	}

	return 0.0
}
