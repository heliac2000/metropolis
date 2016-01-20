//
// cmvec_reduce.go
//

package functions

// Get the PCA representation of CMvec (output from Coulomb
// vectorise), and return first npcs elements
//
func CMvecReduce(cmVec []float64, npcs0 int) []float64 {
	cmVecPca := XeigPc.Predict([][]float64{cmVec})
	return cmVecPca[0][0:npcs0]
}
