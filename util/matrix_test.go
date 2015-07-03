package util_test

import (
	"reflect"
	"testing"
)
import . "./"

func TestTranspose(t *testing.T) {
	m := [][]float64{{1, 2}, {4, 5}, {7, 8}}
	actual := Transpose(m)
	expected := [][]float64{{1, 4, 7}, {2, 5, 8}}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("\ngot  %v\nwant %v", actual, expected)
		return
	}
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
