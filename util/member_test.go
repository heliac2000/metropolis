package util_test

import "testing"
import . "./"

func TestMember(t *testing.T) {
	m := []int{1, 2, 3, 4, 5}
	testCases := map[int]bool{
		-1: false,
		1:  true,
		3:  true,
		10: false,
	}

	var actual bool
	for n, expected := range testCases {
		actual = Member(n, m)
		if actual != expected {
			t.Errorf("\ngot  %v(n = %d)\nwant %v", actual, n, expected)
			return
		}
	}
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
