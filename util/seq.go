//
// seq.go
//

package util

func Seq(from, to, by float64) []float64 {
	seq := make([]float64, 0, int((to-from)/by)+1)
	for i := from; i <= to; i += by {
		seq = append(seq, i)
	}
	return seq
}

func SeqInt(from, to, by int) []int {
	seq := make([]int, 0, ((to-from)/by)+1)
	for i := from; i <= to; i += by {
		seq = append(seq, i)
	}
	return seq
}
