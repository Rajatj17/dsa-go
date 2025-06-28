package maxsubarraysum

import "math"

// Implement a function that takes an input a vector of integers,
// and prints the maximum subarray sum that can be formed.
// A subarray is defined as consecutive segment of the array.
// If all numbers are negative, then return 0.

// Input
// {-1,2,3,4,-2,6,-8,3}
// Output
// 13

// MaxSubarraySum implements Kadane's algorithm to find maximum subarray sum
// Returns 0 if all numbers are negative
func MaxSubarraySum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	// Check if all numbers are negative
	allNegative := true
	for _, num := range arr {
		if num >= 0 {
			allNegative = false
			break
		}
	}

	if allNegative {
		return 0
	}

	// Kadane's Algorithm
	maxSoFar := math.MinInt32
	maxEndingHere := 0

	for _, num := range arr {
		maxEndingHere = max(num, maxEndingHere+num)
		maxSoFar = max(maxSoFar, maxEndingHere)
	}

	return maxSoFar
}

// Alternative implementation that also tracks the subarray indices
func MaxSubarraySumWithIndices(arr []int) (int, int, int) {
	if len(arr) == 0 {
		return 0, -1, -1
	}

	// Check if all numbers are negative
	allNegative := true
	for _, num := range arr {
		if num >= 0 {
			allNegative = false
			break
		}
	}

	if allNegative {
		return 0, -1, -1
	}

	maxSoFar := math.MinInt32
	maxEndingHere := 0
	start := 0
	end := 0
	tempStart := 0

	for i, num := range arr {
		maxEndingHere += num

		if maxSoFar < maxEndingHere {
			maxSoFar = maxEndingHere
			start = tempStart
			end = i
		}

		if maxEndingHere < 0 {
			maxEndingHere = 0
			tempStart = i + 1
		}
	}

	return maxSoFar, start, end
}

// Helper function for max of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
