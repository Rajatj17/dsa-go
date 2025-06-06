package main

func BinarySearch(arr []int, elemToSearch int) int {
	start := 0
	end := len(arr) - 1

	for start <= end {
		mid := (start + end) / 2

		if arr[mid] == elemToSearch {
			return mid
		} else if elemToSearch < arr[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return -1
}
