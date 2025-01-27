package output

import (
	"fmt"
	"sort"
	"ws/graph"
	"ws/unionfind"
)

func PrintGraph(g *graph.Graph) {
	adj := g.GetAdj()

	// Сохраняем ключи в срез
	keys := make([]int, 0, len(adj))
	for key := range adj {
		keys = append(keys, key)
	}

	// Сортируем ключи
	sort.Ints(keys)

	fmt.Println("===================")
	for _, key := range keys {
		values := adj[key]
		fmt.Printf("%2d: [", key)
		for i, value := range values {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Print(value)
		}
		fmt.Println("]")
	}
	fmt.Println("===================")
}

func PrintDisjointSet(ds *unionfind.DisjointSet) {
	fmt.Println("===================")
	parents := ds.GetParent()
	last_parent := -1

	// Print indices
	fmt.Print("         ")
	for i := 0; i < len(parents); i++ {
		if i > 0 && last_parent != ds.Find(i) {
			fmt.Print(" ")
		}
		fmt.Printf("%d", i)
		last_parent = ds.Find(i)
	}
	fmt.Println()

	// Print values
	last_parent = -1
	fmt.Printf("DS: %d  ->", len(parents))
	for i := 0; i < len(parents); i++ {
		if i > 0 && last_parent != ds.Find(i) {
			fmt.Print(" ")
		}
		fmt.Printf("%d", ds.Find(i))
		last_parent = ds.Find(i)
	}
	fmt.Println("\n===================")
}

func PrintEdges(edges []graph.Edge) {
	fmt.Println("===================")
	for _, edge := range edges {
		u, v := edge.GetUV()
		fmt.Printf("%-2d- %-2d: %d\n", u, v, edge.GetWeight())
	}
	fmt.Println("===================")
}
