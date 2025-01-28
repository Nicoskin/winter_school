package algorithms

import (
	"math"
	"ws/graph"
)

func BellmanFord(g *graph.Graph, start int) ([]int, []int, bool) {
	dist := make([]int, len(g.GetAdj()))
	parent := make([]int, len(g.GetAdj()))
	negativeCycle := false

	for i := range dist {
		dist[i] = math.MaxInt32 // Используем MaxInt32 как "бесконечность"
	}
	dist[start] = 0

	// Основной цикл алгоритма
	for i := 1; i < len(g.GetAdj()); i++ {
		for _, edge := range g.GetAllEdges() {
			u, v := edge.GetUV()
			w := edge.GetWeight()
			if dist[u] != math.MaxInt32 && dist[u]+w < dist[v] {
				dist[v] = dist[u] + w
				parent[v] = u
			}
		}
	}

	// Проверка на наличие отрицательного цикла
	for _, edge := range g.GetAllEdges() {
		u, v := edge.GetUV()
		w := edge.GetWeight()
		if dist[u] != math.MaxInt32 && dist[u]+w < dist[v] {
			negativeCycle = true
			break
		}
	}

	return dist, parent, negativeCycle
}
