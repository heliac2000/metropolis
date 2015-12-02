//
// factorial.go
//

package util

import (
	"math/big"
)

func FactorialBigInt(n int) *big.Int {
	if n <= 1 {
		return big.NewInt(1)
	}

	bigN := big.NewInt(int64(n))
	return bigN.Mul(bigN, FactorialBigInt(n-1))
}

func Factorial(n int) float64 {
	if n < 0 {
		return float64(1)
	} else if n < factorial_n_max {
		return factorial_n[n]
	} else {
		v, _ := new(big.Rat).SetInt(FactorialBigInt(n)).Float64()
		return v
	}
}

// Cache
//
//var factorial_n_max int = int(math.Sqrt(float64(Dcoverage))) + 1
const factorial_n_max int = 51

// Immediate function executes only once
var factorial_n []float64 = func() []float64 {
	arr := make([]float64, factorial_n_max)
	for i := 0; i < factorial_n_max; i++ {
		v, _ := new(big.Rat).SetInt(FactorialBigInt(i)).Float64()
		arr[i] = v
	}
	return arr
}()
