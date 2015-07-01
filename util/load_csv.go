package util

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

//
// Return float64 slice
//

// Load data from a CSV file
//
func LoadFromCsvFile(fname string, sep rune) []float64 {
	csvfile, err := os.Open(fname)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer csvfile.Close()

	reader := csv.NewReader(csvfile)
	reader.Comma = sep
	reader.FieldsPerRecord = -1
	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ret := make([]float64, 0, len(rawCSVdata))
	for _, r := range rawCSVdata {
		for _, f := range r {
			n, err := strconv.ParseFloat(f, 64)
			if err == nil {
				ret = append(ret, n)
			}
		}
	}

	return ret
}

// Load data from a CSV file, return 2-dim array
//
func LoadFromCsvFile2Dim(fname string, sep rune) [][]float64 {
	csvfile, err := os.Open(fname)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer csvfile.Close()

	reader := csv.NewReader(csvfile)
	reader.Comma = sep
	reader.FieldsPerRecord = -1

	ret := make([][]float64, 0, 1000)
	var rr []float64
	for r, err := reader.Read(); err == nil; r, err = reader.Read() {
		rr = make([]float64, 0, len(r))
		for _, f := range r {
			n, err := strconv.ParseFloat(f, 64)
			if err == nil {
				rr = append(rr, n)
			}
		}
		ret = append(ret, rr)
	}

	return ret
}

// Load data from a CSV file
//
func LoadFromCsvFileInt(fname string) []int {
	csvfile, err := os.Open(fname)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer csvfile.Close()

	reader := csv.NewReader(csvfile)
	reader.FieldsPerRecord = -1
	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ret := make([]int, 0, len(rawCSVdata))
	for _, r := range rawCSVdata {
		for _, f := range r {
			n, err := strconv.Atoi(f)
			if err == nil {
				ret = append(ret, n)
			}
		}
	}

	return ret
}
