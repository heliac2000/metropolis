//
// approx_calc.go
//

package util

import "math"

// Approximate calculation of nCr
//
func ApproXnCr(n, r int) float64 {
	nf, rf := float64(n), float64(r)
	return (math.Pow((nf/(nf-rf)), 0.5) *
		math.Pow((nf/(nf-rf)), nf) * math.Exp(-rf) * math.Pow((nf-rf), rf)) /
		Factorial(r)
}
