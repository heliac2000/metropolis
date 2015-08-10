//
// qabcd_block.go
//

package functions

import "math"

// Compute the probability q_{ab,cd}^{k-->j}. Inputs are islands
// (positions, characters, and orientations). Extb is the extension of
// b, Lb is the number of elements)
//
func QabcdBlock(pcab []int, pcbb, ccbb []int, ocbb []float64,
	pccb, cccb []int, occb []float64, pcdb, ccdb []int, ocdb []float64,
	preda, creda [][]int, oreda [][]float64,
	canon [][]int, extb [][][]int, lb, i1 int) float64 {

	// 1. Compute the probability q_{ab,cd}^{k-->j} that CaB is reduced
	// to give CcB, and CbB is extended to give CdB.
	//
	// Turn to true if Cc is in the reduction set of Ca, and Cd is in
	// the extension set of Cb
	//
	// Compute the contribution to the probability
	//
	qtot := 0.0
	if isReductionSet(preda, creda, oreda, pccb, cccb, occb) &&
		InExt(pcdb, ccdb, ocdb, pcbb, ccbb, ocbb, extb) {
		qtot = computeContribution(pcbb, pcab, canon, i1, lb, len(preda))
	}

	// 2. Compute the probability q_{ab,dc}^{k-->j} that Ca is reduced
	// to give Cd, and Cb is extended to give Cc.
	//
	// Turn to true if Cd is in the reduction set of Ca, and Cc is in
	// the extension set of Cb
	//
	// Compute the contribution to the probability
	//
	if isReductionSet(preda, creda, oreda, pcdb, ccdb, ocdb) &&
		InExt(pccb, cccb, occb, pcbb, ccbb, ocbb, extb) {
		qtot += computeContribution(pccb, pcab, canon, i1, lb, len(preda))
	}

	return qtot
}

// Return true if blk is in the reduction set of Ca
//
func isReductionSet(preda, creda [][]int, oreda [][]float64,
	pblk, cblk []int, oblk []float64) bool {
	for k := 0; k < len(preda); k++ {
		if IsomorphIslandsBlock(preda[k], creda[k], oreda[k], pblk, cblk, oblk) {
			return true
		}
	}
	return false
}

// Compute the contribution to the probability
//
func computeContribution(blk1, blk2 []int, canon [][]int, i1, lb, lreda int) float64 {
	// Compute the denominators of Pa and Qba
	denPa, denQba := 0.0, 0.0
	// Deal with the zero parts afterwards. Make sure the canonical
	// representations are in the correct order with only one zero
	for k := 0; k < len(canon)-1; k++ {
		denPa += math.Pow(float64(len(canon[k])), Alpha1)
		denQba += 1.0 + math.Pow(float64(len(canon[k])), Alpha2)
	}
	// Include the zero part
	denQba -= math.Pow(float64(len(canon[i1])), Alpha2)
	pa := math.Pow(float64(len(canon[i1])), Alpha1) / denPa

	// Incase a zero is chosen
	qba := 1.0
	if len(blk1) != 1 || blk1[0] != 0 {
		qba += math.Pow(float64(len(blk1)), Alpha2)
	}
	qba /= denQba

	return pa * qba /
		((float64(lreda) / IslandSymmetryBlock(blk2)) *
			(float64(lb) / IslandSymmetryBlock(blk1)))
}
