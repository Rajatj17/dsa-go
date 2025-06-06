package main

// FIND FLOOR OF AN ELEMENT IN A SORTED ARRAY:

// Given a sorted array and a value x, the floor of x is the largest element in array smaller than or equal to x. Write efficient functions to find floor of x.

// Example:

// Input : arr[] = {1, 2, 8, 10, 10, 12, 19}, x = 5
// Output : 2
// 2 is the largest element in arr[] smaller than 5

func FindFloor(arr []int, x int) int {
	start := 0
	end := len(arr) - 1
	res := -1
	for start <= end {
		mid := (start + end) / 2

		if arr[mid] == x {
			return mid
		} else if arr[mid] < x {
			res = mid
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return res
}