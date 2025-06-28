package largestisland

// Largest Island
// You are given a two dimensional grid, containing only 0 s and 1s. Each 1 represents land, and 0 represents water.
// The adjacent 1s form an island. Each land piece (x,y) is connected to its 4 neighbours (Left, Right, Up and Down).
// Your task is to find the size of largest island - size of island is given by number of 1s in that island.
// Your code should return 0 is no island is present.

// Sample Input

// grid = [
//   [1, 0, 0, 1, 0],
//   [1, 0, 1, 0, 0],
//   [0, 0, 1, 0, 1],
//   [1, 0, 1, 1, 1],
//   [1, 0, 1, 1, 0]
// ]
// Sample Output

// 8

// Explanation

// There are 4 islands (connected components) of sizes 2,1,8,2 out of which 8 is largest.

func dfsHelper(grid [][]int, i int, j int, m int, n int, visited [][]int) int {
	if grid[i][j] == 0 {
		return 0
	}

	count := 1

	dx := []int{0, -1, 0, 1}
	dy := []int{-1, 0, 1, 0}

	for k := 0; k < 4; k++ {
		nx := i + dx[k]
		ny := j + dy[k]

		if nx >= 0 && nx < m &&
			ny >= 0 && ny < n &&
			grid[nx][ny] == 1 &&
			visited[nx][ny] == 0 {

			subcomponent := dfsHelper(grid, nx, ny, m, n, visited)
			count += subcomponent
		}
	}

	return count
}

func LargestIsland(grid [][]int) {

}
