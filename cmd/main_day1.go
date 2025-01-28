package main

import (
	"fmt"
	"ws/graph"
	"ws/output"
)

func main() {
	g := graph.NewGraph()
	g.AddEdge(1, 2, true)
	g.AddEdge(1, 3, true)
	g.AddEdge(1, 4, true)
	g.AddEdge(2, 4, true)
	output.PrintGraph(g)

	// Дружба
	fmt.Println("Дружба 1-3:", graph.HasEdge(g, 1, 3))
	fmt.Println("Дружба 2-3:", graph.HasEdge(g, 2, 3))
	fmt.Println()

	//BFS
	bfs := graph.BFS(g, 1)
	fmt.Println("BFS:", bfs) // Вывод: [1 2 3 4]

	//DFS
	dfs := graph.DFS(g, 1)
	fmt.Println("DFS:", dfs) // Вывод: [1 2 3 4]

	// Новый граф
	g = graph.NewGraph()
	g.AddEdge(0, 1, true)
	g.AddEdge(1, 2, true)
	g.AddEdge(2, 3, true)
	g.AddEdge(3, 0, true)

	g.AddEdge(4, 5, true)
	g.AddEdge(5, 6, true)
	g.AddEdge(6, 4, true)
	output.PrintGraph(g)

	fmt.Println(graph.ConnectedComponents(g))
	fmt.Println()

}
