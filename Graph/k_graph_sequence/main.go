package graphsequence

// Graph Sequence (Graph + DP)
// Complete the following function which takes in an implicit graph (2D Matrix) and returns the length of longest increasing path sequence inside it.
// A  path can start from any node (x,y), and is increasing if each of its elements is strictly greater than previous element.
// (Assume - 4 way connectivity, each cell is connected with Top, Left, Up & Down)

// Sample Input

//  [
//   [0,  2,  4,  3,  2],
//   [7,  6,  5,  5,  1],
//   [8,  9,  7, 18, 14],
//   [5, 10, 11, 12, 13],
// ]
// Sample output

// 15

// Explanation

// Look at the following path that starts from 1 and ends at 18. It has 15 nodes.

// Expected Time Complexity
// O(MN) where M is the number or rows, N is the number of column

func dfsHelper(grid [][]int, i int, j int, m int, n int, visited [][]int, cache [][]int) {
	dx := []int{0, -1, 0, 1}
	dy := []int{-1, 0, 1, 0}

	count := 0
	for k := 0; k < 4; k++ {
		nx := i + dx[k]
		ny := j + dy[k]

		if nx >= 0 && nx < m &&
			ny >= 0 && ny < n &&
			grid[nx][ny] > grid[i][j] {
			if visited[nx][ny] == 1 {
				count = max(count, 1+cache[nx][ny])
			} else {
				dfsHelper(grid, nx, ny, m, n, visited, cache)
				count = max(count, 1+cache[nx][ny])
			}
		}
	}

	cache[i][j] = count
}
