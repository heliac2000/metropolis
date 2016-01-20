//
// cmvec_reduce.go
//

package functions

// Get the PCA representation of CMvec (output from Coulomb
// vectorise), and return first npcs elements
//
func CMvecReduce(cmvec []float64, npcs0 int) []float64 {
	cmVecPca := XeigPc.Predict(cmvec)
	return cmVecPca[0:npcs0]
}
