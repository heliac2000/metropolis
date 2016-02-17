//
// energy_pair_reduce.go
//

package functions

import "math"

// Compute the interaction energy of a molecule at unit cell k1, in
// character c1, in orientation o1, and a molecule at unit cel k2,
// character c2, in orientation o2. Orientations in degrees. Uses the
// reduced representation based on PCA and four-way classification
// (see notes pg. 264 - 267)
//
func EnergyPairReduce(k1, k2, c1, c2 int, o1, o2 float64) float64 {
	cmTest, outside := CoulombIntMatrix(k1, k2, c1, c2, o1, o2)

	// outside = false means that the molecules are within the cut-off range
	if outside {
		return 0.0
	}

	cmVec := CoulombVectorise(cmTest)
	cmVecRed := CMvecReduce(cmVec, Npcs)

	// Zero SVM
	typ := "zero"
	opClass := SvmModelOp.Predict(cmVecRed)
	// Attractive SVM
	if opClass == "nonzero" {
		switch a := SvmModelNzp.Predict(cmVecRed); a {
		case "attractive":
			typ = a
		case "repulsive":
			r := SvmModelUsp.Predict(cmVecRed)
			if r == "repulsive" || r == "unstable" {
				typ = r
			}
		}
	}

	switch typ {
	case "unstable":
		return Eunstable
	case "attractive":
		return KernelRegsAtt.Predict(cmVecRed)
	case "repulsive":
		return math.Exp(KernelRegsRepLog.Predict(cmVecRed))
	}

	return 0.0
}
