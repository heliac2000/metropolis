//
// qabcd_block.go
//

package functions

import "math"

// Compute the probability q_{ab,cd}^{k-->j}. Inputs are islands
// (positions, characters, and orientations). Extb is the extension of
// b, Lb is the number of elements)
//
func QabcdBlock(pcab, ccab []int, ocab []float64,
	pcbb, ccbb []int, ocbb []float64, pccb, cccb []int, occb []float64,
	pcdb, ccdb []int, ocdb []float64, preda, creda [][]int, oreda [][]float64,
	canon [][]int, extb [][][]int, lb, i1 int) float64 {

	// 1. Compute the probability q_{ab,cd}^{k-->j} that CaB is reduced
	// to give CcB, and CbB is extended to give CdB.

	// Turn to true if Cc is in the reduction set of Ca, and Cd is in
	// the extension set of Cb
	in_reda := false

	// Check if Cc is in the reduction set of Ca
	for k := 0; k < len(preda); k++ {
		if IsomorphIslandsBlock(preda[k], creda[k], oreda[k], pccb, cccb, occb) {
			in_reda = true
			break
		}
	}

	// Check if Cd is in the extension set of Cb
	//
	// Compute the contribution to the probability
	//
	if !in_reda || !InExt(pcdb, ccdb, ocdb, pcbb, ccbb, ocbb, extb) {
		return 0
	}

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
	if len(pcbb) != 1 || pcbb[0] != 0 {
		qba += math.Pow(float64(len(pcbb)), Alpha2)
	}
	qba /= denQba

	qtot := pa * qba /
		((float64(len(preda)) / IslandSymmetryBlock(pcab)) *
			(float64(lb) / IslandSymmetryBlock(pcbb)))

	// 2. Compute the probability q_{ab,dc}^{k-->j} that Ca is reduced
	// to give Cd, and Cb is extended to give Cc.
	in_reda = false

	// Turn to true if Cd is in the reduction set of Ca, and Cc is in the extension set of Cb
	// Check if Cc is in the reduction set of Ca
	for k := 0; k < len(preda); k++ {
		if IsomorphIslandsBlock(preda[k], creda[k], oreda[k], pcdb, ccdb, ocdb) {
			in_reda = true
			break
		}
	}

	// Check if Cc is in the extension set of Cb
	//
	// Compute the contribution to the probability
	//
	if !in_reda || !InExt(pccb, cccb, occb, pcbb, ccbb, ocbb, extb) {
		return qtot
	}

	// Compute the denominators of Pa and Qba
	denPa, denQba = 0.0, 0.0
	// Deal with the zero parts afterwards. Make sure the canonical
	// representations are in the correct order with only one zero
	for k := 0; k < len(canon)-1; k++ {
		denPa += math.Pow(float64(len(canon[k])), Alpha1)
		denQba += 1.0 + math.Pow(float64(len(canon[k])), Alpha2)
	}
	// Include the zero part
	denQba -= math.Pow(float64(len(canon[i1])), Alpha2)
	pa = math.Pow(float64(len(canon[i1])), Alpha1) / denPa

	// Incase a zero is chosen
	qba = 1.0
	if len(pcbb) != 1 || pcbb[0] != 0 {
		qba += math.Pow(float64(len(pcbb)), Alpha2)
	}
	qba /= denQba

	return qtot + pa*qba/
		((float64(len(preda))/IslandSymmetryBlock(pcab))*
			(float64(lb)/IslandSymmetryBlock(pcbb)))
}
