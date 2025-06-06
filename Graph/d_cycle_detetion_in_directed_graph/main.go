package cycledetetionindirectedgraph

import (
	basic "DSA/Graph/a_basic"
)

// Detect Cycle in a Directed Graph
// Given the number of vertices V and a list of directed edges, determine whether the graph contains a cycle or not.

// Examples:

// Input:  V = 4, edges[][] = [[0, 1], [0, 2], [1, 2], [2, 0], [2, 3]]

// 1
// Cycle: 0 → 2 → 0
// Output:  true
// Explanation:  The diagram clearly shows a cycle 0 → 2 → 0

// Input: V = 4, edges[][] = [[0, 1], [0, 2], [1, 2], [2, 3]

// directed-1
// No Cycle
// Output:  false
// Explanation:  The diagram clearly shows no cycle.

func CycleDetectionInDirectedGraph(g *basic.Graph, node int, visited []bool, stack []bool) bool {
	// Arrive At node
	visited[node] = true
	stack[node] = true

	for _, nbr := range g.AdjList[node] {
		if stack[nbr.To] {
			return true
		} else if !visited[nbr.To] {
			nbrFoundACycle := CycleDetectionInDirectedGraph(g, nbr.To, visited, stack)
			if nbrFoundACycle {
				return true
			}
		}

	}

	stack[node] = false
	return false
}
