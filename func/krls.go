//
// krls.go
//

package functions

import (
	"log"

	. "../util"
)

// R's KRLS object
//
type Krls struct {
	K               [][]float64
	Coeffs          [][]float64
	Looe            [][]float64
	Fitted          [][]float64
	X               [][]float64
	Y               [][]float64
	Sigma           float64
	Lambda          [][]float64
	R2              [][]float64
	VcovC           [][]float64 `json:"vcov.c"`
	VcovFitted      [][]float64 `json:"vcov.fitted"`
	BinaryIndicator [][]bool
}

// R's predict generic function
//
func (kls *Krls) Predict(newd [][]float64) float64 {
	// dimension check
	if len(kls.X[0]) != len(newd[0]) {
		log.Fatalln("PredictKrls: ncol(newdata) differs from ncol(krls.X).")
	}

	// scale
	xMeans, xSd := ColMeans(kls.X), ColSd(kls.X)
	x := Scale(kls.X, xMeans, xSd)

	// scale test data by means and sd of training data
	newData := Scale(newd, xMeans, xSd)

	// predict based on new kernel matrix
	// kernel distances for test points (simply recompute all pairwise
	// distances here because dist() is so fast)
	nn, l := len(newData), len(x)
	bind := make([][]float64, 0, nn+l)
	bind = append(append(bind, newData...), x...)

	// [R] gausskernel(rbind(newdata,X),sigma=object$sigma)[1:nn , (nn+1):(nn+nrow(X))]
	newDataK := GaussKernel(bind, kls.Sigma)
	m := make([]float64, 0, nn*l)
	for i := 0; i < nn; i++ {
		for j := nn; j < (nn + l); j++ {
			m = append(m, newDataK[i][j])
		}
	}

	// [R] matrix(gausskernel(...), nrow=nrow(newdata), byrow=FALSE)
	mm := Create2DimArrayFloat(nn, l)
	for i := 0; i < l; i++ {
		for j := 0; j < nn; j++ {
			mm[j][i] = m[i*nn+j]
		}
	}

	// predict fitted
	if len(mm[0]) != len(kls.Coeffs) {
		log.Fatalln("PredictKrls: ncol(newdataK) differs from ncol(krls.coeffs).")
	}

	// bring back to original scale
	return MatrixMultiplyFloat(mm, kls.Coeffs)[0][0]*ColSd(kls.Y)[0] + ColMeans(kls.Y)[0]
}
