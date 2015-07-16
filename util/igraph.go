//
// igraph.go
//

package util

import "log"

// Creates a graph object from an adjacency matrix
//
func GraphAdjacency(adj [][]int, mode int) *IGraph {
	return igraphCreate(graphAdjacencyDense(adj, mode), len(adj), false)
}

// Calculates the (weakly or strongly) connected components in a graph.
//
// igraph/src/components.c:
// int igraph_clusters(
//       const igraph_t *graph, igraph_vector_t *membership,
//       igraph_vector_t *csize, igraph_integer_t *no,
//       igraph_connectedness_t mode
//     )
//
// rigraph/R/zzz-deprecate.R:
// #' @export clusters
// deprecated("clusters", components)
//
// rigraph/R/structural.properties.R:
// components <- function(graph, mode=c("weak", "strong"))
//
func GraphClusters(graph *IGraph) ([]int, []int, int) {
	no_of_nodes := graph.N
	var qlen int
	if no_of_nodes > 100000 {
		qlen = 10000
	} else {
		qlen = no_of_nodes / 10
	}
	q := make([]int, 0, qlen)

	act_cluster_size, no_of_clusters := 0, 1
	already_added := make([]int, no_of_nodes)
	membership := make([]int, no_of_nodes)
	csize := make([]int, 0, no_of_nodes)
	var neis []int
	var act_node int

	for first_node := 0; first_node < no_of_nodes; first_node++ {
		if already_added[first_node] == 1 {
			continue
		}
		already_added[first_node] = 1
		act_cluster_size = 1
		membership[first_node] = no_of_clusters - 1
		q = append(q, first_node)

		for len(q) > 0 {
			act_node = q[0] // pop
			if len(q) > 1 {
				q = q[1:]
			} else {
				q = nil
			}
			neis = igraphNeighbors(graph, act_node)
			for i := 0; i < len(neis); i++ {
				neighbor := neis[i]
				if already_added[neighbor] == 1 {
					continue
				}
				q = append(q, neighbor)
				already_added[neighbor] = 1
				act_cluster_size++
				membership[neighbor] = no_of_clusters - 1
			}
		}
		no_of_clusters++
		csize = append(csize, act_cluster_size)
	}

	for i := 0; i < len(membership); i++ {
		membership[i]++
	}

	return membership, csize, no_of_clusters - 1
}

// Creates a graph object from an adjacency matrix
//
// igraph/src/structure_generators.c:
// int igraph_i_adjacency_max(igraph_matrix_t *adjmatrix, igraph_vector_t *edges)
//
func graphAdjacencyDense(adj [][]int, mode int) []int {
	if mode != ADJ_UNDIRECTED || len(adj) != len(adj[0]) {
		// Other modes are not implemented
		return nil
	}

	no_of_nodes := len(adj)
	edges := make([]int, 0, no_of_nodes)

	for i := 0; i < no_of_nodes; i++ {
		for j := i; j < no_of_nodes; j++ {
			for k := 0; k < Max(adj[i][j], adj[j][i]); k++ {
				edges = append(edges, i, j)
			}
		}
	}

	return edges
}

// Creates a graph with the specified edges
//
// igraph/src/structure_generators.c:
// int igraphCreate(
//        igraph_t *graph, const igraph_vector_t *edges,
//        igraph_integer_t n, igraph_bool_t directed
//     )
//
func igraphCreate(edges []int, no_of_nodes int, directed bool) *IGraph {
	// Check edges
	if no_of_nodes < 0 {
		log.Fatalln("cannot create empty graph with negative number of vertices.")
	} else if (len(edges) % 2) != 0 {
		log.Fatalln("igraphCreate: Invalid (odd) edges vector")
	}

	for _, v := range edges {
		if v < 0 {
			log.Fatalln("igraphCreate: Invalid (negative) vertex id")
		}
	}

	ig := newIGraph(directed, no_of_nodes)
	ig.igraphAddVertices(no_of_nodes)
	if len(edges) > 0 {
		if ig.N < Max(edges...)+1 {
			ig.igraphAddVertices(no_of_nodes)
		}
		ig.igraphAddEdges(edges)
	}

	return ig
}

