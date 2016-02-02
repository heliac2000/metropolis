//
// extension_reduction_block.go
//

package functions

import (
	"math"

	. "../util"
)

// Perform an extension-reduction step, with Canon test an appropriate
// canonical representation.
//
// xtest: Position component
// ctest: Character component
// otest: Orientation component
//
func ExtensionReductionBlock(xtest, ctest [][]int, otest [][]float64) (
	[][]int, [][]int, [][]float64, int, int) {

	// Probabilities to pick an island from Xtest to do the reduction.
	l := len(xtest)
	pred := make([]float64, l)
	qext := make([]float64, l)

	for k := 0; k < l; k++ {
		// Number of blocks in island Ctest[[k]]
		if ik := len(xtest[k]); ik == 1 && xtest[k][0] == 0 {
			pred[k], qext[k] = 0, 1
		} else {
			pred[k], qext[k] =
				math.Pow(float64(ik), Alpha1), 1+math.Pow(float64(ik), Alpha2)
		}
	}

	// Choose an island to reduce
	sRed := SamplingWithProbabilityFloat(Rnd, pred)
	rmn := make([]int, 0, l-1)
	for i := 0; i < l; i++ {
		if i != sRed {
			rmn = append(rmn, i)
		}
	}
	qext = append(qext[:sRed], qext[(sRed+1):]...)

	sExt := 0
	if len(rmn) == 1 {
		sExt = rmn[0]
	} else {
		sExt = rmn[SamplingWithProbabilityFloat(Rnd, qext)]
	}

	// Generate a list of extended islands.
	var xExtend [][][]int
	if len(xtest[sExt]) == 1 && xtest[sExt][0] == 0 {
		xExtendA := Copy2DimArrayInt(Inp.CharactersOrientations)
		for i := 0; i < len(xExtendA); i++ {
			xExtendA[i][0] = UCcenter
		}
		xExtend = make([][][]int, 1)
		xExtend[0] = xExtendA
	} else {
		xExtend, _ =
			ExtensionBlock(xtest[sExt], CoordsIsland(xtest[sExt], ctest[sExt], otest[sExt]))
	}

	var xReduce, cReduce [][]int
	var oReduce [][]float64
	if len(xtest[sRed]) == 1 {
		xReduce, cReduce, oReduce = [][]int{{0}}, [][]int{{0}}, [][]float64{{0}}
	} else {
		xReduce, cReduce, oReduce = ReductionBlock(xtest[sRed], ctest[sRed], otest[sRed])
	}

	xout, cout, oout := Copy2DimArrayInt(xtest), Copy2DimArrayInt(ctest), Copy2DimArrayFloat(otest)

	// Choose element from CExtend to replace Cout[[sExt]]
	if chE1, chE2 := 0, 0; len(xExtend) == 1 {
		chE2 = Rnd.Intn(len(xExtend[0]))
		xout[sExt] = []int{xExtend[0][chE2][0]}
		cout[sExt] = []int{xExtend[0][chE2][1]}
		oout[sExt] = []float64{float64(xExtend[0][chE2][2])}
	} else {
		chE1 = Rnd.Intn(len(xExtend))
		chE2 = Rnd.Intn(len(xExtend[chE1]))
		xout[sExt] = append(xout[sExt], xExtend[chE1][chE2][0])
		cout[sExt] = append(cout[sExt], xExtend[chE1][chE2][1])
		oout[sExt] = append(oout[sExt], float64(xExtend[chE1][chE2][2]))
	}

	chR1 := 0
	if len(xReduce) > 1 {
		chR1 = Rnd.Intn(len(xReduce))
	}

	xout[sRed] = xReduce[chR1]
	cout[sRed] = cReduce[chR1]
	oout[sRed] = oReduce[chR1]

	canonXout, canonCout, canonOout := CanonicalOrder(xout, cout, oout)

	return canonXout, canonCout, canonOout, sRed, sExt
}
