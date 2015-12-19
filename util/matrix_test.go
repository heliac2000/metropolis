package util_test

import (
	"reflect"
	"testing"
)
import . "./"

func TestTranspose(t *testing.T) {
	m := [][]float64{{1.0, 2.5}, {4.1, 5.8}, {7.7, 8.9}}
	actual := Transpose(m)
	expected := [][]float64{{1.0, 4.1, 7.7}, {2.5, 5.8, 8.9}}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("\ngot  %v\nwant %v", actual, expected)
		return
	}
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
