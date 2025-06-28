package boxstackingproblem

import "sort"

func BoxStackingProblem(arr [][]int) int {
	n := len(arr)
	if n < 1 {
		return 0
	}

	dp := make([]int, n)

	// Sort the array
	sort.SliceStable(arr, func(i, j int) bool {
		return arr[i][2] < arr[j][2]
	})

	dp[0] = arr[0][2]
	len := 1

	for i := 1; i < n; i++ {
		for j := range i {
			if arr[i][0] > arr[j][0] && // Compare width
				arr[i][1] > arr[j][1] && // Compare breadth
				arr[i][2] > arr[j][2] { // Compare heght
				dp[i] = max(dp[i], arr[i][2]+dp[j])
			}
		}
	}

	return len
}
