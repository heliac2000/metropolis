package util_test

import (
	"reflect"
	"testing"
)
import . "./"

func TestMatrixMultiply(t *testing.T) {
	m := [][]int{{1, 4, 3}, {2, 1, 4}, {3, 2, 1}}
	actual := MatrixMultiply(m, m)
	expected := [][]int{{18, 14, 22}, {16, 17, 14}, {10, 16, 18}}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("\ngot  %v\nwant %v", actual, expected)
		return
	}
}

func TestMatrixMultiply1(t *testing.T) {
	m := [][]int{{1, 4, 3}, {2, 1, 4}, {3, 2, 1}}
	actual := MatrixMultiply1(m, m)
	expected := [][]int{{18, 14, 22}, {16, 17, 14}, {10, 16, 18}}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("\ngot  %v\nwant %v", actual, expected)
		return
	}
}

func TestIntersection(t *testing.T) {
	u, v := []int{4, 3, 2, 1}, []int{6, 5, 3, 4}
	actual := Intersection(u, v)
	expected := []int{4, 3}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("\ngot  %v\nwant %v", actual, expected)
		return
	}
}

func TestTranspose(t *testing.T) {
	m := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	actual := Transpose(m)
	expected := [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("\ngot  %v\nwant %v", actual, expected)
		return
	}
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
