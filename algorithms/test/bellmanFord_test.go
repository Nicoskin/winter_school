package algorithms

import (
	"fmt"
	"testing"
	"ws/algorithms"
	"ws/graph"
)

func TestMain1(t *testing.T) {
	gr := graph.NewGraph()
	gr.AddEdge(0, 1, true, 2)
	gr.AddEdge(0, 5, true, 6)
	gr.AddEdge(0, 2, true, 4)
	gr.AddEdge(0, 3, true, 1)
	gr.AddEdge(2, 3, true, 2)
	gr.AddEdge(1, 5, true, 2)
	gr.AddEdge(4, 5, true, 1)
	gr.AddEdge(3, 4, true, 3)

	dist, parents, cycle := algorithms.BellmanFord(gr, 3)
	fmt.Println("Dist:   ", dist)
	fmt.Println("Parents:", parents)
	fmt.Println("Cycle:", cycle)
}
