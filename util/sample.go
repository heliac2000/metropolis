package util

import "math/rand"

// R's sample() function
//
// Weighted random generation in Python
// http://eli.thegreenplace.net/2010/01/22/weighted-random-generation-in-python
//
func SamplingWithProbability(r *rand.Rand, prob []int) int {
	total, totals := 0, make([]int, 0, len(prob))
	for _, v := range prob {
		total += v
		totals = append(totals, total)
	}

	rnd := r.Intn(total)
	for i, v := range totals {
		if rnd < v {
			return i
		}
	}

	return 0
}

func SamplingWithProbabilityFloat(r *rand.Rand, prob []float64) int {
	total, totals := 0.0, make([]float64, 0, len(prob))
	for _, v := range prob {
		total += v
		totals = append(totals, total)
	}

	rnd := r.Float64() * total
	for i, v := range totals {
		if rnd < v {
			return i
		}
	}

	return 0
}
