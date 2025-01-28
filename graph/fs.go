package graph

import (
	"ws/graph/stack_queue"
)

func BFS(g *Graph, start int) []int {
	if len(g.adj) == 0 {
		return nil
	}

	queue := stack_queue.NewQueue()

	visited := make(map[int]bool) // Массив для отслеживания посещенных вершин
	visited[start] = true

	var result []int

	queue.Enqueue(start)

	// Пока очередь не пуста
	for {
		current, ok := queue.Dequeue() // Извлекаем вершину из очереди
		if !ok {
			break // Очередь пуста, завершаем обход
		}

		// Добавляем текущую вершину в результат
		result = append(result, current)

		// Перебираем всех соседей текущей вершины
		for _, neighbor := range g.adj[current] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue.Enqueue(neighbor)
			}
		}
	}

	return result
}

func DFS(g *Graph, start int) []int {
	if len(g.adj) == 0 {
		return nil
	}

	stack := stack_queue.NewStack()
	visited := make(map[int]bool)
	var result []int
	stack.Push(start)

	// Пока стек не пуст
	for {
		current, ok := stack.Pop() // Извлекаем вершину из стека
		if !ok {
			break // Стек пуст, завершаем обход
		}

		// Если вершина уже посещена, пропускаем ее
		if visited[current] {
			continue
		}

		visited[current] = true

		result = append(result, current)

		// Перебираем всех соседей текущей вершины в обратном порядке
		// (чтобы сохранить порядок обхода, как при рекурсивном DFS)
		neighbors := g.adj[current]
		for i := len(neighbors) - 1; i >= 0; i-- {
			neighbor := neighbors[i]
			if !visited[neighbor] {
				stack.Push(neighbor)
			}
		}
	}

	return result
}

// // Рекурсивная функция для обхода графа в глубину DFS
// func (g *Graph) dfsUtil(v int, visited map[int]bool, order *[]int) {
// 	visited[v] = true
// 	*order = append(*order, v)
// 	for _, u := range g.adj[v] {
// 		if !visited[u] {
// 			g.dfsUtil(u, visited, order)
// 		}
// 	}
// }

// func DFS(g *Graph, start int) []int {
// 	visited := make(map[int]bool)
// 	var order []int
// 	g.dfsUtil(start, visited, &order)
// 	return order
// }
