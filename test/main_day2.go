package main

import (
	"fmt"
	"ws/algorithms"
	"ws/graph"
	"ws/mst"
	"ws/output"
	"ws/unionfind"
)

func main() {
	ds := unionfind.NewDisjointSet(6)
	ds.Union(0, 1)
	ds.Union(1, 2)
	ds.Union(3, 4)
	fmt.Println("\n3 компоненты")
	output.PrintDisjointSet(ds)
	fmt.Println("\nПосле Union(2,3)")
	ds.Union(2, 3)
	output.PrintDisjointSet(ds)
	fmt.Println("Parent:", ds.GetParent())
	fmt.Println("Rank:  ", ds.GetRank())

	// MST
	gr := graph.NewGraph()
	gr.AddEdge(1, 2, true, 12)
	gr.AddEdge(1, 3, true, 19)
	gr.AddEdge(1, 4, true, 6)
	gr.AddEdge(4, 5, true, 4)
	gr.AddEdge(5, 1, true, 2)
	gr.AddEdge(5, 2, true, 7)

	fmt.Println("\nГраф для MST")
	output.PrintGraph(gr)
	mst, weigth := mst.MST(6, gr.GetAllEdges())
	fmt.Println("\nMST")
	output.PrintEdges(mst)
	fmt.Println("weigth:", weigth)

	// Di BF
	gr = graph.NewGraph()
	gr.AddEdge(0, 1, true, 2)
	gr.AddEdge(0, 5, true, 6)
	gr.AddEdge(0, 2, true, 4)
	gr.AddEdge(0, 3, true, 1)
	gr.AddEdge(2, 3, true, 2)
	gr.AddEdge(1, 5, true, 2)
	gr.AddEdge(4, 5, true, 1)
	gr.AddEdge(3, 4, true, 3)

	fmt.Println("\nBellman-Ford от 3 вершины")
	dist, parents, cycle := algorithms.BellmanFord(gr, 3)
	fmt.Println("Dist:   ", dist)
	fmt.Println("Parents:", parents)
	fmt.Println("Cycle:", cycle)

	fmt.Println("\nDijkstra от 3 вершины")
	dist, parents = algorithms.Dijkstra(gr, 3)
	fmt.Println("Dist:   ", dist)
	fmt.Println("Parents:", parents)
}
