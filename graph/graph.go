package graph

type Graph struct {
	adj   map[int][]int // Список смежности
	edges []Edge
}
type Edge struct{ u, v, w int }

func NewGraph() *Graph {
	return &Graph{adj: make(map[int][]int), edges: []Edge{}}
}

func (g *Graph) AddEdge(u, v int, undirected bool, weight ...int) {
	w := 0 // Вес по умолчанию
	if len(weight) > 0 {
		w = weight[0]
	}

	// Добавляем ребро u->v
	g.edges = append(g.edges, Edge{u, v, w})
	g.adj[u] = append(g.adj[u], v)

	// Если граф ненаправленный, добавляем также ребро v->u
	if undirected {
		g.edges = append(g.edges, Edge{v, u, w})
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
	visited := make(map[int]bool)
	comp = make(map[int]int)
	count = 0

	for v := range g.adj {
		if !visited[v] {
			count++
			order := BFS(g, v)

			for _, u := range order {
				comp[u] = count
			}
		}
	}

	return count, comp
}

func (g *Graph) GetAllEdges() []Edge {
	return g.edges // Возвращаем все рёбра без изменения
}

func (e *Edge) GetWeight() int {
	return e.w
}

func (e *Edge) GetUV() (int, int) {
	return e.u, e.v
}

func (g *Graph) SetAdj(adj map[int][]int) {
	g.adj = adj
}

func (g *Graph) GetNeighbors(u int) []struct{ V, W int } {
	neighbors := []struct{ V, W int }{}

	// Проверка существования вершины и получение всех соседей
	if edges, ok := g.adj[u]; ok {
		for _, v := range edges {
			// Находим вес ребра между u и v
			for _, edge := range g.edges {
				if edge.u == u && edge.v == v {
					neighbors = append(neighbors, struct{ V, W int }{V: edge.v, W: edge.w})
					break
				}
			}
		}
	}

	return neighbors
}
