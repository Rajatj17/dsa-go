package minarrayjumps

import "math"

func MinArrayJump(arr []int, index int, dp []int) int {
	if index <= 0 {
		return 0
	}

	min_ans := math.MaxInt32
	for jump := range arr[index] {
		min_ans = min(min_ans, 1+MinArrayJump(arr, index+jump, dp))
	}

	dp[index] = min_ans

	return min_ans
}
