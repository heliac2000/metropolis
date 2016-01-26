package util_test

import (
	"testing"

	. "./"
)

type testCasesApproXnCr struct {
	n, r     int
	expected float64
}

func TestApproXnCr(t *testing.T) {
	testCases := []testCasesApproXnCr{
		{
			// [R] format(nCr(2, 1), digits=22, trim=T)
			n: 2, r: 1, expected: 2.081040380091555785924,
		},
		{
			n: 3, r: 2, expected: 3.16450241940425947007,
		},
	}

	var actual float64
	for _, tc := range testCases {
		actual = ApproXnCr(tc.n, tc.r)
		if actual != tc.expected {
			t.Errorf("\ngot  %v\nwant %v", actual, tc.expected)
			return
		}
	}
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
