package graph

import (
	"fmt"
	"testing"
	"ws/graph"
)

func TestMain(t *testing.T) {
	gr := graph.NewGraph()
	gr.AddEdge(1, 2, true)
	gr.AddEdge(1, 3, true)
	gr.AddEdge(1, 4, true)
	gr.AddEdge(2, 4, true)
	fmt.Println(gr.GetAdj())
	fmt.Println(graph.HasEdge(gr, 1, 3))
	fmt.Println(graph.HasEdge(gr, 2, 3))
}

func TestMain1(t *testing.T) {
	g := graph.NewGraph()
	g.AddEdge(0, 1, true)
	g.AddEdge(1, 2, true)
	g.AddEdge(2, 3, true)
	g.AddEdge(3, 0, true)

	g.AddEdge(4, 5, true)
	g.AddEdge(5, 6, true)
	g.AddEdge(6, 4, true)
	fmt.Println(g.GetAdj())

	fmt.Println(graph.ConnectedComponents(g))
}
