package algorithms

import (
	"math"
	"ws/graph"
)

func Dijkstra(g *graph.Graph, start int) ([]int, []int) {
	dist := make([]int, len(g.GetAdj()))
	parent := make([]int, len(g.GetAdj()))
	pq := NewPriorityQueue()

	for i := range dist {
		dist[i] = math.MaxInt32 // Используем MaxInt32 как "бесконечность"
	}
	dist[start] = 0

	// Добавляем начальную вершину в очередь
	item := Item{vertex: start, dist: 0}
	pq.Push(item)

	visited := make([]bool, len(g.GetAdj()))

	for len(pq.items) > 0 {
		minItem := pq.Pop()
		curVertex := minItem.vertex

		if visited[curVertex] {
			continue // Пропускаем уже посещенные вершины
		}
		visited[curVertex] = true

		// Получаем всех соседей текущей вершины
		neighbors := g.GetNeighbors(curVertex)
		//fmt.Println("neighbors", curVertex, ":", neighbors)
		for _, neighbor := range neighbors {
			v, w := neighbor.V, neighbor.W
			alt := dist[curVertex] + w // Правильная формула для альтернативной длины пути

			if alt < dist[v] { // Если найден более короткий путь
				parent[v] = curVertex
				dist[v] = alt
				item := Item{vertex: v, dist: alt}
				pq.Push(item)
			}
		}
	}

	//fmt.Println(dist, parent)

	return dist, parent
}
