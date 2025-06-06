package basic

type Edge struct {
	To     int
	Weight int
}

type Graph struct {
	Vertices int
	AdjList  [][]Edge
	Directed bool
}

func NewGraph(vertices int, directed bool) *Graph {
	return &Graph{
		Vertices: vertices,
		AdjList:  make([][]Edge, vertices),
		Directed: directed,
	}
}

func (g *Graph) AddEdge(from, to, weight int) {
	g.AdjList[from] = append(g.AdjList[from], Edge{To: to, Weight: weight})

	// Add reverse Edge
	if !g.Directed {
		g.AdjList[to] = append(g.AdjList[to], Edge{To: from, Weight: weight})
	}
}

func removeEdgeFromSlice(edges []Edge, target int) []Edge {
	for k, v := range edges {
		if v.To == target {
			return append(edges[:k], edges[k+1:]...)
		}
	}

	return edges
}

func (g *Graph) RemoveEdge(from, to int) {
	g.AdjList[from] = removeEdgeFromSlice(g.AdjList[from], to)

	if !g.Directed {
		g.AdjList[to] = removeEdgeFromSlice(g.AdjList[to], from)
	}
}
