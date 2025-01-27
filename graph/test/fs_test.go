package graph

import (
	"fmt"
	"testing"
	"ws/graph"
)

// go test -v graph.go fs_test.go fs.go
func TestMain0(t *testing.T) {
	gr := graph.NewGraph()
	gr.AddEdge(1, 2, true)
	gr.AddEdge(1, 3, true)
	gr.AddEdge(1, 4, true)
	gr.AddEdge(4, 5, true)
	gr.AddEdge(4, 6, true)
	gr.AddEdge(6, 7, true)
	gr.AddEdge(2, 9, true)
	gr.AddEdge(9, 10, true)

	fmt.Println(gr.GetAdj())

	bfs := graph.BFS(gr, 1)
	fmt.Println(bfs)

	dfs := graph.DFS(gr, 1)
	fmt.Println(dfs)
}
