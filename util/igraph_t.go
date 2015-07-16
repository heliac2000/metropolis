//
// igraph_t.go
//

package util

import "log"

// Mode for GraphAdjacency
//
// NOTICE: Only ADJ_UNDIRECTED is valid.
//
const (
	ADJ_DIRECTED int = iota
	ADJ_UNDIRECTED
	ADJ_MAX
	ADJ_MIN
	ADJ_UPPER
	ADJ_LOWER
	ADJ_PLUS
)

// IGraph: The internal data structure for storing graphs.
//
// igraph/include/igraph_datatype.h:
// typedef struct igraph_s {...} igraph_t;
//
type IGraph struct {
	N                        int
	Directed                 bool
	From, To, Oi, Ii, Os, Is []int
}

// Factory: Creates an empty graph with some vertices and no edges.
//
// igraph/src/type_indexededgelist.c:
// int igraph_empty_attrs(
//       igraph_t *graph, igraph_integer_t n,
//       igraph_bool_t directed, void* attr
//     )
//
func newIGraph(directed bool, cap int) *IGraph {
	//ig := new(IGraph)
	return &IGraph{
		N:        0,
		Directed: directed,
		From:     make([]int, 0, cap),
		To:       make([]int, 0, cap),
		Oi:       make([]int, 0, cap),
		Ii:       make([]int, 0, cap),
		Os:       make([]int, 1, cap),
		Is:       make([]int, 1, cap),
	}
}

// Method: Adds vertices to a graph
//
// igraph/src/type_indexededgelist.c:
// int igraphAddVertices(
//       igraph_t *graph, igraph_integer_t nv, void *attr
//     )
//
func (ig *IGraph) igraphAddVertices(nv int) {
	if nv < 0 {
		log.Fatalln("igraphAddVertices: cannot add negative number of vertices")
	}

	add := make([]int, nv+1)
	ec := len(ig.From)
	for i := 1; i <= nv; i++ {
		add[i] = ec
	}
	ig.Os = append(ig.Os, add...)
	ig.Is = append(ig.Is, add...)
	ig.N += nv

	return
}

// Method: Adds edges to a graph object
//
// igraph/src/type_indexededgelist.c:
// int igraphAddEdges(
//       igraph_t *graph, const igraph_vector_t *edges,
//       void *attr
//     )
//
func (ig *IGraph) igraphAddEdges(edges []int) {
	if len(edges)%2 != 0 {
		log.Fatalln("igraphAddEdges: invalid (odd) length of edges vector")
	}
	for i := 0; i < len(edges); i++ {
		if edges[i] < 0 || edges[i] >= ig.N {
			log.Fatalln("igraphAddEdges: cannot add edges")
		}
	}

	for i := 0; i < len(edges); {
		if ig.Directed || edges[i] > edges[i+1] {
			ig.From = append(ig.From, edges[i])
			i++
			ig.To = append(ig.To, edges[i])
			i++
		} else {
			ig.To = append(ig.To, edges[i])
			i++
			ig.From = append(ig.From, edges[i])
			i++
		}
	}

	/* oi & ii */
	ig.Oi = igraphVectorOrder(ig.From, ig.To, ig.N)
	ig.Ii = igraphVectorOrder(ig.To, ig.From, ig.N)

	/* os & is, its length does not change, error safe */
	ig.Os = igraphICreateStart(ig.From, ig.Oi, ig.N)
	ig.Is = igraphICreateStart(ig.To, ig.Ii, ig.N)

	return
}
