package util_test

import "testing"

import . "./"

type testCasesAverage struct {
	xs       []float64
	expected float64
}

func TestAverage(t *testing.T) {
	testCases := []testCasesAverage{
		{
			xs:       []float64{},
			expected: 0.0,
		},
		{
			xs:       []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0},
			expected: 5.5,
		},
	}

	var actual float64
	for _, tc := range testCases {
		actual = Average(tc.xs)
		if actual != tc.expected {
			t.Errorf("\ngot  %v\nwant %v", actual, tc.expected)
			return
		}
	}
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
