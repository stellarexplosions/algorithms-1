package bfs

import (
	"github.com/arnauddri/algorithms/data-structures/graph"
	"testing"
)

func TestBfsShortestPath(t *testing.T) {
	h := graph.NewGraph()

	for i := 0; i < 10; i++ {
		v := graph.VertexId(i)
		h.AddVertex(v)
	}

	for i := 0; i < 9; i++ {
		h.AddEdge(graph.VertexId(i), graph.VertexId(i+1))
	}

	distance := bfsShortestPath(h, graph.VertexId(0))

	for i := 0; i < len(distance); i++ {
		if distance[graph.VertexId(i)] != i {
			t.Error()
		}

		if bfsDist(h, graph.VertexId(0), graph.VertexId(i)) != i {
			t.Error()
		}
	}
}
