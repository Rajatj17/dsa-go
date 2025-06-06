package snake_and_ladder

import (
	basic "DSA/Graph/a_basic"
)

func MinDiceThrows(n int, snakes [][]int, ladders [][]int) {
	board := []int{}

	// Calculate the edge if the user is bitten by snake, if there is snake on 14 & it bites, it makes user fall down to 3,
	// user reaching 14 should point to 3

	for _, snake := range snakes {
		snakeStartPoint := snake[0]
		snakeEndPoint := snake[1]
		board[snakeStartPoint] = snakeEndPoint - snakeStartPoint
	}

	for _, ladders := range ladders {
		ladderStartPoint := ladders[0]
		ladderEndPoint := ladders[1]
		board[ladderStartPoint] = ladderEndPoint - ladderStartPoint
	}

	g := basic.NewGraph(n+1, true)
	for i := 1; i < n; i++ {
		for dice := 0; dice <= 6; dice++ {
			v := i + dice
			v += board[v]
			if v <= n {
				g.AddEdge(i, v, 1)
			}
		}
	}

	// Now do shortest path using dfs or bfs
}
