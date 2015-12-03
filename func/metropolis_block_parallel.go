//
// metropolis_block_parallel.go
//

package functions

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	. "../util"
)

// Do Metropolis-Hastings sampling with 1-step shifts (slow, uses
// exact PF; uses parallel tempering)
//
func MetropolisBlockParallel(N int, eout, cout string) {
	// random seed
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Truncate files
	fn := getFuncName()
	truncateFile(cout, fn)
	truncateFile(eout, fn)

	tick := 1000
	if N < tick {
		tick = N
	}

	coutP, eoutP := make([][]Canonical, Nparallel), make([][]float64, Nparallel)
	for k := 0; k < Nparallel; k++ {
		coutP[k] = make([]Canonical, tick)
		eoutP[k] = make([]float64, tick)
		coutP[k][0] = NewCanonical(CanonicalOrder(CanonicalGen()))
		eoutP[k][0] = EnergyCanonical(coutP[k][0].Explode())
	}

	for j := 0; j < N/tick; j++ {
		for k := 1; k < tick; k++ {
			if k%100 == 0 {
				log.Printf("n = %4d/N = %5d\n", j*tick+k, N)
			}

			// Draw a uniform random variable to decide to do the tempering
			if randFloat64(rnd) < 0.5 {
				// Do an ordinary extension of a uniformly chosen element
				schoose := rnd.Intn(Nparallel)
				cout := make([]Canonical, len(coutP[schoose]))
				for i := 0; i < k; i++ {
					cout[i] = coutP[schoose][i].Dup()
				}
				eout := CopyVector(eoutP[schoose]).([]float64)
				for h := 0; h < Nparallel; h++ {
					if h != schoose {
						coutP[h][k] = coutP[h][k-1].Dup()
						eoutP[h][k] = eoutP[h][k-1]
					}
				}

				pos, chr, ori, _, _ := ExtensionReductionBlock(cout[k-1].Explode())
				coutTemp, _ := CanonicalImplode(CanonicalOrder(pos, chr, ori))
				qC1C2 := ExtensionReductionProbabilityReaction(cout[k-1].pos, cout[k-1].chr, cout[k-1].ori, coutTemp.pos, coutTemp.chr, coutTemp.ori)
				qC2C1 := ExtensionReductionProbabilityReaction(coutTemp.pos, coutTemp.chr, coutTemp.ori, cout[k-1].pos, cout[k-1].chr, cout[k-1].ori)

				ax, ay := eout[k-1], EnergyCanonical(coutTemp.Explode())

				lctemp := 0
				for r := 0; r < len(coutTemp.pos); r++ {
					if len(coutTemp.pos[r]) > 1 || coutTemp.pos[r][0] != 0 {
						lctemp++
					}
				}

				lcprev := 0
				for r := 0; r < len(cout[k-1].pos); r++ {
					if len(cout[k-1].pos[r]) > 1 || cout[k-1].pos[r][0] != 0 {
						lcprev++
					}
				}

				fx := Factorial(lcprev) * Factorial(Nuc) / (Factorial(lcprev) * Factorial(Nuc-lcprev))
				fy := Factorial(lctemp) * Factorial(Nuc) / (Factorial(lctemp) * Factorial(Nuc-lctemp))
				// Stationary distribution in invisible islands approximation.
				pix := IslandSymmetryCanonicalB(cout[k-1].pos) * fx * fx
				piy := IslandSymmetryCanonicalB(coutTemp.pos) * fy * fy

				alpha := piy * qC2C1 * math.Exp(-(ay-ax)/(KB*Temp)) / (pix * qC1C2)
				if alpha > 1.0 {
					alpha = 1.0
				}

				if randFloat64(rnd) < alpha {
					coutP[schoose][k] = coutTemp
					eoutP[schoose][k] = ay
				} else {
					// Incase of rejection
					coutP[schoose][k] = cout[k-1]
					eoutP[schoose][k] = ax
				}
			} else {
				// Do the swap update
				s1, s2 := rnd.Intn(Nparallel), 0
				temp1 := TempS[s1]
				switch {
				case s1 == 0:
					s2 = 1
				case s1 == (Nparallel - 1):
					s2 = Nparallel - 2
				default:
					rmn := []int{s1 - 1, s1 + 1}
					s2 = rmn[rnd.Intn(len(rmn))]
				}

				temp2 := TempS[s2]
				// Get the new states
				c1, c2 := coutP[s1][k-1], coutP[s2][k-1]
				// Get the lengths
				l1, l2 := 0, 0
				for h := 0; h < len(c1.pos); h++ {
					if len(c1.pos[h]) > 1 || c1.pos[h][0] != 0 {
						l1++
					}
				}
				for h := 0; h < len(c2.pos); h++ {
					if len(c2.pos[h]) > 1 || c2.pos[h][0] != 0 {
						l2++
					}
				}

				// Energy of state 1 with temperature 1
				// Energy of state 2 with temperature 2
				e11, e22 := eoutP[s1][k-1], eoutP[s2][k-1]

				// Energy of state 2 with temperature 1
				// Energy of state 1 with temperature 2
				e12 := EnergyCanonical(c1.Explode())
				e21 := EnergyCanonical(c2.Explode())

				//r12 := h1x2 * h2x1 / (h1x1 * h2x2)
				r12 := math.Exp(-(e21-e11)/(KB*temp1) - (e12-e22)/(KB*temp2))
				alpha := 1.0
				if r12 < alpha {
					alpha = r12
				}

				for h := 0; h < Nparallel; h++ {
					coutP[h][k] = coutP[h][k-1].Dup()
					eoutP[h][k] = eoutP[h][k-1]
				}

				if randFloat64(rnd) < alpha {
					coutP[s1][k], coutP[s2][k] = c2.Dup(), c1.Dup()
					eoutP[s1][k], eoutP[s2][k] = e21, e12
				}
			}
		}

		// Save data
		for i := 0; i < Nparallel; i++ {
			appendArrayToFile(eoutP[i], eout)
			//appendListWithCSV(coutP[i], cout)
		}

		// Reset
		for i := 0; i < Nparallel; i++ {
			coutP[i][0], eoutP[i][0] = coutP[i][tick-1].Dup(), eoutP[i][tick-1]
			for j := 1; j < tick; j++ {
				coutP[i][j], eoutP[i][j] = Canonical{}, 0
			}
		}
	}
}

