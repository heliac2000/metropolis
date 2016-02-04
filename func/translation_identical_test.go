package functions_test

import (
	"testing"

	. "./"
)

type testCasesTranslationIdentical struct {
	pos1, chr1 []int
	ori1       []float64
	pos2, chr2 []int
	ori2       []float64
	same       bool
}

func TestTranslationIdentical(t *testing.T) {
	SetInitData("./data")

	// NOTICE: R is 1-base index, golang is 0-base.
	testCases := []testCasesTranslationIdentical{
		{
			// [R]
			// TranslationIdentical(list(c(5, 40, 200), c(1,1,1), c(1,1,1)), list(c(5,40,200), c(1,1,1), c(1,1,1)))
			//
			pos1: []int{4, 39, 199}, chr1: []int{0, 0, 0}, ori1: []float64{1, 1, 1},
			pos2: []int{4, 39, 199}, chr2: []int{0, 0, 0}, ori2: []float64{1, 1, 1},
			same: true,
		},
		{
			pos1: []int{4, 39, 199}, chr1: []int{0, 0, 0}, ori1: []float64{1, 1, 1},
			pos2: []int{4, 39, 199}, chr2: []int{1, 0, 0}, ori2: []float64{1, 1, 1},
			same: false,
		},
	}

	for _, tc := range testCases {
		same := TranslationIdentical(tc.pos1, tc.chr1, tc.ori1, tc.pos2, tc.chr2, tc.ori2)
		if same != tc.same {
			t.Errorf("\ngot  %v\nwant %v", same, tc.same)
			return
		}
	}

	return
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
