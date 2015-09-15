package util_test

import (
	"reflect"
	"testing"
)

import . "./"

type testJSON struct {
	X []int
	Y []float64
	Z string
}

func TestLoadDataFromJSONFile(t *testing.T) {
	file := "./data/test.json"
	expected := testJSON{
		X: []int{1, 2, 3},
		Y: []float64{1.5, 2.3, 3.1},
		Z: "Hello World!",
	}

	var actual testJSON
	LoadDataFromJSONFile(&actual, file)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("\ngot  %v\nwant %v", actual, expected)
		return
	}
}

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
