package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	. "./func"
	. "./util"
)

func main() {
	// Parse command line options
	N, eout, cout, dataDir, err := parseOption()
	if err != nil {
		os.Exit(2)
	}

	// Load initial data
	SetInitData(dataDir)

	// Start
	log.Printf("N = %d\n", N)
	log.Printf("TempS = %#v\n", TempS)
	log.Printf("Nparallel = %d\n", Nparallel)
	log.Printf("Start.")

	start := time.Now()
	MetropolisBlockParallel(N, eout, cout)
	end := time.Now()

	// End
	log.Printf("End.")
	elapsed := float64(end.Sub(start).Nanoseconds()) / 1000000000
	log.Printf("Execution time = %.2f s/%.2f m/%.2f h\n", elapsed, elapsed/60, elapsed/3600)
}

// Parse option switchs
//
func parseOption() (n int, eout, cout, dataDir string, err *error) {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `
Usage of %s:

  %s [options] -N Step_Number

`, path.Base(os.Args[0]), os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\n")
	}

	// Parse
	nPtr := flag.Int("N", 0, "Step number.")
	tempPtr := flag.String("Temp", "100,500,35", "For parallel tempering.")
	eoutPtr := flag.String("Eout", "Eout.dat", "Eout file.")
	coutPtr := flag.String("Cout", "Cout.csv", "Cout file.")
	dataDirPtr := flag.String("DataDir", "./data", "Input data directory.")
	flag.Parse()

	// `-N': required option
	t := false
	for _, arg := range os.Args[1:] {
		if arg == "-N" {
			t = true
			break
		}
	}
	if !t {
		fmt.Fprintf(os.Stderr,
			"\nStep number should be specified with `-N' option.\n")
		flag.Usage()
		err = new(error)
		return
	}

	n, eout, cout, dataDir = *nPtr, *eoutPtr, *coutPtr, *dataDirPtr

	if n < 2 {
		fmt.Fprintf(os.Stderr,
			"\nStep number should be greater than 1.\n")
		flag.Usage()
		err = new(error)
	}

	parseTemp(*tempPtr)

	return
}

func parseTemp(temp string) {
	temps := make([]float64, 3)
	for i, v := range strings.Split(temp, ",") {
		temps[i], _ = strconv.ParseFloat(v, 64)
	}

	TempS = Seq(temps[0], temps[1], temps[2])
	Nparallel = len(TempS)
}
