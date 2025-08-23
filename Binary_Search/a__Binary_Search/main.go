package main

/*
Problem: Binary Search

Description:
Given a sorted array of integers and a target value, return the index of the target in the array.
If the target is not found, return -1.

Algorithm:
Binary search works by repeatedly dividing the search interval in half.
If the target value is less than the middle element, search the left half.
If the target value is greater than the middle element, search the right half.
Continue until the target is found or the interval is empty.

Example 1:
Input: arr = [1, 3, 5, 7, 9, 11], target = 7
Output: 3
Explanation: 7 is found at index 3

Example 2:
Input: arr = [1, 3, 5, 7, 9, 11], target = 4
Output: -1
Explanation: 4 is not in the array

Example 3:
Input: arr = [2, 4, 6, 8, 10], target = 2
Output: 0
Explanation: 2 is found at index 0

Constraints:
- 1 <= arr.length <= 10^4
- -10^4 <= arr[i] <= 10^4
- arr is sorted in ascending order
- All elements in arr are unique
- -10^4 <= target <= 10^4

Time Complexity: O(log n)
Space Complexity: O(1)

Companies: Google, Microsoft, Amazon, Facebook, Apple
*/

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