// Golang's rand.Float64() method returns a pseudo-random number in [0.0, 1.0).
// randFloat64 function returns a pseudo-random number in (0.0, 1.0), same as R's
// runif() function.
//
func randFloat64(r *rand.Rand) float64 {
	for {
		if f := float64(r.Int63()) / (1 << 63); f < 1.0 {
			return f
		}
	}
}

func getFuncName() string {
	pc := make([]uintptr, 2)
	runtime.Callers(2, pc)
	fname := runtime.FuncForPC(pc[0]).Name()
	strs := strings.Split(fname, ".")
	if l := len(strs); l > 0 {
		fname = strs[l-1]
	}

	return fname
}

func truncateFile(f, fn string) {
	if FileExists(f) {
		if fi, err := os.Stat(f); err == nil && fi.Mode().IsRegular() {
			if err := os.Truncate(f, 0); err != nil {
				log.Fatalf("[%s] %s", fn, err.Error())
			}
		}
	}
}

// Save data
//
func appendArrayToFile(arr []float64, eout string) {
	file, err := os.OpenFile(eout, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		log.Fatalf("[%s] %s", getFuncName(), err.Error())
	}

	defer file.Close()
	for _, v := range arr {
		fmt.Fprintf(file, "%.22f\n", v)
	}
}

func appendListWithCSV(lst [][][][]int, cout string) {
	file, err := os.OpenFile(cout, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		log.Fatalf("[%s] %s", getFuncName(), err.Error())
	}

	defer file.Close()
	writer := csv.NewWriter(file)
	for i := 0; i < len(lst); i++ {
		for j := 0; j < len(lst[i]); j++ {
			r, c := len(lst[i][j]), len(lst[i][j][0])
			l := make([]string, 0, r*c+2)
			l = append(l, strconv.Itoa(r), strconv.Itoa(c))
			for k := 0; k < c; k++ {
				for n := 0; n < r; n++ {
					l = append(l, strconv.Itoa(lst[i][j][n][k]))
				}
			}
			writer.Write(l)
		}
		writer.Write([]string{""})
		writer.Flush()
	}
}
