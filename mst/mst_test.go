package mst

import (
	"fmt"
	"testing"
	"ws/graph"
	"ws/output"
)

func TestMain(t *testing.T) {
	gr := graph.NewGraph()
	gr.AddEdge(1, 2, true, 12)
	gr.AddEdge(1, 3, true, 19)
	gr.AddEdge(1, 4, true, 6)
	gr.AddEdge(4, 5, true, 4)
	gr.AddEdge(4, 6, true, 9)
	gr.AddEdge(6, 7, true, 11)
	gr.AddEdge(2, 9, true, 2)
	gr.AddEdge(9, 10, true, 10)
	gr.AddEdge(2, 3, true, 1)

	output.PrintGraph(gr)
	//sort_edges := algorithms.Mergesort(gr.GetAllEdges())
	//output.PrintEdges(sort_edges)
	mst, weigth := MST(11, gr.GetAllEdges())
	fmt.Println(mst, weigth) // все рёбра кроме 1-3
}
