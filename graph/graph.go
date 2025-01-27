package graph

type Graph struct {
	adj map[int][]int // Список смежности

	// При необходимости хранить веса:
	// edges []Edge
}

func NewGraph() *Graph {
	return &Graph{adj: make(map[int][]int)}
}

func (g *Graph) AddEdge(u, v int, undirected bool) {
	g.adj[u] = append(g.adj[u], v)
	if undirected {
		g.adj[v] = append(g.adj[v], u)
	}
}

func HasEdge(g *Graph, u, v int) bool {
	for _, vertex := range g.adj[u] {
		if vertex == v {
			return true
		}
	}
	return false
}

func (g *Graph) GetAdj() map[int][]int {
	return g.adj
}

func ConnectedComponents(g *Graph) (count int, comp map[int]int) {
	// Инициализация
	visited := make(map[int]bool)
	comp = make(map[int]int)
	count = 0

	// Для каждой вершины в графе
	for v := range g.adj {
		if !visited[v] {
			// Начинаем новую компоненту
			count++
			// Запускаем DFS от текущей вершины
			order := BFS(g, v)

			// Помечаем все вершины из order как принадлежащие текущей компоненте
			for _, u := range order {
				comp[u] = count
			}
		}
	}

	return count, comp
}
