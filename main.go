package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"time"

	. "./func"
)

func main() {
	// Parse command line options
	N, temp, eout, cout, err := parseOption()
	if err != nil {
		os.Exit(2)
	}
	Temp = temp

	// Load initial data
	SetInitData()

	// Start
	log.Printf("N = %d\n", N)
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
func parseOption() (n int, temp float64, eout, cout string, err *error) {
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
	tempPtr := flag.Float64("T", Temp, "Temperature in kelvin.")
	eoutPtr := flag.String("Eout", "Eout.dat", "Eout file.")
	coutPtr := flag.String("Cout", "Cout.csv", "Cout file.")
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

	n, temp, eout, cout = *nPtr, *tempPtr, *eoutPtr, *coutPtr

	if n < 2 {
		fmt.Fprintf(os.Stderr,
			"\nStep number should be greater than 1.\n")
		flag.Usage()
		err = new(error)
	}

	return
}
