package util_test

import (
	"math"
	"testing"

	. "./"
)

func TestFactorial(t *testing.T) {
	testCases := map[int]float64{
		-1: float64(1),
		0:  float64(1),
		8:  float64(40320),
		16: float64(20922789888000),
		32: float64(263130836933693530167218012160000000),
		50: float64(30414093201713378043612608166064768844377641568960512000000000000),
	}

	var actual float64
	for n, expected := range testCases {
		actual = Factorial(n)
		if math.Abs(actual-expected)/expected > 1.0E-6 {
			t.Errorf("\ngot  %v\nwant %v", actual, expected)
			return
		}
	}
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
