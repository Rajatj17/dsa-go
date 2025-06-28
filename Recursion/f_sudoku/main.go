package sudoku

func isSafe(mat [][]int, i int, j int, number int, n int) bool {
	// Check for row and col
	for k := 0; k < n; k++ {
		if mat[k][j] == number || mat[i][k] == number {
			return false
		}
	}

	// Check for subgrid If (4,1) then it will give (3,0)
	sx := (i / 3) * 3
	sy := (j / 3) * 3

	for x := sx; x < sx+3; x++ {
		for y := sy; y < sy+3; y++ {
			if mat[x][y] == number {
				return false
			}
		}
	}

	return true
}

func SolveSudoku(mat [][]int, i int, j int, n int) bool {
	if i == n {
		// Print the solution later
		return true
	}

	if j == n {
		return SolveSudoku(mat, i+1, 0, n)
	}

	if mat[i][j] != 0 {
		return SolveSudoku(mat, i, j+1, n)
	}

	for no := 1; no <= n; no++ {
		if isSafe(mat, i, j, no, n) {
			mat[i][j] = no
			subProblem := SolveSudoku(mat, i, j, n)
			if subProblem {
				return true
			}
		}
	}

	mat[i][j] = 0

	return false
}
