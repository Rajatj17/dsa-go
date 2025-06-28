package minarrayjumps

import "math"

// Given an array of positive integers, where each element represents
// the max no of steps you can jump forward in the array.
// Find min jumps needed to reach the final index
// Input
// arr = [3, 4, 2, 1, 2, 3, 10, 1, 1, 1, 2, 5]
// Output
// 4

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
