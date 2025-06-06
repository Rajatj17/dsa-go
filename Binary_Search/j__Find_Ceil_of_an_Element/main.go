package main

// CEILING OF AN ELEMENT IN A SORTED ARRAY:

// Given a sorted array and a value x, the ceiling of x is the smallest element in array greater than or equal to x, and the floor is the greatest element smaller than or equal to x. Assume than the array is sorted in non-decreasing order. Write efficient functions to find floor and ceiling of x.

// Examples :

// For example, let the input array be {1, 2, 8, 10, 10, 12, 19}
// For x = 0:    floor doesn't exist in array,  ceil  = 1
// For x = 1:    floor  = 1,  ceil  = 1
// For x = 5:    floor  = 2,  ceil  = 8
// For x = 20:   floor  = 19,  ceil doesn't exist in array

func FindCeil(arr []int, x int) int {
	start := 0
	end := len(arr) - 1
	res := -1
	for start <= end {
		mid := (start + end) / 2

		if arr[mid] == x {
			return mid
		} else if arr[mid] < x {
			start = mid + 1
		} else {
			res = mid
			end = mid - 1
		}
	}

	return res
}