// Calculate the order of the elements in a vector
//
// igraph/src/vector.c:
// int igraphVectorOrder(
//       const igraph_vector_t* v,
//       const igraph_vector_t *v2,
//       igraph_vector_t* res, igraph_real_t nodes
//     )
//
func igraphVectorOrder(v, v2 []int, nodes int) []int {
	if v == nil {
		return nil
	}

	edges := len(v)
	ptr, rad, res := make([]int, nodes+1), make([]int, edges), make([]int, 0, edges)

	var radix int
	for i := 0; i < edges; i++ {
		if radix = v2[i]; ptr[radix] != 0 {
			rad[i] = ptr[radix]
		}
		ptr[radix] = i + 1
	}

	var next int
	for i := 0; i < nodes+1; i++ {
		if ptr[i] != 0 {
			next = ptr[i] - 1
			res = append(res, next)
			for rad[next] != 0 {
				next = rad[next] - 1
				res = append(res, next)
			}
		}
	}

	ptr, rad = make([]int, nodes+1), make([]int, edges)
	var edge int
	for i := 0; i < edges; i++ {
		edge = res[edges-i-1]
		radix = v[edge]
		if ptr[radix] != 0 {
			rad[edge] = ptr[radix]
		}
		ptr[radix] = edge + 1
	}

	res = make([]int, 0, edges)
	for i := 0; i < nodes+1; i++ {
		if ptr[i] != 0 {
			next = ptr[i] - 1
			res = append(res, next)
			for rad[next] != 0 {
				next = rad[next] - 1
				res = append(res, next)
			}
		}
	}

	return res
}

// Internal
//
// igraph/src/type_indexededgelist.c
// int igraphICreateStart(
//       igraph_vector_t *res, igraph_vector_t *el, igraph_vector_t *iindex,
//       igraph_integer_t nodes
//     )
//
func igraphICreateStart(el, iindex []int, nodes int) []int {
	no_of_edges, idx := len(el), -1
	res := make([]int, nodes+1)

	/* create the index */
	if no_of_edges > 0 {
		for i := 0; i <= el[iindex[0]]; i++ {
			idx++
			res[idx] = 0
		}
		for i := 1; i < no_of_edges; i++ {
			n := el[iindex[i]] - el[iindex[res[idx]]]
			for j := 0; j < n; j++ {
				idx++
				res[idx] = i
			}
		}
		for i, j := 0, el[iindex[res[idx]]]; i < nodes-j; i++ {
			idx++
			res[idx] = no_of_edges
		}
	}

	return res
}

// Adjacent vertices to a vertex
//
// igraph/src/type_indexededgelist.c:
// int igraphNeighbors(
//       const igraph_t *graph, igraph_vector_t *neis, igraph_integer_t pnode,
//       igraph_neimode_t mode
//     )
//
func igraphNeighbors(graph *IGraph, pnode int) []int {
	if pnode < 0 || pnode > graph.N-1 {
		log.Fatalln("igraphNeighbors: cannot get neighbors")
	}

	// Calculate needed space first & allocate it
	length := (graph.Os[pnode+1] - graph.Os[pnode]) +
		(graph.Is[pnode+1] - graph.Is[pnode])

	neis, idx := make([]int, length), 0
	for i := graph.Os[pnode]; i < graph.Os[pnode+1]; i++ {
		neis[idx] = graph.To[graph.Oi[i]]
		idx++
	}
	for i := graph.Is[pnode]; i < graph.Is[pnode+1]; i++ {
		neis[idx] = graph.From[graph.Ii[i]]
		idx++
	}

	return neis
}
