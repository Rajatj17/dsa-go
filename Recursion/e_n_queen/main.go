package nqueen

import "fmt"

func canPlace(board [][]int, i int, j int) bool {
	return false
}

func PrintBoard(n int, board [][]int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(board[i][j])
		}
	}
}

func NQueen(n int, board [][]int, i int) bool {
	if i == n {
		PrintBoard(n, board)
		return true
	}

	for j := range n {
		if canPlace(board, i, j) {
			board[i][j] = 1
			success := NQueen(n, board, i+1)
			if success {
				return true
			}

			board[i][j] = 0
		}
	}

	return false
}
