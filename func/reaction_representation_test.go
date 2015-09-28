package functions_test

import (
	"reflect"
	"testing"

	. "./"
)

type testCasesReactionRepresentation struct {
	pcanon, ccanon             [][]int
	ocanon                     [][]float64
	pcanon_out, ccanon_out     [][]int
	ocanon_out                 [][]float64
	pil, cil                   [][]int
	oil                        [][]float64
	coeffsRct, coeffsPdt, diff []int
}

func TestReactionRepresentation(t *testing.T) {
	SetInitData(
		"./data/PrecursorUnitCell.csv", "./data/UnitCell2.csv", "./data/PrecursorUnitCellAxes.csv",
		"./data/kernelregS_Rep_log.json", "./data/kernelregS_Att.json", "./data/svm_model.json")

	// NOTICE: R is 1-base index, golang is 0-base.
	//
	testCases := []testCasesReactionRepresentation{
		{
			// [R]
			// canon = Canonical.Order(Canonical.Gen())
			// canon_out = Canonical.Order(ExtensionReductionBlock(canon)[[1]])
			// =>
			// canon = list(
			//           list(313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 0),
			//           list(4, 4, 6, 6, 5, 5, 5, 6, 4, 5, 0),
			//           list(90, 0, 60, 120, 30, 120, 60, 0, 90, 30, 0))
			//
			// canon_out = list(
			//           list(313, 313, 313, 313, 313, 313, 313, 313, 313, 313, 0),
			//           list(4, 4, 6, 6, 5, 5, 5, 6, 4, 5, 0),
			//           list(90, 0, 60, 120, 30, 120, 60, 0, 90, 30, 0))
			//
			pcanon:     [][]int{{313}, {313}, {313}, {313}, {313}, {313}, {313}, {313}, {313}, {313}, {0}},
			ccanon:     [][]int{{3}, {3}, {5}, {5}, {4}, {4}, {4}, {5}, {3}, {4}, {0}},
			ocanon:     [][]float64{{90}, {0}, {60}, {120}, {30}, {120}, {60}, {0}, {90}, {30}, {0}},
			pcanon_out: [][]int{{313}, {313}, {313}, {313}, {313}, {313}, {313}, {313}, {313}, {313}, {0}},
			ccanon_out: [][]int{{3}, {3}, {5}, {5}, {4}, {4}, {4}, {5}, {3}, {4}, {0}},
			ocanon_out: [][]float64{{90}, {0}, {60}, {120}, {30}, {120}, {60}, {0}, {90}, {30}, {0}},
			pil:        [][]int{{313}, {313}, {313}, {313}, {313}, {313}, {313}, {313}, {0}},
			cil:        [][]int{{3}, {3}, {5}, {5}, {4}, {4}, {4}, {5}, {0}},
			oil:        [][]float64{{90}, {0}, {60}, {120}, {30}, {120}, {60}, {0}, {0}},
			coeffsRct:  []int{2, 1, 1, 1, 2, 1, 1, 1, 1},
			coeffsPdt:  []int{2, 1, 1, 1, 2, 1, 1, 1, 1},
			diff:       []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	for _, tc := range testCases {
		pil, cil, oil, coeffsRct, coeffsPdt, diff := ReactionRepresentation(
			tc.pcanon, tc.ccanon, tc.ocanon, tc.pcanon_out, tc.ccanon_out, tc.ocanon_out)
		if !reflect.DeepEqual(pil, tc.pil) {
			t.Errorf("\n[PIL] got  %v\nwant %v", pil, tc.pil)
			return
		} else if !reflect.DeepEqual(cil, tc.cil) {
			t.Errorf("\n[CIL] got  %v\nwant %v", cil, tc.cil)
			return
		} else if !reflect.DeepEqual(oil, tc.oil) {
			t.Errorf("\n[OIL] got  %v\nwant %v", oil, tc.oil)
			return
		} else if !reflect.DeepEqual(coeffsRct, tc.coeffsRct) {
			t.Errorf("\n[CoeffsRct] got  %v\nwant %v", coeffsRct, tc.coeffsRct)
			return
		} else if !reflect.DeepEqual(coeffsPdt, tc.coeffsPdt) {
			t.Errorf("\n[CoeffsPdt] got  %v\nwant %v", coeffsPdt, tc.coeffsPdt)
			return
		} else if !reflect.DeepEqual(diff, tc.diff) {
			t.Errorf("\n[Diff] got  %v\nwant %v", diff, tc.diff)
			return
		}
	}

	return
}

// Local Variables:
// compile-command: (concat "go test -gcflags='-B' -v " (file-name-nondirectory buffer-file-name))
// End:
