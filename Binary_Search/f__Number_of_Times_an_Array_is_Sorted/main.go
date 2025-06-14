// Find the Rotation Count in Rotated Sorted array
// Consider an array of distinct numbers sorted in increasing order. The array has been rotated (clockwise) k number of times. Given such an array, find the value of k.

// Examples:

// Input : arr[] = {15, 18, 2, 3, 6, 12}
// Output: 2
// Explanation : Initial array must be {2, 3,
// 6, 12, 15, 18}. We get the given array after
// rotating the initial array twice.

// Input : arr[] = {7, 9, 11, 12, 5}
// Output: 4

// Input: arr[] = {7, 9, 11, 12, 15};
// Output: 0
package main

func FindPivotElement(arr []int) int {
	start := 0
	end := len(arr) - 1
	res := -1
	for start <= end {
		mid := (start + end) / 2

		if arr[mid] < arr[mid+1] && arr[mid] > arr[mid-1] {
			return mid
		} else if arr[mid] >= arr[start] {
			start = mid + 1
		} else if arr[end] >= arr[mid] {
			end = mid - 1
		}
	}

	return res
}
