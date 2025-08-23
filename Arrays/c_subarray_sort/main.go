package subarraysort

import "math"

/*
Problem: Shortest Unsorted Continuous Subarray

Description:
Given an integer array nums, you need to find one continuous subarray
such that if you only sort this subarray in non-decreasing order,
then the whole array will be sorted in non-decreasing order.

Return the start and end indices of the shortest such subarray.

Example 1:
Input: nums = [2,6,4,8,10,9,15]
Output: (1, 5)
Explanation: You need to sort [6, 4, 8, 10, 9] to make the whole array sorted.

Example 2:
Input: nums = [1,2,3,4]
Output: (-1, 1)
Explanation: Array is already sorted, no subarray needs sorting.

Example 3:
Input: nums = [1]
Output: (-1, 1)
Explanation: Single element array is already sorted.

Constraints:
- 1 <= nums.length <= 10^4
- -10^5 <= nums[i] <= 10^5

Algorithm:
1. Find all elements that are out of order (not in their correct sorted position)
2. Among out-of-order elements, find the minimum and maximum values
3. Find the leftmost position where the minimum out-of-order element should be placed
4. Find the rightmost position where the maximum out-of-order element should be placed
5. Return these positions as the subarray boundaries

Time Complexity: O(n)
Space Complexity: O(1)

Companies: Amazon, Microsoft, Google, Facebook
*/

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

	for i := range n {
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
