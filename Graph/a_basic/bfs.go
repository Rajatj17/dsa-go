package basic

func (g *Graph) BFS(start int) []int {
	if start < 0 || start >= g.Vertices {
		return nil
	}

	visited := make([]bool, g.Vertices)
	queue := []int{start}
	result := []int{}
	distance := []int{}
	parent := []int{}
	visited[start] = true

	for i := 0; i < g.Vertices; i++ {
		parent[i] = -1
	}

	parent[start] = start
	distance[start] = 0

	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]

		result = append(result, vertex)

		for _, edge := range g.AdjList[vertex] {
			if !visited[edge.To] {
				queue = append(queue, edge.To)
				visited[edge.To] = true

				// Parent
				parent[edge.To] = vertex

				// Distance
				distance[edge.To] = distance[vertex] + 1
			}
		}
	}

	return result
}
