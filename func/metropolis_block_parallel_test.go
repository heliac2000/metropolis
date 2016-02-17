//
// metropolis_block_parallel_test.go
//

package functions_test

import (
	"testing"

	. "../util"
	. "./"
)

type testCasesMetoropolisBlockParallel struct {
	N        int
	expected int
}

func TestMetropolisBlockParallel(t *testing.T) {
	SetInitData("./data")
	TempS = Seq(200, 300, 10)
	Nparallel = len(TempS)

	testCases := []testCasesMetoropolisBlockParallel{
		{
			N:        1000,
			expected: 0,
		},
	}

	for _, tc := range testCases {
		MetropolisBlockParallel(tc.N, "Eout.dat", "Cout.csv")
		t.Logf("\nN = %d\n", tc.N)
		// t.Errorf("\ngot  %v\nwant %v", actual2, tc.expected)
	}
}

// Local Variables:
// compile-command: (concat "go test -v -gcflags='-B' " (file-name-nondirectory buffer-file-name))
// End:
//
// go test -v -gcflags='-B' -timeout 1h -cpuprofile cpu.dat metropolis_block_parallel_test.go
