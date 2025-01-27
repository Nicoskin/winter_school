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

	g.adj[u] = append(g.adj[u], v)
	if undirected {
		g.adj[v] = append(g.adj[v], u)
	}
	g.edges = append(g.edges, Edge{u, v, w})
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
	var edges []Edge
	for _, edge := range g.edges {
		if edge.u > edge.v {
			edge.u, edge.v = edge.v, edge.u
		}
		edges = append(edges, edge)
	}
	return edges
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
