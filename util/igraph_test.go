package util_test

import (
	"reflect"
	"testing"

	. "../func"
	. "./"
)

type testCasesGraphAdjacency struct {
	hop        [][]int
	directed   int
	ig         *IGraph
	membership []int
	csize      []int
	no         int
}

func TestGraphAdjacency(t *testing.T) {
	_, ahop, ahop2 := PrepareHopData("../data")
	testCases := []testCasesGraphAdjacency{
		{ // Directed graph is not implemented
			hop:      [][]int{{1, 1}, {1, 1}},
			directed: ADJ_DIRECTED,
			ig: &IGraph{
				2, false, []int{}, []int{}, []int{}, []int{}, []int{0, 0, 0, 0}, []int{0, 0, 0, 0},
			},
		},
		{
			// [R] unclass(graph.adjacency(array(1:5, c(3,3)), mode="undirected"))
			hop:      [][]int{{1, 4, 2}, {2, 5, 3}, {3, 1, 4}},
			directed: ADJ_UNDIRECTED,
			ig: &IGraph{
				3, false,
				[]int{0, 1, 1, 1, 1, 2, 2, 2, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2},
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2},
				[]int{0, 4, 3, 2, 1, 12, 11, 10, 9, 8, 7, 6, 5, 15, 14, 13, 19, 18, 17, 16},
				[]int{0, 4, 3, 2, 1, 7, 6, 5, 12, 11, 10, 9, 8, 15, 14, 13, 19, 18, 17, 16},
				[]int{0, 1, 10, 20},
				[]int{0, 8, 16, 20},
			},
		},
		{
			hop: [][]int{
				{0, 1, 0, 0, 0, 0, 0, 0},
				{1, 0, 1, 1, 0, 0, 0, 0},
				{0, 1, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 1, 0, 1, 1},
				{0, 0, 0, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0, 1, 0, 0}},
			directed: ADJ_UNDIRECTED,
			ig: &IGraph{
				8, false,
				[]int{1, 2, 3, 5, 6, 7},
				[]int{0, 1, 1, 4, 5, 5},
				[]int{0, 1, 2, 3, 4, 5},
				[]int{0, 1, 2, 3, 4, 5},
				[]int{0, 0, 1, 2, 3, 3, 4, 5, 6},
				[]int{0, 1, 3, 3, 3, 4, 6, 6, 6},
			},
		},
		{
			hop:      ahop,
			directed: ADJ_UNDIRECTED,
			ig: &IGraph{
				N: len(ahop), Directed: false,
				From: LoadFromCsvFile("../data/ahop_from.csv"),
				To:   LoadFromCsvFile("../data/ahop_to.csv"),
				Oi:   LoadFromCsvFile("../data/ahop_oi.csv"),
				Ii:   LoadFromCsvFile("../data/ahop_ii.csv"),
				Os:   LoadFromCsvFile("../data/ahop_os.csv"),
				Is:   LoadFromCsvFile("../data/ahop_is.csv"),
			},
		},
		{
			hop:      ahop2,
			directed: ADJ_UNDIRECTED,
			ig: &IGraph{
				N: len(ahop2), Directed: false,
				From: LoadFromCsvFile("../data/ahop2_from.csv"),
				To:   LoadFromCsvFile("../data/ahop2_to.csv"),
				Oi:   LoadFromCsvFile("../data/ahop2_oi.csv"),
				Ii:   LoadFromCsvFile("../data/ahop2_ii.csv"),
				Os:   LoadFromCsvFile("../data/ahop2_os.csv"),
				Is:   LoadFromCsvFile("../data/ahop2_is.csv"),
			},
		},
	}

	var actual *IGraph
	for _, tc := range testCases {
		actual = GraphAdjacency(tc.hop, tc.directed)
		if !reflect.DeepEqual(actual, tc.ig) {
			t.Errorf("\ngot  %v\nwant %v", actual, tc.ig)
			return
		}
	}
}

func TestGraphClusters(t *testing.T) {
	_, ahop, ahop2 := PrepareHopData("../data")
	testCases := []testCasesGraphAdjacency{
		{
			// [R] unclass(graph.adjacency(array(1:5, c(3,3)), mode="undirected"))
			hop:        [][]int{{1, 4, 2}, {2, 5, 3}, {3, 1, 4}},
			directed:   ADJ_UNDIRECTED,
			membership: []int{1, 1, 1},
			csize:      []int{3},
			no:         1,
		},
		{
			hop: [][]int{
				{0, 1, 0, 0, 0, 0, 0, 0},
				{1, 0, 1, 1, 0, 0, 0, 0},
				{0, 1, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 1, 0, 1, 1},
				{0, 0, 0, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0, 1, 0, 0}},
			directed:   ADJ_UNDIRECTED,
			membership: []int{1, 1, 1, 1, 2, 2, 2, 2},
			csize:      []int{4, 4},
			no:         2,
		},
		{
			hop:        ahop,
			directed:   ADJ_UNDIRECTED,
			membership: LoadFromCsvFile("../data/ahop_cl_ms.csv"),
			csize:      []int{900},
			no:         1,
		},
		{
			hop:        ahop2,
			directed:   ADJ_UNDIRECTED,
			membership: LoadFromCsvFile("../data/ahop2_cl_ms.csv"),
			csize:      []int{450, 450},
			no:         2,
		},
	}

	for _, tc := range testCases {
		membership, csize, no := GraphClusters(GraphAdjacency(tc.hop, tc.directed))
		if !(reflect.DeepEqual(tc.membership, membership) &&
			reflect.DeepEqual(tc.csize, csize) && tc.no == no) {
			t.Errorf("\nmembership: %v\ncsize: %v\nno %v\n", membership, csize, no)
			return
		}
	}
}

// Benchmark
//
func BenchmarkGraphAdjacency(b *testing.B) {
	_, ahop, _ := PrepareHopData("../data")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GraphAdjacency(ahop, ADJ_UNDIRECTED)
	}
}

/*
[R] Write a CSV file to a file

xxx = unclass(graph.adjacency(Ahop, mode="undirected"))
write(xxx[[3]], file="/home/takahashi/work/AIMR/Development/Golang/data/ahop_from.csv", ncol=length(xxx[[3]]),sep=",")
write(xxx[[4]], file="/home/takahashi/work/AIMR/Development/Golang/data/ahop_to.csv", ncol=length(xxx[[4]]),sep=",")
write(xxx[[5]], file="/home/takahashi/work/AIMR/Development/Golang/data/ahop_oi.csv", ncol=length(xxx[[5]]),sep=",")
write(xxx[[6]], file="/home/takahashi/work/AIMR/Development/Golang/data/ahop_ii.csv", ncol=length(xxx[[6]]),sep=",")
write(xxx[[7]], file="/home/takahashi/work/AIMR/Development/Golang/data/ahop_os.csv", ncol=length(xxx[[7]]),sep=",")
write(xxx[[8]], file="/home/takahashi/work/AIMR/Development/Golang/data/ahop_is.csv", ncol=length(xxx[[8]]),sep=",")

For clusters:

write(clusters(graph.adjacency(Ahop, mode="undirected"))$membership, file="/home/takahashi/work/AIMR/Development/Golang/data/ahop_cl_ms.csv", ncol=1000, sep=",")
write(clusters(graph.adjacency(Ahop2, mode="undirected"))$membership, file="/home/takahashi/work/AIMR/Development/Golang/data/ahop2_cl_ms.csv", ncol=1000, sep=",")

*/

// Local Variables:
// compile-command: (concat "go test -v " (file-name-nondirectory buffer-file-name))
// End:
