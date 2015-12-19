//
// metropolis_block_parallel_test.go
//

package functions_test

import (
	"testing"

	. "./"
)

type testCasesMetoropolisBlockParallel struct {
	N        int
	expected int
}

func TestMetropolisBlockParallel(t *testing.T) {
	SetInitData(
		"./data/PrecursorUnitCell.csv", "./data/UnitCell2.csv", "./data/PrecursorUnitCellAxes.csv",
		"./data/kernelregS_Rep_log.json", "./data/kernelregS_Att.json", "./data/svm_model.json")

	testCases := []testCasesMetoropolisBlockParallel{
		{
			N:        100,
			expected: 0,
		},
	}

	for _, tc := range testCases {
		MetropolisBlockParallel(tc.N, "/dev/null", "/dev/null")
		t.Logf("\nN = %d\n", tc.N)
		// t.Errorf("\ngot  %v\nwant %v", actual2, tc.expected)
	}
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
//
// go test -v -gcflags='-B' -timeout 1h -cpuprofile cpu.dat metropolis_block_parallel_test.go
