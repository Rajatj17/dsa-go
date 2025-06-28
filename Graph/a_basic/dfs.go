package basic

func (g *Graph) dfsHelper(vertex int, visited []bool, result *[]int) {
	visited[vertex] = true
	*result = append(*result, vertex)

	for _, edge := range g.AdjList[vertex] {
		if !visited[edge.To] {
			g.dfsHelper(edge.To, visited, result)
		}
	}
}

func (g *Graph) DFS(start int) []int {
	if start < 0 || start >= g.Vertices {
		return nil
	}

	visited := make([]bool, g.Vertices)
	results := []int{}

	g.dfsHelper(start, visited, &results)

	return results
}
