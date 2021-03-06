package functions_test

import (
	"testing"

	. "./"
)

type testExtensionReductionBlock struct {
	xtest, ctest [][]int
	otest        [][]float64
}

func TestExtensionRedunctionBlock(t *testing.T) {
	SetInitData("./data")

	testCases := []testExtensionReductionBlock{
		{
			// NOTICE: R is 1-base index, golang is 0-base.
			//   ctest(characters) are 1-base in R but 0-base in Golang
			//
			// [R]
			// lst1 = Canonical.Gen(); ExtensionRedunctionBlock(lst1)
			// lst1 = list(
			//   list(c(313), c(313), c(313), c(313), c(313), c(313), c(313), c(313), c(313), c(313)),
			//   list(c(5), c(6), c(4), c(4), c(5), c(7), c(4), c(6), c(5), c(6)),
			//   list(c(60), c(0), c(0), c(0), c(120), c(0), c(120), c(120), c(120), c(0)))
			//
			xtest: [][]int{{313}, {313}, {313}, {313}, {313}, {313}, {313}, {313}, {313}, {313}},
			ctest: [][]int{{4}, {5}, {3}, {3}, {4}, {6}, {3}, {5}, {4}, {5}},
			otest: [][]float64{{60}, {0}, {0}, {0}, {120}, {0}, {120}, {120}, {120}, {0}},
		},
	}

	// sample() 関数でランダム値を取得して計算を行っているため、同じ入力値
	// でも出力結果が異なる
	//
	for _, tc := range testCases {
		xout, cout, oout, sRed, sExt := ExtensionReductionBlock(tc.xtest, tc.ctest, tc.otest)
		t.Logf("\nXout = %v\nCout = %v\nOout = %v\n", xout, cout, oout)
		t.Logf("sRed = %v, sExt = %v\n", sRed, sExt)
	}
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
