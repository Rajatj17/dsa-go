package subarraysort

import "math"

// Given an integer array nums, you need to find one continuous subarray such that if you only sort this subarray in non-decreasing order, then the whole array will be sorted in non-decreasing order.

// Return the shortest such subarray and output its length.

// Example 1:

// Input: nums = [2,6,4,8,10,9,15]
// Output: 5
// Explanation: You need to sort [6, 4, 8, 10, 9] in ascending order to make the whole array sorted in ascending order.
// Example 2:

// Input: nums = [1,2,3,4]
// Output: 0
// Example 3:

// Input: nums = [1]
// Output: 0

func outOfOrder(arr []int, i int) bool {
	x := arr[i]
	if i == 0 {
		return x > arr[1]
	}
	if i == len(arr)-1 {
		return x < arr[i-1]
	}

	return x > arr[i+1] || x < arr[i-1]
}

func SubArraySort(arr []int) (int, int) {
	smallest := math.MaxInt32
	largest := math.MinInt32
	n := len(arr)

	for i := 0; i < n; i++ {
		x := arr[i]
		if outOfOrder(arr, i) {
			smallest = min(smallest, x)
			largest = max(largest, x)
		}
	}

	if smallest == math.MaxInt32 {
		return -1, 1
	}

	left := 0
	for {
		if smallest < arr[left] {
			break
		}

		left++
	}

	right := n - 1
	for {
		if largest < arr[right] {
			break
		}

		right--
	}

	return left, right
}
