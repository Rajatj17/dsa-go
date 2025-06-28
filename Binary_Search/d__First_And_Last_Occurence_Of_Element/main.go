// FIND FIRST AND LAST POSITIONS OF AN ELEMENT IN A SORTED ARRAY:

// Given a sorted array with possibly duplicate elements, the task is to find indexes of first and last occurrences of an element x in the given array.

// Example:

// Input : arr[] = {1, 3, 5, 5, 5, 5 ,67, 123, 125}
//
//	x = 5
//
// Output : First Occurrence = 2
//
//	Last Occurrence = 5
package main

func FirstOccurence(arr []int, elemToSearch int) int {
	start := 0
	end := len(arr) - 1
	res := -1
	for start <= end {
		mid := (start + end) / 2

		if arr[mid] == elemToSearch {
			res = mid
			end = mid - 1
		} else if elemToSearch < arr[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return res
}

func LastOccurence(arr []int, elemToSearch int) int {
	start := 0
	end := len(arr) - 1
	res := -1
	for start <= end {
		mid := (start + end) / 2

		if arr[mid] == elemToSearch {
			res = mid
			start = mid + 1
		} else if elemToSearch < arr[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return res
}
