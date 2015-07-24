package util_test

import (
	"reflect"
	"testing"
)
import . "./"

func TestTranspose(t *testing.T) {
	// float
	m_float64 := [][]float64{{1.0, 2.5}, {4.1, 5.8}, {7.7, 8.9}}
	actual_float64 := Transpose(m_float64).([][]float64)
	expected_float64 := [][]float64{{1.0, 4.1, 7.7}, {2.5, 5.8, 8.9}}

	if !reflect.DeepEqual(actual_float64, expected_float64) {
		t.Errorf("\ngot  %v\nwant %v", actual_float64, expected_float64)
		return
	}

	// int
	m_int := [][]int{{1, 2}, {4, 5}, {7, 8}}
	actual_int := Transpose(m_int).([][]int)
	expected_int := [][]int{{1, 4, 7}, {2, 5, 8}}

	if !reflect.DeepEqual(actual_int, expected_int) {
		t.Errorf("\ngot  %v\nwant %v", actual_int, expected_int)
		return
	}
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
