package boxstackingproblem

import "sort"

func BoxStackingProblem(arr [][]int) int {
	n := len(arr)
	dp := make([]int, n)

	// Sort the array
	sort.SliceStable(arr, func(i, j int) bool {
		return arr[i][2] < arr[j][2] 
	})

	/
}